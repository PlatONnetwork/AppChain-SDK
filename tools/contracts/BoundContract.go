package contracts

import (
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
)

type BoundContract struct {
	abi      *abi.ABI
	contract *vm.Contract
	evm      *vm.EVM
}

func (c *BoundContract) Caller(to common.Address, results *[]interface{}, method string, params ...interface{}) error {
	if results == nil {
		results = new([]interface{})
	}
	// Pack the input, call and unpack the results
	input, err := c.abi.Pack(method, params...)
	if err != nil {
		return err
	}
	output, err := contracts.Call(c.evm, c.contract, to, input, c.contract.Gas, c.contract.Value())
	if err != nil {
		return err
	}
	if len(*results) == 0 {
		res, err := c.abi.Unpack(method, output)
		*results = res
		return err
	}
	res := *results
	return c.abi.UnpackIntoInterface(res[0], method, output)
}

func (c *BoundContract) DelegateCaller(to common.Address, results *[]interface{}, method string, params ...interface{}) error {
	if results == nil {
		results = new([]interface{})
	}
	// Pack the input, call and unpack the results
	input, err := c.abi.Pack(method, params...)
	if err != nil {
		return err
	}
	output, err := contracts.DelegateCall(c.evm, c.contract, to, input, c.contract.Gas)
	if err != nil {
		return err
	}
	if len(*results) == 0 {
		res, err := c.abi.Unpack(method, output)
		*results = res
		return err
	}
	res := *results
	return c.abi.UnpackIntoInterface(res[0], method, output)
}
