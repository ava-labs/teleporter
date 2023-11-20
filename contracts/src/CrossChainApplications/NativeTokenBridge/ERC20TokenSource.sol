// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ReentrancyGuard} from "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import {IWarpMessenger} from "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import {IERC20TokenSource} from "./IERC20TokenSource.sol";
import {ITokenSource} from "./ITokenSource.sol";
import {ITeleporterMessenger, TeleporterMessageInput, TeleporterFeeInfo} from "../../Teleporter/ITeleporterMessenger.sol";
import {TeleporterOwnerUpgradeable} from "../../Teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
import {ITeleporterReceiver} from "../../Teleporter/ITeleporterReceiver.sol";
import {SafeERC20TransferFrom} from "../../Teleporter/SafeERC20TransferFrom.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract ERC20TokenSource is
    ITeleporterReceiver,
    TeleporterOwnerUpgradeable,
    IERC20TokenSource,
    ITokenSource,
    ReentrancyGuard
{
    // Designated Blackhole Address. Tokens are sent here to be "burned" before sending an unlock
    // message to the source chain. Different from the burned tx fee address so they can be
    // tracked separately.
    address public constant BLACKHOLE_ADDRESS =
        0x0100000000000000000000000000000000000001;
    uint256 public constant MINT_NATIVE_TOKENS_REQUIRED_GAS = 100_000;
    // Used to keep track of tokens burned through transactions on the destination chain. They can
    // be reported to this contract to burn an equivalent number of tokens on this chain.
    uint256 public destinationChainBurnedBalance = 0;
    bytes32 public immutable destinationBlockchainID;
    address public immutable nativeTokenDestinationAddress;
    address public immutable erc20ContractAddress;

    constructor(
        address teleporterRegistryAddress,
        bytes32 destinationBlockchainID_,
        address nativeTokenDestinationAddress_,
        address erc20ContractAddress_
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress) {
        require(
            teleporterRegistryAddress != address(0),
            "NativeTokenSource: zero TeleporterRegistry address"
        );

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
    ) external onlyAllowedTeleporter nonReentrant {
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
            uint256 newBurnBalance = abi.decode(actionData, (uint256));
            _updateDestinationChainBurnedBalance(newBurnBalance);
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
        ITeleporterMessenger teleporterMessenger = teleporterRegistry
            .getLatestTeleporter();

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
                destinationChainID: destinationBlockchainID,
                destinationAddress: nativeTokenDestinationAddress,
                feeInfo: TeleporterFeeInfo({
                    contractAddress: erc20ContractAddress,
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
     * @dev Sends tokens to BLACKHOLE_ADDRESS.
     */
    function _burnTokens(uint256 amount) private {
        SafeERC20.safeTransfer(
            IERC20(erc20ContractAddress),
            BLACKHOLE_ADDRESS,
            amount
        );
        emit BurnTokens(amount);
    }

    /**
     * @dev Update destinationChainBurnedBalance sent from destination chain
     */
    function _updateDestinationChainBurnedBalance(
        uint256 newBurnBalance
    ) private {
        if (newBurnBalance > destinationChainBurnedBalance) {
            uint256 difference = newBurnBalance - destinationChainBurnedBalance;
            _burnTokens(difference);
            destinationChainBurnedBalance = newBurnBalance;
        }
    }
}
