package contracts

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"math/big"
	"runtime/debug"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_              = vm.EVM{}
	_              = errors.New
	_              = big.NewInt
	_              = strings.NewReader
	_              = platon.NotFound
	_              = bind.Bind
	_              = common.Big1
	_              = math.ReadBits
	_              = binary.BigEndian
	_              = types.BloomLookup
	_              = event.NewSubscription
	versionKey     = []byte("__version")
	createBlockKey = []byte("__createBlock")
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
	Signature    []byte
	ValidatorSet BitArray
}

// StateSync is an auto generated low-level Go binding around an user-defined struct.
type StateSync struct {
	Id       *big.Int
	Sender   common.Address
	Receiver common.Address
	Data     []byte
}

// StateSyncCommitment is an auto generated low-level Go binding around an user-defined struct.
type StateSyncCommitment struct {
	StartId *big.Int
	EndId   *big.Int
	Root    common.Hash
}

var (
	ABI    = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"NewCommitment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"StateSyncResult\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32[][]\",\"name\":\"proofs\",\"type\":\"bytes32[][]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structStateSync[]\",\"name\":\"objs\",\"type\":\"tuple[]\"}],\"name\":\"batchExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structStateSyncCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"viewNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"blockIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"extendHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"bits\",\"type\":\"uint32\"},{\"internalType\":\"uint64[]\",\"name\":\"Elems\",\"type\":\"uint64[]\"}],\"internalType\":\"structBitArray\",\"name\":\"validatorSet\",\"type\":\"tuple\"}],\"internalType\":\"structQuorumCert\",\"name\":\"qc\",\"type\":\"tuple\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structStateSync\",\"name\":\"obj\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getCommitmentByStateSyncId\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structStateSyncCommitment\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutedId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getRootByStateSyncId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateSyncId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	Abi, _ = abi.JSON(strings.NewReader(ABI))
)

func (c *StateReceiver) Run(input []byte) (ret []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Trace(string(debug.Stack()))
			switch e := r.(type) {
			case error:
				if r, ok := e.(*typesdk.RevertError); ok {
					ret, err = r.ReturnData, vm.ErrExecutionReverted
				} else {
					ret, err = nil, e
				}
			default:
				ret, err = typesdk.UndefinedError, vm.ErrExecutionReverted
			}
		}
	}()
	if err := c.loadMethodABI(); err != nil {
		return nil, errors.New("load version failed")
	}

	if len(input) < 4 {
		return nil, errors.New("input too short")
	}
	id := input[0:4]
	entry, ok := c.methodEntry[hex.EncodeToString(id)]
	if !ok {
		if c.fallback != nil {
			return c.fallback(input)
		}
		return nil, errors.New("methods not found")
	}
	ret, err = entry(input[4:])
	if err != nil {
		log.Trace("Execute failed", "err", err)
	}
	return ret, err
}

func (c *StateReceiver) initABI() {
	V0 := uint64(0)
	c.abis[V0] = &Abi
}

func (c *StateReceiver) initMethodEntry() {

	methodEntry := map[string]func([]byte) ([]byte, error){
		"eb70ef44": c.GetCommitmentByStateSyncIdEntry,
		"d46904ec": c.GetExecutedIdEntry,
		"196f1b2d": c.GetRootByStateSyncIdEntry,

		"9017c127": c.BatchExecuteEntry,
		"22a704a8": c.CommitEntry,
		"50d5b95b": c.ExecuteEntry,
		"d1673d87": c.GetStateSyncIdEntry,
	}
	V0 := uint64(0)
	c.methodEntries[V0] = methodEntry

}
func (c *StateReceiver) loadMethodABI() error {
	version := c.GetVersion()
	entries, ok := c.methodEntries[version]
	if !ok {
		return errors.New("unknown version")
	}
	c.methodEntry = entries
	abi, ok := c.abis[version]
	if !ok {
		return errors.New("unknown version")
	}
	c.abi = abi
	return nil
}
func (c *StateReceiver) InitGenesis(blockNumber uint64) {
	c.SetCreateBlock(blockNumber)
	c.stateDb.SetCode(c.contract.Address(), []byte("code"))
}
func (c *StateReceiver) SetCreateBlock(blockNumber uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], blockNumber)
	c.evm.StateDB.SetState(c.contract.Address(), createBlockKey, data[:])
}

func (c *StateReceiver) GetCreateBlock() uint64 {
	blockNumber := c.evm.StateDB.GetState(c.contract.Address(), createBlockKey)
	if len(blockNumber) == 0 {
		return math.MaxUint64
	}
	return binary.BigEndian.Uint64(blockNumber)
}

func (c *StateReceiver) SetVersion(version uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], version)
	c.evm.StateDB.SetState(c.contract.Address(), versionKey, data[:])
}

func (c *StateReceiver) GetVersion() uint64 {
	version := c.evm.StateDB.GetState(c.contract.Address(), versionKey)
	if len(version) == 0 {
		return 0
	}
	return binary.BigEndian.Uint64(version)
}

func (c *StateReceiver) GetCommitmentByStateSyncIdEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["getCommitmentByStateSyncId"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.GetCommitmentByStateSyncId(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *StateReceiver) GetExecutedIdEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["getExecutedId"]

	var err error

	res0, err := c.GetExecutedId()
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *StateReceiver) GetRootByStateSyncIdEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["getRootByStateSyncId"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.GetRootByStateSyncId(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *StateReceiver) BatchExecuteEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["batchExecute"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.BatchExecute(*abi.ConvertType(args[0], new([][]common.Hash)).(*[][]common.Hash), *abi.ConvertType(args[1], new([]StateSync)).(*[]StateSync))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *StateReceiver) CommitEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["commit"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.Commit(*abi.ConvertType(args[0], new(StateSyncCommitment)).(*StateSyncCommitment), *abi.ConvertType(args[1], new(uint64)).(*uint64), *abi.ConvertType(args[2], new([]common.Hash)).(*[]common.Hash), *abi.ConvertType(args[3], new(QuorumCert)).(*QuorumCert))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *StateReceiver) ExecuteEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["execute"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.Execute(*abi.ConvertType(args[0], new([]common.Hash)).(*[]common.Hash), *abi.ConvertType(args[1], new(StateSync)).(*StateSync))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *StateReceiver) GetStateSyncIdEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["getStateSyncId"]

	var err error

	res0, err := c.GetStateSyncId()
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *StateReceiver) NewCommitmentEvent(startId *big.Int, endId *big.Int, root common.Hash) (*types.Log, error) {
	event := c.abi.Events["NewCommitment"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, startId, endId, root)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, startId, endId, root)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address:     c.contract.Address(),
		Topics:      hashes,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	}, nil
}
func (c *StateReceiver) EmitNewCommitmentEvent(startId *big.Int, endId *big.Int, root common.Hash) {
	log, err := c.NewCommitmentEvent(startId, endId, root)
	contracts.Require(err == nil, "StateReceiver: emit NewCommitment event failed")
	c.stateDb.AddLog(log)
}

func (c *StateReceiver) StateSyncResultEvent(counter *big.Int, status bool, message []byte) (*types.Log, error) {
	event := c.abi.Events["StateSyncResult"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, counter, status, message)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, counter, status, message)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address:     c.contract.Address(),
		Topics:      hashes,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	}, nil
}
func (c *StateReceiver) EmitStateSyncResultEvent(counter *big.Int, status bool, message []byte) {
	log, err := c.StateSyncResultEvent(counter, status, message)
	contracts.Require(err == nil, "StateReceiver: emit StateSyncResult event failed")
	c.stateDb.AddLog(log)
}
