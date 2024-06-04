package l2reward

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
		Name:               "l2.reward",
		Usage:              "l2.reward",
		ArgsUsage:          "",
		Flags:              flags.DefaultFlag,
		Category:           "L2 REWARD COMMANDS",
		Description:        ``,
		HelpName:           "cast l2.reward",
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func run(ctx *cli.Context) error {
	return flags.ExecuteCommand(ctx, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewReward(address, client)
		return reflect.TypeOf(&manager.RewardTransactor), reflect.ValueOf(&manager.RewardTransactor)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewReward(address, client)
		return reflect.TypeOf(&manager.RewardCaller), reflect.ValueOf(&manager.RewardCaller)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		filter, _ := NewRewardFilterer(address, client)
		return reflect.TypeOf(filter), reflect.ValueOf(filter)
	}, RewardABI)
}
