package baseapp

import (
	"encoding/json"
	"math/big"

	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
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

func (app *BaseApp) Start() error {
	return nil
}

func (app *BaseApp) Stop() error {
	return nil
}

func (app *BaseApp) InitChain(ctx sdk.InitContext) error {
	app.chainID, _ = ctx.Backend().ChainId()
	return app.manager.InitChain(ctx)
}

func (app *BaseApp) Contracts() []sdk.SDKContract {
	return app.manager.Contracts()
}

func (app *BaseApp) CheckTx(ctx sdk.Context, tx *types.Transaction) error {
	return app.manager.CheckTx(ctx, tx)
}

func (app *BaseApp) FilterPendingTxs(ctx sdk.Context, txs map[common.Address]types.Transactions) map[common.Address]types.Transactions {
	return app.manager.FilterPendingTxs(ctx, txs)
}

func (app *BaseApp) APIs() []rpc.API {
	return app.manager.APIs()
}

func (app *BaseApp) Protocols() []basep2p.Protocol {
	return app.manager.Protocols()
}

func (app *BaseApp) ExtendData(ctx sdk.ConsensusContext) []byte {
	return app.manager.ExtendData(ctx)
}

func (app *BaseApp) VerifyExtendData(ctx sdk.ConsensusContext, data []byte) (common.Hash, error) {
	return app.manager.VerifyExtendData(ctx, data)
}

func (app *BaseApp) PrepareQC(ctx sdk.ConsensusContext, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote) {
	app.manager.PrepareQC(ctx, block, votes)
}

func (app *BaseApp) NewHeader(ctx sdk.ConsensusContext, header *types.Header) error {
	return app.manager.NewHeader(ctx, header)
}

func (app *BaseApp) GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {
	return app.manager.GetLastNumber(ctx, blockNumber)
}

func (app *BaseApp) GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	return app.manager.GetValidator(ctx, blockNumber)
}

func (app *BaseApp) IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool {
	return app.manager.IsCandidateNode(ctx, nodeID)
}

func (app *BaseApp) OnCommit(ctx sdk.ConsensusContext, block *types.Block) error {
	return app.manager.OnCommit(ctx, block)
}

func (app *BaseApp) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data map[string]json.RawMessage) error {
	app.chainID = chainConfig.ChainID
	return app.manager.InitGenesis(ctx, db, chainConfig, data)
}

func (app *BaseApp) BeginBlock(ctx sdk.WorkerContext) error {
	return app.manager.BeginBlock(ctx)
}

func (app *BaseApp) EndBlock(ctx sdk.WorkerContext) error {
	return app.manager.EndBlock(ctx)
}

func (app *BaseApp) SortTxs(ctx sdk.WorkerContext, local, remote map[common.Address]types.Transactions) (types.Transactions, error) {
	return app.manager.SortTxs(ctx, local, remote)
}
