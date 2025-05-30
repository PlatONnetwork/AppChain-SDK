package baseapp

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/PlatONnetwork/AppChain-SDK/store"
	sdktypes "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	ctypes "github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	basep2p "github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/rpc"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

const (
	datadirSDKDatabase = "sdk"
)

type BaseApp struct {
	name    string
	version string
	chainID *big.Int

	store   store.Store
	manager *module.Manager

	initChainer        sdktypes.InitChaner
	contractser        sdktypes.Contractser
	starter            sdktypes.Starter
	stoper             sdktypes.Stoper
	checkTxer          sdktypes.CheckTxer
	filterPendingTxser sdktypes.FilterPendingTxser
	apiser             sdktypes.APIser
	protocolser        sdktypes.Protocolser
	extendDataer       sdktypes.ExtendDataer
	verifyExtendDataer sdktypes.VerifyExtendDataer
	prepareQCer        sdktypes.PrepareQCer
	newHeaderer        sdktypes.NewHeaderer
	getLastNumberer    sdktypes.GetLastNumberer
	getValidatorer     sdktypes.GetValidatorer
	isCandidateNoder   sdktypes.IsCandidateNoder
	onCommiter         sdktypes.OnCommiter
	initGenesiser      sdktypes.InitGenesiser
	beginBlocker       sdktypes.BeginBlocker
	endBlocker         sdktypes.EndBlocker
	addTxser           sdktypes.AddTxser
	sortTxser          sdktypes.SortTxser
}

func NewBaseApp(name string, store store.Store, manager *module.Manager) (*BaseApp, error) {
	return &BaseApp{
		name:    name,
		store:   store,
		manager: manager,
	}, nil
}

func (app *BaseApp) Name() string {
	return app.name
}

func (app *BaseApp) Version() string {
	return app.version
}

func (app *BaseApp) ChainID() *big.Int {
	return app.chainID
}

func (app *BaseApp) SetContractser(contractser sdktypes.Contractser) {
	app.contractser = contractser
}

func (app *BaseApp) SetInitChaner(initChainer sdktypes.InitChaner) {
	app.initChainer = initChainer
}

func (app *BaseApp) SetStarter(starter sdktypes.Starter) {
	app.starter = starter
}

func (app *BaseApp) SetStoper(stoper sdktypes.Stoper) {
	app.stoper = stoper
}

func (app *BaseApp) SetCheckTxer(checkTxer sdktypes.CheckTxer) {
	app.checkTxer = checkTxer
}

func (app *BaseApp) SetFilterPendingTxser(filterPendingTxser sdktypes.FilterPendingTxser) {
	app.filterPendingTxser = filterPendingTxser
}

func (app *BaseApp) SetAPIser(apiser sdktypes.APIser) {
	app.apiser = apiser
}

func (app *BaseApp) SetProtocolser(protocolser sdktypes.Protocolser) {
	app.protocolser = protocolser
}

func (app *BaseApp) SetExtendDataer(extendDataer sdktypes.ExtendDataer) {
	app.extendDataer = extendDataer
}

func (app *BaseApp) SetVerifyExtendDataer(verifyExtendDataer sdktypes.VerifyExtendDataer) {
	app.verifyExtendDataer = verifyExtendDataer
}

func (app *BaseApp) SetPrepareQCer(prepareQCer sdktypes.PrepareQCer) {
	app.prepareQCer = prepareQCer
}

func (app *BaseApp) SetNewHeaderer(newHeaderer sdktypes.NewHeaderer) {
	app.newHeaderer = newHeaderer
}

func (app *BaseApp) SetGetLastNumberer(getLastNumberer sdktypes.GetLastNumberer) {
	app.getLastNumberer = getLastNumberer
}

func (app *BaseApp) SetGetValidatorer(getValidatorer sdktypes.GetValidatorer) {
	app.getValidatorer = getValidatorer
}

func (app *BaseApp) SetIsCandidateNoder(isCandidateNoder sdktypes.IsCandidateNoder) {
	app.isCandidateNoder = isCandidateNoder
}

func (app *BaseApp) SetOnCommiter(onCommiter sdktypes.OnCommiter) {
	app.onCommiter = onCommiter
}

func (app *BaseApp) SetInitGenesiser(initGenesiser sdktypes.InitGenesiser) {
	app.initGenesiser = initGenesiser
}

func (app *BaseApp) SetBeginBlocker(beginBlocker sdktypes.BeginBlocker) {
	app.beginBlocker = beginBlocker
}

func (app *BaseApp) SetEndBlocker(endBlocker sdktypes.EndBlocker) {
	app.endBlocker = endBlocker
}

func (app *BaseApp) SetAddTxser(addTxser sdktypes.AddTxser) {
	app.addTxser = addTxser
}

func (app *BaseApp) SetSortTxser(sortTxser sdktypes.SortTxser) {
	app.sortTxser = sortTxser
}

func (app *BaseApp) Start() error {
	if app.starter != nil {
		app.starter()
	}
	return nil
}

func (app *BaseApp) Stop() error {
	if app.stoper != nil {
		app.stoper()
	}
	return nil
}

func (app *BaseApp) InitChain(ctx sdk.InitContext) error {
	app.chainID, _ = ctx.Backend().ChainId()
	if app.initChainer != nil {
		return app.initChainer(ctx)
	}
	return app.manager.InitChain(ctx)
}

