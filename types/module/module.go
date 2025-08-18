package module

import (
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/rpc"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"sort"
	"sync"
	"time"
)

type ModuleValidChecker func(db sdk.StateDBReader, name string, blockNumber uint64) bool
type UpgradeHandler func(ctx sdk.WorkerContext) error

// VersionMap is map of moduleName -> version
type VersionMap map[string]uint64

type ModuleVersion struct {
	Name    string
	Version uint64
}

type contractCache struct {
	blockNumber uint64
	contracts   []sdk.SDKContract
}

type ModuleVersionList []ModuleVersion

func (vm VersionMap) AsSliceSorted() ModuleVersionList {
	keys := make(sort.StringSlice, 0)
	for name, _ := range vm {
		keys = append(keys, name)
	}
	keys.Sort()

	l := make(ModuleVersionList, len(keys))
	for i, key := range keys {
		l[i] = ModuleVersion{
			Name:    key,
			Version: vm[key],
		}
	}
	return l
}

// ValidNumberMap is map of moduleName -> blockNumber
type ValidNumberMap map[string]uint64

type ModuleValidNumber struct {
	Name        string
	ValidNumber uint64
}

func (vn ValidNumberMap) AsSliceSorted() []ModuleValidNumber {
	keys := make(sort.StringSlice, 0)
	for name, _ := range vn {
		keys = append(keys, name)
	}
	keys.Sort()

	l := make([]ModuleValidNumber, len(keys))
	for i, key := range keys {
		l[i] = ModuleValidNumber{
			Name:        key,
			ValidNumber: vn[key],
		}
	}
	return l
}

type Module interface {
	Name() string
	Version() uint64
}

type InitModule interface {
	Module
	Init(ctx sdk.InitContext) error
}

type ContractModule interface {
	Module
	sdk.SDKContract
}

type TxPoolModule interface {
	Module
	CheckTx(ctx sdk.Context, tx *types.Transaction) error
	FilterPendingTxs(ctx sdk.Context, txs map[common.Address]types.Transactions) map[common.Address]types.Transactions
}

type RpcModule interface {
	Module
	APIs() []rpc.API
}

type P2PModule interface {
	Module
	Protocols() []p2p.Protocol
}

type ConsensusExtendModule interface {
	Module
	ExtendData(ctx sdk.ConsensusContext) []byte
	VerifyExtendData(ctx sdk.ConsensusContext, data []byte) (common.Hash, error)
	PrepareQC(ctx sdk.ConsensusContext, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote)
}

type ViewChangeModule interface {
	Module
	ViewChange(ctx sdk.ConsensusContext, validators []*cbfttypes.ValidateNode)
}

type ConsensusNetworkModule interface {
	Module
	CreateConsensusNetworkEngine(ctx sdk.ConsensusNetworkContext) (sdk.ConsensusNetworkEngine, error)
}
type ConsensusBlockTimeModule interface {
	Module
	CalcBlockDeadline(ctx sdk.ConsensusBlockTimeContext, timePoint time.Time) time.Time
	CalcNextBlockTime(ctx sdk.ConsensusBlockTimeContext, blockTime time.Time) time.Time
}

type BlockCommitterModule interface {
	Module
	OnCommit(ctx sdk.ConsensusContext, block *types.Block) error
}

type ElectionModule interface {
	Module
	NewHeader(ctx sdk.ConsensusContext, header *types.Header) error
	GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64
	GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error)
	IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool
}

type GenesisModule interface {
	Module
	InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error
}

type BeginBlockerModule interface {
	Module
	BeginBlock(ctx sdk.WorkerContext) error
}

type EndBlockerModule interface {
	Module
	EndBlock(ctx sdk.WorkerContext) error
}

type WorkerModule interface {
	Module
	SortTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions, remote map[common.Address]types.Transactions) (types.Transactions, error)
}

type TransactionModule interface {
	Module
	AddTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions) (map[common.Address]types.Transactions, error)
}

type FillTransactionsModule interface {
	Module
	FillTransactions(ctx sdk.WorkerContext, cb sdk.TxApplyCallbackApp) (types.Transactions, types.Receipts, error)
}

