package main

import (
	"encoding/json"
	"fmt"
	common2 "github.com/PlatONnetwork/AppChain-SDK/example/common"
	"github.com/PlatONnetwork/AppChain-SDK/example/election"
	"github.com/PlatONnetwork/AppChain-SDK/example/oracle"
	"github.com/PlatONnetwork/AppChain-SDK/example/oracle/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/store/memorydb"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/node"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	ServerCommand = cli.Command{
		Name:   "server",
		Action: Server,
	}
	ClientCommand = cli.Command{
		Name:   "client",
		Action: Client,
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
	store := memorydb.New()
	oracleModule := oracle.NewModule(store, testutil.DefaultAccount[0].NodePrivateKey(), oracle.NewTimerRate(time.Second*2))
	extraVote := extravote.NewExtraVote(store, []extravote.ExtraVerifier{oracleModule})
	extraVote.AddEnableVerifiers(oracleModule.Name())
	vals := election.NewModule()
	manager := module.NewManager(oracleModule, vals, extraVote)
	manager.SetElection(vals.Name())
	manager.SetOrderGenesis(vals.Name(), extraVote.Name(), oracleModule.Name())
	manager.SetConsensusExtend(extraVote.Name())
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
		EpochSize:        250,
		ElectionDistance: 20,
	}
	s, _ := json.Marshal(config)
	og, _ := json.Marshal(oracle.GenesisConfig{Decimals: 3, BlockNumber: 0})
	var stack []*node.Node
	var err error
	if stack, _, err = testutil.CreateCluster(testutil.DefaultAccount[0:1], testutil.DefaultAccount[0:1], []sdk.App{app}, map[string]json.RawMessage{
		vals.Name():         s,
		oracleModule.Name(): og,
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
	lc, err := contracts.NewRate(oracle.RateAddr, cli)

	for i := 0; i < 10; i++ {
		blockNumber, err := lc.BlockNumber(nil)
		if err != nil {
			return err
		}
		rate, err := lc.Rate(nil)
		if err != nil {
			return err
		}
		fmt.Println("blockNumber:", blockNumber, "rate:", rate)
		time.Sleep(5 * time.Second)
	}

	return nil
}
