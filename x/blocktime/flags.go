package blocktime

import "gopkg.in/urfave/cli.v1"

var (
	NextBlockTimeFlag = cli.Uint64Flag{
		Name:  "blocktime.nextblocktime",
		Usage: "How long until the next block is produced based on the last block",
		Value: 400,
	}
)

func AddBlockTimeFlags(app *cli.App) {
	app.Flags = append(app.Flags, NextBlockTimeFlag)
}
