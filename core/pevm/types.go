package pevm

import (
	"bytes"
	"math/big"
	"sync"
	"unsafe"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/zeebo/xxh3"
)

var emptyCodeHash = common.Hash(crypto.Keccak256(nil))
var ripemd = common.HexToAddress("0000000000000000000000000000000000000003")
var codeHashSuffix = []byte("code_hash")

var (
	basicHashCache sync.Map // map[common.Address]MemoryLocationHash
	stateHashCache sync.Map // map[common.Address]map[string]MemoryLocationHash
	codeHashCache  sync.Map
)

var xxhashPool = sync.Pool{
	New: func() interface{} {
		h := xxh3.New()
		return h
	},
}

type IncarnationStatus uint8

const (
	ReadyToExecute IncarnationStatus = iota
	Executing
	Executed
	Validated
	Aborting
)

type FinishExecFlags uint8

const (
	NeedValidation FinishExecFlags = 1 << iota
	WroteNewLocation
)

func NewFinishExecFlags(flags ...FinishExecFlags) FinishExecFlags {
	var result FinishExecFlags
	for _, flag := range flags {
		result |= flag
	}
	return result
}

func (f *FinishExecFlags) Set(flags FinishExecFlags) {
	*f |= flags
}

func (f FinishExecFlags) Has(flags FinishExecFlags) bool {
	return (f & flags) == flags
}

type MemoryLocationHash = uint64

type TxStatus struct {
	Incarnation int32
	Status      IncarnationStatus
}

type TxVersion struct {
	TxIdx         int32
	TxIncarnation int32
}

func BasicLoc(addr common.Address) MemoryLocationHash {
	if hash, ok := basicHashCache.Load(addr); ok {
		return hash.(MemoryLocationHash)
	}

	hash := xxh3.Hash(addr[:])

	basicHashCache.Store(addr, hash)
	return hash
}

func CodeHashLoc(addr common.Address) MemoryLocationHash {
	if hash, ok := codeHashCache.Load(addr); ok {
		return hash.(MemoryLocationHash)
	}

	// 从池中获取hasher
	h := xxhashPool.Get().(*xxh3.Hasher)
	h.Reset()

	h.Write(addr[:])
	h.Write(codeHashSuffix[:])
	hash := h.Sum64()

	// 放回池中
	xxhashPool.Put(h)

	codeHashCache.Store(addr, hash)
	return hash
}

func StateLoc(addr common.Address, key []byte) MemoryLocationHash {
	// 创建复合键
	var k struct {
		addr common.Address
		key  string
	}
	k.addr = addr

	if len(key) > 0 {
		k.key = unsafe.String((*byte)(unsafe.Pointer(&key[0])), len(key))
	}

	// 尝试从全局缓存获取
	if hash, ok := stateHashCache.Load(k); ok {
		return hash.(MemoryLocationHash)
	}

	// 从池中获取hasher
	h := xxhashPool.Get().(*xxh3.Hasher)
	h.Reset()

	h.Write(addr[:])
	h.Write(key)
	hash := h.Sum64()

	// 放回池中
	xxhashPool.Put(h)

	// 存储到缓存
	stateHashCache.Store(k, hash)
	return hash
}

type MemoryValue interface {
	isMemoryValue()
}

type MemoryEntry interface {
	isMemoryEntry()
}

type memoryEntry struct{}

func (memoryEntry) isMemoryEntry() {}

type DataEntry struct {
	memoryEntry

	TxIncarnation int32
	Value         MemoryValue
}

func NewDataEntry(i int32, value MemoryValue) MemoryEntry {
	return &DataEntry{
		TxIncarnation: i,
		Value:         value,
	}
}

type EstimateMarker struct {
	memoryEntry

	_ struct{}
}

var estimateInstance = &EstimateMarker{}

func NewEstimate() MemoryEntry {
	return estimateInstance
}

type memoryValue struct{}

func (memoryValue) isMemoryValue() {}

type AccountBase struct {
	Addr     common.Address
	Nonce    uint64
	Balance  *big.Int
	CodeHash common.Hash
	CodeSize int
	Code     []byte
	Suicided bool
	NewCode  bool
}

func NewEmptyAccountBase(addr common.Address) *AccountBase {
	return &AccountBase{
		Addr:     addr,
		Balance:  common.Big0,
		CodeHash: common.Hash(emptyCodeHash),
	}
}

func (ab *AccountBase) Empty() bool {
	return ab.Nonce == 0 &&
		ab.Balance.Sign() == 0 &&
		bytes.Equal(ab.CodeHash[:], emptyCodeHash[:])
}

func (ab *AccountBase) Touch() bool {
	return ab.Addr == ripemd
}

func (ab *AccountBase) Clone() *AccountBase {
	return &AccountBase{
		Addr:     common.Address(bytes.Clone(ab.Addr[:])),
		Nonce:    ab.Nonce,
		Balance:  new(big.Int).Set(ab.Balance),
		CodeHash: common.Hash(bytes.Clone(ab.CodeHash[:])),
		CodeSize: ab.CodeSize,
		Code:     bytes.Clone(ab.Code),
		Suicided: ab.Suicided,
		NewCode:  ab.NewCode,
	}
}

