// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2statesender

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

// StatesenderABI is the input ABI used to generate the binding from.
const StatesenderABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"L2StateSynced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"syncState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Statesender is an auto generated Go binding around an platon contract.
type Statesender struct {
	StatesenderCaller     // Read-only binding to the contract
	StatesenderTransactor // Write-only binding to the contract
	StatesenderFilterer   // Log filterer for contract events
}

// StatesenderCaller is an auto generated read-only Go binding around an platon contract.
type StatesenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatesenderTransactor is an auto generated write-only Go binding around an platon contract.
type StatesenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatesenderFilterer is an auto generated log filtering Go binding around an platon contract events.
type StatesenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatesenderSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type StatesenderSession struct {
	Contract     *Statesender      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StatesenderCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type StatesenderCallerSession struct {
	Contract *StatesenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StatesenderTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type StatesenderTransactorSession struct {
	Contract     *StatesenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StatesenderRaw is an auto generated low-level Go binding around an platon contract.
type StatesenderRaw struct {
	Contract *Statesender // Generic contract binding to access the raw methods on
}

// StatesenderCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type StatesenderCallerRaw struct {
	Contract *StatesenderCaller // Generic read-only contract binding to access the raw methods on
}

// StatesenderTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type StatesenderTransactorRaw struct {
	Contract *StatesenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStatesender creates a new instance of Statesender, bound to a specific deployed contract.
func NewStatesender(address common.Address, backend bind.ContractBackend) (*Statesender, error) {
	contract, err := bindStatesender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Statesender{StatesenderCaller: StatesenderCaller{contract: contract}, StatesenderTransactor: StatesenderTransactor{contract: contract}, StatesenderFilterer: StatesenderFilterer{contract: contract}}, nil
}

// NewStatesenderCaller creates a new read-only instance of Statesender, bound to a specific deployed contract.
func NewStatesenderCaller(address common.Address, caller bind.ContractCaller) (*StatesenderCaller, error) {
	contract, err := bindStatesender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StatesenderCaller{contract: contract}, nil
}

// NewStatesenderTransactor creates a new write-only instance of Statesender, bound to a specific deployed contract.
func NewStatesenderTransactor(address common.Address, transactor bind.ContractTransactor) (*StatesenderTransactor, error) {
	contract, err := bindStatesender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StatesenderTransactor{contract: contract}, nil
}

// NewStatesenderFilterer creates a new log filterer instance of Statesender, bound to a specific deployed contract.
func NewStatesenderFilterer(address common.Address, filterer bind.ContractFilterer) (*StatesenderFilterer, error) {
	contract, err := bindStatesender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StatesenderFilterer{contract: contract}, nil
}

// bindStatesender binds a generic wrapper to an already deployed contract.
func bindStatesender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StatesenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Statesender *StatesenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Statesender.Contract.StatesenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Statesender *StatesenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Statesender.Contract.StatesenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Statesender *StatesenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Statesender.Contract.StatesenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Statesender *StatesenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Statesender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Statesender *StatesenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Statesender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Statesender *StatesenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Statesender.Contract.contract.Transact(opts, method, params...)
}

// MAXLENGTH is a free data retrieval call binding the contract method 0xa6f9885c.
//
// Solidity: function MAX_LENGTH() view returns(uint256)
func (_Statesender *StatesenderCaller) MAXLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Statesender.contract.Call(opts, &out, "MAX_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLENGTH is a free data retrieval call binding the contract method 0xa6f9885c.
//
// Solidity: function MAX_LENGTH() view returns(uint256)
func (_Statesender *StatesenderSession) MAXLENGTH() (*big.Int, error) {
	return _Statesender.Contract.MAXLENGTH(&_Statesender.CallOpts)
}

// MAXLENGTH is a free data retrieval call binding the contract method 0xa6f9885c.
//
// Solidity: function MAX_LENGTH() view returns(uint256)
func (_Statesender *StatesenderCallerSession) MAXLENGTH() (*big.Int, error) {
	return _Statesender.Contract.MAXLENGTH(&_Statesender.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Statesender *StatesenderCaller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Statesender.contract.Call(opts, &out, "counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Statesender *StatesenderSession) Counter() (*big.Int, error) {
	return _Statesender.Contract.Counter(&_Statesender.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Statesender *StatesenderCallerSession) Counter() (*big.Int, error) {
	return _Statesender.Contract.Counter(&_Statesender.CallOpts)
}

// SyncState is a paid mutator transaction binding the contract method 0x16f19831.
//
// Solidity: function syncState(address receiver, bytes data) returns()
func (_Statesender *StatesenderTransactor) SyncState(opts *bind.TransactOpts, receiver common.Address, data []byte) (*types.Transaction, error) {
	return _Statesender.contract.Transact(opts, "syncState", receiver, data)
}

// SyncState is a paid mutator transaction binding the contract method 0x16f19831.
//
// Solidity: function syncState(address receiver, bytes data) returns()
func (_Statesender *StatesenderSession) SyncState(receiver common.Address, data []byte) (*types.Transaction, error) {
	return _Statesender.Contract.SyncState(&_Statesender.TransactOpts, receiver, data)
}

// SyncState is a paid mutator transaction binding the contract method 0x16f19831.
//
// Solidity: function syncState(address receiver, bytes data) returns()
func (_Statesender *StatesenderTransactorSession) SyncState(receiver common.Address, data []byte) (*types.Transaction, error) {
	return _Statesender.Contract.SyncState(&_Statesender.TransactOpts, receiver, data)
}

// StatesenderL2StateSyncedIterator is returned from FilterL2StateSynced and is used to iterate over the raw logs and unpacked data for L2StateSynced events raised by the Statesender contract.
type StatesenderL2StateSyncedIterator struct {
	Event *StatesenderL2StateSynced // Event containing the contract specifics and raw log

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
func (it *StatesenderL2StateSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatesenderL2StateSynced)
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
		it.Event = new(StatesenderL2StateSynced)
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
func (it *StatesenderL2StateSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatesenderL2StateSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatesenderL2StateSynced represents a L2StateSynced event raised by the Statesender contract.
type StatesenderL2StateSynced struct {
	Id       *big.Int
	Sender   common.Address
	Receiver common.Address
	CallData []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterL2StateSynced is a free log retrieval operation binding the contract event 0xedaf3c471ebd67d60c29efe34b639ede7d6a1d92eaeb3f503e784971e67118a5.
//
// Solidity: event L2StateSynced(uint256 indexed id, address indexed sender, address indexed receiver, bytes callData)
func (_Statesender *StatesenderFilterer) FilterL2StateSynced(opts *bind.FilterOpts, id []*big.Int, sender []common.Address, receiver []common.Address) (*StatesenderL2StateSyncedIterator, error) {

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

	logs, sub, err := _Statesender.contract.FilterLogs(opts, "L2StateSynced", idRule, senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &StatesenderL2StateSyncedIterator{contract: _Statesender.contract, event: "L2StateSynced", logs: logs, sub: sub}, nil
}

// WatchL2StateSynced is a free log subscription operation binding the contract event 0xedaf3c471ebd67d60c29efe34b639ede7d6a1d92eaeb3f503e784971e67118a5.
//
// Solidity: event L2StateSynced(uint256 indexed id, address indexed sender, address indexed receiver, bytes callData)
func (_Statesender *StatesenderFilterer) WatchL2StateSynced(opts *bind.WatchOpts, sink chan<- *StatesenderL2StateSynced, id []*big.Int, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Statesender.contract.WatchLogs(opts, "L2StateSynced", idRule, senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatesenderL2StateSynced)
				if err := _Statesender.contract.UnpackLog(event, "L2StateSynced", log); err != nil {
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
// Solidity: event L2StateSynced(uint256 indexed id, address indexed sender, address indexed receiver, bytes callData)
func (_Statesender *StatesenderFilterer) ParseL2StateSynced(log types.Log) (*StatesenderL2StateSynced, error) {
	event := new(StatesenderL2StateSynced)
	if err := _Statesender.contract.UnpackLog(event, "L2StateSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
