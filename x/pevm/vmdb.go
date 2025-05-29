package pevm

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/PlatONnetwork/AppChain-SDK/x/pevm/types"
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
	txIdx        uint32
	tx           *coretypes.Transaction
	fromHash     types.MemoryLocationHash
	toHash       types.MemoryLocationHash
	toCodeHash   common.Hash
	isLazy       bool
	readSet      types.ReadSet
	readAccounts map[types.MemoryLocationHash]*types.AccountBase
	dirties      map[common.Address]struct{}
	states       map[common.Address]map[string][]byte

	refund     uint64
	logs       map[common.Hash][]*coretypes.Log
	logSize    uint
	accessList *accessList
}

func NewVmDB(
	vm *Vm,
	txIdx uint32,
	tx *coretypes.Transaction,
	fromHash, toHash types.MemoryLocationHash) *VmDB {
	db := &VmDB{
		vm:           vm,
		txIdx:        txIdx,
		tx:           tx,
		fromHash:     fromHash,
		toHash:       toHash,
		readSet:      make(types.ReadSet, 0),
		readAccounts: make(map[types.MemoryLocationHash]*types.AccountBase, 0),
		dirties:      make(map[common.Address]struct{}, 0),
		states:       make(map[common.Address]map[string][]byte, 0),
		refund:       0,
		logs:         make(map[common.Hash][]*coretypes.Log, 0),
		accessList:   newAccessList(),
	}
	if tx.To() != nil {
		db.toCodeHash = vm.statedb.GetCodeHash(*tx.To())
		db.isLazy = db.toCodeHash == common.ZeroHash &&
			(vm.mvMemory.data.Has(fromHash) || vm.mvMemory.data.Has(toHash))
	}
	return db
}

func (db *VmDB) pushOrigin(readOrigins *types.ReadOrigins, readOrigin types.ReadOrigin) {
	if readOrigins.Len() > 0 {
		last := readOrigins.Last()
		if !reflect.DeepEqual(last, readOrigin) {
			panic("inconsistent read")
		}
	}
	readOrigins.Push(readOrigin)
}

func (db *VmDB) hashBasic(addr common.Address) types.MemoryLocationHash {
	if addr == db.tx.FromAddr(coretypes.NewEIP155Signer(db.vm.ctx.ChainConfig().ChainID)) {
		return db.fromHash
	}
	if db.tx.To() != nil && addr == *db.tx.To() {
		return db.toHash
	}
	return types.BasicLoc(addr)
}

func (db *VmDB) getAccountBasic(addr common.Address) *types.AccountBase {
	var (
		locationHash   = db.hashBasic(addr)
		readOrigins    = db.readSet[locationHash]
		hasPrevOrigins = readOrigins.Len() > 0
		newOrigins     = types.NewReadOrigins()
		finalAccount   *types.AccountBase
	)

	if basic, ok := db.readAccounts[locationHash]; ok {
		return basic
	}

	if db.txIdx > 0 {
		if writtenTxs, ok := db.vm.mvMemory.data.Get(locationHash); ok {
			it := writtenTxs.AscendRange(&DataEntry{TxIdx: db.txIdx})
			// Now only deal with account basic, doing lazy calculate in the further.
			// So if we have a `types.DataEntry`, it must be a `types.AccountBasic`.
			entry := it.NextBack().(*DataEntry)
			switch entry.Entry.(type) {
			case *types.DataEntry:
				de := entry.Entry.(*types.DataEntry)
				origin := types.NewMvMemory(types.TxVersion{
					TxIdx:         entry.TxIdx,
					TxIncarnation: de.TxIncarnation,
				})
				if hasPrevOrigins {
					if !reflect.DeepEqual(origin, readOrigins.Get(0)) {
						return nil
					}
				} else {
					newOrigins.Push(origin)
				}
				switch de.Value.(type) {
				case *types.Basic:
					basic := de.Value.(*types.Basic)
					finalAccount = basic.Account
				default:
					return nil
				}
			default:
				return nil
			}
		}
	}

	if finalAccount == nil {
		if !hasPrevOrigins {
			newOrigins.Push(types.NewStorage())
		} else if readOrigins.Len() != newOrigins.Len()+1 ||
			!reflect.DeepEqual(readOrigins.Last(), types.NewStorage()) {
			return nil
		}
		finalAccount = &types.AccountBase{
			Addr:     addr,
			Nonce:    db.vm.statedb.GetNonce(addr),
			Balance:  db.vm.statedb.GetBalance(addr),
			CodeHash: db.vm.statedb.GetCodeHash(addr),
			CodeSize: db.vm.statedb.GetCodeSize(addr),
			Code:     db.vm.statedb.GetCode(addr),
		}
	}

	if !hasPrevOrigins {
		db.readSet[locationHash] = newOrigins
	}

	db.readAccounts[locationHash] = finalAccount
	return finalAccount
}

func (db *VmDB) GetBalance(addr common.Address) *big.Int {
	if basic := db.getAccountBasic(addr); basic != nil {
		return basic.Balance
	}
	return new(big.Int)
}

func (db *VmDB) GetNonce(addr common.Address) uint64 {
	if basic := db.getAccountBasic(addr); basic != nil {
		return basic.Nonce
	}
	return 0
}

