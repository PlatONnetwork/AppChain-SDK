pragma solidity ^0.8.20;

struct PeriodEdge {
    uint256 startBlock;
    uint256 endBlock;
}

interface IStageManager {
    /// @notice Query the list of period edge for a certain period
    /// @dev For the convenience of expanding the list of datas with multiple period properties
    /// @param periodType represents a period of a certain type
    /// @param period represents the number of intervals
    /// @return PeriodEdge array for query
    function getPeriodEdge(uint8 periodType, uint256 period) external view returns (PeriodEdge memory);

    /// @notice Query the list of all period edges
    /// @dev Support pagination to query the list of period edges
    /// @param periodType represents a period of a certain type
    /// @param start represents the starting query ID. When a value of 0 is passed, it defaults to starting from the first Id
    /// @param size page size
    /// @return uint256 of next start
    /// @return PeriodEdge array for query
    function getPeriodEdges(uint8 periodType, uint256 start, uint256 size)
        external
        view
        returns (uint256, uint256[] memory, PeriodEdge[] memory);

    /// @notice Query the period number where the block height is located
    /// @dev For the convenience of expanding the expansion of serial number with multiple period properties
    /// @param periodType represents a period of a certain type
    /// @param blockNumber represents the block sequence number
    /// @return uint256 the period number for query
    function getPeriodByBlockNumber(uint8 periodType, uint256 blockNumber) external view returns (uint256);
}
