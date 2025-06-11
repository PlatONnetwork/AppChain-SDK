package main

import (
	"crypto/ecdsa"
	"flag"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/PlatONnetwork/AppChain-SDK/core/pevm"
	"github.com/PlatONnetwork/AppChain-SDK/x/benchmark"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/mock"
	"github.com/PlatONnetwork/PlatON-Go/consensus"
	coresdk "github.com/PlatONnetwork/PlatON-Go/core/sdk"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

var (
	forceSequential  bool
	numTxs           int
	concurrencyLevel int
	batchTxs         int
)

type chainContext struct {
	engine *consensus.BftMock
}

func newChainContext() *chainContext {
	return &chainContext{
		engine: consensus.NewFaker(),
	}
}

func (cc *chainContext) Engine() coresdk.Engine {
	return cc.engine
}

func (cc *chainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	b := cc.engine.GetBlockByHashAndNum(hash, number)
	if b != nil {
		return b.Header()
	}
	return nil
}

type mockContractsApp struct{}

func newMockContractsApp() *mockContractsApp { return &mockContractsApp{} }
func (m *mockContractsApp) Contracts(sdk.StateDBReader, uint64) []sdk.SDKContract {
	return []sdk.SDKContract{}
}

type account struct {
	addr       common.Address
	privateKey *ecdsa.PrivateKey
}

func prepareAccounts(n int) []*account {
	accounts := make([]*account, n)
	for i := 0; i < n; i++ {
		key, _ := crypto.GenerateKey()
		accounts[i] = &account{
			addr:       crypto.PubkeyToAddress(key.PublicKey),
			privateKey: key,
		}
	}
	return accounts
}

func prepareStateDB(accounts []*account, balance *big.Int) sdk.StateDB {
	statedb := mock.NewMockStateDB()
	for _, acc := range accounts {
		statedb.Balance[acc.addr] = new(big.Int).Set(balance)
		statedb.Nonce[acc.addr] = 0
	}
	return statedb
}

func prepareHeader(number int64) *types.Header {
	return &types.Header{
		ParentHash: common.ZeroHash,
		Number:     new(big.Int).SetInt64(number),
		GasLimit:   5000000000,
		Time:       uint64(time.Now().UnixMilli()),
		BaseFee:    big.NewInt(1000),
		Coinbase:   common.HexToAddress("0x0000012312312321"),
	}
}

func prepareTransactions(n int, accounts []*account, amount *big.Int) types.Transactions {
	txs := make(types.Transactions, n)
	signer := types.NewEIP155Signer(big.NewInt(params.TestChainConfig.ChainID.Int64()))
	for i := 0; i < n; i++ {
		from := accounts[i%len(accounts)]
		//to := accounts[(i+1)%len(accounts)]
		to := benchmark.GenRecipient(uint64(i + 1000))
		tx, _ := types.SignNewTx(from.privateKey, signer, &types.LegacyTx{
			Nonce:    0,
			To:       &to,
			Value:    new(big.Int).Set(amount),
			Gas:      21000,
			GasPrice: big.NewInt(5000),
		})
		txs[i] = tx
	}
	return txs
}

func main() {
	flag.IntVar(&numTxs, "num-txs", 10, "Number of transactions to generates")
	flag.IntVar(&concurrencyLevel, "concurrency-level", 4, "Number of goroutines for parallel executions")
	flag.IntVar(&batchTxs, "batch-txs", 10, "Number of batch transactions to parallel execution")
	flag.BoolVar(&forceSequential, "force-sequential", false, "Force to sequential execution")
	flag.Parse()

	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	accounts := prepareAccounts(numTxs)
	txs := prepareTransactions(numTxs, accounts, big.NewInt(100))
	for i := 0; i < 1000; i++ {
		env1 := &pevm.Env{
			Header:        prepareHeader(10),
			StateDB:       prepareStateDB(accounts, big.NewInt(30000000000)),
			ChainConfig:   params.TestChainConfig,
			VMConfig:      vm.Config{},
			ChainContext:  newChainContext(),
			IsWorker:      true,
			BlockDeadline: time.Now().Add(time.Minute),
		}
		pevm1 := pevm.NewPEVM(forceSequential, concurrencyLevel, batchTxs, log.New("module", "pevmtest"), env1, newMockContractsApp())
		_, err1 := pevm1.Run(txs, false)
		if err1 != nil {
			panic(err1)
		}
	}
}
