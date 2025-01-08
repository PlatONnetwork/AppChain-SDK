package contracts

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"math/big"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_              = vm.EVM{}
	_              = errors.New
	_              = big.NewInt
	_              = strings.NewReader
	_              = platon.NotFound
	_              = bind.Bind
	_              = common.Big1
	_              = math.ReadBits
	_              = binary.BigEndian
	_              = types.BloomLookup
	_              = event.NewSubscription
	versionKey     = []byte("__version")
	createBlockKey = []byte("__createBlock")
)

var (
	ABI    = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProposalThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProposalThreshold\",\"type\":\"uint256\"}],\"name\":\"ProposalThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldQuorumNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"QuorumNumeratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingDelay\",\"type\":\"uint256\"}],\"name\":\"VotingDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingPeriod\",\"type\":\"uint256\"}],\"name\":\"VotingPeriodSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"castVoteBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProposalThreshold\",\"type\":\"uint256\"}],\"name\":\"setProposalThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVotingDelay\",\"type\":\"uint256\"}],\"name\":\"setVotingDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVotingPeriod\",\"type\":\"uint256\"}],\"name\":\"setVotingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"updateQuorumNumerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
	Abi, _ = abi.JSON(strings.NewReader(ABI))
)

func (c *Governance) Run(input []byte) (ret []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case error:
				if r, ok := e.(*typesdk.RevertError); ok {
					ret, err = r.ReturnData, vm.ErrExecutionReverted
				} else {
					ret, err = nil, e
				}
			default:
				ret, err = typesdk.UndefinedError, vm.ErrExecutionReverted
			}
		}
	}()
	if err := c.loadMethodABI(); err != nil {
		return nil, errors.New("load version failed")
	}

	if len(input) < 4 {
		return nil, errors.New("input too short")
	}
	id := input[0:4]
	entry, ok := c.methodEntry[hex.EncodeToString(id)]
	if !ok {
		if c.fallback != nil {
			return c.fallback(input)
		}
		return nil, errors.New("methods not found")
	}
	return entry(input[4:])
}

func (c *Governance) initABI() {
	V0 := uint64(0)
	c.abis[V0] = &Abi
}

func (c *Governance) initMethodEntry() {

	methodEntry := map[string]func([]byte) ([]byte, error){
		"eb9019d4": c.GetVotesEntry,
		"43859632": c.HasVotedEntry,
		"c59057e4": c.HashProposalEntry,
		"06fdde03": c.NameEntry,
		"c01f9e37": c.ProposalDeadlineEntry,
		"2d63f693": c.ProposalSnapshotEntry,
		"b58131b0": c.ProposalThresholdEntry,
		"f8ce560a": c.QuorumEntry,
		"97c3d334": c.QuorumDenominatorEntry,
		"a7713a70": c.QuorumNumeratorEntry,
		"3e4f49e6": c.StateEntry,
		"54fd4d50": c.VersionEntry,
		"3932abb1": c.VotingDelayEntry,
		"02a251a3": c.VotingPeriodEntry,

		"452115d6": c.CancelEntry,
		"56781388": c.CastVoteEntry,
		"3bccf4fd": c.CastVoteBySigEntry,
		"2656227d": c.ExecuteEntry,
		"7d5e81e2": c.ProposeEntry,
		"ece40cc1": c.SetProposalThresholdEntry,
		"70b0f660": c.SetVotingDelayEntry,
		"ea0217cf": c.SetVotingPeriodEntry,
		"06f3f9e6": c.UpdateQuorumNumeratorEntry,
	}
	V0 := uint64(0)
	c.methodEntries[V0] = methodEntry

}
func (c *Governance) loadMethodABI() error {
	version := c.GetVersion()
	entries, ok := c.methodEntries[version]
	if !ok {
		return errors.New("unknown version")
	}
	c.methodEntry = entries
	abi, ok := c.abis[version]
	if !ok {
		return errors.New("unknown version")
	}
	c.abi = abi
	return nil
}
func (c *Governance) InitGenesis(blockNumber uint64) {
	c.SetCreateBlock(blockNumber)
	c.stateDb.SetCode(c.contract.Address(), []byte("code"))
}
func (c *Governance) SetCreateBlock(blockNumber uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], blockNumber)
	c.evm.StateDB.SetState(c.contract.Address(), createBlockKey, data[:])
}

func (c *Governance) GetCreateBlock() uint64 {
	blockNumber := c.evm.StateDB.GetState(c.contract.Address(), createBlockKey)
	if len(blockNumber) == 0 {
		return math.MaxUint64
	}
	return binary.BigEndian.Uint64(blockNumber)
}

func (c *Governance) SetVersion(version uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], version)
	c.evm.StateDB.SetState(c.contract.Address(), versionKey, data[:])
}

