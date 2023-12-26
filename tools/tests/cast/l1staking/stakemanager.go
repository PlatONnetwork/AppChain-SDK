// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l1staking

import (
	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = platon.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ValidatorStake is an auto generated low-level Go binding around an user-defined struct.
type ValidatorStake struct {
	Owner          common.Address
	StakeAmount    *big.Int
	CommissionRate *big.Int
	PubKey         hexutil.Bytes
	BlsKey         hexutil.Bytes
}

// StakemanagerABI is the input ABI used to generate the binding from.
const StakemanagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegationRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegationWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ValidatorStakeSlashed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegater\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegateFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"delegationOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minDelegate\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registryManager\",\"outputs\":[{\"internalType\":\"contractIRegistryManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegater\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"releaseDelegationOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"releaseStakeOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"slashStakeOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structValidatorStake\",\"name\":\"stake\",\"type\":\"tuple\"}],\"name\":\"stakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"stakeOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"validatorDelegationOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawableDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawableStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Stakemanager is an auto generated Go binding around an platon contract.
type Stakemanager struct {
	StakemanagerCaller     // Read-only binding to the contract
	StakemanagerTransactor // Write-only binding to the contract
	StakemanagerFilterer   // Log filterer for contract events
}

// StakemanagerCaller is an auto generated read-only Go binding around an platon contract.
type StakemanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakemanagerTransactor is an auto generated write-only Go binding around an platon contract.
type StakemanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakemanagerFilterer is an auto generated log filtering Go binding around an platon contract events.
type StakemanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakemanagerSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type StakemanagerSession struct {
	Contract     *Stakemanager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakemanagerCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type StakemanagerCallerSession struct {
	Contract *StakemanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StakemanagerTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type StakemanagerTransactorSession struct {
	Contract     *StakemanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StakemanagerRaw is an auto generated low-level Go binding around an platon contract.
type StakemanagerRaw struct {
	Contract *Stakemanager // Generic contract binding to access the raw methods on
}

// StakemanagerCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type StakemanagerCallerRaw struct {
	Contract *StakemanagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakemanagerTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type StakemanagerTransactorRaw struct {
	Contract *StakemanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakemanager creates a new instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanager(address common.Address, backend bind.ContractBackend) (*Stakemanager, error) {
	contract, err := bindStakemanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stakemanager{StakemanagerCaller: StakemanagerCaller{contract: contract}, StakemanagerTransactor: StakemanagerTransactor{contract: contract}, StakemanagerFilterer: StakemanagerFilterer{contract: contract}}, nil
}

// NewStakemanagerCaller creates a new read-only instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanagerCaller(address common.Address, caller bind.ContractCaller) (*StakemanagerCaller, error) {
	contract, err := bindStakemanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakemanagerCaller{contract: contract}, nil
}

// NewStakemanagerTransactor creates a new write-only instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakemanagerTransactor, error) {
	contract, err := bindStakemanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakemanagerTransactor{contract: contract}, nil
}

// NewStakemanagerFilterer creates a new log filterer instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakemanagerFilterer, error) {
	contract, err := bindStakemanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakemanagerFilterer{contract: contract}, nil
}

// bindStakemanager binds a generic wrapper to an already deployed contract.
func bindStakemanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakemanagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakemanager *StakemanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakemanager.Contract.StakemanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakemanager *StakemanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.Contract.StakemanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakemanager *StakemanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakemanager.Contract.StakemanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakemanager *StakemanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakemanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakemanager *StakemanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakemanager *StakemanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakemanager.Contract.contract.Transact(opts, method, params...)
}

// DelegationOf is a free data retrieval call binding the contract method 0xf837123e.
//
// Solidity: function delegationOf(address validator) view returns(uint256 amount)
func (_Stakemanager *StakemanagerCaller) DelegationOf(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "delegationOf", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationOf is a free data retrieval call binding the contract method 0xf837123e.
//
// Solidity: function delegationOf(address validator) view returns(uint256 amount)
func (_Stakemanager *StakemanagerSession) DelegationOf(validator common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.DelegationOf(&_Stakemanager.CallOpts, validator)
}

// DelegationOf is a free data retrieval call binding the contract method 0xf837123e.
//
// Solidity: function delegationOf(address validator) view returns(uint256 amount)
func (_Stakemanager *StakemanagerCallerSession) DelegationOf(validator common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.DelegationOf(&_Stakemanager.CallOpts, validator)
}

// MinDelegate is a free data retrieval call binding the contract method 0x45255c05.
//
// Solidity: function minDelegate() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) MinDelegate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "minDelegate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegate is a free data retrieval call binding the contract method 0x45255c05.
//
// Solidity: function minDelegate() view returns(uint256)
func (_Stakemanager *StakemanagerSession) MinDelegate() (*big.Int, error) {
	return _Stakemanager.Contract.MinDelegate(&_Stakemanager.CallOpts)
}

// MinDelegate is a free data retrieval call binding the contract method 0x45255c05.
//
// Solidity: function minDelegate() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) MinDelegate() (*big.Int, error) {
	return _Stakemanager.Contract.MinDelegate(&_Stakemanager.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "minStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Stakemanager *StakemanagerSession) MinStake() (*big.Int, error) {
	return _Stakemanager.Contract.MinStake(&_Stakemanager.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) MinStake() (*big.Int, error) {
	return _Stakemanager.Contract.MinStake(&_Stakemanager.CallOpts)
}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_Stakemanager *StakemanagerCaller) RegistryManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "registryManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_Stakemanager *StakemanagerSession) RegistryManager() (common.Address, error) {
	return _Stakemanager.Contract.RegistryManager(&_Stakemanager.CallOpts)
}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_Stakemanager *StakemanagerCallerSession) RegistryManager() (common.Address, error) {
	return _Stakemanager.Contract.RegistryManager(&_Stakemanager.CallOpts)
}

