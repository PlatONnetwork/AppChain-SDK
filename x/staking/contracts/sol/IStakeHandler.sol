pragma solidity ^0.8.20;

import "../IL1StateReceiver.sol";

struct ValidatorInit {
    address benefit;
    // uint256[2] signerkey; // ecdsa keySize: 64 bytes
    // uint128[3] blsKey; // bls keySize: 48 bytes
    uint256 stakePower;
    uint256 delegationPower;
    bytes signerKey; // ecdsa keySize: 64 bytes
    uint256[2] blsKey; // bls keySize: 48 bytes
}

interface IStakeHandler is IL1StateReceiver {
    event Slashed(uint256 indexed exitId, address[] validators, uint256[] amounts);
    event StakeAdded(address indexed validator, uint256 amount);
    event DelegationAdded(address indexed delegater, address indexed validator, uint256 amount);
    event UnStaked(address indexed validator, uint256 amount);
    event UnDelegated(address indexed delegater, address indexed validator, uint256 amount);
    event StakeWithdrawalRegistered(address indexed validator, uint256 amount);
    event StakeWithdrawal(address indexed validator, uint256 amount);
    event DelegateWithdrawalRegistered(address indexed delegater, address indexed validator, uint256 amount);
    event DelegateWithdrawal(address indexed delegater, address indexed validator, uint256 amount);

    /// @notice initialises slashing process
    /// @dev system call,
    /// @dev given list of validators are slashed on L2
    /// subsequently after their stake is slashed on L1
    function slash() external;

    /// @notice allows a validator to announce their intention to withdraw a given amount of tokens
    /// @dev initializes a waiting period before the tokens can be withdrawn
    function unstake(address validator, uint256 amount) external;

    function undelegate(address validator, uint256 amount) external;

    /// @notice allows a validator to complete a withdrawal
    function withdrawUnstake(address validator) external; // only owner of validator call

    function withdrawUndelegate(address validator) external; // only delegater call

    /**
     * @notice Calculates how much can be withdrawn for account in this epoch.
     * @param validator The account to calculate amount for
     * @return Amount withdrawable
     */
    function withdrawableOfStake(address validator) external view returns (uint256);

    function withdrawableOfDelegate(address validator, address delegater) external view returns (uint256);

    /**
     * @notice Calculates how much is yet to become withdrawable for account.
     * @param validator The validator to calculate amount for
     * @return Amount not yet withdrawable
     */
    function pendingWithdrawalsOfStake(address validator) external view returns (uint256);

    function pendingWithdrawalsOfDelegate(address validator, address delegater) external view returns (uint256);
}
