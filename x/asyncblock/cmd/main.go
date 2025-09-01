package main

import (
	"encoding/json"
	"flag"
	common2 "github.com/PlatONnetwork/AppChain-SDK/example/common"
	"github.com/PlatONnetwork/AppChain-SDK/store/memorydb"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/asyncblock"
	"github.com/PlatONnetwork/AppChain-SDK/x/benchmark"
	xconsensus "github.com/PlatONnetwork/AppChain-SDK/x/consensusnetwork"
	"github.com/PlatONnetwork/PlatON-Go/node"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	go func() {
		http.ListenAndServe(":6060", nil) // 可自定义端口
	}()
	//glogger := log.NewGlogHandler(log.StreamHandler(os.Stdin, log.TerminalFormat(false)))
	//glogger.Verbosity(log.LvlDebug)
	//log.Root().SetHandler(glogger)
	var apps []sdk.App
	var benchmarkModules []*benchmark.Module
	for i := 0; i < 4; i++ {
		asyncBlockModule := asyncblock.NewModule()
		asyncBlockModule.SetBlsKey(testutil.DefaultAccount[i].BlsKey)
		vals, _ := testutil.NewValidator(testutil.DefaultAccount[0:4])
		store := memorydb.New()
		benchmarkModule := benchmark.NewModule(cli.NewContext(cli.NewApp(), &flag.FlagSet{}, nil), store)
		consensusNetworkModule := xconsensus.NewModule(nil)
		manager := module.NewManager(vals, asyncBlockModule, consensusNetworkModule, benchmarkModule)
		manager.SetElection(vals.Name())
		manager.SetBlockExecutor(asyncBlockModule.Name())
		manager.SetTxFiller(asyncBlockModule.Name())
		manager.SetOrderTxPool(benchmarkModule.Name())
		manager.SetConsensusNetwork(consensusNetworkModule.Name())
		benchmarkModules = append(benchmarkModules, benchmarkModule)
		apps = append(apps, testutil.NewApp(manager))
	}
	var stack []*node.Node
	var err error
	if stack, _, err = testutil.CreateCluster(
		testutil.DefaultAccount[0:4],
		testutil.DefaultAccount[0:4],
		apps,
		map[string]json.RawMessage{
			"benchmark": json.RawMessage("{}"),
		},
		common2.UserAddrs); err != nil {
	}
	rpc := benchmark.NewRPC(benchmarkModules[0])
	go func() {
		rpc.GenTxs(0, 10, 100, 0, 20000)
		time.Sleep(time.Second * 10)
		rpc.Start(10, true)
	}()

	go stack[0].Start()
	go stack[1].Start()
	go stack[2].Start()
	go stack[3].Start()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
}
