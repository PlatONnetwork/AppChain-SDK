package pevm

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/PlatONnetwork/PlatON-Go/common"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

var (
	_             sdk.StateDB = (*VmDB)(nil)
	maxUint256, _             = new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)

	vmdbPool = sync.Pool{
		New: func() interface{} {
			return &VmDB{
				readSet:             NewReadSet(),
				readAccounts:        make(map[common.Address]*AccountBase, 3),
				dirties:             make(map[common.Address]struct{}, 3),
				states:              make(map[common.Address]map[string][]byte, 3),
				readStates:          make(map[common.Address]map[string][]byte, 3),
				addBalances:         make(map[common.Address]*big.Int, 3),
				subBalances:         make(map[common.Address]*big.Int, 3),
				readCodeHash:        make(map[common.Address]common.Hash, 1),
				refund:              0,
				logs:                make(map[common.Hash][]*coretypes.Log, 0),
				accessList:          newAccessList(),
				snapshotAddBalances: make(map[common.Address]*big.Int, 3),
				snapshotSubBalances: make(map[common.Address]*big.Int, 3),
			}
		},
	}
)

const initPoolSize = 64

func init() {
	for i := 0; i < initPoolSize; i++ {
		vmdbPool.Put(vmdbPool.New())
	}
}

func AcquireVmDB() *VmDB {
	db := vmdbPool.Get().(*VmDB)
	db.reset()
	return db
}

func ReleaseVmDB(db *VmDB) {
	db.reset()
	vmdbPool.Put(db)
}

type VmDB struct {
	vm           *Vm
	txIdx        int32
	tx           *coretypes.Transaction
	fromAddr     common.Address
	toAddr       common.Address
	fromHash     MemoryLocationHash
	toHash       MemoryLocationHash
	toCodeHash   common.Hash
	isLazy       bool
	readSet      *ReadSet
	readAccounts map[common.Address]*AccountBase
	dirties      map[common.Address]struct{}
	states       map[common.Address]map[string][]byte
	readStates   map[common.Address]map[string][]byte
	addBalances  map[common.Address]*big.Int
	subBalances  map[common.Address]*big.Int

	readCodeHash map[common.Address]common.Hash

	refund     uint64
	logs       map[common.Hash][]*coretypes.Log
	logSize    uint
	accessList *accessList

	snapshotAddBalances map[common.Address]*big.Int
	snapshotSubBalances map[common.Address]*big.Int

	abortErr error
}

func NewVmDB(
	vm *Vm,
	txIdx int32,
	tx *coretypes.Transaction,
	fromAddr common.Address,
	fromHash, toHash MemoryLocationHash) *VmDB {
	db := &VmDB{
		vm:                  vm,
		txIdx:               txIdx,
		tx:                  tx,
		fromAddr:            fromAddr,
		toAddr:              common.ZeroAddr,
		fromHash:            fromHash,
		toHash:              toHash,
		readSet:             NewReadSet(),
		readAccounts:        make(map[common.Address]*AccountBase, 3),
		dirties:             make(map[common.Address]struct{}, 3),
		states:              make(map[common.Address]map[string][]byte, 3),
		readStates:          make(map[common.Address]map[string][]byte, 3),
		addBalances:         make(map[common.Address]*big.Int, 3),
		subBalances:         make(map[common.Address]*big.Int, 3),
		readCodeHash:        make(map[common.Address]common.Hash, 1),
		refund:              0,
		logs:                make(map[common.Hash][]*coretypes.Log, 0),
		accessList:          newAccessList(),
		snapshotAddBalances: make(map[common.Address]*big.Int, 3),
		snapshotSubBalances: make(map[common.Address]*big.Int, 3),
	}
	if tx.To() != nil {
		db.toAddr = *tx.To()
		db.toCodeHash = db.getCodeHash(db.toAddr)
		db.isLazy = db.toCodeHash == emptyCodeHash &&
			(vm.mvMemory.data.Has(fromHash) || vm.mvMemory.data.Has(toHash))
	}
	return db
}

