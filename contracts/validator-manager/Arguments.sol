// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

struct InitializeValidatorRegistrationArgs {
    uint16 delegationFeeBips;
    uint64 minStakeDuration;
}

struct InitializeEndValidationArgs {
        bool force;
        bool includeUptimeProof;
        uint32 messageIndex;
        address rewardRecipient;
}

library Arguments {
    function decodeInitializeValidatorRegistrationArgs(bytes calldata args) internal pure returns (InitializeValidatorRegistrationArgs memory) {
        return abi.decode(args, (InitializeValidatorRegistrationArgs));
    }

    function decodeInitializeEndValidationArgs(bytes calldata args) internal pure returns (InitializeEndValidationArgs memory) {
        return abi.decode(args, (InitializeEndValidationArgs));
    }
}