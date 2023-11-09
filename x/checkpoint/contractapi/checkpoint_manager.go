// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractapi

import (
	"errors"
	"math/big"
	"strings"

	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = platon.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ICheckpointManagerCheckpoint is an auto generated low-level Go binding around an user-defined struct.
type ICheckpointManagerCheckpoint struct {
	Epoch       uint64
	ViewNumber  uint64
	BlockNumber uint64
	EventRoot   [32]byte
	ExtendRoot  [32]byte
}

// ICheckpointManagerCheckpointMetadata is an auto generated low-level Go binding around an user-defined struct.
type ICheckpointManagerCheckpointMetadata struct {
	BlockHash               [32]byte
	BlockIndex              uint32
	CurrentValidatorSetHash [32]byte
}

// ICheckpointManagerValidator is an auto generated low-level Go binding around an user-defined struct.
type ICheckpointManagerValidator struct {
	Address common.Address
	BlsKey  [2]*big.Int
}

// ContractapiMetaData contains all meta data concerning the Contractapi contract.
var ContractapiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"INITIALIZER\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bls\",\"outputs\":[{\"internalType\":\"contractBLS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"checkpointBlockNumbers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"checkpoints\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"viewNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"eventRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extendRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentCheckpointBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentValidatorSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentValidatorSetHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentValidatorSetLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getCheckpointBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"getEventMembershipByBlockNumber\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"getEventMembershipByEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEventRootByBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBLS\",\"name\":\"newBls\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"blsKey\",\"type\":\"uint256[2]\"}],\"internalType\":\"structICheckpointManager.Validator[]\",\"name\":\"newValidatorSet\",\"type\":\"tuple[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blockIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"currentValidatorSetHash\",\"type\":\"bytes32\"}],\"internalType\":\"structICheckpointManager.CheckpointMetadata\",\"name\":\"checkpointMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"viewNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"eventRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extendRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structICheckpointManager.Checkpoint\",\"name\":\"checkpoint\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bitmap\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"blsKey\",\"type\":\"uint256[2]\"}],\"internalType\":\"structICheckpointManager.Validator[]\",\"name\":\"newValidatorSet\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractapiABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractapiMetaData.ABI instead.
var ContractapiABI = ContractapiMetaData.ABI

// Contractapi is an auto generated Go binding around an platon contract.
type Contractapi struct {
	ContractapiCaller     // Read-only binding to the contract
	ContractapiTransactor // Write-only binding to the contract
	ContractapiFilterer   // Log filterer for contract events
}

// ContractapiCaller is an auto generated read-only Go binding around an platon contract.
type ContractapiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractapiTransactor is an auto generated write-only Go binding around an platon contract.
type ContractapiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractapiFilterer is an auto generated log filtering Go binding around an platon contract events.
type ContractapiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractapiSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type ContractapiSession struct {
	Contract     *Contractapi      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractapiCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type ContractapiCallerSession struct {
	Contract *ContractapiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ContractapiTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type ContractapiTransactorSession struct {
	Contract     *ContractapiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractapiRaw is an auto generated low-level Go binding around an platon contract.
type ContractapiRaw struct {
	Contract *Contractapi // Generic contract binding to access the raw methods on
}

// ContractapiCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type ContractapiCallerRaw struct {
	Contract *ContractapiCaller // Generic read-only contract binding to access the raw methods on
}

// ContractapiTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type ContractapiTransactorRaw struct {
	Contract *ContractapiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractapi creates a new instance of Contractapi, bound to a specific deployed contract.
func NewContractapi(address common.Address, backend bind.ContractBackend) (*Contractapi, error) {
	contract, err := bindContractapi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contractapi{ContractapiCaller: ContractapiCaller{contract: contract}, ContractapiTransactor: ContractapiTransactor{contract: contract}, ContractapiFilterer: ContractapiFilterer{contract: contract}}, nil
}

// NewContractapiCaller creates a new read-only instance of Contractapi, bound to a specific deployed contract.
func NewContractapiCaller(address common.Address, caller bind.ContractCaller) (*ContractapiCaller, error) {
	contract, err := bindContractapi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractapiCaller{contract: contract}, nil
}

// NewContractapiTransactor creates a new write-only instance of Contractapi, bound to a specific deployed contract.
func NewContractapiTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractapiTransactor, error) {
	contract, err := bindContractapi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractapiTransactor{contract: contract}, nil
}

// NewContractapiFilterer creates a new log filterer instance of Contractapi, bound to a specific deployed contract.
func NewContractapiFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractapiFilterer, error) {
	contract, err := bindContractapi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractapiFilterer{contract: contract}, nil
}

// bindContractapi binds a generic wrapper to an already deployed contract.
func bindContractapi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractapiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractapi *ContractapiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractapi.Contract.ContractapiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractapi *ContractapiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractapi.Contract.ContractapiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractapi *ContractapiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractapi.Contract.ContractapiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractapi *ContractapiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractapi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractapi *ContractapiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractapi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractapi *ContractapiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractapi.Contract.contract.Transact(opts, method, params...)
}

