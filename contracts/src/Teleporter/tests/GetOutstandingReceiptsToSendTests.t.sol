// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: BSD-3-Clause

pragma solidity 0.8.18;

import "./TeleporterMessengerTest.t.sol";

contract GetOutstandingReceiptsToSendTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    // getOutstandingReceipts to send is private to the teleporter contract, but we can test
    // its correctness by looking at the receipts included in messages that are submitted to be
    // sent. First, mock receiving 3 messages such that we have 3 outstanding receipts. Then, send
    // a message that should contain the receipts. Finally, send another message that should not
    // have any more receipts since they were included in the previous message.
    function testSuccess() public {
        // Assemble mock receipt information that we will expect to be
        // included in a subsequent message sent to another chain.
        TeleporterMessageReceipt[]
            memory expectedReceipts = new TeleporterMessageReceipt[](3);

        expectedReceipts[0] = TeleporterMessageReceipt(
            13,
            0xF1DFE63909C027Ed814Dd92C5a3644590abf4850
        );
        expectedReceipts[1] = TeleporterMessageReceipt(
            42,
            0x52A258ED593C793251a89bfd36caE158EE9fC4F8
        );
        expectedReceipts[2] = TeleporterMessageReceipt(
            94,
            0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542
        );

        // Mock receiving each of the messages corresponding to the receipts.
        bytes32 chainID = bytes32(
            hex"11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff"
        );
        for (uint256 i = 0; i < expectedReceipts.length; i++) {
            _receiveTestMessage(
                chainID,
                expectedReceipts[i].receivedMessageID,
                expectedReceipts[i].relayerRewardAddress,
                new TeleporterMessageReceipt[](0)
            );
        }

        // Now that we have "received" 3 mock messages, when we send a message back to the
        // other chain, we should expect to see the 3 receipts included in the message metadata.
        TeleporterMessage memory expectedMessage = _createMockTeleporterMessage(
            1,
            hex"deadbeef"
        );
        expectedMessage.receipts = expectedReceipts;
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationChainID: hex"11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff",
            destinationAddress: expectedMessage.destinationAddress,
            feeInfo: TeleporterFeeInfo(address(0), 0),
            requiredGasLimit: expectedMessage.requiredGasLimit,
            allowedRelayerAddresses: expectedMessage.allowedRelayerAddresses,
            message: expectedMessage.message
        });

        // Expect the exact message to be passed to the precompile.
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.sendWarpMessage,
                (
                    messageInput.destinationChainID,
                    address(teleporterMessenger),
                    abi.encode(expectedMessage)
                )
            )
        );

        // Expect the SendCrossChainMessage event to be emitted.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            messageInput.destinationChainID,
            expectedMessage.messageID,
            expectedMessage
        );

        // Submit the message.
        assertEq(teleporterMessenger.sendCrossChainMessage(messageInput), 1);

        // Submit another message to be sent to check that it does not contain any more receipts.
        TeleporterMessage
            memory nextExpectedMessage = _createMockTeleporterMessage(
                2,
                hex"deadbeef"
            );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.sendWarpMessage,
                (
                    messageInput.destinationChainID,
                    address(teleporterMessenger),
                    abi.encode(nextExpectedMessage)
                )
            )
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            messageInput.destinationChainID,
            nextExpectedMessage.messageID,
            nextExpectedMessage
        );

        // Submit the new message.
        assertEq(teleporterMessenger.sendCrossChainMessage(messageInput), 2);
    }

    // Test that when there are more than the maximum limit of receipts to include in a message (5),
    // the first batch of receipts on the first outbound message has the maximum batch size, and
    // next contains the remaining receipts.
    function testExceedsLimit() public {
        // Assemble mock receipt information that we will expect to be
        // included in a subsequent message sent to another chain.
        TeleporterMessageReceipt[]
            memory expectedReceiptsBatch1 = new TeleporterMessageReceipt[](5); // the limit of receipts per message is 5.
        expectedReceiptsBatch1[0] = TeleporterMessageReceipt(
            13,
            0xF1DFE63909C027Ed814Dd92C5a3644590abf4850
        );
        expectedReceiptsBatch1[1] = TeleporterMessageReceipt(
            42,
            0x52A258ED593C793251a89bfd36caE158EE9fC4F8
        );
        expectedReceiptsBatch1[2] = TeleporterMessageReceipt(
            94,
            0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542
        );
        expectedReceiptsBatch1[3] = TeleporterMessageReceipt(
            3,
            0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542
        );
        expectedReceiptsBatch1[4] = TeleporterMessageReceipt(
            53,
            0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542
        );

        TeleporterMessageReceipt[]
            memory expectedReceiptsBatch2 = new TeleporterMessageReceipt[](2); // the limit of receipts per message is 5.
        expectedReceiptsBatch2[0] = TeleporterMessageReceipt(
            75,
            0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542
        );
        expectedReceiptsBatch2[1] = TeleporterMessageReceipt(
            80,
            0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542
        );

        // Mock receiving each of the messages corresponding to the receipts.
        bytes32 chainID = bytes32(
            hex"11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff"
        );
        for (uint256 i = 0; i < expectedReceiptsBatch1.length; i++) {
            _receiveTestMessage(
                chainID,
                expectedReceiptsBatch1[i].receivedMessageID,
                expectedReceiptsBatch1[i].relayerRewardAddress,
                new TeleporterMessageReceipt[](0)
            );
        }
        for (uint256 i = 0; i < expectedReceiptsBatch2.length; i++) {
            _receiveTestMessage(
                chainID,
                expectedReceiptsBatch2[i].receivedMessageID,
                expectedReceiptsBatch2[i].relayerRewardAddress,
                new TeleporterMessageReceipt[](0)
            );
        }

        // Now that we have "received" 7 mock messages, when we send a message back to the
        // other chain, we should expect to see the 5 receipts included in the message metadata because
        // that is the max receipt batch size limit.
        TeleporterMessage memory expectedMessage = _createMockTeleporterMessage(
            1,
            hex"deadbeef"
        );
        expectedMessage.receipts = expectedReceiptsBatch1;
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationChainID: hex"11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff",
            destinationAddress: expectedMessage.destinationAddress,
            feeInfo: TeleporterFeeInfo(address(0), 0),
            requiredGasLimit: expectedMessage.requiredGasLimit,
            allowedRelayerAddresses: expectedMessage.allowedRelayerAddresses,
            message: expectedMessage.message
        });

        // Expect the exact message to be passed to the precompile.
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.sendWarpMessage,
                (
                    messageInput.destinationChainID,
                    address(teleporterMessenger),
                    abi.encode(expectedMessage)
                )
            )
        );

        // Expect the SendCrossChainMessage event to be emitted.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            messageInput.destinationChainID,
            expectedMessage.messageID,
            expectedMessage
        );

        // Submit the message.
        assertEq(teleporterMessenger.sendCrossChainMessage(messageInput), 1);

        // Submit another message to be sent to check that it contains the remaining 2 receipts to be sent.
        TeleporterMessage
            memory nextExpectedMessage = _createMockTeleporterMessage(
                2,
                hex"deadbeef"
            );
        nextExpectedMessage.receipts = expectedReceiptsBatch2;
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.sendWarpMessage,
                (
                    messageInput.destinationChainID,
                    address(teleporterMessenger),
                    abi.encode(nextExpectedMessage)
                )
            )
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            messageInput.destinationChainID,
            nextExpectedMessage.messageID,
            nextExpectedMessage
        );

        // Submit the new message.
        assertEq(teleporterMessenger.sendCrossChainMessage(messageInput), 2);
    }
}
