// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITeleporterMessenger, TeleporterMessageInput, TeleporterFeeInfo} from "../../Teleporter/ITeleporterMessenger.sol";
import {SafeERC20TransferFrom, SafeERC20} from "../../Teleporter/SafeERC20TransferFrom.sol";
import {TeleporterOwnerUpgradeable} from "../../Teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/security/ReentrancyGuard.sol";

/**
 * @dev ExampleCrossChainMessenger is an example contract that demonstrates how to send and receive
 * messages cross chain.
 */
contract ExampleCrossChainMessenger is
    ReentrancyGuard,
    TeleporterOwnerUpgradeable
{
    using SafeERC20 for IERC20;

    // Messages sent to this contract.
    struct Message {
        address sender;
        string message;
    }

    mapping(bytes32 originBlockchainID => Message message) private _messages;

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
        bytes32 indexed originBlockchainID,
        address indexed originSenderAddress,
        string message
    );

    constructor(
        address teleporterRegistryAddress
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress) {}

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
    ) external nonReentrant returns (uint256) {
        ITeleporterMessenger teleporterMessenger = teleporterRegistry
            .getLatestTeleporter();
        // For non-zero fee amounts, transfer the fee into the control of this contract first, and then
        // allow the Teleporter contract to spend it.
        uint256 adjustedFeeAmount = 0;
        if (feeAmount > 0) {
            adjustedFeeAmount = SafeERC20TransferFrom.safeTransferFrom(
                IERC20(feeTokenAddress),
                feeAmount
            );
            IERC20(feeTokenAddress).safeIncreaseAllowance(
                address(teleporterMessenger),
                adjustedFeeAmount
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
        return
            teleporterMessenger.sendCrossChainMessage(
                TeleporterMessageInput({
                    destinationBlockchainID: destinationBlockchainID,
                    destinationAddress: destinationAddress,
                    feeInfo: TeleporterFeeInfo({
                        feeTokenAddress: feeTokenAddress,
                        amount: adjustedFeeAmount
                    }),
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
    function getCurrentMessage(
        bytes32 originBlockchainID
    ) external view returns (address, string memory) {
        Message memory messageInfo = _messages[originBlockchainID];
        return (messageInfo.sender, messageInfo.message);
    }

    /**
     * @dev See {TeleporterUpgradeable-receiveTeleporterMessage}.
     *
     * Receives a message from another chain.
     */
    function _receiveTeleporterMessage(
        bytes32 originBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) internal override {
        // Store the message.
        string memory messageString = abi.decode(message, (string));
        _messages[originBlockchainID] = Message(
            originSenderAddress,
            messageString
        );
        emit ReceiveMessage(
            originBlockchainID,
            originSenderAddress,
            messageString
        );
    }
}
