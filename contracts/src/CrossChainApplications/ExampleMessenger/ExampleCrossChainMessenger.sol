// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "../../Teleporter/ITeleporterMessenger.sol";
import "../../Teleporter/SafeERC20TransferFrom.sol";
import "../../Teleporter/ITeleporterReceiver.sol";
import "../../Teleporter/TeleporterRegistry.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

/**
 * @dev ExampleCrossChainMessenger is an example contract that demonstrates how to send and receive
 * messages cross chain.
 */
contract ExampleCrossChainMessenger is ITeleporterReceiver, ReentrancyGuard {
    using SafeERC20 for IERC20;

    // Messages sent to this contract.
    struct Message {
        address sender;
        string message;
    }

    TeleporterRegistry public immutable teleporterRegistry;
    uint256 internal _minTeleporterVersion;

    mapping(bytes32 => Message) private _messages;

    /**
     * @dev Emitted when a message is submited to be sent.
     */
    event SendMessage(
        bytes32 indexed destinationChainID,
        address indexed destinationAddress,
        address feeAsset,
        uint256 feeAmount,
        uint256 requiredGasLimit,
        string message
    );

    /**
     * @dev Emitted when a new message is received from a given chain ID.
     */
    event ReceiveMessage(
        bytes32 indexed originChainID,
        address indexed originSenderAddress,
        string message
    );

    // Errors
    error InvalidTeleporterRegistryAddress();
    error Unauthorized();

    constructor(address teleporterRegistryAddress) {
        if (teleporterRegistryAddress == address(0)) {
            revert InvalidTeleporterRegistryAddress();
        }

        teleporterRegistry = TeleporterRegistry(teleporterRegistryAddress);
        _minTeleporterVersion = teleporterRegistry.getLatestVersion();
    }

    /**
     * @dev See {ITeleporterReceiver-receiveTeleporterMessage}.
     *
     * Receives a message from another chain.
     */
    function receiveTeleporterMessage(
        bytes32 originChainID,
        address originSenderAddress,
        bytes calldata message
    ) external {
        // Only the Teleporter receiver can deliver a message.
        if (
            teleporterRegistry.getAddressToVersion(msg.sender) <
            _minTeleporterVersion
        ) {
            revert Unauthorized();
        }

        // Store the message.
        string memory messageString = abi.decode(message, (string));
        _messages[originChainID] = Message(originSenderAddress, messageString);
        emit ReceiveMessage(originChainID, originSenderAddress, messageString);
    }

    /**
     * @dev Sends a message to another chain.
     */
    function sendMessage(
        bytes32 destinationChainID,
        address destinationAddress,
        address feeContractAddress,
        uint256 feeAmount,
        uint256 requiredGasLimit,
        string calldata message
    ) external nonReentrant returns (uint256 messageID) {
        ITeleporterMessenger teleporterMessenger = teleporterRegistry
            .getLatestTeleporter();
        // For non-zero fee amounts, transfer the fee into the control of this contract first, and then
        // allow the Teleporter contract to spend it.
        uint256 adjustedFeeAmount = 0;
        if (feeAmount > 0) {
            adjustedFeeAmount = SafeERC20TransferFrom.safeTransferFrom(
                IERC20(feeContractAddress),
                feeAmount
            );
            IERC20(feeContractAddress).safeIncreaseAllowance(
                address(teleporterMessenger),
                adjustedFeeAmount
            );
        }

        emit SendMessage({
            destinationChainID: destinationChainID,
            destinationAddress: destinationAddress,
            feeAsset: feeContractAddress,
            feeAmount: adjustedFeeAmount,
            requiredGasLimit: requiredGasLimit,
            message: message
        });
        return
            teleporterMessenger.sendCrossChainMessage(
                TeleporterMessageInput({
                    destinationChainID: destinationChainID,
                    destinationAddress: destinationAddress,
                    feeInfo: TeleporterFeeInfo({
                        contractAddress: feeContractAddress,
                        amount: adjustedFeeAmount
                    }),
                    requiredGasLimit: requiredGasLimit,
                    allowedRelayerAddresses: new address[](0),
                    message: abi.encode(message)
                })
            );
    }

    function updateMinTeleporterVersion() external {
        _minTeleporterVersion = teleporterRegistry.getLatestVersion();
    }

    /**
     * @dev Returns the current message from another chain.
     */
    function getCurrentMessage(
        bytes32 originChainID
    ) external view returns (address sender, string memory message) {
        Message memory messageInfo = _messages[originChainID];
        return (messageInfo.sender, messageInfo.message);
    }
}