// Bls is a free data retrieval call binding the contract method 0x95b0b027.
//
// Solidity: function bls() view returns(address)
func (_Contractapi *ContractapiCaller) Bls(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contractapi.contract.Call(opts, &out, "bls")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bls is a free data retrieval call binding the contract method 0x95b0b027.
//
// Solidity: function bls() view returns(address)
func (_Contractapi *ContractapiSession) Bls() (common.Address, error) {
	return _Contractapi.Contract.Bls(&_Contractapi.CallOpts)
}

// Bls is a free data retrieval call binding the contract method 0x95b0b027.
//
// Solidity: function bls() view returns(address)
func (_Contractapi *ContractapiCallerSession) Bls() (common.Address, error) {
	return _Contractapi.Contract.Bls(&_Contractapi.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Contractapi *ContractapiCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractapi.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Contractapi *ContractapiSession) ChainId() (*big.Int, error) {
	return _Contractapi.Contract.ChainId(&_Contractapi.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Contractapi *ContractapiCallerSession) ChainId() (*big.Int, error) {
	return _Contractapi.Contract.ChainId(&_Contractapi.CallOpts)
}

// CheckpointBlockNumbers is a free data retrieval call binding the contract method 0xe9193d2b.
//
// Solidity: function checkpointBlockNumbers(uint256 ) view returns(uint256)
func (_Contractapi *ContractapiCaller) CheckpointBlockNumbers(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contractapi.contract.Call(opts, &out, "checkpointBlockNumbers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckpointBlockNumbers is a free data retrieval call binding the contract method 0xe9193d2b.
//
// Solidity: function checkpointBlockNumbers(uint256 ) view returns(uint256)
func (_Contractapi *ContractapiSession) CheckpointBlockNumbers(arg0 *big.Int) (*big.Int, error) {
	return _Contractapi.Contract.CheckpointBlockNumbers(&_Contractapi.CallOpts, arg0)
}

// CheckpointBlockNumbers is a free data retrieval call binding the contract method 0xe9193d2b.
//
// Solidity: function checkpointBlockNumbers(uint256 ) view returns(uint256)
func (_Contractapi *ContractapiCallerSession) CheckpointBlockNumbers(arg0 *big.Int) (*big.Int, error) {
	return _Contractapi.Contract.CheckpointBlockNumbers(&_Contractapi.CallOpts, arg0)
}

// Checkpoints is a free data retrieval call binding the contract method 0xb8a24252.
//
// Solidity: function checkpoints(uint256 ) view returns(uint64 epoch, uint64 viewNumber, uint64 blockNumber, bytes32 eventRoot, bytes32 extendRoot)
func (_Contractapi *ContractapiCaller) Checkpoints(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Epoch       uint64
	ViewNumber  uint64
	BlockNumber uint64
	EventRoot   [32]byte
	ExtendRoot  [32]byte
}, error) {
	var out []interface{}
	err := _Contractapi.contract.Call(opts, &out, "checkpoints", arg0)

	outstruct := new(struct {
		Epoch       uint64
		ViewNumber  uint64
		BlockNumber uint64
		EventRoot   [32]byte
		ExtendRoot  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epoch = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ViewNumber = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.BlockNumber = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.EventRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.ExtendRoot = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xb8a24252.
//
// Solidity: function checkpoints(uint256 ) view returns(uint64 epoch, uint64 viewNumber, uint64 blockNumber, bytes32 eventRoot, bytes32 extendRoot)
func (_Contractapi *ContractapiSession) Checkpoints(arg0 *big.Int) (struct {
	Epoch       uint64
	ViewNumber  uint64
	BlockNumber uint64
	EventRoot   [32]byte
	ExtendRoot  [32]byte
}, error) {
	return _Contractapi.Contract.Checkpoints(&_Contractapi.CallOpts, arg0)
}

// Checkpoints is a free data retrieval call binding the contract method 0xb8a24252.
//
// Solidity: function checkpoints(uint256 ) view returns(uint64 epoch, uint64 viewNumber, uint64 blockNumber, bytes32 eventRoot, bytes32 extendRoot)
func (_Contractapi *ContractapiCallerSession) Checkpoints(arg0 *big.Int) (struct {
	Epoch       uint64
	ViewNumber  uint64
	BlockNumber uint64
	EventRoot   [32]byte
	ExtendRoot  [32]byte
}, error) {
	return _Contractapi.Contract.Checkpoints(&_Contractapi.CallOpts, arg0)
}

// CurrentCheckpointBlockNumber is a free data retrieval call binding the contract method 0xe416d677.
//
// Solidity: function currentCheckpointBlockNumber() view returns(uint256)
func (_Contractapi *ContractapiCaller) CurrentCheckpointBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractapi.contract.Call(opts, &out, "currentCheckpointBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentCheckpointBlockNumber is a free data retrieval call binding the contract method 0xe416d677.
//
// Solidity: function currentCheckpointBlockNumber() view returns(uint256)
func (_Contractapi *ContractapiSession) CurrentCheckpointBlockNumber() (*big.Int, error) {
	return _Contractapi.Contract.CurrentCheckpointBlockNumber(&_Contractapi.CallOpts)
}

// CurrentCheckpointBlockNumber is a free data retrieval call binding the contract method 0xe416d677.
//
// Solidity: function currentCheckpointBlockNumber() view returns(uint256)
func (_Contractapi *ContractapiCallerSession) CurrentCheckpointBlockNumber() (*big.Int, error) {
	return _Contractapi.Contract.CurrentCheckpointBlockNumber(&_Contractapi.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contractapi *ContractapiCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractapi.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contractapi *ContractapiSession) CurrentEpoch() (*big.Int, error) {
	return _Contractapi.Contract.CurrentEpoch(&_Contractapi.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contractapi *ContractapiCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Contractapi.Contract.CurrentEpoch(&_Contractapi.CallOpts)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address _address)
func (_Contractapi *ContractapiCaller) CurrentValidatorSet(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contractapi.contract.Call(opts, &out, "currentValidatorSet", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address _address)
func (_Contractapi *ContractapiSession) CurrentValidatorSet(arg0 *big.Int) (common.Address, error) {
	return _Contractapi.Contract.CurrentValidatorSet(&_Contractapi.CallOpts, arg0)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address _address)
func (_Contractapi *ContractapiCallerSession) CurrentValidatorSet(arg0 *big.Int) (common.Address, error) {
	return _Contractapi.Contract.CurrentValidatorSet(&_Contractapi.CallOpts, arg0)
}

// CurrentValidatorSetHash is a free data retrieval call binding the contract method 0xf896f1a5.
//
// Solidity: function currentValidatorSetHash() view returns(bytes32)
func (_Contractapi *ContractapiCaller) CurrentValidatorSetHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contractapi.contract.Call(opts, &out, "currentValidatorSetHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentValidatorSetHash is a free data retrieval call binding the contract method 0xf896f1a5.
//
// Solidity: function currentValidatorSetHash() view returns(bytes32)
func (_Contractapi *ContractapiSession) CurrentValidatorSetHash() ([32]byte, error) {
	return _Contractapi.Contract.CurrentValidatorSetHash(&_Contractapi.CallOpts)
}

// CurrentValidatorSetHash is a free data retrieval call binding the contract method 0xf896f1a5.
//
// Solidity: function currentValidatorSetHash() view returns(bytes32)
func (_Contractapi *ContractapiCallerSession) CurrentValidatorSetHash() ([32]byte, error) {
	return _Contractapi.Contract.CurrentValidatorSetHash(&_Contractapi.CallOpts)
}

// CurrentValidatorSetLength is a free data retrieval call binding the contract method 0x1d1d4f26.
//
// Solidity: function currentValidatorSetLength() view returns(uint256)
func (_Contractapi *ContractapiCaller) CurrentValidatorSetLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractapi.contract.Call(opts, &out, "currentValidatorSetLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValidatorSetLength is a free data retrieval call binding the contract method 0x1d1d4f26.
//
// Solidity: function currentValidatorSetLength() view returns(uint256)
func (_Contractapi *ContractapiSession) CurrentValidatorSetLength() (*big.Int, error) {
	return _Contractapi.Contract.CurrentValidatorSetLength(&_Contractapi.CallOpts)
}

// CurrentValidatorSetLength is a free data retrieval call binding the contract method 0x1d1d4f26.
//
// Solidity: function currentValidatorSetLength() view returns(uint256)
func (_Contractapi *ContractapiCallerSession) CurrentValidatorSetLength() (*big.Int, error) {
	return _Contractapi.Contract.CurrentValidatorSetLength(&_Contractapi.CallOpts)
}

// GetCheckpointBlock is a free data retrieval call binding the contract method 0x22fd1818.
//
// Solidity: function getCheckpointBlock(uint256 blockNumber) view returns(bool, uint256)
func (_Contractapi *ContractapiCaller) GetCheckpointBlock(opts *bind.CallOpts, blockNumber *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _Contractapi.contract.Call(opts, &out, "getCheckpointBlock", blockNumber)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetCheckpointBlock is a free data retrieval call binding the contract method 0x22fd1818.
//
// Solidity: function getCheckpointBlock(uint256 blockNumber) view returns(bool, uint256)
func (_Contractapi *ContractapiSession) GetCheckpointBlock(blockNumber *big.Int) (bool, *big.Int, error) {
	return _Contractapi.Contract.GetCheckpointBlock(&_Contractapi.CallOpts, blockNumber)
}

// GetCheckpointBlock is a free data retrieval call binding the contract method 0x22fd1818.
//
// Solidity: function getCheckpointBlock(uint256 blockNumber) view returns(bool, uint256)
func (_Contractapi *ContractapiCallerSession) GetCheckpointBlock(blockNumber *big.Int) (bool, *big.Int, error) {
	return _Contractapi.Contract.GetCheckpointBlock(&_Contractapi.CallOpts, blockNumber)
}

// GetEventMembershipByBlockNumber is a free data retrieval call binding the contract method 0x61a02208.
//
// Solidity: function getEventMembershipByBlockNumber(uint256 blockNumber, bytes32 leaf, uint256 leafIndex, bytes32[] proof) view returns(bool)
func (_Contractapi *ContractapiCaller) GetEventMembershipByBlockNumber(opts *bind.CallOpts, blockNumber *big.Int, leaf [32]byte, leafIndex *big.Int, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _Contractapi.contract.Call(opts, &out, "getEventMembershipByBlockNumber", blockNumber, leaf, leafIndex, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetEventMembershipByBlockNumber is a free data retrieval call binding the contract method 0x61a02208.
//
// Solidity: function getEventMembershipByBlockNumber(uint256 blockNumber, bytes32 leaf, uint256 leafIndex, bytes32[] proof) view returns(bool)
func (_Contractapi *ContractapiSession) GetEventMembershipByBlockNumber(blockNumber *big.Int, leaf [32]byte, leafIndex *big.Int, proof [][32]byte) (bool, error) {
	return _Contractapi.Contract.GetEventMembershipByBlockNumber(&_Contractapi.CallOpts, blockNumber, leaf, leafIndex, proof)
}

// GetEventMembershipByBlockNumber is a free data retrieval call binding the contract method 0x61a02208.
//
// Solidity: function getEventMembershipByBlockNumber(uint256 blockNumber, bytes32 leaf, uint256 leafIndex, bytes32[] proof) view returns(bool)
func (_Contractapi *ContractapiCallerSession) GetEventMembershipByBlockNumber(blockNumber *big.Int, leaf [32]byte, leafIndex *big.Int, proof [][32]byte) (bool, error) {
	return _Contractapi.Contract.GetEventMembershipByBlockNumber(&_Contractapi.CallOpts, blockNumber, leaf, leafIndex, proof)
}

// GetEventMembershipByEpoch is a free data retrieval call binding the contract method 0x729e7c6e.
//
// Solidity: function getEventMembershipByEpoch(uint256 epoch, bytes32 leaf, uint256 leafIndex, bytes32[] proof) view returns(bool)
func (_Contractapi *ContractapiCaller) GetEventMembershipByEpoch(opts *bind.CallOpts, epoch *big.Int, leaf [32]byte, leafIndex *big.Int, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _Contractapi.contract.Call(opts, &out, "getEventMembershipByEpoch", epoch, leaf, leafIndex, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetEventMembershipByEpoch is a free data retrieval call binding the contract method 0x729e7c6e.
//
// Solidity: function getEventMembershipByEpoch(uint256 epoch, bytes32 leaf, uint256 leafIndex, bytes32[] proof) view returns(bool)
func (_Contractapi *ContractapiSession) GetEventMembershipByEpoch(epoch *big.Int, leaf [32]byte, leafIndex *big.Int, proof [][32]byte) (bool, error) {
	return _Contractapi.Contract.GetEventMembershipByEpoch(&_Contractapi.CallOpts, epoch, leaf, leafIndex, proof)
}

// GetEventMembershipByEpoch is a free data retrieval call binding the contract method 0x729e7c6e.
//
// Solidity: function getEventMembershipByEpoch(uint256 epoch, bytes32 leaf, uint256 leafIndex, bytes32[] proof) view returns(bool)
func (_Contractapi *ContractapiCallerSession) GetEventMembershipByEpoch(epoch *big.Int, leaf [32]byte, leafIndex *big.Int, proof [][32]byte) (bool, error) {
	return _Contractapi.Contract.GetEventMembershipByEpoch(&_Contractapi.CallOpts, epoch, leaf, leafIndex, proof)
}

// GetEventRootByBlock is a free data retrieval call binding the contract method 0x3569ed93.
//
// Solidity: function getEventRootByBlock(uint256 blockNumber) view returns(bytes32)
func (_Contractapi *ContractapiCaller) GetEventRootByBlock(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Contractapi.contract.Call(opts, &out, "getEventRootByBlock", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEventRootByBlock is a free data retrieval call binding the contract method 0x3569ed93.
//
// Solidity: function getEventRootByBlock(uint256 blockNumber) view returns(bytes32)
func (_Contractapi *ContractapiSession) GetEventRootByBlock(blockNumber *big.Int) ([32]byte, error) {
	return _Contractapi.Contract.GetEventRootByBlock(&_Contractapi.CallOpts, blockNumber)
}

// GetEventRootByBlock is a free data retrieval call binding the contract method 0x3569ed93.
//
// Solidity: function getEventRootByBlock(uint256 blockNumber) view returns(bytes32)
func (_Contractapi *ContractapiCallerSession) GetEventRootByBlock(blockNumber *big.Int) ([32]byte, error) {
	return _Contractapi.Contract.GetEventRootByBlock(&_Contractapi.CallOpts, blockNumber)
}

// Initialize is a paid mutator transaction binding the contract method 0x24d4fad7.
//
// Solidity: function initialize(address newBls, uint256 chainId_, (address,uint256[2])[] newValidatorSet) returns()
func (_Contractapi *ContractapiTransactor) Initialize(opts *bind.TransactOpts, newBls common.Address, chainId_ *big.Int, newValidatorSet []ICheckpointManagerValidator) (*types.Transaction, error) {
	return _Contractapi.contract.Transact(opts, "initialize", newBls, chainId_, newValidatorSet)
}

// Initialize is a paid mutator transaction binding the contract method 0x24d4fad7.
//
// Solidity: function initialize(address newBls, uint256 chainId_, (address,uint256[2])[] newValidatorSet) returns()
func (_Contractapi *ContractapiSession) Initialize(newBls common.Address, chainId_ *big.Int, newValidatorSet []ICheckpointManagerValidator) (*types.Transaction, error) {
	return _Contractapi.Contract.Initialize(&_Contractapi.TransactOpts, newBls, chainId_, newValidatorSet)
}

// Initialize is a paid mutator transaction binding the contract method 0x24d4fad7.
//
// Solidity: function initialize(address newBls, uint256 chainId_, (address,uint256[2])[] newValidatorSet) returns()
func (_Contractapi *ContractapiTransactorSession) Initialize(newBls common.Address, chainId_ *big.Int, newValidatorSet []ICheckpointManagerValidator) (*types.Transaction, error) {
	return _Contractapi.Contract.Initialize(&_Contractapi.TransactOpts, newBls, chainId_, newValidatorSet)
}

// Submit is a paid mutator transaction binding the contract method 0x4bde58a1.
//
// Solidity: function submit((bytes32,uint32,bytes32) checkpointMetadata, (uint64,uint64,uint64,bytes32,bytes32) checkpoint, bytes signature, bytes bitmap, (address,uint256[2])[] newValidatorSet, uint256 leafIndex, bytes32[] proof) returns()
func (_Contractapi *ContractapiTransactor) Submit(opts *bind.TransactOpts, checkpointMetadata ICheckpointManagerCheckpointMetadata, checkpoint ICheckpointManagerCheckpoint, signature []byte, bitmap []byte, newValidatorSet []ICheckpointManagerValidator, leafIndex *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Contractapi.contract.Transact(opts, "submit", checkpointMetadata, checkpoint, signature, bitmap, newValidatorSet, leafIndex, proof)
}

// Submit is a paid mutator transaction binding the contract method 0x4bde58a1.
//
// Solidity: function submit((bytes32,uint32,bytes32) checkpointMetadata, (uint64,uint64,uint64,bytes32,bytes32) checkpoint, bytes signature, bytes bitmap, (address,uint256[2])[] newValidatorSet, uint256 leafIndex, bytes32[] proof) returns()
func (_Contractapi *ContractapiSession) Submit(checkpointMetadata ICheckpointManagerCheckpointMetadata, checkpoint ICheckpointManagerCheckpoint, signature []byte, bitmap []byte, newValidatorSet []ICheckpointManagerValidator, leafIndex *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Contractapi.Contract.Submit(&_Contractapi.TransactOpts, checkpointMetadata, checkpoint, signature, bitmap, newValidatorSet, leafIndex, proof)
}

// Submit is a paid mutator transaction binding the contract method 0x4bde58a1.
//
// Solidity: function submit((bytes32,uint32,bytes32) checkpointMetadata, (uint64,uint64,uint64,bytes32,bytes32) checkpoint, bytes signature, bytes bitmap, (address,uint256[2])[] newValidatorSet, uint256 leafIndex, bytes32[] proof) returns()
func (_Contractapi *ContractapiTransactorSession) Submit(checkpointMetadata ICheckpointManagerCheckpointMetadata, checkpoint ICheckpointManagerCheckpoint, signature []byte, bitmap []byte, newValidatorSet []ICheckpointManagerValidator, leafIndex *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Contractapi.Contract.Submit(&_Contractapi.TransactOpts, checkpointMetadata, checkpoint, signature, bitmap, newValidatorSet, leafIndex, proof)
}

// ContractapiInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contractapi contract.
type ContractapiInitializedIterator struct {
	Event *ContractapiInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  platon.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractapiInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractapiInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractapiInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractapiInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractapiInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractapiInitialized represents a Initialized event raised by the Contractapi contract.
type ContractapiInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contractapi *ContractapiFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractapiInitializedIterator, error) {

	logs, sub, err := _Contractapi.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractapiInitializedIterator{contract: _Contractapi.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contractapi *ContractapiFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractapiInitialized) (event.Subscription, error) {

	logs, sub, err := _Contractapi.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractapiInitialized)
				if err := _Contractapi.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contractapi *ContractapiFilterer) ParseInitialized(log types.Log) (*ContractapiInitialized, error) {
	event := new(ContractapiInitialized)
	if err := _Contractapi.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
