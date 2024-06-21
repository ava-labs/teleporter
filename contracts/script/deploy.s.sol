// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "forge-std/Script.sol";
import {ERC20TokenHome} from "../src/TokenHome/ERC20TokenHome.sol";
import {TransparentUpgradeableProxy} from
    "@openzeppelin/contracts@4.8.1/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts@4.8.1/proxy/transparent/ProxyAdmin.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";

contract DeployScript is Script {
    event Deployed(address instance);

    function run() external {
        vm.startBroadcast();
        bytes32 blockchainID =
            IWarpMessenger(0x0200000000000000000000000000000000000005).getBlockchainID();
        console.logBytes32(blockchainID);

        ERC20TokenHome instance = new ERC20TokenHome();
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(instance),
            address(proxyAdmin),
            abi.encodeCall(
                instance.initialize,
                (
                    vm.envAddress("TELEPORTER_REGISTRY_ADDRESS"),
                    vm.envAddress("TELEPORTER_MANAGER"),
                    vm.envAddress("TOKEN_ADDRESS"),
                    uint8(18)
                )
            )
        );
        emit Deployed(address(proxy));

        vm.stopBroadcast();
    }
}
