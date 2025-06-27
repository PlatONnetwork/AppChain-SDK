package asyncblock

import "gopkg.in/urfave/cli.v1"

const (
	BLsKeyFlagName = "cbft.blskey"
)

var (
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
)

func AddAsyncBlockFlags(app *cli.App) {
	app.Flags = append(app.Flags, EntrySizeFlag)
	app.Flags = append(app.Flags, SplitThresholdFlag)
}
