// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: BSD-3-Clause

pragma solidity 0.8.18;

import "./TeleporterMessengerTest.t.sol";

contract GetNextMessageIDTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testFirstMessageID() public {
        bytes32 chainID = bytes32(
            hex"11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff"
        );

        assertEq(teleporterMessenger.getNextMessageID(chainID), 1);
    }

    function testSecondMessageID() public {
        bytes32 chainID = bytes32(
            hex"11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff"
        );

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(WarpMessenger.sendWarpMessage.selector),
            new bytes(0)
        );
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationChainID: chainID,
            destinationAddress: address(0),
            feeInfo: TeleporterFeeInfo({
                contractAddress: address(0),
                amount: uint256(0)
            }),
            requiredGasLimit: 1e6,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        uint256 first = teleporterMessenger.sendCrossChainMessage(messageInput);
        uint256 second = teleporterMessenger.getNextMessageID(chainID);

        assertEq(first, 1);
        assertEq(second, 2);
    }

    function testOtherDestinationSubnetID() public {
        bytes32 chainID = bytes32(
            hex"11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff"
        );
        bytes32 otherChainID = bytes32(
            hex"00000000556677889900aabbccddeeff11223344556677889900aabbccddeeff"
        );

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(WarpMessenger.sendWarpMessage.selector),
            new bytes(0)
        );
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationChainID: chainID,
            destinationAddress: address(0),
            feeInfo: TeleporterFeeInfo({
                contractAddress: address(0),
                amount: uint256(0)
            }),
            requiredGasLimit: 1e6,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        uint256 first = teleporterMessenger.sendCrossChainMessage(messageInput);
        uint256 other = teleporterMessenger.getNextMessageID(otherChainID);

        assertEq(first, 1);
        assertEq(other, 1);
    }
}
