package contracts

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/tools/contracts"
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

type NodeCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewNodeCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*NodeCaller, error) {
	s := &NodeCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *NodeCaller) AddNode(name string, host string, port uint16) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "addNode", name, host, port)

	if err != nil {
		return err
	}

	return err

}

type NodeDelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewNodeDelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*NodeDelegateCaller, error) {
	s := &NodeDelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *NodeDelegateCaller) AddNode(name string, host string, port uint16) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "addNode", name, host, port)

	if err != nil {
		return err
	}

	return err

}
