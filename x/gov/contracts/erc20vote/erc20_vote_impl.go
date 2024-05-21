package erc20vote

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/db/container"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/ecdsa"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/gov/contracts/eip712"
	"github.com/PlatONnetwork/AppChain-SDK/x/gov/contracts/erc20"
	contracts2 "github.com/PlatONnetwork/AppChain-SDK/x/gov/contracts/ownable"
	"github.com/PlatONnetwork/AppChain-SDK/x/gov/contracts/permit"
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
	delegateKey              = []byte("delegate")
	checkpointsKey           = []byte("checkpoints")
	totalSupplyCheckpointKey = []byte("totalSupplyCheckpoint")
	DELEGATION_TYPEHASH      = crypto.Keccak256Hash([]byte("Delegation(address delegatee,uint256 nonce,uint256 expiry)"))
)

type EIP712 interface {
	DomainSeparator() common.Hash
	HashTypedData(structHash common.Hash) common.Hash
}
type ERC20Permit interface {
	UseNonce(owner common.Address) *big.Int
}
type Storage struct {
	Delegate               *container.Map[common.Address]
	Checkpoints            *container.Map[*container.Array[Checkpoint]]
	TotalSupplyCheckpoints *container.Array[Checkpoint]
}
type ERC20Vote struct {
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
	eip712        *eip712.EIP712
	erc20Permit   *permit.ERC20Permit
	erc20         *erc20.ERC20
	*contracts2.Ownable
}

func NewERC20Vote(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*ERC20Vote, error) {
	s := &ERC20Vote{
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
	s.storage = Storage{
		Delegate:               container.NewMap[common.Address](delegateKey, contract.Address(), s.stateDb),
		Checkpoints:            container.NewMap[*container.Array[Checkpoint]](checkpointsKey, contract.Address(), s.stateDb),
		TotalSupplyCheckpoints: container.NewArray[Checkpoint](totalSupplyCheckpointKey, contract.Address(), s.stateDb),
	}
	var err error
	s.eip712, err = eip712.NewEIP712(evm, contract, readOnly)
	if err != nil {
		return nil, err
	}
	s.erc20, err = erc20.NewERC20(evm, contract, readOnly)
	if err != nil {
		return nil, err
	}
	s.erc20Permit, err = permit.NewERC20Permit(evm, contract, readOnly, s.eip712, s.erc20)
	if err != nil {
		return nil, err
	}
	s.initABI()
	s.initMethodEntry()
	s.loadMethodABI()
	return s, nil
}

func (c *ERC20Vote) Init(name, symbol, version string, owner common.Address) {
	c.eip712.Init(name, version)
	c.erc20.Init(name, symbol, owner)
}

func (c *ERC20Vote) Checkpoints(account common.Address, pos uint32) (Checkpoint, error) {
	cs, err := c.storage.Checkpoints.Get(account)
	if err != nil {
		return Checkpoint{}, err
	}
	return cs.Index(pos)
}

func (c *ERC20Vote) getDelegate(account common.Address) common.Address {
	delegate := c.storage.Delegate.MustGet(account)
	if common.ZeroAddr == delegate {
		return account
	}
	return delegate
}
func (c *ERC20Vote) Delegates(account common.Address) (common.Address, error) {
	return c.storage.Delegate.MustGet(account), nil
}

func (c *ERC20Vote) GetPastTotalSupply(blockNumber *big.Int) (*big.Int, error) {
	contracts.Require(blockNumber.Cmp(c.evm.Context.BlockNumber) < 0, "block not yet mined")
	return c.checkpointsLookup(c.storage.TotalSupplyCheckpoints, blockNumber.Uint64()), nil
}

func (c *ERC20Vote) GetPastVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	contracts.Require(blockNumber.Cmp(c.evm.Context.BlockNumber) < 0, "block not yet mined")
	cpts, err := c.storage.Checkpoints.Get(account)
	if err != nil {
		return big.NewInt(0), nil
	}
	return c.checkpointsLookup(cpts, blockNumber.Uint64()), nil
}

func (c *ERC20Vote) GetVotes(account common.Address) (*big.Int, error) {
	cpts, err := c.storage.Checkpoints.Get(account)
	if err != nil {
		return big.NewInt(0), nil
	}
	if cpts.Length() == 0 {
		return big.NewInt(0), nil
	}
	cp, err := cpts.Index(cpts.Length() - 1)
	if err != nil {
		return big.NewInt(0), nil
	}
	return cp.Votes, nil
}

func (c *ERC20Vote) NumCheckpoints(account common.Address) (uint32, error) {
	cs, err := c.storage.Checkpoints.Get(account)
	if err != nil {
		return 0, err
	}
	return cs.Length(), nil
}

func (c *ERC20Vote) AfterTokenTransfer(from common.Address, to common.Address, amount *big.Int) error {
	src := c.getDelegate(from)
	dst := c.getDelegate(to)
	c.moveVotingPower(src, dst, amount)
	return nil
}

func (c *ERC20Vote) BeforeTokenTransfer(from common.Address, to common.Address, amount *big.Int) error {
	return nil
}

func (c *ERC20Vote) Delegate(delegatee common.Address) error {
	c.delegate(c.context.Caller(), delegatee)
	return nil
}

func (c *ERC20Vote) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r common.Hash, s common.Hash) error {
	contracts.Require(c.context.Timestamp() <= expiry.Int64(), "ERC20Vote: signature expired")
	data, err := abi2.Encode([]interface{}{DELEGATION_TYPEHASH, delegatee, nonce, expiry}, abi2.MustNewType("tuple(bytes32 hash, address delegatee, uint256 nonce, uint256 expiry)"))
	contracts.Require(err == nil, "ERC20Vote: encode failed")
	signer := ecdsa.Recover(c.eip712.HashTypedData(crypto.Keccak256Hash(data)), v, r, s)
	contracts.Require(nonce.Cmp(c.erc20Permit.UseNonce(signer)) == 0, "ERC20Vote: invalid nonce")
	c.delegate(signer, delegatee)
	return nil
}

