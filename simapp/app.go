package main

import (
	"encoding/json"
	"path/filepath"
	"time"

	"gopkg.in/urfave/cli.v1"

	xconsensus "github.com/PlatONnetwork/AppChain-SDK/x/consensus"
	"github.com/PlatONnetwork/AppChain-SDK/x/gov"
	"github.com/PlatONnetwork/AppChain-SDK/x/votetoken"
	ctypes "github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"

	"github.com/PlatONnetwork/AppChain-SDK/baseapp"
	"github.com/PlatONnetwork/AppChain-SDK/store/storage"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint"
	"github.com/PlatONnetwork/AppChain-SDK/x/deposit"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/AppChain-SDK/x/l1"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking"
	"github.com/PlatONnetwork/AppChain-SDK/x/stateevent"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesender"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesync"
	"github.com/PlatONnetwork/AppChain-SDK/x/txrelayer"
	"github.com/PlatONnetwork/AppChain-SDK/x/upgrade"
	"github.com/PlatONnetwork/AppChain-SDK/x/upgrade/testcontract"
	"github.com/PlatONnetwork/AppChain-SDK/x/upgrade/testmod"
	"github.com/PlatONnetwork/AppChain-SDK/x/vrf"
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

type SimApp struct {
	*baseapp.BaseApp

	l1                 *l1.L1Module
	stateEvent         *stateevent.Module
	stage              *stage.StageModule
	vrf                *vrf.VRFModule
	staking            *staking.StakeModule
	reward             *reward.RewardModule
	deposit            *deposit.DepositModule
	l2StateSender      *statesender.StateSenderModule
	stateSync          *statesync.StateSync
	rootchainTxRelayer *txrelayer.Module
	checkpoint         *checkpoint.Module
	extraVote          *extravote.ExtraVote
	upgrade            *upgrade.Module
	voteToken          *votetoken.Module
	gov                *gov.Module
	consensusNetwork   *xconsensus.ConsensusNetworkModule
	manager            *module.Manager
}

func NewSimApp(ctx *cli.Context) (*SimApp, error) {
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

	app := &SimApp{}

	app.l1 = l1.NewModule(store)

	app.stateEvent = stateevent.NewModule(store)

	app.stage = stage.NewModule(ctx)
	app.vrf = vrf.NewModule(ctx, app.stage)
	app.staking = staking.NewModule(ctx, app.l1, app.stage)
	app.stateSync, err = statesync.NewModule(ctx, app.l1, app.staking, store, extravote.NewExtraVoteDB(store))
	if err != nil {
		return nil, err
	}
	app.reward = reward.NewModule(ctx, app.stage)
	app.deposit = deposit.NewModule(ctx, app.l1)
	app.l2StateSender = statesender.NewModule(ctx)

	app.vrf.SetStakeModule(app.staking)
	app.staking.SetRewardModule(app.reward)
	app.staking.SetVRFModule(app.vrf)
	app.reward.SetStakeModule(app.staking)

	rootchainRpc := ctx.GlobalString(x.RootchainNodeRPCFlag.Name)
	app.rootchainTxRelayer = txrelayer.NewModule(rootchainRpc, txrelayer.DefaultReceiptTimeout, txrelayer.DefaultNumRetries)

	app.checkpoint, err = checkpoint.NewModule(
		ctx,
		store,
		app.staking,
		app.rootchainTxRelayer,
		extravote.NewExtraVoteDB(store),
		app.stateEvent,
		app.l1)
	if err != nil {
		return nil, err
	}

	app.extraVote = extravote.NewExtraVote(store, []extravote.ExtraVerifier{app.stateSync, app.checkpoint})

	app.upgrade = upgrade.NewModule(store)

	app.voteToken, _ = votetoken.NewModule()
	app.gov, _ = gov.NewModule(ctx)
	tm := testmod.NewModule()
	tc := testcontract.NewModule()
	app.consensusNetwork = xconsensus.NewModule(ctx)

	manager := module.NewManager(
		app.stateSync,
		app.stateEvent,
		app.l1,
		app.extraVote,
		app.rootchainTxRelayer,
		app.checkpoint,
		app.stage,
		app.vrf,
		app.staking,
		app.reward,
		app.deposit,
		app.l2StateSender,
		app.upgrade,
		app.gov,
		app.voteToken,
		app.consensusNetwork,
		tm, tc)
	manager.SetElection(app.staking.Name())
	manager.SetConsensusExtend(app.extraVote.Name())
	manager.SetOrderTransaction(app.stateSync.Name(), app.vrf.Name(), app.staking.Name(), app.gov.Name())
	manager.SetOrderBeginBlocker(app.upgrade.Name(), app.stage.Name(), app.staking.Name(), app.reward.Name(), tm.Name())
	manager.SetOrderEndBlocker(app.upgrade.Name(), app.stage.Name(), app.vrf.Name(), app.staking.Name(), app.reward.Name())
	manager.SetOrderBlockCommitter(app.staking.Name(), app.stateEvent.Name(), app.checkpoint.Name())

	manager.SetOrderInit(
		app.stateSync.Name(),
		app.rootchainTxRelayer.Name(),
		app.checkpoint.Name(),
		app.vrf.Name(),
		app.staking.Name(),
		app.reward.Name(),
		app.upgrade.Name(),
		app.gov.Name(),
		app.consensusNetwork.Name(),
	)

	manager.SetOrderGenesis(
		app.l1.Name(),
		app.stage.Name(),
		app.vrf.Name(),
		app.staking.Name(),
		app.reward.Name(),
		app.deposit.Name(),
		app.l2StateSender.Name(),
		app.stateSync.Name(),
		app.upgrade.Name(),
		app.voteToken.Name(),
		app.gov.Name(),
		tm.Name(),
		tc.Name(),
	)

	manager.SetModuleValidChecker(app.upgrade.IsModuleValid)

	manager.RegisterUpgradeHandler(app.upgrade)
	app.upgrade.SetIsContractModule(manager.IsContractModule)

	baseApp, err := baseapp.NewBaseApp("simapp", store, manager)
	if err != nil {
		return nil, err
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

	return app, nil
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
	if err := s.upgrade.SetModuleValidNumberMap(db, s.manager.GetModuleInitValidNumberMap(chainConfig, db)); err != nil {
		return err
	}

	vm, err := module.GetVersionMapFromGenesis(chainConfig.Modules)
	if err != nil {
		return err
	}
	return s.upgrade.SetModuleVersionMap(db, vm)
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

func (s *SimApp) StartNetworkEngine() {
	// TODO
}
func (s *SimApp) RemoveMessageHash(id string, msgHash common.Hash) {
	// TODO
}
func (s *SimApp) Broadcast(msg ctypes.Message) {
	// TODO
}
func (s *SimApp) PartBroadcast(msg ctypes.Message) {
	// TODO
}
func (s *SimApp) Forwarding(nodeID string, msg ctypes.Message) error {
	// TODO
	return nil
}
func (s *SimApp) Send(peerID string, msg ctypes.Message) {
	// TODO
}
func (s *SimApp) AvgLatency() time.Duration {
	// TODO
	return time.Second
}
func (s *SimApp) PeerSetting(peerID string, bType uint64, blockNumber uint64) error {
	// TODO
	return nil
}
func (s *SimApp) RemovePeer(id string) {
	// TODO
}

func (s *SimApp) FillTransactions(ctx sdk.WorkerContext, cb sdk.TxApplyCallbackApp) (types.Transactions, types.Receipts, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SimApp) ExecuteTxs(ctx sdk.WorkerContext, cApp vm.ContractsApp, txs types.Transactions) (types.Receipts, uint64, error) {
	//TODO implement me
	panic("implement me")
}
