// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package governancehandler

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

// GovernanceHandlerMetaData contains all meta data concerning the GovernanceHandler contract.
var GovernanceHandlerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"newChildGovernanceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newExitHelper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onL2StateReceive\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalExecuted\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"},{\"name\":\"description\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false}]",
	Bin: "0x608080604052346100bf575f549060ff8260081c1661006d575060ff80821603610033575b60405161081390816100c48239f35b60ff90811916175f557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160ff8152a15f610024565b62461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608490fd5b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c908163485cc95514610554575063f43cda8b14610032575f80fd5b346104715760603660031901126104715761004b6106c3565b67ffffffffffffffff6044351161047157366023604435011215610471576044356004013567ffffffffffffffff8111610471576044350190366024830111610471576001546001600160a01b03163303610503575f5460101c6001600160a01b039081169116036104a5576080604435820312610471576024604435013567ffffffffffffffff811161047157604435019060248101604383011215610471576024820135906101036100fe836106ff565b6106d9565b9260208484815201906044829460051b8201019060248401821161047157604401915b8183106104855750505060448035013567ffffffffffffffff8111610471576044350191602482016043840112156104715760248301356101696100fe826106ff565b93602085838152016044819360051b8301019160248601831161047157604401905b82821061047557505050606460443501359067ffffffffffffffff821161047157602484016043836044350101121561047157602482604435010135936101d46100fe866106ff565b946020868281520180946024840160448460051b838235010101116104715760448181350101915b60448460051b8382350101018310610404575050505050604051926080840160808552875180915260a0850191905f5b8181106103e557505050838103602085015260208651918281520191905f5b8181106103cf5750505082810360408401528351808252602082019160208260051b82010193925f915b8383106103a2576044356084013560608801528989897fd228733f9d7ff127e0dca16765efb34ebdc4e2dfe78b3e12aea94aeb0ed210b38a8a038ba16040516060810181811067ffffffffffffffff82111761038e57604052603081527f476f7665726e616e63654d616e616765723a2063616c6c20726576657274656460208201526f20776974686f7574206d65737361676560801b60408201525f5b845181101561038c5760019061037d835f806001600160a01b03610337868c610771565b5116610343868b610771565b519061034f878b610771565b5191602083519301915af13d15610384573d9061036e6100fe83610717565b9182523d5f602084013e610799565b5001610313565b606090610799565b005b634e487b7160e01b5f52604160045260245ffd5b90919293946020806103c0600193601f198682030187528951610733565b97019301930191939290610275565b825184526020938401939092019160010161024b565b82516001600160a01b031684526020938401939092019160010161022c565b823567ffffffffffffffff811161047157826044350101602486016063820112156104715760448101359061043b6100fe83610717565b928284526024880160648484010111610471576044935f60208581966064839701838601378301015281520193019290506101fc565b5f80fd5b813581526020918201910161018b565b82356001600160a01b038116810361047157815260209283019201610126565b60405162461bcd60e51b815260206004820152603060248201527f476f7665726e616e636548616e646c65723a204f4e4c595f4348494c445f474f60448201526f2b22a92720a721a2afa420a7222622a960811b6064820152608490fd5b60405162461bcd60e51b815260206004820152602360248201527f476f7665726e616e636548616e646c65723a204f4e4c595f455849545f48454c6044820152622822a960e91b6064820152608490fd5b34610471576040366003190112610471576001600160a01b0390600435908282168203610471576105836106c3565b925f5460ff8160081c1615928380946106b6575b801561069f575b15610646575060ff1981166001175f5582610635575b505f549262010000600160b01b039060101b16938462010000600160b01b03198516175f55166bffffffffffffffffffffffff60a01b60015416176001556105f857005b610100600160b01b031916175f55604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a1005b61ffff1916610101175f55846105b4565b62461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b50303b15801561059e5750600160ff83161461059e565b50600160ff831610610597565b602435906001600160a01b038216820361047157565b6040519190601f01601f1916820167ffffffffffffffff81118382101761038e57604052565b67ffffffffffffffff811161038e5760051b60200190565b67ffffffffffffffff811161038e57601f01601f191660200190565b91908251928382525f5b84811061075d575050825f602080949584010152601f8019910116010190565b60208183018101518483018201520161073d565b80518210156107855760209160051b010190565b634e487b7160e01b5f52603260045260245ffd5b909190156107a5575090565b8151156107b55750805190602001fd5b60405162461bcd60e51b8152602060048201529081906107d9906024830190610733565b0390fdfea264697066735822122094755c36e4b478ef641c691a757909953630bd4d3ad3a76e6a50abf955e3a1f864736f6c63430008160033",
}

// GovernanceHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernanceHandlerMetaData.ABI instead.
var GovernanceHandlerABI = GovernanceHandlerMetaData.ABI

// GovernanceHandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovernanceHandlerMetaData.Bin instead.
var GovernanceHandlerBin = GovernanceHandlerMetaData.Bin

// DeployGovernanceHandler deploys a new platon contract, binding an instance of GovernanceHandler to it.
func DeployGovernanceHandler(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GovernanceHandler, error) {
	parsed, err := GovernanceHandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovernanceHandlerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GovernanceHandler{GovernanceHandlerCaller: GovernanceHandlerCaller{contract: contract}, GovernanceHandlerTransactor: GovernanceHandlerTransactor{contract: contract}, GovernanceHandlerFilterer: GovernanceHandlerFilterer{contract: contract}}, nil
}

// GovernanceHandler is an auto generated Go binding around an platon contract.
type GovernanceHandler struct {
	GovernanceHandlerCaller     // Read-only binding to the contract
	GovernanceHandlerTransactor // Write-only binding to the contract
	GovernanceHandlerFilterer   // Log filterer for contract events
}

// GovernanceHandlerCaller is an auto generated read-only Go binding around an platon contract.
type GovernanceHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceHandlerTransactor is an auto generated write-only Go binding around an platon contract.
type GovernanceHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceHandlerFilterer is an auto generated log filtering Go binding around an platon contract events.
type GovernanceHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceHandlerSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type GovernanceHandlerSession struct {
	Contract     *GovernanceHandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GovernanceHandlerCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type GovernanceHandlerCallerSession struct {
	Contract *GovernanceHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// GovernanceHandlerTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type GovernanceHandlerTransactorSession struct {
	Contract     *GovernanceHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// GovernanceHandlerRaw is an auto generated low-level Go binding around an platon contract.
type GovernanceHandlerRaw struct {
	Contract *GovernanceHandler // Generic contract binding to access the raw methods on
}

// GovernanceHandlerCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type GovernanceHandlerCallerRaw struct {
	Contract *GovernanceHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceHandlerTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type GovernanceHandlerTransactorRaw struct {
	Contract *GovernanceHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernanceHandler creates a new instance of GovernanceHandler, bound to a specific deployed contract.
func NewGovernanceHandler(address common.Address, backend bind.ContractBackend) (*GovernanceHandler, error) {
	contract, err := bindGovernanceHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernanceHandler{GovernanceHandlerCaller: GovernanceHandlerCaller{contract: contract}, GovernanceHandlerTransactor: GovernanceHandlerTransactor{contract: contract}, GovernanceHandlerFilterer: GovernanceHandlerFilterer{contract: contract}}, nil
}

// NewGovernanceHandlerCaller creates a new read-only instance of GovernanceHandler, bound to a specific deployed contract.
func NewGovernanceHandlerCaller(address common.Address, caller bind.ContractCaller) (*GovernanceHandlerCaller, error) {
	contract, err := bindGovernanceHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceHandlerCaller{contract: contract}, nil
}

// NewGovernanceHandlerTransactor creates a new write-only instance of GovernanceHandler, bound to a specific deployed contract.
func NewGovernanceHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceHandlerTransactor, error) {
	contract, err := bindGovernanceHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceHandlerTransactor{contract: contract}, nil
}

// NewGovernanceHandlerFilterer creates a new log filterer instance of GovernanceHandler, bound to a specific deployed contract.
func NewGovernanceHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceHandlerFilterer, error) {
	contract, err := bindGovernanceHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceHandlerFilterer{contract: contract}, nil
}

// bindGovernanceHandler binds a generic wrapper to an already deployed contract.
func bindGovernanceHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceHandler *GovernanceHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceHandler.Contract.GovernanceHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceHandler *GovernanceHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceHandler.Contract.GovernanceHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceHandler *GovernanceHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceHandler.Contract.GovernanceHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceHandler *GovernanceHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceHandler *GovernanceHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceHandler *GovernanceHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceHandler.Contract.contract.Transact(opts, method, params...)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address newChildGovernanceManager, address newExitHelper) returns()
func (_GovernanceHandler *GovernanceHandlerTransactor) Initialize(opts *bind.TransactOpts, newChildGovernanceManager common.Address, newExitHelper common.Address) (*types.Transaction, error) {
	return _GovernanceHandler.contract.Transact(opts, "initialize", newChildGovernanceManager, newExitHelper)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address newChildGovernanceManager, address newExitHelper) returns()
func (_GovernanceHandler *GovernanceHandlerSession) Initialize(newChildGovernanceManager common.Address, newExitHelper common.Address) (*types.Transaction, error) {
	return _GovernanceHandler.Contract.Initialize(&_GovernanceHandler.TransactOpts, newChildGovernanceManager, newExitHelper)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address newChildGovernanceManager, address newExitHelper) returns()
