package contracts

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"math/big"
)

type Context struct {
	evm      *vm.EVM
	contract *vm.Contract
	statedb  *StateDB
}

func NewContext(evm *vm.EVM, contract *vm.Contract) *Context {
	return &Context{
		evm:      evm,
		contract: contract,
		statedb:  NewStateDB(evm, contract),
	}
}
func (c *Context) ChainID() *big.Int {
	return c.evm.ChainConfig().ChainID
}
func (c *Context) GasPrice() *big.Int {
	return c.evm.GasPrice
}

func (c *Context) BlockNumber() *big.Int {
	return c.evm.Context.BlockNumber
}

func (c *Context) GasLimit() uint64 {
	return c.evm.Context.GasLimit
}

func (c *Context) Gas() uint64 {
	return c.contract.Gas
}

func (c *Context) Timestamp() int64 {
	return c.evm.Context.Time.Int64()
}

func (c *Context) Coinbase() common.Address {
	return c.evm.Context.Coinbase
}

func (c *Context) Origin() common.Address {
	return c.evm.Origin
}

func (c *Context) Caller() common.Address {
	return c.contract.Caller()
}

func (c *Context) Value() *big.Int {
	return c.contract.Value()
}

func (c *Context) Address() common.Address {
	return c.contract.Address()
}

func (c *Context) Nonce() uint64 {
	addr := c.contract.Caller()
	return c.statedb.GetNonce(addr)
}