func (db *VmDB) Init(vm *Vm,
	txIdx int32,
	tx *coretypes.Transaction,
	fromAddr common.Address,
	fromHash, toHash MemoryLocationHash) {
	db.vm = vm
	db.txIdx = txIdx
	db.tx = tx
	db.fromAddr = fromAddr
	db.fromHash = fromHash
	db.toHash = toHash
	toAddr := tx.To()
	if toAddr != nil {
		db.toAddr = *toAddr
		db.toCodeHash = db.getCodeHash(db.toAddr)
		db.isLazy = db.toCodeHash == emptyCodeHash &&
			(vm.mvMemory.data.Has(fromHash) || vm.mvMemory.data.Has(toHash))
	}
	if !db.isLazy {
		db.getAccountBasic(db.fromAddr)
		if toAddr != nil {
			db.getAccountBasic(db.toAddr)
		}
	}
}

func (db *VmDB) reset() {
	db.readSet = NewReadSet()
	db.accessList.Reset()
	for k, _ := range db.readAccounts {
		delete(db.readAccounts, k)
	}
	for k, _ := range db.dirties {
		delete(db.dirties, k)
	}
	for k, _ := range db.states {
		delete(db.states, k)
	}
	for k, _ := range db.readStates {
		delete(db.readStates, k)
	}
	for k, _ := range db.addBalances {
		delete(db.addBalances, k)
	}
	for k, _ := range db.subBalances {
		delete(db.subBalances, k)
	}
	for k, _ := range db.readCodeHash {
		delete(db.readCodeHash, k)
	}
	for k, _ := range db.logs {
		delete(db.logs, k)
	}
	for k, _ := range db.snapshotAddBalances {
		delete(db.snapshotAddBalances, k)
	}
	for k, _ := range db.snapshotSubBalances {
		delete(db.snapshotSubBalances, k)
	}
	db.logSize = 0
	db.refund = 0
	db.vm = nil
	db.tx = nil
	db.isLazy = false
	db.abortErr = nil
}

func (db *VmDB) pushOrigin(readOrigins *ReadOrigins, readOrigin ReadOrigin) {
	if readOrigins.Len() > 0 {
		last := readOrigins.Last()
		if !last.Equal(readOrigin) {
			db.abortErr = ErrInconsistentRead
			return
		}
	}
	readOrigins.Push(readOrigin)
}

func (db *VmDB) hashBasic(addr common.Address) MemoryLocationHash {
	if addr == db.fromAddr {
		return db.fromHash
	}
	if db.toAddr != common.ZeroAddr && addr == db.toAddr {
		return db.toHash
	}
	return BasicLoc(addr)
}

func (db *VmDB) getAccountBasic(addr common.Address) *AccountBase {
	if db.abortErr != nil {
		return nil
	}

	if db.isLazy {
		if addr == db.fromAddr || addr == db.toAddr {
			return nil
		}
	}
	if basic, ok := db.readAccounts[addr]; ok {
		return basic
	}

	var (
		locationHash           = db.hashBasic(addr)
		readOrigins            = db.readSet.GetOrDefault(locationHash)
		hasPrevOrigins         = readOrigins.Len() > 0
		newOrigins             = NewReadOrigins()
		balanceAddition        = new(big.Int)
		nonceAddtion    uint64 = 0
		finalAccount    *AccountBase
	)

	if db.txIdx > 0 {
		if wh := db.vm.mvMemory.data.Get(locationHash); wh != nil {
			wh.ScanHistory(db.txIdx, func(entry *item) bool {
				switch entry.Entry.(type) {
				case *DataEntry:
					de := entry.Entry.(*DataEntry)
					// About to push a new origin
					// Inconsistent: new origin will be longer than the previous!
					if hasPrevOrigins && readOrigins.Len() == newOrigins.Len() {
						db.abortErr = ErrInconsistentRead
						return false
					}

					origin := NewMemory(TxVersion{
						TxIdx:         entry.TxIdx,
						TxIncarnation: de.TxIncarnation,
					})
					if hasPrevOrigins {
						if !origin.Equal(readOrigins.Get(newOrigins.Len())) {
							db.abortErr = ErrInconsistentRead
							return false
						}
					} else {
						newOrigins.Push(origin)
					}
					switch de.Value.(type) {
					case *Basic:
						basic := de.Value.(*Basic)
						finalAccount = basic.Account
						return false
					case *LazySender:
						lazySender := de.Value.(*LazySender)
						balanceAddition.Sub(balanceAddition, lazySender.Balance)
						nonceAddtion += 1
					case *LazyRecipient:
						lazyRecipient := de.Value.(*LazyRecipient)
						balanceAddition.Add(balanceAddition, lazyRecipient.Balance)
					default:
						db.abortErr = ErrInvalidMemoryValueType
						return false
					}
				case *EstimateMarker:
					db.abortErr = BlockingError{Addr: addr, TxIdx: entry.TxIdx}
					return false
				}
				return true
			})
		}
	}

	if finalAccount == nil {
		if !hasPrevOrigins {
			newOrigins.Push(NewStorage())
		} else if readOrigins.Len() != newOrigins.Len()+1 ||
			!readOrigins.Last().Equal(NewStorage()) {
			db.abortErr = ErrInconsistentRead
			return nil
		}
		finalAccount = &AccountBase{
			Addr:    addr,
			Nonce:   db.vm.statedb.GetNonce(addr),
			Balance: new(big.Int).Set(db.vm.statedb.GetBalance(addr)),
			//CodeHash: db.GetCodeHash(addr),
			//CodeSize: db.GetCodeSize(addr),
			//Code:     db.GetCode(addr),
			//Suicided: db.vm.statedb.HasSuicided(addr),
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
			db.abortErr = BlockingError{Addr: addr, TxIdx: db.txIdx - 1}
		} else {
			log.Error("Get account basic: invalid nonce", "txIdx", db.txIdx, "txHash", db.tx.Hash(),
				"from", db.fromAddr.Hex(), "txNonce", db.tx.Nonce(), "curNonce", finalAccount.Nonce,
				"nonceAddition", nonceAddtion)
			db.abortErr = InvalidNonceError{db.txIdx, db.tx.Nonce(), finalAccount.Nonce}
		}
		return nil
	}
	finalAccount.Balance.Add(finalAccount.Balance, balanceAddition)
	// TODO: get code
	/*
		var codeHash common.Hash
		if locationHash == db.toHash {
			codeHash = db.toCodeHash
		} else {
			codeHash = db.GetCodeHash(addr)
	}*/

	db.readAccounts[addr] = finalAccount
	return finalAccount
}

