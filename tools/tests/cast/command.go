package cast

import (
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/erc20"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l1deposit"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l1exithelper"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l1staking"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l2deposit"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l2gov"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l2reward"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l2staking"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l2statesender"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l2statesync"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l2testcontract"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l2upgrade"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l2validator"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/l2votetoken"
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
			l2statesync.Command,
			l2deposit.Command,
			l2statesender.Command,
			l2reward.Command,
			l2staking.Command,
			l2withdraw.Command,
			l2validator.Command,
			l2upgrade.Command,
			l2testcontract.Command,
			l2gov.Command,
			l2votetoken.Command,
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
