// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: BSD-3-Clause

pragma solidity 0.8.18;

import "./TeleporterMessengerTest.t.sol";

contract AddFeeAmountTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testSuccess() public {
        // First submit a message with a small fee
        uint256 originalFeeAmount = 10;
        uint256 messageID = _sendTestMessageWithFee(
            DEFAULT_DESTINATION_CHAIN_ID,
            originalFeeAmount
        );

        // Add to the fee
        uint256 additionalFeeAmount = 131313;
        uint256 totalFeeAmount = originalFeeAmount + additionalFeeAmount;
        vm.expectCall(
            address(_mockFeeAsset),
            abi.encodeCall(
                IERC20.transferFrom,
                (
                    address(this),
                    address(teleporterMessenger),
                    additionalFeeAmount
                )
            )
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit AddFeeAmount(
            DEFAULT_DESTINATION_CHAIN_ID,
            messageID,
            TeleporterFeeInfo({
                contractAddress: address(_mockFeeAsset),
                amount: totalFeeAmount
            })
        );
        teleporterMessenger.addFeeAmount(
            DEFAULT_DESTINATION_CHAIN_ID,
            messageID,
            address(_mockFeeAsset),
            additionalFeeAmount
        );

        // Get the fee info to make sure it is properly updated.
        (address actualFeeAsset, uint256 actualFeeAmount) = teleporterMessenger
            .getFeeInfo(DEFAULT_DESTINATION_CHAIN_ID, messageID);
        assertEq(actualFeeAsset, address(_mockFeeAsset));
        assertEq(actualFeeAmount, totalFeeAmount);
    }

    function testInvalidMessage() public {
        // Add to the fee amount of a message that doesn't exist. Expect revert.
        uint256 additionalFeeAmount = 131313;
        uint256 fakeMessageID = 13;
        vm.expectRevert("Message not found or already delivered.");
        teleporterMessenger.addFeeAmount(
            DEFAULT_DESTINATION_CHAIN_ID,
            fakeMessageID,
            address(_mockFeeAsset),
            additionalFeeAmount
        );
    }

    function testMessageAlreadyDelivered() public {
        // First submit a message with a small fee
        uint256 originalFeeAmount = 10;
        uint256 messageID = _sendTestMessageWithFee(
            DEFAULT_DESTINATION_CHAIN_ID,
            originalFeeAmount
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

        // Now try to add to the fee of the message. Should revert since the message receipt was received already.
        uint256 additionalFeeAmount = 131313;
        vm.expectRevert("Message not found or already delivered.");
        teleporterMessenger.addFeeAmount(
            DEFAULT_DESTINATION_CHAIN_ID,
            messageID,
            address(_mockFeeAsset),
            additionalFeeAmount
        );
    }

    function testInvalidAmount() public {
        // First submit a message with a small fee
        uint256 originalFeeAmount = 10;
        uint256 messageID = _sendTestMessageWithFee(
            DEFAULT_DESTINATION_CHAIN_ID,
            originalFeeAmount
        );

        // Expect revert when adding 0 additional amount.
        uint256 additionalFeeAmount = 0;
        vm.expectRevert("Invalid additional fee amount.");
        teleporterMessenger.addFeeAmount(
            DEFAULT_DESTINATION_CHAIN_ID,
            messageID,
            address(_mockFeeAsset),
            additionalFeeAmount
        );
    }

    function testMismatchFeeAsset() public {
        // First submit a message with a small fee
        uint256 originalFeeAmount = 10;
        uint256 messageID = _sendTestMessageWithFee(
            DEFAULT_DESTINATION_CHAIN_ID,
            originalFeeAmount
        );

        // Expect revert when using a different fee asset than originally used.
        uint256 additionalFeeAmount = 131313;
        address differentFeeAsset = 0xA7D7079b0FEaD91F3e65f86E8915Cb59c1a4C664;
        vm.expectRevert("Mismatched fee asset contract address.");
        teleporterMessenger.addFeeAmount(
            DEFAULT_DESTINATION_CHAIN_ID,
            messageID,
            differentFeeAsset,
            additionalFeeAmount
        );
    }

    function testInvalidFeeAsset() public {
        // First submit a message with a small fee
        uint256 originalFeeAmount = 10;
        uint256 messageID = _sendTestMessageWithFee(
            DEFAULT_DESTINATION_CHAIN_ID,
            originalFeeAmount
        );

        // Expect revert when using an invalid fee asset.
        uint256 additionalFeeAmount = 131313;
        address invalidFeeAsset = address(0);
        vm.expectRevert("Invalid fee asset contract address.");
        teleporterMessenger.addFeeAmount(
            DEFAULT_DESTINATION_CHAIN_ID,
            messageID,
            invalidFeeAsset,
            additionalFeeAmount
        );
    }

    function testInsufficientBalance() public {
        // First submit a message with a small fee
        uint256 originalFeeAmount = 10;
        uint256 messageID = _sendTestMessageWithFee(
            DEFAULT_DESTINATION_CHAIN_ID,
            originalFeeAmount
        );

        // Add to the fee, but mock the ERC20 contract returning an error from transferFrom
        uint256 additionalFeeAmount = 131313;
        vm.expectCall(
            address(_mockFeeAsset),
            abi.encodeCall(
                IERC20.transferFrom,
                (
                    address(this),
                    address(teleporterMessenger),
                    additionalFeeAmount
                )
            )
        );
        vm.mockCall(
            address(_mockFeeAsset),
            abi.encodeCall(
                IERC20.transferFrom,
                (
                    address(this),
                    address(teleporterMessenger),
                    additionalFeeAmount
                )
            ),
            abi.encode(false)
        );
        vm.expectRevert("SafeERC20: ERC20 operation did not succeed");

        teleporterMessenger.addFeeAmount(
            DEFAULT_DESTINATION_CHAIN_ID,
            messageID,
            address(_mockFeeAsset),
            additionalFeeAmount
        );
    }
}
