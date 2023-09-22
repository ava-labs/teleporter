//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./IAllowList.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

// AllowList is a base contract to use AllowList precompile capabilities.
contract AllowList is Ownable {
  // Precompiled Allow List Contract Address
  IAllowList private _allowList;

  uint256 constant _STATUS_NONE = 0;
  uint256 constant _STATUS_ENABLED = 1;
  uint256 constant _STATUS_ADMIN = 2;

  enum Role {
    None,
    Enabled,
    Admin
  }

  constructor(address precompileAddr) Ownable() {
    _allowList = IAllowList(precompileAddr);
  }

  modifier onlyEnabled() {
    require(isEnabled(msg.sender), "not enabled");
    _;
  }

  function isAdmin(address addr) public view returns (bool) {
    uint256 result = _allowList.readAllowList(addr);
    return result == _STATUS_ADMIN;
  }

  function isEnabled(address addr) public view returns (bool) {
    uint256 result = _allowList.readAllowList(addr);
    // if address is ENABLED or ADMIN it can deploy
    // in other words, if it's not NONE it can deploy.
    return result != _STATUS_NONE;
  }

  function setAdmin(address addr) public virtual onlyOwner {
    _setAdmin(addr);
  }

  function _setAdmin(address addr) private {
    _allowList.setAdmin(addr);
  }

  function setEnabled(address addr) public virtual onlyOwner {
    _setEnabled(addr);
  }

  function _setEnabled(address addr) private {
    _allowList.setEnabled(addr);
  }

  function revoke(address addr) public virtual onlyOwner {
    _revoke(addr);
  }

  function _revoke(address addr) private {
    require(msg.sender != addr, "cannot revoke own role");
    _allowList.setNone(addr);
  }
}
