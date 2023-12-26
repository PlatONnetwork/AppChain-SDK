package cast

import (
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/erc20"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l1deposit"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l1exithelper"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l1staking"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l2reward"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l2staking"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l2withdraw"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"gopkg.in/urfave/cli.v1"
)

var (
	Command = cli.Command{
		Action:    utils.MigrateFlags(run),
		Name:      "cast",
		Usage:     "cast",
		ArgsUsage: "",
		Subcommands: []cli.Command{
			erc20.Command,
			l1deposit.Command,
			l1staking.Command,
			l1exithelper.Command,
			l2reward.Command,
			l2staking.Command,
			l2withdraw.Command,
		},
		Flags: []cli.Flag{
			flags.RPCFlags,
			flags.KeyFlags,
		},
		Category:    "CAST COMMANDS",
		Description: ``,
	}
)

func run(context2 *cli.Context) error {
	return nil
}
