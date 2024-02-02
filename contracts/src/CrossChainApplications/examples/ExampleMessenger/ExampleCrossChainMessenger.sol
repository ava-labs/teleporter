// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";
import {SafeERC20TransferFrom, SafeERC20} from "@teleporter/SafeERC20TransferFrom.sol";
import {TeleporterOwnerUpgradeable} from "@teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts@4.8.1/security/ReentrancyGuard.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev ExampleCrossChainMessenger is an example contract that demonstrates how to send and receive
 * messages cross chain.
 */
contract ExampleCrossChainMessenger is ReentrancyGuard, TeleporterOwnerUpgradeable {
    using SafeERC20 for IERC20;

    // Messages sent to this contract.
    struct Message {
        address sender;
        string message;
    }

    mapping(bytes32 sourceBlockchainID => Message message) private _messages;

    /**
     * @dev Emitted when a message is submited to be sent.
     */
    event SendMessage(
        bytes32 indexed destinationBlockchainID,
        address indexed destinationAddress,
        address feeTokenAddress,
        uint256 feeAmount,
        uint256 requiredGasLimit,
        string message
    );

    /**
     * @dev Emitted when a new message is received from a given chain ID.
     */
    event ReceiveMessage(
        bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, string message
    );

    constructor(address teleporterRegistryAddress)
        TeleporterOwnerUpgradeable(teleporterRegistryAddress)
    {}

    /**
     * @dev Sends a message to another chain.
     * @return The message ID of the newly sent message.
     */
    function sendMessage(
        bytes32 destinationBlockchainID,
        address destinationAddress,
        address feeTokenAddress,
        uint256 feeAmount,
        uint256 requiredGasLimit,
        string calldata message
    ) external nonReentrant returns (bytes32) {
        ITeleporterMessenger teleporterMessenger = _getTeleporterMessenger();
        // For non-zero fee amounts, first transfer the fee to this contract, and then
        // allow the Teleporter contract to spend it.
        uint256 adjustedFeeAmount;
        if (feeAmount > 0) {
            adjustedFeeAmount =
                SafeERC20TransferFrom.safeTransferFrom(IERC20(feeTokenAddress), feeAmount);
            IERC20(feeTokenAddress).safeIncreaseAllowance(
                address(teleporterMessenger), adjustedFeeAmount
            );
        }

        emit SendMessage({
            destinationBlockchainID: destinationBlockchainID,
            destinationAddress: destinationAddress,
            feeTokenAddress: feeTokenAddress,
            feeAmount: adjustedFeeAmount,
            requiredGasLimit: requiredGasLimit,
            message: message
        });
        return teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationBlockchainID: destinationBlockchainID,
                destinationAddress: destinationAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: feeTokenAddress, amount: adjustedFeeAmount}),
                requiredGasLimit: requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );
    }

    /**
     * @dev Returns the current message from another chain.
     * @return The sender of the message, and the message itself.
     */
    function getCurrentMessage(bytes32 sourceBlockchainID)
        external
        view
        returns (address, string memory)
    {
        Message memory messageInfo = _messages[sourceBlockchainID];
        return (messageInfo.sender, messageInfo.message);
    }

    /**
     * @dev See {TeleporterUpgradeable-receiveTeleporterMessage}.
     *
     * Receives a message from another chain.
     */
    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) internal override {
        // Store the message.
        string memory messageString = abi.decode(message, (string));
        _messages[sourceBlockchainID] = Message(originSenderAddress, messageString);
        emit ReceiveMessage(sourceBlockchainID, originSenderAddress, messageString);
    }
}
