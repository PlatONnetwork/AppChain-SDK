package l2statesender

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
		Name:      "l2.statesender",
		Usage:     "l2.statesender",
		ArgsUsage: "",
		Flags: []cli.Flag{
			flags.RPCFlags,
			flags.KeyFlags,
			flags.AddressFlags,
			flags.MethodFlags,
			flags.TypeFlags,
		},
		Category:           "L2 STATESENDER COMMANDS",
		Description:        ``,
		HelpName:           "cast l2.statesender",
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func run(ctx *cli.Context) error {
	return flags.ExecuteCommand(ctx, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewStatesender(address, client)
		return reflect.TypeOf(&manager.StatesenderTransactor), reflect.ValueOf(&manager.StatesenderTransactor)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewStatesender(address, client)
		return reflect.TypeOf(&manager.StatesenderCaller), reflect.ValueOf(&manager.StatesenderCaller)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		filter, _ := NewStatesenderFilterer(address, client)
		return reflect.TypeOf(filter), reflect.ValueOf(filter)
	}, StatesenderABI)
}
