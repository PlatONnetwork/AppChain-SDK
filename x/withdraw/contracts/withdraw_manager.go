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
	ABI    = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"constants\",\"name\":\"recipient\",\"type\":\"constants\"},{\"indexed\":false,\"internalType\":\"constants\",\"name\":\"depositor\",\"type\":\"constants\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"L2MintableCoinDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"constants\",\"name\":\"recipient\",\"type\":\"constants\"},{\"indexed\":false,\"internalType\":\"constants\",\"name\":\"withdrawer\",\"type\":\"constants\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"L2MintableCoinWithdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"constants\",\"name\":\"recipient\",\"type\":\"constants\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"constants\",\"name\":\"sender\",\"type\":\"constants\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onStateReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	Abi, _ = abi.JSON(strings.NewReader(ABI))
)

func (c *WithdrawManager) Run(input []byte) ([]byte, error) {
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
func (c *WithdrawManager) initMethodEntry() {

	c.methodEntry = map[string]func([]byte) ([]byte, error){

		"47e7ef24": c.DepositEntry,
		"eeb49945": c.OnStateReceiveEntry,
	}

}

func (c *WithdrawManager) DepositEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["deposit"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.Deposit(*abi.ConvertType(args[0], new(common.Address)).(*common.Address), *abi.ConvertType(args[1], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *WithdrawManager) OnStateReceiveEntry(input []byte) ([]byte, error) {

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

func (c *WithdrawManager) EmitL2MintableCoinDepositEvent(recipient common.Address, depositor common.Address, amount *big.Int) (*types.Log, error) {
	event := c.abi.Events["L2MintableCoinDeposit"]
	hashes, err := abi.PackTopics(event.Inputs, recipient, depositor, amount)
	if err != nil {
		return nil, err
	}
	data, err := event.Inputs.Pack(recipient, depositor, amount)
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

func (c *WithdrawManager) EmitL2MintableCoinWithdrawEvent(recipient common.Address, withdrawer common.Address, amount *big.Int) (*types.Log, error) {
	event := c.abi.Events["L2MintableCoinWithdraw"]
	hashes, err := abi.PackTopics(event.Inputs, recipient, withdrawer, amount)
	if err != nil {
		return nil, err
	}
	data, err := event.Inputs.Pack(recipient, withdrawer, amount)
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
