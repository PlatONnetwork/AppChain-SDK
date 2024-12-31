package main

import (
	"encoding/json"
	"fmt"
	common2 "github.com/PlatONnetwork/AppChain-SDK/example/common"
	"github.com/PlatONnetwork/AppChain-SDK/example/election"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/eth"
	"github.com/PlatONnetwork/PlatON-Go/node"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
	"os"
	"os/signal"
	"syscall"
)

var (
	ServerCommand = cli.Command{
		Name:   "server",
		Action: Server,
	}
)

func main() {
	app := cli.NewApp()
	app.HideVersion = true // we have a command to print the version
	app.Commands = []cli.Command{
		ServerCommand,
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
		AdminAddress:     common2.UserAddrs[0],
		ElectionDistance: 20,
	}
	s, _ := json.Marshal(config)
	var stack []*node.Node
	var backend []*eth.Ethereum
	var err error
	if stack, backend, err = testutil.CreateCluster(testutil.DefaultAccount[0:1], []sdk.App{app}, map[string]json.RawMessage{
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