type TxExecutorModule interface {
	Module
	ExecuteTxs(ctx sdk.WorkerContext, cApp sdk.ContractsApp, txs types.Transactions) (types.Receipts, uint64, error)
}
type BlockExecutorModule interface {
	Module
	CreateBlockExecutor(ctx sdk.BlockchainContext) (sdk.BlockExecutor, error)
}
type UpgradeRegistrar interface {
	RegisterUpgradeHandler(module string, version uint64, handler UpgradeHandler) error
}

type RegistryModule interface {
	Module
	RegistryUpgradeHandler(registrar UpgradeRegistrar) error
}

type Manager struct {
	sync.Mutex
	Modules             map[string]Module
	ConsensusExtend     string
	Election            string
	Worker              string
	BlockExecutor       string
	TxFiller            string
	TxExecutor          string
	ConsensusNetwork    string
	ConsensusBlockTime  string
	OrderInit           []string
	OrderTxPool         []string
	OrderBlockCommitter []string
	OrderGenesis        []string
	OrderBeginBlocker   []string
	OrderEndBlocker     []string
	OrderBlocker        []string
	OrderTransaction    []string

	moduleValidChecker     ModuleValidChecker
	contractValidNumberMap ValidNumberMap
	contractCache          *contractCache
	checkForgotten         bool
}

func NewManager(modules ...Module) *Manager {
	moduleMap := make(map[string]Module)
	moduleStr := make([]string, 0, len(modules))
	for _, module := range modules {
		moduleMap[module.Name()] = module
		moduleStr = append(moduleStr, module.Name())
	}
	return &Manager{
		Modules:                moduleMap,
		OrderInit:              moduleStr,
		OrderTxPool:            moduleStr,
		OrderBlockCommitter:    moduleStr,
		OrderGenesis:           moduleStr,
		OrderBeginBlocker:      moduleStr,
		OrderEndBlocker:        moduleStr,
		OrderBlocker:           moduleStr,
		OrderTransaction:       moduleStr,
		contractValidNumberMap: ValidNumberMap{},
	}
}

func (m *Manager) SetOrderInit(moduleNames ...string) {
	m.assertNoForgottenModules("SetOrderInit", moduleNames, func(moduleName string) bool {
		module := m.Modules[moduleName]
		_, hasInit := module.(InitModule)
		return !hasInit
	})
	m.OrderInit = moduleNames
}

func (m *Manager) SetOrderTxPool(moduleNames ...string) {
	m.assertNoForgottenModules("SetOrderTxPool", moduleNames, func(moduleName string) bool {
		module := m.Modules[moduleName]
		_, hasTxPool := module.(TxPoolModule)
		return !hasTxPool
	})
	m.OrderTxPool = moduleNames
}

func (m *Manager) SetConsensusExtend(moduleName string) {
	mod := m.Modules[moduleName]
	if _, has := mod.(ConsensusExtendModule); !has {
		panic(fmt.Sprintf("ConsensusExtendModule %s missing", moduleName))
	}
	m.ConsensusExtend = moduleName
}

func (m *Manager) SetElection(moduleName string) {
	mod := m.Modules[moduleName]
	if _, has := mod.(ElectionModule); !has {
		panic(fmt.Sprintf("EelectionModule %s missing", moduleName))
	}
	m.Election = moduleName
}

func (m *Manager) SetWorker(moduleName string) {
	mod := m.Modules[moduleName]
	if _, has := mod.(WorkerModule); !has {
		panic(fmt.Sprintf("WorkerModule %s missing", moduleName))
	}
	m.Worker = moduleName
}

func (m *Manager) SetTxFiller(moduleName string) {
	mod := m.Modules[moduleName]
	if _, has := mod.(FillTransactionsModule); !has {
		panic(fmt.Sprintf("FillTransactionModule %s missing", moduleName))
	}
	m.TxFiller = moduleName
}

func (m *Manager) SetTxExecutor(moduleName string) {
	mod := m.Modules[moduleName]
	if _, has := mod.(TxExecutorModule); !has {
		panic(fmt.Sprintf("TxExecutorModule %s missing", moduleName))
	}
	m.TxExecutor = moduleName
}
func (m *Manager) SetBlockExecutor(moduleName string) {
	mod := m.Modules[moduleName]
	if _, has := mod.(BlockExecutorModule); !has {
		panic(fmt.Sprintf("TxExecutorModule %s missing", moduleName))
	}
	m.BlockExecutor = moduleName
}
func (m *Manager) SetConsensusNetwork(moduleName string) {
	mod := m.Modules[moduleName]
	if _, has := mod.(ConsensusNetworkModule); !has {
		panic(fmt.Sprintf("ConsensusNetworkModule %s missing", moduleName))
	}
	m.ConsensusNetwork = moduleName
}
func (m *Manager) SetConsensusBlockTime(moduleName string) {
	mod := m.Modules[moduleName]
	if _, has := mod.(ConsensusBlockTimeModule); !has {
		panic(fmt.Sprintf("ConsensusBlockTimeModule %s missing", moduleName))
	}
	m.ConsensusBlockTime = moduleName
}

