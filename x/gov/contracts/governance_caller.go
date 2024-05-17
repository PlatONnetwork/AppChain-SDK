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

type GovernanceCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewGovernanceCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*GovernanceCaller, error) {
	s := &GovernanceCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *GovernanceCaller) GetExecutableProposal() ([]Proposal, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "getExecutableProposal")

	if err != nil {
		return *new([]Proposal), err
	}

	out0 := *abi.ConvertType(out[0], new([]Proposal)).(*[]Proposal)

	return out0, err

}

func (c *GovernanceCaller) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "hasVoted", proposalId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *GovernanceCaller) HashProposal(proposalType uint8, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash common.Hash) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "hashProposal", proposalType, targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceCaller) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "proposalDeadline", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceCaller) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "proposalSnapshot", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceCaller) ProposalThreshold() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "proposalThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceCaller) State(proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (c *GovernanceCaller) Cancel(proposalType uint8, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash common.Hash) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "cancel", proposalType, targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceCaller) CastVote(proposalId *big.Int, support uint8) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "castVote", proposalId, support)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceCaller) CastVoteBySig(proposalId *big.Int, support uint8, v uint8, r common.Hash, s common.Hash) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "castVoteBySig", proposalId, support, v, r, s)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceCaller) Execute(proposalType uint8, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash common.Hash) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "execute", proposalType, targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceCaller) Propose(proposalType uint8, targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "propose", proposalType, targets, values, calldatas, description)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

type GovernanceDelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewGovernanceDelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*GovernanceDelegateCaller, error) {
	s := &GovernanceDelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *GovernanceDelegateCaller) GetExecutableProposal() ([]Proposal, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "getExecutableProposal")

	if err != nil {
		return *new([]Proposal), err
	}

	out0 := *abi.ConvertType(out[0], new([]Proposal)).(*[]Proposal)

	return out0, err

}

func (c *GovernanceDelegateCaller) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "hasVoted", proposalId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (c *GovernanceDelegateCaller) HashProposal(proposalType uint8, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash common.Hash) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "hashProposal", proposalType, targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceDelegateCaller) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "proposalDeadline", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceDelegateCaller) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "proposalSnapshot", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceDelegateCaller) ProposalThreshold() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "proposalThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceDelegateCaller) State(proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (c *GovernanceDelegateCaller) Cancel(proposalType uint8, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash common.Hash) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "cancel", proposalType, targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceDelegateCaller) CastVote(proposalId *big.Int, support uint8) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "castVote", proposalId, support)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceDelegateCaller) CastVoteBySig(proposalId *big.Int, support uint8, v uint8, r common.Hash, s common.Hash) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "castVoteBySig", proposalId, support, v, r, s)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceDelegateCaller) Execute(proposalType uint8, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash common.Hash) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "execute", proposalType, targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceDelegateCaller) Propose(proposalType uint8, targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "propose", proposalType, targets, values, calldatas, description)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}
