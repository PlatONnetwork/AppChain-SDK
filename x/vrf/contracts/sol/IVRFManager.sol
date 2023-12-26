pragma solidity ^0.8.20;

interface IVRFManager {
    event VRFNonceAdded(uint256 indexed block, bytes nonce);

    /// @notice push vrf nonce
    /// @dev validator call,
    /// @param nonceAndProof 81 byte, nonce and proof, flag |nonce |proof, 1byte|32byte|48byte
    function pushNonceAndProof(bytes calldata nonceAndProof) external;
}
