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
    bytes32 public immutable currentBlockchainID;
    bytes32 public immutable destinationBlockchainID;
    address public immutable destinationContractAddress;

    // Used for sending an receiving Teleporter messages.
    ITeleporterMessenger public immutable teleporterMessenger;

    constructor(
        address teleporterMessengerAddress,
        bytes32 destinationBlockchainID_,
        address destinationContractAddress_
    ) {
        currentBlockchainID = WarpMessenger(WARP_PRECOMPILE_ADDRESS)
            .getBlockchainID();

        require(teleporterMessengerAddress != address(0), "Invalid Teleporter Messenger Address");
        teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);

        require(destinationBlockchainID_ != bytes32(0), "Invalid Destination Chain ID");
        require(destinationBlockchainID_ != currentBlockchainID, "Cannot Bridge With Same Blockchain");
        destinationBlockchainID = destinationBlockchainID_;

        require(destinationContractAddress_ != address(0), "Invalid Destination Contract Address");
        destinationContractAddress = destinationContractAddress_;
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
        require(msg.sender == address(teleporterMessenger), "Unauthorized teleporter contract");

        // Only allow messages from the destination chain.
        require(senderBlockchainID == destinationBlockchainID, "Invalid Destination Chain");

        // Only allow the partner contract to send messages.
        require(senderAddress == destinationContractAddress, "Unauthorized Sender");

        (address recipient, uint256 amount) = abi.decode(
            message,
            (address, uint256)
        );
        require(recipient != address(0), "Invalid Recipient Address");

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
        require(recipient != address(0), "Invalid Recipient Address");

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
                destinationChainID: destinationBlockchainID,
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
            sender: msg.sender,
            recipient: recipient,
            amount: msg.value,
            feeContractAddress: feeContractAddress,
            feeAmount: adjustedAmount,
            teleporterMessageID: messageID
        });
    }
}
