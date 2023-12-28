package l2deposit

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
		Action:    utils.MigrateFlags(run),
		Name:      "l2.deposit",
		Usage:     "l2.deposit",
		ArgsUsage: "",
		Flags: []cli.Flag{
			flags.RPCFlags,
			flags.KeyFlags,
			flags.AddressFlags,
			flags.MethodFlags,
			flags.TypeFlags,
		},
		Category:           "L2 DEPOSIT COMMANDS",
		Description:        ``,
		HelpName:           "cast l2.deposit",
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func run(ctx *cli.Context) error {
	return flags.ExecuteCommand(ctx, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewDeposit(address, client)
		return reflect.TypeOf(&manager.DepositTransactor), reflect.ValueOf(&manager.DepositTransactor)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewDeposit(address, client)
		return reflect.TypeOf(&manager.DepositCaller), reflect.ValueOf(&manager.DepositCaller)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		filter, _ := NewDepositFilterer(address, client)
		return reflect.TypeOf(filter), reflect.ValueOf(filter)
	}, DepositABI)
}
