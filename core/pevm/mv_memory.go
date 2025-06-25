package pevm

import (
	"sync"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/google/btree"
	cmap "github.com/orcaman/concurrent-map/v2"
)

var itemPool = sync.Pool{
	New: func() interface{} {
		return new(item)
	},
}

func getItem() *item {
	return itemPool.Get().(*item)
}

func putItem(it *item) {
	it.TxIdx = 0
	it.Entry = nil
	itemPool.Put(it)
}

const writeHistoryShards = 256

var shardPool = sync.Pool{
	New: func() interface{} {
		return NewWriteHistoryShard()
	},
}

func getShard() *WriteHistoryShard {
	return shardPool.Get().(*WriteHistoryShard)
}

func putShard(s *WriteHistoryShard) {
	// 清理 shard 的状态，避免数据污染
	for _, wh := range s.histories {
		wh.wh.Release()
	}
	s.histories = make(map[MemoryLocationHash]*writeHistoryEntry)
	shardPool.Put(s)
}

type writeHistoryEntry struct {
	wh   *WriteHistory
	once sync.Once
}

func (e *writeHistoryEntry) GetOrCreate() *WriteHistory {
	e.once.Do(func() {
		e.wh = NewWriteHistory()
	})
	return e.wh
}

type WriteHistoryShard struct {
	histories map[MemoryLocationHash]*writeHistoryEntry
	mu        sync.RWMutex
}

func NewWriteHistoryShard() *WriteHistoryShard {
	return &WriteHistoryShard{
		histories: make(map[MemoryLocationHash]*writeHistoryEntry),
	}
}

func (s *WriteHistoryShard) GetOrCreate(loc MemoryLocationHash) *WriteHistory {
	s.mu.RLock()
	if entry, ok := s.histories[loc]; ok {
		s.mu.RUnlock()
		return entry.wh
	}
	s.mu.RUnlock()

	s.mu.Lock()
	defer s.mu.Unlock()

	if entry, ok := s.histories[loc]; ok {
		return entry.GetOrCreate()
	}

	entry := &writeHistoryEntry{}
	s.histories[loc] = entry
	return entry.GetOrCreate()
}

func (s *WriteHistoryShard) Get(loc MemoryLocationHash) *WriteHistory {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if entry, ok := s.histories[loc]; ok {
		return entry.wh
	}
	return nil
}

func (s *WriteHistoryShard) Has(loc MemoryLocationHash) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.histories[loc]
	return ok
}

func (s *WriteHistoryShard) Release() {
	for _, entry := range s.histories {
		entry.wh.Release()
	}
}

type ShardedWriteHistory struct {
	shards [writeHistoryShards]*WriteHistoryShard
}

func NewShardedWriteHistory() *ShardedWriteHistory {
	s := &ShardedWriteHistory{}
	for i := range s.shards {
		s.shards[i] = getShard()
	}
	return s
}

func (s *ShardedWriteHistory) GetOrCreate(loc MemoryLocationHash) *WriteHistory {
	shardIdx := loc % writeHistoryShards
	return s.shards[shardIdx].GetOrCreate(loc)
}

func (s *ShardedWriteHistory) Get(loc MemoryLocationHash) *WriteHistory {
	shardIdx := loc % writeHistoryShards
	return s.shards[shardIdx].Get(loc)
}

func (s *ShardedWriteHistory) Has(loc MemoryLocationHash) bool {
	shardIdx := loc % writeHistoryShards
	return s.shards[shardIdx].Has(loc)
}

func (s *ShardedWriteHistory) Release() {
	for _, shard := range s.shards {
		putShard(shard)
	}
}

type WriteHistory struct {
	items *btree.BTree
	mu    sync.RWMutex
}

func NewWriteHistory() *WriteHistory {
	return &WriteHistory{
		items: btree.New(2),
	}
}

func (wh *WriteHistory) AscendRange(txIdx int32) *ItemIterator {
	wh.mu.RLock()
	defer wh.mu.RUnlock()

	var items []*item
	wh.items.AscendLessThan(&item{TxIdx: txIdx}, func(entry btree.Item) bool {
		items = append(items, entry.(*item))
		return true
	})

	return &ItemIterator{
		items: items,
		index: len(items) - 1,
	}
}

func (wh *WriteHistory) Ascend(f func(*item) bool) {
	wh.mu.RLock()
	defer wh.mu.RUnlock()

	wh.items.Ascend(func(entry btree.Item) bool {
		return f(entry.(*item))
	})
}

func (wh *WriteHistory) ReplaceOrInsert(entry *item) {
	wh.mu.Lock()
	defer wh.mu.Unlock()

	old := wh.items.ReplaceOrInsert(entry)
	if old != nil {
		putItem(old.(*item))
	}
}

