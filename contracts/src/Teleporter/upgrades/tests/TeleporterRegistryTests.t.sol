// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "forge-std/Test.sol";
import "../TeleporterRegistry.sol";
import "../../TeleporterMessenger.sol";

contract TeleporterRegistryTest is Test {
    TeleporterRegistry public teleporterRegistry;
    address public teleporterAddress;

    bytes32 public constant MOCK_BLOCK_CHAIN_ID = bytes32(uint256(123456));
    address public constant WARP_PRECOMPILE_ADDRESS =
        0x0200000000000000000000000000000000000005;

    event AddProtocolVersion(
        uint256 indexed version,
        address indexed protocolAddress
    );

    function setUp() public {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(WarpMessenger.getBlockchainID.selector),
            abi.encode(MOCK_BLOCK_CHAIN_ID)
        );
        teleporterRegistry = new TeleporterRegistry(
            new uint256[](0),
            new address[](0)
        );
        assertEq(0, teleporterRegistry.getLatestVersion());

        teleporterAddress = address(new TeleporterMessenger());
    }

    function testRegistryInitializationMismatchFails() public {
        uint256[] memory initialVersions = new uint256[](1);
        initialVersions[0] = 1;
        address[] memory initialProtocolAddresses = new address[](0);

        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(WarpMessenger.getBlockchainID, ())
        );

        vm.expectRevert(
            WarpProtocolRegistry.InvalidRegistryInitialization.selector
        );
        new TeleporterRegistry(initialVersions, initialProtocolAddresses);
    }

    function testRegistryInitializationDuplicateFails() public {
        uint256[] memory initialVersions = new uint256[](2);
        initialVersions[0] = 1;
        initialVersions[1] = 2;
        address[] memory initialProtocolAddresses = new address[](2);
        initialProtocolAddresses[0] = teleporterAddress;
        initialProtocolAddresses[1] = teleporterAddress;

        vm.expectRevert(TeleporterRegistry.DuplicateProtocolAddress.selector);
        new TeleporterRegistry(initialVersions, initialProtocolAddresses);
    }

    function testAddProtocolVersionSuccess() public {
        uint256 latestVersion = teleporterRegistry.getLatestVersion();
        assertEq(0, latestVersion);

        _addProtocolVersion(teleporterRegistry);
        assertEq(latestVersion + 1, teleporterRegistry.getLatestVersion());
        assertEq(
            teleporterAddress,
            address(teleporterRegistry.getLatestTeleporter())
        );
        assertEq(
            teleporterRegistry.getAddressToVersion(teleporterAddress),
            teleporterRegistry.getLatestVersion()
        );
    }

    function testAddToRegistryFails() public {
        uint256 latestVersion = teleporterRegistry.getLatestVersion();
        uint32 messageIndex = 0;

        // Create a warp out of band message with same version as latest version
        WarpMessage memory warpMessage = _createWarpOutofBandMessage(
            latestVersion,
            teleporterAddress,
            address(teleporterRegistry)
        );

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.getVerifiedWarpMessage,
                (messageIndex)
            ),
            abi.encode(warpMessage, true)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(WarpMessenger.getVerifiedWarpMessage, (messageIndex))
        );

        // Check that adding a protocol version with the same version fails
        vm.expectRevert(WarpProtocolRegistry.InvalidProtocolVersion.selector);
        teleporterRegistry.addProtocolVersion(messageIndex);

        // Check that adding a protocol version with a version that is not the increment of the latest version fails
        warpMessage = _createWarpOutofBandMessage(
            latestVersion + 2,
            teleporterAddress,
            address(teleporterRegistry)
        );

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.getVerifiedWarpMessage,
                (messageIndex)
            ),
            abi.encode(warpMessage, true)
        );

        vm.expectRevert(WarpProtocolRegistry.InvalidProtocolVersion.selector);
        teleporterRegistry.addProtocolVersion(messageIndex);

        // Check that adding a protocol address that is not a contract fails
        warpMessage = _createWarpOutofBandMessage(
            latestVersion + 1,
            address(0),
            address(teleporterRegistry)
        );

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.getVerifiedWarpMessage,
                (messageIndex)
            ),
            abi.encode(warpMessage, true)
        );

        vm.expectRevert(WarpProtocolRegistry.InvalidProtocolAddress.selector);
        teleporterRegistry.addProtocolVersion(messageIndex);
    }

    function testAddToRegistryDuplicateFails() public {
        _addProtocolVersion(teleporterRegistry);
        uint256 latestVersion = teleporterRegistry.getLatestVersion();
        uint32 messageIndex = 0;

        // Check that adding a protocol address that is already registered fails
        WarpMessage memory warpMessage = _createWarpOutofBandMessage(
            latestVersion + 1,
            teleporterAddress,
            address(teleporterRegistry)
        );

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.getVerifiedWarpMessage,
                (messageIndex)
            ),
            abi.encode(warpMessage, true)
        );

        vm.expectRevert(TeleporterRegistry.DuplicateProtocolAddress.selector);
        teleporterRegistry.addProtocolVersion(messageIndex);
    }

    function testGetVersionToAddress() public {
        _addProtocolVersion(teleporterRegistry);
        uint256 latestVersion = teleporterRegistry.getLatestVersion();

        // First test success case
        assertEq(
            teleporterAddress,
            teleporterRegistry.getVersionToAddress(latestVersion)
        );

        // Check that getting version 0 fails
        vm.expectRevert(WarpProtocolRegistry.InvalidProtocolVersion.selector);
        teleporterRegistry.getVersionToAddress(0);

        // Check that getting a version that doesn't exist fails
        vm.expectRevert(WarpProtocolRegistry.InvalidProtocolVersion.selector);
        teleporterRegistry.getVersionToAddress(latestVersion + 1);
    }

    function testGetAddressToVersion() public {
        _addProtocolVersion(teleporterRegistry);
        uint256 latestVersion = teleporterRegistry.getLatestVersion();

        // First test success case
        assertEq(
            latestVersion,
            teleporterRegistry.getAddressToVersion(teleporterAddress)
        );

        // Check that getting a version of an address that doesn't exist returns 0
        assertEq(0, teleporterRegistry.getAddressToVersion(address(this)));
    }

    function testInvalidWarpMessage() public {
        uint256 latestVersion = teleporterRegistry.getLatestVersion();
        uint32 messageIndex = 0;
        WarpMessage memory warpMessage = _createWarpOutofBandMessage(
            latestVersion + 1,
            teleporterAddress,
            address(teleporterRegistry)
        );

        // First check if warp message is invalid from getVerifiedWarpMessage
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.getVerifiedWarpMessage,
                (messageIndex)
            ),
            abi.encode(warpMessage, false)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(WarpMessenger.getVerifiedWarpMessage, (messageIndex))
        );

        vm.expectRevert(WarpProtocolRegistry.InvalidWarpMessage.selector);
        teleporterRegistry.addProtocolVersion(messageIndex);

        // Check if we have an invalid source chain ID
        warpMessage.sourceChainID = bytes32(uint256(1234567));
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.getVerifiedWarpMessage,
                (messageIndex)
            ),
            abi.encode(warpMessage, true)
        );

        vm.expectRevert(WarpProtocolRegistry.InvalidSourceChainID.selector);
        teleporterRegistry.addProtocolVersion(messageIndex);

        // Check if we have an invalid origin sender address
        warpMessage.sourceChainID = MOCK_BLOCK_CHAIN_ID;
        warpMessage.originSenderAddress = address(this);
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.getVerifiedWarpMessage,
                (messageIndex)
            ),
            abi.encode(warpMessage, true)
        );

        vm.expectRevert(
            WarpProtocolRegistry.InvalidOriginSenderAddress.selector
        );
        teleporterRegistry.addProtocolVersion(messageIndex);

        // Check if we have an invalid destination chain ID
        warpMessage.originSenderAddress = teleporterRegistry
            .VALIDATORS_SOURCE_ADDRESS();
        warpMessage.destinationChainID = bytes32(uint256(1234567));
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.getVerifiedWarpMessage,
                (messageIndex)
            ),
            abi.encode(warpMessage, true)
        );

        vm.expectRevert(
            WarpProtocolRegistry.InvalidDestinationChainID.selector
        );
        teleporterRegistry.addProtocolVersion(messageIndex);

        // Check if we have an invalid destination address
        warpMessage.destinationChainID = MOCK_BLOCK_CHAIN_ID;
        warpMessage.destinationAddress = address(this);
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.getVerifiedWarpMessage,
                (messageIndex)
            ),
            abi.encode(warpMessage, true)
        );

        vm.expectRevert(
            WarpProtocolRegistry.InvalidDestinationAddress.selector
        );
        teleporterRegistry.addProtocolVersion(messageIndex);
    }

    function _addProtocolVersion(TeleporterRegistry registry) internal {
        uint256 latestVersion = registry.getLatestVersion();
        uint32 messageIndex = 0;
        WarpMessage memory warpMessage = _createWarpOutofBandMessage(
            latestVersion + 1,
            teleporterAddress,
            address(registry)
        );
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.getVerifiedWarpMessage,
                (messageIndex)
            ),
            abi.encode(warpMessage, true)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(WarpMessenger.getVerifiedWarpMessage, (messageIndex))
        );

        vm.expectEmit(true, true, false, false, address(registry));
        emit AddProtocolVersion(latestVersion + 1, teleporterAddress);
        registry.addProtocolVersion(messageIndex);
    }

    function _createWarpOutofBandMessage(
        uint256 version,
        address protocolAddress,
        address registryAddress
    ) internal view returns (WarpMessage memory) {
        return
            WarpMessage({
                sourceChainID: MOCK_BLOCK_CHAIN_ID,
                originSenderAddress: TeleporterRegistry(registryAddress)
                    .VALIDATORS_SOURCE_ADDRESS(),
                destinationChainID: MOCK_BLOCK_CHAIN_ID,
                destinationAddress: address(registryAddress),
                payload: abi.encode(version, protocolAddress)
            });
    }
}
