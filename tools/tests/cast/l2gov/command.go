package l2gov

import (
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"gopkg.in/urfave/cli.v1"
	"reflect"
)

var (
	Command = cli.Command{
		Action:             utils.MigrateFlags(run),
		Name:               "l2.gov",
		Usage:              "l2.gov",
		ArgsUsage:          "",
		Flags:              flags.DefaultFlag,
		Category:           "L2 GOV COMMANDS",
		Description:        ``,
		HelpName:           "cast l2.gov",
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func run(ctx *cli.Context) error {
	return flags.ExecuteCommand(ctx, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewL2gov(address, client)
		return reflect.TypeOf(&manager.L2govTransactor), reflect.ValueOf(&manager.L2govTransactor)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewL2gov(address, client)
		return reflect.TypeOf(&manager.L2govCaller), reflect.ValueOf(&manager.L2govCaller)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		filter, _ := NewL2govFilterer(address, client)
		return reflect.TypeOf(filter), reflect.ValueOf(filter)
	}, L2govABI)
}
