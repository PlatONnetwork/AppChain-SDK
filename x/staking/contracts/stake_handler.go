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

// Epoch is an auto generated low-level Go binding around an user-defined struct.
type Epoch struct {
	StartBlock *big.Int
	EndBlock   *big.Int
	EpochRoot  [32]byte
}

var (
	ABI    = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"epochRoot\",\"type\":\"bytes32\"}],\"name\":\"NewEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exitId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegater\",\"type\":\"address\"}],\"name\":\"UnDelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"UnStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"epochRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structEpoch\",\"name\":\"epoch\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"epochSize\",\"type\":\"uint256\"}],\"name\":\"commitEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onStateReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"pendingWithdrawalsOfDelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"pendingWithdrawalsOfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawableOfDelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawableOfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
	Abi, _ = abi.JSON(strings.NewReader(ABI))
)

func (c *StakeHandler) Run(input []byte) ([]byte, error) {
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
		"a69cc41e": c.PendingWithdrawalsOfDelegateEntry,
		"46d8c653": c.PendingWithdrawalsOfStakeEntry,
		"9dd5dd41": c.WithdrawableOfDelegateEntry,
		"68a11156": c.WithdrawableOfStakeEntry,

		"dab567de": c.CommitEpochEntry,
		"eeb49945": c.OnStateReceiveEntry,
		"8b8c24c1": c.SlashEntry,
		"4d99dd16": c.UndelegateEntry,
		"2e17de78": c.UnstakeEntry,
		"c635e813": c.WithdrawUndelegateEntry,
		"01d2700a": c.WithdrawUnstakeEntry,
	}

}

func (c *StakeHandler) PendingWithdrawalsOfDelegateEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["pendingWithdrawalsOfDelegate"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.PendingWithdrawalsOfDelegate(*abi.ConvertType(args[0], new(common.Address)).(*common.Address))
	if err != nil {
		if r := err.(*typesdk.RevertError); r != nil {
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
		if r := err.(*typesdk.RevertError); r != nil {
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

	res0, err := c.WithdrawableOfDelegate(*abi.ConvertType(args[0], new(common.Address)).(*common.Address))
	if err != nil {
		if r := err.(*typesdk.RevertError); r != nil {
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
		if r := err.(*typesdk.RevertError); r != nil {
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

func (c *StakeHandler) CommitEpochEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["commitEpoch"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.CommitEpoch(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int), *abi.ConvertType(args[1], new(Epoch)).(*Epoch), *abi.ConvertType(args[2], new(*big.Int)).(**big.Int))
	if err != nil {
		if r := err.(*typesdk.RevertError); r != nil {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

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
		if r := err.(*typesdk.RevertError); r != nil {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *StakeHandler) SlashEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["slash"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.Slash(*abi.ConvertType(args[0], new([]common.Address)).(*[]common.Address))
	if err != nil {
		if r := err.(*typesdk.RevertError); r != nil {
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
		if r := err.(*typesdk.RevertError); r != nil {
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

	err = c.Unstake(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int))
	if err != nil {
		if r := err.(*typesdk.RevertError); r != nil {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *StakeHandler) WithdrawUndelegateEntry(input []byte) ([]byte, error) {

	var err error

	err = c.WithdrawUndelegate()
	if err != nil {
		if r := err.(*typesdk.RevertError); r != nil {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *StakeHandler) WithdrawUnstakeEntry(input []byte) ([]byte, error) {

	var err error

	err = c.WithdrawUnstake()
	if err != nil {
		if r := err.(*typesdk.RevertError); r != nil {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *StakeHandler) EmitNewEpochEvent(id *big.Int, startBlock *big.Int, endBlock *big.Int, epochRoot [32]byte) (*types.Log, error) {
	event := c.abi.Events["NewEpoch"]
	hashes, err := abi.PackTopics(event.Inputs, id, startBlock, endBlock, epochRoot)
	if err != nil {
		return nil, err
	}
	data, err := event.Inputs.Pack(id, startBlock, endBlock, epochRoot)
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

func (c *StakeHandler) EmitSlashedEvent(exitId *big.Int, validators []common.Address) (*types.Log, error) {
	event := c.abi.Events["Slashed"]
	hashes, err := abi.PackTopics(event.Inputs, exitId, validators)
	if err != nil {
		return nil, err
	}
	data, err := event.Inputs.Pack(exitId, validators)
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

func (c *StakeHandler) EmitUnDelegatedEvent(validator common.Address, delegater common.Address) (*types.Log, error) {
	event := c.abi.Events["UnDelegated"]
	hashes, err := abi.PackTopics(event.Inputs, validator, delegater)
	if err != nil {
		return nil, err
	}
	data, err := event.Inputs.Pack(validator, delegater)
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

func (c *StakeHandler) EmitUnStakedEvent(validator common.Address) (*types.Log, error) {
	event := c.abi.Events["UnStaked"]
	hashes, err := abi.PackTopics(event.Inputs, validator)
	if err != nil {
		return nil, err
	}
	data, err := event.Inputs.Pack(validator)
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

func (c *StakeHandler) EmitWithdrawalEvent(account common.Address, amount *big.Int) (*types.Log, error) {
	event := c.abi.Events["Withdrawal"]
	hashes, err := abi.PackTopics(event.Inputs, account, amount)
	if err != nil {
		return nil, err
	}
	data, err := event.Inputs.Pack(account, amount)
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

func (c *StakeHandler) EmitWithdrawalRegisteredEvent(account common.Address, amount *big.Int) (*types.Log, error) {
	event := c.abi.Events["WithdrawalRegistered"]
	hashes, err := abi.PackTopics(event.Inputs, account, amount)
	if err != nil {
		return nil, err
	}
	data, err := event.Inputs.Pack(account, amount)
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