func (c *ERC20Vote) delegate(delegator, delegatee common.Address) {
	currentDelegate, _ := c.Delegates(delegator)
	delegatorBalance, _ := c.erc20.BalanceOf(delegator)
	c.storage.Delegate.MustSet(delegator, delegatee)
	c.EmitDelegateChangedEvent(delegator, currentDelegate, delegatee)
	c.moveVotingPower(currentDelegate, delegatee, delegatorBalance)
}

func (c *ERC20Vote) moveVotingPower(src, dst common.Address, amount *big.Int) {
	if src != dst && amount.Cmp(big.NewInt(0)) > 0 {
		if src != common.ZeroAddr {
			oldWeight, newWeight := c.writeCheckpoint(c.storage.Checkpoints.MustGet(src), sub, amount)
			c.EmitDelegateVotesChangedEvent(src, oldWeight, newWeight)
		}
		if dst != common.ZeroAddr {
			oldWeight, newWeight := c.writeCheckpoint(c.storage.Checkpoints.MustGet(dst), add, amount)
			c.EmitDelegateVotesChangedEvent(dst, oldWeight, newWeight)
		}
	}
}

func (c *ERC20Vote) writeCheckpoint(ckpts *container.Array[Checkpoint], op func(*big.Int, *big.Int) *big.Int, delta *big.Int) (*big.Int, *big.Int) {
	pos := ckpts.Length()
	oldWeight := big.NewInt(0)
	var oldCheckpoint Checkpoint
	if pos != 0 {
		oldCheckpoint = ckpts.MustIndex(pos - 1)
		oldWeight = oldCheckpoint.Votes
	}
	newWeight := op(oldWeight, delta)
	if pos > 0 && oldCheckpoint.FromBlock == c.evm.Context.BlockNumber.Uint64() {
		oldCheckpoint.Votes = newWeight
		ckpts.Replace(pos-1, oldCheckpoint)
	} else {
		ckpts.Push(Checkpoint{
			FromBlock: c.evm.Context.BlockNumber.Uint64(),
			Votes:     newWeight,
		})
	}
	return oldWeight, newWeight
}

func (c *ERC20Vote) checkpointsLookup(ckpts *container.Array[Checkpoint], blockNumber uint64) *big.Int {
	high := ckpts.Length()
	low := uint32(0)
	for low < high {
		mid := (low + high) / 2
		cs, err := ckpts.Index(mid)
		contracts.Require(err == nil, "checkpoint is null")
		if cs.FromBlock > blockNumber {
			high = mid
		} else {
			low = mid + 1
		}
	}
	if high == 0 {
		return big.NewInt(0)
	} else {
		cs, err := ckpts.Index(high - 1)
		contracts.Require(err == nil, "checkpoint is null")
		return cs.Votes
	}
}

func (c *ERC20Vote) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return c.erc20.Allowance(owner, spender)
}

func (c *ERC20Vote) BalanceOf(account common.Address) (*big.Int, error) {
	return c.erc20.BalanceOf(account)
}

func (c *ERC20Vote) Decimals() (uint8, error) {
	return c.erc20.Decimals()
}

func (c *ERC20Vote) Name() (string, error) {
	return c.erc20.Name()
}

func (c *ERC20Vote) Symbol() (string, error) {
	return c.erc20.Symbol()
}

func (c *ERC20Vote) TotalSupply() (*big.Int, error) {
	return c.erc20.TotalSupply()
}

func (c *ERC20Vote) Approve(spender common.Address, amount *big.Int) (bool, error) {
	return c.erc20.Approve(spender, amount)
}

func (c *ERC20Vote) Burn(account common.Address, amount *big.Int) error {
	c.BeforeTokenTransfer(account, common.Address{}, amount)

	c.erc20.Burn(account, amount)
	c.AfterTokenTransfer(account, common.Address{}, amount)
	c.writeCheckpoint(c.storage.TotalSupplyCheckpoints, sub, amount)

	return nil
}

func (c *ERC20Vote) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (bool, error) {
	return c.erc20.DecreaseAllowance(spender, subtractedValue)
}

func (c *ERC20Vote) IncreaseAllowance(spender common.Address, addedValue *big.Int) (bool, error) {
	return c.erc20.IncreaseAllowance(spender, addedValue)
}

func (c *ERC20Vote) Mint(account common.Address, amount *big.Int) error {
	c.BeforeTokenTransfer(common.Address{}, account, amount)
	c.erc20.Mint(account, amount)
	c.AfterTokenTransfer(common.Address{}, account, amount)
	c.writeCheckpoint(c.storage.TotalSupplyCheckpoints, add, amount)
	return nil
}

func (c *ERC20Vote) Transfer(recipient common.Address, amount *big.Int) (bool, error) {
	c.BeforeTokenTransfer(c.context.Caller(), recipient, amount)
	c.erc20.Transfer(recipient, amount)
	c.AfterTokenTransfer(c.context.Caller(), recipient, amount)
	return true, nil
}

func (c *ERC20Vote) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (bool, error) {
	c.BeforeTokenTransfer(sender, recipient, amount)
	c.erc20.TransferFrom(sender, recipient, amount)
	c.AfterTokenTransfer(sender, recipient, amount)
	return true, nil
}
func add(b *big.Int, b2 *big.Int) *big.Int {
	return new(big.Int).Add(b, b2)
}
func sub(b *big.Int, b2 *big.Int) *big.Int {
	return new(big.Int).Sub(b, b2)
}
