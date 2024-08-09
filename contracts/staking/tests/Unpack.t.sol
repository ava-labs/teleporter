// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// GENERATED CODE - Do not edit

import {Test} from "@forge-std/Test.sol";
import {console} from "@forge-std/console.sol";
import {Unpack} from "../Unpack.sol";

/* solhint-disable func-name-mixedcase */

contract UnpackTest is Test {
    function testUnpack_4_32_32_8_8_Dyn_Destructive(
        bytes4 v0,
        bytes32 v1,
        bytes32 v2,
        bytes8 v3,
        bytes8 v4,
        bytes memory v5
    ) external pure {
        bytes memory packed = abi.encodePacked(v0, v1, v2, v3, v4, v5);
        console.logBytes(packed);

        (bytes4 got0, bytes32 got1, bytes32 got2, bytes8 got3, bytes8 got4, bytes memory got5) =
            Unpack.unpack_4_32_32_8_8_Dyn_Destructive(packed);
        assertEq(v0, got0, "output[0] bytes4");
        assertEq(v1, got1, "output[1] bytes32");
        assertEq(v2, got2, "output[2] bytes32");
        assertEq(v3, got3, "output[3] bytes8");
        assertEq(v4, got4, "output[4] bytes8");
        assertEq(v5, got5, "output[5] bytes memory");
    }

    function testUnpack_4_32_1(bytes4 v0, bytes32 v1, bytes1 v2) external pure {
        bytes memory packed = abi.encodePacked(v0, v1, v2);
        console.logBytes(packed);

        (bytes4 got0, bytes32 got1, bytes1 got2) = Unpack.unpack_4_32_1(packed);
        assertEq(v0, got0, "output[0] bytes4");
        assertEq(v1, got1, "output[1] bytes32");
        assertEq(v2, got2, "output[2] bytes1");
    }

    function testUnpack_4_32_8_8(bytes4 v0, bytes32 v1, bytes8 v2, bytes8 v3) external pure {
        bytes memory packed = abi.encodePacked(v0, v1, v2, v3);
        console.logBytes(packed);

        (bytes4 got0, bytes32 got1, bytes8 got2, bytes8 got3) = Unpack.unpack_4_32_8_8(packed);
        assertEq(v0, got0, "output[0] bytes4");
        assertEq(v1, got1, "output[1] bytes32");
        assertEq(v2, got2, "output[2] bytes8");
        assertEq(v3, got3, "output[3] bytes8");
    }

    function testUnpack_4_32_8(bytes4 v0, bytes32 v1, bytes8 v2) external pure {
        bytes memory packed = abi.encodePacked(v0, v1, v2);
        console.logBytes(packed);

        (bytes4 got0, bytes32 got1, bytes8 got2) = Unpack.unpack_4_32_8(packed);
        assertEq(v0, got0, "output[0] bytes4");
        assertEq(v1, got1, "output[1] bytes32");
        assertEq(v2, got2, "output[2] bytes8");
    }

    function testUnpack_32_32_8_8_Dyn_Destructive(
        bytes32 v0,
        bytes32 v1,
        bytes8 v2,
        bytes8 v3,
        bytes memory v4
    ) external pure {
        bytes memory packed = abi.encodePacked(v0, v1, v2, v3, v4);
        console.logBytes(packed);

        (bytes32 got0, bytes32 got1, bytes8 got2, bytes8 got3, bytes memory got4) =
            Unpack.unpack_32_32_8_8_Dyn_Destructive(packed);
        assertEq(v0, got0, "output[0] bytes32");
        assertEq(v1, got1, "output[1] bytes32");
        assertEq(v2, got2, "output[2] bytes8");
        assertEq(v3, got3, "output[3] bytes8");
        assertEq(v4, got4, "output[4] bytes memory");
    }
}
