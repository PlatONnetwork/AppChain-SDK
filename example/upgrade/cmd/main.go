package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	common2 "github.com/PlatONnetwork/AppChain-SDK/example/common"
	"github.com/PlatONnetwork/AppChain-SDK/example/election"
	"github.com/PlatONnetwork/AppChain-SDK/example/oracle"
	"github.com/PlatONnetwork/AppChain-SDK/example/oracle/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/AppChain-SDK/x/consensus"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	types2 "github.com/PlatONnetwork/PlatON-Go/core/types"
	"math/big"
	"math/rand"

	//exampleupgrade "github.com/PlatONnetwork/AppChain-SDK/example/upgrade"
	examplenode "github.com/PlatONnetwork/AppChain-SDK/example/upgrade/node"
	oracle2 "github.com/PlatONnetwork/AppChain-SDK/example/upgrade/oracle"
	"github.com/PlatONnetwork/AppChain-SDK/store/memorydb"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/AppChain-SDK/x/upgrade"
	"github.com/PlatONnetwork/AppChain-SDK/x/upgrade/types"
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
	GasPrice      = big.NewInt(11000000000)
	GasLimit      = uint64(1000000)
	ServerCommand = cli.Command{
		Name:   "server",
		Action: Server,
	}
	ClientCommand = cli.Command{
		Name: "client",
		Subcommands: []cli.Command{{
			Name:               "oracle-upgrade",
			Usage:              "upgrade oracle module",
			Flags:              []cli.Flag{upgradeBlockFlag},
			Action:             upgradeOracle,
			CustomHelpTemplate: flags.CommandHelpTemplate,
		}, {
			Name:               "query-rate",
			Usage:              "query rate",
			Flags:              []cli.Flag{},
			Action:             queryRate,
			CustomHelpTemplate: flags.CommandHelpTemplate,
		}, {
			Name:               "node-upgrade",
			Usage:              "query rate",
			Flags:              []cli.Flag{upgradeBlockFlag, upgradeVersionFlag},
			Action:             upgradeNode,
			CustomHelpTemplate: flags.CommandHelpTemplate,
		}, {
			Name:               "check-node",
			Usage:              "check node module",
			Flags:              []cli.Flag{nameFlag, hostFlag, portFlag},
			Action:             checkNode,
			CustomHelpTemplate: flags.CommandHelpTemplate,
		},
		},
	}
	upgradeBlockFlag = cli.Uint64Flag{
		Name:  "upgrade-block",
		Value: 10,
	}
	upgradeVersionFlag = cli.Uint64Flag{
		Name:  "upgrade-version",
		Value: 0,
	}
	nameFlag = cli.StringFlag{
		Name: "name",
	}
	hostFlag = cli.StringFlag{
		Name:  "host",
		Value: "127.0.0.1",
	}
	portFlag = cli.Uint64Flag{
		Name:  "port",
		Value: 6789,
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
	oracleModule := oracle2.NewModule(store, testutil.DefaultAccount[0].NodePrivateKey(), oracle.NewTimerRate(time.Second*2))
	extraVote := extravote.NewExtraVote(store, []extravote.ExtraVerifier{oracleModule})
	oracleModule.SetExtraVote(extraVote)
	vals := election.NewModule()
	network := consensus.NewModule(ctx)

	upgradeModule := upgrade.NewModule(store)
	nodeModule := examplenode.NewModule()
	manager := module.NewManager(vals, extraVote, upgradeModule, oracleModule, nodeModule, network)
	manager.SetElection(vals.Name())
	manager.SetOrderGenesis(vals.Name(), extraVote.Name(), upgradeModule.Name())
	manager.SetConsensusExtend(extraVote.Name())
	manager.SetConsensusNetwork(network.Name())
	manager.RegisterUpgradeHandler(upgradeModule)
	manager.SetModuleValidChecker(upgradeModule.IsModuleValid)
	upgradeModule.SetIsContractModule(manager.IsContractModule)
	oracleModule.RegistryUpgradeHandler(upgradeModule)
	nodeModule.RegistryUpgradeHandler(upgradeModule)
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
		EpochSize:        50,
		ElectionDistance: 20,
	}
	s, _ := json.Marshal(config)
	upgradeConfig, _ := json.Marshal(&types.GenesisConfig{
		ModuleGenesisConfig: module.ModuleGenesisConfig{},
		Owner:               common2.UserAddrs[9],
	})
	var stack []*node.Node
	var err error
	if stack, _, err = testutil.CreateCluster(testutil.DefaultAccount[0:1], testutil.DefaultAccount[0:1], []sdk.App{app}, map[string]json.RawMessage{
		vals.Name():          s,
		upgradeModule.Name(): upgradeConfig,
	}, common2.UserAddrs); err != nil {
		return err
	}

	stack[0].Start()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
	return nil
}

func upgradeOracle(ctx *cli.Context) error {
	cli, err := ethclient.Dial("http://127.0.0.1:8801")
	if err != nil {
		return err
	}
	height := ctx.Uint64(upgradeBlockFlag.Name)
	addr := common2.UserAddrs[9]
	key := common2.UserPrivateKeys[addr]
	nonce, err := cli.NonceAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}
	up, err := NewUpgrade(constants.UpgradeAddress, cli)
	if err != nil {
		return err
	}
	tx, err := up.AddUpgradePlan(createOpt(addr, key, nonce), IUpgradePlan{
		Name: "oracle-upgrade",
		Modules: []IUpgradeModule{
			{
				ModuleName: "oracle",
				Version:    0,
			},
		},
		Info:   "init genesis oracle module",
		Height: height,
		Status: 0,
	})
	if err != nil {
		return err
	}
	if err = waitTx(cli, tx); err != nil {
		return err
	}

	queryPlan(up, addr, height)
	return nil
}

