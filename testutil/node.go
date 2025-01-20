package testutil

import (
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	"github.com/PlatONnetwork/PlatON-Go/core"
	"github.com/PlatONnetwork/PlatON-Go/eth"
	"github.com/PlatONnetwork/PlatON-Go/eth/downloader"
	"github.com/PlatONnetwork/PlatON-Go/eth/ethconfig"
	"github.com/PlatONnetwork/PlatON-Go/eth/tracers"
	"github.com/PlatONnetwork/PlatON-Go/ethdb"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/miner"
	"github.com/PlatONnetwork/PlatON-Go/node"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func StartPProf(address string) {
	log.Info("Starting pprof server", "addr", fmt.Sprintf("http://%s/debug/pprof", address))
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Error("Failure in running pprof server", "err", err)
	}
}
func InitLog(level log.Lvl) {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stdout, log.TerminalFormat(false)))
	glogger.Verbosity(level)
	log.Root().SetHandler(glogger)
	log.PrintOrigins(true)
}
func CreateCluster(accounts []*Account, genesisAccount []*Account, apps []sdk.App, genesisModule map[string]json.RawMessage, allocs []common.Address) ([]*node.Node, []*eth.Ethereum, error) {
	var nodes []*node.Node
	var backends []*eth.Ethereum
	genesisJson, _ := GenerateGenesis(genesisAccount, genesisModule, allocs)
	for i, acc := range accounts {
		n, b, err := NewMemoryNode(acc, string(genesisJson), apps[i])
		if err != nil {
			return nil, nil, err
		}
		nodes = append(nodes, n)
		backends = append(backends, b)
	}
	return nodes, backends, nil
}

func NewMemoryNode(account *Account, genesisJson string, app sdk.App) (*node.Node, *eth.Ethereum, error) {
	nodePriKey, blsPriKey := account.NodePrivateKey(), account.BlsSecretKey()
	p2pConfig := node.DefaultConfig.P2P
	p2pConfig.ListenAddr = fmt.Sprintf("%s:%d", account.Host, account.P2PPort)
	p2pConfig.PrivateKey = nodePriKey
	p2pConfig.BlsPublicKey = *blsPriKey.GetPublicKey()
	stack, err := node.New(&node.Config{Name: "root", DataDir: "", P2P: p2pConfig, HTTPHost: "0.0.0.0", HTTPPort: account.HTTP, HTTPModules: []string{"platon", "debug", "personal", "admin", "net", "web3", "txpool"}})
	chaindbs, err := InitGenesis(stack, genesisJson, app)
	if err != nil {
		return nil, nil, err
	}
	ethConfig := getEthConfig(account)
	if err != nil {
		return nil, nil, err
	}
	backend, err := eth.NewMem(stack, chaindbs[0], ethConfig, app)
	if err != nil {
		return nil, nil, err
	}
	stack.RegisterAPIs(tracers.APIs(backend.APIBackend))
	return stack, backend, nil
}

func InitGenesis(stack *node.Node, genesisJson string, app sdk.App) ([]ethdb.Database, error) {
	genesis := new(core.Genesis)
	genesis.SetApp(app)
	if err := genesis.InitGenesisConfigFromJson(genesisJson); err != nil {
		return nil, err
	}
	var chaindbs []ethdb.Database
	for _, name := range []string{"chaindata", "lightchaindata"} {
		ns := fmt.Sprintf("%s/%s/", "mem", name)
		chaindb, err := stack.OpenDatabase(ns, 0, 0, "", false)
		if err != nil {
			utils.Fatalf("Failed to open database: %v", err)
		}

		_, hash, err := core.SetupGenesisBlock(chaindb, genesis)
		if err != nil {
			utils.Fatalf("Failed to write genesis block: %v", err)
		}
		log.Info("Successfully wrote genesis state", "database", ns, "hash", hash.Hex())
		chaindbs = append(chaindbs, chaindb)
	}
	return chaindbs, nil
}
func getEthConfig(acc *Account) *ethconfig.Config {
	nodePriKey, nod, blsPriKey := acc.NodePrivateKey(), acc.Node(), acc.BlsSecretKey()

	return &ethconfig.Config{
		CbftConfig: types.OptionsConfig{
			NodePriKey:        nodePriKey,
			NodeID:            nod.IDv0(),
			Node:              nod,
			BlsPriKey:         blsPriKey,
			WalMode:           false,
			PeerMsgQueueSize:  1024,
			EvidenceDir:       "",
			MaxPingLatency:    5000,
			MaxQueuesLimit:    4096,
			BlacklistDeadline: 60,
			Period:            400,
			Amount:            10,
		},
		SyncMode:                downloader.FullSync,
		DatabaseCache:           768,
		TrieCache:               32,
		TrieTimeout:             60 * time.Minute,
		SnapshotCache:           256,
		TrieDBCache:             512,
		DBDisabledGC:            true,
		DBGCInterval:            86400,
		DBGCTimeout:             time.Minute,
		DBGCMpt:                 false,
		DBGCBlock:               256,
		VMWasmType:              "wagon",
		VmTimeoutDuration:       0, // default 0 ms for vm exec timeout
		TrieCleanCache:          154,
		TrieCleanCacheJournal:   "triecache",
		TrieCleanCacheRejournal: 60 * time.Minute,
		TrieDirtyCache:          256,
		Miner: miner.Config{
			GasFloor: params.GenesisGasLimit,
			GasPrice: big.NewInt(params.GVon),
			Recommit: 3 * time.Second,
		},

		MiningLogAtDepth:       7,
		TxChanSize:             4096,
		ChainHeadChanSize:      10,
		ChainSideChanSize:      10,
		ResultQueueSize:        10,
		ResubmitAdjustChanSize: 10,
		MinRecommitInterval:    1 * time.Second,
		MaxRecommitInterval:    15 * time.Second,
		IntervalAdjustRatio:    0.1,
		IntervalAdjustBias:     200 * 1000.0 * 1000.0,
		StaleThreshold:         7,
		DefaultCommitRatio:     0.95,

		BodyCacheLimit:    256,
		BlockCacheLimit:   256,
		MaxFutureBlocks:   256,
		TriesInMemory:     128,
		BlockChainVersion: 3,

		TxPool:        core.DefaultTxPoolConfig,
		RPCGasCap:     50000000,
		RPCEVMTimeout: 5 * time.Second,
		RPCTxFeeCap:   1, // 1 lat
	}
}
