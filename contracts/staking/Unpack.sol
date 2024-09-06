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
     * @dev The originally packed inputs MUST be of the same types as those returned by this function, in the precise
     * order, otherwise the output is undefined.
     * @dev This function takes ownership of and corrupts the `input` parameter to avoid memory expansion for the
     * returned `bytes memory`. The returned array is owned by the caller.
     * @param input Buffer returned by `abi.encodePacked()`
     */
    // slither-disable-next-line naming-convention
    function unpack_4_32_32_8_8_Dyn_Destructive(bytes memory input)
        internal
        pure
        returns (bytes4 v0, bytes32 v1, bytes32 v2, bytes8 v3, bytes8 v4, bytes memory v5)
    {
        if (input.length < 84) {
            revert InsufficientInputLength(84, input.length);
        }

        assembly ("memory-safe") {
            v0 := mload(add(input, 0x20))
            v1 := mload(add(input, 0x24))
            v2 := mload(add(input, 0x44))
            v3 := mload(add(input, 0x64))
            v4 := mload(add(input, 0x6c))
            let length := mload(input)
            let dynLength := sub(length, 84)
            v5 := add(input, 84)
            mstore(v5, dynLength)
        }
    }

    /**
     * @notice Inverts `abi.encodePacked()` for the specific return types.
     * @dev The originally packed inputs MUST be of the same types as those returned by this function, in the precise
     * order, otherwise the output is undefined.
     * @param input Buffer returned by `abi.encodePacked()`
     */
    // slither-disable-next-line naming-convention
    function unpack_4_32_1(bytes memory input)
        internal
        pure
        returns (bytes4 v0, bytes32 v1, bytes1 v2)
    {
        if (input.length != 37) {
            revert IncorrectInputLength(input.length, 37);
        }

        assembly ("memory-safe") {
            v0 := mload(add(input, 0x20))
            v1 := mload(add(input, 0x24))
            v2 := mload(add(input, 0x44))
        }
    }

    /**
     * @notice Inverts `abi.encodePacked()` for the specific return types.
     * @dev The originally packed inputs MUST be of the same types as those returned by this function, in the precise
     * order, otherwise the output is undefined.
     * @param input Buffer returned by `abi.encodePacked()`
     */
    // slither-disable-next-line naming-convention
    function unpack_4_32_8_8(bytes memory input)
        internal
        pure
        returns (bytes4 v0, bytes32 v1, bytes8 v2, bytes8 v3)
    {
        if (input.length != 52) {
            revert IncorrectInputLength(input.length, 52);
        }

        assembly ("memory-safe") {
            v0 := mload(add(input, 0x20))
            v1 := mload(add(input, 0x24))
            v2 := mload(add(input, 0x44))
            v3 := mload(add(input, 0x4c))
        }
    }

    /**
     * @notice Inverts `abi.encodePacked()` for the specific return types.
     * @dev The originally packed inputs MUST be of the same types as those returned by this function, in the precise
     * order, otherwise the output is undefined.
     * @param input Buffer returned by `abi.encodePacked()`
     */
    // slither-disable-next-line naming-convention
    function unpack_4_32_8(bytes memory input)
        internal
        pure
        returns (bytes4 v0, bytes32 v1, bytes8 v2)
    {
        if (input.length != 44) {
            revert IncorrectInputLength(input.length, 44);
        }

        assembly ("memory-safe") {
            v0 := mload(add(input, 0x20))
            v1 := mload(add(input, 0x24))
            v2 := mload(add(input, 0x44))
        }
    }

    /**
     * @notice Inverts `abi.encodePacked()` for the specific return types.
     * @dev The originally packed inputs MUST be of the same types as those returned by this function, in the precise
     * order, otherwise the output is undefined.
     * @dev This function takes ownership of and corrupts the `input` parameter to avoid memory expansion for the
     * returned `bytes memory`. The returned array is owned by the caller.
     * @param input Buffer returned by `abi.encodePacked()`
     */
    // slither-disable-next-line naming-convention
    function unpack_32_32_8_8_Dyn_Destructive(bytes memory input)
        internal
        pure
        returns (bytes32 v0, bytes32 v1, bytes8 v2, bytes8 v3, bytes memory v4)
    {
        if (input.length < 80) {
            revert InsufficientInputLength(80, input.length);
        }

        assembly ("memory-safe") {
            v0 := mload(add(input, 0x20))
            v1 := mload(add(input, 0x40))
            v2 := mload(add(input, 0x60))
            v3 := mload(add(input, 0x68))
            let length := mload(input)
            let dynLength := sub(length, 80)
            v4 := add(input, 80)
            mstore(v4, dynLength)
        }
    }
}
