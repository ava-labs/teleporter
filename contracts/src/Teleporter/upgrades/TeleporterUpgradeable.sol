// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterRegistry} from "./TeleporterRegistry.sol";
import {ITeleporterReceiver} from "../ITeleporterReceiver.sol";
import {ITeleporterMessenger} from "../ITeleporterMessenger.sol";

/**
 * @dev TeleporterUpgradeable provides upgrade utility for applications built on top
 * of the Teleporter protocol by integrating with the {TeleporterRegistry}.
 *
 * This contract is intended to be inherited by other contracts that wish to use the
 * upgrade mechanism. It provides an interface that restricts access to only Teleporter
 * versions that are greater than or equal to `minTeleporterVersion`.
 */
abstract contract TeleporterUpgradeable is ITeleporterReceiver {
    TeleporterRegistry public immutable teleporterRegistry;

    /**
     * @dev A mapping that keeps track of paused Teleporter addresses.
     */
    mapping(address teleporterAddress => bool paused)
        private _pausedTeleporterAddresses;

    /**
     * @dev The minimum required Teleporter version that the contract is allowed
     * to receive messages from. Should only be updated by `_setMinTeleporterVersion`
     */
    uint256 private _minTeleporterVersion;

    /**
     * @dev Emitted when `minTeleporterVersion` is updated.
     */
    event MinTeleporterVersionUpdated(
        uint256 indexed oldMinTeleporterVersion,
        uint256 indexed newMinTeleporterVersion
    );

    /**
     * @dev Emitted when a Teleporter address is paused.
     */
    event TeleporterAddressPaused(address indexed teleporterAddress);

    /**
     * @dev Initializes the {TeleporterUpgradeable} contract by getting `teleporterRegistry`
     * instance and setting `_minTeleporterVersion`.
     */
    constructor(address teleporterRegistryAddress) {
        require(
            teleporterRegistryAddress != address(0),
            "TeleporterUpgradeable: zero teleporter registry address"
        );

        teleporterRegistry = TeleporterRegistry(teleporterRegistryAddress);
        _minTeleporterVersion = teleporterRegistry.latestVersion();
    }

    /**
     * @dev See {ITeleporterReceiver-receiveTeleporterMessage}
     * Requirements:
     *
     * - `msg.sender` must be a Teleporter version greater than or equal to `minTeleporterVersion`.
     */
    function receiveTeleporterMessage(
        bytes32 originBlockchainID,
        address originSenderAddress,
        bytes calldata message
    ) external {
        // Checks that `msg.sender` matches a Teleporter version greater than or equal to `minTeleporterVersion`.
        require(
            teleporterRegistry.getVersionFromAddress(msg.sender) >=
                _minTeleporterVersion,
            "TeleporterUpgradeable: invalid Teleporter sender"
        );

        // Check against the paused Teleporter addresses.
        require(
            !isTeleporterAddressPaused(msg.sender),
            "TeleporterUpgradeable: Teleporter address paused"
        );

        _receiveTeleporterMessage(
            originBlockchainID,
            originSenderAddress,
            message
        );
    }

    /**
     * @dev Updates the minimum Teleporter version allowed for delivering Teleporer messages
     * to this contract.
     *
     * To prevent anyone from being able to call this function, which would disallow messages
     * from old Teleporter versions from being received, this function should be safeguarded with access
     * controls. This is done by overriding the implementation of {_checkTeleporterUpgradeAccess}.
     */
    function updateMinTeleporterVersion(uint256 version) public virtual {
        _checkTeleporterUpgradeAccess();
        _setMinTeleporterVersion(version);
    }

    /**
     * @dev Pauses a Teleporter address from interacting with this contract.
     * After pausing a Teleporter address, it will no longer be able to deliver messages
     * to this contract, and this contract will not send messages through that Teleporter address.
     * The address does not need to be registered with the Teleporter registry.
     * Emits a {TeleporterAddressPaused} event if successfully paused.
     * Requirements:
     *
     * - `msg.sender` must have Teleporter upgrade access.
     * - `teleporterAddress` is not already paused.
     */
    function pauseTeleporterAddress(address teleporterAddress) public virtual {
        _checkTeleporterUpgradeAccess();
        require(
            !isTeleporterAddressPaused(teleporterAddress),
            "TeleporterUpgradeable: address already paused"
        );
        _pausedTeleporterAddresses[teleporterAddress] = true;
        emit TeleporterAddressPaused(teleporterAddress);
    }

    /**
     * @dev Public getter for `_minTeleporterVersion`.
     */
    function getMinTeleporterVersion() public view returns (uint256) {
        return _minTeleporterVersion;
    }

    /**
     * @dev Checks if a Teleporter address is paused.
     */
    function isTeleporterAddressPaused(
        address teleporterAddress
    ) public view virtual returns (bool) {
        return _pausedTeleporterAddresses[teleporterAddress];
    }

    /**
     * @dev Sets the minimum Teleporter version allowed for delivering Teleporter messages.
     * Emits a {MinTeleporterVersionUpdated} event if the minimum Teleporter version was updated.
     * Requirements:
     *
     * - `version` must be less than or equal to the latest Teleporter version.
     * - `version` must be greater than the current minimum Teleporter version.
     *
     */
    function _setMinTeleporterVersion(uint256 version) internal virtual {
        uint256 latestTeleporterVersion = teleporterRegistry.latestVersion();
        uint256 oldMinTeleporterVersion = _minTeleporterVersion;

        require(
            version <= latestTeleporterVersion,
            "TeleporterUpgradeable: invalid Teleporter version"
        );
        require(
            version > oldMinTeleporterVersion,
            "TeleporterUpgradeable: not greater than current minimum version"
        );

        _minTeleporterVersion = version;
        emit MinTeleporterVersionUpdated(oldMinTeleporterVersion, version);
    }

    /**
     * @dev Receives Teleporter messages and handles accordingly.
     * This function should be overridden by contracts that inherit from this contract.
     */
    function _receiveTeleporterMessage(
        bytes32 originBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) internal virtual;

    /**
     * @dev Checks that the caller has access to update the minimum Teleporter version
     * allowed for delivering Teleporter messages to this contract.
     *
     * This function should be overridden by contracts that inherit from this contract.
     */
    function _checkTeleporterUpgradeAccess() internal virtual;

    /**
     * @dev Returns the Teleporter messenger used to send Teleporter messages,
     * and checks that the Teleporter messenger is not paused.
     *
     * By default returns the latest Teleporter messenger, but can be overriden to
     * return a Teleporter messenger of specific version.
     */
    function _getTeleporterMessenger()
        internal
        view
        virtual
        returns (ITeleporterMessenger)
    {
        ITeleporterMessenger teleporter = teleporterRegistry
            .getLatestTeleporter();
        require(
            !isTeleporterAddressPaused(address(teleporter)),
            "TeleporterUpgradeable: Teleporter sending paused"
        );

        return teleporter;
    }
}
