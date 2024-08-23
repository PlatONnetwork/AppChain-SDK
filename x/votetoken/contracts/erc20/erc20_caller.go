package erc20

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

type ERC20Caller struct {
	contracts.BoundContract
	to common.Address
}

func NewERC20Caller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*ERC20Caller, error) {
	s := &ERC20Caller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *ERC20Caller) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20Caller) BalanceOf(account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20Caller) Decimals() (uint8, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (c *ERC20Caller) Name() (string, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (c *ERC20Caller) Symbol() (string, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (c *ERC20Caller) TotalSupply() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20Caller) Approve(spender common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "approve", spender, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *ERC20Caller) Burn(account common.Address, amount *big.Int) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "burn", account, amount)

	if err != nil {
		return err
	}

	return err

}

func (c *ERC20Caller) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "decreaseAllowance", spender, subtractedValue)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *ERC20Caller) IncreaseAllowance(spender common.Address, addedValue *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "increaseAllowance", spender, addedValue)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *ERC20Caller) Mint(account common.Address, amount *big.Int) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "mint", account, amount)

	if err != nil {
		return err
	}

	return err

}

func (c *ERC20Caller) Transfer(recipient common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "transfer", recipient, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *ERC20Caller) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "transferFrom", sender, recipient, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

type ERC20DelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewERC20DelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*ERC20DelegateCaller, error) {
	s := &ERC20DelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *ERC20DelegateCaller) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20DelegateCaller) BalanceOf(account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20DelegateCaller) Decimals() (uint8, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (c *ERC20DelegateCaller) Name() (string, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (c *ERC20DelegateCaller) Symbol() (string, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (c *ERC20DelegateCaller) TotalSupply() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20DelegateCaller) Approve(spender common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "approve", spender, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *ERC20DelegateCaller) Burn(account common.Address, amount *big.Int) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "burn", account, amount)

	if err != nil {
		return err
	}

	return err

}

func (c *ERC20DelegateCaller) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "decreaseAllowance", spender, subtractedValue)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *ERC20DelegateCaller) IncreaseAllowance(spender common.Address, addedValue *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "increaseAllowance", spender, addedValue)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *ERC20DelegateCaller) Mint(account common.Address, amount *big.Int) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "mint", account, amount)

	if err != nil {
		return err
	}

	return err

}

func (c *ERC20DelegateCaller) Transfer(recipient common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "transfer", recipient, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *ERC20DelegateCaller) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "transferFrom", sender, recipient, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}
