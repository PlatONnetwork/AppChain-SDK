pragma solidity ^0.8.20;

import "./IDepositHandler.sol";

contract DepositHandler is IDepositHandler {
    function onStateReceive(uint256 id, address sender, bytes calldata data) external {}

    function withdraw(address recipient, uint256 amount) external {}
}
