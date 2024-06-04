package contracts

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/db"
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
	ownerKey = []byte("owner")
)

type Storage struct {
	Owner *db.Base[common.Address]
}

type Ownable struct {
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
	storage       Storage
}

func NewOwnable(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*Ownable, error) {
	s := &Ownable{
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
	store := db.NewStore([]byte{}, contract.Address(), s.stateDb)
	s.storage = Storage{Owner: db.NewBase[common.Address](ownerKey, store)}
	s.initABI()
	s.initMethodEntry()
	return s, nil
}
func (c *Ownable) Init(owner common.Address) {
	c.storage.Owner.MustSet(owner)
}
func (c *Ownable) Owner() (common.Address, error) {
	return c.storage.Owner.MustGet(), nil
}

func (c *Ownable) OnlyOwner() {
	msgSender := c.context.Caller()
	owner := c.storage.Owner.MustGet()
	contracts.Require(msgSender == owner, "Ownable: caller is not the owner")
}

func (c *Ownable) RenounceOwnership() error {
	c.OnlyOwner()
	c.transferOwnership(common.Address{})
	return nil
}

func (c *Ownable) TransferOwnership(newOwner common.Address) error {
	contracts.Require(newOwner != common.Address{}, "Ownable: new owner is the zero address")
	c.transferOwnership(newOwner)
	return nil
}

func (c *Ownable) transferOwnership(newOwner common.Address) error {
	oldOwner := c.storage.Owner.MustGet()
	c.storage.Owner.MustSet(newOwner)
	c.EmitOwnershipTransferredEvent(oldOwner, newOwner)
	return nil
}
