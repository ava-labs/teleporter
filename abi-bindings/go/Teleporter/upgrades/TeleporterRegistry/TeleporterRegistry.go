// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teleporterregistry

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ProtocolRegistryEntry is an auto generated low-level Go binding around an user-defined struct.
type ProtocolRegistryEntry struct {
	Version            *big.Int
	DestinationAddress common.Address
	ProtocolAddress    common.Address
}

// TeleporterRegistryMetaData contains all meta data concerning the TeleporterRegistry contract.
var TeleporterRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"internalType\":\"structProtocolRegistryEntry[]\",\"name\":\"initialEntries\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"name\":\"AddProtocolVersion\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VALIDATORS_SOURCE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"addProtocolVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getAddressFromVersion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestTeleporter\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getTeleporterFromVersion\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"name\":\"getVersionFromAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TeleporterRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleporterRegistryMetaData.ABI instead.
var TeleporterRegistryABI = TeleporterRegistryMetaData.ABI

// TeleporterRegistry is an auto generated Go binding around an Ethereum contract.
type TeleporterRegistry struct {
	TeleporterRegistryCaller     // Read-only binding to the contract
	TeleporterRegistryTransactor // Write-only binding to the contract
	TeleporterRegistryFilterer   // Log filterer for contract events
}

// TeleporterRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeleporterRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeleporterRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeleporterRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeleporterRegistrySession struct {
	Contract     *TeleporterRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TeleporterRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeleporterRegistryCallerSession struct {
	Contract *TeleporterRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TeleporterRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeleporterRegistryTransactorSession struct {
	Contract     *TeleporterRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TeleporterRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeleporterRegistryRaw struct {
	Contract *TeleporterRegistry // Generic contract binding to access the raw methods on
}

// TeleporterRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeleporterRegistryCallerRaw struct {
	Contract *TeleporterRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TeleporterRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeleporterRegistryTransactorRaw struct {
	Contract *TeleporterRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeleporterRegistry creates a new instance of TeleporterRegistry, bound to a specific deployed contract.
func NewTeleporterRegistry(address common.Address, backend bind.ContractBackend) (*TeleporterRegistry, error) {
	contract, err := bindTeleporterRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistry{TeleporterRegistryCaller: TeleporterRegistryCaller{contract: contract}, TeleporterRegistryTransactor: TeleporterRegistryTransactor{contract: contract}, TeleporterRegistryFilterer: TeleporterRegistryFilterer{contract: contract}}, nil
}

// NewTeleporterRegistryCaller creates a new read-only instance of TeleporterRegistry, bound to a specific deployed contract.
func NewTeleporterRegistryCaller(address common.Address, caller bind.ContractCaller) (*TeleporterRegistryCaller, error) {
	contract, err := bindTeleporterRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistryCaller{contract: contract}, nil
}

// NewTeleporterRegistryTransactor creates a new write-only instance of TeleporterRegistry, bound to a specific deployed contract.
func NewTeleporterRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TeleporterRegistryTransactor, error) {
	contract, err := bindTeleporterRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistryTransactor{contract: contract}, nil
}

// NewTeleporterRegistryFilterer creates a new log filterer instance of TeleporterRegistry, bound to a specific deployed contract.
func NewTeleporterRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TeleporterRegistryFilterer, error) {
	contract, err := bindTeleporterRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistryFilterer{contract: contract}, nil
}

// bindTeleporterRegistry binds a generic wrapper to an already deployed contract.
func bindTeleporterRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeleporterRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterRegistry *TeleporterRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterRegistry.Contract.TeleporterRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterRegistry *TeleporterRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.TeleporterRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterRegistry *TeleporterRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.TeleporterRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterRegistry *TeleporterRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterRegistry *TeleporterRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterRegistry *TeleporterRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.contract.Transact(opts, method, params...)
}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCaller) VALIDATORSSOURCEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "VALIDATORS_SOURCE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_TeleporterRegistry *TeleporterRegistrySession) VALIDATORSSOURCEADDRESS() (common.Address, error) {
	return _TeleporterRegistry.Contract.VALIDATORSSOURCEADDRESS(&_TeleporterRegistry.CallOpts)
}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) VALIDATORSSOURCEADDRESS() (common.Address, error) {
	return _TeleporterRegistry.Contract.VALIDATORSSOURCEADDRESS(&_TeleporterRegistry.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterRegistry *TeleporterRegistrySession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterRegistry.Contract.WARPMESSENGER(&_TeleporterRegistry.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterRegistry.Contract.WARPMESSENGER(&_TeleporterRegistry.CallOpts)
}

// GetAddressFromVersion is a free data retrieval call binding the contract method 0x46f9ef49.
//
// Solidity: function getAddressFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCaller) GetAddressFromVersion(opts *bind.CallOpts, version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "getAddressFromVersion", version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressFromVersion is a free data retrieval call binding the contract method 0x46f9ef49.
//
// Solidity: function getAddressFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistrySession) GetAddressFromVersion(version *big.Int) (common.Address, error) {
	return _TeleporterRegistry.Contract.GetAddressFromVersion(&_TeleporterRegistry.CallOpts, version)
}

// GetAddressFromVersion is a free data retrieval call binding the contract method 0x46f9ef49.
//
// Solidity: function getAddressFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) GetAddressFromVersion(version *big.Int) (common.Address, error) {
	return _TeleporterRegistry.Contract.GetAddressFromVersion(&_TeleporterRegistry.CallOpts, version)
}

