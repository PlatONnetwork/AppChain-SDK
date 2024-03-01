// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2upgrade

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

// L2upgradeMetaData contains all meta data concerning the L2upgrade contract.
var L2upgradeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"moduleName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"internalType\":\"structIUpgrade.Module[]\",\"name\":\"modules\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"enumIUpgrade.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIUpgrade.Plan\",\"name\":\"plan\",\"type\":\"tuple\"}],\"name\":\"addUpgradePlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getUpgradePlan\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"moduleName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"internalType\":\"structIUpgrade.Module[]\",\"name\":\"modules\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"enumIUpgrade.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIUpgrade.Plan[]\",\"name\":\"plan\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"setUpgradePlanDone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// L2upgradeABI is the input ABI used to generate the binding from.
// Deprecated: Use L2upgradeMetaData.ABI instead.
var L2upgradeABI = L2upgradeMetaData.ABI

// L2upgrade is an auto generated Go binding around an platon contract.
type L2upgrade struct {
	L2upgradeCaller     // Read-only binding to the contract
	L2upgradeTransactor // Write-only binding to the contract
	L2upgradeFilterer   // Log filterer for contract events
}

// L2upgradeCaller is an auto generated read-only Go binding around an platon contract.
type L2upgradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2upgradeTransactor is an auto generated write-only Go binding around an platon contract.
type L2upgradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2upgradeFilterer is an auto generated log filtering Go binding around an platon contract events.
type L2upgradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2upgradeSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type L2upgradeSession struct {
	Contract     *L2upgrade        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2upgradeCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type L2upgradeCallerSession struct {
	Contract *L2upgradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// L2upgradeTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type L2upgradeTransactorSession struct {
	Contract     *L2upgradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// L2upgradeRaw is an auto generated low-level Go binding around an platon contract.
type L2upgradeRaw struct {
	Contract *L2upgrade // Generic contract binding to access the raw methods on
}

// L2upgradeCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type L2upgradeCallerRaw struct {
	Contract *L2upgradeCaller // Generic read-only contract binding to access the raw methods on
}

// L2upgradeTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type L2upgradeTransactorRaw struct {
	Contract *L2upgradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2upgrade creates a new instance of L2upgrade, bound to a specific deployed contract.
func NewL2upgrade(address common.Address, backend bind.ContractBackend) (*L2upgrade, error) {
	contract, err := bindL2upgrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2upgrade{L2upgradeCaller: L2upgradeCaller{contract: contract}, L2upgradeTransactor: L2upgradeTransactor{contract: contract}, L2upgradeFilterer: L2upgradeFilterer{contract: contract}}, nil
}

// NewL2upgradeCaller creates a new read-only instance of L2upgrade, bound to a specific deployed contract.
func NewL2upgradeCaller(address common.Address, caller bind.ContractCaller) (*L2upgradeCaller, error) {
	contract, err := bindL2upgrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2upgradeCaller{contract: contract}, nil
}

// NewL2upgradeTransactor creates a new write-only instance of L2upgrade, bound to a specific deployed contract.
func NewL2upgradeTransactor(address common.Address, transactor bind.ContractTransactor) (*L2upgradeTransactor, error) {
	contract, err := bindL2upgrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2upgradeTransactor{contract: contract}, nil
}

// NewL2upgradeFilterer creates a new log filterer instance of L2upgrade, bound to a specific deployed contract.
func NewL2upgradeFilterer(address common.Address, filterer bind.ContractFilterer) (*L2upgradeFilterer, error) {
	contract, err := bindL2upgrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2upgradeFilterer{contract: contract}, nil
}

// bindL2upgrade binds a generic wrapper to an already deployed contract.
func bindL2upgrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2upgradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2upgrade *L2upgradeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2upgrade.Contract.L2upgradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2upgrade *L2upgradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2upgrade.Contract.L2upgradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2upgrade *L2upgradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2upgrade.Contract.L2upgradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2upgrade *L2upgradeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2upgrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2upgrade *L2upgradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2upgrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2upgrade *L2upgradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2upgrade.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_L2upgrade *L2upgradeCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2upgrade.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_L2upgrade *L2upgradeSession) GetOwner() (common.Address, error) {
	return _L2upgrade.Contract.GetOwner(&_L2upgrade.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_L2upgrade *L2upgradeCallerSession) GetOwner() (common.Address, error) {
	return _L2upgrade.Contract.GetOwner(&_L2upgrade.CallOpts)
}

// GetUpgradePlan is a free data retrieval call binding the contract method 0x39f278c7.
//
// Solidity: function getUpgradePlan(uint64 height) view returns((string,(string,uint64)[],string,uint64,uint8)[] plan)
func (_L2upgrade *L2upgradeCaller) GetUpgradePlan(opts *bind.CallOpts, height uint64) ([]IUpgradePlan, error) {
	var out []interface{}
	err := _L2upgrade.contract.Call(opts, &out, "getUpgradePlan", height)

	if err != nil {
		return *new([]IUpgradePlan), err
	}

	out0 := *abi.ConvertType(out[0], new([]IUpgradePlan)).(*[]IUpgradePlan)

	return out0, err

}

// GetUpgradePlan is a free data retrieval call binding the contract method 0x39f278c7.
//
// Solidity: function getUpgradePlan(uint64 height) view returns((string,(string,uint64)[],string,uint64,uint8)[] plan)
func (_L2upgrade *L2upgradeSession) GetUpgradePlan(height uint64) ([]IUpgradePlan, error) {
	return _L2upgrade.Contract.GetUpgradePlan(&_L2upgrade.CallOpts, height)
}

// GetUpgradePlan is a free data retrieval call binding the contract method 0x39f278c7.
//
// Solidity: function getUpgradePlan(uint64 height) view returns((string,(string,uint64)[],string,uint64,uint8)[] plan)
func (_L2upgrade *L2upgradeCallerSession) GetUpgradePlan(height uint64) ([]IUpgradePlan, error) {
	return _L2upgrade.Contract.GetUpgradePlan(&_L2upgrade.CallOpts, height)
}

// AddUpgradePlan is a paid mutator transaction binding the contract method 0xc906633d.
//
// Solidity: function addUpgradePlan((string,(string,uint64)[],string,uint64,uint8) plan) returns()
func (_L2upgrade *L2upgradeTransactor) AddUpgradePlan(opts *bind.TransactOpts, plan IUpgradePlan) (*types.Transaction, error) {
	return _L2upgrade.contract.Transact(opts, "addUpgradePlan", plan)
}

// AddUpgradePlan is a paid mutator transaction binding the contract method 0xc906633d.
//
// Solidity: function addUpgradePlan((string,(string,uint64)[],string,uint64,uint8) plan) returns()
func (_L2upgrade *L2upgradeSession) AddUpgradePlan(plan IUpgradePlan) (*types.Transaction, error) {
	return _L2upgrade.Contract.AddUpgradePlan(&_L2upgrade.TransactOpts, plan)
}

// AddUpgradePlan is a paid mutator transaction binding the contract method 0xc906633d.
//
// Solidity: function addUpgradePlan((string,(string,uint64)[],string,uint64,uint8) plan) returns()
func (_L2upgrade *L2upgradeTransactorSession) AddUpgradePlan(plan IUpgradePlan) (*types.Transaction, error) {
	return _L2upgrade.Contract.AddUpgradePlan(&_L2upgrade.TransactOpts, plan)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_L2upgrade *L2upgradeTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2upgrade.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_L2upgrade *L2upgradeSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _L2upgrade.Contract.SetOwner(&_L2upgrade.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_L2upgrade *L2upgradeTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _L2upgrade.Contract.SetOwner(&_L2upgrade.TransactOpts, newOwner)
}

// SetUpgradePlanDone is a paid mutator transaction binding the contract method 0xab4f4c2f.
//
// Solidity: function setUpgradePlanDone(uint64 height) returns()
func (_L2upgrade *L2upgradeTransactor) SetUpgradePlanDone(opts *bind.TransactOpts, height uint64) (*types.Transaction, error) {
	return _L2upgrade.contract.Transact(opts, "setUpgradePlanDone", height)
}

// SetUpgradePlanDone is a paid mutator transaction binding the contract method 0xab4f4c2f.
//
// Solidity: function setUpgradePlanDone(uint64 height) returns()
func (_L2upgrade *L2upgradeSession) SetUpgradePlanDone(height uint64) (*types.Transaction, error) {
	return _L2upgrade.Contract.SetUpgradePlanDone(&_L2upgrade.TransactOpts, height)
}

// SetUpgradePlanDone is a paid mutator transaction binding the contract method 0xab4f4c2f.
//
// Solidity: function setUpgradePlanDone(uint64 height) returns()
func (_L2upgrade *L2upgradeTransactorSession) SetUpgradePlanDone(height uint64) (*types.Transaction, error) {
	return _L2upgrade.Contract.SetUpgradePlanDone(&_L2upgrade.TransactOpts, height)
}
