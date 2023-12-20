package contracts

import (
	"encoding/hex"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
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

// DelegationInfo is an auto generated low-level Go binding around an user-defined struct.
type DelegationInfo struct {
	ValidatorAddr common.Address
	DelegatorAddr common.Address
	Amount        *big.Int
	StakeEpoch    *big.Int
	DelegateEpoch *big.Int
}

// ValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type ValidatorInfo struct {
	ValidatorAddr  common.Address
	Owner          common.Address
	StakeAmount    *big.Int
	DelegateAmount *big.Int
	CommissionRate *big.Int
	Status         *big.Int
	Epoch          *big.Int
	StakeIndex     *big.Int
	PubKey         []byte
	BlsKey         []byte
}

var (
	ABI    = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegateWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegateWithdrawalRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exitId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawalRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnDelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"UpdateValidatorStatus\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getDelegationsWithValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegateEpoch\",\"type\":\"uint256\"}],\"internalType\":\"structDelegationInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"periodType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"getValidatorAddrs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"start\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegateAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structValidatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"getValidatorsWithAddr\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegateAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structValidatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onStateReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"pendingWithdrawalsOfDelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"pendingWithdrawalsOfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"validatorIndexs\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signatues\",\"type\":\"bytes\"}],\"name\":\"verifyAggregateSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signatues\",\"type\":\"bytes\"}],\"name\":\"verifyAggregateSignatureByValidators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"withdrawableOfDelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawableOfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
	Abi, _ = abi.JSON(strings.NewReader(ABI))
)

