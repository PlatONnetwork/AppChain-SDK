package main

import (
	"path/filepath"

	"github.com/PlatONnetwork/AppChain-SDK/x/deposit"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesender"
	"github.com/PlatONnetwork/AppChain-SDK/x/vrf"

	"github.com/PlatONnetwork/AppChain-SDK/baseapp"
	"github.com/PlatONnetwork/AppChain-SDK/store/storage"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/AppChain-SDK/x/l1"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking"
	"github.com/PlatONnetwork/AppChain-SDK/x/stateevent"
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

	l1Module := l1.NewModule(store)

	stateEvent := stateevent.NewModule(store)

	stageModule := stage.NewModule(ctx)
	vrfModule := vrf.NewModule(ctx, stageModule)
	stakeModule := staking.NewModule(ctx, l1Module, stageModule)
	stateSync, err := statesync.NewStateSync(ctx, l1Module, stakeModule, store, extravote.NewExtraVoteDB(store))
	if err != nil {
		return nil, err
	}
	rewardModule := reward.NewModule(ctx, stageModule)
	depositModule := deposit.NewModule(ctx, l1Module)
	l2StateSender := statesender.NewModule(ctx)

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
		l1Module)

	extraVote := extravote.NewExtraVote(store, []extravote.ExtraVerifier{stateSync, checkpoint})

	manager := module.NewManager(stateSync, stateEvent, l1Module, extraVote, rootchainTxRelayer, checkpoint, stageModule, vrfModule, stakeModule, rewardModule, depositModule, l2StateSender)
	manager.SetElection(stakeModule.Name())
	manager.SetConsensusExtend(extraVote.Name())
	//manager.SetWorker(stateSync.Name())
	manager.SetOrderTransaction(stateSync.Name(), vrfModule.Name(), stakeModule.Name())
	manager.SetOrderInit(stateSync.Name(), rootchainTxRelayer.Name(), checkpoint.Name(), vrfModule.Name(), stakeModule.Name(), rewardModule.Name())
	manager.SetOrderGenesis(l1Module.Name(), stageModule.Name(), vrfModule.Name(), stakeModule.Name(), rewardModule.Name(), depositModule.Name(), l2StateSender.Name(), stateSync.Name())
	manager.SetOrderBeginBlocker(stageModule.Name(), stakeModule.Name(), rewardModule.Name())
	manager.SetOrderEndBlocker(stageModule.Name(), vrfModule.Name(), stakeModule.Name(), rewardModule.Name())
	manager.SetOrderBlockCommiter(stakeModule.Name(), stateEvent.Name(), checkpoint.Name())

	app := &SimApp{}
	baseApp, err := baseapp.NewBaseApp("simapp", store, manager)
	if err != nil {
		return nil, err
	}
	app.BaseApp = baseApp

	return app, nil
}
