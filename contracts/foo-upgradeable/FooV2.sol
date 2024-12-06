pragma solidity 0.8.25;

import "./BarLibrary.sol";
import {ContextUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ContextUpgradeable.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";

contract FooV2 is Initializable, ContextUpgradeable {
    using BarLibrary for BarLibrary.Storage;

    struct FooStorage {
        uint256 varFoo;
        BarLibrary.Storage barStorage;
    }

    struct BarStorage {
        uint256 varBar;
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

    // keccak256(abi.encode(uint256(keccak256("bar")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant BAR_STORAGE_LOCATION =
        0xab087b855a66dd5f7ac611eac86e092a2f3fc08a745c96afb2bc0553ff79ea00;

    function _getBarStorage()
        private
        pure
        returns (BarStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := BAR_STORAGE_LOCATION
        }
    }

    function initialize() public reinitializer(2) {
        // Migrate the existing Bar state to the new BarLibrary
        BarStorage storage barStorage = _getBarStorage();
        FooStorage storage fooStorage = _getFooStorage();
        fooStorage.barStorage.varBar = barStorage.varBar;
    }

    function setFoo(uint256 value) external {
        FooStorage storage $ = _getFooStorage();
        $.varFoo = value;
    }

    function getFoo() external view returns (uint256) {
        return _getFooStorage().varFoo;
    }

    function setBar(uint256 value) external {
        FooStorage storage $ = _getFooStorage();
        $.barStorage.setBar(value);
    }

    function getBar() external view returns (uint256) {
        return _getFooStorage().barStorage.getBar();
    }
}