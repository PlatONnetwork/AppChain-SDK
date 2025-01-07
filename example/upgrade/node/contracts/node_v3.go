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
	ABIV3    = "[{\"type\":\"function\",\"name\":\"addNode\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"host\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"port\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delNode\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getNode\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structNodeInfo\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"host\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"port\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AddNode\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"host\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"port\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelNode\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]"
	AbiV3, _ = abi.JSON(strings.NewReader(ABIV3))
)

func (c *Node) initABIV3() {
	V3 := uint64(3)
	c.abis[V3] = &AbiV3
}

func (c *Node) initMethodV3Entry() {
	methodEntry := map[string]func([]byte) ([]byte, error){

		"11c90305": c.DelNodeEntry,
		"9428522a": c.GetNodeEntry,
		"d63df6a2": c.AddNodeEntry,
	}
	V3 := uint64(3)
	c.methodEntries[V3] = methodEntry
}

func (c *Node) DelNodeEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["delNode"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.DelNode(*abi.ConvertType(args[0], new(string)).(*string))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}
func (c *Node) DelNode(name string) error {
	c.storage.Nodes.MustSet(name, nil)
	c.EmitDelNodeEvent(name)
	return nil
}

func (c *Node) DelNodeEvent(name string) (*types.Log, error) {
	event := c.abi.Events["DelNode"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, name)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, name)
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
func (c *Node) EmitDelNodeEvent(name string) {
	log, err := c.DelNodeEvent(name)
	contracts.Require(err == nil, "Node: emit DelNode event failed")
	c.stateDb.AddLog(log)
}
