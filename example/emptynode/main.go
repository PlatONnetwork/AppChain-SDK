package main

import (
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/example/blacklist"
	common2 "github.com/PlatONnetwork/AppChain-SDK/example/common"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/consensus"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/node"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	PprofFlag = cli.IntFlag{
		Name:  "pprofport",
		Value: 7801,
	}
	P2pFlag = cli.IntFlag{
		Name:  "p2pport",
		Value: 19001,
	}
	HttpFlag = cli.IntFlag{
		Name:  "httpport",
		Value: 8901,
	}
	AllocAddrsFlag = cli.StringFlag{
		Name:  "alloc",
		Usage: "Alloc account, separated by commas",
		Value: "",
	}
	ServerCommand = cli.Command{
		Name: "server",
		Flags: []cli.Flag{
			PprofFlag,
			P2pFlag,
			HttpFlag,
			AllocAddrsFlag,
		},
		Description: `Default genesis address:{" +
			"0x8a0f3F8389F79a05Dd03027dE0Bcbb06cADB3F9c":"5856fb978afcb48fd023a0e6ba76eb90ed35de98ca4a6873d28d1c5b9a47f39e",
			"0x1FD13cA28ccc423e0565fD3C9b64187ad91Ae7D6":"fa4ca9b1cc136fbaab4ab1bf55a3718a5d9ff2cc0b6ffbb098682e5d257b416b",
			"0x27E1a2e05186c9e4C4A3A957F8748fA35EA19076":"94c67767ede78befe7352bd14a8338a124147db0f01c1a011401340b94832af1",
			"0x6151aE08d3acD4c8194bB029633Ac7059066C75F":"bc22d5de77e9b50f0635158edd282cfe6a5c944f3dee6b4d68a38dcfd5e68784",
			"0xE9565da74A5149e43475e62646dC4f47c2FcEa6f":"f0014181639f603f938647aa7d38663c3e867b2c25c1da7784d38752deaca67e",
			"0xd1391ba0Be0eECA4031e952402b92e26417D5F38":"6eb85bfd34859d81db7f3d0a0c3f72c662eb1ea92de435dffc41497bbb0fbbcc",
			"0xEdA947963D73F7Ca96c94843E2bd091d31Ec9F90":"0d1d0a522c1765b2fadbe86ba2bf45b3ab9a9122f448c5f59bcee76dfe3790e1",
			"0xD03ae6Da0708D073f3f920Ea1AAEbfdc847971b3":"6fc120b574690e6019a0c91fef94ad6fb0c196264fb986497ce4b46e662c2251",
			"0xD9bE8f83736b508514183363a1fec4c0Bd1E80Db":"b3b71d3e1c48e20b9b6d548a20a5ba4db5a7a3666d2e35ec911538be4dfd006d",
			"0x0504d04808022aC69bBb37206FD97Edaef9268e9":"9eb4fda04c360a5b3ea046f2317006bb2f05802215ac75ae316b9dd7d03e2c15",
`,
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
	blacklistModule := blacklist.Module{}
	account := testutil.DefaultAccount[0:1]

	vals, _ := testutil.NewValidator(account)
	network := consensus.NewModule(ctx)
	manager := module.NewManager(vals, &blacklistModule, network)
	manager.SetElection(vals.Name())
	manager.SetOrderTxPool(blacklistModule.Name())
	manager.SetConsensusNetwork(network.Name())
	app := testutil.NewApp(manager)

	var stack []*node.Node
	var err error
	account[0].Pprof = ctx.Int(PprofFlag.Name)
	account[0].P2PPort = ctx.Int(P2pFlag.Name)
	account[0].HTTP = ctx.Int(HttpFlag.Name)
	addrs := append(common2.UserAddrs, splitAddress(ctx.String(AllocAddrsFlag.Name))...)
	if stack, _, err = testutil.CreateCluster(
		account,
		account,
		[]sdk.App{app},
		map[string]json.RawMessage{},
		addrs); err != nil {
		return err
	}
	types.HttpEthCompatible = true

	stack[0].Start()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
	return nil
}

func splitAddress(s string) []common.Address {
	addrs := strings.Split(s, ",")
	var res []common.Address
	for _, a := range addrs {
		res = append(res, common.HexToAddress(a))
	}
	return res
}
