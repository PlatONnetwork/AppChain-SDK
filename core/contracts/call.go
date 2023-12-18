package contracts

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"math/big"
)

func Call(evm *vm.EVM, contract *vm.Contract, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, err error) {
	ret, returnGas, err := evm.Call(contract, addr, input, gas, value)
	contract.UseGas(gas - returnGas)
	return ret, err
}

func DelegateCall(evm *vm.EVM, contract *vm.Contract, addr common.Address, input []byte, gas uint64) (ret []byte, err error) {
	ret, returnGas, err := evm.DelegateCall(contract, addr, input, gas)
	contract.UseGas(gas - returnGas)
	return ret, err
}

func StaticCall(evm *vm.EVM, contract *vm.Contract, addr common.Address, input []byte, gas uint64) (ret []byte, err error) {
	ret, returnGas, err := evm.StaticCall(contract, addr, input, gas)
	contract.UseGas(gas - returnGas)
	return ret, err
}
