pragma solidity ^0.8.7;

import "./IVRFManager.sol";

contract VRFManager is IVRFManager {
    function pushNonceAndProof(bytes calldata nonceAndProof) external {}
}
