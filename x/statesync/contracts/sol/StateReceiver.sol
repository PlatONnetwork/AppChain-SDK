// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;
    struct StateSync {
        uint256 id;
        address sender;
        address receiver;
        bytes data;
    }

    struct StateSyncCommitment {
        uint256 startId;
        uint256 endId;
        bytes32 root;
    }
    struct BitArray {
        uint32 bits;
        uint64[] Elems;
    }
    struct QuorumCert {
        uint64 epoch;
        uint64 viewNumber;
        bytes32 blockHash;
        uint64 blockNumber;
        uint32 blockIndex;
        bytes32 extendHash;
        bytes signature;
        BitArray validatorSet;
    }
contract StateReceiver {


    event StateSyncResult(uint256 indexed counter, bool indexed status, bytes message);
    event NewCommitment(uint256 indexed startId, uint256 indexed endId, bytes32 root);

    function commit(
        StateSyncCommitment calldata commitment,
        uint64 index,
        byte32[] proof,
        QuorumCert calldata qc,
    ) external {}

    function execute(bytes32[] calldata proof, StateSync calldata obj) external {}
    function batchExecute(bytes32[][] calldata proofs, StateSync[] calldata objs) external{}
    function getExecutedId() external view returns(uint256){}
    function getStateSyncId() external returns(uint256){}
    function getRootByStateSyncId(uint256 id) external view returns (bytes32){}
    function getCommitmentByStateSyncId(uint256 id) public view returns (StateSyncCommitment memory){}
}