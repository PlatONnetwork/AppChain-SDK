pragma solidity ^0.8.20;

import "./IStakeHandler.sol";

contract StakeHandler is IStakeHandler {
    function onStateReceive(uint256 id, address sender, bytes calldata data) external {}

    /// @notice initialises slashing process
    /// @dev system call,
    /// @dev given list of validators are slashed on L2
    /// subsequently after their stake is slashed on L1
    function slash() external {}

    /// @notice allows a validator to announce their intention to withdraw a given amount of tokens
    /// @dev initializes a waiting period before the tokens can be withdrawn
    /// @param validator the validator unstake for
    /// @param amount the amount unstake for
    function unstake(address validator, uint256 amount) external {}

    /// @notice Allow the client to announce their intention to withdraw a fixed quantity of tokens to the validator from the stake
    /// @dev initializes a waiting period before the tokens can be withdrawn
    /// @param validator the validator undelegate for
    /// @param amount the amount undelegate for
    function undelegate(address validator, uint256 amount) external {}

    /// @notice allows a validator to complete a withdrawal
    /// @param validator The validator to withdraw amount for
    function withdrawUnstake(address validator) external {} // only owner of validator call

    /// @notice allows a delegator to complete a withdrawal with validator
    /// @param validator The validator to withdraw amount for
    function withdrawUndelegate(address validator) external {} // only delegator call

    /// @notice Calculates how much can be withdrawn for account in this epoch.
    /// @param validator The account to calculate amount for
    /// @return Amount withdrawable
    function withdrawableOfStake(address validator) external view returns (uint256) {
        return 0;
    }

    /// @notice Calculates how much can be withdrawn for account in this epoch.
    /// @param validator The validator to calculate amount for
    /// @param validator The delegator to calculate amount for
    /// @return Amount withdrawable
    function withdrawableOfDelegate(address validator, address delegator) external view returns (uint256) {
        return 0;
    }

    /// @notice Calculates how much is yet to become withdrawable for account.
    /// @param validator The validator to calculate amount for
    /// @return Amount not yet withdrawable
    function pendingWithdrawalsOfStake(address validator) external view returns (uint256) {
        return 0;
    }

    /// @notice Calculates how much is yet to become withdrawable for account.
    /// @param validator The validator to calculate amount for
    /// @param delegator The delegator to calculate amount for
    /// @return Amount not yet withdrawable
    function pendingWithdrawalsOfDelegate(address validator, address delegator) external view returns (uint256) {
        return 0;
    }

    /// @notice Verify the aggregated signature of the validators.
    /// @param blockNumber The number of the block to which the validator list belongs in a period
    /// @param validatorIndexs The index in the list of validators for aggregate signatures
    /// @param data Signature Data Hash
    /// @param signatues Aggregated signatures for validators
    /// @return True is successful
    function verifyAggregateSignature(
        uint256 blockNumber,
        uint256[] calldata validatorIndexs,
        bytes32 data,
        bytes calldata signatues
    ) external view returns (bool) {
        return true;
    }

    /// @notice Verify the aggregated signature of the validators.
    /// @param validators List of validators for aggregated signatures
    /// @param data Signature Data Hash
    /// @param signatues Aggregated signatures for validators
    /// @return True is successful
    function verifyAggregateSignatureByValidators(address[] calldata validators, bytes32 data, bytes calldata signatues)
        external
        view
        returns (bool)
    {
        return true;
    }

    /// @notice Query the delegation information of the delegator on the validators
    /// @dev Query the delegation information of the delegator on these validators based on the addr list of validators
    /// @param validators addr of validators
    /// @param delegator the delegator
    /// @return DelegationInfo array for query
    function getDelegationsWithValidator(address[] calldata validators, address delegator)
        external
        view
        returns (DelegationInfo[] memory)
    {
        return new DelegationInfo[](0);
    }

    /// @notice Query the list of validators for a certain period
    /// @dev For the convenience of expanding the list of validators with multiple period properties
    /// @param periodType represents a period of a certain type
    /// @param period represents the number of intervals
    /// @return validator address array
    function getValidatorAddrs(uint8 periodType, uint256 period) external view returns (address[] memory) {
        return new address[](0);
    }

    /// @notice Query the list of all validators
    /// @dev Support pagination to query the list of all validators
    /// @param start represents the starting query ID. When passing empty bytes, it defaults to starting from the first Id
    /// @param size page size
    /// @return bytes of next start
    /// @return ValidatorInfo array for query
    function getValidators(bytes calldata start, uint256 size)
        external
        view
        returns (bytes memory, ValidatorInfo[] memory)
    {
        return (bytes(""), new ValidatorInfo[](0));
    }

    /// @notice Query the list of validators by addrs
    /// @dev Support to query the list of validators by addr of validators
    /// @param validators addr of validators
    /// @return ValidatorInfo array for query
    function getValidatorsWithAddr(address[] calldata validators) external view returns (ValidatorInfo[] memory) {
        return new ValidatorInfo[](0);
    }

    /// @notice Query the list of information on the number of sealed blocks of validators for a certain period
    /// @dev For the convenience of expanding the list of validators with multiple period properties
    /// @param periodType represents a period of a certain type
    /// @param period represents the number of intervals
    /// @return BlocksOfValidator array for query
    function getBlocksOfValidators(uint8 periodType, uint256 period)
        external
        view
        returns (BlocksOfValidator[] memory)
    {
        return new BlocksOfValidator[](0);
    }
}
