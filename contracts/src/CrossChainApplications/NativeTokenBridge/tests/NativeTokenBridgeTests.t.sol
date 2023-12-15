// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";

contract NativeTokenBridgeTest is Test {
    address public constant MOCK_TELEPORTER_MESSENGER_ADDRESS =
        0x644E5b7c5D4Bc8073732CEa72c66e0BB90dFC00f;
    address public constant WARP_PRECOMPILE_ADDRESS =
        address(0x0200000000000000000000000000000000000005);
    address public constant NATIVE_MINTER_PRECOMPILE_ADDRESS =
        address(0x0200000000000000000000000000000000000001);
    bytes32 internal constant _MOCK_BLOCKCHAIN_ID = bytes32(uint256(123456));
    bytes32 internal constant _DEFAULT_OTHER_CHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    address internal constant _DEFAULT_OTHER_BRIDGE_ADDRESS =
        0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    uint256 internal constant _DEFAULT_INITIAL_RESERVE_IMBALANCE = 1000000000;
    address internal constant _DEFAULT_RECIPIENT = 0xa4CEE7d1aF6aDdDD33E3b1cC680AB84fdf1b6d1d;
    uint256 internal constant _DEFAULT_TRANSFER_AMOUNT = 1e18;
    uint256 internal constant _DEFAULT_FEE_AMOUNT = 123456;

    function _createMessageID(uint256 messageNonce) internal pure returns (bytes32) {
        return
            sha256(abi.encode(MOCK_TELEPORTER_MESSENGER_ADDRESS, _MOCK_BLOCKCHAIN_ID, messageNonce));
    }
}
