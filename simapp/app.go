package main

import (
	"os"
	"path/filepath"

	"github.com/PlatONnetwork/AppChain-SDK/baseapp"
	"github.com/PlatONnetwork/AppChain-SDK/store/storage"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/AppChain-SDK/x/l1"
	stateevent "github.com/PlatONnetwork/AppChain-SDK/x/state_event"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesync"
	"github.com/PlatONnetwork/AppChain-SDK/x/txrelayer"
	"github.com/PlatONnetwork/PlatON-Go/accounts/keystore"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
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

	dbfile := filepath.Join(datadir, "sdk")
	store, err := storage.NewStorage(dbfile, 256, 1024, "sdk")
	if err != nil {
		return nil, err
	}

	stateSync, err := statesync.NewStateSync("", nil, store, nil, extravote.NewExtraVoteDB(store), nil)
	if err != nil {
		return nil, err
	}

	l1Module := l1.NewL1(store)

	stateEvent := stateevent.NewModule(store)

	rootchainRpc := ctx.GlobalString(checkpoint.RootchainNodeRPCFlag.Name)
	rootchainTxRelayer, err := txrelayer.NewModule(rootchainRpc, txrelayer.DefaultReceiptTimeout, txrelayer.DefaultNumRetries)
	if err != nil {
		return nil, err
	}

	ksFile := ctx.GlobalString(checkpoint.KeystoreFlag.Name)
	ksPaswordFile := ctx.GlobalString(checkpoint.PasswordFlag.Name)
	key, err := decryptKey(ksFile, ksPaswordFile)
	if err != nil {
		return nil, err
	}

	checkpoint, err := checkpoint.NewModule(key,
		store,
		nil,
		rootchainTxRelayer,
		extravote.NewExtraVoteDB(store),
		stateEvent,
		l1.NewL1GenesisDB(store))

	extraVote := extravote.NewExtraVote(store, []module.ConsensusExtendModule{stateSync, checkpoint})

	manager := module.NewManager(stateSync, stateEvent, l1Module, extraVote, checkpoint)
	// TODO: 设置staking模块为election回调模块
	//manager.SetElection(staking.Name())
	manager.SetConsensusExtend(extraVote.Name())
	manager.SetWorker(stateSync.Name())
	manager.SetOrderGenesis(l1Module.Name())

	app := &SimApp{}
	baseApp, err := baseapp.NewBaseApp("simapp", store, manager)
	if err != nil {
		return nil, err
	}
	app.BaseApp = baseApp

	return app, nil
}

func decryptKey(ksFile, pwdFile string) (*keystore.Key, error) {
	json, err := os.ReadFile(ksFile)
	if err != nil {
		return nil, err
	}
	passphrase, err := os.ReadFile(pwdFile)
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(json, string(passphrase))
	if err != nil {
		return nil, err
	}
	return key, nil
}
