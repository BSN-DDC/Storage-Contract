// SPDX-License-Identifier: MIT
pragma solidity ^0.8.2;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";

contract StoragePlus is
    Initializable,
    UUPSUpgradeable,
    OwnableUpgradeable,
    PausableUpgradeable
{
    event StoreData(string key, string value);
    event DeleteData(string key);
    mapping(string => string) private _keyValues;

    function initialize() public initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();
        __Pausable_init();
    }

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() initializer {}

    function _authorizeUpgrade(address) internal override onlyOwner {}

    function storeData(
        string memory key,
        string memory value
    ) public onlyOwner whenNotPaused {
        _keyValues[key] = value;
        emit StoreData(key, value);
    }

    function queryData(string memory key) public view whenNotPaused returns (string memory) {
        return _keyValues[key];
    }

    function deleteData(string memory key) public onlyOwner whenNotPaused {
        delete _keyValues[key];
        emit DeleteData(key);
    }

    /**
     * @dev All methods of suspending the contract.
     */
    function pause() public onlyOwner {
        _pause();
    }

    /**
     * @dev All methods of starting the contract.
     */
    function unpause() public onlyOwner {
        _unpause();
    }
}
