// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: BSD-3-Clause

pragma solidity 0.8.18;

import "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import "@openzeppelin/contracts/utils/Address.sol";

abstract contract WarpProtocolRegistry {
    // Address that the out-of-band warp message sets as the "source" address.
    // The address is obviously not owned by any EOA or smart contract account, so it
    // can not possibly be the source address of any other warp message emitted by the VM.
    address public constant VALIDATORS_SOURCE_ADDRESS = address(0);

    WarpMessenger public constant WARP_MESSENGER =
        WarpMessenger(0x0200000000000000000000000000000000000005);

    bytes32 internal immutable _chainID;

    uint256 internal _latestVersion;

    mapping(uint256 => address) internal _versionToAddress;

    // Errors
    error InvalidWarpMessage();
    error InvalidOriginChainID();
    error InvalidOriginSenderAddress();
    error InvalidDestinationChainID();
    error InvalidDestinationAddress();
    error InvalidProtocolAddress();
    error InvalidProtocolVersion();

    constructor() {
        _latestVersion = 0;
        _chainID = WARP_MESSENGER.getBlockchainID();
    }

    function addProtocolVersion() external {
        _addProtocolVersion();
    }

    function getVersionToAddress(
        uint256 version
    ) external view returns (address) {
        return _getVersionToAddress(version);
    }

    function getLatestVersion() external view returns (uint256) {
        return _latestVersion;
    }

    function _addProtocolVersion() internal virtual {
        (WarpMessage memory message, bool valid) = WARP_MESSENGER
            .getVerifiedWarpMessage();
        if (!valid) {
            revert InvalidWarpMessage();
        }
        if (message.originChainID != _chainID) {
            revert InvalidOriginChainID();
        }
        if (message.originSenderAddress != VALIDATORS_SOURCE_ADDRESS) {
            revert InvalidOriginSenderAddress();
        }
        if (message.destinationChainID != _chainID) {
            revert InvalidDestinationChainID();
        }
        if (message.destinationAddress != address(this)) {
            revert InvalidDestinationAddress();
        }

        (uint256 version, address protocolAddress) = abi.decode(
            message.payload,
            (uint256, address)
        );

        if (version != _latestVersion) {
            revert InvalidProtocolVersion();
        }
        if (!Address.isContract(protocolAddress)) {
            revert InvalidProtocolAddress();
        }

        _latestVersion++;
        _versionToAddress[_latestVersion] = protocolAddress;
    }

    function _getVersionToAddress(
        uint256 version
    ) internal view virtual returns (address) {
        if (!(0 < version && version < _latestVersion)) {
            revert InvalidProtocolVersion();
        }
        return _versionToAddress[version];
    }
}
