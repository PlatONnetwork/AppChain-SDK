package benchmark

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
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
	"path/filepath"
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
	Deployer = common.BigToAddress(big.NewInt(1))

	TxFileFlag = cli.StringFlag{
		Name:  "benchmark.txfile",
		Usage: "Transaction file",
	}
)

func AddBenchmarkFlags(app *cli.App) {
	app.Flags = append(app.Flags, TxFileFlag)

}

type Account struct {
	key   *ecdsa.PrivateKey
	addr  common.Address
	nonce *uint64
}

type TxPool interface {
	Nonce(addr common.Address) uint64
	AddLocal(tx *types.Transaction) error
}

type TxPoolModule interface {
	AddLocalBatch(addr common.Address, txs []*types.Transaction)
	Pending(getNonce func(addr common.Address) uint64, limit int, txsPerAccount int) types.Transactions
	Total() uint64
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
	cache   atomic.Uint64
	send    atomic.Uint64
	confirm atomic.Uint64
}
type Module struct {
	Params
	Statistics
	sync.Mutex
	lastSortTxBlock uint64
	logger          log.Logger
	db              *DB
	keys            []*Account
	txCache         map[common.Address][]*types.Transaction
	sent            sync.Map //map[common.Hash]uint64
	signer          types.Signer
	mmapTxFile      *MmapTxFile
	readyCh         chan struct{}
	txPoolModule    TxPoolModule
	txPool          TxPool
	starting        atomic.Bool
	stopC           chan struct{}
	wg              sync.WaitGroup
}

