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
	"math/big"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_              = contracts.Context{}
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

var (
	ABI    = "[{\"type\":\"function\",\"name\":\"addNode\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"host\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"port\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]"
	Abi, _ = abi.JSON(strings.NewReader(ABI))
)

func (c *Node) Run(input []byte) (ret []byte, err error) {
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

func (c *Node) initABI() {
	V0 := uint64(0)
	c.abis[V0] = &Abi
}

func (c *Node) initMethodEntry() {

	methodEntry := map[string]func([]byte) ([]byte, error){

		"d63df6a2": c.AddNodeEntry,
	}
	V0 := uint64(0)
	c.methodEntries[V0] = methodEntry

}
func (c *Node) loadMethodABI() error {
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

func (c *Node) InitGenesis(blockNumber uint64) {
	c.stateDb.SetNonce(c.contract.Address(), 1)
	c.SetCreateBlock(blockNumber)
	c.stateDb.SetCode(c.contract.Address(), []byte("code"))
}

func (c *Node) SetCreateBlock(blockNumber uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], blockNumber)
	c.evm.StateDB.SetState(c.contract.Address(), createBlockKey, data[:])
}

func (c *Node) GetCreateBlock() uint64 {
	blockNumber := c.evm.StateDB.GetState(c.contract.Address(), createBlockKey)
	if len(blockNumber) == 0 {
		return math.MaxUint64
	}
	return binary.BigEndian.Uint64(blockNumber)
}

func (c *Node) SetVersion(version uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], version)
	c.evm.StateDB.SetState(c.contract.Address(), versionKey, data[:])
}

func (c *Node) GetVersion() uint64 {
	version := c.evm.StateDB.GetState(c.contract.Address(), versionKey)
	if len(version) == 0 {
		return 0
	}
	return binary.BigEndian.Uint64(version)
}

func (c *Node) AddNodeEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["addNode"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.AddNode(*abi.ConvertType(args[0], new(string)).(*string), *abi.ConvertType(args[1], new(string)).(*string), *abi.ConvertType(args[2], new(uint16)).(*uint16))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}
