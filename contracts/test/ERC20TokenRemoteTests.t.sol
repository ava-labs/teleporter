// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.20;

import {ERC20TokenTransferrerTest} from "./ERC20TokenTransferrerTests.t.sol";
import {TokenRemoteTest} from "./TokenRemoteTests.t.sol";
import {IERC20SendAndCallReceiver} from "../src/interfaces/IERC20SendAndCallReceiver.sol";
import {TokenRemote} from "../src/TokenRemote/TokenRemote.sol";
import {TokenRemoteSettings} from "../src/TokenRemote/interfaces/ITokenRemote.sol";
import {ERC20TokenRemote} from "../src/TokenRemote/ERC20TokenRemote.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {ExampleERC20} from "../lib/teleporter/contracts/src/Mocks/ExampleERC20.sol";
import {SendTokensInput} from "../src/interfaces/ITokenTransferrer.sol";
import {Ownable} from "@openzeppelin/contracts@5.0.2/access/Ownable.sol";

contract ERC20TokenRemoteTest is ERC20TokenTransferrerTest, TokenRemoteTest {
    using SafeERC20 for IERC20;

    string public constant MOCK_TOKEN_NAME = "Test Token";
    string public constant MOCK_TOKEN_SYMBOL = "TST";

    ERC20TokenRemote public app;

    function setUp() public virtual override {
        TokenRemoteTest.setUp();

        tokenDecimals = 14;
        tokenHomeDecimals = 18;
        app = ERC20TokenRemote(address(_createNewRemoteInstance()));

        erc20TokenTransferrer = app;
        tokenRemote = app;
        tokenTransferrer = app;

        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(0), address(this), 10e18);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_HOME_ADDRESS,
            _encodeSingleHopSendMessage(10e18, address(this))
        );
    }

    /**
     * Initialization unit tests
     */
    function testZeroTeleporterRegistryAddress() public {
        _invalidInitialization(
            TokenRemoteSettings({
                teleporterRegistryAddress: address(0),
                teleporterManager: address(this),
                tokenHomeBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals,
            "TeleporterUpgradeable: zero teleporter registry address"
        );
    }

    function testZeroTeleporterManagerAddress() public {
        _invalidInitialization(
            TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(0),
                tokenHomeBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals,
            abi.encodeWithSelector(Ownable.OwnableInvalidOwner.selector, address(0))
        );
    }

    function testZeroTokenHomeBlockchainID() public {
        _invalidInitialization(
            TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                tokenHomeBlockchainID: bytes32(0),
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals,
            _formatErrorMessage("zero token home blockchain ID")
        );
    }

    function testDeployToSameBlockchain() public {
        _invalidInitialization(
            TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                tokenHomeBlockchainID: DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals,
            _formatErrorMessage("cannot deploy to same blockchain as token home")
        );
    }

    function testZeroTokenHomeAddress() public {
        _invalidInitialization(
            TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                tokenHomeBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: address(0),
                tokenHomeDecimals: 18
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            18,
            _formatErrorMessage("zero token home address")
        );
    }

    function testDecimals() public {
        uint8 res = app.decimals();
        assertEq(tokenDecimals, res);
    }

    function testSendWithSeparateFeeAsset() public {
        uint256 amount = 200_000;
        uint256 feeAmount = 100;
        ExampleERC20 separateFeeAsset = new ExampleERC20();
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFeeTokenAddress = address(separateFeeAsset);
        input.primaryFee = feeAmount;

        IERC20(separateFeeAsset).safeIncreaseAllowance(address(tokenTransferrer), feeAmount);
        vm.expectCall(
            address(separateFeeAsset),
            abi.encodeCall(
                IERC20.transferFrom, (address(this), address(tokenTransferrer), feeAmount)
            )
        );
        // Increase the allowance of the token transferrer to transfer the funds from the user
        IERC20(app).safeIncreaseAllowance(address(tokenTransferrer), amount);

        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(this), address(0), amount);
        _checkExpectedTeleporterCallsForSend(_createSingleHopTeleporterMessageInput(input, amount));
        vm.expectEmit(true, true, true, true, address(tokenTransferrer));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, amount);
        _send(input, amount);
    }

    function _createNewRemoteInstance() internal override returns (TokenRemote) {
        ERC20TokenRemote instance = new ERC20TokenRemote();
        instance.initialize(
            TokenRemoteSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                tokenHomeBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: DEFAULT_TOKEN_HOME_ADDRESS,
                tokenHomeDecimals: tokenHomeDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals
        );
        return instance;
    }

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal override {
        vm.expectEmit(true, true, true, true, address(tokenRemote));
        emit TokensWithdrawn(recipient, amount);
        vm.expectEmit(true, true, true, true, address(tokenRemote));
        emit Transfer(address(0), recipient, amount);
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
        // The transferred tokens will be minted to the contract itself
        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(0), address(tokenRemote), amount);

        // Then recipient contract is then approved to spend them
        vm.expectEmit(true, true, true, true, address(app));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

        if (targetHasCode) {
            vm.etch(recipient, new bytes(1));

            bytes memory expectedCalldata = abi.encodeCall(
                IERC20SendAndCallReceiver.receiveTokens,
                (
                    sourceBlockchainID,
                    originInfo.tokenTransferrerAddress,
                    originInfo.senderAddress,
                    address(app),
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

    function _setUpExpectedZeroAmountRevert() internal override {
        vm.expectRevert(_formatErrorMessage("insufficient tokens to transfer"));
    }

    function _setUpExpectedDeposit(uint256 amount, uint256 feeAmount) internal virtual override {
        // Transfer the fee to the token transferrer if it is greater than 0
        if (feeAmount > 0) {
            IERC20(app).safeIncreaseAllowance(address(tokenTransferrer), feeAmount);
        }

        // Increase the allowance of the token transferrer to transfer the funds from the user
        IERC20(app).safeIncreaseAllowance(address(tokenTransferrer), amount);

        uint256 currentAllowance = app.allowance(address(this), address(tokenTransferrer));
        if (feeAmount > 0) {
            vm.expectEmit(true, true, true, true, address(app));
            emit Approval(address(this), address(tokenTransferrer), currentAllowance - feeAmount);
            vm.expectEmit(true, true, true, true, address(app));
            emit Transfer(address(this), address(tokenTransferrer), feeAmount);
        }
        vm.expectEmit(true, true, true, true, address(app));
        emit Approval(
            address(this), address(tokenTransferrer), currentAllowance - feeAmount - amount
        );
        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(this), address(0), amount);
    }

    function _getTotalSupply() internal view override returns (uint256) {
        return app.totalSupply();
    }

    function _setUpMockMint(address, uint256) internal pure override {
        // Don't need to mock the minting of an ERC20TokenRemote since it is an internal call
        // on the remote contract.
        return;
    }

    function _invalidInitialization(
        TokenRemoteSettings memory settings,
        string memory tokenName,
        string memory tokenSymbol,
        uint8 tokenDecimals_,
        bytes memory expectedErrorMessage
    ) private {
        app = new ERC20TokenRemote();
        vm.expectRevert(expectedErrorMessage);
        app.initialize(settings, tokenName, tokenSymbol, tokenDecimals_);
    }
}
