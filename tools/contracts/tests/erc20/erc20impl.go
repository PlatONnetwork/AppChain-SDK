package erc20

import (
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

type Erc20 struct {
	abi         *abi.ABI
	methodEntry map[string]func([]byte) ([]byte, error)
	readOnly    bool
	contract    *vm.Contract
	evm         *vm.EVM
	fallback    func(input []byte) ([]byte, error)
}

func NewErc20(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*Erc20, error) {
	s := &Erc20{
		abi:      &Abi,
		evm:      evm,
		contract: contract,
		readOnly: readOnly,
	}
	s.initMethodEntry()
	return s, nil
}

func (c *Erc20) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	buf := c.evm.StateDB.GetState(c.contract.Address(), append(owner.Bytes(), spender.Bytes()...))
	value := new(big.Int).SetBytes(buf)
	return value, nil
}

func (c *Erc20) BalanceOf(account common.Address) (*big.Int, error) {
	buf := c.evm.StateDB.GetState(c.contract.Address(), account.Bytes())
	value := new(big.Int).SetBytes(buf)
	return value, nil
}

func (c *Erc20) Decimals() (uint8, error) {
	return 18, nil
}

func (c *Erc20) Name() (string, error) {
	return "TestName", nil
}

func (c *Erc20) Symbol() (string, error) {
	return "TEU", nil
}

func (c *Erc20) TotalSupply() (*big.Int, error) {
	buf := c.evm.StateDB.GetState(c.contract.Address(), []byte("totalSupply"))
	value := new(big.Int).SetBytes(buf)
	return value, nil
}

func (c *Erc20) Approve(spender common.Address, amount *big.Int) (bool, error) {
	c.evm.StateDB.SetState(c.contract.Address(), append(c.contract.Caller().Bytes(), spender.Bytes()...), amount.Bytes())

	return true, nil
}

func (c *Erc20) Mint(account common.Address, value *big.Int) error {
	c.evm.StateDB.SetState(c.contract.Address(), account.Bytes(), value.Bytes())
	had, err := c.TotalSupply()
	if err != nil {
		return err
	}
	c.evm.StateDB.SetState(c.contract.Address(), []byte("totalSupply"), had.Add(had, value).Bytes())
	return nil
}

func (c *Erc20) Transfer(recipient common.Address, amount *big.Int) (bool, error) {
	from := new(big.Int).SetBytes(c.evm.StateDB.GetState(c.contract.Address(), c.contract.Caller().Bytes()))
	to := new(big.Int).SetBytes(c.evm.StateDB.GetState(c.contract.Address(), recipient.Bytes()))
	if from.Cmp(amount) < 0 {
		return false, errors.New("balance not enough")
	}
	from.Sub(from, amount)
	to.Add(to, amount)
	c.EmitTransferEvent(c.evm.Origin, recipient, amount)
	return true, nil
}

func (c *Erc20) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (bool, error) {
	approve := new(big.Int).SetBytes(c.evm.StateDB.GetState(c.contract.Address(), append(sender.Bytes(), c.contract.Caller().Bytes()...)))
	if approve.Cmp(amount) < 0 {
		return false, errors.New("approve not enough")
	}

	from := new(big.Int).SetBytes(c.evm.StateDB.GetState(c.contract.Address(), sender.Bytes()))
	to := new(big.Int).SetBytes(c.evm.StateDB.GetState(c.contract.Address(), recipient.Bytes()))
	if from.Cmp(amount) < 0 {
		return false, errors.New("balance not enough")
	}
	from.Sub(from, amount)
	to.Add(to, amount)
	c.EmitTransferEvent(sender, recipient, amount)
	c.evm.StateDB.SetState(c.contract.Address(), append(sender.Bytes(), c.evm.Origin.Bytes()...), approve.Sub(approve, amount).Bytes())
	return true, nil
}
