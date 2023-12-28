// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2deposit

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

// DepositABI is the input ABI used to generate the binding from.
const DepositABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"L2CoinDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"L2CoinWithdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onStateReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Deposit is an auto generated Go binding around an platon contract.
type Deposit struct {
	DepositCaller     // Read-only binding to the contract
	DepositTransactor // Write-only binding to the contract
	DepositFilterer   // Log filterer for contract events
}

// DepositCaller is an auto generated read-only Go binding around an platon contract.
type DepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositTransactor is an auto generated write-only Go binding around an platon contract.
type DepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositFilterer is an auto generated log filtering Go binding around an platon contract events.
type DepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type DepositSession struct {
	Contract     *Deposit          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type DepositCallerSession struct {
	Contract *DepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DepositTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type DepositTransactorSession struct {
	Contract     *DepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DepositRaw is an auto generated low-level Go binding around an platon contract.
type DepositRaw struct {
	Contract *Deposit // Generic contract binding to access the raw methods on
}

// DepositCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type DepositCallerRaw struct {
	Contract *DepositCaller // Generic read-only contract binding to access the raw methods on
}

// DepositTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type DepositTransactorRaw struct {
	Contract *DepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeposit creates a new instance of Deposit, bound to a specific deployed contract.
func NewDeposit(address common.Address, backend bind.ContractBackend) (*Deposit, error) {
	contract, err := bindDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deposit{DepositCaller: DepositCaller{contract: contract}, DepositTransactor: DepositTransactor{contract: contract}, DepositFilterer: DepositFilterer{contract: contract}}, nil
}

// NewDepositCaller creates a new read-only instance of Deposit, bound to a specific deployed contract.
func NewDepositCaller(address common.Address, caller bind.ContractCaller) (*DepositCaller, error) {
	contract, err := bindDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositCaller{contract: contract}, nil
}

// NewDepositTransactor creates a new write-only instance of Deposit, bound to a specific deployed contract.
func NewDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositTransactor, error) {
	contract, err := bindDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositTransactor{contract: contract}, nil
}

// NewDepositFilterer creates a new log filterer instance of Deposit, bound to a specific deployed contract.
func NewDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositFilterer, error) {
	contract, err := bindDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositFilterer{contract: contract}, nil
}

// bindDeposit binds a generic wrapper to an already deployed contract.
func bindDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deposit *DepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deposit.Contract.DepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deposit *DepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deposit.Contract.DepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deposit *DepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deposit.Contract.DepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deposit *DepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deposit *DepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deposit *DepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deposit.Contract.contract.Transact(opts, method, params...)
}

// OnStateReceive is a paid mutator transaction binding the contract method 0xeeb49945.
//
// Solidity: function onStateReceive(uint256 , address sender, bytes data) returns()
func (_Deposit *DepositTransactor) OnStateReceive(opts *bind.TransactOpts, arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "onStateReceive", arg0, sender, data)
}

// OnStateReceive is a paid mutator transaction binding the contract method 0xeeb49945.
//
// Solidity: function onStateReceive(uint256 , address sender, bytes data) returns()
func (_Deposit *DepositSession) OnStateReceive(arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _Deposit.Contract.OnStateReceive(&_Deposit.TransactOpts, arg0, sender, data)
}

// OnStateReceive is a paid mutator transaction binding the contract method 0xeeb49945.
//
// Solidity: function onStateReceive(uint256 , address sender, bytes data) returns()
func (_Deposit *DepositTransactorSession) OnStateReceive(arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _Deposit.Contract.OnStateReceive(&_Deposit.TransactOpts, arg0, sender, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_Deposit *DepositTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "withdraw", recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_Deposit *DepositSession) Withdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Deposit.Contract.Withdraw(&_Deposit.TransactOpts, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_Deposit *DepositTransactorSession) Withdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Deposit.Contract.Withdraw(&_Deposit.TransactOpts, recipient, amount)
}

