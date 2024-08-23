pragma solidity ^0.8.20;

interface ICounter {
    function name() external view returns(string memory);
    function count() external view returns(uint64);
    function incr() external;

    // Version 1
    function dec() external;

    // Version 2
    function add(uint256 n) external;

    // Version 3
    function minus(uint256 n) external;
}