func (_GovernanceHandler *GovernanceHandlerTransactorSession) Initialize(newChildGovernanceManager common.Address, newExitHelper common.Address) (*types.Transaction, error) {
	return _GovernanceHandler.Contract.Initialize(&_GovernanceHandler.TransactOpts, newChildGovernanceManager, newExitHelper)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 , address sender, bytes data) returns()
func (_GovernanceHandler *GovernanceHandlerTransactor) OnL2StateReceive(opts *bind.TransactOpts, arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _GovernanceHandler.contract.Transact(opts, "onL2StateReceive", arg0, sender, data)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 , address sender, bytes data) returns()
func (_GovernanceHandler *GovernanceHandlerSession) OnL2StateReceive(arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _GovernanceHandler.Contract.OnL2StateReceive(&_GovernanceHandler.TransactOpts, arg0, sender, data)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 , address sender, bytes data) returns()
func (_GovernanceHandler *GovernanceHandlerTransactorSession) OnL2StateReceive(arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _GovernanceHandler.Contract.OnL2StateReceive(&_GovernanceHandler.TransactOpts, arg0, sender, data)
}

// GovernanceHandlerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GovernanceHandler contract.
type GovernanceHandlerInitializedIterator struct {
	Event *GovernanceHandlerInitialized // Event containing the contract specifics and raw log

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
func (it *GovernanceHandlerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceHandlerInitialized)
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
		it.Event = new(GovernanceHandlerInitialized)
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
func (it *GovernanceHandlerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceHandlerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceHandlerInitialized represents a Initialized event raised by the GovernanceHandler contract.
type GovernanceHandlerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GovernanceHandler *GovernanceHandlerFilterer) FilterInitialized(opts *bind.FilterOpts) (*GovernanceHandlerInitializedIterator, error) {

	logs, sub, err := _GovernanceHandler.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GovernanceHandlerInitializedIterator{contract: _GovernanceHandler.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GovernanceHandler *GovernanceHandlerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GovernanceHandlerInitialized) (event.Subscription, error) {

	logs, sub, err := _GovernanceHandler.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceHandlerInitialized)
				if err := _GovernanceHandler.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GovernanceHandler *GovernanceHandlerFilterer) ParseInitialized(log types.Log) (*GovernanceHandlerInitialized, error) {
	event := new(GovernanceHandlerInitialized)
	if err := _GovernanceHandler.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceHandlerProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the GovernanceHandler contract.
type GovernanceHandlerProposalExecutedIterator struct {
	Event *GovernanceHandlerProposalExecuted // Event containing the contract specifics and raw log

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
func (it *GovernanceHandlerProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceHandlerProposalExecuted)
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
		it.Event = new(GovernanceHandlerProposalExecuted)
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
func (it *GovernanceHandlerProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceHandlerProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceHandlerProposalExecuted represents a ProposalExecuted event raised by the GovernanceHandler contract.
type GovernanceHandlerProposalExecuted struct {
	Targets     []common.Address
	Values      []*big.Int
	Calldatas   [][]byte
	Description [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0xd228733f9d7ff127e0dca16765efb34ebdc4e2dfe78b3e12aea94aeb0ed210b3.
//
// Solidity: event ProposalExecuted(address[] targets, uint256[] values, bytes[] calldatas, bytes32 description)
func (_GovernanceHandler *GovernanceHandlerFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*GovernanceHandlerProposalExecutedIterator, error) {

	logs, sub, err := _GovernanceHandler.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &GovernanceHandlerProposalExecutedIterator{contract: _GovernanceHandler.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0xd228733f9d7ff127e0dca16765efb34ebdc4e2dfe78b3e12aea94aeb0ed210b3.
//
// Solidity: event ProposalExecuted(address[] targets, uint256[] values, bytes[] calldatas, bytes32 description)
func (_GovernanceHandler *GovernanceHandlerFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *GovernanceHandlerProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _GovernanceHandler.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceHandlerProposalExecuted)
				if err := _GovernanceHandler.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0xd228733f9d7ff127e0dca16765efb34ebdc4e2dfe78b3e12aea94aeb0ed210b3.
//
// Solidity: event ProposalExecuted(address[] targets, uint256[] values, bytes[] calldatas, bytes32 description)
func (_GovernanceHandler *GovernanceHandlerFilterer) ParseProposalExecuted(log types.Log) (*GovernanceHandlerProposalExecuted, error) {
	event := new(GovernanceHandlerProposalExecuted)
	if err := _GovernanceHandler.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
