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
	PendingLimitFlag = cli.IntFlag{
		Name:  "nontxpool.pendinglimit",
		Usage: "How many transactions are packaged",
		Value: 10000,
	}
	TxsPerAccountFlag = cli.IntFlag{
		Name:  "nontxpool.txsperaccount",
		Usage: "How many transactions are packaged per account",
		Value: 100,
	}
	BroadcastFlag = cli.BoolFlag{
		Name:  "nontxpool.broadcast",
		Usage: "Broadcast transaction",
	}
	GlobalTxCountFlag = cli.IntFlag{
		Name:  "nontxpool.globaltxcount",
		Usage: "Maximum number of transactions for package",
		Value: 10000,
	}
)

func AddNonTxPoolFlags(app *cli.App) {
	app.Flags = append(app.Flags, TxsCacheSizeFlag)
	app.Flags = append(app.Flags, BroadcastIntervalFlag)
	app.Flags = append(app.Flags, ValidateTxFlag)
	app.Flags = append(app.Flags, PendingLimitFlag)
	app.Flags = append(app.Flags, TxsPerAccountFlag)
	app.Flags = append(app.Flags, BroadcastFlag)
	app.Flags = append(app.Flags, GlobalTxCountFlag)

}
