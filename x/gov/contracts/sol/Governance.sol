pragma solidity ^0.8.7;

struct GovernanceParams{
    uint64 voteDelay;
    uint64 votePeriod;
    uint256 proposalThreshold;
    uint256 emergencyCouncilVoteRatio;
    uint256 constitutionCouncilVoteRatio;
    uint256 nonConstitutionCouncilVoteRatio;
    uint256 constitutionVoteRatio;
    uint256 nonConstitutionVoteRatio;
}

enum ProposalType{
    L1Constitution,
    L1NonConstitution,
    L2Constitution,
    L2NonConstitution,
    L1Emergency,
    L2Emergency
}

enum ProposalState {
    Pending,
    Active,
    Canceled,
    Defeated,
    Succeeded,
    Executed
}

struct Proposal {
    ProposalType proposalType;
    address[] targets;
    uint256[] values;
    bytes[] calldatas;
    bytes32 descriptionHash;
}

struct ProposalCore {
    Proposal proposal;
    uint64 startVote;
    uint64 endVote;
    ProposalState state;
    ProposalVote vote;
}

enum VoteType {
    Against,
    For,
    Abstain
}

struct ProposalVote {
    uint256 voteRatio;
    uint256 againstVotes;
    uint256 forVotes;
    uint256 abstainVotes;
    mapping(address => bool) hasVoted;
}

interface Governance {
    event ProposalCreated(uint256 proposalId, address proposer, ProposalType proposalType, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 startBlock, uint256 endBlock, string description);
    event ProposalCanceled(uint256 proposalId);
    event ProposalExecuted(uint256 proposalId);
    event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason);
    event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay);
    event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod);
    event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold);
    event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator);

    function name() external pure returns (string memory);
    function version() external pure returns (string memory);
    function getVotes(address account, uint256 blockNumber) external view returns (uint256);
    function quorumNumerator() external view returns (uint256);
    function quorumDenominator() external view returns (uint256);
    function quorum(uint256 blockNumber) external view returns (uint256);
    function updateQuorumNumerator(uint256 newQuorumNumerator) external;
    function votingDelay() external view  returns (uint256);
    function votingPeriod() external view  returns (uint256);
    function setVotingDelay(uint256 newVotingDelay) external;
    function setVotingPeriod(uint256 newVotingPeriod) external;
    function setProposalThreshold(uint256 newProposalThreshold) external;

    function state(uint256 proposalId) external view  returns (ProposalState);
    function proposalSnapshot(uint256 proposalId) external view returns (uint256);
    function proposalDeadline(uint256 proposalId) external view returns (uint256);
    function proposalThreshold() external view  returns (uint256);

    function propose(
        ProposalType proposalType,
        address[] memory targets,
        uint256[] memory values,
        bytes[] memory calldatas,
        string memory description
    ) external returns (uint256 proposalId);

    function execute(
        ProposalType proposalType,
        address[] memory targets,
        uint256[] memory values,
        bytes[] memory calldatas,
        bytes32 descriptionHash
    ) external payable  returns (uint256 proposalId);

    function cancel(
        ProposalType proposalType,
        address[] memory targets,
        uint256[] memory values,
        bytes[] memory calldatas,
        bytes32 descriptionHash
    ) external  returns (uint256);

    function castVote(uint256 proposalId, uint8 support) external  returns (uint256 balance);
    function castVoteBySig(uint256 proposalId, uint8 support, uint8 v, bytes32 r, bytes32 s) external returns (uint256);
    function hashProposal(
        ProposalType proposalType,
        address[] memory targets,
        uint256[] memory values,
        bytes[] memory calldatas,
        bytes32 descriptionHash
    ) external pure  returns (uint256);

    function hasVoted(uint256 proposalId, address account) external view  returns (bool);

    function getExecutableProposal() external view returns( Proposal[]memory);
}
