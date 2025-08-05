package main

import (
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/x/asyncblock"
	"github.com/PlatONnetwork/AppChain-SDK/x/benchmark"
	"github.com/PlatONnetwork/AppChain-SDK/x/blocktime"
	"github.com/PlatONnetwork/AppChain-SDK/x/nontxpool"
	"os"

	"github.com/PlatONnetwork/AppChain-SDK/utils"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesync"

	"gopkg.in/urfave/cli.v1"

	"github.com/PlatONnetwork/AppChain-SDK/x/miner"

	"github.com/PlatONnetwork/AppChain-SDK/x"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/PlatONnetwork/PlatON-Go/sdk/app"
)

func main() {
	cliApp := cli.NewApp()
	x.AddModuleInitFlags(cliApp)
	checkpoint.AddModuleInitFlags(cliApp)
	statesync.AddModuleInitFlags(cliApp)
	miner.AddModuleInitFlags(cliApp)
	utils.AddModuleInitFlags(cliApp)
	benchmark.AddBenchmarkFlags(cliApp)
	nontxpool.AddNonTxPoolFlags(cliApp)
	asyncblock.AddAsyncBlockFlags(cliApp)
	blocktime.AddBlockTimeFlags(cliApp)
	app.InitApp(cliApp, func(ctx *cli.Context) sdk.App {
		simApp, err := NewSimApp(ctx)
		//simApp, err := NewConsensusApp(ctx)
		if err != nil {
			panic(fmt.Sprintf("Create simple app error: %v", err))
		}
		return simApp
	}, nil, nil, nil)

	if err := cliApp.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
