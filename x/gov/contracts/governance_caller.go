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

func (c *GovernanceCaller) GetVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "getVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

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

func (c *GovernanceCaller) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash common.Hash) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "hashProposal", targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceCaller) Name() (string, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

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

func (c *GovernanceCaller) Quorum(blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "quorum", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceCaller) QuorumDenominator() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "quorumDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceCaller) QuorumNumerator() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "quorumNumerator")

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

func (c *GovernanceCaller) Version() (string, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (c *GovernanceCaller) VotingDelay() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "votingDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceCaller) VotingPeriod() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "votingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceCaller) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash common.Hash) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "cancel", targets, values, calldatas, descriptionHash)

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

func (c *GovernanceCaller) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash common.Hash) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "execute", targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceCaller) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "propose", targets, values, calldatas, description)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceCaller) SetProposalThreshold(newProposalThreshold *big.Int) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "setProposalThreshold", newProposalThreshold)

	if err != nil {
		return err
	}

	return err

}

func (c *GovernanceCaller) SetVotingDelay(newVotingDelay *big.Int) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "setVotingDelay", newVotingDelay)

	if err != nil {
		return err
	}

	return err

}

func (c *GovernanceCaller) SetVotingPeriod(newVotingPeriod *big.Int) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "setVotingPeriod", newVotingPeriod)

	if err != nil {
		return err
	}

	return err

}

func (c *GovernanceCaller) UpdateQuorumNumerator(newQuorumNumerator *big.Int) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "updateQuorumNumerator", newQuorumNumerator)

	if err != nil {
		return err
	}

	return err

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

func (c *GovernanceDelegateCaller) GetVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "getVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

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

func (c *GovernanceDelegateCaller) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash common.Hash) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "hashProposal", targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceDelegateCaller) Name() (string, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

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

func (c *GovernanceDelegateCaller) Quorum(blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "quorum", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceDelegateCaller) QuorumDenominator() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "quorumDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceDelegateCaller) QuorumNumerator() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "quorumNumerator")

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

func (c *GovernanceDelegateCaller) Version() (string, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (c *GovernanceDelegateCaller) VotingDelay() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "votingDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceDelegateCaller) VotingPeriod() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "votingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceDelegateCaller) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash common.Hash) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "cancel", targets, values, calldatas, descriptionHash)

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

func (c *GovernanceDelegateCaller) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash common.Hash) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "execute", targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceDelegateCaller) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "propose", targets, values, calldatas, description)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *GovernanceDelegateCaller) SetProposalThreshold(newProposalThreshold *big.Int) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "setProposalThreshold", newProposalThreshold)

	if err != nil {
		return err
	}

	return err

}

func (c *GovernanceDelegateCaller) SetVotingDelay(newVotingDelay *big.Int) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "setVotingDelay", newVotingDelay)

	if err != nil {
		return err
	}

	return err

}

func (c *GovernanceDelegateCaller) SetVotingPeriod(newVotingPeriod *big.Int) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "setVotingPeriod", newVotingPeriod)

	if err != nil {
		return err
	}

	return err

}

func (c *GovernanceDelegateCaller) UpdateQuorumNumerator(newQuorumNumerator *big.Int) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "updateQuorumNumerator", newQuorumNumerator)

	if err != nil {
		return err
	}

	return err

}
