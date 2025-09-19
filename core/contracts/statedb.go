package contracts

import (
	"math/big"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
)

type StateDB struct {
	Storage
}

func NewStateDB(evm *vm.EVM, contract *vm.Contract) *StateDB {
	return &StateDB{
		Storage: Storage{
			evm:       evm,
			contract:  contract,
			burn:      NewBurner(contract),
			jumpTable: newJumpTable(),
		},
	}
}

func (s *StateDB) Prepare(txHash common.Hash, i int) {
	s.evm.StateDB.Prepare(txHash, i)
}

func (s *StateDB) AddRefund(u uint64) {
	s.evm.StateDB.AddRefund(u)
}

func (s *StateDB) SubRefund(u uint64) {
	s.evm.StateDB.SubRefund(u)
}

func (s *StateDB) GetRefund() uint64 {
	return s.evm.StateDB.GetRefund()
}

func (s *StateDB) GetCommittedState(address common.Address, bytes []byte) []byte {
	return s.evm.StateDB.GetCommittedState(address, bytes)
}

func (s *StateDB) Suicide(address common.Address) bool {
	return s.evm.StateDB.Suicide(address)
}

func (s *StateDB) HasSuicided(address common.Address) bool {
	return s.evm.StateDB.HasSuicided(address)
}

func (s *StateDB) Exist(address common.Address) bool {
	return s.evm.StateDB.Exist(address)
}

func (s *StateDB) Empty(address common.Address) bool {
	return s.evm.StateDB.Empty(address)
}

func (s *StateDB) PrepareAccessList(sender common.Address, dest *common.Address, precompiles []common.Address, txAccesses types.AccessList) {
	s.evm.StateDB.PrepareAccessList(sender, dest, precompiles, txAccesses)
}

func (s *StateDB) AddressInAccessList(addr common.Address) bool {
	return s.evm.StateDB.AddressInAccessList(addr)
}

func (s *StateDB) SlotInAccessList(addr common.Address, slot common.Hash) (addressOk bool, slotOk bool) {
	return s.evm.StateDB.SlotInAccessList(addr, slot)
}

func (s *StateDB) AddAddressToAccessList(addr common.Address) {
	s.evm.StateDB.AddAddressToAccessList(addr)
}

func (s *StateDB) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	s.evm.StateDB.AddSlotToAccessList(addr, slot)
}

func (s *StateDB) RevertToSnapshot(i int) {
	s.evm.StateDB.RevertToSnapshot(i)
}

func (s *StateDB) Snapshot() int {
	return s.evm.StateDB.Snapshot()
}

func (s *StateDB) GetLogs(hash common.Hash, blockHash common.Hash) []*types.Log {
	return s.evm.StateDB.GetLogs(hash, blockHash)
}

func (s *StateDB) AddPreimage(hash common.Hash, bytes []byte) {
	s.evm.StateDB.AddPreimage(hash, bytes)
}

func (s *StateDB) ForEachStorage(address common.Address, f func([]byte, []byte) bool) {
	s.ForEachStorage(address, f)
}

func (s *StateDB) MigrateStorage(from, to common.Address) {
	s.evm.StateDB.MigrateStorage(from, to)
}

func (s *StateDB) TxHash() common.Hash {
	return s.evm.StateDB.TxHash()
}

func (s *StateDB) TxIndex() int {
	return s.evm.StateDB.TxIndex()
}

func (s *StateDB) IntermediateRoot(deleteEmptyObjects bool) common.Hash {
	return s.evm.StateDB.IntermediateRoot(deleteEmptyObjects)
}

func (s *StateDB) Finalise(deleteEmptyObjects bool) {
	s.evm.StateDB.Finalise(deleteEmptyObjects)
}

func (s *StateDB) SetBalance(addr common.Address, balance *big.Int) {
	s.evm.StateDB.SetBalance(addr, balance)
}
