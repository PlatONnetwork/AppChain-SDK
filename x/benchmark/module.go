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
	"math/big"
	"sync"
	"sync/atomic"
	"time"
)

const (
	ModuleName    = "benchmark"
	ModuleVersion = 0
	AccountLimit  = 100000
)

var (
	Deployer = common.BigToAddress(big.NewInt(1))
)

type Account struct {
	key  *ecdsa.PrivateKey
	addr common.Address
}

type TxPool interface {
	Nonce(addr common.Address) uint64
	AddLocal(tx *types.Transaction) error
}
type MockTxPool struct {
}

func (MockTxPool) Nonce(addr common.Address) uint64 {
	return 0
}
func (MockTxPool) AddLocal(tx *types.Transaction) error {
	return nil
}

type Params struct {
	startIndex        int
	endIndex          int
	rawTxPercent      int
	contractTxPercent int
	sendTxPool        bool
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
	logger  log.Logger
	db      *DB
	keys    []*Account
	txCache map[common.Address][]*types.Transaction
	sent    sync.Map //map[common.Hash]uint64
	signer  types.Signer

	txPool   TxPool
	starting atomic.Bool
	stopC    chan struct{}
}

func NewModule(store store.Store) *Module {
	m := &Module{
		logger:  log.New("module", ModuleName),
		db:      NewDB(store),
		txCache: make(map[common.Address][]*types.Transaction),
		stopC:   make(chan struct{}),
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
		db.AddBalance(m.keys[i].addr, big.NewInt(1000000000000000000))
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
	rawEndIndex := (m.endIndex - m.startIndex) * m.rawTxPercent / (m.rawTxPercent + m.contractTxPercent)
	contractStartIndex := rawEndIndex + 1
	contractEndIndex := m.endIndex
	sum := uint64(0)
	for amount > sum {
		for i := rawStartIndex; i <= rawEndIndex && amount > sum; i, sum = i+1, sum+1 {
			k := m.keys[i]
			txs := m.txCache[k.addr]
			nonce := uint64(0)
			if len(txs) == 0 {
				nonce = m.txPool.Nonce(k.addr)
			} else {
				nonce = txs[len(txs)-1].Nonce() + 1
			}
			tx, err := createRawTransfer(m.signer, k.key, GenRecipient(uint64(i)), nonce)
			if err != nil {
				return err
			}
			txs = append(txs, tx)
			m.txCache[k.addr] = txs
		}

		for i := contractStartIndex; i <= contractEndIndex && amount > sum; i, sum = i+1, sum+1 {
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
				return err
			}
			txs = append(txs, tx)
			m.txCache[k.addr] = txs
		}
	}
	return nil
}

func (m *Module) start(amount uint64, sendTxPool bool) error {
	m.Lock()
	defer m.Unlock()
	m.sendTxPool = sendTxPool
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
		m.stopC <- struct{}{}
	}
	return nil
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
			sum := uint64(0)
			for sum < amount {
				pos := index%total + m.startIndex
				txs := m.txCache[m.keys[pos].addr]
				if len(txs) > 0 {
					err := m.txPool.AddLocal(txs[0])
					if err != nil {
						m.logger.Warn("Add local tx failed", "error", err)
						continue
					} else {
						m.sent.Store(txs[0].Hash(), uint64(time.Now().UnixMilli()))
						m.txCache[m.keys[pos].addr] = txs[1:]
						m.send.Add(1)
					}
				}
				sum++
				index++
			}
		case <-m.stopC:
			m.starting.Store(false)
			m.Statistics.send.Store(0)
			m.Statistics.confirm.Store(0)
			return
		}
	}
}
