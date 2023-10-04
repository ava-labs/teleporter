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
// This is the address where the burned base fee go. We use this to send the amount of
// burned tokens back to the source chain.
address constant BURNED_TX_FEES_ADDRESS = 0x0100000000000000000000000000000000000000;
// Designated Blackhole Address. This is where we burn token before sending an unlock
// message to the source chain. Different from the burned tx fee address so we can
// track each separately.
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
    bytes32 public immutable currentBlockchainID;
    bytes32 public immutable sourceBlockchainID;
    address public immutable sourceContractAddress;

    // Used for sending an receiving Teleporter messages.
    ITeleporterMessenger public immutable teleporterMessenger;

    // TODO we probably want to add the original token supply from this chain to the constructor.
    constructor(
        address teleporterMessengerAddress,
        bytes32 sourceBlockchainID_,
        address sourceContractAddress_
    ) {
        currentBlockchainID = WarpMessenger(WARP_PRECOMPILE_ADDRESS)
            .getBlockchainID();

        require(teleporterMessengerAddress != address(0), "Invalid Teleporter Messenger Address");
        teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);

        require(sourceBlockchainID_ != bytes32(0), "Invalid Source Chain ID");
        require(sourceBlockchainID_ != currentBlockchainID, "Cannot Bridge With Same Blockchain");
        sourceBlockchainID = sourceBlockchainID_;

        require(sourceContractAddress_ != address(0), "Invalid Source Contract Address");
        sourceContractAddress = sourceContractAddress_;
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

        // Only allow messages from the source chain.
        require(senderBlockchainID == sourceBlockchainID, "Invalid Source Chain");

        // Only allow the partner contract to send messages.
        require(senderAddress == sourceContractAddress, "Unauthorized Sender");

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

        // Burn native token by sending to BLACKHOLE_ADDRESS
        payable(BLACKHOLE_ADDRESS).transfer(msg.value);

        // Send Teleporter message.
        bytes memory messageData = abi.encode(recipient, msg.value);

        uint256 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationChainID: sourceBlockchainID,
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
            sender: msg.sender,
            recipient: recipient,
            amount: msg.value,
            feeContractAddress: feeContractAddress,
            feeAmount: adjustedAmount,
            teleporterMessageID: messageID
        });
    }
}
