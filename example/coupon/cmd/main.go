package main

import (
	"context"
	"encoding/json"
	"fmt"
	common2 "github.com/PlatONnetwork/AppChain-SDK/example/common"
	"github.com/PlatONnetwork/AppChain-SDK/example/coupon"
	"github.com/PlatONnetwork/AppChain-SDK/example/coupon/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/eth"
	"github.com/PlatONnetwork/PlatON-Go/node"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
	"math/big"
	"os"
	"os/signal"
	"strings"
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
		Name:   "server",
		Action: Server,
	}
	ClientCommand = cli.Command{
		Name:   "client",
		Action: Client,
		Flags: []cli.Flag{
			priorityFlag,
		},
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
	priorityFlag = cli.StringFlag{
		Name:  "priority",
		Value: "0x0504d04808022aC69bBb37206FD97Edaef9268e9",
	}
)

func Server(ctx *cli.Context) error {
	couponModule := coupon.Module{}

	vals, _ := testutil.NewValidator(testutil.DefaultAccount[0:1])

	manager := module.NewManager(vals, &couponModule)
	manager.SetElection(vals.Name())
	manager.SetOrderGenesis(couponModule.Name())
	manager.SetWorker(couponModule.Name())
	app := testutil.NewApp(manager)
	config := coupon.GenesisConfig{
		Name:   "Token",
		Symbol: "USDC",
	}
	s, _ := json.Marshal(config)
	var stack []*node.Node
	var backend []*eth.Ethereum
	var err error
	if stack, backend, err = testutil.CreateCluster(testutil.DefaultAccount[0:1], []sdk.App{app}, map[string]json.RawMessage{
		couponModule.Name(): s,
	}, common2.UserAddrs); err != nil {
		return err
	}

	stack[0].Config().HTTPModules = append(stack[0].Config().HTTPModules, "coupon")

	stack[0].Start()
	backend[0].Start()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
	return nil
}

func Client(ctx *cli.Context) error {

	cli, err := coupon.NewClient("http://127.0.0.1:8801")
	if err != nil {
		return err
	}
	addrs := decodeAddr(ctx)
	err = cli.SetPriority(context.Background(), addrs)
	if err != nil {
		return err
	}
	addrs, err = cli.GetPriority(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("addrs:", len(addrs))

	coupon, err := contracts.NewContracts(coupon.CouponAddress, cli)
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
		tx, err := coupon.ApplyCoupon(&bind.TransactOpts{
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
		})
		if err != nil {
			return err
		}
		fmt.Println("tx:", tx.Hash().String())
		txs = append(txs, tx)
	}
	time.Sleep(time.Second * 5)
	for _, tx := range txs {
		receipt, _ := cli.TransactionReceipt(context.Background(), tx.Hash())
		for _, log := range receipt.Logs {
			event, err := coupon.ParseApply(*log)
			if err == nil {
				fmt.Println("block:", receipt.BlockNumber, "index:", receipt.TransactionIndex, "address:", event.To.Hex(), "value:", event.Value, "balance:", event.Balance)
				break
			}
		}
	}
	return nil
}
func decodeAddr(ctx *cli.Context) []common.Address {
	var addrs []common.Address
	as := strings.Split(ctx.String(priorityFlag.Name), ",")
	for _, s := range as {
		addrs = append(addrs, common.HexToAddress(s))
	}
	return addrs
}
