package asyncblock

import "gopkg.in/urfave/cli.v1"

const (
	BLsKeyFlagName = "cbft.blskey"
)

var (
	ConcurrencyLevelFlag = cli.IntFlag{
		Name:  "asyncblock.concurrency_level",
		Usage: "Number of goroutines use to parallel execute transactions",
		Value: 4,
	}
	ForceSequentialFlag = cli.BoolFlag{
		Name:  "asyncblock.force_sequential",
		Usage: "Force sequential execute transactions",
	}
	TxsBatchFlag = cli.IntFlag{
		Name:  "asyncblock.txs_batch",
		Usage: "Number of transactions for a batch to parallel exectue",
		Value: 64,
	}

	EntrySizeFlag = cli.IntFlag{
		Name:  "asyncblock.entrysize",
		Usage: "How many transactions can be bundled at most",
		Value: 1000,
	}
	SplitThresholdFlag = cli.IntFlag{
		Name:  "asyncblock.splitthreshold",
		Usage: "How many transactions need to be split",
		Value: 2000,
	}
	ComputeSenderThreadFlag = cli.IntFlag{
		Name:  "asyncblock.computersenderthread",
		Usage: "How many threads compute transaction sender",
		Value: 4,
	}
)

func AddAsyncBlockFlags(app *cli.App) {
	app.Flags = append(app.Flags, ConcurrencyLevelFlag)
	app.Flags = append(app.Flags, ForceSequentialFlag)
	app.Flags = append(app.Flags, TxsBatchFlag)
	app.Flags = append(app.Flags, EntrySizeFlag)
	app.Flags = append(app.Flags, SplitThresholdFlag)
	app.Flags = append(app.Flags, ComputeSenderThreadFlag)
}
