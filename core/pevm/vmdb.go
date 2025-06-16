package pevm

import (
	"fmt"
	"math"
	"math/big"
	"reflect"

	"github.com/PlatONnetwork/PlatON-Go/common"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

var (
	_ sdk.StateDB = (*VmDB)(nil)
)

type VmDB struct {
	vm           *Vm
	txIdx        int32
	tx           *coretypes.Transaction
	fromAddr     common.Address
	fromHash     MemoryLocationHash
	toHash       MemoryLocationHash
	toCodeHash   common.Hash
	isLazy       bool
	readSet      *ReadSet
	readAccounts map[MemoryLocationHash]*AccountBase
	dirties      map[common.Address]struct{}
	states       map[common.Address]map[string][]byte
	readStates   map[common.Address]map[string][]byte
	addBalances  map[common.Address]*big.Int
	subBalances  map[common.Address]*big.Int

	refund     uint64
	logs       map[common.Hash][]*coretypes.Log
	logSize    uint
	accessList *accessList
}

func NewVmDB(
	vm *Vm,
	txIdx int32,
	tx *coretypes.Transaction,
	fromAddr common.Address,
	fromHash, toHash MemoryLocationHash) *VmDB {
	db := &VmDB{
		vm:           vm,
		txIdx:        txIdx,
		tx:           tx,
		fromAddr:     fromAddr,
		fromHash:     fromHash,
		toHash:       toHash,
		readSet:      NewReadSet(),
		readAccounts: make(map[MemoryLocationHash]*AccountBase),
		dirties:      make(map[common.Address]struct{}),
		states:       make(map[common.Address]map[string][]byte),
		readStates:   make(map[common.Address]map[string][]byte),
		addBalances:  make(map[common.Address]*big.Int),
		subBalances:  make(map[common.Address]*big.Int),
		refund:       0,
		logs:         make(map[common.Hash][]*coretypes.Log, 0),
		accessList:   newAccessList(),
	}
	if tx.To() != nil {
		db.toCodeHash = vm.statedb.GetCodeHash(*tx.To())
		db.isLazy = db.toCodeHash == emptyCodeHash &&
			(vm.mvMemory.data.Has(fromHash) || vm.mvMemory.data.Has(toHash))
	}
	return db
}

func (db *VmDB) pushOrigin(readOrigins *ReadOrigins, readOrigin ReadOrigin) {
	if readOrigins.Len() > 0 {
		last := readOrigins.Last()
		if !reflect.DeepEqual(last, readOrigin) {
			panic(ErrInconsistentRead)
		}
	}
	readOrigins.Push(readOrigin)
}

func (db *VmDB) hashBasic(addr common.Address) MemoryLocationHash {
	if addr == db.fromAddr {
		return db.fromHash
	}
	if db.tx.To() != nil && addr == *db.tx.To() {
		return db.toHash
	}
	return BasicLoc(addr)
}

func (db *VmDB) getAccountBasic(addr common.Address) *AccountBase {
	var (
		locationHash           = db.hashBasic(addr)
		readOrigins            = db.readSet.GetOrDefault(locationHash)
		hasPrevOrigins         = readOrigins.Len() > 0
		newOrigins             = NewReadOrigins()
		balanceAddition        = new(big.Int)
		nonceAddtion    uint64 = 0
		finalAccount    *AccountBase
	)

	if db.isLazy {
		if locationHash == db.fromHash || locationHash == db.toHash {
			return nil
		}
	}

	if basic, ok := db.readAccounts[locationHash]; ok {
		return basic
	}

	if db.txIdx > 0 {
		if writtenTxs, ok := db.vm.mvMemory.data.Get(locationHash); ok {
			it := writtenTxs.AscendRange(&item{TxIdx: db.txIdx})
		itLoop:
			for {
				entry := it.NextBack()
				if entry == nil {
					break itLoop
				}
				entryItem := entry.(*item)
				switch entryItem.Entry.(type) {
				case *DataEntry:
					de := entryItem.Entry.(*DataEntry)

					// About to push a new origin
					// Inconsistent: new origin will be longer than the previous!
					if hasPrevOrigins && readOrigins.Len() == newOrigins.Len() {
						panic(ErrInconsistentRead)
					}

					origin := NewMemory(TxVersion{
						TxIdx:         entryItem.TxIdx,
						TxIncarnation: de.TxIncarnation,
					})
					if hasPrevOrigins {
						if !reflect.DeepEqual(origin, readOrigins.Get(newOrigins.Len())) {
							panic(ErrInconsistentRead)
						}
					} else {
						newOrigins.Push(origin)
					}
					switch de.Value.(type) {
					case *Basic:
						basic := de.Value.(*Basic)
						finalAccount = basic.Account
						break itLoop
					case *LazySender:
						lazySender := de.Value.(*LazySender)
						balanceAddition = new(big.Int).Sub(balanceAddition, lazySender.Balance)
						nonceAddtion += 1
					case *LazyRecipient:
						lazyRecipient := de.Value.(*LazyRecipient)
						balanceAddition = new(big.Int).Add(balanceAddition, lazyRecipient.Balance)
					default:
						panic(ErrInvalidMemoryValueType)
					}
				case *EstimateMarker:
					panic(BlockingError{Addr: addr, TxIdx: entryItem.TxIdx})
				}

			}
		}
	}

	if finalAccount == nil {
		if !hasPrevOrigins {
			newOrigins.Push(NewStorage())
		} else if readOrigins.Len() != newOrigins.Len()+1 ||
			!reflect.DeepEqual(readOrigins.Last(), NewStorage()) {
			panic(ErrInconsistentRead)
		}
		finalAccount = &AccountBase{
			Addr:     addr,
			Nonce:    db.vm.statedb.GetNonce(addr),
			Balance:  db.vm.statedb.GetBalance(addr),
			CodeHash: db.vm.statedb.GetCodeHash(addr),
			CodeSize: db.vm.statedb.GetCodeSize(addr),
			Code:     db.vm.statedb.GetCode(addr),
			Suicided: db.vm.statedb.HasSuicided(addr),
		}
	} else {
		finalAccount = finalAccount.Clone()
	}

	if !hasPrevOrigins {
		db.readSet.Set(locationHash, newOrigins)
	}

	finalAccount.Nonce += nonceAddtion
	if locationHash == db.fromHash && db.tx.Nonce() != finalAccount.Nonce {
		if db.txIdx > 0 {
			panic(BlockingError{Addr: addr, TxIdx: db.txIdx - 1})
		} else {
			panic(InvalidNonceError{db.txIdx})
		}
	}
	finalAccount.Balance = new(big.Int).Add(finalAccount.Balance, balanceAddition)
	// TODO: get code
	/*
		var codeHash common.Hash
		if locationHash == db.toHash {
			codeHash = db.toCodeHash
		} else {
			codeHash = db.GetCodeHash(addr)
	}*/

	db.readAccounts[locationHash] = finalAccount
	return finalAccount
}

func (db *VmDB) GetBalance(addr common.Address) *big.Int {
	locationHash := BasicLoc(addr)
	if db.isLazy {
		if db.fromHash == locationHash {
			return new(big.Int).SetUint64(math.MaxUint64)
		}
		if db.toHash == locationHash {
			return common.Big0
		}
	}

	if basic := db.getAccountBasic(addr); basic != nil {
		return basic.Balance
	}
	return common.Big0
}

func (db *VmDB) GetNonce(addr common.Address) uint64 {
	locationHash := BasicLoc(addr)
	if db.isLazy {
		if db.fromHash == locationHash {
			return db.tx.Nonce()
		}
	}

	if basic := db.getAccountBasic(addr); basic != nil {
		return basic.Nonce
	}
	return 0
}

func (db *VmDB) GetCodeHash(addr common.Address) common.Hash {
	if acc := db.getAccountBasic(addr); acc != nil {
		return acc.CodeHash
	}
	return db.vm.statedb.GetCodeHash(addr)
}

func (db *VmDB) GetCode(addr common.Address) []byte {
	codeHash := db.GetCodeHash(addr)
	if code, ok := db.vm.mvMemory.newByteCodes.Get(codeHash); ok {
		return code
	}
	return db.vm.statedb.GetCode(addr)
}

func (db *VmDB) GetCodeSize(addr common.Address) int {
	return len(db.GetCode(addr))
}

func (db *VmDB) GetRefund() uint64 {
	return db.refund
}

func (db *VmDB) GetCommittedState(addr common.Address, key []byte) []byte {
	return db.vm.statedb.GetCommittedState(addr, key)
}

func (db *VmDB) GetState(addr common.Address, key []byte) []byte {
	locationHash := StateLoc(addr, key)
	readOrigins := db.readSet.GetOrDefault(locationHash)

	if val, exist := db.getStateFromCache(addr, key); exist {
		return val
	}

	// Try reading from multi-version data
	if db.txIdx > 0 {
		if writtenTxs, ok := db.vm.mvMemory.data.Get(locationHash); ok {
			it := writtenTxs.AscendRange(&item{TxIdx: db.txIdx})
			entry := it.NextBack()
			if entry != nil {
				entry := entry.(*item)
				switch entry.Entry.(type) {
				case *DataEntry:
					de := entry.Entry.(*DataEntry)
					db.pushOrigin(readOrigins, NewMemory(TxVersion{
						TxIdx:         entry.TxIdx,
						TxIncarnation: de.TxIncarnation,
					}))
					val := de.Value.(*State).Value
					db.setReadState(addr, key, val)
					return val
				case *EstimateMarker:
					panic(BlockingError{Addr: addr, TxIdx: entry.TxIdx})
				default:
					panic(ErrInvalidMemoryValueType)
				}
			}
		}
	}

	// Fall back to storage
	db.pushOrigin(readOrigins, NewStorage())
	val := db.vm.statedb.GetState(addr, key)
	db.setReadState(addr, key, val)
	return val
}

func (db *VmDB) HasSuicided(addr common.Address) bool {
	if acc := db.getAccountBasic(addr); acc != nil {
		return acc.Suicided
	}
	return db.vm.statedb.HasSuicided(addr)
}

func (db *VmDB) Exist(addr common.Address) bool {
	_, exist := db.readAccounts[BasicLoc(addr)]
	if exist {
		return exist
	}
	return db.vm.statedb.Exist(addr)
}

func (db *VmDB) Empty(addr common.Address) bool {
	if acc := db.getAccountBasic(addr); acc != nil {
		return acc.Empty()
	}
	return db.vm.statedb.Empty(addr)
}

func (db *VmDB) GetLogs(hash common.Hash, blockHash common.Hash) []*coretypes.Log {
	logs := db.logs[hash]
	for _, l := range logs {
		l.BlockHash = blockHash
	}
	return logs
}

func (db *VmDB) CreateAccount(addr common.Address) {
	db.dirties[addr] = struct{}{}
	if db.getAccountBasic(addr) == nil {
		basic := NewEmptyAccountBase(addr)
		db.readAccounts[BasicLoc(addr)] = basic
	}
}

func (db *VmDB) SubBalance(addr common.Address, amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}

	locationHash := BasicLoc(addr)
	isLazy := (db.isLazy && locationHash == db.fromHash) || addr == db.vm.env.Header.Coinbase
	if isLazy {
		if balance, exist := db.subBalances[addr]; exist {
			db.subBalances[addr] = balance.Add(balance, amount)
		} else {
			db.subBalances[addr] = new(big.Int).Set(amount)
		}
		return
	}

	if basic := db.getAccountBasic(addr); basic != nil {
		db.dirties[addr] = struct{}{}
		basic.Balance = basic.Balance.Sub(basic.Balance, amount)
	}
}

