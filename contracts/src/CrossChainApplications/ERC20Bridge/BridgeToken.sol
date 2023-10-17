// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

/**
 * @dev BridgeToken is an ERC20Burnable token contract that is associated with a specific native chain bridge and asset, and is only mintable by the bridge contract on this chain.
 */
contract BridgeToken is ERC20Burnable {
    address public immutable bridgeContract;

    bytes32 public immutable nativeChainID;
    address public immutable nativeBridge;
    address public immutable nativeAsset;

    uint8 private immutable _decimals;

    /**
     * @dev Initializes a BridgeToken instance.
     */
    constructor(
        bytes32 sourceChainID,
        address sourceBridge,
        address sourceAsset,
        string memory tokenName,
        string memory tokenSymbol,
        uint8 tokenDecimals
    ) ERC20(tokenName, tokenSymbol) {
        require(
            sourceChainID != bytes32(0),
            "BridgeToken: zero source chain id"
        );
        require(
            sourceBridge != address(0),
            "BridgeToken: zero source bridge address"
        );
        require(
            sourceAsset != address(0),
            "BridgeToken: zero source asset address"
        );
        bridgeContract = msg.sender;
        nativeChainID = sourceChainID;
        nativeBridge = sourceBridge;
        nativeAsset = sourceAsset;
        _decimals = tokenDecimals;
    }

    /**
     * @dev Mints tokens to `account` if called by original `bridgeContract`.
     */
    function mint(address account, uint256 amount) public {
        require(msg.sender == bridgeContract, "BridgeToken: unauthorized");
        _mint(account, amount);
    }

    function decimals() public view virtual override returns (uint8) {
        return _decimals;
    }
}
