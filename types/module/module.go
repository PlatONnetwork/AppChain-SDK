package module

import (
	"encoding/json"
	"fmt"
	"sort"

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
)

type Module interface {
	Name() string
}

type InitModule interface {
	Init() error
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
	ExtendData(ctx sdk.Context) []byte
	VerifyExtendData(ctx sdk.Context, data []byte) (common.Hash, error)
	PrepareQC(ctx sdk.Context, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote)
}

type BlockCommiter interface {
	Module
	OnCommit(ctx sdk.Context, block *types.Block) error
}

type ElectionModule interface {
	Module
	NewHeader(ctx sdk.Context, header *types.Header) error
	GetLastNumber(ctx sdk.Context, blockNumber uint64) uint64
	GetValidator(ctx sdk.Context, blockNumber uint64) (*cbfttypes.Validators, error)
	IsCandidateNode(ctx sdk.Context, nodeID enode.IDv0) bool
}

type GenesisModule interface {
	Module
	InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage)
}

type BlockerModule interface {
	Module
	BeginBlock(ctx sdk.Context)
	EndBlock(ctx sdk.Context)
}

type WorkerModule interface {
	Module
	SortTxs(ctx sdk.Context, local map[common.Address]types.Transactions, remote map[common.Address]types.Transactions) (types.Transactions, error)
}

type TransactionModule interface {
	Module
	AddTxs(ctx sdk.Context, local, remote map[common.Address]types.Transactions) (map[common.Address]types.Transactions, map[common.Address]types.Transactions)
}

type Manager struct {
	Modules map[string]interface{}

	ConsensusExtend    string
	Election           string
	Worker             string
	OrderInit          []string
	OrderTxPool        []string
	OrderBlockCommiter []string
	OrderGenesis       []string
	OrderBlocker       []string
	OrderTransaction   []string
}