func queryRate(ctx *cli.Context) error {
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
func upgradeNode(ctx *cli.Context) error {
	cli, err := ethclient.Dial("http://127.0.0.1:8801")
	if err != nil {
		return err
	}
	height := ctx.Uint64(upgradeBlockFlag.Name)
	version := ctx.Uint64(upgradeVersionFlag.Name)
	addr := common2.UserAddrs[9]
	key := common2.UserPrivateKeys[addr]
	nonce, err := cli.NonceAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}
	up, err := NewUpgrade(constants.UpgradeAddress, cli)
	if err != nil {
		return err
	}
	tx, err := up.AddUpgradePlan(createOpt(addr, key, nonce), IUpgradePlan{
		Name: "node-upgrade-" + randomString(6),
		Modules: []IUpgradeModule{
			{
				ModuleName: "node",
				Version:    version,
			},
		},
		Info:   fmt.Sprintf("upgrade node module, version:%d", version),
		Height: height,
		Status: 0,
	})
	if err != nil {
		return err
	}
	if err = waitTx(cli, tx); err != nil {
		return err
	}
	queryPlan(up, addr, height)
	return nil
}

func checkNode(ctx *cli.Context) (err error) {
	addr := common2.UserAddrs[9]
	key := common2.UserPrivateKeys[addr]
	cli, err := ethclient.Dial("http://127.0.0.1:8801")
	if err != nil {
		return err
	}
	nonce, err := cli.NonceAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}

	node, err := NewNode(examplenode.NodeAddress, cli)
	if err != nil {
		return err
	}
	name := ctx.String(nameFlag.Name)
	if len(name) == 0 {
		name = randomString(6)
	}

	//check addNode method
	fmt.Println("check addNode method")
	tx, err := node.AddNode(createOpt(addr, key, nonce), name, ctx.String(hostFlag.Name), uint16(ctx.Uint64(portFlag.Name)))
	if err != nil {
		return err
	}

	err = waitTx(cli, tx)
	if err != nil {
		return err
	}
	receipt, err := cli.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return nil
	}
	fmt.Println("add node success", "name:", name, "event", len(receipt.Logs) != 0)

	//check getNode method
	fmt.Println("check getNode method")
	info, err := node.GetNode(&bind.CallOpts{
		From: addr,
	}, name)
	if err != nil {
		return err
	}
	fmt.Println("get node:%s", name, "info:", info)

	//check delNode method
	fmt.Println("check delNode method")
	tx, err = node.DelNode(createOpt(addr, key, nonce+1), name)
	if err != nil {
		return err
	}

	err = waitTx(cli, tx)
	if err != nil {
		return err
	}
	info, err = node.GetNode(&bind.CallOpts{
		From: addr,
	}, name)
	if err != nil {
		return err
	}
	fmt.Println("get node:%s", name, "info:", info)
	return nil
}

func randomString(n int) string {

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)

}

func createOpt(addr common.Address, key *ecdsa.PrivateKey, nonce uint64) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:  addr,
		Nonce: new(big.Int).SetUint64(nonce),
		Signer: func(address common.Address, tx *types2.Transaction) (*types2.Transaction, error) {
			signer := types2.NewLondonSigner(big.NewInt(123083))
			signature, err := crypto.Sign(signer.Hash(tx, nil).Bytes(), key)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
		Value:     nil,
		GasPrice:  big.NewInt(11000000000),
		GasFeeCap: nil,
		GasTipCap: nil,
		GasLimit:  1000000,
		Context:   nil,
		NoSend:    true,
	}
}

func queryPlan(up *Upgrade, addr common.Address, height uint64) {
	plan, _ := up.GetUpgradePlan(&bind.CallOpts{From: addr}, height)

	for _, p := range plan {
		for _, m := range p.Modules {
			fmt.Println("set plan success", "name:", p.Name, "module:", m.ModuleName, "version:", m.Version, "info:", p.Info, "height:", p.Height)
		}
	}
}
func callTx(cli *ethclient.Client, tx *types2.Transaction) error {
	_, err := cli.CallContract(context.Background(), platon.CallMsg{
		From:     tx.FromAddr(types2.NewLondonSigner(tx.ChainId())),
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}, nil)
	return err
}
func waitTx(cli *ethclient.Client, tx *types2.Transaction) error {
	if err := callTx(cli, tx); err != nil {
		return err
	}
	if err := cli.SendTransaction(context.Background(), tx); err != nil {
		return err
	}

	fmt.Println("try call tx success, hash:", tx.Hash().Hex())
	var receipt *types2.Receipt
	var err error
	for i := 0; i < 5; i++ {
		receipt, err = cli.TransactionReceipt(context.Background(), tx.Hash())
		if receipt != nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if receipt == nil || err != nil {
		return errors.New("get receipt failed")
	}
	return nil
}

func require(err error) {
	if err != nil {
		panic(err)
	}
}
