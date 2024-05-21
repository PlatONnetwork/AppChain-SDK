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

type StageManagerCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewStageManagerCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*StageManagerCaller, error) {
	s := &StageManagerCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *StageManagerCaller) GetPeriodByBlockNumber(periodType uint8, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "getPeriodByBlockNumber", periodType, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *StageManagerCaller) GetPeriodEdge(periodType uint8, period *big.Int) (PeriodEdge, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "getPeriodEdge", periodType, period)

	if err != nil {
		return *new(PeriodEdge), err
	}

	out0 := *abi.ConvertType(out[0], new(PeriodEdge)).(*PeriodEdge)

	return out0, err

}

func (c *StageManagerCaller) GetPeriodEdges(periodType uint8, start *big.Int, size *big.Int) (*big.Int, []*big.Int, []PeriodEdge, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "getPeriodEdges", periodType, start, size)

	if err != nil {
		return *new(*big.Int), *new([]*big.Int), *new([]PeriodEdge), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	out2 := *abi.ConvertType(out[2], new([]PeriodEdge)).(*[]PeriodEdge)

	return out0, out1, out2, err

}

type StageManagerDelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewStageManagerDelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*StageManagerDelegateCaller, error) {
	s := &StageManagerDelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *StageManagerDelegateCaller) GetPeriodByBlockNumber(periodType uint8, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "getPeriodByBlockNumber", periodType, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *StageManagerDelegateCaller) GetPeriodEdge(periodType uint8, period *big.Int) (PeriodEdge, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "getPeriodEdge", periodType, period)

	if err != nil {
		return *new(PeriodEdge), err
	}

	out0 := *abi.ConvertType(out[0], new(PeriodEdge)).(*PeriodEdge)

	return out0, err

}

func (c *StageManagerDelegateCaller) GetPeriodEdges(periodType uint8, start *big.Int, size *big.Int) (*big.Int, []*big.Int, []PeriodEdge, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "getPeriodEdges", periodType, start, size)

	if err != nil {
		return *new(*big.Int), *new([]*big.Int), *new([]PeriodEdge), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	out2 := *abi.ConvertType(out[2], new([]PeriodEdge)).(*[]PeriodEdge)

	return out0, out1, out2, err

}