func (db *VmDB) GetCodeHash(addr common.Address) common.Hash {
	locationHash := types.CodeHashLoc(addr)
	readOrigins := db.readSet[locationHash]

	if writtenTxs, ok := db.vm.mvMemory.data.Get(locationHash); ok {
		it := writtenTxs.AscendRange(&DataEntry{TxIdx: db.txIdx})
		entry := it.NextBack().(*DataEntry)
		if dataEntry, ok := entry.Entry.(*types.DataEntry); ok {
			switch dataEntry.Value.(type) {
			case *types.CodeHash:
				ch := dataEntry.Value.(*types.CodeHash)
				db.pushOrigin(readOrigins, types.NewMvMemory(types.TxVersion{
					TxIdx:         entry.TxIdx,
					TxIncarnation: dataEntry.TxIncarnation,
				}))
				return ch.CodeHash
			case *types.SelfDestructed:
				return common.ZeroHash
			}
		}
	}
	db.pushOrigin(readOrigins, types.NewStorage())
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
	locationHash := types.StateLoc(addr, key)
	readOrigins := db.readSet[locationHash]

	// Try reading from multi-version data
	if db.txIdx > 0 {
		if writtenTxs, ok := db.vm.mvMemory.data.Get(locationHash); ok {
			it := writtenTxs.AscendRange(&DataEntry{TxIdx: db.txIdx})
			entry := it.NextBack().(*DataEntry)
			switch entry.Entry.(type) {
			case *types.DataEntry:
				de := entry.Entry.(*types.DataEntry)
				db.pushOrigin(readOrigins, types.NewMvMemory(types.TxVersion{
					TxIdx:         entry.TxIdx,
					TxIncarnation: de.TxIncarnation,
				}))
				return de.Value.(*types.State).Value
			default:
				return []byte{}
			}
		}
	}

	// Fall back to storage
	db.pushOrigin(readOrigins, types.NewStorage())
	return db.vm.statedb.GetState(addr, key)
}

func (db *VmDB) HasSuicided(addr common.Address) bool {
	locationHash := types.CodeHashLoc(addr)
	readOrigins := db.readSet[locationHash]

	if db.txIdx > 0 {
		if writtenTxs, ok := db.vm.mvMemory.data.Get(locationHash); ok {
			it := writtenTxs.AscendRange(&DataEntry{TxIdx: db.txIdx})
			entry := it.NextBack().(*DataEntry)
			switch entry.Entry.(type) {
			case *types.DataEntry:
				de := entry.Entry.(*types.DataEntry)
				if _, ok := de.Value.(*types.SelfDestructed); ok {
					db.pushOrigin(readOrigins, types.NewMvMemory(types.TxVersion{
						TxIdx:         entry.TxIdx,
						TxIncarnation: de.TxIncarnation,
					}))
					return true
				}
			}
		}
	}

	db.pushOrigin(readOrigins, types.NewStorage())
	return db.vm.statedb.HasSuicided(addr)
}

func (db *VmDB) Exist(addr common.Address) bool {
	_, exist := db.readAccounts[types.BasicLoc(addr)]
	if exist {
		return exist
	}
	return db.vm.statedb.Exist(addr)
}

func (db *VmDB) Empty(addr common.Address) bool {
	var (
		locationHash   = db.hashBasic(addr)
		readOrigins    = db.readSet[locationHash]
		hasPrevOrigins = readOrigins.Len() > 0
		finalAccount   *types.AccountBase
	)
	if db.txIdx > 0 {
		if writtenTxs, ok := db.vm.mvMemory.data.Get(locationHash); ok {
			it := writtenTxs.AscendRange(&DataEntry{TxIdx: db.txIdx})
			// Now only deal with account basic, doing lazy calculate in the further.
			// So if we have a `types.DataEntry`, it must be a `types.AccountBasic`.
			entry := it.NextBack().(*DataEntry)
			switch entry.Entry.(type) {
			case *types.DataEntry:
				de := entry.Entry.(*types.DataEntry)
				origin := types.NewMvMemory(types.TxVersion{
					TxIdx:         entry.TxIdx,
					TxIncarnation: de.TxIncarnation,
				})
				if !hasPrevOrigins {
					db.pushOrigin(readOrigins, origin)
				}
				switch de.Value.(type) {
				case *types.Basic:
					basic := de.Value.(*types.Basic)
					finalAccount = basic.Account
				}
			}
		}
	}
	if finalAccount != nil {
		return finalAccount.Empty()
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
		basic := types.NewEmptyAccountBase(addr)
		db.readAccounts[types.BasicLoc(addr)] = basic
	}
}

func (db *VmDB) SubBalance(addr common.Address, amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	if basic := db.getAccountBasic(addr); basic != nil {
		db.dirties[addr] = struct{}{}
		basic.Balance = new(big.Int).Sub(basic.Balance, amount)
	}
}

func (db *VmDB) AddBalance(addr common.Address, amount *big.Int) {
	if basic := db.getAccountBasic(addr); basic != nil {
		if amount.Sign() == 0 {
			if basic.Empty() && basic.Touch() {
				db.dirties[addr] = struct{}{}
			}
			return
		}
		basic.Balance = new(big.Int).Add(basic.Balance, amount)
	}
}

func (db *VmDB) SetBalance(common.Address, *big.Int) {
	panic("not implement")
}

func (db *VmDB) SetNonce(addr common.Address, nonce uint64) {
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

func (db *VmDB) Finalise(bool)                     { panic("not implement") }
func (db *VmDB) IntermediateRoot(bool) common.Hash { panic("not implement") }
