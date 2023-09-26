//SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import "./INativeTokenReceiver.sol";
import "../../Teleporter/ITeleporterMessenger.sol";
import "../../Teleporter/ITeleporterReceiver.sol";
import "../../Teleporter/SafeERC20TransferFrom.sol";

contract NativeTokenMinter is ITeleporterReceiver, INativeTokenReceiver, ReentrancyGuard {
  address public constant WARP_PRECOMPILE_ADDRESS =
      0x0200000000000000000000000000000000000005;

  uint256 public constant MINT_NATIVE_TOKENS_REQUIRED_GAS = 200_000; // TODO this is a placeholder
  bytes32 public immutable currentChainID;
  bytes32 public immutable partnerChainID;

  // Used for sending an receiving Teleporter messages.
  ITeleporterMessenger public immutable teleporterMessenger;

  error InvalidTeleporterMessengerAddress();
  error InvalidRecipientAddress();
  error InvalidSourceChain();
  error InvalidPartnerContractAddress();
  error CannotBridgeTokenWithinSameChain();
  error Unauthorized();
  error InsufficientPayment();
  error InsufficientAdjustedAmount(uint256 adjustedAmount, uint256 feeAmount);

  constructor(address teleporterMessengerAddress, bytes32 partnerChainID_) {
    if (teleporterMessengerAddress == address(0)) {
        revert InvalidTeleporterMessengerAddress();
    }
    teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);
    currentChainID = WarpMessenger(WARP_PRECOMPILE_ADDRESS)
        .getBlockchainID();

    if (partnerChainID_ == currentChainID) {
        revert CannotBridgeTokenWithinSameChain();
    }
    partnerChainID = partnerChainID_;
  }

  /**
    * @dev See {ITeleporterReceiver-receiveTeleporterMessage}.
    *
    * Receives a Teleporter message and routes to the appropriate internal function call.
    */
  function receiveTeleporterMessage(
    bytes32 nativeChainID,
    address nativeBridgeAddress,
    bytes calldata message
  ) external nonReentrant() {

    // Only allow the Teleporter messenger to deliver messages.
    if (msg.sender != address(teleporterMessenger)) {
      revert Unauthorized();
    }
    // Only allow messages from the partner chain.
    if (nativeChainID != partnerChainID) {
      revert InvalidSourceChain();
    }
    // Only allow the partner contract to send messages.
    if (nativeBridgeAddress != address(this)) {
      revert InvalidPartnerContractAddress();
    }

    (address recipient, uint256 amount) = abi.decode(message, (address, uint256));

    // TODO set up starting threshold.

    // Send to recipient
    payable(recipient).transfer(amount);

    emit UnlockTokens(recipient, amount);
  }

    /**
    * @dev See {IERC20Bridge-bridgeTokens}.
    *
    * Requirements:
    *
    * - `msg.value` must be greater than the fee amount.
    */
    function bridgeTokens(
        address recipient,
        address feeTokenContractAddress,
        uint256 feeAmount
    ) external payable nonReentrant {
        // The recipient cannot be the zero address.
        if (recipient == address(0)) {
            revert InvalidRecipientAddress();
        }

        // Lock tokens in this bridge instance. Supports "fee/burn on transfer" ERC20 token
        // implementations by only bridging the actual balance increase reflected by the call
        // to transferFrom.
        uint256 adjustedAmount = SafeERC20TransferFrom.safeTransferFrom(
            IERC20(feeTokenContractAddress),
            feeAmount
        );

        // Ensure that the adjusted amount is greater than the fee to be paid.
        // The secondary fee amount is not used in this case (and can assumed to be 0) since bridging
        // a native token to another chain only ever involves a single cross-chain message.
        if (adjustedAmount <= feeAmount) {
            revert InsufficientAdjustedAmount(adjustedAmount, feeAmount);
        }

        // Send Teleporter message.
        bytes memory messageData = abi.encode(recipient, msg.value);

        uint256 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationChainID: partnerChainID,
                destinationAddress: address(this),
                feeInfo: TeleporterFeeInfo({
                    contractAddress: feeTokenContractAddress,
                    amount: feeAmount
                }),
                requiredGasLimit: MINT_NATIVE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: new address[](0),
                message: messageData
            })
        );

        emit BridgeTokens({
            tokenContractAddress: feeTokenContractAddress,
            teleporterMessageID: messageID,
            destinationChainID: partnerChainID,
            destinationBridgeAddress: address(this),
            recipient: recipient,
            transferAmount: msg.value,
            feeAmount: feeAmount
        });
    }
}