// StakeOf is a free data retrieval call binding the contract method 0x42623360.
//
// Solidity: function stakeOf(address validator) view returns(uint256 amount)
func (_Stakemanager *StakemanagerCaller) StakeOf(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "stakeOf", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeOf is a free data retrieval call binding the contract method 0x42623360.
//
// Solidity: function stakeOf(address validator) view returns(uint256 amount)
func (_Stakemanager *StakemanagerSession) StakeOf(validator common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.StakeOf(&_Stakemanager.CallOpts, validator)
}

// StakeOf is a free data retrieval call binding the contract method 0x42623360.
//
// Solidity: function stakeOf(address validator) view returns(uint256 amount)
func (_Stakemanager *StakemanagerCallerSession) StakeOf(validator common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.StakeOf(&_Stakemanager.CallOpts, validator)
}

// TotalDelegation is a free data retrieval call binding the contract method 0xe3c3ae58.
//
// Solidity: function totalDelegation() view returns(uint256 amount)
func (_Stakemanager *StakemanagerCaller) TotalDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "totalDelegation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDelegation is a free data retrieval call binding the contract method 0xe3c3ae58.
//
// Solidity: function totalDelegation() view returns(uint256 amount)
func (_Stakemanager *StakemanagerSession) TotalDelegation() (*big.Int, error) {
	return _Stakemanager.Contract.TotalDelegation(&_Stakemanager.CallOpts)
}

// TotalDelegation is a free data retrieval call binding the contract method 0xe3c3ae58.
//
// Solidity: function totalDelegation() view returns(uint256 amount)
func (_Stakemanager *StakemanagerCallerSession) TotalDelegation() (*big.Int, error) {
	return _Stakemanager.Contract.TotalDelegation(&_Stakemanager.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256 amount)
func (_Stakemanager *StakemanagerCaller) TotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "totalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256 amount)
func (_Stakemanager *StakemanagerSession) TotalStake() (*big.Int, error) {
	return _Stakemanager.Contract.TotalStake(&_Stakemanager.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256 amount)
func (_Stakemanager *StakemanagerCallerSession) TotalStake() (*big.Int, error) {
	return _Stakemanager.Contract.TotalStake(&_Stakemanager.CallOpts)
}

// ValidatorDelegationOf is a free data retrieval call binding the contract method 0xef86e74c.
//
// Solidity: function validatorDelegationOf(address validator) view returns(uint256 amount)
func (_Stakemanager *StakemanagerCaller) ValidatorDelegationOf(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "validatorDelegationOf", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorDelegationOf is a free data retrieval call binding the contract method 0xef86e74c.
//
// Solidity: function validatorDelegationOf(address validator) view returns(uint256 amount)
func (_Stakemanager *StakemanagerSession) ValidatorDelegationOf(validator common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.ValidatorDelegationOf(&_Stakemanager.CallOpts, validator)
}

// ValidatorDelegationOf is a free data retrieval call binding the contract method 0xef86e74c.
//
// Solidity: function validatorDelegationOf(address validator) view returns(uint256 amount)
func (_Stakemanager *StakemanagerCallerSession) ValidatorDelegationOf(validator common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.ValidatorDelegationOf(&_Stakemanager.CallOpts, validator)
}

// WithdrawableDelegation is a free data retrieval call binding the contract method 0xf7b9a9a2.
//
// Solidity: function withdrawableDelegation(address validator) view returns(uint256 amount)
func (_Stakemanager *StakemanagerCaller) WithdrawableDelegation(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "withdrawableDelegation", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableDelegation is a free data retrieval call binding the contract method 0xf7b9a9a2.
//
// Solidity: function withdrawableDelegation(address validator) view returns(uint256 amount)
func (_Stakemanager *StakemanagerSession) WithdrawableDelegation(validator common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.WithdrawableDelegation(&_Stakemanager.CallOpts, validator)
}

// WithdrawableDelegation is a free data retrieval call binding the contract method 0xf7b9a9a2.
//
// Solidity: function withdrawableDelegation(address validator) view returns(uint256 amount)
func (_Stakemanager *StakemanagerCallerSession) WithdrawableDelegation(validator common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.WithdrawableDelegation(&_Stakemanager.CallOpts, validator)
}

// WithdrawableStake is a free data retrieval call binding the contract method 0xd5364bbf.
//
// Solidity: function withdrawableStake(address validator) view returns(uint256 amount)
func (_Stakemanager *StakemanagerCaller) WithdrawableStake(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "withdrawableStake", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableStake is a free data retrieval call binding the contract method 0xd5364bbf.
//
// Solidity: function withdrawableStake(address validator) view returns(uint256 amount)
func (_Stakemanager *StakemanagerSession) WithdrawableStake(validator common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.WithdrawableStake(&_Stakemanager.CallOpts, validator)
}

// WithdrawableStake is a free data retrieval call binding the contract method 0xd5364bbf.
//
// Solidity: function withdrawableStake(address validator) view returns(uint256 amount)
func (_Stakemanager *StakemanagerCallerSession) WithdrawableStake(validator common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.WithdrawableStake(&_Stakemanager.CallOpts, validator)
}

// AddStake is a paid mutator transaction binding the contract method 0x6374299e.
//
// Solidity: function addStake(address validator, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactor) AddStake(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "addStake", validator, amount)
}

// AddStake is a paid mutator transaction binding the contract method 0x6374299e.
//
// Solidity: function addStake(address validator, uint256 amount) returns()
func (_Stakemanager *StakemanagerSession) AddStake(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.AddStake(&_Stakemanager.TransactOpts, validator, amount)
}

// AddStake is a paid mutator transaction binding the contract method 0x6374299e.
//
// Solidity: function addStake(address validator, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) AddStake(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.AddStake(&_Stakemanager.TransactOpts, validator, amount)
}

// DelegateFor is a paid mutator transaction binding the contract method 0x74bbe1c8.
//
// Solidity: function delegateFor(address validator, address delegater, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactor) DelegateFor(opts *bind.TransactOpts, validator common.Address, delegater common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "delegateFor", validator, delegater, amount)
}

// DelegateFor is a paid mutator transaction binding the contract method 0x74bbe1c8.
//
// Solidity: function delegateFor(address validator, address delegater, uint256 amount) returns()
func (_Stakemanager *StakemanagerSession) DelegateFor(validator common.Address, delegater common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.DelegateFor(&_Stakemanager.TransactOpts, validator, delegater, amount)
}

// DelegateFor is a paid mutator transaction binding the contract method 0x74bbe1c8.
//
// Solidity: function delegateFor(address validator, address delegater, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) DelegateFor(validator common.Address, delegater common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.DelegateFor(&_Stakemanager.TransactOpts, validator, delegater, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _registryManager, uint256 _minStake, uint256 _minDelegate) returns()
func (_Stakemanager *StakemanagerTransactor) Initialize(opts *bind.TransactOpts, _registryManager common.Address, _minStake *big.Int, _minDelegate *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "initialize", _registryManager, _minStake, _minDelegate)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _registryManager, uint256 _minStake, uint256 _minDelegate) returns()
func (_Stakemanager *StakemanagerSession) Initialize(_registryManager common.Address, _minStake *big.Int, _minDelegate *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Initialize(&_Stakemanager.TransactOpts, _registryManager, _minStake, _minDelegate)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _registryManager, uint256 _minStake, uint256 _minDelegate) returns()
func (_Stakemanager *StakemanagerTransactorSession) Initialize(_registryManager common.Address, _minStake *big.Int, _minDelegate *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Initialize(&_Stakemanager.TransactOpts, _registryManager, _minStake, _minDelegate)
}

// ReleaseDelegationOf is a paid mutator transaction binding the contract method 0x5a601d91.
//
// Solidity: function releaseDelegationOf(address validator, address delegater, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactor) ReleaseDelegationOf(opts *bind.TransactOpts, validator common.Address, delegater common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "releaseDelegationOf", validator, delegater, amount)
}

// ReleaseDelegationOf is a paid mutator transaction binding the contract method 0x5a601d91.
//
// Solidity: function releaseDelegationOf(address validator, address delegater, uint256 amount) returns()
func (_Stakemanager *StakemanagerSession) ReleaseDelegationOf(validator common.Address, delegater common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ReleaseDelegationOf(&_Stakemanager.TransactOpts, validator, delegater, amount)
}

// ReleaseDelegationOf is a paid mutator transaction binding the contract method 0x5a601d91.
//
// Solidity: function releaseDelegationOf(address validator, address delegater, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) ReleaseDelegationOf(validator common.Address, delegater common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ReleaseDelegationOf(&_Stakemanager.TransactOpts, validator, delegater, amount)
}

