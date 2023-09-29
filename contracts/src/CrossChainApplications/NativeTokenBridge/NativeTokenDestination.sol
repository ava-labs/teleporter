//SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import "@subnet-evm-contracts/interfaces/IAllowList.sol";
import "@subnet-evm-contracts/interfaces/INativeMinter.sol";
import "./INativeTokenDestination.sol";
import "../../Teleporter/ITeleporterMessenger.sol";
import "../../Teleporter/ITeleporterReceiver.sol";
import "../../Teleporter/SafeERC20TransferFrom.sol";

// Precompiled Native Minter Contract Address
address constant MINTER_ADDRESS = 0x0200000000000000000000000000000000000001;
// Designated Blackhole Address
address constant BLACKHOLE_ADDRESS = 0x0100000000000000000000000000000000000001;
// Precompiled Warp address
address constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;

contract NativeTokenDestination is
    ITeleporterReceiver,
    INativeTokenDestination,
    ReentrancyGuard
{
    using SafeERC20 for IERC20;

    INativeMinter private immutable _nativeMinter =
        INativeMinter(MINTER_ADDRESS);

    uint256 public constant TRANSFER_NATIVE_TOKENS_REQUIRED_GAS = 150_000; // TODO this is a placeholder
    bytes32 public immutable currentChainID;
    bytes32 public immutable sourceChainID;
    address public immutable sourceContractAddress;

    // Used for sending an receiving Teleporter messages.
    ITeleporterMessenger public immutable teleporterMessenger;

    // TODO we probably want to add the original token supply from this chain to the constructor.
    constructor(
        address teleporterMessengerAddress,
        bytes32 sourceChainID_,
        address sourceContractAddress_
    ) {
        currentChainID = WarpMessenger(WARP_PRECOMPILE_ADDRESS)
            .getBlockchainID();

        if (teleporterMessengerAddress == address(0)) {
            revert InvalidTeleporterMessengerAddress();
        }
        teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);

        if (sourceContractAddress_ == address(0)) {
            revert InvalidSourceContractAddress();
        }
        sourceContractAddress = sourceContractAddress_;

        if (sourceChainID_ == currentChainID) {
            revert CannotBridgeTokenWithinSameChain();
        }
        sourceChainID = sourceChainID_;
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
        if (senderChainID != sourceChainID) {
            revert InvalidSourceChain();
        }
        // Only allow the partner contract to send messages.
        if (senderAddress != sourceContractAddress) {
            revert InvalidPartnerContractAddress();
        }

        (address recipient, uint256 amount) = abi.decode(
            message,
            (address, uint256)
        );

        // Calls NativeMinter precompile through INativeMinter interface.
        _nativeMinter.mintNativeCoin(recipient, amount);
        emit MintNativeTokens(recipient, amount);
    }

    /**
     * @dev See {INativeTokenMinter-bridgeTokens}.
     */
    function transferToSource(
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

        // Burn native token by sending to BLACKHOLE_ADDRESS
        payable(BLACKHOLE_ADDRESS).transfer(msg.value);

        // Send Teleporter message.
        bytes memory messageData = abi.encode(recipient, msg.value);

        uint256 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationChainID: sourceChainID,
                destinationAddress: sourceContractAddress,
                feeInfo: TeleporterFeeInfo({
                    contractAddress: feeContractAddress,
                    amount: adjustedAmount
                }),
                requiredGasLimit: TRANSFER_NATIVE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: new address[](0),
                message: messageData
            })
        );

        emit TransferToSource({
            teleporterMessageID: messageID,
            recipient: recipient,
            transferAmount: msg.value,
            feeContractAddress: feeContractAddress,
            feeAmount: adjustedAmount
        });
    }
}
