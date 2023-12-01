pragma solidity ^0.8.22;

import "./IWithdrawManager.sol";

contract WithdrawManager is IWithdrawManager {
    function onStateReceive(uint256 id, address sender, bytes calldata data) external {}

    function deposit(address recipient, uint256 amount) external {}
}
