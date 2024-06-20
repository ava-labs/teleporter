// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ERC20TokenBridgeTest} from "./ERC20TokenBridgeTests.t.sol";
import {TokenHomeTest} from "./TokenHomeTests.t.sol";
import {IERC20SendAndCallReceiver} from "../src/interfaces/IERC20SendAndCallReceiver.sol";
import {SendTokensInput} from "../src/interfaces/ITokenBridge.sol";
import {ERC20TokenHome} from "../src/TokenHome/ERC20TokenHome.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {ExampleERC20} from "../lib/teleporter/contracts/src/Mocks/ExampleERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";
import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TokenScalingUtils} from "../src/utils/TokenScalingUtils.sol";

contract ERC20TokenHomeTest is ERC20TokenBridgeTest, TokenHomeTest {
    using SafeERC20 for IERC20;

    ERC20TokenHome public app;
    IERC20 public mockERC20;

    function setUp() public override {
        TokenHomeTest.setUp();

        mockERC20 = new ExampleERC20();
        tokenHomeDecimals = 6;
        app = new ERC20TokenHome(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            address(mockERC20),
            tokenHomeDecimals
        );
        erc20Bridge = app;
        tokenHome = app;
        tokenBridge = app;

        bridgedToken = mockERC20;
    }

    /**
     * Initialization unit tests
     */
    function testZeroTeleporterRegistryAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");
        new ERC20TokenHome(address(0), address(this), address(mockERC20), tokenHomeDecimals);
    }

    function testZeroTeleporterManagerAddress() public {
        vm.expectRevert("Ownable: new owner is the zero address");
        new ERC20TokenHome(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(0),
            address(mockERC20),
            tokenHomeDecimals
        );
    }

    function testZeroFeeTokenAddress() public {
        vm.expectRevert(_formatErrorMessage("zero token address"));
        new ERC20TokenHome(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            address(0),
            tokenHomeDecimals
        );
    }

    function testTokenDecimalsTooHigh() public {
        vm.expectRevert(_formatErrorMessage("token decimals too high"));
        new ERC20TokenHome(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            address(mockERC20),
            uint8(TokenScalingUtils.MAX_TOKEN_DECIMALS) + 1
        );
    }

    function testReceiveZeroHomeTokenAmount() public {
        // Set up a registered remote that will scale down the received amount
        // to zero home tokens.
        uint256 tokenMultiplier = 100_000;
        _setUpRegisteredRemote(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, 0, tokenMultiplier, true);

        // Send over home token to the remote
        // and check for expected calls for scaled amount of tokens sent.
        SendTokensInput memory input = _createDefaultSendTokensInput();
        uint256 amount = 1;
        _setUpExpectedDeposit(amount, input.primaryFee);

        uint256 scaledAmount = tokenMultiplier * amount;
        _checkExpectedTeleporterCallsForSend(_createSingleHopTeleporterMessageInput(input, scaledAmount));
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, scaledAmount);
        _send(input, amount);

        // Receive an amount from remote less than `scaledAmount`
        // which will be scaled down to zero home tokens.
        vm.expectRevert(_formatErrorMessage("zero token amount"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            _encodeSingleHopSendMessage(scaledAmount - 1, DEFAULT_RECIPIENT_ADDRESS)
        );
    }

    function testRegisterDestinationRoundUpCollateralNeeded() public {
        _setUpRegisteredRemote(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, 11, 10, true);
        (, uint256 collateralNeeded,,) =
            tokenHome.registeredRemotes(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS);
        assertEq(collateralNeeded, 2);
    }

    function testSendScaledUpAmount() public {
        uint256 amount = 100;
        uint256 feeAmount = 1;

        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = feeAmount;

        // Raw amount sent over wire should be multipled by 1e2.
        uint256 tokenMultiplier = 1e2;
        _setUpRegisteredRemote(input.destinationBlockchainID, input.destinationBridgeAddress, 0, tokenMultiplier, true);
        _setUpExpectedDeposit(amount, input.primaryFee);
        TeleporterMessageInput memory expectedMessage = TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationBridgeAddress,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(bridgedToken), amount: input.primaryFee}),
            requiredGasLimit: input.requiredGasLimit,
            allowedRelayerAddresses: new address[](0),
            message: _encodeSingleHopSendMessage(amount * tokenMultiplier, DEFAULT_RECIPIENT_ADDRESS)
        });
        _checkExpectedTeleporterCallsForSend(expectedMessage);
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, amount * tokenMultiplier);
        _send(input, amount);
    }

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal override {
        vm.expectEmit(true, true, true, true, address(tokenHome));
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
        // The recipient contract will be approved to spend the tokens from the token home contract.
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

    function _addCollateral(bytes32 remoteBlockchainID, address remoteBridgeAddress, uint256 amount)
        internal
        override
    {
        app.addCollateral(remoteBlockchainID, remoteBridgeAddress, amount);
    }

    function _setUpDeposit(uint256 amount) internal virtual override {
        // Increase the allowance of the bridge to transfer the funds from the user
        bridgedToken.safeIncreaseAllowance(address(tokenBridge), amount);
    }

    function _setUpExpectedZeroAmountRevert() internal override {
        vm.expectRevert("SafeERC20TransferFrom: balance not increased");
    }

    function _setUpExpectedDeposit(uint256 amount, uint256 feeAmount) internal virtual override {
        // Transfer the fee to the bridge if it is greater than 0
        if (feeAmount > 0) {
            bridgedToken.safeIncreaseAllowance(address(tokenBridge), feeAmount);
            vm.expectCall(
                address(bridgedToken),
                abi.encodeCall(IERC20.transferFrom, (address(this), address(tokenBridge), feeAmount))
            );
        }
        // Increase the allowance of the bridge to transfer the funds from the user
        bridgedToken.safeIncreaseAllowance(address(tokenBridge), amount);

        // Check that transferFrom is called to deposit the funds sent from the user to the bridge
        // This is the case because for the {ERC20TokenHome) is not the fee token itself
        vm.expectCall(
            address(bridgedToken), abi.encodeCall(IERC20.transferFrom, (address(this), address(tokenBridge), amount))
        );
        vm.expectEmit(true, true, true, true, address(bridgedToken));
        emit Transfer(address(this), address(tokenBridge), amount);
    }
}
