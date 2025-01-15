package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	common2 "github.com/PlatONnetwork/AppChain-SDK/example/common"
	"github.com/PlatONnetwork/AppChain-SDK/example/election"
	"github.com/PlatONnetwork/AppChain-SDK/example/election/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/eth"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/node"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
	"math/big"
	"os"
	"os/signal"
	"syscall"
)

var (
	adminFlag = cli.StringFlag{
		Name:  "admin",
		Value: common2.UserAddrs[0].Hex(),
	}
	genesisFlag = cli.BoolFlag{
		Name: "genesis",
	}

	nodeFlag = cli.IntFlag{
		Name:  "node",
		Value: 1,
	}

	adminKeyFlag = cli.StringFlag{
		Name:        "admin-key",
		Value:       hex.EncodeToString(crypto.FromECDSA(common2.UserPrivateKeys[common2.UserAddrs[0]])),
		Destination: nil,
	}

	ServerCommand = cli.Command{
		Name:   "server",
		Action: Server,
		Flags: []cli.Flag{
			adminFlag,
			genesisFlag,
			nodeFlag,
		},
	}
	ClientCommand = cli.Command{
		Name:   "client",
		Action: Client,
		Flags: []cli.Flag{
			adminKeyFlag,
			nodeFlag,
		},
	}
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

func Server(ctx *cli.Context) error {
	vals := election.NewModule()
	manager := module.NewManager(vals)
	manager.SetElection(vals.Name())
	manager.SetOrderGenesis(vals.Name())
	app := testutil.NewApp(manager)
	nodeNumber := ctx.Int(nodeFlag.Name) - 1
	genesis := ctx.Bool(genesisFlag.Name)
	if genesis {
		nodeNumber = 0
	}
	testutil.InitLog(log.LvlDebug)
	config := election.GenesisConfig{
		InitialNodes: election.Nodes{
			election.Node{
				Name:        "node1",
				Owner:       testutil.DefaultAccount[0].NodeAddress(),
				Desc:        "node1",
				PublicKey:   crypto.FromECDSAPub(&testutil.DefaultAccount[0].NodePrivateKey().PublicKey),
				BlsPubKey:   testutil.DefaultAccount[0].BlsSecretKey().GetPublicKey().Serialize(),
				HostAddress: "127.0.0.1",
				RpcPort:     uint16(testutil.DefaultAccount[0].HTTP),
				P2pPort:     uint16(testutil.DefaultAccount[0].P2PPort),
			},
		},
		AdminAddress:     common.HexToAddress(ctx.String(adminFlag.Name)),
		EpochSize:        10,
		ElectionDistance: 5,
	}
	s, _ := json.Marshal(config)
	var stack []*node.Node
	var backend []*eth.Ethereum
	var err error
	go testutil.StartPProf(fmt.Sprintf("10.2.11.24:%d", testutil.DefaultAccount[nodeNumber].Pprof))
	if stack, backend, err = testutil.CreateCluster([]*testutil.Account{testutil.DefaultAccount[nodeNumber]}, testutil.DefaultAccount[0:1], []sdk.App{app}, map[string]json.RawMessage{
		vals.Name(): s,
	}, common2.UserAddrs); err != nil {
		return err
	}

	stack[0].Start()
	backend[0].Start()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
	return nil
}

func Client(ctx *cli.Context) error {
	nodeNumber := ctx.Int(nodeFlag.Name) - 1
	if nodeNumber < 2 || nodeNumber > 3 {
		return errors.New("wrong node number")
	}
	adminKey, err := crypto.HexToECDSA(ctx.String(adminKeyFlag.Name))
	if err != nil {
		return err
	}
	cli, err := ethclient.Dial("http://127.0.0.1:8801")
	if err != nil {
		return err
	}
	election, err := contracts.NewElection(election.ProxyAddress, cli)
	if err != nil {
		return err
	}
	tx, err := election.AddNode(&bind.TransactOpts{
		From: crypto.PubkeyToAddress(adminKey.PublicKey),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signer := types.NewLondonSigner(big.NewInt(123083))
			signature, err := crypto.Sign(signer.Hash(tx, nil).Bytes(), adminKey)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}, contracts.Node{
		Name:        fmt.Sprintf("node%d", nodeNumber),
		Owner:       testutil.DefaultAccount[nodeNumber].NodeAddress(),
		Desc:        fmt.Sprintf("node%d", nodeNumber),
		PublicKey:   crypto.FromECDSAPub(&testutil.DefaultAccount[nodeNumber].NodePrivateKey().PublicKey),
		BlsPubKey:   testutil.DefaultAccount[nodeNumber].BlsSecretKey().GetPublicKey().Serialize(),
		HostAddress: "127.0.0.1",
		RpcPort:     uint16(testutil.DefaultAccount[nodeNumber].HTTP),
		P2pPort:     uint16(testutil.DefaultAccount[nodeNumber].HTTP),
	})
	fmt.Println("tx:", tx.Hash().Hex())
	return common2.WaitTx(cli, tx)
}
