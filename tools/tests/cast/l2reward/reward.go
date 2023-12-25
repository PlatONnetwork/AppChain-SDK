// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2reward

import (
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

// RewardABI is the input ABI used to generate the binding from.
const RewardABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"BlockReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"DelegatorRewardWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"EpochReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"}],\"name\":\"RewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"ValidatorRewardWithdrawal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"paidRewardPerEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"pendingDelegatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"pendingValidatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawDelegatorReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawValidatorReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Reward is an auto generated Go binding around an platon contract.
type Reward struct {
	RewardCaller     // Read-only binding to the contract
	RewardTransactor // Write-only binding to the contract
	RewardFilterer   // Log filterer for contract events
}

// RewardCaller is an auto generated read-only Go binding around an platon contract.
type RewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardTransactor is an auto generated write-only Go binding around an platon contract.
type RewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardFilterer is an auto generated log filtering Go binding around an platon contract events.
type RewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type RewardSession struct {
	Contract     *Reward           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type RewardCallerSession struct {
	Contract *RewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RewardTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type RewardTransactorSession struct {
	Contract     *RewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardRaw is an auto generated low-level Go binding around an platon contract.
type RewardRaw struct {
	Contract *Reward // Generic contract binding to access the raw methods on
}

// RewardCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type RewardCallerRaw struct {
	Contract *RewardCaller // Generic read-only contract binding to access the raw methods on
}

// RewardTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type RewardTransactorRaw struct {
	Contract *RewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReward creates a new instance of Reward, bound to a specific deployed contract.
func NewReward(address common.Address, backend bind.ContractBackend) (*Reward, error) {
	contract, err := bindReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reward{RewardCaller: RewardCaller{contract: contract}, RewardTransactor: RewardTransactor{contract: contract}, RewardFilterer: RewardFilterer{contract: contract}}, nil
}

// NewRewardCaller creates a new read-only instance of Reward, bound to a specific deployed contract.
func NewRewardCaller(address common.Address, caller bind.ContractCaller) (*RewardCaller, error) {
	contract, err := bindReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardCaller{contract: contract}, nil
}

// NewRewardTransactor creates a new write-only instance of Reward, bound to a specific deployed contract.
func NewRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardTransactor, error) {
	contract, err := bindReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardTransactor{contract: contract}, nil
}

// NewRewardFilterer creates a new log filterer instance of Reward, bound to a specific deployed contract.
func NewRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardFilterer, error) {
	contract, err := bindReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardFilterer{contract: contract}, nil
}

// bindReward binds a generic wrapper to an already deployed contract.
func bindReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reward *RewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reward.Contract.RewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reward *RewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reward.Contract.RewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reward *RewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reward.Contract.RewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reward *RewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reward *RewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reward *RewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reward.Contract.contract.Transact(opts, method, params...)
}

