// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts@4.8.1/security/ReentrancyGuard.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev Message format for the WarpMessage payload to be forwarded to the target contract
 *
 * validatorSetSigAddress: Address of the ValidatorSetSig contract this message is intended for
 * targetContractAddress: Address of the contract that the payload should be forwarded to
 * targetBlockchainID: Blockchain ID of the chain the message is intended for
 * nonce: Unique nonce for the target contract address to provide replay protection
 * payload: Payload to be forwarded to the target contract. Usually ABI encoded function call with parameters.
 */
struct ValidatorSetSigMessage {
    address validatorSetSigAddress;
    address targetContractAddress;
    bytes32 targetBlockchainID;
    uint256 nonce;
    bytes payload;
}

/**
 * @dev Contract that verifies that a set threshold of validators from a given blockchainID
 * have signed an off-chain Warp message and forwards the payload to the target contract specified in the ValidatorSetSigMessage.
 * The threshold itself is set by the validator themselves in their Warp configs:
 * https://github.com/ava-labs/subnet-evm/blob/master/precompile/contracts/warp/config.go#L50
 *
 * This is intended to be used for safe off-chain governance of enabled contracts. An example usecase would be
 * to deploy a `Ownable` target contract that is owned by an instance of this contract and adding the
 * `onlyOwner` modifier to the functions that should be governed.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter/blob/main/SECURITY.md
 */
contract ValidatorSetSig is ReentrancyGuard {
    /**
     * @notice A blockchain ID whose validators need to sign the message for it to be considered valid.
     */
    bytes32 public immutable validatorBlockchainID;

    /**
     * @dev Tracks nonces for messages sent to a specific contract address to provide replay protection.
     * The target contract address is used as the key instead of keeping a global nonce to prevent
     * one party from consuming a global nonce that another was intending to use
     * in case a single instance of ValidatorSetSig contract is used to manage multiple downstream contracts.
     */
    mapping(address targetContractAddress => uint256 nonce) public nonces;

    /**
     * @notice Address that the off-chain Warp message sets as the "source" address.
     * @dev The address is not owned by any EOA or smart contract account, so it
     * cannot possibly be the source address of any other Warp message emitted by the VM.
     */
    address public constant VALIDATORS_SOURCE_ADDRESS = address(0);

    /**
     * @notice The blockchain ID of the chain the contract is deployed on.
     */
    bytes32 public immutable blockchainID;

    /**
     * @notice Warp precompile used for sending and receiving Warp messages.
     */
    IWarpMessenger public constant WARP_MESSENGER =
        IWarpMessenger(0x0200000000000000000000000000000000000005);

    /**
     * @notice Emited when the payload is successfully delivered to the target contract.
     */
    event Delivered(address indexed targetContractAddress, uint256 indexed nonce);

    constructor(bytes32 validatorBlockchainID_) {
        validatorBlockchainID = validatorBlockchainID_;
        blockchainID = WARP_MESSENGER.getBlockchainID();
    }

    function executeCall(uint32 messageIndex) external nonReentrant {
        // Get the WarpMessage from the WarpMessenger pre-compile and verify that it is valid
        (WarpMessage memory message, bool valid) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        require(valid, "ValidatorSetSig: invalid warp message");

        // Ensure that the sourceChainID of the validator set that signed the message is the authorized one
        require(
            message.sourceChainID == validatorBlockchainID, "ValidatorSetSig: invalid sourceChainID"
        );
        // Ensure that the originSenderAddress is the zero address to prevent on-chain messages from being forwarded
        require(
            message.originSenderAddress == VALIDATORS_SOURCE_ADDRESS,
            "ValidatorSetSig: non-zero originSenderAddress"
        );

        ValidatorSetSigMessage memory validatorSetSigMessage =
            abi.decode(message.payload, (ValidatorSetSigMessage));

        validateMessage(validatorSetSigMessage);

        nonces[validatorSetSigMessage.targetContractAddress] = validatorSetSigMessage.nonce;

        // We don't need to protect against return bomb vectors below here since the caller is expected to have full control over the contract called.
        (bool success,) =
        // solhint-disable-next-line avoid-low-level-calls
         validatorSetSigMessage.targetContractAddress.call(validatorSetSigMessage.payload);

        // Use require to revert the transaction if the call fails. This is to prevent consuming the nonce if the call fails due to OOG
        // and requiring re-signing of the message with a new nonce.
        require(success, "ValidatorSetSig: call failed");
        emit Delivered(validatorSetSigMessage.targetContractAddress, validatorSetSigMessage.nonce);
    }

    function validateMessage(ValidatorSetSigMessage memory message) public view {
        require(
            message.validatorSetSigAddress == address(this),
            "ValidatorSetSig: invalid validatorSetSigAddress"
        );
        require(
            message.targetBlockchainID == blockchainID,
            "ValidatorSetSig: invalid targetBlockchainID"
        );
        require(
            nonces[message.targetContractAddress] + 1 == message.nonce,
            "ValidatorSetSig: invalid nonce"
        );
    }
}