func (m *Manager) SetOrderBlockCommitter(moduleNames ...string) {
	m.assertNoForgottenModules("SetOrderBlockCommitter", moduleNames, func(moduleName string) bool {
		module := m.Modules[moduleName]
		_, has := module.(BlockCommitterModule)
		return !has
	})
	m.OrderBlockCommitter = moduleNames
}

func (m *Manager) SetOrderGenesis(moduleNames ...string) {
	m.assertNoForgottenModules("SetOrderGenesis", moduleNames, func(moduleName string) bool {
		module := m.Modules[moduleName]
		_, has := module.(GenesisModule)
		return !has
	})
	m.OrderGenesis = moduleNames
}

func (m *Manager) SetOrderBeginBlocker(moduleNames ...string) {
	m.assertNoForgottenModules("SetOrderBeginBlocker", moduleNames, func(moduleName string) bool {
		module := m.Modules[moduleName]
		_, has := module.(BeginBlockerModule)
		return !has
	})
	m.OrderBeginBlocker = moduleNames
}

func (m *Manager) SetOrderEndBlocker(moduleNames ...string) {
	m.assertNoForgottenModules("SetOrderEndBlocker", moduleNames, func(moduleName string) bool {
		module := m.Modules[moduleName]
		_, has := module.(EndBlockerModule)
		return !has
	})
	m.OrderEndBlocker = moduleNames
}

func (m *Manager) SetOrderTransaction(moduleNames ...string) {
	m.assertNoForgottenModules("SetOrderTransaction", moduleNames, func(moduleName string) bool {
		module := m.Modules[moduleName]
		_, has := module.(TransactionModule)
		return !has
	})
	m.OrderTransaction = moduleNames
}

func (m *Manager) SetModuleValidChecker(moduleValidChecker ModuleValidChecker) {
	m.moduleValidChecker = moduleValidChecker
}

func (m *Manager) InitChain(ctx sdk.InitContext) error {
	log.Debug("Init modules for sdk")
	for _, moduleName := range m.OrderInit {

		mod := m.Modules[moduleName]
		if module, ok := mod.(InitModule); ok {
			log.Debug("Init for module", "module", moduleName)
			if err := module.Init(ctx); err != nil {
				log.Error("Failed to init module", "module", moduleName, "err", err)
				return err
			}
		}
	}
	return nil
}

func (m *Manager) Contracts(statedb sdk.StateDBReader, blockNumber uint64) []sdk.SDKContract {
	m.Lock()
	defer m.Unlock()
	if m.contractCache != nil && m.contractCache.blockNumber == blockNumber {
		return m.contractCache.contracts
	}
	contracts := make([]sdk.SDKContract, 0, len(m.contractValidNumberMap))
	for name, mod := range m.Modules {
		if module, ok := mod.(ContractModule); ok {
			var number = uint64(0)
			var ok bool
			if number, ok = m.contractValidNumberMap[name]; !ok {
				number = module.ContractCreateBlockNumber(statedb)
				m.contractValidNumberMap[name] = number
			}
			if number <= blockNumber {
				contracts = append(contracts, module)
			}
		}
	}
	m.contractCache = &contractCache{
		blockNumber: blockNumber,
		contracts:   contracts,
	}
	return contracts
}

func (m *Manager) APIs() []rpc.API {
	apis := make([]rpc.API, 0)
	for _, mod := range m.Modules {
		if module, ok := mod.(RpcModule); ok {
			log.Debug("Load rpc module", "module", module.Name())
			apis = append(apis, module.APIs()...)
		}
	}
	return apis
}

func (m *Manager) Protocols() []p2p.Protocol {
	protocols := make([]p2p.Protocol, 0)
	for _, mod := range m.Modules {
		if module, ok := mod.(P2PModule); ok {
			protocols = append(protocols, module.Protocols()...)
		}
	}
	return protocols
}