func (db *VmDB) AddBalance(addr common.Address, amount *big.Int) {
	locationHash := BasicLoc(addr)
	isLazy := (db.isLazy && locationHash == db.toHash) || addr == db.vm.env.Header.Coinbase
	if isLazy {
		if balance, exist := db.addBalances[addr]; exist {
			db.addBalances[addr] = balance.Add(balance, amount)
		} else {
			db.addBalances[addr] = new(big.Int).Set(amount)
		}
		return
	}

	if basic := db.getAccountBasic(addr); basic != nil {
		if amount.Sign() == 0 {
			if basic.Empty() && basic.Touch() {
				db.dirties[addr] = struct{}{}
			}
			return
		}
		db.dirties[addr] = struct{}{}
		basic.Balance = basic.Balance.Add(basic.Balance, amount)
	}
}

func (db *VmDB) SetBalance(common.Address, *big.Int) {
	panic("not implement")
}

func (db *VmDB) SetNonce(addr common.Address, nonce uint64) {
	locationHash := BasicLoc(addr)
	if db.isLazy && locationHash == db.fromHash {
		return // Lazy cumulative
	}

	if basic := db.getAccountBasic(addr); basic != nil {
		db.dirties[addr] = struct{}{}
		basic.Nonce = nonce
	}
}

