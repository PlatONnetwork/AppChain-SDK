package l2testcontract

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
		Action:    utils.MigrateFlags(run),
		Name:      "l2.testcontract",
		Usage:     "l2.testcontract",
		ArgsUsage: "",
		Flags: []cli.Flag{
			flags.RPCFlags,
			flags.KeyFlags,
			flags.AddressFlags,
			flags.MethodFlags,
			flags.TypeFlags,
		},
		Category:           "L2 TEST CONTRACT",
		Description:        "",
		HelpName:           "cast l2.testcontract",
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func run(ctx *cli.Context) error {
	return flags.ExecuteCommand(ctx, func(a common.Address, c *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewL2testcontract(a, c)
		return reflect.TypeOf(&manager.L2testcontractTransactor), reflect.ValueOf(&manager.L2testcontractTransactor)
	}, func(a common.Address, c *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewL2testcontract(a, c)
		return reflect.TypeOf(&manager.L2testcontractCaller), reflect.ValueOf(&manager.L2testcontractCaller)
	}, func(a common.Address, c *ethclient.Client) (reflect.Type, reflect.Value) {
		filter, _ := NewL2testcontractFilterer(a, c)
		return reflect.TypeOf(filter), reflect.ValueOf(filter)
	}, L2testcontractABI)
}
