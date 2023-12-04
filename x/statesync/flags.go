package statesync

import (
	"gopkg.in/urfave/cli.v1"
)

var (
	StartBlockFlag = cli.Uint64Flag{
		Name:  "statesync.startblock",
		Usage: "Starting block height for scanning",
	}
)

func AddModuleInitFlags(app *cli.App) {
	app.Flags = append(app.Flags, StartBlockFlag)
}
