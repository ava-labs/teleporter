// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {WarpMessage, WarpMessenger} from "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import {TeleporterMessageReceipt, TeleporterMessageInput, TeleporterMessage, TeleporterFeeInfo, ITeleporterMessenger} from "./ITeleporterMessenger.sol";
import {ReceiptQueue} from "./ReceiptQueue.sol";
import {SafeERC20TransferFrom} from "./SafeERC20TransferFrom.sol";
import {ITeleporterReceiver} from "./ITeleporterReceiver.sol";
import {ReentrancyGuards} from "./ReentrancyGuards.sol";

/**
 * @dev Implementation of the {ITeleporterMessenger} interface.
 *
 * This implementation is used to send messages cross chain using the WarpMessenger precompile,
 * and to receive messages sent from other chains. Teleporter contracts should be deployed through Nick's method
 * of universal deployer, such that the same contract is deployed at the same address on all chains.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter/blob/main/SECURITY.md
 */
contract TeleporterMessenger is ITeleporterMessenger, ReentrancyGuards {
    using SafeERC20 for IERC20;
    using ReceiptQueue for ReceiptQueue.TeleporterMessageReceiptQueue;

    // SentMessageInfo includes the fee information for a given message submitted
    // to be sent, along with the hash of the message itself.
    struct SentMessageInfo {
        bytes32 messageHash;
        TeleporterFeeInfo feeInfo;
    }

    WarpMessenger public constant WARP_MESSENGER =
        WarpMessenger(0x0200000000000000000000000000000000000005);

    // The blockchain ID of the chain the contract is deployed on. Initialized lazily on the first call of `receiveCrossChainMessage`
    bytes32 public blockchainID;

    // Tracks the latest message ID used for a given destination chain.
    // Key is the destination blockchain ID, and the value is the last message ID used for that chain.
    // Note that the first message ID used for each chain will be 1 (not 0).
    mapping(bytes32 destinationChainID => uint256 messageID)
        public latestMessageIDs;

    // Tracks the outstanding receipts to send back to a given chain in subsequent messages sent to that chain.
    // Key is the blockchain ID of the other chain, and the value is a queue of pending receipts for messages
    // received from that chain.
    mapping(bytes32 sourceChainID => ReceiptQueue.TeleporterMessageReceiptQueue receiptQueue)
        public receiptQueues;

    // Tracks the message hash and fee information for each message sent that has yet to be acknowledged
    // with a receipt. The messages are tracked per chain and keyed by message ID.
    // The first key is the blockchain ID, the second key is the message ID, and the value is the info
    // for the uniquely identified message.
    mapping(bytes32 destinationChainID => mapping(uint256 messageID => SentMessageInfo messageInfo))
        public sentMessageInfo;

    // Tracks the hash of messages that have been received but whose execution has never succeeded.
    // Enables retrying of failed messages with higher gas limits. Message execution is guaranteed to
    // succeed at most once. The first key is the blockchain ID, the second key is the message ID, and
    // the value is the hash of the uniquely identified message whose execution failed.
    mapping(bytes32 sourceChainID => mapping(uint256 messageID => bytes32 messageHash))
        public receivedFailedMessageHashes;

    // Tracks the relayer reward address for each message delivered from a given chain.
    // Note that these values are also used to determine if a given message has been delivered or not.
    // The first key is the blockchain ID, the second key is the message ID, and the value is the reward address
    // provided by the deliverer of the uniquely identified message.
    mapping(bytes32 sourceChainID => mapping(uint256 messageID => address relayerRewardAddress))
        internal _relayerRewardAddresses;

    // Tracks the fee amounts for a given asset able to be redeemed by a given relayer.
    // The first key is the relayer address, the second key is the fee token contract address,
    // and the value is the amount of the asset redeemable by the relayer.
    mapping(address relayerAddress => mapping(address feeTokenContract => uint256 redeemableRewardAmount))
        internal _relayerRewardAmounts;

    /**
     * @dev See {ITeleporterMessenger-sendCrossChainMessage}
     *
     * When executed, it will kick off an asynchronous event to have the validators of the chain create an
     * aggregate BLS signature of the message.
     *
     * Emits a {SendCrossChainMessage} event when message successfully gets sent.
     */
    function sendCrossChainMessage(
        TeleporterMessageInput calldata messageInput
    ) external senderNonReentrant returns (uint256) {
        // Get the outstanding receipts for messages that have been previously received
        // from the destination chain but not yet acknowledged, and attach the receipts
        // to the Teleporter message to be sent.
        return
            _sendTeleporterMessage(
                messageInput,
                receiptQueues[messageInput.destinationChainID]
                    .getOutstandingReceiptsToSend()
            );
    }

    /**
     * @dev See {ITeleporterMessenger-retrySendCrossChainMessage}
     *
     * Emits {SendCrossChainMessage} event.
     * Requirements:
     *
     * - `message` must have been previously sent to the given destination chain ID.
     * - `message` encoding must match previously sent message.
     */
    function retrySendCrossChainMessage(
        bytes32 destinationChainID,
        TeleporterMessage calldata message
    ) external senderNonReentrant {
        // Get the previously sent message hash.
        SentMessageInfo memory existingMessageInfo = sentMessageInfo[
            destinationChainID
        ][message.messageID];
        // If the message hash is zero, the message was never sent.
        require(
            existingMessageInfo.messageHash != bytes32(0),
            "TeleporterMessenger: message not found"
        );

        // Check that the hash of the provided message matches the one that was originally submitted.
        bytes memory messageBytes = abi.encode(message);
        require(
            keccak256(messageBytes) == existingMessageInfo.messageHash,
            "TeleporterMessenger: invalid message hash"
        );

        // Emit and make state variable changes before external calls when possible,
        // though this function is protected by sender reentrancy guard.
        emit SendCrossChainMessage(
            destinationChainID,
            message.messageID,
            message,
            existingMessageInfo.feeInfo
        );

        // Resubmit the message to the warp message precompile now that it is known that
        // the exact message was already submitted in the past.
        WARP_MESSENGER.sendWarpMessage(
            destinationChainID,
            address(this),
            messageBytes
        );
    }

    /**
     * @dev See {ITeleporterMessenger-addFeeAmount}
     *
     * Emits {AddFeeAmount} event.
     * Requirements:
     *
     * - `additionalFeeAmount` must be non-zero.
     * - message must exist and not have been delivered yet.
     * - `feeContractAddress` must match the fee asset contract address used in the original call to sendCrossChainMessage.
     */
    function addFeeAmount(
        bytes32 destinationChainID,
        uint256 messageID,
        address feeContractAddress,
        uint256 additionalFeeAmount
    ) external senderNonReentrant {
        // The additional fee amount must be non-zero.
        require(
            additionalFeeAmount > 0,
            "TeleporterMessenger: zero additional fee amount"
        );

        // Do not allow adding a fee asset with contract address zero.
        require(
            feeContractAddress != address(0),
            "TeleporterMessenger: zero fee asset contract address"
        );

        // If a receipt has been received for this message, its hash and fee information
        // will be cleared from state. At this point, you can not add to its fee. This is also the
        // case if the given message never existed.
        require(
            sentMessageInfo[destinationChainID][messageID].messageHash !=
                bytes32(0),
            "TeleporterMessenger: message not found"
        );

        // Check that the fee contract address matches the one that was originally used. Only a single
        // fee asset can be used to incentivize the delivery of a given message.
        // We require users to explicitly pass the same fee asset contract address here rather than just using
        // the previously submitted asset type as a defensive measure to avoid having users accidentally confuse
        // which asset they are paying.
        require(
            sentMessageInfo[destinationChainID][messageID]
                .feeInfo
                .contractAddress == feeContractAddress,
            "TeleporterMessenger: invalid fee asset contract address"
        );

        // Transfer the additional fee amount to this Teleporter instance.
        uint256 adjustedAmount = SafeERC20TransferFrom.safeTransferFrom(
            IERC20(feeContractAddress),
            additionalFeeAmount
        );

        // Store the updated fee amount, and emit it as an event.
        sentMessageInfo[destinationChainID][messageID]
            .feeInfo
            .amount += adjustedAmount;

        emit AddFeeAmount(
            destinationChainID,
            messageID,
            sentMessageInfo[destinationChainID][messageID].feeInfo
        );
    }

    /**
     * @dev See {ITeleporterMessenger-receiveCrossChainMessage}
     *
     * Emits {ReceiveCrossChainMessage} event.
     * Re-entrancy is explicitly disallowed between receiving functions. One message is not able to receive another message.
     * Requirements:
     *
     * - `relayerRewardAddress` must not be the zero address.
     * - `messageIndex` must specify a valid warp message in the transaction's storage slots.
     * - Valid warp message provided in storage slots, and sender address matches the address of this contract.
     * - Warp message `destinationChainID` must match the `blockchainID` of this contract.
     * - Warp message `destinationAddress` must match the address of this contract.
     * - Teleporter message was not previously delivered.
     * - Transaction was sent by an allowed relayer for corresponding teleporter message.
     */
    function receiveCrossChainMessage(
        uint32 messageIndex,
        address relayerRewardAddress
    ) external receiverNonReentrant {
        // The relayer reward address is not allowed to be the zero address because it is how the
        // contract tracks whether or not a message has been delivered.
        require(
            relayerRewardAddress != address(0),
            "TeleporterMessenger: zero relayer reward address"
        );

        // Verify and parse the cross chain message included in the transaction access list
        // using the warp message precompile.
        (WarpMessage memory warpMessage, bool success) = WARP_MESSENGER
            .getVerifiedWarpMessage(messageIndex);
        require(success, "TeleporterMessenger: invalid warp message");

        // Only allow for messages to be received from the same address as this teleporter contract.
        // The contract should be deployed using the universal deployer pattern, such that it knows messages
        // received from the same address on other chains were constructed using the same bytecode of this contract.
        // This allows for trusting the message format and uniqueness as specified by sendCrossChainMessage.
        require(
            warpMessage.originSenderAddress == address(this),
            "TeleporterMessenger: invalid origin sender address"
        );

        // If the blockchain ID has yet to be initialized, do so now.
        if (blockchainID == bytes32(0)) {
            blockchainID = WARP_MESSENGER.getBlockchainID();
        }

        // Require that the message was intended for this blockchain and teleporter contract.
        require(
            warpMessage.destinationChainID == blockchainID,
            "TeleporterMessenger: invalid destination chain ID"
        );
        require(
            warpMessage.destinationAddress == address(this),
            "TeleporterMessenger: invalid destination address"
        );

        // Parse the payload of the message.
        TeleporterMessage memory teleporterMessage = abi.decode(
            warpMessage.payload,
            (TeleporterMessage)
        );

        // Check the message has not been delivered before by checking that there is no relayer reward
        // address stored for it already.
        require(
            _relayerRewardAddresses[warpMessage.sourceChainID][
                teleporterMessage.messageID
            ] == address(0),
            "TeleporterMessenger: message already delivered"
        );

        // Check that the caller is allowed to deliver this message.
        require(
            _checkIsAllowedRelayer(
                msg.sender,
                teleporterMessage.allowedRelayerAddresses
            ),
            "TeleporterMessenger: unauthorized relayer"
        );

        // Store the relayer reward address provided, effectively marking the message as received.
        _relayerRewardAddresses[warpMessage.sourceChainID][
            teleporterMessage.messageID
        ] = relayerRewardAddress;

        // Execute the message.
        if (teleporterMessage.message.length > 0) {
            _handleInitialMessageExecution(
                warpMessage.sourceChainID,
                teleporterMessage
            );
        }

        // Process the receipts that were included in the teleporter message by paying the
        // fee for the messages are reward to the given relayers.
        uint256 length = teleporterMessage.receipts.length;
        for (uint256 i = 0; i < length; ++i) {
            TeleporterMessageReceipt memory receipt = teleporterMessage
                .receipts[i];
            _markReceipt(
                warpMessage.sourceChainID,
                receipt.receivedMessageID,
                receipt.relayerRewardAddress
            );
        }

        // Store the receipt of this message delivery.
        ReceiptQueue.TeleporterMessageReceiptQueue
            storage receiptsQueue = receiptQueues[warpMessage.sourceChainID];

        receiptsQueue.enqueue(
            TeleporterMessageReceipt({
                receivedMessageID: teleporterMessage.messageID,
                relayerRewardAddress: relayerRewardAddress
            })
        );

        emit ReceiveCrossChainMessage(
            warpMessage.sourceChainID,
            teleporterMessage.messageID,
            msg.sender,
            relayerRewardAddress,
            teleporterMessage
        );
    }

    /**
     * @dev See {ITeleporterMessenger-retryMessageExecution}
     *
     * Reverts if the message execution fails again on specified message.
     * Emits a {MessageExecuted} event if the retry is successful.
     * Requirements:
     *
     * - `message` must have previously failed to execute, and matches the hash of the failed message.
     */
    function retryMessageExecution(
        bytes32 originChainID,
        TeleporterMessage calldata message
    ) external receiverNonReentrant {
        // Check that the hash of the payload provided matches the hash of the payload that previously failed to execute.
        bytes32 failedMessageHash = receivedFailedMessageHashes[originChainID][
            message.messageID
        ];
        require(
            failedMessageHash != bytes32(0),
            "TeleporterMessenger: message not found"
        );
        require(
            keccak256(abi.encode(message)) == failedMessageHash,
            "TeleporterMessenger: invalid message hash"
        );

        // Check that the target address has fully initialized contract code prior to calling it.
        // If the target address does not have code, the execution automatically fails.
        require(
            message.destinationAddress.code.length > 0,
            "TeleporterMessenger: destination address has no code"
        );

        // Clear the failed message hash from state prior to retrying its execution to redundantly prevent
        // reentrance attacks (on top of the nonReentrant guard).
        emit MessageExecuted(originChainID, message.messageID);
        delete receivedFailedMessageHashes[originChainID][message.messageID];

        // Reattempt the message execution with all of the gas left available for execution of this transaction.
        // Use all of the gas left because this message has already been successfully delivered, and it is the
        // responsibility of the caller to provide as much gas is needed. Compared to the initial delivery, where
        // the relayer should still receive their reward even if the message execution takes more gas than expected.
        // Require that the call be successful such that in the failure case this transaction reverts and the
        // message can be retried again if desired.
        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = message.destinationAddress.call(
            abi.encodeCall(
                ITeleporterReceiver.receiveTeleporterMessage,
                (originChainID, message.senderAddress, message.message)
            )
        );

        require(success, "TeleporterMessenger: retry execution failed");
    }

    /**
     * @dev See {ITeleporterMessenger-sendSpecifiedReceipts}
     *
     * There is no explicit limit to the number of receipts able to be sent by a {sendSpecifiedReceipts} message because
     * this method is intended to be used by relayers themselves to ensure their receipts get returned.
     * There is no fee associated with the empty message, and the same relayer is expected to relay it
     * themselves in order to claim their rewards, so it is their responsibility to ensure that the necessary
     * gas is provided for however many receipts are being retried.
     *
     * These specified receipts are not removed from their corresponding receipt queue because there
     * is no efficient way to remove a specific receipt from an arbitrary position in the queue, and it is
     * harmless for receipts to be sent multiple times within the protocol.
     *
     * Emits {SendCrossChainMessage} event.
     * Requirements:
     * - `messageIDs` must all be valid and have existing receipts.
     */
    function sendSpecifiedReceipts(
        bytes32 originChainID,
        uint256[] calldata messageIDs,
        TeleporterFeeInfo calldata feeInfo,
        address[] calldata allowedRelayerAddresses
    ) external senderNonReentrant returns (uint256) {
        // Iterate through the specified message IDs and create teleporter receipts to send back.
        TeleporterMessageReceipt[]
            memory receiptsToSend = new TeleporterMessageReceipt[](
                messageIDs.length
            );
        for (uint256 i = 0; i < messageIDs.length; i++) {
            uint256 receivedMessageID = messageIDs[i];
            // Check the relayer reward address for this message.
            address relayerRewardAddress = _relayerRewardAddresses[
                originChainID
            ][receivedMessageID];
            require(
                relayerRewardAddress != address(0),
                "TeleporterMessenger: receipt not found"
            );

            receiptsToSend[i] = TeleporterMessageReceipt({
                receivedMessageID: receivedMessageID,
                relayerRewardAddress: relayerRewardAddress
            });
        }

        return
            _sendTeleporterMessage(
                TeleporterMessageInput({
                    destinationChainID: originChainID,
                    destinationAddress: address(0),
                    feeInfo: feeInfo,
                    requiredGasLimit: uint256(0),
                    allowedRelayerAddresses: allowedRelayerAddresses,
                    message: new bytes(0)
                }),
                receiptsToSend
            );
    }

    /**
     * @dev See {ITeleporterMessenger-redeemRelayerRewards}
     *
     * Requirements:
     *
     * - `rewardAmount` must be non-zero.
     */
    function redeemRelayerRewards(address feeAsset) external {
        uint256 rewardAmount = _relayerRewardAmounts[msg.sender][feeAsset];
        require(rewardAmount > 0, "TeleporterMessenger: no reward to redeem");

        // Zero the reward balance before calling the external ERC20 to transfer the
        // reward to prevent any possible re-entrancy.
        delete _relayerRewardAmounts[msg.sender][feeAsset];

        emit RelayerRewardsRedeemed(msg.sender, feeAsset, rewardAmount);

        // "Fee on transfer" tokens do not require a special case here because
        // the amount credited to the caller does not affect this contract's accounting.
        // The reward is considered paid in full in all cases.
        IERC20(feeAsset).safeTransfer(msg.sender, rewardAmount);
    }

    /**
     * See {ITeleporterMessenger-getMessageHash}
     */
    function getMessageHash(
        bytes32 destinationChainID,
        uint256 messageID
    ) external view returns (bytes32) {
        return sentMessageInfo[destinationChainID][messageID].messageHash;
    }

    /**
     * @dev See {ITeleporterMessenger-messageReceived}
     */
    function messageReceived(
        bytes32 originChainID,
        uint256 messageID
    ) external view returns (bool) {
        return _relayerRewardAddresses[originChainID][messageID] != address(0);
    }

    /**
     * @dev See {ITeleporterMessenger-getRelayerRewardAddress}
     */
    function getRelayerRewardAddress(
        bytes32 originChainID,
        uint256 messageID
    ) external view returns (address) {
        return _relayerRewardAddresses[originChainID][messageID];
    }

    /**
     * @dev See {ITeleporterMessenger-checkRelayerRewardAmount}
     */
    function checkRelayerRewardAmount(
        address relayer,
        address feeAsset
    ) external view returns (uint256) {
        return _relayerRewardAmounts[relayer][feeAsset];
    }

    /**
     * @dev See {ITeleporterMessenger-getFeeInfo}
     */
    function getFeeInfo(
        bytes32 destinationChainID,
        uint256 messageID
    ) external view returns (address, uint256) {
        TeleporterFeeInfo memory feeInfo = sentMessageInfo[destinationChainID][
            messageID
        ].feeInfo;
        return (feeInfo.contractAddress, feeInfo.amount);
    }

    /**
     * @dev Returns the next message ID to be used to send a message to the given chain ID.
     */
    function getNextMessageID(bytes32 chainID) external view returns (uint256) {
        return _getNextMessageID(chainID);
    }

    /**
     * @dev See {ITeleporterMessenger-getReceiptQueueSize}
     */
    function getReceiptQueueSize(
        bytes32 chainID
    ) external view returns (uint256) {
        return receiptQueues[chainID].size();
    }

    /**
     * @dev See {ITeleporterMessenger-getReceiptAtIndex}
     */
    function getReceiptAtIndex(
        bytes32 chainID,
        uint256 index
    ) external view returns (TeleporterMessageReceipt memory) {
        return receiptQueues[chainID].getReceiptAtIndex(index);
    }

    /**
     * @dev Checks whether `delivererAddress` is allowed to deliver the message.
     */
    function _checkIsAllowedRelayer(
        address delivererAddress,
        address[] memory allowedRelayers
    ) internal pure returns (bool) {
        // An empty allowed relayers list means anyone is allowed to deliver the message.
        if (allowedRelayers.length == 0) {
            return true;
        }

        // Otherwise, the deliverer address must be included in allowedRelayers.
        for (uint256 i = 0; i < allowedRelayers.length; ++i) {
            if (allowedRelayers[i] == delivererAddress) {
                return true;
            }
        }
        return false;
    }

    /**
     * @dev Helper function for sending a teleporter message cross chain.
     * Constructs the teleporter message and sends it through the warp messenger precompile,
     * and performs fee transfer if necessary.
     *
     * Emits a {SendCrossChainMessage} event.
     */
    function _sendTeleporterMessage(
        TeleporterMessageInput memory messageInput,
        TeleporterMessageReceipt[] memory receipts
    ) private returns (uint256) {
        // Get the message ID to use for this message.
        uint256 messageID = _getNextMessageID(messageInput.destinationChainID);

        // Construct and serialize the message.
        TeleporterMessage memory teleporterMessage = TeleporterMessage({
            messageID: messageID,
            senderAddress: msg.sender,
            destinationAddress: messageInput.destinationAddress,
            requiredGasLimit: messageInput.requiredGasLimit,
            allowedRelayerAddresses: messageInput.allowedRelayerAddresses,
            receipts: receipts,
            message: messageInput.message
        });
        bytes memory teleporterMessageBytes = abi.encode(teleporterMessage);

        // Set the message ID value as being used.
        latestMessageIDs[messageInput.destinationChainID] = messageID;

        // If the fee amount is non-zero, transfer the asset into control of this TeleporterMessenger contract instance.
        // The fee is allowed to be 0 because it's possible for someone to run their own relayer and deliver their own messages,
        // which does not require further incentivization. They still must pay the transaction fee to submit the message, so
        // this is not a DOS vector in terms of being able to submit zero-fee messages.
        uint256 adjustedFeeAmount = 0;
        if (messageInput.feeInfo.amount > 0) {
            // If the fee amount is non-zero, check that the contract address is not address(0)
            require(
                messageInput.feeInfo.contractAddress != address(0),
                "TeleporterMessenger: zero fee asset contract address"
            );

            adjustedFeeAmount = SafeERC20TransferFrom.safeTransferFrom(
                IERC20(messageInput.feeInfo.contractAddress),
                messageInput.feeInfo.amount
            );
        }

        // Store the fee asset and amount to be paid to the relayer of this message upon receiving the receipt.
        // Also store the message hash so that it can be retried until a receipt of its delivery is received back.
        TeleporterFeeInfo memory adjustedFeeInfo = TeleporterFeeInfo({
            contractAddress: messageInput.feeInfo.contractAddress,
            amount: adjustedFeeAmount
        });
        sentMessageInfo[messageInput.destinationChainID][
            messageID
        ] = SentMessageInfo({
            messageHash: keccak256(teleporterMessageBytes),
            feeInfo: adjustedFeeInfo
        });

        emit SendCrossChainMessage(
            messageInput.destinationChainID,
            messageID,
            teleporterMessage,
            adjustedFeeInfo
        );

        // Submit the message to the AWM precompile.
        WARP_MESSENGER.sendWarpMessage(
            messageInput.destinationChainID,
            address(this),
            teleporterMessageBytes
        );

        return messageID;
    }

    /**
     * @dev Marks the receipt of a message from the given destination chain ID with the given message ID.
     *
     * If the receipt was already previously received for this message, returns early.
     * If existing message is found and not yet delivered, deletes it from state and increment the fee/reward
     */
    function _markReceipt(
        bytes32 destinationChainID,
        uint256 messageID,
        address relayerRewardAddress
    ) private {
        // Get the information about the sent message to be marked as received.
        SentMessageInfo memory messageInfo = sentMessageInfo[
            destinationChainID
        ][messageID];

        // If the message hash does not exist, it could be the case that the receipt was already
        // received for this message (it's possible for receipts to be sent more than once)
        // or that the other chain sent an invalid receipt. Return early since this is an expected
        // case where there is no fee to be paid for the given message.
        if (messageInfo.messageHash == bytes32(0)) {
            return;
        }

        // Delete the message information from state now that it is known to be delivered.
        delete sentMessageInfo[destinationChainID][messageID];

        // Increment the fee/reward amount owed to the relayer for having delivered
        // the message identified in this receipt.
        _relayerRewardAmounts[relayerRewardAddress][
            messageInfo.feeInfo.contractAddress
        ] += messageInfo.feeInfo.amount;
    }

    /**
     * @dev Attempts to execute the newly delivered message.
     *
     * Only revert in the event that the message deliverer (relayer) did not provide enough gas to handle the execution
     * (including possibly storing a failed message in state). All execution specific errors (i.e. invalid call data, etc)
     * that are not in the relayers control are caught and handled properly.
     *
     * Emits a {MessageExecuted} event if the call on destination address is successful.
     * Emits a {MessageExecutionFailed} event if the call on destination address fails with formatted call data.
     * Requirements:
     *
     * - There is enough gas left to cover `message.requiredGasLimit`.
     */
    function _handleInitialMessageExecution(
        bytes32 originChainID,
        TeleporterMessage memory message
    ) private {
        // Check that the message delivery was provided the required gas amount as specified by the sender.
        // If the required gas amount is provided, the message will be considered delivered whether or not
        // its execution succeeds, such that the relayer can claim their fee reward. However, if the message
        // execution fails, the message hash will be stored in state such that anyone can try to provide more
        // gas to successfully execute the message.
        require(
            gasleft() >= message.requiredGasLimit,
            "TeleporterMessenger: insufficient gas"
        );

        // The destination address must have fully initialized contract code in order for the message
        // to call it. If the destination address does not have code, store the message as a failed
        // execution so that it can be retried in the future should a contract be later deployed to
        // the address.
        if (message.destinationAddress.code.length == 0) {
            _storeFailedMessageExecution(originChainID, message);
            return;
        }

        // Call the destination address of the message with the formatted call data.
        // Only provide the required gas limit to the sub-call so that the end application
        // cannot consume an arbitrary amount gas.
        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = message.destinationAddress.call{
            gas: message.requiredGasLimit
        }(
            abi.encodeCall(
                ITeleporterReceiver.receiveTeleporterMessage,
                (originChainID, message.senderAddress, message.message)
            )
        );

        // If the execution failed, store a hash of the message in state such that its
        // execution can be retried again in the future with a higher gas limit (paid by whoever
        // retries).Either way, the message will now be considered "delivered" since the relayer
        // provided enough gas to meet the required gas limit.
        if (!success) {
            _storeFailedMessageExecution(originChainID, message);
            return;
        }

        emit MessageExecuted(originChainID, message.messageID);
    }

    /**
     * @dev Stores the hash of a message that has been successfully delivered but fails to execute properly
     * such that the message execution can be retried by anyone in the future.
     */
    function _storeFailedMessageExecution(
        bytes32 originChainID,
        TeleporterMessage memory message
    ) private {
        receivedFailedMessageHashes[originChainID][
            message.messageID
        ] = keccak256(abi.encode(message));

        // Emit a failed execution event for anyone monitoring unsuccessful messages to retry.
        emit MessageExecutionFailed(originChainID, message.messageID, message);
    }

    /**
     * @dev Returns the next message ID to be used to send a message to the given chain ID.
     */
    function _getNextMessageID(bytes32 chainID) private view returns (uint256) {
        return latestMessageIDs[chainID] + 1;
    }
}
