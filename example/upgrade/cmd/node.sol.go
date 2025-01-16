// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// NodeInfo is an auto generated low-level Go binding around an user-defined struct.
type NodeInfo struct {
	Name string
	Host string
	Port uint16
}

// NodeMetaData contains all meta data concerning the Node contract.
var NodeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addNode\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"host\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"port\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delNode\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getNode\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structNodeInfo\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"host\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"port\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AddNode\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"host\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"port\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelNode\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]",
}

// NodeABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeMetaData.ABI instead.
var NodeABI = NodeMetaData.ABI

// Node is an auto generated Go binding around an platon contract.
type Node struct {
	NodeCaller     // Read-only binding to the contract
	NodeTransactor // Write-only binding to the contract
	NodeFilterer   // Log filterer for contract events
}

// NodeCaller is an auto generated read-only Go binding around an platon contract.
type NodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeTransactor is an auto generated write-only Go binding around an platon contract.
type NodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFilterer is an auto generated log filtering Go binding around an platon contract events.
type NodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type NodeSession struct {
	Contract     *Node             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type NodeCallerSession struct {
	Contract *NodeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NodeTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type NodeTransactorSession struct {
	Contract     *NodeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeRaw is an auto generated low-level Go binding around an platon contract.
type NodeRaw struct {
	Contract *Node // Generic contract binding to access the raw methods on
}

// NodeCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type NodeCallerRaw struct {
	Contract *NodeCaller // Generic read-only contract binding to access the raw methods on
}

// NodeTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type NodeTransactorRaw struct {
	Contract *NodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNode creates a new instance of Node, bound to a specific deployed contract.
func NewNode(address common.Address, backend bind.ContractBackend) (*Node, error) {
	contract, err := bindNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// NewNodeCaller creates a new read-only instance of Node, bound to a specific deployed contract.
func NewNodeCaller(address common.Address, caller bind.ContractCaller) (*NodeCaller, error) {
	contract, err := bindNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeCaller{contract: contract}, nil
}

// NewNodeTransactor creates a new write-only instance of Node, bound to a specific deployed contract.
func NewNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeTransactor, error) {
	contract, err := bindNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeTransactor{contract: contract}, nil
}

// NewNodeFilterer creates a new log filterer instance of Node, bound to a specific deployed contract.
func NewNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeFilterer, error) {
	contract, err := bindNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeFilterer{contract: contract}, nil
}

// bindNode binds a generic wrapper to an already deployed contract.
func bindNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.NodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.contract.Transact(opts, method, params...)
}

// GetNode is a free data retrieval call binding the contract method 0x9428522a.
//
// Solidity: function getNode(string name) view returns((string,string,uint16))
func (_Node *NodeCaller) GetNode(opts *bind.CallOpts, name string) (NodeInfo, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getNode", name)

	if err != nil {
		return *new(NodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(NodeInfo)).(*NodeInfo)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x9428522a.
//
// Solidity: function getNode(string name) view returns((string,string,uint16))
func (_Node *NodeSession) GetNode(name string) (NodeInfo, error) {
	return _Node.Contract.GetNode(&_Node.CallOpts, name)
}

// GetNode is a free data retrieval call binding the contract method 0x9428522a.
//
// Solidity: function getNode(string name) view returns((string,string,uint16))
func (_Node *NodeCallerSession) GetNode(name string) (NodeInfo, error) {
	return _Node.Contract.GetNode(&_Node.CallOpts, name)
}

// AddNode is a paid mutator transaction binding the contract method 0xd63df6a2.
//
// Solidity: function addNode(string name, string host, uint16 port) returns()
func (_Node *NodeTransactor) AddNode(opts *bind.TransactOpts, name string, host string, port uint16) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "addNode", name, host, port)
}

// AddNode is a paid mutator transaction binding the contract method 0xd63df6a2.
//
// Solidity: function addNode(string name, string host, uint16 port) returns()
func (_Node *NodeSession) AddNode(name string, host string, port uint16) (*types.Transaction, error) {
	return _Node.Contract.AddNode(&_Node.TransactOpts, name, host, port)
}

// AddNode is a paid mutator transaction binding the contract method 0xd63df6a2.
//
// Solidity: function addNode(string name, string host, uint16 port) returns()
func (_Node *NodeTransactorSession) AddNode(name string, host string, port uint16) (*types.Transaction, error) {
	return _Node.Contract.AddNode(&_Node.TransactOpts, name, host, port)
}

// DelNode is a paid mutator transaction binding the contract method 0x11c90305.
//
// Solidity: function delNode(string name) returns()
func (_Node *NodeTransactor) DelNode(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "delNode", name)
}

