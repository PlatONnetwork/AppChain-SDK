package pevm

import (
	"sync"

	"github.com/PlatONnetwork/AppChain-SDK/x/pevm/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/google/btree"
	cmap "github.com/orcaman/concurrent-map/v2"
)

type LastLocations struct {
	sync.RWMutex

	read  types.ReadSet
	write []types.MemoryLocationHash
}

func (l *LastLocations) SetRead(rs types.ReadSet) {
	l.Lock()
	defer l.Unlock()

	l.read = rs
}

func (l *LastLocations) RangeRead(f func(types.MemoryLocationHash, types.ReadOrigins)) {
	l.RLock()
	defer l.RUnlock()

	for h, ro := range l.read {
		f(h, ro)
	}
}

func (l *LastLocations) GetWrite(idx int) types.MemoryLocationHash {
	l.RLock()
	defer l.RUnlock()
	return l.write[idx]
}

func (l *LastLocations) RangeWrite(f func(types.MemoryLocationHash)) {
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

func (l *LastLocations) HasWrite(h types.MemoryLocationHash) bool {
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

func (l *LastLocations) AppendWrite(h types.MemoryLocationHash) {
	l.Lock()
	defer l.Unlock()

	l.write = append(l.write, h)
}

func newLastLocations() *LastLocations {
	return &LastLocations{
		read:  make(types.ReadSet, 0),
		write: make([]types.MemoryLocationHash, 0),
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

type DataEntry struct {
	TxIdx uint32
	Entry types.MemoryEntry
}

func (e DataEntry) Less(rh btree.Item) bool {
	return e.TxIdx < rh.(DataEntry).TxIdx
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
	data          cmap.ConcurrentMap[types.MemoryLocationHash, *ConcurrencyBTree]
	lastLocations []*LastLocations
	lazyAddresses *LazyAddresses
	newByteCodes  cmap.ConcurrentMap[common.Hash, []byte]
}

func NewMvMemory(
	blockSize int,
	estimatedLoactions map[types.MemoryLocationHash][]uint32,
	lazyAddrs []common.Address) *MvMemory {
	m := &MvMemory{
		data:          cmap.NewWithCustomShardingFunction[types.MemoryLocationHash, *ConcurrencyBTree](memHashShard),
		lastLocations: initializeLastLocations(blockSize),
		lazyAddresses: NewLazyAddresses(),
		newByteCodes:  cmap.NewWithCustomShardingFunction[common.Hash, []byte](hashShard),
	}

	for h, txIdxs := range estimatedLoactions {
		tree := NewConcurrentBTree()
		for _, txIdx := range txIdxs {
			tree.BTree.ReplaceOrInsert(&DataEntry{
				TxIdx: txIdx,
				Entry: types.NewEstimate(),
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

func (m *MvMemory) Record(txVersion *types.TxVersion, readSet types.ReadSet, writeSet types.WriteSet) bool {
	lastLocation := m.lastLocations[txVersion.TxIdx]
	lastLocation.SetRead(readSet)

	lastLocationIdx := 0
	for lastLocationIdx < len(lastLocation.write) {
		prevLocation := lastLocation.GetWrite(lastLocationIdx)
		// Remove old locations that aren't written to anymore.
		if _, found := writeSet.Find(prevLocation); !found {
			if writtenTxs, has := m.data.Get(prevLocation); has {
				writtenTxs.BTree.Delete(&DataEntry{
					TxIdx: txVersion.TxIdx,
				})
			}
			lastLocation.SwapRemoveWriteAt(lastLocationIdx)
		} else {
			lastLocationIdx++
		}
	}

	wroteNewLocation := false
	writeSet.Range(func(h types.MemoryLocationHash, value types.MemoryValue) {
		m.data.Upsert(h, nil, func(exist bool, valueInMap, newValue *ConcurrencyBTree) *ConcurrencyBTree {
			if !exist {
				valueInMap = NewConcurrentBTree()
			}
			valueInMap.BTree.ReplaceOrInsert(&DataEntry{
				TxIdx: txVersion.TxIdx,
				Entry: types.NewDataEntry(txVersion.TxIdx, value),
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

func (m *MvMemory) ValidateReadLocations(txIdx uint32) bool {
	validated := true
	lastLocation := m.lastLocations[txIdx]
	lastLocation.RangeRead(func(location types.MemoryLocationHash, priorOrigins types.ReadOrigins) {
		if writtenTxs, found := m.data.Get(location); found {
			it := writtenTxs.AscendRange(&DataEntry{TxIdx: txIdx})

			for _, priorOrign := range priorOrigins {
				switch priorOrign.(type) {
				case *types.MvMemory:
					priordVer := priorOrign.(*types.MvMemory)
					entry := it.NextBack().(*DataEntry)
					dataEntry, isData := entry.Entry.(*types.DataEntry)
					if !isData {
						validated = false
						break
					}
					if priordVer.Version.TxIdx != entry.TxIdx ||
						dataEntry.TxIncarnation != priordVer.Version.TxIncarnation {
						validated = false
						break
					}
				case *types.Storage:
					if it.NextBack() != nil {
						validated = false
						break
					}
				}
			}
		} else {
			_, ok := priorOrigins[len(priorOrigins)-1].(*types.Storage)
			if len(priorOrigins) != 1 || !ok {
				validated = false
			}
		}
	})
	return validated
}

func (m *MvMemory) ConvertWritesToEstimate(txIdx uint32) {
	lastLocation := m.lastLocations[txIdx]
	lastLocation.RangeWrite(func(location types.MemoryLocationHash) {
		if writtenTxs, ok := m.data.Get(location); ok {
			writtenTxs.Lock()
			writtenTxs.BTree.ReplaceOrInsert(&DataEntry{
				TxIdx: txIdx,
				Entry: types.NewEstimate(),
			})
			writtenTxs.Unlock()
		}
	})
}

func (m *MvMemory) ConsumeLazyAddresses(f func(common.Address) bool) {
	m.lazyAddresses.Range(f)
}

func memHashShard(h types.MemoryLocationHash) uint32 {
	return uint32(h[31]) % 31
}

func hashShard(h common.Hash) uint32 {
	return uint32(h[31]) % 31
}
