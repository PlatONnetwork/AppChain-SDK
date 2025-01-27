package main

import (
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/utils"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesync"
	"os"

	"github.com/PlatONnetwork/AppChain-SDK/x"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/PlatONnetwork/PlatON-Go/sdk/app"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	cliApp := cli.NewApp()
	x.AddModuleInitFlags(cliApp)
	checkpoint.AddModuleInitFlags(cliApp)
	statesync.AddModuleInitFlags(cliApp)
	utils.AddModuleInitFlags(cliApp)
	app.InitApp(cliApp, func(ctx *cli.Context) sdk.App {
		simApp, err := NewSimApp(ctx)
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