func (db *VmDB) GetBalance(addr common.Address) *big.Int {
	if db.abortErr != nil {
		return big.NewInt(0)
	}

	if db.isLazy {
		if db.fromAddr == addr {
			return new(big.Int).Set(maxUint256)
		}
		if db.toAddr == addr {
			return big.NewInt(0)
		}
	}

	if basic := db.getAccountBasic(addr); basic != nil {
		return basic.Balance
	}
	return big.NewInt(0)
}

func (db *VmDB) GetNonce(addr common.Address) uint64 {
	if db.abortErr != nil {
		return 0
	}

	if db.isLazy {
		if db.fromAddr == addr {
			return db.tx.Nonce()
		}
	}

	if basic := db.getAccountBasic(addr); basic != nil {
		return basic.Nonce
	}
	return 0
}

func (db *VmDB) GetCodeHash(addr common.Address) common.Hash {
	if db.abortErr != nil {
		return emptyCodeHash
	}

	if addr == db.toAddr {
		return db.toCodeHash
	}
	return db.getCodeHash(addr)
}

func (db *VmDB) getCodeHash(addr common.Address) common.Hash {
	if db.abortErr != nil {
		return emptyCodeHash
	}

	if h, ok := db.readCodeHash[addr]; ok {
		return h
	}

	var (
		locationHash = CodeHashLoc(addr)
		readOrigins  = db.readSet.GetOrDefault(locationHash)
		h            = emptyCodeHash
		got          bool
	)

	if wh := db.vm.mvMemory.data.Get(locationHash); wh != nil {
		wh.ScanHistory(db.txIdx, func(entryItem *item) bool {
			if entry, ok := entryItem.Entry.(*DataEntry); ok {
				switch entry.Value.(type) {
				case *SelfDestructed:
					db.abortErr = ErrSelfDestructedAccount
					got = true
					return false
				case *CodeHash:
					codeHash := entry.Value.(*CodeHash)
					db.pushOrigin(readOrigins, NewMemory(TxVersion{
						TxIdx:         entryItem.TxIdx,
						TxIncarnation: entry.TxIncarnation,
					}))
					db.readCodeHash[addr] = codeHash.CodeHash
					h = codeHash.CodeHash
					got = true
					return false
				}
			}
			return false
		})
	}
	if got {
		return h
	}

	// Fallback to storage
	db.pushOrigin(readOrigins, NewStorage())
	h = db.vm.statedb.GetCodeHash(addr)
	db.readCodeHash[addr] = h
	return h
}

