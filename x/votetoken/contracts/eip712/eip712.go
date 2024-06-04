package eip712

import (
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/db"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/ecdsa"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/umbracle/ethgo/abi"
	"math/big"
)

var (
	cachedDomainSeparatorKey = []byte("cachedDomainSeparator")
	cachedChainIdKey         = []byte("cachedChainId")
	cachedThisKey            = []byte("cachedThis")
	hashedNameKey            = []byte("hashedName")
	hashedVersionKey         = []byte("hashedVersion")
	typeHashKey              = []byte("typeHash")
	DomainSeparator          = abi.MustNewType("tuple(bytes32 typeHash, bytes32 nameHash, bytes32 versionHash, uint256 chainId, address this)")
)

type Storage struct {
	CACHED_DOMAIN_SEPARATOR *db.Base[common.Hash]
	CACHED_CHAIN_ID         *db.Base[*big.Int]
	CACHED_THIS             *db.Base[common.Address]
	HASHED_NAME             *db.Base[common.Hash]
	HASHED_VERSION          *db.Base[common.Hash]
	TYPE_HASH               *db.Base[common.Hash]
}
type EIP712 struct {
	readOnly bool
	contract *vm.Contract
	evm      *vm.EVM
	burner   contracts.Burn
	stateDb  *contracts.StateDB
	fallback func(input []byte) ([]byte, error)
	storage  Storage
	context  *contracts.Context
}

func NewEIP712(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*EIP712, error) {
	s := &EIP712{
		evm:      evm,
		contract: contract,
		burner:   contracts.NewBurner(contract),
		stateDb:  contracts.NewStateDB(evm, contract),
		context:  contracts.NewContext(evm, contract),
		readOnly: readOnly,
	}
	store := db.NewStore([]byte{}, contract.Address(), s.stateDb)
	s.storage = Storage{
		CACHED_DOMAIN_SEPARATOR: db.NewBase[common.Hash](cachedDomainSeparatorKey, store),
		CACHED_CHAIN_ID:         db.NewBase[*big.Int](cachedChainIdKey, store),
		CACHED_THIS:             db.NewBase[common.Address](cachedThisKey, store),
		HASHED_NAME:             db.NewBase[common.Hash](hashedNameKey, store),
		HASHED_VERSION:          db.NewBase[common.Hash](hashedVersionKey, store),
		TYPE_HASH:               db.NewBase[common.Hash](typeHashKey, store),
	}
	return s, nil
}

func (c *EIP712) Init(name, version string) {
	hashedName := crypto.Keccak256Hash([]byte(name))
	hashedVersion := crypto.Keccak256Hash([]byte(version))
	typeHash := crypto.Keccak256Hash(
		[]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"))
	c.storage.HASHED_NAME.MustSet(hashedName)
	c.storage.HASHED_VERSION.MustSet(hashedVersion)
	c.storage.CACHED_CHAIN_ID.MustSet(c.context.ChainID())
	c.storage.CACHED_DOMAIN_SEPARATOR.MustSet(c.BuildDomainSeparator(typeHash, hashedName, hashedVersion))
	c.storage.CACHED_THIS.MustSet(c.context.Address())
	c.storage.TYPE_HASH.MustSet(typeHash)
}
func (c *EIP712) DomainSeparator() common.Hash {
	cacheThis := c.storage.CACHED_THIS.MustGet()
	chainId := c.storage.CACHED_CHAIN_ID.MustGet()
	if cacheThis == c.context.Address() && c.context.ChainID().Cmp(chainId) == 0 {
		return c.storage.CACHED_DOMAIN_SEPARATOR.MustGet()
	} else {
		return c.BuildDomainSeparator(c.storage.TYPE_HASH.MustGet(), c.storage.HASHED_NAME.MustGet(), c.storage.HASHED_VERSION.MustGet())
	}
}

func (c *EIP712) BuildDomainSeparator(typeHash, nameHash, versionHash common.Hash) common.Hash {
	data, err := abi.Encode([]interface{}{typeHash, nameHash, versionHash, c.context.ChainID(), c.context.Address()}, DomainSeparator)
	contracts.Require(err == nil, "EIP712: build domain separator failed")
	return crypto.Keccak256Hash(data)
}

func (c *EIP712) HashTypedData(structHash common.Hash) common.Hash {
	return ecdsa.ToTypedDataHash(c.DomainSeparator(), structHash)
}
