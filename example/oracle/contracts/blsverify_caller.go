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

type BlsVerifyCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewBlsVerifyCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*BlsVerifyCaller, error) {
	s := &BlsVerifyCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *BlsVerifyCaller) VerifyAggSignature(message common.Hash, signature []byte, pubs [][]byte) (bool, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "verifyAggSignature", message, signature, pubs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

type BlsVerifyDelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewBlsVerifyDelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*BlsVerifyDelegateCaller, error) {
	s := &BlsVerifyDelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *BlsVerifyDelegateCaller) VerifyAggSignature(message common.Hash, signature []byte, pubs [][]byte) (bool, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "verifyAggSignature", message, signature, pubs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}
