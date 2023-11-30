pragma solidity ^0.8.20;

import "./IVRFHandler.sol";

contract VRFHandler is IVRFHandler {
    function pushNonceAndProof(bytes calldata nonceAndProof) external {}
}
