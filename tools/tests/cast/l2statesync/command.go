package l2statesync

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
		Name:               "l2.statesync",
		Usage:              "l2.statesync",
		ArgsUsage:          "",
		Flags:              flags.DefaultFlag,
		Category:           "L2 STATESYNC COMMANDS",
		Description:        ``,
		HelpName:           "cast l2.statesync",
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func run(ctx *cli.Context) error {
	return flags.ExecuteCommand(ctx, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewStatesync(address, client)
		return reflect.TypeOf(&manager.StatesyncTransactor), reflect.ValueOf(&manager.StatesyncTransactor)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewStatesync(address, client)
		return reflect.TypeOf(&manager.StatesyncCaller), reflect.ValueOf(&manager.StatesyncCaller)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		filter, _ := NewStatesyncFilterer(address, client)
		return reflect.TypeOf(filter), reflect.ValueOf(filter)
	}, StatesyncABI)
}
