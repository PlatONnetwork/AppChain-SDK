package main

import (
	"fmt"
	deploytools "github.com/PlatONnetwork/AppChain-SDK/tools/deploy"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := deploytools.CreateApp()
	app.Commands = append(app.Commands, []cli.Command{
		DeployTemplateCommand,
		*createNodeCommand(),
		DeployChildChainCommand,
		CreateGenesisCommand,
		deploytools.CreateAnsibleCommand,
		*createAnsibleNodeCommand(),
	}...)
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
