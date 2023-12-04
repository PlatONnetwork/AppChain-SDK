pragma solidity ^0.8.20;

interface IDepositHandler {
    event L2CoinDeposit(address indexed recipient, address depositor, uint256 amount);
    event L2CoinWithdraw(address indexed recipient, address withdrawer, uint256 amount);

    function onStateReceive(uint256, /* id */ address sender, bytes calldata data) external;

    function withdraw(address recipient, uint256 amount) external;
}
