package permit

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/db/container"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/ecdsa"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/event"
	abi2 "github.com/umbracle/ethgo/abi"
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
	nonceKey        = []byte("permit_nonce")
	PERMIT_TYPEHASH = crypto.Keccak256Hash([]byte("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)"))
	StructHash      = abi2.MustNewType("tuple(bytes32 typeHash, address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)")
)

type EIP712 interface {
	DomainSeparator() common.Hash
	HashTypedData(structHash common.Hash) common.Hash
}

type ERC20 interface {
	Approve(owner, spender common.Address, amount *big.Int) (bool, error)
}

type Storage struct {
	Nonces *container.Map[*big.Int]
}

type ERC20Permit struct {
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
	eip712        EIP712
	erc20         ERC20
}

func NewERC20Permit(evm *vm.EVM, contract *vm.Contract, readOnly bool, eip712 EIP712, erc20 ERC20) (*ERC20Permit, error) {
	s := &ERC20Permit{
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
		eip712:        eip712,
		erc20:         erc20,
	}
	s.storage = Storage{
		Nonces: container.NewMap[*big.Int](nonceKey, contract.Address(), s.stateDb),
	}
	s.initABI()
	s.initMethodEntry()
	return s, nil
}

func (c *ERC20Permit) DOMAINSEPARATOR() (common.Hash, error) {
	return c.eip712.DomainSeparator(), nil
}

func (c *ERC20Permit) Nonces(owner common.Address) (*big.Int, error) {
	nonce := c.storage.Nonces.MustGet(owner)
	return nonce, nil
}

func (c *ERC20Permit) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r common.Hash, s common.Hash) error {
	contracts.Require(c.context.Timestamp() <= deadline.Int64(), "ERC20Permit: expired deadline")
	data, err := abi2.Encode([]interface{}{PERMIT_TYPEHASH, owner, spender, value, c.UseNonce(owner), deadline}, StructHash)
	contracts.Require(err != nil, "ERC20Permit: encode failed")
	structHash := crypto.Keccak256Hash(data)
	hash := c.eip712.HashTypedData(structHash)
	signer := ecdsa.Recover(hash, v, r, s)
	contracts.Require(signer == owner, "ERC20Permit: invalid signature")
	c.erc20.Approve(owner, spender, value)
	return nil
}

func (c *ERC20Permit) UseNonce(owner common.Address) *big.Int {
	nonce := c.storage.Nonces.MustGet(owner)
	c.storage.Nonces.MustSet(owner, nonce.Add(nonce, big.NewInt(1)))
	return nonce
}
