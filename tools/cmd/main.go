package main

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"os"

	"github.com/PlatONnetwork/AppChain-SDK/tools/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast"
	"gopkg.in/urfave/cli.v1"
)

var LogLevelFlag = cli.IntFlag{Name: "loglevel", Usage: "Log level, 1:error,2:warn,3:info,4:debug", Value: 0}

func initLog(lv int) {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stdout, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(lv))
	log.Root().SetHandler(glogger)
}
func main() {
	app := cli.NewApp()
	app.Action = initTool
	app.HideVersion = true // we have a command to print the version
	app.Copyright = "Copyright 2023 The AppChain-SDK Authors"
	app.Commands = []cli.Command{
		contracts.ContractCommand,
		cast.Command,
	}
	app.Before = func(ctx *cli.Context) error {
		initLog(ctx.Int(LogLevelFlag.Name))
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func initTool(ctx *cli.Context) error {
	return nil
}
