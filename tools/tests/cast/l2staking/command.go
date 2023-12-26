package l2staking

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
		Category:    "L2 STAKING COMMANDS",
		Description: ``,
	}
)

func initGlobal(ctx *cli.Context) (*ethclient.Client, *Staking, *StakingFilterer, *bind.TransactOpts, error) {
	client, stakeAddr, opt, err := flags.InitGlobal(ctx)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	manager, err := NewStaking(stakeAddr, client)
	filter, err := NewStakingFilterer(stakeAddr, client)
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
		tx, err := transaction.Send(method, inputs, reflect.TypeOf(&manager.StakingTransactor), reflect.ValueOf(&manager.StakingTransactor), opt)
		if err != nil {
			return err
		}
		_, err = transaction.WaitTx(client, tx.Hash())
		if err != nil {
			return err
		}
		fmt.Println("send success", tx.Hash())
	} else if typeName == "call" {
		result, err := transaction.Call(method, inputs, reflect.TypeOf(&manager.StakingCaller), reflect.ValueOf(&manager.StakingCaller), &bind.CallOpts{})
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
