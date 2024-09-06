// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: UNLICENSED
// slither-disable-next-line solc-version
pragma solidity "0.8.25"; // solhint-disable-line compiler-version

// GENERATED CODE - Do not edit

// Console logging is useful for debug information in _tests_.
// solhint-disable no-console
// solhint-disable func-name-mixedcase

import {Test} from "@forge-std/Test.sol";
import {console} from "@forge-std/console.sol";
import {Unpack} from "../Unpack.sol";

contract UnpackTest is Test {
    function testUnpack_2_4_32_32_8_Dyn_8_Destructive(
        bytes2 v0,
        bytes4 v1,
        bytes32 v2,
        bytes32 v3,
        bytes8 v4,
        bytes memory v5,
        bytes8 v6
    ) external pure {
        bytes memory packed = abi.encodePacked(v0, v1, v2, v3, v4, v5, v6);
        console.logBytes(packed);

        (
            bytes2 got0,
            bytes4 got1,
            bytes32 got2,
            bytes32 got3,
            bytes8 got4,
            bytes memory got5,
            bytes8 got6
        ) = Unpack.unpack_2_4_32_32_8_Dyn_8_Destructive(packed);
        assertEq(v0, got0, "output[0] bytes2");
        assertEq(v1, got1, "output[1] bytes4");
        assertEq(v2, got2, "output[2] bytes32");
        assertEq(v3, got3, "output[3] bytes32");
        assertEq(v4, got4, "output[4] bytes8");
        assertEq(v5, got5, "output[5] bytes memory");
        assertEq(v6, got6, "output[6] bytes8");
    }

    function testUnpack_2_4_32_1(bytes2 v0, bytes4 v1, bytes32 v2, bytes1 v3) external pure {
        bytes memory packed = abi.encodePacked(v0, v1, v2, v3);
        console.logBytes(packed);

        (bytes2 got0, bytes4 got1, bytes32 got2, bytes1 got3) = Unpack.unpack_2_4_32_1(packed);
        assertEq(v0, got0, "output[0] bytes2");
        assertEq(v1, got1, "output[1] bytes4");
        assertEq(v2, got2, "output[2] bytes32");
        assertEq(v3, got3, "output[3] bytes1");
    }

    function testUnpack_2_4_32_8_8(
        bytes2 v0,
        bytes4 v1,
        bytes32 v2,
        bytes8 v3,
        bytes8 v4
    ) external pure {
        bytes memory packed = abi.encodePacked(v0, v1, v2, v3, v4);
        console.logBytes(packed);

        (bytes2 got0, bytes4 got1, bytes32 got2, bytes8 got3, bytes8 got4) =
            Unpack.unpack_2_4_32_8_8(packed);
        assertEq(v0, got0, "output[0] bytes2");
        assertEq(v1, got1, "output[1] bytes4");
        assertEq(v2, got2, "output[2] bytes32");
        assertEq(v3, got3, "output[3] bytes8");
        assertEq(v4, got4, "output[4] bytes8");
    }

    function testUnpack_2_4_32_8(bytes2 v0, bytes4 v1, bytes32 v2, bytes8 v3) external pure {
        bytes memory packed = abi.encodePacked(v0, v1, v2, v3);
        console.logBytes(packed);

        (bytes2 got0, bytes4 got1, bytes32 got2, bytes8 got3) = Unpack.unpack_2_4_32_8(packed);
        assertEq(v0, got0, "output[0] bytes2");
        assertEq(v1, got1, "output[1] bytes4");
        assertEq(v2, got2, "output[2] bytes32");
        assertEq(v3, got3, "output[3] bytes8");
    }
}