func (wh *WriteHistory) Delete(txIdx int32) {
	wh.mu.Lock()
	defer wh.mu.Unlock()

	entry := wh.items.Delete(&item{TxIdx: txIdx})
	if entry != nil {
		putItem(entry.(*item))
	}
}

func (wh *WriteHistory) Release() {
	wh.items.Ascend(func(entry btree.Item) bool {
		putItem(entry.(*item))
		return true
	})
	wh.items.Clear(false)
}

type ItemIterator struct {
	items []*item
	index int
}

func (it *ItemIterator) NextBack() *item {
	if it.index < 0 {
		return nil
	}
	item := it.items[it.index]
	it.index--
	return item
}

type LastLocations struct {
	sync.RWMutex

	read  *ReadSet
	write []MemoryLocationHash
}

func (l *LastLocations) SetRead(rs *ReadSet) {
	l.Lock()
	defer l.Unlock()

	l.read = rs
}

func (l *LastLocations) RangeRead(f func(MemoryLocationHash, *ReadOrigins) bool) {
	l.RLock()
	defer l.RUnlock()

	l.read.Range(f)
}

func (l *LastLocations) GetWrite(idx int) MemoryLocationHash {
	l.RLock()
	defer l.RUnlock()
	return l.write[idx]
}

func (l *LastLocations) RangeWrite(f func(MemoryLocationHash)) {
	l.RLock()
	defer l.RUnlock()

	for _, h := range l.write {
		f(h)
	}
}

func (l *LastLocations) SwapRemoveWriteAt(idx int) {
	l.Lock()
	defer l.Unlock()

	last := len(l.write) - 1
	l.write[idx], l.write[last] = l.write[last], l.write[idx]
	l.write = l.write[:last]
}

func (l *LastLocations) HasWrite(h MemoryLocationHash) bool {
	l.RLock()
	defer l.RUnlock()

	found := false
	for _, hh := range l.write {
		if hh == h {
			found = true
			break
		}
	}
	return found
}

func (l *LastLocations) AppendWrite(h MemoryLocationHash) {
	l.Lock()
	defer l.Unlock()

	l.write = append(l.write, h)
}

func newLastLocations() *LastLocations {
	return &LastLocations{
		read:  NewReadSet(),
		write: make([]MemoryLocationHash, 0),
	}
}

func initializeLastLocations(blockSize int) []*LastLocations {
	ll := make([]*LastLocations, blockSize)
	for i := range ll {
		ll[i] = newLastLocations()
	}
	return ll
}

type LazyAddresses struct {
	sync.RWMutex

	addresses map[common.Address]struct{}
}

func NewLazyAddresses() *LazyAddresses {
	return &LazyAddresses{
		addresses: make(map[common.Address]struct{}, 0),
	}
}

func (la *LazyAddresses) Insert(addr common.Address) {
	la.Lock()
	defer la.Unlock()
	la.addresses[addr] = struct{}{}
}

func (la *LazyAddresses) Remove(addr common.Address) {
	la.Lock()
	defer la.Unlock()
	delete(la.addresses, addr)
}

func (la *LazyAddresses) Has(addr common.Address) bool {
	la.RLock()
	defer la.RUnlock()
	_, exist := la.addresses[addr]
	return exist
}

func (la *LazyAddresses) Range(f func(common.Address) bool) {
	la.RLock()
	defer la.RLock()

	for addr, _ := range la.addresses {
		if !f(addr) {
			break
		}
	}
}

type item struct {
	TxIdx int32
	Entry MemoryEntry
}

func (it item) Less(rh btree.Item) bool {
	return it.TxIdx < rh.(*item).TxIdx
}

type MvMemory struct {
	data          *ShardedWriteHistory
	lastLocations []*LastLocations
	lazyAddresses *LazyAddresses
	newByteCodes  cmap.ConcurrentMap[common.Hash, []byte]
}

func NewMvMemory(
	blockSize int,
	estimatedLoactions map[MemoryLocationHash][]int32,
	lazyAddrs []common.Address) *MvMemory {
	m := &MvMemory{
		data:          NewShardedWriteHistory(),
		lastLocations: initializeLastLocations(blockSize),
		lazyAddresses: NewLazyAddresses(),
		newByteCodes:  cmap.NewWithCustomShardingFunction[common.Hash, []byte](hashShard),
	}

	for h, txIdxs := range estimatedLoactions {
		his := m.data.GetOrCreate(h)
		for _, txIdx := range txIdxs {
			his.ReplaceOrInsert(&item{
				TxIdx: txIdx,
				Entry: NewEstimate(),
			})
		}
	}
	if len(lazyAddrs) > 0 {
		m.AddLazyAddresses(lazyAddrs)
	}
	return m
}

