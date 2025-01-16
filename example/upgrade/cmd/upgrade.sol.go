// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// IUpgradeModule is an auto generated low-level Go binding around an user-defined struct.
type IUpgradeModule struct {
	ModuleName string
	Version    uint64
}

// IUpgradePlan is an auto generated low-level Go binding around an user-defined struct.
type IUpgradePlan struct {
	Name    string
	Modules []IUpgradeModule
	Info    string
	Height  uint64
	Status  uint8
}

// UpgradeMetaData contains all meta data concerning the Upgrade contract.
var UpgradeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addUpgradePlan\",\"inputs\":[{\"name\":\"plan\",\"type\":\"tuple\",\"internalType\":\"structIUpgrade.Plan\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"modules\",\"type\":\"tuple[]\",\"internalType\":\"structIUpgrade.Module[]\",\"components\":[{\"name\":\"moduleName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"info\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIUpgrade.Status\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUpgradePlan\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"plan\",\"type\":\"tuple[]\",\"internalType\":\"structIUpgrade.Plan[]\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"modules\",\"type\":\"tuple[]\",\"internalType\":\"structIUpgrade.Module[]\",\"components\":[{\"name\":\"moduleName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"info\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIUpgrade.Status\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setOwner\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUpgradePlanDone\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// UpgradeABI is the input ABI used to generate the binding from.
// Deprecated: Use UpgradeMetaData.ABI instead.
var UpgradeABI = UpgradeMetaData.ABI

// Upgrade is an auto generated Go binding around an platon contract.
type Upgrade struct {
	UpgradeCaller     // Read-only binding to the contract
	UpgradeTransactor // Write-only binding to the contract
	UpgradeFilterer   // Log filterer for contract events
}

