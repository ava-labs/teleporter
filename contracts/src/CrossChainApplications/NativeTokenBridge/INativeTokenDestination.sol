// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterFeeInfo} from "../../Teleporter/ITeleporterMessenger.sol";

/**
 * @dev Interface that describes functionalities for a contract that can mint native tokens when
 * paired with a {INativeTokenSource} or {IERC20TokenSource} contract that will lock tokens on another chain.
 */
interface INativeTokenDestination {
    /**
     * @dev Emitted when tokens are burned in the destination contract and to be unlocked on the source contract.
     */
    event TransferToSource(
        address indexed sender,
        address indexed recipient,
        uint256 indexed teleporterMessageID,
        uint256 amount
    );

    /**
     * @dev Emitted when tokens are not minted in order to collateralize the source contract.
     */
    event CollateralAdded(uint256 amount, uint256 remaining, address recipient);

    /**
     * @dev Emitted when minting native tokens.
     */
    event NativeTokensMinted(address indexed recipient, uint256 amount);

    /**
     * @dev Emitted when reporting total burned tx fees to source chain.
     */
    event ReportTotalBurnedTxFees(
        uint256 indexed teleporterMessageID,
        uint256 burnAddressBalance
    );

    /**
     * @dev Burns native tokens on the destination contract chain, and sends a message to the source
     * contract to unlock corresponding tokens.
     */
    function transferToSource(
        address recipient,
        TeleporterFeeInfo calldata feeInfo,
        address[] calldata allowedRelayerAddresses
    ) external payable;

    /**
     * @dev Returns true if the reserve imbalance for this contract has been accounted for.
     *  When this is true, all tokens sent to this chain will be minted, and sending tokens
     *  to the source chain is allowed.
     */
    function isCollateralized() external view returns (bool);

    /**
     * @dev Returns a best-estimate (upper bound) of tokens in circulation on this chain.
     */
    function totalSupply() external view returns (uint256);
}
