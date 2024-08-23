package l1staking

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
		Name:               "l1.staking",
		Usage:              "l1.staking",
		ArgsUsage:          "",
		Flags:              flags.DefaultFlag,
		Category:           "L1 STAKING COMMANDS",
		Description:        ``,
		HelpName:           "cast l1.staking",
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func run(ctx *cli.Context) error {
	return flags.ExecuteCommand(ctx, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewStakemanager(address, client)
		return reflect.TypeOf(&manager.StakemanagerTransactor), reflect.ValueOf(&manager.StakemanagerTransactor)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewStakemanager(address, client)
		return reflect.TypeOf(&manager.StakemanagerCaller), reflect.ValueOf(&manager.StakemanagerCaller)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		filter, _ := NewStakemanagerFilterer(address, client)
		return reflect.TypeOf(filter), reflect.ValueOf(filter)
	}, StakemanagerABI)
}
