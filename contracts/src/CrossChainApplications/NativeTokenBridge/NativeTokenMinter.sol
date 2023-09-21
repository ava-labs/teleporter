//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@subnet-evm-contracts/AllowList.sol";
import "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import "./INativeTokenMinter.sol";
import "../../Teleporter/ITeleporterMessenger.sol";
import "../../Teleporter/ITeleporterReceiver.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

// Precompiled Native Minter Contract Address
address constant MINTER_ADDRESS = 0x0200000000000000000000000000000000000001;
// Designated Blackhole Address
address constant BLACKHOLE_ADDRESS = 0x0100000000000000000000000000000000000000;

contract NativeTokenMinter is ITeleporterReceiver, INativeTokenMinter, AllowList {
  INativeTokenMinter _nativeMinter = INativeTokenMinter(MINTER_ADDRESS);

  address public constant WARP_PRECOMPILE_ADDRESS =
      0x0200000000000000000000000000000000000005;
  bytes32 public immutable currentChainID;

  // Used for sending an receiving Teleporter messages.
  ITeleporterMessenger public immutable teleporterMessenger;

  error InvalidTeleporterMessengerAddress();
  error Unauthorized();

  constructor(address teleporterMessengerAddress) AllowList(MINTER_ADDRESS) {
    if (teleporterMessengerAddress == address(0)) {
        revert InvalidTeleporterMessengerAddress();
    }

    teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);
    currentChainID = WarpMessenger(WARP_PRECOMPILE_ADDRESS)
        .getBlockchainID();
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
  ) external {
      // Only allow the Teleporter messenger to deliver messages.
      if (msg.sender != address(teleporterMessenger)) {
          revert Unauthorized();
      }

      // Decode the payload to recover the action and corresponding function parameters
      (BridgeAction action, bytes memory actionData) = abi.decode(
          message,
          (BridgeAction, bytes)
      );

      // Route to the appropriate function.
      if (action == BridgeAction.Create) {
          (
              address nativeContractAddress,
              string memory nativeName,
              string memory nativeSymbol,
              uint8 nativeDecimals
          ) = abi.decode(actionData, (address, string, string, uint8));
          _createBridgeToken({
              nativeChainID: nativeChainID,
              nativeBridgeAddress: nativeBridgeAddress,
              nativeContractAddress: nativeContractAddress,
              nativeName: nativeName,
              nativeSymbol: nativeSymbol,
              nativeDecimals: nativeDecimals
          });
      } else if (action == BridgeAction.Mint) {
          (
              address nativeContractAddress,
              address recipient,
              uint256 amount
          ) = abi.decode(actionData, (address, address, uint256));
          _mintBridgeTokens(
              nativeChainID,
              nativeBridgeAddress,
              nativeContractAddress,
              recipient,
              amount
          );
      } else if (action == BridgeAction.Transfer) {
          (
              bytes32 destinationChainID,
              address destinationBridgeAddress,
              address nativeContractAddress,
              address recipient,
              uint256 totalAmount,
              uint256 secondaryFeeAmount
          ) = abi.decode(
                  actionData,
                  (bytes32, address, address, address, uint256, uint256)
              );
          _transferBridgeTokens({
              sourceChainID: nativeChainID,
              sourceBridgeAddress: nativeBridgeAddress,
              destinationChainID: destinationChainID,
              destinationBridgeAddress: destinationBridgeAddress,
              nativeContractAddress: nativeContractAddress,
              recipient: recipient,
              totalAmount: totalAmount,
              secondaryFeeAmount: secondaryFeeAmount
          });
      } else {
          revert InvalidAction();
      }
  }

  // Mints [amount] number of ERC20 token to [to] address.
  function mint(address to, uint256 amount) external onlyOwner {
    // Mints [amount] number of native coins (gas coin) to [msg.sender] address.
    // Calls NativeMinter precompile through INativeMinter interface.
    _nativeMinter.mintNativeCoin(to, amount);
    emit Mintdrawal(_msgSender(), amount);
  }

  // Burns [amount] number of ERC20 token from [from] address.
  function burn(address from, uint256 amount) external onlyOwner {
    // Burn native token by sending to BLACKHOLE_ADDRESS
    payable(BLACKHOLE_ADDRESS).transfer(msg.value);
    // Mint ERC20 token.
    emit Deposit(_msgSender(), msg.value);
  }
}