func NewModule(ctx *cli.Context, store store.Store, txPoolModule TxPoolModule, datadir string) *Module {
	m := &Module{
		logger:       log.New("module", ModuleName),
		db:           NewDB(store),
		txPoolModule: txPoolModule,
		txCache:      make(map[common.Address][]*types.Transaction),
		readyCh:      make(chan struct{}),
		stopC:        make(chan struct{}),
	}
	var err error
	txFile := ctx.GlobalString(TxFileFlag.Name)
	create := false
	if len(txFile) == 0 {
		txFile = filepath.Join(datadir, "benchmarktxs")
		create = true
	}
	if m.mmapTxFile, err = NewMmapTxFile(txFile, create); err != nil {
		m.logger.Crit("Create mmap file failed", "err", err)
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
	txsCache := make(map[common.Address][]*types.Transaction)
	m.logger.Debug("create tx", "rawStartIndex", rawStartIndex, "rawEndIndex", rawEndIndex, "contractStartIndex", contractStartIndex, "contractEndIndex", contractEndIndex, "amount", amount)
	for amount > sum {
		for i := rawStartIndex; i < rawEndIndex && amount > sum; i, sum = i+1, sum+1 {
			m.Lock()
			k := m.keys[i]
			txs := txsCache[k.addr]
			nonce := uint64(0)
			if k.nonce == nil {
				nonce = m.txPool.Nonce(k.addr)
				m.logger.Debug("create tx", "address", k.addr, "nonce", nonce)
			} else {
				nonce = *k.nonce + 1
			}
			k.nonce = &nonce
			m.keys[i] = k
			tx, err := createRawTransfer(m.signer, k.key, GenRecipient(uint64(i)), nonce)
			if err != nil {
				m.Unlock()
				return err
			}
			tx.CacheFromAddr(m.signer, k.addr)
			txs = append(txs, tx)
			txsCache[k.addr] = txs
			m.cache.Add(1)

			m.Unlock()
		}

		for i := contractStartIndex; i < contractEndIndex && amount > sum; i, sum = i+1, sum+1 {
			m.Lock()
			k := m.keys[i]
			txs := txsCache[k.addr]
			nonce := uint64(0)
			if k.nonce == nil {
				nonce = m.txPool.Nonce(k.addr)
			} else {
				nonce = *k.nonce + 1
			}
			k.nonce = &nonce
			m.keys[i] = k
			tx, err := createTokenTransfer(m.signer, k.key, GenContract(uint64(i)), GenRecipient(uint64(i)), nonce)
			if err != nil {
				m.Unlock()
				return err
			}
			tx.CacheFromAddr(m.signer, k.addr)
			txs = append(txs, tx)
			txsCache[k.addr] = txs
			m.cache.Add(1)
			m.Unlock()
		}
	}
	for len(txsCache) != 0 {
		for k, v := range txsCache {
			end := 100
			if len(v) < 100 {
				end = len(v)
			}
			m.mmapTxFile.WriteTxs(k, v[:end])
			txsCache[k] = v[end:]
			if len(v[end:]) == 0 {
				delete(txsCache, k)
			}
		}
	}
	return nil
}

func (m *Module) start(amount uint64, txsPerAccount int, sendTxPool bool) error {
	m.Lock()
	defer m.Unlock()
	if m.starting.Load() {
		return errors.New("had started")
	}
	m.starting.Store(true)
	m.txsPerAccount = txsPerAccount
	m.sendTxPool = sendTxPool
	m.amount = int(amount)
	m.Statistics.start = time.Now()
	m.Statistics.send.Store(0)
	m.Statistics.confirm.Store(0)
	m.mmapTxFile.UnMmap()
	m.mmapTxFile.Mmap()
	m.wg.Add(1)
	go m.decodeTxLoop(amount)
	if m.sendTxPool {
		go m.sendLoop(amount)
	}
	return nil
}

func (m *Module) stop() error {
	if m.starting.Load() {
		close(m.stopC)
		m.readReady()
		m.wg.Wait()
		m.stopC = make(chan struct{})
		m.starting.Store(false)
		m.logger.Info("Stop benchmark")
	}
	return nil
}
func (m *Module) decodeTxLoop(amount uint64) {
	defer m.wg.Done()
	if m.mmapTxFile == nil {
		m.logger.Info("Mmap unmmap, stop decode tx")
		return
	}
	sum := uint64(0)
	start := time.Now()
	m.logger.Debug("Start decode Tx")
	for {
		select {
		case <-m.stopC:
			m.logger.Info("Receive stop signal, stop decode txs")
			return
		default:
			addr, txs, err := m.mmapTxFile.ReadTxs()
			if err != nil {
				m.logger.Error("Read txs failed", "err", err)
				m.starting.Store(false)
				return
			}
			for _, tx := range txs {
				tx.CacheFromAddr(m.signer, addr)
			}
			if m.sendTxPool {
				m.Lock()
				cache := m.txCache[addr]
				cache = append(cache, txs...)
				m.txCache[addr] = cache
				m.Unlock()
			} else {
				m.send.Add(uint64(len(txs)))
				m.txPoolModule.AddLocalBatch(addr, txs)
			}
			sum += uint64(len(txs))
			if sum >= amount {
				m.logger.Debug("Had ready txs, send ready signal", "cost", time.Since(start), "sum", sum, "amount", amount)
				sum = 0
				m.readyCh <- struct{}{}
				start = time.Now()
				//for _, tx := range txs {
				//	m.sent.Store(tx.Hash(), uint64(time.Now().UnixMilli()))
				//}

			}
		}

	}
	m.logger.Debug("Stop decode Tx")

}
func (m *Module) readReady() {
	select {
	case <-m.readyCh:
	default:
	}
}
func (m *Module) SortTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions, remote map[common.Address]types.Transactions) (types.Transactions, error) {
	statedb := ctx.StateDB()
	if !m.starting.Load() {
		return m.txPoolModule.Pending(func(addr common.Address) uint64 {
			return statedb.GetNonce(addr)
		}, 20000, 200), nil
	}
	return m.txPoolModule.Pending(func(addr common.Address) uint64 {
		return statedb.GetNonce(addr)
	}, m.amount, m.txsPerAccount), nil
}
func (m *Module) AddTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions) (map[common.Address]types.Transactions, error) {
	m.logger.Debug("Read ready signal")
	if m.txPoolModule.Total() < uint64(m.amount*6) {
		m.readReady()
	}
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
		//for _, tx := range block.Transactions() {
		//	if t, ok := m.sent.Load(tx.Hash()); ok {
		//		elapsed += now - t.(uint64)
		//		count++
		//		m.sent.Delete(tx.Hash())
		//	}
		//}
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
			m.readReady()
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
