// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.23;

import {Test} from "forge-std/Test.sol";
// solhint-disable no-console
import {ERC20TokenHome} from "../src/TokenHome/ERC20TokenHome.sol";
import {ERC20TokenRemote} from "../src/TokenRemote/ERC20TokenRemote.sol";
import {NativeTokenHome} from "../src/TokenHome/NativeTokenHome.sol";
import {NativeTokenRemote} from "../src/TokenRemote/NativeTokenRemote.sol";

contract StorageSlotTest is Test {
    function testERC20TokenHomeStorageSlot() public {
        bytes32 slot = _erc7201StorageSlot("ERC20TokenHome");
        assertEq(new ERC20TokenHome().ERC20_TOKEN_HOME_STORAGE_LOCATION(), slot);
    }

    function testERC20TokenRemoteStorageSlot() public {
        bytes32 slot = _erc7201StorageSlot("ERC20TokenRemote");
        assertEq(new ERC20TokenRemote().ERC20_TOKEN_REMOTE_STORAGE_LOCATION(), slot);
    }

    function testNativeTokenHomeStorageSlot() public {
        bytes32 slot = _erc7201StorageSlot("NativeTokenHome");
        assertEq(new NativeTokenHome().NATIVE_TOKEN_HOME_STORAGE_LOCATION(), slot);
    }

    function testNativeTokenRemoteStorageSlot() public {
        bytes32 slot = _erc7201StorageSlot("NativeTokenRemote");
        assertEq(new NativeTokenRemote().NATIVE_TOKEN_REMOTE_STORAGE_LOCATION(), slot);
    }

    function testTokenHomeStorageSlot() public {
        bytes32 slot = _erc7201StorageSlot("TokenHome");
        assertEq(new ERC20TokenHome().TOKEN_HOME_STORAGE_LOCATION(), slot);
    }

    function testTokenRemoteStorageSlot() public {
        bytes32 slot = _erc7201StorageSlot("TokenRemote");
        assertEq(new NativeTokenRemote().TOKEN_REMOTE_STORAGE_LOCATION(), slot);
    }

    function _erc7201StorageSlot(bytes memory storageName) private pure returns (bytes32) {
        return keccak256(
            abi.encode(
                uint256(keccak256(abi.encodePacked("avalanche-ictt.storage.", storageName))) - 1
            )
        ) & ~bytes32(uint256(0xff));
    }
}
