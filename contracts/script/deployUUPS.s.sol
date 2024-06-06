// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "forge-std/Script.sol";
import "../src/TokenHub/ERC20TokenHub.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts@4.8.1/proxy/ERC1967/ERC1967Proxy.sol";

contract DeployUUPS is Script {
    event Deployed(address instance);

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        ERC20TokenHub instance = new ERC20TokenHub();
        ERC1967Proxy proxy = new ERC1967Proxy(address(instance), "");
        emit Deployed(address(proxy));

        vm.stopBroadcast();
    }
}
