package cast

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/log"
	ethabi "github.com/umbracle/ethgo/abi"
	"gopkg.in/urfave/cli.v1"
	"math/big"
)

var (
	exitIDFlag = cli.Uint64Flag{
		Name:  "exit-id",
		Usage: "Exit event identifier",
	}
	ExitHelperCommand = cli.Command{
		Name: "l1.exithelper",
		Flags: []cli.Flag{
			flags.RootchainRPCFlags,
			flags.AppchainRPCFlag,
			exitIDFlag,
			flags.KeyFlags,
			flags.AddressFlags,
			flags.GasLimitFlags,
			flags.GasPriceFlags,
		},
		Category:           "L1 EXIT HELPER",
		Action:             exitHelper,
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func exitHelper(ctx *cli.Context) error {
	appCli, rootCli, addr, opt, err := initEnv(ctx)
	if err != nil {
		log.Error("Init Env failed", "err", err)
		return err
	}
	cpCli := checkpoint.NewClient(appCli)
	exitID, ok := new(big.Int).SetString(ctx.String(exitIDFlag.Name), 10)
	if !ok {
		return fmt.Errorf("invalid exit event identifier")
	}

	proof, err := cpCli.GenerateExitProof(context.Background(), exitID.Uint64())
	if err != nil {
		return fmt.Errorf("Generate exit proof error: %v\n", err)
	}

	blockNumber := new(big.Int).SetUint64(proof.Metadata.CheckpointBlock)
	leafIndex := new(big.Int).SetUint64(proof.Metadata.LeafIndex)

	unhashedLeaf := common.Hex2Bytes(proof.Metadata.ExitEvent)
	log.Debug("inputs", "number", blockNumber, "leaf", hex.EncodeToString(unhashedLeaf), "leafindex", leafIndex)
	var proofs [][32]byte
	for _, d := range proof.Data {
		proofs = append(proofs, d)
		log.Debug("inputs", "proof", hex.EncodeToString(d[:]))

	}
	abi, _ := ethabi.NewABI(Modules["l1.exithelper"])

	tx, err := Send(abi, "exit", []interface{}{
		blockNumber, leafIndex, unhashedLeaf, proofs,
	}, addr, rootCli, opt)
	if err != nil {
		return err
	}

	fmt.Println("send success", tx.Hash())
	tx, _, err = rootCli.TransactionByHash(context.Background(), tx.Hash())
	if err != nil {
		return err
	}
	bytes, _ := json.MarshalIndent(tx, " ", " ")
	fmt.Println("transaction:", string(bytes))
	receipt, err := rootCli.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return err
	}
	bytes, _ = json.MarshalIndent(receipt, " ", " ")

	fmt.Println("receipt:", string(bytes))
	return nil
}

func initEnv(ctx *cli.Context) (*ethclient.Client, *ethclient.Client, common.Address, *bind.TransactOpts, error) {
	url := ctx.String(flags.AppchainRPCFlag.Name)
	appCli, err := ethclient.Dial(url)
	if err != nil {
		return nil, nil, common.Address{}, nil, err
	}
	rootUrl := ctx.String(flags.RootchainRPCFlags.Name)
	rootCli, err := ethclient.Dial(rootUrl)
	if err != nil {
		return nil, nil, common.Address{}, nil, err
	}
	address := common.HexToAddress(ctx.String(flags.AddressFlags.Name))
	var opt *bind.TransactOpts
	key, err := crypto.HexToECDSA(ctx.String(flags.KeyFlags.Name))
	if err != nil {
		return nil, nil, common.Address{}, nil, err
	}
	chainId, err := rootCli.ChainID(context.Background())
	if err != nil {
		return nil, nil, common.Address{}, nil, err
	}
	opt, err = bind.NewKeyedTransactorWithChainID(key, chainId)
	log.Debug("Chainid", "value", chainId)
	if err != nil {
		return nil, nil, common.Address{}, nil, err
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

	log.Debug("Init args", "contractAddress", address.Hex(), "url", url)
	return appCli, rootCli, address, opt, nil
}
