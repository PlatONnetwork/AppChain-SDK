package blocktime

import "gopkg.in/urfave/cli.v1"

var (
	NextBlockTimeFlag = cli.Uint64Flag{
		Name:  "blocktime.nextblocktime",
		Usage: "How long until the next block is produced based on the last block",
		Value: 400,
	}
	BlockProductionTimeoutFlag = cli.Uint64Flag{
		Name:  "blocktime.blockproductiontimeout",
		Usage: "A parameter that defines the maximum allowed time (in milliseconds) for a block production process to complete",
		Value: 1000,
	}
)

func AddModuleInitFlags(app *cli.App) {
	app.Flags = append(app.Flags, NextBlockTimeFlag)
	app.Flags = append(app.Flags, BlockProductionTimeoutFlag)
}
