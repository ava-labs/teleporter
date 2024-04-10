// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ERC20BridgeTest} from "./ERC20BridgeTests.t.sol";
import {TeleporterTokenDestinationTest} from "./TeleporterTokenDestinationTests.t.sol";
import {IERC20SendAndCallReceiver} from "../src/interfaces/IERC20SendAndCallReceiver.sol";
import {ERC20Destination} from "../src/ERC20Destination.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";

contract ERC20DestinationTest is ERC20BridgeTest, TeleporterTokenDestinationTest {
    using SafeERC20 for IERC20;

    string public constant MOCK_TOKEN_NAME = "Test Token";
    string public constant MOCK_TOKEN_SYMBOL = "TST";
    uint8 public constant MOCK_TOKEN_DECIMALS = 18;

    ERC20Destination public app;

    function setUp() public virtual override {
        TeleporterTokenDestinationTest.setUp();

        app = new ERC20Destination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: address(this),
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            tokenName: MOCK_TOKEN_NAME,
            tokenSymbol: MOCK_TOKEN_SYMBOL,
            tokenDecimals: MOCK_TOKEN_DECIMALS
        });

        erc20Bridge = app;
        tokenDestination = app;
        tokenBridge = app;
        feeToken = IERC20(app);

        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(0), address(this), 10e18);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            _encodeSingleHopSendMessage(10e18, address(this))
        );
    }

    /**
     * Initialization unit tests
     */
    function testZeroTeleporterRegistryAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");
        new ERC20Destination({
            teleporterRegistryAddress: address(0),
            teleporterManager: address(this),
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            tokenName: MOCK_TOKEN_NAME,
            tokenSymbol: MOCK_TOKEN_SYMBOL,
            tokenDecimals: MOCK_TOKEN_DECIMALS
        });
    }

    function testZeroTeleporterManagerAddress() public {
        vm.expectRevert("Ownable: new owner is the zero address");
        new ERC20Destination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: address(0),
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            tokenName: MOCK_TOKEN_NAME,
            tokenSymbol: MOCK_TOKEN_SYMBOL,
            tokenDecimals: MOCK_TOKEN_DECIMALS
        });
    }

    function testZeroSourceBlockchainID() public {
        vm.expectRevert(_formatErrorMessage("zero source blockchain ID"));
        new ERC20Destination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: address(this),
            sourceBlockchainID: bytes32(0),
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            tokenName: MOCK_TOKEN_NAME,
            tokenSymbol: MOCK_TOKEN_SYMBOL,
            tokenDecimals: MOCK_TOKEN_DECIMALS
        });
    }

    function testDecimals() public {
        uint8 res = app.decimals();
        assertEq(MOCK_TOKEN_DECIMALS, res);
    }

    function testDeployToSameBlockchain() public {
        vm.expectRevert(_formatErrorMessage("cannot deploy to same blockchain as source"));
        new ERC20Destination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: address(this),
            sourceBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            tokenName: MOCK_TOKEN_NAME,
            tokenSymbol: MOCK_TOKEN_SYMBOL,
            tokenDecimals: MOCK_TOKEN_DECIMALS
        });
    }

    function testZeroTokenSourceAddress() public {
        vm.expectRevert(_formatErrorMessage("zero token source address"));
        new ERC20Destination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: address(this),
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: address(0),
            tokenName: MOCK_TOKEN_NAME,
            tokenSymbol: MOCK_TOKEN_SYMBOL,
            tokenDecimals: MOCK_TOKEN_DECIMALS
        });
    }

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal override {
        vm.expectEmit(true, true, true, true, address(tokenDestination));
        emit Transfer(address(0), recipient, amount);
    }

    function _setUpExpectedSendAndCall(
        address recipient,
        uint256 amount,
        bytes memory payload,
        uint256 gasLimit,
        bool targetHasCode,
        bool expectSuccess
    ) internal override {
        // The bridge tokens will be minted to the contract itself
        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(0), address(tokenDestination), amount);

        // Then recipient contract is then approved to spend them
        vm.expectEmit(true, true, true, true, address(app));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

        if (targetHasCode) {
            vm.etch(recipient, new bytes(1));

            bytes memory expectedCalldata = abi.encodeCall(
                IERC20SendAndCallReceiver.receiveTokens, (address(app), amount, payload)
            );
            if (expectSuccess) {
                vm.mockCall(recipient, expectedCalldata, new bytes(0));
            } else {
                vm.mockCallRevert(recipient, expectedCalldata, new bytes(0));
            }
            vm.expectCall(recipient, 0, uint64(gasLimit), expectedCalldata);
        } else {
            vm.etch(recipient, new bytes(0));
        }

        // Then recipient contract approval is reset
        vm.expectEmit(true, true, true, true, address(app));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, 0);

        if (targetHasCode && expectSuccess) {
            // The call should have succeeded.
            vm.expectEmit(true, true, true, true, address(app));
            emit CallSucceeded(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);
        } else {
            // The call should have failed.
            vm.expectEmit(true, true, true, true, address(app));
            emit CallFailed(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

            // Then the amount should be sent to the fallback recipient.
            vm.expectEmit(true, true, true, true, address(app));
            emit Transfer(address(app), address(DEFAULT_FALLBACK_RECIPIENT_ADDRESS), amount);
        }
    }

    function _setUpMockMint(address, uint256) internal pure override {
        // Don't need to mock the minting of an ERC20 destination since it is an internal call
        // on the destination contract.
        return;
    }

    function _scaleTokens(uint256 amount, bool) internal pure override returns (uint256) {
        return amount;
    }
}
