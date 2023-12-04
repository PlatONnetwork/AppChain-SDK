pragma solidity ^0.8.20;

interface IWithdrawManager {
    event L2MintableCoinDeposit(address indexed recipient, address depositor, uint256 amount);
    event L2MintableCoinWithdraw(address indexed recipient, address withdrawer, uint256 amount);

    function onStateReceive(uint256, /* id */ address sender, bytes calldata data) external;

    function deposit(address recipient, uint256 amount) external;
}
