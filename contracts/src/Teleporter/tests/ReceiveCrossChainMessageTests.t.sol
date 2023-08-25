// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: BSD-3-Clause

pragma solidity 0.8.18;

import "./TeleporterMessengerTest.t.sol";

// Tests of the logic in receiveCrossChainMessage.
// Tests of the execution and receipts helper methods are split out into independent
// test files in HandleInitialMessageExecutionTests and MarkReceiptAndPayRelayerTests.
contract ReceiveCrossChainMessagedTest is TeleporterMessengerTest {
    bytes public constant DEFAULT_MESSAGE_PAYLOAD =
        hex"cafebabe11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff11223344556677889900aabbccddeeffdeadbeef";

    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testSuccess() public {
        // This test contract must be an allowed relayer since it is what
        // will call receiveCrossSubnetMessage.
        address[] memory allowedRelayers = new address[](2);
        allowedRelayers[0] = address(this);
        allowedRelayers[1] = DEFAULT_RELAYER_REWARD_ADDRESS;

        // Construct the test message to be received.
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageID: 42,
            senderAddress: address(this),
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: allowedRelayers,
            receipts: new TeleporterMessageReceipt[](0),
            message: DEFAULT_MESSAGE_PAYLOAD
        });
        WarpMessage memory warpMessage = _createDefaultWarpMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            abi.encode(messageToReceive)
        );

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(warpMessage);

        // Receive the message.
        teleporterMessenger.receiveCrossChainMessage(
            DEFAULT_RELAYER_REWARD_ADDRESS
        );
    }

    function testNoValidMessage() public {
        // Mock the call to the warp precompile to get the message failing.
        WarpMessage memory emptyMessage = WarpMessage({
            originChainID: bytes32(0),
            originSenderAddress: address(0),
            destinationChainID: bytes32(0),
            destinationAddress: address(0),
            payload: new bytes(0)
        });
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(WarpMessenger.getVerifiedWarpMessage, ()),
            abi.encode(emptyMessage, false)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(WarpMessenger.getVerifiedWarpMessage, ())
        );

        vm.expectRevert("Failed to verify or parse cross chain message.");
        teleporterMessenger.receiveCrossChainMessage(address(1));
    }

    function testInvalidOriginSenderAddress() public {
        // Construct the test message to be received.
        TeleporterMessage
            memory messageToReceive = _createMockTeleporterMessage(
                1,
                new bytes(0)
            );

        WarpMessage memory warpMessage = _createDefaultWarpMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            abi.encode(messageToReceive)
        );
        address invalidSenderAddress = 0xb73aD7e0FF026a805D1f1186EAB89E41bf01835D;
        warpMessage.originSenderAddress = invalidSenderAddress;

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(warpMessage);

        vm.expectRevert("Invalid cross chain message sender address.");
        teleporterMessenger.receiveCrossChainMessage(
            DEFAULT_RELAYER_REWARD_ADDRESS
        );
    }

    function testInvalidDestinationChainID() public {
        // Construct the test message to be received.
        TeleporterMessage
            memory messageToReceive = _createMockTeleporterMessage(
                1,
                new bytes(0)
            );

        WarpMessage memory warpMessage = _createDefaultWarpMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            abi.encode(messageToReceive)
        );
        bytes32 invalidDestinationChainID = bytes32(
            hex"deadbeefcafebabedeadbeefcafebabedeadbeefcafebabedeadbeefcafebabe"
        );
        warpMessage.destinationChainID = invalidDestinationChainID;

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(warpMessage);

        vm.expectRevert("Invalid destination chain ID.");
        teleporterMessenger.receiveCrossChainMessage(
            DEFAULT_RELAYER_REWARD_ADDRESS
        );
    }

    function testInvalidDestinationAddress() public {
        // Construct the test message to be received.
        TeleporterMessage
            memory messageToReceive = _createMockTeleporterMessage(
                1,
                new bytes(0)
            );

        WarpMessage memory warpMessage = _createDefaultWarpMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            abi.encode(messageToReceive)
        );
        address invalidDestinationAddress = 0xb73aD7e0FF026a805D1f1186EAB89E41bf01835D;
        warpMessage.destinationAddress = invalidDestinationAddress;

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(warpMessage);

        vm.expectRevert("Invalid destination address of cross chain message.");
        teleporterMessenger.receiveCrossChainMessage(
            DEFAULT_RELAYER_REWARD_ADDRESS
        );
    }

    function testInvalidRelayerAddress() public {
        vm.expectRevert("Invalid relayer reward address.");
        teleporterMessenger.receiveCrossChainMessage(address(0));
    }

    function testMessageAlreadyReceived() public {
        // Receive a message successfully.
        testSuccess();

        // Check you can't deliver it again.
        vm.expectRevert("Teleporter message previously delivered.");
        teleporterMessenger.receiveCrossChainMessage(
            DEFAULT_RELAYER_REWARD_ADDRESS
        );
    }

    function testUnauthorizedRelayer() public {
        // It doesn't matter if the reward address is an allowed relayer.
        // It is based on who calls receiveCrossChainMessage.
        address[] memory allowedRelayers = new address[](1);
        allowedRelayers[0] = DEFAULT_RELAYER_REWARD_ADDRESS;

        // Construct the test message to be received.
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageID: 42,
            senderAddress: address(this),
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: allowedRelayers,
            receipts: new TeleporterMessageReceipt[](0),
            message: DEFAULT_MESSAGE_PAYLOAD
        });
        WarpMessage memory warpMessage = _createDefaultWarpMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            abi.encode(messageToReceive)
        );

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(warpMessage);

        // Receive the message.
        vm.expectRevert("Unauthorized relayer.");
        teleporterMessenger.receiveCrossChainMessage(
            DEFAULT_RELAYER_REWARD_ADDRESS
        );
    }

    function testMessageSentToEOADoesNotExecute() public {
        _receiveMessageSentToEOA();
    }
}
