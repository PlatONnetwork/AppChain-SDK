pragma solidity ^0.8.20;

import "./IRewardManager.sol";

contract RewardManager is IRewardManager {
    /// @notice withdraws pending rewards for the sender (validator)
    function withdrawValidatorReward(address validator) external {}

    function withdrawDelegaterReward(address validator) external {}

    /// @notice returns the total reward (epoch reward and blocks reward) paid for the given epoch
    function paidRewardPerEpoch(uint256 epochId) external view returns (uint256) {
        return 0;
    }

    /// @notice returns the pending reward for the given account(validator)
    function pendingValidatorRewards(address account) external view returns (uint256) {
        return 0;
    }

    function pendingDelegaterRewards(address account) external view returns (uint256) {
        return 0;
    }
}
