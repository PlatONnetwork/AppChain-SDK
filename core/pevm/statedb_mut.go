package pevm

import (
	"math/big"
	"sync"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

var _ sdk.StateDB = (*StateDBMut)(nil)

type StateDBMut struct {
	sync.Mutex

	statedb sdk.StateDB
}

func NewStateDBMut(statedb sdk.StateDB) *StateDBMut {
	return &StateDBMut{
		statedb: statedb,
	}
}

func (s *StateDBMut) PrepareAccessList(common.Address, *common.Address, []common.Address, types.AccessList) {
	panic("not implement")
}
func (s *StateDBMut) AddressInAccessList(common.Address) bool { panic("not implement") }
func (s *StateDBMut) SlotInAccessList(common.Address, common.Hash) (bool, bool) {
	panic("not implement")
}
func (s *StateDBMut) AddAddressToAccessList(common.Address)           { panic("not implement") }
func (s *StateDBMut) AddSlotToAccessList(common.Address, common.Hash) { panic("not implement") }
func (s *StateDBMut) RevertToSnapshot(int)                            { panic("not implement") }
func (s *StateDBMut) Snapshot() int                                   { panic("not implement") }
func (s *StateDBMut) ForEachStorage(common.Address, func([]byte, []byte) bool) {
	panic("not implemente")
}
func (s *StateDBMut) MigrateStorage(common.Address, common.Address) { panic("not implement") }
func (s *StateDBMut) Prepare(common.Hash, int)                      { panic("not implement") }
func (s *StateDBMut) TxHash() common.Hash                           { panic("not implement") }
func (s *StateDBMut) TxIndex() int                                  { panic("not implement") }
func (s *StateDBMut) Finalise(bool)                                 { panic("not implement") }
func (s *StateDBMut) IntermediateRoot(bool) common.Hash             { panic("not implement") }
func (s *StateDBMut) CreateAccount(common.Address)                  { panic("not implement") }
func (s *StateDBMut) SubBalance(common.Address, *big.Int)           { panic("not implement") }
func (s *StateDBMut) AddBalance(common.Address, *big.Int)           { panic("not implement") }
func (s *StateDBMut) SetBalance(common.Address, *big.Int)           { panic("not implement") }
func (s *StateDBMut) SetNonce(common.Address, uint64)               { panic("not implement") }
func (s *StateDBMut) SetCode(common.Address, []byte)                { panic("not implement") }
func (s *StateDBMut) AddRefund(uint64)                              { panic("not implement") }
func (s *StateDBMut) SubRefund(uint64)                              { panic("not implement") }
func (s *StateDBMut) SetState(common.Address, []byte, []byte)       { panic("not implement") }
func (s *StateDBMut) Suicide(common.Address) bool                   { panic("not implement") }
func (s *StateDBMut) AddLog(*types.Log)                             { panic("not implement") }
func (s *StateDBMut) AddPreimage(common.Hash, []byte)               { panic("not implement") }

func (s *StateDBMut) GetBalance(addr common.Address) *big.Int {
	s.Lock()
	defer s.Unlock()
	return s.statedb.GetBalance(addr)
}

func (s *StateDBMut) GetNonce(addr common.Address) uint64 {
	s.Lock()
	defer s.Unlock()
	return s.statedb.GetNonce(addr)
}

func (s *StateDBMut) GetCodeHash(addr common.Address) common.Hash {
	s.Lock()
	defer s.Unlock()
	return s.statedb.GetCodeHash(addr)
}

func (s *StateDBMut) GetCode(addr common.Address) []byte {
	s.Lock()
	defer s.Unlock()
	return s.statedb.GetCode(addr)
}

func (s *StateDBMut) GetCodeSize(addr common.Address) int {
	s.Lock()
	defer s.Unlock()
	return s.statedb.GetCodeSize(addr)
}

func (s *StateDBMut) GetRefund() uint64 {
	s.Lock()
	defer s.Unlock()
	return s.statedb.GetRefund()
}

func (s *StateDBMut) GetCommittedState(addr common.Address, key []byte) []byte {
	s.Lock()
	defer s.Unlock()
	return s.statedb.GetCommittedState(addr, key)
}

func (s *StateDBMut) GetState(addr common.Address, key []byte) []byte {
	s.Lock()
	defer s.Unlock()
	return s.statedb.GetState(addr, key)
}

func (s *StateDBMut) HasSuicided(addr common.Address) bool {
	s.Lock()
	defer s.Unlock()
	return s.statedb.HasSuicided(addr)
}

func (s *StateDBMut) Exist(addr common.Address) bool {
	s.Lock()
	defer s.Unlock()
	return s.statedb.Exist(addr)
}

func (s *StateDBMut) Empty(addr common.Address) bool {
	s.Lock()
	defer s.Unlock()
	return s.statedb.Empty(addr)
}

func (s *StateDBMut) GetLogs(hash common.Hash, blockHash common.Hash) []*types.Log {
	s.Lock()
	defer s.Unlock()
	return s.statedb.GetLogs(hash, blockHash)
}
