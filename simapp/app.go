package main

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/deposit"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage"
	"github.com/PlatONnetwork/AppChain-SDK/x/vrf"
	"path/filepath"

	"github.com/PlatONnetwork/AppChain-SDK/baseapp"
	"github.com/PlatONnetwork/AppChain-SDK/store/storage"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/AppChain-SDK/x/l1"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking"
	stateevent "github.com/PlatONnetwork/AppChain-SDK/x/state_event"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesync"
	"github.com/PlatONnetwork/AppChain-SDK/x/txrelayer"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/node"
	"gopkg.in/urfave/cli.v1"
)

type SimApp struct {
	*baseapp.BaseApp
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

	stateSync, err := statesync.NewStateSync(ctx, store, extravote.NewExtraVoteDB(store))
	if err != nil {
		return nil, err
	}

	l1Module := l1.NewL1(store)
	stateEvent := stateevent.NewModule(store)

	stageModule := stage.NewStageModule(ctx)
	vrfModule := vrf.NewVRFModule(ctx, stageModule)
	stakeModule := staking.NewStakeModule(ctx, stageModule)
	rewardModule := reward.NewRewardModule(ctx, stageModule)
	depositModule := deposit.NewDepositModule(ctx)

	vrfModule.SetStakeModule(stakeModule)
	stakeModule.SetRewardModule(rewardModule)
	stakeModule.SetVRFModule(vrfModule)
	rewardModule.SetStakeModule(stakeModule)

	rootchainRpc := ctx.GlobalString(x.RootchainNodeRPCFlag.Name)
	rootchainTxRelayer := txrelayer.NewModule(rootchainRpc, txrelayer.DefaultReceiptTimeout, txrelayer.DefaultNumRetries)

	checkpoint, err := checkpoint.NewModule(
		ctx,
		store,
		stakeModule,
		rootchainTxRelayer,
		extravote.NewExtraVoteDB(store),
		stateEvent,
		l1.NewL1GenesisDB(store))

	extraVote := extravote.NewExtraVote(store, []extravote.ExtraVerifier{stateSync, checkpoint})

	manager := module.NewManager(stateSync, stateEvent, l1Module, extraVote, checkpoint, stageModule, vrfModule, stakeModule, rewardModule, depositModule)
	manager.SetElection(stakeModule.Name())
	manager.SetConsensusExtend(extraVote.Name())
	//manager.SetWorker(stateSync.Name())
	manager.SetOrderInit(stateSync.Name(), checkpoint.Name())
	manager.SetOrderGenesis(l1Module.Name(), stageModule.Name(), vrfModule.Name(), stakeModule.Name(), rewardModule.Name())
	manager.SetOrderBeginBlocker(stageModule.Name(), stakeModule.Name(), rewardModule.Name())
	manager.SetOrderEndBlocker(stageModule.Name(), vrfModule.Name(), stakeModule.Name(), rewardModule.Name())
	manager.SetOrderBlockCommiter(stakeModule.Name())

	app := &SimApp{}
	baseApp, err := baseapp.NewBaseApp("simapp", store, manager)
	if err != nil {
		return nil, err
	}
	app.BaseApp = baseApp

	return app, nil
}
