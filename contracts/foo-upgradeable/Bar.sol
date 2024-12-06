pragma solidity 0.8.25;

import {ContextUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ContextUpgradeable.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";

contract Bar is Initializable, ContextUpgradeable{
    struct BarStorage {
        uint256 varBar;
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

    function __Bar_init() internal onlyInitializing {
        __Context_init();
    }

    function setBar(uint256 value) external {
        BarStorage storage $ = _getBarStorage();
        $.varBar = value;
    }

    function getBar() external view returns (uint256) {
        return _getBarStorage().varBar;
    }
}