func (db *VmDB) GetCode(addr common.Address) []byte {
	if db.abortErr != nil {
		return []byte{}
	}

	codeHash := db.GetCodeHash(addr)
	if codeHash == emptyCodeHash {
		return []byte{}
	}
	if code, ok := db.vm.mvMemory.newByteCodes.Load(codeHash); ok {
		return code.([]byte)
	}
	return db.vm.statedb.GetCode(addr)
}

func (db *VmDB) GetCodeSize(addr common.Address) int {
	if db.abortErr != nil {
		return 0
	}
	return len(db.GetCode(addr))
}

func (db *VmDB) GetRefund() uint64 {
	return db.refund
}

func (db *VmDB) GetCommittedState(addr common.Address, key []byte) []byte {
	if db.txIdx > 0 {
		var (
			val []byte
			got bool
		)
		locationHash := StateLoc(addr, key)
		if wh := db.vm.mvMemory.data.Get(locationHash); wh != nil {
			wh.ScanHistory(db.txIdx, func(entry *item) bool {
				switch entry.Entry.(type) {
				case *DataEntry:
					de := entry.Entry.(*DataEntry)
					val = de.Value.(*State).Value
					db.setReadState(addr, key, val)
					got = true
					return false
				case *EstimateMarker:
					// Here `GetState` should be throw the same error, so discard
					// abort error in this case.
					got = true
					return false
				default:
					// Here `GetState` should be throw the same error, so discard
					// abort error in this case.
					got = true
					return false
				}
			})
		}
		if got {
			return val
		}
	}
	return db.vm.statedb.GetCommittedState(addr, key)
}

func (db *VmDB) GetState(addr common.Address, key []byte) []byte {
	if db.abortErr != nil {
		return []byte{}
	}

	if val, exist := db.getStateFromCache(addr, key); exist {
		return val
	}

	locationHash := StateLoc(addr, key)
	readOrigins := db.readSet.GetOrDefault(locationHash)
	var (
		val []byte
		got bool
	)

	// Try reading from multi-version data
	if db.txIdx > 0 {
		if wh := db.vm.mvMemory.data.Get(locationHash); wh != nil {
			wh.ScanHistory(db.txIdx, func(entry *item) bool {
				switch entry.Entry.(type) {
				case *DataEntry:
					de := entry.Entry.(*DataEntry)
					db.pushOrigin(readOrigins, NewMemory(TxVersion{
						TxIdx:         entry.TxIdx,
						TxIncarnation: de.TxIncarnation,
					}))
					val = de.Value.(*State).Value
					db.setReadState(addr, key, val)
					got = true
					return false
				case *EstimateMarker:
					db.abortErr = BlockingError{Addr: addr, TxIdx: entry.TxIdx}
					got = true
					return false
				default:
					db.abortErr = ErrInvalidMemoryValueType
					got = true
					return false
				}
			})
		}
	}
	if got {
		return val
	}

	// Fall back to storage
	db.pushOrigin(readOrigins, NewStorage())
	val = db.vm.statedb.GetState(addr, key)
	db.setReadState(addr, key, val)
	return val
}

func (db *VmDB) HasSuicided(addr common.Address) bool {
	if db.abortErr != nil {
		return true
	}

	if acc := db.getAccountBasic(addr); acc != nil {
		if acc.Suicided == nil {
			suicided := db.vm.statedb.HasSuicided(addr)
			acc.Suicided = &suicided
		}
		return *acc.Suicided
	}
	return db.vm.statedb.HasSuicided(addr)
}

func (db *VmDB) Exist(addr common.Address) bool {
	if db.abortErr != nil {
		return false
	}

	_, exist := db.readAccounts[addr]
	if exist {
		return exist
	}
	return db.vm.statedb.Exist(addr)
}

func (db *VmDB) Empty(addr common.Address) bool {
	if db.abortErr != nil {
		return true
	}

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
	if db.abortErr != nil {
		return
	}
	if db.getAccountBasic(addr) == nil {
		basic := NewEmptyAccountBase(addr)
		db.readAccounts[addr] = basic
	}
}

func (db *VmDB) SubBalance(addr common.Address, amount *big.Int) {
	if db.abortErr != nil {
		return
	}

	if amount.Sign() == 0 {
		return
	}

	if db.lazy(addr) {
		if balance, exist := db.subBalances[addr]; exist {
			balance.Add(balance, amount)
		} else {
			db.subBalances[addr] = new(big.Int).Set(amount)
		}
		return
	}

	if basic := db.getAccountBasic(addr); basic != nil {
		if _, ok := db.dirties[addr]; !ok {
			db.dirties[addr] = struct{}{}
		}
		basic.Balance.Sub(basic.Balance, amount)
	}
}

