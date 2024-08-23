pragma solidity ^0.8.7;
interface Ownable {
    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);
    function owner() external returns (address);
    function renounceOwnership() external;
    function transferOwnership(address newOwner) external;
}