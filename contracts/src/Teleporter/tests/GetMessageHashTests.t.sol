// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterMessengerTest, TeleporterMessage, TeleporterMessageReceipt} from "./TeleporterMessengerTest.t.sol";

contract GetMessageHashTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testSuccess() public {
        // Submit a message
        bytes32 messageID = _sendTestMessageWithNoFee(
            DEFAULT_DESTINATION_CHAIN_ID
        );
        TeleporterMessage memory expectedMessage = TeleporterMessage({
            messageID: messageID,
            senderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_CHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: new bytes(0)
        });
        bytes memory expectedMessageBytes = abi.encode(expectedMessage);
        bytes32 expectedMessageHash = keccak256(expectedMessageBytes);

        // Get its stored hash
        bytes32 actualMessageHash = teleporterMessenger.getMessageHash(
            DEFAULT_DESTINATION_CHAIN_ID,
            messageID
        );
        assertEq(actualMessageHash, expectedMessageHash);
    }

    function testMessageDoesNotExist() public {
        assertEq(
            teleporterMessenger.getMessageHash(
                DEFAULT_DESTINATION_CHAIN_ID,
                bytes32(uint256(42))
            ),
            bytes32(0)
        );
    }

    function testMessageAlreadyReceived() public {
        // Submit a message
        bytes32 messageID = _sendTestMessageWithNoFee(
            DEFAULT_DESTINATION_CHAIN_ID
        );

        // Now mock receiving a message back from that subnet with a receipt of the above message.
        address relayerRewardAddress = 0xA66884fAdC0D4d7B7eedcF61Eb863Ff413bB6234;
        TeleporterMessageReceipt[]
            memory receipts = new TeleporterMessageReceipt[](1);
        receipts[0] = TeleporterMessageReceipt({
            receivedMessageID: messageID,
            relayerRewardAddress: relayerRewardAddress
        });
        _receiveTestMessage(
            DEFAULT_DESTINATION_CHAIN_ID,
            messageID,
            relayerRewardAddress,
            receipts
        );

        // Now the message hash should be cleared.
        assertEq(
            teleporterMessenger.getMessageHash(
                DEFAULT_DESTINATION_CHAIN_ID,
                messageID
            ),
            bytes32(0)
        );
    }
}
