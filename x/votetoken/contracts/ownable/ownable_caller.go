package contracts

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/tools/contracts"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"math/big"
	"strings"
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

type OwnableCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewOwnableCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*OwnableCaller, error) {
	s := &OwnableCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *OwnableCaller) Owner() (common.Address, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (c *OwnableCaller) RenounceOwnership() error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "renounceOwnership")

	if err != nil {
		return err
	}

	return err

}

func (c *OwnableCaller) TransferOwnership(newOwner common.Address) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "transferOwnership", newOwner)

	if err != nil {
		return err
	}

	return err

}

type OwnableDelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewOwnableDelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*OwnableDelegateCaller, error) {
	s := &OwnableDelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *OwnableDelegateCaller) Owner() (common.Address, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (c *OwnableDelegateCaller) RenounceOwnership() error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "renounceOwnership")

	if err != nil {
		return err
	}

	return err

}

func (c *OwnableDelegateCaller) TransferOwnership(newOwner common.Address) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "transferOwnership", newOwner)

	if err != nil {
		return err
	}

	return err

}
