// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2statesync

import (
	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
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

// BitArray is an auto generated low-level Go binding around an user-defined struct.
type BitArray struct {
	Bits  uint32
	Elems []uint64
}

// QuorumCert is an auto generated low-level Go binding around an user-defined struct.
type QuorumCert struct {
	Epoch        uint64
	ViewNumber   uint64
	BlockHash    common.Hash
	BlockNumber  uint64
	BlockIndex   uint32
	ExtendHash   common.Hash
	Signature    hexutil.Bytes
	ValidatorSet BitArray
}

// StateSync is an auto generated low-level Go binding around an user-defined struct.
type StateSync struct {
	Id       *big.Int
	Sender   common.Address
	Receiver common.Address
	Data     hexutil.Bytes
}

// StateSyncCommitment is an auto generated low-level Go binding around an user-defined struct.
type StateSyncCommitment struct {
	StartId *big.Int
	EndId   *big.Int
	Root    common.Hash
}

// StatesyncABI is the input ABI used to generate the binding from.
const StatesyncABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"NewCommitment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"StateSyncResult\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32[][]\",\"name\":\"proofs\",\"type\":\"bytes32[][]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structStateSync[]\",\"name\":\"objs\",\"type\":\"tuple[]\"}],\"name\":\"batchExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structStateSyncCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"viewNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"blockIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"extendHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"bits\",\"type\":\"uint32\"},{\"internalType\":\"uint64[]\",\"name\":\"Elems\",\"type\":\"uint64[]\"}],\"internalType\":\"structBitArray\",\"name\":\"validatorSet\",\"type\":\"tuple\"}],\"internalType\":\"structQuorumCert\",\"name\":\"qc\",\"type\":\"tuple\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structStateSync\",\"name\":\"obj\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getCommitmentByStateSyncId\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structStateSyncCommitment\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutedId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getRootByStateSyncId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateSyncId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Statesync is an auto generated Go binding around an platon contract.
type Statesync struct {
	StatesyncCaller     // Read-only binding to the contract
	StatesyncTransactor // Write-only binding to the contract
	StatesyncFilterer   // Log filterer for contract events
}

// StatesyncCaller is an auto generated read-only Go binding around an platon contract.
type StatesyncCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatesyncTransactor is an auto generated write-only Go binding around an platon contract.
type StatesyncTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatesyncFilterer is an auto generated log filtering Go binding around an platon contract events.
type StatesyncFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatesyncSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type StatesyncSession struct {
	Contract     *Statesync        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StatesyncCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type StatesyncCallerSession struct {
	Contract *StatesyncCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StatesyncTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type StatesyncTransactorSession struct {
	Contract     *StatesyncTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StatesyncRaw is an auto generated low-level Go binding around an platon contract.
type StatesyncRaw struct {
	Contract *Statesync // Generic contract binding to access the raw methods on
}

// StatesyncCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type StatesyncCallerRaw struct {
	Contract *StatesyncCaller // Generic read-only contract binding to access the raw methods on
}

// StatesyncTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type StatesyncTransactorRaw struct {
	Contract *StatesyncTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStatesync creates a new instance of Statesync, bound to a specific deployed contract.
func NewStatesync(address common.Address, backend bind.ContractBackend) (*Statesync, error) {
	contract, err := bindStatesync(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Statesync{StatesyncCaller: StatesyncCaller{contract: contract}, StatesyncTransactor: StatesyncTransactor{contract: contract}, StatesyncFilterer: StatesyncFilterer{contract: contract}}, nil
}

// NewStatesyncCaller creates a new read-only instance of Statesync, bound to a specific deployed contract.
func NewStatesyncCaller(address common.Address, caller bind.ContractCaller) (*StatesyncCaller, error) {
	contract, err := bindStatesync(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StatesyncCaller{contract: contract}, nil
}

// NewStatesyncTransactor creates a new write-only instance of Statesync, bound to a specific deployed contract.
func NewStatesyncTransactor(address common.Address, transactor bind.ContractTransactor) (*StatesyncTransactor, error) {
	contract, err := bindStatesync(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StatesyncTransactor{contract: contract}, nil
}

// NewStatesyncFilterer creates a new log filterer instance of Statesync, bound to a specific deployed contract.
func NewStatesyncFilterer(address common.Address, filterer bind.ContractFilterer) (*StatesyncFilterer, error) {
	contract, err := bindStatesync(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StatesyncFilterer{contract: contract}, nil
}

// bindStatesync binds a generic wrapper to an already deployed contract.
func bindStatesync(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StatesyncABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Statesync *StatesyncRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Statesync.Contract.StatesyncCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Statesync *StatesyncRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Statesync.Contract.StatesyncTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Statesync *StatesyncRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Statesync.Contract.StatesyncTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Statesync *StatesyncCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Statesync.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Statesync *StatesyncTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Statesync.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Statesync *StatesyncTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Statesync.Contract.contract.Transact(opts, method, params...)
}

// GetCommitmentByStateSyncId is a free data retrieval call binding the contract method 0xeb70ef44.
//
// Solidity: function getCommitmentByStateSyncId(uint256 id) view returns((uint256,uint256,bytes32))
func (_Statesync *StatesyncCaller) GetCommitmentByStateSyncId(opts *bind.CallOpts, id *big.Int) (StateSyncCommitment, error) {
	var out []interface{}
	err := _Statesync.contract.Call(opts, &out, "getCommitmentByStateSyncId", id)

	if err != nil {
		return *new(StateSyncCommitment), err
	}

	out0 := *abi.ConvertType(out[0], new(StateSyncCommitment)).(*StateSyncCommitment)

	return out0, err

}

// GetCommitmentByStateSyncId is a free data retrieval call binding the contract method 0xeb70ef44.
//
// Solidity: function getCommitmentByStateSyncId(uint256 id) view returns((uint256,uint256,bytes32))
func (_Statesync *StatesyncSession) GetCommitmentByStateSyncId(id *big.Int) (StateSyncCommitment, error) {
	return _Statesync.Contract.GetCommitmentByStateSyncId(&_Statesync.CallOpts, id)
}

// GetCommitmentByStateSyncId is a free data retrieval call binding the contract method 0xeb70ef44.
//
// Solidity: function getCommitmentByStateSyncId(uint256 id) view returns((uint256,uint256,bytes32))
func (_Statesync *StatesyncCallerSession) GetCommitmentByStateSyncId(id *big.Int) (StateSyncCommitment, error) {
	return _Statesync.Contract.GetCommitmentByStateSyncId(&_Statesync.CallOpts, id)
}

// GetExecutedId is a free data retrieval call binding the contract method 0xd46904ec.
//
// Solidity: function getExecutedId() view returns(uint256)
func (_Statesync *StatesyncCaller) GetExecutedId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Statesync.contract.Call(opts, &out, "getExecutedId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExecutedId is a free data retrieval call binding the contract method 0xd46904ec.
//
// Solidity: function getExecutedId() view returns(uint256)
func (_Statesync *StatesyncSession) GetExecutedId() (*big.Int, error) {
	return _Statesync.Contract.GetExecutedId(&_Statesync.CallOpts)
}

// GetExecutedId is a free data retrieval call binding the contract method 0xd46904ec.
//
// Solidity: function getExecutedId() view returns(uint256)
func (_Statesync *StatesyncCallerSession) GetExecutedId() (*big.Int, error) {
	return _Statesync.Contract.GetExecutedId(&_Statesync.CallOpts)
}

// GetRootByStateSyncId is a free data retrieval call binding the contract method 0x196f1b2d.
//
// Solidity: function getRootByStateSyncId(uint256 id) view returns(bytes32)
func (_Statesync *StatesyncCaller) GetRootByStateSyncId(opts *bind.CallOpts, id *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Statesync.contract.Call(opts, &out, "getRootByStateSyncId", id)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRootByStateSyncId is a free data retrieval call binding the contract method 0x196f1b2d.
//
// Solidity: function getRootByStateSyncId(uint256 id) view returns(bytes32)
func (_Statesync *StatesyncSession) GetRootByStateSyncId(id *big.Int) ([32]byte, error) {
	return _Statesync.Contract.GetRootByStateSyncId(&_Statesync.CallOpts, id)
}

// GetRootByStateSyncId is a free data retrieval call binding the contract method 0x196f1b2d.
//
// Solidity: function getRootByStateSyncId(uint256 id) view returns(bytes32)
func (_Statesync *StatesyncCallerSession) GetRootByStateSyncId(id *big.Int) ([32]byte, error) {
	return _Statesync.Contract.GetRootByStateSyncId(&_Statesync.CallOpts, id)
}

// GetStateSyncId is a free data retrieval call binding the contract method 0xd1673d87.
//
// Solidity: function getStateSyncId() view returns(uint256)
func (_Statesync *StatesyncCaller) GetStateSyncId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Statesync.contract.Call(opts, &out, "getStateSyncId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStateSyncId is a free data retrieval call binding the contract method 0xd1673d87.
//
// Solidity: function getStateSyncId() view returns(uint256)
func (_Statesync *StatesyncSession) GetStateSyncId() (*big.Int, error) {
	return _Statesync.Contract.GetStateSyncId(&_Statesync.CallOpts)
}

// GetStateSyncId is a free data retrieval call binding the contract method 0xd1673d87.
//
// Solidity: function getStateSyncId() view returns(uint256)
func (_Statesync *StatesyncCallerSession) GetStateSyncId() (*big.Int, error) {
	return _Statesync.Contract.GetStateSyncId(&_Statesync.CallOpts)
}

// BatchExecute is a paid mutator transaction binding the contract method 0x9017c127.
//
// Solidity: function batchExecute(bytes32[][] proofs, (uint256,address,address,bytes)[] objs) returns()
func (_Statesync *StatesyncTransactor) BatchExecute(opts *bind.TransactOpts, proofs [][][32]byte, objs []StateSync) (*types.Transaction, error) {
	return _Statesync.contract.Transact(opts, "batchExecute", proofs, objs)
}

// BatchExecute is a paid mutator transaction binding the contract method 0x9017c127.
//
// Solidity: function batchExecute(bytes32[][] proofs, (uint256,address,address,bytes)[] objs) returns()
func (_Statesync *StatesyncSession) BatchExecute(proofs [][][32]byte, objs []StateSync) (*types.Transaction, error) {
	return _Statesync.Contract.BatchExecute(&_Statesync.TransactOpts, proofs, objs)
}

// BatchExecute is a paid mutator transaction binding the contract method 0x9017c127.
//
// Solidity: function batchExecute(bytes32[][] proofs, (uint256,address,address,bytes)[] objs) returns()
func (_Statesync *StatesyncTransactorSession) BatchExecute(proofs [][][32]byte, objs []StateSync) (*types.Transaction, error) {
	return _Statesync.Contract.BatchExecute(&_Statesync.TransactOpts, proofs, objs)
}

// Commit is a paid mutator transaction binding the contract method 0x22a704a8.
//
// Solidity: function commit((uint256,uint256,bytes32) commitment, uint64 index, bytes32[] proof, (uint64,uint64,bytes32,uint64,uint32,bytes32,bytes,(uint32,uint64[])) qc) returns()
func (_Statesync *StatesyncTransactor) Commit(opts *bind.TransactOpts, commitment StateSyncCommitment, index uint64, proof [][32]byte, qc QuorumCert) (*types.Transaction, error) {
	return _Statesync.contract.Transact(opts, "commit", commitment, index, proof, qc)
}

// Commit is a paid mutator transaction binding the contract method 0x22a704a8.
//
// Solidity: function commit((uint256,uint256,bytes32) commitment, uint64 index, bytes32[] proof, (uint64,uint64,bytes32,uint64,uint32,bytes32,bytes,(uint32,uint64[])) qc) returns()
func (_Statesync *StatesyncSession) Commit(commitment StateSyncCommitment, index uint64, proof [][32]byte, qc QuorumCert) (*types.Transaction, error) {
	return _Statesync.Contract.Commit(&_Statesync.TransactOpts, commitment, index, proof, qc)
}

// Commit is a paid mutator transaction binding the contract method 0x22a704a8.
//
// Solidity: function commit((uint256,uint256,bytes32) commitment, uint64 index, bytes32[] proof, (uint64,uint64,bytes32,uint64,uint32,bytes32,bytes,(uint32,uint64[])) qc) returns()
func (_Statesync *StatesyncTransactorSession) Commit(commitment StateSyncCommitment, index uint64, proof [][32]byte, qc QuorumCert) (*types.Transaction, error) {
	return _Statesync.Contract.Commit(&_Statesync.TransactOpts, commitment, index, proof, qc)
}

// Execute is a paid mutator transaction binding the contract method 0x50d5b95b.
//
// Solidity: function execute(bytes32[] proof, (uint256,address,address,bytes) obj) returns()
func (_Statesync *StatesyncTransactor) Execute(opts *bind.TransactOpts, proof [][32]byte, obj StateSync) (*types.Transaction, error) {
	return _Statesync.contract.Transact(opts, "execute", proof, obj)
}

// Execute is a paid mutator transaction binding the contract method 0x50d5b95b.
//
// Solidity: function execute(bytes32[] proof, (uint256,address,address,bytes) obj) returns()
func (_Statesync *StatesyncSession) Execute(proof [][32]byte, obj StateSync) (*types.Transaction, error) {
	return _Statesync.Contract.Execute(&_Statesync.TransactOpts, proof, obj)
}

// Execute is a paid mutator transaction binding the contract method 0x50d5b95b.
//
// Solidity: function execute(bytes32[] proof, (uint256,address,address,bytes) obj) returns()
func (_Statesync *StatesyncTransactorSession) Execute(proof [][32]byte, obj StateSync) (*types.Transaction, error) {
	return _Statesync.Contract.Execute(&_Statesync.TransactOpts, proof, obj)
}

// StatesyncNewCommitmentIterator is returned from FilterNewCommitment and is used to iterate over the raw logs and unpacked data for NewCommitment events raised by the Statesync contract.
type StatesyncNewCommitmentIterator struct {
	Event *StatesyncNewCommitment // Event containing the contract specifics and raw log

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
func (it *StatesyncNewCommitmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatesyncNewCommitment)
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
		it.Event = new(StatesyncNewCommitment)
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
func (it *StatesyncNewCommitmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatesyncNewCommitmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatesyncNewCommitment represents a NewCommitment event raised by the Statesync contract.
type StatesyncNewCommitment struct {
	StartId *big.Int
	EndId   *big.Int
	Root    common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewCommitment is a free log retrieval operation binding the contract event 0x11efd893530b26afc66d488ff54cb15df117cb6e0e4a08c6dcb166d766c3bf3b.
//
// Solidity: event NewCommitment(uint256 indexed startId, uint256 indexed endId, bytes32 root)
func (_Statesync *StatesyncFilterer) FilterNewCommitment(opts *bind.FilterOpts, startId []*big.Int, endId []*big.Int) (*StatesyncNewCommitmentIterator, error) {

	var startIdRule []interface{}
	for _, startIdItem := range startId {
		startIdRule = append(startIdRule, startIdItem)
	}
	var endIdRule []interface{}
	for _, endIdItem := range endId {
		endIdRule = append(endIdRule, endIdItem)
	}

	logs, sub, err := _Statesync.contract.FilterLogs(opts, "NewCommitment", startIdRule, endIdRule)
	if err != nil {
		return nil, err
	}
	return &StatesyncNewCommitmentIterator{contract: _Statesync.contract, event: "NewCommitment", logs: logs, sub: sub}, nil
}

// WatchNewCommitment is a free log subscription operation binding the contract event 0x11efd893530b26afc66d488ff54cb15df117cb6e0e4a08c6dcb166d766c3bf3b.
//
// Solidity: event NewCommitment(uint256 indexed startId, uint256 indexed endId, bytes32 root)
func (_Statesync *StatesyncFilterer) WatchNewCommitment(opts *bind.WatchOpts, sink chan<- *StatesyncNewCommitment, startId []*big.Int, endId []*big.Int) (event.Subscription, error) {

	var startIdRule []interface{}
	for _, startIdItem := range startId {
		startIdRule = append(startIdRule, startIdItem)
	}
	var endIdRule []interface{}
	for _, endIdItem := range endId {
		endIdRule = append(endIdRule, endIdItem)
	}

	logs, sub, err := _Statesync.contract.WatchLogs(opts, "NewCommitment", startIdRule, endIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatesyncNewCommitment)
				if err := _Statesync.contract.UnpackLog(event, "NewCommitment", log); err != nil {
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

// ParseNewCommitment is a log parse operation binding the contract event 0x11efd893530b26afc66d488ff54cb15df117cb6e0e4a08c6dcb166d766c3bf3b.
//
// Solidity: event NewCommitment(uint256 indexed startId, uint256 indexed endId, bytes32 root)
func (_Statesync *StatesyncFilterer) ParseNewCommitment(log types.Log) (*StatesyncNewCommitment, error) {
	event := new(StatesyncNewCommitment)
	if err := _Statesync.contract.UnpackLog(event, "NewCommitment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatesyncStateSyncResultIterator is returned from FilterStateSyncResult and is used to iterate over the raw logs and unpacked data for StateSyncResult events raised by the Statesync contract.
type StatesyncStateSyncResultIterator struct {
	Event *StatesyncStateSyncResult // Event containing the contract specifics and raw log

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
func (it *StatesyncStateSyncResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatesyncStateSyncResult)
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
		it.Event = new(StatesyncStateSyncResult)
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
func (it *StatesyncStateSyncResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatesyncStateSyncResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatesyncStateSyncResult represents a StateSyncResult event raised by the Statesync contract.
type StatesyncStateSyncResult struct {
	Counter *big.Int
	Status  bool
	Message hexutil.Bytes
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStateSyncResult is a free log retrieval operation binding the contract event 0x31c652130602f3ce96ceaf8a4c2b8b49f049166c6fcf2eb31943a75ec7c936ae.
//
// Solidity: event StateSyncResult(uint256 indexed counter, bool indexed status, bytes message)
func (_Statesync *StatesyncFilterer) FilterStateSyncResult(opts *bind.FilterOpts, counter []*big.Int, status []bool) (*StatesyncStateSyncResultIterator, error) {

	var counterRule []interface{}
	for _, counterItem := range counter {
		counterRule = append(counterRule, counterItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Statesync.contract.FilterLogs(opts, "StateSyncResult", counterRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &StatesyncStateSyncResultIterator{contract: _Statesync.contract, event: "StateSyncResult", logs: logs, sub: sub}, nil
}

// WatchStateSyncResult is a free log subscription operation binding the contract event 0x31c652130602f3ce96ceaf8a4c2b8b49f049166c6fcf2eb31943a75ec7c936ae.
//
// Solidity: event StateSyncResult(uint256 indexed counter, bool indexed status, bytes message)
func (_Statesync *StatesyncFilterer) WatchStateSyncResult(opts *bind.WatchOpts, sink chan<- *StatesyncStateSyncResult, counter []*big.Int, status []bool) (event.Subscription, error) {

	var counterRule []interface{}
	for _, counterItem := range counter {
		counterRule = append(counterRule, counterItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Statesync.contract.WatchLogs(opts, "StateSyncResult", counterRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatesyncStateSyncResult)
				if err := _Statesync.contract.UnpackLog(event, "StateSyncResult", log); err != nil {
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

// ParseStateSyncResult is a log parse operation binding the contract event 0x31c652130602f3ce96ceaf8a4c2b8b49f049166c6fcf2eb31943a75ec7c936ae.
//
// Solidity: event StateSyncResult(uint256 indexed counter, bool indexed status, bytes message)
func (_Statesync *StatesyncFilterer) ParseStateSyncResult(log types.Log) (*StatesyncStateSyncResult, error) {
	event := new(StatesyncStateSyncResult)
	if err := _Statesync.contract.UnpackLog(event, "StateSyncResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
