pragma solidity ^0.8.7;

import "../../../statereceiver/contracts/sol/IL1StateReceiver.sol";

interface IDepositHandler is IL1StateReceiver {
    event L2CoinDeposit(address indexed recipient, address depositor, uint256 amount);
    event L2CoinWithdraw(address indexed recipient, address withdrawer, uint256 amount);

    function withdraw(address recipient, uint256 amount) external;
}
