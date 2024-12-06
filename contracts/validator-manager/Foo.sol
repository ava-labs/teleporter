pragma solidity 0.8.25;

import "./Bar.sol";

contract Foo {
    using Bar for Bar.Storage;

    uint256 internal varFoo;
    Bar.Storage private barStorage;

    function setFoo(uint256 value) external {
        varFoo = value;
    }

    function getFoo() external view returns (uint256) {
        return varFoo;
    }

    function setBar(uint256 value) external {
        barStorage.setBar(value);
    }

    function getBar() external view returns (uint256) {
        return barStorage.getBar();
    }
}