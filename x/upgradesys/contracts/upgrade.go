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

// BitArray is an auto generated low-level Go binding around an user-defined struct.
type BitArray struct {
	Bits  uint32
	Elems []uint64
}

// QuorumCert is an auto generated low-level Go binding around an user-defined struct.
type QuorumCert struct {
	Epoch        uint64
	ViewNumber   uint64
	BlockHash    [32]byte
	BlockNumber  uint64
	BlockIndex   uint32
	ExtendHash   [32]byte
	Signature    []byte
	ValidatorSet BitArray
}

// UpgradeCommitment is an auto generated low-level Go binding around an user-defined struct.
type UpgradeCommitment struct {
	Origin  common.Address
	Upgrade common.Address
	Data    []byte
}

var (
	ABI    = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"upgrade\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgrade\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structUpgradeCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"viewNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"blockIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"extendHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"bits\",\"type\":\"uint32\"},{\"internalType\":\"uint64[]\",\"name\":\"Elems\",\"type\":\"uint64[]\"}],\"internalType\":\"structBitArray\",\"name\":\"validatorSet\",\"type\":\"tuple\"}],\"internalType\":\"structQuorumCert\",\"name\":\"qc\",\"type\":\"tuple\"}],\"name\":\"commitUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implement\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implement\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	Abi, _ = abi.JSON(strings.NewReader(ABI))
)

func (c *Upgrade) Run(input []byte) ([]byte, error) {
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
func (c *Upgrade) initMethodEntry() {

	c.methodEntry = map[string]func([]byte) ([]byte, error){
		"aa254851": c.ImplementEntry,

		"d302f799": c.CommitUpgradeEntry,
		"c4d66de8": c.InitializeEntry,
	}

}

func (c *Upgrade) ImplementEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["implement"]

	var err error

	res0, err := c.Implement()
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

func (c *Upgrade) CommitUpgradeEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["commitUpgrade"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.CommitUpgrade(*abi.ConvertType(args[0], new(UpgradeCommitment)).(*UpgradeCommitment), *abi.ConvertType(args[1], new(uint64)).(*uint64), *abi.ConvertType(args[2], new([][32]byte)).(*[][32]byte), *abi.ConvertType(args[3], new(QuorumCert)).(*QuorumCert))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *Upgrade) InitializeEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["initialize"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.Initialize(*abi.ConvertType(args[0], new(common.Address)).(*common.Address))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *Upgrade) EmitUpgradedEvent(origin common.Address, upgrade common.Address) (*types.Log, error) {
	event := c.abi.Events["Upgraded"]
	hashes, err := abi.PackTopics(event.Inputs, origin, upgrade)
	if err != nil {
		return nil, err
	}
	data, err := event.Inputs.Pack(origin, upgrade)
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
