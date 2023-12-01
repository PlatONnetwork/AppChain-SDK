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

type WithdrawManagerCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewWithdrawManagerCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*WithdrawManagerCaller, error) {
	s := &WithdrawManagerCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *WithdrawManagerCaller) Deposit(recipient common.Address, amount *big.Int) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "deposit", recipient, amount)

	if err != nil {
		return err
	}

	return err

}

func (c *WithdrawManagerCaller) OnStateReceive(id *big.Int, sender common.Address, data []byte) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "onStateReceive", id, sender, data)

	if err != nil {
		return err
	}

	return err

}

type WithdrawManagerDelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewWithdrawManagerDelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*WithdrawManagerDelegateCaller, error) {
	s := &WithdrawManagerDelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *WithdrawManagerDelegateCaller) Deposit(recipient common.Address, amount *big.Int) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "deposit", recipient, amount)

	if err != nil {
		return err
	}

	return err

}

func (c *WithdrawManagerDelegateCaller) OnStateReceive(id *big.Int, sender common.Address, data []byte) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "onStateReceive", id, sender, data)

	if err != nil {
		return err
	}

	return err

}
