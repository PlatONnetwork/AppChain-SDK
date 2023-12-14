package contracts

import (
	"encoding/hex"
	"errors"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"math/big"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
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
	ABI    = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nonce\",\"type\":\"bytes\"}],\"name\":\"VRFNonceAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nonceAndProof\",\"type\":\"bytes\"}],\"name\":\"pushNonceAndProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	Abi, _ = abi.JSON(strings.NewReader(ABI))
)

func (c *VRFHandler) Run(input []byte) ([]byte, error) {
	if len(input) < 4 {
		return nil, errors.New("input too short")
	}
	id := input[0:4]
	log.Info("Start execute vrf contract", "blockNumber", c.evm.Context.BlockNumber.Uint64(), "methId", hex.EncodeToString(id), "input", input)
	entry, ok := c.methodEntry[hex.EncodeToString(id)]
	if !ok {
		if c.fallback != nil {
			return c.fallback(input)
		}
		return nil, errors.New("methods not found")
	}
	return entry(input[4:])
}
func (c *VRFHandler) initMethodEntry() {

	c.methodEntry = map[string]func([]byte) ([]byte, error){

		"5bf28775": c.PushNonceAndProofEntry,
	}

}

func (c *VRFHandler) PushNonceAndProofEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["pushNonceAndProof"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.PushNonceAndProof(*abi.ConvertType(args[0], new([]byte)).(*[]byte))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *VRFHandler) EmitVRFNonceAddedEvent(block *big.Int, nonce []byte) (*types.Log, error) {
	event := c.abi.Events["VRFNonceAdded"]
	hashes, err := abi.PackTopics(event.Inputs, block, nonce)
	if err != nil {
		return nil, err
	}
	data, err := event.Inputs.Pack(block, nonce)
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
