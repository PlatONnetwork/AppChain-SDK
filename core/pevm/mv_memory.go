package pevm

import (
	"sort"
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

const writeHistoryShards = 64

type WriteHistoryShard struct {
	histories map[MemoryLocationHash]*WriteHistory
	mu        sync.RWMutex
}

func NewWriteHistoryShard() *WriteHistoryShard {
	return &WriteHistoryShard{
		histories: make(map[MemoryLocationHash]*WriteHistory),
	}
}

func (s *WriteHistoryShard) GetOrCreate(loc MemoryLocationHash) *WriteHistory {
	s.mu.RLock()
	if wh, ok := s.histories[loc]; ok {
		s.mu.RUnlock()
		return wh
	}
	s.mu.RUnlock()

	s.mu.Lock()
	defer s.mu.Unlock()

	if wh, ok := s.histories[loc]; ok {
		return wh
	}

	wh := NewWriteHistory()
	s.histories[loc] = wh
	return wh
}

func (s *WriteHistoryShard) Get(loc MemoryLocationHash) *WriteHistory {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.histories[loc]
}

func (s *WriteHistoryShard) Has(loc MemoryLocationHash) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.histories[loc]
	return ok
}

func (s *WriteHistoryShard) Release() {
	for _, wh := range s.histories {
		wh.Release()
	}
}

type ShardedWriteHistory struct {
	shards [writeHistoryShards]*WriteHistoryShard
}

func NewShardedWriteHistory() *ShardedWriteHistory {
	s := &ShardedWriteHistory{}
	for i := range s.shards {
		s.shards[i] = NewWriteHistoryShard()
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
		shard.Release()
	}
}

type WriteHistory struct {
	// Ascend sorted
	items []*item
	mu    sync.RWMutex
}

func NewWriteHistory() *WriteHistory {
	return &WriteHistory{
		items: make([]*item, 0, 128),
	}
}

// 优化 AscendRange 方法
func (wh *WriteHistory) AscendRange(txIdx int32) *ItemIterator {
	wh.mu.RLock()
	defer wh.mu.RUnlock()

	n := len(wh.items)
	if n == 0 {
		return &ItemIterator{items: nil}
	}

	// 二分查找起始位置（第一个 >= txIdx 的索引）
	start := sort.Search(n, func(i int) bool {
		return wh.items[i].TxIdx < txIdx
	})

	// 返回从 start 到结尾的切片（降序排列）
	return &ItemIterator{
		items: wh.items[start:],
		index: 0,
	}
}

func (wh *WriteHistory) Ascend(f func(*item) bool) {
	wh.mu.RLock()
	defer wh.mu.RUnlock()

	i := len(wh.items) - 1
	for ; i >= 0; i-- {
		if !f(wh.items[i]) {
			break
		}
	}
}

func (wh *WriteHistory) ReplaceOrInsert(entry *item) {
	wh.mu.Lock()
	defer wh.mu.Unlock()

	n := len(wh.items)

	// 快速路径：空列表或新项最大
	if n == 0 || entry.TxIdx > wh.items[0].TxIdx {
		// 在开头插入
		wh.items = append([]*item{entry}, wh.items...)
		return
	}

	// 快速路径：新项最小
	if entry.TxIdx < wh.items[n-1].TxIdx {
		wh.items = append(wh.items, entry)
		return
	}

	// 二分查找插入位置（降序）
	idx := sort.Search(n, func(i int) bool {
		return wh.items[i].TxIdx <= entry.TxIdx
	})

	if idx < n && wh.items[idx].TxIdx == entry.TxIdx {
		// 替换
		putItem(wh.items[idx])
		wh.items[idx] = entry
	} else {
		// 插入
		wh.items = append(wh.items, nil)
		if idx < n {
			copy(wh.items[idx+1:], wh.items[idx:])
		}
		wh.items[idx] = entry
	}
}

func (wh *WriteHistory) Delete(txIdx int32) {
	wh.mu.Lock()
	defer wh.mu.Unlock()

	idx := sort.Search(len(wh.items), func(i int) bool {
		return wh.items[i].TxIdx <= txIdx
	})

	if idx < len(wh.items) && wh.items[idx].TxIdx == txIdx {
		putItem(wh.items[idx])
		copy(wh.items[idx:], wh.items[idx+1:])
		wh.items = wh.items[:len(wh.items)-1]
	}
}

func (wh *WriteHistory) Release() {
	for _, entry := range wh.items {
		cpy := entry
		putItem(cpy)
	}
}

type ItemIterator struct {
	items []*item
	index int
}

func (it *ItemIterator) NextBack() *item {
	if it.index >= len(it.items) {
		return nil
	}
	item := it.items[it.index]
	it.index++
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

	for _, entry := range writeSet {
		h := entry.Hash
		value := entry.Value

		wh := m.data.GetOrCreate(h)
		entry := getItem()
		entry.TxIdx = txVersion.TxIdx
		entry.Entry = NewDataEntry(txVersion.TxIncarnation, value)
		wh.ReplaceOrInsert(entry)

		foundInWriteSet = false
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
	const batchSize = 8 // CPU缓存行大小

	lastLocation := m.lastLocations[txIdx]

	locations := make([]MemoryLocationHash, 0, 32)
	originsMap := make([]*ReadOrigins, 0, 32)

	lastLocation.RangeRead(func(loc MemoryLocationHash, priorOrigins *ReadOrigins) bool {
		locations = append(locations, loc)
		originsMap = append(originsMap, priorOrigins)
		return true
	})

	// 批处理验证 (8个一组)
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
			switch po := priorOrigin.(type) {
			case *Memory:
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
			case *Storage:
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
