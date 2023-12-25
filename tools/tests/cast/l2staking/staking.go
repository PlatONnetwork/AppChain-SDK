// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2staking

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
	PubKey         hexutil.Bytes
	BlsKey         hexutil.Bytes
}

// StakingABI is the input ABI used to generate the binding from.
const StakingABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegateWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegateWithdrawalRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exitId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawalRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnDelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"UpdateValidatorStatus\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getDelegationsWithValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegateEpoch\",\"type\":\"uint256\"}],\"internalType\":\"structDelegationInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"periodType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"getValidatorAddrs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"start\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegateAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structValidatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"getValidatorsWithAddr\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegateAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structValidatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onStateReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"pendingWithdrawalsOfDelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"pendingWithdrawalsOfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"validatorIndexs\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signatues\",\"type\":\"bytes\"}],\"name\":\"verifyAggregateSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signatues\",\"type\":\"bytes\"}],\"name\":\"verifyAggregateSignatureByValidators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"withdrawableOfDelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawableOfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Staking is an auto generated Go binding around an platon contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an platon contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an platon contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an platon contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an platon contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// GetDelegationsWithValidator is a free data retrieval call binding the contract method 0xabec6a74.
//
// Solidity: function getDelegationsWithValidator(address[] validators, address delegator) view returns((address,address,uint256,uint256,uint256)[])
func (_Staking *StakingCaller) GetDelegationsWithValidator(opts *bind.CallOpts, validators []common.Address, delegator common.Address) ([]DelegationInfo, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getDelegationsWithValidator", validators, delegator)

	if err != nil {
		return *new([]DelegationInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]DelegationInfo)).(*[]DelegationInfo)

	return out0, err

}

// GetDelegationsWithValidator is a free data retrieval call binding the contract method 0xabec6a74.
//
// Solidity: function getDelegationsWithValidator(address[] validators, address delegator) view returns((address,address,uint256,uint256,uint256)[])
func (_Staking *StakingSession) GetDelegationsWithValidator(validators []common.Address, delegator common.Address) ([]DelegationInfo, error) {
	return _Staking.Contract.GetDelegationsWithValidator(&_Staking.CallOpts, validators, delegator)
}

// GetDelegationsWithValidator is a free data retrieval call binding the contract method 0xabec6a74.
//
// Solidity: function getDelegationsWithValidator(address[] validators, address delegator) view returns((address,address,uint256,uint256,uint256)[])
func (_Staking *StakingCallerSession) GetDelegationsWithValidator(validators []common.Address, delegator common.Address) ([]DelegationInfo, error) {
	return _Staking.Contract.GetDelegationsWithValidator(&_Staking.CallOpts, validators, delegator)
}

// GetValidatorAddrs is a free data retrieval call binding the contract method 0x23bc38f4.
//
// Solidity: function getValidatorAddrs(uint8 periodType, uint256 period) view returns(address[])
func (_Staking *StakingCaller) GetValidatorAddrs(opts *bind.CallOpts, periodType uint8, period *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorAddrs", periodType, period)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidatorAddrs is a free data retrieval call binding the contract method 0x23bc38f4.
//
// Solidity: function getValidatorAddrs(uint8 periodType, uint256 period) view returns(address[])
func (_Staking *StakingSession) GetValidatorAddrs(periodType uint8, period *big.Int) ([]common.Address, error) {
	return _Staking.Contract.GetValidatorAddrs(&_Staking.CallOpts, periodType, period)
}

// GetValidatorAddrs is a free data retrieval call binding the contract method 0x23bc38f4.
//
// Solidity: function getValidatorAddrs(uint8 periodType, uint256 period) view returns(address[])
func (_Staking *StakingCallerSession) GetValidatorAddrs(periodType uint8, period *big.Int) ([]common.Address, error) {
	return _Staking.Contract.GetValidatorAddrs(&_Staking.CallOpts, periodType, period)
}

// GetValidators is a free data retrieval call binding the contract method 0x229006c5.
//
// Solidity: function getValidators(bytes start, uint256 size) view returns(bytes, (address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[])
func (_Staking *StakingCaller) GetValidators(opts *bind.CallOpts, start []byte, size *big.Int) ([]byte, []ValidatorInfo, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidators", start, size)

	if err != nil {
		return *new([]byte), *new([]ValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]ValidatorInfo)).(*[]ValidatorInfo)

	return out0, out1, err

}

// GetValidators is a free data retrieval call binding the contract method 0x229006c5.
//
// Solidity: function getValidators(bytes start, uint256 size) view returns(bytes, (address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[])
func (_Staking *StakingSession) GetValidators(start []byte, size *big.Int) ([]byte, []ValidatorInfo, error) {
	return _Staking.Contract.GetValidators(&_Staking.CallOpts, start, size)
}

