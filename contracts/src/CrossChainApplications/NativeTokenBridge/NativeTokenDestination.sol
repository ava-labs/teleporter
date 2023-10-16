// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import "@subnet-evm-contracts/interfaces/IAllowList.sol";
import "@subnet-evm-contracts/interfaces/INativeMinter.sol";
import "./INativeTokenDestination.sol";
import "../../Teleporter/ITeleporterMessenger.sol";
import "../../Teleporter/ITeleporterReceiver.sol";
import "../../Teleporter/SafeERC20TransferFrom.sol";

// Native Minter Precompile Contract Address
address constant MINTER_ADDRESS = 0x0200000000000000000000000000000000000001;
// This is the address where the burned base fee go. We use this to send the amount of
// burned tokens back to the source chain.
// TODO implement mechanism to report burned tx fees to source chian.
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
    address public immutable nativeTokenSourceAddress;
    // We will not mint the first `tokenReserve` tokens sent to this subnet.
    // This should be constructed to match the initial token supply of this subnet.
    // This will mean we will not mint tokens until the source contact is collateralized.
    uint256 public tokenReserve;

    // Used for sending an receiving Teleporter messages.
    ITeleporterMessenger public immutable teleporterMessenger;

    // TODO we probably want to add the original token supply from this chain to the constructor.
    constructor(
        address teleporterMessengerAddress,
        bytes32 sourceBlockchainID_,
        address nativeTokenSourceAddress_,
        uint256 tokenReserve_
    ) {
        currentBlockchainID = WarpMessenger(WARP_PRECOMPILE_ADDRESS)
            .getBlockchainID();

        require(
            teleporterMessengerAddress != address(0),
            "Invalid Teleporter Messenger Address"
        );
        teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);

        require(sourceBlockchainID_ != bytes32(0), "Invalid Source Chain ID");
        require(
            sourceBlockchainID_ != currentBlockchainID,
            "Cannot Bridge With Same Blockchain"
        );
        sourceBlockchainID = sourceBlockchainID_;

        require(
            nativeTokenSourceAddress_ != address(0),
            "Invalid Source Contract Address"
        );
        nativeTokenSourceAddress = nativeTokenSourceAddress_;

        tokenReserve = tokenReserve_;
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
            "Unauthorized teleporter contract"
        );

        // Only allow messages from the source chain.
        require(
            senderBlockchainID == sourceBlockchainID,
            "Invalid Source Chain"
        );

        // Only allow the partner contract to send messages.
        require(senderAddress == nativeTokenSourceAddress, "Unauthorized Sender");

        (address recipient, uint256 amount) = abi.decode(
            message,
            (address, uint256)
        );
        require(recipient != address(0), "Invalid Recipient Address");

        uint256 adjustedAmount = amount;
        if (tokenReserve > 0) {
            if (amount > tokenReserve) {
                emit CollateralAdded({amount: tokenReserve, remaining: 0});
                adjustedAmount = amount - tokenReserve;
                tokenReserve = 0;
            } else {
                tokenReserve -= amount;
                emit CollateralAdded({amount: amount, remaining: tokenReserve});
                return;
            }
        }

        // Calls NativeMinter precompile through INativeMinter interface.
        _nativeMinter.mintNativeCoin(recipient, adjustedAmount);
        emit MintNativeTokens(recipient, adjustedAmount);
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

        require(
            tokenReserve == 0,
            "Cannot release tokens until source contract is collateralized"
        );

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

        uint256 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationChainID: sourceBlockchainID,
                destinationAddress: nativeTokenSourceAddress,
                feeInfo: TeleporterFeeInfo({
                    contractAddress: feeContractAddress,
                    amount: adjustedAmount
                }),
                requiredGasLimit: TRANSFER_NATIVE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(recipient, msg.value)
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
