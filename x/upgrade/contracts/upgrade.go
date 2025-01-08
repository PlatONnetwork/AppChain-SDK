package contracts

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"

	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
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

// IUpgradeModule is an auto generated low-level Go binding around an user-defined struct.
type IUpgradeModule struct {
	ModuleName string
	Version    uint64
}

// IUpgradePlan is an auto generated low-level Go binding around an user-defined struct.
type IUpgradePlan struct {
	Name    string
	Modules []IUpgradeModule
	Info    string
	Height  uint64
	Status  uint8
}

var (
	ABI    = "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"moduleName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"internalType\":\"structIUpgrade.Module[]\",\"name\":\"modules\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"enumIUpgrade.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIUpgrade.Plan\",\"name\":\"plan\",\"type\":\"tuple\"}],\"name\":\"addUpgradePlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getUpgradePlan\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"moduleName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"internalType\":\"structIUpgrade.Module[]\",\"name\":\"modules\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"enumIUpgrade.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIUpgrade.Plan[]\",\"name\":\"plan\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"setUpgradePlanDone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	Abi, _ = abi.JSON(strings.NewReader(ABI))
)

func (c *Upgrade) Run(input []byte) (ret []byte, err error) {
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

func (c *Upgrade) initABI() {
	V0 := uint64(0)
	c.abis[V0] = &Abi
}

func (c *Upgrade) initMethodEntry() {

	methodEntry := map[string]func([]byte) ([]byte, error){
		"893d20e8": c.GetOwnerEntry,
		"39f278c7": c.GetUpgradePlanEntry,

		"c906633d": c.AddUpgradePlanEntry,
		"13af4035": c.SetOwnerEntry,
		"ab4f4c2f": c.SetUpgradePlanDoneEntry,
	}
	V0 := uint64(0)
	c.methodEntries[V0] = methodEntry

}
func (c *Upgrade) loadMethodABI() error {
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
func (c *Upgrade) InitGenesis(blockNumber uint64) {
	c.SetCreateBlock(blockNumber)
	c.stateDb.SetCode(c.contract.Address(), []byte("code"))
}
func (c *Upgrade) SetCreateBlock(blockNumber uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], blockNumber)
	c.evm.StateDB.SetState(c.contract.Address(), createBlockKey, data[:])
}

func (c *Upgrade) GetCreateBlock() uint64 {
	blockNumber := c.evm.StateDB.GetState(c.contract.Address(), createBlockKey)
	if len(blockNumber) == 0 {
		return math.MaxUint64
	}
	return binary.BigEndian.Uint64(blockNumber)
}

func (c *Upgrade) SetVersion(version uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], version)
	c.evm.StateDB.SetState(c.contract.Address(), versionKey, data[:])
}

func (c *Upgrade) GetVersion() uint64 {
	version := c.evm.StateDB.GetState(c.contract.Address(), versionKey)
	if len(version) == 0 {
		return 0
	}
	return binary.BigEndian.Uint64(version)
}

func (c *Upgrade) GetOwnerEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["getOwner"]

	var err error

	res0, err := c.GetOwner()
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

func (c *Upgrade) GetUpgradePlanEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["getUpgradePlan"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.GetUpgradePlan(*abi.ConvertType(args[0], new(uint64)).(*uint64))
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

func (c *Upgrade) AddUpgradePlanEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["addUpgradePlan"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.AddUpgradePlan(*abi.ConvertType(args[0], new(IUpgradePlan)).(*IUpgradePlan))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *Upgrade) SetOwnerEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["setOwner"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.SetOwner(*abi.ConvertType(args[0], new(common.Address)).(*common.Address))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *Upgrade) SetUpgradePlanDoneEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["setUpgradePlanDone"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.SetUpgradePlanDone(*abi.ConvertType(args[0], new(uint64)).(*uint64))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}
