package testutil

import (
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/baseapp"
	"github.com/PlatONnetwork/AppChain-SDK/store/storage"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/rpc"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/PlatONnetwork/AppChain-SDK/types/module"
)

type SimApp struct {
	*baseapp.BaseApp
	manager *module.Manager
}

func NewApp(manager *module.Manager) *SimApp {
	app := &SimApp{}

	store, err := storage.NewStorage(filepath.Join(os.TempDir(), fmt.Sprintf("sdk_%d", rand.Int())), 256, 512, "sdk")

	baseApp, err := baseapp.NewBaseApp("simapp", store, manager)
	if err != nil {
		panic(fmt.Sprintf("create base app failed:%s", err.Error()))
	}
	app.BaseApp = baseApp
	app.manager = manager

	app.SetInitChaner(app.InitChain)
	app.SetContractser(app.Contracts)
	app.SetStarter(app.Start)
	app.SetStoper(app.Stop)
	app.SetCheckTxer(app.CheckTx)
	app.SetFilterPendingTxser(app.FilterPendingTxs)
	app.SetAPIser(app.APIs)
	app.SetProtocolser(app.Protocols)
	app.SetExtendDataer(app.ExtendData)
	app.SetVerifyExtendDataer(app.VerifyExtendData)
	app.SetPrepareQCer(app.PrepareQC)
	app.SetNewHeaderer(app.NewHeader)
	app.SetGetLastNumberer(app.GetLastNumber)
	app.SetGetValidatorer(app.GetValidator)
	app.SetIsCandidateNoder(app.IsCandidateNode)
	app.SetOnCommiter(app.OnCommit)
	app.SetInitGenesiser(app.InitGenesis)
	app.SetBeginBlocker(app.BeginBlock)
	app.SetEndBlocker(app.EndBlock)
	app.SetAddTxser(app.AddTxs)
	app.SetSortTxser(app.SortTxs)
	return app
}
func (s *SimApp) Start() error {
	return nil
}

func (s *SimApp) Stop() error {
	return nil
}

func (s *SimApp) InitChain(ctx sdk.InitContext) error {
	return s.manager.InitChain(ctx)
}

func (s *SimApp) Contracts(statedb sdk.StateDBReader, blockNumber uint64) []sdk.SDKContract {
	return s.manager.Contracts(statedb, blockNumber)
}

func (s *SimApp) CheckTx(ctx sdk.Context, tx *types.Transaction) error {
	return s.manager.CheckTx(ctx, tx)
}

func (s *SimApp) FilterPendingTxs(ctx sdk.Context, txs map[common.Address]types.Transactions) map[common.Address]types.Transactions {
	return s.manager.FilterPendingTxs(ctx, txs)
}

func (s *SimApp) APIs() []rpc.API {
	return s.manager.APIs()
}

func (s *SimApp) Protocols() []p2p.Protocol {
	return s.manager.Protocols()
}

func (s *SimApp) ExtendData(ctx sdk.ConsensusContext) []byte {
	return s.manager.ExtendData(ctx)
}

func (s *SimApp) VerifyExtendData(ctx sdk.ConsensusContext, data []byte) (common.Hash, error) {
	return s.manager.VerifyExtendData(ctx, data)
}

func (s *SimApp) PrepareQC(ctx sdk.ConsensusContext, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote) {
	s.manager.PrepareQC(ctx, block, votes)
}

func (s *SimApp) ViewChange(ctx sdk.ConsensusContext, validators []*cbfttypes.ValidateNode) {
	s.manager.ViewChange(ctx, validators)
}

func (s *SimApp) NewHeader(ctx sdk.ConsensusContext, header *types.Header) error {
	return s.manager.NewHeader(ctx, header)
}

func (s *SimApp) GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {
	return s.manager.GetLastNumber(ctx, blockNumber)
}

func (s *SimApp) GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	return s.manager.GetValidator(ctx, blockNumber)
}

func (s *SimApp) IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool {
	return s.manager.IsCandidateNode(ctx, nodeID)
}

func (s *SimApp) OnCommit(ctx sdk.ConsensusContext, block *types.Block) error {
	return s.manager.OnCommit(ctx, block)
}

func (s *SimApp) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data map[string]json.RawMessage) error {
	if err := s.manager.InitGenesis(ctx, db, chainConfig, data); err != nil {
		return err
	}
	return nil
}

func (s *SimApp) BeginBlock(ctx sdk.WorkerContext) error {
	return s.manager.BeginBlock(ctx)
}

func (s *SimApp) EndBlock(ctx sdk.WorkerContext) error {
	return s.manager.EndBlock(ctx)
}

func (s *SimApp) AddTxs(ctx sdk.WorkerContext) (types.Transactions, error) {
	return s.manager.AddTxs(ctx)
}

func (s *SimApp) SortTxs(ctx sdk.WorkerContext, local, remote map[common.Address]types.Transactions) (types.Transactions, error) {
	return s.manager.SortTxs(ctx, local, remote)
}

func (s *SimApp) FillTransactions(ctx sdk.WorkerContext, cb sdk.TxApplyCallbackApp) (types.Transactions, types.Receipts, error) {
	return s.manager.FillTransactions(ctx, cb)
}

func (s *SimApp) ExecuteTxs(ctx sdk.WorkerContext, cApp vm.ContractsApp, txs types.Transactions) (types.Receipts, uint64, error) {
	return s.manager.ExecuteTxs(ctx, cApp, txs)
}
func (s *SimApp) CreateBlockExecutor(ctx sdk.BlockchainContext) (sdk.BlockExecutor, error) {
	return s.manager.CreateBlockExecutor(ctx)
}
func (s *SimApp) CalcBlockDeadline(ctx sdk.ConsensusBlockTimeContext, timePoint time.Time) time.Time {
	return s.manager.CalcBlockDeadline(ctx, timePoint)
}
func (s *SimApp) CalcNextBlockTime(ctx sdk.ConsensusBlockTimeContext, blockTime time.Time) time.Time {
	return s.manager.CalcNextBlockTime(ctx, blockTime)
}
func (s *SimApp) CreateConsensusNetworkEngine(ctx sdk.ConsensusNetworkContext) (sdk.ConsensusNetworkEngine, error) {
	return s.manager.CreateConsensusNetworkEngine(ctx)
}
