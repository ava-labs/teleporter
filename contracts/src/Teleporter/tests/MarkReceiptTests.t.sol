// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {
    TeleporterMessengerTest,
    TeleporterMessageReceipt,
    TeleporterMessage,
    WarpMessage
} from "./TeleporterMessengerTest.t.sol";

contract MarkReceiptTest is TeleporterMessengerTest {
    struct FeeRewardInfo {
        uint256 feeAmount;
        address relayerRewardAddress;
    }

    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    // This test first mocks sending 3 messages (each with fees) to a different subnet.
    // It then mocks receiving a message back from that subnet with the receipt information
    // for each of the 3 messages sent, and checks that the relayers can redeem their rewards.
    function testCheckRelayersUponReceipt() public {
        // Submit a few mock messages to be sent.
        FeeRewardInfo[3] memory feeRewardInfos = [
            FeeRewardInfo(1111111111111111, 0x52A258ED593C793251a89bfd36caE158EE9fC4F8),
            FeeRewardInfo(2222222222222222, 0xeF6ed43EB8Ff15E336D64d1468947cA1046824E6),
            FeeRewardInfo(3333333333333333, 0xdc00AB1cF6942cE0891eF1AC5ff686833Fa0C542)
        ];

        for (uint256 i = 0; i < feeRewardInfos.length; i++) {
            assertEq(_sendTestMessageWithFee(DEFAULT_ORIGIN_CHAIN_ID, feeRewardInfos[i].feeAmount), i + 1);
        }

        // Mock receiving a message with the 3 receipts from the mock messages sent above.
        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](
                feeRewardInfos.length
            );
        for (uint256 i = 0; i < receipts.length; i++) {
            receipts[i] = TeleporterMessageReceipt({
                receivedMessageID: i + 1,
                relayerRewardAddress: feeRewardInfos[i].relayerRewardAddress
            });
        }
        TeleporterMessage memory messageToReceive = _createMockTeleporterMessage(1, new bytes(0));
        messageToReceive.receipts = receipts;
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_ORIGIN_CHAIN_ID, abi.encode(messageToReceive));

        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the mock message.
        address expectedRelayerRewardAddress = 0x93753a9eA4C9D6eeed9f64eA92E97ce1f5FBAeDe;
        teleporterMessenger.receiveCrossChainMessage(0, expectedRelayerRewardAddress);

        // Check that the relayers have redeemable balances
        for (uint256 i = 0; i < feeRewardInfos.length; i++) {
            assertEq(
                teleporterMessenger.checkRelayerRewardAmount(
                    feeRewardInfos[i].relayerRewardAddress, address(_mockFeeAsset)
                ),
                feeRewardInfos[i].feeAmount
            );
        }

        // Check that the message received is considered delivered, and that the relayer reward address is stored.
        assertEq(teleporterMessenger.getRelayerRewardAddress(DEFAULT_ORIGIN_CHAIN_ID, 1), expectedRelayerRewardAddress);
        assertTrue(teleporterMessenger.messageReceived(DEFAULT_ORIGIN_CHAIN_ID, 1));

        // Check that the message hashes for the message receipts we received have been cleared.
        for (uint256 i = 0; i < receipts.length; i++) {
            assertEq(teleporterMessenger.getMessageHash(DEFAULT_ORIGIN_CHAIN_ID, i + 1), bytes32(0));
        }
    }

    function testReceiptForNoFeeMessage() public {
        // Submit a a mock message with no fee.
        assertEq(_sendTestMessageWithNoFee(DEFAULT_ORIGIN_CHAIN_ID), 1);

        // Mock receiving a message with the a receipts of the mock message sent above.
        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](1);
        receipts[0] =
            TeleporterMessageReceipt({receivedMessageID: 1, relayerRewardAddress: DEFAULT_RELAYER_REWARD_ADDRESS});
        TeleporterMessage memory messageToReceive = _createMockTeleporterMessage(1, new bytes(0));
        messageToReceive.receipts = receipts;
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_ORIGIN_CHAIN_ID, abi.encode(messageToReceive));

        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the mock message.
        address expectedRelayerRewardAddress = 0x2F20537C2F5c57231866DE9D0CE33d0681a200D4;
        teleporterMessenger.receiveCrossChainMessage(0, expectedRelayerRewardAddress);

        // Check that the message received is considered delivered, and that the relayer reward address is stored.
        assertEq(teleporterMessenger.getRelayerRewardAddress(DEFAULT_ORIGIN_CHAIN_ID, 1), expectedRelayerRewardAddress);
        assertTrue(teleporterMessenger.messageReceived(DEFAULT_ORIGIN_CHAIN_ID, 1));
    }

    function testDuplicateReceiptAllowed() public {
        // Submit a mock message to be sent.
        FeeRewardInfo memory feeRewardInfo = FeeRewardInfo(1111111111111111, 0x52A258ED593C793251a89bfd36caE158EE9fC4F8);
        assertEq(_sendTestMessageWithFee(DEFAULT_ORIGIN_CHAIN_ID, feeRewardInfo.feeAmount), 1);

        // Mock receiving a message with the 2 receipts for the  same mock message above.
        TeleporterMessageReceipt[] memory receipts = new TeleporterMessageReceipt[](2);
        TeleporterMessageReceipt memory receipt =
            TeleporterMessageReceipt({receivedMessageID: 1, relayerRewardAddress: feeRewardInfo.relayerRewardAddress});
        receipts[0] = receipt;
        receipts[1] = receipt;
        TeleporterMessage memory messageToReceive = _createMockTeleporterMessage(1, new bytes(0));
        messageToReceive.receipts = receipts;
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_ORIGIN_CHAIN_ID, abi.encode(messageToReceive));

        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the mock message.
        address expectedRelayerRewardAddress = 0x6DAEF0D63ea44C801b655Fd97fe3804B9bFCC097;
        teleporterMessenger.receiveCrossChainMessage(0, expectedRelayerRewardAddress);

        // Check that the relayer redeemable balances was only added once.
        assertEq(
            teleporterMessenger.checkRelayerRewardAmount(feeRewardInfo.relayerRewardAddress, address(_mockFeeAsset)),
            feeRewardInfo.feeAmount
        );

        // Check that the message received is considered delivered, and that the relayer reward address is stored.
        assertEq(teleporterMessenger.getRelayerRewardAddress(DEFAULT_ORIGIN_CHAIN_ID, 1), expectedRelayerRewardAddress);
        assertTrue(teleporterMessenger.messageReceived(DEFAULT_ORIGIN_CHAIN_ID, 1));

        // Check that the message hashes for the message receipts we received have been cleared.
        assertEq(teleporterMessenger.getMessageHash(DEFAULT_ORIGIN_CHAIN_ID, 1), bytes32(0));
    }
}
