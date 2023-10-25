// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "./TeleporterMessengerTest.t.sol";

contract RetryReceiptTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testSuccess() public {
        // Mock receiving 3 messages from another chain
        address[3] memory relayerRewardAddresses = [
            0xa4CEE7d1aF6aDdDD33E3b1cC680AB84fdf1b6d1d,
            0x22d4002028f537599bE9f666d1c4Fa138522f9c8,
            0xb85e0651385876C3E3A1fF81ABf410ecf12f52f4
        ];

        TeleporterMessageReceipt[]
            memory expectedReceipts = new TeleporterMessageReceipt[](3);
        for (uint256 i = 0; i < relayerRewardAddresses.length; i++) {
            _receiveTestMessage(
                DEFAULT_DESTINATION_CHAIN_ID,
                i + 1,
                relayerRewardAddresses[i],
                new TeleporterMessageReceipt[](0)
            );
            expectedReceipts[i] = TeleporterMessageReceipt({
                receivedMessageID: i + 1,
                relayerRewardAddress: relayerRewardAddresses[i]
            });
        }

        // Mock sending a message back to that chain, which should include the 3 receipts.
        TeleporterMessage memory expectedMessage = TeleporterMessage({
            messageID: 1,
            senderAddress: address(this),
            destinationChainID: DEFAULT_DESTINATION_CHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: expectedReceipts,
            message: new bytes(0)
        });
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        TeleporterFeeInfo memory feeInfo = TeleporterFeeInfo(address(0), 0);
        emit SendCrossChainMessage(
            DEFAULT_DESTINATION_CHAIN_ID,
            1,
            expectedMessage,
            feeInfo
        );
        uint256 outboundMessageID = _sendTestMessageWithNoFee(
            DEFAULT_DESTINATION_CHAIN_ID
        );
        assertEq(outboundMessageID, 1);
        assertEq(
            teleporterMessenger.getRelayerRewardAddress(
                DEFAULT_DESTINATION_CHAIN_ID,
                1
            ),
            relayerRewardAddresses[0]
        );

        TeleporterMessageReceipt[]
            memory newExpectedReceipts = new TeleporterMessageReceipt[](2);
        newExpectedReceipts[0] = TeleporterMessageReceipt({
            receivedMessageID: 3,
            relayerRewardAddress: relayerRewardAddresses[2]
        });
        newExpectedReceipts[1] = TeleporterMessageReceipt({
            receivedMessageID: 1,
            relayerRewardAddress: relayerRewardAddresses[0]
        });

        TeleporterMessage memory newExpectedMessage = TeleporterMessage({
            messageID: 2,
            senderAddress: address(this),
            destinationChainID: DEFAULT_DESTINATION_CHAIN_ID,
            destinationAddress: address(0),
            requiredGasLimit: uint256(0),
            allowedRelayerAddresses: new address[](0),
            receipts: newExpectedReceipts,
            message: new bytes(0)
        });

        // Retry two of the receipts.
        uint256[] memory receiptIDs = new uint256[](2);
        receiptIDs[0] = 3;
        receiptIDs[1] = 1;

        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            DEFAULT_DESTINATION_CHAIN_ID,
            2,
            newExpectedMessage,
            feeInfo
        );

        outboundMessageID = _retryTestReceiptsWithNoFee(
            DEFAULT_DESTINATION_CHAIN_ID,
            receiptIDs
        );

        assertEq(outboundMessageID, 2);
    }

    function testDuplicateAllowed() public {
        // Mock receiving a message from another chain
        _receiveTestMessage(
            DEFAULT_DESTINATION_CHAIN_ID,
            1,
            DEFAULT_RELAYER_REWARD_ADDRESS,
            new TeleporterMessageReceipt[](0)
        );

        // Mock sending a message to that chain, which should include the 2 copies of the same receipt.
        TeleporterMessageReceipt
            memory expectedReceipt = TeleporterMessageReceipt({
                receivedMessageID: 1,
                relayerRewardAddress: DEFAULT_RELAYER_REWARD_ADDRESS
            });
        TeleporterMessageReceipt[]
            memory expectedReceipts = new TeleporterMessageReceipt[](2);
        expectedReceipts[0] = expectedReceipt;
        expectedReceipts[1] = expectedReceipt;
        TeleporterMessage memory expectedMessage = TeleporterMessage({
            messageID: 1,
            senderAddress: address(this),
            destinationChainID: DEFAULT_DESTINATION_CHAIN_ID,
            destinationAddress: address(0),
            requiredGasLimit: uint256(0),
            allowedRelayerAddresses: new address[](0),
            receipts: expectedReceipts,
            message: new bytes(0)
        });

        uint256[] memory messageIDs = new uint256[](2);
        messageIDs[0] = 1;
        messageIDs[1] = 1;

        // Test retryReceipts when there are duplicate message IDs in the input.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            DEFAULT_DESTINATION_CHAIN_ID,
            1,
            expectedMessage,
            TeleporterFeeInfo(address(0), 0)
        );

        uint256 outboundMessageID = _retryTestReceiptsWithNoFee(
            DEFAULT_DESTINATION_CHAIN_ID,
            messageIDs
        );
        assertEq(outboundMessageID, 1);
    }

    function testMissingMessage() public {
        // Mock receiving a message from another chain
        _receiveTestMessage(
            DEFAULT_DESTINATION_CHAIN_ID,
            1,
            DEFAULT_RELAYER_REWARD_ADDRESS,
            new TeleporterMessageReceipt[](0)
        );

        uint256[] memory missingIDs = new uint256[](1);
        missingIDs[0] = 21;

        // Try to retry a receipt for an unreceived message from that chain - should fail.
        vm.expectRevert(_formatTeleporterErrorMessage("receipt not found"));
        _retryTestReceiptsWithNoFee(DEFAULT_DESTINATION_CHAIN_ID, missingIDs);
    }

    function _retryTestReceiptsWithFee(
        bytes32 chainID,
        uint256[] memory messageIDs,
        address feeAddress,
        uint256 feeAmount
    ) private returns (uint256) {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(IWarpMessenger.sendWarpMessage.selector),
            new bytes(0)
        );
        if (feeAmount > 0) {
            vm.mockCall(
                feeAddress,
                abi.encodeCall(
                    IERC20.transferFrom,
                    (address(this), address(teleporterMessenger), feeAmount)
                ),
                abi.encode(true)
            );
        }

        TeleporterFeeInfo memory feeInfo = TeleporterFeeInfo({
            contractAddress: feeAddress,
            amount: feeAmount
        });

        if (feeAmount > 0) {
            // Expect the ERC20 contract transferFrom method to be called to transfer the fee.
            vm.expectCall(
                feeAddress,
                abi.encodeCall(
                    IERC20.transferFrom,
                    (address(this), address(teleporterMessenger), feeAmount)
                )
            );
        }

        return
            teleporterMessenger.retryReceipts(
                chainID,
                messageIDs,
                feeInfo,
                new address[](0)
            );
    }

    function _retryTestReceiptsWithNoFee(
        bytes32 chainID,
        uint256[] memory messageIDs
    ) private returns (uint256) {
        return _retryTestReceiptsWithFee(chainID, messageIDs, address(0), 0);
    }
}
