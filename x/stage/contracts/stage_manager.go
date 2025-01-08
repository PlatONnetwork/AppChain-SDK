package contracts

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"math/big"
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

// PeriodEdge is an auto generated low-level Go binding around an user-defined struct.
type PeriodEdge struct {
	StartBlock *big.Int
	EndBlock   *big.Int
}

var (
	ABI    = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"periodType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPeriodByBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"periodType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"getPeriodEdge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"internalType\":\"structPeriodEdge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"periodType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"getPeriodEdges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"internalType\":\"structPeriodEdge[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
	Abi, _ = abi.JSON(strings.NewReader(ABI))
)

func (c *StageManager) Run(input []byte) (ret []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
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
	return entry(input[4:])
}

func (c *StageManager) initABI() {
	V0 := uint64(0)
	c.abis[V0] = &Abi
}

func (c *StageManager) initMethodEntry() {

	methodEntry := map[string]func([]byte) ([]byte, error){
		"f466e3d3": c.GetPeriodByBlockNumberEntry,
		"ccb9ce96": c.GetPeriodEdgeEntry,
		"bff4d5ab": c.GetPeriodEdgesEntry,
	}
	V0 := uint64(0)
	c.methodEntries[V0] = methodEntry

}
func (c *StageManager) loadMethodABI() error {
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
func (c *StageManager) InitGenesis(blockNumber uint64) {
	c.SetCreateBlock(blockNumber)
	c.stateDb.SetCode(c.contract.Address(), []byte("code"))
}
func (c *StageManager) SetCreateBlock(blockNumber uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], blockNumber)
	c.evm.StateDB.SetState(c.contract.Address(), createBlockKey, data[:])
}

func (c *StageManager) GetCreateBlock() uint64 {
	blockNumber := c.evm.StateDB.GetState(c.contract.Address(), createBlockKey)
	if len(blockNumber) == 0 {
		return math.MaxUint64
	}
	return binary.BigEndian.Uint64(blockNumber)
}

func (c *StageManager) SetVersion(version uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], version)
	c.evm.StateDB.SetState(c.contract.Address(), versionKey, data[:])
}

func (c *StageManager) GetVersion() uint64 {
	version := c.evm.StateDB.GetState(c.contract.Address(), versionKey)
	if len(version) == 0 {
		return 0
	}
	return binary.BigEndian.Uint64(version)
}

func (c *StageManager) GetPeriodByBlockNumberEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["getPeriodByBlockNumber"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.GetPeriodByBlockNumber(*abi.ConvertType(args[0], new(uint8)).(*uint8), *abi.ConvertType(args[1], new(*big.Int)).(**big.Int))
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

func (c *StageManager) GetPeriodEdgeEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["getPeriodEdge"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.GetPeriodEdge(*abi.ConvertType(args[0], new(uint8)).(*uint8), *abi.ConvertType(args[1], new(*big.Int)).(**big.Int))
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

func (c *StageManager) GetPeriodEdgesEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["getPeriodEdges"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, res1, res2, err := c.GetPeriodEdges(*abi.ConvertType(args[0], new(uint8)).(*uint8), *abi.ConvertType(args[1], new(*big.Int)).(**big.Int), *abi.ConvertType(args[2], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0, res1, res2)
	if err != nil {
		return nil, err
	}

	return output, err
}
