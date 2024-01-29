# StateSender

StateSender 负责产生 L2-L1 的事件，目前有 Deposit 和 Staking 使用。

```solidity title="x/statesender/contracts/sol/L2StateSender.sol"
contract L2StateSender  {
    event L2StateSynced(uint256 indexed id, address indexed sender, address indexed receiver, bytes callData);

    /**
     * @notice Emits an event which is indexed by v3 validators and submitted as a commitment on L1
     * allowing for lazy execution
     * @param receiver Address of the message recipient on L1
     * @param data Data to use in message call to recipient
     */
    function syncState(address receiver, bytes calldata data) external
}
```

每个向 L1 提交信息的需要调用 syncState，后触发 L2StateSynced 事件。L2StateSynced 事件由 Checkpoint 模块提交到 L1