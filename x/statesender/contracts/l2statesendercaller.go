package contracts

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/tools/contracts"
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

type L2StateSenderCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewL2StateSenderCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*L2StateSenderCaller, error) {
	s := &L2StateSenderCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *L2StateSenderCaller) MAXLENGTH() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "MAX_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *L2StateSenderCaller) Counter() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *L2StateSenderCaller) SyncState(receiver common.Address, data []byte) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "syncState", receiver, data)

	if err != nil {
		return err
	}

	return err

}

type L2StateSenderDelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewL2StateSenderDelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*L2StateSenderDelegateCaller, error) {
	s := &L2StateSenderDelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *L2StateSenderDelegateCaller) MAXLENGTH() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "MAX_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *L2StateSenderDelegateCaller) Counter() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *L2StateSenderDelegateCaller) SyncState(receiver common.Address, data []byte) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "syncState", receiver, data)

	if err != nil {
		return err
	}

	return err

}
