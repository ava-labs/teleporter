// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: UNLICENSED
// slither-disable-next-line solc-version
pragma solidity "0.8.25"; // solhint-disable-line compiler-version

// GENERATED CODE - Do not edit

// solhint-disable no-inline-assembly
// slither-disable-start assembly
// Mixed-case can't apply to numbers and an underscore is the most natural separator
// solhint-disable func-name-mixedcase

library Unpack {
    /// @dev Thrown if the input buffer to an unpack function is of the wrong length.
    error IncorrectInputLength(uint256 expected, uint256 actual);

    /// @dev Thrown if the input buffer to an unpack function with dynamically sized output is too short.
    error InsufficientInputLength(uint256 min, uint256 actual);

    /**
     * @notice Inverts `abi.encodePacked()` for the specific return types.
     * @dev The originally packed inputs MUST be of the same byte lengths as those returned by this function, in the
     * precise order, otherwise the output is undefined. `bytes<N>` and `uint<8N>` are interchangeable.
     * @dev This function takes ownership of and corrupts the `input` parameter to avoid memory expansion for the
     * returned `bytes memory`. The returned array is owned by the caller.
     * @param input Buffer returned by `abi.encodePacked()`
     */
    // slither-disable-next-line naming-convention
    function unpack_2_4_32_32_8_Dyn_8_Destructive(bytes memory input)
        internal
        pure
        returns (
            bytes2 v0,
            bytes4 v1,
            bytes32 v2,
            bytes32 v3,
            bytes8 v4,
            bytes memory v5,
            bytes8 v6
        )
    {
        if (input.length < 86) {
            revert InsufficientInputLength(86, input.length);
        }

        assembly ("memory-safe") {
            v0 := mload(add(input, 0x20))
            v1 := mload(add(input, 0x22))
            v2 := mload(add(input, 0x26))
            v3 := mload(add(input, 0x46))
            v4 := mload(add(input, 0x66))
            let length := mload(input)
            let dynLength := sub(length, 86)
            v5 := add(input, 78)
            mstore(v5, dynLength)
            let end := add(input, add(length, 0x20))
            v6 := mload(sub(end, 0x08))
        }
    }

    /**
     * @notice Inverts `abi.encodePacked()` for the specific return types.
     * @dev The originally packed inputs MUST be of the same byte lengths as those returned by this function, in the
     * precise order, otherwise the output is undefined. `bytes<N>` and `uint<8N>` are interchangeable.
     * @param input Buffer returned by `abi.encodePacked()`
     */
    // slither-disable-next-line naming-convention
    function unpack_2_4_32_1(bytes memory input)
        internal
        pure
        returns (bytes2 v0, bytes4 v1, bytes32 v2, bytes1 v3)
    {
        if (input.length != 39) {
            revert IncorrectInputLength(input.length, 39);
        }

        assembly ("memory-safe") {
            v0 := mload(add(input, 0x20))
            v1 := mload(add(input, 0x22))
            v2 := mload(add(input, 0x26))
            v3 := mload(add(input, 0x46))
        }
    }

    /**
     * @notice Inverts `abi.encodePacked()` for the specific return types.
     * @dev The originally packed inputs MUST be of the same byte lengths as those returned by this function, in the
     * precise order, otherwise the output is undefined. `bytes<N>` and `uint<8N>` are interchangeable.
     * @param input Buffer returned by `abi.encodePacked()`
     */
    // slither-disable-next-line naming-convention
    function unpack_2_4_32_8_8(bytes memory input)
        internal
        pure
        returns (bytes2 v0, bytes4 v1, bytes32 v2, bytes8 v3, bytes8 v4)
    {
        if (input.length != 54) {
            revert IncorrectInputLength(input.length, 54);
        }

        assembly ("memory-safe") {
            v0 := mload(add(input, 0x20))
            v1 := mload(add(input, 0x22))
            v2 := mload(add(input, 0x26))
            v3 := mload(add(input, 0x46))
            v4 := mload(add(input, 0x4e))
        }
    }

    /**
     * @notice Inverts `abi.encodePacked()` for the specific return types.
     * @dev The originally packed inputs MUST be of the same byte lengths as those returned by this function, in the
     * precise order, otherwise the output is undefined. `bytes<N>` and `uint<8N>` are interchangeable.
     * @param input Buffer returned by `abi.encodePacked()`
     */
    // slither-disable-next-line naming-convention
    function unpack_2_4_32_8(bytes memory input)
        internal
        pure
        returns (bytes2 v0, bytes4 v1, bytes32 v2, bytes8 v3)
    {
        if (input.length != 46) {
            revert IncorrectInputLength(input.length, 46);
        }

        assembly ("memory-safe") {
            v0 := mload(add(input, 0x20))
            v1 := mload(add(input, 0x22))
            v2 := mload(add(input, 0x26))
            v3 := mload(add(input, 0x46))
        }
    }
}