// DelNode is a paid mutator transaction binding the contract method 0x11c90305.
//
// Solidity: function delNode(string name) returns()
func (_Node *NodeSession) DelNode(name string) (*types.Transaction, error) {
	return _Node.Contract.DelNode(&_Node.TransactOpts, name)
}

// DelNode is a paid mutator transaction binding the contract method 0x11c90305.
//
// Solidity: function delNode(string name) returns()
func (_Node *NodeTransactorSession) DelNode(name string) (*types.Transaction, error) {
	return _Node.Contract.DelNode(&_Node.TransactOpts, name)
}

// NodeAddNodeIterator is returned from FilterAddNode and is used to iterate over the raw logs and unpacked data for AddNode events raised by the Node contract.
type NodeAddNodeIterator struct {
	Event *NodeAddNode // Event containing the contract specifics and raw log

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
func (it *NodeAddNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeAddNode)
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
		it.Event = new(NodeAddNode)
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
func (it *NodeAddNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeAddNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeAddNode represents a AddNode event raised by the Node contract.
type NodeAddNode struct {
	Name string
	Host string
	Port uint16
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddNode is a free log retrieval operation binding the contract event 0x83cf1b3ad5026dda90efb51c908c90da82cbbe77ce13b60bcd70f7ac48d19836.
//
// Solidity: event AddNode(string name, string host, uint16 port)
func (_Node *NodeFilterer) FilterAddNode(opts *bind.FilterOpts) (*NodeAddNodeIterator, error) {

	logs, sub, err := _Node.contract.FilterLogs(opts, "AddNode")
	if err != nil {
		return nil, err
	}
	return &NodeAddNodeIterator{contract: _Node.contract, event: "AddNode", logs: logs, sub: sub}, nil
}

// WatchAddNode is a free log subscription operation binding the contract event 0x83cf1b3ad5026dda90efb51c908c90da82cbbe77ce13b60bcd70f7ac48d19836.
//
// Solidity: event AddNode(string name, string host, uint16 port)
func (_Node *NodeFilterer) WatchAddNode(opts *bind.WatchOpts, sink chan<- *NodeAddNode) (event.Subscription, error) {

	logs, sub, err := _Node.contract.WatchLogs(opts, "AddNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeAddNode)
				if err := _Node.contract.UnpackLog(event, "AddNode", log); err != nil {
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

// ParseAddNode is a log parse operation binding the contract event 0x83cf1b3ad5026dda90efb51c908c90da82cbbe77ce13b60bcd70f7ac48d19836.
//
// Solidity: event AddNode(string name, string host, uint16 port)
func (_Node *NodeFilterer) ParseAddNode(log types.Log) (*NodeAddNode, error) {
	event := new(NodeAddNode)
	if err := _Node.contract.UnpackLog(event, "AddNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeDelNodeIterator is returned from FilterDelNode and is used to iterate over the raw logs and unpacked data for DelNode events raised by the Node contract.
type NodeDelNodeIterator struct {
	Event *NodeDelNode // Event containing the contract specifics and raw log

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
func (it *NodeDelNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeDelNode)
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
		it.Event = new(NodeDelNode)
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
func (it *NodeDelNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeDelNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeDelNode represents a DelNode event raised by the Node contract.
type NodeDelNode struct {
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDelNode is a free log retrieval operation binding the contract event 0xc26fa984f1d210e6e6b48250428999276a7f94d26fc0349b2eed37b3e29164ec.
//
// Solidity: event DelNode(string name)
func (_Node *NodeFilterer) FilterDelNode(opts *bind.FilterOpts) (*NodeDelNodeIterator, error) {

	logs, sub, err := _Node.contract.FilterLogs(opts, "DelNode")
	if err != nil {
		return nil, err
	}
	return &NodeDelNodeIterator{contract: _Node.contract, event: "DelNode", logs: logs, sub: sub}, nil
}

// WatchDelNode is a free log subscription operation binding the contract event 0xc26fa984f1d210e6e6b48250428999276a7f94d26fc0349b2eed37b3e29164ec.
//
// Solidity: event DelNode(string name)
func (_Node *NodeFilterer) WatchDelNode(opts *bind.WatchOpts, sink chan<- *NodeDelNode) (event.Subscription, error) {

	logs, sub, err := _Node.contract.WatchLogs(opts, "DelNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeDelNode)
				if err := _Node.contract.UnpackLog(event, "DelNode", log); err != nil {
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

// ParseDelNode is a log parse operation binding the contract event 0xc26fa984f1d210e6e6b48250428999276a7f94d26fc0349b2eed37b3e29164ec.
//
// Solidity: event DelNode(string name)
func (_Node *NodeFilterer) ParseDelNode(log types.Log) (*NodeDelNode, error) {
	event := new(NodeDelNode)
	if err := _Node.contract.UnpackLog(event, "DelNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

