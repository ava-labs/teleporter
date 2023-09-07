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
    bytes32 public constant VALIDATORS_SOURCE_ADDRESS =
        0x0000000000000000000000000000000000000000000000000000000000000000;

    WarpMessenger public constant WARP_MESSENGER =
        WarpMessenger(0x0200000000000000000000000000000000000005);

    bytes32 private immutable _chainID;

    uint256 private _latestVersion;

    mapping(uint256 version => address) private _protocolAddresses;

    constructor() {
        _latestVersion = 0;
        _chainID = WARP_MESSENGER.getBlockchainID();
    }

    function addProtocolVersion() external virtual {
        (WarpMessage memory message, bool exists) = WARP_MESSENGER
            .getVerifiedWarpMessage();
        require(exists, "WarpProtocolRegistry: no warp message");
        require(
            message.originChainID == _chainID,
            "WarpProtocolRegistry: invalid origin chain ID"
        );
        require(
            message.originSenderAddress == VALIDATORS_SOURCE_ADDRESS,
            "WarpProtocolRegistry: invalid origin sender address"
        );
        require(
            message.destinationChainID == _chainID,
            "WarpProtocolRegistry: invalid destination chain ID"
        );
        require(
            message.destinationAddress == address(this),
            "WarpProtocolRegistry: invalid destination address"
        );

        (uint256 version, address protocolAddress) = abi.decode(
            message.payload,
            (uint256, address)
        );

        require(
            version == _latestVersion,
            "WarpProtocolRegistry: invalid nonce"
        );
        _latestVersion++;

        require(
            Address.isContract(protocolAddress),
            "WarpProtocolRegistry: not a contract"
        );

        _protocolAddresses[_latestVersion] = protocolAddress;
    }

    function getProtocolAddress(
        uint256 version
    ) external view virtual returns (address) {
        return _getProtocolAddress(version);
    }

    function _getProtocolAddress(
        uint256 version
    ) internal view returns (address) {
        require(
            0 < version <= _latestVersion,
            "WarpProtocolRegistry: invalid version"
        );
        return _protocolAddresses[version];
    }
}
