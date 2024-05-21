package contracts

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/db"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/db/container"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/ecdsa"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/gov/contracts/erc20vote"
	contracts2 "github.com/PlatONnetwork/AppChain-SDK/x/gov/contracts/ownable"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/event"
	abi2 "github.com/umbracle/ethgo/abi"
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

var (
	nameKey         = []byte("name")
	proposalKey     = []byte("proposal")
	BALLOT_TYPEHASH = crypto.Keccak256Hash([]byte("Ballot(uint256 proposalId,uint8 support)"))
)

type EIP712 interface {
	HashTypedData(structHash common.Hash) common.Hash
}
type ProposalCore struct {
	VoteStart BlockNumber
	VoteEnd   BlockNumber
	Executed  bool
	Canceled  bool
}

type Storage struct {
	Name              *db.Base[string]
	Version           *db.Base[string]
	Proposals         *container.Map[ProposalCore]
	QuorumNumerator   *db.Base[*big.Int]
	ProposalVotes     *container.Map[ProposalVote]
	HasVoted          *container.Map[*container.Map[bool]]
	VoteDelay         *db.Base[*big.Int]
	VotePeriod        *db.Base[*big.Int]
	ProposalThreshold *db.Base[*big.Int]
}

type Governance struct {
	abi           *abi.ABI
	abis          map[uint64]*abi.ABI
	methodEntry   map[string]func([]byte) ([]byte, error)
	methodEntries map[uint64]map[string]func([]byte) ([]byte, error)
	readOnly      bool
	contract      *vm.Contract
	evm           *vm.EVM
	burner        contracts.Burn
	stateDb       *contracts.StateDB
	context       *contracts.Context
	fallback      func(input []byte) ([]byte, error)
	storage       Storage
	voteCaller    *erc20vote.ERC20VoteCaller
	//eip712        EIP712
	*contracts2.Ownable
}

func NewGovernance(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*Governance, error) {
	s := &Governance{
		abi:           nil,
		abis:          make(map[uint64]*abi.ABI),
		methodEntry:   make(map[string]func([]byte) ([]byte, error)),
		methodEntries: make(map[uint64]map[string]func([]byte) ([]byte, error)),
		evm:           evm,
		contract:      contract,
		burner:        contracts.NewBurner(contract),
		stateDb:       contracts.NewStateDB(evm, contract),
		context:       contracts.NewContext(evm, contract),
		readOnly:      readOnly,
	}
	store := db.NewStore([]byte{}, contract.Address(), s.stateDb)
	s.storage = Storage{
		Name:      db.NewBase[string](nameKey, store),
		Proposals: container.NewMap[ProposalCore](proposalKey, common.Address{}, s.stateDb),
	}
	s.initABI()
	s.initMethodEntry()
	s.loadMethodABI()
	return s, nil
}

func (c *Governance) Init(name, version string, voteDelay, votePeriod, quorumNumerator, proposalThreshold *big.Int, owner common.Address) {
	c.storage.Name.MustSet(name)
	c.storage.Version.MustSet(version)
	c.storage.VoteDelay.MustSet(voteDelay)
	c.storage.VotePeriod.MustSet(votePeriod)
	c.storage.QuorumNumerator.MustSet(quorumNumerator)
	c.storage.ProposalThreshold.MustSet(proposalThreshold)
	c.Ownable.Init(owner)
}
func (c *Governance) onlyGovernance() {
	owner, _ := c.Ownable.Owner()
	contracts.Require(c.context.Caller() == owner, "Governance: caller is not Governor")
}
func (c *Governance) Name() (string, error) {
	return c.storage.Name.MustGet(), nil
}

func (c *Governance) Version() (string, error) {
	return c.storage.Version.MustGet(), nil
}

func (c *Governance) GetExecutableProposal() ([]Proposal, error) {
	panic("implement")
}

func (c *Governance) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	hasVoted := c.storage.HasVoted.MustGet(proposalId)
	return hasVoted.MustGet(account), nil
}

func (c *Governance) HashProposal(proposalType uint8, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash common.Hash) (*big.Int, error) {
	data, _ := abi2.Encode([]interface{}{proposalType, targets, values, calldatas, descriptionHash}, abi2.MustNewType(""))
	return crypto.Keccak256Hash(data).Big(), nil
}

func (c *Governance) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return c.storage.Proposals.MustGet(proposalId).VoteEnd.getDeadline(), nil
}

func (c *Governance) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return c.storage.Proposals.MustGet(proposalId).VoteStart.getDeadline(), nil
}

