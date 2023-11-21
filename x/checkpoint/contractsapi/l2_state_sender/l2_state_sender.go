// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2_state_sender

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

// L2StateSenderMetaData contains all meta data concerning the L2StateSender contract.
var L2StateSenderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"L2StateSynced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"syncState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// L2StateSenderABI is the input ABI used to generate the binding from.
// Deprecated: Use L2StateSenderMetaData.ABI instead.
var L2StateSenderABI = L2StateSenderMetaData.ABI

// L2StateSender is an auto generated Go binding around an platon contract.
type L2StateSender struct {
	L2StateSenderCaller     // Read-only binding to the contract
	L2StateSenderTransactor // Write-only binding to the contract
	L2StateSenderFilterer   // Log filterer for contract events
}

// L2StateSenderCaller is an auto generated read-only Go binding around an platon contract.
type L2StateSenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2StateSenderTransactor is an auto generated write-only Go binding around an platon contract.
type L2StateSenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2StateSenderFilterer is an auto generated log filtering Go binding around an platon contract events.
type L2StateSenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2StateSenderSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type L2StateSenderSession struct {
	Contract     *L2StateSender    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2StateSenderCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type L2StateSenderCallerSession struct {
	Contract *L2StateSenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L2StateSenderTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type L2StateSenderTransactorSession struct {
	Contract     *L2StateSenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L2StateSenderRaw is an auto generated low-level Go binding around an platon contract.
type L2StateSenderRaw struct {
	Contract *L2StateSender // Generic contract binding to access the raw methods on
}

// L2StateSenderCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type L2StateSenderCallerRaw struct {
	Contract *L2StateSenderCaller // Generic read-only contract binding to access the raw methods on
}

// L2StateSenderTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type L2StateSenderTransactorRaw struct {
	Contract *L2StateSenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2StateSender creates a new instance of L2StateSender, bound to a specific deployed contract.
func NewL2StateSender(address common.Address, backend bind.ContractBackend) (*L2StateSender, error) {
	contract, err := bindL2StateSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2StateSender{L2StateSenderCaller: L2StateSenderCaller{contract: contract}, L2StateSenderTransactor: L2StateSenderTransactor{contract: contract}, L2StateSenderFilterer: L2StateSenderFilterer{contract: contract}}, nil
}

// NewL2StateSenderCaller creates a new read-only instance of L2StateSender, bound to a specific deployed contract.
func NewL2StateSenderCaller(address common.Address, caller bind.ContractCaller) (*L2StateSenderCaller, error) {
	contract, err := bindL2StateSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2StateSenderCaller{contract: contract}, nil
}

// NewL2StateSenderTransactor creates a new write-only instance of L2StateSender, bound to a specific deployed contract.
func NewL2StateSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*L2StateSenderTransactor, error) {
	contract, err := bindL2StateSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2StateSenderTransactor{contract: contract}, nil
}

// NewL2StateSenderFilterer creates a new log filterer instance of L2StateSender, bound to a specific deployed contract.
func NewL2StateSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*L2StateSenderFilterer, error) {
	contract, err := bindL2StateSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2StateSenderFilterer{contract: contract}, nil
}

// bindL2StateSender binds a generic wrapper to an already deployed contract.
func bindL2StateSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2StateSenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2StateSender *L2StateSenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2StateSender.Contract.L2StateSenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2StateSender *L2StateSenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2StateSender.Contract.L2StateSenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2StateSender *L2StateSenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2StateSender.Contract.L2StateSenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2StateSender *L2StateSenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2StateSender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2StateSender *L2StateSenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2StateSender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2StateSender *L2StateSenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2StateSender.Contract.contract.Transact(opts, method, params...)
}

