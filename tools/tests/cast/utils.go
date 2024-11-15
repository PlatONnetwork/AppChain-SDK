package cast

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/umbracle/ethgo"
	ethabi "github.com/umbracle/ethgo/abi"
	"gopkg.in/urfave/cli.v1"
	"math/big"
	"os"
	"time"
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
	url := ctx.String(flags.RPCFlags.Name)
	cli, err := ethclient.Dial(url)
	if err != nil {
		return nil, common.Address{}, nil, err
	}
	address := ctx.String(flags.AddressFlags.Name)
	stakeAddr := common.HexToAddress(address)
	var opt *bind.TransactOpts
	if ctx.String(flags.TypeFlags.Name) == "send" {
		key, err := crypto.HexToECDSA(ctx.String(flags.KeyFlags.Name))
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
		opt.GasLimit = flags.GasLimitFlags.Value
		if ctx.IsSet(flags.GasLimitFlags.Name) {
			opt.GasLimit = ctx.Uint64(flags.GasLimitFlags.Name)
		}
		opt.GasPrice = new(big.Int).SetUint64(flags.GasPriceFlags.Value)
		if ctx.IsSet(flags.GasPriceFlags.Name) {
			opt.GasPrice = new(big.Int).SetUint64(ctx.Uint64(flags.GasPriceFlags.Name))
		}
		if ctx.IsSet(flags.NonceFlags.Name) {
			opt.Nonce = new(big.Int).SetUint64(ctx.Uint64(flags.NonceFlags.Name))
		}
	}
	return cli, stakeAddr, opt, nil
}

func GetAbi(ctx *cli.Context) (*ethabi.ABI, error) {
	module := ctx.String(flags.ModuleFlags.Name)
	fileName := ctx.String(flags.AbiFileFlags.Name)
	abiJson := ""
	if fileName == "" {
		abiJson = Modules[module]
		if len(abiJson) == 0 {
			return nil, fmt.Errorf("find abi failed:%s", module)
		}
	} else {
		abiBytes, err := os.ReadFile(fileName)
		if err != nil {
			return nil, err
		}
		abiJson = string(abiBytes)
	}

	abi, err := ethabi.NewABI(abiJson)

	return abi, err
}

