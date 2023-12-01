// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package checkpoint_manager

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
	BlsKey  []byte
}

// CheckpointManagerMetaData contains all meta data concerning the CheckpointManager contract.
var CheckpointManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"INITIALIZER\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bls\",\"outputs\":[{\"internalType\":\"contractIBLS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"checkpointBlockNumbers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"checkpoints\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"viewNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"eventRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extendRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentCheckpointBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentValidatorSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentValidatorSetHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentValidatorSetLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getCheckpointBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"getEventMembershipByBlockNumber\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"getEventMembershipByEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEventRootByBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBLS\",\"name\":\"newBls\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structICheckpointManager.Validator[]\",\"name\":\"newValidatorSet\",\"type\":\"tuple[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blockIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"currentValidatorSetHash\",\"type\":\"bytes32\"}],\"internalType\":\"structICheckpointManager.CheckpointMetadata\",\"name\":\"checkpointMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"viewNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"eventRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extendRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structICheckpointManager.Checkpoint\",\"name\":\"checkpoint\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bitmap\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structICheckpointManager.Validator[]\",\"name\":\"newValidatorSet\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CheckpointManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use CheckpointManagerMetaData.ABI instead.
var CheckpointManagerABI = CheckpointManagerMetaData.ABI

// CheckpointManager is an auto generated Go binding around an platon contract.
type CheckpointManager struct {
	CheckpointManagerCaller     // Read-only binding to the contract
	CheckpointManagerTransactor // Write-only binding to the contract
	CheckpointManagerFilterer   // Log filterer for contract events
}

// CheckpointManagerCaller is an auto generated read-only Go binding around an platon contract.
type CheckpointManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointManagerTransactor is an auto generated write-only Go binding around an platon contract.
type CheckpointManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointManagerFilterer is an auto generated log filtering Go binding around an platon contract events.
type CheckpointManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointManagerSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type CheckpointManagerSession struct {
	Contract     *CheckpointManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CheckpointManagerCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type CheckpointManagerCallerSession struct {
	Contract *CheckpointManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CheckpointManagerTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type CheckpointManagerTransactorSession struct {
	Contract     *CheckpointManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CheckpointManagerRaw is an auto generated low-level Go binding around an platon contract.
type CheckpointManagerRaw struct {
	Contract *CheckpointManager // Generic contract binding to access the raw methods on
}

// CheckpointManagerCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type CheckpointManagerCallerRaw struct {
	Contract *CheckpointManagerCaller // Generic read-only contract binding to access the raw methods on
}

// CheckpointManagerTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type CheckpointManagerTransactorRaw struct {
	Contract *CheckpointManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCheckpointManager creates a new instance of CheckpointManager, bound to a specific deployed contract.
func NewCheckpointManager(address common.Address, backend bind.ContractBackend) (*CheckpointManager, error) {
	contract, err := bindCheckpointManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CheckpointManager{CheckpointManagerCaller: CheckpointManagerCaller{contract: contract}, CheckpointManagerTransactor: CheckpointManagerTransactor{contract: contract}, CheckpointManagerFilterer: CheckpointManagerFilterer{contract: contract}}, nil
}

// NewCheckpointManagerCaller creates a new read-only instance of CheckpointManager, bound to a specific deployed contract.
func NewCheckpointManagerCaller(address common.Address, caller bind.ContractCaller) (*CheckpointManagerCaller, error) {
	contract, err := bindCheckpointManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CheckpointManagerCaller{contract: contract}, nil
}

// NewCheckpointManagerTransactor creates a new write-only instance of CheckpointManager, bound to a specific deployed contract.
func NewCheckpointManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*CheckpointManagerTransactor, error) {
	contract, err := bindCheckpointManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CheckpointManagerTransactor{contract: contract}, nil
}

// NewCheckpointManagerFilterer creates a new log filterer instance of CheckpointManager, bound to a specific deployed contract.
func NewCheckpointManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*CheckpointManagerFilterer, error) {
	contract, err := bindCheckpointManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CheckpointManagerFilterer{contract: contract}, nil
}

