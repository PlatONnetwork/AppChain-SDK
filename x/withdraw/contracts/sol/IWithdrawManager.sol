pragma solidity ^0.8.20;

import "../../../statereceiver/contracts/sol/IL1StateReceiver.sol";

interface IWithdrawManager is IL1StateReceiver {
    event L2MintableCoinDeposit(address indexed recipient, address depositor, uint256 amount);
    event L2MintableCoinWithdraw(address indexed recipient, address withdrawer, uint256 amount);

    function deposit(address recipient, uint256 amount) external;
}
