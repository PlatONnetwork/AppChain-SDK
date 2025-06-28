package asyncblock

import (
	"errors"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/core/pevm"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

const (
	UNINITIALIZED = uint32(0)
	READY         = uint32(1)
	RUNNING       = uint32(2)
	PAUSE         = uint32(3)
	DONE          = uint32(4)
	INTERRUPT     = uint32(5)
)

type EnvCreateFn func(header *types.Header, parentSealHash, parentHash common.Hash, parentNumber uint64) (sdk.StateDB, *pevm.PEVM, sdk.WorkerContext, error)
type FinalizeFn func(executor *EntryExecutor) error

type ConsensusState struct {
	sync.Mutex
	Self enode.ID

	LeaderIndex int
	Epoch       uint64
	ViewNumber  uint64
	Validator   []*cbfttypes.ValidateNode
}

func (c *ConsensusState) EpochView() (uint64, uint64) {
	c.Lock()
	defer c.Unlock()
	return c.Epoch, c.ViewNumber
}

func (c *ConsensusState) Update(epoch, viewNumber uint64, validator []*cbfttypes.ValidateNode, leaderIndex int) {
	c.Lock()
	defer c.Unlock()
	c.Epoch = epoch
	c.ViewNumber = viewNumber
	c.Validator = validator
	c.LeaderIndex = leaderIndex
}

func (c *ConsensusState) VerifyEntry(entry *Entry) error {
	c.Lock()
	defer c.Unlock()
	if entry.Epoch == c.Epoch && (entry.View == c.ViewNumber) {
		return entry.VerifySign(c.Validator[c.LeaderIndex].BlsPubKey)
	} else if entry.Epoch == c.Epoch && (entry.View == c.ViewNumber+1) {
		return entry.VerifySign(c.Validator[(c.LeaderIndex+1)%len(c.Validator)].BlsPubKey)
	}
	return fmt.Errorf("unexcepted consensus, epoch:%d, view:%d, actual epoch:%d, view:%d", c.Epoch, c.ViewNumber, entry.Epoch, entry.View)
}

type EntryExecutorTree struct {
	sync.Mutex
	tree        map[uint64]map[uint64]*EntryExecutor
	envCreateFn EnvCreateFn
	finalizeFn  FinalizeFn
}

func NewEntryExecutorTree(envCreateFn EnvCreateFn, finalizeFn FinalizeFn) *EntryExecutorTree {
	return &EntryExecutorTree{
		tree:        make(map[uint64]map[uint64]*EntryExecutor),
		envCreateFn: envCreateFn,
		finalizeFn:  finalizeFn,
	}
}
func (e *EntryExecutorTree) AddEntry(entry *Entry) {
	e.Lock()
	defer e.Unlock()
	executors := e.tree[entry.BlockNumber]
	if executors == nil {
		executors = make(map[uint64]*EntryExecutor)
		executor := NewEntryExecutor(e.envCreateFn, e.finalizeFn)
		executor.AddEntry(entry)
		executors[entry.View] = executor
		e.tree[entry.BlockNumber] = executors
	} else if executor, ok := executors[entry.View]; !ok {
		executor = NewEntryExecutor(e.envCreateFn, e.finalizeFn)
		executor.AddEntry(entry)
		executors[entry.View] = executor
	} else {
		executor.AddEntry(entry)
		executors[entry.View] = executor
	}
}
func (e *EntryExecutorTree) FullEntry(blockNumber uint64) bool {
	e.Lock()
	defer e.Unlock()
	if fragment, ok := e.tree[blockNumber]; ok {
		for _, f := range fragment {
			if f.LastEntry().Ending == 1 {
				return true
			}
		}
	}
	return false
}
func (e *EntryExecutorTree) TryFindBody(view uint64, blockNumber uint64) (types.Transactions, error) {
	e.Lock()
	defer e.Unlock()
	if fragment, ok := e.tree[blockNumber]; ok {
		if ex, ok := fragment[view]; ok {
			length := ex.entryList.Len()
			if length > 0 &&
				ex.entryList.Get(length-1).Ending == 1 {
				return ex.Transactions(), nil
			}
		}
	}

	return nil, errors.New("no block body")
}
func (e *EntryExecutorTree) GetEntry(blockNumber uint64, view uint64, entryNumber uint32) *Entry {
	e.Lock()
	defer e.Unlock()
	if fragment, ok := e.tree[blockNumber]; ok {
		if ex, ok := fragment[view]; ok {
			if uint32(ex.entryList.Len()) > entryNumber {
				return ex.entryList.Get(int(entryNumber))
			}
		}
	}
	return nil
}
func (e *EntryExecutorTree) GetBlockExecutor(blockNumber uint64) []*EntryExecutor {
	e.Lock()
	defer e.Unlock()
	if fragment, ok := e.tree[blockNumber]; ok {
		var exes []*EntryExecutor
		for _, v := range fragment {
			exes = append(exes, v)
		}
		return exes
	}
	return nil
}

func (e *EntryExecutorTree) Clean(epoch, view uint64, finishFn func(e *EntryExecutor)) {
	e.Lock()
	defer e.Unlock()
	for k, exes := range e.tree {
		for _, exe := range exes {
			if exe.epoch < epoch || (exe.epoch == epoch && exe.view < view) {
				//core.CleanStateDB(exe.statedb)
				if exe.running.Load() == DONE {
					finishFn(exe)
				}
				delete(exes, exe.view)
			}
			if exe.running.Load() != DONE {
				exe.Stop()
			}
		}
		if len(exes) == 0 {
			delete(e.tree, k)
		}
	}
}

func (e *EntryExecutorTree) FindBlockNumbers() []uint64 {
	e.Lock()
	defer e.Unlock()
	nums := make([]uint64, 0, len(e.tree))
	for k, _ := range e.tree {
		nums = append(nums, k)
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums
}

type BlockStateCache struct {
	sync.Mutex
	logger      log.Logger
	executing   atomic.Value
	splitBlock  sync.Map //map[uint64]struct{}
	blocks      sync.Map
	fragments   *EntryExecutorTree
	envCreateFn EnvCreateFn
	finalizeFn  FinalizeFn
}

func NewBlockStateCache(envCreateFn EnvCreateFn, finalizeFn FinalizeFn) *BlockStateCache {
	return &BlockStateCache{
		logger:      log.New("module", ModuleName, "component", "blockstatecache"),
		fragments:   NewEntryExecutorTree(envCreateFn, finalizeFn),
		envCreateFn: envCreateFn,
		finalizeFn:  finalizeFn,
	}
}
func (b *BlockStateCache) IsSplit(num uint64) bool {
	_, ok := b.splitBlock.Load(num)
	return ok
}
func (b *BlockStateCache) AddSplit(num uint64) {
	b.splitBlock.Store(num, struct{}{})
}
func (b *BlockStateCache) AddFinishExecutor(executor *EntryExecutor) {
	b.blocks.Store(executor.block.Header().SealHash(), executor)
}
func (b *BlockStateCache) FindTransactions(hash common.Hash) (types.Transactions, error) {
	b.Lock()
	defer b.Unlock()
	if entry, ok := b.blocks.Load(hash); ok {
		return entry.(*EntryExecutor).block.Transactions(), nil
	}
	return nil, errors.New("no block body")
}
func (b *BlockStateCache) FindStateDB(hash common.Hash) (sdk.StateDB, types.Receipts) {
	if entry, ok := b.blocks.Load(hash); ok {
		return entry.(*EntryExecutor).statedb, entry.(*EntryExecutor).receipts
	}
	return nil, nil
}

func (b *BlockStateCache) FindHeader(hash common.Hash) *types.Header {
	b.Lock()
	defer b.Unlock()
	if entry, ok := b.blocks.Load(hash); ok {
		return &entry.(*EntryExecutor).header
	}
	return nil
}

func (b *BlockStateCache) AddEntry(entry *Entry) {
	b.fragments.AddEntry(entry)

}
func (b *BlockStateCache) TryFindBody(view uint64, blockNumber uint64) (types.Transactions, error) {
	return b.fragments.TryFindBody(view, blockNumber)
}
func (b *BlockStateCache) GetEntry(blockNumber uint64, view uint64, entryNumber uint32) *Entry {
	return b.fragments.GetEntry(blockNumber, view, entryNumber)
}
func (b *BlockStateCache) HadFullEntry(blockNumber uint64) bool {
	return b.fragments.FullEntry(blockNumber)
}
func (b *BlockStateCache) RemoveBlock(sealHash common.Hash) {
	b.blocks.Delete(sealHash)
}
func (b *BlockStateCache) Clean(epoch, view uint64) {
	b.fragments.Clean(epoch, view, func(e *EntryExecutor) {
		b.blocks.Delete(e.block.Header().SealHash())
	})
}

func (b *BlockStateCache) FindExecutableEntry() {
	exe := b.executing.Load()
	if exe != nil && exe.(*EntryExecutor).Status() == RUNNING {
		return
	}
	b.Lock()
	defer b.Unlock()
	nums := b.fragments.FindBlockNumbers()
	skip := false
	b.logger.Debug("Find block number entry", "nums", nums)
	for _, n := range nums {
		if skip {
			break
		}
		for _, v := range b.fragments.GetBlockExecutor(n) {
			status := v.Status()
			b.logger.Debug("Block executor status", "blockNumber", n, "status", status)
			switch status {
			case READY, PAUSE:
				b.executing.Store(v)
				v.Execute()
				if v.Status() != DONE {
					skip = true
				}
				b.logger.Debug("Block executor ready, execute enable entry", "blockNumber", n, "finish", !skip)

				break
			case RUNNING:
				skip = true
				break
			}
		}
	}
}

type EntryList struct {
	sync.Mutex
	entries    []*Entry
	entryCache map[uint32]*Entry
}

func NewEntryList() *EntryList {
	return &EntryList{
		entryCache: make(map[uint32]*Entry),
	}
}
func (e *EntryList) Get(index int) *Entry {
	e.Lock()
	defer e.Unlock()
	return e.entries[index]
}
func (e *EntryList) Len() int {
	e.Lock()
	defer e.Unlock()
	return len(e.entries)
}

func (e *EntryList) Transactions() types.Transactions {
	e.Lock()
	defer e.Unlock()
	var txs types.Transactions
	for _, entry := range e.entries {
		txs = append(txs, entry.Transactions...)
	}
	return txs
}
func (e *EntryList) AddEntry(entry *Entry) {
	e.Lock()
	defer e.Unlock()
	e.entryCache[entry.EntryNumber] = entry
	if (len(e.entries) == 0 && entry.EntryNumber == 0) || (len(e.entries) > 0 && e.entries[len(e.entries)-1].EntryNumber+1 == entry.EntryNumber) {
		e.entries = append(e.entries, entry)
	}
}
func (e *EntryList) LastEntry() *Entry {
	e.Lock()
	defer e.Unlock()
	return e.entries[len(e.entries)-1]
}

type EntryExecutor struct {
	sync.Mutex
	start          time.Time
	logger         log.Logger
	running        atomic.Uint32
	block          *types.Block
	parentSealHash common.Hash
	header         types.Header
	workCtx        sdk.WorkerContext
	epoch          uint64
	view           uint64
	entryIndex     int
	entryList      *EntryList
	statedb        sdk.StateDB
	vm             *pevm.PEVM
	receipts       types.Receipts
	envCreateFn    EnvCreateFn
	finalizeFn     FinalizeFn
}

func NewEntryExecutor(envCreateFn EnvCreateFn, finalizeFn FinalizeFn) *EntryExecutor {
	return &EntryExecutor{
		start:       time.Now(),
		logger:      log.New("module", ModuleName, "component", "entryexecutor"),
		entryList:   NewEntryList(),
		envCreateFn: envCreateFn,
		finalizeFn:  finalizeFn,
	}
}

func (e *EntryExecutor) Status() uint32 {
	e.Lock()
	defer e.Unlock()
	if e.statedb == nil && e.entryList.Len() != 0 {
		e.logger.Debug("Get ready to create env")
		var err error
		e.statedb, e.vm, e.workCtx, err = e.envCreateFn(&e.header, e.parentSealHash, e.header.ParentHash, e.header.Number.Uint64()-1)
		if err == nil {
			e.running.Store(READY)
		}
	}
	return e.running.Load()
}
func (e *EntryExecutor) Transactions() types.Transactions {
	return e.entryList.Transactions()
}
func (e *EntryExecutor) Stop() {
	e.running.Store(INTERRUPT)
}

func (e *EntryExecutor) AddEntry(entry *Entry) {
	if entry.EntryNumber == 0 {
		e.logger = log.New("module", ModuleName, "epoch", entry.Epoch, "view", entry.View, "number", entry.BlockNumber)
	}
	e.entryList.AddEntry(entry)
	if entry.EntryNumber == 0 {
		e.parentSealHash = entry.ParentSealHash
		e.header = *entry.Header
		e.epoch = entry.Epoch
		e.view = entry.View
	}

}

func (e *EntryExecutor) LastEntry() *Entry {
	return e.entryList.LastEntry()
}

func (e *EntryExecutor) AsyncExecute() {
	if e.running.Load() == RUNNING {
		return
	}
	go func() {
		e.Execute()
	}()
}

func (e *EntryExecutor) Execute() {
	if e.running.Load() == RUNNING {
		return
	}
	e.Lock()
	defer e.Unlock()
	e.running.Store(RUNNING)
	e.logger.Trace("Try to execute entry", "index", e.entryIndex)
	status := PAUSE
	for e.entryList.Len() > e.entryIndex {
		entry := e.entryList.Get(e.entryIndex)
		result, err := e.vm.Run(entry.Transactions, false)
		if err != nil {
			e.logger.Warn("Execute entry failed, entry got interrupt", "index", e.entryIndex, "err", err)
			status = INTERRUPT
			break
		}
		e.receipts = append(e.receipts, result.Receipts...)
		e.header.GasUsed += result.GasUsed
		e.entryIndex++
		if entry.Ending == 1 {
			status = DONE
			if err := e.finalizeFn(e); err != nil {
				e.logger.Warn("Finalize entry failed, entry got interrupt", "index", e.entryIndex, "err", err)
				status = INTERRUPT
			}
			e.logger.Debug("Execute finish", "index", e.entryIndex)
		}
	}
	e.running.Store(status)
}

type NewFinalizeBlockEvent struct {
	Block    *types.Block
	StateDB  sdk.StateDB
	Receipts types.Receipts
}
type FinalizeFeed struct {
	txFeed event.Feed
	scope  event.SubscriptionScope
}

func (f *FinalizeFeed) SendFinalizeBlock(event NewFinalizeBlockEvent) {
	f.txFeed.Send(event)
}

func (f *FinalizeFeed) SubscribeNewFinalizeBlockEvent(ev chan NewFinalizeBlockEvent) event.Subscription {
	return f.scope.Track(f.txFeed.Subscribe(ev))
}

func (f *FinalizeFeed) Close() {
	f.scope.Close()
}
