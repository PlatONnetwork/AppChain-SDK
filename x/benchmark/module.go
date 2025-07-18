package benchmark

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/x/benchmark/contracts"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/rpc"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
	"math/big"
	"sync"
	"sync/atomic"
	"time"
)

const (
	ModuleName    = "benchmark"
	ModuleVersion = 0
	AccountLimit  = 2000
)

var (
	Deployer         = common.BigToAddress(big.NewInt(1))
	PendingLimitFlag = cli.Uint64Flag{
		Name:  "benchmark.pendinglimit",
		Usage: "How many transactions are packaged after sending a transaction",
	}
)

func AddBenchmarkFlags(app *cli.App) {
	app.Flags = append(app.Flags, PendingLimitFlag)
}

type Account struct {
	key  *ecdsa.PrivateKey
	addr common.Address
}

type TxPool interface {
	Nonce(addr common.Address) uint64
	AddLocal(tx *types.Transaction) error
}

type Params struct {
	startIndex        int
	endIndex          int
	rawTxPercent      int
	contractTxPercent int
	sendTxPool        bool
	amount            int
	txsPerAccount     int
}
type Statistics struct {
	start   time.Time
	send    atomic.Uint64
	confirm atomic.Uint64
}
type Module struct {
	Params
	Statistics
	sync.Mutex
	pendingLimit uint64
	logger       log.Logger
	db           *DB
	keys         []*Account
	txCache      map[common.Address][]*types.Transaction
	sent         sync.Map //map[common.Hash]uint64
	signer       types.Signer

	txPool   TxPool
	starting atomic.Bool
	stopC    chan struct{}
}

func NewModule(ctx *cli.Context, store store.Store) *Module {
	m := &Module{
		logger:       log.New("module", ModuleName),
		db:           NewDB(store),
		txCache:      make(map[common.Address][]*types.Transaction),
		stopC:        make(chan struct{}),
		pendingLimit: ctx.GlobalUint64(PendingLimitFlag.Name),
	}
	m.initAccount()
	return m
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}
func (m *Module) Init(ctx sdk.InitContext) error {
	m.txPool = ctx.Backend().TxPool()
	chainId, err := ctx.Backend().ChainId()
	if err != nil {
		return err
	}
	m.signer = types.NewLondonSigner(chainId)
	return nil
}
func (m *Module) initAccount() error {
	var err error
	keys, err := GenKeys(AccountLimit)
	for _, k := range keys {
		m.keys = append(m.keys, &Account{
			key:  k,
			addr: crypto.PubkeyToAddress(k.PublicKey),
		})
	}
	return err
}

func (m *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	//var genesis GenesisConfig
	//if err := json.Unmarshal(data, &genesis); err != nil {
	//	return err
	//}

	for i := 0; i < AccountLimit; i++ {
		db.AddBalance(m.keys[i].addr, new(big.Int).Mul(big.NewInt(1000000000000000000), big.NewInt(10000000)))
		caller, err := contracts.NewBenchTokenGenesisCaller(ctx, db, chainConfig)
		if err != nil {
			return err
		}
		contractAddr := GenContract(uint64(i))
		caller.WithCaller(Deployer).WithTo(contractAddr).DeployBenchToken(m.keys[i].addr, big.NewInt(1000000000000))
		_, err = caller.WithCaller(m.keys[i].addr).WithTo(contractAddr).Transfer(GenRecipient(uint64(i)), big.NewInt(1))
		if err != nil {
			panic(err)
		}
	}
	return nil
}
func (m *Module) APIs() []rpc.API {
	fmt.Println("benchmark rpc")
	return []rpc.API{
		rpc.API{
			Namespace: "benchmark",
			Version:   "1",
			Service:   NewRPC(m),
			Public:    true,
		},
	}
}

func (m *Module) createTransactions(amount uint64) error {
	rawStartIndex := m.startIndex
	rawEndIndex := m.startIndex + (m.endIndex+1-m.startIndex)*m.rawTxPercent/(m.rawTxPercent+m.contractTxPercent)
	contractStartIndex := rawEndIndex
	contractEndIndex := m.endIndex + 1
	sum := uint64(0)
	m.logger.Debug("create tx", "rawStartIndex", rawStartIndex, "rawEndIndex", rawEndIndex, "contractStartIndex", contractStartIndex, "contractEndIndex", contractEndIndex, "amount", amount)
	for amount > sum {
		for i := rawStartIndex; i < rawEndIndex && amount > sum; i, sum = i+1, sum+1 {
			m.Lock()
			k := m.keys[i]
			txs := m.txCache[k.addr]
			nonce := uint64(0)
			if len(txs) == 0 {
				nonce = m.txPool.Nonce(k.addr)
				m.logger.Debug("create tx", "address", k.addr, "nonce", nonce)
			} else {
				nonce = txs[len(txs)-1].Nonce() + 1
			}
			tx, err := createRawTransfer(m.signer, k.key, GenRecipient(uint64(i)), nonce)
			if err != nil {
				m.Unlock()
				return err
			}
			types.Sender(m.signer, tx)
			txs = append(txs, tx)
			m.txCache[k.addr] = txs
			m.Unlock()
		}

		for i := contractStartIndex; i < contractEndIndex && amount > sum; i, sum = i+1, sum+1 {
			m.Lock()
			k := m.keys[i]
			txs := m.txCache[k.addr]
			nonce := uint64(0)
			if len(txs) == 0 {
				nonce = m.txPool.Nonce(k.addr)
			} else {
				nonce = txs[len(txs)-1].Nonce() + 1
			}
			tx, err := createTokenTransfer(m.signer, k.key, GenContract(uint64(i)), GenRecipient(uint64(i)), nonce)
			if err != nil {
				m.Unlock()
				return err
			}
			types.Sender(m.signer, tx)
			txs = append(txs, tx)
			m.txCache[k.addr] = txs
			m.Unlock()
		}
	}
	for k, v := range m.txCache {
		m.logger.Debug("create tx result", "address", k, "len", len(v), "nonce", v[0].Nonce())
	}
	return nil
}