func (c *Governance) GetVersion() uint64 {
	version := c.evm.StateDB.GetState(c.contract.Address(), versionKey)
	if len(version) == 0 {
		return 0
	}
	return binary.BigEndian.Uint64(version)
}

func (c *Governance) GetVotesEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["getVotes"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.GetVotes(*abi.ConvertType(args[0], new(common.Address)).(*common.Address), *abi.ConvertType(args[1], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) HasVotedEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["hasVoted"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.HasVoted(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int), *abi.ConvertType(args[1], new(common.Address)).(*common.Address))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) HashProposalEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["hashProposal"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.HashProposal(*abi.ConvertType(args[0], new([]common.Address)).(*[]common.Address), *abi.ConvertType(args[1], new([]*big.Int)).(*[]*big.Int), *abi.ConvertType(args[2], new([][]byte)).(*[][]byte), *abi.ConvertType(args[3], new(common.Hash)).(*common.Hash))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) NameEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["name"]

	var err error

	res0, err := c.Name()
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) ProposalDeadlineEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["proposalDeadline"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.ProposalDeadline(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) ProposalSnapshotEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["proposalSnapshot"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.ProposalSnapshot(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) ProposalThresholdEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["proposalThreshold"]

	var err error

	res0, err := c.ProposalThreshold()
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) QuorumEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["quorum"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.Quorum(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) QuorumDenominatorEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["quorumDenominator"]

	var err error

	res0, err := c.QuorumDenominator()
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) QuorumNumeratorEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["quorumNumerator"]

	var err error

	res0, err := c.QuorumNumerator()
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) StateEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["state"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.State(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) VersionEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["version"]

	var err error

	res0, err := c.Version()
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) VotingDelayEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["votingDelay"]

	var err error

	res0, err := c.VotingDelay()
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) VotingPeriodEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["votingPeriod"]

	var err error

	res0, err := c.VotingPeriod()
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) CancelEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["cancel"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.Cancel(*abi.ConvertType(args[0], new([]common.Address)).(*[]common.Address), *abi.ConvertType(args[1], new([]*big.Int)).(*[]*big.Int), *abi.ConvertType(args[2], new([][]byte)).(*[][]byte), *abi.ConvertType(args[3], new(common.Hash)).(*common.Hash))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) CastVoteEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["castVote"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.CastVote(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int), *abi.ConvertType(args[1], new(uint8)).(*uint8))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) CastVoteBySigEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["castVoteBySig"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.CastVoteBySig(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int), *abi.ConvertType(args[1], new(uint8)).(*uint8), *abi.ConvertType(args[2], new(uint8)).(*uint8), *abi.ConvertType(args[3], new(common.Hash)).(*common.Hash), *abi.ConvertType(args[4], new(common.Hash)).(*common.Hash))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) ExecuteEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["execute"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.Execute(*abi.ConvertType(args[0], new([]common.Address)).(*[]common.Address), *abi.ConvertType(args[1], new([]*big.Int)).(*[]*big.Int), *abi.ConvertType(args[2], new([][]byte)).(*[][]byte), *abi.ConvertType(args[3], new(common.Hash)).(*common.Hash))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) ProposeEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["propose"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	res0, err := c.Propose(*abi.ConvertType(args[0], new([]common.Address)).(*[]common.Address), *abi.ConvertType(args[1], new([]*big.Int)).(*[]*big.Int), *abi.ConvertType(args[2], new([][]byte)).(*[][]byte), *abi.ConvertType(args[3], new(string)).(*string))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Governance) SetProposalThresholdEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["setProposalThreshold"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.SetProposalThreshold(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *Governance) SetVotingDelayEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["setVotingDelay"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.SetVotingDelay(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *Governance) SetVotingPeriodEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["setVotingPeriod"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.SetVotingPeriod(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *Governance) UpdateQuorumNumeratorEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["updateQuorumNumerator"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.UpdateQuorumNumerator(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *Governance) ProposalCanceledEvent(proposalId *big.Int) (*types.Log, error) {
	event := c.abi.Events["ProposalCanceled"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, proposalId)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, proposalId)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address:     c.contract.Address(),
		Topics:      hashes,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	}, nil
}
func (c *Governance) EmitProposalCanceledEvent(proposalId *big.Int) {
	log, err := c.ProposalCanceledEvent(proposalId)
	contracts.Require(err == nil, "Governance: emit ProposalCanceled event failed")
	c.stateDb.AddLog(log)
}

func (c *Governance) ProposalCreatedEvent(proposalId *big.Int, proposer common.Address, targets []common.Address, values []*big.Int, signatures []string, calldatas [][]byte, startBlock *big.Int, endBlock *big.Int, description string) (*types.Log, error) {
	event := c.abi.Events["ProposalCreated"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, proposalId, proposer, targets, values, signatures, calldatas, startBlock, endBlock, description)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, proposalId, proposer, targets, values, signatures, calldatas, startBlock, endBlock, description)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address:     c.contract.Address(),
		Topics:      hashes,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	}, nil
}
func (c *Governance) EmitProposalCreatedEvent(proposalId *big.Int, proposer common.Address, targets []common.Address, values []*big.Int, signatures []string, calldatas [][]byte, startBlock *big.Int, endBlock *big.Int, description string) {
	log, err := c.ProposalCreatedEvent(proposalId, proposer, targets, values, signatures, calldatas, startBlock, endBlock, description)
	contracts.Require(err == nil, "Governance: emit ProposalCreated event failed")
	c.stateDb.AddLog(log)
}

