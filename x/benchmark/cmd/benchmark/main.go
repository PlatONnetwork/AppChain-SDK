package main

import (
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/store/storage"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/benchmark"
	xconsensus "github.com/PlatONnetwork/AppChain-SDK/x/consensus"
	"github.com/PlatONnetwork/AppChain-SDK/x/nontxpool"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/node"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/PlatONnetwork/PlatON-Go/sdk/app"
	"gopkg.in/urfave/cli.v1"
	"os"
	"path/filepath"

	cmdutils "github.com/PlatONnetwork/PlatON-Go/cmd/utils"
)

func main() {
	cliApp := cli.NewApp()
	benchmark.AddBenchmarkFlags(cliApp)
	nontxpool.AddNonTxPoolFlags(cliApp)
	app.InitApp(cliApp, func(ctx *cli.Context) sdk.App {
		datadir := node.DefaultDataDir()
		if ctx.GlobalIsSet(cmdutils.DataDirFlag.Name) {
			datadir = ctx.GlobalString(cmdutils.DataDirFlag.Name)
		}
		if datadir != "" {
			absdatadir, err := filepath.Abs(datadir)
			if err != nil {
				log.Crit("get data dir failed", "err", err)
			}
			datadir = absdatadir
		}

		dbfile := filepath.Join(datadir, "sdk")
		store, err := storage.NewStorage(dbfile, 256, 512, "sdk")
		if err != nil {
			log.Crit("failed to new storage", "err", err)
		}

		election := NewElection()
		benchmarkModule := benchmark.NewModule(ctx, store)
		nonTxPoolModule := nontxpool.NewModule(ctx)
		consensusNetworkModule := xconsensus.NewModule(ctx)
		manager := module.NewManager(election, benchmarkModule, nonTxPoolModule, consensusNetworkModule)
		manager.SetElection(election.Name())
		manager.SetWorker(benchmarkModule.Name())
		manager.SetOrderTxPool(benchmarkModule.Name())
		manager.SetConsensusNetwork(consensusNetworkModule.Name())
		app := testutil.NewApp(manager)
		return app
	}, nil, nil, nil)

	if err := cliApp.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