func FilterLog(abi *ethabi.ABI, method string, inputs []string, to common.Address, cli *ethclient.Client, opt *bind.FilterOpts) ([]map[string]interface{}, error) {
	event, ok := abi.Events[method]
	if !ok {
		return nil, fmt.Errorf("event not found")
	}
	topics, err := MakeTopics(event.Inputs, inputs)
	if err != nil {
		return nil, err
	}

	config := platon.FilterQuery{
		Addresses: []common.Address{to},
		Topics:    topics,
		FromBlock: new(big.Int).SetUint64(opt.Start),
	}
	if opt.End != nil {
		config.ToBlock = new(big.Int).SetUint64(*opt.End)
	}

	logs, err := cli.FilterLogs(ensureContext(nil), config)
	if err != nil {
		return nil, err
	}
	var res []map[string]interface{}
	for _, log := range logs {
		topics := make([]ethgo.Hash, len(log.Topics))
		for i, h := range log.Topics {
			topics[i] = ethgo.BytesToHash(h.Bytes())
		}
		r, err := ethabi.ParseLog(event.Inputs, &ethgo.Log{
			Removed:          log.Removed,
			LogIndex:         uint64(log.Index),
			TransactionIndex: uint64(log.Index),
			TransactionHash:  ethgo.BytesToHash(log.TxHash.Bytes()),
			BlockHash:        ethgo.BytesToHash(log.BlockHash.Bytes()),
			BlockNumber:      log.BlockNumber,
			Address:          ethgo.BytesToAddress(log.Address.Bytes()),
			Data:             log.Data,
		})

		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}
func MakeTopics(t *ethabi.Type, inputs []string) ([][]common.Hash, error) {
	var res [][]common.Hash
	for i, elem := range t.TupleElems() {
		if elem.Indexed {
			if len(inputs) <= i {
				return res, nil
			}
			r, err := makeTopic(elem.Elem, inputs[i])
			if err != nil {
				return nil, err
			}
			res = append(res, r)
		}
	}
	return res, nil
}

func makeTopic(t *ethabi.Type, input string) ([]common.Hash, error) {
	var res []common.Hash
	switch t.Kind() {
	case ethabi.KindBool:
		var array []bool
		err := json.Unmarshal([]byte(input), &array)
		if err != nil {
			var value bool
			err := json.Unmarshal([]byte(input), &value)
			if err != nil {
				return nil, err
			}
			hash, err := ethabi.EncodeTopic(t, value)
			if err != nil {
				return nil, err
			}
			res = append(res, common.BytesToHash(hash.Bytes()))
		} else {
			for _, a := range array {
				hash, err := ethabi.EncodeTopic(t, a)
				if err != nil {
					return nil, err
				}
				res = append(res, common.BytesToHash(hash.Bytes()))
			}
		}

	case ethabi.KindUInt, ethabi.KindInt:
		var array []*big.Int
		err := json.Unmarshal([]byte(input), &array)
		if err != nil {
			var value []*big.Int
			err := json.Unmarshal([]byte(input), &value)
			if err != nil {
				return nil, err
			}
			hash, err := ethabi.EncodeTopic(t, value)
			if err != nil {
				return nil, err
			}
			res = append(res, common.BytesToHash(hash.Bytes()))
		} else {
			for _, a := range array {
				hash, err := ethabi.EncodeTopic(t, a)
				if err != nil {
					return nil, err
				}
				res = append(res, common.BytesToHash(hash.Bytes()))
			}
		}

	case ethabi.KindAddress:
		var array []string
		err := json.Unmarshal([]byte(input), &array)
		if err != nil {
			var value []string
			err := json.Unmarshal([]byte(input), &value)
			if err != nil {
				return nil, err
			}
			hash, err := ethabi.EncodeTopic(t, value)
			if err != nil {
				return nil, err
			}
			res = append(res, common.BytesToHash(hash.Bytes()))
		} else {
			for _, a := range array {
				hash, err := ethabi.EncodeTopic(t, a)
				if err != nil {
					return nil, err
				}
				res = append(res, common.BytesToHash(hash.Bytes()))
			}
		}

	}
	return res, nil
}

func Send(abi *ethabi.ABI, methodName string, inputs []string, to common.Address, cli *ethclient.Client, opt *bind.TransactOpts) (*types.Transaction, error) {
	method := abi.GetMethod(methodName)
	input, err := method.Encode(inputs)
	if err != nil {
		return nil, fmt.Errorf("encode input failed:%s", err.Error())
	}
	rawTx, err := createLegacyTx(opt, cli, &to, input)
	if err != nil {
		return nil, err
	}
	signedTx, err := opt.Signer(opt.From, rawTx)
	if err != nil {
		return nil, err
	}
	err = cli.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}
func Call(abi *ethabi.ABI, methodName string, inputs []string, to common.Address, cli *ethclient.Client, opt *bind.CallOpts) ([]byte, error) {
	method := abi.GetMethod(methodName)
	input, err := method.Encode(inputs)
	if err != nil {
		return nil, fmt.Errorf("encode input failed:%s", err.Error())
	}
	return cli.CallContract(context.Background(), platon.CallMsg{
		From: opt.From,
		To:   &to,
		Data: input,
	}, nil)
}
func createLegacyTx(opts *bind.TransactOpts, cli *ethclient.Client, contract *common.Address, input []byte) (*types.Transaction, error) {
	if opts.GasFeeCap != nil || opts.GasTipCap != nil {
		return nil, fmt.Errorf("maxFeePerGas or maxPriorityFeePerGas specified but london is not active yet")
	}
	// Normalize value
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}
	// Estimate GasPrice
	gasPrice := opts.GasPrice
	if gasPrice == nil {
		price, err := cli.SuggestGasPrice(ensureContext(opts.Context))
		if err != nil {
			return nil, err
		}
		gasPrice = price
	}
	// Estimate GasLimit
	gasLimit := opts.GasLimit
	if opts.GasLimit == 0 {
		var err error
		gasLimit, err = estimateGasLimit(opts, cli, contract, input, gasPrice, nil, nil, value)
		if err != nil {
			return nil, err
		}
	}
	// create the transaction
	nonce, err := getNonce(opts, cli)
	if err != nil {
		return nil, err
	}
	baseTx := &types.LegacyTx{
		To:       contract,
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		Value:    value,
		Data:     input,
	}
	return types.NewTx(baseTx), nil
}
func estimateGasLimit(opts *bind.TransactOpts, cli *ethclient.Client, contract *common.Address, input []byte, gasPrice, gasTipCap, gasFeeCap, value *big.Int) (uint64, error) {

	msg := platon.CallMsg{
		From:      opts.From,
		To:        contract,
		GasPrice:  gasPrice,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     value,
		Data:      input,
	}
	return cli.EstimateGas(ensureContext(opts.Context), msg)
}

func getNonce(opts *bind.TransactOpts, cli *ethclient.Client) (uint64, error) {
	if opts.Nonce == nil {
		return cli.PendingNonceAt(ensureContext(opts.Context), opts.From)
	} else {
		return opts.Nonce.Uint64(), nil
	}
}
func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return ctx
}
func WaitTx(client *ethclient.Client, hash common.Hash) (*types.Receipt, error) {
	for i := 0; i < 20; i++ {
		receipt, _ := client.TransactionReceipt(context.Background(), hash)
		if receipt != nil {
			return receipt, nil
		}
		time.Sleep(time.Second)
	}
	return nil, fmt.Errorf("wait tx time out")
}
