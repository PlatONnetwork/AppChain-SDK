pragma solidity ^0.8.7;

interface IL1StateReceiver {
    /**
     * @notice Called by exit helper when state is received from L1
     * @param sender Address of the sender on the root chain
     * @param data Data sent by the sender
     */
    function onStateReceive(uint256 id, address sender, bytes calldata data) external;
}