// ReleaseStakeOf is a paid mutator transaction binding the contract method 0x3651bb1d.
//
// Solidity: function releaseStakeOf(address validator, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactor) ReleaseStakeOf(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "releaseStakeOf", validator, amount)
}

// ReleaseStakeOf is a paid mutator transaction binding the contract method 0x3651bb1d.
//
// Solidity: function releaseStakeOf(address validator, uint256 amount) returns()
func (_Stakemanager *StakemanagerSession) ReleaseStakeOf(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ReleaseStakeOf(&_Stakemanager.TransactOpts, validator, amount)
}

// ReleaseStakeOf is a paid mutator transaction binding the contract method 0x3651bb1d.
//
// Solidity: function releaseStakeOf(address validator, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) ReleaseStakeOf(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ReleaseStakeOf(&_Stakemanager.TransactOpts, validator, amount)
}

// SlashStakeOf is a paid mutator transaction binding the contract method 0x8028a6db.
//
// Solidity: function slashStakeOf(address validator, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactor) SlashStakeOf(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "slashStakeOf", validator, amount)
}

// SlashStakeOf is a paid mutator transaction binding the contract method 0x8028a6db.
//
// Solidity: function slashStakeOf(address validator, uint256 amount) returns()
func (_Stakemanager *StakemanagerSession) SlashStakeOf(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.SlashStakeOf(&_Stakemanager.TransactOpts, validator, amount)
}

