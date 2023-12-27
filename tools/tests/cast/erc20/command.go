package erc20

import (
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"gopkg.in/urfave/cli.v1"
	"reflect"
)

var (
	Command = cli.Command{
		Action:    utils.MigrateFlags(run),
		Name:      "erc20",
		Usage:     "erc20",
		ArgsUsage: "",
		Flags: []cli.Flag{
			flags.RPCFlags,
			flags.KeyFlags,
			flags.AddressFlags,
			flags.MethodFlags,
			flags.TypeFlags,
		},
		Category:           "ERC20 COMMANDS",
		Description:        ``,
		HelpName:           "cast erc20",
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func initGlobal(ctx *cli.Context) (*ethclient.Client, *Erc20, *bind.TransactOpts, error) {
	client, stakeAddr, opt, err := flags.InitGlobal(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	manager, err := NewErc20(stakeAddr, client)
	return client, manager, opt, nil
}

func run(ctx *cli.Context) error {
	return flags.ExecuteCommand(ctx, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewErc20(address, client)
		return reflect.TypeOf(&manager.Erc20Transactor), reflect.ValueOf(&manager.Erc20Transactor)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		manager, _ := NewErc20(address, client)
		return reflect.TypeOf(&manager.Erc20Caller), reflect.ValueOf(&manager.Erc20Caller)
	}, func(address common.Address, client *ethclient.Client) (reflect.Type, reflect.Value) {
		filter, _ := NewErc20Filterer(address, client)
		return reflect.TypeOf(filter), reflect.ValueOf(filter)
	}, Erc20ABI)
}
