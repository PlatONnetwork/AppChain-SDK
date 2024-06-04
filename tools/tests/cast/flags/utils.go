package flags

import (
	"context"
	"errors"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/transaction"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/json"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"gopkg.in/urfave/cli.v1"
	"math/big"
	"reflect"
	"strings"
)

func FindMethodArgs(ctx *cli.Context) (string, []string, error) {
	args := ctx.Args()
	return "", args, nil
	for i, arg := range args {
		if arg == "--method" {
			method := args[i+1]
			var inputs []string
			if len(args) > i+2 {
				inputs = args[i+2:]

				return method, inputs, nil
			}
		}
	}
	return "", nil, errors.New("not found method")
}

func InitGlobal(ctx *cli.Context) (*ethclient.Client, common.Address, *bind.TransactOpts, error) {
	url := ctx.String(RPCFlags.Name)
	cli, err := ethclient.Dial(url)
	if err != nil {
		return nil, common.Address{}, nil, err
	}
	address := ctx.String(AddressFlags.Name)
	stakeAddr := common.HexToAddress(address)
	var opt *bind.TransactOpts
	if ctx.String(TypeFlags.Name) == "send" {
		key, err := crypto.HexToECDSA(ctx.String(KeyFlags.Name))
		if err != nil {
			return nil, common.Address{}, nil, err
		}
		chainId, err := cli.ChainID(context.Background())
		if err != nil {
			return nil, common.Address{}, nil, err
		}
		opt, err = bind.NewKeyedTransactorWithChainID(key, chainId)
		if err != nil {
			return nil, common.Address{}, nil, err
		}
		opt.GasLimit = GasLimitFlags.Value
		if ctx.IsSet(GasLimitFlags.Name) {
			opt.GasLimit = ctx.Uint64(GasLimitFlags.Name)
		}
		opt.GasPrice = new(big.Int).SetUint64(GasPriceFlags.Value)
		if ctx.IsSet(GasPriceFlags.Name) {
			opt.GasPrice = new(big.Int).SetUint64(ctx.Uint64(GasPriceFlags.Name))
		}
		if ctx.IsSet(NonceFlags.Name) {
			opt.Nonce = new(big.Int).SetUint64(ctx.Uint64(NonceFlags.Name))
		}
	}
	return cli, stakeAddr, opt, nil
}

func ExecuteCommand(ctx *cli.Context,
	findTransactor func(common.Address, *ethclient.Client) (reflect.Type, reflect.Value),
	findCaller func(common.Address, *ethclient.Client) (reflect.Type, reflect.Value),
	findFilter func(common.Address, *ethclient.Client) (reflect.Type, reflect.Value),
	abiJson string) error {
	_, inputs, err := FindMethodArgs(ctx)
	if err != nil {
		return err
	}
	method := ctx.String(MethodFlags.Name)
	typeName := ctx.String(TypeFlags.Name)

	switch typeName {
	case "send":
		client, addr, opt, err := InitGlobal(ctx)
		if err != nil {
			return err
		}
		transactorType, transactorValue := findTransactor(addr, client)
		tx, err := transaction.Send(method, inputs, transactorType, transactorValue, opt)
		if err != nil {
			return err
		}
		_, err = transaction.WaitTx(client, tx.Hash())
		if err != nil {
			fmt.Println("send failed", tx.Hash())
			return err
		}
		fmt.Println("send success", tx.Hash())
		tx, _, err = client.TransactionByHash(context.Background(), tx.Hash())
		if err != nil {
			return err
		}
		bytes, _ := json.MarshalIndent(tx, " ", " ")
		fmt.Println("transaction:", string(bytes))
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			return err
		}
		bytes, _ = json.MarshalIndent(receipt, " ", " ")

		fmt.Println("receipt:", string(bytes))

	case "call":
		client, addr, _, err := InitGlobal(ctx)
		if err != nil {
			return err
		}
		callerType, callerValue := findCaller(addr, client)
		result, err := transaction.Call(method, inputs, callerType, callerValue, &bind.CallOpts{})
		if err != nil {
			return err
		}

		fmt.Println(string(result))

	case "logs":
		startBlock := ctx.Uint64(StartFlags.Name)
		var endBlock *uint64
		end := ctx.Uint64(EndFlags.Name)
		if end != 0 {
			endBlock = &end
		}
		client, addr, _, err := InitGlobal(ctx)
		if err != nil {
			return err
		}
		filterType, filterValue := findFilter(addr, client)
		result, err := transaction.FilterLog(method, inputs, filterType, filterValue, &bind.FilterOpts{
			Start:   startBlock,
			End:     endBlock,
			Context: context.Background(),
		})
		if err != nil {
			return err
		}

		fmt.Println(string(result))
	case "abi":
		parsed, err := abi.JSON(strings.NewReader(abiJson))
		if err != nil {
			return err
		}
		fmt.Println("methods:")
		for _, method := range parsed.Methods {
			fmt.Println("  ", method.Sig)
		}

		fmt.Println("events:")
		for _, event := range parsed.Events {
			var args []string
			for _, arg := range event.Inputs {
				index := ""
				if arg.Indexed {
					index = " indexed"
				}
				args = append(args, fmt.Sprintf("%s%s", arg.Type.String(), index))
			}
			fmt.Println("  ", fmt.Sprintf("%s(%s)", event.Name, strings.Join(args, ",")))
		}
	}
	return nil
}
