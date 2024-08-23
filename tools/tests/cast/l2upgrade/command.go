package l2upgrade

import (
	"reflect"

	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"gopkg.in/urfave/cli.v1"
)

var (
	Command = cli.Command{
		Action:             utils.MigrateFlags(run),
		Name:               "l2.upgrade",
		Usage:              "l2.upgrade",
		ArgsUsage:          "",
		Flags:              flags.DefaultFlag,
		Category:           "L2 UPGRADE COMMANDS",
		Description:        "",
		HelpName:           "cast l2.upgrade",
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func run(ctx *cli.Context) error {
	return flags.ExecuteCommand(ctx, func(a common.Address, c *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewL2upgrade(a, c)
		return reflect.TypeOf(&manager.L2upgradeTransactor), reflect.ValueOf(&manager.L2upgradeTransactor)
	}, func(a common.Address, c *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewL2upgrade(a, c)
		return reflect.TypeOf(&manager.L2upgradeCaller), reflect.ValueOf(&manager.L2upgradeCaller)
	}, func(a common.Address, c *ethclient.Client) (reflect.Type, reflect.Value) {
		filter, _ := NewL2upgradeFilterer(a, c)
		return reflect.TypeOf(filter), reflect.ValueOf(filter)
	}, L2upgradeABI)
}
