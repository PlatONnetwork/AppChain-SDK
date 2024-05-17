pragma solidity ^0.8.7;
import "./IVotes.sol";
import "./ERC20.sol";

    struct Checkpoint {
        uint64 fromBlock;
        uint224 votes;
    }

interface ERC20Vote is IVotes, ERC20{
    /**
 * @dev Get the `pos`-th checkpoint for `account`.
     */
    function checkpoints(address account, uint32 pos) external view virtual returns (Checkpoint memory);

    /**
     * @dev Get number of checkpoints for `account`.
     */
    function numCheckpoints(address account) external view virtual returns (uint32);

    // function _mint(address account, uint256 amount) internal virtual;
    // function _burn(address account, uint256 amount) internal virtual;
    // function _transfer(address from, address to, uint256 amount) internal virtual;


}