// MAXLENGTH is a free data retrieval call binding the contract method 0xa6f9885c.
//
// Solidity: function MAX_LENGTH() view returns(uint256)
func (_L2StateSender *L2StateSenderCaller) MAXLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2StateSender.contract.Call(opts, &out, "MAX_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLENGTH is a free data retrieval call binding the contract method 0xa6f9885c.
//
// Solidity: function MAX_LENGTH() view returns(uint256)
func (_L2StateSender *L2StateSenderSession) MAXLENGTH() (*big.Int, error) {
	return _L2StateSender.Contract.MAXLENGTH(&_L2StateSender.CallOpts)
}

// MAXLENGTH is a free data retrieval call binding the contract method 0xa6f9885c.
//
// Solidity: function MAX_LENGTH() view returns(uint256)
func (_L2StateSender *L2StateSenderCallerSession) MAXLENGTH() (*big.Int, error) {
	return _L2StateSender.Contract.MAXLENGTH(&_L2StateSender.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_L2StateSender *L2StateSenderCaller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2StateSender.contract.Call(opts, &out, "counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_L2StateSender *L2StateSenderSession) Counter() (*big.Int, error) {
	return _L2StateSender.Contract.Counter(&_L2StateSender.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_L2StateSender *L2StateSenderCallerSession) Counter() (*big.Int, error) {
	return _L2StateSender.Contract.Counter(&_L2StateSender.CallOpts)
}

// SyncState is a paid mutator transaction binding the contract method 0x16f19831.
//
// Solidity: function syncState(address receiver, bytes data) returns()
func (_L2StateSender *L2StateSenderTransactor) SyncState(opts *bind.TransactOpts, receiver common.Address, data []byte) (*types.Transaction, error) {
	return _L2StateSender.contract.Transact(opts, "syncState", receiver, data)
}

// SyncState is a paid mutator transaction binding the contract method 0x16f19831.
//
// Solidity: function syncState(address receiver, bytes data) returns()
func (_L2StateSender *L2StateSenderSession) SyncState(receiver common.Address, data []byte) (*types.Transaction, error) {
	return _L2StateSender.Contract.SyncState(&_L2StateSender.TransactOpts, receiver, data)
}

// SyncState is a paid mutator transaction binding the contract method 0x16f19831.
//
// Solidity: function syncState(address receiver, bytes data) returns()
func (_L2StateSender *L2StateSenderTransactorSession) SyncState(receiver common.Address, data []byte) (*types.Transaction, error) {
	return _L2StateSender.Contract.SyncState(&_L2StateSender.TransactOpts, receiver, data)
}

// L2StateSenderL2StateSyncedIterator is returned from FilterL2StateSynced and is used to iterate over the raw logs and unpacked data for L2StateSynced events raised by the L2StateSender contract.
type L2StateSenderL2StateSyncedIterator struct {
	Event *L2StateSenderL2StateSynced // Event containing the contract specifics and raw log

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
func (it *L2StateSenderL2StateSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StateSenderL2StateSynced)
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
		it.Event = new(L2StateSenderL2StateSynced)
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
func (it *L2StateSenderL2StateSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StateSenderL2StateSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StateSenderL2StateSynced represents a L2StateSynced event raised by the L2StateSender contract.
type L2StateSenderL2StateSynced struct {
	Id       *big.Int
	Sender   common.Address
	Receiver common.Address
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterL2StateSynced is a free log retrieval operation binding the contract event 0xedaf3c471ebd67d60c29efe34b639ede7d6a1d92eaeb3f503e784971e67118a5.
//
// Solidity: event L2StateSynced(uint256 indexed id, address indexed sender, address indexed receiver, bytes data)
func (_L2StateSender *L2StateSenderFilterer) FilterL2StateSynced(opts *bind.FilterOpts, id []*big.Int, sender []common.Address, receiver []common.Address) (*L2StateSenderL2StateSyncedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _L2StateSender.contract.FilterLogs(opts, "L2StateSynced", idRule, senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &L2StateSenderL2StateSyncedIterator{contract: _L2StateSender.contract, event: "L2StateSynced", logs: logs, sub: sub}, nil
}

// WatchL2StateSynced is a free log subscription operation binding the contract event 0xedaf3c471ebd67d60c29efe34b639ede7d6a1d92eaeb3f503e784971e67118a5.
//
// Solidity: event L2StateSynced(uint256 indexed id, address indexed sender, address indexed receiver, bytes data)
func (_L2StateSender *L2StateSenderFilterer) WatchL2StateSynced(opts *bind.WatchOpts, sink chan<- *L2StateSenderL2StateSynced, id []*big.Int, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _L2StateSender.contract.WatchLogs(opts, "L2StateSynced", idRule, senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StateSenderL2StateSynced)
				if err := _L2StateSender.contract.UnpackLog(event, "L2StateSynced", log); err != nil {
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

// ParseL2StateSynced is a log parse operation binding the contract event 0xedaf3c471ebd67d60c29efe34b639ede7d6a1d92eaeb3f503e784971e67118a5.
//
// Solidity: event L2StateSynced(uint256 indexed id, address indexed sender, address indexed receiver, bytes data)
func (_L2StateSender *L2StateSenderFilterer) ParseL2StateSynced(log types.Log) (*L2StateSenderL2StateSynced, error) {
	event := new(L2StateSenderL2StateSynced)
	if err := _L2StateSender.contract.UnpackLog(event, "L2StateSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