func (c *StakeHandler) Run(input []byte) (ret []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
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
func (c *StakeHandler) initMethodEntry() {

	c.methodEntry = map[string]func([]byte) ([]byte, error){
		"abec6a74": c.GetDelegationsWithValidatorEntry,
		"23bc38f4": c.GetValidatorAddrsEntry,
		"229006c5": c.GetValidatorsEntry,
		"8f1956f6": c.GetValidatorsWithAddrEntry,
		"b445f494": c.PendingWithdrawalsOfDelegateEntry,
		"46d8c653": c.PendingWithdrawalsOfStakeEntry,
		"cce6860b": c.VerifyAggregateSignatureEntry,
		"34ccc389": c.VerifyAggregateSignatureByValidatorsEntry,
		"a315ebb6": c.WithdrawableOfDelegateEntry,
		"68a11156": c.WithdrawableOfStakeEntry,

		"eeb49945": c.OnStateReceiveEntry,
		"2da25de3": c.SlashEntry,
		"4d99dd16": c.UndelegateEntry,
		"c2a672e0": c.UnstakeEntry,
		"b4065e75": c.WithdrawUndelegateEntry,
		"c76d485f": c.WithdrawUnstakeEntry,
	}

}

func (c *StakeHandler) GetDelegationsWithValidatorEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["getDelegationsWithValidator"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.GetDelegationsWithValidator(*abi.ConvertType(args[0], new([]common.Address)).(*[]common.Address), *abi.ConvertType(args[1], new(common.Address)).(*common.Address))
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

func (c *StakeHandler) GetValidatorAddrsEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["getValidatorAddrs"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.GetValidatorAddrs(*abi.ConvertType(args[0], new(uint8)).(*uint8), *abi.ConvertType(args[1], new(*big.Int)).(**big.Int))
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

func (c *StakeHandler) GetValidatorsEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["getValidators"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, res1, err := c.GetValidators(*abi.ConvertType(args[0], new([]byte)).(*[]byte), *abi.ConvertType(args[1], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0, res1)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *StakeHandler) GetValidatorsWithAddrEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["getValidatorsWithAddr"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.GetValidatorsWithAddr(*abi.ConvertType(args[0], new([]common.Address)).(*[]common.Address))
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

func (c *StakeHandler) PendingWithdrawalsOfDelegateEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["pendingWithdrawalsOfDelegate"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.PendingWithdrawalsOfDelegate(*abi.ConvertType(args[0], new(common.Address)).(*common.Address), *abi.ConvertType(args[1], new(common.Address)).(*common.Address))
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

func (c *StakeHandler) PendingWithdrawalsOfStakeEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["pendingWithdrawalsOfStake"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.PendingWithdrawalsOfStake(*abi.ConvertType(args[0], new(common.Address)).(*common.Address))
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

func (c *StakeHandler) VerifyAggregateSignatureEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["verifyAggregateSignature"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.VerifyAggregateSignature(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int), *abi.ConvertType(args[1], new([]*big.Int)).(*[]*big.Int), *abi.ConvertType(args[2], new(common.Hash)).(*common.Hash), *abi.ConvertType(args[3], new([]byte)).(*[]byte))
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

func (c *StakeHandler) VerifyAggregateSignatureByValidatorsEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["verifyAggregateSignatureByValidators"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.VerifyAggregateSignatureByValidators(*abi.ConvertType(args[0], new([]common.Address)).(*[]common.Address), *abi.ConvertType(args[1], new(common.Hash)).(*common.Hash), *abi.ConvertType(args[2], new([]byte)).(*[]byte))
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

func (c *StakeHandler) WithdrawableOfDelegateEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["withdrawableOfDelegate"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.WithdrawableOfDelegate(*abi.ConvertType(args[0], new(common.Address)).(*common.Address), *abi.ConvertType(args[1], new(common.Address)).(*common.Address))
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

func (c *StakeHandler) WithdrawableOfStakeEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["withdrawableOfStake"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.WithdrawableOfStake(*abi.ConvertType(args[0], new(common.Address)).(*common.Address))
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

func (c *StakeHandler) OnStateReceiveEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["onStateReceive"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.OnStateReceive(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int), *abi.ConvertType(args[1], new(common.Address)).(*common.Address), *abi.ConvertType(args[2], new([]byte)).(*[]byte))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *StakeHandler) SlashEntry(input []byte) ([]byte, error) {

	var err error

	err = c.Slash()
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *StakeHandler) UndelegateEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["undelegate"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.Undelegate(*abi.ConvertType(args[0], new(common.Address)).(*common.Address), *abi.ConvertType(args[1], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *StakeHandler) UnstakeEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["unstake"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.Unstake(*abi.ConvertType(args[0], new(common.Address)).(*common.Address), *abi.ConvertType(args[1], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *StakeHandler) WithdrawUndelegateEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["withdrawUndelegate"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.WithdrawUndelegate(*abi.ConvertType(args[0], new(common.Address)).(*common.Address))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *StakeHandler) WithdrawUnstakeEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["withdrawUnstake"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.WithdrawUnstake(*abi.ConvertType(args[0], new(common.Address)).(*common.Address))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *StakeHandler) EmitDelegateWithdrawalEvent(delegator common.Address, validator common.Address, amount *big.Int) (*types.Log, error) {
	event := c.abi.Events["DelegateWithdrawal"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, delegator, validator, amount)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, delegator, validator, amount)
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

func (c *StakeHandler) EmitDelegateWithdrawalRegisteredEvent(delegator common.Address, validator common.Address, amount *big.Int) (*types.Log, error) {
	event := c.abi.Events["DelegateWithdrawalRegistered"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, delegator, validator, amount)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, delegator, validator, amount)
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

func (c *StakeHandler) EmitDelegationAddedEvent(delegator common.Address, validator common.Address, amount *big.Int) (*types.Log, error) {
	event := c.abi.Events["DelegationAdded"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, delegator, validator, amount)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, delegator, validator, amount)
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

func (c *StakeHandler) EmitSlashedEvent(exitId *big.Int, validators []common.Address, amounts []*big.Int) (*types.Log, error) {
	event := c.abi.Events["Slashed"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, exitId, validators, amounts)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, exitId, validators, amounts)
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

func (c *StakeHandler) EmitStakeAddedEvent(validator common.Address, amount *big.Int) (*types.Log, error) {
	event := c.abi.Events["StakeAdded"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, validator, amount)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, validator, amount)
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

func (c *StakeHandler) EmitStakeWithdrawalEvent(validator common.Address, amount *big.Int) (*types.Log, error) {
	event := c.abi.Events["StakeWithdrawal"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, validator, amount)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, validator, amount)
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

func (c *StakeHandler) EmitStakeWithdrawalRegisteredEvent(validator common.Address, amount *big.Int) (*types.Log, error) {
	event := c.abi.Events["StakeWithdrawalRegistered"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, validator, amount)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, validator, amount)
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

func (c *StakeHandler) EmitUnDelegatedEvent(delegator common.Address, validator common.Address, amount *big.Int) (*types.Log, error) {
	event := c.abi.Events["UnDelegated"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, delegator, validator, amount)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, delegator, validator, amount)
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

func (c *StakeHandler) EmitUnStakedEvent(validator common.Address, amount *big.Int) (*types.Log, error) {
	event := c.abi.Events["UnStaked"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, validator, amount)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, validator, amount)
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

func (c *StakeHandler) EmitUpdateValidatorStatusEvent(validator common.Address, status *big.Int) (*types.Log, error) {
	event := c.abi.Events["UpdateValidatorStatus"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, validator, status)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, validator, status)
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
