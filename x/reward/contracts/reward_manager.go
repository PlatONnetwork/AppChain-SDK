package contracts

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"math/big"
	"runtime/debug"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_              = vm.EVM{}
	_              = errors.New
	_              = big.NewInt
	_              = strings.NewReader
	_              = platon.NotFound
	_              = bind.Bind
	_              = common.Big1
	_              = math.ReadBits
	_              = binary.BigEndian
	_              = types.BloomLookup
	_              = event.NewSubscription
	versionKey     = []byte("__version")
	createBlockKey = []byte("__createBlock")
)

var (
	ABI    = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"BlockReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"DelegatorRewardWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"EpochReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"}],\"name\":\"RewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"ValidatorRewardWithdrawal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"paidRewardPerEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"pendingDelegatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"pendingValidatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawDelegatorRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawValidatorRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	Abi, _ = abi.JSON(strings.NewReader(ABI))
)

func (c *RewardManager) Run(input []byte) (ret []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Trace(string(debug.Stack()))
			switch e := r.(type) {
			case error:
				if r, ok := e.(*typesdk.RevertError); ok {
					ret, err = r.ReturnData, vm.ErrExecutionReverted
				} else {
					ret, err = nil, e
				}
			default:
				ret, err = typesdk.UndefinedError, vm.ErrExecutionReverted
			}
		}
	}()
	if err := c.loadMethodABI(); err != nil {
		return nil, errors.New("load version failed")
	}

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
	ret, err = entry(input[4:])
	if err != nil {
		log.Trace("Execute failed", "err", err)
	}
	return ret, err
}

func (c *RewardManager) initABI() {
	V0 := uint64(0)
	c.abis[V0] = &Abi
}

func (c *RewardManager) initMethodEntry() {

	methodEntry := map[string]func([]byte) ([]byte, error){
		"07358b99": c.PaidRewardPerEpochEntry,
		"129e656a": c.PendingDelegatorRewardsEntry,
		"a617627c": c.PendingValidatorRewardsEntry,

		"1095cf98": c.WithdrawDelegatorRewardsEntry,
		"91b28216": c.WithdrawValidatorRewardsEntry,
	}
	V0 := uint64(0)
	c.methodEntries[V0] = methodEntry

}
func (c *RewardManager) loadMethodABI() error {
	version := c.GetVersion()
	entries, ok := c.methodEntries[version]
	if !ok {
		return errors.New("unknown version")
	}
	c.methodEntry = entries
	abi, ok := c.abis[version]
	if !ok {
		return errors.New("unknown version")
	}
	c.abi = abi
	return nil
}
func (c *RewardManager) InitGenesis(blockNumber uint64) {
	c.SetCreateBlock(blockNumber)
	c.stateDb.SetCode(c.contract.Address(), []byte("code"))
}

func (c *RewardManager) SetCreateBlock(blockNumber uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], blockNumber)
	c.evm.StateDB.SetState(c.contract.Address(), createBlockKey, data[:])
}

func (c *RewardManager) GetCreateBlock() uint64 {
	blockNumber := c.evm.StateDB.GetState(c.contract.Address(), createBlockKey)
	if len(blockNumber) == 0 {
		return math.MaxUint64
	}
	return binary.BigEndian.Uint64(blockNumber)
}

func (c *RewardManager) SetVersion(version uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], version)
	c.evm.StateDB.SetState(c.contract.Address(), versionKey, data[:])
}

func (c *RewardManager) GetVersion() uint64 {
	version := c.evm.StateDB.GetState(c.contract.Address(), versionKey)
	if len(version) == 0 {
		return 0
	}
	return binary.BigEndian.Uint64(version)
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

func (c *RewardManager) PendingDelegatorRewardsEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["pendingDelegatorRewards"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.PendingDelegatorRewards(*abi.ConvertType(args[0], new(common.Address)).(*common.Address), *abi.ConvertType(args[1], new(common.Address)).(*common.Address))
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

func (c *RewardManager) WithdrawDelegatorRewardsEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["withdrawDelegatorRewards"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.WithdrawDelegatorRewards(*abi.ConvertType(args[0], new(common.Address)).(*common.Address))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *RewardManager) WithdrawValidatorRewardsEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["withdrawValidatorRewards"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.WithdrawValidatorRewards(*abi.ConvertType(args[0], new(common.Address)).(*common.Address))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *RewardManager) BlockRewardEvent(epochId *big.Int, validators []common.Address, amounts []*big.Int) (*types.Log, error) {
	event := c.abi.Events["BlockReward"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, epochId, validators, amounts)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, epochId, validators, amounts)
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
func (c *RewardManager) EmitBlockRewardEvent(epochId *big.Int, validators []common.Address, amounts []*big.Int) {
	log, err := c.BlockRewardEvent(epochId, validators, amounts)
	contracts.Require(err == nil, "RewardManager: emit BlockReward event failed")
	c.stateDb.AddLog(log)
}

func (c *RewardManager) DelegatorRewardWithdrawalEvent(validator common.Address, amount *big.Int, caller common.Address) (*types.Log, error) {
	event := c.abi.Events["DelegatorRewardWithdrawal"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, validator, amount, caller)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, validator, amount, caller)
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
func (c *RewardManager) EmitDelegatorRewardWithdrawalEvent(validator common.Address, amount *big.Int, caller common.Address) {
	log, err := c.DelegatorRewardWithdrawalEvent(validator, amount, caller)
	contracts.Require(err == nil, "RewardManager: emit DelegatorRewardWithdrawal event failed")
	c.stateDb.AddLog(log)
}

func (c *RewardManager) EpochRewardEvent(epochId *big.Int, validators []common.Address, amounts []*big.Int) (*types.Log, error) {
	event := c.abi.Events["EpochReward"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, epochId, validators, amounts)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, epochId, validators, amounts)
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
func (c *RewardManager) EmitEpochRewardEvent(epochId *big.Int, validators []common.Address, amounts []*big.Int) {
	log, err := c.EpochRewardEvent(epochId, validators, amounts)
	contracts.Require(err == nil, "RewardManager: emit EpochReward event failed")
	c.stateDb.AddLog(log)
}

func (c *RewardManager) RewardDistributedEvent(epochId *big.Int, totalReward *big.Int) (*types.Log, error) {
	event := c.abi.Events["RewardDistributed"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, epochId, totalReward)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, epochId, totalReward)
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
func (c *RewardManager) EmitRewardDistributedEvent(epochId *big.Int, totalReward *big.Int) {
	log, err := c.RewardDistributedEvent(epochId, totalReward)
	contracts.Require(err == nil, "RewardManager: emit RewardDistributed event failed")
	c.stateDb.AddLog(log)
}

func (c *RewardManager) ValidatorRewardWithdrawalEvent(validator common.Address, amount *big.Int, caller common.Address) (*types.Log, error) {
	event := c.abi.Events["ValidatorRewardWithdrawal"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, validator, amount, caller)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, validator, amount, caller)
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
func (c *RewardManager) EmitValidatorRewardWithdrawalEvent(validator common.Address, amount *big.Int, caller common.Address) {
	log, err := c.ValidatorRewardWithdrawalEvent(validator, amount, caller)
	contracts.Require(err == nil, "RewardManager: emit ValidatorRewardWithdrawal event failed")
	c.stateDb.AddLog(log)
}
