package erc20

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/db"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/db/container"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	contracts2 "github.com/PlatONnetwork/AppChain-SDK/x/gov/contracts/ownable"
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

type Storage struct {
	Balance     *container.Map[*big.Int]
	Allowances  *container.Map[*container.Map[*big.Int]]
	TotalSupply *db.Base[*big.Int]
	Name        *db.Base[string]
	Symbol      *db.Base[string]
}

type ERC20 struct {
	abi           *abi.ABI
	abis          map[uint64]*abi.ABI
	methodEntry   map[string]func([]byte) ([]byte, error)
	methodEntries map[uint64]map[string]func([]byte) ([]byte, error)
	readOnly      bool
	contract      *vm.Contract
	evm           *vm.EVM
	burner        contracts.Burn
	stateDb       *contracts.StateDB
	fallback      func(input []byte) ([]byte, error)
	storage       *Storage
	context       *contracts.Context
	contracts2.Ownable
}

func NewERC20(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*ERC20, error) {
	s := &ERC20{
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
	return s, nil
}

func (c *ERC20) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return c.storage.Allowances.
		MustGet(owner).
		MustGet(spender), nil
}

func (c *ERC20) BalanceOf(account common.Address) (*big.Int, error) {
	return c.storage.Balance.MustGet(account), nil
}

func (c *ERC20) Decimals() (uint8, error) {
	return 18, nil
}

func (c *ERC20) Name() (string, error) {
	return c.storage.Name.MustGet(), nil
}

func (c *ERC20) Symbol() (string, error) {
	return c.storage.Symbol.MustGet(), nil
}

func (c *ERC20) TotalSupply() (*big.Int, error) {
	return c.storage.TotalSupply.MustGet(), nil

}

func (c *ERC20) Approve(spender common.Address, amount *big.Int) (bool, error) {
	c.approve(c.context.Caller(), spender, amount)
	return true, nil
	panic("implement")
}

func (c *ERC20) approve(owner, spender common.Address, amount *big.Int) {
	contracts.Require(owner != common.Address{}, "ERC20: approve from the zero address")
	contracts.Require(spender != common.Address{}, "ERC20: approve to the zero address")
	c.storage.Allowances.MustGet(owner).MustSet(spender, amount)
	c.EmitApprovalEvent(owner, spender, amount)
}

func (c *ERC20) Burn(account common.Address, amount *big.Int) error {
	c.OnlyOwner()
	contracts.Require(account != common.Address{}, "ERC20: burn from the zero address")
	accountBalance := c.storage.Balance.MustGet(account)
	contracts.Require(accountBalance.Cmp(amount) > 0, "ERC20: burn amount exceeds balance")
	accountBalance.Sub(accountBalance, amount)
	c.storage.Balance.MustSet(account, accountBalance)
	totalSupply := c.storage.TotalSupply.MustGet()
	totalSupply.Sub(totalSupply, amount)
	c.storage.TotalSupply.MustSet(totalSupply)

	c.EmitTransferEvent(account, common.Address{}, amount)

	return nil
}

func (c *ERC20) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (bool, error) {
	msgSender := c.context.Caller()
	currentAllowance := c.storage.Allowances.MustGet(msgSender).MustGet(spender)
	contracts.Require(currentAllowance.Cmp(subtractedValue) > 0, "ERC20: decreased allowance below zero")
	c.approve(msgSender, spender, currentAllowance.Sub(currentAllowance, subtractedValue))
	return true, nil
}

func (c *ERC20) IncreaseAllowance(spender common.Address, addedValue *big.Int) (bool, error) {
	msgSender := c.context.Caller()
	allowance := c.storage.Allowances.MustGet(msgSender).MustGet(spender)
	c.approve(msgSender, spender, allowance.Add(allowance, addedValue))
	return true, nil
}

func (c *ERC20) Mint(account common.Address, amount *big.Int) error {
	c.OnlyOwner()
	contracts.Require(account != common.Address{}, "ERC20: mint from the zero address")
	accountBalance := c.storage.Balance.MustGet(account)
	accountBalance.Add(accountBalance, amount)
	c.storage.Balance.MustSet(account, accountBalance)
	totalSupply := c.storage.TotalSupply.MustGet()
	totalSupply.Add(totalSupply, amount)
	c.storage.TotalSupply.MustSet(totalSupply)

	c.EmitTransferEvent(common.Address{}, account, amount)
	return nil
}

func (c *ERC20) Transfer(recipient common.Address, amount *big.Int) (bool, error) {
	msgSender := c.context.Caller()
	c.transfer(msgSender, recipient, amount)
	return true, nil
}
func (c *ERC20) transfer(sender, recipient common.Address, amount *big.Int) {
	senderBalance := c.storage.Balance.MustGet(sender)
	contracts.Require(senderBalance.Cmp(amount) > 0, "ERC20: transfer amount exceeds balance")
	senderBalance.Sub(senderBalance, amount)
	c.storage.Balance.MustSet(sender, senderBalance)
	recipientBalance := c.storage.Balance.MustGet(recipient)
	recipientBalance.Add(recipientBalance, amount)
	c.storage.Balance.MustSet(recipient, recipientBalance)
	c.EmitTransferEvent(sender, recipient, amount)

}

func (c *ERC20) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (bool, error) {
	c.transfer(sender, recipient, amount)
	msgSender := c.context.Caller()
	currentAllowance := c.storage.Allowances.MustGet(sender).MustGet(msgSender)
	contracts.Require(currentAllowance.Cmp(amount) > 0, "ERC20: transfer amount exceeds allowance")
	c.approve(sender, msgSender, currentAllowance.Sub(currentAllowance, amount))
	return true, nil
}
