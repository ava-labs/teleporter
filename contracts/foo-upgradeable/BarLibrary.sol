pragma solidity 0.8.25;

library BarLibrary {
    struct Storage {
        uint256 varBar;
    }

    function setBar(Storage storage self, uint256 newValue) external {
        self.varBar = newValue;
    }

    function getBar(Storage storage self) external view returns (uint256) {
        return self.varBar;
    }
}