func (app *BaseApp) Contracts(statedb sdk.StateDBReader, blockNumber uint64) []sdk.SDKContract {
	if app.contractser != nil {
		return app.contractser(statedb, blockNumber)
	}
	return app.manager.Contracts(statedb, blockNumber)
}

func (app *BaseApp) CheckTx(ctx sdk.Context, tx *types.Transaction) error {
	if app.checkTxer != nil {
		return app.checkTxer(ctx, tx)
	}
	return app.manager.CheckTx(ctx, tx)
}

func (app *BaseApp) FilterPendingTxs(ctx sdk.Context, txs map[common.Address]types.Transactions) map[common.Address]types.Transactions {
	if app.filterPendingTxser != nil {
		return app.filterPendingTxser(ctx, txs)
	}
	return app.manager.FilterPendingTxs(ctx, txs)
}

func (app *BaseApp) APIs() []rpc.API {
	if app.apiser != nil {
		return app.apiser()
	}
	return app.manager.APIs()
}

func (app *BaseApp) Protocols() []basep2p.Protocol {
	if app.protocolser != nil {
		return app.protocolser()
	}
	return app.manager.Protocols()
}

func (app *BaseApp) ExtendData(ctx sdk.ConsensusContext) []byte {
	if app.extendDataer != nil {
		return app.extendDataer(ctx)
	}
	return app.manager.ExtendData(ctx)
}

func (app *BaseApp) VerifyExtendData(ctx sdk.ConsensusContext, data []byte) (common.Hash, error) {
	if app.verifyExtendDataer != nil {
		return app.verifyExtendDataer(ctx, data)
	}
	return app.manager.VerifyExtendData(ctx, data)
}

func (app *BaseApp) PrepareQC(ctx sdk.ConsensusContext, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote) {
	if app.prepareQCer != nil {
		app.prepareQCer(ctx, block, votes)
	} else {
		app.manager.PrepareQC(ctx, block, votes)
	}
}

func (app *BaseApp) ViewChange(ctx sdk.ConsensusContext, validators []*cbfttypes.ValidateNode) {
	app.manager.ViewChange(ctx, validators)
}

func (app *BaseApp) NewHeader(ctx sdk.ConsensusContext, header *types.Header) error {
	if app.newHeaderer != nil {
		return app.newHeaderer(ctx, header)
	}
	return app.manager.NewHeader(ctx, header)
}

func (app *BaseApp) GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {
	if app.getLastNumberer != nil {
		return app.getLastNumberer(ctx, blockNumber)
	}
	return app.manager.GetLastNumber(ctx, blockNumber)
}

func (app *BaseApp) GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	if app.getValidatorer != nil {
		return app.getValidatorer(ctx, blockNumber)
	}
	return app.manager.GetValidator(ctx, blockNumber)
}

func (app *BaseApp) IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool {
	if app.isCandidateNoder != nil {
		return app.isCandidateNoder(ctx, nodeID)
	}
	return app.manager.IsCandidateNode(ctx, nodeID)
}

func (app *BaseApp) OnCommit(ctx sdk.ConsensusContext, block *types.Block) error {
	if app.onCommiter != nil {
		return app.onCommiter(ctx, block)
	}
	return app.manager.OnCommit(ctx, block)
}

func (app *BaseApp) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data map[string]json.RawMessage) error {
	app.chainID = chainConfig.ChainID
	if app.initGenesiser != nil {
		return app.initGenesiser(ctx, db, chainConfig, data)
	}
	return app.manager.InitGenesis(ctx, db, chainConfig, data)
}

func (app *BaseApp) BeginBlock(ctx sdk.WorkerContext) error {
	if app.beginBlocker != nil {
		return app.beginBlocker(ctx)
	}
	return app.manager.BeginBlock(ctx)
}

func (app *BaseApp) EndBlock(ctx sdk.WorkerContext) error {
	if app.endBlocker != nil {
		return app.endBlocker(ctx)
	}
	return app.manager.EndBlock(ctx)
}

func (app *BaseApp) AddTxs(ctx sdk.WorkerContext) (types.Transactions, error) {
	if app.addTxser != nil {
		return app.addTxser(ctx)
	}
	return app.manager.AddTxs(ctx)
}

func (app *BaseApp) SortTxs(ctx sdk.WorkerContext, local, remote map[common.Address]types.Transactions) (types.Transactions, error) {
	if app.sortTxser != nil {
		return app.sortTxser(ctx, local, remote)
	}
	return app.manager.SortTxs(ctx, local, remote)
}

func (app *BaseApp) StartNetworkEngine() {
	// TODO
}
func (app *BaseApp) Broadcast(msg ctypes.Message) {
	// TODO
}
func (app *BaseApp) PartBroadcast(msg ctypes.Message) {
	// TODO
}
func (app *BaseApp) Forwarding(nodeID string, msg ctypes.Message) error {
	// TODO
	return nil
}
func (app *BaseApp) Send(peerID string, msg ctypes.Message) {
	// TODO
}

func (app *BaseApp) AvgLatency() time.Duration {
	// TODO
	return time.Second
}
func (app *BaseApp) PeerSetting(peerID string, bType uint64, blockNumber uint64) error {
	// TODO
	return nil
}
func (app *BaseApp) RemovePeer(id string) {
	// TODO
}