// GetLatestTeleporter is a free data retrieval call binding the contract method 0xd820e64f.
//
// Solidity: function getLatestTeleporter() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCaller) GetLatestTeleporter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "getLatestTeleporter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLatestTeleporter is a free data retrieval call binding the contract method 0xd820e64f.
//
// Solidity: function getLatestTeleporter() view returns(address)
func (_TeleporterRegistry *TeleporterRegistrySession) GetLatestTeleporter() (common.Address, error) {
	return _TeleporterRegistry.Contract.GetLatestTeleporter(&_TeleporterRegistry.CallOpts)
}

// GetLatestTeleporter is a free data retrieval call binding the contract method 0xd820e64f.
//
// Solidity: function getLatestTeleporter() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) GetLatestTeleporter() (common.Address, error) {
	return _TeleporterRegistry.Contract.GetLatestTeleporter(&_TeleporterRegistry.CallOpts)
}

// GetLatestVersion is a free data retrieval call binding the contract method 0x0e6d1de9.
//
// Solidity: function getLatestVersion() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCaller) GetLatestVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "getLatestVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestVersion is a free data retrieval call binding the contract method 0x0e6d1de9.
//
// Solidity: function getLatestVersion() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistrySession) GetLatestVersion() (*big.Int, error) {
	return _TeleporterRegistry.Contract.GetLatestVersion(&_TeleporterRegistry.CallOpts)
}

// GetLatestVersion is a free data retrieval call binding the contract method 0x0e6d1de9.
//
// Solidity: function getLatestVersion() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) GetLatestVersion() (*big.Int, error) {
	return _TeleporterRegistry.Contract.GetLatestVersion(&_TeleporterRegistry.CallOpts)
}

// GetTeleporterFromVersion is a free data retrieval call binding the contract method 0x215abce9.
//
// Solidity: function getTeleporterFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCaller) GetTeleporterFromVersion(opts *bind.CallOpts, version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "getTeleporterFromVersion", version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTeleporterFromVersion is a free data retrieval call binding the contract method 0x215abce9.
//
// Solidity: function getTeleporterFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistrySession) GetTeleporterFromVersion(version *big.Int) (common.Address, error) {
	return _TeleporterRegistry.Contract.GetTeleporterFromVersion(&_TeleporterRegistry.CallOpts, version)
}

// GetTeleporterFromVersion is a free data retrieval call binding the contract method 0x215abce9.
//
// Solidity: function getTeleporterFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) GetTeleporterFromVersion(version *big.Int) (common.Address, error) {
	return _TeleporterRegistry.Contract.GetTeleporterFromVersion(&_TeleporterRegistry.CallOpts, version)
}

// GetVersionFromAddress is a free data retrieval call binding the contract method 0x4c1f08ce.
//
// Solidity: function getVersionFromAddress(address protocolAddress) view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCaller) GetVersionFromAddress(opts *bind.CallOpts, protocolAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "getVersionFromAddress", protocolAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVersionFromAddress is a free data retrieval call binding the contract method 0x4c1f08ce.
//
// Solidity: function getVersionFromAddress(address protocolAddress) view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistrySession) GetVersionFromAddress(protocolAddress common.Address) (*big.Int, error) {
	return _TeleporterRegistry.Contract.GetVersionFromAddress(&_TeleporterRegistry.CallOpts, protocolAddress)
}

