//SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import "./IAllowList.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

// AllowList is a base contract to use AllowList precompile capabilities.
contract AllowList is Ownable {
  enum Role {
    None,
    Enabled,
    Admin
  }

  // Precompiled Allow List Contract Address
  IAllowList private _allowList;

  uint256 public constant STATUS_NONE = 0;
  uint256 public constant STATUS_ENABLED = 1;
  uint256 public constant STATUS_ADMIN = 2;

  error NotEnabled();
  error CannotRevokeOwnRule();

  modifier onlyEnabled() {
    if (!isEnabled(msg.sender)) {
      revert NotEnabled();
    }
    _;
  }

  constructor(address precompileAddr) Ownable() {
    _allowList = IAllowList(precompileAddr);
  }

  function isAdmin(address addr) public view returns (bool) {
    uint256 result = _allowList.readAllowList(addr);
    return result == STATUS_ADMIN;
  }

  function isEnabled(address addr) public view returns (bool) {
    uint256 result = _allowList.readAllowList(addr);
    // if address is ENABLED or ADMIN it can deploy
    // in other words, if it's not NONE it can deploy.
    return result != STATUS_NONE;
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
    if (msg.sender == addr) {
      revert CannotRevokeOwnRule();
    }
    _allowList.setNone(addr);
  }
}
