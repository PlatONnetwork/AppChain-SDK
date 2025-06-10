package pevm

import (
	"sync"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/google/btree"
	cmap "github.com/orcaman/concurrent-map/v2"
)

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

func (l *LastLocations) RangeRead(f func(MemoryLocationHash, *ReadOrigins)) {
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

type ConcurrencyBTree struct {
	*btree.BTree
	sync.RWMutex
}

func NewConcurrentBTree() *ConcurrencyBTree {
	return &ConcurrencyBTree{
		BTree: btree.New(32),
	}
}

func (cbt *ConcurrencyBTree) AscendRange(lessThan btree.Item) *Iterator {
	cbt.RLock()
	defer cbt.RUnlock()
	items := make([]btree.Item, 0)
	cbt.AscendLessThan(lessThan, func(item btree.Item) bool {
		items = append(items, item)
		return true
	})
	return NewIterator(items)
}

type Iterator struct {
	items []btree.Item
	last  int
}

func NewIterator(items []btree.Item) *Iterator {
	return &Iterator{
		items: items,
		last:  len(items) - 1,
	}
}

func (it *Iterator) NextBack() btree.Item {
	if it.last < 0 {
		return nil
	}
	item := it.items[it.last]
	it.items = it.items[:it.last]
	it.last--
	return item
}

type MvMemory struct {
	data          cmap.ConcurrentMap[MemoryLocationHash, *ConcurrencyBTree]
	lastLocations []*LastLocations
	lazyAddresses *LazyAddresses
	newByteCodes  cmap.ConcurrentMap[common.Hash, []byte]
}

func NewMvMemory(
	blockSize int,
	estimatedLoactions map[MemoryLocationHash][]int32,
	lazyAddrs []common.Address) *MvMemory {
	m := &MvMemory{
		data:          cmap.NewWithCustomShardingFunction[MemoryLocationHash, *ConcurrencyBTree](memHashShard),
		lastLocations: initializeLastLocations(blockSize),
		lazyAddresses: NewLazyAddresses(),
		newByteCodes:  cmap.NewWithCustomShardingFunction[common.Hash, []byte](hashShard),
	}

	for h, txIdxs := range estimatedLoactions {
		tree := NewConcurrentBTree()
		for _, txIdx := range txIdxs {
			tree.BTree.ReplaceOrInsert(&item{
				TxIdx: txIdx,
				Entry: NewEstimate(),
			})
		}
		m.data.Set(h, tree)
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

	lastLocationIdx := 0
	for lastLocationIdx < len(lastLocation.write) {
		prevLocation := lastLocation.GetWrite(lastLocationIdx)
		// Remove old locations that aren't written to anymore.
		if _, found := writeSet.Find(prevLocation); !found {
			if writtenTxs, has := m.data.Get(prevLocation); has {
				writtenTxs.BTree.Delete(&item{
					TxIdx: txVersion.TxIdx,
				})
			}
			lastLocation.SwapRemoveWriteAt(lastLocationIdx)
		} else {
			lastLocationIdx++
		}
	}

	wroteNewLocation := false
	writeSet.Range(func(h MemoryLocationHash, value MemoryValue) {
		m.data.Upsert(h, nil, func(exist bool, valueInMap, newValue *ConcurrencyBTree) *ConcurrencyBTree {
			if !exist {
				valueInMap = NewConcurrentBTree()
			}
			valueInMap.BTree.ReplaceOrInsert(&item{
				TxIdx: txVersion.TxIdx,
				Entry: NewDataEntry(txVersion.TxIdx, value),
			})
			return valueInMap
		})

		if !lastLocation.HasWrite(h) {
			lastLocation.AppendWrite(h)
			wroteNewLocation = true
		}
	})
	return wroteNewLocation
}

func (m *MvMemory) ValidateReadLocations(txIdx int32) bool {
	validated := true
	lastLocation := m.lastLocations[txIdx]
	lastLocation.RangeRead(func(location MemoryLocationHash, priorOrigins *ReadOrigins) {
		if writtenTxs, found := m.data.Get(location); found {
			it := writtenTxs.AscendRange(&item{TxIdx: txIdx})

			priorOrigins.Range(func(priorOrigin ReadOrigin) bool {
				switch priorOrigin.(type) {
				case *Memory:
					priorVer := priorOrigin.(*Memory)
					entry := it.NextBack().(*item)
					dataEntry, isData := entry.Entry.(*DataEntry)
					if !isData {
						validated = false
						return false
					}
					if priorVer.Version.TxIdx != entry.TxIdx ||
						dataEntry.TxIncarnation != priorVer.Version.TxIncarnation {
						validated = false
						return false
					}
				case *Storage:
					if it.NextBack() != nil {
						validated = false
						return false
					}
				}
				return true
			})
		} else {
			_, ok := priorOrigins.Last().(*Storage)
			if priorOrigins.Len() != 1 || !ok {
				validated = false
			}
		}
	})
	return validated
}

func (m *MvMemory) ConvertWritesToEstimate(txIdx int32) {
	lastLocation := m.lastLocations[txIdx]
	lastLocation.RangeWrite(func(location MemoryLocationHash) {
		if writtenTxs, ok := m.data.Get(location); ok {
			writtenTxs.Lock()
			writtenTxs.BTree.ReplaceOrInsert(&item{
				TxIdx: txIdx,
				Entry: NewEstimate(),
			})
			writtenTxs.Unlock()
		}
	})
}

func (m *MvMemory) ConsumeLazyAddresses(f func(common.Address) bool) {
	m.lazyAddresses.Range(f)
}

func memHashShard(h MemoryLocationHash) uint32 {
	return uint32(h % 31)
}

func hashShard(h common.Hash) uint32 {
	return uint32(h[31]) % 31
}
