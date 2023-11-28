// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ReentrancyGuard} from "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import {IWarpMessenger} from "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import {IERC20TokenSource} from "./IERC20TokenSource.sol";
import {ITokenSource} from "./ITokenSource.sol";
import {ITeleporterMessenger, TeleporterMessageInput, TeleporterFeeInfo} from "../../Teleporter/ITeleporterMessenger.sol";
import {ITeleporterReceiver} from "../../Teleporter/ITeleporterReceiver.sol";
import {SafeERC20TransferFrom} from "../../Teleporter/SafeERC20TransferFrom.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract ERC20TokenSource is
    ITeleporterReceiver,
    IERC20TokenSource,
    ITokenSource,
    ReentrancyGuard
{
    // The address where the burned transaction fees are credited.
    // Defined as BLACKHOLE_ADDRESS at
    // https://github.com/ava-labs/subnet-evm/blob/e23ab058d039ff9c8469c89b139d21d52c4bd283/constants/constants.go
    address public constant BURNED_TX_FEES_ADDRESS =
        0x0100000000000000000000000000000000000000;
    uint256 public constant MINT_NATIVE_TOKENS_REQUIRED_GAS = 100_000;
    // Used to keep track of tokens burned through transactions on the destination chain. They can
    // be reported to this contract to burn an equivalent number of tokens on this chain.
    uint256 public destinationBurnedTotal;
    bytes32 public immutable destinationBlockchainID;
    address public immutable nativeTokenDestinationAddress;
    address public immutable erc20ContractAddress;

    // Used for sending and receiving Teleporter messages.
    ITeleporterMessenger public immutable teleporterMessenger;

    constructor(
        address teleporterMessengerAddress,
        bytes32 destinationBlockchainID_,
        address nativeTokenDestinationAddress_,
        address erc20ContractAddress_
    ) {
        require(
            teleporterMessengerAddress != address(0),
            "ERC20TokenSource: zero TeleporterMessenger address"
        );
        teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);

        require(
            destinationBlockchainID_ != bytes32(0),
            "ERC20TokenSource: zero destination chain ID"
        );
        require(
            destinationBlockchainID_ !=
                IWarpMessenger(0x0200000000000000000000000000000000000005)
                    .getBlockchainID(),
            "ERC20TokenSource: cannot bridge with same blockchain"
        );
        destinationBlockchainID = destinationBlockchainID_;

        require(
            nativeTokenDestinationAddress_ != address(0),
            "ERC20TokenSource: zero destination contract address"
        );
        nativeTokenDestinationAddress = nativeTokenDestinationAddress_;

        require(
            erc20ContractAddress_ != address(0),
            "ERC20TokenSource: invalid ERC20 contract address"
        );
        erc20ContractAddress = erc20ContractAddress_;
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
            "ERC20TokenSource: unauthorized TeleporterMessenger contract"
        );

        // Only allow messages from the destination chain.
        require(
            senderBlockchainID == destinationBlockchainID,
            "ERC20TokenSource: invalid destination chain"
        );

        // Only allow the partner contract to send messages.
        require(
            senderAddress == nativeTokenDestinationAddress,
            "ERC20TokenSource: unauthorized sender"
        );

        // Decode the payload to recover the action and corresponding function parameters
        (SourceAction action, bytes memory actionData) = abi.decode(
            message,
            (SourceAction, bytes)
        );

        // Route to the appropriate function.
        if (action == SourceAction.Unlock) {
            (address recipient, uint256 amount) = abi.decode(
                actionData,
                (address, uint256)
            );
            _unlockTokens(recipient, amount);
        } else if (action == SourceAction.Burn) {
            uint256 newBurnTotal = abi.decode(actionData, (uint256));
            _updateDestinationBurnedTotal(newBurnTotal);
        } else {
            revert("ERC20TokenSource: invalid action");
        }
    }

    /**
     * @dev See {IERC20TokenSource-transferToDestination}.
     */
    function transferToDestination(
        address recipient,
        uint256 totalAmount,
        uint256 feeAmount,
        address[] calldata allowedRelayerAddresses
    ) external nonReentrant {
        // The recipient cannot be the zero address.
        require(
            recipient != address(0),
            "ERC20TokenSource: zero recipient address"
        );

        // Lock tokens in this contract. Supports "fee/burn on transfer" ERC20 token
        // implementations by only bridging the actual balance increase reflected by the call
        // to transferFrom.
        uint256 adjustedAmount = SafeERC20TransferFrom.safeTransferFrom(
            IERC20(erc20ContractAddress),
            totalAmount
        );

        // Ensure that the adjusted amount is greater than the fee to be paid.
        require(
            adjustedAmount > feeAmount,
            "ERC20TokenSource: insufficient adjusted amount"
        );

        // Allow the Teleporter messenger to spend the fee amount.
        if (feeAmount > 0) {
            SafeERC20.safeIncreaseAllowance(
                IERC20(erc20ContractAddress),
                address(teleporterMessenger),
                feeAmount
            );
        }

        uint256 transferAmount = adjustedAmount - feeAmount;

        uint256 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationBlockchainID: destinationBlockchainID,
                destinationAddress: nativeTokenDestinationAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: erc20ContractAddress,
                    amount: feeAmount
                }),
                requiredGasLimit: MINT_NATIVE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: allowedRelayerAddresses,
                message: abi.encode(recipient, transferAmount)
            })
        );

        emit TransferToDestination({
            sender: msg.sender,
            recipient: recipient,
            transferAmount: transferAmount,
            feeAmount: feeAmount,
            teleporterMessageID: messageID
        });
    }

    /**
     * @dev Unlocks tokens to recipient.
     */
    function _unlockTokens(address recipient, uint256 amount) private {
        require(
            recipient != address(0),
            "ERC20TokenSource: zero recipient address"
        );

        // Transfer to recipient
        SafeERC20.safeTransfer(IERC20(erc20ContractAddress), recipient, amount);

        emit UnlockTokens(recipient, amount);
    }

    /**
     * @dev Sends tokens to BURN_FOR_TRANSFER_ADDRESS.
     */
    function _burnTokens(uint256 amount) private {
        SafeERC20.safeTransfer(
            IERC20(erc20ContractAddress),
            BURNED_TX_FEES_ADDRESS,
            amount
        );
        emit BurnTokens(amount);
    }

    /**
     * @dev Update destinationBurnedTotal sent from destination chain
     */
    function _updateDestinationBurnedTotal(uint256 newBurnTotal) private {
        if (newBurnTotal > destinationBurnedTotal) {
            uint256 difference = newBurnTotal - destinationBurnedTotal;
            _burnTokens(difference);
            destinationBurnedTotal = newBurnTotal;
        }
    }
}