func (db *VmDB) SetCode(addr common.Address, code []byte) {
	if basic := db.getAccountBasic(addr); basic != nil {
		db.dirties[addr] = struct{}{}
		basic.Code = code
		basic.CodeHash = crypto.Keccak256Hash(code)
		basic.CodeSize = len(code)
		basic.NewCode = true
	}
}

func (db *VmDB) AddRefund(gas uint64) {
	// FIXME: lazy caculate?
	db.refund += gas
}

func (db *VmDB) SubRefund(gas uint64) {
	// FIXME: lazy caculate?
	if gas > db.refund {
		panic(fmt.Sprintf("Refund counter below zero (gas: %d > refund: %d", gas, db.refund))
	}
	db.refund -= gas
}

func (db *VmDB) SetState(addr common.Address, key, val []byte) {
	db.dirties[addr] = struct{}{} // FIXME: is need to set dirty?
	if _, ok := db.states[addr]; !ok {
		db.states[addr] = make(map[string][]byte, 0)
	}
	db.states[addr][string(key)] = val
}

func (db *VmDB) Suicide(addr common.Address) bool {
	if basic := db.getAccountBasic(addr); basic != nil {
		db.dirties[addr] = struct{}{}
		basic.Suicided = true
		basic.Balance = new(big.Int)
		return true
	}
	return false
}

