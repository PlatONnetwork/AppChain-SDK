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

type VRFCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewVRFCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*VRFCaller, error) {
	s := &VRFCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *VRFCaller) Hash(nonceProof []byte) (common.Hash, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "hash", nonceProof)

	if err != nil {
		return *new(common.Hash), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Hash)).(*common.Hash)

	return out0, err

}

func (c *VRFCaller) Verify(pubKey []byte, pi []byte, m []byte) (bool, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "verify", pubKey, pi, m)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

type VRFDelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewVRFDelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*VRFDelegateCaller, error) {
	s := &VRFDelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *VRFDelegateCaller) Hash(nonceProof []byte) (common.Hash, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "hash", nonceProof)

	if err != nil {
		return *new(common.Hash), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Hash)).(*common.Hash)

	return out0, err

}

func (c *VRFDelegateCaller) Verify(pubKey []byte, pi []byte, m []byte) (bool, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "verify", pubKey, pi, m)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}
