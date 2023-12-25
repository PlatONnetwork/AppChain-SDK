// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2withdraw

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

// WithdrawABI is the input ABI used to generate the binding from.
const WithdrawABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"L2MintableCoinDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"L2MintableCoinWithdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Withdraw is an auto generated Go binding around an platon contract.
type Withdraw struct {
	WithdrawCaller     // Read-only binding to the contract
	WithdrawTransactor // Write-only binding to the contract
	WithdrawFilterer   // Log filterer for contract events
}

// WithdrawCaller is an auto generated read-only Go binding around an platon contract.
type WithdrawCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawTransactor is an auto generated write-only Go binding around an platon contract.
type WithdrawTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawFilterer is an auto generated log filtering Go binding around an platon contract events.
type WithdrawFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type WithdrawSession struct {
	Contract     *Withdraw         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WithdrawCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type WithdrawCallerSession struct {
	Contract *WithdrawCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// WithdrawTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type WithdrawTransactorSession struct {
	Contract     *WithdrawTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WithdrawRaw is an auto generated low-level Go binding around an platon contract.
type WithdrawRaw struct {
	Contract *Withdraw // Generic contract binding to access the raw methods on
}

// WithdrawCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type WithdrawCallerRaw struct {
	Contract *WithdrawCaller // Generic read-only contract binding to access the raw methods on
}

// WithdrawTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type WithdrawTransactorRaw struct {
	Contract *WithdrawTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdraw creates a new instance of Withdraw, bound to a specific deployed contract.
func NewWithdraw(address common.Address, backend bind.ContractBackend) (*Withdraw, error) {
	contract, err := bindWithdraw(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Withdraw{WithdrawCaller: WithdrawCaller{contract: contract}, WithdrawTransactor: WithdrawTransactor{contract: contract}, WithdrawFilterer: WithdrawFilterer{contract: contract}}, nil
}

// NewWithdrawCaller creates a new read-only instance of Withdraw, bound to a specific deployed contract.
func NewWithdrawCaller(address common.Address, caller bind.ContractCaller) (*WithdrawCaller, error) {
	contract, err := bindWithdraw(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawCaller{contract: contract}, nil
}

// NewWithdrawTransactor creates a new write-only instance of Withdraw, bound to a specific deployed contract.
func NewWithdrawTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawTransactor, error) {
	contract, err := bindWithdraw(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawTransactor{contract: contract}, nil
}

// NewWithdrawFilterer creates a new log filterer instance of Withdraw, bound to a specific deployed contract.
func NewWithdrawFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawFilterer, error) {
	contract, err := bindWithdraw(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithdrawFilterer{contract: contract}, nil
}

// bindWithdraw binds a generic wrapper to an already deployed contract.
func bindWithdraw(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WithdrawABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdraw *WithdrawRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Withdraw.Contract.WithdrawCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdraw *WithdrawRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdraw.Contract.WithdrawTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdraw *WithdrawRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdraw.Contract.WithdrawTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdraw *WithdrawCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Withdraw.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdraw *WithdrawTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdraw.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdraw *WithdrawTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdraw.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address recipient, uint256 amount) returns()
func (_Withdraw *WithdrawTransactor) Deposit(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Withdraw.contract.Transact(opts, "deposit", recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address recipient, uint256 amount) returns()
func (_Withdraw *WithdrawSession) Deposit(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Withdraw.Contract.Deposit(&_Withdraw.TransactOpts, recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address recipient, uint256 amount) returns()
func (_Withdraw *WithdrawTransactorSession) Deposit(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Withdraw.Contract.Deposit(&_Withdraw.TransactOpts, recipient, amount)
}

// WithdrawL2MintableCoinDepositIterator is returned from FilterL2MintableCoinDeposit and is used to iterate over the raw logs and unpacked data for L2MintableCoinDeposit events raised by the Withdraw contract.
type WithdrawL2MintableCoinDepositIterator struct {
	Event *WithdrawL2MintableCoinDeposit // Event containing the contract specifics and raw log

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
func (it *WithdrawL2MintableCoinDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawL2MintableCoinDeposit)
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
		it.Event = new(WithdrawL2MintableCoinDeposit)
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
func (it *WithdrawL2MintableCoinDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawL2MintableCoinDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawL2MintableCoinDeposit represents a L2MintableCoinDeposit event raised by the Withdraw contract.
type WithdrawL2MintableCoinDeposit struct {
	Recipient common.Address
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterL2MintableCoinDeposit is a free log retrieval operation binding the contract event 0x0ec6f36ecdd4a102a39d966bb57706b9b1823fc660361d025999d6c26364d3bc.
//
// Solidity: event L2MintableCoinDeposit(address indexed recipient, address depositor, uint256 amount)
func (_Withdraw *WithdrawFilterer) FilterL2MintableCoinDeposit(opts *bind.FilterOpts, recipient []common.Address) (*WithdrawL2MintableCoinDepositIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Withdraw.contract.FilterLogs(opts, "L2MintableCoinDeposit", recipientRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawL2MintableCoinDepositIterator{contract: _Withdraw.contract, event: "L2MintableCoinDeposit", logs: logs, sub: sub}, nil
}

// WatchL2MintableCoinDeposit is a free log subscription operation binding the contract event 0x0ec6f36ecdd4a102a39d966bb57706b9b1823fc660361d025999d6c26364d3bc.
//
// Solidity: event L2MintableCoinDeposit(address indexed recipient, address depositor, uint256 amount)
func (_Withdraw *WithdrawFilterer) WatchL2MintableCoinDeposit(opts *bind.WatchOpts, sink chan<- *WithdrawL2MintableCoinDeposit, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Withdraw.contract.WatchLogs(opts, "L2MintableCoinDeposit", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawL2MintableCoinDeposit)
				if err := _Withdraw.contract.UnpackLog(event, "L2MintableCoinDeposit", log); err != nil {
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

// ParseL2MintableCoinDeposit is a log parse operation binding the contract event 0x0ec6f36ecdd4a102a39d966bb57706b9b1823fc660361d025999d6c26364d3bc.
//
// Solidity: event L2MintableCoinDeposit(address indexed recipient, address depositor, uint256 amount)
func (_Withdraw *WithdrawFilterer) ParseL2MintableCoinDeposit(log types.Log) (*WithdrawL2MintableCoinDeposit, error) {
	event := new(WithdrawL2MintableCoinDeposit)
	if err := _Withdraw.contract.UnpackLog(event, "L2MintableCoinDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawL2MintableCoinWithdrawIterator is returned from FilterL2MintableCoinWithdraw and is used to iterate over the raw logs and unpacked data for L2MintableCoinWithdraw events raised by the Withdraw contract.
type WithdrawL2MintableCoinWithdrawIterator struct {
	Event *WithdrawL2MintableCoinWithdraw // Event containing the contract specifics and raw log

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
func (it *WithdrawL2MintableCoinWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawL2MintableCoinWithdraw)
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
		it.Event = new(WithdrawL2MintableCoinWithdraw)
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
func (it *WithdrawL2MintableCoinWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawL2MintableCoinWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawL2MintableCoinWithdraw represents a L2MintableCoinWithdraw event raised by the Withdraw contract.
type WithdrawL2MintableCoinWithdraw struct {
	Recipient  common.Address
	Withdrawer common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterL2MintableCoinWithdraw is a free log retrieval operation binding the contract event 0x4eae782f2ab5ccea689c03097f004950b274138ab2b25b75f6f673537f37c3d7.
//
// Solidity: event L2MintableCoinWithdraw(address indexed recipient, address withdrawer, uint256 amount)
func (_Withdraw *WithdrawFilterer) FilterL2MintableCoinWithdraw(opts *bind.FilterOpts, recipient []common.Address) (*WithdrawL2MintableCoinWithdrawIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Withdraw.contract.FilterLogs(opts, "L2MintableCoinWithdraw", recipientRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawL2MintableCoinWithdrawIterator{contract: _Withdraw.contract, event: "L2MintableCoinWithdraw", logs: logs, sub: sub}, nil
}

// WatchL2MintableCoinWithdraw is a free log subscription operation binding the contract event 0x4eae782f2ab5ccea689c03097f004950b274138ab2b25b75f6f673537f37c3d7.
//
// Solidity: event L2MintableCoinWithdraw(address indexed recipient, address withdrawer, uint256 amount)
func (_Withdraw *WithdrawFilterer) WatchL2MintableCoinWithdraw(opts *bind.WatchOpts, sink chan<- *WithdrawL2MintableCoinWithdraw, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Withdraw.contract.WatchLogs(opts, "L2MintableCoinWithdraw", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawL2MintableCoinWithdraw)
				if err := _Withdraw.contract.UnpackLog(event, "L2MintableCoinWithdraw", log); err != nil {
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

// ParseL2MintableCoinWithdraw is a log parse operation binding the contract event 0x4eae782f2ab5ccea689c03097f004950b274138ab2b25b75f6f673537f37c3d7.
//
// Solidity: event L2MintableCoinWithdraw(address indexed recipient, address withdrawer, uint256 amount)
func (_Withdraw *WithdrawFilterer) ParseL2MintableCoinWithdraw(log types.Log) (*WithdrawL2MintableCoinWithdraw, error) {
	event := new(WithdrawL2MintableCoinWithdraw)
	if err := _Withdraw.contract.UnpackLog(event, "L2MintableCoinWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
