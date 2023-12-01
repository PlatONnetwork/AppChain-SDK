pragma solidity ^0.8.20;

interface IDepositHandler {
    function onStateReceive(uint256, /* id */ address sender, bytes calldata data) external;

    function withdraw(address recipient, uint256 amount) external;
}
