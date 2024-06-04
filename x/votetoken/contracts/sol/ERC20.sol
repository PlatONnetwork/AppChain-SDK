pragma solidity ^0.8.7;
import "./ERC20.sol";
import "./IERC20Metadata.sol";
import "./Ownable.sol";
interface ERC20 is Ownable, IERC20, IERC20Metadata {
    function increaseAllowance(address spender, uint256 addedValue) external returns (bool);
    function decreaseAllowance(address spender, uint256 subtractedValue) external returns (bool);
    function mint(address account, uint256 amount) external;
    function burn(address account, uint256 amount) external;
//    function beforeTokenTransfer(address from, address to, uint256 amount) external virtual;
//    function afterTokenTransfer(address from, address to,uint256 amount) external virtual;
}