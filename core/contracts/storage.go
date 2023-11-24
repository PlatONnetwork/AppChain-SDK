package contracts

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"math/big"
)

type Storage struct {
	evm       *vm.EVM
	contract  *vm.Contract
	burn      Burn
	jumpTable JumpTable
}

func NewStorage(evm *vm.EVM, contract *vm.Contract) *Storage {
	return &Storage{
		evm:       evm,
		contract:  contract,
		burn:      NewBurner(contract),
		jumpTable: newJumpTable(),
	}
}

func (s *Storage) CreateAccount(address common.Address) {
	s.burn.UseGas(s.jumpTable[CREATEACCOUNT].constantGas)
	s.evm.StateDB.CreateAccount(address)
}

func (s *Storage) SubBalance(address common.Address, b *big.Int) {
	s.burn.UseGas(s.jumpTable[SUBBALANCE].constantGas)

	s.evm.StateDB.SubBalance(address, b)
}

func (s *Storage) AddBalance(address common.Address, b *big.Int) {
	s.burn.UseGas(s.jumpTable[ADDBALANCE].constantGas)
	s.evm.StateDB.AddBalance(address, b)
}

func (s *Storage) GetBalance(address common.Address) *big.Int {
	s.burn.UseGas(s.jumpTable[GETBALANCE].constantGas)
	return s.evm.StateDB.GetBalance(address)
}

func (s *Storage) GetNonce(address common.Address) uint64 {
	s.burn.UseGas(s.jumpTable[GETNONCE].constantGas)
	return s.evm.StateDB.GetNonce(address)
}

func (s *Storage) SetNonce(address common.Address, u uint64) {
	s.burn.UseGas(s.jumpTable[SETNONCE].constantGas)
	s.evm.StateDB.SetNonce(address, u)
}

func (s *Storage) GetCodeHash(address common.Address) common.Hash {
	s.burn.UseGas(s.jumpTable[GETCODEHASH].constantGas)
	return s.evm.StateDB.GetCodeHash(address)
}

func (s *Storage) GetCode(address common.Address) []byte {
	s.burn.UseGas(s.jumpTable[GETCODE].constantGas)
	return s.evm.StateDB.GetCode(address)
}

func (s *Storage) SetCode(address common.Address, bytes []byte) {
	s.burn.UseGas(s.jumpTable[SETCODE].constantGas)
	s.burn.UseGas(s.jumpTable[SETCODE].dynamicGas(uint64(len(bytes))))

	s.evm.StateDB.SetCode(address, bytes)
}

func (s *Storage) GetCodeSize(address common.Address) int {
	s.burn.UseGas(s.jumpTable[GETCODESIZE].constantGas)
	return s.evm.StateDB.GetCodeSize(address)
}

func (s *Storage) GetState(address common.Address, bytes []byte) []byte {
	s.burn.UseGas(s.jumpTable[GETSTATE].constantGas)

	return s.evm.StateDB.GetState(address, bytes)
}

func (s *Storage) SetState(address common.Address, bytes []byte, bytes2 []byte) {
	s.burn.UseGas(s.jumpTable[SETSTATE].constantGas)
	s.burn.UseGas(s.jumpTable[SETSTATE].dynamicGas(s.evm, address, bytes, bytes2))
	s.evm.StateDB.SetState(address, bytes, bytes2)
}

func (s *Storage) AddLog(log *types.Log) {
	s.burn.UseGas(s.jumpTable[ADDLOG].constantGas)
	s.burn.UseGas(s.jumpTable[ADDLOG].dynamicGas(log))
	s.evm.StateDB.AddLog(log)
}
