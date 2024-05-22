// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ERC20TokenBridgeTest} from "./ERC20TokenBridgeTests.t.sol";
import {TokenHubTest} from "./TokenHubTests.t.sol";
import {IERC20SendAndCallReceiver} from "../src/interfaces/IERC20SendAndCallReceiver.sol";
import {ERC20TokenHub} from "../src/TokenHub/ERC20TokenHub.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {ExampleERC20} from "../lib/teleporter/contracts/src/Mocks/ExampleERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";

contract ERC20TokenHubTest is ERC20TokenBridgeTest, TokenHubTest {
    using SafeERC20 for IERC20;

    ERC20TokenHub public app;
    IERC20 public mockERC20;

    function setUp() public override {
        TokenHubTest.setUp();

        mockERC20 = new ExampleERC20();
        app = new ERC20TokenHub(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            address(mockERC20)
        );
        erc20Bridge = app;
        tokenHub = app;
        tokenBridge = app;

        bridgedToken = mockERC20;
    }

    /**
     * Initialization unit tests
     */
    function testZeroTeleporterRegistryAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");
        new ERC20TokenHub(address(0), address(this), address(mockERC20));
    }

    function testZeroTeleporterManagerAddress() public {
        vm.expectRevert("Ownable: new owner is the zero address");
        new ERC20TokenHub(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(0),
            address(mockERC20)
        );
    }

    function testZeroFeeTokenAddress() public {
        vm.expectRevert(_formatErrorMessage("zero token address"));
        new ERC20TokenHub(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            address(0)
        );
    }

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal override {
        vm.expectEmit(true, true, true, true, address(tokenHub));
        emit TokensWithdrawn(recipient, amount);
        vm.expectCall(address(mockERC20), abi.encodeCall(IERC20.transfer, (address(recipient), amount)));
        vm.expectEmit(true, true, true, true, address(mockERC20));
        emit Transfer(address(app), recipient, amount);
    }

    function _setUpExpectedSendAndCall(
        bytes32 sourceBlockchainID,
        OriginSenderInfo memory originInfo,
        address recipient,
        uint256 amount,
        bytes memory payload,
        uint256 gasLimit,
        bool targetHasCode,
        bool expectSuccess
    ) internal override {
        // The recipient contract will be approved to spend the tokens from the token hub contract.
        vm.expectEmit(true, true, true, true, address(mockERC20));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

        if (targetHasCode) {
            vm.etch(recipient, new bytes(1));

            bytes memory expectedCalldata = abi.encodeCall(
                IERC20SendAndCallReceiver.receiveTokens,
                (
                    sourceBlockchainID,
                    originInfo.bridgeAddress,
                    originInfo.senderAddress,
                    address(mockERC20),
                    amount,
                    payload
                )
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
        vm.expectEmit(true, true, true, true, address(mockERC20));
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
            vm.expectEmit(true, true, true, true, address(mockERC20));
            emit Transfer(address(app), address(DEFAULT_FALLBACK_RECIPIENT_ADDRESS), amount);
        }
    }

    function _addCollateral(bytes32 spokeBlockchainID, address spokeBridgeAddress, uint256 amount) internal override {
        app.addCollateral(spokeBlockchainID, spokeBridgeAddress, amount);
    }

    function _setUpDeposit(uint256 amount) internal virtual override {
        // Increase the allowance of the bridge to transfer the funds from the user
        bridgedToken.safeIncreaseAllowance(address(tokenBridge), amount);
    }

    function _setUpExpectedZeroAmountRevert() internal override {
        vm.expectRevert("SafeERC20TransferFrom: balance not increased");
    }
}