func (c *Governance) ProposalThreshold() (*big.Int, error) {
	return big.NewInt(0), nil
}

func (c *Governance) State(proposalId *big.Int) (uint8, error) {
	proposal := c.storage.Proposals.MustGet(proposalId)
	if proposal.Executed {
		return Executed, nil
	}
	if proposal.Canceled {
		return Canceled, nil
	}
	snapshot, _ := c.ProposalSnapshot(proposalId)
	contracts.Require(snapshot.Cmp(big.NewInt(0)) == 0, "Governance: unknown proposal id")

	if snapshot.Uint64() > c.context.BlockNumber().Uint64() {
		return Pending, nil
	}
	deadline, _ := c.ProposalDeadline(proposalId)
	if deadline.Uint64() >= c.context.BlockNumber().Uint64() {
		return Active, nil
	}
	if c.quorumReached(proposalId) && c.voteSucceeded(proposalId) {
		return Succeeded, nil
	} else {
		return Defeated, nil
	}

}
func (c *Governance) quorumReached(proposalId *big.Int) bool {
	proposalVote := c.storage.ProposalVotes.MustGet(proposalId)
	blockNumber, _ := c.ProposalSnapshot(proposalId)
	quorum, _ := c.Quorum(blockNumber)
	return quorum.Cmp(new(big.Int).Add(proposalVote.ForVotes, proposalVote.AbstainVotes)) >= 0

}

func (c *Governance) voteSucceeded(proposalId *big.Int) bool {
	proposalVote := c.storage.ProposalVotes.MustGet(proposalId)
	return proposalVote.ForVotes.Cmp(proposalVote.AgainstVotes) > 0
}

func (c *Governance) countVote(proposalId *big.Int, account common.Address, support uint8, weight *big.Int) {
	proposalVote := c.storage.ProposalVotes.MustGet(proposalId)
	hasVoted := c.storage.HasVoted.MustGet(proposalId)
	contracts.Require(!hasVoted.MustGet(account), "Governance: vote already cast")
	hasVoted.MustSet(account, true)
	if support == Against {
		proposalVote.AgainstVotes.Add(proposalVote.AgainstVotes, weight)
	} else if support == For {
		proposalVote.ForVotes.Add(proposalVote.ForVotes, weight)
	} else if support == Abstain {
		proposalVote.AbstainVotes.Add(proposalVote.AbstainVotes, weight)
	} else {
		contracts.Require(false, "Governance: invalid value for enum VoteType")
	}
	c.storage.ProposalVotes.MustSet(proposalId, proposalVote)
}

func (c *Governance) Cancel(proposalType uint8, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash common.Hash) (*big.Int, error) {
	proposalId, _ := c.HashProposal(proposalType, targets, values, calldatas, descriptionHash)
	status, _ := c.State(proposalId)
	contracts.Require(status != Canceled && status != Expired || status != Executed, "Governance: proposal not active")
	proposal := c.storage.Proposals.MustGet(proposalId)
	proposal.Canceled = true
	c.storage.Proposals.MustSet(proposalId, proposal)
	c.EmitProposalCanceledEvent(proposalId)
	return proposalId, nil
}

func (c *Governance) CastVote(proposalId *big.Int, support uint8) (*big.Int, error) {
	voter := c.context.Caller()
	return c.castVote(proposalId, voter, support, ""), nil
}
func (c *Governance) castVote(proposalId *big.Int, account common.Address, support uint8, reason string) *big.Int {
	proposal := c.storage.Proposals.MustGet(proposalId)
	status, _ := c.State(proposalId)
	contracts.Require(status == Active, "Governance: vote not currently active")
	weight, _ := c.GetVotes(account, new(big.Int).SetUint64(uint64(proposal.VoteStart)))
	c.countVote(proposalId, account, support, weight)
	c.EmitVoteCastEvent(account, proposalId, support, weight, reason)
	return weight
}

func (c *Governance) CastVoteBySig(proposalId *big.Int, support uint8, v uint8, r common.Hash, s common.Hash) (*big.Int, error) {
	voter := ecdsa.Recover(common.Hash{}, v, r, s)
	return c.castVote(proposalId, voter, support, ""), nil
}

func (c *Governance) Execute(proposalType uint8, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash common.Hash) (*big.Int, error) {
	proposalId, _ := c.HashProposal(proposalType, targets, values, calldatas, descriptionHash)
	status, _ := c.State(proposalId)
	contracts.Require(status == Succeeded || status == Queued, "Governance: proposal not successful")
	proposal := c.storage.Proposals.MustGet(proposalId)
	proposal.Executed = true
	c.storage.Proposals.MustSet(proposalId, proposal)
	c.EmitProposalExecutedEvent(proposalId)
	c.execute(proposalType, targets, values, calldatas, descriptionHash)
	return proposalId, nil
}

