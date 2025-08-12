// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package logicproxy

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

// LogicProxyMetaData contains all meta data concerning the LogicProxy contract.
var LogicProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"implementationContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"changeAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Received\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60803461011557601f6106ae38819003918201601f19168301916001600160401b038311848410176101195780849260209460405283398101031261011557516001600160a01b038116810361011557803b156100aa577f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c355337f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b55604051610580908161012e8239f35b60405162461bcd60e51b815260206004820152603b60248201527f43616e6e6f742073657420612070726f787920696d706c656d656e746174696f60448201527f6e20746f2061206e6f6e2d636f6e7472616374206164647265737300000000006064820152608490fd5b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe6080604052600436101561001d575b366103a05761001b61036b565b005b5f3560e01c80633659cfe61461006c5780634f1ef286146100675780635c60da1b146100625780638f2839701461005d5763f851a4400361000e576102e9565b6101f6565b6101af565b6100cd565b346100b35760203660031901126100b3576100856100b7565b5f8051602061052b83398151915254336001600160a01b03909116036100ae5761001b9061045c565b6103a0565b5f80fd5b600435906001600160a01b03821682036100b357565b60403660031901126100b3576100e16100b7565b60243567ffffffffffffffff918282116100b357366023830112156100b3578160040135908382116100b35736602483850101116100b3575f8051602061052b83398151915254336001600160a01b03909116036100ae575f92839261014860249361045c565b806040519384930183378101838152039034305af13d156101a5573d8281116101a057601f199060405191603f81601f8401160116820193828510908511176101a05761001b9360405281525f60203d92013e610331565b61031d565b61001b9150610331565b346100b3575f3660031901126100b3577f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c3546040516001600160a01b039091168152602090f35b346100b35760203660031901126100b35761020f6100b7565b5f8051602061052b83398151915280546001600160a01b0392919033908416036100ae578282169081156102855761001b937f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f9260409254918351921682526020820152a15f8051602061052b83398151915255565b60405162461bcd60e51b815260206004820152603660248201527f43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f604482015275787920746f20746865207a65726f206164647265737360501b6064820152608490fd5b346100b3575f3660031901126100b3575f8051602061052b833981519152546040516001600160a01b039091168152602090f35b634e487b7160e01b5f52604160045260245ffd5b1561033857565b60405162461bcd60e51b815260206004820152600b60248201526a18d85b1b0819985a5b195960aa1b6044820152606490fd5b3461037257565b6040513481527f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f8852587460203392a2565b5f8051602061052b833981519152546001600160a01b031633146103fc575f807f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c354368280378136915af43d5f803e156103f8573d5ff35b3d5ffd5b60405162461bcd60e51b815260206004820152603260248201527f43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e20667260448201527137b6903a343290383937bc3c9030b236b4b760711b6064820152608490fd5b803b156104bf577f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c38190556040516001600160a01b0390911681527fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90602090a1565b60405162461bcd60e51b815260206004820152603b60248201527f43616e6e6f742073657420612070726f787920696d706c656d656e746174696f60448201527f6e20746f2061206e6f6e2d636f6e7472616374206164647265737300000000006064820152608490fdfe10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390ba26469706673582212207f72183e1c384b92cf43c61e5f3b51f2aa013c983096aae08b565eeee8b3af7064736f6c63430008160033",
}

// LogicProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use LogicProxyMetaData.ABI instead.
var LogicProxyABI = LogicProxyMetaData.ABI

// LogicProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LogicProxyMetaData.Bin instead.
var LogicProxyBin = LogicProxyMetaData.Bin

// DeployLogicProxy deploys a new platon contract, binding an instance of LogicProxy to it.
func DeployLogicProxy(auth *bind.TransactOpts, backend bind.ContractBackend, implementationContract common.Address) (common.Address, *types.Transaction, *LogicProxy, error) {
	parsed, err := LogicProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LogicProxyBin), backend, implementationContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LogicProxy{LogicProxyCaller: LogicProxyCaller{contract: contract}, LogicProxyTransactor: LogicProxyTransactor{contract: contract}, LogicProxyFilterer: LogicProxyFilterer{contract: contract}}, nil
}

// LogicProxy is an auto generated Go binding around an platon contract.
type LogicProxy struct {
	LogicProxyCaller     // Read-only binding to the contract
	LogicProxyTransactor // Write-only binding to the contract
	LogicProxyFilterer   // Log filterer for contract events
}

