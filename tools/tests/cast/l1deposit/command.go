package l1deposit

import (
	"context"
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
		Name:      "l1.deposit",
		Usage:     "l1.deposit",
		ArgsUsage: "",
		Flags: []cli.Flag{
			flags.RPCFlags,
			flags.KeyFlags,
			flags.AddressFlags,
			flags.MethodFlags,
			flags.TypeFlags,
		},
		Category:    "L1 DEPOSIT COMMANDS",
		Description: ``,
	}
)

func initGlobal(ctx *cli.Context) (*ethclient.Client, *Deposit, *DepositFilterer, *bind.TransactOpts, error) {
	client, stakeAddr, opt, err := flags.InitGlobal(ctx)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	manager, err := NewDeposit(stakeAddr, client)
	filter, err := NewDepositFilterer(stakeAddr, client)
	return client, manager, filter, opt, nil
}

func run(ctx *cli.Context) error {
	typeName := ctx.String(flags.TypeFlags.Name)
	client, manager, filter, opt, err := initGlobal(ctx)
	if err != nil {
		return err
	}
	_, inputs, err := flags.FindMethodArgs(ctx)
	if err != nil {
		return err
	}
	method := ctx.String(flags.MethodFlags.Name)
	if typeName == "send" {
		tx, err := transaction.Send(method, inputs, reflect.TypeOf(&manager.DepositTransactor), reflect.ValueOf(&manager.DepositTransactor), opt)
		if err != nil {
			return err
		}
		_, err = transaction.WaitTx(client, tx.Hash())
		if err != nil {
			return err
		}
		fmt.Println("send success", tx.Hash())
	} else if typeName == "call" {
		result, err := transaction.Call(method, inputs, reflect.TypeOf(&manager.DepositCaller), reflect.ValueOf(&manager.DepositCaller), &bind.CallOpts{})
		if err != nil {
			return err
		}

		fmt.Println(string(result))
	} else if typeName == "logs" {
		startBlock := ctx.Uint64(flags.StartFlags.Name)
		var endBlock *uint64
		end := ctx.Uint64(flags.EndFlags.Name)
		if end != 0 {
			endBlock = &end
		}
		result, err := transaction.FilterLog(method, inputs, reflect.TypeOf(filter), reflect.ValueOf(filter), &bind.FilterOpts{
			Start:   startBlock,
			End:     endBlock,
			Context: context.Background(),
		})
		if err != nil {
			return err
		}

		fmt.Println(string(result))
	}
	return nil
}
