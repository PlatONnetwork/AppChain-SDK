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

type RewardManagerCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewRewardManagerCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*RewardManagerCaller, error) {
	s := &RewardManagerCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *RewardManagerCaller) PaidRewardPerEpoch(epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "paidRewardPerEpoch", epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *RewardManagerCaller) PendingDelegaterRewards(account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "pendingDelegaterRewards", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *RewardManagerCaller) PendingValidatorRewards(account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "pendingValidatorRewards", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *RewardManagerCaller) WithdrawDelegaterReward(validator common.Address) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "withdrawDelegaterReward", validator)

	if err != nil {
		return err
	}

	return err

}

func (c *RewardManagerCaller) WithdrawValidatorReward(validator common.Address) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "withdrawValidatorReward", validator)

	if err != nil {
		return err
	}

	return err

}

type RewardManagerDelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewRewardManagerDelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*RewardManagerDelegateCaller, error) {
	s := &RewardManagerDelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *RewardManagerDelegateCaller) PaidRewardPerEpoch(epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "paidRewardPerEpoch", epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *RewardManagerDelegateCaller) PendingDelegaterRewards(account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "pendingDelegaterRewards", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *RewardManagerDelegateCaller) PendingValidatorRewards(account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "pendingValidatorRewards", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *RewardManagerDelegateCaller) WithdrawDelegaterReward(validator common.Address) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "withdrawDelegaterReward", validator)

	if err != nil {
		return err
	}

	return err

}

func (c *RewardManagerDelegateCaller) WithdrawValidatorReward(validator common.Address) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "withdrawValidatorReward", validator)

	if err != nil {
		return err
	}

	return err

}
