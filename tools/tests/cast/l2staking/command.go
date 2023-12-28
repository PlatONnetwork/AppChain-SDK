package l2staking

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
		Name:      "l2.staking",
		Usage:     "l2.staking",
		ArgsUsage: "",
		Flags: []cli.Flag{
			flags.RPCFlags,
			flags.KeyFlags,
			flags.AddressFlags,
			flags.MethodFlags,
			flags.TypeFlags,
		},
		Category:           "L2 STAKING COMMANDS",
		Description:        ``,
		HelpName:           "cast l2.staking",
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func run(ctx *cli.Context) error {
	return flags.ExecuteCommand(ctx, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewStaking(address, client)
		return reflect.TypeOf(&manager.StakingTransactor), reflect.ValueOf(&manager.StakingTransactor)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewStaking(address, client)
		return reflect.TypeOf(&manager.StakingCaller), reflect.ValueOf(&manager.StakingCaller)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		filter, _ := NewStakingFilterer(address, client)
		return reflect.TypeOf(filter), reflect.ValueOf(filter)
	}, StakingABI)
}
