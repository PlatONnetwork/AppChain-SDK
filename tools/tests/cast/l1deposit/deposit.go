// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l1deposit

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
const DepositABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSIT_SIG\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_SIG\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRegistryManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newStateSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newChildDepositHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newChildChainHelper\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onL2StateReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registryManager\",\"outputs\":[{\"internalType\":\"contractIRegistryManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// DEPOSITSIG is a free data retrieval call binding the contract method 0xd41f1771.
//
// Solidity: function DEPOSIT_SIG() view returns(bytes32)
func (_Deposit *DepositCaller) DEPOSITSIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Deposit.contract.Call(opts, &out, "DEPOSIT_SIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITSIG is a free data retrieval call binding the contract method 0xd41f1771.
//
// Solidity: function DEPOSIT_SIG() view returns(bytes32)
func (_Deposit *DepositSession) DEPOSITSIG() ([32]byte, error) {
	return _Deposit.Contract.DEPOSITSIG(&_Deposit.CallOpts)
}

// DEPOSITSIG is a free data retrieval call binding the contract method 0xd41f1771.
//
// Solidity: function DEPOSIT_SIG() view returns(bytes32)
func (_Deposit *DepositCallerSession) DEPOSITSIG() ([32]byte, error) {
	return _Deposit.Contract.DEPOSITSIG(&_Deposit.CallOpts)
}

// WITHDRAWSIG is a free data retrieval call binding the contract method 0xb1768065.
//
// Solidity: function WITHDRAW_SIG() view returns(bytes32)
func (_Deposit *DepositCaller) WITHDRAWSIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Deposit.contract.Call(opts, &out, "WITHDRAW_SIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWSIG is a free data retrieval call binding the contract method 0xb1768065.
//
// Solidity: function WITHDRAW_SIG() view returns(bytes32)
func (_Deposit *DepositSession) WITHDRAWSIG() ([32]byte, error) {
	return _Deposit.Contract.WITHDRAWSIG(&_Deposit.CallOpts)
}

// WITHDRAWSIG is a free data retrieval call binding the contract method 0xb1768065.
//
// Solidity: function WITHDRAW_SIG() view returns(bytes32)
func (_Deposit *DepositCallerSession) WITHDRAWSIG() ([32]byte, error) {
	return _Deposit.Contract.WITHDRAWSIG(&_Deposit.CallOpts)
}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_Deposit *DepositCaller) RegistryManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deposit.contract.Call(opts, &out, "registryManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_Deposit *DepositSession) RegistryManager() (common.Address, error) {
	return _Deposit.Contract.RegistryManager(&_Deposit.CallOpts)
}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_Deposit *DepositCallerSession) RegistryManager() (common.Address, error) {
	return _Deposit.Contract.RegistryManager(&_Deposit.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address recipient, uint256 amount) returns()
func (_Deposit *DepositTransactor) Deposit(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "deposit", recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address recipient, uint256 amount) returns()
func (_Deposit *DepositSession) Deposit(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Deposit.Contract.Deposit(&_Deposit.TransactOpts, recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address recipient, uint256 amount) returns()
func (_Deposit *DepositTransactorSession) Deposit(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Deposit.Contract.Deposit(&_Deposit.TransactOpts, recipient, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address newRegistryManager, address newStateSender, address newChildDepositHandler, address newChildChainHelper) returns()
func (_Deposit *DepositTransactor) Initialize(opts *bind.TransactOpts, newRegistryManager common.Address, newStateSender common.Address, newChildDepositHandler common.Address, newChildChainHelper common.Address) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "initialize", newRegistryManager, newStateSender, newChildDepositHandler, newChildChainHelper)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address newRegistryManager, address newStateSender, address newChildDepositHandler, address newChildChainHelper) returns()
func (_Deposit *DepositSession) Initialize(newRegistryManager common.Address, newStateSender common.Address, newChildDepositHandler common.Address, newChildChainHelper common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.Initialize(&_Deposit.TransactOpts, newRegistryManager, newStateSender, newChildDepositHandler, newChildChainHelper)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address newRegistryManager, address newStateSender, address newChildDepositHandler, address newChildChainHelper) returns()
func (_Deposit *DepositTransactorSession) Initialize(newRegistryManager common.Address, newStateSender common.Address, newChildDepositHandler common.Address, newChildChainHelper common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.Initialize(&_Deposit.TransactOpts, newRegistryManager, newStateSender, newChildDepositHandler, newChildChainHelper)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 , address sender, bytes data) returns()
func (_Deposit *DepositTransactor) OnL2StateReceive(opts *bind.TransactOpts, arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "onL2StateReceive", arg0, sender, data)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 , address sender, bytes data) returns()
func (_Deposit *DepositSession) OnL2StateReceive(arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _Deposit.Contract.OnL2StateReceive(&_Deposit.TransactOpts, arg0, sender, data)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 , address sender, bytes data) returns()
func (_Deposit *DepositTransactorSession) OnL2StateReceive(arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _Deposit.Contract.OnL2StateReceive(&_Deposit.TransactOpts, arg0, sender, data)
}

// DepositInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Deposit contract.
type DepositInitializedIterator struct {
	Event *DepositInitialized // Event containing the contract specifics and raw log

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
func (it *DepositInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositInitialized)
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
		it.Event = new(DepositInitialized)
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
func (it *DepositInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositInitialized represents a Initialized event raised by the Deposit contract.
type DepositInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Deposit *DepositFilterer) FilterInitialized(opts *bind.FilterOpts) (*DepositInitializedIterator, error) {

	logs, sub, err := _Deposit.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DepositInitializedIterator{contract: _Deposit.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Deposit *DepositFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DepositInitialized) (event.Subscription, error) {

	logs, sub, err := _Deposit.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositInitialized)
				if err := _Deposit.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Deposit *DepositFilterer) ParseInitialized(log types.Log) (*DepositInitialized, error) {
	event := new(DepositInitialized)
	if err := _Deposit.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTokenDepositIterator is returned from FilterTokenDeposit and is used to iterate over the raw logs and unpacked data for TokenDeposit events raised by the Deposit contract.
type DepositTokenDepositIterator struct {
	Event *DepositTokenDeposit // Event containing the contract specifics and raw log

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
func (it *DepositTokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTokenDeposit)
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
		it.Event = new(DepositTokenDeposit)
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
func (it *DepositTokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTokenDeposit represents a TokenDeposit event raised by the Deposit contract.
type DepositTokenDeposit struct {
	Token     common.Address
	Recipient common.Address
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenDeposit is a free log retrieval operation binding the contract event 0xedde939981675ab6f1bccc178784c22dde9c3da9c409aaa5c222157f656f9a05.
//
// Solidity: event TokenDeposit(address indexed token, address indexed recipient, address depositor, uint256 amount)
func (_Deposit *DepositFilterer) FilterTokenDeposit(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*DepositTokenDepositIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Deposit.contract.FilterLogs(opts, "TokenDeposit", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &DepositTokenDepositIterator{contract: _Deposit.contract, event: "TokenDeposit", logs: logs, sub: sub}, nil
}

// WatchTokenDeposit is a free log subscription operation binding the contract event 0xedde939981675ab6f1bccc178784c22dde9c3da9c409aaa5c222157f656f9a05.
//
// Solidity: event TokenDeposit(address indexed token, address indexed recipient, address depositor, uint256 amount)
func (_Deposit *DepositFilterer) WatchTokenDeposit(opts *bind.WatchOpts, sink chan<- *DepositTokenDeposit, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Deposit.contract.WatchLogs(opts, "TokenDeposit", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTokenDeposit)
				if err := _Deposit.contract.UnpackLog(event, "TokenDeposit", log); err != nil {
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

// ParseTokenDeposit is a log parse operation binding the contract event 0xedde939981675ab6f1bccc178784c22dde9c3da9c409aaa5c222157f656f9a05.
//
// Solidity: event TokenDeposit(address indexed token, address indexed recipient, address depositor, uint256 amount)
func (_Deposit *DepositFilterer) ParseTokenDeposit(log types.Log) (*DepositTokenDeposit, error) {
	event := new(DepositTokenDeposit)
	if err := _Deposit.contract.UnpackLog(event, "TokenDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the Deposit contract.
type DepositTokenWithdrawIterator struct {
	Event *DepositTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *DepositTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositTokenWithdraw)
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
		it.Event = new(DepositTokenWithdraw)
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
func (it *DepositTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositTokenWithdraw represents a TokenWithdraw event raised by the Deposit contract.
type DepositTokenWithdraw struct {
	Token      common.Address
	Recipient  common.Address
	Withdrawer common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x0cc8e26c5fe2f82346d6755e53b81d7301a600a4ada570aa7281c99f1c983d5c.
//
// Solidity: event TokenWithdraw(address indexed token, address indexed recipient, address withdrawer, uint256 amount)
func (_Deposit *DepositFilterer) FilterTokenWithdraw(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*DepositTokenWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Deposit.contract.FilterLogs(opts, "TokenWithdraw", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &DepositTokenWithdrawIterator{contract: _Deposit.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x0cc8e26c5fe2f82346d6755e53b81d7301a600a4ada570aa7281c99f1c983d5c.
//
// Solidity: event TokenWithdraw(address indexed token, address indexed recipient, address withdrawer, uint256 amount)
func (_Deposit *DepositFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *DepositTokenWithdraw, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Deposit.contract.WatchLogs(opts, "TokenWithdraw", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositTokenWithdraw)
				if err := _Deposit.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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

// ParseTokenWithdraw is a log parse operation binding the contract event 0x0cc8e26c5fe2f82346d6755e53b81d7301a600a4ada570aa7281c99f1c983d5c.
//
// Solidity: event TokenWithdraw(address indexed token, address indexed recipient, address withdrawer, uint256 amount)
func (_Deposit *DepositFilterer) ParseTokenWithdraw(log types.Log) (*DepositTokenWithdraw, error) {
	event := new(DepositTokenWithdraw)
	if err := _Deposit.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
