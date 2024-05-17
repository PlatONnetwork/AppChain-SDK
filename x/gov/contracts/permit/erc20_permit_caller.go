package permit

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

type ERC20PermitCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewERC20PermitCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*ERC20PermitCaller, error) {
	s := &ERC20PermitCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *ERC20PermitCaller) DOMAINSEPARATOR() (common.Hash, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new(common.Hash), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Hash)).(*common.Hash)

	return out0, err

}

func (c *ERC20PermitCaller) Nonces(owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20PermitCaller) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r common.Hash, s common.Hash) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "permit", owner, spender, value, deadline, v, r, s)

	if err != nil {
		return err
	}

	return err

}

type ERC20PermitDelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewERC20PermitDelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*ERC20PermitDelegateCaller, error) {
	s := &ERC20PermitDelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *ERC20PermitDelegateCaller) DOMAINSEPARATOR() (common.Hash, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new(common.Hash), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Hash)).(*common.Hash)

	return out0, err

}

func (c *ERC20PermitDelegateCaller) Nonces(owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *ERC20PermitDelegateCaller) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r common.Hash, s common.Hash) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "permit", owner, spender, value, deadline, v, r, s)

	if err != nil {
		return err
	}

	return err

}
