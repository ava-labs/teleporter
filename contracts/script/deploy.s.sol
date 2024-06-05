// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "forge-std/Script.sol";
import "../src/TokenHub/ERC20TokenHub.sol";
import {TransparentUpgradeableProxy} from
    "@openzeppelin/contracts@4.8.1/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts@4.8.1/proxy/transparent/ProxyAdmin.sol";

contract DeployScript is Script {
    event Deployed(address instance);

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        ERC20TokenHub instance = new ERC20TokenHub();
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(instance),
            address(proxyAdmin),
            ""
        );
        emit Deployed(address(proxy));

        vm.stopBroadcast();
    }
}
