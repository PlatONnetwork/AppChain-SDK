package contracts

import (
	"errors"
	"math/big"
	"strings"

	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"

	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
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

type Upgrade struct {
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
}

func NewUpgrade(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*Upgrade, error) {
	s := &Upgrade{
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
	s.initABI()
	s.initMethodEntry()
	s.loadMethodABI()
	return s, nil
}

func (c *Upgrade) GetOwner() (common.Address, error) {
	return c.getOwner(), nil
}

func (c *Upgrade) GetUpgradePlan(height uint64) ([]IUpgradePlan, error) {
	return c.getUpgradePlan(height)
}

func (c *Upgrade) AddUpgradePlan(plan IUpgradePlan) error {
	c.onlyOwner()
	contracts.Require(plan.Height > c.evm.Context.BlockNumber.Uint64(), "invalid height")
	return c.addUpgradePlan(plan)
}

func (c *Upgrade) SetOwner(newOwner common.Address) error {
	c.onlyOwner()
	c.setOwner(newOwner)
	return nil
}

func (c *Upgrade) SetUpgradePlanDone(height uint64) error {
	c.onlyOwner()
	contracts.Require(height == c.evm.Context.BlockNumber.Uint64(), "invalid height")
	return c.setUpgradePlanDone(height)
}
