// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";
import {
    ValidatorSetSig,
    WarpMessage,
    IWarpMessenger,
    ValidatorSetSigMessage
} from "../ValidatorSetSig.sol";

import {UnitTestMockERC20} from "@mocks/UnitTestMockERC20.sol";

contract ValidatorSetSigTest is Test {
    address public constant WARP_PRECOMPILE_ADDRESS =
        address(0x0200000000000000000000000000000000000005);
    bytes32 private constant _MOCK_BLOCKCHAIN_ID = bytes32(uint256(123456));
    bytes32 private constant _MOCK_VALIDATOR_CHAIN_ID = bytes32(uint256(234567));
    address private constant _DEFAULT_TARGET_ADDRESS = 0x1234512345123451234512345123451234512345;
    address private constant _OTHER_TARGET_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    address private constant _ERC20_RECIPIENT_ADDRESS = 0xa4CEE7d1aF6aDdDD33E3b1cC680AB84fdf1b6d1d;

    ValidatorSetSig public validatorSetSig;
    UnitTestMockERC20 public mockERC20;

    function setUp() public virtual {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(_MOCK_BLOCKCHAIN_ID)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector)
        );

        validatorSetSig = new ValidatorSetSig(_MOCK_VALIDATOR_CHAIN_ID);
        assertEq(validatorSetSig.blockchainID(), _MOCK_BLOCKCHAIN_ID);
        assertEq(validatorSetSig.validatorBlockchainID(), _MOCK_VALIDATOR_CHAIN_ID);

        mockERC20 = new UnitTestMockERC20();
        mockERC20.transfer(_ERC20_RECIPIENT_ADDRESS, 10);
        assertEq(mockERC20.balanceOf(_ERC20_RECIPIENT_ADDRESS), 10);
    }

    function testInvalidMessage() public {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(
                WarpMessage({
                    sourceChainID: _MOCK_BLOCKCHAIN_ID,
                    originSenderAddress: address(0),
                    payload: ""
                }),
                false
            )
        );
        vm.expectRevert("ValidatorSetSig: invalid warp message");
        validatorSetSig.executeCall(0);
    }

    function testInvalidSourceChainID() public {
        WarpMessage memory invalidSourceChainMessage = _getValidWarpMessage();
        invalidSourceChainMessage.sourceChainID = _MOCK_BLOCKCHAIN_ID;
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(invalidSourceChainMessage, true)
        );
        vm.expectRevert("ValidatorSetSig: invalid sourceChainID");
        validatorSetSig.executeCall(0);
    }

    function testNonZeroSourceAddress() public {
        WarpMessage memory nonZeroSourceAddressMessage = _getValidWarpMessage();
        nonZeroSourceAddressMessage.originSenderAddress = address(1);
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(nonZeroSourceAddressMessage, true)
        );
        vm.expectRevert("ValidatorSetSig: non-zero originSenderAddress");
        validatorSetSig.executeCall(0);
    }

    function testInvalidValidatorSetSigContractAddress() public {
        WarpMessage memory invalidValidatorSetSigContractAddressMessage = _getValidWarpMessage();
        ValidatorSetSigMessage memory validatorSetSigMessage = _getValidValidatorSetSigMessage();
        validatorSetSigMessage.validatorSetSigAddress = address(1);
        invalidValidatorSetSigContractAddressMessage.payload = abi.encode(validatorSetSigMessage);
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(invalidValidatorSetSigContractAddressMessage, true)
        );
        vm.expectRevert("ValidatorSetSig: invalid validatorSetSigAddress");
        validatorSetSig.executeCall(0);
    }

    function testInvalidTargetBlockChainID() public {
        WarpMessage memory invalidTargetBlockChainIDMessage = _getValidWarpMessage();
        ValidatorSetSigMessage memory validatorSetSigMessage = _getValidValidatorSetSigMessage();
        validatorSetSigMessage.targetBlockchainID = _MOCK_VALIDATOR_CHAIN_ID;
        invalidTargetBlockChainIDMessage.payload = abi.encode(validatorSetSigMessage);
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(invalidTargetBlockChainIDMessage, true)
        );
        vm.expectRevert("ValidatorSetSig: invalid targetBlockchainID");
        validatorSetSig.executeCall(0);
    }

    function testNonces() public {
        WarpMessage memory warpMessage = _getValidWarpMessage();
        ValidatorSetSigMessage memory validatorSetSigMessage = _getValidValidatorSetSigMessage();
        // Set nonce to invalid value, since first message should have nonce 0
        validatorSetSigMessage.nonce = 1;
        warpMessage.payload = abi.encode(validatorSetSigMessage);
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(warpMessage, true)
        );
        vm.expectRevert("ValidatorSetSig: invalid nonce");
        validatorSetSig.executeCall(0);
        // Set nonce to 2 which is still invalid and verify that it still fails
        validatorSetSigMessage.nonce = 2;
        warpMessage.payload = abi.encode(validatorSetSigMessage);
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(warpMessage, true)
        );
        vm.expectRevert("ValidatorSetSig: invalid nonce");
        validatorSetSig.executeCall(0);

        // Set nonce to 0 which is valid
        validatorSetSigMessage.nonce = 0;
        warpMessage.payload = abi.encode(validatorSetSigMessage);
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(warpMessage, true)
        );
        assertEq(validatorSetSig.nonces(_DEFAULT_TARGET_ADDRESS), 0);
        validatorSetSig.executeCall(0);
        assertEq(validatorSetSig.nonces(_DEFAULT_TARGET_ADDRESS), 1);

        // Set nonce to 1 which is valid
        validatorSetSigMessage.nonce = 1;
        warpMessage.payload = abi.encode(validatorSetSigMessage);
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(warpMessage, true)
        );
        validatorSetSig.executeCall(0);
        assertEq(validatorSetSig.nonces(_DEFAULT_TARGET_ADDRESS), 2);

        // finally test different target address
        validatorSetSigMessage.targetContractAddress = _OTHER_TARGET_ADDRESS;
        validatorSetSigMessage.nonce = 0;
        warpMessage.payload = abi.encode(validatorSetSigMessage);
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(warpMessage, true)
        );
        validatorSetSig.executeCall(0);
        assertEq(validatorSetSig.nonces(_DEFAULT_TARGET_ADDRESS), 2);
        assertEq(validatorSetSig.nonces(_OTHER_TARGET_ADDRESS), 1);
    }

    function testReentrancy() public {
        WarpMessage memory warpMessage = _getValidWarpMessage();
        ValidatorSetSigMessage memory validatorSetSigMessage = _getValidValidatorSetSigMessage();
        // Set the targetContractAddress and the txPayload to the ValidatorSetSig contract itself
        validatorSetSigMessage.payload =
            abi.encodeWithSelector(validatorSetSig.executeCall.selector, uint32(0));
        validatorSetSigMessage.targetContractAddress = address(validatorSetSig);
        warpMessage.payload = abi.encode(validatorSetSigMessage);
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(warpMessage, true)
        );
        // Checking for the outermost revert message since the reentrancy guard happens inside the call
        vm.expectRevert("ValidatorSetSig: call failed");
        validatorSetSig.executeCall(0);
    }

    // Uses the validatorSetSig contract to call the mockERC20 contract
    // and verifies that the balance got updated correctly
    function testSuccess() public {
        WarpMessage memory warpMessage = _getValidWarpMessage();
        ValidatorSetSigMessage memory validatorSetSigMessage = _getValidValidatorSetSigMessage();
        // Set the targetContractAddress and the txPayload to the mockERC20 contract
        validatorSetSigMessage.targetContractAddress = address(mockERC20);
        validatorSetSigMessage.payload =
            abi.encodeWithSelector(mockERC20.transfer.selector, _ERC20_RECIPIENT_ADDRESS, 5);
        warpMessage.payload = abi.encode(validatorSetSigMessage);
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(warpMessage, true)
        );
        assertEq(mockERC20.balanceOf(_ERC20_RECIPIENT_ADDRESS), 10);
        validatorSetSig.executeCall(0);
        assertEq(mockERC20.balanceOf(_ERC20_RECIPIENT_ADDRESS), 15);
    }

    function _getValidWarpMessage() private view returns (WarpMessage memory) {
        return WarpMessage({
            sourceChainID: _MOCK_VALIDATOR_CHAIN_ID,
            originSenderAddress: address(0),
            payload: abi.encode(_getValidValidatorSetSigMessage())
        });
    }

    function _getValidValidatorSetSigMessage()
        private
        view
        returns (ValidatorSetSigMessage memory)
    {
        return ValidatorSetSigMessage({
            targetBlockchainID: _MOCK_BLOCKCHAIN_ID,
            validatorSetSigAddress: address(validatorSetSig),
            targetContractAddress: _DEFAULT_TARGET_ADDRESS,
            nonce: 0,
            payload: ""
        });
    }
}