// GetValidators is a free data retrieval call binding the contract method 0x229006c5.
//
// Solidity: function getValidators(bytes start, uint256 size) view returns(bytes, (address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[])
func (_Staking *StakingCallerSession) GetValidators(start []byte, size *big.Int) ([]byte, []ValidatorInfo, error) {
	return _Staking.Contract.GetValidators(&_Staking.CallOpts, start, size)
}

// GetValidatorsWithAddr is a free data retrieval call binding the contract method 0x8f1956f6.
//
// Solidity: function getValidatorsWithAddr(address[] validators) view returns((address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[])
func (_Staking *StakingCaller) GetValidatorsWithAddr(opts *bind.CallOpts, validators []common.Address) ([]ValidatorInfo, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorsWithAddr", validators)

	if err != nil {
		return *new([]ValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]ValidatorInfo)).(*[]ValidatorInfo)

	return out0, err

}

// GetValidatorsWithAddr is a free data retrieval call binding the contract method 0x8f1956f6.
//
// Solidity: function getValidatorsWithAddr(address[] validators) view returns((address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[])
func (_Staking *StakingSession) GetValidatorsWithAddr(validators []common.Address) ([]ValidatorInfo, error) {
	return _Staking.Contract.GetValidatorsWithAddr(&_Staking.CallOpts, validators)
}

// GetValidatorsWithAddr is a free data retrieval call binding the contract method 0x8f1956f6.
//
// Solidity: function getValidatorsWithAddr(address[] validators) view returns((address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[])
func (_Staking *StakingCallerSession) GetValidatorsWithAddr(validators []common.Address) ([]ValidatorInfo, error) {
	return _Staking.Contract.GetValidatorsWithAddr(&_Staking.CallOpts, validators)
}

// PendingWithdrawalsOfDelegate is a free data retrieval call binding the contract method 0xb445f494.
//
// Solidity: function pendingWithdrawalsOfDelegate(address validator, address delegator) view returns(uint256)
func (_Staking *StakingCaller) PendingWithdrawalsOfDelegate(opts *bind.CallOpts, validator common.Address, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "pendingWithdrawalsOfDelegate", validator, delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingWithdrawalsOfDelegate is a free data retrieval call binding the contract method 0xb445f494.
//
// Solidity: function pendingWithdrawalsOfDelegate(address validator, address delegator) view returns(uint256)
func (_Staking *StakingSession) PendingWithdrawalsOfDelegate(validator common.Address, delegator common.Address) (*big.Int, error) {
	return _Staking.Contract.PendingWithdrawalsOfDelegate(&_Staking.CallOpts, validator, delegator)
}

// PendingWithdrawalsOfDelegate is a free data retrieval call binding the contract method 0xb445f494.
//
// Solidity: function pendingWithdrawalsOfDelegate(address validator, address delegator) view returns(uint256)
func (_Staking *StakingCallerSession) PendingWithdrawalsOfDelegate(validator common.Address, delegator common.Address) (*big.Int, error) {
	return _Staking.Contract.PendingWithdrawalsOfDelegate(&_Staking.CallOpts, validator, delegator)
}

// PendingWithdrawalsOfStake is a free data retrieval call binding the contract method 0x46d8c653.
//
// Solidity: function pendingWithdrawalsOfStake(address validator) view returns(uint256)
func (_Staking *StakingCaller) PendingWithdrawalsOfStake(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "pendingWithdrawalsOfStake", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingWithdrawalsOfStake is a free data retrieval call binding the contract method 0x46d8c653.
//
// Solidity: function pendingWithdrawalsOfStake(address validator) view returns(uint256)
func (_Staking *StakingSession) PendingWithdrawalsOfStake(validator common.Address) (*big.Int, error) {
	return _Staking.Contract.PendingWithdrawalsOfStake(&_Staking.CallOpts, validator)
}

// PendingWithdrawalsOfStake is a free data retrieval call binding the contract method 0x46d8c653.
//
// Solidity: function pendingWithdrawalsOfStake(address validator) view returns(uint256)
func (_Staking *StakingCallerSession) PendingWithdrawalsOfStake(validator common.Address) (*big.Int, error) {
	return _Staking.Contract.PendingWithdrawalsOfStake(&_Staking.CallOpts, validator)
}

// VerifyAggregateSignature is a free data retrieval call binding the contract method 0xcce6860b.
//
// Solidity: function verifyAggregateSignature(uint256 blockNumber, uint256[] validatorIndexs, bytes32 data, bytes signatues) view returns(bool)
func (_Staking *StakingCaller) VerifyAggregateSignature(opts *bind.CallOpts, blockNumber *big.Int, validatorIndexs []*big.Int, data [32]byte, signatues []byte) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "verifyAggregateSignature", blockNumber, validatorIndexs, data, signatues)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAggregateSignature is a free data retrieval call binding the contract method 0xcce6860b.
//
// Solidity: function verifyAggregateSignature(uint256 blockNumber, uint256[] validatorIndexs, bytes32 data, bytes signatues) view returns(bool)
func (_Staking *StakingSession) VerifyAggregateSignature(blockNumber *big.Int, validatorIndexs []*big.Int, data [32]byte, signatues []byte) (bool, error) {
	return _Staking.Contract.VerifyAggregateSignature(&_Staking.CallOpts, blockNumber, validatorIndexs, data, signatues)
}

// VerifyAggregateSignature is a free data retrieval call binding the contract method 0xcce6860b.
//
// Solidity: function verifyAggregateSignature(uint256 blockNumber, uint256[] validatorIndexs, bytes32 data, bytes signatues) view returns(bool)
func (_Staking *StakingCallerSession) VerifyAggregateSignature(blockNumber *big.Int, validatorIndexs []*big.Int, data [32]byte, signatues []byte) (bool, error) {
	return _Staking.Contract.VerifyAggregateSignature(&_Staking.CallOpts, blockNumber, validatorIndexs, data, signatues)
}

// VerifyAggregateSignatureByValidators is a free data retrieval call binding the contract method 0x34ccc389.
//
// Solidity: function verifyAggregateSignatureByValidators(address[] validators, bytes32 data, bytes signatues) view returns(bool)
func (_Staking *StakingCaller) VerifyAggregateSignatureByValidators(opts *bind.CallOpts, validators []common.Address, data [32]byte, signatues []byte) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "verifyAggregateSignatureByValidators", validators, data, signatues)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAggregateSignatureByValidators is a free data retrieval call binding the contract method 0x34ccc389.
//
// Solidity: function verifyAggregateSignatureByValidators(address[] validators, bytes32 data, bytes signatues) view returns(bool)
func (_Staking *StakingSession) VerifyAggregateSignatureByValidators(validators []common.Address, data [32]byte, signatues []byte) (bool, error) {
	return _Staking.Contract.VerifyAggregateSignatureByValidators(&_Staking.CallOpts, validators, data, signatues)
}

