// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2testcontract

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

// L2testcontractMetaData contains all meta data concerning the L2testcontract contract.
var L2testcontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dec\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// L2testcontractABI is the input ABI used to generate the binding from.
// Deprecated: Use L2testcontractMetaData.ABI instead.
var L2testcontractABI = L2testcontractMetaData.ABI

// L2testcontract is an auto generated Go binding around an platon contract.
type L2testcontract struct {
	L2testcontractCaller     // Read-only binding to the contract
	L2testcontractTransactor // Write-only binding to the contract
	L2testcontractFilterer   // Log filterer for contract events
}

// L2testcontractCaller is an auto generated read-only Go binding around an platon contract.
type L2testcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2testcontractTransactor is an auto generated write-only Go binding around an platon contract.
type L2testcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2testcontractFilterer is an auto generated log filtering Go binding around an platon contract events.
type L2testcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2testcontractSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type L2testcontractSession struct {
	Contract     *L2testcontract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2testcontractCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type L2testcontractCallerSession struct {
	Contract *L2testcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// L2testcontractTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type L2testcontractTransactorSession struct {
	Contract     *L2testcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// L2testcontractRaw is an auto generated low-level Go binding around an platon contract.
type L2testcontractRaw struct {
	Contract *L2testcontract // Generic contract binding to access the raw methods on
}

// L2testcontractCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type L2testcontractCallerRaw struct {
	Contract *L2testcontractCaller // Generic read-only contract binding to access the raw methods on
}

// L2testcontractTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type L2testcontractTransactorRaw struct {
	Contract *L2testcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2testcontract creates a new instance of L2testcontract, bound to a specific deployed contract.
func NewL2testcontract(address common.Address, backend bind.ContractBackend) (*L2testcontract, error) {
	contract, err := bindL2testcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2testcontract{L2testcontractCaller: L2testcontractCaller{contract: contract}, L2testcontractTransactor: L2testcontractTransactor{contract: contract}, L2testcontractFilterer: L2testcontractFilterer{contract: contract}}, nil
}

// NewL2testcontractCaller creates a new read-only instance of L2testcontract, bound to a specific deployed contract.
func NewL2testcontractCaller(address common.Address, caller bind.ContractCaller) (*L2testcontractCaller, error) {
	contract, err := bindL2testcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2testcontractCaller{contract: contract}, nil
}

// NewL2testcontractTransactor creates a new write-only instance of L2testcontract, bound to a specific deployed contract.
func NewL2testcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*L2testcontractTransactor, error) {
	contract, err := bindL2testcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2testcontractTransactor{contract: contract}, nil
}

// NewL2testcontractFilterer creates a new log filterer instance of L2testcontract, bound to a specific deployed contract.
func NewL2testcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*L2testcontractFilterer, error) {
	contract, err := bindL2testcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2testcontractFilterer{contract: contract}, nil
}

// bindL2testcontract binds a generic wrapper to an already deployed contract.
func bindL2testcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2testcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2testcontract *L2testcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2testcontract.Contract.L2testcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2testcontract *L2testcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2testcontract.Contract.L2testcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2testcontract *L2testcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2testcontract.Contract.L2testcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2testcontract *L2testcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2testcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2testcontract *L2testcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2testcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2testcontract *L2testcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2testcontract.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint64)
func (_L2testcontract *L2testcontractCaller) Count(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _L2testcontract.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint64)
func (_L2testcontract *L2testcontractSession) Count() (uint64, error) {
	return _L2testcontract.Contract.Count(&_L2testcontract.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint64)
func (_L2testcontract *L2testcontractCallerSession) Count() (uint64, error) {
	return _L2testcontract.Contract.Count(&_L2testcontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L2testcontract *L2testcontractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2testcontract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L2testcontract *L2testcontractSession) Name() (string, error) {
	return _L2testcontract.Contract.Name(&_L2testcontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L2testcontract *L2testcontractCallerSession) Name() (string, error) {
	return _L2testcontract.Contract.Name(&_L2testcontract.CallOpts)
}

// Dec is a paid mutator transaction binding the contract method 0xb3bcfa82.
//
// Solidity: function dec() returns()
func (_L2testcontract *L2testcontractTransactor) Dec(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2testcontract.contract.Transact(opts, "dec")
}

// Dec is a paid mutator transaction binding the contract method 0xb3bcfa82.
//
// Solidity: function dec() returns()
func (_L2testcontract *L2testcontractSession) Dec() (*types.Transaction, error) {
	return _L2testcontract.Contract.Dec(&_L2testcontract.TransactOpts)
}

// Dec is a paid mutator transaction binding the contract method 0xb3bcfa82.
//
// Solidity: function dec() returns()
func (_L2testcontract *L2testcontractTransactorSession) Dec() (*types.Transaction, error) {
	return _L2testcontract.Contract.Dec(&_L2testcontract.TransactOpts)
}

// Incr is a paid mutator transaction binding the contract method 0x119fbbd4.
//
// Solidity: function incr() returns()
func (_L2testcontract *L2testcontractTransactor) Incr(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2testcontract.contract.Transact(opts, "incr")
}

// Incr is a paid mutator transaction binding the contract method 0x119fbbd4.
//
// Solidity: function incr() returns()
func (_L2testcontract *L2testcontractSession) Incr() (*types.Transaction, error) {
	return _L2testcontract.Contract.Incr(&_L2testcontract.TransactOpts)
}

// Incr is a paid mutator transaction binding the contract method 0x119fbbd4.
//
// Solidity: function incr() returns()
func (_L2testcontract *L2testcontractTransactorSession) Incr() (*types.Transaction, error) {
	return _L2testcontract.Contract.Incr(&_L2testcontract.TransactOpts)
}
