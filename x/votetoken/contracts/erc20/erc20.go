package erc20

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
	"math/big"
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
	versionKey     = []byte("version")
	createBlockKey = []byte("createBlock")
)

var (
	ABI    = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	Abi, _ = abi.JSON(strings.NewReader(ABI))
)

func (c *ERC20) Run(input []byte) (ret []byte, err error) {
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
	return entry(input[4:])
}

func (c *ERC20) initABI() {
	V0 := uint64(0)
	c.abis[V0] = &Abi
}

func (c *ERC20) initMethodEntry() {

	methodEntry := map[string]func([]byte) ([]byte, error){
		"dd62ed3e": c.AllowanceEntry,
		"70a08231": c.BalanceOfEntry,
		"313ce567": c.DecimalsEntry,
		"06fdde03": c.NameEntry,
		"95d89b41": c.SymbolEntry,
		"18160ddd": c.TotalSupplyEntry,

		"095ea7b3": c.ApproveEntry,
		"9dc29fac": c.BurnEntry,
		"a457c2d7": c.DecreaseAllowanceEntry,
		"39509351": c.IncreaseAllowanceEntry,
		"40c10f19": c.MintEntry,
		"8da5cb5b": c.OwnerEntry,
		"715018a6": c.RenounceOwnershipEntry,
		"a9059cbb": c.TransferEntry,
		"23b872dd": c.TransferFromEntry,
		"f2fde38b": c.TransferOwnershipEntry,
	}
	V0 := uint64(0)
	c.methodEntries[V0] = methodEntry

}
func (c *ERC20) loadMethodABI() error {
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
func (c *ERC20) SetCreateBlock(blockNumber uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], blockNumber)
	c.evm.StateDB.SetState(c.contract.Address(), createBlockKey, data[:])
}

func (c *ERC20) GetCreateBlock() uint64 {
	blockNumber := c.evm.StateDB.GetState(c.contract.Address(), createBlockKey)
	if len(blockNumber) == 0 {
		return math.MaxUint64
	}
	return binary.BigEndian.Uint64(blockNumber)
}

func (c *ERC20) SetVersion(version uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], version)
	c.evm.StateDB.SetState(c.contract.Address(), versionKey, data[:])
}

func (c *ERC20) GetVersion() uint64 {
	version := c.evm.StateDB.GetState(c.contract.Address(), versionKey)
	if len(version) == 0 {
		return 0
	}
	return binary.BigEndian.Uint64(version)
}

func (c *ERC20) AllowanceEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["allowance"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.Allowance(*abi.ConvertType(args[0], new(common.Address)).(*common.Address), *abi.ConvertType(args[1], new(common.Address)).(*common.Address))
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

func (c *ERC20) BalanceOfEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["balanceOf"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.BalanceOf(*abi.ConvertType(args[0], new(common.Address)).(*common.Address))
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

func (c *ERC20) DecimalsEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["decimals"]

	var err error

	res0, err := c.Decimals()
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

func (c *ERC20) NameEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["name"]

	var err error

	res0, err := c.Name()
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

func (c *ERC20) SymbolEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["symbol"]

	var err error

	res0, err := c.Symbol()
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

func (c *ERC20) TotalSupplyEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["totalSupply"]

	var err error

	res0, err := c.TotalSupply()
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

func (c *ERC20) ApproveEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["approve"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.Approve(*abi.ConvertType(args[0], new(common.Address)).(*common.Address), *abi.ConvertType(args[1], new(*big.Int)).(**big.Int))
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

func (c *ERC20) BurnEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["burn"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.Burn(*abi.ConvertType(args[0], new(common.Address)).(*common.Address), *abi.ConvertType(args[1], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *ERC20) DecreaseAllowanceEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["decreaseAllowance"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.DecreaseAllowance(*abi.ConvertType(args[0], new(common.Address)).(*common.Address), *abi.ConvertType(args[1], new(*big.Int)).(**big.Int))
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

func (c *ERC20) IncreaseAllowanceEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["increaseAllowance"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.IncreaseAllowance(*abi.ConvertType(args[0], new(common.Address)).(*common.Address), *abi.ConvertType(args[1], new(*big.Int)).(**big.Int))
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

func (c *ERC20) MintEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["mint"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.Mint(*abi.ConvertType(args[0], new(common.Address)).(*common.Address), *abi.ConvertType(args[1], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *ERC20) OwnerEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["owner"]

	var err error

	res0, err := c.Owner()
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

func (c *ERC20) RenounceOwnershipEntry(input []byte) ([]byte, error) {

	var err error

	err = c.RenounceOwnership()
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *ERC20) TransferEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["transfer"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.Transfer(*abi.ConvertType(args[0], new(common.Address)).(*common.Address), *abi.ConvertType(args[1], new(*big.Int)).(**big.Int))
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

func (c *ERC20) TransferFromEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["transferFrom"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.TransferFrom(*abi.ConvertType(args[0], new(common.Address)).(*common.Address), *abi.ConvertType(args[1], new(common.Address)).(*common.Address), *abi.ConvertType(args[2], new(*big.Int)).(**big.Int))
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

func (c *ERC20) TransferOwnershipEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["transferOwnership"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.TransferOwnership(*abi.ConvertType(args[0], new(common.Address)).(*common.Address))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *ERC20) ApprovalEvent(owner common.Address, spender common.Address, value *big.Int) (*types.Log, error) {
	event := c.abi.Events["Approval"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, owner, spender, value)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, owner, spender, value)
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
func (c *ERC20) EmitApprovalEvent(owner common.Address, spender common.Address, value *big.Int) {
	log, err := c.ApprovalEvent(owner, spender, value)
	contracts.Require(err == nil, "ERC20: emit Approval event failed")
	c.stateDb.AddLog(log)
}

func (c *ERC20) OwnershipTransferredEvent(previousOwner common.Address, newOwner common.Address) (*types.Log, error) {
	event := c.abi.Events["OwnershipTransferred"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, previousOwner, newOwner)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, previousOwner, newOwner)
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
func (c *ERC20) EmitOwnershipTransferredEvent(previousOwner common.Address, newOwner common.Address) {
	log, err := c.OwnershipTransferredEvent(previousOwner, newOwner)
	contracts.Require(err == nil, "ERC20: emit OwnershipTransferred event failed")
	c.stateDb.AddLog(log)
}

func (c *ERC20) TransferEvent(from common.Address, to common.Address, value *big.Int) (*types.Log, error) {
	event := c.abi.Events["Transfer"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, from, to, value)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, from, to, value)
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
func (c *ERC20) EmitTransferEvent(from common.Address, to common.Address, value *big.Int) {
	log, err := c.TransferEvent(from, to, value)
	contracts.Require(err == nil, "ERC20: emit Transfer event failed")
	c.stateDb.AddLog(log)
}
