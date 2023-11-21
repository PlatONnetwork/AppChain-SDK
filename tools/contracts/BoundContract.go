package contracts

import (
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
)

type BoundContract struct {
	Abi      *abi.ABI
	Contract *vm.Contract
	Evm      *vm.EVM
}

func (c *BoundContract) Caller(to common.Address, results *[]interface{}, method string, params ...interface{}) error {
	if results == nil {
		results = new([]interface{})
	}
	// Pack the input, call and unpack the results
	input, err := c.Abi.Pack(method, params...)
	if err != nil {
		return err
	}
	output, err := contracts.Call(c.Evm, c.Contract, to, input, c.Contract.Gas, c.Contract.Value())
	if err != nil {
		return err
	}
	if len(*results) == 0 {
		res, err := c.Abi.Unpack(method, output)
		*results = res
		return err
	}
	res := *results
	return c.Abi.UnpackIntoInterface(res[0], method, output)
}

func (c *BoundContract) DelegateCaller(to common.Address, results *[]interface{}, method string, params ...interface{}) error {
	if results == nil {
		results = new([]interface{})
	}
	// Pack the input, call and unpack the results
	input, err := c.Abi.Pack(method, params...)
	if err != nil {
		return err
	}
	output, err := contracts.DelegateCall(c.Evm, c.Contract, to, input, c.Contract.Gas)
	if err != nil {
		return err
	}
	if len(*results) == 0 {
		res, err := c.Abi.Unpack(method, output)
		*results = res
		return err
	}
	res := *results
	return c.Abi.UnpackIntoInterface(res[0], method, output)
}