func NewManager(modules ...Module) *Manager {
	moduleMap := make(map[string]interface{})
	moduleStr := make([]string, 0, len(modules))
	for _, module := range modules {
		moduleMap[module.Name()] = module
		moduleStr = append(moduleStr, module.Name())
	}
	return &Manager{
		Modules:            moduleMap,
		OrderInit:          moduleStr,
		OrderTxPool:        moduleStr,
		OrderBlockCommiter: moduleStr,
		OrderGenesis:       moduleStr,
		OrderBlocker:       moduleStr,
		OrderTransaction:   moduleStr,
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

func (m *Manager) SetOrderBlockCommiter(moduleNames ...string) {
	m.assertNoForgottenModules("SetOrderBlockCommiter", moduleNames, func(moduleName string) bool {
		module := m.Modules[moduleName]
		_, has := module.(BlockCommiter)
		return !has
	})
	m.OrderBlockCommiter = moduleNames
}

func (m *Manager) SetOrderGenesis(moduleNames ...string) {
	m.assertNoForgottenModules("SetOrderGenesis", moduleNames, func(moduleName string) bool {
		module := m.Modules[moduleName]
		_, has := module.(GenesisModule)
		return !has
	})
	m.OrderGenesis = moduleNames
}

func (m *Manager) SetOrderBlocker(moduleNames ...string) {
	m.assertNoForgottenModules("SetOrderBlocker", moduleNames, func(moduleName string) bool {
		module := m.Modules[moduleName]
		_, has := module.(BlockerModule)
		return !has
	})
	m.OrderBlocker = moduleNames
}

func (m *Manager) SetOrderTransaction(moduleNames ...string) {
	m.assertNoForgottenModules("SetOrderTransaction", moduleNames, func(moduleName string) bool {
		module := m.Modules[moduleName]
		_, has := module.(TransactionModule)
		return !has
	})
	m.OrderTransaction = moduleNames
}

func (m *Manager) Init() error {
	log.Info("Init modules for sdk")
	for _, moduleName := range m.OrderInit {
		mod := m.Modules[moduleName]
		if module, ok := mod.(InitModule); ok {
			log.Info("Init for module", "module", moduleName)
			if err := module.Init(); err != nil {
				log.Error("Failed to init module", "module", moduleName, "err", err)
				return err
			}
		}
	}
	return nil
}

func (m *Manager) Contracts() []sdk.SDKContract {
	contracts := make([]sdk.SDKContract, 0)
	for _, mod := range m.Modules {
		if module, ok := mod.(ContractModule); ok {
			contracts = append(contracts, module)
		}
	}
	return contracts
}

func (m *Manager) APIs() []rpc.API {
	apis := make([]rpc.API, 0)
	for _, mod := range m.Modules {
		if module, ok := mod.(RpcModule); ok {
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
	log.Info("Check transaction for tx pool", "hash", tx.Hash())
	for _, moduleName := range m.OrderTxPool {
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
	log.Info("Filter pending transactions for tx pool")
	filterTxs := txs
	for _, moduleName := range m.OrderTxPool {
		module := m.Modules[moduleName]
		if txPoolModule, ok := module.(TxPoolModule); ok {
			log.Debug("Filter pending transaction for module", "module", moduleName)
			filterTxs = txPoolModule.FilterPendingTxs(ctx, filterTxs)
		}
	}
	return filterTxs
}

func (m *Manager) ExtendData(ctx sdk.Context) []byte {
	log.Info("Extend data for consensus engine")
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

func (m *Manager) VerifyExtendData(ctx sdk.Context, data []byte) (common.Hash, error) {
	log.Info("Verify extend data for consensus engine")
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

func (m *Manager) PrepareQC(ctx sdk.Context, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote) {
	log.Info("Notify prepare qc")
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

func (m *Manager) NewHeader(ctx sdk.Context, header *types.Header) error {
	log.Info("New header for election app")
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

func (m *Manager) GetLastNumber(ctx sdk.Context, blockNumber uint64) uint64 {
	log.Info("Get last number for election app")
	if m.Modules[m.Election] == nil {
		return 0
	}

	mod := m.Modules[m.Election]
	if module, ok := mod.(ElectionModule); ok {
		log.Debug("Get last number for module", "module", m.Election)
		return module.GetLastNumber(ctx, blockNumber)
	}
	return 0
}

func (m *Manager) GetValidator(ctx sdk.Context, blockNumber uint64) (*cbfttypes.Validators, error) {
	log.Info("Get validator for election app")
	if m.Modules[m.Election] == nil {
		return nil, nil
	}

	mod := m.Modules[m.Election]
	if module, ok := mod.(ElectionModule); ok {
		log.Debug("Get last number for module", "module", m.Election)
		return module.GetValidator(ctx, blockNumber)
	}
	return nil, nil
}

func (m *Manager) IsCandidateNode(ctx sdk.Context, nodeID enode.IDv0) bool {
	log.Info("Check node if a candidate node for election app")
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

func (m *Manager) OnCommit(ctx sdk.Context, block *types.Block) error {
	log.Info("Notify block commit for election app")
	for _, moduleName := range m.OrderBlockCommiter {
		mod := m.Modules[moduleName]
		if module, ok := mod.(BlockCommiter); ok {
			log.Debug("Notify block commit for module", "module", moduleName)
			if err := module.OnCommit(ctx, block); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Manager) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data map[string]json.RawMessage) error {
	log.Info("Init blockchain state from genesis.json")
	for _, moduleName := range m.OrderGenesis {
		mod := m.Modules[moduleName]
		if module, ok := mod.(GenesisModule); ok {
			log.Info("Running initialization for module ", "module", moduleName)
			module.InitGenesis(ctx, db, chainConfig, data[moduleName])
		}
	}
	return nil
}

func (m *Manager) BeginBlock(ctx sdk.Context) {
	for _, moduleName := range m.OrderBlocker {
		if module, ok := m.Modules[moduleName].(BlockerModule); ok {
			module.BeginBlock(ctx)
		} else {
			continue
		}
	}
}

func (m *Manager) EndBlock(ctx sdk.Context) {
	for _, moduleName := range m.OrderBlocker {
		if module, ok := m.Modules[moduleName].(BlockerModule); ok {
			module.EndBlock(ctx)
		} else {
			continue
		}
	}
}

func (m *Manager) SortTxs(ctx sdk.Context, local, remote map[common.Address]types.Transactions) (types.Transactions, error) {
	for _, moduleName := range m.OrderTransaction {
		if module, ok := m.Modules[moduleName].(TransactionModule); ok {
			local, remote = module.AddTxs(ctx, local, remote)
		} else {
			continue
		}
	}

	if module, ok := m.Modules[m.Worker].(WorkerModule); ok {
		return module.SortTxs(ctx, local, remote)
	}

	allTxs := make(types.Transactions, 0)
	for _, txs := range local {
		allTxs = append(allTxs, txs...)
	}
	for _, txs := range remote {
		allTxs = append(allTxs, txs...)
	}
	return allTxs, nil
}

func (m *Manager) assertNoForgottenModules(setOrderFnName string, moduleNames []string, pass func(moduleName string) bool) {
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