// GetVersionFromAddress is a free data retrieval call binding the contract method 0x4c1f08ce.
//
// Solidity: function getVersionFromAddress(address protocolAddress) view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) GetVersionFromAddress(protocolAddress common.Address) (*big.Int, error) {
	return _TeleporterRegistry.Contract.GetVersionFromAddress(&_TeleporterRegistry.CallOpts, protocolAddress)
}

// AddProtocolVersion is a paid mutator transaction binding the contract method 0x41f34ed9.
//
// Solidity: function addProtocolVersion(uint32 messageIndex) returns()
func (_TeleporterRegistry *TeleporterRegistryTransactor) AddProtocolVersion(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _TeleporterRegistry.contract.Transact(opts, "addProtocolVersion", messageIndex)
}

// AddProtocolVersion is a paid mutator transaction binding the contract method 0x41f34ed9.
//
// Solidity: function addProtocolVersion(uint32 messageIndex) returns()
func (_TeleporterRegistry *TeleporterRegistrySession) AddProtocolVersion(messageIndex uint32) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.AddProtocolVersion(&_TeleporterRegistry.TransactOpts, messageIndex)
}

// AddProtocolVersion is a paid mutator transaction binding the contract method 0x41f34ed9.
//
// Solidity: function addProtocolVersion(uint32 messageIndex) returns()
func (_TeleporterRegistry *TeleporterRegistryTransactorSession) AddProtocolVersion(messageIndex uint32) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.AddProtocolVersion(&_TeleporterRegistry.TransactOpts, messageIndex)
}

// TeleporterRegistryAddProtocolVersionIterator is returned from FilterAddProtocolVersion and is used to iterate over the raw logs and unpacked data for AddProtocolVersion events raised by the TeleporterRegistry contract.
type TeleporterRegistryAddProtocolVersionIterator struct {
	Event *TeleporterRegistryAddProtocolVersion // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterRegistryAddProtocolVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterRegistryAddProtocolVersion)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeleporterRegistryAddProtocolVersion)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeleporterRegistryAddProtocolVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterRegistryAddProtocolVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterRegistryAddProtocolVersion represents a AddProtocolVersion event raised by the TeleporterRegistry contract.
type TeleporterRegistryAddProtocolVersion struct {
	Version         *big.Int
	ProtocolAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddProtocolVersion is a free log retrieval operation binding the contract event 0xa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a.
//
// Solidity: event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress)
func (_TeleporterRegistry *TeleporterRegistryFilterer) FilterAddProtocolVersion(opts *bind.FilterOpts, version []*big.Int, protocolAddress []common.Address) (*TeleporterRegistryAddProtocolVersionIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var protocolAddressRule []interface{}
	for _, protocolAddressItem := range protocolAddress {
		protocolAddressRule = append(protocolAddressRule, protocolAddressItem)
	}

	logs, sub, err := _TeleporterRegistry.contract.FilterLogs(opts, "AddProtocolVersion", versionRule, protocolAddressRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistryAddProtocolVersionIterator{contract: _TeleporterRegistry.contract, event: "AddProtocolVersion", logs: logs, sub: sub}, nil
}

// WatchAddProtocolVersion is a free log subscription operation binding the contract event 0xa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a.
//
// Solidity: event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress)
func (_TeleporterRegistry *TeleporterRegistryFilterer) WatchAddProtocolVersion(opts *bind.WatchOpts, sink chan<- *TeleporterRegistryAddProtocolVersion, version []*big.Int, protocolAddress []common.Address) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var protocolAddressRule []interface{}
	for _, protocolAddressItem := range protocolAddress {
		protocolAddressRule = append(protocolAddressRule, protocolAddressItem)
	}

	logs, sub, err := _TeleporterRegistry.contract.WatchLogs(opts, "AddProtocolVersion", versionRule, protocolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterRegistryAddProtocolVersion)
				if err := _TeleporterRegistry.contract.UnpackLog(event, "AddProtocolVersion", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddProtocolVersion is a log parse operation binding the contract event 0xa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a.
//
// Solidity: event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress)
func (_TeleporterRegistry *TeleporterRegistryFilterer) ParseAddProtocolVersion(log types.Log) (*TeleporterRegistryAddProtocolVersion, error) {
	event := new(TeleporterRegistryAddProtocolVersion)
	if err := _TeleporterRegistry.contract.UnpackLog(event, "AddProtocolVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