// DepositL2CoinDepositIterator is returned from FilterL2CoinDeposit and is used to iterate over the raw logs and unpacked data for L2CoinDeposit events raised by the Deposit contract.
type DepositL2CoinDepositIterator struct {
	Event *DepositL2CoinDeposit // Event containing the contract specifics and raw log

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
func (it *DepositL2CoinDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositL2CoinDeposit)
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
		it.Event = new(DepositL2CoinDeposit)
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
func (it *DepositL2CoinDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositL2CoinDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositL2CoinDeposit represents a L2CoinDeposit event raised by the Deposit contract.
type DepositL2CoinDeposit struct {
	Recipient common.Address
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterL2CoinDeposit is a free log retrieval operation binding the contract event 0x84ef242c1a7dc47c0b6c5be725f7daff6231e45e140c08d04b53db2199c80f18.
//
// Solidity: event L2CoinDeposit(address indexed recipient, address depositor, uint256 amount)
func (_Deposit *DepositFilterer) FilterL2CoinDeposit(opts *bind.FilterOpts, recipient []common.Address) (*DepositL2CoinDepositIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Deposit.contract.FilterLogs(opts, "L2CoinDeposit", recipientRule)
	if err != nil {
		return nil, err
	}
	return &DepositL2CoinDepositIterator{contract: _Deposit.contract, event: "L2CoinDeposit", logs: logs, sub: sub}, nil
}

// WatchL2CoinDeposit is a free log subscription operation binding the contract event 0x84ef242c1a7dc47c0b6c5be725f7daff6231e45e140c08d04b53db2199c80f18.
//
// Solidity: event L2CoinDeposit(address indexed recipient, address depositor, uint256 amount)
func (_Deposit *DepositFilterer) WatchL2CoinDeposit(opts *bind.WatchOpts, sink chan<- *DepositL2CoinDeposit, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Deposit.contract.WatchLogs(opts, "L2CoinDeposit", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositL2CoinDeposit)
				if err := _Deposit.contract.UnpackLog(event, "L2CoinDeposit", log); err != nil {
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

// ParseL2CoinDeposit is a log parse operation binding the contract event 0x84ef242c1a7dc47c0b6c5be725f7daff6231e45e140c08d04b53db2199c80f18.
//
// Solidity: event L2CoinDeposit(address indexed recipient, address depositor, uint256 amount)
func (_Deposit *DepositFilterer) ParseL2CoinDeposit(log types.Log) (*DepositL2CoinDeposit, error) {
	event := new(DepositL2CoinDeposit)
	if err := _Deposit.contract.UnpackLog(event, "L2CoinDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositL2CoinWithdrawIterator is returned from FilterL2CoinWithdraw and is used to iterate over the raw logs and unpacked data for L2CoinWithdraw events raised by the Deposit contract.
type DepositL2CoinWithdrawIterator struct {
	Event *DepositL2CoinWithdraw // Event containing the contract specifics and raw log

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
func (it *DepositL2CoinWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositL2CoinWithdraw)
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
		it.Event = new(DepositL2CoinWithdraw)
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
func (it *DepositL2CoinWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositL2CoinWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositL2CoinWithdraw represents a L2CoinWithdraw event raised by the Deposit contract.
type DepositL2CoinWithdraw struct {
	Recipient  common.Address
	Withdrawer common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterL2CoinWithdraw is a free log retrieval operation binding the contract event 0xe5b45e26026bfc93fbc88850ea58ecadcedab3e10261392bbbd4a9bdb097257d.
//
// Solidity: event L2CoinWithdraw(address indexed recipient, address withdrawer, uint256 amount)
func (_Deposit *DepositFilterer) FilterL2CoinWithdraw(opts *bind.FilterOpts, recipient []common.Address) (*DepositL2CoinWithdrawIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Deposit.contract.FilterLogs(opts, "L2CoinWithdraw", recipientRule)
	if err != nil {
		return nil, err
	}
	return &DepositL2CoinWithdrawIterator{contract: _Deposit.contract, event: "L2CoinWithdraw", logs: logs, sub: sub}, nil
}

// WatchL2CoinWithdraw is a free log subscription operation binding the contract event 0xe5b45e26026bfc93fbc88850ea58ecadcedab3e10261392bbbd4a9bdb097257d.
//
// Solidity: event L2CoinWithdraw(address indexed recipient, address withdrawer, uint256 amount)
func (_Deposit *DepositFilterer) WatchL2CoinWithdraw(opts *bind.WatchOpts, sink chan<- *DepositL2CoinWithdraw, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Deposit.contract.WatchLogs(opts, "L2CoinWithdraw", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositL2CoinWithdraw)
				if err := _Deposit.contract.UnpackLog(event, "L2CoinWithdraw", log); err != nil {
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

// ParseL2CoinWithdraw is a log parse operation binding the contract event 0xe5b45e26026bfc93fbc88850ea58ecadcedab3e10261392bbbd4a9bdb097257d.
//
// Solidity: event L2CoinWithdraw(address indexed recipient, address withdrawer, uint256 amount)
func (_Deposit *DepositFilterer) ParseL2CoinWithdraw(log types.Log) (*DepositL2CoinWithdraw, error) {
	event := new(DepositL2CoinWithdraw)
	if err := _Deposit.contract.UnpackLog(event, "L2CoinWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