// VerifyAggregateSignatureByValidators is a free data retrieval call binding the contract method 0x34ccc389.
//
// Solidity: function verifyAggregateSignatureByValidators(address[] validators, bytes32 data, bytes signatues) view returns(bool)
func (_Staking *StakingCallerSession) VerifyAggregateSignatureByValidators(validators []common.Address, data [32]byte, signatues []byte) (bool, error) {
	return _Staking.Contract.VerifyAggregateSignatureByValidators(&_Staking.CallOpts, validators, data, signatues)
}

// WithdrawableOfDelegate is a free data retrieval call binding the contract method 0xa315ebb6.
//
// Solidity: function withdrawableOfDelegate(address validator, address delegator) view returns(uint256)
func (_Staking *StakingCaller) WithdrawableOfDelegate(opts *bind.CallOpts, validator common.Address, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "withdrawableOfDelegate", validator, delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableOfDelegate is a free data retrieval call binding the contract method 0xa315ebb6.
//
// Solidity: function withdrawableOfDelegate(address validator, address delegator) view returns(uint256)
func (_Staking *StakingSession) WithdrawableOfDelegate(validator common.Address, delegator common.Address) (*big.Int, error) {
	return _Staking.Contract.WithdrawableOfDelegate(&_Staking.CallOpts, validator, delegator)
}

// WithdrawableOfDelegate is a free data retrieval call binding the contract method 0xa315ebb6.
//
// Solidity: function withdrawableOfDelegate(address validator, address delegator) view returns(uint256)
func (_Staking *StakingCallerSession) WithdrawableOfDelegate(validator common.Address, delegator common.Address) (*big.Int, error) {
	return _Staking.Contract.WithdrawableOfDelegate(&_Staking.CallOpts, validator, delegator)
}

// WithdrawableOfStake is a free data retrieval call binding the contract method 0x68a11156.
//
// Solidity: function withdrawableOfStake(address validator) view returns(uint256)
func (_Staking *StakingCaller) WithdrawableOfStake(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "withdrawableOfStake", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableOfStake is a free data retrieval call binding the contract method 0x68a11156.
//
// Solidity: function withdrawableOfStake(address validator) view returns(uint256)
func (_Staking *StakingSession) WithdrawableOfStake(validator common.Address) (*big.Int, error) {
	return _Staking.Contract.WithdrawableOfStake(&_Staking.CallOpts, validator)
}

// WithdrawableOfStake is a free data retrieval call binding the contract method 0x68a11156.
//
// Solidity: function withdrawableOfStake(address validator) view returns(uint256)
func (_Staking *StakingCallerSession) WithdrawableOfStake(validator common.Address) (*big.Int, error) {
	return _Staking.Contract.WithdrawableOfStake(&_Staking.CallOpts, validator)
}

// OnStateReceive is a paid mutator transaction binding the contract method 0xeeb49945.
//
// Solidity: function onStateReceive(uint256 id, address sender, bytes data) returns()
func (_Staking *StakingTransactor) OnStateReceive(opts *bind.TransactOpts, id *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "onStateReceive", id, sender, data)
}

// OnStateReceive is a paid mutator transaction binding the contract method 0xeeb49945.
//
// Solidity: function onStateReceive(uint256 id, address sender, bytes data) returns()
func (_Staking *StakingSession) OnStateReceive(id *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _Staking.Contract.OnStateReceive(&_Staking.TransactOpts, id, sender, data)
}

// OnStateReceive is a paid mutator transaction binding the contract method 0xeeb49945.
//
// Solidity: function onStateReceive(uint256 id, address sender, bytes data) returns()
func (_Staking *StakingTransactorSession) OnStateReceive(id *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _Staking.Contract.OnStateReceive(&_Staking.TransactOpts, id, sender, data)
}

// Slash is a paid mutator transaction binding the contract method 0x2da25de3.
//
// Solidity: function slash() returns()
func (_Staking *StakingTransactor) Slash(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "slash")
}

