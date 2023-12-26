package l1staking

import (
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/transaction"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"gopkg.in/urfave/cli.v1"
	"reflect"
)

var (
	Command = cli.Command{
		Action:    utils.MigrateFlags(run),
		Name:      "l1.staking",
		Usage:     "l1.staking",
		ArgsUsage: "",
		Flags: []cli.Flag{
			flags.RPCFlags,
			flags.KeyFlags,
			flags.AddressFlags,
			flags.MethodFlags,
			flags.TypeFlags,
		},
		Category:    "L1 STAKING COMMANDS",
		Description: ``,
	}
)

func initGlobal(ctx *cli.Context) (*ethclient.Client, *Stakemanager, *bind.TransactOpts, error) {
	client, stakeAddr, opt, err := flags.InitGlobal(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	manager, err := NewStakemanager(stakeAddr, client)
	return client, manager, opt, nil
}

func run(ctx *cli.Context) error {
	typeName := ctx.String(flags.TypeFlags.Name)
	client, manager, opt, err := initGlobal(ctx)
	if err != nil {
		return err
	}
	_, inputs, err := flags.FindMethodArgs(ctx)
	if err != nil {
		return err
	}
	method := ctx.String(flags.MethodFlags.Name)
	if typeName == "send" {
		tx, err := transaction.Send(method, inputs, reflect.TypeOf(&manager.StakemanagerTransactor), reflect.ValueOf(&manager.StakemanagerTransactor), opt)
		if err != nil {
			return err
		}
		_, err = transaction.WaitTx(client, tx.Hash())
		if err != nil {
			return err
		}
		fmt.Println("send success", tx.Hash())
	} else {
		result, err := transaction.Call(method, inputs, reflect.TypeOf(&manager.StakemanagerCaller), reflect.ValueOf(&manager.StakemanagerCaller), &bind.CallOpts{})
		if err != nil {
			return err
		}

		fmt.Println(string(result))
	}
	return nil
}