// bindCheckpointManager binds a generic wrapper to an already deployed contract.
func bindCheckpointManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CheckpointManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CheckpointManager *CheckpointManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CheckpointManager.Contract.CheckpointManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CheckpointManager *CheckpointManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckpointManager.Contract.CheckpointManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CheckpointManager *CheckpointManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CheckpointManager.Contract.CheckpointManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CheckpointManager *CheckpointManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CheckpointManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CheckpointManager *CheckpointManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckpointManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CheckpointManager *CheckpointManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CheckpointManager.Contract.contract.Transact(opts, method, params...)
}

// Bls is a free data retrieval call binding the contract method 0x95b0b027.
//
// Solidity: function bls() view returns(address)
func (_CheckpointManager *CheckpointManagerCaller) Bls(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "bls")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bls is a free data retrieval call binding the contract method 0x95b0b027.
//
// Solidity: function bls() view returns(address)
func (_CheckpointManager *CheckpointManagerSession) Bls() (common.Address, error) {
	return _CheckpointManager.Contract.Bls(&_CheckpointManager.CallOpts)
}

// Bls is a free data retrieval call binding the contract method 0x95b0b027.
//
// Solidity: function bls() view returns(address)
func (_CheckpointManager *CheckpointManagerCallerSession) Bls() (common.Address, error) {
	return _CheckpointManager.Contract.Bls(&_CheckpointManager.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_CheckpointManager *CheckpointManagerCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_CheckpointManager *CheckpointManagerSession) ChainId() (*big.Int, error) {
	return _CheckpointManager.Contract.ChainId(&_CheckpointManager.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_CheckpointManager *CheckpointManagerCallerSession) ChainId() (*big.Int, error) {
	return _CheckpointManager.Contract.ChainId(&_CheckpointManager.CallOpts)
}

// CheckpointBlockNumbers is a free data retrieval call binding the contract method 0xe9193d2b.
//
// Solidity: function checkpointBlockNumbers(uint256 ) view returns(uint256)
func (_CheckpointManager *CheckpointManagerCaller) CheckpointBlockNumbers(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "checkpointBlockNumbers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckpointBlockNumbers is a free data retrieval call binding the contract method 0xe9193d2b.
//
// Solidity: function checkpointBlockNumbers(uint256 ) view returns(uint256)
func (_CheckpointManager *CheckpointManagerSession) CheckpointBlockNumbers(arg0 *big.Int) (*big.Int, error) {
	return _CheckpointManager.Contract.CheckpointBlockNumbers(&_CheckpointManager.CallOpts, arg0)
}

// CheckpointBlockNumbers is a free data retrieval call binding the contract method 0xe9193d2b.
//
// Solidity: function checkpointBlockNumbers(uint256 ) view returns(uint256)
func (_CheckpointManager *CheckpointManagerCallerSession) CheckpointBlockNumbers(arg0 *big.Int) (*big.Int, error) {
	return _CheckpointManager.Contract.CheckpointBlockNumbers(&_CheckpointManager.CallOpts, arg0)
}

// Checkpoints is a free data retrieval call binding the contract method 0xb8a24252.
//
// Solidity: function checkpoints(uint256 ) view returns(uint64 epoch, uint64 viewNumber, uint64 blockNumber, bytes32 eventRoot, bytes32 extendRoot)
func (_CheckpointManager *CheckpointManagerCaller) Checkpoints(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Epoch       uint64
	ViewNumber  uint64
	BlockNumber uint64
	EventRoot   [32]byte
	ExtendRoot  [32]byte
}, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "checkpoints", arg0)

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
func (_CheckpointManager *CheckpointManagerSession) Checkpoints(arg0 *big.Int) (struct {
	Epoch       uint64
	ViewNumber  uint64
	BlockNumber uint64
	EventRoot   [32]byte
	ExtendRoot  [32]byte
}, error) {
	return _CheckpointManager.Contract.Checkpoints(&_CheckpointManager.CallOpts, arg0)
}

// Checkpoints is a free data retrieval call binding the contract method 0xb8a24252.
//
// Solidity: function checkpoints(uint256 ) view returns(uint64 epoch, uint64 viewNumber, uint64 blockNumber, bytes32 eventRoot, bytes32 extendRoot)
func (_CheckpointManager *CheckpointManagerCallerSession) Checkpoints(arg0 *big.Int) (struct {
	Epoch       uint64
	ViewNumber  uint64
	BlockNumber uint64
	EventRoot   [32]byte
	ExtendRoot  [32]byte
}, error) {
	return _CheckpointManager.Contract.Checkpoints(&_CheckpointManager.CallOpts, arg0)
}

// CurrentCheckpointBlockNumber is a free data retrieval call binding the contract method 0xe416d677.
//
// Solidity: function currentCheckpointBlockNumber() view returns(uint256)
func (_CheckpointManager *CheckpointManagerCaller) CurrentCheckpointBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "currentCheckpointBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentCheckpointBlockNumber is a free data retrieval call binding the contract method 0xe416d677.
//
// Solidity: function currentCheckpointBlockNumber() view returns(uint256)
func (_CheckpointManager *CheckpointManagerSession) CurrentCheckpointBlockNumber() (*big.Int, error) {
	return _CheckpointManager.Contract.CurrentCheckpointBlockNumber(&_CheckpointManager.CallOpts)
}

// CurrentCheckpointBlockNumber is a free data retrieval call binding the contract method 0xe416d677.
//
// Solidity: function currentCheckpointBlockNumber() view returns(uint256)
func (_CheckpointManager *CheckpointManagerCallerSession) CurrentCheckpointBlockNumber() (*big.Int, error) {
	return _CheckpointManager.Contract.CurrentCheckpointBlockNumber(&_CheckpointManager.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_CheckpointManager *CheckpointManagerCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_CheckpointManager *CheckpointManagerSession) CurrentEpoch() (*big.Int, error) {
	return _CheckpointManager.Contract.CurrentEpoch(&_CheckpointManager.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_CheckpointManager *CheckpointManagerCallerSession) CurrentEpoch() (*big.Int, error) {
	return _CheckpointManager.Contract.CurrentEpoch(&_CheckpointManager.CallOpts)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address _address, bytes blsKey)
func (_CheckpointManager *CheckpointManagerCaller) CurrentValidatorSet(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Address common.Address
	BlsKey  []byte
}, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "currentValidatorSet", arg0)

	outstruct := new(struct {
		Address common.Address
		BlsKey  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Address = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BlsKey = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address _address, bytes blsKey)
func (_CheckpointManager *CheckpointManagerSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	Address common.Address
	BlsKey  []byte
}, error) {
	return _CheckpointManager.Contract.CurrentValidatorSet(&_CheckpointManager.CallOpts, arg0)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address _address, bytes blsKey)
func (_CheckpointManager *CheckpointManagerCallerSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	Address common.Address
	BlsKey  []byte
}, error) {
	return _CheckpointManager.Contract.CurrentValidatorSet(&_CheckpointManager.CallOpts, arg0)
}

// CurrentValidatorSetHash is a free data retrieval call binding the contract method 0xf896f1a5.
//
// Solidity: function currentValidatorSetHash() view returns(bytes32)
func (_CheckpointManager *CheckpointManagerCaller) CurrentValidatorSetHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "currentValidatorSetHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentValidatorSetHash is a free data retrieval call binding the contract method 0xf896f1a5.
//
// Solidity: function currentValidatorSetHash() view returns(bytes32)
func (_CheckpointManager *CheckpointManagerSession) CurrentValidatorSetHash() ([32]byte, error) {
	return _CheckpointManager.Contract.CurrentValidatorSetHash(&_CheckpointManager.CallOpts)
}

// CurrentValidatorSetHash is a free data retrieval call binding the contract method 0xf896f1a5.
//
// Solidity: function currentValidatorSetHash() view returns(bytes32)
func (_CheckpointManager *CheckpointManagerCallerSession) CurrentValidatorSetHash() ([32]byte, error) {
	return _CheckpointManager.Contract.CurrentValidatorSetHash(&_CheckpointManager.CallOpts)
}

// CurrentValidatorSetLength is a free data retrieval call binding the contract method 0x1d1d4f26.
//
// Solidity: function currentValidatorSetLength() view returns(uint256)
func (_CheckpointManager *CheckpointManagerCaller) CurrentValidatorSetLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "currentValidatorSetLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValidatorSetLength is a free data retrieval call binding the contract method 0x1d1d4f26.
//
// Solidity: function currentValidatorSetLength() view returns(uint256)
func (_CheckpointManager *CheckpointManagerSession) CurrentValidatorSetLength() (*big.Int, error) {
	return _CheckpointManager.Contract.CurrentValidatorSetLength(&_CheckpointManager.CallOpts)
}

// CurrentValidatorSetLength is a free data retrieval call binding the contract method 0x1d1d4f26.
//
// Solidity: function currentValidatorSetLength() view returns(uint256)
func (_CheckpointManager *CheckpointManagerCallerSession) CurrentValidatorSetLength() (*big.Int, error) {
	return _CheckpointManager.Contract.CurrentValidatorSetLength(&_CheckpointManager.CallOpts)
}

// GetCheckpointBlock is a free data retrieval call binding the contract method 0x22fd1818.
//
// Solidity: function getCheckpointBlock(uint256 blockNumber) view returns(bool, uint256)
func (_CheckpointManager *CheckpointManagerCaller) GetCheckpointBlock(opts *bind.CallOpts, blockNumber *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "getCheckpointBlock", blockNumber)

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
func (_CheckpointManager *CheckpointManagerSession) GetCheckpointBlock(blockNumber *big.Int) (bool, *big.Int, error) {
	return _CheckpointManager.Contract.GetCheckpointBlock(&_CheckpointManager.CallOpts, blockNumber)
}

// GetCheckpointBlock is a free data retrieval call binding the contract method 0x22fd1818.
//
// Solidity: function getCheckpointBlock(uint256 blockNumber) view returns(bool, uint256)
func (_CheckpointManager *CheckpointManagerCallerSession) GetCheckpointBlock(blockNumber *big.Int) (bool, *big.Int, error) {
	return _CheckpointManager.Contract.GetCheckpointBlock(&_CheckpointManager.CallOpts, blockNumber)
}

// GetEventMembershipByBlockNumber is a free data retrieval call binding the contract method 0x61a02208.
//
// Solidity: function getEventMembershipByBlockNumber(uint256 blockNumber, bytes32 leaf, uint256 leafIndex, bytes32[] proof) view returns(bool)
func (_CheckpointManager *CheckpointManagerCaller) GetEventMembershipByBlockNumber(opts *bind.CallOpts, blockNumber *big.Int, leaf [32]byte, leafIndex *big.Int, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "getEventMembershipByBlockNumber", blockNumber, leaf, leafIndex, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetEventMembershipByBlockNumber is a free data retrieval call binding the contract method 0x61a02208.
//
// Solidity: function getEventMembershipByBlockNumber(uint256 blockNumber, bytes32 leaf, uint256 leafIndex, bytes32[] proof) view returns(bool)
func (_CheckpointManager *CheckpointManagerSession) GetEventMembershipByBlockNumber(blockNumber *big.Int, leaf [32]byte, leafIndex *big.Int, proof [][32]byte) (bool, error) {
	return _CheckpointManager.Contract.GetEventMembershipByBlockNumber(&_CheckpointManager.CallOpts, blockNumber, leaf, leafIndex, proof)
}

// GetEventMembershipByBlockNumber is a free data retrieval call binding the contract method 0x61a02208.
//
// Solidity: function getEventMembershipByBlockNumber(uint256 blockNumber, bytes32 leaf, uint256 leafIndex, bytes32[] proof) view returns(bool)
func (_CheckpointManager *CheckpointManagerCallerSession) GetEventMembershipByBlockNumber(blockNumber *big.Int, leaf [32]byte, leafIndex *big.Int, proof [][32]byte) (bool, error) {
	return _CheckpointManager.Contract.GetEventMembershipByBlockNumber(&_CheckpointManager.CallOpts, blockNumber, leaf, leafIndex, proof)
}

// GetEventMembershipByEpoch is a free data retrieval call binding the contract method 0x729e7c6e.
//
// Solidity: function getEventMembershipByEpoch(uint256 epoch, bytes32 leaf, uint256 leafIndex, bytes32[] proof) view returns(bool)
func (_CheckpointManager *CheckpointManagerCaller) GetEventMembershipByEpoch(opts *bind.CallOpts, epoch *big.Int, leaf [32]byte, leafIndex *big.Int, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "getEventMembershipByEpoch", epoch, leaf, leafIndex, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetEventMembershipByEpoch is a free data retrieval call binding the contract method 0x729e7c6e.
//
// Solidity: function getEventMembershipByEpoch(uint256 epoch, bytes32 leaf, uint256 leafIndex, bytes32[] proof) view returns(bool)
func (_CheckpointManager *CheckpointManagerSession) GetEventMembershipByEpoch(epoch *big.Int, leaf [32]byte, leafIndex *big.Int, proof [][32]byte) (bool, error) {
	return _CheckpointManager.Contract.GetEventMembershipByEpoch(&_CheckpointManager.CallOpts, epoch, leaf, leafIndex, proof)
}

// GetEventMembershipByEpoch is a free data retrieval call binding the contract method 0x729e7c6e.
//
// Solidity: function getEventMembershipByEpoch(uint256 epoch, bytes32 leaf, uint256 leafIndex, bytes32[] proof) view returns(bool)
func (_CheckpointManager *CheckpointManagerCallerSession) GetEventMembershipByEpoch(epoch *big.Int, leaf [32]byte, leafIndex *big.Int, proof [][32]byte) (bool, error) {
	return _CheckpointManager.Contract.GetEventMembershipByEpoch(&_CheckpointManager.CallOpts, epoch, leaf, leafIndex, proof)
}

// GetEventRootByBlock is a free data retrieval call binding the contract method 0x3569ed93.
//
// Solidity: function getEventRootByBlock(uint256 blockNumber) view returns(bytes32)
func (_CheckpointManager *CheckpointManagerCaller) GetEventRootByBlock(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "getEventRootByBlock", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEventRootByBlock is a free data retrieval call binding the contract method 0x3569ed93.
//
// Solidity: function getEventRootByBlock(uint256 blockNumber) view returns(bytes32)
func (_CheckpointManager *CheckpointManagerSession) GetEventRootByBlock(blockNumber *big.Int) ([32]byte, error) {
	return _CheckpointManager.Contract.GetEventRootByBlock(&_CheckpointManager.CallOpts, blockNumber)
}

// GetEventRootByBlock is a free data retrieval call binding the contract method 0x3569ed93.
//
// Solidity: function getEventRootByBlock(uint256 blockNumber) view returns(bytes32)
func (_CheckpointManager *CheckpointManagerCallerSession) GetEventRootByBlock(blockNumber *big.Int) ([32]byte, error) {
	return _CheckpointManager.Contract.GetEventRootByBlock(&_CheckpointManager.CallOpts, blockNumber)
}

// Initialize is a paid mutator transaction binding the contract method 0x9b2a80a1.
//
// Solidity: function initialize(address newBls, uint256 chainId_, (address,bytes)[] newValidatorSet) returns()
func (_CheckpointManager *CheckpointManagerTransactor) Initialize(opts *bind.TransactOpts, newBls common.Address, chainId_ *big.Int, newValidatorSet []ICheckpointManagerValidator) (*types.Transaction, error) {
	return _CheckpointManager.contract.Transact(opts, "initialize", newBls, chainId_, newValidatorSet)
}

// Initialize is a paid mutator transaction binding the contract method 0x9b2a80a1.
//
// Solidity: function initialize(address newBls, uint256 chainId_, (address,bytes)[] newValidatorSet) returns()
func (_CheckpointManager *CheckpointManagerSession) Initialize(newBls common.Address, chainId_ *big.Int, newValidatorSet []ICheckpointManagerValidator) (*types.Transaction, error) {
	return _CheckpointManager.Contract.Initialize(&_CheckpointManager.TransactOpts, newBls, chainId_, newValidatorSet)
}

// Initialize is a paid mutator transaction binding the contract method 0x9b2a80a1.
//
// Solidity: function initialize(address newBls, uint256 chainId_, (address,bytes)[] newValidatorSet) returns()
func (_CheckpointManager *CheckpointManagerTransactorSession) Initialize(newBls common.Address, chainId_ *big.Int, newValidatorSet []ICheckpointManagerValidator) (*types.Transaction, error) {
	return _CheckpointManager.Contract.Initialize(&_CheckpointManager.TransactOpts, newBls, chainId_, newValidatorSet)
}

// Submit is a paid mutator transaction binding the contract method 0x836b9e7f.
//
// Solidity: function submit((bytes32,uint32,bytes32) checkpointMetadata, (uint64,uint64,uint64,bytes32,bytes32) checkpoint, bytes signature, bytes bitmap, (address,bytes)[] newValidatorSet, uint256 leafIndex, bytes32[] proof) returns()
func (_CheckpointManager *CheckpointManagerTransactor) Submit(opts *bind.TransactOpts, checkpointMetadata ICheckpointManagerCheckpointMetadata, checkpoint ICheckpointManagerCheckpoint, signature []byte, bitmap []byte, newValidatorSet []ICheckpointManagerValidator, leafIndex *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _CheckpointManager.contract.Transact(opts, "submit", checkpointMetadata, checkpoint, signature, bitmap, newValidatorSet, leafIndex, proof)
}

// Submit is a paid mutator transaction binding the contract method 0x836b9e7f.
//
// Solidity: function submit((bytes32,uint32,bytes32) checkpointMetadata, (uint64,uint64,uint64,bytes32,bytes32) checkpoint, bytes signature, bytes bitmap, (address,bytes)[] newValidatorSet, uint256 leafIndex, bytes32[] proof) returns()
func (_CheckpointManager *CheckpointManagerSession) Submit(checkpointMetadata ICheckpointManagerCheckpointMetadata, checkpoint ICheckpointManagerCheckpoint, signature []byte, bitmap []byte, newValidatorSet []ICheckpointManagerValidator, leafIndex *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _CheckpointManager.Contract.Submit(&_CheckpointManager.TransactOpts, checkpointMetadata, checkpoint, signature, bitmap, newValidatorSet, leafIndex, proof)
}

// Submit is a paid mutator transaction binding the contract method 0x836b9e7f.
//
// Solidity: function submit((bytes32,uint32,bytes32) checkpointMetadata, (uint64,uint64,uint64,bytes32,bytes32) checkpoint, bytes signature, bytes bitmap, (address,bytes)[] newValidatorSet, uint256 leafIndex, bytes32[] proof) returns()
func (_CheckpointManager *CheckpointManagerTransactorSession) Submit(checkpointMetadata ICheckpointManagerCheckpointMetadata, checkpoint ICheckpointManagerCheckpoint, signature []byte, bitmap []byte, newValidatorSet []ICheckpointManagerValidator, leafIndex *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _CheckpointManager.Contract.Submit(&_CheckpointManager.TransactOpts, checkpointMetadata, checkpoint, signature, bitmap, newValidatorSet, leafIndex, proof)
}

// CheckpointManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CheckpointManager contract.
type CheckpointManagerInitializedIterator struct {
	Event *CheckpointManagerInitialized // Event containing the contract specifics and raw log

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
func (it *CheckpointManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CheckpointManagerInitialized)
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
		it.Event = new(CheckpointManagerInitialized)
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
func (it *CheckpointManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CheckpointManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CheckpointManagerInitialized represents a Initialized event raised by the CheckpointManager contract.
type CheckpointManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CheckpointManager *CheckpointManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*CheckpointManagerInitializedIterator, error) {

	logs, sub, err := _CheckpointManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CheckpointManagerInitializedIterator{contract: _CheckpointManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CheckpointManager *CheckpointManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CheckpointManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _CheckpointManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CheckpointManagerInitialized)
				if err := _CheckpointManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CheckpointManager *CheckpointManagerFilterer) ParseInitialized(log types.Log) (*CheckpointManagerInitialized, error) {
	event := new(CheckpointManagerInitialized)
	if err := _CheckpointManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
