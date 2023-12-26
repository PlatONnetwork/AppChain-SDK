// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l1exithelper

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

// IExitHelperBatchExitInput is an auto generated low-level Go binding around an user-defined struct.
type IExitHelperBatchExitInput struct {
	BlockNumber  *big.Int
	LeafIndex    *big.Int
	UnhashedLeaf []byte
	Proof        [][32]byte
}

// ExithelperMetaData contains all meta data concerning the Exithelper contract.
var ExithelperMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExitProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"unhashedLeaf\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIExitHelper.BatchExitInput[]\",\"name\":\"inputs\",\"type\":\"tuple[]\"}],\"name\":\"batchExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"caller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkpointManager\",\"outputs\":[{\"internalType\":\"contractICheckpointManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leafIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"unhashedLeaf\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCheckpointManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"processedExits\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ExithelperABI is the input ABI used to generate the binding from.
// Deprecated: Use ExithelperMetaData.ABI instead.
var ExithelperABI = ExithelperMetaData.ABI

// Exithelper is an auto generated Go binding around an platon contract.
type Exithelper struct {
	ExithelperCaller     // Read-only binding to the contract
	ExithelperTransactor // Write-only binding to the contract
	ExithelperFilterer   // Log filterer for contract events
}

// ExithelperCaller is an auto generated read-only Go binding around an platon contract.
type ExithelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExithelperTransactor is an auto generated write-only Go binding around an platon contract.
type ExithelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExithelperFilterer is an auto generated log filtering Go binding around an platon contract events.
type ExithelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExithelperSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type ExithelperSession struct {
	Contract     *Exithelper       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExithelperCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type ExithelperCallerSession struct {
	Contract *ExithelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ExithelperTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type ExithelperTransactorSession struct {
	Contract     *ExithelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ExithelperRaw is an auto generated low-level Go binding around an platon contract.
type ExithelperRaw struct {
	Contract *Exithelper // Generic contract binding to access the raw methods on
}

// ExithelperCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type ExithelperCallerRaw struct {
	Contract *ExithelperCaller // Generic read-only contract binding to access the raw methods on
}

// ExithelperTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type ExithelperTransactorRaw struct {
	Contract *ExithelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExithelper creates a new instance of Exithelper, bound to a specific deployed contract.
func NewExithelper(address common.Address, backend bind.ContractBackend) (*Exithelper, error) {
	contract, err := bindExithelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exithelper{ExithelperCaller: ExithelperCaller{contract: contract}, ExithelperTransactor: ExithelperTransactor{contract: contract}, ExithelperFilterer: ExithelperFilterer{contract: contract}}, nil
}

// NewExithelperCaller creates a new read-only instance of Exithelper, bound to a specific deployed contract.
func NewExithelperCaller(address common.Address, caller bind.ContractCaller) (*ExithelperCaller, error) {
	contract, err := bindExithelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExithelperCaller{contract: contract}, nil
}

// NewExithelperTransactor creates a new write-only instance of Exithelper, bound to a specific deployed contract.
func NewExithelperTransactor(address common.Address, transactor bind.ContractTransactor) (*ExithelperTransactor, error) {
	contract, err := bindExithelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExithelperTransactor{contract: contract}, nil
}

// NewExithelperFilterer creates a new log filterer instance of Exithelper, bound to a specific deployed contract.
func NewExithelperFilterer(address common.Address, filterer bind.ContractFilterer) (*ExithelperFilterer, error) {
	contract, err := bindExithelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExithelperFilterer{contract: contract}, nil
}

// bindExithelper binds a generic wrapper to an already deployed contract.
func bindExithelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExithelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exithelper *ExithelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exithelper.Contract.ExithelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exithelper *ExithelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exithelper.Contract.ExithelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exithelper *ExithelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exithelper.Contract.ExithelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exithelper *ExithelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exithelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exithelper *ExithelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exithelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exithelper *ExithelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exithelper.Contract.contract.Transact(opts, method, params...)
}

// Caller is a free data retrieval call binding the contract method 0xfc9c8d39.
//
// Solidity: function caller() view returns(address)
func (_Exithelper *ExithelperCaller) Caller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exithelper.contract.Call(opts, &out, "caller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Caller is a free data retrieval call binding the contract method 0xfc9c8d39.
//
// Solidity: function caller() view returns(address)
func (_Exithelper *ExithelperSession) Caller() (common.Address, error) {
	return _Exithelper.Contract.Caller(&_Exithelper.CallOpts)
}

// Caller is a free data retrieval call binding the contract method 0xfc9c8d39.
//
// Solidity: function caller() view returns(address)
func (_Exithelper *ExithelperCallerSession) Caller() (common.Address, error) {
	return _Exithelper.Contract.Caller(&_Exithelper.CallOpts)
}

// CheckpointManager is a free data retrieval call binding the contract method 0xc0857ba0.
//
// Solidity: function checkpointManager() view returns(address)
func (_Exithelper *ExithelperCaller) CheckpointManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exithelper.contract.Call(opts, &out, "checkpointManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CheckpointManager is a free data retrieval call binding the contract method 0xc0857ba0.
//
// Solidity: function checkpointManager() view returns(address)
func (_Exithelper *ExithelperSession) CheckpointManager() (common.Address, error) {
	return _Exithelper.Contract.CheckpointManager(&_Exithelper.CallOpts)
}

// CheckpointManager is a free data retrieval call binding the contract method 0xc0857ba0.
//
// Solidity: function checkpointManager() view returns(address)
func (_Exithelper *ExithelperCallerSession) CheckpointManager() (common.Address, error) {
	return _Exithelper.Contract.CheckpointManager(&_Exithelper.CallOpts)
}

// ProcessedExits is a free data retrieval call binding the contract method 0xbd88ea79.
//
// Solidity: function processedExits(uint256 ) view returns(bool)
func (_Exithelper *ExithelperCaller) ProcessedExits(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Exithelper.contract.Call(opts, &out, "processedExits", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedExits is a free data retrieval call binding the contract method 0xbd88ea79.
//
// Solidity: function processedExits(uint256 ) view returns(bool)
func (_Exithelper *ExithelperSession) ProcessedExits(arg0 *big.Int) (bool, error) {
	return _Exithelper.Contract.ProcessedExits(&_Exithelper.CallOpts, arg0)
}

// ProcessedExits is a free data retrieval call binding the contract method 0xbd88ea79.
//
// Solidity: function processedExits(uint256 ) view returns(bool)
func (_Exithelper *ExithelperCallerSession) ProcessedExits(arg0 *big.Int) (bool, error) {
	return _Exithelper.Contract.ProcessedExits(&_Exithelper.CallOpts, arg0)
}

// BatchExit is a paid mutator transaction binding the contract method 0x50607b35.
//
// Solidity: function batchExit((uint256,uint256,bytes,bytes32[])[] inputs) returns()
func (_Exithelper *ExithelperTransactor) BatchExit(opts *bind.TransactOpts, inputs []IExitHelperBatchExitInput) (*types.Transaction, error) {
	return _Exithelper.contract.Transact(opts, "batchExit", inputs)
}

// BatchExit is a paid mutator transaction binding the contract method 0x50607b35.
//
// Solidity: function batchExit((uint256,uint256,bytes,bytes32[])[] inputs) returns()
func (_Exithelper *ExithelperSession) BatchExit(inputs []IExitHelperBatchExitInput) (*types.Transaction, error) {
	return _Exithelper.Contract.BatchExit(&_Exithelper.TransactOpts, inputs)
}

// BatchExit is a paid mutator transaction binding the contract method 0x50607b35.
//
// Solidity: function batchExit((uint256,uint256,bytes,bytes32[])[] inputs) returns()
func (_Exithelper *ExithelperTransactorSession) BatchExit(inputs []IExitHelperBatchExitInput) (*types.Transaction, error) {
	return _Exithelper.Contract.BatchExit(&_Exithelper.TransactOpts, inputs)
}

// Exit is a paid mutator transaction binding the contract method 0xaa209cc3.
//
// Solidity: function exit(uint256 blockNumber, uint256 leafIndex, bytes unhashedLeaf, bytes32[] proof) returns()
func (_Exithelper *ExithelperTransactor) Exit(opts *bind.TransactOpts, blockNumber *big.Int, leafIndex *big.Int, unhashedLeaf []byte, proof [][32]byte) (*types.Transaction, error) {
	return _Exithelper.contract.Transact(opts, "exit", blockNumber, leafIndex, unhashedLeaf, proof)
}

// Exit is a paid mutator transaction binding the contract method 0xaa209cc3.
//
// Solidity: function exit(uint256 blockNumber, uint256 leafIndex, bytes unhashedLeaf, bytes32[] proof) returns()
func (_Exithelper *ExithelperSession) Exit(blockNumber *big.Int, leafIndex *big.Int, unhashedLeaf []byte, proof [][32]byte) (*types.Transaction, error) {
	return _Exithelper.Contract.Exit(&_Exithelper.TransactOpts, blockNumber, leafIndex, unhashedLeaf, proof)
}

// Exit is a paid mutator transaction binding the contract method 0xaa209cc3.
//
// Solidity: function exit(uint256 blockNumber, uint256 leafIndex, bytes unhashedLeaf, bytes32[] proof) returns()
func (_Exithelper *ExithelperTransactorSession) Exit(blockNumber *big.Int, leafIndex *big.Int, unhashedLeaf []byte, proof [][32]byte) (*types.Transaction, error) {
	return _Exithelper.Contract.Exit(&_Exithelper.TransactOpts, blockNumber, leafIndex, unhashedLeaf, proof)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newCheckpointManager) returns()
func (_Exithelper *ExithelperTransactor) Initialize(opts *bind.TransactOpts, newCheckpointManager common.Address) (*types.Transaction, error) {
	return _Exithelper.contract.Transact(opts, "initialize", newCheckpointManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newCheckpointManager) returns()
func (_Exithelper *ExithelperSession) Initialize(newCheckpointManager common.Address) (*types.Transaction, error) {
	return _Exithelper.Contract.Initialize(&_Exithelper.TransactOpts, newCheckpointManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newCheckpointManager) returns()
func (_Exithelper *ExithelperTransactorSession) Initialize(newCheckpointManager common.Address) (*types.Transaction, error) {
	return _Exithelper.Contract.Initialize(&_Exithelper.TransactOpts, newCheckpointManager)
}

// ExithelperExitProcessedIterator is returned from FilterExitProcessed and is used to iterate over the raw logs and unpacked data for ExitProcessed events raised by the Exithelper contract.
type ExithelperExitProcessedIterator struct {
	Event *ExithelperExitProcessed // Event containing the contract specifics and raw log

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
func (it *ExithelperExitProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExithelperExitProcessed)
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
		it.Event = new(ExithelperExitProcessed)
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
func (it *ExithelperExitProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExithelperExitProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExithelperExitProcessed represents a ExitProcessed event raised by the Exithelper contract.
type ExithelperExitProcessed struct {
	Id         *big.Int
	Success    bool
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExitProcessed is a free log retrieval operation binding the contract event 0x8bbfa0c9bee3785c03700d2a909592286efb83fc7e7002be5764424b9842f7ec.
//
// Solidity: event ExitProcessed(uint256 indexed id, bool indexed success, bytes returnData)
func (_Exithelper *ExithelperFilterer) FilterExitProcessed(opts *bind.FilterOpts, id []*big.Int, success []bool) (*ExithelperExitProcessedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _Exithelper.contract.FilterLogs(opts, "ExitProcessed", idRule, successRule)
	if err != nil {
		return nil, err
	}
	return &ExithelperExitProcessedIterator{contract: _Exithelper.contract, event: "ExitProcessed", logs: logs, sub: sub}, nil
}

// WatchExitProcessed is a free log subscription operation binding the contract event 0x8bbfa0c9bee3785c03700d2a909592286efb83fc7e7002be5764424b9842f7ec.
//
// Solidity: event ExitProcessed(uint256 indexed id, bool indexed success, bytes returnData)
func (_Exithelper *ExithelperFilterer) WatchExitProcessed(opts *bind.WatchOpts, sink chan<- *ExithelperExitProcessed, id []*big.Int, success []bool) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _Exithelper.contract.WatchLogs(opts, "ExitProcessed", idRule, successRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExithelperExitProcessed)
				if err := _Exithelper.contract.UnpackLog(event, "ExitProcessed", log); err != nil {
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

// ParseExitProcessed is a log parse operation binding the contract event 0x8bbfa0c9bee3785c03700d2a909592286efb83fc7e7002be5764424b9842f7ec.
//
// Solidity: event ExitProcessed(uint256 indexed id, bool indexed success, bytes returnData)
func (_Exithelper *ExithelperFilterer) ParseExitProcessed(log types.Log) (*ExithelperExitProcessed, error) {
	event := new(ExithelperExitProcessed)
	if err := _Exithelper.contract.UnpackLog(event, "ExitProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExithelperInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Exithelper contract.
type ExithelperInitializedIterator struct {
	Event *ExithelperInitialized // Event containing the contract specifics and raw log

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
func (it *ExithelperInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExithelperInitialized)
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
		it.Event = new(ExithelperInitialized)
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
func (it *ExithelperInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExithelperInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExithelperInitialized represents a Initialized event raised by the Exithelper contract.
type ExithelperInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Exithelper *ExithelperFilterer) FilterInitialized(opts *bind.FilterOpts) (*ExithelperInitializedIterator, error) {

	logs, sub, err := _Exithelper.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ExithelperInitializedIterator{contract: _Exithelper.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Exithelper *ExithelperFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ExithelperInitialized) (event.Subscription, error) {

	logs, sub, err := _Exithelper.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExithelperInitialized)
				if err := _Exithelper.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Exithelper *ExithelperFilterer) ParseInitialized(log types.Log) (*ExithelperInitialized, error) {
	event := new(ExithelperInitialized)
	if err := _Exithelper.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
