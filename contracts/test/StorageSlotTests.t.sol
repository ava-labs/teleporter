// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.23;

import {Test} from "forge-std/Test.sol";
// solhint-disable no-console
import {ERC20TokenHomeUpgradeable} from "../src/TokenHome/ERC20TokenHomeUpgradeable.sol";
import {ERC20TokenRemoteUpgradeable} from "../src/TokenRemote/ERC20TokenRemoteUpgradeable.sol";
import {NativeTokenHomeUpgradeable} from "../src/TokenHome/NativeTokenHomeUpgradeable.sol";
import {NativeTokenRemoteUpgradeable} from "../src/TokenRemote/NativeTokenRemoteUpgradeable.sol";
import {Initializable} from "../src/utils/Initializable.sol";

contract StorageSlotTest is Test {
    function testERC20TokenHomeStorageSlot() public {
        bytes32 slot = _erc7201StorageSlot("ERC20TokenHome");
        assertEq(
            new ERC20TokenHomeUpgradeable(Initializable.Disallowed)
                .ERC20_TOKEN_HOME_STORAGE_LOCATION(),
            slot
        );
    }

    function testERC20TokenRemoteStorageSlot() public {
        bytes32 slot = _erc7201StorageSlot("ERC20TokenRemote");
        assertEq(
            new ERC20TokenRemoteUpgradeable(Initializable.Disallowed)
                .ERC20_TOKEN_REMOTE_STORAGE_LOCATION(),
            slot
        );
    }

    function testNativeTokenHomeStorageSlot() public {
        bytes32 slot = _erc7201StorageSlot("NativeTokenHome");
        assertEq(
            new NativeTokenHomeUpgradeable(Initializable.Disallowed)
                .NATIVE_TOKEN_HOME_STORAGE_LOCATION(),
            slot
        );
    }

    function testNativeTokenRemoteStorageSlot() public {
        bytes32 slot = _erc7201StorageSlot("NativeTokenRemote");
        assertEq(
            new NativeTokenRemoteUpgradeable(Initializable.Disallowed)
                .NATIVE_TOKEN_REMOTE_STORAGE_LOCATION(),
            slot
        );
    }

    function testTokenHomeStorageSlot() public {
        bytes32 slot = _erc7201StorageSlot("TokenHome");
        assertEq(
            new ERC20TokenHomeUpgradeable(Initializable.Disallowed).TOKEN_HOME_STORAGE_LOCATION(),
            slot
        );
    }

    function testTokenRemoteStorageSlot() public {
        bytes32 slot = _erc7201StorageSlot("TokenRemote");
        assertEq(
            new NativeTokenRemoteUpgradeable(Initializable.Disallowed).TOKEN_REMOTE_STORAGE_LOCATION(
            ),
            slot
        );
    }

    function _erc7201StorageSlot(bytes memory storageName) private pure returns (bytes32) {
        return keccak256(
            abi.encode(
                uint256(keccak256(abi.encodePacked("avalanche-ictt.storage.", storageName))) - 1
            )
        ) & ~bytes32(uint256(0xff));
    }
}