type Basic struct {
	memoryValue

	Addr    common.Address
	Account *AccountBase
}

func NewBasic(addr common.Address, account *AccountBase) MemoryValue {
	return &Basic{
		Addr:    addr,
		Account: account,
	}
}

type LazySender struct {
	memoryValue

	Addr    common.Address
	Balance *big.Int
}

func NewLazySender(addr common.Address, balance *big.Int) MemoryValue {
	return &LazySender{
		Addr:    addr,
		Balance: balance,
	}
}

type LazyRecipient struct {
	memoryValue

	Addr    common.Address
	Balance *big.Int
}

func NewLazyRecipient(addr common.Address, balance *big.Int) MemoryValue {
	return &LazyRecipient{
		Addr:    addr,
		Balance: balance,
	}
}

type State struct {
	memoryValue

	Addr  common.Address
	Key   []byte
	Value []byte
}

func NewState(addr common.Address, key, value []byte) MemoryValue {
	return &State{
		Addr:  addr,
		Key:   key,
		Value: value,
	}
}

type CodeHash struct {
	memoryValue

	Addr     common.Address
	CodeHash common.Hash
}

func NewCodeHash(addr common.Address, codeHash common.Hash) MemoryValue {
	return &CodeHash{
		Addr:     addr,
		CodeHash: codeHash,
	}
}

type SelfDestructed struct {
	memoryValue

	Addr common.Address
}

func NewSelfDestructed(addr common.Address) MemoryValue {
	return &SelfDestructed{
		Addr: addr,
	}
}

type ReadOrigin interface {
	isReadOrigin()
}

type readOrigin struct{}

func (readOrigin) isReadOrigin() {}

type Memory struct {
	readOrigin

	Version TxVersion
}

func NewMemory(ver TxVersion) ReadOrigin { return &Memory{Version: ver} }

type Storage struct {
	readOrigin
}

var storageInstance = &Storage{}

func NewStorage() ReadOrigin { return storageInstance }

type ReadOrigins struct {
	origins []ReadOrigin
}

func NewReadOrigins() *ReadOrigins {
	return &ReadOrigins{
		origins: make([]ReadOrigin, 0),
	}
}

func (ro *ReadOrigins) Len() int {
	return len(ro.origins)
}

func (ro *ReadOrigins) Get(idx int) ReadOrigin {
	if ro.Len() >= 0 && idx <= (ro.Len()-1) {
		return ro.origins[idx]
	}
	return nil
}

func (ro *ReadOrigins) Push(origin ReadOrigin) {
	ro.origins = append(ro.origins, origin)
}

func (ro *ReadOrigins) Last() ReadOrigin {
	last := len(ro.origins) - 1
	if last >= 0 {
		return ro.origins[last]
	}
	return nil
}

func (ro *ReadOrigins) Range(f func(ReadOrigin) bool) {
	for _, origin := range ro.origins {
		if !f(origin) {
			break
		}
	}
}

type ReadSet struct {
	readOrigins map[MemoryLocationHash]*ReadOrigins
}

func NewReadSet() *ReadSet {
	return &ReadSet{
		readOrigins: make(map[MemoryLocationHash]*ReadOrigins),
	}
}

func (rs *ReadSet) GetOrDefault(locationHash MemoryLocationHash) *ReadOrigins {
	ro, exist := rs.readOrigins[locationHash]
	if !exist || ro == nil {
		ro = NewReadOrigins()
		rs.readOrigins[locationHash] = ro
	}

	return ro
}

func (rs *ReadSet) Set(locationHash MemoryLocationHash, ro *ReadOrigins) {
	rs.readOrigins[locationHash] = ro
}

func (rs *ReadSet) Range(f func(MemoryLocationHash, *ReadOrigins) bool) {
	for h, ro := range rs.readOrigins {
		if !f(h, ro) {
			break
		}
	}
}

type WriteEntry struct {
	Hash  MemoryLocationHash
	Value MemoryValue
}

type WriteSet []WriteEntry

func NewWriteSet() WriteSet {
	return make(WriteSet, 0)
}

func (ws *WriteSet) Add(hash MemoryLocationHash, value MemoryValue) {
	*ws = append(*ws, WriteEntry{Hash: hash, Value: value})
}

func (ws *WriteSet) Find(hash MemoryLocationHash) (MemoryValue, bool) {
	for _, entry := range *ws {
		if entry.Hash == hash {
			return entry.Value, true
		}
	}
	return nil, false
}

func (ws *WriteSet) Range(f func(MemoryLocationHash, MemoryValue)) {
	for _, entry := range *ws {
		f(entry.Hash, entry.Value)
	}
}

type Task interface {
	isTask()
	Version() TxVersion
}

type Execution struct {
	version TxVersion
}

func NewExection(version TxVersion) Task {
	return &Execution{
		version: version,
	}
}

func (e *Execution) isTask()            {}
func (e *Execution) Version() TxVersion { return e.version }

type Validation struct {
	version TxVersion
}

func NewValidation(version TxVersion) Task {
	return &Validation{
		version: version,
	}
}

func (v *Validation) isTask()            {}
func (v *Validation) Version() TxVersion { return v.version }
