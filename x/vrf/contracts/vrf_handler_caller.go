package contracts

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/tools/contracts"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
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

type VRFHandlerCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewVRFHandlerCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*VRFHandlerCaller, error) {
	s := &VRFHandlerCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *VRFHandlerCaller) PushNonceAndProof(nonceAndProof []byte) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "pushNonceAndProof", nonceAndProof)

	if err != nil {
		return err
	}

	return err

}

type VRFHandlerDelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewVRFHandlerDelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*VRFHandlerDelegateCaller, error) {
	s := &VRFHandlerDelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *VRFHandlerDelegateCaller) PushNonceAndProof(nonceAndProof []byte) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "pushNonceAndProof", nonceAndProof)

	if err != nil {
		return err
	}

	return err

}
