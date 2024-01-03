// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {
    TeleporterMessengerTest,
    TeleporterMessage,
    WarpMessage,
    TeleporterMessageReceipt
} from "./TeleporterMessengerTest.t.sol";

enum FallbackReceiveAction {
    Fail,
    Succeed
}

contract FallbackReceiveApp {
    bytes32 public originChainID;
    address public originSenderAddress;
    uint256 public nonce;

    constructor(bytes32 originChainID_, address originSenderAddress_) {
        originChainID = originChainID_;
        originSenderAddress = originSenderAddress_;
    }

    // solhint-disable-next-line no-complex-fallback, payable-fallback
    fallback(bytes calldata data) external returns (bytes memory) {
        bytes4 selector = bytes4(data[:4]);
        // Check that the function selector was for receiveTeleporterMessage.
        require(
            selector == bytes4(keccak256("receiveTeleporterMessage(bytes32,address,bytes)")),
            "FallbackReceiveApp: Invalid selector"
        );

        (bytes32 originChainID_, address originSenderAddress_, bytes memory message) =
            abi.decode(data[4:], (bytes32, address, bytes));

        require(originChainID == originChainID_, "FallbackReceiveApp: Invalid origin chain ID");
        require(
            originSenderAddress == originSenderAddress_,
            "FallbackReceiveApp: Invalid origin sender address"
        );

        FallbackReceiveAction action = abi.decode(message, (FallbackReceiveAction));

        if (action == FallbackReceiveAction.Fail) {
            require(false, "FallbackReceiveApp: Fallback failed");
        } else if (action == FallbackReceiveAction.Succeed) {
            nonce++;
        } else {
            revert("FallbackReceiveApp: Invalid action");
        }

        return abi.encode(nonce);
    }
}

contract FallbackReceiveTest is TeleporterMessengerTest {
    FallbackReceiveApp public destinationContract;

    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
        destinationContract = new FallbackReceiveApp(
            DEFAULT_ORIGIN_CHAIN_ID,
            address(this)
        );
    }

    function testFallbackSuccess() public {
        uint256 nonce = destinationContract.nonce();

        // Construct the mock message to be received.
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageID: 1,
            senderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_CHAIN_ID,
            destinationAddress: address(destinationContract),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: abi.encode(FallbackReceiveAction.Succeed)
        });
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_ORIGIN_CHAIN_ID, abi.encode(messageToReceive));

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the message and check that message execution was successful.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit ReceiveCrossChainMessage(
            warpMessage.sourceChainID,
            messageToReceive.messageID,
            address(this),
            DEFAULT_RELAYER_REWARD_ADDRESS,
            messageToReceive
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit MessageExecuted(DEFAULT_ORIGIN_CHAIN_ID, messageToReceive.messageID);
        teleporterMessenger.receiveCrossChainMessage(0, DEFAULT_RELAYER_REWARD_ADDRESS);

        // Check that the nonce was incremented.
        assertEq(destinationContract.nonce(), nonce + 1);
    }

    function testFallbackFail() public {
        uint256 nonce = destinationContract.nonce();

        // Construct the mock message to be received.
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageID: 1,
            senderAddress: address(this),
            destinationBlockchainID: DEFAULT_DESTINATION_CHAIN_ID,
            destinationAddress: address(destinationContract),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: abi.encode(FallbackReceiveAction.Fail)
        });
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_ORIGIN_CHAIN_ID, abi.encode(messageToReceive));

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the message and check that message execution failed.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit ReceiveCrossChainMessage(
            warpMessage.sourceChainID,
            messageToReceive.messageID,
            address(this),
            DEFAULT_RELAYER_REWARD_ADDRESS,
            messageToReceive
        );
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit MessageExecutionFailed(
            DEFAULT_ORIGIN_CHAIN_ID, messageToReceive.messageID, messageToReceive
        );
        teleporterMessenger.receiveCrossChainMessage(0, DEFAULT_RELAYER_REWARD_ADDRESS);

        // Check that the nonce was not changed.
        assertEq(destinationContract.nonce(), nonce);
    }
}