func (db *VmDB) AddLog(logInfo *coretypes.Log) {
	logInfo.TxHash = db.TxHash()
	logInfo.TxIndex = uint(db.TxIndex())
	logInfo.Index = db.logSize
	db.logs[logInfo.TxHash] = append(db.logs[logInfo.TxHash], logInfo)
	db.logSize++
}

func (db *VmDB) AddPreimage(common.Hash, []byte) {
	panic("not implement")
}

func (db *VmDB) PrepareAccessList(sender common.Address, dest *common.Address, precompiles []common.Address, list coretypes.AccessList) {
	db.AddAddressToAccessList(sender)
	if dest != nil {
		db.AddAddressToAccessList(*dest)
	}
	for _, addr := range precompiles {
		db.AddAddressToAccessList(addr)
	}
	for _, el := range list {
		db.AddAddressToAccessList(el.Address)
		for _, key := range el.StorageKeys {
			db.AddSlotToAccessList(el.Address, key)
		}
	}
}

func (db *VmDB) AddressInAccessList(addr common.Address) bool {
	return db.accessList.ContainsAddress(addr)
}

func (db *VmDB) SlotInAccessList(addr common.Address, slot common.Hash) (bool, bool) {
	return db.accessList.Contains(addr, slot)
}

func (db *VmDB) AddAddressToAccessList(addr common.Address) {
}

func (db *VmDB) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	db.accessList.AddSlot(addr, slot)
}

func (db *VmDB) RevertToSnapshot(int)                                     { panic("not implement") }
func (db *VmDB) Snapshot() int                                            { panic("not implement") }
func (db *VmDB) ForEachStorage(common.Address, func([]byte, []byte) bool) { panic("not implement") }
func (db *VmDB) MigrateStorage(common.Address, common.Address)            { panic("not implement") }

func (db *VmDB) Prepare(common.Hash, int) {}
func (db *VmDB) TxHash() common.Hash      { return db.tx.Hash() }
func (db *VmDB) TxIndex() int             { return int(db.txIdx) }

func (db *VmDB) Finalise(bool)                     {}
func (db *VmDB) IntermediateRoot(bool) common.Hash { return common.ZeroHash }

func (db *VmDB) setReadState(addr common.Address, key, val []byte) {
	if states, exist := db.readStates[addr]; exist {
		states[string(key)] = val
		db.readStates[addr] = states
	} else {
		states = make(map[string][]byte)
		states[string(key)] = val
		db.readStates[addr] = states
	}
}

func (db *VmDB) getStateFromCache(addr common.Address, key []byte) ([]byte, bool) {
	if states, exist := db.states[addr]; exist {
		if val, valExist := states[string(key)]; valExist {
			return val, valExist
		}
	}
	if states, exist := db.readStates[addr]; exist {
		val, valExist := states[string(key)]
		return val, valExist
	}
	return []byte{}, false
}
