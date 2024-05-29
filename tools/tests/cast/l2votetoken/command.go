package l2votetoken

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
		Name:      "l2.votetoken",
		Usage:     "l2.votetoken",
		ArgsUsage: "",
		Flags: []cli.Flag{
			flags.RPCFlags,
			flags.KeyFlags,
			flags.AddressFlags,
			flags.MethodFlags,
			flags.TypeFlags,
		},
		Category:           "L2 VOTETOKEN COMMANDS",
		Description:        ``,
		HelpName:           "cast l2.votetoken",
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func run(ctx *cli.Context) error {
	return flags.ExecuteCommand(ctx, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewL2votetoken(address, client)
		return reflect.TypeOf(&manager.L2votetokenTransactor), reflect.ValueOf(&manager.L2votetokenTransactor)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewL2votetoken(address, client)
		return reflect.TypeOf(&manager.L2votetokenCaller), reflect.ValueOf(&manager.L2votetokenCaller)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		filter, _ := NewL2votetokenFilterer(address, client)
		return reflect.TypeOf(filter), reflect.ValueOf(filter)
	}, L2votetokenABI)
}
