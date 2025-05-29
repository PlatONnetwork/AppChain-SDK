package types

import (
	"bytes"
	"math/big"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
)

var emptyCodeHash = crypto.Keccak256(nil)
var ripemd = common.HexToAddress("0000000000000000000000000000000000000003")

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

type MemoryLocationHash = common.Hash

type TxStatus struct {
	Incarnation uint32
	Status      IncarnationStatus
}

type TxVersion struct {
	TxIdx         uint32
	TxIncarnation uint32
}

func BasicLoc(addr common.Address) MemoryLocationHash { return crypto.Keccak256Hash(addr[:]) }
func CodeHashLoc(addr common.Address) MemoryLocationHash {
	return crypto.Keccak256Hash(addr[:], []byte("code_hash"))
}
func StateLoc(addr common.Address, key []byte) MemoryLocationHash {
	return crypto.Keccak256Hash(addr[:], key)
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

	TxIncarnation uint32
	Value         MemoryValue
}

func NewDataEntry(i uint32, value MemoryValue) MemoryEntry {
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
		bytes.Equal(ab.CodeHash[:], emptyCodeHash)
}

func (ab *AccountBase) Touch() bool {
	return ab.Addr == ripemd
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

type MvMemory struct {
	readOrigin

	Version TxVersion
}

func NewMvMemory(ver TxVersion) ReadOrigin { return &MvMemory{Version: ver} }

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

type ReadSet = map[MemoryLocationHash]*ReadOrigins

type WriteEntry struct {
	Hash  MemoryLocationHash
	Value MemoryValue
}

type WriteSet []WriteEntry

func NewWriteSet(initialSize int) WriteSet {
	return make(WriteSet, initialSize)
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