func (c *Governance) execute(proposalType uint8, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash common.Hash) {
	for i := 0; i < len(targets); i++ {
		_, err := contracts.Call(c.evm, c.contract, targets[i], calldatas[i], c.context.Gas(), values[i])
		contracts.Require(err == nil, "Governance: call reverted without message")
	}
}

func (c *Governance) Propose(proposalType uint8, targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*big.Int, error) {
	votes, _ := c.GetVotes(c.context.Caller(), c.context.BlockNumber())
	threshold, _ := c.ProposalThreshold()
	contracts.Require(votes.Cmp(threshold) >= 0, "Governance: proposer votes below proposal threshold")
	proposalId, _ := c.HashProposal(proposalType, targets, values, calldatas, crypto.Keccak256Hash([]byte(description)))
	contracts.Require(len(targets) == len(values), "Governance: invalid proposal length")
	contracts.Require(len(targets) == len(calldatas), "Governance: invalid proposal length")
	contracts.Require(len(targets) > 0, "Governance: empty proposal")
	proposal := c.storage.Proposals.MustGet(proposalId)
	contracts.Require(proposal.VoteStart == 0, "Governance: proposal already exists")
	snapshot := c.context.BlockNumber().Uint64() + c.storage.VoteDelay.MustGet().Uint64()
	deadline := snapshot + c.storage.VotePeriod.MustGet().Uint64()
	proposal.VoteStart = BlockNumber(snapshot)
	proposal.VoteEnd = BlockNumber(deadline)
	c.EmitProposalCreatedEvent(proposalId, c.context.Caller(), proposalType, targets, values, make([]string, len(targets)), calldatas, new(big.Int).SetUint64(snapshot), new(big.Int).SetUint64(deadline), description)
	return proposalId, nil
}

func (c *Governance) GetVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return c.voteCaller.GetPastVotes(account, blockNumber)
}

func (c *Governance) Quorum(blockNumber *big.Int) (*big.Int, error) {
	totalSupply, _ := c.voteCaller.GetPastTotalSupply(blockNumber)
	quorumNumerator, _ := c.QuorumNumerator()
	quorumDenominator, _ := c.QuorumDenominator()
	totalSupply.Mul(totalSupply, quorumNumerator)
	return new(big.Int).Div(totalSupply, quorumDenominator), nil
}

func (c *Governance) QuorumDenominator() (*big.Int, error) {
	return c.storage.QuorumNumerator.MustGet(), nil
}

func (c *Governance) QuorumNumerator() (*big.Int, error) {
	return big.NewInt(100), nil
}

func (c *Governance) UpdateQuorumNumerator(newQuorumNumerator *big.Int) error {
	oldQuorumNumerator := c.storage.QuorumNumerator.MustGet()
	contracts.Require(newQuorumNumerator.Cmp(oldQuorumNumerator) > 0, "Governance: quorumNumerator over quorumDenominator")
	c.storage.QuorumNumerator.MustSet(newQuorumNumerator)
	c.EmitQuorumNumeratorUpdatedEvent(oldQuorumNumerator, newQuorumNumerator)
	return nil
}

func (c *Governance) VotingDelay() (*big.Int, error) {
	return c.storage.VoteDelay.MustGet(), nil
}

func (c *Governance) VotingPeriod() (*big.Int, error) {
	return c.storage.VotePeriod.MustGet(), nil
}

func (c *Governance) SetVotingDelay(newVotingDelay *big.Int) error {
	c.EmitVotingDelaySetEvent(c.storage.VoteDelay.MustGet(), newVotingDelay)
	c.storage.VoteDelay.MustSet(newVotingDelay)
	return nil
}

func (c *Governance) SetVotingPeriod(newVotingPeriod *big.Int) error {
	contracts.Require(newVotingPeriod.Cmp(big.NewInt(0)) > 0, "Governance: voting period too low")
	c.EmitVotingPeriodSetEvent(c.storage.VotePeriod.MustGet(), newVotingPeriod)
	c.storage.VotePeriod.MustSet(newVotingPeriod)
	return nil
}
func (c *Governance) SetProposalThreshold(newProposalThreshold *big.Int) error {
	c.EmitProposalThresholdSetEvent(c.storage.ProposalThreshold.MustGet(), newProposalThreshold)
	c.storage.ProposalThreshold.MustSet(newProposalThreshold)
	return nil
}
