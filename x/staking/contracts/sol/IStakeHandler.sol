pragma solidity ^0.8.20;

import "../IL1StateReceiver.sol";

interface IStakeHandler is IL1StateReceiver {
    event Slashed(uint256 indexed exitId, address[] validators, uint256[] amounts);
    event StakeAdded(address indexed validator, uint256 amount);
    event DelegationAdded(address indexed delegator, address indexed validator, uint256 amount);
    event UnStaked(address indexed validator, uint256 amount);
    event UnDelegated(address indexed delegator, address indexed validator, uint256 amount);
    event StakeWithdrawalRegistered(address indexed validator, uint256 amount);
    event StakeWithdrawal(address indexed validator, uint256 amount);
    event DelegateWithdrawalRegistered(address indexed delegator, address indexed validator, uint256 amount);
    event DelegateWithdrawal(address indexed delegator, address indexed validator, uint256 amount);

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

    function withdrawUndelegate(address validator) external; // only delegator call

    /**
     * @notice Calculates how much can be withdrawn for account in this epoch.
     * @param validator The account to calculate amount for
     * @return Amount withdrawable
     */
    function withdrawableOfStake(address validator) external view returns (uint256);

    function withdrawableOfDelegate(address validator, address delegator) external view returns (uint256);

    /**
     * @notice Calculates how much is yet to become withdrawable for account.
     * @param validator The validator to calculate amount for
     * @return Amount not yet withdrawable
     */
    function pendingWithdrawalsOfStake(address validator) external view returns (uint256);

    function pendingWithdrawalsOfDelegate(address validator, address delegator) external view returns (uint256);

    /**
     * @notice Verify the aggregated signature of the validators.
     * @param blockNumber The number of the block to which the validator list belongs in a period
     * @param validatorIndexs The index in the list of validators for aggregate signatures
     * @param data Signature Data Hash
     * @param signatues Aggregated signatures for validators
     * @return True is successful
     */
    function verifyAggregateSignature(
        uint256 blockNumber,
        uint256[] calldata validatorIndexs,
        bytes32 data,
        bytes calldata signatues
    ) external view returns (bool);

    /**
     * @notice Verify the aggregated signature of the validators.
     * @param validators List of validators for aggregated signatures
     * @param data Signature Data Hash
     * @param signatues Aggregated signatures for validators
     * @return True is successful
     */
    function verifyAggregateSignatureByValidators(address[] calldata validators, bytes32 data, bytes calldata signatues)
        external
        view
        returns (bool);
}
