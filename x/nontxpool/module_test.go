package nontxpool

import (
	"crypto/ecdsa"
	"encoding/json"
	"flag"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	xconsensus "github.com/PlatONnetwork/AppChain-SDK/x/consensusnetwork"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	crypto "github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/node"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"

	"math/big"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

type MockModule struct {
	txpool   sdk.TxPool
	sortTxs  func(ctx sdk.WorkerContext, local map[common.Address]types.Transactions, remote map[common.Address]types.Transactions) (types.Transactions, error)
	onCommit func(ctx sdk.ConsensusContext, block *types.Block) error
}

func (m *MockModule) Name() string {
	return "mockmodule"
}

func (m *MockModule) Version() uint64 {
	return 1
}

func (m *MockModule) Init(ctx sdk.InitContext) error {
	m.txpool = ctx.Backend().TxPool()
	return nil
}

func (m *MockModule) SortTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions, remote map[common.Address]types.Transactions) (types.Transactions, error) {
	return m.sortTxs(ctx, local, remote)
}

func (m *MockModule) OnCommit(ctx sdk.ConsensusContext, block *types.Block) error {
	return m.onCommit(ctx, block)
}

var (
	gasPrice = big.NewInt(1000000000)
)

func createRawTransfer(signer types.Signer, from *ecdsa.PrivateKey, to common.Address, nonce uint64) (*types.Transaction, error) {
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Value:    big.NewInt(1),
		Gas:      30000,
		GasPrice: gasPrice,
		Data:     nil,
	})
	return types.SignTx(tx, signer, from)
}

func TestTxPoolPending(t *testing.T) {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.LvlDebug)
	log.Root().SetHandler(glogger)
	var apps []sdk.App
	var mockModules []*MockModule
	flags := &flag.FlagSet{}
	cliApp := cli.NewApp()
	AddModuleInitFlags(cliApp)
	for i := 0; i < 1; i++ {
		vals, _ := testutil.NewValidator(testutil.DefaultAccount[0:1])
		consensusNetworkModule := xconsensus.NewModule(nil)
		nontxpoolModule := NewModule(cli.NewContext(cliApp, flags, nil))
		mockModule := &MockModule{}
		mockModules = append(mockModules, mockModule)
		manager := module.NewManager(vals, consensusNetworkModule, nontxpoolModule, mockModule)
		manager.SetWorker(nontxpoolModule.Name())
		manager.SetElection(vals.Name())
		manager.SetConsensusNetwork(consensusNetworkModule.Name())
		apps = append(apps, testutil.NewApp(manager))
	}
	var stack []*node.Node
	var err error
	sk, _ := crypto.GenerateKey()
	if stack, _, err = testutil.CreateCluster(
		testutil.DefaultAccount[0:1],
		testutil.DefaultAccount[0:1],
		apps,
		map[string]json.RawMessage{
			"nontxpool":  json.RawMessage("{}"),
			"mockmodule": json.RawMessage("{}"),
		},
		[]common.Address{crypto.PubkeyToAddress(sk.PublicKey)}); err != nil {
		t.Fatal(err)
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSTOP)
	signer := types.NewLondonSigner(testutil.DefaultGenesisTemplate.Config.ChainID)
	tx, _ := createRawTransfer(signer, sk, common.Address{}, 0)

	mockModules[0].sortTxs = func(ctx sdk.WorkerContext, local map[common.Address]types.Transactions, remote map[common.Address]types.Transactions) (types.Transactions, error) {
		return nil, nil
	}
	mockModules[0].onCommit = func(ctx sdk.ConsensusContext, block *types.Block) error {
		if block.NumberU64() == 10 {
			go mockModules[0].txpool.AddLocal(tx)
		} else {
			if block.Transaction(tx.Hash()) != nil {
				stack[0].Close()
				sigc <- syscall.SIGSTOP
			}
		}
		return nil
	}

	go func() {
		stack[0].Start()
	}()
	<-sigc
}