func (c *Governance) ProposalExecutedEvent(proposalId *big.Int) (*types.Log, error) {
	event := c.abi.Events["ProposalExecuted"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, proposalId)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, proposalId)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address:     c.contract.Address(),
		Topics:      hashes,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	}, nil
}
func (c *Governance) EmitProposalExecutedEvent(proposalId *big.Int) {
	log, err := c.ProposalExecutedEvent(proposalId)
	contracts.Require(err == nil, "Governance: emit ProposalExecuted event failed")
	c.stateDb.AddLog(log)
}

func (c *Governance) ProposalThresholdSetEvent(oldProposalThreshold *big.Int, newProposalThreshold *big.Int) (*types.Log, error) {
	event := c.abi.Events["ProposalThresholdSet"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, oldProposalThreshold, newProposalThreshold)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, oldProposalThreshold, newProposalThreshold)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address:     c.contract.Address(),
		Topics:      hashes,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	}, nil
}
func (c *Governance) EmitProposalThresholdSetEvent(oldProposalThreshold *big.Int, newProposalThreshold *big.Int) {
	log, err := c.ProposalThresholdSetEvent(oldProposalThreshold, newProposalThreshold)
	contracts.Require(err == nil, "Governance: emit ProposalThresholdSet event failed")
	c.stateDb.AddLog(log)
}

func (c *Governance) QuorumNumeratorUpdatedEvent(oldQuorumNumerator *big.Int, newQuorumNumerator *big.Int) (*types.Log, error) {
	event := c.abi.Events["QuorumNumeratorUpdated"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, oldQuorumNumerator, newQuorumNumerator)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, oldQuorumNumerator, newQuorumNumerator)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address:     c.contract.Address(),
		Topics:      hashes,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	}, nil
}
func (c *Governance) EmitQuorumNumeratorUpdatedEvent(oldQuorumNumerator *big.Int, newQuorumNumerator *big.Int) {
	log, err := c.QuorumNumeratorUpdatedEvent(oldQuorumNumerator, newQuorumNumerator)
	contracts.Require(err == nil, "Governance: emit QuorumNumeratorUpdated event failed")
	c.stateDb.AddLog(log)
}

func (c *Governance) VoteCastEvent(voter common.Address, proposalId *big.Int, support uint8, weight *big.Int, reason string) (*types.Log, error) {
	event := c.abi.Events["VoteCast"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, voter, proposalId, support, weight, reason)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, voter, proposalId, support, weight, reason)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address:     c.contract.Address(),
		Topics:      hashes,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	}, nil
}
func (c *Governance) EmitVoteCastEvent(voter common.Address, proposalId *big.Int, support uint8, weight *big.Int, reason string) {
	log, err := c.VoteCastEvent(voter, proposalId, support, weight, reason)
	contracts.Require(err == nil, "Governance: emit VoteCast event failed")
	c.stateDb.AddLog(log)
}

func (c *Governance) VotingDelaySetEvent(oldVotingDelay *big.Int, newVotingDelay *big.Int) (*types.Log, error) {
	event := c.abi.Events["VotingDelaySet"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, oldVotingDelay, newVotingDelay)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, oldVotingDelay, newVotingDelay)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address:     c.contract.Address(),
		Topics:      hashes,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	}, nil
}
func (c *Governance) EmitVotingDelaySetEvent(oldVotingDelay *big.Int, newVotingDelay *big.Int) {
	log, err := c.VotingDelaySetEvent(oldVotingDelay, newVotingDelay)
	contracts.Require(err == nil, "Governance: emit VotingDelaySet event failed")
	c.stateDb.AddLog(log)
}

func (c *Governance) VotingPeriodSetEvent(oldVotingPeriod *big.Int, newVotingPeriod *big.Int) (*types.Log, error) {
	event := c.abi.Events["VotingPeriodSet"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, oldVotingPeriod, newVotingPeriod)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, oldVotingPeriod, newVotingPeriod)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address:     c.contract.Address(),
		Topics:      hashes,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	}, nil
}
func (c *Governance) EmitVotingPeriodSetEvent(oldVotingPeriod *big.Int, newVotingPeriod *big.Int) {
	log, err := c.VotingPeriodSetEvent(oldVotingPeriod, newVotingPeriod)
	contracts.Require(err == nil, "Governance: emit VotingPeriodSet event failed")
	c.stateDb.AddLog(log)
}