func (db *VmDB) AddBalance(addr common.Address, amount *big.Int) {
	if db.abortErr != nil {
		return
	}

	if db.lazy(addr) {
		if addr == db.fromAddr {
			// refund
			if balance, ok := db.subBalances[addr]; ok {
				if amount.Cmp(balance) > 0 {
					panic(fmt.Sprintf("invalid balance(addr: %s, balance: %s, amount: %s)", addr.Hex(), amount, balance))
				}
				balance.Sub(balance, amount)
			} else {
				panic(fmt.Sprintf("AddBalance: unreachable(%s balance not found)", addr.Hex()))
			}
			return
		}

		if balance, exist := db.addBalances[addr]; exist {
			balance.Add(balance, amount)
		} else {
			db.addBalances[addr] = new(big.Int).Set(amount)
		}
		return
	}

	if basic := db.getAccountBasic(addr); basic != nil {
		if amount.Sign() == 0 {
			if basic.Empty() && basic.Touch() {
				if _, ok := db.dirties[addr]; !ok {
					db.dirties[addr] = struct{}{}
				}
			}
			return
		}
		if _, ok := db.dirties[addr]; !ok {
			db.dirties[addr] = struct{}{}
		}
		basic.Balance.Add(basic.Balance, amount)
	}
}

func (db *VmDB) SetBalance(common.Address, *big.Int) {
	panic("not implement")
}

func (db *VmDB) SetNonce(addr common.Address, nonce uint64) {
	if db.abortErr != nil {
		return
	}

	if db.isLazy && addr == db.fromAddr {
		return // Lazy cumulative
	}

	if basic := db.getAccountBasic(addr); basic != nil {
		if _, ok := db.dirties[addr]; !ok {
			db.dirties[addr] = struct{}{}
		}
		basic.Nonce = nonce
	}
}

func (db *VmDB) SetCode(addr common.Address, code []byte) {
	if db.abortErr != nil {
		return
	}
	if basic := db.getAccountBasic(addr); basic != nil {
		if _, ok := db.dirties[addr]; !ok {
			db.dirties[addr] = struct{}{}
		}
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
	if _, ok := db.states[addr]; !ok {
		db.states[addr] = make(map[string][]byte, 0)
	}
	db.states[addr][string(key)] = val
}

func (db *VmDB) Suicide(addr common.Address) bool {
	if db.abortErr != nil {
		return false
	}
	if basic := db.getAccountBasic(addr); basic != nil {
		if _, ok := db.dirties[addr]; !ok {
			db.dirties[addr] = struct{}{}
		}
		suicided := true
		basic.Suicided = &suicided
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
	db.accessList.AddAddress(addr)
}

func (db *VmDB) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	db.accessList.AddSlot(addr, slot)
}

func (db *VmDB) RevertToSnapshot(int) {
	db.refund = 0
	for k, _ := range db.dirties {
		delete(db.dirties, k)
	}
	for k, _ := range db.states {
		delete(db.states, k)
	}
	for k, _ := range db.logs {
		delete(db.logs, k)
	}
	for k, _ := range db.addBalances {
		delete(db.addBalances, k)
	}
	for k, _ := range db.subBalances {
		delete(db.subBalances, k)
	}

	for k, v := range db.snapshotAddBalances {
		db.addBalances[k] = new(big.Int).Set(v)
		delete(db.snapshotAddBalances, k)
	}

	for k, v := range db.snapshotSubBalances {
		db.subBalances[k] = new(big.Int).Set(v)
		delete(db.snapshotSubBalances, k)
	}
}

func (db *VmDB) Snapshot() int {
	for k, v := range db.addBalances {
		db.snapshotAddBalances[k] = new(big.Int).Set(v)
	}
	for k, v := range db.subBalances {
		db.snapshotSubBalances[k] = new(big.Int).Set(v)
	}
	return 0
}

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

func (db *VmDB) lazy(addr common.Address) bool {
	lazyAddr := db.isLazy && (addr == db.fromAddr || addr == db.toAddr)
	return lazyAddr || addr == db.vm.env.Header.Coinbase
}