// Slash is a paid mutator transaction binding the contract method 0x2da25de3.
//
// Solidity: function slash() returns()
func (_Staking *StakingSession) Slash() (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0x2da25de3.
//
// Solidity: function slash() returns()
func (_Staking *StakingTransactorSession) Slash() (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_Staking *StakingTransactor) Undelegate(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "undelegate", validator, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_Staking *StakingSession) Undelegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Undelegate(&_Staking.TransactOpts, validator, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_Staking *StakingTransactorSession) Undelegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Undelegate(&_Staking.TransactOpts, validator, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address validator, uint256 amount) returns()
func (_Staking *StakingTransactor) Unstake(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unstake", validator, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address validator, uint256 amount) returns()
func (_Staking *StakingSession) Unstake(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts, validator, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address validator, uint256 amount) returns()
func (_Staking *StakingTransactorSession) Unstake(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts, validator, amount)
}

// WithdrawUndelegate is a paid mutator transaction binding the contract method 0xb4065e75.
//
// Solidity: function withdrawUndelegate(address validator) returns()
func (_Staking *StakingTransactor) WithdrawUndelegate(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdrawUndelegate", validator)
}

// WithdrawUndelegate is a paid mutator transaction binding the contract method 0xb4065e75.
//
// Solidity: function withdrawUndelegate(address validator) returns()
func (_Staking *StakingSession) WithdrawUndelegate(validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.WithdrawUndelegate(&_Staking.TransactOpts, validator)
}

// WithdrawUndelegate is a paid mutator transaction binding the contract method 0xb4065e75.
//
// Solidity: function withdrawUndelegate(address validator) returns()
func (_Staking *StakingTransactorSession) WithdrawUndelegate(validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.WithdrawUndelegate(&_Staking.TransactOpts, validator)
}

// WithdrawUnstake is a paid mutator transaction binding the contract method 0xc76d485f.
//
// Solidity: function withdrawUnstake(address validator) returns()
func (_Staking *StakingTransactor) WithdrawUnstake(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdrawUnstake", validator)
}

// WithdrawUnstake is a paid mutator transaction binding the contract method 0xc76d485f.
//
// Solidity: function withdrawUnstake(address validator) returns()
func (_Staking *StakingSession) WithdrawUnstake(validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.WithdrawUnstake(&_Staking.TransactOpts, validator)
}

// WithdrawUnstake is a paid mutator transaction binding the contract method 0xc76d485f.
//
// Solidity: function withdrawUnstake(address validator) returns()
func (_Staking *StakingTransactorSession) WithdrawUnstake(validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.WithdrawUnstake(&_Staking.TransactOpts, validator)
}

// StakingDelegateWithdrawalIterator is returned from FilterDelegateWithdrawal and is used to iterate over the raw logs and unpacked data for DelegateWithdrawal events raised by the Staking contract.
type StakingDelegateWithdrawalIterator struct {
	Event *StakingDelegateWithdrawal // Event containing the contract specifics and raw log

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
func (it *StakingDelegateWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegateWithdrawal)
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
		it.Event = new(StakingDelegateWithdrawal)
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
func (it *StakingDelegateWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegateWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegateWithdrawal represents a DelegateWithdrawal event raised by the Staking contract.
type StakingDelegateWithdrawal struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegateWithdrawal is a free log retrieval operation binding the contract event 0x1fbf88541245ec027637253fd351d7237c1244a04eaa3352ca77e8aaff599a59.
//
// Solidity: event DelegateWithdrawal(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) FilterDelegateWithdrawal(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*StakingDelegateWithdrawalIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DelegateWithdrawal", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegateWithdrawalIterator{contract: _Staking.contract, event: "DelegateWithdrawal", logs: logs, sub: sub}, nil
}

// WatchDelegateWithdrawal is a free log subscription operation binding the contract event 0x1fbf88541245ec027637253fd351d7237c1244a04eaa3352ca77e8aaff599a59.
//
// Solidity: event DelegateWithdrawal(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) WatchDelegateWithdrawal(opts *bind.WatchOpts, sink chan<- *StakingDelegateWithdrawal, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DelegateWithdrawal", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegateWithdrawal)
				if err := _Staking.contract.UnpackLog(event, "DelegateWithdrawal", log); err != nil {
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

// ParseDelegateWithdrawal is a log parse operation binding the contract event 0x1fbf88541245ec027637253fd351d7237c1244a04eaa3352ca77e8aaff599a59.
//
// Solidity: event DelegateWithdrawal(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) ParseDelegateWithdrawal(log types.Log) (*StakingDelegateWithdrawal, error) {
	event := new(StakingDelegateWithdrawal)
	if err := _Staking.contract.UnpackLog(event, "DelegateWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDelegateWithdrawalRegisteredIterator is returned from FilterDelegateWithdrawalRegistered and is used to iterate over the raw logs and unpacked data for DelegateWithdrawalRegistered events raised by the Staking contract.
type StakingDelegateWithdrawalRegisteredIterator struct {
	Event *StakingDelegateWithdrawalRegistered // Event containing the contract specifics and raw log

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
func (it *StakingDelegateWithdrawalRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegateWithdrawalRegistered)
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
		it.Event = new(StakingDelegateWithdrawalRegistered)
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
func (it *StakingDelegateWithdrawalRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegateWithdrawalRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegateWithdrawalRegistered represents a DelegateWithdrawalRegistered event raised by the Staking contract.
type StakingDelegateWithdrawalRegistered struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegateWithdrawalRegistered is a free log retrieval operation binding the contract event 0xaf7d40b00b0df8607eafd5bfbe9a61372370c29c3a1d2004a917c52b0c14099c.
//
// Solidity: event DelegateWithdrawalRegistered(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) FilterDelegateWithdrawalRegistered(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*StakingDelegateWithdrawalRegisteredIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DelegateWithdrawalRegistered", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegateWithdrawalRegisteredIterator{contract: _Staking.contract, event: "DelegateWithdrawalRegistered", logs: logs, sub: sub}, nil
}

// WatchDelegateWithdrawalRegistered is a free log subscription operation binding the contract event 0xaf7d40b00b0df8607eafd5bfbe9a61372370c29c3a1d2004a917c52b0c14099c.
//
// Solidity: event DelegateWithdrawalRegistered(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) WatchDelegateWithdrawalRegistered(opts *bind.WatchOpts, sink chan<- *StakingDelegateWithdrawalRegistered, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DelegateWithdrawalRegistered", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegateWithdrawalRegistered)
				if err := _Staking.contract.UnpackLog(event, "DelegateWithdrawalRegistered", log); err != nil {
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

// ParseDelegateWithdrawalRegistered is a log parse operation binding the contract event 0xaf7d40b00b0df8607eafd5bfbe9a61372370c29c3a1d2004a917c52b0c14099c.
//
// Solidity: event DelegateWithdrawalRegistered(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) ParseDelegateWithdrawalRegistered(log types.Log) (*StakingDelegateWithdrawalRegistered, error) {
	event := new(StakingDelegateWithdrawalRegistered)
	if err := _Staking.contract.UnpackLog(event, "DelegateWithdrawalRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDelegationAddedIterator is returned from FilterDelegationAdded and is used to iterate over the raw logs and unpacked data for DelegationAdded events raised by the Staking contract.
type StakingDelegationAddedIterator struct {
	Event *StakingDelegationAdded // Event containing the contract specifics and raw log

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
func (it *StakingDelegationAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegationAdded)
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
		it.Event = new(StakingDelegationAdded)
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
func (it *StakingDelegationAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegationAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegationAdded represents a DelegationAdded event raised by the Staking contract.
type StakingDelegationAdded struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegationAdded is a free log retrieval operation binding the contract event 0x52467f14b857734001c77e6f125dac41b45798837c9fc9adfe3a5b394c77a0e9.
//
// Solidity: event DelegationAdded(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) FilterDelegationAdded(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*StakingDelegationAddedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DelegationAdded", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegationAddedIterator{contract: _Staking.contract, event: "DelegationAdded", logs: logs, sub: sub}, nil
}

// WatchDelegationAdded is a free log subscription operation binding the contract event 0x52467f14b857734001c77e6f125dac41b45798837c9fc9adfe3a5b394c77a0e9.
//
// Solidity: event DelegationAdded(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) WatchDelegationAdded(opts *bind.WatchOpts, sink chan<- *StakingDelegationAdded, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DelegationAdded", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegationAdded)
				if err := _Staking.contract.UnpackLog(event, "DelegationAdded", log); err != nil {
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
// Solidity: event DelegationAdded(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) ParseDelegationAdded(log types.Log) (*StakingDelegationAdded, error) {
	event := new(StakingDelegationAdded)
	if err := _Staking.contract.UnpackLog(event, "DelegationAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the Staking contract.
type StakingSlashedIterator struct {
	Event *StakingSlashed // Event containing the contract specifics and raw log

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
func (it *StakingSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSlashed)
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
		it.Event = new(StakingSlashed)
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
func (it *StakingSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSlashed represents a Slashed event raised by the Staking contract.
type StakingSlashed struct {
	ExitId     *big.Int
	Validators []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0xd2f2b50d0c108d01a95cfb6ee87668e30a20c08be7facf9f28146548f82a8ab7.
//
// Solidity: event Slashed(uint256 indexed exitId, address[] validators, uint256[] amounts)
func (_Staking *StakingFilterer) FilterSlashed(opts *bind.FilterOpts, exitId []*big.Int) (*StakingSlashedIterator, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Slashed", exitIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingSlashedIterator{contract: _Staking.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0xd2f2b50d0c108d01a95cfb6ee87668e30a20c08be7facf9f28146548f82a8ab7.
//
// Solidity: event Slashed(uint256 indexed exitId, address[] validators, uint256[] amounts)
func (_Staking *StakingFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *StakingSlashed, exitId []*big.Int) (event.Subscription, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Slashed", exitIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSlashed)
				if err := _Staking.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0xd2f2b50d0c108d01a95cfb6ee87668e30a20c08be7facf9f28146548f82a8ab7.
//
// Solidity: event Slashed(uint256 indexed exitId, address[] validators, uint256[] amounts)
func (_Staking *StakingFilterer) ParseSlashed(log types.Log) (*StakingSlashed, error) {
	event := new(StakingSlashed)
	if err := _Staking.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeAddedIterator is returned from FilterStakeAdded and is used to iterate over the raw logs and unpacked data for StakeAdded events raised by the Staking contract.
type StakingStakeAddedIterator struct {
	Event *StakingStakeAdded // Event containing the contract specifics and raw log

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
func (it *StakingStakeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeAdded)
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
		it.Event = new(StakingStakeAdded)
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
func (it *StakingStakeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeAdded represents a StakeAdded event raised by the Staking contract.
type StakingStakeAdded struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeAdded is a free log retrieval operation binding the contract event 0x7c717985ac273e663b7f3050f5b15a4388ff6ed952338954f650e2093e13937f.
//
// Solidity: event StakeAdded(address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) FilterStakeAdded(opts *bind.FilterOpts, validator []common.Address) (*StakingStakeAddedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakeAddedIterator{contract: _Staking.contract, event: "StakeAdded", logs: logs, sub: sub}, nil
}

// WatchStakeAdded is a free log subscription operation binding the contract event 0x7c717985ac273e663b7f3050f5b15a4388ff6ed952338954f650e2093e13937f.
//
// Solidity: event StakeAdded(address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) WatchStakeAdded(opts *bind.WatchOpts, sink chan<- *StakingStakeAdded, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeAdded)
				if err := _Staking.contract.UnpackLog(event, "StakeAdded", log); err != nil {
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
func (_Staking *StakingFilterer) ParseStakeAdded(log types.Log) (*StakingStakeAdded, error) {
	event := new(StakingStakeAdded)
	if err := _Staking.contract.UnpackLog(event, "StakeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeWithdrawalIterator is returned from FilterStakeWithdrawal and is used to iterate over the raw logs and unpacked data for StakeWithdrawal events raised by the Staking contract.
type StakingStakeWithdrawalIterator struct {
	Event *StakingStakeWithdrawal // Event containing the contract specifics and raw log

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
func (it *StakingStakeWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeWithdrawal)
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
		it.Event = new(StakingStakeWithdrawal)
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
func (it *StakingStakeWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeWithdrawal represents a StakeWithdrawal event raised by the Staking contract.
type StakingStakeWithdrawal struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawal is a free log retrieval operation binding the contract event 0xf0ed97f7b968f9d8268bc8d104a11b3586ceeadd0e0af5f73769e2b479f9d0ae.
//
// Solidity: event StakeWithdrawal(address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) FilterStakeWithdrawal(opts *bind.FilterOpts, validator []common.Address) (*StakingStakeWithdrawalIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeWithdrawal", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakeWithdrawalIterator{contract: _Staking.contract, event: "StakeWithdrawal", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawal is a free log subscription operation binding the contract event 0xf0ed97f7b968f9d8268bc8d104a11b3586ceeadd0e0af5f73769e2b479f9d0ae.
//
// Solidity: event StakeWithdrawal(address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) WatchStakeWithdrawal(opts *bind.WatchOpts, sink chan<- *StakingStakeWithdrawal, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeWithdrawal", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeWithdrawal)
				if err := _Staking.contract.UnpackLog(event, "StakeWithdrawal", log); err != nil {
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

// ParseStakeWithdrawal is a log parse operation binding the contract event 0xf0ed97f7b968f9d8268bc8d104a11b3586ceeadd0e0af5f73769e2b479f9d0ae.
//
// Solidity: event StakeWithdrawal(address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) ParseStakeWithdrawal(log types.Log) (*StakingStakeWithdrawal, error) {
	event := new(StakingStakeWithdrawal)
	if err := _Staking.contract.UnpackLog(event, "StakeWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeWithdrawalRegisteredIterator is returned from FilterStakeWithdrawalRegistered and is used to iterate over the raw logs and unpacked data for StakeWithdrawalRegistered events raised by the Staking contract.
type StakingStakeWithdrawalRegisteredIterator struct {
	Event *StakingStakeWithdrawalRegistered // Event containing the contract specifics and raw log

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
func (it *StakingStakeWithdrawalRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeWithdrawalRegistered)
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
		it.Event = new(StakingStakeWithdrawalRegistered)
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
func (it *StakingStakeWithdrawalRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeWithdrawalRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeWithdrawalRegistered represents a StakeWithdrawalRegistered event raised by the Staking contract.
type StakingStakeWithdrawalRegistered struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawalRegistered is a free log retrieval operation binding the contract event 0x53fd52fe077bc27429744caa52a6d5476c5608a0c3c99bec17b770d3705dc7d0.
//
// Solidity: event StakeWithdrawalRegistered(address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) FilterStakeWithdrawalRegistered(opts *bind.FilterOpts, validator []common.Address) (*StakingStakeWithdrawalRegisteredIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeWithdrawalRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakeWithdrawalRegisteredIterator{contract: _Staking.contract, event: "StakeWithdrawalRegistered", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawalRegistered is a free log subscription operation binding the contract event 0x53fd52fe077bc27429744caa52a6d5476c5608a0c3c99bec17b770d3705dc7d0.
//
// Solidity: event StakeWithdrawalRegistered(address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) WatchStakeWithdrawalRegistered(opts *bind.WatchOpts, sink chan<- *StakingStakeWithdrawalRegistered, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeWithdrawalRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeWithdrawalRegistered)
				if err := _Staking.contract.UnpackLog(event, "StakeWithdrawalRegistered", log); err != nil {
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

// ParseStakeWithdrawalRegistered is a log parse operation binding the contract event 0x53fd52fe077bc27429744caa52a6d5476c5608a0c3c99bec17b770d3705dc7d0.
//
// Solidity: event StakeWithdrawalRegistered(address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) ParseStakeWithdrawalRegistered(log types.Log) (*StakingStakeWithdrawalRegistered, error) {
	event := new(StakingStakeWithdrawalRegistered)
	if err := _Staking.contract.UnpackLog(event, "StakeWithdrawalRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnDelegatedIterator is returned from FilterUnDelegated and is used to iterate over the raw logs and unpacked data for UnDelegated events raised by the Staking contract.
type StakingUnDelegatedIterator struct {
	Event *StakingUnDelegated // Event containing the contract specifics and raw log

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
func (it *StakingUnDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnDelegated)
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
		it.Event = new(StakingUnDelegated)
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
func (it *StakingUnDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnDelegated represents a UnDelegated event raised by the Staking contract.
type StakingUnDelegated struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnDelegated is a free log retrieval operation binding the contract event 0x33f37c4c8173c3f236d2b74f93b425280ea2f7f2924a98ec73744fbc04f9ee35.
//
// Solidity: event UnDelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) FilterUnDelegated(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*StakingUnDelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "UnDelegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingUnDelegatedIterator{contract: _Staking.contract, event: "UnDelegated", logs: logs, sub: sub}, nil
}

// WatchUnDelegated is a free log subscription operation binding the contract event 0x33f37c4c8173c3f236d2b74f93b425280ea2f7f2924a98ec73744fbc04f9ee35.
//
// Solidity: event UnDelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) WatchUnDelegated(opts *bind.WatchOpts, sink chan<- *StakingUnDelegated, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "UnDelegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnDelegated)
				if err := _Staking.contract.UnpackLog(event, "UnDelegated", log); err != nil {
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

// ParseUnDelegated is a log parse operation binding the contract event 0x33f37c4c8173c3f236d2b74f93b425280ea2f7f2924a98ec73744fbc04f9ee35.
//
// Solidity: event UnDelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) ParseUnDelegated(log types.Log) (*StakingUnDelegated, error) {
	event := new(StakingUnDelegated)
	if err := _Staking.contract.UnpackLog(event, "UnDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnStakedIterator is returned from FilterUnStaked and is used to iterate over the raw logs and unpacked data for UnStaked events raised by the Staking contract.
type StakingUnStakedIterator struct {
	Event *StakingUnStaked // Event containing the contract specifics and raw log

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
func (it *StakingUnStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnStaked)
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
		it.Event = new(StakingUnStaked)
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
func (it *StakingUnStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnStaked represents a UnStaked event raised by the Staking contract.
type StakingUnStaked struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnStaked is a free log retrieval operation binding the contract event 0x79d3df6837cc49ff0e09fd3258e6e45594e0703445bb06825e9d75156eaee8f0.
//
// Solidity: event UnStaked(address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) FilterUnStaked(opts *bind.FilterOpts, validator []common.Address) (*StakingUnStakedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "UnStaked", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingUnStakedIterator{contract: _Staking.contract, event: "UnStaked", logs: logs, sub: sub}, nil
}

// WatchUnStaked is a free log subscription operation binding the contract event 0x79d3df6837cc49ff0e09fd3258e6e45594e0703445bb06825e9d75156eaee8f0.
//
// Solidity: event UnStaked(address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) WatchUnStaked(opts *bind.WatchOpts, sink chan<- *StakingUnStaked, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "UnStaked", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnStaked)
				if err := _Staking.contract.UnpackLog(event, "UnStaked", log); err != nil {
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

// ParseUnStaked is a log parse operation binding the contract event 0x79d3df6837cc49ff0e09fd3258e6e45594e0703445bb06825e9d75156eaee8f0.
//
// Solidity: event UnStaked(address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) ParseUnStaked(log types.Log) (*StakingUnStaked, error) {
	event := new(StakingUnStaked)
	if err := _Staking.contract.UnpackLog(event, "UnStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUpdateValidatorStatusIterator is returned from FilterUpdateValidatorStatus and is used to iterate over the raw logs and unpacked data for UpdateValidatorStatus events raised by the Staking contract.
type StakingUpdateValidatorStatusIterator struct {
	Event *StakingUpdateValidatorStatus // Event containing the contract specifics and raw log

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
func (it *StakingUpdateValidatorStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUpdateValidatorStatus)
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
		it.Event = new(StakingUpdateValidatorStatus)
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
func (it *StakingUpdateValidatorStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUpdateValidatorStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUpdateValidatorStatus represents a UpdateValidatorStatus event raised by the Staking contract.
type StakingUpdateValidatorStatus struct {
	Validator common.Address
	Status    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateValidatorStatus is a free log retrieval operation binding the contract event 0x85ff997a3e90354ca8883205ac49293eed56e342aa01c0223bd70027118943c2.
//
// Solidity: event UpdateValidatorStatus(address indexed validator, uint256 status)
func (_Staking *StakingFilterer) FilterUpdateValidatorStatus(opts *bind.FilterOpts, validator []common.Address) (*StakingUpdateValidatorStatusIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "UpdateValidatorStatus", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingUpdateValidatorStatusIterator{contract: _Staking.contract, event: "UpdateValidatorStatus", logs: logs, sub: sub}, nil
}

// WatchUpdateValidatorStatus is a free log subscription operation binding the contract event 0x85ff997a3e90354ca8883205ac49293eed56e342aa01c0223bd70027118943c2.
//
// Solidity: event UpdateValidatorStatus(address indexed validator, uint256 status)
func (_Staking *StakingFilterer) WatchUpdateValidatorStatus(opts *bind.WatchOpts, sink chan<- *StakingUpdateValidatorStatus, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "UpdateValidatorStatus", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUpdateValidatorStatus)
				if err := _Staking.contract.UnpackLog(event, "UpdateValidatorStatus", log); err != nil {
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

// ParseUpdateValidatorStatus is a log parse operation binding the contract event 0x85ff997a3e90354ca8883205ac49293eed56e342aa01c0223bd70027118943c2.
//
// Solidity: event UpdateValidatorStatus(address indexed validator, uint256 status)
func (_Staking *StakingFilterer) ParseUpdateValidatorStatus(log types.Log) (*StakingUpdateValidatorStatus, error) {
	event := new(StakingUpdateValidatorStatus)
	if err := _Staking.contract.UnpackLog(event, "UpdateValidatorStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
