package l2withdraw

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
		Name:               "l2.withdraw",
		Usage:              "l2.withdraw",
		ArgsUsage:          "",
		Flags:              flags.DefaultFlag,
		Category:           "L2 WITHDRAW COMMANDS",
		Description:        ``,
		HelpName:           "cast l2.withdraw",
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func run(ctx *cli.Context) error {
	return flags.ExecuteCommand(ctx, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewWithdraw(address, client)
		return reflect.TypeOf(&manager.WithdrawTransactor), reflect.ValueOf(&manager.WithdrawTransactor)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewWithdraw(address, client)
		return reflect.TypeOf(&manager.WithdrawCaller), reflect.ValueOf(&manager.WithdrawCaller)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		filter, _ := NewWithdrawFilterer(address, client)
		return reflect.TypeOf(filter), reflect.ValueOf(filter)
	}, WithdrawABI)
}
