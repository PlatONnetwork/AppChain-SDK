package x

import "gopkg.in/urfave/cli.v1"

var (
	RootchainNodeRPCFlag = cli.StringFlag{
		Name:  "rootchain-node-rpc",
		Usage: "Rootchain node RPC endpoint",
	}
)

func AddModuleInitFlags(app *cli.App) {
	app.Flags = append(app.Flags, RootchainNodeRPCFlag)
}
