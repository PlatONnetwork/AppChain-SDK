package flags

import (
	"context"
	"errors"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"gopkg.in/urfave/cli.v1"
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
	key, err := crypto.HexToECDSA(ctx.String(KeyFlags.Name))
	if err != nil {
		return nil, common.Address{}, nil, err
	}
	chainId, err := cli.ChainID(context.Background())
	if err != nil {
		return nil, common.Address{}, nil, err
	}
	opt, err := bind.NewKeyedTransactorWithChainID(key, chainId)
	if err != nil {
		return nil, common.Address{}, nil, err
	}
	return cli, stakeAddr, opt, nil
}
