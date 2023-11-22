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

type StakeHandlerCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewStakeHandlerCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*StakeHandlerCaller, error) {
	s := &StakeHandlerCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *StakeHandlerCaller) PendingWithdrawalsOfDelegate(validator common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "pendingWithdrawalsOfDelegate", validator, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *StakeHandlerCaller) PendingWithdrawalsOfStake(validator common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "pendingWithdrawalsOfStake", validator, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *StakeHandlerCaller) WithdrawableOfDelegate(validator common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "withdrawableOfDelegate", validator, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *StakeHandlerCaller) WithdrawableOfStake(validator common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "withdrawableOfStake", validator, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *StakeHandlerCaller) CommitEpoch(id *big.Int, epoch Epoch, epochSize *big.Int) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "commitEpoch", id, epoch, epochSize)

	if err != nil {
		return err
	}

	return err

}

func (c *StakeHandlerCaller) OnStateReceive(id *big.Int, sender common.Address, data []byte) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "onStateReceive", id, sender, data)

	if err != nil {
		return err
	}

	return err

}

func (c *StakeHandlerCaller) Slash(validators []common.Address) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "slash", validators)

	if err != nil {
		return err
	}

	return err

}

func (c *StakeHandlerCaller) Undelegate(validator common.Address, amount *big.Int) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "undelegate", validator, amount)

	if err != nil {
		return err
	}

	return err

}

func (c *StakeHandlerCaller) Unstake(validator common.Address, amount *big.Int) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "unstake", validator, amount)

	if err != nil {
		return err
	}

	return err

}

func (c *StakeHandlerCaller) WithdrawUndelegate(validator common.Address) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "withdrawUndelegate", validator)

	if err != nil {
		return err
	}

	return err

}

func (c *StakeHandlerCaller) WithdrawUnstake(validator common.Address) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "withdrawUnstake", validator)

	if err != nil {
		return err
	}

	return err

}

type StakeHandlerDelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewStakeHandlerDelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*StakeHandlerDelegateCaller, error) {
	s := &StakeHandlerDelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *StakeHandlerDelegateCaller) PendingWithdrawalsOfDelegate(validator common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "pendingWithdrawalsOfDelegate", validator, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *StakeHandlerDelegateCaller) PendingWithdrawalsOfStake(validator common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "pendingWithdrawalsOfStake", validator, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *StakeHandlerDelegateCaller) WithdrawableOfDelegate(validator common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "withdrawableOfDelegate", validator, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *StakeHandlerDelegateCaller) WithdrawableOfStake(validator common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "withdrawableOfStake", validator, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *StakeHandlerDelegateCaller) CommitEpoch(id *big.Int, epoch Epoch, epochSize *big.Int) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "commitEpoch", id, epoch, epochSize)

	if err != nil {
		return err
	}

	return err

}

func (c *StakeHandlerDelegateCaller) OnStateReceive(id *big.Int, sender common.Address, data []byte) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "onStateReceive", id, sender, data)

	if err != nil {
		return err
	}

	return err

}

func (c *StakeHandlerDelegateCaller) Slash(validators []common.Address) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "slash", validators)

	if err != nil {
		return err
	}

	return err

}

func (c *StakeHandlerDelegateCaller) Undelegate(validator common.Address, amount *big.Int) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "undelegate", validator, amount)

	if err != nil {
		return err
	}

	return err

}

func (c *StakeHandlerDelegateCaller) Unstake(validator common.Address, amount *big.Int) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "unstake", validator, amount)

	if err != nil {
		return err
	}

	return err

}

func (c *StakeHandlerDelegateCaller) WithdrawUndelegate(validator common.Address) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "withdrawUndelegate", validator)

	if err != nil {
		return err
	}

	return err

}

func (c *StakeHandlerDelegateCaller) WithdrawUnstake(validator common.Address) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "withdrawUnstake", validator)

	if err != nil {
		return err
	}

	return err

}