// PaidRewardPerEpoch is a free data retrieval call binding the contract method 0x07358b99.
//
// Solidity: function paidRewardPerEpoch(uint256 epochId) view returns(uint256)
func (_Reward *RewardCaller) PaidRewardPerEpoch(opts *bind.CallOpts, epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Reward.contract.Call(opts, &out, "paidRewardPerEpoch", epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PaidRewardPerEpoch is a free data retrieval call binding the contract method 0x07358b99.
//
// Solidity: function paidRewardPerEpoch(uint256 epochId) view returns(uint256)
func (_Reward *RewardSession) PaidRewardPerEpoch(epochId *big.Int) (*big.Int, error) {
	return _Reward.Contract.PaidRewardPerEpoch(&_Reward.CallOpts, epochId)
}

// PaidRewardPerEpoch is a free data retrieval call binding the contract method 0x07358b99.
//
// Solidity: function paidRewardPerEpoch(uint256 epochId) view returns(uint256)
func (_Reward *RewardCallerSession) PaidRewardPerEpoch(epochId *big.Int) (*big.Int, error) {
	return _Reward.Contract.PaidRewardPerEpoch(&_Reward.CallOpts, epochId)
}

// PendingDelegatorRewards is a free data retrieval call binding the contract method 0x08abda72.
//
// Solidity: function pendingDelegatorRewards(address validator) view returns(uint256)
func (_Reward *RewardCaller) PendingDelegatorRewards(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Reward.contract.Call(opts, &out, "pendingDelegatorRewards", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingDelegatorRewards is a free data retrieval call binding the contract method 0x08abda72.
//
// Solidity: function pendingDelegatorRewards(address validator) view returns(uint256)
func (_Reward *RewardSession) PendingDelegatorRewards(validator common.Address) (*big.Int, error) {
	return _Reward.Contract.PendingDelegatorRewards(&_Reward.CallOpts, validator)
}

// PendingDelegatorRewards is a free data retrieval call binding the contract method 0x08abda72.
//
// Solidity: function pendingDelegatorRewards(address validator) view returns(uint256)
func (_Reward *RewardCallerSession) PendingDelegatorRewards(validator common.Address) (*big.Int, error) {
	return _Reward.Contract.PendingDelegatorRewards(&_Reward.CallOpts, validator)
}

// PendingValidatorRewards is a free data retrieval call binding the contract method 0xa617627c.
//
// Solidity: function pendingValidatorRewards(address validator) view returns(uint256)
func (_Reward *RewardCaller) PendingValidatorRewards(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Reward.contract.Call(opts, &out, "pendingValidatorRewards", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingValidatorRewards is a free data retrieval call binding the contract method 0xa617627c.
//
// Solidity: function pendingValidatorRewards(address validator) view returns(uint256)
func (_Reward *RewardSession) PendingValidatorRewards(validator common.Address) (*big.Int, error) {
	return _Reward.Contract.PendingValidatorRewards(&_Reward.CallOpts, validator)
}

// PendingValidatorRewards is a free data retrieval call binding the contract method 0xa617627c.
//
// Solidity: function pendingValidatorRewards(address validator) view returns(uint256)
func (_Reward *RewardCallerSession) PendingValidatorRewards(validator common.Address) (*big.Int, error) {
	return _Reward.Contract.PendingValidatorRewards(&_Reward.CallOpts, validator)
}

// WithdrawDelegatorReward is a paid mutator transaction binding the contract method 0x346683f0.
//
// Solidity: function withdrawDelegatorReward(address validator) returns()
func (_Reward *RewardTransactor) WithdrawDelegatorReward(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "withdrawDelegatorReward", validator)
}

// WithdrawDelegatorReward is a paid mutator transaction binding the contract method 0x346683f0.
//
// Solidity: function withdrawDelegatorReward(address validator) returns()
func (_Reward *RewardSession) WithdrawDelegatorReward(validator common.Address) (*types.Transaction, error) {
	return _Reward.Contract.WithdrawDelegatorReward(&_Reward.TransactOpts, validator)
}

// WithdrawDelegatorReward is a paid mutator transaction binding the contract method 0x346683f0.
//
// Solidity: function withdrawDelegatorReward(address validator) returns()
func (_Reward *RewardTransactorSession) WithdrawDelegatorReward(validator common.Address) (*types.Transaction, error) {
	return _Reward.Contract.WithdrawDelegatorReward(&_Reward.TransactOpts, validator)
}

// WithdrawValidatorReward is a paid mutator transaction binding the contract method 0x5bc0f928.
//
// Solidity: function withdrawValidatorReward(address validator) returns()
func (_Reward *RewardTransactor) WithdrawValidatorReward(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Reward.contract.Transact(opts, "withdrawValidatorReward", validator)
}

// WithdrawValidatorReward is a paid mutator transaction binding the contract method 0x5bc0f928.
//
// Solidity: function withdrawValidatorReward(address validator) returns()
func (_Reward *RewardSession) WithdrawValidatorReward(validator common.Address) (*types.Transaction, error) {
	return _Reward.Contract.WithdrawValidatorReward(&_Reward.TransactOpts, validator)
}

// WithdrawValidatorReward is a paid mutator transaction binding the contract method 0x5bc0f928.
//
// Solidity: function withdrawValidatorReward(address validator) returns()
func (_Reward *RewardTransactorSession) WithdrawValidatorReward(validator common.Address) (*types.Transaction, error) {
	return _Reward.Contract.WithdrawValidatorReward(&_Reward.TransactOpts, validator)
}

// RewardBlockRewardIterator is returned from FilterBlockReward and is used to iterate over the raw logs and unpacked data for BlockReward events raised by the Reward contract.
type RewardBlockRewardIterator struct {
	Event *RewardBlockReward // Event containing the contract specifics and raw log

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
func (it *RewardBlockRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardBlockReward)
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
		it.Event = new(RewardBlockReward)
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
func (it *RewardBlockRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardBlockRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardBlockReward represents a BlockReward event raised by the Reward contract.
type RewardBlockReward struct {
	EpochId    *big.Int
	Validators []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBlockReward is a free log retrieval operation binding the contract event 0xd603e2579e24fac887272c2477c3a04b78380676222de5ddd2da918124e7a878.
//
// Solidity: event BlockReward(uint256 indexed epochId, address[] validators, uint256[] amounts)
func (_Reward *RewardFilterer) FilterBlockReward(opts *bind.FilterOpts, epochId []*big.Int) (*RewardBlockRewardIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Reward.contract.FilterLogs(opts, "BlockReward", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &RewardBlockRewardIterator{contract: _Reward.contract, event: "BlockReward", logs: logs, sub: sub}, nil
}

// WatchBlockReward is a free log subscription operation binding the contract event 0xd603e2579e24fac887272c2477c3a04b78380676222de5ddd2da918124e7a878.
//
// Solidity: event BlockReward(uint256 indexed epochId, address[] validators, uint256[] amounts)
func (_Reward *RewardFilterer) WatchBlockReward(opts *bind.WatchOpts, sink chan<- *RewardBlockReward, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Reward.contract.WatchLogs(opts, "BlockReward", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardBlockReward)
				if err := _Reward.contract.UnpackLog(event, "BlockReward", log); err != nil {
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

// ParseBlockReward is a log parse operation binding the contract event 0xd603e2579e24fac887272c2477c3a04b78380676222de5ddd2da918124e7a878.
//
// Solidity: event BlockReward(uint256 indexed epochId, address[] validators, uint256[] amounts)
func (_Reward *RewardFilterer) ParseBlockReward(log types.Log) (*RewardBlockReward, error) {
	event := new(RewardBlockReward)
	if err := _Reward.contract.UnpackLog(event, "BlockReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardDelegatorRewardWithdrawalIterator is returned from FilterDelegatorRewardWithdrawal and is used to iterate over the raw logs and unpacked data for DelegatorRewardWithdrawal events raised by the Reward contract.
type RewardDelegatorRewardWithdrawalIterator struct {
	Event *RewardDelegatorRewardWithdrawal // Event containing the contract specifics and raw log

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
func (it *RewardDelegatorRewardWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardDelegatorRewardWithdrawal)
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
		it.Event = new(RewardDelegatorRewardWithdrawal)
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
func (it *RewardDelegatorRewardWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardDelegatorRewardWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardDelegatorRewardWithdrawal represents a DelegatorRewardWithdrawal event raised by the Reward contract.
type RewardDelegatorRewardWithdrawal struct {
	Validator common.Address
	Amount    *big.Int
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRewardWithdrawal is a free log retrieval operation binding the contract event 0xa03fb819557c2f18e46480598e54ee3f5fa40811c2019a51c56647ad8636d782.
//
// Solidity: event DelegatorRewardWithdrawal(address indexed validator, uint256 amount, address caller)
func (_Reward *RewardFilterer) FilterDelegatorRewardWithdrawal(opts *bind.FilterOpts, validator []common.Address) (*RewardDelegatorRewardWithdrawalIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Reward.contract.FilterLogs(opts, "DelegatorRewardWithdrawal", validatorRule)
	if err != nil {
		return nil, err
	}
	return &RewardDelegatorRewardWithdrawalIterator{contract: _Reward.contract, event: "DelegatorRewardWithdrawal", logs: logs, sub: sub}, nil
}

// WatchDelegatorRewardWithdrawal is a free log subscription operation binding the contract event 0xa03fb819557c2f18e46480598e54ee3f5fa40811c2019a51c56647ad8636d782.
//
// Solidity: event DelegatorRewardWithdrawal(address indexed validator, uint256 amount, address caller)
func (_Reward *RewardFilterer) WatchDelegatorRewardWithdrawal(opts *bind.WatchOpts, sink chan<- *RewardDelegatorRewardWithdrawal, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Reward.contract.WatchLogs(opts, "DelegatorRewardWithdrawal", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardDelegatorRewardWithdrawal)
				if err := _Reward.contract.UnpackLog(event, "DelegatorRewardWithdrawal", log); err != nil {
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

// ParseDelegatorRewardWithdrawal is a log parse operation binding the contract event 0xa03fb819557c2f18e46480598e54ee3f5fa40811c2019a51c56647ad8636d782.
//
// Solidity: event DelegatorRewardWithdrawal(address indexed validator, uint256 amount, address caller)
func (_Reward *RewardFilterer) ParseDelegatorRewardWithdrawal(log types.Log) (*RewardDelegatorRewardWithdrawal, error) {
	event := new(RewardDelegatorRewardWithdrawal)
	if err := _Reward.contract.UnpackLog(event, "DelegatorRewardWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardEpochRewardIterator is returned from FilterEpochReward and is used to iterate over the raw logs and unpacked data for EpochReward events raised by the Reward contract.
type RewardEpochRewardIterator struct {
	Event *RewardEpochReward // Event containing the contract specifics and raw log

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
func (it *RewardEpochRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardEpochReward)
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
		it.Event = new(RewardEpochReward)
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
func (it *RewardEpochRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardEpochRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardEpochReward represents a EpochReward event raised by the Reward contract.
type RewardEpochReward struct {
	EpochId    *big.Int
	Validators []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEpochReward is a free log retrieval operation binding the contract event 0x9d3a37db5471c2d283c8d98edcafe0f296aef6faed15859feb33a9a012490649.
//
// Solidity: event EpochReward(uint256 indexed epochId, address[] validators, uint256[] amounts)
func (_Reward *RewardFilterer) FilterEpochReward(opts *bind.FilterOpts, epochId []*big.Int) (*RewardEpochRewardIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Reward.contract.FilterLogs(opts, "EpochReward", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &RewardEpochRewardIterator{contract: _Reward.contract, event: "EpochReward", logs: logs, sub: sub}, nil
}

// WatchEpochReward is a free log subscription operation binding the contract event 0x9d3a37db5471c2d283c8d98edcafe0f296aef6faed15859feb33a9a012490649.
//
// Solidity: event EpochReward(uint256 indexed epochId, address[] validators, uint256[] amounts)
func (_Reward *RewardFilterer) WatchEpochReward(opts *bind.WatchOpts, sink chan<- *RewardEpochReward, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Reward.contract.WatchLogs(opts, "EpochReward", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardEpochReward)
				if err := _Reward.contract.UnpackLog(event, "EpochReward", log); err != nil {
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

// ParseEpochReward is a log parse operation binding the contract event 0x9d3a37db5471c2d283c8d98edcafe0f296aef6faed15859feb33a9a012490649.
//
// Solidity: event EpochReward(uint256 indexed epochId, address[] validators, uint256[] amounts)
func (_Reward *RewardFilterer) ParseEpochReward(log types.Log) (*RewardEpochReward, error) {
	event := new(RewardEpochReward)
	if err := _Reward.contract.UnpackLog(event, "EpochReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardRewardDistributedIterator is returned from FilterRewardDistributed and is used to iterate over the raw logs and unpacked data for RewardDistributed events raised by the Reward contract.
type RewardRewardDistributedIterator struct {
	Event *RewardRewardDistributed // Event containing the contract specifics and raw log

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
func (it *RewardRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardRewardDistributed)
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
		it.Event = new(RewardRewardDistributed)
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
func (it *RewardRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardRewardDistributed represents a RewardDistributed event raised by the Reward contract.
type RewardRewardDistributed struct {
	EpochId     *big.Int
	TotalReward *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardDistributed is a free log retrieval operation binding the contract event 0xeaf3d57629d9b1ce95715ccd98d6f5bf48023be1d5a06e09f64ab7f6d8be01d5.
//
// Solidity: event RewardDistributed(uint256 indexed epochId, uint256 totalReward)
func (_Reward *RewardFilterer) FilterRewardDistributed(opts *bind.FilterOpts, epochId []*big.Int) (*RewardRewardDistributedIterator, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Reward.contract.FilterLogs(opts, "RewardDistributed", epochIdRule)
	if err != nil {
		return nil, err
	}
	return &RewardRewardDistributedIterator{contract: _Reward.contract, event: "RewardDistributed", logs: logs, sub: sub}, nil
}

// WatchRewardDistributed is a free log subscription operation binding the contract event 0xeaf3d57629d9b1ce95715ccd98d6f5bf48023be1d5a06e09f64ab7f6d8be01d5.
//
// Solidity: event RewardDistributed(uint256 indexed epochId, uint256 totalReward)
func (_Reward *RewardFilterer) WatchRewardDistributed(opts *bind.WatchOpts, sink chan<- *RewardRewardDistributed, epochId []*big.Int) (event.Subscription, error) {

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Reward.contract.WatchLogs(opts, "RewardDistributed", epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardRewardDistributed)
				if err := _Reward.contract.UnpackLog(event, "RewardDistributed", log); err != nil {
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

// ParseRewardDistributed is a log parse operation binding the contract event 0xeaf3d57629d9b1ce95715ccd98d6f5bf48023be1d5a06e09f64ab7f6d8be01d5.
//
// Solidity: event RewardDistributed(uint256 indexed epochId, uint256 totalReward)
func (_Reward *RewardFilterer) ParseRewardDistributed(log types.Log) (*RewardRewardDistributed, error) {
	event := new(RewardRewardDistributed)
	if err := _Reward.contract.UnpackLog(event, "RewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardValidatorRewardWithdrawalIterator is returned from FilterValidatorRewardWithdrawal and is used to iterate over the raw logs and unpacked data for ValidatorRewardWithdrawal events raised by the Reward contract.
type RewardValidatorRewardWithdrawalIterator struct {
	Event *RewardValidatorRewardWithdrawal // Event containing the contract specifics and raw log

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
func (it *RewardValidatorRewardWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardValidatorRewardWithdrawal)
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
		it.Event = new(RewardValidatorRewardWithdrawal)
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
func (it *RewardValidatorRewardWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardValidatorRewardWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardValidatorRewardWithdrawal represents a ValidatorRewardWithdrawal event raised by the Reward contract.
type RewardValidatorRewardWithdrawal struct {
	Validator common.Address
	Amount    *big.Int
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRewardWithdrawal is a free log retrieval operation binding the contract event 0xb1208165cd498a8357a1a86bbfc0d721060a056e8071e8d4a251428b5efb24d6.
//
// Solidity: event ValidatorRewardWithdrawal(address indexed validator, uint256 amount, address caller)
func (_Reward *RewardFilterer) FilterValidatorRewardWithdrawal(opts *bind.FilterOpts, validator []common.Address) (*RewardValidatorRewardWithdrawalIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Reward.contract.FilterLogs(opts, "ValidatorRewardWithdrawal", validatorRule)
	if err != nil {
		return nil, err
	}
	return &RewardValidatorRewardWithdrawalIterator{contract: _Reward.contract, event: "ValidatorRewardWithdrawal", logs: logs, sub: sub}, nil
}

// WatchValidatorRewardWithdrawal is a free log subscription operation binding the contract event 0xb1208165cd498a8357a1a86bbfc0d721060a056e8071e8d4a251428b5efb24d6.
//
// Solidity: event ValidatorRewardWithdrawal(address indexed validator, uint256 amount, address caller)
func (_Reward *RewardFilterer) WatchValidatorRewardWithdrawal(opts *bind.WatchOpts, sink chan<- *RewardValidatorRewardWithdrawal, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Reward.contract.WatchLogs(opts, "ValidatorRewardWithdrawal", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardValidatorRewardWithdrawal)
				if err := _Reward.contract.UnpackLog(event, "ValidatorRewardWithdrawal", log); err != nil {
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

// ParseValidatorRewardWithdrawal is a log parse operation binding the contract event 0xb1208165cd498a8357a1a86bbfc0d721060a056e8071e8d4a251428b5efb24d6.
//
// Solidity: event ValidatorRewardWithdrawal(address indexed validator, uint256 amount, address caller)
func (_Reward *RewardFilterer) ParseValidatorRewardWithdrawal(log types.Log) (*RewardValidatorRewardWithdrawal, error) {
	event := new(RewardValidatorRewardWithdrawal)
	if err := _Reward.contract.UnpackLog(event, "ValidatorRewardWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
