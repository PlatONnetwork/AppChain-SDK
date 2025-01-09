package contracts

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/db/container"
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
	_ = typesdk.RevertError{}
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
	nodeKey = []byte("node")
)

type Info struct {
	Host string
	port uint16
}

type Storage struct {
	Nodes *container.Map[*Info]
}
type Node struct {
	abi           *abi.ABI
	abis          map[uint64]*abi.ABI
	methodEntry   map[string]func([]byte) ([]byte, error)
	methodEntries map[uint64]map[string]func([]byte) ([]byte, error)
	readOnly      bool
	contract      *vm.Contract
	evm           *vm.EVM
	burner        contracts.Burn
	stateDb       *contracts.StateDB
	context       *contracts.Context
	fallback      func(input []byte) ([]byte, error)
	storage       *Storage
}

func NewNode(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*Node, error) {
	s := &Node{
		abi:           nil,
		abis:          make(map[uint64]*abi.ABI),
		methodEntry:   make(map[string]func([]byte) ([]byte, error)),
		methodEntries: make(map[uint64]map[string]func([]byte) ([]byte, error)),
		evm:           evm,
		contract:      contract,
		burner:        contracts.NewBurner(contract),
		stateDb:       contracts.NewStateDB(evm, contract),
		context:       contracts.NewContext(evm, contract),
		readOnly:      readOnly,
	}
	s.storage = &Storage{
		Nodes: container.NewMap[*Info](nodeKey, contract.Address(), s.stateDb),
	}
	s.initABI()
	s.initABIV1()
	s.initABIV2()
	s.initABIV3()
	s.initMethodEntry()
	s.initMethodV1Entry()
	s.initMethodV2Entry()
	s.initMethodV3Entry()
	s.loadMethodABI()
	return s, nil
}

func (c *Node) AddNode(name string, host string, port uint16) error {
	contracts.Require(len(host) != 0, "Node: invalid host")
	contracts.Require(port != 0, "Node: invalid port")
	c.storage.Nodes.MustSet(name, &Info{
		Host: host,
		port: port,
	})
	return nil
}
