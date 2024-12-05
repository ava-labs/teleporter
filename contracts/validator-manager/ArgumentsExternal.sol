// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {InitializeEndValidationArgs, InitializeValidatorRegistrationArgs} from "./Arguments.sol";

contract ArgumentsExternal {
    function decodeInitializeEndValidationArgs(bytes calldata args) external pure returns (InitializeEndValidationArgs memory) {
        return abi.decode(args, (InitializeEndValidationArgs));
    }

    function decodeInitializeValidatorRegistrationArgs(bytes calldata args) external pure returns (InitializeValidatorRegistrationArgs memory) {
        return abi.decode(args, (InitializeValidatorRegistrationArgs));
    }
}