// UpgradeCaller is an auto generated read-only Go binding around an platon contract.
type UpgradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeTransactor is an auto generated write-only Go binding around an platon contract.
type UpgradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeFilterer is an auto generated log filtering Go binding around an platon contract events.
type UpgradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type UpgradeSession struct {
	Contract     *Upgrade          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradeCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type UpgradeCallerSession struct {
	Contract *UpgradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// UpgradeTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type UpgradeTransactorSession struct {
	Contract     *UpgradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UpgradeRaw is an auto generated low-level Go binding around an platon contract.
type UpgradeRaw struct {
	Contract *Upgrade // Generic contract binding to access the raw methods on
}

// UpgradeCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type UpgradeCallerRaw struct {
	Contract *UpgradeCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type UpgradeTransactorRaw struct {
	Contract *UpgradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgrade creates a new instance of Upgrade, bound to a specific deployed contract.
func NewUpgrade(address common.Address, backend bind.ContractBackend) (*Upgrade, error) {
	contract, err := bindUpgrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Upgrade{UpgradeCaller: UpgradeCaller{contract: contract}, UpgradeTransactor: UpgradeTransactor{contract: contract}, UpgradeFilterer: UpgradeFilterer{contract: contract}}, nil
}

// NewUpgradeCaller creates a new read-only instance of Upgrade, bound to a specific deployed contract.
func NewUpgradeCaller(address common.Address, caller bind.ContractCaller) (*UpgradeCaller, error) {
	contract, err := bindUpgrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeCaller{contract: contract}, nil
}

// NewUpgradeTransactor creates a new write-only instance of Upgrade, bound to a specific deployed contract.
func NewUpgradeTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeTransactor, error) {
	contract, err := bindUpgrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeTransactor{contract: contract}, nil
}

// NewUpgradeFilterer creates a new log filterer instance of Upgrade, bound to a specific deployed contract.
func NewUpgradeFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeFilterer, error) {
	contract, err := bindUpgrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeFilterer{contract: contract}, nil
}

// bindUpgrade binds a generic wrapper to an already deployed contract.
func bindUpgrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Upgrade *UpgradeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Upgrade.Contract.UpgradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Upgrade *UpgradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Upgrade.Contract.UpgradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Upgrade *UpgradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Upgrade.Contract.UpgradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Upgrade *UpgradeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Upgrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Upgrade *UpgradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Upgrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Upgrade *UpgradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Upgrade.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Upgrade *UpgradeCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Upgrade.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Upgrade *UpgradeSession) GetOwner() (common.Address, error) {
	return _Upgrade.Contract.GetOwner(&_Upgrade.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Upgrade *UpgradeCallerSession) GetOwner() (common.Address, error) {
	return _Upgrade.Contract.GetOwner(&_Upgrade.CallOpts)
}

// GetUpgradePlan is a free data retrieval call binding the contract method 0x39f278c7.
//
// Solidity: function getUpgradePlan(uint64 height) view returns((string,(string,uint64)[],string,uint64,uint8)[] plan)
func (_Upgrade *UpgradeCaller) GetUpgradePlan(opts *bind.CallOpts, height uint64) ([]IUpgradePlan, error) {
	var out []interface{}
	err := _Upgrade.contract.Call(opts, &out, "getUpgradePlan", height)

	if err != nil {
		return *new([]IUpgradePlan), err
	}

	out0 := *abi.ConvertType(out[0], new([]IUpgradePlan)).(*[]IUpgradePlan)

	return out0, err

}

// GetUpgradePlan is a free data retrieval call binding the contract method 0x39f278c7.
//
// Solidity: function getUpgradePlan(uint64 height) view returns((string,(string,uint64)[],string,uint64,uint8)[] plan)
func (_Upgrade *UpgradeSession) GetUpgradePlan(height uint64) ([]IUpgradePlan, error) {
	return _Upgrade.Contract.GetUpgradePlan(&_Upgrade.CallOpts, height)
}

// GetUpgradePlan is a free data retrieval call binding the contract method 0x39f278c7.
//
// Solidity: function getUpgradePlan(uint64 height) view returns((string,(string,uint64)[],string,uint64,uint8)[] plan)
func (_Upgrade *UpgradeCallerSession) GetUpgradePlan(height uint64) ([]IUpgradePlan, error) {
	return _Upgrade.Contract.GetUpgradePlan(&_Upgrade.CallOpts, height)
}

// AddUpgradePlan is a paid mutator transaction binding the contract method 0xc906633d.
//
// Solidity: function addUpgradePlan((string,(string,uint64)[],string,uint64,uint8) plan) returns()
func (_Upgrade *UpgradeTransactor) AddUpgradePlan(opts *bind.TransactOpts, plan IUpgradePlan) (*types.Transaction, error) {
	return _Upgrade.contract.Transact(opts, "addUpgradePlan", plan)
}

// AddUpgradePlan is a paid mutator transaction binding the contract method 0xc906633d.
//
// Solidity: function addUpgradePlan((string,(string,uint64)[],string,uint64,uint8) plan) returns()
func (_Upgrade *UpgradeSession) AddUpgradePlan(plan IUpgradePlan) (*types.Transaction, error) {
	return _Upgrade.Contract.AddUpgradePlan(&_Upgrade.TransactOpts, plan)
}

// AddUpgradePlan is a paid mutator transaction binding the contract method 0xc906633d.
//
// Solidity: function addUpgradePlan((string,(string,uint64)[],string,uint64,uint8) plan) returns()
func (_Upgrade *UpgradeTransactorSession) AddUpgradePlan(plan IUpgradePlan) (*types.Transaction, error) {
	return _Upgrade.Contract.AddUpgradePlan(&_Upgrade.TransactOpts, plan)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Upgrade *UpgradeTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Upgrade.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Upgrade *UpgradeSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Upgrade.Contract.SetOwner(&_Upgrade.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Upgrade *UpgradeTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Upgrade.Contract.SetOwner(&_Upgrade.TransactOpts, newOwner)
}

// SetUpgradePlanDone is a paid mutator transaction binding the contract method 0xab4f4c2f.
//
// Solidity: function setUpgradePlanDone(uint64 height) returns()
func (_Upgrade *UpgradeTransactor) SetUpgradePlanDone(opts *bind.TransactOpts, height uint64) (*types.Transaction, error) {
	return _Upgrade.contract.Transact(opts, "setUpgradePlanDone", height)
}

// SetUpgradePlanDone is a paid mutator transaction binding the contract method 0xab4f4c2f.
//
// Solidity: function setUpgradePlanDone(uint64 height) returns()
func (_Upgrade *UpgradeSession) SetUpgradePlanDone(height uint64) (*types.Transaction, error) {
	return _Upgrade.Contract.SetUpgradePlanDone(&_Upgrade.TransactOpts, height)
}

// SetUpgradePlanDone is a paid mutator transaction binding the contract method 0xab4f4c2f.
//
// Solidity: function setUpgradePlanDone(uint64 height) returns()
func (_Upgrade *UpgradeTransactorSession) SetUpgradePlanDone(height uint64) (*types.Transaction, error) {
	return _Upgrade.Contract.SetUpgradePlanDone(&_Upgrade.TransactOpts, height)
}

