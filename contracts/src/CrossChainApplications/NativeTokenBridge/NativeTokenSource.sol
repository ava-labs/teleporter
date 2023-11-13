// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ReentrancyGuard} from "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import {IWarpMessenger} from "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import {INativeTokenSource} from "./INativeTokenSource.sol";
import {ITeleporterMessenger, TeleporterFeeInfo, TeleporterMessageInput} from "../../Teleporter/ITeleporterMessenger.sol";
import {ITeleporterReceiver} from "../../Teleporter/ITeleporterReceiver.sol";
import {SafeERC20TransferFrom} from "../../Teleporter/SafeERC20TransferFrom.sol";
import {SafeERC20TransferFrom} from "../../Teleporter/SafeERC20TransferFrom.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract NativeTokenSource is
    ITeleporterReceiver,
    INativeTokenSource,
    ReentrancyGuard
{
    uint256 public constant MINT_NATIVE_TOKENS_REQUIRED_GAS = 150_000; // TODO this is a placeholder
    bytes32 public immutable destinationBlockchainID;
    address public immutable nativeTokenDestinationAddress;

    // Used for sending and receiving Teleporter messages.
    ITeleporterMessenger public immutable teleporterMessenger;

    constructor(
        address teleporterMessengerAddress,
        bytes32 destinationBlockchainID_,
        address nativeTokenDestinationAddress_
    ) {
        require(
            teleporterMessengerAddress != address(0),
            "NativeTokenSource: zero TeleporterMessenger address"
        );
        teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);

        require(
            destinationBlockchainID_ != bytes32(0),
            "NativeTokenSource: zero destination chain ID"
        );
        require(
            destinationBlockchainID_ !=
                IWarpMessenger(0x0200000000000000000000000000000000000005)
                    .getBlockchainID(),
            "NativeTokenSource: cannot bridge with same blockchain"
        );
        destinationBlockchainID = destinationBlockchainID_;

        require(
            nativeTokenDestinationAddress_ != address(0),
            "NativeTokenSource: zero destination contract address"
        );
        nativeTokenDestinationAddress = nativeTokenDestinationAddress_;
    }

    /**
     * @dev See {ITeleporterReceiver-receiveTeleporterMessage}.
     *
     * Receives a Teleporter message and routes to the appropriate internal function call.
     */
    function receiveTeleporterMessage(
        bytes32 senderBlockchainID,
        address senderAddress,
        bytes calldata message
    ) external nonReentrant {
        // Only allow the Teleporter messenger to deliver messages.
        require(
            msg.sender == address(teleporterMessenger),
            "NativeTokenSource: unauthorized teleporterMessenger contract"
        );

        // Only allow messages from the destination chain.
        require(
            senderBlockchainID == destinationBlockchainID,
            "NativeTokenSource: invalid destination chain"
        );

        // Only allow the partner contract to send messages.
        require(
            senderAddress == nativeTokenDestinationAddress,
            "NativeTokenSource: unauthorized sender"
        );

        (address recipient, uint256 amount) = abi.decode(
            message,
            (address, uint256)
        );
        require(
            recipient != address(0),
            "NativeTokenSource: zero recipient address"
        );

        // Send to recipient
        payable(recipient).transfer(amount);

        emit UnlockTokens(recipient, amount);
    }

    /**
     * @dev See {INativeTokenSource-transferToDestination}.
     */
    function transferToDestination(
        address recipient,
        TeleporterFeeInfo calldata feeInfo,
        address[] calldata allowedRelayerAddresses
    ) external payable nonReentrant {
        // The recipient cannot be the zero address.
        require(
            recipient != address(0),
            "NativeTokenSource: zero ercipient address"
        );

        // Lock tokens in this bridge instance. Supports "fee/burn on transfer" ERC20 token
        // implementations by only bridging the actual balance increase reflected by the call
        // to transferFrom.
        uint256 adjustedFeeAmount = 0;
        if (feeInfo.amount > 0) {
            adjustedFeeAmount = SafeERC20TransferFrom.safeTransferFrom(
                IERC20(feeInfo.contractAddress),
                feeInfo.amount
            );
            SafeERC20.safeIncreaseAllowance(IERC20(feeInfo.contractAddress), address(teleporterMessenger), adjustedFeeAmount);
        }

        uint256 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationChainID: destinationBlockchainID,
                destinationAddress: nativeTokenDestinationAddress,
                feeInfo: feeInfo,
                requiredGasLimit: MINT_NATIVE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: allowedRelayerAddresses,
                message: abi.encode(recipient, msg.value)
            })
        );

        emit TransferToDestination({
            sender: msg.sender,
            recipient: recipient,
            amount: msg.value,
            teleporterMessageID: messageID
        });
    }
}
