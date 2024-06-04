pragma solidity ^0.8.20;

interface IUpgrade {
    enum Status {
        READY,
        DONE
    }

    struct Module {
        string moduleName;
        uint64 version;
    }

    struct Plan {
        string name;
        Module[] modules;
        string info;
        uint64 height;
        Status status;
    }

    function setOwner(address newOwner) external;
    function getOwner() external view returns (address);

    function addUpgradePlan(Plan calldata plan) external;
    function getUpgradePlan(uint64 height) external view returns (Plan[] memory plan);
    function setUpgradePlanDone(uint64 height) external;
}