// SlashStakeOf is a paid mutator transaction binding the contract method 0x8028a6db.
//
// Solidity: function slashStakeOf(address validator, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) SlashStakeOf(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.SlashStakeOf(&_Stakemanager.TransactOpts, validator, amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x512b91b6.
//
// Solidity: function stakeFor((address,uint256,uint256,bytes,bytes) stake) returns()
func (_Stakemanager *StakemanagerTransactor) StakeFor(opts *bind.TransactOpts, stake ValidatorStake) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "stakeFor", stake)
}

// StakeFor is a paid mutator transaction binding the contract method 0x512b91b6.
//
// Solidity: function stakeFor((address,uint256,uint256,bytes,bytes) stake) returns()
func (_Stakemanager *StakemanagerSession) StakeFor(stake ValidatorStake) (*types.Transaction, error) {
	return _Stakemanager.Contract.StakeFor(&_Stakemanager.TransactOpts, stake)
}

// StakeFor is a paid mutator transaction binding the contract method 0x512b91b6.
//
// Solidity: function stakeFor((address,uint256,uint256,bytes,bytes) stake) returns()
func (_Stakemanager *StakemanagerTransactorSession) StakeFor(stake ValidatorStake) (*types.Transaction, error) {
	return _Stakemanager.Contract.StakeFor(&_Stakemanager.TransactOpts, stake)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x3d9a6dcb.
//
// Solidity: function withdrawDelegation(address validator, address to, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactor) WithdrawDelegation(opts *bind.TransactOpts, validator common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "withdrawDelegation", validator, to, amount)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x3d9a6dcb.
//
// Solidity: function withdrawDelegation(address validator, address to, uint256 amount) returns()
func (_Stakemanager *StakemanagerSession) WithdrawDelegation(validator common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.WithdrawDelegation(&_Stakemanager.TransactOpts, validator, to, amount)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x3d9a6dcb.
//
// Solidity: function withdrawDelegation(address validator, address to, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) WithdrawDelegation(validator common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.WithdrawDelegation(&_Stakemanager.TransactOpts, validator, to, amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x0c1e8bf7.
//
// Solidity: function withdrawStake(address validator, address to, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactor) WithdrawStake(opts *bind.TransactOpts, validator common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "withdrawStake", validator, to, amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x0c1e8bf7.
//
// Solidity: function withdrawStake(address validator, address to, uint256 amount) returns()
func (_Stakemanager *StakemanagerSession) WithdrawStake(validator common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.WithdrawStake(&_Stakemanager.TransactOpts, validator, to, amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x0c1e8bf7.
//
// Solidity: function withdrawStake(address validator, address to, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) WithdrawStake(validator common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.WithdrawStake(&_Stakemanager.TransactOpts, validator, to, amount)
}

// StakemanagerDelegationAddedIterator is returned from FilterDelegationAdded and is used to iterate over the raw logs and unpacked data for DelegationAdded events raised by the Stakemanager contract.
type StakemanagerDelegationAddedIterator struct {
	Event *StakemanagerDelegationAdded // Event containing the contract specifics and raw log

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
func (it *StakemanagerDelegationAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerDelegationAdded)
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
		it.Event = new(StakemanagerDelegationAdded)
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
func (it *StakemanagerDelegationAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerDelegationAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerDelegationAdded represents a DelegationAdded event raised by the Stakemanager contract.
type StakemanagerDelegationAdded struct {
	Validator common.Address
	Delegater common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegationAdded is a free log retrieval operation binding the contract event 0x52467f14b857734001c77e6f125dac41b45798837c9fc9adfe3a5b394c77a0e9.
//
// Solidity: event DelegationAdded(address indexed validator, address indexed delegater, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) FilterDelegationAdded(opts *bind.FilterOpts, validator []common.Address, delegater []common.Address) (*StakemanagerDelegationAddedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegaterRule []interface{}
	for _, delegaterItem := range delegater {
		delegaterRule = append(delegaterRule, delegaterItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "DelegationAdded", validatorRule, delegaterRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerDelegationAddedIterator{contract: _Stakemanager.contract, event: "DelegationAdded", logs: logs, sub: sub}, nil
}

// WatchDelegationAdded is a free log subscription operation binding the contract event 0x52467f14b857734001c77e6f125dac41b45798837c9fc9adfe3a5b394c77a0e9.
//
// Solidity: event DelegationAdded(address indexed validator, address indexed delegater, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) WatchDelegationAdded(opts *bind.WatchOpts, sink chan<- *StakemanagerDelegationAdded, validator []common.Address, delegater []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegaterRule []interface{}
	for _, delegaterItem := range delegater {
		delegaterRule = append(delegaterRule, delegaterItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "DelegationAdded", validatorRule, delegaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerDelegationAdded)
				if err := _Stakemanager.contract.UnpackLog(event, "DelegationAdded", log); err != nil {
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

// ParseDelegationAdded is a log parse operation binding the contract event 0x52467f14b857734001c77e6f125dac41b45798837c9fc9adfe3a5b394c77a0e9.
//
// Solidity: event DelegationAdded(address indexed validator, address indexed delegater, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) ParseDelegationAdded(log types.Log) (*StakemanagerDelegationAdded, error) {
	event := new(StakemanagerDelegationAdded)
	if err := _Stakemanager.contract.UnpackLog(event, "DelegationAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerDelegationRemovedIterator is returned from FilterDelegationRemoved and is used to iterate over the raw logs and unpacked data for DelegationRemoved events raised by the Stakemanager contract.
type StakemanagerDelegationRemovedIterator struct {
	Event *StakemanagerDelegationRemoved // Event containing the contract specifics and raw log

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
func (it *StakemanagerDelegationRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerDelegationRemoved)
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
		it.Event = new(StakemanagerDelegationRemoved)
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
func (it *StakemanagerDelegationRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerDelegationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerDelegationRemoved represents a DelegationRemoved event raised by the Stakemanager contract.
type StakemanagerDelegationRemoved struct {
	Validator common.Address
	Delegater common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegationRemoved is a free log retrieval operation binding the contract event 0xbf340c6e47f6acc1fa5fcad9ef75c1e4bd8d91e7313667c3c9859f230fc7f883.
//
// Solidity: event DelegationRemoved(address indexed validator, address indexed delegater, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) FilterDelegationRemoved(opts *bind.FilterOpts, validator []common.Address, delegater []common.Address) (*StakemanagerDelegationRemovedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegaterRule []interface{}
	for _, delegaterItem := range delegater {
		delegaterRule = append(delegaterRule, delegaterItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "DelegationRemoved", validatorRule, delegaterRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerDelegationRemovedIterator{contract: _Stakemanager.contract, event: "DelegationRemoved", logs: logs, sub: sub}, nil
}

// WatchDelegationRemoved is a free log subscription operation binding the contract event 0xbf340c6e47f6acc1fa5fcad9ef75c1e4bd8d91e7313667c3c9859f230fc7f883.
//
// Solidity: event DelegationRemoved(address indexed validator, address indexed delegater, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) WatchDelegationRemoved(opts *bind.WatchOpts, sink chan<- *StakemanagerDelegationRemoved, validator []common.Address, delegater []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegaterRule []interface{}
	for _, delegaterItem := range delegater {
		delegaterRule = append(delegaterRule, delegaterItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "DelegationRemoved", validatorRule, delegaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerDelegationRemoved)
				if err := _Stakemanager.contract.UnpackLog(event, "DelegationRemoved", log); err != nil {
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

// ParseDelegationRemoved is a log parse operation binding the contract event 0xbf340c6e47f6acc1fa5fcad9ef75c1e4bd8d91e7313667c3c9859f230fc7f883.
//
// Solidity: event DelegationRemoved(address indexed validator, address indexed delegater, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) ParseDelegationRemoved(log types.Log) (*StakemanagerDelegationRemoved, error) {
	event := new(StakemanagerDelegationRemoved)
	if err := _Stakemanager.contract.UnpackLog(event, "DelegationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerDelegationWithdrawnIterator is returned from FilterDelegationWithdrawn and is used to iterate over the raw logs and unpacked data for DelegationWithdrawn events raised by the Stakemanager contract.
type StakemanagerDelegationWithdrawnIterator struct {
	Event *StakemanagerDelegationWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakemanagerDelegationWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerDelegationWithdrawn)
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
		it.Event = new(StakemanagerDelegationWithdrawn)
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
func (it *StakemanagerDelegationWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerDelegationWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerDelegationWithdrawn represents a DelegationWithdrawn event raised by the Stakemanager contract.
type StakemanagerDelegationWithdrawn struct {
	Validator common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegationWithdrawn is a free log retrieval operation binding the contract event 0xaa0001b0e2e76b9a1b925257fefa58b756aebc52d2d7c8a85ea5beacc77a2100.
//
// Solidity: event DelegationWithdrawn(address indexed validator, address indexed recipient, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) FilterDelegationWithdrawn(opts *bind.FilterOpts, validator []common.Address, recipient []common.Address) (*StakemanagerDelegationWithdrawnIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "DelegationWithdrawn", validatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerDelegationWithdrawnIterator{contract: _Stakemanager.contract, event: "DelegationWithdrawn", logs: logs, sub: sub}, nil
}

// WatchDelegationWithdrawn is a free log subscription operation binding the contract event 0xaa0001b0e2e76b9a1b925257fefa58b756aebc52d2d7c8a85ea5beacc77a2100.
//
// Solidity: event DelegationWithdrawn(address indexed validator, address indexed recipient, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) WatchDelegationWithdrawn(opts *bind.WatchOpts, sink chan<- *StakemanagerDelegationWithdrawn, validator []common.Address, recipient []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "DelegationWithdrawn", validatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerDelegationWithdrawn)
				if err := _Stakemanager.contract.UnpackLog(event, "DelegationWithdrawn", log); err != nil {
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

// ParseDelegationWithdrawn is a log parse operation binding the contract event 0xaa0001b0e2e76b9a1b925257fefa58b756aebc52d2d7c8a85ea5beacc77a2100.
//
// Solidity: event DelegationWithdrawn(address indexed validator, address indexed recipient, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) ParseDelegationWithdrawn(log types.Log) (*StakemanagerDelegationWithdrawn, error) {
	event := new(StakemanagerDelegationWithdrawn)
	if err := _Stakemanager.contract.UnpackLog(event, "DelegationWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Stakemanager contract.
type StakemanagerInitializedIterator struct {
	Event *StakemanagerInitialized // Event containing the contract specifics and raw log

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
func (it *StakemanagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerInitialized)
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
		it.Event = new(StakemanagerInitialized)
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
func (it *StakemanagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerInitialized represents a Initialized event raised by the Stakemanager contract.
type StakemanagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Stakemanager *StakemanagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakemanagerInitializedIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakemanagerInitializedIterator{contract: _Stakemanager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Stakemanager *StakemanagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakemanagerInitialized) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerInitialized)
				if err := _Stakemanager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Stakemanager *StakemanagerFilterer) ParseInitialized(log types.Log) (*StakemanagerInitialized, error) {
	event := new(StakemanagerInitialized)
	if err := _Stakemanager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerStakeAddedIterator is returned from FilterStakeAdded and is used to iterate over the raw logs and unpacked data for StakeAdded events raised by the Stakemanager contract.
type StakemanagerStakeAddedIterator struct {
	Event *StakemanagerStakeAdded // Event containing the contract specifics and raw log

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
func (it *StakemanagerStakeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerStakeAdded)
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
		it.Event = new(StakemanagerStakeAdded)
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
func (it *StakemanagerStakeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerStakeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerStakeAdded represents a StakeAdded event raised by the Stakemanager contract.
type StakemanagerStakeAdded struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeAdded is a free log retrieval operation binding the contract event 0x7c717985ac273e663b7f3050f5b15a4388ff6ed952338954f650e2093e13937f.
//
// Solidity: event StakeAdded(address indexed validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) FilterStakeAdded(opts *bind.FilterOpts, validator []common.Address) (*StakemanagerStakeAddedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "StakeAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerStakeAddedIterator{contract: _Stakemanager.contract, event: "StakeAdded", logs: logs, sub: sub}, nil
}

// WatchStakeAdded is a free log subscription operation binding the contract event 0x7c717985ac273e663b7f3050f5b15a4388ff6ed952338954f650e2093e13937f.
//
// Solidity: event StakeAdded(address indexed validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) WatchStakeAdded(opts *bind.WatchOpts, sink chan<- *StakemanagerStakeAdded, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "StakeAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerStakeAdded)
				if err := _Stakemanager.contract.UnpackLog(event, "StakeAdded", log); err != nil {
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

// ParseStakeAdded is a log parse operation binding the contract event 0x7c717985ac273e663b7f3050f5b15a4388ff6ed952338954f650e2093e13937f.
//
// Solidity: event StakeAdded(address indexed validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) ParseStakeAdded(log types.Log) (*StakemanagerStakeAdded, error) {
	event := new(StakemanagerStakeAdded)
	if err := _Stakemanager.contract.UnpackLog(event, "StakeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerStakeRemovedIterator is returned from FilterStakeRemoved and is used to iterate over the raw logs and unpacked data for StakeRemoved events raised by the Stakemanager contract.
type StakemanagerStakeRemovedIterator struct {
	Event *StakemanagerStakeRemoved // Event containing the contract specifics and raw log

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
func (it *StakemanagerStakeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerStakeRemoved)
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
		it.Event = new(StakemanagerStakeRemoved)
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
func (it *StakemanagerStakeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerStakeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerStakeRemoved represents a StakeRemoved event raised by the Stakemanager contract.
type StakemanagerStakeRemoved struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeRemoved is a free log retrieval operation binding the contract event 0xa018dcbc822f59fb0d0c3e7a86c8e4259b9676cdea9e5fc26279b9c4c5d86eef.
//
// Solidity: event StakeRemoved(address indexed validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) FilterStakeRemoved(opts *bind.FilterOpts, validator []common.Address) (*StakemanagerStakeRemovedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "StakeRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerStakeRemovedIterator{contract: _Stakemanager.contract, event: "StakeRemoved", logs: logs, sub: sub}, nil
}

// WatchStakeRemoved is a free log subscription operation binding the contract event 0xa018dcbc822f59fb0d0c3e7a86c8e4259b9676cdea9e5fc26279b9c4c5d86eef.
//
// Solidity: event StakeRemoved(address indexed validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) WatchStakeRemoved(opts *bind.WatchOpts, sink chan<- *StakemanagerStakeRemoved, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "StakeRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerStakeRemoved)
				if err := _Stakemanager.contract.UnpackLog(event, "StakeRemoved", log); err != nil {
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

// ParseStakeRemoved is a log parse operation binding the contract event 0xa018dcbc822f59fb0d0c3e7a86c8e4259b9676cdea9e5fc26279b9c4c5d86eef.
//
// Solidity: event StakeRemoved(address indexed validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) ParseStakeRemoved(log types.Log) (*StakemanagerStakeRemoved, error) {
	event := new(StakemanagerStakeRemoved)
	if err := _Stakemanager.contract.UnpackLog(event, "StakeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the Stakemanager contract.
type StakemanagerStakeWithdrawnIterator struct {
	Event *StakemanagerStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakemanagerStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerStakeWithdrawn)
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
		it.Event = new(StakemanagerStakeWithdrawn)
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
func (it *StakemanagerStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerStakeWithdrawn represents a StakeWithdrawn event raised by the Stakemanager contract.
type StakemanagerStakeWithdrawn struct {
	Validator common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed validator, address indexed recipient, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, validator []common.Address, recipient []common.Address) (*StakemanagerStakeWithdrawnIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "StakeWithdrawn", validatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerStakeWithdrawnIterator{contract: _Stakemanager.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed validator, address indexed recipient, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *StakemanagerStakeWithdrawn, validator []common.Address, recipient []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "StakeWithdrawn", validatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerStakeWithdrawn)
				if err := _Stakemanager.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed validator, address indexed recipient, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) ParseStakeWithdrawn(log types.Log) (*StakemanagerStakeWithdrawn, error) {
	event := new(StakemanagerStakeWithdrawn)
	if err := _Stakemanager.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerValidatorStakeSlashedIterator is returned from FilterValidatorStakeSlashed and is used to iterate over the raw logs and unpacked data for ValidatorStakeSlashed events raised by the Stakemanager contract.
type StakemanagerValidatorStakeSlashedIterator struct {
	Event *StakemanagerValidatorStakeSlashed // Event containing the contract specifics and raw log

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
func (it *StakemanagerValidatorStakeSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerValidatorStakeSlashed)
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
		it.Event = new(StakemanagerValidatorStakeSlashed)
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
func (it *StakemanagerValidatorStakeSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerValidatorStakeSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerValidatorStakeSlashed represents a ValidatorStakeSlashed event raised by the Stakemanager contract.
type StakemanagerValidatorStakeSlashed struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorStakeSlashed is a free log retrieval operation binding the contract event 0x2fd10b18da60ce9090915fbb72edd4d6550168ad1915bd58038d803b9faba2d7.
//
// Solidity: event ValidatorStakeSlashed(address indexed validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) FilterValidatorStakeSlashed(opts *bind.FilterOpts, validator []common.Address) (*StakemanagerValidatorStakeSlashedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ValidatorStakeSlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerValidatorStakeSlashedIterator{contract: _Stakemanager.contract, event: "ValidatorStakeSlashed", logs: logs, sub: sub}, nil
}

// WatchValidatorStakeSlashed is a free log subscription operation binding the contract event 0x2fd10b18da60ce9090915fbb72edd4d6550168ad1915bd58038d803b9faba2d7.
//
// Solidity: event ValidatorStakeSlashed(address indexed validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) WatchValidatorStakeSlashed(opts *bind.WatchOpts, sink chan<- *StakemanagerValidatorStakeSlashed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ValidatorStakeSlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerValidatorStakeSlashed)
				if err := _Stakemanager.contract.UnpackLog(event, "ValidatorStakeSlashed", log); err != nil {
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

// ParseValidatorStakeSlashed is a log parse operation binding the contract event 0x2fd10b18da60ce9090915fbb72edd4d6550168ad1915bd58038d803b9faba2d7.
//
// Solidity: event ValidatorStakeSlashed(address indexed validator, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) ParseValidatorStakeSlashed(log types.Log) (*StakemanagerValidatorStakeSlashed, error) {
	event := new(StakemanagerValidatorStakeSlashed)
	if err := _Stakemanager.contract.UnpackLog(event, "ValidatorStakeSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
