//SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import "./INativeTokenSource.sol";
import "../../Teleporter/ITeleporterMessenger.sol";
import "../../Teleporter/ITeleporterReceiver.sol";
import "../../Teleporter/SafeERC20TransferFrom.sol";

// Precompiled Warp address
address constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;

contract NativeTokenSource is
    ITeleporterReceiver,
    INativeTokenSource,
    ReentrancyGuard
{
    using SafeERC20 for IERC20;

    uint256 public constant MINT_NATIVE_TOKENS_REQUIRED_GAS = 150_000; // TODO this is a placeholder
    bytes32 public immutable currentChainID;
    bytes32 public immutable destinationChainID;
    address public immutable destinationContractAddress;

    // Used for sending an receiving Teleporter messages.
    ITeleporterMessenger public immutable teleporterMessenger;

    constructor(
        address teleporterMessengerAddress,
        bytes32 destinationChainID_,
        address destinationContractAddress_
    ) {
        currentChainID = WarpMessenger(WARP_PRECOMPILE_ADDRESS)
            .getBlockchainID();

        if (teleporterMessengerAddress == address(0)) {
            revert InvalidTeleporterMessengerAddress();
        }
        teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);

        if (destinationContractAddress_ == address(0)) {
            revert InvalidDestinationContractAddress();
        }
        destinationContractAddress = destinationContractAddress_;

        if (destinationChainID_ == currentChainID) {
            revert CannotBridgeTokenWithinSameChain();
        }
        destinationChainID = destinationChainID_;
    }

    /**
     * @dev See {ITeleporterReceiver-receiveTeleporterMessage}.
     *
     * Receives a Teleporter message and routes to the appropriate internal function call.
     */
    function receiveTeleporterMessage(
        bytes32 senderChainID,
        address senderAddress,
        bytes calldata message
    ) external nonReentrant {
        // Only allow the Teleporter messenger to deliver messages.
        if (msg.sender != address(teleporterMessenger)) {
            revert Unauthorized();
        }
        // Only allow messages from the partner chain.
        if (senderChainID != destinationChainID) {
            revert InvalidSourceChain();
        }
        // Only allow the partner contract to send messages.
        if (senderAddress != destinationContractAddress) {
            revert InvalidPartnerContractAddress();
        }

        (address recipient, uint256 amount) = abi.decode(
            message,
            (address, uint256)
        );
        if (recipient == address(0)) {
            revert InvalidRecipient();
        }

        // TODO set up starting threshold.

        // Send to recipient
        payable(recipient).transfer(amount);

        emit UnlockTokens(recipient, amount);
    }

    /**
     * @dev See {IERC20Bridge-bridgeTokens}.
     */
    function transferToDestination(
        address recipient,
        address feeContractAddress,
        uint256 feeAmount
    ) external payable nonReentrant {
        // The recipient cannot be the zero address.
        if (recipient == address(0)) {
            revert InvalidRecipientAddress();
        }

        // Lock tokens in this bridge instance. Supports "fee/burn on transfer" ERC20 token
        // implementations by only bridging the actual balance increase reflected by the call
        // to transferFrom.
        uint256 adjustedAmount = 0;
        if (feeAmount > 0) {
            adjustedAmount = SafeERC20TransferFrom.safeTransferFrom(
                IERC20(feeContractAddress),
                feeAmount
            );
            IERC20(feeContractAddress).safeIncreaseAllowance(
                address(teleporterMessenger),
                adjustedAmount
            );
        }

        // Send Teleporter message.
        bytes memory messageData = abi.encode(recipient, msg.value);

        uint256 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationChainID: destinationChainID,
                destinationAddress: destinationContractAddress,
                feeInfo: TeleporterFeeInfo({
                    contractAddress: feeContractAddress,
                    amount: adjustedAmount
                }),
                requiredGasLimit: MINT_NATIVE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: new address[](0),
                message: messageData
            })
        );

        emit TransferToDestination({
            teleporterMessageID: messageID,
            recipient: recipient,
            transferAmount: msg.value,
            feeContractAddress: feeContractAddress,
            feeAmount: adjustedAmount
        });
    }
}
