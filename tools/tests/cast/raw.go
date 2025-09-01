package cast

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"gopkg.in/urfave/cli.v1"
	"math/big"
)

var (
	ToFlag = cli.StringFlag{
		Name:  "to",
		Usage: "Payee address",
	}
	ValueFlag = cli.StringFlag{
		Name:  "value",
		Usage: "Transfer value",
	}
	RawSendCommand = cli.Command{
		Name: "raw",
		Flags: []cli.Flag{
			ToFlag,
			ValueFlag,
			flags.AppchainRPCFlag,
			flags.GasLimitFlags,
			flags.GasPriceFlags,
			flags.KeyFlags,
		},
		Category:           "Raw transfer",
		Action:             rawSend,
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func rawSend(ctx *cli.Context) error {
	url := ctx.String(flags.AppchainRPCFlag.Name)
	client, err := ethclient.Dial(url)
	to := common.HexToAddress(ctx.String(ToFlag.Name))
	value, _ := new(big.Int).SetString(ctx.String(ValueFlag.Name), 10)
	fmt.Println("11111")
	key, err := crypto.HexToECDSA(ctx.String(flags.KeyFlags.Name))
	if err != nil {
		return err
	}
	fmt.Println("11111")
	addr := crypto.PubkeyToAddress(key.PublicKey)
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("11111")
	nonce, _ := client.PendingNonceAt(context.Background(), addr)
	gasPrice := big.NewInt(0).SetUint64(ctx.Uint64(flags.GasPriceFlags.Name))
	if gasPrice.Uint64() == 0 {
		gasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			return err
		}
	}
	fmt.Println("11111")
	gasLimit := ctx.Uint64(flags.GasLimitFlags.Name)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     nil,
	})
	signer := types.NewLondonSigner(chainId)
	tx, err = types.SignTx(tx, signer, key)
	if err != nil {
		return err
	}
	bytes, _ := json.MarshalIndent(tx, " ", " ")
	fmt.Println("transaction:", string(bytes))
	if err := client.SendTransaction(context.Background(), tx); err != nil {
		return err
	}

	receipt, err := WaitTx(client, tx.Hash())
	if err != nil {
		return err
	}
	bytes, _ = json.MarshalIndent(receipt, " ", " ")

	fmt.Println("receipt:", string(bytes))
	return nil
}
