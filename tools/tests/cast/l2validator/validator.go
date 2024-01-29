package l2validator

import (
	"context"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"gopkg.in/urfave/cli.v1"
)

var (
	Command = cli.Command{
		Action:    utils.MigrateFlags(run),
		Name:      "l2.validator",
		Usage:     "l2.validator",
		ArgsUsage: "",
		Flags: []cli.Flag{
			flags.RPCFlags,
			flags.RoundFlags,
			flags.EpochFlags,
		},
		Category:           "L2 VALIDATOR COMMANDS",
		Description:        ``,
		HelpName:           "cast l2.validator",
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func run(ctx *cli.Context) error {
	url := ctx.String(flags.RPCFlags.Name)
	cli, err := ethclient.Dial(url)
	if err != nil {
		return err
	}
	blockNumber, err := cli.BlockNumber(context.Background())
	if err != nil {
		return err
	}
	round := ctx.Uint64(flags.RoundFlags.Name)
	epoch := ctx.Uint64(flags.EpochFlags.Name)

	fmt.Println("round:", blockNumber/round, "epoch", blockNumber/epoch)
	return nil
}
