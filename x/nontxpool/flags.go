package nontxpool

import (
	"gopkg.in/urfave/cli.v1"
)

var (
	TxsCacheSizeFlag = cli.IntFlag{
		Name:  "nontxpool.txscachesize",
		Usage: "How many cached transactions will be broadcast?",
		Value: 200,
	}
	BroadcastIntervalFlag = cli.IntFlag{
		Name:  "nontxpool.broadcastinterval",
		Usage: "How long is the interval for broadcasting transactions?",
		Value: 1000,
	}
	ValidateTxFlag = cli.BoolFlag{
		Name:  "nontxpool.validatetx",
		Usage: "Verify remote transactions?",
	}
)

func AddNonTxPoolFlags(app *cli.App) {
	app.Flags = append(app.Flags, TxsCacheSizeFlag)
	app.Flags = append(app.Flags, BroadcastIntervalFlag)
	app.Flags = append(app.Flags, ValidateTxFlag)
}
