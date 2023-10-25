// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "./TeleporterMessengerTest.t.sol";
import "../ITeleporterReceiver.sol";

enum FlakyMessageReceiverAction {
    ReceiveMessage,
    RetryReceive
}

// A mock contract for testing retrying message execution.
// Any message delivered in block with an even timestamp will
// revert. Message delivered in block with an odd timestamp will succeed.
contract FlakyMessageReceiver is ITeleporterReceiver {
    address public immutable teleporterContract;
    string public latestMessage;
    bytes32 public latestMessageSenderSubnetID;
    address public latestMessageSenderAddress;

    constructor(address teleporterContractAddress) {
        teleporterContract = teleporterContractAddress;
    }

    function receiveTeleporterMessage(
        bytes32 originChainID,
        address originSenderAddress,
        bytes calldata messageBytes
    ) external {
        require(msg.sender == teleporterContract, "unauthorized");
        // Decode the payload to recover the action and corresponding function parameters
        (FlakyMessageReceiverAction action, bytes memory actionData) = abi
            .decode(messageBytes, (FlakyMessageReceiverAction, bytes));
        if (action == FlakyMessageReceiverAction.ReceiveMessage) {
            string memory message = abi.decode(actionData, (string));
            _receiveMessage(originChainID, originSenderAddress, message);
        } else if (action == FlakyMessageReceiverAction.RetryReceive) {
            string memory message = abi.decode(actionData, (string));
            _retryReceive(originChainID, originSenderAddress, message);
        } else {
            revert("invalid action");
        }
    }

    // Stores the message in this contract to be fetched by anyone.
    function _receiveMessage(
        bytes32 originChainID,
        address originSenderAddress,
        string memory message
    ) internal {
        require(msg.sender == teleporterContract, "unauthorized");
        require(block.number % 2 != 0, "even block number");
        latestMessage = message;
        latestMessageSenderSubnetID = originChainID;
        latestMessageSenderAddress = originSenderAddress;
    }

    // Tries to call the teleporterContract to receive a message, which should always fail.
    function _retryReceive(
        bytes32 originChainID,
        address originSenderAddress,
        string memory message
    ) internal {
        require(msg.sender == teleporterContract, "unauthorized");
        require(block.number % 2 != 0, "even block number");

        ITeleporterMessenger messenger = ITeleporterMessenger(
            teleporterContract
        );
        messenger.receiveCrossChainMessage(0, address(42));
        latestMessage = message;
        latestMessageSenderSubnetID = originChainID;
        latestMessageSenderAddress = originSenderAddress;
    }
}