func (m *Module) start(amount uint64, txsPerAccount int, sendTxPool bool) error {
	m.Lock()
	defer m.Unlock()
	m.txsPerAccount = txsPerAccount
	m.sendTxPool = sendTxPool
	m.amount = int(amount)
	if m.starting.Load() {
		return errors.New("had started")
	}
	m.starting.Store(true)
	m.Statistics.start = time.Now()
	if m.sendTxPool {
		go m.sendLoop(amount)
	}
	return nil
}
func (m *Module) stop() error {
	if m.starting.Load() {
		m.starting.Store(false)
		m.Statistics.send.Store(0)
		m.Statistics.confirm.Store(0)
	}
	return nil
}

func (m *Module) SortTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions, remote map[common.Address]types.Transactions) (types.Transactions, error) {
	m.logger.Debug("benchmark sort txs", "local", len(local), "remote", len(remote), "number", ctx.Header().Number)
	if m.starting.Load() {
		m.Lock()
		defer m.Unlock()
		var txs []types.Transactions
		sum := 0
		for k, v := range m.txCache {
			nonce := ctx.StateDB().GetNonce(k)
			start := int(nonce - v[0].Nonce())
			end := start + m.txsPerAccount
			if end > len(v) {
				end = len(v)
			}
			if sum+end-start > m.amount {
				end = start + m.amount - sum
			}
			m.logger.Debug("Sort Txs", "address", k, "len", len(v), "nonce", nonce, "firstNonce", v[0].Nonce(), "start", start, "end", end, "amount", m.amount, "txsPerAccount", m.txsPerAccount, "sum", sum)
			txs = append(txs, v[start:end])

			sum += end - start
			m.send.Add(uint64(end - start))
			if sum >= m.amount {
				break
			}
			m.txCache[k] = v[start:]
		}
		res := make(types.Transactions, 0, sum)
		for _, s := range txs {
			res = append(res, s...)
		}

		return res, nil
	}
	if m.send.Load() >= m.pendingLimit {
		allTxs := make(types.Transactions, 0)
		for _, txs := range local {
			allTxs = append(allTxs, txs...)
		}
		for _, txs := range remote {
			allTxs = append(allTxs, txs...)
		}
		return allTxs, nil
	} else {
		return make(types.Transactions, 0), nil
	}

}
func (m *Module) AddTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions) (map[common.Address]types.Transactions, error) {

	if !m.sendTxPool {
		//
	}
	return local, nil
}

func (m *Module) OnCommit(ctx sdk.ConsensusContext, block *types.Block) error {
	handleBlock := func() {
		m.Lock()
		defer m.Unlock()
		now := block.Time()
		elapsed := uint64(0)
		count := 0
		for _, tx := range block.Transactions() {
			if t, ok := m.sent.Load(tx.Hash()); ok {
				elapsed += now - t.(uint64)
				count++
			}
		}
		m.confirm.Add(uint64(count))
		blockInfo := &BlockInfo{
			ProduceTime: now,
			Number:      block.NumberU64(),
			TotalLength: block.Transactions().Len(),
			TxLength:    count,
			TimeUse:     elapsed,
		}
		m.db.Set(blockInfo)
	}
	go handleBlock()
	return nil
}

func (m *Module) sendLoop(amount uint64) {
	tick := time.NewTicker(time.Second)
	total := m.endIndex - m.startIndex + 1
	index := 0
	for {
		select {
		case <-tick.C:
			if !m.starting.Load() {
				continue
			}
			m.Lock()
			sum := uint64(0)
			for sum < amount && len(m.txCache) != 0 {
				pos := index%total + m.startIndex
				txs := m.txCache[m.keys[pos].addr]
				if len(txs) > 0 {
					err := m.txPool.AddLocal(txs[0])
					if err != nil {
						m.logger.Warn("Add local tx failed", "error", err, m.keys[pos].addr.Hex())
						continue
					} else {
						m.sent.Store(txs[0].Hash(), uint64(time.Now().UnixMilli()))
						m.txCache[m.keys[pos].addr] = txs[1:]
						m.send.Add(1)
					}
				} else {
					delete(m.txCache, m.keys[pos].addr)
					m.logger.Warn("txs is empty", "addr", m.keys[pos].addr.Hex(), "pos", pos, "total", total, "index", index, "startIndex", m.startIndex)
				}
				sum++
				index++
			}
			m.Unlock()
		}
	}
}