// LogicProxyCaller is an auto generated read-only Go binding around an platon contract.
type LogicProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LogicProxyTransactor is an auto generated write-only Go binding around an platon contract.
type LogicProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LogicProxyFilterer is an auto generated log filtering Go binding around an platon contract events.
type LogicProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LogicProxySession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type LogicProxySession struct {
	Contract     *LogicProxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LogicProxyCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type LogicProxyCallerSession struct {
	Contract *LogicProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LogicProxyTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type LogicProxyTransactorSession struct {
	Contract     *LogicProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LogicProxyRaw is an auto generated low-level Go binding around an platon contract.
type LogicProxyRaw struct {
	Contract *LogicProxy // Generic contract binding to access the raw methods on
}

// LogicProxyCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type LogicProxyCallerRaw struct {
	Contract *LogicProxyCaller // Generic read-only contract binding to access the raw methods on
}

// LogicProxyTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type LogicProxyTransactorRaw struct {
	Contract *LogicProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLogicProxy creates a new instance of LogicProxy, bound to a specific deployed contract.
func NewLogicProxy(address common.Address, backend bind.ContractBackend) (*LogicProxy, error) {
	contract, err := bindLogicProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LogicProxy{LogicProxyCaller: LogicProxyCaller{contract: contract}, LogicProxyTransactor: LogicProxyTransactor{contract: contract}, LogicProxyFilterer: LogicProxyFilterer{contract: contract}}, nil
}

// NewLogicProxyCaller creates a new read-only instance of LogicProxy, bound to a specific deployed contract.
func NewLogicProxyCaller(address common.Address, caller bind.ContractCaller) (*LogicProxyCaller, error) {
	contract, err := bindLogicProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LogicProxyCaller{contract: contract}, nil
}

// NewLogicProxyTransactor creates a new write-only instance of LogicProxy, bound to a specific deployed contract.
func NewLogicProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*LogicProxyTransactor, error) {
	contract, err := bindLogicProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LogicProxyTransactor{contract: contract}, nil
}

// NewLogicProxyFilterer creates a new log filterer instance of LogicProxy, bound to a specific deployed contract.
func NewLogicProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*LogicProxyFilterer, error) {
	contract, err := bindLogicProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LogicProxyFilterer{contract: contract}, nil
}

