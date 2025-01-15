package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/example/blacklist"
	common2 "github.com/PlatONnetwork/AppChain-SDK/example/common"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
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
			blacklistFlag,
		},
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
	blacklistFlag = cli.StringFlag{
		Name:  "blacklist",
		Value: "0x8a0f3F8389F79a05Dd03027dE0Bcbb06cADB3F9c",
	}
)

func Server(ctx *cli.Context) error {
	blacklistModule := blacklist.Module{}

	vals, _ := testutil.NewValidator(testutil.DefaultAccount[0:1])

	manager := module.NewManager(vals, &blacklistModule)
	manager.SetElection(vals.Name())
	manager.SetOrderTxPool(blacklistModule.Name())
	app := testutil.NewApp(manager)

	var stack []*node.Node
	var backend []*eth.Ethereum
	var err error
	if stack, backend, err = testutil.CreateCluster(
		testutil.DefaultAccount[0:1],
		testutil.DefaultAccount[0:1],
		[]sdk.App{app},
		map[string]json.RawMessage{},
		common2.UserAddrs); err != nil {
		return err
	}

	stack[0].Config().HTTPModules = append(stack[0].Config().HTTPModules, "blacklist")

	stack[0].Start()
	backend[0].Start()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
	return nil
}

func Client(ctx *cli.Context) error {

	cli, err := blacklist.NewClient("http://127.0.0.1:8801")
	if err != nil {
		return err
	}
	addrs := decodeAddr(ctx)
	fmt.Println("addrs:", len(addrs))

	err = cli.SetBlacklist(context.Background(), addrs)
	if err != nil {
		return err
	}
	addrs, err = cli.GetBlacklist(context.Background())
	if err != nil {
		return err
	}

	var nonces []uint64
	for _, addr := range common2.UserAddrs {
		nonce, err := cli.NonceAt(context.Background(), addr, nil)
		if err != nil {
			return err
		}
		nonces = append(nonces, nonce)
	}
	chainId, err := cli.ChainID(context.Background())
	if err != nil {
		return err
	}
	for i, addr := range common2.UserAddrs {
		tx := types.NewTransaction(nonces[i], common.BigToAddress(big.NewInt(1111)), big.NewInt(1000), 100000, big.NewInt(11000000000), nil)
		signer := types.NewLondonSigner(chainId)
		signature, err := crypto.Sign(signer.Hash(tx, nil).Bytes(), common2.UserPrivateKeys[addr])
		if err != nil {
			return err
		}
		tx, _ = tx.WithSignature(signer, signature)
		fmt.Println("tx:", tx.Hash().String())
		err = cli.SendTransaction(context.Background(), tx)
		if err != nil {
			fmt.Println("send failed, from", addr.Hex(), "err:", err.Error())
		}
		time.Sleep(time.Second * 2)
	}

	return nil
}
func decodeAddr(ctx *cli.Context) []common.Address {
	var addrs []common.Address
	as := strings.Split(ctx.String(blacklistFlag.Name), ",")
	for _, s := range as {
		addrs = append(addrs, common.HexToAddress(s))
	}
	return addrs
}
