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
	ABIV2    = "[{\"type\":\"function\",\"name\":\"addNode\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"host\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"port\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getNode\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structNodeInfo\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"host\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"port\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"}]"
	AbiV2, _ = abi.JSON(strings.NewReader(ABIV2))
)

// NodeInfo is an auto generated low-level Go binding around an user-defined struct.
type NodeInfo struct {
	Name string
	Host string
	Port uint16
}

func (c *Node) initABIV2() {
	V2 := uint64(2)
	c.abis[V2] = &AbiV2
}

func (c *Node) initMethodV2Entry() {
	methodEntry := map[string]func([]byte) ([]byte, error){

		"9428522a": c.GetNodeEntry,
		"d63df6a2": c.AddNodeEntry,
	}
	V2 := uint64(2)
	c.methodEntries[V2] = methodEntry
}

func (c *Node) GetNodeEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["getNode"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.GetNode(*abi.ConvertType(args[0], new(string)).(*string))
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
func (c *Node) GetNode(name string) (NodeInfo, error) {
	var nodeInfo NodeInfo
	if node := c.storage.Nodes.MustGet(name); node != nil {
		nodeInfo = NodeInfo{
			Name: name,
			Host: node.Host,
			Port: node.port,
		}
	}
	return nodeInfo, nil
}