contract RetryMessageExecutionTest is TeleporterMessengerTest {
    FlakyMessageReceiver public destinationContract;

    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
        destinationContract = new FlakyMessageReceiver(
            address(teleporterMessenger)
        );
    }

    function testSuccess() public {
        _successfullyRetryMessage();
    }

    function testExecutionFailsAgain() public {
        // First submit a message whose execution fails.
        (
            bytes32 originChainID,
            TeleporterMessage memory message,

        ) = _receiveFailedMessage(false);

        // Now retry it in another block with an even timestamp so that it fails again.
        vm.expectRevert(
            _formatTeleporterErrorMessage("retry execution failed")
        );
        teleporterMessenger.retryMessageExecution(originChainID, message);
    }

    function testMessageHashNotFound() public {
        // Retrying a message that never was delivered should always fail.
        TeleporterMessage memory fakeMessage = TeleporterMessage({
            messageID: 12345,
            senderAddress: address(this),
            destinationAddress: address(destinationContract),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: new bytes(0)
        });

        vm.expectRevert(_formatTeleporterErrorMessage("message not found"));
        teleporterMessenger.retryMessageExecution(
            DEFAULT_ORIGIN_CHAIN_ID,
            fakeMessage
        );
    }

    function testInvalidMessageHash() public {
        // First submit a message whose execution fails.
        (
            bytes32 originChainID,
            TeleporterMessage memory message,

        ) = _receiveFailedMessage(false);

        // Alter the message before retrying it.
        message.message = "altered message";
        vm.expectRevert(_formatTeleporterErrorMessage("invalid message hash"));
        teleporterMessenger.retryMessageExecution(originChainID, message);
    }

    function testCanNotRetryAgainAfterSuccess() public {
        // Successfully retry a message execution.
        (
            bytes32 originChainID,
            TeleporterMessage memory message
        ) = _successfullyRetryMessage();

        // Now try again and make sure it's been cleared from state
        vm.expectRevert(_formatTeleporterErrorMessage("message not found"));
        teleporterMessenger.retryMessageExecution(originChainID, message);
    }

    function testCanNotReceiveMessageWhileRetrying() public {
        // First submit a message whose execution fails.
        (
            bytes32 originChainID,
            TeleporterMessage memory message,

        ) = _receiveFailedMessage(false);

        // Now retry it within a block with an odd timestamp.
        // It should still fail because it tries to re-enter the teleporter contract.
        vm.expectRevert();
        teleporterMessenger.retryMessageExecution(originChainID, message);
    }

    function testEOAFailsThenRetrySucceeds() public {
        (
            bytes32 originChainID,
            address destinationAddress,
            TeleporterMessage memory message
        ) = _receiveMessageSentToEOA();

        // Now mock a contract being deployed to the destination address. It will still return no
        // data from the method being called, but code (with non-zero length) exists at the address.
        vm.etch(destinationAddress, new bytes(10));

        // Retrying the message execution should not revert.
        teleporterMessenger.retryMessageExecution(originChainID, message);
    }

    function testEOAFailsAgainOnRetry() public {
        (
            bytes32 originChainID,
            ,
            TeleporterMessage memory message
        ) = _receiveMessageSentToEOA();

        // Retrying the message execution should revert since there is still no contract deployed
        // to the destination address.
        vm.expectRevert(
            _formatTeleporterErrorMessage("destination address has no code")
        );
        teleporterMessenger.retryMessageExecution(originChainID, message);
    }

    function _receiveFailedMessage(
        bool retryReceive
    ) internal returns (bytes32, TeleporterMessage memory, string memory) {
        // Construct the mock message to be received.
        string memory messageString = "Testing successful message";
        FlakyMessageReceiverAction action;
        if (retryReceive) {
            action = FlakyMessageReceiverAction.RetryReceive;
        } else {
            action = FlakyMessageReceiverAction.ReceiveMessage;
        }
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageID: 42,
            senderAddress: address(this),
            destinationAddress: address(destinationContract),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: abi.encode(action, abi.encode(messageString))
        });
        WarpMessage memory warpMessage = _createDefaultWarpMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            abi.encode(messageToReceive)
        );

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the message in a block with an even height.
        // The message should be successfully received, but its execution should fail.
        vm.roll(12);
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit MessageExecutionFailed(
            DEFAULT_ORIGIN_CHAIN_ID,
            messageToReceive.messageID,
            messageToReceive
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit ReceiveCrossChainMessage(
            warpMessage.sourceChainID,
            messageToReceive.messageID,
            address(this),
            DEFAULT_RELAYER_REWARD_ADDRESS,
            messageToReceive
        );
        teleporterMessenger.receiveCrossChainMessage(
            0,
            DEFAULT_RELAYER_REWARD_ADDRESS
        );

        // Check that the message execution didn't have any effect, but
        // the message was marked as marked as delivered.
        assertEq(destinationContract.latestMessage(), "");
        assertEq(destinationContract.latestMessageSenderSubnetID(), bytes32(0));
        assertEq(destinationContract.latestMessageSenderAddress(), address(0));
        assertEq(
            teleporterMessenger.getRelayerRewardAddress(
                DEFAULT_ORIGIN_CHAIN_ID,
                messageToReceive.messageID
            ),
            DEFAULT_RELAYER_REWARD_ADDRESS
        );
        assertTrue(
            teleporterMessenger.messageReceived(
                DEFAULT_ORIGIN_CHAIN_ID,
                messageToReceive.messageID
            )
        );

        return (DEFAULT_ORIGIN_CHAIN_ID, messageToReceive, messageString);
    }

    function _successfullyRetryMessage()
        internal
        returns (bytes32, TeleporterMessage memory)
    {
        // First submit a message whose execution fails.
        (
            bytes32 originChainID,
            TeleporterMessage memory message,
            string memory messageString
        ) = _receiveFailedMessage(false);

        // Now retry the message execution in a block with an odd height, which should succeed.
        vm.roll(13);
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit MessageExecuted(originChainID, message.messageID, true);
        teleporterMessenger.retryMessageExecution(originChainID, message);

        // Check that the message had the proper affect on the destination contract.
        assertEq(destinationContract.latestMessage(), messageString);
        assertEq(
            destinationContract.latestMessageSenderSubnetID(),
            originChainID
        );
        assertEq(
            destinationContract.latestMessageSenderAddress(),
            address(this)
        );

        return (originChainID, message);
    }
}
