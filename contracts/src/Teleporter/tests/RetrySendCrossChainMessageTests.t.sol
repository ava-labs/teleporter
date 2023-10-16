// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "./TeleporterMessengerTest.t.sol";

contract RetrySendCrossChainMessageTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testSuccess() public {
        // Send a message
        uint256 messageID = _sendTestMessageWithFee(
            DEFAULT_DESTINATION_CHAIN_ID,
            654456
        );
        TeleporterMessage memory expectedMessage = TeleporterMessage({
            messageID: messageID,
            senderAddress: address(this),
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: new bytes(0)
        });

        // Retry it
        teleporterMessenger.retrySendCrossChainMessage(
            DEFAULT_DESTINATION_CHAIN_ID,
            expectedMessage
        );
    }

    function testMessageNotFound() public {
        TeleporterMessage memory fakeMessage = TeleporterMessage({
            messageID: 354,
            senderAddress: address(this),
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: new bytes(0)
        });
        vm.expectRevert(_formatErrorMessage("message hash not found"));
        teleporterMessenger.retrySendCrossChainMessage(
            DEFAULT_DESTINATION_CHAIN_ID,
            fakeMessage
        );
    }

    function testInvalidMessageHash() public {
        // Send a message, then try to alter it's contents.
        uint256 messageID = _sendTestMessageWithFee(
            DEFAULT_DESTINATION_CHAIN_ID,
            654456
        );
        TeleporterMessage memory alteredMessage = TeleporterMessage({
            messageID: messageID,
            senderAddress: address(this),
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: hex"cafebabe" // The message would be expected to be empty
        });

        // Retry it - should fail.
        vm.expectRevert(_formatErrorMessage("invalid message hash"));
        teleporterMessenger.retrySendCrossChainMessage(
            DEFAULT_DESTINATION_CHAIN_ID,
            alteredMessage
        );
    }
}