func (m *MvMemory) AddLazyAddresses(addrs []common.Address) {
	for _, addr := range addrs {
		m.lazyAddresses.Insert(addr)
	}
}

func (m *MvMemory) Record(txVersion *TxVersion, readSet *ReadSet, writeSet WriteSet) bool {
	lastLocation := m.lastLocations[txVersion.TxIdx]
	lastLocation.SetRead(readSet)

	oldWrites := lastLocation.write
	newWrites := oldWrites[:0]

	wroteNewLocation := false
	foundInWriteSet := false

	for _, loc := range oldWrites {
		if _, found := writeSet.Find(loc); found {
			newWrites = append(newWrites, loc)
		} else {
			if wh := m.data.Get(loc); wh != nil {
				wh.Delete(txVersion.TxIdx)
			}
		}
	}

	// Create a map to store WriteHistory instances for each location
	writeHistories := make(map[MemoryLocationHash]*WriteHistory)

	// First, get or create all the WriteHistory instances
	for _, entry := range writeSet {
		h := entry.Hash
		if _, ok := writeHistories[h]; !ok {
			writeHistories[h] = m.data.GetOrCreate(h)
		}
	}

	// Now, iterate through the writeSet and insert the entries
	for _, entry := range writeSet {
		h := entry.Hash
		value := entry.Value

		wh := writeHistories[h]
		entry := getItem()
		entry.TxIdx = txVersion.TxIdx
		entry.Entry = NewDataEntry(txVersion.TxIncarnation, value)
		wh.ReplaceOrInsert(entry)

		for _, existing := range newWrites {
			if existing == h {
				foundInWriteSet = true
				break
			}
		}

		if !foundInWriteSet {
			newWrites = append(newWrites, h)
			wroteNewLocation = true
		}
	}

	lastLocation.write = newWrites
	return wroteNewLocation
}

func (m *MvMemory) ValidateReadLocations(txIdx int32) bool {
	const batchSize = 8

	lastLocation := m.lastLocations[txIdx]

	locations := make([]MemoryLocationHash, 0, 32)
	originsMap := make([]*ReadOrigins, 0, 32)

	lastLocation.RangeRead(func(loc MemoryLocationHash, priorOrigins *ReadOrigins) bool {
		locations = append(locations, loc)
		originsMap = append(originsMap, priorOrigins)
		return true
	})

	valid := true
	for i := 0; i < len(locations); i += batchSize {
		end := i + batchSize
		if end > len(locations) {
			end = len(locations)
		}

		batchValid := m.validateBatch(txIdx, locations[i:end], originsMap[i:end])
		valid = valid && batchValid
		if !valid {
			break
		}
	}
	return valid
}

func (m *MvMemory) validateBatch(txIdx int32, locs []MemoryLocationHash, priorOrigins []*ReadOrigins) bool {
	valid := true
	for i, loc := range locs {
		wh := m.data.Get(loc)
		origins := priorOrigins[i]
		if wh == nil {
			_, ok := origins.Last().(*Storage)
			return origins.Len() == 1 && ok
		}

		it := wh.AscendRange(txIdx)

		origins.Range(func(priorOrigin ReadOrigin) bool {
			if po, ok := priorOrigin.(*Memory); ok {
				entry := it.NextBack()
				if entry == nil {
					valid = false
					return false
				}
				if de, ok := entry.Entry.(*DataEntry); ok {
					if po.Version.TxIdx != entry.TxIdx || de.TxIncarnation != po.Version.TxIncarnation {
						valid = false
						return false
					}
				} else {
					valid = false
					return false
				}
			} else if _, ok := priorOrigin.(*Storage); ok {
				if it.NextBack() != nil {
					valid = false
					return false
				}
			}
			return true
		})
		if !valid {
			break
		}
	}

	return valid
}

func (m *MvMemory) ConvertWritesToEstimate(txIdx int32) {
	lastLocation := m.lastLocations[txIdx]
	lastLocation.RangeWrite(func(location MemoryLocationHash) {
		if wh := m.data.Get(location); wh != nil {
			wh.ReplaceOrInsert(&item{
				TxIdx: txIdx,
				Entry: NewEstimate(),
			})
		}
	})
}

func (m *MvMemory) ConsumeLazyAddresses(f func(common.Address) bool) {
	m.lazyAddresses.Range(f)
}

func (m *MvMemory) Release() {
	m.data.Release()
}

func memHashShard(h MemoryLocationHash) uint32 {
	return uint32(h % 31)
}

func hashShard(h common.Hash) uint32 {
	return uint32(h[31]) % 31
}
