package contracts

import (
	"encoding/hex"
	"errors"
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

var (
	ABI    = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"constants[]\",\"name\":\"validators\",\"type\":\"constants[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"BlockReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"constants\",\"name\":\"validator\",\"type\":\"constants\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"constants\",\"name\":\"caller\",\"type\":\"constants\"}],\"name\":\"DelegaterRewardWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"constants[]\",\"name\":\"validators\",\"type\":\"constants[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"EpochReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"}],\"name\":\"RewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"constants\",\"name\":\"validator\",\"type\":\"constants\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"constants\",\"name\":\"caller\",\"type\":\"constants\"}],\"name\":\"ValidatorRewardWithdrawal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"paidRewardPerEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"constants\",\"name\":\"validator\",\"type\":\"constants\"}],\"name\":\"pendingDelegaterRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"constants\",\"name\":\"validator\",\"type\":\"constants\"}],\"name\":\"pendingValidatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"constants\",\"name\":\"validator\",\"type\":\"constants\"}],\"name\":\"withdrawDelegaterReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"constants\",\"name\":\"validator\",\"type\":\"constants\"}],\"name\":\"withdrawValidatorReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	Abi, _ = abi.JSON(strings.NewReader(ABI))
)

func (c *RewardManager) Run(input []byte) ([]byte, error) {
	if len(input) < 4 {
		return nil, errors.New("input too short")
	}
	id := input[0:4]
	entry, ok := c.methodEntry[hex.EncodeToString(id)]
	if !ok {
		if c.fallback != nil {
			return c.fallback(input)
		}
		return nil, errors.New("methods not found")
	}
	return entry(input[4:])
}
func (c *RewardManager) initMethodEntry() {

	c.methodEntry = map[string]func([]byte) ([]byte, error){
		"07358b99": c.PaidRewardPerEpochEntry,
		"8b3f0301": c.PendingDelegaterRewardsEntry,
		"a617627c": c.PendingValidatorRewardsEntry,

		"eecd58a1": c.WithdrawDelegaterRewardEntry,
		"5bc0f928": c.WithdrawValidatorRewardEntry,
	}

}

func (c *RewardManager) PaidRewardPerEpochEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["paidRewardPerEpoch"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.PaidRewardPerEpoch(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *RewardManager) PendingDelegaterRewardsEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["pendingDelegaterRewards"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.PendingDelegaterRewards(*abi.ConvertType(args[0], new(common.Address)).(*common.Address))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *RewardManager) PendingValidatorRewardsEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["pendingValidatorRewards"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.PendingValidatorRewards(*abi.ConvertType(args[0], new(common.Address)).(*common.Address))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *RewardManager) WithdrawDelegaterRewardEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["withdrawDelegaterReward"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.WithdrawDelegaterReward(*abi.ConvertType(args[0], new(common.Address)).(*common.Address))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *RewardManager) WithdrawValidatorRewardEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["withdrawValidatorReward"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.WithdrawValidatorReward(*abi.ConvertType(args[0], new(common.Address)).(*common.Address))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *RewardManager) EmitBlockRewardEvent(epochId *big.Int, validators []common.Address, amounts []*big.Int) (*types.Log, error) {
	event := c.abi.Events["BlockReward"]
	hashes, err := abi.PackTopics(event.Inputs, epochId, validators, amounts)
	if err != nil {
		return nil, err
	}
	data, err := event.Inputs.Pack(epochId, validators, amounts)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address:     c.contract.Address(),
		Topics:      hashes,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	}, nil
}

func (c *RewardManager) EmitDelegaterRewardWithdrawalEvent(validator common.Address, amount *big.Int, caller common.Address) (*types.Log, error) {
	event := c.abi.Events["DelegaterRewardWithdrawal"]
	hashes, err := abi.PackTopics(event.Inputs, validator, amount, caller)
	if err != nil {
		return nil, err
	}
	data, err := event.Inputs.Pack(validator, amount, caller)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address:     c.contract.Address(),
		Topics:      hashes,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	}, nil
}

func (c *RewardManager) EmitEpochRewardEvent(epochId *big.Int, validators []common.Address, amounts []*big.Int) (*types.Log, error) {
	event := c.abi.Events["EpochReward"]
	hashes, err := abi.PackTopics(event.Inputs, epochId, validators, amounts)
	if err != nil {
		return nil, err
	}
	data, err := event.Inputs.Pack(epochId, validators, amounts)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address:     c.contract.Address(),
		Topics:      hashes,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	}, nil
}

func (c *RewardManager) EmitRewardDistributedEvent(epochId *big.Int, totalReward *big.Int) (*types.Log, error) {
	event := c.abi.Events["RewardDistributed"]
	hashes, err := abi.PackTopics(event.Inputs, epochId, totalReward)
	if err != nil {
		return nil, err
	}
	data, err := event.Inputs.Pack(epochId, totalReward)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address:     c.contract.Address(),
		Topics:      hashes,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	}, nil
}

func (c *RewardManager) EmitValidatorRewardWithdrawalEvent(validator common.Address, amount *big.Int, caller common.Address) (*types.Log, error) {
	event := c.abi.Events["ValidatorRewardWithdrawal"]
	hashes, err := abi.PackTopics(event.Inputs, validator, amount, caller)
	if err != nil {
		return nil, err
	}
	data, err := event.Inputs.Pack(validator, amount, caller)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address:     c.contract.Address(),
		Topics:      hashes,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	}, nil
}
