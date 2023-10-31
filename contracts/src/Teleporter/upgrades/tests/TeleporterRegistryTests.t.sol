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

    function setUp() public virtual {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(WarpMessenger.getBlockchainID.selector),
            abi.encode(MOCK_BLOCK_CHAIN_ID)
        );
        teleporterRegistry = new TeleporterRegistry(
            new ProtocolRegistryEntry[](0)
        );
        assertEq(0, teleporterRegistry.latestVersion());

        teleporterAddress = address(new TeleporterMessenger());
    }

    function testAddProtocolVersionBasic() public {
        uint256 latestVersion = teleporterRegistry.latestVersion();
        uint32 messageIndex = 0;

        _addProtocolVersion(teleporterRegistry, teleporterAddress);
        assertEq(latestVersion + 1, teleporterRegistry.latestVersion());
        assertEq(
            teleporterAddress,
            address(teleporterRegistry.getLatestTeleporter())
        );
        assertEq(
            teleporterRegistry.getVersionFromAddress(teleporterAddress),
            teleporterRegistry.latestVersion()
        );

        // Check that adding a protocol version with a version that is not the increment of the latest version succeeds
        latestVersion = teleporterRegistry.latestVersion();
        WarpMessage memory warpMessage = _createWarpOutofBandMessage(
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

        teleporterRegistry.addProtocolVersion(messageIndex);
        assertEq(latestVersion + 2, teleporterRegistry.latestVersion());
        assertEq(
            teleporterAddress,
            address(teleporterRegistry.getLatestTeleporter())
        );
    }

    function testAddNonContractAddress() public {
        // Check that adding a protocol version with a protocol address that is not a contract address succeeds
        uint256 latestVersion = teleporterRegistry.latestVersion();
        uint32 messageIndex = 0;
        WarpMessage memory warpMessage = _createWarpOutofBandMessage(
            latestVersion + 2,
            address(this),
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

        teleporterRegistry.addProtocolVersion(messageIndex);
        assertEq(latestVersion + 2, teleporterRegistry.latestVersion());
        assertEq(
            address(this),
            teleporterRegistry.getAddressFromVersion(latestVersion + 2)
        );
    }

    function testAddOldVersion() public {
        // Check that adding a protocol version that has not been registered, but is less than the latest version succeeds
        // First add to latest version by skipping a version.
        uint256 latestVersion = teleporterRegistry.latestVersion();
        uint32 messageIndex = 0;
        WarpMessage memory warpMessage = _createWarpOutofBandMessage(
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

        teleporterRegistry.addProtocolVersion(messageIndex);
        assertEq(latestVersion + 2, teleporterRegistry.latestVersion());
        assertEq(
            teleporterAddress,
            teleporterRegistry.getAddressFromVersion(latestVersion + 2)
        );

        // latestVersion + 1 was skipped in previous check, is not registered, and is less than latestVersion()
        uint256 oldVersion = latestVersion + 1;
        warpMessage = _createWarpOutofBandMessage(
            oldVersion,
            address(this),
            address(teleporterRegistry)
        );

        // Make sure that oldVersion is not registered, and is less than latestVersion()
        assertEq(oldVersion, teleporterRegistry.latestVersion() - 1);
        assertEq(
            address(0),
            teleporterRegistry.getAddressFromVersion(oldVersion)
        );

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.getVerifiedWarpMessage,
                (messageIndex)
            ),
            abi.encode(warpMessage, true)
        );

        teleporterRegistry.addProtocolVersion(messageIndex);
        assertEq(
            address(this),
            teleporterRegistry.getAddressFromVersion(oldVersion)
        );
        assertEq(oldVersion + 1, teleporterRegistry.latestVersion());
    }

    function testAddExistingVersion() public {
        uint256 latestVersion = teleporterRegistry.latestVersion();
        uint32 messageIndex = 0;

        // Add a new version to the registiry
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
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(WarpMessenger.getVerifiedWarpMessage, (messageIndex))
        );

        teleporterRegistry.addProtocolVersion(messageIndex);
        assertEq(latestVersion + 1, teleporterRegistry.latestVersion());
        assertEq(
            teleporterAddress,
            teleporterRegistry.getAddressFromVersion(latestVersion + 1)
        );

        // Check that adding a protocol version with the same version fails
        vm.expectRevert(_formatRegistryErrorMessage("version already exists"));
        teleporterRegistry.addProtocolVersion(messageIndex);
    }

    function testAddZeroProtocolAddress() public {
        uint256 latestVersion = teleporterRegistry.latestVersion();
        uint32 messageIndex = 0;

        // Check that adding an invalid protocol address of address(0) fails
        WarpMessage memory warpMessage = _createWarpOutofBandMessage(
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

        vm.expectRevert(_formatRegistryErrorMessage("zero protocol address"));
        teleporterRegistry.addProtocolVersion(messageIndex);
    }

    function testAddZeroVersion() public {
        uint32 messageIndex = 0;

        // Check that adding an invalid version of 0 fails
        WarpMessage memory warpMessage = _createWarpOutofBandMessage(
            0,
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

        vm.expectRevert(_formatRegistryErrorMessage("zero version"));
        teleporterRegistry.addProtocolVersion(messageIndex);
    }

    function testGetAddressFromVersion() public {
        _addProtocolVersion(teleporterRegistry, teleporterAddress);
        uint256 latestVersion = teleporterRegistry.latestVersion();

        // First test success case
        assertEq(
            teleporterAddress,
            teleporterRegistry.getAddressFromVersion(latestVersion)
        );

        // Check that getting version 0 fails
        vm.expectRevert(_formatRegistryErrorMessage("zero version"));
        teleporterRegistry.getAddressFromVersion(0);

        // Check that getting a version that doesn't exist fails
        vm.expectRevert(_formatRegistryErrorMessage("invalid version"));
        teleporterRegistry.getAddressFromVersion(latestVersion + 1);
    }

    function testGetVersionFromAddress() public {
        _addProtocolVersion(teleporterRegistry, teleporterAddress);
        uint256 latestVersion = teleporterRegistry.latestVersion();

        // First test success case
        assertEq(
            latestVersion,
            teleporterRegistry.getVersionFromAddress(teleporterAddress)
        );

        // Check that getting a version of an address that doesn't exist returns 0
        assertEq(0, teleporterRegistry.getVersionFromAddress(address(this)));
    }

    function testInvalidWarpMessage() public {
        uint256 latestVersion = teleporterRegistry.latestVersion();
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

        vm.expectRevert(_formatRegistryErrorMessage("invalid warp message"));
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

        vm.expectRevert(_formatRegistryErrorMessage("invalid source chain ID"));
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
            _formatRegistryErrorMessage("invalid origin sender address")
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
            _formatRegistryErrorMessage("invalid destination chain ID")
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
            _formatRegistryErrorMessage("invalid destination address")
        );
        teleporterRegistry.addProtocolVersion(messageIndex);
    }

    function _addProtocolVersion(
        TeleporterRegistry registry,
        address newProtocolAddress
    ) internal {
        uint256 latestVersion = registry.latestVersion();
        uint32 messageIndex = 0;
        WarpMessage memory warpMessage = _createWarpOutofBandMessage(
            latestVersion + 1,
            newProtocolAddress,
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
        emit AddProtocolVersion(latestVersion + 1, newProtocolAddress);
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
                payload: abi.encode(
                    ProtocolRegistryEntry({
                        version: version,
                        protocolAddress: protocolAddress
                    })
                )
            });
    }

    function _formatRegistryErrorMessage(
        string memory errorMessage
    ) private pure returns (bytes memory) {
        return bytes(string.concat("TeleporterRegistry: ", errorMessage));
    }
}
