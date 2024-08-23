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

type UpgradeCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewUpgradeCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*UpgradeCaller, error) {
	s := &UpgradeCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *UpgradeCaller) GetOwner() (common.Address, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (c *UpgradeCaller) GetUpgradePlan(height uint64) ([]IUpgradePlan, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "getUpgradePlan", height)

	if err != nil {
		return *new([]IUpgradePlan), err
	}

	out0 := *abi.ConvertType(out[0], new([]IUpgradePlan)).(*[]IUpgradePlan)

	return out0, err

}

func (c *UpgradeCaller) AddUpgradePlan(plan IUpgradePlan) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "addUpgradePlan", plan)

	if err != nil {
		return err
	}

	return err

}

func (c *UpgradeCaller) SetOwner(newOwner common.Address) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "setOwner", newOwner)

	if err != nil {
		return err
	}

	return err

}

func (c *UpgradeCaller) SetUpgradePlanDone(height uint64) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "setUpgradePlanDone", height)

	if err != nil {
		return err
	}

	return err

}

type UpgradeDelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewUpgradeDelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*UpgradeDelegateCaller, error) {
	s := &UpgradeDelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *UpgradeDelegateCaller) GetOwner() (common.Address, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (c *UpgradeDelegateCaller) GetUpgradePlan(height uint64) ([]IUpgradePlan, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "getUpgradePlan", height)

	if err != nil {
		return *new([]IUpgradePlan), err
	}

	out0 := *abi.ConvertType(out[0], new([]IUpgradePlan)).(*[]IUpgradePlan)

	return out0, err

}

func (c *UpgradeDelegateCaller) AddUpgradePlan(plan IUpgradePlan) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "addUpgradePlan", plan)

	if err != nil {
		return err
	}

	return err

}

func (c *UpgradeDelegateCaller) SetOwner(newOwner common.Address) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "setOwner", newOwner)

	if err != nil {
		return err
	}

	return err

}

func (c *UpgradeDelegateCaller) SetUpgradePlanDone(height uint64) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "setUpgradePlanDone", height)

	if err != nil {
		return err
	}

	return err

}
