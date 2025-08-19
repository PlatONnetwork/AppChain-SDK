package miner

import "gopkg.in/urfave/cli.v1"

var (
	ConcurrencyLevelFlag = cli.IntFlag{
		Name:  "miner.concurrency_level",
		Usage: "Number of goroutines use to parallel execute transactions",
		Value: 4,
	}
	ForceSequentialFlag = cli.BoolFlag{
		Name:  "miner.force_sequential",
		Usage: "Force sequential execute transactions",
	}
	TxsBatchFlag = cli.IntFlag{
		Name:  "miner.txs_batch",
		Usage: "Number of transactions for a batch to parallel execute",
		Value: 64,
	}
)

func AddModuleInitFlags(app *cli.App) {
	app.Flags = append(app.Flags, ConcurrencyLevelFlag)
	app.Flags = append(app.Flags, ForceSequentialFlag)
	app.Flags = append(app.Flags, TxsBatchFlag)
}
