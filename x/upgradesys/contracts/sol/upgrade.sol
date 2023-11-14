// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

contract upgrade {
    struct UpgradeCommitment {
        address origin;
        uint256 validNumber;
        address upgrade;
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
    function initialized(address implement) external {}
    function commit(
        UpgradeCommitment calldata commitment,
        uint64 index,
        bytes32[] calldata proof,
        QuorumCert calldata qc
    ) external {}
}
