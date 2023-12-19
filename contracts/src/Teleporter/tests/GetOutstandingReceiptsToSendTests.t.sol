// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {
    TeleporterMessengerTest,
    TeleporterMessage,
    TeleporterMessageInput,
    TeleporterFeeInfo,
    TeleporterMessageReceipt,
    IWarpMessenger
} from "./TeleporterMessengerTest.t.sol";

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
        bytes32 blockchainID =
            bytes32(hex"11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff");

        // Assemble mock receipt information that we will expect to be
        // included in a subsequent message sent to another chain.
        uint256[3] memory receiptMessageNonces = [uint256(13), 42, 94];

        TeleporterMessageReceipt[] memory expectedReceipts = new TeleporterMessageReceipt[](3);
        expectedReceipts[0] = TeleporterMessageReceipt(
            _createMessageID(blockchainID, receiptMessageNonces[0]),
            0xF1DFE63909C027Ed814Dd92C5a3644590abf4850
        );
        expectedReceipts[1] = TeleporterMessageReceipt(
            _createMessageID(blockchainID, receiptMessageNonces[1]),
            0xF1DFE63909C027Ed814Dd92C5a3644590abf4850
        );
        expectedReceipts[2] = TeleporterMessageReceipt(
            _createMessageID(blockchainID, receiptMessageNonces[2]),
            0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542
        );

        // Mock receiving each of the messages corresponding to the receipts.
        for (uint256 i = 0; i < expectedReceipts.length; i++) {
            _receiveTestMessage(
                blockchainID,
                receiptMessageNonces[i],
                expectedReceipts[i].relayerRewardAddress,
                new TeleporterMessageReceipt[](0)
            );
        }

        // Now that we have "received" 3 mock messages, when we send a message back to the
        // other chain, we should expect to see the 3 receipts included in the message metadata.
        TeleporterMessage memory expectedMessage =
            _createMockTeleporterMessage(teleporterMessenger.messageNonce(), hex"deadbeef");
        bytes32 expectedMessageID =
            _createMessageID(DEFAULT_DESTINATION_BLOCKCHAIN_ID, expectedMessage.messageNonce);
        expectedMessage.receipts = expectedReceipts;
        expectedMessage.destinationBlockchainID = blockchainID;
        TeleporterFeeInfo memory feeInfo = TeleporterFeeInfo(address(0), 0);
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: expectedMessage.destinationBlockchainID,
            destinationAddress: expectedMessage.destinationAddress,
            feeInfo: feeInfo,
            requiredGasLimit: expectedMessage.requiredGasLimit,
            allowedRelayerAddresses: expectedMessage.allowedRelayerAddresses,
            message: expectedMessage.message
        });

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(IWarpMessenger.sendWarpMessage.selector),
            abi.encode(expectedMessageID)
        );

        // Expect the exact message to be passed to the precompile.
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.sendWarpMessage, (abi.encode(expectedMessage)))
        );

        // Expect the SendCrossChainMessage event to be emitted.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            messageInput.destinationBlockchainID, expectedMessageID, expectedMessage, feeInfo
        );

        // Submit the message.
        teleporterMessenger.sendCrossChainMessage(messageInput);

        // Submit another message to be sent to check that it does not contain any more receipts.
        TeleporterMessage memory nextExpectedMessage =
            _createMockTeleporterMessage(teleporterMessenger.messageNonce(), hex"deadbeef");
        nextExpectedMessage.destinationBlockchainID = blockchainID;
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.sendWarpMessage, (abi.encode(nextExpectedMessage)))
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            messageInput.destinationBlockchainID,
            _createMessageID(DEFAULT_DESTINATION_BLOCKCHAIN_ID, nextExpectedMessage.messageNonce),
            nextExpectedMessage,
            feeInfo
        );

        // Submit the new message.
        teleporterMessenger.sendCrossChainMessage(messageInput);
    }

    // Test that when there are more than the maximum limit of receipts to include in a message (5),
    // the first batch of receipts on the first outbound message has the maximum batch size, and
    // next contains the remaining receipts.
    function testExceedsLimit() public {
        bytes32 blockchainID =
            bytes32(hex"11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff");

        // Assemble mock receipt information that we will expect to be
        // included in a subsequent message sent to another chain.
        uint256[5] memory receiptMessageNonces1 = [uint256(13), 42, 94, 3, 53];
        TeleporterMessageReceipt[] memory expectedReceiptsBatch1 = new TeleporterMessageReceipt[](5); // the limit of receipts per message is 5.
        expectedReceiptsBatch1[0] = TeleporterMessageReceipt(
            _createMessageID(blockchainID, receiptMessageNonces1[0]),
            0xF1DFE63909C027Ed814Dd92C5a3644590abf4850
        );
        expectedReceiptsBatch1[1] = TeleporterMessageReceipt(
            _createMessageID(blockchainID, receiptMessageNonces1[1]),
            0x52A258ED593C793251a89bfd36caE158EE9fC4F8
        );
        expectedReceiptsBatch1[2] = TeleporterMessageReceipt(
            _createMessageID(blockchainID, receiptMessageNonces1[2]),
            0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542
        );
        expectedReceiptsBatch1[3] = TeleporterMessageReceipt(
            _createMessageID(blockchainID, receiptMessageNonces1[3]),
            0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542
        );
        expectedReceiptsBatch1[4] = TeleporterMessageReceipt(
            _createMessageID(blockchainID, receiptMessageNonces1[4]),
            0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542
        );

        uint256[2] memory receiptMessageNonces2 = [uint256(75), 80];
        TeleporterMessageReceipt[] memory expectedReceiptsBatch2 = new TeleporterMessageReceipt[](2); // the limit of receipts per message is 5.
        expectedReceiptsBatch2[0] = TeleporterMessageReceipt(
            _createMessageID(blockchainID, receiptMessageNonces2[0]),
            0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542
        );
        expectedReceiptsBatch2[1] = TeleporterMessageReceipt(
            _createMessageID(blockchainID, receiptMessageNonces2[1]),
            0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542
        );

        // Mock receiving each of the messages corresponding to the receipts.
        for (uint256 i = 0; i < expectedReceiptsBatch1.length; i++) {
            _receiveTestMessage(
                blockchainID,
                receiptMessageNonces1[i],
                expectedReceiptsBatch1[i].relayerRewardAddress,
                new TeleporterMessageReceipt[](0)
            );
        }
        for (uint256 i = 0; i < expectedReceiptsBatch2.length; i++) {
            _receiveTestMessage(
                blockchainID,
                receiptMessageNonces2[i],
                expectedReceiptsBatch2[i].relayerRewardAddress,
                new TeleporterMessageReceipt[](0)
            );
        }

        // Now that we have "received" 7 mock messages, when we send a message back to the
        // other chain, we should expect to see the 5 receipts included in the message metadata because
        // that is the max receipt batch size limit.
        TeleporterMessage memory expectedMessage =
            _createMockTeleporterMessage(teleporterMessenger.messageNonce(), hex"deadbeef");
        expectedMessage.receipts = expectedReceiptsBatch1;
        expectedMessage.destinationBlockchainID = blockchainID;
        TeleporterFeeInfo memory feeInfo = TeleporterFeeInfo(address(0), 0);
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: expectedMessage.destinationBlockchainID,
            destinationAddress: expectedMessage.destinationAddress,
            feeInfo: feeInfo,
            requiredGasLimit: expectedMessage.requiredGasLimit,
            allowedRelayerAddresses: expectedMessage.allowedRelayerAddresses,
            message: expectedMessage.message
        });

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(IWarpMessenger.sendWarpMessage.selector),
            abi.encode(bytes32(0))
        );

        // Expect the exact message to be passed to the precompile.
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.sendWarpMessage, (abi.encode(expectedMessage)))
        );

        // Expect the SendCrossChainMessage event to be emitted.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            messageInput.destinationBlockchainID,
            _createMessageID(DEFAULT_DESTINATION_BLOCKCHAIN_ID, expectedMessage.messageNonce),
            expectedMessage,
            feeInfo
        );

        // Submit the message.
        teleporterMessenger.sendCrossChainMessage(messageInput);

        // Submit another message to be sent to check that it contains the remaining 2 receipts to be sent.
        TeleporterMessage memory nextExpectedMessage =
            _createMockTeleporterMessage(teleporterMessenger.messageNonce(), hex"deadbeef");
        nextExpectedMessage.receipts = expectedReceiptsBatch2;
        nextExpectedMessage.destinationBlockchainID = blockchainID;
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.sendWarpMessage, (abi.encode(nextExpectedMessage)))
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            messageInput.destinationBlockchainID,
            _createMessageID(DEFAULT_DESTINATION_BLOCKCHAIN_ID, nextExpectedMessage.messageNonce),
            nextExpectedMessage,
            feeInfo
        );

        // Submit the new message.
        teleporterMessenger.sendCrossChainMessage(messageInput);
    }
}
