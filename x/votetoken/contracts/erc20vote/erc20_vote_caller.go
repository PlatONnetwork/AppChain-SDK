package erc20vote

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

type ERC20VoteCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewERC20VoteCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*ERC20VoteCaller, error) {
	s := &ERC20VoteCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *ERC20VoteCaller) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20VoteCaller) BalanceOf(account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20VoteCaller) Checkpoints(account common.Address, pos uint32) (Checkpoint, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "checkpoints", account, pos)

	if err != nil {
		return *new(Checkpoint), err
	}

	out0 := *abi.ConvertType(out[0], new(Checkpoint)).(*Checkpoint)

	return out0, err

}

func (c *ERC20VoteCaller) Decimals() (uint8, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (c *ERC20VoteCaller) Delegates(account common.Address) (common.Address, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "delegates", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (c *ERC20VoteCaller) GetPastTotalSupply(blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "getPastTotalSupply", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20VoteCaller) GetPastVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "getPastVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20VoteCaller) GetVotes(account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "getVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20VoteCaller) Name() (string, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (c *ERC20VoteCaller) NumCheckpoints(account common.Address) (uint32, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "numCheckpoints", account)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (c *ERC20VoteCaller) Symbol() (string, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (c *ERC20VoteCaller) TotalSupply() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20VoteCaller) Approve(spender common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "approve", spender, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *ERC20VoteCaller) Burn(account common.Address, amount *big.Int) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "burn", account, amount)

	if err != nil {
		return err
	}

	return err

}

func (c *ERC20VoteCaller) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "decreaseAllowance", spender, subtractedValue)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *ERC20VoteCaller) Delegate(delegatee common.Address) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "delegate", delegatee)

	if err != nil {
		return err
	}

	return err

}

func (c *ERC20VoteCaller) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r common.Hash, s common.Hash) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "delegateBySig", delegatee, nonce, expiry, v, r, s)

	if err != nil {
		return err
	}

	return err

}

func (c *ERC20VoteCaller) IncreaseAllowance(spender common.Address, addedValue *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "increaseAllowance", spender, addedValue)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *ERC20VoteCaller) Mint(account common.Address, amount *big.Int) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "mint", account, amount)

	if err != nil {
		return err
	}

	return err

}

func (c *ERC20VoteCaller) Transfer(recipient common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "transfer", recipient, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *ERC20VoteCaller) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "transferFrom", sender, recipient, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

type ERC20VoteDelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewERC20VoteDelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*ERC20VoteDelegateCaller, error) {
	s := &ERC20VoteDelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *ERC20VoteDelegateCaller) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20VoteDelegateCaller) BalanceOf(account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20VoteDelegateCaller) Checkpoints(account common.Address, pos uint32) (Checkpoint, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "checkpoints", account, pos)

	if err != nil {
		return *new(Checkpoint), err
	}

	out0 := *abi.ConvertType(out[0], new(Checkpoint)).(*Checkpoint)

	return out0, err

}

func (c *ERC20VoteDelegateCaller) Decimals() (uint8, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (c *ERC20VoteDelegateCaller) Delegates(account common.Address) (common.Address, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "delegates", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (c *ERC20VoteDelegateCaller) GetPastTotalSupply(blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "getPastTotalSupply", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20VoteDelegateCaller) GetPastVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "getPastVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20VoteDelegateCaller) GetVotes(account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "getVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20VoteDelegateCaller) Name() (string, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (c *ERC20VoteDelegateCaller) NumCheckpoints(account common.Address) (uint32, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "numCheckpoints", account)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (c *ERC20VoteDelegateCaller) Symbol() (string, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (c *ERC20VoteDelegateCaller) TotalSupply() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20VoteDelegateCaller) Approve(spender common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "approve", spender, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *ERC20VoteDelegateCaller) Burn(account common.Address, amount *big.Int) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "burn", account, amount)

	if err != nil {
		return err
	}

	return err

}

func (c *ERC20VoteDelegateCaller) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "decreaseAllowance", spender, subtractedValue)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *ERC20VoteDelegateCaller) Delegate(delegatee common.Address) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "delegate", delegatee)

	if err != nil {
		return err
	}

	return err

}

func (c *ERC20VoteDelegateCaller) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r common.Hash, s common.Hash) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "delegateBySig", delegatee, nonce, expiry, v, r, s)

	if err != nil {
		return err
	}

	return err

}

func (c *ERC20VoteDelegateCaller) IncreaseAllowance(spender common.Address, addedValue *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "increaseAllowance", spender, addedValue)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *ERC20VoteDelegateCaller) Mint(account common.Address, amount *big.Int) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "mint", account, amount)

	if err != nil {
		return err
	}

	return err

}

func (c *ERC20VoteDelegateCaller) Transfer(recipient common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "transfer", recipient, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *ERC20VoteDelegateCaller) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "transferFrom", sender, recipient, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}
