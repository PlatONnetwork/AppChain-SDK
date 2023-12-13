pragma solidity ^0.8.20;

interface IRewardManager {
    event RewardDistributed(uint256 indexed epochId, uint256 totalReward);
    event EpochReward(uint256 indexed epochId, address[] validators, uint256[] amounts);
    event BlockReward(uint256 indexed epochId, address[] validators, uint256[] amounts);
    event ValidatorRewardWithdrawal(address indexed validator, uint256 amount, address caller);
    event DelegatorRewardWithdrawal(address indexed validator, uint256 amount, address caller);

    /// @notice withdraws pending rewards for the sender (validator)
    function withdrawValidatorReward(address validator) external; // only owner of validator call

    function withdrawDelegatorReward(address validator) external; // only delegator call

    /// @notice returns the total reward (epoch reward and blocks reward) paid for the given epoch
    function paidRewardPerEpoch(uint256 epochId) external view returns (uint256);

    /// @notice returns the pending reward for the given account(validator)
    function pendingValidatorRewards(address validator) external view returns (uint256);

    function pendingDelegatorRewards(address validator) external view returns (uint256);
}
