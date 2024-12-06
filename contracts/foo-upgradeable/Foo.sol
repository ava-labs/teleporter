pragma solidity 0.8.25;

import {Bar} from "./Bar.sol";


contract Foo is Bar {
    struct FooStorage {
        uint256 varFoo;
    }

    // keccak256(abi.encode(uint256(keccak256("foo.bar")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant FOO_STORAGE_LOCATION =
        0x92c6022ab2e4d8395662f3b814d5b9a489d615215f1269fc366f81bf3a17f700;

    function _getFooStorage()
        private
        pure
        returns (FooStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := FOO_STORAGE_LOCATION
        }
    }

    function initialize() external initializer {
        __Bar_init();
    }

    function setFoo(uint256 value) external {
        FooStorage storage $ = _getFooStorage();
        $.varFoo = value;
    }

    function getFoo() external view returns (uint256) {
        return _getFooStorage().varFoo;
    }
}