func (m *Manager) CheckTx(ctx sdk.Context, tx *types.Transaction) error {
	log.Trace("Check transaction for tx pool", "hash", tx.Hash())
	for _, moduleName := range m.OrderTxPool {
		statedb, _ := ctx.Backend().State()
		if !m.isModuleValid(statedb, moduleName, ctx.Backend().CurrentHeader().Number.Uint64()) {
			continue
		}

		module := m.Modules[moduleName]
		if txPoolModule, ok := module.(TxPoolModule); ok {
			log.Debug("Check transaction for module", "module", moduleName, "hash", tx.Hash())
			if err := txPoolModule.CheckTx(ctx, tx); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Manager) FilterPendingTxs(ctx sdk.Context, txs map[common.Address]types.Transactions) map[common.Address]types.Transactions {
	log.Debug("Filter pending transactions for tx pool")
	filterTxs := txs
	for _, moduleName := range m.OrderTxPool {
		statedb, _ := ctx.Backend().State()
		if !m.isModuleValid(statedb, moduleName, ctx.Backend().CurrentHeader().Number.Uint64()) {
			continue
		}

		module := m.Modules[moduleName]
		if txPoolModule, ok := module.(TxPoolModule); ok {
			log.Debug("Filter pending transaction for module", "module", moduleName)
			filterTxs = txPoolModule.FilterPendingTxs(ctx, filterTxs)
		}
	}
	return filterTxs
}

func (m *Manager) ExtendData(ctx sdk.ConsensusContext) []byte {
	log.Debug("Extend data for consensus engine")
	if m.Modules[m.ConsensusExtend] == nil {
		return []byte{}
	}

	mod := m.Modules[m.ConsensusExtend]
	if module, ok := mod.(ConsensusExtendModule); ok {
		log.Debug("Extend data for module", "module", m.ConsensusExtend)
		return module.ExtendData(ctx)
	}
	return []byte{}
}

func (m *Manager) VerifyExtendData(ctx sdk.ConsensusContext, data []byte) (common.Hash, error) {
	log.Debug("Verify extend data for consensus engine")
	if m.Modules[m.ConsensusExtend] == nil {
		return common.ZeroHash, nil
	}

	mod := m.Modules[m.ConsensusExtend]
	if module, ok := mod.(ConsensusExtendModule); ok {
		log.Debug("Verify extend data for module", "module", m.ConsensusExtend)
		return module.VerifyExtendData(ctx, data)
	}
	return common.ZeroHash, nil
}

func (m *Manager) PrepareQC(ctx sdk.ConsensusContext, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote) {
	log.Debug("Notify prepare qc")
	if m.Modules[m.ConsensusExtend] == nil {
		return
	}

	mod := m.Modules[m.ConsensusExtend]
	if module, ok := mod.(ConsensusExtendModule); ok {
		log.Debug("Notify prepare qc for module", "module", m.ConsensusExtend)
		module.PrepareQC(ctx, block, votes)
	}
	return
}

func (m *Manager) ViewChange(ctx sdk.ConsensusContext, validators []*cbfttypes.ValidateNode) {
	log.Debug("Notify viewchange")

	for name, mod := range m.Modules {
		if module, ok := mod.(ViewChangeModule); ok {
			log.Debug("Notify viewchange for module", "module", name)
			module.ViewChange(ctx, validators)
		}
	}
}

func (m *Manager) NewHeader(ctx sdk.ConsensusContext, header *types.Header) error {
	log.Debug("New header for election app", "blockNumber", header.Number.Uint64())
	if m.Modules[m.Election] == nil {
		return nil
	}

	mod := m.Modules[m.Election]
	if module, ok := mod.(ElectionModule); ok {
		log.Debug("New header for election module", "module", m.Election)
		return module.NewHeader(ctx, header)
	}
	return nil
}

func (m *Manager) GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {
	log.Debug("Get last number for election app", "blockNumber", blockNumber)
	if m.Modules[m.Election] == nil {
		return 0
	}

	mod := m.Modules[m.Election]
	if module, ok := mod.(ElectionModule); ok {
		lastNumber := module.GetLastNumber(ctx, blockNumber)
		log.Debug("Get last number for module", "module", m.Election, "blockNumber", blockNumber, "lastNumber", lastNumber)
		return lastNumber
	}
	return 0
}

func (m *Manager) GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	log.Debug("Get validator for election app", "blockNumber", blockNumber)
	if m.Modules[m.Election] == nil {
		return nil, nil
	}

	mod := m.Modules[m.Election]
	if module, ok := mod.(ElectionModule); ok {
		vals, err := module.GetValidator(ctx, blockNumber)
		if err == nil {
			log.Debug("Get last number for module", "module", m.Election, "blockNumber", blockNumber, "validators", vals.String())
		}
		return vals, err
	}
	return nil, nil
}

func (m *Manager) IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool {
	log.Debug("Check node if a candidate node for election app")
	if m.Modules[m.Election] == nil {
		return false
	}

	mod := m.Modules[m.Election]
	if module, ok := mod.(ElectionModule); ok {
		log.Debug("Check node if a candidate node for module", "module", m.Election)
		return module.IsCandidateNode(ctx, nodeID)
	}
	return false
}

func (m *Manager) OnCommit(ctx sdk.ConsensusContext, block *types.Block) error {
	log.Debug("Notify block commit for election app")
	for _, moduleName := range m.OrderBlockCommitter {
		if !m.isModuleValid(ctx.StateDB(), moduleName, ctx.Header().Number.Uint64()) {
			continue
		}

		mod := m.Modules[moduleName]
		if module, ok := mod.(BlockCommitterModule); ok {
			log.Debug("Notify block commit for module", "module", moduleName)
			if err := module.OnCommit(ctx, block); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Manager) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data map[string]json.RawMessage) error {
	log.Debug("Init blockchain state from genesis.json")
	for _, moduleName := range m.OrderGenesis {

		mod := m.Modules[moduleName]
		if module, ok := mod.(GenesisModule); ok {
			log.Debug("Running initialization for module ", "module", moduleName)
			conf := data[moduleName]
			if conf == nil {
				conf = json.RawMessage("{}")
			}
			if err := module.InitGenesis(ctx, db, chainConfig, data[moduleName]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Manager) BeginBlock(ctx sdk.WorkerContext) error {
	for _, moduleName := range m.OrderBeginBlocker {
		if !m.isModuleValid(ctx.StateDB(), moduleName, ctx.Header().Number.Uint64()) {
			continue
		}
		if module, ok := m.Modules[moduleName].(BeginBlockerModule); ok {
			if err := module.BeginBlock(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Manager) EndBlock(ctx sdk.WorkerContext) error {
	for _, moduleName := range m.OrderEndBlocker {
		if !m.isModuleValid(ctx.StateDB(), moduleName, ctx.Header().Number.Uint64()) {
			continue
		}
		if module, ok := m.Modules[moduleName].(EndBlockerModule); ok {
			if err := module.EndBlock(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Manager) AddTxs(ctx sdk.WorkerContext) (types.Transactions, error) {
	local := make(map[common.Address]types.Transactions, 0)
	var err error
	for _, module := range m.Modules {
		if !m.isModuleValid(ctx.StateDB(), module.Name(), ctx.Header().Number.Uint64()) {
			continue
		}
		if txModule, ok := module.(TransactionModule); ok {
			local, err = txModule.AddTxs(ctx, local)
			if err != nil {
				return nil, err
			}
		}
	}

	allTxs := make(types.Transactions, 0)
	for _, txs := range local {
		allTxs = append(allTxs, txs...)
	}
	return allTxs, nil
}

func (m *Manager) SortTxs(ctx sdk.WorkerContext, local, remote map[common.Address]types.Transactions) (types.Transactions, error) {
	isValid := m.isModuleValid(ctx.StateDB(), m.Worker, ctx.Header().Number.Uint64())
	if module, ok := m.Modules[m.Worker].(WorkerModule); ok && isValid {
		return module.SortTxs(ctx, local, remote)
	}

	signer := types.MakeSigner(ctx.ChainConfig(), ctx.Header().Number)
	allTxs := make(types.Transactions, 0)
	if len(local) > 0 {
		peeker := types.NewTransactionsByPriceAndNonce(signer, local, ctx.Header().BaseFee)
		allTxs = append(allTxs, peeker.PeekAll()...)
	}
	if len(remote) > 0 {
		peeker := types.NewTransactionsByPriceAndNonce(signer, remote, ctx.Header().BaseFee)
		allTxs = append(allTxs, peeker.PeekAll()...)
	}
	return allTxs, nil
}

func (m *Manager) FillTransactions(ctx sdk.WorkerContext, cb sdk.TxApplyCallbackApp) (types.Transactions, types.Receipts, error) {
	isValid := m.isModuleValid(ctx.StateDB(), m.TxFiller, ctx.Header().Number.Uint64())
	if module, ok := m.Modules[m.TxFiller].(FillTransactionsModule); ok && isValid {
		return module.FillTransactions(ctx, cb)
	}
	return nil, nil, nil
}

func (m *Manager) ExecuteTxs(ctx sdk.WorkerContext, cApp sdk.ContractsApp, txs types.Transactions) (types.Receipts, uint64, error) {
	isValid := m.isModuleValid(ctx.StateDB(), m.TxExecutor, ctx.Header().Number.Uint64())
	if module, ok := m.Modules[m.TxExecutor].(TxExecutorModule); ok && isValid {
		return module.ExecuteTxs(ctx, cApp, txs)
	}
	return nil, 0, nil
}

func (m *Manager) CreateBlockExecutor(ctx sdk.BlockchainContext) (sdk.BlockExecutor, error) {
	if module, ok := m.Modules[m.BlockExecutor].(BlockExecutorModule); ok {
		return module.CreateBlockExecutor(ctx)
	}
	return nil, nil
}
func (m *Manager) RegisterUpgradeHandler(registrar UpgradeRegistrar) error {
	for _, module := range m.Modules {
		if mod, ok := module.(RegistryModule); ok {
			if err := mod.RegistryUpgradeHandler(registrar); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Manager) GetModuleInitValidNumberMap(chainConfig *params.ChainConfig, db sdk.StateDBReader) ValidNumberMap {
	vn := ValidNumberMap{}
	for name, _ := range chainConfig.Modules {
		if m.Modules[name] == nil {
			continue
		}

		vn[name] = 0
		mod := m.Modules[name]
		if module, ok := mod.(ContractModule); ok {
			vn[name] = module.ContractCreateBlockNumber(db)
		}
	}
	return vn
}

func (m *Manager) IsContractModule(moduleName string) bool {
	yesOrNo := false
	if module, ok := m.Modules[moduleName]; ok {
		_, yesOrNo = module.(ContractModule)
	}
	return yesOrNo
}

func (m *Manager) SetCheckForgotten(check bool) {
	m.checkForgotten = check
}

func (m *Manager) assertNoForgottenModules(setOrderFnName string, moduleNames []string, pass func(moduleName string) bool) {
	if !m.checkForgotten {
		return
	}
	ms := make(map[string]bool)
	for _, m := range moduleNames {
		ms[m] = true
	}
	var missing []string
	for m := range m.Modules {
		m := m
		if pass != nil && pass(m) {
			continue
		}

		if !ms[m] {
			missing = append(missing, m)
		}
	}
	if len(missing) != 0 {
		sort.Strings(missing)
		panic(fmt.Sprintf(
			"all modules must be defined when setting %s, missing: %v", setOrderFnName, missing))
	}
}

func (m *Manager) isModuleValid(db sdk.StateDBReader, name string, blockNumber uint64) bool {
	if m.moduleValidChecker != nil {
		return m.moduleValidChecker(db, name, blockNumber)
	}
	return true
}
func (m *Manager) CreateConsensusNetworkEngine(ctx sdk.ConsensusNetworkContext) (sdk.ConsensusNetworkEngine, error) {
	if module, ok := m.Modules[m.ConsensusNetwork].(ConsensusNetworkModule); ok {
		return module.CreateConsensusNetworkEngine(ctx)
	}
	return nil, nil
}

func (m *Manager) CalcBlockDeadline(ctx sdk.ConsensusBlockTimeContext, timePoint time.Time) time.Time {
	if module, ok := m.Modules[m.ConsensusBlockTime].(ConsensusBlockTimeModule); ok {
		return module.CalcBlockDeadline(ctx, timePoint)
	}
	if ctx.Deadline().Sub(timePoint) > ctx.ProduceInterval() {
		return timePoint.Add(ctx.ProduceInterval())
	}
	return ctx.Deadline()
}
func (m *Manager) CalcNextBlockTime(ctx sdk.ConsensusBlockTimeContext, blockTime time.Time) time.Time {
	if module, ok := m.Modules[m.ConsensusBlockTime].(ConsensusBlockTimeModule); ok {
		return module.CalcNextBlockTime(ctx, blockTime)
	}
	return blockTime.Add(ctx.ProduceInterval())
}
