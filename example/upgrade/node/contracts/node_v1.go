package contracts

import (
	"encoding/hex"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"math/big"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = hex.ErrLength
	_ = contracts.Context{}
	_ = vm.EVM{}
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = platon.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

var (
	ABIV1    = "[{\"type\":\"function\",\"name\":\"addNode\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"host\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"port\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AddNode\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"host\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"port\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false}]"
	AbiV1, _ = abi.JSON(strings.NewReader(ABIV1))
)

func (c *Node) initABIV1() {
	V1 := uint64(1)
	c.abis[V1] = &AbiV1
}

func (c *Node) initMethodV1Entry() {
	methodEntry := map[string]func([]byte) ([]byte, error){

		"d63df6a2": c.AddNodeV1Entry,
	}
	V1 := uint64(1)
	c.methodEntries[V1] = methodEntry
}

func (c *Node) AddNodeV1Entry(input []byte) ([]byte, error) {

	method := c.abi.Methods["addNode"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.AddNodeV1(*abi.ConvertType(args[0], new(string)).(*string), *abi.ConvertType(args[1], new(string)).(*string), *abi.ConvertType(args[2], new(uint16)).(*uint16))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}
func (c *Node) AddNodeV1(name string, host string, port uint16) error {
	contracts.Require(len(host) != 0, "Node: invalid host")
	contracts.Require(port != 0, "Node: invalid port")
	c.AddNode(name, host, port)
	if c.GetVersion() > 0 {
		c.EmitAddNodeEvent(name, host, port)
	}
	return nil
}

func (c *Node) AddNodeEvent(name string, host string, port uint16) (*types.Log, error) {
	event := c.abi.Events["AddNode"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, name, host, port)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, name, host, port)
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
func (c *Node) EmitAddNodeEvent(name string, host string, port uint16) {
	log, err := c.AddNodeEvent(name, host, port)
	contracts.Require(err == nil, "Node: emit AddNode event failed")
	c.stateDb.AddLog(log)
}
