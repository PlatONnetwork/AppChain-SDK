package contracts

import (
	"errors"
	"math/big"
	"strings"

	"github.com/PlatONnetwork/AppChain-SDK/tools/contracts"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = typesdk.RevertError{}
	_ = vm.EVM{}
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = platon.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

type CounterCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewCounterCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*CounterCaller, error) {
	s := &CounterCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *CounterCaller) Count() (uint64, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "count")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (c *CounterCaller) Name() (string, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (c *CounterCaller) Incr() error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "incr")

	if err != nil {
		return err
	}

	return err

}

type CounterDelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewCounterDelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*CounterDelegateCaller, error) {
	s := &CounterDelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *CounterDelegateCaller) Count() (uint64, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "count")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (c *CounterDelegateCaller) Name() (string, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (c *CounterDelegateCaller) Incr() error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "incr")

	if err != nil {
		return err
	}

	return err

}
