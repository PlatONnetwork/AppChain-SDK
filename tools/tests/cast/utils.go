package cast

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/umbracle/ethgo"
	ethabi "github.com/umbracle/ethgo/abi"
	"gopkg.in/urfave/cli.v1"
	"math/big"
	"os"
	"strings"
	"time"
)

func FindMethodArgs(ctx *cli.Context) (string, []interface{}, error) {
	var args []interface{}
	for i, arg := range ctx.Args() {
		if len(arg) != 0 && arg[0] == byte('{') && json.Valid([]byte(arg)) {
			var argMap map[string]interface{}
			if err := json.Unmarshal([]byte(arg), &argMap); err != nil {
				return "", nil, fmt.Errorf("decode args [%d][%s] failed: %s", i, arg, err.Error())
			}
			args = append(args, argMap)
		} else if len(arg) != 0 && arg[0] == byte('[') && json.Valid([]byte(arg)) {
			var argSlice []string
			if err := json.Unmarshal([]byte(arg), &argSlice); err != nil {
				return "", nil, fmt.Errorf("decode args [%d][%s] failed: %s", i, arg, err.Error())
			}
			args = append(args, argSlice)
		} else {
			args = append(args, arg)
		}
	}
	return "", args, nil
}

func InitGlobal(ctx *cli.Context) (*ethclient.Client, common.Address, *bind.TransactOpts, error) {
	url := ctx.String(flags.AppchainRPCFlag.Name)
	cli, err := ethclient.Dial(url)
	if err != nil {
		return nil, common.Address{}, nil, err
	}
	address := common.HexToAddress(ctx.String(flags.AddressFlags.Name))
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
		log.Debug("Chainid", "value", chainId)
		if err != nil {
			return nil, common.Address{}, nil, err
		}
		if ctx.IsSet(flags.GasLimitFlags.Name) {
			opt.GasLimit = ctx.Uint64(flags.GasLimitFlags.Name)
			log.Debug("GasLimit", "value", opt.GasLimit)
		}
		opt.GasPrice = new(big.Int).SetUint64(flags.GasPriceFlags.Value)
		if ctx.IsSet(flags.GasPriceFlags.Name) {
			opt.GasPrice = new(big.Int).SetUint64(ctx.Uint64(flags.GasPriceFlags.Name))
			log.Debug("GasPrice", "value", opt.GasPrice)
		}
		if ctx.IsSet(flags.NonceFlags.Name) {
			opt.Nonce = new(big.Int).SetUint64(ctx.Uint64(flags.NonceFlags.Name))
			log.Debug("Nonce", "value", opt.Nonce)

		}
	}
	log.Debug("Init args", "contractAddress", address.Hex(), "url", url)
	return cli, address, opt, nil
}

func GetAbi(ctx *cli.Context) (*ethabi.ABI, error) {
	module := ctx.String(flags.ModuleFlags.Name)
	fileName := ctx.String(flags.AbiFileFlags.Name)
	log.Debug("Get abi", "module", module, "fileName", fileName)
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

func FilterLog(abi *ethabi.ABI, method string, inputs []interface{}, to common.Address, cli *ethclient.Client, opt *bind.FilterOpts) ([]map[string]interface{}, error) {
	log.Debug("Filterlog", "method", method, "inputs", inputs, "to", to.Hex(), "start", opt.Start, "end", opt.End)
	event, ok := abi.Events[method]
	if !ok {
		return nil, fmt.Errorf("event not found:%s", method)
	}
	topics, err := MakeTopics(event, inputs)
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
			Topics:           topics,
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
func MakeTopics(event *ethabi.Event, inputs []interface{}) ([][]common.Hash, error) {
	name := []common.Hash{common.BytesToHash(event.ID().Bytes())}
	res := [][]common.Hash{name}
	for i, elem := range event.Inputs.TupleElems() {
		if elem.Indexed {
			if len(inputs) <= i {
				break
			}
			r, err := makeTopic(elem.Elem, inputs[i])
			if err != nil {
				return nil, err
			}
			res = append(res, []common.Hash{r})
		}
	}
	return res, nil
}

func makeTopic(t *ethabi.Type, input interface{}) (common.Hash, error) {
	var hash ethgo.Hash
	var err error
	switch t.Kind() {
	case ethabi.KindBool:
		switch v := input.(type) {
		case string:
			if strings.ToLower(v) == "true" {
				hash, err = ethabi.EncodeTopic(t, true)
			} else {
				hash, err = ethabi.EncodeTopic(t, false)
			}
		}
	default:
		hash, err = ethabi.EncodeTopic(t, input)
	}
	if err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(hash.Bytes()), nil
}

func Send(abi *ethabi.ABI, methodName string, inputs interface{}, to common.Address, cli *ethclient.Client, opt *bind.TransactOpts) (*types.Transaction, error) {
	log.Debug("Send", "method", methodName, "inputs", inputs, "to", to.Hex())
	method := abi.GetMethod(methodName)
	if method == nil {
		return nil, fmt.Errorf("method not found:%s", methodName)
	}
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
		fmt.Println("send failed", err, signedTx.Gas(), signedTx.GasPrice())
		return nil, err
	}
	return signedTx, nil
}
func Call(abi *ethabi.ABI, methodName string, inputs interface{}, to common.Address, cli *ethclient.Client, opt *bind.CallOpts) (string, error) {
	log.Debug("Call", "method", methodName, "inputs", inputs, "to", to.Hex())

	method := abi.GetMethod(methodName)
	if method == nil {
		return "", fmt.Errorf("method not found:%s", methodName)
	}
	input, err := method.Encode(inputs)
	if err != nil {
		return "", fmt.Errorf("encode input failed:%s", err.Error())
	}
	result, err := cli.CallContract(context.Background(), platon.CallMsg{
		From: opt.From,
		To:   &to,
		Data: input,
	}, nil)
	if err != nil {
		return "", err
	}
	return formatReturn(result, method.Outputs)
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
	for i := 0; i < 60; i++ {
		receipt, _ := client.TransactionReceipt(context.Background(), hash)
		if receipt != nil {
			return receipt, nil
		}
		time.Sleep(time.Second)
	}
	return nil, fmt.Errorf("wait tx time out")
}

func formatReturn(input []byte, p *ethabi.Type) (string, error) {
	output, err := p.Decode(input)
	if err != nil {
		return "", err
	}
	var result string
	if val, ok := output.(map[string]interface{}); ok {
		var args []string
		for _, v := range val {
			args = append(args, fmt.Sprintf("%v", v))
		}
		result = strings.Join(args, ",")
	} else {
		result = fmt.Sprintf("%v", output)
	}

	return result, nil
}
