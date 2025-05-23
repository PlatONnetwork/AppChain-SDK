package pevm

import (
	"math/big"
	"reflect"

	"github.com/PlatONnetwork/AppChain-SDK/x/pevm/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

var (
	_ sdk.StateDB = (*VmDB)(nil)
)

type VmDB struct {
	vm           *VM
	txIdx        uint32
	tx           *coretypes.Transaction
	fromHash     types.MemoryLocationHash
	toHash       types.MemoryLocationHash
	toCodeHash   common.Hash
	isLazy       bool
	readSet      types.ReadSet
	readAccounts map[types.MemoryLocationHash]*types.AccountBase

	logs       map[common.Hash][]*coretypes.Log
	logSize    uint
	accessList *accessList
}

func NewVmDB(
	vm *VM,
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
	l := len(*readOrigins)
	if l > 0 {
		last := (*readOrigins)[l-1]
		if !reflect.DeepEqual(last, readOrigin) {
			panic("inconsistent read")
		}
	}
	*readOrigins = append(*readOrigins, readOrigin)
}

func (db *VmDB) hashBaisc(addr common.Address) types.MemoryLocationHash {
	if addr == db.tx.FromAddr(coretypes.NewEIP155Signer(db.vm.ctx.ChainConfig().ChainID)) {
		return db.fromHash
	}
	if db.tx.To() != nil && addr == *db.tx.To() {
		return db.toHash
	}
	return types.BasicLoc(addr)
}

func (db *VmDB) GetBalance(addr common.Address) *big.Int {
	locationHash := db.hashBaisc(addr)
	readOrigins := db.readSet[locationHash]
	hasPrevOrigins := len(readOrigins) > 0
	newOrigins := make(types.ReadOrigins, 0)

	return nil
}

func (db *VmDB) GetNonce(addr common.Address) uint64 {
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
				db.pushOrigin(&readOrigins, types.NewMvMemory(types.TxVersion{
					TxIdx:         entry.TxIdx,
					TxIncarnation: dataEntry.TxIncarnation,
				}))
				return ch.CodeHash
			case *types.SelfDestructed:
				return common.ZeroHash
			}
		}
	}
	db.pushOrigin(&readOrigins, types.NewStorage())
	return db.vm.statedb.GetCodeHash(addr)
}

func (db *VmDB) GetCode(addr common.Address) []byte {
	return nil
}

func (db *VmDB) GetCodeSize(addr common.Address) int {
	return 0
}

func (db *VmDB) GetRefund() uint64 {
	return 0
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
				db.pushOrigin(&readOrigins, types.NewMvMemory(types.TxVersion{
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
	db.pushOrigin(&readOrigins, types.NewStorage())
	return db.vm.statedb.GetState(addr, key)
}

func (db *VmDB) HasSuicided(addr common.Address) bool {
	locationHash := types.SuicideLoc(addr)
	readOrigins := db.readSet[locationHash]

	if db.txIdx > 0 {
		if writtenTxs, ok := db.vm.mvMemory.data.Get(locationHash); ok {
			it := writtenTxs.AscendRange(&DataEntry{TxIdx: db.txIdx})
			entry := it.NextBack().(*DataEntry)
			switch entry.Entry.(type) {
			case *types.DataEntry:
				// SelfDestructed
				de := entry.Entry.(*types.DataEntry)
				db.pushOrigin(&readOrigins, types.NewMvMemory(types.TxVersion{
					TxIdx:         entry.TxIdx,
					TxIncarnation: de.TxIncarnation,
				}))
				return true
			}
		}
	}

	db.pushOrigin(&readOrigins, types.NewStorage())
	return db.vm.statedb.HasSuicided(addr)
}

func (db *VmDB) Exist(addr common.Address) bool {
	return true
}

func (db *VmDB) Empty(addr common.Address) bool {
	return false
}

func (db *VmDB) GetLogs(hash common.Hash, blockHash common.Hash) []*coretypes.Log {
	logs := db.logs[hash]
	for _, l := range logs {
		l.BlockHash = blockHash
	}
	return logs
}

func (db *VmDB) CreateAccount(addr common.Address) {

}

func (db *VmDB) SubBalance(addr common.Address, amount *big.Int) {

}

func (db *VmDB) AddBalance(addr common.Address, amount *big.Int) {

}

func (db *VmDB) SetNonce(addr common.Address, nonce uint64) {

}

func (db *VmDB) SetCode(addr common.Address, code []byte) {

}

func (db *VmDB) AddRefund(amount uint64) {

}

func (db *VmDB) SubRefund(amount uint64) {

}

func (db *VmDB) SetState(addr common.Address, key, val []byte) {

}

func (db *VmDB) Suicide(addr common.Address) bool {
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
