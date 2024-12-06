// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import "./Foo.sol";
import {Test} from "@forge-std/Test.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts@5.0.2/proxy/transparent/TransparentUpgradeableProxy.sol";

contract TestFooBarUpgradeable is Test {
    Foo public foo;

    function setUp() public {
        // Deploy Foo
        Foo impl = new Foo();

        // Deploy a transparent upgradeable proxy
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(address(impl), address(this), "");

        foo = Foo(address(proxy));
        foo.initialize();
    }

    // Test: Setting and Getting varFoo
    function testSetAndGetFoo() public {
        uint256 expectedValue = 42;

        // Call setFoo to set varFoo
        foo.setFoo(expectedValue);

        // Verify that varFoo was updated
        uint256 actualValue = foo.getFoo();
        assert(actualValue == expectedValue);
    }

    // Test: Setting and Getting varBar via barStorage
    function testSetAndGetBar() public {
        uint256 expectedValue = 99;

        // Call setVarB to set varBar in barStorage
        foo.setBar(expectedValue);

        // Verify that varBar in barStorage was updated
        uint256 actualValue = foo.getBar();
        assert(actualValue == expectedValue);
    }
}