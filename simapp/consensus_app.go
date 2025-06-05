package main

import (
	"encoding/json"
	"path/filepath"
	"time"

	"gopkg.in/urfave/cli.v1"

	xconsensus "github.com/PlatONnetwork/AppChain-SDK/x/consensus"
	"github.com/PlatONnetwork/AppChain-SDK/x/l1"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking"
	ctypes "github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"

	"github.com/PlatONnetwork/AppChain-SDK/baseapp"
	"github.com/PlatONnetwork/AppChain-SDK/store/storage"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/node"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/rpc"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type ConsensusApp struct {
	*baseapp.BaseApp

	l1               *l1.L1Module
	stage            *stage.StageModule
	staking          *staking.StakeModule
	consensusNetwork *xconsensus.ConsensusNetworkModule
	manager          *module.Manager
}

func NewConsensusApp(ctx *cli.Context) (*ConsensusApp, error) {
	datadir := node.DefaultDataDir()
	if ctx.GlobalIsSet(utils.DataDirFlag.Name) {
		datadir = ctx.GlobalString(utils.DataDirFlag.Name)
	}
	if datadir != "" {
		absdatadir, err := filepath.Abs(datadir)
		if err != nil {
			return nil, err
		}
		datadir = absdatadir
	}

	dbfile := filepath.Join(datadir, "sdk")
	store, err := storage.NewStorage(dbfile, 256, 512, "sdk")
	if err != nil {
		log.Error("failed to new storage", "err", err)
		return nil, err
	}

	app := &ConsensusApp{}
	app.l1 = l1.NewModule(store)
	app.stage = stage.NewModule(ctx)
	app.staking = staking.NewModule(ctx, app.l1, app.stage)
	app.consensusNetwork = xconsensus.NewModule(ctx)

	manager := module.NewManager(
		app.l1,
		app.stage,
		app.staking,
		app.consensusNetwork,
	)
	manager.SetElection(app.staking.Name())
	manager.SetOrderBeginBlocker(app.stage.Name(), app.staking.Name())
	manager.SetOrderEndBlocker(app.stage.Name(), app.staking.Name())
	manager.SetOrderBlockCommitter(app.staking.Name())
	manager.SetOrderInit(
		app.staking.Name(),
		app.consensusNetwork.Name(),
	)
	manager.SetOrderGenesis(
		app.l1.Name(),
		app.staking.Name(),
		app.stage.Name(),
	)

	baseApp, err := baseapp.NewBaseApp("consensusapp", store, manager)
	if err != nil {
		return nil, err
	}
	app.BaseApp = baseApp
	app.manager = manager

	return app, nil
}

func (s *ConsensusApp) Start() error {
	return nil
}

func (s *ConsensusApp) Stop() error {
	return nil
}

func (s *ConsensusApp) InitChain(ctx sdk.InitContext) error {
	return s.manager.InitChain(ctx)
}

func (s *ConsensusApp) Contracts(statedb sdk.StateDBReader, blockNumber uint64) []sdk.SDKContract {
	return s.manager.Contracts(statedb, blockNumber)
}

func (s *ConsensusApp) CheckTx(ctx sdk.Context, tx *types.Transaction) error {
	return s.manager.CheckTx(ctx, tx)
}

func (s *ConsensusApp) FilterPendingTxs(ctx sdk.Context, txs map[common.Address]types.Transactions) map[common.Address]types.Transactions {
	return s.manager.FilterPendingTxs(ctx, txs)
}

func (s *ConsensusApp) APIs() []rpc.API {
	return s.manager.APIs()
}

func (s *ConsensusApp) Protocols() []p2p.Protocol {
	return s.manager.Protocols()
}

func (s *ConsensusApp) ExtendData(ctx sdk.ConsensusContext) []byte {
	return s.manager.ExtendData(ctx)
}

func (s *ConsensusApp) VerifyExtendData(ctx sdk.ConsensusContext, data []byte) (common.Hash, error) {
	return s.manager.VerifyExtendData(ctx, data)
}

func (s *ConsensusApp) PrepareQC(ctx sdk.ConsensusContext, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote) {
	s.manager.PrepareQC(ctx, block, votes)
}

func (s *ConsensusApp) ViewChange(ctx sdk.ConsensusContext, validators []*cbfttypes.ValidateNode) {
	s.manager.ViewChange(ctx, validators)
}

func (s *ConsensusApp) NewHeader(ctx sdk.ConsensusContext, header *types.Header) error {
	return s.manager.NewHeader(ctx, header)
}

func (s *ConsensusApp) GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {
	return s.manager.GetLastNumber(ctx, blockNumber)
}

func (s *ConsensusApp) GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	return s.manager.GetValidator(ctx, blockNumber)
}

func (s *ConsensusApp) IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool {
	return s.manager.IsCandidateNode(ctx, nodeID)
}

func (s *ConsensusApp) OnCommit(ctx sdk.ConsensusContext, block *types.Block) error {
	return s.manager.OnCommit(ctx, block)
}

func (s *ConsensusApp) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data map[string]json.RawMessage) error {
	if err := s.manager.InitGenesis(ctx, db, chainConfig, data); err != nil {
		return err
	}
	//if err := s.upgrade.SetModuleValidNumberMap(db, s.manager.GetModuleInitValidNumberMap(chainConfig, db)); err != nil {
	//	return err
	//}

	//vm, err := module.GetVersionMapFromGenesis(chainConfig.Modules)
	//if err != nil {
	//	return err
	//}
	//return s.upgrade.SetModuleVersionMap(db, vm)
	return nil
}
func (s *ConsensusApp) BeginBlock(ctx sdk.WorkerContext) error {
	return s.manager.BeginBlock(ctx)
}

func (s *ConsensusApp) EndBlock(ctx sdk.WorkerContext) error {
	return s.manager.EndBlock(ctx)
}

func (s *ConsensusApp) AddTxs(ctx sdk.WorkerContext) (types.Transactions, error) {
	return s.manager.AddTxs(ctx)
}

func (s *ConsensusApp) SortTxs(ctx sdk.WorkerContext, local, remote map[common.Address]types.Transactions) (types.Transactions, error) {
	return s.manager.SortTxs(ctx, local, remote)
}

func (s *ConsensusApp) StartNetworkEngine() {
	s.manager.StartNetworkEngine()
}
func (s *ConsensusApp) Broadcast(msg ctypes.Message) {
	s.manager.Broadcast(msg)
}
func (s *ConsensusApp) PartBroadcast(msg ctypes.Message) {
	s.manager.PartBroadcast(msg)
}
func (s *ConsensusApp) Forwarding(nodeID string, msg ctypes.Message) error {
	return s.manager.Forwarding(nodeID, msg)
}
func (s *ConsensusApp) Send(peerID string, msg ctypes.Message) {
	s.manager.Send(peerID, msg)
}
func (s *ConsensusApp) AvgLatency() time.Duration {
	return s.manager.AvgLatency()
}
func (s *ConsensusApp) PeerSetting(peerID string, bType uint64, blockNumber uint64) error {
	return s.manager.PeerSetting(peerID, bType, blockNumber)
}
func (s *ConsensusApp) RemovePeer(id string) {
	s.RemovePeer(id)
}

func (s *ConsensusApp) FillTransactions(ctx sdk.WorkerContext, cb sdk.TxApplyCallbackApp) (types.Transactions, types.Receipts, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ConsensusApp) ExecuteTxs(ctx sdk.WorkerContext, cApp vm.ContractsApp, txs types.Transactions) (types.Receipts, uint64, error) {
	//TODO implement me
	panic("implement me")
}