// bindLogicProxy binds a generic wrapper to an already deployed contract.
func bindLogicProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LogicProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LogicProxy *LogicProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LogicProxy.Contract.LogicProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LogicProxy *LogicProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogicProxy.Contract.LogicProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LogicProxy *LogicProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LogicProxy.Contract.LogicProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LogicProxy *LogicProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LogicProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LogicProxy *LogicProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogicProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LogicProxy *LogicProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LogicProxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_LogicProxy *LogicProxyCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LogicProxy.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_LogicProxy *LogicProxySession) Admin() (common.Address, error) {
	return _LogicProxy.Contract.Admin(&_LogicProxy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_LogicProxy *LogicProxyCallerSession) Admin() (common.Address, error) {
	return _LogicProxy.Contract.Admin(&_LogicProxy.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LogicProxy *LogicProxyCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LogicProxy.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LogicProxy *LogicProxySession) Implementation() (common.Address, error) {
	return _LogicProxy.Contract.Implementation(&_LogicProxy.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LogicProxy *LogicProxyCallerSession) Implementation() (common.Address, error) {
	return _LogicProxy.Contract.Implementation(&_LogicProxy.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_LogicProxy *LogicProxyTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _LogicProxy.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_LogicProxy *LogicProxySession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _LogicProxy.Contract.ChangeAdmin(&_LogicProxy.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_LogicProxy *LogicProxyTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _LogicProxy.Contract.ChangeAdmin(&_LogicProxy.TransactOpts, newAdmin)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LogicProxy *LogicProxyTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _LogicProxy.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LogicProxy *LogicProxySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LogicProxy.Contract.UpgradeTo(&_LogicProxy.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LogicProxy *LogicProxyTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LogicProxy.Contract.UpgradeTo(&_LogicProxy.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LogicProxy *LogicProxyTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LogicProxy.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LogicProxy *LogicProxySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LogicProxy.Contract.UpgradeToAndCall(&_LogicProxy.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LogicProxy *LogicProxyTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LogicProxy.Contract.UpgradeToAndCall(&_LogicProxy.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_LogicProxy *LogicProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _LogicProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_LogicProxy *LogicProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _LogicProxy.Contract.Fallback(&_LogicProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_LogicProxy *LogicProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _LogicProxy.Contract.Fallback(&_LogicProxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LogicProxy *LogicProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogicProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LogicProxy *LogicProxySession) Receive() (*types.Transaction, error) {
	return _LogicProxy.Contract.Receive(&_LogicProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LogicProxy *LogicProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _LogicProxy.Contract.Receive(&_LogicProxy.TransactOpts)
}

// LogicProxyAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the LogicProxy contract.
type LogicProxyAdminChangedIterator struct {
	Event *LogicProxyAdminChanged // Event containing the contract specifics and raw log

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
func (it *LogicProxyAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogicProxyAdminChanged)
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
		it.Event = new(LogicProxyAdminChanged)
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
func (it *LogicProxyAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogicProxyAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogicProxyAdminChanged represents a AdminChanged event raised by the LogicProxy contract.
type LogicProxyAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LogicProxy *LogicProxyFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*LogicProxyAdminChangedIterator, error) {

	logs, sub, err := _LogicProxy.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &LogicProxyAdminChangedIterator{contract: _LogicProxy.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LogicProxy *LogicProxyFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *LogicProxyAdminChanged) (event.Subscription, error) {

	logs, sub, err := _LogicProxy.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogicProxyAdminChanged)
				if err := _LogicProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LogicProxy *LogicProxyFilterer) ParseAdminChanged(log types.Log) (*LogicProxyAdminChanged, error) {
	event := new(LogicProxyAdminChanged)
	if err := _LogicProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LogicProxyReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the LogicProxy contract.
type LogicProxyReceivedIterator struct {
	Event *LogicProxyReceived // Event containing the contract specifics and raw log

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
func (it *LogicProxyReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogicProxyReceived)
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
		it.Event = new(LogicProxyReceived)
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
func (it *LogicProxyReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogicProxyReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogicProxyReceived represents a Received event raised by the LogicProxy contract.
type LogicProxyReceived struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed sender, uint256 value)
func (_LogicProxy *LogicProxyFilterer) FilterReceived(opts *bind.FilterOpts, sender []common.Address) (*LogicProxyReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LogicProxy.contract.FilterLogs(opts, "Received", senderRule)
	if err != nil {
		return nil, err
	}
	return &LogicProxyReceivedIterator{contract: _LogicProxy.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed sender, uint256 value)
func (_LogicProxy *LogicProxyFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *LogicProxyReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LogicProxy.contract.WatchLogs(opts, "Received", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogicProxyReceived)
				if err := _LogicProxy.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed sender, uint256 value)
func (_LogicProxy *LogicProxyFilterer) ParseReceived(log types.Log) (*LogicProxyReceived, error) {
	event := new(LogicProxyReceived)
	if err := _LogicProxy.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LogicProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the LogicProxy contract.
type LogicProxyUpgradedIterator struct {
	Event *LogicProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *LogicProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogicProxyUpgraded)
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
		it.Event = new(LogicProxyUpgraded)
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
func (it *LogicProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogicProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogicProxyUpgraded represents a Upgraded event raised by the LogicProxy contract.
type LogicProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address implementation)
func (_LogicProxy *LogicProxyFilterer) FilterUpgraded(opts *bind.FilterOpts) (*LogicProxyUpgradedIterator, error) {

	logs, sub, err := _LogicProxy.contract.FilterLogs(opts, "Upgraded")
	if err != nil {
		return nil, err
	}
	return &LogicProxyUpgradedIterator{contract: _LogicProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address implementation)
func (_LogicProxy *LogicProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LogicProxyUpgraded) (event.Subscription, error) {

	logs, sub, err := _LogicProxy.contract.WatchLogs(opts, "Upgraded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogicProxyUpgraded)
				if err := _LogicProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address implementation)
func (_LogicProxy *LogicProxyFilterer) ParseUpgraded(log types.Log) (*LogicProxyUpgraded, error) {
	event := new(LogicProxyUpgraded)
	if err := _LogicProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
