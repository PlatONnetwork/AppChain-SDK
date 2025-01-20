package main

import (
	"context"
	"encoding/json"
	"fmt"
	common2 "github.com/PlatONnetwork/AppChain-SDK/example/common"
	"github.com/PlatONnetwork/AppChain-SDK/example/lottery"
	"github.com/PlatONnetwork/AppChain-SDK/example/lottery/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/node"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app := cli.NewApp()
	app.HideVersion = true // we have a command to print the version
	app.Commands = []cli.Command{
		ServerCommand,
		ClientCommand,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var (
	ServerCommand = cli.Command{
		Name:               "server",
		Action:             Server,
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
	ClientCommand = cli.Command{
		Name:               "client",
		Action:             Client,
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func Server(ctx *cli.Context) error {
	lotteryModule := lottery.NewModule(testutil.DefaultAccount[0].NodePrivateKey())
	vals, _ := testutil.NewValidator(testutil.DefaultAccount[0:1])
	vals.SetCoinbase(testutil.DefaultAccount[0].NodePrivateKey())
	manager := module.NewManager(vals, lotteryModule)
	manager.SetElection(vals.Name())
	manager.SetOrderGenesis(lotteryModule.Name())
	app := testutil.NewApp(manager)
	config := lottery.GenesisConfig{
		Name:   "Token",
		Symbol: "USDC",
		Nonce:  "0x03f3b376f00863de14440eff826835d16ffa3c8b0fc7ad1402beee7ccf076aa9282f795c3e92d2f61e45e85abe5fcfef134ae6700a50b0885942a92d92b9c88a280450a416880be13a23e449f41ef12f43",
	}
	s, _ := json.Marshal(config)
	var stack []*node.Node
	var err error
	if stack, _, err = testutil.CreateCluster(testutil.DefaultAccount[0:1], testutil.DefaultAccount[0:1], []sdk.App{app}, map[string]json.RawMessage{
		lotteryModule.Name(): s,
	}, common2.UserAddrs); err != nil {
		return err
	}

	stack[0].Start()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
	return nil
}

func Client(ctx *cli.Context) error {

	cli, err := ethclient.Dial("http://127.0.0.1:8801")
	if err != nil {
		return err
	}
	lc, err := contracts.NewContract(lottery.LotteryAddr, cli)

	for i := 0; i < 10; i++ {
		err := sendTx(cli, lc)
		if err != nil {
			return err
		}
	}

	return nil
}

func sendTx(cli *ethclient.Client, lc *contracts.Contract) error {

	var nonces []uint64
	for _, addr := range common2.UserAddrs {
		nonce, err := cli.NonceAt(context.Background(), addr, nil)
		if err != nil {
			return err
		}
		nonces = append(nonces, nonce)
	}
	var txs types.Transactions
	for i, addr := range common2.UserAddrs {
		tx, err := lc.Guessing(&bind.TransactOpts{
			From:  addr,
			Nonce: new(big.Int).SetUint64(nonces[i]),
			Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
				signer := types.NewLondonSigner(big.NewInt(123083))
				signature, err := crypto.Sign(signer.Hash(tx, nil).Bytes(), common2.UserPrivateKeys[addr])
				if err != nil {
					return nil, err
				}
				return tx.WithSignature(signer, signature)
			},
			Value:     nil,
			GasPrice:  big.NewInt(11000000000),
			GasFeeCap: nil,
			GasTipCap: nil,
			GasLimit:  0,
			Context:   nil,
			NoSend:    false,
		}, big.NewInt(int64(i)))
		if err != nil {
			return err
		}
		//fmt.Println("tx:", tx.Hash().String())
		txs = append(txs, tx)
	}
	time.Sleep(time.Second * 3)
	blocks := make(map[uint64]struct{})
	for _, tx := range txs {
		receipt, _ := cli.TransactionReceipt(context.Background(), tx.Hash())
		for _, log := range receipt.Logs {
			event, err := lc.ParseGuessing(*log)
			if err == nil {
				fmt.Println("block:", receipt.BlockNumber, "owner:", event.Owner.Hex(), "guess", event.Number)
				blocks[receipt.BlockNumber.Uint64()] = struct{}{}
			}
		}
	}
	for blockNumber, _ := range blocks {
		block, err := cli.BlockByNumber(context.Background(), new(big.Int).SetUint64(blockNumber+1))
		if err != nil {
			panic(err)
		}
		for _, tx := range block.Transactions() {
			if *tx.To() == lottery.LotteryAddr {
				receipt, _ := cli.TransactionReceipt(context.Background(), tx.Hash())
				for _, log := range receipt.Logs {
					event, err := lc.ParseDrawing(*log)
					if err == nil {
						fmt.Println("drawing", "block:", receipt.BlockNumber,
							"address:", event.Owner.Hex(), "number:", event.Number, "nonce:", event.Nonce, "bonus", event.Bonus)
					}
				}
			}
		}
	}
	return nil
}
