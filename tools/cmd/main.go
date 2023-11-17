package main

import (
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/tools/contracts"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Action = initTool
	app.HideVersion = true // we have a command to print the version
	app.Copyright = "Copyright 2023 The AppChain-SDK Authors"
	app.Commands = []cli.Command{
		contracts.ContractCommand,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func initTool(ctx *cli.Context) error {
	return nil
}
