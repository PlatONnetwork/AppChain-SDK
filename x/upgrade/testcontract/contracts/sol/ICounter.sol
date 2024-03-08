pragma solidity ^0.8.20;

interface ICounter {
    function name() external view returns(string memory);
    function count() external view returns(uint64);
    function incr() external;

    // Version 1
    function dec() external;
}
