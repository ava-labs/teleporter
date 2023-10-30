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

// The address where the burned transaction fees are credited.
// TODO implement mechanism to report burned tx fees to source chian.
address constant BURNED_TX_FEES_ADDRESS = 0x0100000000000000000000000000000000000000;
// Designated Blackhole Address. Tokens are sent here to be "burned" before sending an unlock
// message to the source chain. Different from the burned tx fee address so they can be
// tracked separately.
address constant BLACKHOLE_ADDRESS = 0x0100000000000000000000000000000000000001;

contract NativeTokenDestination is
    ITeleporterReceiver,
    INativeTokenDestination,
    ReentrancyGuard
{
    using SafeERC20 for IERC20;

    INativeMinter private immutable _nativeMinter =
        INativeMinter(0x0200000000000000000000000000000000000001);

    uint256 public constant TRANSFER_NATIVE_TOKENS_REQUIRED_GAS = 150_000; // TODO this is a placeholder
    bytes32 public immutable currentBlockchainID;
    bytes32 public immutable sourceBlockchainID;
    address public immutable nativeTokenSourceAddress;
    // The first `initialReserveImbalance` tokens sent to this subnet will not be minted.
    // `initialReserveImbalance` should be constructed to match the initial token supply of this subnet.
    // This means tokens will not be minted until the source contact is collateralized.
    uint256 public immutable initialReserveImbalance;
    uint256 public currentReserveImbalance;

    // Used for sending and receiving Teleporter messages.
    ITeleporterMessenger public immutable teleporterMessenger;

    constructor(
        address teleporterMessengerAddress,
        bytes32 sourceBlockchainID_,
        address nativeTokenSourceAddress_,
        uint256 initialReserveImbalance_
    ) {
        currentBlockchainID = WarpMessenger(
            0x0200000000000000000000000000000000000005
        ).getBlockchainID();

        require(
            teleporterMessengerAddress != address(0),
            "NativeTokenDestination: zero TeleporterMessenger Address"
        );
        teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);

        require(
            sourceBlockchainID_ != bytes32(0),
            "NativeTokenDestination: zero Source Chain ID"
        );
        require(
            sourceBlockchainID_ != currentBlockchainID,
            "NativeTokenDestination: Cannot Bridge With Same Blockchain"
        );
        sourceBlockchainID = sourceBlockchainID_;

        require(
            nativeTokenSourceAddress_ != address(0),
            "NativeTokenDestination: zero Source Contract Address"
        );
        nativeTokenSourceAddress = nativeTokenSourceAddress_;

        initialReserveImbalance = initialReserveImbalance_;
        currentReserveImbalance = initialReserveImbalance_;
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
            "NativeTokenDestination: Unauthorized TeleporterMessenger contract"
        );

        // Only allow messages from the source chain.
        require(
            senderBlockchainID == sourceBlockchainID,
            "NativeTokenDestination: Invalid Source Chain"
        );

        // Only allow the partner contract to send messages.
        require(
            senderAddress == nativeTokenSourceAddress,
            "NativeTokenDestination: Unauthorized Sender"
        );

        (address recipient, uint256 amount) = abi.decode(
            message,
            (address, uint256)
        );
        require(
            recipient != address(0),
            "NativeTokenDestination: zero Recipient Address"
        );
        require(amount != 0, "NativeTokenDestination: Transfer value of 0");

        uint256 adjustedAmount = amount;
        if (currentReserveImbalance > 0) {
            if (amount > currentReserveImbalance) {
                emit CollateralAdded({amount: currentReserveImbalance, remaining: 0});
                adjustedAmount = amount - currentReserveImbalance;
                currentReserveImbalance = 0;
            } else {
                currentReserveImbalance -= amount;
                emit CollateralAdded({amount: amount, remaining: currentReserveImbalance});
                return;
            }
        }

        // Calls NativeMinter precompile through INativeMinter interface.
        _nativeMinter.mintNativeCoin(recipient, adjustedAmount);
        emit NativeTokensMinted(recipient, adjustedAmount);
    }

    /**
     * @dev See {INativeTokenDestination-transferToSource}.
     */
    function transferToSource(
        address recipient,
        address feeContractAddress,
        uint256 feeAmount,
        address[] calldata allowedRelayerAddresses
    ) external payable nonReentrant {
        // The recipient cannot be the zero address.
        require(
            recipient != address(0),
            "NativeTokenDestination: zero Recipient Address"
        );

        require(
            currentReserveImbalance == 0,
            "NativeTokenDestination: Cannot release tokens until source contract is collateralized"
        );

        // Lock tokens in this bridge instance. Supports "fee/burn on transfer" ERC20 token
        // implementations by only bridging the actual balance increase reflected by the call
        // to transferFrom.
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

        // Burn native token by sending to BLACKHOLE_ADDRESS
        payable(BLACKHOLE_ADDRESS).transfer(msg.value);

        uint256 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationChainID: sourceBlockchainID,
                destinationAddress: nativeTokenSourceAddress,
                feeInfo: TeleporterFeeInfo({
                    contractAddress: feeContractAddress,
                    amount: adjustedFeeAmount
                }),
                requiredGasLimit: TRANSFER_NATIVE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: allowedRelayerAddresses,
                message: abi.encode(recipient, msg.value)
            })
        );

        emit TransferToSource({
            sender: msg.sender,
            recipient: recipient,
            amount: msg.value,
            teleporterMessageID: messageID
        });
    }

    function isCollateralized() public view returns (bool) {
        return currentReserveImbalance == 0;
    }
}
