// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokenstakingmanager

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

// InitialValidator is an auto generated low-level Go binding around an user-defined struct.
type InitialValidator struct {
	NodeID       []byte
	BlsPublicKey []byte
	Weight       uint64
}

// PChainOwner is an auto generated low-level Go binding around an user-defined struct.
type PChainOwner struct {
	Threshold uint32
	Addresses []common.Address
}

// PoSValidatorManagerSettings is an auto generated low-level Go binding around an user-defined struct.
type PoSValidatorManagerSettings struct {
	BaseSettings             ValidatorManagerSettings
	MinimumStakeAmount       *big.Int
	MaximumStakeAmount       *big.Int
	MinimumStakeDuration     uint64
	MinimumDelegationFeeBips uint16
	MaximumStakeMultiplier   uint8
	RewardCalculator         common.Address
}

// SubnetConversionData is an auto generated low-level Go binding around an user-defined struct.
type SubnetConversionData struct {
	SubnetID                     [32]byte
	ValidatorManagerBlockchainID [32]byte
	ValidatorManagerAddress      common.Address
	InitialValidators            []InitialValidator
}

// Validator is an auto generated low-level Go binding around an user-defined struct.
type Validator struct {
	Status         uint8
	NodeID         []byte
	StartingWeight uint64
	MessageNonce   uint64
	Weight         uint64
	StartedAt      uint64
	EndedAt        uint64
}

// ValidatorManagerSettings is an auto generated low-level Go binding around an user-defined struct.
type ValidatorManagerSettings struct {
	SubnetID               [32]byte
	ChurnPeriodSeconds     uint64
	MaximumChurnPercentage uint8
}

// ValidatorMessagesValidationPeriod is an auto generated low-level Go binding around an user-defined struct.
type ValidatorMessagesValidationPeriod struct {
	SubnetID              [32]byte
	NodeID                []byte
	BlsPublicKey          []byte
	RegistrationExpiry    uint64
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
	Weight                uint64
}

// ValidatorRegistrationInput is an auto generated low-level Go binding around an user-defined struct.
type ValidatorRegistrationInput struct {
	NodeID                []byte
	BlsPublicKey          []byte
	RegistrationExpiry    uint64
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
}

// WarpBlockHash is an auto generated low-level Go binding around an user-defined struct.
type WarpBlockHash struct {
	SourceChainID [32]byte
	BlockHash     [32]byte
}

// WarpMessage is an auto generated low-level Go binding around an user-defined struct.
type WarpMessage struct {
	SourceChainID       [32]byte
	OriginSenderAddress common.Address
	Payload             []byte
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"}]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea26469706673582212209d2dabc5051a0f1ba488def6f101c5ffcdd34e71b16d3b6b944eda1ed980237464736f6c63430008190033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// ContextUpgradeableMetaData contains all meta data concerning the ContextUpgradeable contract.
var ContextUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// ContextUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextUpgradeableMetaData.ABI instead.
var ContextUpgradeableABI = ContextUpgradeableMetaData.ABI

// ContextUpgradeable is an auto generated Go binding around an Ethereum contract.
type ContextUpgradeable struct {
	ContextUpgradeableCaller     // Read-only binding to the contract
	ContextUpgradeableTransactor // Write-only binding to the contract
	ContextUpgradeableFilterer   // Log filterer for contract events
}

// ContextUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextUpgradeableSession struct {
	Contract     *ContextUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContextUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextUpgradeableCallerSession struct {
	Contract *ContextUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContextUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextUpgradeableTransactorSession struct {
	Contract     *ContextUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContextUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextUpgradeableRaw struct {
	Contract *ContextUpgradeable // Generic contract binding to access the raw methods on
}

// ContextUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextUpgradeableCallerRaw struct {
	Contract *ContextUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ContextUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactorRaw struct {
	Contract *ContextUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContextUpgradeable creates a new instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeable(address common.Address, backend bind.ContractBackend) (*ContextUpgradeable, error) {
	contract, err := bindContextUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeable{ContextUpgradeableCaller: ContextUpgradeableCaller{contract: contract}, ContextUpgradeableTransactor: ContextUpgradeableTransactor{contract: contract}, ContextUpgradeableFilterer: ContextUpgradeableFilterer{contract: contract}}, nil
}

// NewContextUpgradeableCaller creates a new read-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ContextUpgradeableCaller, error) {
	contract, err := bindContextUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableCaller{contract: contract}, nil
}

// NewContextUpgradeableTransactor creates a new write-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextUpgradeableTransactor, error) {
	contract, err := bindContextUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableTransactor{contract: contract}, nil
}

// NewContextUpgradeableFilterer creates a new log filterer instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextUpgradeableFilterer, error) {
	contract, err := bindContextUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableFilterer{contract: contract}, nil
}

// bindContextUpgradeable binds a generic wrapper to an already deployed contract.
func bindContextUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContextUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.ContextUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ContextUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContextUpgradeable contract.
type ContextUpgradeableInitializedIterator struct {
	Event *ContextUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *ContextUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContextUpgradeableInitialized)
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
		it.Event = new(ContextUpgradeableInitialized)
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
func (it *ContextUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContextUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContextUpgradeableInitialized represents a Initialized event raised by the ContextUpgradeable contract.
type ContextUpgradeableInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContextUpgradeableInitializedIterator, error) {

	logs, sub, err := _ContextUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableInitializedIterator{contract: _ContextUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContextUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _ContextUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContextUpgradeableInitialized)
				if err := _ContextUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) ParseInitialized(log types.Log) (*ContextUpgradeableInitialized, error) {
	event := new(ContextUpgradeableInitialized)
	if err := _ContextUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllowListMetaData contains all meta data concerning the IAllowList contract.
var IAllowListMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRole\",\"type\":\"uint256\"}],\"name\":\"RoleSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"readAllowList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAllowListABI is the input ABI used to generate the binding from.
// Deprecated: Use IAllowListMetaData.ABI instead.
var IAllowListABI = IAllowListMetaData.ABI

// IAllowList is an auto generated Go binding around an Ethereum contract.
type IAllowList struct {
	IAllowListCaller     // Read-only binding to the contract
	IAllowListTransactor // Write-only binding to the contract
	IAllowListFilterer   // Log filterer for contract events
}

// IAllowListCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAllowListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAllowListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAllowListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAllowListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAllowListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAllowListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAllowListSession struct {
	Contract     *IAllowList       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAllowListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAllowListCallerSession struct {
	Contract *IAllowListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IAllowListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAllowListTransactorSession struct {
	Contract     *IAllowListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IAllowListRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAllowListRaw struct {
	Contract *IAllowList // Generic contract binding to access the raw methods on
}

// IAllowListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAllowListCallerRaw struct {
	Contract *IAllowListCaller // Generic read-only contract binding to access the raw methods on
}

// IAllowListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAllowListTransactorRaw struct {
	Contract *IAllowListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAllowList creates a new instance of IAllowList, bound to a specific deployed contract.
func NewIAllowList(address common.Address, backend bind.ContractBackend) (*IAllowList, error) {
	contract, err := bindIAllowList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAllowList{IAllowListCaller: IAllowListCaller{contract: contract}, IAllowListTransactor: IAllowListTransactor{contract: contract}, IAllowListFilterer: IAllowListFilterer{contract: contract}}, nil
}

// NewIAllowListCaller creates a new read-only instance of IAllowList, bound to a specific deployed contract.
func NewIAllowListCaller(address common.Address, caller bind.ContractCaller) (*IAllowListCaller, error) {
	contract, err := bindIAllowList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAllowListCaller{contract: contract}, nil
}

// NewIAllowListTransactor creates a new write-only instance of IAllowList, bound to a specific deployed contract.
func NewIAllowListTransactor(address common.Address, transactor bind.ContractTransactor) (*IAllowListTransactor, error) {
	contract, err := bindIAllowList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAllowListTransactor{contract: contract}, nil
}

// NewIAllowListFilterer creates a new log filterer instance of IAllowList, bound to a specific deployed contract.
func NewIAllowListFilterer(address common.Address, filterer bind.ContractFilterer) (*IAllowListFilterer, error) {
	contract, err := bindIAllowList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAllowListFilterer{contract: contract}, nil
}

// bindIAllowList binds a generic wrapper to an already deployed contract.
func bindIAllowList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAllowListMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAllowList *IAllowListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAllowList.Contract.IAllowListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAllowList *IAllowListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAllowList.Contract.IAllowListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAllowList *IAllowListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAllowList.Contract.IAllowListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAllowList *IAllowListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAllowList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAllowList *IAllowListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAllowList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAllowList *IAllowListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAllowList.Contract.contract.Transact(opts, method, params...)
}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256 role)
func (_IAllowList *IAllowListCaller) ReadAllowList(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IAllowList.contract.Call(opts, &out, "readAllowList", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256 role)
func (_IAllowList *IAllowListSession) ReadAllowList(addr common.Address) (*big.Int, error) {
	return _IAllowList.Contract.ReadAllowList(&_IAllowList.CallOpts, addr)
}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256 role)
func (_IAllowList *IAllowListCallerSession) ReadAllowList(addr common.Address) (*big.Int, error) {
	return _IAllowList.Contract.ReadAllowList(&_IAllowList.CallOpts, addr)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_IAllowList *IAllowListTransactor) SetAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _IAllowList.contract.Transact(opts, "setAdmin", addr)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_IAllowList *IAllowListSession) SetAdmin(addr common.Address) (*types.Transaction, error) {
	return _IAllowList.Contract.SetAdmin(&_IAllowList.TransactOpts, addr)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_IAllowList *IAllowListTransactorSession) SetAdmin(addr common.Address) (*types.Transaction, error) {
	return _IAllowList.Contract.SetAdmin(&_IAllowList.TransactOpts, addr)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_IAllowList *IAllowListTransactor) SetEnabled(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _IAllowList.contract.Transact(opts, "setEnabled", addr)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_IAllowList *IAllowListSession) SetEnabled(addr common.Address) (*types.Transaction, error) {
	return _IAllowList.Contract.SetEnabled(&_IAllowList.TransactOpts, addr)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_IAllowList *IAllowListTransactorSession) SetEnabled(addr common.Address) (*types.Transaction, error) {
	return _IAllowList.Contract.SetEnabled(&_IAllowList.TransactOpts, addr)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address addr) returns()
func (_IAllowList *IAllowListTransactor) SetManager(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _IAllowList.contract.Transact(opts, "setManager", addr)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address addr) returns()
func (_IAllowList *IAllowListSession) SetManager(addr common.Address) (*types.Transaction, error) {
	return _IAllowList.Contract.SetManager(&_IAllowList.TransactOpts, addr)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address addr) returns()
func (_IAllowList *IAllowListTransactorSession) SetManager(addr common.Address) (*types.Transaction, error) {
	return _IAllowList.Contract.SetManager(&_IAllowList.TransactOpts, addr)
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_IAllowList *IAllowListTransactor) SetNone(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _IAllowList.contract.Transact(opts, "setNone", addr)
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_IAllowList *IAllowListSession) SetNone(addr common.Address) (*types.Transaction, error) {
	return _IAllowList.Contract.SetNone(&_IAllowList.TransactOpts, addr)
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_IAllowList *IAllowListTransactorSession) SetNone(addr common.Address) (*types.Transaction, error) {
	return _IAllowList.Contract.SetNone(&_IAllowList.TransactOpts, addr)
}

// IAllowListRoleSetIterator is returned from FilterRoleSet and is used to iterate over the raw logs and unpacked data for RoleSet events raised by the IAllowList contract.
type IAllowListRoleSetIterator struct {
	Event *IAllowListRoleSet // Event containing the contract specifics and raw log

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
func (it *IAllowListRoleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllowListRoleSet)
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
		it.Event = new(IAllowListRoleSet)
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
func (it *IAllowListRoleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllowListRoleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllowListRoleSet represents a RoleSet event raised by the IAllowList contract.
type IAllowListRoleSet struct {
	Role    *big.Int
	Account common.Address
	Sender  common.Address
	OldRole *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleSet is a free log retrieval operation binding the contract event 0xcdb7ea01f00a414d78757bdb0f6391664ba3fedf987eed280927c1e7d695be3e.
//
// Solidity: event RoleSet(uint256 indexed role, address indexed account, address indexed sender, uint256 oldRole)
func (_IAllowList *IAllowListFilterer) FilterRoleSet(opts *bind.FilterOpts, role []*big.Int, account []common.Address, sender []common.Address) (*IAllowListRoleSetIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAllowList.contract.FilterLogs(opts, "RoleSet", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IAllowListRoleSetIterator{contract: _IAllowList.contract, event: "RoleSet", logs: logs, sub: sub}, nil
}

// WatchRoleSet is a free log subscription operation binding the contract event 0xcdb7ea01f00a414d78757bdb0f6391664ba3fedf987eed280927c1e7d695be3e.
//
// Solidity: event RoleSet(uint256 indexed role, address indexed account, address indexed sender, uint256 oldRole)
func (_IAllowList *IAllowListFilterer) WatchRoleSet(opts *bind.WatchOpts, sink chan<- *IAllowListRoleSet, role []*big.Int, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAllowList.contract.WatchLogs(opts, "RoleSet", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllowListRoleSet)
				if err := _IAllowList.contract.UnpackLog(event, "RoleSet", log); err != nil {
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

// ParseRoleSet is a log parse operation binding the contract event 0xcdb7ea01f00a414d78757bdb0f6391664ba3fedf987eed280927c1e7d695be3e.
//
// Solidity: event RoleSet(uint256 indexed role, address indexed account, address indexed sender, uint256 oldRole)
func (_IAllowList *IAllowListFilterer) ParseRoleSet(log types.Log) (*IAllowListRoleSet, error) {
	event := new(IAllowListRoleSet)
	if err := _IAllowList.contract.UnpackLog(event, "RoleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeMinterMetaData contains all meta data concerning the INativeMinter contract.
var INativeMinterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NativeCoinMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRole\",\"type\":\"uint256\"}],\"name\":\"RoleSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintNativeCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"readAllowList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// INativeMinterABI is the input ABI used to generate the binding from.
// Deprecated: Use INativeMinterMetaData.ABI instead.
var INativeMinterABI = INativeMinterMetaData.ABI

// INativeMinter is an auto generated Go binding around an Ethereum contract.
type INativeMinter struct {
	INativeMinterCaller     // Read-only binding to the contract
	INativeMinterTransactor // Write-only binding to the contract
	INativeMinterFilterer   // Log filterer for contract events
}

// INativeMinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type INativeMinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INativeMinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INativeMinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INativeMinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INativeMinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INativeMinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INativeMinterSession struct {
	Contract     *INativeMinter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INativeMinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INativeMinterCallerSession struct {
	Contract *INativeMinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// INativeMinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INativeMinterTransactorSession struct {
	Contract     *INativeMinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// INativeMinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type INativeMinterRaw struct {
	Contract *INativeMinter // Generic contract binding to access the raw methods on
}

// INativeMinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INativeMinterCallerRaw struct {
	Contract *INativeMinterCaller // Generic read-only contract binding to access the raw methods on
}

// INativeMinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INativeMinterTransactorRaw struct {
	Contract *INativeMinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINativeMinter creates a new instance of INativeMinter, bound to a specific deployed contract.
func NewINativeMinter(address common.Address, backend bind.ContractBackend) (*INativeMinter, error) {
	contract, err := bindINativeMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INativeMinter{INativeMinterCaller: INativeMinterCaller{contract: contract}, INativeMinterTransactor: INativeMinterTransactor{contract: contract}, INativeMinterFilterer: INativeMinterFilterer{contract: contract}}, nil
}

// NewINativeMinterCaller creates a new read-only instance of INativeMinter, bound to a specific deployed contract.
func NewINativeMinterCaller(address common.Address, caller bind.ContractCaller) (*INativeMinterCaller, error) {
	contract, err := bindINativeMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INativeMinterCaller{contract: contract}, nil
}

// NewINativeMinterTransactor creates a new write-only instance of INativeMinter, bound to a specific deployed contract.
func NewINativeMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*INativeMinterTransactor, error) {
	contract, err := bindINativeMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INativeMinterTransactor{contract: contract}, nil
}

// NewINativeMinterFilterer creates a new log filterer instance of INativeMinter, bound to a specific deployed contract.
func NewINativeMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*INativeMinterFilterer, error) {
	contract, err := bindINativeMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INativeMinterFilterer{contract: contract}, nil
}

// bindINativeMinter binds a generic wrapper to an already deployed contract.
func bindINativeMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := INativeMinterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INativeMinter *INativeMinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INativeMinter.Contract.INativeMinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INativeMinter *INativeMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INativeMinter.Contract.INativeMinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INativeMinter *INativeMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INativeMinter.Contract.INativeMinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INativeMinter *INativeMinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INativeMinter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INativeMinter *INativeMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INativeMinter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INativeMinter *INativeMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INativeMinter.Contract.contract.Transact(opts, method, params...)
}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256 role)
func (_INativeMinter *INativeMinterCaller) ReadAllowList(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _INativeMinter.contract.Call(opts, &out, "readAllowList", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256 role)
func (_INativeMinter *INativeMinterSession) ReadAllowList(addr common.Address) (*big.Int, error) {
	return _INativeMinter.Contract.ReadAllowList(&_INativeMinter.CallOpts, addr)
}

// ReadAllowList is a free data retrieval call binding the contract method 0xeb54dae1.
//
// Solidity: function readAllowList(address addr) view returns(uint256 role)
func (_INativeMinter *INativeMinterCallerSession) ReadAllowList(addr common.Address) (*big.Int, error) {
	return _INativeMinter.Contract.ReadAllowList(&_INativeMinter.CallOpts, addr)
}

// MintNativeCoin is a paid mutator transaction binding the contract method 0x4f5aaaba.
//
// Solidity: function mintNativeCoin(address addr, uint256 amount) returns()
func (_INativeMinter *INativeMinterTransactor) MintNativeCoin(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _INativeMinter.contract.Transact(opts, "mintNativeCoin", addr, amount)
}

// MintNativeCoin is a paid mutator transaction binding the contract method 0x4f5aaaba.
//
// Solidity: function mintNativeCoin(address addr, uint256 amount) returns()
func (_INativeMinter *INativeMinterSession) MintNativeCoin(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _INativeMinter.Contract.MintNativeCoin(&_INativeMinter.TransactOpts, addr, amount)
}

// MintNativeCoin is a paid mutator transaction binding the contract method 0x4f5aaaba.
//
// Solidity: function mintNativeCoin(address addr, uint256 amount) returns()
func (_INativeMinter *INativeMinterTransactorSession) MintNativeCoin(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _INativeMinter.Contract.MintNativeCoin(&_INativeMinter.TransactOpts, addr, amount)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_INativeMinter *INativeMinterTransactor) SetAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _INativeMinter.contract.Transact(opts, "setAdmin", addr)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_INativeMinter *INativeMinterSession) SetAdmin(addr common.Address) (*types.Transaction, error) {
	return _INativeMinter.Contract.SetAdmin(&_INativeMinter.TransactOpts, addr)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address addr) returns()
func (_INativeMinter *INativeMinterTransactorSession) SetAdmin(addr common.Address) (*types.Transaction, error) {
	return _INativeMinter.Contract.SetAdmin(&_INativeMinter.TransactOpts, addr)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_INativeMinter *INativeMinterTransactor) SetEnabled(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _INativeMinter.contract.Transact(opts, "setEnabled", addr)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_INativeMinter *INativeMinterSession) SetEnabled(addr common.Address) (*types.Transaction, error) {
	return _INativeMinter.Contract.SetEnabled(&_INativeMinter.TransactOpts, addr)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x0aaf7043.
//
// Solidity: function setEnabled(address addr) returns()
func (_INativeMinter *INativeMinterTransactorSession) SetEnabled(addr common.Address) (*types.Transaction, error) {
	return _INativeMinter.Contract.SetEnabled(&_INativeMinter.TransactOpts, addr)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address addr) returns()
func (_INativeMinter *INativeMinterTransactor) SetManager(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _INativeMinter.contract.Transact(opts, "setManager", addr)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address addr) returns()
func (_INativeMinter *INativeMinterSession) SetManager(addr common.Address) (*types.Transaction, error) {
	return _INativeMinter.Contract.SetManager(&_INativeMinter.TransactOpts, addr)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address addr) returns()
func (_INativeMinter *INativeMinterTransactorSession) SetManager(addr common.Address) (*types.Transaction, error) {
	return _INativeMinter.Contract.SetManager(&_INativeMinter.TransactOpts, addr)
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_INativeMinter *INativeMinterTransactor) SetNone(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _INativeMinter.contract.Transact(opts, "setNone", addr)
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_INativeMinter *INativeMinterSession) SetNone(addr common.Address) (*types.Transaction, error) {
	return _INativeMinter.Contract.SetNone(&_INativeMinter.TransactOpts, addr)
}

// SetNone is a paid mutator transaction binding the contract method 0x8c6bfb3b.
//
// Solidity: function setNone(address addr) returns()
func (_INativeMinter *INativeMinterTransactorSession) SetNone(addr common.Address) (*types.Transaction, error) {
	return _INativeMinter.Contract.SetNone(&_INativeMinter.TransactOpts, addr)
}

// INativeMinterNativeCoinMintedIterator is returned from FilterNativeCoinMinted and is used to iterate over the raw logs and unpacked data for NativeCoinMinted events raised by the INativeMinter contract.
type INativeMinterNativeCoinMintedIterator struct {
	Event *INativeMinterNativeCoinMinted // Event containing the contract specifics and raw log

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
func (it *INativeMinterNativeCoinMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeMinterNativeCoinMinted)
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
		it.Event = new(INativeMinterNativeCoinMinted)
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
func (it *INativeMinterNativeCoinMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeMinterNativeCoinMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeMinterNativeCoinMinted represents a NativeCoinMinted event raised by the INativeMinter contract.
type INativeMinterNativeCoinMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNativeCoinMinted is a free log retrieval operation binding the contract event 0x400cd392f3d56fd10bb1dbd5839fdda8298208ddaa97b368faa053e1850930ee.
//
// Solidity: event NativeCoinMinted(address indexed sender, address indexed recipient, uint256 amount)
func (_INativeMinter *INativeMinterFilterer) FilterNativeCoinMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*INativeMinterNativeCoinMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _INativeMinter.contract.FilterLogs(opts, "NativeCoinMinted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &INativeMinterNativeCoinMintedIterator{contract: _INativeMinter.contract, event: "NativeCoinMinted", logs: logs, sub: sub}, nil
}

// WatchNativeCoinMinted is a free log subscription operation binding the contract event 0x400cd392f3d56fd10bb1dbd5839fdda8298208ddaa97b368faa053e1850930ee.
//
// Solidity: event NativeCoinMinted(address indexed sender, address indexed recipient, uint256 amount)
func (_INativeMinter *INativeMinterFilterer) WatchNativeCoinMinted(opts *bind.WatchOpts, sink chan<- *INativeMinterNativeCoinMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _INativeMinter.contract.WatchLogs(opts, "NativeCoinMinted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeMinterNativeCoinMinted)
				if err := _INativeMinter.contract.UnpackLog(event, "NativeCoinMinted", log); err != nil {
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

// ParseNativeCoinMinted is a log parse operation binding the contract event 0x400cd392f3d56fd10bb1dbd5839fdda8298208ddaa97b368faa053e1850930ee.
//
// Solidity: event NativeCoinMinted(address indexed sender, address indexed recipient, uint256 amount)
func (_INativeMinter *INativeMinterFilterer) ParseNativeCoinMinted(log types.Log) (*INativeMinterNativeCoinMinted, error) {
	event := new(INativeMinterNativeCoinMinted)
	if err := _INativeMinter.contract.UnpackLog(event, "NativeCoinMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeMinterRoleSetIterator is returned from FilterRoleSet and is used to iterate over the raw logs and unpacked data for RoleSet events raised by the INativeMinter contract.
type INativeMinterRoleSetIterator struct {
	Event *INativeMinterRoleSet // Event containing the contract specifics and raw log

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
func (it *INativeMinterRoleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeMinterRoleSet)
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
		it.Event = new(INativeMinterRoleSet)
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
func (it *INativeMinterRoleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeMinterRoleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeMinterRoleSet represents a RoleSet event raised by the INativeMinter contract.
type INativeMinterRoleSet struct {
	Role    *big.Int
	Account common.Address
	Sender  common.Address
	OldRole *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleSet is a free log retrieval operation binding the contract event 0xcdb7ea01f00a414d78757bdb0f6391664ba3fedf987eed280927c1e7d695be3e.
//
// Solidity: event RoleSet(uint256 indexed role, address indexed account, address indexed sender, uint256 oldRole)
func (_INativeMinter *INativeMinterFilterer) FilterRoleSet(opts *bind.FilterOpts, role []*big.Int, account []common.Address, sender []common.Address) (*INativeMinterRoleSetIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _INativeMinter.contract.FilterLogs(opts, "RoleSet", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &INativeMinterRoleSetIterator{contract: _INativeMinter.contract, event: "RoleSet", logs: logs, sub: sub}, nil
}

// WatchRoleSet is a free log subscription operation binding the contract event 0xcdb7ea01f00a414d78757bdb0f6391664ba3fedf987eed280927c1e7d695be3e.
//
// Solidity: event RoleSet(uint256 indexed role, address indexed account, address indexed sender, uint256 oldRole)
func (_INativeMinter *INativeMinterFilterer) WatchRoleSet(opts *bind.WatchOpts, sink chan<- *INativeMinterRoleSet, role []*big.Int, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _INativeMinter.contract.WatchLogs(opts, "RoleSet", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeMinterRoleSet)
				if err := _INativeMinter.contract.UnpackLog(event, "RoleSet", log); err != nil {
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

// ParseRoleSet is a log parse operation binding the contract event 0xcdb7ea01f00a414d78757bdb0f6391664ba3fedf987eed280927c1e7d695be3e.
//
// Solidity: event RoleSet(uint256 indexed role, address indexed account, address indexed sender, uint256 oldRole)
func (_INativeMinter *INativeMinterFilterer) ParseRoleSet(log types.Log) (*INativeMinterRoleSet, error) {
	event := new(INativeMinterRoleSet)
	if err := _INativeMinter.contract.UnpackLog(event, "RoleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenStakingManagerMetaData contains all meta data concerning the INativeTokenStakingManager contract.
var INativeTokenStakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"DelegationEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"delegatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"DelegatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"InitialValidatorCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"UptimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"ValidationPeriodCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ValidationPeriodEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidationPeriodRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"ValidationRewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"ValidatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorWeightUpdate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"claimDelegationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"claimValidationRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeDelegatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initializeDelegatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"internalType\":\"structValidatorRegistrationInput\",\"name\":\"registrationInput\",\"type\":\"tuple\"},{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"}],\"name\":\"initializeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structSubnetConversionData\",\"name\":\"subnetConversionData\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"messsageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendEndValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendRegisterValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"resendUpdateDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"submitUptimeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// INativeTokenStakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use INativeTokenStakingManagerMetaData.ABI instead.
var INativeTokenStakingManagerABI = INativeTokenStakingManagerMetaData.ABI

// INativeTokenStakingManager is an auto generated Go binding around an Ethereum contract.
type INativeTokenStakingManager struct {
	INativeTokenStakingManagerCaller     // Read-only binding to the contract
	INativeTokenStakingManagerTransactor // Write-only binding to the contract
	INativeTokenStakingManagerFilterer   // Log filterer for contract events
}

// INativeTokenStakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type INativeTokenStakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INativeTokenStakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INativeTokenStakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INativeTokenStakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INativeTokenStakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INativeTokenStakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INativeTokenStakingManagerSession struct {
	Contract     *INativeTokenStakingManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// INativeTokenStakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INativeTokenStakingManagerCallerSession struct {
	Contract *INativeTokenStakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// INativeTokenStakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INativeTokenStakingManagerTransactorSession struct {
	Contract     *INativeTokenStakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// INativeTokenStakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type INativeTokenStakingManagerRaw struct {
	Contract *INativeTokenStakingManager // Generic contract binding to access the raw methods on
}

// INativeTokenStakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INativeTokenStakingManagerCallerRaw struct {
	Contract *INativeTokenStakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// INativeTokenStakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INativeTokenStakingManagerTransactorRaw struct {
	Contract *INativeTokenStakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINativeTokenStakingManager creates a new instance of INativeTokenStakingManager, bound to a specific deployed contract.
func NewINativeTokenStakingManager(address common.Address, backend bind.ContractBackend) (*INativeTokenStakingManager, error) {
	contract, err := bindINativeTokenStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INativeTokenStakingManager{INativeTokenStakingManagerCaller: INativeTokenStakingManagerCaller{contract: contract}, INativeTokenStakingManagerTransactor: INativeTokenStakingManagerTransactor{contract: contract}, INativeTokenStakingManagerFilterer: INativeTokenStakingManagerFilterer{contract: contract}}, nil
}

// NewINativeTokenStakingManagerCaller creates a new read-only instance of INativeTokenStakingManager, bound to a specific deployed contract.
func NewINativeTokenStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*INativeTokenStakingManagerCaller, error) {
	contract, err := bindINativeTokenStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INativeTokenStakingManagerCaller{contract: contract}, nil
}

// NewINativeTokenStakingManagerTransactor creates a new write-only instance of INativeTokenStakingManager, bound to a specific deployed contract.
func NewINativeTokenStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*INativeTokenStakingManagerTransactor, error) {
	contract, err := bindINativeTokenStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INativeTokenStakingManagerTransactor{contract: contract}, nil
}

// NewINativeTokenStakingManagerFilterer creates a new log filterer instance of INativeTokenStakingManager, bound to a specific deployed contract.
func NewINativeTokenStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*INativeTokenStakingManagerFilterer, error) {
	contract, err := bindINativeTokenStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INativeTokenStakingManagerFilterer{contract: contract}, nil
}

// bindINativeTokenStakingManager binds a generic wrapper to an already deployed contract.
func bindINativeTokenStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := INativeTokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INativeTokenStakingManager *INativeTokenStakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INativeTokenStakingManager.Contract.INativeTokenStakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INativeTokenStakingManager *INativeTokenStakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.INativeTokenStakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INativeTokenStakingManager *INativeTokenStakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.INativeTokenStakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INativeTokenStakingManager *INativeTokenStakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INativeTokenStakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.contract.Transact(opts, method, params...)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) ClaimDelegationFees(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "claimDelegationFees", validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.ClaimDelegationFees(&_INativeTokenStakingManager.TransactOpts, validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.ClaimDelegationFees(&_INativeTokenStakingManager.TransactOpts, validationID)
}

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) ClaimValidationRewards(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "claimValidationRewards", validationID, messageIndex)
}

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) ClaimValidationRewards(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.ClaimValidationRewards(&_INativeTokenStakingManager.TransactOpts, validationID, messageIndex)
}

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) ClaimValidationRewards(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.ClaimValidationRewards(&_INativeTokenStakingManager.TransactOpts, validationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "completeDelegatorRegistration", delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.CompleteDelegatorRegistration(&_INativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.CompleteDelegatorRegistration(&_INativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) CompleteEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "completeEndDelegation", delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.CompleteEndDelegation(&_INativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.CompleteEndDelegation(&_INativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) CompleteEndValidation(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "completeEndValidation", messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.CompleteEndValidation(&_INativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.CompleteEndValidation(&_INativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.CompleteValidatorRegistration(&_INativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.CompleteValidatorRegistration(&_INativeTokenStakingManager.TransactOpts, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) ForceInitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "forceInitializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.ForceInitializeEndDelegation(&_INativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.ForceInitializeEndDelegation(&_INativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) ForceInitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "forceInitializeEndValidation", validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.ForceInitializeEndValidation(&_INativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.ForceInitializeEndValidation(&_INativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0xc599e24f.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID) payable returns(bytes32)
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) InitializeDelegatorRegistration(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "initializeDelegatorRegistration", validationID)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0xc599e24f.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID) payable returns(bytes32)
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) InitializeDelegatorRegistration(validationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.InitializeDelegatorRegistration(&_INativeTokenStakingManager.TransactOpts, validationID)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0xc599e24f.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID) payable returns(bytes32)
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) InitializeDelegatorRegistration(validationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.InitializeDelegatorRegistration(&_INativeTokenStakingManager.TransactOpts, validationID)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) InitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "initializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.InitializeEndDelegation(&_INativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.InitializeEndDelegation(&_INativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) InitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "initializeEndValidation", validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.InitializeEndValidation(&_INativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.InitializeEndValidation(&_INativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0xaf2f5feb.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration) payable returns(bytes32 validationID)
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) InitializeValidatorRegistration(opts *bind.TransactOpts, registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "initializeValidatorRegistration", registrationInput, delegationFeeBips, minStakeDuration)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0xaf2f5feb.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration) payable returns(bytes32 validationID)
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) InitializeValidatorRegistration(registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.InitializeValidatorRegistration(&_INativeTokenStakingManager.TransactOpts, registrationInput, delegationFeeBips, minStakeDuration)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0xaf2f5feb.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration) payable returns(bytes32 validationID)
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) InitializeValidatorRegistration(registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.InitializeValidatorRegistration(&_INativeTokenStakingManager.TransactOpts, registrationInput, delegationFeeBips, minStakeDuration)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messsageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) InitializeValidatorSet(opts *bind.TransactOpts, subnetConversionData SubnetConversionData, messsageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "initializeValidatorSet", subnetConversionData, messsageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messsageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) InitializeValidatorSet(subnetConversionData SubnetConversionData, messsageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.InitializeValidatorSet(&_INativeTokenStakingManager.TransactOpts, subnetConversionData, messsageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messsageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) InitializeValidatorSet(subnetConversionData SubnetConversionData, messsageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.InitializeValidatorSet(&_INativeTokenStakingManager.TransactOpts, subnetConversionData, messsageIndex)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) ResendEndValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "resendEndValidatorMessage", validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.ResendEndValidatorMessage(&_INativeTokenStakingManager.TransactOpts, validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.ResendEndValidatorMessage(&_INativeTokenStakingManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) ResendRegisterValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "resendRegisterValidatorMessage", validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.ResendRegisterValidatorMessage(&_INativeTokenStakingManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.ResendRegisterValidatorMessage(&_INativeTokenStakingManager.TransactOpts, validationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) ResendUpdateDelegation(opts *bind.TransactOpts, delegationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "resendUpdateDelegation", delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.ResendUpdateDelegation(&_INativeTokenStakingManager.TransactOpts, delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.ResendUpdateDelegation(&_INativeTokenStakingManager.TransactOpts, delegationID)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactor) SubmitUptimeProof(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.contract.Transact(opts, "submitUptimeProof", validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.SubmitUptimeProof(&_INativeTokenStakingManager.TransactOpts, validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_INativeTokenStakingManager *INativeTokenStakingManagerTransactorSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenStakingManager.Contract.SubmitUptimeProof(&_INativeTokenStakingManager.TransactOpts, validationID, messageIndex)
}

// INativeTokenStakingManagerDelegationEndedIterator is returned from FilterDelegationEnded and is used to iterate over the raw logs and unpacked data for DelegationEnded events raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerDelegationEndedIterator struct {
	Event *INativeTokenStakingManagerDelegationEnded // Event containing the contract specifics and raw log

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
func (it *INativeTokenStakingManagerDelegationEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenStakingManagerDelegationEnded)
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
		it.Event = new(INativeTokenStakingManagerDelegationEnded)
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
func (it *INativeTokenStakingManagerDelegationEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenStakingManagerDelegationEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenStakingManagerDelegationEnded represents a DelegationEnded event raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerDelegationEnded struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Rewards      *big.Int
	Fees         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegationEnded is a free log retrieval operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) FilterDelegationEnded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*INativeTokenStakingManagerDelegationEndedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.FilterLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenStakingManagerDelegationEndedIterator{contract: _INativeTokenStakingManager.contract, event: "DelegationEnded", logs: logs, sub: sub}, nil
}

// WatchDelegationEnded is a free log subscription operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) WatchDelegationEnded(opts *bind.WatchOpts, sink chan<- *INativeTokenStakingManagerDelegationEnded, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.WatchLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenStakingManagerDelegationEnded)
				if err := _INativeTokenStakingManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
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

// ParseDelegationEnded is a log parse operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) ParseDelegationEnded(log types.Log) (*INativeTokenStakingManagerDelegationEnded, error) {
	event := new(INativeTokenStakingManagerDelegationEnded)
	if err := _INativeTokenStakingManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenStakingManagerDelegatorAddedIterator is returned from FilterDelegatorAdded and is used to iterate over the raw logs and unpacked data for DelegatorAdded events raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerDelegatorAddedIterator struct {
	Event *INativeTokenStakingManagerDelegatorAdded // Event containing the contract specifics and raw log

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
func (it *INativeTokenStakingManagerDelegatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenStakingManagerDelegatorAdded)
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
		it.Event = new(INativeTokenStakingManagerDelegatorAdded)
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
func (it *INativeTokenStakingManagerDelegatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenStakingManagerDelegatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenStakingManagerDelegatorAdded represents a DelegatorAdded event raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerDelegatorAdded struct {
	DelegationID       [32]byte
	ValidationID       [32]byte
	DelegatorAddress   common.Address
	Nonce              uint64
	ValidatorWeight    uint64
	DelegatorWeight    uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDelegatorAdded is a free log retrieval operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) FilterDelegatorAdded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (*INativeTokenStakingManagerDelegatorAddedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.FilterLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenStakingManagerDelegatorAddedIterator{contract: _INativeTokenStakingManager.contract, event: "DelegatorAdded", logs: logs, sub: sub}, nil
}

// WatchDelegatorAdded is a free log subscription operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) WatchDelegatorAdded(opts *bind.WatchOpts, sink chan<- *INativeTokenStakingManagerDelegatorAdded, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.WatchLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenStakingManagerDelegatorAdded)
				if err := _INativeTokenStakingManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
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

// ParseDelegatorAdded is a log parse operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) ParseDelegatorAdded(log types.Log) (*INativeTokenStakingManagerDelegatorAdded, error) {
	event := new(INativeTokenStakingManagerDelegatorAdded)
	if err := _INativeTokenStakingManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenStakingManagerDelegatorRegisteredIterator is returned from FilterDelegatorRegistered and is used to iterate over the raw logs and unpacked data for DelegatorRegistered events raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerDelegatorRegisteredIterator struct {
	Event *INativeTokenStakingManagerDelegatorRegistered // Event containing the contract specifics and raw log

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
func (it *INativeTokenStakingManagerDelegatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenStakingManagerDelegatorRegistered)
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
		it.Event = new(INativeTokenStakingManagerDelegatorRegistered)
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
func (it *INativeTokenStakingManagerDelegatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenStakingManagerDelegatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenStakingManagerDelegatorRegistered represents a DelegatorRegistered event raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerDelegatorRegistered struct {
	DelegationID [32]byte
	ValidationID [32]byte
	StartTime    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRegistered is a free log retrieval operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) FilterDelegatorRegistered(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*INativeTokenStakingManagerDelegatorRegisteredIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.FilterLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenStakingManagerDelegatorRegisteredIterator{contract: _INativeTokenStakingManager.contract, event: "DelegatorRegistered", logs: logs, sub: sub}, nil
}

// WatchDelegatorRegistered is a free log subscription operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) WatchDelegatorRegistered(opts *bind.WatchOpts, sink chan<- *INativeTokenStakingManagerDelegatorRegistered, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.WatchLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenStakingManagerDelegatorRegistered)
				if err := _INativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
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

// ParseDelegatorRegistered is a log parse operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) ParseDelegatorRegistered(log types.Log) (*INativeTokenStakingManagerDelegatorRegistered, error) {
	event := new(INativeTokenStakingManagerDelegatorRegistered)
	if err := _INativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenStakingManagerDelegatorRemovalInitializedIterator is returned from FilterDelegatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for DelegatorRemovalInitialized events raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerDelegatorRemovalInitializedIterator struct {
	Event *INativeTokenStakingManagerDelegatorRemovalInitialized // Event containing the contract specifics and raw log

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
func (it *INativeTokenStakingManagerDelegatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenStakingManagerDelegatorRemovalInitialized)
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
		it.Event = new(INativeTokenStakingManagerDelegatorRemovalInitialized)
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
func (it *INativeTokenStakingManagerDelegatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenStakingManagerDelegatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenStakingManagerDelegatorRemovalInitialized represents a DelegatorRemovalInitialized event raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerDelegatorRemovalInitialized struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRemovalInitialized is a free log retrieval operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) FilterDelegatorRemovalInitialized(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*INativeTokenStakingManagerDelegatorRemovalInitializedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.FilterLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenStakingManagerDelegatorRemovalInitializedIterator{contract: _INativeTokenStakingManager.contract, event: "DelegatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchDelegatorRemovalInitialized is a free log subscription operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) WatchDelegatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *INativeTokenStakingManagerDelegatorRemovalInitialized, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.WatchLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenStakingManagerDelegatorRemovalInitialized)
				if err := _INativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
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

// ParseDelegatorRemovalInitialized is a log parse operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) ParseDelegatorRemovalInitialized(log types.Log) (*INativeTokenStakingManagerDelegatorRemovalInitialized, error) {
	event := new(INativeTokenStakingManagerDelegatorRemovalInitialized)
	if err := _INativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenStakingManagerInitialValidatorCreatedIterator is returned from FilterInitialValidatorCreated and is used to iterate over the raw logs and unpacked data for InitialValidatorCreated events raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerInitialValidatorCreatedIterator struct {
	Event *INativeTokenStakingManagerInitialValidatorCreated // Event containing the contract specifics and raw log

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
func (it *INativeTokenStakingManagerInitialValidatorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenStakingManagerInitialValidatorCreated)
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
		it.Event = new(INativeTokenStakingManagerInitialValidatorCreated)
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
func (it *INativeTokenStakingManagerInitialValidatorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenStakingManagerInitialValidatorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenStakingManagerInitialValidatorCreated represents a InitialValidatorCreated event raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerInitialValidatorCreated struct {
	ValidationID [32]byte
	NodeID       common.Hash
	Weight       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialValidatorCreated is a free log retrieval operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) FilterInitialValidatorCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte) (*INativeTokenStakingManagerInitialValidatorCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.FilterLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenStakingManagerInitialValidatorCreatedIterator{contract: _INativeTokenStakingManager.contract, event: "InitialValidatorCreated", logs: logs, sub: sub}, nil
}

// WatchInitialValidatorCreated is a free log subscription operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) WatchInitialValidatorCreated(opts *bind.WatchOpts, sink chan<- *INativeTokenStakingManagerInitialValidatorCreated, validationID [][32]byte, nodeID [][]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.WatchLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenStakingManagerInitialValidatorCreated)
				if err := _INativeTokenStakingManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
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

// ParseInitialValidatorCreated is a log parse operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) ParseInitialValidatorCreated(log types.Log) (*INativeTokenStakingManagerInitialValidatorCreated, error) {
	event := new(INativeTokenStakingManagerInitialValidatorCreated)
	if err := _INativeTokenStakingManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenStakingManagerUptimeUpdatedIterator is returned from FilterUptimeUpdated and is used to iterate over the raw logs and unpacked data for UptimeUpdated events raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerUptimeUpdatedIterator struct {
	Event *INativeTokenStakingManagerUptimeUpdated // Event containing the contract specifics and raw log

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
func (it *INativeTokenStakingManagerUptimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenStakingManagerUptimeUpdated)
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
		it.Event = new(INativeTokenStakingManagerUptimeUpdated)
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
func (it *INativeTokenStakingManagerUptimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenStakingManagerUptimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenStakingManagerUptimeUpdated represents a UptimeUpdated event raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerUptimeUpdated struct {
	ValidationID [32]byte
	Uptime       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUptimeUpdated is a free log retrieval operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) FilterUptimeUpdated(opts *bind.FilterOpts, validationID [][32]byte) (*INativeTokenStakingManagerUptimeUpdatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.FilterLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenStakingManagerUptimeUpdatedIterator{contract: _INativeTokenStakingManager.contract, event: "UptimeUpdated", logs: logs, sub: sub}, nil
}

// WatchUptimeUpdated is a free log subscription operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) WatchUptimeUpdated(opts *bind.WatchOpts, sink chan<- *INativeTokenStakingManagerUptimeUpdated, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.WatchLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenStakingManagerUptimeUpdated)
				if err := _INativeTokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
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

// ParseUptimeUpdated is a log parse operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) ParseUptimeUpdated(log types.Log) (*INativeTokenStakingManagerUptimeUpdated, error) {
	event := new(INativeTokenStakingManagerUptimeUpdated)
	if err := _INativeTokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenStakingManagerValidationPeriodCreatedIterator is returned from FilterValidationPeriodCreated and is used to iterate over the raw logs and unpacked data for ValidationPeriodCreated events raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerValidationPeriodCreatedIterator struct {
	Event *INativeTokenStakingManagerValidationPeriodCreated // Event containing the contract specifics and raw log

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
func (it *INativeTokenStakingManagerValidationPeriodCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenStakingManagerValidationPeriodCreated)
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
		it.Event = new(INativeTokenStakingManagerValidationPeriodCreated)
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
func (it *INativeTokenStakingManagerValidationPeriodCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenStakingManagerValidationPeriodCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenStakingManagerValidationPeriodCreated represents a ValidationPeriodCreated event raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerValidationPeriodCreated struct {
	ValidationID                [32]byte
	NodeID                      common.Hash
	RegisterValidationMessageID [32]byte
	Weight                      *big.Int
	RegistrationExpiry          uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodCreated is a free log retrieval operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) FilterValidationPeriodCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (*INativeTokenStakingManagerValidationPeriodCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenStakingManagerValidationPeriodCreatedIterator{contract: _INativeTokenStakingManager.contract, event: "ValidationPeriodCreated", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodCreated is a free log subscription operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) WatchValidationPeriodCreated(opts *bind.WatchOpts, sink chan<- *INativeTokenStakingManagerValidationPeriodCreated, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenStakingManagerValidationPeriodCreated)
				if err := _INativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
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

// ParseValidationPeriodCreated is a log parse operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) ParseValidationPeriodCreated(log types.Log) (*INativeTokenStakingManagerValidationPeriodCreated, error) {
	event := new(INativeTokenStakingManagerValidationPeriodCreated)
	if err := _INativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenStakingManagerValidationPeriodEndedIterator is returned from FilterValidationPeriodEnded and is used to iterate over the raw logs and unpacked data for ValidationPeriodEnded events raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerValidationPeriodEndedIterator struct {
	Event *INativeTokenStakingManagerValidationPeriodEnded // Event containing the contract specifics and raw log

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
func (it *INativeTokenStakingManagerValidationPeriodEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenStakingManagerValidationPeriodEnded)
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
		it.Event = new(INativeTokenStakingManagerValidationPeriodEnded)
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
func (it *INativeTokenStakingManagerValidationPeriodEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenStakingManagerValidationPeriodEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenStakingManagerValidationPeriodEnded represents a ValidationPeriodEnded event raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerValidationPeriodEnded struct {
	ValidationID [32]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodEnded is a free log retrieval operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) FilterValidationPeriodEnded(opts *bind.FilterOpts, validationID [][32]byte, status []uint8) (*INativeTokenStakingManagerValidationPeriodEndedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenStakingManagerValidationPeriodEndedIterator{contract: _INativeTokenStakingManager.contract, event: "ValidationPeriodEnded", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodEnded is a free log subscription operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) WatchValidationPeriodEnded(opts *bind.WatchOpts, sink chan<- *INativeTokenStakingManagerValidationPeriodEnded, validationID [][32]byte, status []uint8) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenStakingManagerValidationPeriodEnded)
				if err := _INativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
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

// ParseValidationPeriodEnded is a log parse operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) ParseValidationPeriodEnded(log types.Log) (*INativeTokenStakingManagerValidationPeriodEnded, error) {
	event := new(INativeTokenStakingManagerValidationPeriodEnded)
	if err := _INativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenStakingManagerValidationPeriodRegisteredIterator is returned from FilterValidationPeriodRegistered and is used to iterate over the raw logs and unpacked data for ValidationPeriodRegistered events raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerValidationPeriodRegisteredIterator struct {
	Event *INativeTokenStakingManagerValidationPeriodRegistered // Event containing the contract specifics and raw log

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
func (it *INativeTokenStakingManagerValidationPeriodRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenStakingManagerValidationPeriodRegistered)
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
		it.Event = new(INativeTokenStakingManagerValidationPeriodRegistered)
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
func (it *INativeTokenStakingManagerValidationPeriodRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenStakingManagerValidationPeriodRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenStakingManagerValidationPeriodRegistered represents a ValidationPeriodRegistered event raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerValidationPeriodRegistered struct {
	ValidationID [32]byte
	Weight       *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodRegistered is a free log retrieval operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) FilterValidationPeriodRegistered(opts *bind.FilterOpts, validationID [][32]byte) (*INativeTokenStakingManagerValidationPeriodRegisteredIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenStakingManagerValidationPeriodRegisteredIterator{contract: _INativeTokenStakingManager.contract, event: "ValidationPeriodRegistered", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodRegistered is a free log subscription operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) WatchValidationPeriodRegistered(opts *bind.WatchOpts, sink chan<- *INativeTokenStakingManagerValidationPeriodRegistered, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenStakingManagerValidationPeriodRegistered)
				if err := _INativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
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

// ParseValidationPeriodRegistered is a log parse operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) ParseValidationPeriodRegistered(log types.Log) (*INativeTokenStakingManagerValidationPeriodRegistered, error) {
	event := new(INativeTokenStakingManagerValidationPeriodRegistered)
	if err := _INativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenStakingManagerValidationRewardsClaimedIterator is returned from FilterValidationRewardsClaimed and is used to iterate over the raw logs and unpacked data for ValidationRewardsClaimed events raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerValidationRewardsClaimedIterator struct {
	Event *INativeTokenStakingManagerValidationRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *INativeTokenStakingManagerValidationRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenStakingManagerValidationRewardsClaimed)
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
		it.Event = new(INativeTokenStakingManagerValidationRewardsClaimed)
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
func (it *INativeTokenStakingManagerValidationRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenStakingManagerValidationRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenStakingManagerValidationRewardsClaimed represents a ValidationRewardsClaimed event raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerValidationRewardsClaimed struct {
	ValidationID [32]byte
	Reward       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationRewardsClaimed is a free log retrieval operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) FilterValidationRewardsClaimed(opts *bind.FilterOpts, validationID [][32]byte, reward []*big.Int) (*INativeTokenStakingManagerValidationRewardsClaimedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.FilterLogs(opts, "ValidationRewardsClaimed", validationIDRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenStakingManagerValidationRewardsClaimedIterator{contract: _INativeTokenStakingManager.contract, event: "ValidationRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchValidationRewardsClaimed is a free log subscription operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) WatchValidationRewardsClaimed(opts *bind.WatchOpts, sink chan<- *INativeTokenStakingManagerValidationRewardsClaimed, validationID [][32]byte, reward []*big.Int) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.WatchLogs(opts, "ValidationRewardsClaimed", validationIDRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenStakingManagerValidationRewardsClaimed)
				if err := _INativeTokenStakingManager.contract.UnpackLog(event, "ValidationRewardsClaimed", log); err != nil {
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

// ParseValidationRewardsClaimed is a log parse operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) ParseValidationRewardsClaimed(log types.Log) (*INativeTokenStakingManagerValidationRewardsClaimed, error) {
	event := new(INativeTokenStakingManagerValidationRewardsClaimed)
	if err := _INativeTokenStakingManager.contract.UnpackLog(event, "ValidationRewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenStakingManagerValidatorRemovalInitializedIterator is returned from FilterValidatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for ValidatorRemovalInitialized events raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerValidatorRemovalInitializedIterator struct {
	Event *INativeTokenStakingManagerValidatorRemovalInitialized // Event containing the contract specifics and raw log

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
func (it *INativeTokenStakingManagerValidatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenStakingManagerValidatorRemovalInitialized)
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
		it.Event = new(INativeTokenStakingManagerValidatorRemovalInitialized)
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
func (it *INativeTokenStakingManagerValidatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenStakingManagerValidatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenStakingManagerValidatorRemovalInitialized represents a ValidatorRemovalInitialized event raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerValidatorRemovalInitialized struct {
	ValidationID       [32]byte
	SetWeightMessageID [32]byte
	Weight             *big.Int
	EndTime            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemovalInitialized is a free log retrieval operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) FilterValidatorRemovalInitialized(opts *bind.FilterOpts, validationID [][32]byte, setWeightMessageID [][32]byte) (*INativeTokenStakingManagerValidatorRemovalInitializedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.FilterLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenStakingManagerValidatorRemovalInitializedIterator{contract: _INativeTokenStakingManager.contract, event: "ValidatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchValidatorRemovalInitialized is a free log subscription operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) WatchValidatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *INativeTokenStakingManagerValidatorRemovalInitialized, validationID [][32]byte, setWeightMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.WatchLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenStakingManagerValidatorRemovalInitialized)
				if err := _INativeTokenStakingManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
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

// ParseValidatorRemovalInitialized is a log parse operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) ParseValidatorRemovalInitialized(log types.Log) (*INativeTokenStakingManagerValidatorRemovalInitialized, error) {
	event := new(INativeTokenStakingManagerValidatorRemovalInitialized)
	if err := _INativeTokenStakingManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenStakingManagerValidatorWeightUpdateIterator is returned from FilterValidatorWeightUpdate and is used to iterate over the raw logs and unpacked data for ValidatorWeightUpdate events raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerValidatorWeightUpdateIterator struct {
	Event *INativeTokenStakingManagerValidatorWeightUpdate // Event containing the contract specifics and raw log

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
func (it *INativeTokenStakingManagerValidatorWeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenStakingManagerValidatorWeightUpdate)
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
		it.Event = new(INativeTokenStakingManagerValidatorWeightUpdate)
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
func (it *INativeTokenStakingManagerValidatorWeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenStakingManagerValidatorWeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenStakingManagerValidatorWeightUpdate represents a ValidatorWeightUpdate event raised by the INativeTokenStakingManager contract.
type INativeTokenStakingManagerValidatorWeightUpdate struct {
	ValidationID       [32]byte
	Nonce              uint64
	ValidatorWeight    uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorWeightUpdate is a free log retrieval operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) FilterValidatorWeightUpdate(opts *bind.FilterOpts, validationID [][32]byte, nonce []uint64) (*INativeTokenStakingManagerValidatorWeightUpdateIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.FilterLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenStakingManagerValidatorWeightUpdateIterator{contract: _INativeTokenStakingManager.contract, event: "ValidatorWeightUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorWeightUpdate is a free log subscription operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) WatchValidatorWeightUpdate(opts *bind.WatchOpts, sink chan<- *INativeTokenStakingManagerValidatorWeightUpdate, validationID [][32]byte, nonce []uint64) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _INativeTokenStakingManager.contract.WatchLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenStakingManagerValidatorWeightUpdate)
				if err := _INativeTokenStakingManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
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

// ParseValidatorWeightUpdate is a log parse operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_INativeTokenStakingManager *INativeTokenStakingManagerFilterer) ParseValidatorWeightUpdate(log types.Log) (*INativeTokenStakingManagerValidatorWeightUpdate, error) {
	event := new(INativeTokenStakingManagerValidatorWeightUpdate)
	if err := _INativeTokenStakingManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoSValidatorManagerMetaData contains all meta data concerning the IPoSValidatorManager contract.
var IPoSValidatorManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"DelegationEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"delegatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"DelegatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"InitialValidatorCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"UptimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"ValidationPeriodCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ValidationPeriodEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidationPeriodRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"ValidationRewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"ValidatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorWeightUpdate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"claimDelegationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"claimValidationRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeDelegatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structSubnetConversionData\",\"name\":\"subnetConversionData\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"messsageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendEndValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendRegisterValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"resendUpdateDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"submitUptimeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPoSValidatorManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IPoSValidatorManagerMetaData.ABI instead.
var IPoSValidatorManagerABI = IPoSValidatorManagerMetaData.ABI

// IPoSValidatorManager is an auto generated Go binding around an Ethereum contract.
type IPoSValidatorManager struct {
	IPoSValidatorManagerCaller     // Read-only binding to the contract
	IPoSValidatorManagerTransactor // Write-only binding to the contract
	IPoSValidatorManagerFilterer   // Log filterer for contract events
}

// IPoSValidatorManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPoSValidatorManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoSValidatorManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPoSValidatorManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoSValidatorManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPoSValidatorManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoSValidatorManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPoSValidatorManagerSession struct {
	Contract     *IPoSValidatorManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IPoSValidatorManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPoSValidatorManagerCallerSession struct {
	Contract *IPoSValidatorManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IPoSValidatorManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPoSValidatorManagerTransactorSession struct {
	Contract     *IPoSValidatorManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IPoSValidatorManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPoSValidatorManagerRaw struct {
	Contract *IPoSValidatorManager // Generic contract binding to access the raw methods on
}

// IPoSValidatorManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPoSValidatorManagerCallerRaw struct {
	Contract *IPoSValidatorManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IPoSValidatorManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPoSValidatorManagerTransactorRaw struct {
	Contract *IPoSValidatorManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPoSValidatorManager creates a new instance of IPoSValidatorManager, bound to a specific deployed contract.
func NewIPoSValidatorManager(address common.Address, backend bind.ContractBackend) (*IPoSValidatorManager, error) {
	contract, err := bindIPoSValidatorManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPoSValidatorManager{IPoSValidatorManagerCaller: IPoSValidatorManagerCaller{contract: contract}, IPoSValidatorManagerTransactor: IPoSValidatorManagerTransactor{contract: contract}, IPoSValidatorManagerFilterer: IPoSValidatorManagerFilterer{contract: contract}}, nil
}

// NewIPoSValidatorManagerCaller creates a new read-only instance of IPoSValidatorManager, bound to a specific deployed contract.
func NewIPoSValidatorManagerCaller(address common.Address, caller bind.ContractCaller) (*IPoSValidatorManagerCaller, error) {
	contract, err := bindIPoSValidatorManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPoSValidatorManagerCaller{contract: contract}, nil
}

// NewIPoSValidatorManagerTransactor creates a new write-only instance of IPoSValidatorManager, bound to a specific deployed contract.
func NewIPoSValidatorManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IPoSValidatorManagerTransactor, error) {
	contract, err := bindIPoSValidatorManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPoSValidatorManagerTransactor{contract: contract}, nil
}

// NewIPoSValidatorManagerFilterer creates a new log filterer instance of IPoSValidatorManager, bound to a specific deployed contract.
func NewIPoSValidatorManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IPoSValidatorManagerFilterer, error) {
	contract, err := bindIPoSValidatorManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPoSValidatorManagerFilterer{contract: contract}, nil
}

// bindIPoSValidatorManager binds a generic wrapper to an already deployed contract.
func bindIPoSValidatorManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPoSValidatorManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoSValidatorManager *IPoSValidatorManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoSValidatorManager.Contract.IPoSValidatorManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoSValidatorManager *IPoSValidatorManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.IPoSValidatorManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoSValidatorManager *IPoSValidatorManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.IPoSValidatorManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoSValidatorManager *IPoSValidatorManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoSValidatorManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.contract.Transact(opts, method, params...)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactor) ClaimDelegationFees(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _IPoSValidatorManager.contract.Transact(opts, "claimDelegationFees", validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.ClaimDelegationFees(&_IPoSValidatorManager.TransactOpts, validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.ClaimDelegationFees(&_IPoSValidatorManager.TransactOpts, validationID)
}

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactor) ClaimValidationRewards(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.contract.Transact(opts, "claimValidationRewards", validationID, messageIndex)
}

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerSession) ClaimValidationRewards(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.ClaimValidationRewards(&_IPoSValidatorManager.TransactOpts, validationID, messageIndex)
}

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorSession) ClaimValidationRewards(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.ClaimValidationRewards(&_IPoSValidatorManager.TransactOpts, validationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.contract.Transact(opts, "completeDelegatorRegistration", delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.CompleteDelegatorRegistration(&_IPoSValidatorManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.CompleteDelegatorRegistration(&_IPoSValidatorManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactor) CompleteEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.contract.Transact(opts, "completeEndDelegation", delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.CompleteEndDelegation(&_IPoSValidatorManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.CompleteEndDelegation(&_IPoSValidatorManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactor) CompleteEndValidation(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.contract.Transact(opts, "completeEndValidation", messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.CompleteEndValidation(&_IPoSValidatorManager.TransactOpts, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.CompleteEndValidation(&_IPoSValidatorManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.CompleteValidatorRegistration(&_IPoSValidatorManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.CompleteValidatorRegistration(&_IPoSValidatorManager.TransactOpts, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactor) ForceInitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.contract.Transact(opts, "forceInitializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.ForceInitializeEndDelegation(&_IPoSValidatorManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.ForceInitializeEndDelegation(&_IPoSValidatorManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactor) ForceInitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.contract.Transact(opts, "forceInitializeEndValidation", validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.ForceInitializeEndValidation(&_IPoSValidatorManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.ForceInitializeEndValidation(&_IPoSValidatorManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactor) InitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.contract.Transact(opts, "initializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.InitializeEndDelegation(&_IPoSValidatorManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.InitializeEndDelegation(&_IPoSValidatorManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactor) InitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.contract.Transact(opts, "initializeEndValidation", validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.InitializeEndValidation(&_IPoSValidatorManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.InitializeEndValidation(&_IPoSValidatorManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messsageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactor) InitializeValidatorSet(opts *bind.TransactOpts, subnetConversionData SubnetConversionData, messsageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.contract.Transact(opts, "initializeValidatorSet", subnetConversionData, messsageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messsageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerSession) InitializeValidatorSet(subnetConversionData SubnetConversionData, messsageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.InitializeValidatorSet(&_IPoSValidatorManager.TransactOpts, subnetConversionData, messsageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messsageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorSession) InitializeValidatorSet(subnetConversionData SubnetConversionData, messsageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.InitializeValidatorSet(&_IPoSValidatorManager.TransactOpts, subnetConversionData, messsageIndex)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactor) ResendEndValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _IPoSValidatorManager.contract.Transact(opts, "resendEndValidatorMessage", validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.ResendEndValidatorMessage(&_IPoSValidatorManager.TransactOpts, validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.ResendEndValidatorMessage(&_IPoSValidatorManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactor) ResendRegisterValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _IPoSValidatorManager.contract.Transact(opts, "resendRegisterValidatorMessage", validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.ResendRegisterValidatorMessage(&_IPoSValidatorManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.ResendRegisterValidatorMessage(&_IPoSValidatorManager.TransactOpts, validationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactor) ResendUpdateDelegation(opts *bind.TransactOpts, delegationID [32]byte) (*types.Transaction, error) {
	return _IPoSValidatorManager.contract.Transact(opts, "resendUpdateDelegation", delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.ResendUpdateDelegation(&_IPoSValidatorManager.TransactOpts, delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.ResendUpdateDelegation(&_IPoSValidatorManager.TransactOpts, delegationID)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactor) SubmitUptimeProof(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.contract.Transact(opts, "submitUptimeProof", validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.SubmitUptimeProof(&_IPoSValidatorManager.TransactOpts, validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_IPoSValidatorManager *IPoSValidatorManagerTransactorSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSValidatorManager.Contract.SubmitUptimeProof(&_IPoSValidatorManager.TransactOpts, validationID, messageIndex)
}

// IPoSValidatorManagerDelegationEndedIterator is returned from FilterDelegationEnded and is used to iterate over the raw logs and unpacked data for DelegationEnded events raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerDelegationEndedIterator struct {
	Event *IPoSValidatorManagerDelegationEnded // Event containing the contract specifics and raw log

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
func (it *IPoSValidatorManagerDelegationEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSValidatorManagerDelegationEnded)
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
		it.Event = new(IPoSValidatorManagerDelegationEnded)
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
func (it *IPoSValidatorManagerDelegationEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSValidatorManagerDelegationEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSValidatorManagerDelegationEnded represents a DelegationEnded event raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerDelegationEnded struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Rewards      *big.Int
	Fees         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegationEnded is a free log retrieval operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) FilterDelegationEnded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*IPoSValidatorManagerDelegationEndedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.FilterLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &IPoSValidatorManagerDelegationEndedIterator{contract: _IPoSValidatorManager.contract, event: "DelegationEnded", logs: logs, sub: sub}, nil
}

// WatchDelegationEnded is a free log subscription operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) WatchDelegationEnded(opts *bind.WatchOpts, sink chan<- *IPoSValidatorManagerDelegationEnded, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.WatchLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSValidatorManagerDelegationEnded)
				if err := _IPoSValidatorManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
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

// ParseDelegationEnded is a log parse operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) ParseDelegationEnded(log types.Log) (*IPoSValidatorManagerDelegationEnded, error) {
	event := new(IPoSValidatorManagerDelegationEnded)
	if err := _IPoSValidatorManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoSValidatorManagerDelegatorAddedIterator is returned from FilterDelegatorAdded and is used to iterate over the raw logs and unpacked data for DelegatorAdded events raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerDelegatorAddedIterator struct {
	Event *IPoSValidatorManagerDelegatorAdded // Event containing the contract specifics and raw log

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
func (it *IPoSValidatorManagerDelegatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSValidatorManagerDelegatorAdded)
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
		it.Event = new(IPoSValidatorManagerDelegatorAdded)
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
func (it *IPoSValidatorManagerDelegatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSValidatorManagerDelegatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSValidatorManagerDelegatorAdded represents a DelegatorAdded event raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerDelegatorAdded struct {
	DelegationID       [32]byte
	ValidationID       [32]byte
	DelegatorAddress   common.Address
	Nonce              uint64
	ValidatorWeight    uint64
	DelegatorWeight    uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDelegatorAdded is a free log retrieval operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) FilterDelegatorAdded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (*IPoSValidatorManagerDelegatorAddedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.FilterLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPoSValidatorManagerDelegatorAddedIterator{contract: _IPoSValidatorManager.contract, event: "DelegatorAdded", logs: logs, sub: sub}, nil
}

// WatchDelegatorAdded is a free log subscription operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) WatchDelegatorAdded(opts *bind.WatchOpts, sink chan<- *IPoSValidatorManagerDelegatorAdded, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.WatchLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSValidatorManagerDelegatorAdded)
				if err := _IPoSValidatorManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
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

// ParseDelegatorAdded is a log parse operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) ParseDelegatorAdded(log types.Log) (*IPoSValidatorManagerDelegatorAdded, error) {
	event := new(IPoSValidatorManagerDelegatorAdded)
	if err := _IPoSValidatorManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoSValidatorManagerDelegatorRegisteredIterator is returned from FilterDelegatorRegistered and is used to iterate over the raw logs and unpacked data for DelegatorRegistered events raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerDelegatorRegisteredIterator struct {
	Event *IPoSValidatorManagerDelegatorRegistered // Event containing the contract specifics and raw log

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
func (it *IPoSValidatorManagerDelegatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSValidatorManagerDelegatorRegistered)
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
		it.Event = new(IPoSValidatorManagerDelegatorRegistered)
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
func (it *IPoSValidatorManagerDelegatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSValidatorManagerDelegatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSValidatorManagerDelegatorRegistered represents a DelegatorRegistered event raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerDelegatorRegistered struct {
	DelegationID [32]byte
	ValidationID [32]byte
	StartTime    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRegistered is a free log retrieval operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) FilterDelegatorRegistered(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*IPoSValidatorManagerDelegatorRegisteredIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.FilterLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &IPoSValidatorManagerDelegatorRegisteredIterator{contract: _IPoSValidatorManager.contract, event: "DelegatorRegistered", logs: logs, sub: sub}, nil
}

// WatchDelegatorRegistered is a free log subscription operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) WatchDelegatorRegistered(opts *bind.WatchOpts, sink chan<- *IPoSValidatorManagerDelegatorRegistered, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.WatchLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSValidatorManagerDelegatorRegistered)
				if err := _IPoSValidatorManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
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

// ParseDelegatorRegistered is a log parse operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) ParseDelegatorRegistered(log types.Log) (*IPoSValidatorManagerDelegatorRegistered, error) {
	event := new(IPoSValidatorManagerDelegatorRegistered)
	if err := _IPoSValidatorManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoSValidatorManagerDelegatorRemovalInitializedIterator is returned from FilterDelegatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for DelegatorRemovalInitialized events raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerDelegatorRemovalInitializedIterator struct {
	Event *IPoSValidatorManagerDelegatorRemovalInitialized // Event containing the contract specifics and raw log

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
func (it *IPoSValidatorManagerDelegatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSValidatorManagerDelegatorRemovalInitialized)
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
		it.Event = new(IPoSValidatorManagerDelegatorRemovalInitialized)
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
func (it *IPoSValidatorManagerDelegatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSValidatorManagerDelegatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSValidatorManagerDelegatorRemovalInitialized represents a DelegatorRemovalInitialized event raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerDelegatorRemovalInitialized struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRemovalInitialized is a free log retrieval operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) FilterDelegatorRemovalInitialized(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*IPoSValidatorManagerDelegatorRemovalInitializedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.FilterLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &IPoSValidatorManagerDelegatorRemovalInitializedIterator{contract: _IPoSValidatorManager.contract, event: "DelegatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchDelegatorRemovalInitialized is a free log subscription operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) WatchDelegatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *IPoSValidatorManagerDelegatorRemovalInitialized, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.WatchLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSValidatorManagerDelegatorRemovalInitialized)
				if err := _IPoSValidatorManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
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

// ParseDelegatorRemovalInitialized is a log parse operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) ParseDelegatorRemovalInitialized(log types.Log) (*IPoSValidatorManagerDelegatorRemovalInitialized, error) {
	event := new(IPoSValidatorManagerDelegatorRemovalInitialized)
	if err := _IPoSValidatorManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoSValidatorManagerInitialValidatorCreatedIterator is returned from FilterInitialValidatorCreated and is used to iterate over the raw logs and unpacked data for InitialValidatorCreated events raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerInitialValidatorCreatedIterator struct {
	Event *IPoSValidatorManagerInitialValidatorCreated // Event containing the contract specifics and raw log

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
func (it *IPoSValidatorManagerInitialValidatorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSValidatorManagerInitialValidatorCreated)
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
		it.Event = new(IPoSValidatorManagerInitialValidatorCreated)
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
func (it *IPoSValidatorManagerInitialValidatorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSValidatorManagerInitialValidatorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSValidatorManagerInitialValidatorCreated represents a InitialValidatorCreated event raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerInitialValidatorCreated struct {
	ValidationID [32]byte
	NodeID       common.Hash
	Weight       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialValidatorCreated is a free log retrieval operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) FilterInitialValidatorCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte) (*IPoSValidatorManagerInitialValidatorCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.FilterLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &IPoSValidatorManagerInitialValidatorCreatedIterator{contract: _IPoSValidatorManager.contract, event: "InitialValidatorCreated", logs: logs, sub: sub}, nil
}

// WatchInitialValidatorCreated is a free log subscription operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) WatchInitialValidatorCreated(opts *bind.WatchOpts, sink chan<- *IPoSValidatorManagerInitialValidatorCreated, validationID [][32]byte, nodeID [][]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.WatchLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSValidatorManagerInitialValidatorCreated)
				if err := _IPoSValidatorManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
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

// ParseInitialValidatorCreated is a log parse operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) ParseInitialValidatorCreated(log types.Log) (*IPoSValidatorManagerInitialValidatorCreated, error) {
	event := new(IPoSValidatorManagerInitialValidatorCreated)
	if err := _IPoSValidatorManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoSValidatorManagerUptimeUpdatedIterator is returned from FilterUptimeUpdated and is used to iterate over the raw logs and unpacked data for UptimeUpdated events raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerUptimeUpdatedIterator struct {
	Event *IPoSValidatorManagerUptimeUpdated // Event containing the contract specifics and raw log

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
func (it *IPoSValidatorManagerUptimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSValidatorManagerUptimeUpdated)
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
		it.Event = new(IPoSValidatorManagerUptimeUpdated)
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
func (it *IPoSValidatorManagerUptimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSValidatorManagerUptimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSValidatorManagerUptimeUpdated represents a UptimeUpdated event raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerUptimeUpdated struct {
	ValidationID [32]byte
	Uptime       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUptimeUpdated is a free log retrieval operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) FilterUptimeUpdated(opts *bind.FilterOpts, validationID [][32]byte) (*IPoSValidatorManagerUptimeUpdatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.FilterLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &IPoSValidatorManagerUptimeUpdatedIterator{contract: _IPoSValidatorManager.contract, event: "UptimeUpdated", logs: logs, sub: sub}, nil
}

// WatchUptimeUpdated is a free log subscription operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) WatchUptimeUpdated(opts *bind.WatchOpts, sink chan<- *IPoSValidatorManagerUptimeUpdated, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.WatchLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSValidatorManagerUptimeUpdated)
				if err := _IPoSValidatorManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
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

// ParseUptimeUpdated is a log parse operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) ParseUptimeUpdated(log types.Log) (*IPoSValidatorManagerUptimeUpdated, error) {
	event := new(IPoSValidatorManagerUptimeUpdated)
	if err := _IPoSValidatorManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoSValidatorManagerValidationPeriodCreatedIterator is returned from FilterValidationPeriodCreated and is used to iterate over the raw logs and unpacked data for ValidationPeriodCreated events raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerValidationPeriodCreatedIterator struct {
	Event *IPoSValidatorManagerValidationPeriodCreated // Event containing the contract specifics and raw log

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
func (it *IPoSValidatorManagerValidationPeriodCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSValidatorManagerValidationPeriodCreated)
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
		it.Event = new(IPoSValidatorManagerValidationPeriodCreated)
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
func (it *IPoSValidatorManagerValidationPeriodCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSValidatorManagerValidationPeriodCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSValidatorManagerValidationPeriodCreated represents a ValidationPeriodCreated event raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerValidationPeriodCreated struct {
	ValidationID                [32]byte
	NodeID                      common.Hash
	RegisterValidationMessageID [32]byte
	Weight                      *big.Int
	RegistrationExpiry          uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodCreated is a free log retrieval operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) FilterValidationPeriodCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (*IPoSValidatorManagerValidationPeriodCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.FilterLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &IPoSValidatorManagerValidationPeriodCreatedIterator{contract: _IPoSValidatorManager.contract, event: "ValidationPeriodCreated", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodCreated is a free log subscription operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) WatchValidationPeriodCreated(opts *bind.WatchOpts, sink chan<- *IPoSValidatorManagerValidationPeriodCreated, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.WatchLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSValidatorManagerValidationPeriodCreated)
				if err := _IPoSValidatorManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
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

// ParseValidationPeriodCreated is a log parse operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) ParseValidationPeriodCreated(log types.Log) (*IPoSValidatorManagerValidationPeriodCreated, error) {
	event := new(IPoSValidatorManagerValidationPeriodCreated)
	if err := _IPoSValidatorManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoSValidatorManagerValidationPeriodEndedIterator is returned from FilterValidationPeriodEnded and is used to iterate over the raw logs and unpacked data for ValidationPeriodEnded events raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerValidationPeriodEndedIterator struct {
	Event *IPoSValidatorManagerValidationPeriodEnded // Event containing the contract specifics and raw log

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
func (it *IPoSValidatorManagerValidationPeriodEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSValidatorManagerValidationPeriodEnded)
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
		it.Event = new(IPoSValidatorManagerValidationPeriodEnded)
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
func (it *IPoSValidatorManagerValidationPeriodEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSValidatorManagerValidationPeriodEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSValidatorManagerValidationPeriodEnded represents a ValidationPeriodEnded event raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerValidationPeriodEnded struct {
	ValidationID [32]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodEnded is a free log retrieval operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) FilterValidationPeriodEnded(opts *bind.FilterOpts, validationID [][32]byte, status []uint8) (*IPoSValidatorManagerValidationPeriodEndedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.FilterLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &IPoSValidatorManagerValidationPeriodEndedIterator{contract: _IPoSValidatorManager.contract, event: "ValidationPeriodEnded", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodEnded is a free log subscription operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) WatchValidationPeriodEnded(opts *bind.WatchOpts, sink chan<- *IPoSValidatorManagerValidationPeriodEnded, validationID [][32]byte, status []uint8) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.WatchLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSValidatorManagerValidationPeriodEnded)
				if err := _IPoSValidatorManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
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

// ParseValidationPeriodEnded is a log parse operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) ParseValidationPeriodEnded(log types.Log) (*IPoSValidatorManagerValidationPeriodEnded, error) {
	event := new(IPoSValidatorManagerValidationPeriodEnded)
	if err := _IPoSValidatorManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoSValidatorManagerValidationPeriodRegisteredIterator is returned from FilterValidationPeriodRegistered and is used to iterate over the raw logs and unpacked data for ValidationPeriodRegistered events raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerValidationPeriodRegisteredIterator struct {
	Event *IPoSValidatorManagerValidationPeriodRegistered // Event containing the contract specifics and raw log

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
func (it *IPoSValidatorManagerValidationPeriodRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSValidatorManagerValidationPeriodRegistered)
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
		it.Event = new(IPoSValidatorManagerValidationPeriodRegistered)
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
func (it *IPoSValidatorManagerValidationPeriodRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSValidatorManagerValidationPeriodRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSValidatorManagerValidationPeriodRegistered represents a ValidationPeriodRegistered event raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerValidationPeriodRegistered struct {
	ValidationID [32]byte
	Weight       *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodRegistered is a free log retrieval operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) FilterValidationPeriodRegistered(opts *bind.FilterOpts, validationID [][32]byte) (*IPoSValidatorManagerValidationPeriodRegisteredIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.FilterLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &IPoSValidatorManagerValidationPeriodRegisteredIterator{contract: _IPoSValidatorManager.contract, event: "ValidationPeriodRegistered", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodRegistered is a free log subscription operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) WatchValidationPeriodRegistered(opts *bind.WatchOpts, sink chan<- *IPoSValidatorManagerValidationPeriodRegistered, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.WatchLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSValidatorManagerValidationPeriodRegistered)
				if err := _IPoSValidatorManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
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

// ParseValidationPeriodRegistered is a log parse operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) ParseValidationPeriodRegistered(log types.Log) (*IPoSValidatorManagerValidationPeriodRegistered, error) {
	event := new(IPoSValidatorManagerValidationPeriodRegistered)
	if err := _IPoSValidatorManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoSValidatorManagerValidationRewardsClaimedIterator is returned from FilterValidationRewardsClaimed and is used to iterate over the raw logs and unpacked data for ValidationRewardsClaimed events raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerValidationRewardsClaimedIterator struct {
	Event *IPoSValidatorManagerValidationRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *IPoSValidatorManagerValidationRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSValidatorManagerValidationRewardsClaimed)
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
		it.Event = new(IPoSValidatorManagerValidationRewardsClaimed)
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
func (it *IPoSValidatorManagerValidationRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSValidatorManagerValidationRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSValidatorManagerValidationRewardsClaimed represents a ValidationRewardsClaimed event raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerValidationRewardsClaimed struct {
	ValidationID [32]byte
	Reward       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationRewardsClaimed is a free log retrieval operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) FilterValidationRewardsClaimed(opts *bind.FilterOpts, validationID [][32]byte, reward []*big.Int) (*IPoSValidatorManagerValidationRewardsClaimedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.FilterLogs(opts, "ValidationRewardsClaimed", validationIDRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &IPoSValidatorManagerValidationRewardsClaimedIterator{contract: _IPoSValidatorManager.contract, event: "ValidationRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchValidationRewardsClaimed is a free log subscription operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) WatchValidationRewardsClaimed(opts *bind.WatchOpts, sink chan<- *IPoSValidatorManagerValidationRewardsClaimed, validationID [][32]byte, reward []*big.Int) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.WatchLogs(opts, "ValidationRewardsClaimed", validationIDRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSValidatorManagerValidationRewardsClaimed)
				if err := _IPoSValidatorManager.contract.UnpackLog(event, "ValidationRewardsClaimed", log); err != nil {
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

// ParseValidationRewardsClaimed is a log parse operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) ParseValidationRewardsClaimed(log types.Log) (*IPoSValidatorManagerValidationRewardsClaimed, error) {
	event := new(IPoSValidatorManagerValidationRewardsClaimed)
	if err := _IPoSValidatorManager.contract.UnpackLog(event, "ValidationRewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoSValidatorManagerValidatorRemovalInitializedIterator is returned from FilterValidatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for ValidatorRemovalInitialized events raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerValidatorRemovalInitializedIterator struct {
	Event *IPoSValidatorManagerValidatorRemovalInitialized // Event containing the contract specifics and raw log

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
func (it *IPoSValidatorManagerValidatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSValidatorManagerValidatorRemovalInitialized)
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
		it.Event = new(IPoSValidatorManagerValidatorRemovalInitialized)
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
func (it *IPoSValidatorManagerValidatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSValidatorManagerValidatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSValidatorManagerValidatorRemovalInitialized represents a ValidatorRemovalInitialized event raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerValidatorRemovalInitialized struct {
	ValidationID       [32]byte
	SetWeightMessageID [32]byte
	Weight             *big.Int
	EndTime            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemovalInitialized is a free log retrieval operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) FilterValidatorRemovalInitialized(opts *bind.FilterOpts, validationID [][32]byte, setWeightMessageID [][32]byte) (*IPoSValidatorManagerValidatorRemovalInitializedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.FilterLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &IPoSValidatorManagerValidatorRemovalInitializedIterator{contract: _IPoSValidatorManager.contract, event: "ValidatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchValidatorRemovalInitialized is a free log subscription operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) WatchValidatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *IPoSValidatorManagerValidatorRemovalInitialized, validationID [][32]byte, setWeightMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.WatchLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSValidatorManagerValidatorRemovalInitialized)
				if err := _IPoSValidatorManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
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

// ParseValidatorRemovalInitialized is a log parse operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) ParseValidatorRemovalInitialized(log types.Log) (*IPoSValidatorManagerValidatorRemovalInitialized, error) {
	event := new(IPoSValidatorManagerValidatorRemovalInitialized)
	if err := _IPoSValidatorManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoSValidatorManagerValidatorWeightUpdateIterator is returned from FilterValidatorWeightUpdate and is used to iterate over the raw logs and unpacked data for ValidatorWeightUpdate events raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerValidatorWeightUpdateIterator struct {
	Event *IPoSValidatorManagerValidatorWeightUpdate // Event containing the contract specifics and raw log

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
func (it *IPoSValidatorManagerValidatorWeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSValidatorManagerValidatorWeightUpdate)
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
		it.Event = new(IPoSValidatorManagerValidatorWeightUpdate)
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
func (it *IPoSValidatorManagerValidatorWeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSValidatorManagerValidatorWeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSValidatorManagerValidatorWeightUpdate represents a ValidatorWeightUpdate event raised by the IPoSValidatorManager contract.
type IPoSValidatorManagerValidatorWeightUpdate struct {
	ValidationID       [32]byte
	Nonce              uint64
	ValidatorWeight    uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorWeightUpdate is a free log retrieval operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) FilterValidatorWeightUpdate(opts *bind.FilterOpts, validationID [][32]byte, nonce []uint64) (*IPoSValidatorManagerValidatorWeightUpdateIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.FilterLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &IPoSValidatorManagerValidatorWeightUpdateIterator{contract: _IPoSValidatorManager.contract, event: "ValidatorWeightUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorWeightUpdate is a free log subscription operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) WatchValidatorWeightUpdate(opts *bind.WatchOpts, sink chan<- *IPoSValidatorManagerValidatorWeightUpdate, validationID [][32]byte, nonce []uint64) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _IPoSValidatorManager.contract.WatchLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSValidatorManagerValidatorWeightUpdate)
				if err := _IPoSValidatorManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
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

// ParseValidatorWeightUpdate is a log parse operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_IPoSValidatorManager *IPoSValidatorManagerFilterer) ParseValidatorWeightUpdate(log types.Log) (*IPoSValidatorManagerValidatorWeightUpdate, error) {
	event := new(IPoSValidatorManagerValidatorWeightUpdate)
	if err := _IPoSValidatorManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRewardCalculatorMetaData contains all meta data concerning the IRewardCalculator contract.
var IRewardCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"validatorStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"stakingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"stakingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"uptimeSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endSupply\",\"type\":\"uint256\"}],\"name\":\"calculateReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IRewardCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use IRewardCalculatorMetaData.ABI instead.
var IRewardCalculatorABI = IRewardCalculatorMetaData.ABI

// IRewardCalculator is an auto generated Go binding around an Ethereum contract.
type IRewardCalculator struct {
	IRewardCalculatorCaller     // Read-only binding to the contract
	IRewardCalculatorTransactor // Write-only binding to the contract
	IRewardCalculatorFilterer   // Log filterer for contract events
}

// IRewardCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRewardCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRewardCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRewardCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRewardCalculatorSession struct {
	Contract     *IRewardCalculator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IRewardCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRewardCalculatorCallerSession struct {
	Contract *IRewardCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IRewardCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRewardCalculatorTransactorSession struct {
	Contract     *IRewardCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IRewardCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRewardCalculatorRaw struct {
	Contract *IRewardCalculator // Generic contract binding to access the raw methods on
}

// IRewardCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRewardCalculatorCallerRaw struct {
	Contract *IRewardCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// IRewardCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRewardCalculatorTransactorRaw struct {
	Contract *IRewardCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRewardCalculator creates a new instance of IRewardCalculator, bound to a specific deployed contract.
func NewIRewardCalculator(address common.Address, backend bind.ContractBackend) (*IRewardCalculator, error) {
	contract, err := bindIRewardCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRewardCalculator{IRewardCalculatorCaller: IRewardCalculatorCaller{contract: contract}, IRewardCalculatorTransactor: IRewardCalculatorTransactor{contract: contract}, IRewardCalculatorFilterer: IRewardCalculatorFilterer{contract: contract}}, nil
}

// NewIRewardCalculatorCaller creates a new read-only instance of IRewardCalculator, bound to a specific deployed contract.
func NewIRewardCalculatorCaller(address common.Address, caller bind.ContractCaller) (*IRewardCalculatorCaller, error) {
	contract, err := bindIRewardCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardCalculatorCaller{contract: contract}, nil
}

// NewIRewardCalculatorTransactor creates a new write-only instance of IRewardCalculator, bound to a specific deployed contract.
func NewIRewardCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IRewardCalculatorTransactor, error) {
	contract, err := bindIRewardCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardCalculatorTransactor{contract: contract}, nil
}

// NewIRewardCalculatorFilterer creates a new log filterer instance of IRewardCalculator, bound to a specific deployed contract.
func NewIRewardCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IRewardCalculatorFilterer, error) {
	contract, err := bindIRewardCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRewardCalculatorFilterer{contract: contract}, nil
}

// bindIRewardCalculator binds a generic wrapper to an already deployed contract.
func bindIRewardCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IRewardCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewardCalculator *IRewardCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewardCalculator.Contract.IRewardCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewardCalculator *IRewardCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewardCalculator.Contract.IRewardCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewardCalculator *IRewardCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewardCalculator.Contract.IRewardCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewardCalculator *IRewardCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewardCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewardCalculator *IRewardCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewardCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewardCalculator *IRewardCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewardCalculator.Contract.contract.Transact(opts, method, params...)
}

// CalculateReward is a free data retrieval call binding the contract method 0x778c06b5.
//
// Solidity: function calculateReward(uint256 stakeAmount, uint64 validatorStartTime, uint64 stakingStartTime, uint64 stakingEndTime, uint64 uptimeSeconds, uint256 initialSupply, uint256 endSupply) view returns(uint256)
func (_IRewardCalculator *IRewardCalculatorCaller) CalculateReward(opts *bind.CallOpts, stakeAmount *big.Int, validatorStartTime uint64, stakingStartTime uint64, stakingEndTime uint64, uptimeSeconds uint64, initialSupply *big.Int, endSupply *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IRewardCalculator.contract.Call(opts, &out, "calculateReward", stakeAmount, validatorStartTime, stakingStartTime, stakingEndTime, uptimeSeconds, initialSupply, endSupply)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateReward is a free data retrieval call binding the contract method 0x778c06b5.
//
// Solidity: function calculateReward(uint256 stakeAmount, uint64 validatorStartTime, uint64 stakingStartTime, uint64 stakingEndTime, uint64 uptimeSeconds, uint256 initialSupply, uint256 endSupply) view returns(uint256)
func (_IRewardCalculator *IRewardCalculatorSession) CalculateReward(stakeAmount *big.Int, validatorStartTime uint64, stakingStartTime uint64, stakingEndTime uint64, uptimeSeconds uint64, initialSupply *big.Int, endSupply *big.Int) (*big.Int, error) {
	return _IRewardCalculator.Contract.CalculateReward(&_IRewardCalculator.CallOpts, stakeAmount, validatorStartTime, stakingStartTime, stakingEndTime, uptimeSeconds, initialSupply, endSupply)
}

// CalculateReward is a free data retrieval call binding the contract method 0x778c06b5.
//
// Solidity: function calculateReward(uint256 stakeAmount, uint64 validatorStartTime, uint64 stakingStartTime, uint64 stakingEndTime, uint64 uptimeSeconds, uint256 initialSupply, uint256 endSupply) view returns(uint256)
func (_IRewardCalculator *IRewardCalculatorCallerSession) CalculateReward(stakeAmount *big.Int, validatorStartTime uint64, stakingStartTime uint64, stakingEndTime uint64, uptimeSeconds uint64, initialSupply *big.Int, endSupply *big.Int) (*big.Int, error) {
	return _IRewardCalculator.Contract.CalculateReward(&_IRewardCalculator.CallOpts, stakeAmount, validatorStartTime, stakingStartTime, stakingEndTime, uptimeSeconds, initialSupply, endSupply)
}

// IValidatorManagerMetaData contains all meta data concerning the IValidatorManager contract.
var IValidatorManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"InitialValidatorCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"ValidationPeriodCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ValidationPeriodEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidationPeriodRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"ValidatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorWeightUpdate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structSubnetConversionData\",\"name\":\"subnetConversionData\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"messsageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendEndValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendRegisterValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IValidatorManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IValidatorManagerMetaData.ABI instead.
var IValidatorManagerABI = IValidatorManagerMetaData.ABI

// IValidatorManager is an auto generated Go binding around an Ethereum contract.
type IValidatorManager struct {
	IValidatorManagerCaller     // Read-only binding to the contract
	IValidatorManagerTransactor // Write-only binding to the contract
	IValidatorManagerFilterer   // Log filterer for contract events
}

// IValidatorManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IValidatorManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValidatorManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IValidatorManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValidatorManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IValidatorManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValidatorManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IValidatorManagerSession struct {
	Contract     *IValidatorManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IValidatorManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IValidatorManagerCallerSession struct {
	Contract *IValidatorManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IValidatorManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IValidatorManagerTransactorSession struct {
	Contract     *IValidatorManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IValidatorManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IValidatorManagerRaw struct {
	Contract *IValidatorManager // Generic contract binding to access the raw methods on
}

// IValidatorManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IValidatorManagerCallerRaw struct {
	Contract *IValidatorManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IValidatorManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IValidatorManagerTransactorRaw struct {
	Contract *IValidatorManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIValidatorManager creates a new instance of IValidatorManager, bound to a specific deployed contract.
func NewIValidatorManager(address common.Address, backend bind.ContractBackend) (*IValidatorManager, error) {
	contract, err := bindIValidatorManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IValidatorManager{IValidatorManagerCaller: IValidatorManagerCaller{contract: contract}, IValidatorManagerTransactor: IValidatorManagerTransactor{contract: contract}, IValidatorManagerFilterer: IValidatorManagerFilterer{contract: contract}}, nil
}

// NewIValidatorManagerCaller creates a new read-only instance of IValidatorManager, bound to a specific deployed contract.
func NewIValidatorManagerCaller(address common.Address, caller bind.ContractCaller) (*IValidatorManagerCaller, error) {
	contract, err := bindIValidatorManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IValidatorManagerCaller{contract: contract}, nil
}

// NewIValidatorManagerTransactor creates a new write-only instance of IValidatorManager, bound to a specific deployed contract.
func NewIValidatorManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IValidatorManagerTransactor, error) {
	contract, err := bindIValidatorManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IValidatorManagerTransactor{contract: contract}, nil
}

// NewIValidatorManagerFilterer creates a new log filterer instance of IValidatorManager, bound to a specific deployed contract.
func NewIValidatorManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IValidatorManagerFilterer, error) {
	contract, err := bindIValidatorManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IValidatorManagerFilterer{contract: contract}, nil
}

// bindIValidatorManager binds a generic wrapper to an already deployed contract.
func bindIValidatorManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IValidatorManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IValidatorManager *IValidatorManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IValidatorManager.Contract.IValidatorManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IValidatorManager *IValidatorManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IValidatorManager.Contract.IValidatorManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IValidatorManager *IValidatorManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IValidatorManager.Contract.IValidatorManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IValidatorManager *IValidatorManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IValidatorManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IValidatorManager *IValidatorManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IValidatorManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IValidatorManager *IValidatorManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IValidatorManager.Contract.contract.Transact(opts, method, params...)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_IValidatorManager *IValidatorManagerTransactor) CompleteEndValidation(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _IValidatorManager.contract.Transact(opts, "completeEndValidation", messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_IValidatorManager *IValidatorManagerSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _IValidatorManager.Contract.CompleteEndValidation(&_IValidatorManager.TransactOpts, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_IValidatorManager *IValidatorManagerTransactorSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _IValidatorManager.Contract.CompleteEndValidation(&_IValidatorManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_IValidatorManager *IValidatorManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _IValidatorManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_IValidatorManager *IValidatorManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _IValidatorManager.Contract.CompleteValidatorRegistration(&_IValidatorManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_IValidatorManager *IValidatorManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _IValidatorManager.Contract.CompleteValidatorRegistration(&_IValidatorManager.TransactOpts, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messsageIndex) returns()
func (_IValidatorManager *IValidatorManagerTransactor) InitializeValidatorSet(opts *bind.TransactOpts, subnetConversionData SubnetConversionData, messsageIndex uint32) (*types.Transaction, error) {
	return _IValidatorManager.contract.Transact(opts, "initializeValidatorSet", subnetConversionData, messsageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messsageIndex) returns()
func (_IValidatorManager *IValidatorManagerSession) InitializeValidatorSet(subnetConversionData SubnetConversionData, messsageIndex uint32) (*types.Transaction, error) {
	return _IValidatorManager.Contract.InitializeValidatorSet(&_IValidatorManager.TransactOpts, subnetConversionData, messsageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messsageIndex) returns()
func (_IValidatorManager *IValidatorManagerTransactorSession) InitializeValidatorSet(subnetConversionData SubnetConversionData, messsageIndex uint32) (*types.Transaction, error) {
	return _IValidatorManager.Contract.InitializeValidatorSet(&_IValidatorManager.TransactOpts, subnetConversionData, messsageIndex)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_IValidatorManager *IValidatorManagerTransactor) ResendEndValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _IValidatorManager.contract.Transact(opts, "resendEndValidatorMessage", validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_IValidatorManager *IValidatorManagerSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _IValidatorManager.Contract.ResendEndValidatorMessage(&_IValidatorManager.TransactOpts, validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_IValidatorManager *IValidatorManagerTransactorSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _IValidatorManager.Contract.ResendEndValidatorMessage(&_IValidatorManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_IValidatorManager *IValidatorManagerTransactor) ResendRegisterValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _IValidatorManager.contract.Transact(opts, "resendRegisterValidatorMessage", validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_IValidatorManager *IValidatorManagerSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _IValidatorManager.Contract.ResendRegisterValidatorMessage(&_IValidatorManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_IValidatorManager *IValidatorManagerTransactorSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _IValidatorManager.Contract.ResendRegisterValidatorMessage(&_IValidatorManager.TransactOpts, validationID)
}

// IValidatorManagerInitialValidatorCreatedIterator is returned from FilterInitialValidatorCreated and is used to iterate over the raw logs and unpacked data for InitialValidatorCreated events raised by the IValidatorManager contract.
type IValidatorManagerInitialValidatorCreatedIterator struct {
	Event *IValidatorManagerInitialValidatorCreated // Event containing the contract specifics and raw log

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
func (it *IValidatorManagerInitialValidatorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IValidatorManagerInitialValidatorCreated)
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
		it.Event = new(IValidatorManagerInitialValidatorCreated)
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
func (it *IValidatorManagerInitialValidatorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IValidatorManagerInitialValidatorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IValidatorManagerInitialValidatorCreated represents a InitialValidatorCreated event raised by the IValidatorManager contract.
type IValidatorManagerInitialValidatorCreated struct {
	ValidationID [32]byte
	NodeID       common.Hash
	Weight       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialValidatorCreated is a free log retrieval operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_IValidatorManager *IValidatorManagerFilterer) FilterInitialValidatorCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte) (*IValidatorManagerInitialValidatorCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _IValidatorManager.contract.FilterLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &IValidatorManagerInitialValidatorCreatedIterator{contract: _IValidatorManager.contract, event: "InitialValidatorCreated", logs: logs, sub: sub}, nil
}

// WatchInitialValidatorCreated is a free log subscription operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_IValidatorManager *IValidatorManagerFilterer) WatchInitialValidatorCreated(opts *bind.WatchOpts, sink chan<- *IValidatorManagerInitialValidatorCreated, validationID [][32]byte, nodeID [][]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _IValidatorManager.contract.WatchLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IValidatorManagerInitialValidatorCreated)
				if err := _IValidatorManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
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

// ParseInitialValidatorCreated is a log parse operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_IValidatorManager *IValidatorManagerFilterer) ParseInitialValidatorCreated(log types.Log) (*IValidatorManagerInitialValidatorCreated, error) {
	event := new(IValidatorManagerInitialValidatorCreated)
	if err := _IValidatorManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IValidatorManagerValidationPeriodCreatedIterator is returned from FilterValidationPeriodCreated and is used to iterate over the raw logs and unpacked data for ValidationPeriodCreated events raised by the IValidatorManager contract.
type IValidatorManagerValidationPeriodCreatedIterator struct {
	Event *IValidatorManagerValidationPeriodCreated // Event containing the contract specifics and raw log

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
func (it *IValidatorManagerValidationPeriodCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IValidatorManagerValidationPeriodCreated)
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
		it.Event = new(IValidatorManagerValidationPeriodCreated)
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
func (it *IValidatorManagerValidationPeriodCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IValidatorManagerValidationPeriodCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IValidatorManagerValidationPeriodCreated represents a ValidationPeriodCreated event raised by the IValidatorManager contract.
type IValidatorManagerValidationPeriodCreated struct {
	ValidationID                [32]byte
	NodeID                      common.Hash
	RegisterValidationMessageID [32]byte
	Weight                      *big.Int
	RegistrationExpiry          uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodCreated is a free log retrieval operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_IValidatorManager *IValidatorManagerFilterer) FilterValidationPeriodCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (*IValidatorManagerValidationPeriodCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _IValidatorManager.contract.FilterLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &IValidatorManagerValidationPeriodCreatedIterator{contract: _IValidatorManager.contract, event: "ValidationPeriodCreated", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodCreated is a free log subscription operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_IValidatorManager *IValidatorManagerFilterer) WatchValidationPeriodCreated(opts *bind.WatchOpts, sink chan<- *IValidatorManagerValidationPeriodCreated, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _IValidatorManager.contract.WatchLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IValidatorManagerValidationPeriodCreated)
				if err := _IValidatorManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
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

// ParseValidationPeriodCreated is a log parse operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_IValidatorManager *IValidatorManagerFilterer) ParseValidationPeriodCreated(log types.Log) (*IValidatorManagerValidationPeriodCreated, error) {
	event := new(IValidatorManagerValidationPeriodCreated)
	if err := _IValidatorManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IValidatorManagerValidationPeriodEndedIterator is returned from FilterValidationPeriodEnded and is used to iterate over the raw logs and unpacked data for ValidationPeriodEnded events raised by the IValidatorManager contract.
type IValidatorManagerValidationPeriodEndedIterator struct {
	Event *IValidatorManagerValidationPeriodEnded // Event containing the contract specifics and raw log

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
func (it *IValidatorManagerValidationPeriodEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IValidatorManagerValidationPeriodEnded)
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
		it.Event = new(IValidatorManagerValidationPeriodEnded)
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
func (it *IValidatorManagerValidationPeriodEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IValidatorManagerValidationPeriodEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IValidatorManagerValidationPeriodEnded represents a ValidationPeriodEnded event raised by the IValidatorManager contract.
type IValidatorManagerValidationPeriodEnded struct {
	ValidationID [32]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodEnded is a free log retrieval operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_IValidatorManager *IValidatorManagerFilterer) FilterValidationPeriodEnded(opts *bind.FilterOpts, validationID [][32]byte, status []uint8) (*IValidatorManagerValidationPeriodEndedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _IValidatorManager.contract.FilterLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &IValidatorManagerValidationPeriodEndedIterator{contract: _IValidatorManager.contract, event: "ValidationPeriodEnded", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodEnded is a free log subscription operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_IValidatorManager *IValidatorManagerFilterer) WatchValidationPeriodEnded(opts *bind.WatchOpts, sink chan<- *IValidatorManagerValidationPeriodEnded, validationID [][32]byte, status []uint8) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _IValidatorManager.contract.WatchLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IValidatorManagerValidationPeriodEnded)
				if err := _IValidatorManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
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

// ParseValidationPeriodEnded is a log parse operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_IValidatorManager *IValidatorManagerFilterer) ParseValidationPeriodEnded(log types.Log) (*IValidatorManagerValidationPeriodEnded, error) {
	event := new(IValidatorManagerValidationPeriodEnded)
	if err := _IValidatorManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IValidatorManagerValidationPeriodRegisteredIterator is returned from FilterValidationPeriodRegistered and is used to iterate over the raw logs and unpacked data for ValidationPeriodRegistered events raised by the IValidatorManager contract.
type IValidatorManagerValidationPeriodRegisteredIterator struct {
	Event *IValidatorManagerValidationPeriodRegistered // Event containing the contract specifics and raw log

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
func (it *IValidatorManagerValidationPeriodRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IValidatorManagerValidationPeriodRegistered)
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
		it.Event = new(IValidatorManagerValidationPeriodRegistered)
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
func (it *IValidatorManagerValidationPeriodRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IValidatorManagerValidationPeriodRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IValidatorManagerValidationPeriodRegistered represents a ValidationPeriodRegistered event raised by the IValidatorManager contract.
type IValidatorManagerValidationPeriodRegistered struct {
	ValidationID [32]byte
	Weight       *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodRegistered is a free log retrieval operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_IValidatorManager *IValidatorManagerFilterer) FilterValidationPeriodRegistered(opts *bind.FilterOpts, validationID [][32]byte) (*IValidatorManagerValidationPeriodRegisteredIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IValidatorManager.contract.FilterLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &IValidatorManagerValidationPeriodRegisteredIterator{contract: _IValidatorManager.contract, event: "ValidationPeriodRegistered", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodRegistered is a free log subscription operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_IValidatorManager *IValidatorManagerFilterer) WatchValidationPeriodRegistered(opts *bind.WatchOpts, sink chan<- *IValidatorManagerValidationPeriodRegistered, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IValidatorManager.contract.WatchLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IValidatorManagerValidationPeriodRegistered)
				if err := _IValidatorManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
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

// ParseValidationPeriodRegistered is a log parse operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_IValidatorManager *IValidatorManagerFilterer) ParseValidationPeriodRegistered(log types.Log) (*IValidatorManagerValidationPeriodRegistered, error) {
	event := new(IValidatorManagerValidationPeriodRegistered)
	if err := _IValidatorManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IValidatorManagerValidatorRemovalInitializedIterator is returned from FilterValidatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for ValidatorRemovalInitialized events raised by the IValidatorManager contract.
type IValidatorManagerValidatorRemovalInitializedIterator struct {
	Event *IValidatorManagerValidatorRemovalInitialized // Event containing the contract specifics and raw log

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
func (it *IValidatorManagerValidatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IValidatorManagerValidatorRemovalInitialized)
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
		it.Event = new(IValidatorManagerValidatorRemovalInitialized)
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
func (it *IValidatorManagerValidatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IValidatorManagerValidatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IValidatorManagerValidatorRemovalInitialized represents a ValidatorRemovalInitialized event raised by the IValidatorManager contract.
type IValidatorManagerValidatorRemovalInitialized struct {
	ValidationID       [32]byte
	SetWeightMessageID [32]byte
	Weight             *big.Int
	EndTime            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemovalInitialized is a free log retrieval operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_IValidatorManager *IValidatorManagerFilterer) FilterValidatorRemovalInitialized(opts *bind.FilterOpts, validationID [][32]byte, setWeightMessageID [][32]byte) (*IValidatorManagerValidatorRemovalInitializedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _IValidatorManager.contract.FilterLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &IValidatorManagerValidatorRemovalInitializedIterator{contract: _IValidatorManager.contract, event: "ValidatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchValidatorRemovalInitialized is a free log subscription operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_IValidatorManager *IValidatorManagerFilterer) WatchValidatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *IValidatorManagerValidatorRemovalInitialized, validationID [][32]byte, setWeightMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _IValidatorManager.contract.WatchLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IValidatorManagerValidatorRemovalInitialized)
				if err := _IValidatorManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
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

// ParseValidatorRemovalInitialized is a log parse operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_IValidatorManager *IValidatorManagerFilterer) ParseValidatorRemovalInitialized(log types.Log) (*IValidatorManagerValidatorRemovalInitialized, error) {
	event := new(IValidatorManagerValidatorRemovalInitialized)
	if err := _IValidatorManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IValidatorManagerValidatorWeightUpdateIterator is returned from FilterValidatorWeightUpdate and is used to iterate over the raw logs and unpacked data for ValidatorWeightUpdate events raised by the IValidatorManager contract.
type IValidatorManagerValidatorWeightUpdateIterator struct {
	Event *IValidatorManagerValidatorWeightUpdate // Event containing the contract specifics and raw log

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
func (it *IValidatorManagerValidatorWeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IValidatorManagerValidatorWeightUpdate)
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
		it.Event = new(IValidatorManagerValidatorWeightUpdate)
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
func (it *IValidatorManagerValidatorWeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IValidatorManagerValidatorWeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IValidatorManagerValidatorWeightUpdate represents a ValidatorWeightUpdate event raised by the IValidatorManager contract.
type IValidatorManagerValidatorWeightUpdate struct {
	ValidationID       [32]byte
	Nonce              uint64
	ValidatorWeight    uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorWeightUpdate is a free log retrieval operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_IValidatorManager *IValidatorManagerFilterer) FilterValidatorWeightUpdate(opts *bind.FilterOpts, validationID [][32]byte, nonce []uint64) (*IValidatorManagerValidatorWeightUpdateIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _IValidatorManager.contract.FilterLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &IValidatorManagerValidatorWeightUpdateIterator{contract: _IValidatorManager.contract, event: "ValidatorWeightUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorWeightUpdate is a free log subscription operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_IValidatorManager *IValidatorManagerFilterer) WatchValidatorWeightUpdate(opts *bind.WatchOpts, sink chan<- *IValidatorManagerValidatorWeightUpdate, validationID [][32]byte, nonce []uint64) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _IValidatorManager.contract.WatchLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IValidatorManagerValidatorWeightUpdate)
				if err := _IValidatorManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
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

// ParseValidatorWeightUpdate is a log parse operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_IValidatorManager *IValidatorManagerFilterer) ParseValidatorWeightUpdate(log types.Log) (*IValidatorManagerValidatorWeightUpdate, error) {
	event := new(IValidatorManagerValidatorWeightUpdate)
	if err := _IValidatorManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWarpMessengerMetaData contains all meta data concerning the IWarpMessenger contract.
var IWarpMessengerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"SendWarpMessage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"getVerifiedWarpBlockHash\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"internalType\":\"structWarpBlockHash\",\"name\":\"warpBlockHash\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"getVerifiedWarpMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structWarpMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"sendWarpMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IWarpMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use IWarpMessengerMetaData.ABI instead.
var IWarpMessengerABI = IWarpMessengerMetaData.ABI

// IWarpMessenger is an auto generated Go binding around an Ethereum contract.
type IWarpMessenger struct {
	IWarpMessengerCaller     // Read-only binding to the contract
	IWarpMessengerTransactor // Write-only binding to the contract
	IWarpMessengerFilterer   // Log filterer for contract events
}

// IWarpMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWarpMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWarpMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWarpMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWarpMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWarpMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWarpMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWarpMessengerSession struct {
	Contract     *IWarpMessenger   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWarpMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWarpMessengerCallerSession struct {
	Contract *IWarpMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IWarpMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWarpMessengerTransactorSession struct {
	Contract     *IWarpMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IWarpMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWarpMessengerRaw struct {
	Contract *IWarpMessenger // Generic contract binding to access the raw methods on
}

// IWarpMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWarpMessengerCallerRaw struct {
	Contract *IWarpMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// IWarpMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWarpMessengerTransactorRaw struct {
	Contract *IWarpMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWarpMessenger creates a new instance of IWarpMessenger, bound to a specific deployed contract.
func NewIWarpMessenger(address common.Address, backend bind.ContractBackend) (*IWarpMessenger, error) {
	contract, err := bindIWarpMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWarpMessenger{IWarpMessengerCaller: IWarpMessengerCaller{contract: contract}, IWarpMessengerTransactor: IWarpMessengerTransactor{contract: contract}, IWarpMessengerFilterer: IWarpMessengerFilterer{contract: contract}}, nil
}

// NewIWarpMessengerCaller creates a new read-only instance of IWarpMessenger, bound to a specific deployed contract.
func NewIWarpMessengerCaller(address common.Address, caller bind.ContractCaller) (*IWarpMessengerCaller, error) {
	contract, err := bindIWarpMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWarpMessengerCaller{contract: contract}, nil
}

// NewIWarpMessengerTransactor creates a new write-only instance of IWarpMessenger, bound to a specific deployed contract.
func NewIWarpMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*IWarpMessengerTransactor, error) {
	contract, err := bindIWarpMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWarpMessengerTransactor{contract: contract}, nil
}

// NewIWarpMessengerFilterer creates a new log filterer instance of IWarpMessenger, bound to a specific deployed contract.
func NewIWarpMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*IWarpMessengerFilterer, error) {
	contract, err := bindIWarpMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWarpMessengerFilterer{contract: contract}, nil
}

// bindIWarpMessenger binds a generic wrapper to an already deployed contract.
func bindIWarpMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IWarpMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWarpMessenger *IWarpMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWarpMessenger.Contract.IWarpMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWarpMessenger *IWarpMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWarpMessenger.Contract.IWarpMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWarpMessenger *IWarpMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWarpMessenger.Contract.IWarpMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWarpMessenger *IWarpMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWarpMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWarpMessenger *IWarpMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWarpMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWarpMessenger *IWarpMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWarpMessenger.Contract.contract.Transact(opts, method, params...)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32 blockchainID)
func (_IWarpMessenger *IWarpMessengerCaller) GetBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IWarpMessenger.contract.Call(opts, &out, "getBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32 blockchainID)
func (_IWarpMessenger *IWarpMessengerSession) GetBlockchainID() ([32]byte, error) {
	return _IWarpMessenger.Contract.GetBlockchainID(&_IWarpMessenger.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32 blockchainID)
func (_IWarpMessenger *IWarpMessengerCallerSession) GetBlockchainID() ([32]byte, error) {
	return _IWarpMessenger.Contract.GetBlockchainID(&_IWarpMessenger.CallOpts)
}

// GetVerifiedWarpBlockHash is a free data retrieval call binding the contract method 0xce7f5929.
//
// Solidity: function getVerifiedWarpBlockHash(uint32 index) view returns((bytes32,bytes32) warpBlockHash, bool valid)
func (_IWarpMessenger *IWarpMessengerCaller) GetVerifiedWarpBlockHash(opts *bind.CallOpts, index uint32) (struct {
	WarpBlockHash WarpBlockHash
	Valid         bool
}, error) {
	var out []interface{}
	err := _IWarpMessenger.contract.Call(opts, &out, "getVerifiedWarpBlockHash", index)

	outstruct := new(struct {
		WarpBlockHash WarpBlockHash
		Valid         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WarpBlockHash = *abi.ConvertType(out[0], new(WarpBlockHash)).(*WarpBlockHash)
	outstruct.Valid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetVerifiedWarpBlockHash is a free data retrieval call binding the contract method 0xce7f5929.
//
// Solidity: function getVerifiedWarpBlockHash(uint32 index) view returns((bytes32,bytes32) warpBlockHash, bool valid)
func (_IWarpMessenger *IWarpMessengerSession) GetVerifiedWarpBlockHash(index uint32) (struct {
	WarpBlockHash WarpBlockHash
	Valid         bool
}, error) {
	return _IWarpMessenger.Contract.GetVerifiedWarpBlockHash(&_IWarpMessenger.CallOpts, index)
}

// GetVerifiedWarpBlockHash is a free data retrieval call binding the contract method 0xce7f5929.
//
// Solidity: function getVerifiedWarpBlockHash(uint32 index) view returns((bytes32,bytes32) warpBlockHash, bool valid)
func (_IWarpMessenger *IWarpMessengerCallerSession) GetVerifiedWarpBlockHash(index uint32) (struct {
	WarpBlockHash WarpBlockHash
	Valid         bool
}, error) {
	return _IWarpMessenger.Contract.GetVerifiedWarpBlockHash(&_IWarpMessenger.CallOpts, index)
}

// GetVerifiedWarpMessage is a free data retrieval call binding the contract method 0x6f825350.
//
// Solidity: function getVerifiedWarpMessage(uint32 index) view returns((bytes32,address,bytes) message, bool valid)
func (_IWarpMessenger *IWarpMessengerCaller) GetVerifiedWarpMessage(opts *bind.CallOpts, index uint32) (struct {
	Message WarpMessage
	Valid   bool
}, error) {
	var out []interface{}
	err := _IWarpMessenger.contract.Call(opts, &out, "getVerifiedWarpMessage", index)

	outstruct := new(struct {
		Message WarpMessage
		Valid   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Message = *abi.ConvertType(out[0], new(WarpMessage)).(*WarpMessage)
	outstruct.Valid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetVerifiedWarpMessage is a free data retrieval call binding the contract method 0x6f825350.
//
// Solidity: function getVerifiedWarpMessage(uint32 index) view returns((bytes32,address,bytes) message, bool valid)
func (_IWarpMessenger *IWarpMessengerSession) GetVerifiedWarpMessage(index uint32) (struct {
	Message WarpMessage
	Valid   bool
}, error) {
	return _IWarpMessenger.Contract.GetVerifiedWarpMessage(&_IWarpMessenger.CallOpts, index)
}

// GetVerifiedWarpMessage is a free data retrieval call binding the contract method 0x6f825350.
//
// Solidity: function getVerifiedWarpMessage(uint32 index) view returns((bytes32,address,bytes) message, bool valid)
func (_IWarpMessenger *IWarpMessengerCallerSession) GetVerifiedWarpMessage(index uint32) (struct {
	Message WarpMessage
	Valid   bool
}, error) {
	return _IWarpMessenger.Contract.GetVerifiedWarpMessage(&_IWarpMessenger.CallOpts, index)
}

// SendWarpMessage is a paid mutator transaction binding the contract method 0xee5b48eb.
//
// Solidity: function sendWarpMessage(bytes payload) returns(bytes32 messageID)
func (_IWarpMessenger *IWarpMessengerTransactor) SendWarpMessage(opts *bind.TransactOpts, payload []byte) (*types.Transaction, error) {
	return _IWarpMessenger.contract.Transact(opts, "sendWarpMessage", payload)
}

// SendWarpMessage is a paid mutator transaction binding the contract method 0xee5b48eb.
//
// Solidity: function sendWarpMessage(bytes payload) returns(bytes32 messageID)
func (_IWarpMessenger *IWarpMessengerSession) SendWarpMessage(payload []byte) (*types.Transaction, error) {
	return _IWarpMessenger.Contract.SendWarpMessage(&_IWarpMessenger.TransactOpts, payload)
}

// SendWarpMessage is a paid mutator transaction binding the contract method 0xee5b48eb.
//
// Solidity: function sendWarpMessage(bytes payload) returns(bytes32 messageID)
func (_IWarpMessenger *IWarpMessengerTransactorSession) SendWarpMessage(payload []byte) (*types.Transaction, error) {
	return _IWarpMessenger.Contract.SendWarpMessage(&_IWarpMessenger.TransactOpts, payload)
}

// IWarpMessengerSendWarpMessageIterator is returned from FilterSendWarpMessage and is used to iterate over the raw logs and unpacked data for SendWarpMessage events raised by the IWarpMessenger contract.
type IWarpMessengerSendWarpMessageIterator struct {
	Event *IWarpMessengerSendWarpMessage // Event containing the contract specifics and raw log

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
func (it *IWarpMessengerSendWarpMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWarpMessengerSendWarpMessage)
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
		it.Event = new(IWarpMessengerSendWarpMessage)
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
func (it *IWarpMessengerSendWarpMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWarpMessengerSendWarpMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWarpMessengerSendWarpMessage represents a SendWarpMessage event raised by the IWarpMessenger contract.
type IWarpMessengerSendWarpMessage struct {
	Sender    common.Address
	MessageID [32]byte
	Message   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSendWarpMessage is a free log retrieval operation binding the contract event 0x56600c567728a800c0aa927500f831cb451df66a7af570eb4df4dfbf4674887d.
//
// Solidity: event SendWarpMessage(address indexed sender, bytes32 indexed messageID, bytes message)
func (_IWarpMessenger *IWarpMessengerFilterer) FilterSendWarpMessage(opts *bind.FilterOpts, sender []common.Address, messageID [][32]byte) (*IWarpMessengerSendWarpMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _IWarpMessenger.contract.FilterLogs(opts, "SendWarpMessage", senderRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &IWarpMessengerSendWarpMessageIterator{contract: _IWarpMessenger.contract, event: "SendWarpMessage", logs: logs, sub: sub}, nil
}

// WatchSendWarpMessage is a free log subscription operation binding the contract event 0x56600c567728a800c0aa927500f831cb451df66a7af570eb4df4dfbf4674887d.
//
// Solidity: event SendWarpMessage(address indexed sender, bytes32 indexed messageID, bytes message)
func (_IWarpMessenger *IWarpMessengerFilterer) WatchSendWarpMessage(opts *bind.WatchOpts, sink chan<- *IWarpMessengerSendWarpMessage, sender []common.Address, messageID [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _IWarpMessenger.contract.WatchLogs(opts, "SendWarpMessage", senderRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWarpMessengerSendWarpMessage)
				if err := _IWarpMessenger.contract.UnpackLog(event, "SendWarpMessage", log); err != nil {
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

// ParseSendWarpMessage is a log parse operation binding the contract event 0x56600c567728a800c0aa927500f831cb451df66a7af570eb4df4dfbf4674887d.
//
// Solidity: event SendWarpMessage(address indexed sender, bytes32 indexed messageID, bytes message)
func (_IWarpMessenger *IWarpMessengerFilterer) ParseSendWarpMessage(log types.Log) (*IWarpMessengerSendWarpMessage, error) {
	event := new(IWarpMessengerSendWarpMessage)
	if err := _IWarpMessenger.contract.UnpackLog(event, "SendWarpMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// InitializableABI is the input ABI used to generate the binding from.
// Deprecated: Use InitializableMetaData.ABI instead.
var InitializableABI = InitializableMetaData.ABI

// Initializable is an auto generated Go binding around an Ethereum contract.
type Initializable struct {
	InitializableCaller     // Read-only binding to the contract
	InitializableTransactor // Write-only binding to the contract
	InitializableFilterer   // Log filterer for contract events
}

// InitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableSession struct {
	Contract     *Initializable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableCallerSession struct {
	Contract *InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableTransactorSession struct {
	Contract     *InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableRaw struct {
	Contract *Initializable // Generic contract binding to access the raw methods on
}

// InitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableCallerRaw struct {
	Contract *InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableTransactorRaw struct {
	Contract *InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializable creates a new instance of Initializable, bound to a specific deployed contract.
func NewInitializable(address common.Address, backend bind.ContractBackend) (*Initializable, error) {
	contract, err := bindInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// NewInitializableCaller creates a new read-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableCaller(address common.Address, caller bind.ContractCaller) (*InitializableCaller, error) {
	contract, err := bindInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableCaller{contract: contract}, nil
}

// NewInitializableTransactor creates a new write-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableTransactor, error) {
	contract, err := bindInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableTransactor{contract: contract}, nil
}

// NewInitializableFilterer creates a new log filterer instance of Initializable, bound to a specific deployed contract.
func NewInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableFilterer, error) {
	contract, err := bindInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableFilterer{contract: contract}, nil
}

// bindInitializable binds a generic wrapper to an already deployed contract.
func bindInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InitializableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transact(opts, method, params...)
}

// InitializableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Initializable contract.
type InitializableInitializedIterator struct {
	Event *InitializableInitialized // Event containing the contract specifics and raw log

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
func (it *InitializableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitializableInitialized)
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
		it.Event = new(InitializableInitialized)
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
func (it *InitializableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitializableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitializableInitialized represents a Initialized event raised by the Initializable contract.
type InitializableInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Initializable *InitializableFilterer) FilterInitialized(opts *bind.FilterOpts) (*InitializableInitializedIterator, error) {

	logs, sub, err := _Initializable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &InitializableInitializedIterator{contract: _Initializable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Initializable *InitializableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *InitializableInitialized) (event.Subscription, error) {

	logs, sub, err := _Initializable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitializableInitialized)
				if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Initializable *InitializableFilterer) ParseInitialized(log types.Log) (*InitializableInitialized, error) {
	event := new(InitializableInitialized)
	if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerMetaData contains all meta data concerning the NativeTokenStakingManager contract.
var NativeTokenStakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumICMInitializable\",\"name\":\"init\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"InvalidBLSKeyLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"}],\"name\":\"InvalidDelegationFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidDelegationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumDelegatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidDelegatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializationStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\"}],\"name\":\"InvalidMaximumChurnPercentage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"}],\"name\":\"InvalidMinStakeDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"InvalidNodeID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addressesLength\",\"type\":\"uint256\"}],\"name\":\"InvalidPChainOwnerThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"InvalidRegistrationExpiry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"InvalidStakeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"}],\"name\":\"InvalidStakeMultiplier\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"encodedSubnetConversionID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedSubnetConversionID\",\"type\":\"bytes32\"}],\"name\":\"InvalidSubnetConversionID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"}],\"name\":\"InvalidValidatorManagerAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidatorManagerBlockchainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidValidatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWarpMessage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"InvalidWarpOriginSenderAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidWarpSourceChainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"churnAmount\",\"type\":\"uint64\"}],\"name\":\"MaxChurnRateExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newValidatorWeight\",\"type\":\"uint64\"}],\"name\":\"MaxWeightExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"name\":\"MinStakeDurationNotPassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PChainOwnerAddressesNotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UnauthorizedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"validRegistration\",\"type\":\"bool\"}],\"name\":\"UnexpectedRegistrationStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorNotPoS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"DelegationEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"delegatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"DelegatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"InitialValidatorCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"UptimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"ValidationPeriodCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ValidationPeriodEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidationPeriodRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"ValidationRewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"ValidatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorWeightUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_LENGTH\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLS_PUBLIC_KEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_CHURN_PERCENTAGE_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_REGISTRATION_EXPIRY_LENGTH\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_MINTER\",\"outputs\":[{\"internalType\":\"contractINativeMinter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POS_VALIDATOR_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"P_CHAIN_BLOCKCHAIN_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPTIME_REWARDS_THRESHOLD_PERCENTAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"claimDelegationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"claimValidationRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeDelegatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"startingWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"messageNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endedAt\",\"type\":\"uint64\"}],\"internalType\":\"structValidator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"churnPeriodSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\"}],\"internalType\":\"structValidatorManagerSettings\",\"name\":\"baseSettings\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minimumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minimumStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"contractIRewardCalculator\",\"name\":\"rewardCalculator\",\"type\":\"address\"}],\"internalType\":\"structPoSValidatorManagerSettings\",\"name\":\"settings\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initializeDelegatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"internalType\":\"structValidatorRegistrationInput\",\"name\":\"registrationInput\",\"type\":\"tuple\"},{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"}],\"name\":\"initializeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structSubnetConversionData\",\"name\":\"subnetConversionData\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"registeredValidators\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendEndValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendRegisterValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"resendUpdateDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"submitUptimeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"valueToWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"weightToValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051615cdc380380615cdc83398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b615b8f8061014d5f395ff3fe6080604052600436106101fc575f3560e01c8063732214f811610113578063b771b3bc1161009d578063c599e24f1161006d578063c599e24f146105cb578063c974d1b6146105de578063d5f20ff6146105f2578063df93d8de1461061e578063fd7ac5e714610634575f80fd5b8063b771b3bc14610553578063ba3a4b971461056d578063bc5fbfec1461058c578063bee0a03f146105ac575f80fd5b806393e24598116100e357806393e24598146104ce578063a3a65e48146104ed578063af2f5feb1461050c578063afb980961461051f578063afba878a1461053f575f80fd5b8063732214f81461046957806376f786211461047c57806380dd672f1461049b5780638280a25a146104ba575f80fd5b80632e2194d811610194578063467ef06f11610164578063467ef06f146103b657806360305d62146103d557806360ad7784146103fe578063620658561461041d57806366435abf1461044a575f80fd5b80632e2194d814610306578063329c3e121461033d57806335455ded1461036f5780633a1cfff614610397575f80fd5b806320d91b7a116101cf57806320d91b7a1461028a57806320e55565146102a957806325e1c776146102c85780632893d077146102e7575f80fd5b80630118acc4146102005780630322ed9814610221578063151d30d1146102405780631ec447241461026b575b5f80fd5b34801561020b575f80fd5b5061021f61021a366004614bec565b610653565b005b34801561022c575f80fd5b5061021f61023b366004614c27565b610665565b34801561024b575f80fd5b50610254600a81565b60405160ff90911681526020015b60405180910390f35b348015610276575f80fd5b5061021f610285366004614bec565b6108f5565b348015610295575f80fd5b5061021f6102a4366004614c3e565b610901565b3480156102b4575f80fd5b5061021f6102c3366004614c8c565b610ec9565b3480156102d3575f80fd5b5061021f6102e2366004614ca3565b610fa6565b3480156102f2575f80fd5b5061021f610301366004614ca3565b61101a565b348015610311575f80fd5b50610325610320366004614c27565b61136a565b6040516001600160401b039091168152602001610262565b348015610348575f80fd5b506103576001600160991b0181565b6040516001600160a01b039091168152602001610262565b34801561037a575f80fd5b5061038461271081565b60405161ffff9091168152602001610262565b3480156103a2575f80fd5b5061021f6103b1366004614bec565b611380565b3480156103c1575f80fd5b5061021f6103d0366004614cc4565b61138c565b3480156103e0575f80fd5b506103e9601481565b60405163ffffffff9091168152602001610262565b348015610409575f80fd5b5061021f610418366004614ca3565b61147a565b348015610428575f80fd5b5061043c610437366004614cf1565b611752565b604051908152602001610262565b348015610455575f80fd5b50610325610464366004614c27565b61176b565b348015610474575f80fd5b5061043c5f81565b348015610487575f80fd5b5061021f610496366004614bec565b61177f565b3480156104a6575f80fd5b5061021f6104b5366004614ca3565b61178c565b3480156104c5575f80fd5b50610254603081565b3480156104d9575f80fd5b5061021f6104e8366004614c27565b6119cd565b3480156104f8575f80fd5b5061021f610507366004614cc4565b611a77565b61043c61051a366004614d1d565b611c6c565b34801561052a575f80fd5b5061043c5f80516020615ada83398151915281565b34801561054a575f80fd5b50610254605081565b34801561055e575f80fd5b506103576005600160991b0181565b348015610578575f80fd5b5061021f610587366004614c27565b611ca0565b348015610597575f80fd5b5061043c5f80516020615afa83398151915281565b3480156105b7575f80fd5b5061021f6105c6366004614c27565b611f0b565b61043c6105d9366004614c27565b612048565b3480156105e9575f80fd5b50610254601481565b3480156105fd575f80fd5b5061061161060c366004614c27565b612079565b6040516102629190614df3565b348015610629575f80fd5b506103256202a30081565b34801561063f575f80fd5b5061043c61064e366004614e73565b6121c8565b6106608383836001612223565b505050565b5f8181525f80516020615b3a8339815191526020526040808220815160e0810190925280545f80516020615afa83398151915293929190829060ff1660058111156106b2576106b2614d7e565b60058111156106c3576106c3614d7e565b81526020016001820180546106d790614ede565b80601f016020809104026020016040519081016040528092919081815260200182805461070390614ede565b801561074e5780601f106107255761010080835404028352916020019161074e565b820191905f5260205f20905b81548152906001019060200180831161073157829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b909104811660808301526003928301541660a090910152909150815160058111156107b9576107b9614d7e565b146107f5575f8381526007830160205260409081902054905163170cc93360e21b81526107ec9160ff1690600401614f10565b60405180910390fd5b60608101516040516342a2e0b560e11b8152600481018590526001600160401b0390911660248201525f60448201526005600160991b019063ee5b48eb9073__$fd0c147b4031eef6079b0498cbafa865f0$__90638545c16a906064015f60405180830381865af415801561086c573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526108939190810190615021565b6040518263ffffffff1660e01b81526004016108af9190615052565b6020604051808303815f875af11580156108cb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108ef9190615064565b50505050565b6106608383835f612223565b7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb09545f80516020615afa8339815191529060ff161561095357604051637fab81e560e01b815260040160405180910390fd5b6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015610996573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109ba9190615064565b8360200135146109e3576040516372b0a7e760e11b8152602084013560048201526024016107ec565b306109f4606085016040860161508f565b6001600160a01b031614610a3757610a12606084016040850161508f565b604051632f88120d60e21b81526001600160a01b0390911660048201526024016107ec565b5f610a4560608501856150aa565b905090505f805b828163ffffffff161015610d2d575f610a6860608801886150aa565b8363ffffffff16818110610a7e57610a7e6150ef565b9050602002810190610a909190615103565b610a999061516e565b80516040519192505f916008880191610ab1916151e9565b90815260200160405180910390205414610ae157805160405163a41f772f60e01b81526107ec9190600401615052565b5f6002885f013584604051602001610b1092919091825260e01b6001600160e01b031916602082015260240190565b60408051601f1981840301815290829052610b2a916151e9565b602060405180830381855afa158015610b45573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610b689190615064565b90508086600801835f0151604051610b8091906151e9565b90815260408051602092819003830181209390935560e0830181526002835284518284015284810180516001600160401b03908116858401525f60608601819052915181166080860152421660a085015260c0840181905284815260078a01909252902081518154829060ff19166001836005811115610c0257610c02614d7e565b021790555060208201516001820190610c1b908261523e565b506040828101516002830180546060860151608087015160a08801516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c0909301516003909201805467ffffffffffffffff191692841692909217909155830151610cc091168561530d565b8251604051919550610cd1916151e9565b60408051918290038220908401516001600160401b031682529082907f9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf9060200160405180910390a3505080610d2690615320565b9050610a4c565b50600483018190555f73__$fd0c147b4031eef6079b0498cbafa865f0$__631e6d9789610d598761254f565b604001516040518263ffffffff1660e01b8152600401610d799190615052565b602060405180830381865af4158015610d94573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610db89190615064565b90505f73__$fd0c147b4031eef6079b0498cbafa865f0$__63862bfa63886040518263ffffffff1660e01b8152600401610df2919061546d565b5f60405180830381865af4158015610e0c573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610e339190810190615021565b90505f600282604051610e4691906151e9565b602060405180830381855afa158015610e61573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610e849190615064565b9050828114610eb057604051631872fc8d60e01b815260048101829052602481018490526044016107ec565b5050506009909201805460ff1916600117905550505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805460029190600160401b900460ff1680610f12575080546001600160401b03808416911610155b15610f305760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b03831617600160401b178155610f5a83612665565b805460ff60401b191681556040516001600160401b03831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050565b610faf82612676565b610fcf576040516330efa98b60e01b8152600481018390526024016107ec565b5f610fd983612079565b5190506002816005811115610ff057610ff0614d7e565b14611010578060405163170cc93360e21b81526004016107ec9190614f10565b6108ef83836126b1565b611022612997565b5f80516020615ada8339815191525f61103a84612079565b905060028151600581111561105157611051614d7e565b1461107257805160405163170cc93360e21b81526107ec9190600401614f10565b61107b84612676565b61109b576040516330efa98b60e01b8152600481018590526024016107ec565b5f8481526004830160205260409020546001600160a01b031633146110e157335b604051636e2ccd7560e11b81526001600160a01b0390911660048201526024016107ec565b5f84815260048301602052604090205460a0820151429161111391600160b01b9091046001600160401b0316906154f8565b6001600160401b0316816001600160401b031610156111505760405163fb6ce63f60e01b81526001600160401b03821660048201526024016107ec565b5f61115b86866126b1565b5f8781526004860160205260408120600101549192509061118c90600160401b90046001600160401b031683615518565b5f888152600487016020526040812060010154919250600160801b9091046001600160401b0316908190036111c2575060a08401515b6111cd8282866129ce565b6111ed57604051635bff683f60e11b8152600481018990526024016107ec565b600386015460408601515f916001600160a01b03169063778c06b59061121290611752565b848589885f806040518863ffffffff1660e01b815260040161123a9796959493929190615538565b602060405180830381865afa158015611255573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112799190615064565b90506064605060ff168760a00151876112929190615518565b61129c9190615576565b6112a691906155b5565b5f8a8152600489016020526040902060018101805477ffffffffffffffffffffffffffffffff00000000000000001916600160401b6001600160401b039485160267ffffffffffffffff60801b191617600160801b93891693909302929092179091555461131d906001600160a01b031682612a19565b60405181908a907f69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a905f90a35050505050505061136660015f80516020615b1a83398151915255565b5050565b5f61137a64e8d4a51000836155da565b92915050565b6106608383835f612a8a565b611394612997565b5f80516020615ada8339815191525f806113ad84612ceb565b915091506113ba82612676565b6113c657505050611461565b5f82815260048085016020526040909120546001600160a01b031690825160058111156113f5576113f5614d7e565b03611446575f838152600785016020526040812080549190556114188282612a19565b604051819085907f69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a905f90a3505b61145c816114578460400151611752565b61309e565b505050505b61147760015f80516020615b1a83398151915255565b50565b5f8281525f80516020615aba8339815191526020526040808220815160e0810190925280545f80516020615ada83398151915293929190829060ff1660038111156114c7576114c7614d7e565b60038111156114d8576114d8614d7e565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091505f61154e82612079565b905060018351600381111561156557611565614d7e565b14611586578251604051633b0d540d60e21b81526107ec91906004016155ed565b60048151600581111561159b5761159b614d7e565b036115b1576115a9866130b1565b505050505050565b5f8073__$fd0c147b4031eef6079b0498cbafa865f0$__639de23d406115d68961254f565b604001516040518263ffffffff1660e01b81526004016115f69190615052565b606060405180830381865af4158015611611573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116359190615607565b509150915081841461166257846040015160405163089938b360e11b81526004016107ec91815260200190565b806001600160401b031683606001516001600160401b0316108061169b5750806001600160401b03168560a001516001600160401b0316115b156116c457604051632e19bc2d60e11b81526001600160401b03821660048201526024016107ec565b5f888152600587016020908152604091829020805460ff1916600290811782550180546001600160401b034216600160401b81026fffffffffffffffff00000000000000001990921691909117909155915191825285918a917f047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6910160405180910390a35050505050505050565b5f61137a6001600160401b03831664e8d4a5100061563c565b5f61177582612079565b6080015192915050565b6106608383836001612a8a565b611794612997565b5f8281525f80516020615aba8339815191526020526040808220815160e0810190925280545f80516020615ada83398151915293929190829060ff1660038111156117e1576117e1614d7e565b60038111156117f2576117f2614d7e565b8152815461010090046001600160a01b03166020820152600182015460408201526002909101546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c090910152905060038151600381111561186b5761186b614d7e565b1461188c578051604051633b0d540d60e21b81526107ec91906004016155ed565b600461189b8260400151612079565b5160058111156118ad576118ad614d7e565b146119ac575f6118bc8461254f565b90505f8073__$fd0c147b4031eef6079b0498cbafa865f0$__639de23d4084604001516040518263ffffffff1660e01b81526004016118fb9190615052565b606060405180830381865af4158015611916573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061193a9190615607565b5091509150818460400151146119665760405163089938b360e11b8152600481018390526024016107ec565b806001600160401b03168460c001516001600160401b031611156119a857604051632e19bc2d60e11b81526001600160401b03821660048201526024016107ec565b5050505b6119b5846130b1565b505061136660015f80516020615b1a83398151915255565b5f80516020615ada8339815191525f6119e583612079565b51905060048160058111156119fc576119fc614d7e565b14611a1c578060405163170cc93360e21b81526004016107ec9190614f10565b5f8381526004830160205260409020546001600160a01b03163314611a4157336110bc565b5f838152600783016020908152604080832080549084905560048601909252909120546108ef906001600160a01b031682612a19565b5f80516020615afa8339815191525f8073__$fd0c147b4031eef6079b0498cbafa865f0$__632e43ceb5611aaa8661254f565b604001516040518263ffffffff1660e01b8152600401611aca9190615052565b6040805180830381865af4158015611ae4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b089190615653565b9150915080611b2e57604051632d07135360e01b815281151560048201526024016107ec565b5f82815260068401602052604090208054611b4890614ede565b90505f03611b6c5760405163089938b360e11b8152600481018390526024016107ec565b60015f83815260078501602052604090205460ff166005811115611b9257611b92614d7e565b14611bc5575f8281526007840160205260409081902054905163170cc93360e21b81526107ec9160ff1690600401614f10565b5f8281526006840160205260408120611bdd91614b40565b5f828152600784016020908152604091829020805460ff1916600290811782550180546001600160401b0342818116600160c01b026001600160c01b0390931692909217928390558451600160801b9093041682529181019190915283917ff8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568910160405180910390a250505050565b5f611c75612997565b611c8184848434613298565b9050611c9960015f80516020615b1a83398151915255565b9392505050565b5f8181525f80516020615aba8339815191526020526040808220815160e0810190925280545f80516020615ada83398151915293929190829060ff166003811115611ced57611ced614d7e565b6003811115611cfe57611cfe614d7e565b8152815461010090046001600160a01b0316602082015260018083015460408301526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015290915081516003811115611d7757611d77614d7e565b14158015611d985750600381516003811115611d9557611d95614d7e565b14155b15611db9578051604051633b0d540d60e21b81526107ec91906004016155ed565b5f611dc78260400151612079565b905080606001516001600160401b03165f03611df9576040516339b894f960e21b8152600481018590526024016107ec565b6040808301516060830151608084015192516342a2e0b560e11b81526005600160991b019363ee5b48eb9373__$fd0c147b4031eef6079b0498cbafa865f0$__93638545c16a93611e6793906004019283526001600160401b03918216602084015216604082015260600190565b5f60405180830381865af4158015611e81573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611ea89190810190615021565b6040518263ffffffff1660e01b8152600401611ec49190615052565b6020604051808303815f875af1158015611ee0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f049190615064565b5050505050565b5f8181527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb066020526040902080545f80516020615afa8339815191529190611f5290614ede565b90505f03611f765760405163089938b360e11b8152600481018390526024016107ec565b60015f83815260078301602052604090205460ff166005811115611f9c57611f9c614d7e565b14611fcf575f8281526007820160205260409081902054905163170cc93360e21b81526107ec9160ff1690600401614f10565b5f82815260068201602052604090819020905163ee5b48eb60e01b81526005600160991b019163ee5b48eb916120089190600401615681565b6020604051808303815f875af1158015612024573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106609190615064565b5f612051612997565b61205c8233346134ab565b905061207460015f80516020615b1a83398151915255565b919050565b612081614b77565b5f8281525f80516020615b3a833981519152602052604090819020815160e0810190925280545f80516020615afa833981519152929190829060ff1660058111156120ce576120ce614d7e565b60058111156120df576120df614d7e565b81526020016001820180546120f390614ede565b80601f016020809104026020016040519081016040528092919081815260200182805461211f90614ede565b801561216a5780601f106121415761010080835404028352916020019161216a565b820191905f5260205f20905b81548152906001019060200180831161214d57829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b9091048116608083015260039092015490911660a0909101529392505050565b6040515f905f80516020615afa833981519152907fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb089061220b908690869061570b565b90815260200160405180910390205491505092915050565b5f8481525f80516020615aba8339815191526020526040808220815160e0810190925280545f80516020615ada83398151915293929190829060ff16600381111561227057612270614d7e565b600381111561228157612281614d7e565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091505f6122f782612079565b905060028351600381111561230e5761230e614d7e565b1461232f578251604051633b0d540d60e21b81526107ec91906004016155ed565b60208301516001600160a01b031633146123d2575f828152600485016020526040902054336001600160a01b03909116036123cc575f82815260048501602052604090205460a082015161239391600160b01b90046001600160401b0316906154f8565b6001600160401b03164210156123c75760405163fb6ce63f60e01b81526001600160401b03421660048201526024016107ec565b6123d2565b336110bc565b6002815160058111156123e7576123e7614d7e565b036124e85786156123fe576123fc82876126b1565b505b5f8881526005850160205260409020805460ff19166003179055606083015160808201516124379184916124329190615518565b613780565b505f898152600586016020526040812060020180546001600160401b03909316600160c01b026001600160c01b039093169290921790915561247884613957565b9050858015612485575080155b156124a657604051631036cf9160e11b8152600481018a90526024016107ec565b5f8981526006860160205260408082208390555184918b917f366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed579190a350612545565b6004815160058111156124fd576124fd614d7e565b036125295761250b83613957565b5f898152600686016020526040902055612524886130b1565b612545565b805160405163170cc93360e21b81526107ec9190600401614f10565b5050505050505050565b60408051606080820183525f8083526020830152918101919091526040516306f8253560e41b815263ffffffff831660048201525f9081906005600160991b0190636f825350906024015f60405180830381865afa1580156125b3573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526125da919081019061571a565b91509150806125fc57604051636b2f19e960e01b815260040160405180910390fd5b815115612622578151604051636ba589a560e01b815260048101919091526024016107ec565b60208201516001600160a01b03161561265e576020820151604051624de75d60e31b81526001600160a01b0390911660048201526024016107ec565b5092915050565b61266d613b20565b61147781613b6b565b5f9081527f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d0460205260409020546001600160a01b0316151590565b6040516306f8253560e41b815263ffffffff821660048201525f90819081906005600160991b0190636f825350906024015f60405180830381865afa1580156126fc573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612723919081019061571a565b915091508061274557604051636b2f19e960e01b815260040160405180910390fd5b6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015612788573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127ac9190615064565b8251146127d2578151604051636ba589a560e01b815260048101919091526024016107ec565b60208201516001600160a01b03161561280e576020820151604051624de75d60e31b81526001600160a01b0390911660048201526024016107ec565b5f8073__$fd0c147b4031eef6079b0498cbafa865f0$__63088c246385604001516040518263ffffffff1660e01b815260040161284b9190615052565b6040805180830381865af4158015612865573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061288991906157aa565b915091508187146128b05760405163089938b360e11b8152600481018890526024016107ec565b5f8781527f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d0460205260409020600101545f80516020615ada833981519152906001600160401b03908116908316111561296d575f888152600482016020908152604091829020600101805467ffffffffffffffff19166001600160401b038616908117909155915191825289917fec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435910160405180910390a261298c565b5f8881526004820160205260409020600101546001600160401b031691505b509695505050505050565b5f80516020615b1a8339815191528054600119016129c857604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f60506129db8484615518565b6129e59190615576565b6001600160401b03166129f9856064615576565b6001600160401b03161015612a0f57505f611c99565b5060019392505050565b6040516327ad555d60e11b81526001600160a01b0383166004820152602481018290526001600160991b0190634f5aaaba906044015f604051808303815f87803b158015612a65575f80fd5b505af11580156115a9573d5f803e3d5ffd5b60015f80516020615b1a83398151915255565b5f80516020615ada8339815191525f612aa286613bd9565b9050612aad86612676565b612ab85750506108ef565b5f8681526004830160205260409020546001600160a01b03163314612add57336110bc565b5f86815260048301602052604090205460a0820151612b0c91600160b01b90046001600160401b0316906154f8565b6001600160401b03168160c001516001600160401b03161015612b535760c081015160405163fb6ce63f60e01b81526001600160401b0390911660048201526024016107ec565b5f8515612b6b57612b6487866126b1565b9050612b89565b505f8681526004830160205260409020600101546001600160401b03165b5f878152600484016020526040812060010154612bb690600160401b90046001600160401b031683615518565b5f898152600486016020526040812060010154919250600160801b9091046001600160401b031690819003612bec575060a08301515b858015612c055750612c0382828660c001516129ce565b155b15612c2657604051635bff683f60e11b8152600481018a90526024016107ec565b600385015460408501515f916001600160a01b03169063778c06b590612c4b90611752565b84858960c00151885f806040518863ffffffff1660e01b8152600401612c779796959493929190615538565b602060405180830381865afa158015612c92573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612cb69190615064565b905080866007015f8c81526020019081526020015f205f828254612cda919061530d565b909155505050505050505050505050565b5f612cf4614b77565b5f80516020615afa8339815191525f8073__$fd0c147b4031eef6079b0498cbafa865f0$__632e43ceb5612d278861254f565b604001516040518263ffffffff1660e01b8152600401612d479190615052565b6040805180830381865af4158015612d61573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612d859190615653565b915091508015612dac57604051632d07135360e01b815281151560048201526024016107ec565b5f828152600784016020526040808220815160e081019092528054829060ff166005811115612ddd57612ddd614d7e565b6005811115612dee57612dee614d7e565b8152602001600182018054612e0290614ede565b80601f0160208091040260200160405190810160405280929190818152602001828054612e2e90614ede565b8015612e795780601f10612e5057610100808354040283529160200191612e79565b820191905f5260205f20905b815481529060010190602001808311612e5c57829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b909104811660808301526003928301541660a09091015290915081516005811115612ee457612ee4614d7e565b14158015612f055750600181516005811115612f0257612f02614d7e565b14155b15612f2657805160405163170cc93360e21b81526107ec9190600401614f10565b600381516005811115612f3b57612f3b614d7e565b03612f495760048152612f4e565b600581525b836008018160200151604051612f6491906151e9565b90815260408051602092819003830190205f908190558581526007870190925290208151815483929190829060ff19166001836005811115612fa857612fa8614d7e565b021790555060208201516001820190612fc1908261523e565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c0909201516003909101805467ffffffffffffffff1916919092161790558051600581111561306757613067614d7e565b60405184907f1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16905f90a39196919550909350505050565b6113666001600160a01b03831682613ebd565b5f8181525f80516020615aba8339815191526020526040808220815160e0810190925280545f80516020615ada83398151915293929190829060ff1660038111156130fe576130fe614d7e565b600381111561310f5761310f614d7e565b8152815461010090046001600160a01b03166020808301919091526001808401546040808501919091526002948501546001600160401b038082166060870152600160401b820481166080870152600160801b8204811660a0870152600160c01b9091041660c090940193909352848301515f89815260058901845284812080546001600160a81b0319168155928301819055919094018190556006870190915290812080549082905592935090919080821561323d575f848152600487016020526040902054612710906131ef90600160a01b900461ffff168561563c565b6131f991906155da565b915081866007015f8681526020019081526020015f205f82825461321d919061530d565b9091555061322d905082846157cd565b905061323d856020015182612a19565b61325285602001516114578760600151611752565b6040805182815260208101849052859189917f8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993910160405180910390a350505050505050565b7f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d02545f905f80516020615ada83398151915290600160401b900461ffff90811690861610806132ec575061271061ffff8616115b1561331057604051635f12e6c360e11b815261ffff861660048201526024016107ec565b60028101546001600160401b03908116908516101561334c576040516202a06d60e11b81526001600160401b03851660048201526024016107ec565b805483108061335e5750806001015483115b1561337f5760405163222d164360e21b8152600481018490526024016107ec565b825f61338a8261136a565b90505f6133978983613f50565b90506040518060c001604052806133ab3390565b6001600160a01b03908116825261ffff808c166020808501919091526001600160401b03808d166040808701919091525f6060808801829052608080890183905260a09889018390528a83526004909d0185529082902088518154958a0151938a01518516600160b01b0267ffffffffffffffff60b01b1994909716600160a01b026001600160b01b03199096169716969096179390931716929092178355840151600190920180549885015194909301518116600160801b0267ffffffffffffffff60801b19948216600160401b026001600160801b031990991692909116919091179690961791909116949094179093555090915050949350505050565b5f5f80516020615ada833981519152816134c48461136a565b90505f6134d087612079565b90506134db87612676565b6134fb576040516330efa98b60e01b8152600481018890526024016107ec565b60028151600581111561351057613510614d7e565b1461353157805160405163170cc93360e21b81526107ec9190600401614f10565b5f82826080015161354291906154f8565b905083600201600a9054906101000a90046001600160401b0316826040015161356b9190615576565b6001600160401b0316816001600160401b031611156135a857604051636d51fe0560e11b81526001600160401b03821660048201526024016107ec565b5f806135b48a84613780565b915091505f8a836040516020016135e292919091825260c01b6001600160c01b031916602082015260280190565b60408051601f19818403018152828252805160209091012060e08301909152915080600181526001600160a01b038c1660208083019190915260408083018f90526001600160401b03808b1660608501525f6080850181905290881660a085015260c090930183905284835260058b01909152902081518154829060ff1916600183600381111561367557613675614d7e565b02179055506020828101518254610100600160a81b0319166101006001600160a01b039283160217835560408085015160018501556060808601516002909501805460808089015160a08a015160c0909a01516001600160401b03998a166001600160801b031990941693909317600160401b918a1691909102176001600160801b0316600160801b998916999099026001600160c01b031698909817600160c01b91881691909102179055815189861681528a861694810194909452938b1690830152918101859052908c16918d9184917fb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426910160405180910390a49a9950505050505050505050565b5f8281525f80516020615b3a833981519152602052604081206002015481905f80516020615afa83398151915290600160801b90046001600160401b03166137c88582614538565b5f6137d287614712565b5f888152600785016020526040808220600201805467ffffffffffffffff60801b1916600160801b6001600160401b038c81169182029290921790925591516342a2e0b560e11b8152600481018c905291841660248301526044820152919250906005600160991b019063ee5b48eb9073__$fd0c147b4031eef6079b0498cbafa865f0$__90638545c16a906064015f60405180830381865af415801561387b573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526138a29190810190615021565b6040518263ffffffff1660e01b81526004016138be9190615052565b6020604051808303815f875af11580156138da573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906138fe9190615064565b604080516001600160401b038a811682526020820184905282519394508516928b927f07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df928290030190a3909450925050505b9250929050565b5f805f80516020615ada83398151915290505f6139778460400151612079565b90505f60038251600581111561398f5761398f614d7e565b14806139ad57506004825160058111156139ab576139ab614d7e565b145b156139bd575060c08101516139fa565b6002825160058111156139d2576139d2614d7e565b036139de5750426139fa565b815160405163170cc93360e21b81526107ec9190600401614f10565b84608001516001600160401b0316816001600160401b031611613a2157505f949350505050565b6040808601515f90815260048501602052206001015460a0830151613a50916001600160401b031690836129ce565b613a5e57505f949350505050565b600383015460608601516001600160a01b039091169063778c06b590613a8390611752565b60a085015160808901516040808b01515f9081526004808b016020528282206001015492516001600160e01b031960e089901b168152613ad8969594938a936001600160401b03909116929091829101615538565b602060405180830381865afa158015613af3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b179190615064565b95945050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16613b6957604051631afcd79f60e31b815260040160405180910390fd5b565b613b73613b20565b613b7c81614787565b613b846147a0565b61147760608201356080830135613ba160c0850160a08601614cf1565b613bb160e0860160c087016157e0565b613bc2610100870160e088016157f9565b613bd46101208801610100890161508f565b6147b0565b613be1614b77565b5f8281525f80516020615b3a8339815191526020526040808220815160e0810190925280545f80516020615afa83398151915293929190829060ff166005811115613c2e57613c2e614d7e565b6005811115613c3f57613c3f614d7e565b8152602001600182018054613c5390614ede565b80601f0160208091040260200160405190810160405280929190818152602001828054613c7f90614ede565b8015613cca5780601f10613ca157610100808354040283529160200191613cca565b820191905f5260205f20905b815481529060010190602001808311613cad57829003601f168201915b50505091835250506002828101546001600160401b038082166020850152600160401b820481166040850152600160801b820481166060850152600160c01b9091048116608084015260039093015490921660a09091015290915081516005811115613d3857613d38614d7e565b14613d6b575f8481526007830160205260409081902054905163170cc93360e21b81526107ec9160ff1690600401614f10565b60038152426001600160401b031660c08201525f84815260078301602052604090208151815483929190829060ff19166001836005811115613daf57613daf614d7e565b021790555060208201516001820190613dc8908261523e565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c0909201516003909101805467ffffffffffffffff1916919092161790555f613e668582613780565b6080840151604080516001600160401b03909216825242602083015291935083925087917f13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67910160405180910390a3509392505050565b80471015613ee05760405163cd78605960e01b81523060048201526024016107ec565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114613f29576040519150601f19603f3d011682016040523d82523d5f602084013e613f2e565b606091505b505090508061066057604051630a12f52160e11b815260040160405180910390fd5b7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb09545f9060ff16613f9457604051637fab81e560e01b815260040160405180910390fd5b5f80516020615afa83398151915242613fb36060860160408701614cf1565b6001600160401b0316111580613fed5750613fd16202a3004261530d565b613fe16060860160408701614cf1565b6001600160401b031610155b15614027576140026060850160408601614cf1565b604051635879da1360e11b81526001600160401b0390911660048201526024016107ec565b61403c6140376060860186615819565b6148df565b61404c6140376080860186615819565b603061405b602086018661582d565b90501461408d5761406f602085018561582d565b6040516326475b2f60e11b81526107ec925060040190815260200190565b614097848061582d565b90505f036140c4576140a9848061582d565b604051633e08a12560e11b81526004016107ec92919061586f565b5f600882016140d3868061582d565b6040516140e192919061570b565b9081526020016040518091039020541461411a576140ff848061582d565b60405163a41f772f60e01b81526004016107ec92919061586f565b614124835f614538565b6040805160e08101909152815481525f90819073__$fd0c147b4031eef6079b0498cbafa865f0$__9063e047b28390602081016141618a8061582d565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505050908252506020908101906141a9908b018b61582d565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505050908252506020016141f260608b0160408c01614cf1565b6001600160401b0316815260200161420d60608b018b615819565b61421690615882565b815260200161422860808b018b615819565b61423190615882565b8152602001886001600160401b03168152506040518263ffffffff1660e01b815260040161425f91906159a4565b5f60405180830381865af4158015614279573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526142a09190810190615a5b565b5f828152600686016020526040902091935091506142be828261523e565b5081600884016142ce888061582d565b6040516142dc92919061570b565b9081526040519081900360200181209190915563ee5b48eb60e01b81525f906005600160991b019063ee5b48eb90614318908590600401615052565b6020604051808303815f875af1158015614334573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906143589190615064565b6040805160e081019091529091508060018152602001614378898061582d565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201829052509385525050506001600160401b0389166020808401829052604080850184905260608501929092526080840183905260a0909301829052868252600788019092522081518154829060ff1916600183600581111561440757614407614d7e565b021790555060208201516001820190614420908261523e565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c0909201516003909101805467ffffffffffffffff191691909216179055806144be888061582d565b6040516144cc92919061570b565b6040518091039020847fb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430898b604001602081019061450a9190614cf1565b604080516001600160401b0393841681529290911660208301520160405180910390a4509095945050505050565b5f80516020615afa8339815191525f6001600160401b03808416908516111561456c576145658385615518565b9050614579565b6145768484615518565b90505b6040805160808101825260028401548082526003850154602083015260048501549282019290925260058401546001600160401b03166060820152429115806145db5750600184015481516145d7916001600160401b03169061530d565b8210155b15614601576001600160401b038316606082015281815260408101516020820152614620565b828160600181815161461391906154f8565b6001600160401b03169052505b6060810151614630906064615576565b602082015160018601546001600160401b03929092169161465b9190600160401b900460ff1661563c565b101561468b57606081015160405163dfae880160e01b81526001600160401b0390911660048201526024016107ec565b856001600160401b0316816040018181516146a6919061530d565b9052506040810180516001600160401b03871691906146c69083906157cd565b905250805160028501556020810151600385015560408101516004850155606001516005909301805467ffffffffffffffff19166001600160401b039094169390931790925550505050565b5f8181525f80516020615b3a8339815191526020526040812060020180545f80516020615afa833981519152919060089061475c90600160401b90046001600160401b0316615a9e565b91906101000a8154816001600160401b0302191690836001600160401b031602179055915050919050565b61478f613b20565b614797614a48565b61147781614a50565b6147a8613b20565b613b69614b38565b6147b8613b20565b5f80516020615ada83398151915261ffff841615806147dc575061271061ffff8516115b1561480057604051635f12e6c360e11b815261ffff851660048201526024016107ec565b858711156148245760405163222d164360e21b8152600481018890526024016107ec565b60ff831615806148375750600a60ff8416115b1561485a5760405163170db35960e31b815260ff841660048201526024016107ec565b95865560018601949094556002850180546001600160401b039490941669ffffffffffffffffffff1990941693909317600160401b61ffff93909316929092029190911767ffffffffffffffff60501b191660ff91909116600160501b02179055600390910180546001600160a01b0319166001600160a01b03909216919091179055565b6148ec6020820182614cc4565b63ffffffff1615801561490c575061490760208201826150aa565b151590505b156149535761491e6020820182614cc4565b61492b60208301836150aa565b60405163c08a0f1d60e01b815263ffffffff90931660048401526024830152506044016107ec565b61496060208201826150aa565b905061496f6020830183614cc4565b63ffffffff1611156149885761491e6020820182614cc4565b60015b61499860208301836150aa565b9050811015611366576149ae60208301836150aa565b6149b96001846157cd565b8181106149c8576149c86150ef565b90506020020160208101906149dd919061508f565b6001600160a01b03166149f360208401846150aa565b83818110614a0357614a036150ef565b9050602002016020810190614a18919061508f565b6001600160a01b03161015614a4057604051630dbc8d5f60e31b815260040160405180910390fd5b60010161498b565b613b69613b20565b614a58613b20565b80355f80516020615afa8339815191529081556014614a7d60608401604085016157f9565b60ff161180614a9c5750614a9760608301604084016157f9565b60ff16155b15614ad057614ab160608301604084016157f9565b604051634a59bbff60e11b815260ff90911660048201526024016107ec565b614ae060608301604084016157f9565b60018201805460ff92909216600160401b0260ff60401b19909216919091179055614b116040830160208401614cf1565b600191909101805467ffffffffffffffff19166001600160401b0390921691909117905550565b612a77613b20565b508054614b4c90614ede565b5f825580601f10614b5b575050565b601f0160209004905f5260205f20908101906114779190614bb4565b6040805160e08101909152805f81526060602082018190525f604083018190529082018190526080820181905260a0820181905260c09091015290565b5b80821115614bc8575f8155600101614bb5565b5090565b8015158114611477575f80fd5b803563ffffffff81168114612074575f80fd5b5f805f60608486031215614bfe575f80fd5b833592506020840135614c1081614bcc565b9150614c1e60408501614bd9565b90509250925092565b5f60208284031215614c37575f80fd5b5035919050565b5f8060408385031215614c4f575f80fd5b82356001600160401b03811115614c64575f80fd5b830160808186031215614c75575f80fd5b9150614c8360208401614bd9565b90509250929050565b5f6101208284031215614c9d575f80fd5b50919050565b5f8060408385031215614cb4575f80fd5b82359150614c8360208401614bd9565b5f60208284031215614cd4575f80fd5b611c9982614bd9565b6001600160401b0381168114611477575f80fd5b5f60208284031215614d01575f80fd5b8135611c9981614cdd565b803561ffff81168114612074575f80fd5b5f805f60608486031215614d2f575f80fd5b83356001600160401b03811115614d44575f80fd5b840160a08187031215614d55575f80fd5b9250614d6360208501614d0c565b91506040840135614d7381614cdd565b809150509250925092565b634e487b7160e01b5f52602160045260245ffd5b60068110614da257614da2614d7e565b9052565b5f5b83811015614dc0578181015183820152602001614da8565b50505f910152565b5f8151808452614ddf816020860160208601614da6565b601f01601f19169290920160200192915050565b60208152614e05602082018351614d92565b5f602083015160e06040840152614e20610100840182614dc8565b905060408401516001600160401b0380821660608601528060608701511660808601528060808701511660a08601528060a08701511660c08601528060c08701511660e086015250508091505092915050565b5f8060208385031215614e84575f80fd5b82356001600160401b0380821115614e9a575f80fd5b818501915085601f830112614ead575f80fd5b813581811115614ebb575f80fd5b866020828501011115614ecc575f80fd5b60209290920196919550909350505050565b600181811c90821680614ef257607f821691505b602082108103614c9d57634e487b7160e01b5f52602260045260245ffd5b6020810161137a8284614d92565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715614f5457614f54614f1e565b60405290565b604080519081016001600160401b0381118282101715614f5457614f54614f1e565b604051601f8201601f191681016001600160401b0381118282101715614fa457614fa4614f1e565b604052919050565b5f6001600160401b03821115614fc457614fc4614f1e565b50601f01601f191660200190565b5f82601f830112614fe1575f80fd5b8151614ff4614fef82614fac565b614f7c565b818152846020838601011115615008575f80fd5b615019826020830160208701614da6565b949350505050565b5f60208284031215615031575f80fd5b81516001600160401b03811115615046575f80fd5b61501984828501614fd2565b602081525f611c996020830184614dc8565b5f60208284031215615074575f80fd5b5051919050565b6001600160a01b0381168114611477575f80fd5b5f6020828403121561509f575f80fd5b8135611c998161507b565b5f808335601e198436030181126150bf575f80fd5b8301803591506001600160401b038211156150d8575f80fd5b6020019150600581901b3603821315613950575f80fd5b634e487b7160e01b5f52603260045260245ffd5b5f8235605e19833603018112615117575f80fd5b9190910192915050565b5f82601f830112615130575f80fd5b813561513e614fef82614fac565b818152846020838601011115615152575f80fd5b816020850160208301375f918101602001919091529392505050565b5f6060823603121561517e575f80fd5b615186614f32565b82356001600160401b038082111561519c575f80fd5b6151a836838701615121565b835260208501359150808211156151bd575f80fd5b506151ca36828601615121565b60208301525060408301356151de81614cdd565b604082015292915050565b5f8251615117818460208701614da6565b601f82111561066057805f5260205f20601f840160051c8101602085101561521f5750805b601f840160051c820191505b81811015611f04575f815560010161522b565b81516001600160401b0381111561525757615257614f1e565b61526b816152658454614ede565b846151fa565b602080601f83116001811461529e575f84156152875750858301515b5f19600386901b1c1916600185901b1785556115a9565b5f85815260208120601f198616915b828110156152cc578886015182559484019460019091019084016152ad565b50858210156152e957878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561137a5761137a6152f9565b5f63ffffffff808316818103615338576153386152f9565b6001019392505050565b5f808335601e19843603018112615357575f80fd5b83016020810192503590506001600160401b03811115615375575f80fd5b803603821315613950575f80fd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8383855260208086019550808560051b830101845f5b8781101561546057848303601f19018952813536889003605e190181126153e7575f80fd5b870160606153f58280615342565b8287526154058388018284615383565b9250505061541586830183615342565b86830388880152615427838284615383565b92505050604080830135925061543c83614cdd565b6001600160401b0392909216949091019390935297830197908301906001016153c2565b5090979650505050505050565b6020815281356020820152602082013560408201525f60408301356154918161507b565b6001600160a01b031660608381019190915283013536849003601e190181126154b8575f80fd5b83016020810190356001600160401b038111156154d3575f80fd5b8060051b36038213156154e4575f80fd5b608080850152613b1760a0850182846153ab565b6001600160401b0381811683821601908082111561265e5761265e6152f9565b6001600160401b0382811682821603908082111561265e5761265e6152f9565b9687526001600160401b03958616602088015293851660408701529184166060860152909216608084015260a083019190915260c082015260e00190565b6001600160401b03818116838216028082169190828114615599576155996152f9565b505092915050565b634e487b7160e01b5f52601260045260245ffd5b5f6001600160401b03808416806155ce576155ce6155a1565b92169190910492915050565b5f826155e8576155e86155a1565b500490565b602081016004831061560157615601614d7e565b91905290565b5f805f60608486031215615619575f80fd5b83519250602084015161562b81614cdd565b6040850151909250614d7381614cdd565b808202811582820484141761137a5761137a6152f9565b5f8060408385031215615664575f80fd5b82519150602083015161567681614bcc565b809150509250929050565b5f60208083525f845461569381614ede565b806020870152604060018084165f81146156b457600181146156d0576156fd565b60ff19851660408a0152604084151560051b8a010195506156fd565b895f5260205f205f5b858110156156f45781548b82018601529083019088016156d9565b8a016040019650505b509398975050505050505050565b818382375f9101908152919050565b5f806040838503121561572b575f80fd5b82516001600160401b0380821115615741575f80fd5b9084019060608287031215615754575f80fd5b61575c614f32565b82518152602083015161576e8161507b565b6020820152604083015182811115615784575f80fd5b61579088828601614fd2565b604083015250809450505050602083015161567681614bcc565b5f80604083850312156157bb575f80fd5b82519150602083015161567681614cdd565b8181038181111561137a5761137a6152f9565b5f602082840312156157f0575f80fd5b611c9982614d0c565b5f60208284031215615809575f80fd5b813560ff81168114611c99575f80fd5b5f8235603e19833603018112615117575f80fd5b5f808335601e19843603018112615842575f80fd5b8301803591506001600160401b0382111561585b575f80fd5b602001915036819003821315613950575f80fd5b602081525f615019602083018486615383565b5f60408236031215615892575f80fd5b61589a614f5a565b6158a383614bd9565b81526020808401356001600160401b03808211156158bf575f80fd5b9085019036601f8301126158d1575f80fd5b8135818111156158e3576158e3614f1e565b8060051b91506158f4848301614f7c565b818152918301840191848101903684111561590d575f80fd5b938501935b8385101561593757843592506159278361507b565b8282529385019390850190615912565b94860194909452509295945050505050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b8083101561298c5784516001600160a01b0316825293830193600192909201919083019061597b565b60208152815160208201525f602083015160e060408401526159ca610100840182614dc8565b90506040840151601f19808584030160608601526159e88383614dc8565b92506001600160401b03606087015116608086015260808601519150808584030160a0860152615a188383615949565b925060a08601519150808584030160c086015250615a368282615949565b91505060c0840151615a5360e08501826001600160401b03169052565b509392505050565b5f8060408385031215615a6c575f80fd5b8251915060208301516001600160401b03811115615a88575f80fd5b615a9485828601614fd2565b9150509250929050565b5f6001600160401b03808316818103615338576153386152f956fe4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d054317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d00e92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00e92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb07a2646970667358221220078679fb2796d8f420467e177ef948de40792fc75c62f7e4a05a7c341da68aae64736f6c63430008190033",
}

// NativeTokenStakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenStakingManagerMetaData.ABI instead.
var NativeTokenStakingManagerABI = NativeTokenStakingManagerMetaData.ABI

// NativeTokenStakingManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenStakingManagerMetaData.Bin instead.
var NativeTokenStakingManagerBin = NativeTokenStakingManagerMetaData.Bin

// DeployNativeTokenStakingManager deploys a new Ethereum contract, binding an instance of NativeTokenStakingManager to it.
func DeployNativeTokenStakingManager(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *NativeTokenStakingManager, error) {
	parsed, err := NativeTokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	validatorMessagesAddr, _, _, _ := DeployValidatorMessages(auth, backend)
	NativeTokenStakingManagerBin = strings.ReplaceAll(NativeTokenStakingManagerBin, "__$fd0c147b4031eef6079b0498cbafa865f0$__", validatorMessagesAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenStakingManagerBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenStakingManager{NativeTokenStakingManagerCaller: NativeTokenStakingManagerCaller{contract: contract}, NativeTokenStakingManagerTransactor: NativeTokenStakingManagerTransactor{contract: contract}, NativeTokenStakingManagerFilterer: NativeTokenStakingManagerFilterer{contract: contract}}, nil
}

// NativeTokenStakingManager is an auto generated Go binding around an Ethereum contract.
type NativeTokenStakingManager struct {
	NativeTokenStakingManagerCaller     // Read-only binding to the contract
	NativeTokenStakingManagerTransactor // Write-only binding to the contract
	NativeTokenStakingManagerFilterer   // Log filterer for contract events
}

// NativeTokenStakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenStakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenStakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenStakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenStakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenStakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenStakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenStakingManagerSession struct {
	Contract     *NativeTokenStakingManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NativeTokenStakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenStakingManagerCallerSession struct {
	Contract *NativeTokenStakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// NativeTokenStakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenStakingManagerTransactorSession struct {
	Contract     *NativeTokenStakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// NativeTokenStakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenStakingManagerRaw struct {
	Contract *NativeTokenStakingManager // Generic contract binding to access the raw methods on
}

// NativeTokenStakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenStakingManagerCallerRaw struct {
	Contract *NativeTokenStakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenStakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenStakingManagerTransactorRaw struct {
	Contract *NativeTokenStakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenStakingManager creates a new instance of NativeTokenStakingManager, bound to a specific deployed contract.
func NewNativeTokenStakingManager(address common.Address, backend bind.ContractBackend) (*NativeTokenStakingManager, error) {
	contract, err := bindNativeTokenStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManager{NativeTokenStakingManagerCaller: NativeTokenStakingManagerCaller{contract: contract}, NativeTokenStakingManagerTransactor: NativeTokenStakingManagerTransactor{contract: contract}, NativeTokenStakingManagerFilterer: NativeTokenStakingManagerFilterer{contract: contract}}, nil
}

// NewNativeTokenStakingManagerCaller creates a new read-only instance of NativeTokenStakingManager, bound to a specific deployed contract.
func NewNativeTokenStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenStakingManagerCaller, error) {
	contract, err := bindNativeTokenStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerCaller{contract: contract}, nil
}

// NewNativeTokenStakingManagerTransactor creates a new write-only instance of NativeTokenStakingManager, bound to a specific deployed contract.
func NewNativeTokenStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenStakingManagerTransactor, error) {
	contract, err := bindNativeTokenStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerTransactor{contract: contract}, nil
}

// NewNativeTokenStakingManagerFilterer creates a new log filterer instance of NativeTokenStakingManager, bound to a specific deployed contract.
func NewNativeTokenStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenStakingManagerFilterer, error) {
	contract, err := bindNativeTokenStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerFilterer{contract: contract}, nil
}

// bindNativeTokenStakingManager binds a generic wrapper to an already deployed contract.
func bindNativeTokenStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenStakingManager *NativeTokenStakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenStakingManager.Contract.NativeTokenStakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenStakingManager *NativeTokenStakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.NativeTokenStakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenStakingManager *NativeTokenStakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.NativeTokenStakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenStakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) ADDRESSLENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "ADDRESS_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ADDRESSLENGTH() (uint32, error) {
	return _NativeTokenStakingManager.Contract.ADDRESSLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) ADDRESSLENGTH() (uint32, error) {
	return _NativeTokenStakingManager.Contract.ADDRESSLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) BLSPUBLICKEYLENGTH(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "BLS_PUBLIC_KEY_LENGTH")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) BLSPUBLICKEYLENGTH() (uint8, error) {
	return _NativeTokenStakingManager.Contract.BLSPUBLICKEYLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) BLSPUBLICKEYLENGTH() (uint8, error) {
	return _NativeTokenStakingManager.Contract.BLSPUBLICKEYLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) MAXIMUMCHURNPERCENTAGELIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "MAXIMUM_CHURN_PERCENTAGE_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) MAXIMUMCHURNPERCENTAGELIMIT() (uint8, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMCHURNPERCENTAGELIMIT(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) MAXIMUMCHURNPERCENTAGELIMIT() (uint8, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMCHURNPERCENTAGELIMIT(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) MAXIMUMDELEGATIONFEEBIPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "MAXIMUM_DELEGATION_FEE_BIPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) MAXIMUMREGISTRATIONEXPIRYLENGTH(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "MAXIMUM_REGISTRATION_EXPIRY_LENGTH")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) MAXIMUMREGISTRATIONEXPIRYLENGTH() (uint64, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMREGISTRATIONEXPIRYLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) MAXIMUMREGISTRATIONEXPIRYLENGTH() (uint64, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMREGISTRATIONEXPIRYLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) MAXIMUMSTAKEMULTIPLIERLIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "MAXIMUM_STAKE_MULTIPLIER_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_NativeTokenStakingManager.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) NATIVEMINTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "NATIVE_MINTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenStakingManager.Contract.NATIVEMINTER(&_NativeTokenStakingManager.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenStakingManager.Contract.NATIVEMINTER(&_NativeTokenStakingManager.CallOpts)
}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) POSVALIDATORMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "POS_VALIDATOR_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) POSVALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.POSVALIDATORMANAGERSTORAGELOCATION(&_NativeTokenStakingManager.CallOpts)
}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) POSVALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.POSVALIDATORMANAGERSTORAGELOCATION(&_NativeTokenStakingManager.CallOpts)
}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) PCHAINBLOCKCHAINID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "P_CHAIN_BLOCKCHAIN_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) PCHAINBLOCKCHAINID() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.PCHAINBLOCKCHAINID(&_NativeTokenStakingManager.CallOpts)
}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) PCHAINBLOCKCHAINID() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.PCHAINBLOCKCHAINID(&_NativeTokenStakingManager.CallOpts)
}

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) UPTIMEREWARDSTHRESHOLDPERCENTAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "UPTIME_REWARDS_THRESHOLD_PERCENTAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) UPTIMEREWARDSTHRESHOLDPERCENTAGE() (uint8, error) {
	return _NativeTokenStakingManager.Contract.UPTIMEREWARDSTHRESHOLDPERCENTAGE(&_NativeTokenStakingManager.CallOpts)
}

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) UPTIMEREWARDSTHRESHOLDPERCENTAGE() (uint8, error) {
	return _NativeTokenStakingManager.Contract.UPTIMEREWARDSTHRESHOLDPERCENTAGE(&_NativeTokenStakingManager.CallOpts)
}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) VALIDATORMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "VALIDATOR_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) VALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.VALIDATORMANAGERSTORAGELOCATION(&_NativeTokenStakingManager.CallOpts)
}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) VALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.VALIDATORMANAGERSTORAGELOCATION(&_NativeTokenStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) WARPMESSENGER() (common.Address, error) {
	return _NativeTokenStakingManager.Contract.WARPMESSENGER(&_NativeTokenStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _NativeTokenStakingManager.Contract.WARPMESSENGER(&_NativeTokenStakingManager.CallOpts)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) GetValidator(opts *bind.CallOpts, validationID [32]byte) (Validator, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "getValidator", validationID)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _NativeTokenStakingManager.Contract.GetValidator(&_NativeTokenStakingManager.CallOpts, validationID)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _NativeTokenStakingManager.Contract.GetValidator(&_NativeTokenStakingManager.CallOpts, validationID)
}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) GetWeight(opts *bind.CallOpts, validationID [32]byte) (uint64, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "getWeight", validationID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) GetWeight(validationID [32]byte) (uint64, error) {
	return _NativeTokenStakingManager.Contract.GetWeight(&_NativeTokenStakingManager.CallOpts, validationID)
}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) GetWeight(validationID [32]byte) (uint64, error) {
	return _NativeTokenStakingManager.Contract.GetWeight(&_NativeTokenStakingManager.CallOpts, validationID)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) RegisteredValidators(opts *bind.CallOpts, nodeID []byte) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "registeredValidators", nodeID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) RegisteredValidators(nodeID []byte) ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.RegisteredValidators(&_NativeTokenStakingManager.CallOpts, nodeID)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) RegisteredValidators(nodeID []byte) ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.RegisteredValidators(&_NativeTokenStakingManager.CallOpts, nodeID)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) pure returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) ValueToWeight(opts *bind.CallOpts, value *big.Int) (uint64, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "valueToWeight", value)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) pure returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _NativeTokenStakingManager.Contract.ValueToWeight(&_NativeTokenStakingManager.CallOpts, value)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) pure returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _NativeTokenStakingManager.Contract.ValueToWeight(&_NativeTokenStakingManager.CallOpts, value)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) pure returns(uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) WeightToValue(opts *bind.CallOpts, weight uint64) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "weightToValue", weight)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) pure returns(uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _NativeTokenStakingManager.Contract.WeightToValue(&_NativeTokenStakingManager.CallOpts, weight)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) pure returns(uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _NativeTokenStakingManager.Contract.WeightToValue(&_NativeTokenStakingManager.CallOpts, weight)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ClaimDelegationFees(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "claimDelegationFees", validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ClaimDelegationFees(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ClaimDelegationFees(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ClaimValidationRewards(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "claimValidationRewards", validationID, messageIndex)
}

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ClaimValidationRewards(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ClaimValidationRewards(&_NativeTokenStakingManager.TransactOpts, validationID, messageIndex)
}

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ClaimValidationRewards(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ClaimValidationRewards(&_NativeTokenStakingManager.TransactOpts, validationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeDelegatorRegistration", delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeEndDelegation", delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteEndValidation(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeEndValidation", messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteEndValidation(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteEndValidation(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteValidatorRegistration(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteValidatorRegistration(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ForceInitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "forceInitializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitializeEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitializeEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ForceInitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "forceInitializeEndValidation", validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitializeEndValidation(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitializeEndValidation(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0x20e55565.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,address) settings) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) Initialize(opts *bind.TransactOpts, settings PoSValidatorManagerSettings) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initialize", settings)
}

// Initialize is a paid mutator transaction binding the contract method 0x20e55565.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,address) settings) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) Initialize(settings PoSValidatorManagerSettings) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.Initialize(&_NativeTokenStakingManager.TransactOpts, settings)
}

// Initialize is a paid mutator transaction binding the contract method 0x20e55565.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,address) settings) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) Initialize(settings PoSValidatorManagerSettings) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.Initialize(&_NativeTokenStakingManager.TransactOpts, settings)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0xc599e24f.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeDelegatorRegistration(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeDelegatorRegistration", validationID)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0xc599e24f.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeDelegatorRegistration(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0xc599e24f.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeDelegatorRegistration(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeEndValidation", validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeEndValidation(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeEndValidation(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0xaf2f5feb.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeValidatorRegistration(opts *bind.TransactOpts, registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeValidatorRegistration", registrationInput, delegationFeeBips, minStakeDuration)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0xaf2f5feb.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeValidatorRegistration(registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeValidatorRegistration(&_NativeTokenStakingManager.TransactOpts, registrationInput, delegationFeeBips, minStakeDuration)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0xaf2f5feb.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeValidatorRegistration(registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeValidatorRegistration(&_NativeTokenStakingManager.TransactOpts, registrationInput, delegationFeeBips, minStakeDuration)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeValidatorSet(opts *bind.TransactOpts, subnetConversionData SubnetConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeValidatorSet", subnetConversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeValidatorSet(subnetConversionData SubnetConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeValidatorSet(&_NativeTokenStakingManager.TransactOpts, subnetConversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeValidatorSet(subnetConversionData SubnetConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeValidatorSet(&_NativeTokenStakingManager.TransactOpts, subnetConversionData, messageIndex)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ResendEndValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "resendEndValidatorMessage", validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendEndValidatorMessage(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendEndValidatorMessage(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ResendRegisterValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "resendRegisterValidatorMessage", validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendRegisterValidatorMessage(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendRegisterValidatorMessage(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ResendUpdateDelegation(opts *bind.TransactOpts, delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "resendUpdateDelegation", delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendUpdateDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendUpdateDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) SubmitUptimeProof(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "submitUptimeProof", validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.SubmitUptimeProof(&_NativeTokenStakingManager.TransactOpts, validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.SubmitUptimeProof(&_NativeTokenStakingManager.TransactOpts, validationID, messageIndex)
}

// NativeTokenStakingManagerDelegationEndedIterator is returned from FilterDelegationEnded and is used to iterate over the raw logs and unpacked data for DelegationEnded events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegationEndedIterator struct {
	Event *NativeTokenStakingManagerDelegationEnded // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerDelegationEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerDelegationEnded)
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
		it.Event = new(NativeTokenStakingManagerDelegationEnded)
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
func (it *NativeTokenStakingManagerDelegationEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerDelegationEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerDelegationEnded represents a DelegationEnded event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegationEnded struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Rewards      *big.Int
	Fees         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegationEnded is a free log retrieval operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterDelegationEnded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*NativeTokenStakingManagerDelegationEndedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerDelegationEndedIterator{contract: _NativeTokenStakingManager.contract, event: "DelegationEnded", logs: logs, sub: sub}, nil
}

// WatchDelegationEnded is a free log subscription operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchDelegationEnded(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerDelegationEnded, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerDelegationEnded)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
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

// ParseDelegationEnded is a log parse operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseDelegationEnded(log types.Log) (*NativeTokenStakingManagerDelegationEnded, error) {
	event := new(NativeTokenStakingManagerDelegationEnded)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerDelegatorAddedIterator is returned from FilterDelegatorAdded and is used to iterate over the raw logs and unpacked data for DelegatorAdded events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorAddedIterator struct {
	Event *NativeTokenStakingManagerDelegatorAdded // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerDelegatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerDelegatorAdded)
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
		it.Event = new(NativeTokenStakingManagerDelegatorAdded)
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
func (it *NativeTokenStakingManagerDelegatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerDelegatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerDelegatorAdded represents a DelegatorAdded event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorAdded struct {
	DelegationID       [32]byte
	ValidationID       [32]byte
	DelegatorAddress   common.Address
	Nonce              uint64
	ValidatorWeight    uint64
	DelegatorWeight    uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDelegatorAdded is a free log retrieval operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterDelegatorAdded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (*NativeTokenStakingManagerDelegatorAddedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerDelegatorAddedIterator{contract: _NativeTokenStakingManager.contract, event: "DelegatorAdded", logs: logs, sub: sub}, nil
}

// WatchDelegatorAdded is a free log subscription operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchDelegatorAdded(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerDelegatorAdded, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerDelegatorAdded)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
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

// ParseDelegatorAdded is a log parse operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseDelegatorAdded(log types.Log) (*NativeTokenStakingManagerDelegatorAdded, error) {
	event := new(NativeTokenStakingManagerDelegatorAdded)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerDelegatorRegisteredIterator is returned from FilterDelegatorRegistered and is used to iterate over the raw logs and unpacked data for DelegatorRegistered events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorRegisteredIterator struct {
	Event *NativeTokenStakingManagerDelegatorRegistered // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerDelegatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerDelegatorRegistered)
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
		it.Event = new(NativeTokenStakingManagerDelegatorRegistered)
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
func (it *NativeTokenStakingManagerDelegatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerDelegatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerDelegatorRegistered represents a DelegatorRegistered event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorRegistered struct {
	DelegationID [32]byte
	ValidationID [32]byte
	StartTime    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRegistered is a free log retrieval operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterDelegatorRegistered(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*NativeTokenStakingManagerDelegatorRegisteredIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerDelegatorRegisteredIterator{contract: _NativeTokenStakingManager.contract, event: "DelegatorRegistered", logs: logs, sub: sub}, nil
}

// WatchDelegatorRegistered is a free log subscription operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchDelegatorRegistered(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerDelegatorRegistered, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerDelegatorRegistered)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
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

// ParseDelegatorRegistered is a log parse operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseDelegatorRegistered(log types.Log) (*NativeTokenStakingManagerDelegatorRegistered, error) {
	event := new(NativeTokenStakingManagerDelegatorRegistered)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerDelegatorRemovalInitializedIterator is returned from FilterDelegatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for DelegatorRemovalInitialized events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorRemovalInitializedIterator struct {
	Event *NativeTokenStakingManagerDelegatorRemovalInitialized // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerDelegatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerDelegatorRemovalInitialized)
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
		it.Event = new(NativeTokenStakingManagerDelegatorRemovalInitialized)
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
func (it *NativeTokenStakingManagerDelegatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerDelegatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerDelegatorRemovalInitialized represents a DelegatorRemovalInitialized event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorRemovalInitialized struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRemovalInitialized is a free log retrieval operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterDelegatorRemovalInitialized(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*NativeTokenStakingManagerDelegatorRemovalInitializedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerDelegatorRemovalInitializedIterator{contract: _NativeTokenStakingManager.contract, event: "DelegatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchDelegatorRemovalInitialized is a free log subscription operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchDelegatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerDelegatorRemovalInitialized, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerDelegatorRemovalInitialized)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
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

// ParseDelegatorRemovalInitialized is a log parse operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseDelegatorRemovalInitialized(log types.Log) (*NativeTokenStakingManagerDelegatorRemovalInitialized, error) {
	event := new(NativeTokenStakingManagerDelegatorRemovalInitialized)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerInitialValidatorCreatedIterator is returned from FilterInitialValidatorCreated and is used to iterate over the raw logs and unpacked data for InitialValidatorCreated events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitialValidatorCreatedIterator struct {
	Event *NativeTokenStakingManagerInitialValidatorCreated // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerInitialValidatorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerInitialValidatorCreated)
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
		it.Event = new(NativeTokenStakingManagerInitialValidatorCreated)
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
func (it *NativeTokenStakingManagerInitialValidatorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerInitialValidatorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerInitialValidatorCreated represents a InitialValidatorCreated event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitialValidatorCreated struct {
	ValidationID [32]byte
	NodeID       common.Hash
	Weight       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialValidatorCreated is a free log retrieval operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterInitialValidatorCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte) (*NativeTokenStakingManagerInitialValidatorCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerInitialValidatorCreatedIterator{contract: _NativeTokenStakingManager.contract, event: "InitialValidatorCreated", logs: logs, sub: sub}, nil
}

// WatchInitialValidatorCreated is a free log subscription operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchInitialValidatorCreated(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerInitialValidatorCreated, validationID [][32]byte, nodeID [][]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerInitialValidatorCreated)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
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

// ParseInitialValidatorCreated is a log parse operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseInitialValidatorCreated(log types.Log) (*NativeTokenStakingManagerInitialValidatorCreated, error) {
	event := new(NativeTokenStakingManagerInitialValidatorCreated)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitializedIterator struct {
	Event *NativeTokenStakingManagerInitialized // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerInitialized)
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
		it.Event = new(NativeTokenStakingManagerInitialized)
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
func (it *NativeTokenStakingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerInitialized represents a Initialized event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*NativeTokenStakingManagerInitializedIterator, error) {

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerInitializedIterator{contract: _NativeTokenStakingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerInitialized)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseInitialized(log types.Log) (*NativeTokenStakingManagerInitialized, error) {
	event := new(NativeTokenStakingManagerInitialized)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerUptimeUpdatedIterator is returned from FilterUptimeUpdated and is used to iterate over the raw logs and unpacked data for UptimeUpdated events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerUptimeUpdatedIterator struct {
	Event *NativeTokenStakingManagerUptimeUpdated // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerUptimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerUptimeUpdated)
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
		it.Event = new(NativeTokenStakingManagerUptimeUpdated)
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
func (it *NativeTokenStakingManagerUptimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerUptimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerUptimeUpdated represents a UptimeUpdated event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerUptimeUpdated struct {
	ValidationID [32]byte
	Uptime       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUptimeUpdated is a free log retrieval operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterUptimeUpdated(opts *bind.FilterOpts, validationID [][32]byte) (*NativeTokenStakingManagerUptimeUpdatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerUptimeUpdatedIterator{contract: _NativeTokenStakingManager.contract, event: "UptimeUpdated", logs: logs, sub: sub}, nil
}

// WatchUptimeUpdated is a free log subscription operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchUptimeUpdated(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerUptimeUpdated, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerUptimeUpdated)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
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

// ParseUptimeUpdated is a log parse operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseUptimeUpdated(log types.Log) (*NativeTokenStakingManagerUptimeUpdated, error) {
	event := new(NativeTokenStakingManagerUptimeUpdated)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerValidationPeriodCreatedIterator is returned from FilterValidationPeriodCreated and is used to iterate over the raw logs and unpacked data for ValidationPeriodCreated events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidationPeriodCreatedIterator struct {
	Event *NativeTokenStakingManagerValidationPeriodCreated // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerValidationPeriodCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerValidationPeriodCreated)
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
		it.Event = new(NativeTokenStakingManagerValidationPeriodCreated)
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
func (it *NativeTokenStakingManagerValidationPeriodCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerValidationPeriodCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerValidationPeriodCreated represents a ValidationPeriodCreated event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidationPeriodCreated struct {
	ValidationID                [32]byte
	NodeID                      common.Hash
	RegisterValidationMessageID [32]byte
	Weight                      *big.Int
	RegistrationExpiry          uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodCreated is a free log retrieval operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterValidationPeriodCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (*NativeTokenStakingManagerValidationPeriodCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerValidationPeriodCreatedIterator{contract: _NativeTokenStakingManager.contract, event: "ValidationPeriodCreated", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodCreated is a free log subscription operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchValidationPeriodCreated(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerValidationPeriodCreated, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerValidationPeriodCreated)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
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

// ParseValidationPeriodCreated is a log parse operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseValidationPeriodCreated(log types.Log) (*NativeTokenStakingManagerValidationPeriodCreated, error) {
	event := new(NativeTokenStakingManagerValidationPeriodCreated)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerValidationPeriodEndedIterator is returned from FilterValidationPeriodEnded and is used to iterate over the raw logs and unpacked data for ValidationPeriodEnded events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidationPeriodEndedIterator struct {
	Event *NativeTokenStakingManagerValidationPeriodEnded // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerValidationPeriodEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerValidationPeriodEnded)
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
		it.Event = new(NativeTokenStakingManagerValidationPeriodEnded)
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
func (it *NativeTokenStakingManagerValidationPeriodEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerValidationPeriodEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerValidationPeriodEnded represents a ValidationPeriodEnded event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidationPeriodEnded struct {
	ValidationID [32]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodEnded is a free log retrieval operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterValidationPeriodEnded(opts *bind.FilterOpts, validationID [][32]byte, status []uint8) (*NativeTokenStakingManagerValidationPeriodEndedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerValidationPeriodEndedIterator{contract: _NativeTokenStakingManager.contract, event: "ValidationPeriodEnded", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodEnded is a free log subscription operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchValidationPeriodEnded(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerValidationPeriodEnded, validationID [][32]byte, status []uint8) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerValidationPeriodEnded)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
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

// ParseValidationPeriodEnded is a log parse operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseValidationPeriodEnded(log types.Log) (*NativeTokenStakingManagerValidationPeriodEnded, error) {
	event := new(NativeTokenStakingManagerValidationPeriodEnded)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerValidationPeriodRegisteredIterator is returned from FilterValidationPeriodRegistered and is used to iterate over the raw logs and unpacked data for ValidationPeriodRegistered events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidationPeriodRegisteredIterator struct {
	Event *NativeTokenStakingManagerValidationPeriodRegistered // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerValidationPeriodRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerValidationPeriodRegistered)
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
		it.Event = new(NativeTokenStakingManagerValidationPeriodRegistered)
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
func (it *NativeTokenStakingManagerValidationPeriodRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerValidationPeriodRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerValidationPeriodRegistered represents a ValidationPeriodRegistered event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidationPeriodRegistered struct {
	ValidationID [32]byte
	Weight       *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodRegistered is a free log retrieval operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterValidationPeriodRegistered(opts *bind.FilterOpts, validationID [][32]byte) (*NativeTokenStakingManagerValidationPeriodRegisteredIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerValidationPeriodRegisteredIterator{contract: _NativeTokenStakingManager.contract, event: "ValidationPeriodRegistered", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodRegistered is a free log subscription operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchValidationPeriodRegistered(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerValidationPeriodRegistered, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerValidationPeriodRegistered)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
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

// ParseValidationPeriodRegistered is a log parse operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseValidationPeriodRegistered(log types.Log) (*NativeTokenStakingManagerValidationPeriodRegistered, error) {
	event := new(NativeTokenStakingManagerValidationPeriodRegistered)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerValidationRewardsClaimedIterator is returned from FilterValidationRewardsClaimed and is used to iterate over the raw logs and unpacked data for ValidationRewardsClaimed events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidationRewardsClaimedIterator struct {
	Event *NativeTokenStakingManagerValidationRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerValidationRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerValidationRewardsClaimed)
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
		it.Event = new(NativeTokenStakingManagerValidationRewardsClaimed)
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
func (it *NativeTokenStakingManagerValidationRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerValidationRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerValidationRewardsClaimed represents a ValidationRewardsClaimed event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidationRewardsClaimed struct {
	ValidationID [32]byte
	Reward       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationRewardsClaimed is a free log retrieval operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterValidationRewardsClaimed(opts *bind.FilterOpts, validationID [][32]byte, reward []*big.Int) (*NativeTokenStakingManagerValidationRewardsClaimedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "ValidationRewardsClaimed", validationIDRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerValidationRewardsClaimedIterator{contract: _NativeTokenStakingManager.contract, event: "ValidationRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchValidationRewardsClaimed is a free log subscription operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchValidationRewardsClaimed(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerValidationRewardsClaimed, validationID [][32]byte, reward []*big.Int) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "ValidationRewardsClaimed", validationIDRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerValidationRewardsClaimed)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidationRewardsClaimed", log); err != nil {
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

// ParseValidationRewardsClaimed is a log parse operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseValidationRewardsClaimed(log types.Log) (*NativeTokenStakingManagerValidationRewardsClaimed, error) {
	event := new(NativeTokenStakingManagerValidationRewardsClaimed)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidationRewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerValidatorRemovalInitializedIterator is returned from FilterValidatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for ValidatorRemovalInitialized events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidatorRemovalInitializedIterator struct {
	Event *NativeTokenStakingManagerValidatorRemovalInitialized // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerValidatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerValidatorRemovalInitialized)
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
		it.Event = new(NativeTokenStakingManagerValidatorRemovalInitialized)
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
func (it *NativeTokenStakingManagerValidatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerValidatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerValidatorRemovalInitialized represents a ValidatorRemovalInitialized event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidatorRemovalInitialized struct {
	ValidationID       [32]byte
	SetWeightMessageID [32]byte
	Weight             *big.Int
	EndTime            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemovalInitialized is a free log retrieval operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterValidatorRemovalInitialized(opts *bind.FilterOpts, validationID [][32]byte, setWeightMessageID [][32]byte) (*NativeTokenStakingManagerValidatorRemovalInitializedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerValidatorRemovalInitializedIterator{contract: _NativeTokenStakingManager.contract, event: "ValidatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchValidatorRemovalInitialized is a free log subscription operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchValidatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerValidatorRemovalInitialized, validationID [][32]byte, setWeightMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerValidatorRemovalInitialized)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
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

// ParseValidatorRemovalInitialized is a log parse operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseValidatorRemovalInitialized(log types.Log) (*NativeTokenStakingManagerValidatorRemovalInitialized, error) {
	event := new(NativeTokenStakingManagerValidatorRemovalInitialized)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerValidatorWeightUpdateIterator is returned from FilterValidatorWeightUpdate and is used to iterate over the raw logs and unpacked data for ValidatorWeightUpdate events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidatorWeightUpdateIterator struct {
	Event *NativeTokenStakingManagerValidatorWeightUpdate // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerValidatorWeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerValidatorWeightUpdate)
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
		it.Event = new(NativeTokenStakingManagerValidatorWeightUpdate)
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
func (it *NativeTokenStakingManagerValidatorWeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerValidatorWeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerValidatorWeightUpdate represents a ValidatorWeightUpdate event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidatorWeightUpdate struct {
	ValidationID       [32]byte
	Nonce              uint64
	ValidatorWeight    uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorWeightUpdate is a free log retrieval operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterValidatorWeightUpdate(opts *bind.FilterOpts, validationID [][32]byte, nonce []uint64) (*NativeTokenStakingManagerValidatorWeightUpdateIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerValidatorWeightUpdateIterator{contract: _NativeTokenStakingManager.contract, event: "ValidatorWeightUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorWeightUpdate is a free log subscription operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchValidatorWeightUpdate(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerValidatorWeightUpdate, validationID [][32]byte, nonce []uint64) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerValidatorWeightUpdate)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
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

// ParseValidatorWeightUpdate is a log parse operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseValidatorWeightUpdate(log types.Log) (*NativeTokenStakingManagerValidatorWeightUpdate, error) {
	event := new(NativeTokenStakingManagerValidatorWeightUpdate)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoSValidatorManagerMetaData contains all meta data concerning the PoSValidatorManager contract.
var PoSValidatorManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"InvalidBLSKeyLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"}],\"name\":\"InvalidDelegationFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidDelegationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumDelegatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidDelegatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializationStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\"}],\"name\":\"InvalidMaximumChurnPercentage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"}],\"name\":\"InvalidMinStakeDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"InvalidNodeID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addressesLength\",\"type\":\"uint256\"}],\"name\":\"InvalidPChainOwnerThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"InvalidRegistrationExpiry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"InvalidStakeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"}],\"name\":\"InvalidStakeMultiplier\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"encodedSubnetConversionID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedSubnetConversionID\",\"type\":\"bytes32\"}],\"name\":\"InvalidSubnetConversionID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"}],\"name\":\"InvalidValidatorManagerAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidatorManagerBlockchainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidValidatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWarpMessage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"InvalidWarpOriginSenderAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidWarpSourceChainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"churnAmount\",\"type\":\"uint64\"}],\"name\":\"MaxChurnRateExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newValidatorWeight\",\"type\":\"uint64\"}],\"name\":\"MaxWeightExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"name\":\"MinStakeDurationNotPassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PChainOwnerAddressesNotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UnauthorizedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"validRegistration\",\"type\":\"bool\"}],\"name\":\"UnexpectedRegistrationStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorNotPoS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"DelegationEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"delegatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"DelegatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"InitialValidatorCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"UptimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"ValidationPeriodCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ValidationPeriodEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidationPeriodRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"ValidationRewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"ValidatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorWeightUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_LENGTH\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLS_PUBLIC_KEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_CHURN_PERCENTAGE_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_REGISTRATION_EXPIRY_LENGTH\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POS_VALIDATOR_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"P_CHAIN_BLOCKCHAIN_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPTIME_REWARDS_THRESHOLD_PERCENTAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"claimDelegationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"claimValidationRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeDelegatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"startingWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"messageNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endedAt\",\"type\":\"uint64\"}],\"internalType\":\"structValidator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structSubnetConversionData\",\"name\":\"subnetConversionData\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"registeredValidators\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendEndValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendRegisterValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"resendUpdateDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"submitUptimeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"valueToWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"weightToValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// PoSValidatorManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PoSValidatorManagerMetaData.ABI instead.
var PoSValidatorManagerABI = PoSValidatorManagerMetaData.ABI

// PoSValidatorManager is an auto generated Go binding around an Ethereum contract.
type PoSValidatorManager struct {
	PoSValidatorManagerCaller     // Read-only binding to the contract
	PoSValidatorManagerTransactor // Write-only binding to the contract
	PoSValidatorManagerFilterer   // Log filterer for contract events
}

// PoSValidatorManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoSValidatorManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoSValidatorManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoSValidatorManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoSValidatorManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoSValidatorManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoSValidatorManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoSValidatorManagerSession struct {
	Contract     *PoSValidatorManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PoSValidatorManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoSValidatorManagerCallerSession struct {
	Contract *PoSValidatorManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PoSValidatorManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoSValidatorManagerTransactorSession struct {
	Contract     *PoSValidatorManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PoSValidatorManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoSValidatorManagerRaw struct {
	Contract *PoSValidatorManager // Generic contract binding to access the raw methods on
}

// PoSValidatorManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoSValidatorManagerCallerRaw struct {
	Contract *PoSValidatorManagerCaller // Generic read-only contract binding to access the raw methods on
}

// PoSValidatorManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoSValidatorManagerTransactorRaw struct {
	Contract *PoSValidatorManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoSValidatorManager creates a new instance of PoSValidatorManager, bound to a specific deployed contract.
func NewPoSValidatorManager(address common.Address, backend bind.ContractBackend) (*PoSValidatorManager, error) {
	contract, err := bindPoSValidatorManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManager{PoSValidatorManagerCaller: PoSValidatorManagerCaller{contract: contract}, PoSValidatorManagerTransactor: PoSValidatorManagerTransactor{contract: contract}, PoSValidatorManagerFilterer: PoSValidatorManagerFilterer{contract: contract}}, nil
}

// NewPoSValidatorManagerCaller creates a new read-only instance of PoSValidatorManager, bound to a specific deployed contract.
func NewPoSValidatorManagerCaller(address common.Address, caller bind.ContractCaller) (*PoSValidatorManagerCaller, error) {
	contract, err := bindPoSValidatorManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManagerCaller{contract: contract}, nil
}

// NewPoSValidatorManagerTransactor creates a new write-only instance of PoSValidatorManager, bound to a specific deployed contract.
func NewPoSValidatorManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PoSValidatorManagerTransactor, error) {
	contract, err := bindPoSValidatorManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManagerTransactor{contract: contract}, nil
}

// NewPoSValidatorManagerFilterer creates a new log filterer instance of PoSValidatorManager, bound to a specific deployed contract.
func NewPoSValidatorManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PoSValidatorManagerFilterer, error) {
	contract, err := bindPoSValidatorManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManagerFilterer{contract: contract}, nil
}

// bindPoSValidatorManager binds a generic wrapper to an already deployed contract.
func bindPoSValidatorManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoSValidatorManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoSValidatorManager *PoSValidatorManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoSValidatorManager.Contract.PoSValidatorManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoSValidatorManager *PoSValidatorManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.PoSValidatorManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoSValidatorManager *PoSValidatorManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.PoSValidatorManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoSValidatorManager *PoSValidatorManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoSValidatorManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoSValidatorManager *PoSValidatorManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoSValidatorManager *PoSValidatorManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_PoSValidatorManager *PoSValidatorManagerCaller) ADDRESSLENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PoSValidatorManager.contract.Call(opts, &out, "ADDRESS_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_PoSValidatorManager *PoSValidatorManagerSession) ADDRESSLENGTH() (uint32, error) {
	return _PoSValidatorManager.Contract.ADDRESSLENGTH(&_PoSValidatorManager.CallOpts)
}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_PoSValidatorManager *PoSValidatorManagerCallerSession) ADDRESSLENGTH() (uint32, error) {
	return _PoSValidatorManager.Contract.ADDRESSLENGTH(&_PoSValidatorManager.CallOpts)
}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_PoSValidatorManager *PoSValidatorManagerCaller) BLSPUBLICKEYLENGTH(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PoSValidatorManager.contract.Call(opts, &out, "BLS_PUBLIC_KEY_LENGTH")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_PoSValidatorManager *PoSValidatorManagerSession) BLSPUBLICKEYLENGTH() (uint8, error) {
	return _PoSValidatorManager.Contract.BLSPUBLICKEYLENGTH(&_PoSValidatorManager.CallOpts)
}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_PoSValidatorManager *PoSValidatorManagerCallerSession) BLSPUBLICKEYLENGTH() (uint8, error) {
	return _PoSValidatorManager.Contract.BLSPUBLICKEYLENGTH(&_PoSValidatorManager.CallOpts)
}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_PoSValidatorManager *PoSValidatorManagerCaller) MAXIMUMCHURNPERCENTAGELIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PoSValidatorManager.contract.Call(opts, &out, "MAXIMUM_CHURN_PERCENTAGE_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_PoSValidatorManager *PoSValidatorManagerSession) MAXIMUMCHURNPERCENTAGELIMIT() (uint8, error) {
	return _PoSValidatorManager.Contract.MAXIMUMCHURNPERCENTAGELIMIT(&_PoSValidatorManager.CallOpts)
}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_PoSValidatorManager *PoSValidatorManagerCallerSession) MAXIMUMCHURNPERCENTAGELIMIT() (uint8, error) {
	return _PoSValidatorManager.Contract.MAXIMUMCHURNPERCENTAGELIMIT(&_PoSValidatorManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_PoSValidatorManager *PoSValidatorManagerCaller) MAXIMUMDELEGATIONFEEBIPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PoSValidatorManager.contract.Call(opts, &out, "MAXIMUM_DELEGATION_FEE_BIPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_PoSValidatorManager *PoSValidatorManagerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _PoSValidatorManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_PoSValidatorManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_PoSValidatorManager *PoSValidatorManagerCallerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _PoSValidatorManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_PoSValidatorManager.CallOpts)
}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_PoSValidatorManager *PoSValidatorManagerCaller) MAXIMUMREGISTRATIONEXPIRYLENGTH(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PoSValidatorManager.contract.Call(opts, &out, "MAXIMUM_REGISTRATION_EXPIRY_LENGTH")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_PoSValidatorManager *PoSValidatorManagerSession) MAXIMUMREGISTRATIONEXPIRYLENGTH() (uint64, error) {
	return _PoSValidatorManager.Contract.MAXIMUMREGISTRATIONEXPIRYLENGTH(&_PoSValidatorManager.CallOpts)
}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_PoSValidatorManager *PoSValidatorManagerCallerSession) MAXIMUMREGISTRATIONEXPIRYLENGTH() (uint64, error) {
	return _PoSValidatorManager.Contract.MAXIMUMREGISTRATIONEXPIRYLENGTH(&_PoSValidatorManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_PoSValidatorManager *PoSValidatorManagerCaller) MAXIMUMSTAKEMULTIPLIERLIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PoSValidatorManager.contract.Call(opts, &out, "MAXIMUM_STAKE_MULTIPLIER_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_PoSValidatorManager *PoSValidatorManagerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _PoSValidatorManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_PoSValidatorManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_PoSValidatorManager *PoSValidatorManagerCallerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _PoSValidatorManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_PoSValidatorManager.CallOpts)
}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_PoSValidatorManager *PoSValidatorManagerCaller) POSVALIDATORMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoSValidatorManager.contract.Call(opts, &out, "POS_VALIDATOR_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_PoSValidatorManager *PoSValidatorManagerSession) POSVALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _PoSValidatorManager.Contract.POSVALIDATORMANAGERSTORAGELOCATION(&_PoSValidatorManager.CallOpts)
}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_PoSValidatorManager *PoSValidatorManagerCallerSession) POSVALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _PoSValidatorManager.Contract.POSVALIDATORMANAGERSTORAGELOCATION(&_PoSValidatorManager.CallOpts)
}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_PoSValidatorManager *PoSValidatorManagerCaller) PCHAINBLOCKCHAINID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoSValidatorManager.contract.Call(opts, &out, "P_CHAIN_BLOCKCHAIN_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_PoSValidatorManager *PoSValidatorManagerSession) PCHAINBLOCKCHAINID() ([32]byte, error) {
	return _PoSValidatorManager.Contract.PCHAINBLOCKCHAINID(&_PoSValidatorManager.CallOpts)
}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_PoSValidatorManager *PoSValidatorManagerCallerSession) PCHAINBLOCKCHAINID() ([32]byte, error) {
	return _PoSValidatorManager.Contract.PCHAINBLOCKCHAINID(&_PoSValidatorManager.CallOpts)
}

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_PoSValidatorManager *PoSValidatorManagerCaller) UPTIMEREWARDSTHRESHOLDPERCENTAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PoSValidatorManager.contract.Call(opts, &out, "UPTIME_REWARDS_THRESHOLD_PERCENTAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_PoSValidatorManager *PoSValidatorManagerSession) UPTIMEREWARDSTHRESHOLDPERCENTAGE() (uint8, error) {
	return _PoSValidatorManager.Contract.UPTIMEREWARDSTHRESHOLDPERCENTAGE(&_PoSValidatorManager.CallOpts)
}

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_PoSValidatorManager *PoSValidatorManagerCallerSession) UPTIMEREWARDSTHRESHOLDPERCENTAGE() (uint8, error) {
	return _PoSValidatorManager.Contract.UPTIMEREWARDSTHRESHOLDPERCENTAGE(&_PoSValidatorManager.CallOpts)
}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_PoSValidatorManager *PoSValidatorManagerCaller) VALIDATORMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoSValidatorManager.contract.Call(opts, &out, "VALIDATOR_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_PoSValidatorManager *PoSValidatorManagerSession) VALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _PoSValidatorManager.Contract.VALIDATORMANAGERSTORAGELOCATION(&_PoSValidatorManager.CallOpts)
}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_PoSValidatorManager *PoSValidatorManagerCallerSession) VALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _PoSValidatorManager.Contract.VALIDATORMANAGERSTORAGELOCATION(&_PoSValidatorManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_PoSValidatorManager *PoSValidatorManagerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoSValidatorManager.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_PoSValidatorManager *PoSValidatorManagerSession) WARPMESSENGER() (common.Address, error) {
	return _PoSValidatorManager.Contract.WARPMESSENGER(&_PoSValidatorManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_PoSValidatorManager *PoSValidatorManagerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _PoSValidatorManager.Contract.WARPMESSENGER(&_PoSValidatorManager.CallOpts)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_PoSValidatorManager *PoSValidatorManagerCaller) GetValidator(opts *bind.CallOpts, validationID [32]byte) (Validator, error) {
	var out []interface{}
	err := _PoSValidatorManager.contract.Call(opts, &out, "getValidator", validationID)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_PoSValidatorManager *PoSValidatorManagerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _PoSValidatorManager.Contract.GetValidator(&_PoSValidatorManager.CallOpts, validationID)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_PoSValidatorManager *PoSValidatorManagerCallerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _PoSValidatorManager.Contract.GetValidator(&_PoSValidatorManager.CallOpts, validationID)
}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_PoSValidatorManager *PoSValidatorManagerCaller) GetWeight(opts *bind.CallOpts, validationID [32]byte) (uint64, error) {
	var out []interface{}
	err := _PoSValidatorManager.contract.Call(opts, &out, "getWeight", validationID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_PoSValidatorManager *PoSValidatorManagerSession) GetWeight(validationID [32]byte) (uint64, error) {
	return _PoSValidatorManager.Contract.GetWeight(&_PoSValidatorManager.CallOpts, validationID)
}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_PoSValidatorManager *PoSValidatorManagerCallerSession) GetWeight(validationID [32]byte) (uint64, error) {
	return _PoSValidatorManager.Contract.GetWeight(&_PoSValidatorManager.CallOpts, validationID)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_PoSValidatorManager *PoSValidatorManagerCaller) RegisteredValidators(opts *bind.CallOpts, nodeID []byte) ([32]byte, error) {
	var out []interface{}
	err := _PoSValidatorManager.contract.Call(opts, &out, "registeredValidators", nodeID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_PoSValidatorManager *PoSValidatorManagerSession) RegisteredValidators(nodeID []byte) ([32]byte, error) {
	return _PoSValidatorManager.Contract.RegisteredValidators(&_PoSValidatorManager.CallOpts, nodeID)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_PoSValidatorManager *PoSValidatorManagerCallerSession) RegisteredValidators(nodeID []byte) ([32]byte, error) {
	return _PoSValidatorManager.Contract.RegisteredValidators(&_PoSValidatorManager.CallOpts, nodeID)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) pure returns(uint64)
func (_PoSValidatorManager *PoSValidatorManagerCaller) ValueToWeight(opts *bind.CallOpts, value *big.Int) (uint64, error) {
	var out []interface{}
	err := _PoSValidatorManager.contract.Call(opts, &out, "valueToWeight", value)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) pure returns(uint64)
func (_PoSValidatorManager *PoSValidatorManagerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _PoSValidatorManager.Contract.ValueToWeight(&_PoSValidatorManager.CallOpts, value)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) pure returns(uint64)
func (_PoSValidatorManager *PoSValidatorManagerCallerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _PoSValidatorManager.Contract.ValueToWeight(&_PoSValidatorManager.CallOpts, value)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) pure returns(uint256)
func (_PoSValidatorManager *PoSValidatorManagerCaller) WeightToValue(opts *bind.CallOpts, weight uint64) (*big.Int, error) {
	var out []interface{}
	err := _PoSValidatorManager.contract.Call(opts, &out, "weightToValue", weight)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) pure returns(uint256)
func (_PoSValidatorManager *PoSValidatorManagerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _PoSValidatorManager.Contract.WeightToValue(&_PoSValidatorManager.CallOpts, weight)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) pure returns(uint256)
func (_PoSValidatorManager *PoSValidatorManagerCallerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _PoSValidatorManager.Contract.WeightToValue(&_PoSValidatorManager.CallOpts, weight)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactor) ClaimDelegationFees(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _PoSValidatorManager.contract.Transact(opts, "claimDelegationFees", validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_PoSValidatorManager *PoSValidatorManagerSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.ClaimDelegationFees(&_PoSValidatorManager.TransactOpts, validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactorSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.ClaimDelegationFees(&_PoSValidatorManager.TransactOpts, validationID)
}

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactor) ClaimValidationRewards(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.contract.Transact(opts, "claimValidationRewards", validationID, messageIndex)
}

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerSession) ClaimValidationRewards(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.ClaimValidationRewards(&_PoSValidatorManager.TransactOpts, validationID, messageIndex)
}

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactorSession) ClaimValidationRewards(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.ClaimValidationRewards(&_PoSValidatorManager.TransactOpts, validationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.contract.Transact(opts, "completeDelegatorRegistration", delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.CompleteDelegatorRegistration(&_PoSValidatorManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactorSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.CompleteDelegatorRegistration(&_PoSValidatorManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactor) CompleteEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.contract.Transact(opts, "completeEndDelegation", delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.CompleteEndDelegation(&_PoSValidatorManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactorSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.CompleteEndDelegation(&_PoSValidatorManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactor) CompleteEndValidation(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.contract.Transact(opts, "completeEndValidation", messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.CompleteEndValidation(&_PoSValidatorManager.TransactOpts, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactorSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.CompleteEndValidation(&_PoSValidatorManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.CompleteValidatorRegistration(&_PoSValidatorManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.CompleteValidatorRegistration(&_PoSValidatorManager.TransactOpts, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactor) ForceInitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.contract.Transact(opts, "forceInitializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.ForceInitializeEndDelegation(&_PoSValidatorManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactorSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.ForceInitializeEndDelegation(&_PoSValidatorManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactor) ForceInitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.contract.Transact(opts, "forceInitializeEndValidation", validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.ForceInitializeEndValidation(&_PoSValidatorManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactorSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.ForceInitializeEndValidation(&_PoSValidatorManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactor) InitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.contract.Transact(opts, "initializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.InitializeEndDelegation(&_PoSValidatorManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactorSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.InitializeEndDelegation(&_PoSValidatorManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactor) InitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.contract.Transact(opts, "initializeEndValidation", validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.InitializeEndValidation(&_PoSValidatorManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactorSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.InitializeEndValidation(&_PoSValidatorManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactor) InitializeValidatorSet(opts *bind.TransactOpts, subnetConversionData SubnetConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.contract.Transact(opts, "initializeValidatorSet", subnetConversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerSession) InitializeValidatorSet(subnetConversionData SubnetConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.InitializeValidatorSet(&_PoSValidatorManager.TransactOpts, subnetConversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactorSession) InitializeValidatorSet(subnetConversionData SubnetConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.InitializeValidatorSet(&_PoSValidatorManager.TransactOpts, subnetConversionData, messageIndex)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactor) ResendEndValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _PoSValidatorManager.contract.Transact(opts, "resendEndValidatorMessage", validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_PoSValidatorManager *PoSValidatorManagerSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.ResendEndValidatorMessage(&_PoSValidatorManager.TransactOpts, validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactorSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.ResendEndValidatorMessage(&_PoSValidatorManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactor) ResendRegisterValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _PoSValidatorManager.contract.Transact(opts, "resendRegisterValidatorMessage", validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_PoSValidatorManager *PoSValidatorManagerSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.ResendRegisterValidatorMessage(&_PoSValidatorManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactorSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.ResendRegisterValidatorMessage(&_PoSValidatorManager.TransactOpts, validationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactor) ResendUpdateDelegation(opts *bind.TransactOpts, delegationID [32]byte) (*types.Transaction, error) {
	return _PoSValidatorManager.contract.Transact(opts, "resendUpdateDelegation", delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_PoSValidatorManager *PoSValidatorManagerSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.ResendUpdateDelegation(&_PoSValidatorManager.TransactOpts, delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactorSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.ResendUpdateDelegation(&_PoSValidatorManager.TransactOpts, delegationID)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactor) SubmitUptimeProof(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.contract.Transact(opts, "submitUptimeProof", validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.SubmitUptimeProof(&_PoSValidatorManager.TransactOpts, validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_PoSValidatorManager *PoSValidatorManagerTransactorSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _PoSValidatorManager.Contract.SubmitUptimeProof(&_PoSValidatorManager.TransactOpts, validationID, messageIndex)
}

// PoSValidatorManagerDelegationEndedIterator is returned from FilterDelegationEnded and is used to iterate over the raw logs and unpacked data for DelegationEnded events raised by the PoSValidatorManager contract.
type PoSValidatorManagerDelegationEndedIterator struct {
	Event *PoSValidatorManagerDelegationEnded // Event containing the contract specifics and raw log

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
func (it *PoSValidatorManagerDelegationEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSValidatorManagerDelegationEnded)
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
		it.Event = new(PoSValidatorManagerDelegationEnded)
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
func (it *PoSValidatorManagerDelegationEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSValidatorManagerDelegationEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSValidatorManagerDelegationEnded represents a DelegationEnded event raised by the PoSValidatorManager contract.
type PoSValidatorManagerDelegationEnded struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Rewards      *big.Int
	Fees         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegationEnded is a free log retrieval operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) FilterDelegationEnded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*PoSValidatorManagerDelegationEndedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.FilterLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManagerDelegationEndedIterator{contract: _PoSValidatorManager.contract, event: "DelegationEnded", logs: logs, sub: sub}, nil
}

// WatchDelegationEnded is a free log subscription operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) WatchDelegationEnded(opts *bind.WatchOpts, sink chan<- *PoSValidatorManagerDelegationEnded, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.WatchLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSValidatorManagerDelegationEnded)
				if err := _PoSValidatorManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
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

// ParseDelegationEnded is a log parse operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) ParseDelegationEnded(log types.Log) (*PoSValidatorManagerDelegationEnded, error) {
	event := new(PoSValidatorManagerDelegationEnded)
	if err := _PoSValidatorManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoSValidatorManagerDelegatorAddedIterator is returned from FilterDelegatorAdded and is used to iterate over the raw logs and unpacked data for DelegatorAdded events raised by the PoSValidatorManager contract.
type PoSValidatorManagerDelegatorAddedIterator struct {
	Event *PoSValidatorManagerDelegatorAdded // Event containing the contract specifics and raw log

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
func (it *PoSValidatorManagerDelegatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSValidatorManagerDelegatorAdded)
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
		it.Event = new(PoSValidatorManagerDelegatorAdded)
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
func (it *PoSValidatorManagerDelegatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSValidatorManagerDelegatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSValidatorManagerDelegatorAdded represents a DelegatorAdded event raised by the PoSValidatorManager contract.
type PoSValidatorManagerDelegatorAdded struct {
	DelegationID       [32]byte
	ValidationID       [32]byte
	DelegatorAddress   common.Address
	Nonce              uint64
	ValidatorWeight    uint64
	DelegatorWeight    uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDelegatorAdded is a free log retrieval operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) FilterDelegatorAdded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (*PoSValidatorManagerDelegatorAddedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.FilterLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManagerDelegatorAddedIterator{contract: _PoSValidatorManager.contract, event: "DelegatorAdded", logs: logs, sub: sub}, nil
}

// WatchDelegatorAdded is a free log subscription operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) WatchDelegatorAdded(opts *bind.WatchOpts, sink chan<- *PoSValidatorManagerDelegatorAdded, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.WatchLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSValidatorManagerDelegatorAdded)
				if err := _PoSValidatorManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
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

// ParseDelegatorAdded is a log parse operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) ParseDelegatorAdded(log types.Log) (*PoSValidatorManagerDelegatorAdded, error) {
	event := new(PoSValidatorManagerDelegatorAdded)
	if err := _PoSValidatorManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoSValidatorManagerDelegatorRegisteredIterator is returned from FilterDelegatorRegistered and is used to iterate over the raw logs and unpacked data for DelegatorRegistered events raised by the PoSValidatorManager contract.
type PoSValidatorManagerDelegatorRegisteredIterator struct {
	Event *PoSValidatorManagerDelegatorRegistered // Event containing the contract specifics and raw log

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
func (it *PoSValidatorManagerDelegatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSValidatorManagerDelegatorRegistered)
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
		it.Event = new(PoSValidatorManagerDelegatorRegistered)
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
func (it *PoSValidatorManagerDelegatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSValidatorManagerDelegatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSValidatorManagerDelegatorRegistered represents a DelegatorRegistered event raised by the PoSValidatorManager contract.
type PoSValidatorManagerDelegatorRegistered struct {
	DelegationID [32]byte
	ValidationID [32]byte
	StartTime    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRegistered is a free log retrieval operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) FilterDelegatorRegistered(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*PoSValidatorManagerDelegatorRegisteredIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.FilterLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManagerDelegatorRegisteredIterator{contract: _PoSValidatorManager.contract, event: "DelegatorRegistered", logs: logs, sub: sub}, nil
}

// WatchDelegatorRegistered is a free log subscription operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) WatchDelegatorRegistered(opts *bind.WatchOpts, sink chan<- *PoSValidatorManagerDelegatorRegistered, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.WatchLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSValidatorManagerDelegatorRegistered)
				if err := _PoSValidatorManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
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

// ParseDelegatorRegistered is a log parse operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) ParseDelegatorRegistered(log types.Log) (*PoSValidatorManagerDelegatorRegistered, error) {
	event := new(PoSValidatorManagerDelegatorRegistered)
	if err := _PoSValidatorManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoSValidatorManagerDelegatorRemovalInitializedIterator is returned from FilterDelegatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for DelegatorRemovalInitialized events raised by the PoSValidatorManager contract.
type PoSValidatorManagerDelegatorRemovalInitializedIterator struct {
	Event *PoSValidatorManagerDelegatorRemovalInitialized // Event containing the contract specifics and raw log

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
func (it *PoSValidatorManagerDelegatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSValidatorManagerDelegatorRemovalInitialized)
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
		it.Event = new(PoSValidatorManagerDelegatorRemovalInitialized)
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
func (it *PoSValidatorManagerDelegatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSValidatorManagerDelegatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSValidatorManagerDelegatorRemovalInitialized represents a DelegatorRemovalInitialized event raised by the PoSValidatorManager contract.
type PoSValidatorManagerDelegatorRemovalInitialized struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRemovalInitialized is a free log retrieval operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) FilterDelegatorRemovalInitialized(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*PoSValidatorManagerDelegatorRemovalInitializedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.FilterLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManagerDelegatorRemovalInitializedIterator{contract: _PoSValidatorManager.contract, event: "DelegatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchDelegatorRemovalInitialized is a free log subscription operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) WatchDelegatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *PoSValidatorManagerDelegatorRemovalInitialized, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.WatchLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSValidatorManagerDelegatorRemovalInitialized)
				if err := _PoSValidatorManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
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

// ParseDelegatorRemovalInitialized is a log parse operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) ParseDelegatorRemovalInitialized(log types.Log) (*PoSValidatorManagerDelegatorRemovalInitialized, error) {
	event := new(PoSValidatorManagerDelegatorRemovalInitialized)
	if err := _PoSValidatorManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoSValidatorManagerInitialValidatorCreatedIterator is returned from FilterInitialValidatorCreated and is used to iterate over the raw logs and unpacked data for InitialValidatorCreated events raised by the PoSValidatorManager contract.
type PoSValidatorManagerInitialValidatorCreatedIterator struct {
	Event *PoSValidatorManagerInitialValidatorCreated // Event containing the contract specifics and raw log

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
func (it *PoSValidatorManagerInitialValidatorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSValidatorManagerInitialValidatorCreated)
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
		it.Event = new(PoSValidatorManagerInitialValidatorCreated)
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
func (it *PoSValidatorManagerInitialValidatorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSValidatorManagerInitialValidatorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSValidatorManagerInitialValidatorCreated represents a InitialValidatorCreated event raised by the PoSValidatorManager contract.
type PoSValidatorManagerInitialValidatorCreated struct {
	ValidationID [32]byte
	NodeID       common.Hash
	Weight       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialValidatorCreated is a free log retrieval operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) FilterInitialValidatorCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte) (*PoSValidatorManagerInitialValidatorCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.FilterLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManagerInitialValidatorCreatedIterator{contract: _PoSValidatorManager.contract, event: "InitialValidatorCreated", logs: logs, sub: sub}, nil
}

// WatchInitialValidatorCreated is a free log subscription operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) WatchInitialValidatorCreated(opts *bind.WatchOpts, sink chan<- *PoSValidatorManagerInitialValidatorCreated, validationID [][32]byte, nodeID [][]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.WatchLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSValidatorManagerInitialValidatorCreated)
				if err := _PoSValidatorManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
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

// ParseInitialValidatorCreated is a log parse operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) ParseInitialValidatorCreated(log types.Log) (*PoSValidatorManagerInitialValidatorCreated, error) {
	event := new(PoSValidatorManagerInitialValidatorCreated)
	if err := _PoSValidatorManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoSValidatorManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PoSValidatorManager contract.
type PoSValidatorManagerInitializedIterator struct {
	Event *PoSValidatorManagerInitialized // Event containing the contract specifics and raw log

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
func (it *PoSValidatorManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSValidatorManagerInitialized)
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
		it.Event = new(PoSValidatorManagerInitialized)
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
func (it *PoSValidatorManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSValidatorManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSValidatorManagerInitialized represents a Initialized event raised by the PoSValidatorManager contract.
type PoSValidatorManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*PoSValidatorManagerInitializedIterator, error) {

	logs, sub, err := _PoSValidatorManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManagerInitializedIterator{contract: _PoSValidatorManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PoSValidatorManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _PoSValidatorManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSValidatorManagerInitialized)
				if err := _PoSValidatorManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) ParseInitialized(log types.Log) (*PoSValidatorManagerInitialized, error) {
	event := new(PoSValidatorManagerInitialized)
	if err := _PoSValidatorManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoSValidatorManagerUptimeUpdatedIterator is returned from FilterUptimeUpdated and is used to iterate over the raw logs and unpacked data for UptimeUpdated events raised by the PoSValidatorManager contract.
type PoSValidatorManagerUptimeUpdatedIterator struct {
	Event *PoSValidatorManagerUptimeUpdated // Event containing the contract specifics and raw log

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
func (it *PoSValidatorManagerUptimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSValidatorManagerUptimeUpdated)
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
		it.Event = new(PoSValidatorManagerUptimeUpdated)
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
func (it *PoSValidatorManagerUptimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSValidatorManagerUptimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSValidatorManagerUptimeUpdated represents a UptimeUpdated event raised by the PoSValidatorManager contract.
type PoSValidatorManagerUptimeUpdated struct {
	ValidationID [32]byte
	Uptime       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUptimeUpdated is a free log retrieval operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) FilterUptimeUpdated(opts *bind.FilterOpts, validationID [][32]byte) (*PoSValidatorManagerUptimeUpdatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.FilterLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManagerUptimeUpdatedIterator{contract: _PoSValidatorManager.contract, event: "UptimeUpdated", logs: logs, sub: sub}, nil
}

// WatchUptimeUpdated is a free log subscription operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) WatchUptimeUpdated(opts *bind.WatchOpts, sink chan<- *PoSValidatorManagerUptimeUpdated, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.WatchLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSValidatorManagerUptimeUpdated)
				if err := _PoSValidatorManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
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

// ParseUptimeUpdated is a log parse operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) ParseUptimeUpdated(log types.Log) (*PoSValidatorManagerUptimeUpdated, error) {
	event := new(PoSValidatorManagerUptimeUpdated)
	if err := _PoSValidatorManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoSValidatorManagerValidationPeriodCreatedIterator is returned from FilterValidationPeriodCreated and is used to iterate over the raw logs and unpacked data for ValidationPeriodCreated events raised by the PoSValidatorManager contract.
type PoSValidatorManagerValidationPeriodCreatedIterator struct {
	Event *PoSValidatorManagerValidationPeriodCreated // Event containing the contract specifics and raw log

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
func (it *PoSValidatorManagerValidationPeriodCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSValidatorManagerValidationPeriodCreated)
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
		it.Event = new(PoSValidatorManagerValidationPeriodCreated)
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
func (it *PoSValidatorManagerValidationPeriodCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSValidatorManagerValidationPeriodCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSValidatorManagerValidationPeriodCreated represents a ValidationPeriodCreated event raised by the PoSValidatorManager contract.
type PoSValidatorManagerValidationPeriodCreated struct {
	ValidationID                [32]byte
	NodeID                      common.Hash
	RegisterValidationMessageID [32]byte
	Weight                      *big.Int
	RegistrationExpiry          uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodCreated is a free log retrieval operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) FilterValidationPeriodCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (*PoSValidatorManagerValidationPeriodCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.FilterLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManagerValidationPeriodCreatedIterator{contract: _PoSValidatorManager.contract, event: "ValidationPeriodCreated", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodCreated is a free log subscription operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) WatchValidationPeriodCreated(opts *bind.WatchOpts, sink chan<- *PoSValidatorManagerValidationPeriodCreated, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.WatchLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSValidatorManagerValidationPeriodCreated)
				if err := _PoSValidatorManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
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

// ParseValidationPeriodCreated is a log parse operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) ParseValidationPeriodCreated(log types.Log) (*PoSValidatorManagerValidationPeriodCreated, error) {
	event := new(PoSValidatorManagerValidationPeriodCreated)
	if err := _PoSValidatorManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoSValidatorManagerValidationPeriodEndedIterator is returned from FilterValidationPeriodEnded and is used to iterate over the raw logs and unpacked data for ValidationPeriodEnded events raised by the PoSValidatorManager contract.
type PoSValidatorManagerValidationPeriodEndedIterator struct {
	Event *PoSValidatorManagerValidationPeriodEnded // Event containing the contract specifics and raw log

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
func (it *PoSValidatorManagerValidationPeriodEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSValidatorManagerValidationPeriodEnded)
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
		it.Event = new(PoSValidatorManagerValidationPeriodEnded)
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
func (it *PoSValidatorManagerValidationPeriodEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSValidatorManagerValidationPeriodEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSValidatorManagerValidationPeriodEnded represents a ValidationPeriodEnded event raised by the PoSValidatorManager contract.
type PoSValidatorManagerValidationPeriodEnded struct {
	ValidationID [32]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodEnded is a free log retrieval operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) FilterValidationPeriodEnded(opts *bind.FilterOpts, validationID [][32]byte, status []uint8) (*PoSValidatorManagerValidationPeriodEndedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.FilterLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManagerValidationPeriodEndedIterator{contract: _PoSValidatorManager.contract, event: "ValidationPeriodEnded", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodEnded is a free log subscription operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) WatchValidationPeriodEnded(opts *bind.WatchOpts, sink chan<- *PoSValidatorManagerValidationPeriodEnded, validationID [][32]byte, status []uint8) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.WatchLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSValidatorManagerValidationPeriodEnded)
				if err := _PoSValidatorManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
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

// ParseValidationPeriodEnded is a log parse operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) ParseValidationPeriodEnded(log types.Log) (*PoSValidatorManagerValidationPeriodEnded, error) {
	event := new(PoSValidatorManagerValidationPeriodEnded)
	if err := _PoSValidatorManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoSValidatorManagerValidationPeriodRegisteredIterator is returned from FilterValidationPeriodRegistered and is used to iterate over the raw logs and unpacked data for ValidationPeriodRegistered events raised by the PoSValidatorManager contract.
type PoSValidatorManagerValidationPeriodRegisteredIterator struct {
	Event *PoSValidatorManagerValidationPeriodRegistered // Event containing the contract specifics and raw log

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
func (it *PoSValidatorManagerValidationPeriodRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSValidatorManagerValidationPeriodRegistered)
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
		it.Event = new(PoSValidatorManagerValidationPeriodRegistered)
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
func (it *PoSValidatorManagerValidationPeriodRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSValidatorManagerValidationPeriodRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSValidatorManagerValidationPeriodRegistered represents a ValidationPeriodRegistered event raised by the PoSValidatorManager contract.
type PoSValidatorManagerValidationPeriodRegistered struct {
	ValidationID [32]byte
	Weight       *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodRegistered is a free log retrieval operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) FilterValidationPeriodRegistered(opts *bind.FilterOpts, validationID [][32]byte) (*PoSValidatorManagerValidationPeriodRegisteredIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.FilterLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManagerValidationPeriodRegisteredIterator{contract: _PoSValidatorManager.contract, event: "ValidationPeriodRegistered", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodRegistered is a free log subscription operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) WatchValidationPeriodRegistered(opts *bind.WatchOpts, sink chan<- *PoSValidatorManagerValidationPeriodRegistered, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.WatchLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSValidatorManagerValidationPeriodRegistered)
				if err := _PoSValidatorManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
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

// ParseValidationPeriodRegistered is a log parse operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) ParseValidationPeriodRegistered(log types.Log) (*PoSValidatorManagerValidationPeriodRegistered, error) {
	event := new(PoSValidatorManagerValidationPeriodRegistered)
	if err := _PoSValidatorManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoSValidatorManagerValidationRewardsClaimedIterator is returned from FilterValidationRewardsClaimed and is used to iterate over the raw logs and unpacked data for ValidationRewardsClaimed events raised by the PoSValidatorManager contract.
type PoSValidatorManagerValidationRewardsClaimedIterator struct {
	Event *PoSValidatorManagerValidationRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *PoSValidatorManagerValidationRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSValidatorManagerValidationRewardsClaimed)
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
		it.Event = new(PoSValidatorManagerValidationRewardsClaimed)
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
func (it *PoSValidatorManagerValidationRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSValidatorManagerValidationRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSValidatorManagerValidationRewardsClaimed represents a ValidationRewardsClaimed event raised by the PoSValidatorManager contract.
type PoSValidatorManagerValidationRewardsClaimed struct {
	ValidationID [32]byte
	Reward       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationRewardsClaimed is a free log retrieval operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) FilterValidationRewardsClaimed(opts *bind.FilterOpts, validationID [][32]byte, reward []*big.Int) (*PoSValidatorManagerValidationRewardsClaimedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.FilterLogs(opts, "ValidationRewardsClaimed", validationIDRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManagerValidationRewardsClaimedIterator{contract: _PoSValidatorManager.contract, event: "ValidationRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchValidationRewardsClaimed is a free log subscription operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) WatchValidationRewardsClaimed(opts *bind.WatchOpts, sink chan<- *PoSValidatorManagerValidationRewardsClaimed, validationID [][32]byte, reward []*big.Int) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.WatchLogs(opts, "ValidationRewardsClaimed", validationIDRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSValidatorManagerValidationRewardsClaimed)
				if err := _PoSValidatorManager.contract.UnpackLog(event, "ValidationRewardsClaimed", log); err != nil {
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

// ParseValidationRewardsClaimed is a log parse operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) ParseValidationRewardsClaimed(log types.Log) (*PoSValidatorManagerValidationRewardsClaimed, error) {
	event := new(PoSValidatorManagerValidationRewardsClaimed)
	if err := _PoSValidatorManager.contract.UnpackLog(event, "ValidationRewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoSValidatorManagerValidatorRemovalInitializedIterator is returned from FilterValidatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for ValidatorRemovalInitialized events raised by the PoSValidatorManager contract.
type PoSValidatorManagerValidatorRemovalInitializedIterator struct {
	Event *PoSValidatorManagerValidatorRemovalInitialized // Event containing the contract specifics and raw log

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
func (it *PoSValidatorManagerValidatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSValidatorManagerValidatorRemovalInitialized)
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
		it.Event = new(PoSValidatorManagerValidatorRemovalInitialized)
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
func (it *PoSValidatorManagerValidatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSValidatorManagerValidatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSValidatorManagerValidatorRemovalInitialized represents a ValidatorRemovalInitialized event raised by the PoSValidatorManager contract.
type PoSValidatorManagerValidatorRemovalInitialized struct {
	ValidationID       [32]byte
	SetWeightMessageID [32]byte
	Weight             *big.Int
	EndTime            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemovalInitialized is a free log retrieval operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) FilterValidatorRemovalInitialized(opts *bind.FilterOpts, validationID [][32]byte, setWeightMessageID [][32]byte) (*PoSValidatorManagerValidatorRemovalInitializedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.FilterLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManagerValidatorRemovalInitializedIterator{contract: _PoSValidatorManager.contract, event: "ValidatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchValidatorRemovalInitialized is a free log subscription operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) WatchValidatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *PoSValidatorManagerValidatorRemovalInitialized, validationID [][32]byte, setWeightMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.WatchLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSValidatorManagerValidatorRemovalInitialized)
				if err := _PoSValidatorManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
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

// ParseValidatorRemovalInitialized is a log parse operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) ParseValidatorRemovalInitialized(log types.Log) (*PoSValidatorManagerValidatorRemovalInitialized, error) {
	event := new(PoSValidatorManagerValidatorRemovalInitialized)
	if err := _PoSValidatorManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoSValidatorManagerValidatorWeightUpdateIterator is returned from FilterValidatorWeightUpdate and is used to iterate over the raw logs and unpacked data for ValidatorWeightUpdate events raised by the PoSValidatorManager contract.
type PoSValidatorManagerValidatorWeightUpdateIterator struct {
	Event *PoSValidatorManagerValidatorWeightUpdate // Event containing the contract specifics and raw log

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
func (it *PoSValidatorManagerValidatorWeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoSValidatorManagerValidatorWeightUpdate)
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
		it.Event = new(PoSValidatorManagerValidatorWeightUpdate)
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
func (it *PoSValidatorManagerValidatorWeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoSValidatorManagerValidatorWeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoSValidatorManagerValidatorWeightUpdate represents a ValidatorWeightUpdate event raised by the PoSValidatorManager contract.
type PoSValidatorManagerValidatorWeightUpdate struct {
	ValidationID       [32]byte
	Nonce              uint64
	ValidatorWeight    uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorWeightUpdate is a free log retrieval operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) FilterValidatorWeightUpdate(opts *bind.FilterOpts, validationID [][32]byte, nonce []uint64) (*PoSValidatorManagerValidatorWeightUpdateIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.FilterLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &PoSValidatorManagerValidatorWeightUpdateIterator{contract: _PoSValidatorManager.contract, event: "ValidatorWeightUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorWeightUpdate is a free log subscription operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) WatchValidatorWeightUpdate(opts *bind.WatchOpts, sink chan<- *PoSValidatorManagerValidatorWeightUpdate, validationID [][32]byte, nonce []uint64) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _PoSValidatorManager.contract.WatchLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoSValidatorManagerValidatorWeightUpdate)
				if err := _PoSValidatorManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
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

// ParseValidatorWeightUpdate is a log parse operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_PoSValidatorManager *PoSValidatorManagerFilterer) ParseValidatorWeightUpdate(log types.Log) (*PoSValidatorManagerValidatorWeightUpdate, error) {
	event := new(PoSValidatorManagerValidatorWeightUpdate)
	if err := _PoSValidatorManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReentrancyGuardUpgradeableMetaData contains all meta data concerning the ReentrancyGuardUpgradeable contract.
var ReentrancyGuardUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// ReentrancyGuardUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ReentrancyGuardUpgradeableMetaData.ABI instead.
var ReentrancyGuardUpgradeableABI = ReentrancyGuardUpgradeableMetaData.ABI

// ReentrancyGuardUpgradeable is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeable struct {
	ReentrancyGuardUpgradeableCaller     // Read-only binding to the contract
	ReentrancyGuardUpgradeableTransactor // Write-only binding to the contract
	ReentrancyGuardUpgradeableFilterer   // Log filterer for contract events
}

// ReentrancyGuardUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardUpgradeableSession struct {
	Contract     *ReentrancyGuardUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ReentrancyGuardUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardUpgradeableCallerSession struct {
	Contract *ReentrancyGuardUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ReentrancyGuardUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardUpgradeableTransactorSession struct {
	Contract     *ReentrancyGuardUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ReentrancyGuardUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableRaw struct {
	Contract *ReentrancyGuardUpgradeable // Generic contract binding to access the raw methods on
}

// ReentrancyGuardUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableCallerRaw struct {
	Contract *ReentrancyGuardUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableTransactorRaw struct {
	Contract *ReentrancyGuardUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuardUpgradeable creates a new instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeable(address common.Address, backend bind.ContractBackend) (*ReentrancyGuardUpgradeable, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeable{ReentrancyGuardUpgradeableCaller: ReentrancyGuardUpgradeableCaller{contract: contract}, ReentrancyGuardUpgradeableTransactor: ReentrancyGuardUpgradeableTransactor{contract: contract}, ReentrancyGuardUpgradeableFilterer: ReentrancyGuardUpgradeableFilterer{contract: contract}}, nil
}

// NewReentrancyGuardUpgradeableCaller creates a new read-only instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardUpgradeableCaller, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableCaller{contract: contract}, nil
}

// NewReentrancyGuardUpgradeableTransactor creates a new write-only instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardUpgradeableTransactor, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableTransactor{contract: contract}, nil
}

// NewReentrancyGuardUpgradeableFilterer creates a new log filterer instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardUpgradeableFilterer, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableFilterer{contract: contract}, nil
}

// bindReentrancyGuardUpgradeable binds a generic wrapper to an already deployed contract.
func bindReentrancyGuardUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReentrancyGuardUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuardUpgradeable.Contract.ReentrancyGuardUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.ReentrancyGuardUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.ReentrancyGuardUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuardUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ReentrancyGuardUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ReentrancyGuardUpgradeable contract.
type ReentrancyGuardUpgradeableInitializedIterator struct {
	Event *ReentrancyGuardUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *ReentrancyGuardUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReentrancyGuardUpgradeableInitialized)
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
		it.Event = new(ReentrancyGuardUpgradeableInitialized)
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
func (it *ReentrancyGuardUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReentrancyGuardUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReentrancyGuardUpgradeableInitialized represents a Initialized event raised by the ReentrancyGuardUpgradeable contract.
type ReentrancyGuardUpgradeableInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ReentrancyGuardUpgradeableInitializedIterator, error) {

	logs, sub, err := _ReentrancyGuardUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableInitializedIterator{contract: _ReentrancyGuardUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ReentrancyGuardUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _ReentrancyGuardUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReentrancyGuardUpgradeableInitialized)
				if err := _ReentrancyGuardUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableFilterer) ParseInitialized(log types.Log) (*ReentrancyGuardUpgradeableInitialized, error) {
	event := new(ReentrancyGuardUpgradeableInitialized)
	if err := _ReentrancyGuardUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerMetaData contains all meta data concerning the ValidatorManager contract.
var ValidatorManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"InvalidBLSKeyLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializationStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\"}],\"name\":\"InvalidMaximumChurnPercentage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"InvalidNodeID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addressesLength\",\"type\":\"uint256\"}],\"name\":\"InvalidPChainOwnerThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"InvalidRegistrationExpiry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"encodedSubnetConversionID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedSubnetConversionID\",\"type\":\"bytes32\"}],\"name\":\"InvalidSubnetConversionID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"}],\"name\":\"InvalidValidatorManagerAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidatorManagerBlockchainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidValidatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWarpMessage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"InvalidWarpOriginSenderAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidWarpSourceChainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"churnAmount\",\"type\":\"uint64\"}],\"name\":\"MaxChurnRateExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PChainOwnerAddressesNotSorted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"validRegistration\",\"type\":\"bool\"}],\"name\":\"UnexpectedRegistrationStatus\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"InitialValidatorCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"ValidationPeriodCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ValidationPeriodEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidationPeriodRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"ValidatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorWeightUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_LENGTH\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLS_PUBLIC_KEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_CHURN_PERCENTAGE_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_REGISTRATION_EXPIRY_LENGTH\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"P_CHAIN_BLOCKCHAIN_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"startingWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"messageNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endedAt\",\"type\":\"uint64\"}],\"internalType\":\"structValidator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structSubnetConversionData\",\"name\":\"subnetConversionData\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"registeredValidators\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendEndValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendRegisterValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ValidatorManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorManagerMetaData.ABI instead.
var ValidatorManagerABI = ValidatorManagerMetaData.ABI

// ValidatorManager is an auto generated Go binding around an Ethereum contract.
type ValidatorManager struct {
	ValidatorManagerCaller     // Read-only binding to the contract
	ValidatorManagerTransactor // Write-only binding to the contract
	ValidatorManagerFilterer   // Log filterer for contract events
}

// ValidatorManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorManagerSession struct {
	Contract     *ValidatorManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorManagerCallerSession struct {
	Contract *ValidatorManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ValidatorManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorManagerTransactorSession struct {
	Contract     *ValidatorManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ValidatorManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorManagerRaw struct {
	Contract *ValidatorManager // Generic contract binding to access the raw methods on
}

// ValidatorManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorManagerCallerRaw struct {
	Contract *ValidatorManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorManagerTransactorRaw struct {
	Contract *ValidatorManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorManager creates a new instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManager(address common.Address, backend bind.ContractBackend) (*ValidatorManager, error) {
	contract, err := bindValidatorManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorManager{ValidatorManagerCaller: ValidatorManagerCaller{contract: contract}, ValidatorManagerTransactor: ValidatorManagerTransactor{contract: contract}, ValidatorManagerFilterer: ValidatorManagerFilterer{contract: contract}}, nil
}

// NewValidatorManagerCaller creates a new read-only instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManagerCaller(address common.Address, caller bind.ContractCaller) (*ValidatorManagerCaller, error) {
	contract, err := bindValidatorManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerCaller{contract: contract}, nil
}

// NewValidatorManagerTransactor creates a new write-only instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorManagerTransactor, error) {
	contract, err := bindValidatorManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerTransactor{contract: contract}, nil
}

// NewValidatorManagerFilterer creates a new log filterer instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorManagerFilterer, error) {
	contract, err := bindValidatorManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerFilterer{contract: contract}, nil
}

// bindValidatorManager binds a generic wrapper to an already deployed contract.
func bindValidatorManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorManager *ValidatorManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorManager.Contract.ValidatorManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorManager *ValidatorManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.Contract.ValidatorManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorManager *ValidatorManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorManager.Contract.ValidatorManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorManager *ValidatorManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorManager *ValidatorManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorManager *ValidatorManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorManager.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_ValidatorManager *ValidatorManagerCaller) ADDRESSLENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "ADDRESS_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_ValidatorManager *ValidatorManagerSession) ADDRESSLENGTH() (uint32, error) {
	return _ValidatorManager.Contract.ADDRESSLENGTH(&_ValidatorManager.CallOpts)
}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_ValidatorManager *ValidatorManagerCallerSession) ADDRESSLENGTH() (uint32, error) {
	return _ValidatorManager.Contract.ADDRESSLENGTH(&_ValidatorManager.CallOpts)
}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_ValidatorManager *ValidatorManagerCaller) BLSPUBLICKEYLENGTH(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "BLS_PUBLIC_KEY_LENGTH")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_ValidatorManager *ValidatorManagerSession) BLSPUBLICKEYLENGTH() (uint8, error) {
	return _ValidatorManager.Contract.BLSPUBLICKEYLENGTH(&_ValidatorManager.CallOpts)
}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_ValidatorManager *ValidatorManagerCallerSession) BLSPUBLICKEYLENGTH() (uint8, error) {
	return _ValidatorManager.Contract.BLSPUBLICKEYLENGTH(&_ValidatorManager.CallOpts)
}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_ValidatorManager *ValidatorManagerCaller) MAXIMUMCHURNPERCENTAGELIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "MAXIMUM_CHURN_PERCENTAGE_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_ValidatorManager *ValidatorManagerSession) MAXIMUMCHURNPERCENTAGELIMIT() (uint8, error) {
	return _ValidatorManager.Contract.MAXIMUMCHURNPERCENTAGELIMIT(&_ValidatorManager.CallOpts)
}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_ValidatorManager *ValidatorManagerCallerSession) MAXIMUMCHURNPERCENTAGELIMIT() (uint8, error) {
	return _ValidatorManager.Contract.MAXIMUMCHURNPERCENTAGELIMIT(&_ValidatorManager.CallOpts)
}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_ValidatorManager *ValidatorManagerCaller) MAXIMUMREGISTRATIONEXPIRYLENGTH(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "MAXIMUM_REGISTRATION_EXPIRY_LENGTH")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_ValidatorManager *ValidatorManagerSession) MAXIMUMREGISTRATIONEXPIRYLENGTH() (uint64, error) {
	return _ValidatorManager.Contract.MAXIMUMREGISTRATIONEXPIRYLENGTH(&_ValidatorManager.CallOpts)
}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_ValidatorManager *ValidatorManagerCallerSession) MAXIMUMREGISTRATIONEXPIRYLENGTH() (uint64, error) {
	return _ValidatorManager.Contract.MAXIMUMREGISTRATIONEXPIRYLENGTH(&_ValidatorManager.CallOpts)
}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_ValidatorManager *ValidatorManagerCaller) PCHAINBLOCKCHAINID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "P_CHAIN_BLOCKCHAIN_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_ValidatorManager *ValidatorManagerSession) PCHAINBLOCKCHAINID() ([32]byte, error) {
	return _ValidatorManager.Contract.PCHAINBLOCKCHAINID(&_ValidatorManager.CallOpts)
}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_ValidatorManager *ValidatorManagerCallerSession) PCHAINBLOCKCHAINID() ([32]byte, error) {
	return _ValidatorManager.Contract.PCHAINBLOCKCHAINID(&_ValidatorManager.CallOpts)
}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ValidatorManager *ValidatorManagerCaller) VALIDATORMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "VALIDATOR_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ValidatorManager *ValidatorManagerSession) VALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ValidatorManager.Contract.VALIDATORMANAGERSTORAGELOCATION(&_ValidatorManager.CallOpts)
}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ValidatorManager *ValidatorManagerCallerSession) VALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ValidatorManager.Contract.VALIDATORMANAGERSTORAGELOCATION(&_ValidatorManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ValidatorManager *ValidatorManagerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ValidatorManager *ValidatorManagerSession) WARPMESSENGER() (common.Address, error) {
	return _ValidatorManager.Contract.WARPMESSENGER(&_ValidatorManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ValidatorManager *ValidatorManagerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _ValidatorManager.Contract.WARPMESSENGER(&_ValidatorManager.CallOpts)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_ValidatorManager *ValidatorManagerCaller) GetValidator(opts *bind.CallOpts, validationID [32]byte) (Validator, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "getValidator", validationID)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_ValidatorManager *ValidatorManagerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _ValidatorManager.Contract.GetValidator(&_ValidatorManager.CallOpts, validationID)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_ValidatorManager *ValidatorManagerCallerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _ValidatorManager.Contract.GetValidator(&_ValidatorManager.CallOpts, validationID)
}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_ValidatorManager *ValidatorManagerCaller) GetWeight(opts *bind.CallOpts, validationID [32]byte) (uint64, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "getWeight", validationID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_ValidatorManager *ValidatorManagerSession) GetWeight(validationID [32]byte) (uint64, error) {
	return _ValidatorManager.Contract.GetWeight(&_ValidatorManager.CallOpts, validationID)
}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_ValidatorManager *ValidatorManagerCallerSession) GetWeight(validationID [32]byte) (uint64, error) {
	return _ValidatorManager.Contract.GetWeight(&_ValidatorManager.CallOpts, validationID)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_ValidatorManager *ValidatorManagerCaller) RegisteredValidators(opts *bind.CallOpts, nodeID []byte) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "registeredValidators", nodeID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_ValidatorManager *ValidatorManagerSession) RegisteredValidators(nodeID []byte) ([32]byte, error) {
	return _ValidatorManager.Contract.RegisteredValidators(&_ValidatorManager.CallOpts, nodeID)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_ValidatorManager *ValidatorManagerCallerSession) RegisteredValidators(nodeID []byte) ([32]byte, error) {
	return _ValidatorManager.Contract.RegisteredValidators(&_ValidatorManager.CallOpts, nodeID)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_ValidatorManager *ValidatorManagerTransactor) CompleteEndValidation(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "completeEndValidation", messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_ValidatorManager *ValidatorManagerSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorManager.Contract.CompleteEndValidation(&_ValidatorManager.TransactOpts, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorManager.Contract.CompleteEndValidation(&_ValidatorManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_ValidatorManager *ValidatorManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_ValidatorManager *ValidatorManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorManager.Contract.CompleteValidatorRegistration(&_ValidatorManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorManager.Contract.CompleteValidatorRegistration(&_ValidatorManager.TransactOpts, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messageIndex) returns()
func (_ValidatorManager *ValidatorManagerTransactor) InitializeValidatorSet(opts *bind.TransactOpts, subnetConversionData SubnetConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "initializeValidatorSet", subnetConversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messageIndex) returns()
func (_ValidatorManager *ValidatorManagerSession) InitializeValidatorSet(subnetConversionData SubnetConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorManager.Contract.InitializeValidatorSet(&_ValidatorManager.TransactOpts, subnetConversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messageIndex) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) InitializeValidatorSet(subnetConversionData SubnetConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorManager.Contract.InitializeValidatorSet(&_ValidatorManager.TransactOpts, subnetConversionData, messageIndex)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_ValidatorManager *ValidatorManagerTransactor) ResendEndValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "resendEndValidatorMessage", validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_ValidatorManager *ValidatorManagerSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _ValidatorManager.Contract.ResendEndValidatorMessage(&_ValidatorManager.TransactOpts, validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _ValidatorManager.Contract.ResendEndValidatorMessage(&_ValidatorManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_ValidatorManager *ValidatorManagerTransactor) ResendRegisterValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "resendRegisterValidatorMessage", validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_ValidatorManager *ValidatorManagerSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _ValidatorManager.Contract.ResendRegisterValidatorMessage(&_ValidatorManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _ValidatorManager.Contract.ResendRegisterValidatorMessage(&_ValidatorManager.TransactOpts, validationID)
}

// ValidatorManagerInitialValidatorCreatedIterator is returned from FilterInitialValidatorCreated and is used to iterate over the raw logs and unpacked data for InitialValidatorCreated events raised by the ValidatorManager contract.
type ValidatorManagerInitialValidatorCreatedIterator struct {
	Event *ValidatorManagerInitialValidatorCreated // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerInitialValidatorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerInitialValidatorCreated)
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
		it.Event = new(ValidatorManagerInitialValidatorCreated)
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
func (it *ValidatorManagerInitialValidatorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerInitialValidatorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerInitialValidatorCreated represents a InitialValidatorCreated event raised by the ValidatorManager contract.
type ValidatorManagerInitialValidatorCreated struct {
	ValidationID [32]byte
	NodeID       common.Hash
	Weight       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialValidatorCreated is a free log retrieval operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_ValidatorManager *ValidatorManagerFilterer) FilterInitialValidatorCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte) (*ValidatorManagerInitialValidatorCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerInitialValidatorCreatedIterator{contract: _ValidatorManager.contract, event: "InitialValidatorCreated", logs: logs, sub: sub}, nil
}

// WatchInitialValidatorCreated is a free log subscription operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_ValidatorManager *ValidatorManagerFilterer) WatchInitialValidatorCreated(opts *bind.WatchOpts, sink chan<- *ValidatorManagerInitialValidatorCreated, validationID [][32]byte, nodeID [][]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerInitialValidatorCreated)
				if err := _ValidatorManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
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

// ParseInitialValidatorCreated is a log parse operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_ValidatorManager *ValidatorManagerFilterer) ParseInitialValidatorCreated(log types.Log) (*ValidatorManagerInitialValidatorCreated, error) {
	event := new(ValidatorManagerInitialValidatorCreated)
	if err := _ValidatorManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ValidatorManager contract.
type ValidatorManagerInitializedIterator struct {
	Event *ValidatorManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerInitialized)
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
		it.Event = new(ValidatorManagerInitialized)
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
func (it *ValidatorManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerInitialized represents a Initialized event raised by the ValidatorManager contract.
type ValidatorManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ValidatorManager *ValidatorManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ValidatorManagerInitializedIterator, error) {

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerInitializedIterator{contract: _ValidatorManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ValidatorManager *ValidatorManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ValidatorManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerInitialized)
				if err := _ValidatorManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ValidatorManager *ValidatorManagerFilterer) ParseInitialized(log types.Log) (*ValidatorManagerInitialized, error) {
	event := new(ValidatorManagerInitialized)
	if err := _ValidatorManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerValidationPeriodCreatedIterator is returned from FilterValidationPeriodCreated and is used to iterate over the raw logs and unpacked data for ValidationPeriodCreated events raised by the ValidatorManager contract.
type ValidatorManagerValidationPeriodCreatedIterator struct {
	Event *ValidatorManagerValidationPeriodCreated // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerValidationPeriodCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerValidationPeriodCreated)
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
		it.Event = new(ValidatorManagerValidationPeriodCreated)
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
func (it *ValidatorManagerValidationPeriodCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerValidationPeriodCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerValidationPeriodCreated represents a ValidationPeriodCreated event raised by the ValidatorManager contract.
type ValidatorManagerValidationPeriodCreated struct {
	ValidationID                [32]byte
	NodeID                      common.Hash
	RegisterValidationMessageID [32]byte
	Weight                      *big.Int
	RegistrationExpiry          uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodCreated is a free log retrieval operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_ValidatorManager *ValidatorManagerFilterer) FilterValidationPeriodCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (*ValidatorManagerValidationPeriodCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerValidationPeriodCreatedIterator{contract: _ValidatorManager.contract, event: "ValidationPeriodCreated", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodCreated is a free log subscription operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_ValidatorManager *ValidatorManagerFilterer) WatchValidationPeriodCreated(opts *bind.WatchOpts, sink chan<- *ValidatorManagerValidationPeriodCreated, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerValidationPeriodCreated)
				if err := _ValidatorManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
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

// ParseValidationPeriodCreated is a log parse operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_ValidatorManager *ValidatorManagerFilterer) ParseValidationPeriodCreated(log types.Log) (*ValidatorManagerValidationPeriodCreated, error) {
	event := new(ValidatorManagerValidationPeriodCreated)
	if err := _ValidatorManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerValidationPeriodEndedIterator is returned from FilterValidationPeriodEnded and is used to iterate over the raw logs and unpacked data for ValidationPeriodEnded events raised by the ValidatorManager contract.
type ValidatorManagerValidationPeriodEndedIterator struct {
	Event *ValidatorManagerValidationPeriodEnded // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerValidationPeriodEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerValidationPeriodEnded)
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
		it.Event = new(ValidatorManagerValidationPeriodEnded)
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
func (it *ValidatorManagerValidationPeriodEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerValidationPeriodEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerValidationPeriodEnded represents a ValidationPeriodEnded event raised by the ValidatorManager contract.
type ValidatorManagerValidationPeriodEnded struct {
	ValidationID [32]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodEnded is a free log retrieval operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_ValidatorManager *ValidatorManagerFilterer) FilterValidationPeriodEnded(opts *bind.FilterOpts, validationID [][32]byte, status []uint8) (*ValidatorManagerValidationPeriodEndedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerValidationPeriodEndedIterator{contract: _ValidatorManager.contract, event: "ValidationPeriodEnded", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodEnded is a free log subscription operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_ValidatorManager *ValidatorManagerFilterer) WatchValidationPeriodEnded(opts *bind.WatchOpts, sink chan<- *ValidatorManagerValidationPeriodEnded, validationID [][32]byte, status []uint8) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerValidationPeriodEnded)
				if err := _ValidatorManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
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

// ParseValidationPeriodEnded is a log parse operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_ValidatorManager *ValidatorManagerFilterer) ParseValidationPeriodEnded(log types.Log) (*ValidatorManagerValidationPeriodEnded, error) {
	event := new(ValidatorManagerValidationPeriodEnded)
	if err := _ValidatorManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerValidationPeriodRegisteredIterator is returned from FilterValidationPeriodRegistered and is used to iterate over the raw logs and unpacked data for ValidationPeriodRegistered events raised by the ValidatorManager contract.
type ValidatorManagerValidationPeriodRegisteredIterator struct {
	Event *ValidatorManagerValidationPeriodRegistered // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerValidationPeriodRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerValidationPeriodRegistered)
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
		it.Event = new(ValidatorManagerValidationPeriodRegistered)
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
func (it *ValidatorManagerValidationPeriodRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerValidationPeriodRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerValidationPeriodRegistered represents a ValidationPeriodRegistered event raised by the ValidatorManager contract.
type ValidatorManagerValidationPeriodRegistered struct {
	ValidationID [32]byte
	Weight       *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodRegistered is a free log retrieval operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_ValidatorManager *ValidatorManagerFilterer) FilterValidationPeriodRegistered(opts *bind.FilterOpts, validationID [][32]byte) (*ValidatorManagerValidationPeriodRegisteredIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerValidationPeriodRegisteredIterator{contract: _ValidatorManager.contract, event: "ValidationPeriodRegistered", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodRegistered is a free log subscription operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_ValidatorManager *ValidatorManagerFilterer) WatchValidationPeriodRegistered(opts *bind.WatchOpts, sink chan<- *ValidatorManagerValidationPeriodRegistered, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerValidationPeriodRegistered)
				if err := _ValidatorManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
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

// ParseValidationPeriodRegistered is a log parse operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_ValidatorManager *ValidatorManagerFilterer) ParseValidationPeriodRegistered(log types.Log) (*ValidatorManagerValidationPeriodRegistered, error) {
	event := new(ValidatorManagerValidationPeriodRegistered)
	if err := _ValidatorManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerValidatorRemovalInitializedIterator is returned from FilterValidatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for ValidatorRemovalInitialized events raised by the ValidatorManager contract.
type ValidatorManagerValidatorRemovalInitializedIterator struct {
	Event *ValidatorManagerValidatorRemovalInitialized // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerValidatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerValidatorRemovalInitialized)
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
		it.Event = new(ValidatorManagerValidatorRemovalInitialized)
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
func (it *ValidatorManagerValidatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerValidatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerValidatorRemovalInitialized represents a ValidatorRemovalInitialized event raised by the ValidatorManager contract.
type ValidatorManagerValidatorRemovalInitialized struct {
	ValidationID       [32]byte
	SetWeightMessageID [32]byte
	Weight             *big.Int
	EndTime            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemovalInitialized is a free log retrieval operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_ValidatorManager *ValidatorManagerFilterer) FilterValidatorRemovalInitialized(opts *bind.FilterOpts, validationID [][32]byte, setWeightMessageID [][32]byte) (*ValidatorManagerValidatorRemovalInitializedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerValidatorRemovalInitializedIterator{contract: _ValidatorManager.contract, event: "ValidatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchValidatorRemovalInitialized is a free log subscription operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_ValidatorManager *ValidatorManagerFilterer) WatchValidatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *ValidatorManagerValidatorRemovalInitialized, validationID [][32]byte, setWeightMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerValidatorRemovalInitialized)
				if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
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

// ParseValidatorRemovalInitialized is a log parse operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_ValidatorManager *ValidatorManagerFilterer) ParseValidatorRemovalInitialized(log types.Log) (*ValidatorManagerValidatorRemovalInitialized, error) {
	event := new(ValidatorManagerValidatorRemovalInitialized)
	if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerValidatorWeightUpdateIterator is returned from FilterValidatorWeightUpdate and is used to iterate over the raw logs and unpacked data for ValidatorWeightUpdate events raised by the ValidatorManager contract.
type ValidatorManagerValidatorWeightUpdateIterator struct {
	Event *ValidatorManagerValidatorWeightUpdate // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerValidatorWeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerValidatorWeightUpdate)
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
		it.Event = new(ValidatorManagerValidatorWeightUpdate)
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
func (it *ValidatorManagerValidatorWeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerValidatorWeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerValidatorWeightUpdate represents a ValidatorWeightUpdate event raised by the ValidatorManager contract.
type ValidatorManagerValidatorWeightUpdate struct {
	ValidationID       [32]byte
	Nonce              uint64
	ValidatorWeight    uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorWeightUpdate is a free log retrieval operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_ValidatorManager *ValidatorManagerFilterer) FilterValidatorWeightUpdate(opts *bind.FilterOpts, validationID [][32]byte, nonce []uint64) (*ValidatorManagerValidatorWeightUpdateIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerValidatorWeightUpdateIterator{contract: _ValidatorManager.contract, event: "ValidatorWeightUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorWeightUpdate is a free log subscription operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_ValidatorManager *ValidatorManagerFilterer) WatchValidatorWeightUpdate(opts *bind.WatchOpts, sink chan<- *ValidatorManagerValidatorWeightUpdate, validationID [][32]byte, nonce []uint64) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerValidatorWeightUpdate)
				if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
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

// ParseValidatorWeightUpdate is a log parse operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_ValidatorManager *ValidatorManagerFilterer) ParseValidatorWeightUpdate(log types.Log) (*ValidatorManagerValidatorWeightUpdate, error) {
	event := new(ValidatorManagerValidatorWeightUpdate)
	if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMessagesMetaData contains all meta data concerning the ValidatorMessages contract.
var ValidatorMessagesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidBLSPublicKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"name\":\"InvalidCodecID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"actual\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"expected\",\"type\":\"uint32\"}],\"name\":\"InvalidMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMessageType\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorMessages.ValidationPeriod\",\"name\":\"validationPeriod\",\"type\":\"tuple\"}],\"name\":\"packRegisterSubnetValidatorMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structSubnetConversionData\",\"name\":\"subnetConversionData\",\"type\":\"tuple\"}],\"name\":\"packSubnetConversionData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnetConversionID\",\"type\":\"bytes32\"}],\"name\":\"packSubnetConversionMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"packSubnetValidatorRegistrationMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"packSubnetValidatorWeightMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"packValidationUptimeMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackRegisterSubnetValidatorMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structValidatorMessages.ValidationPeriod\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackSubnetConversionMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackSubnetValidatorRegistrationMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackSubnetValidatorWeightMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"unpackValidationUptimeMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x612160610034600b8282823980515f1a607314602857634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100b1575f3560e01c8063862bfa6311610079578063862bfa63146101d75780639de23d40146101ea578063a523377014610222578063e047b28314610242578063e1d68f3014610263578063fa1f8dfb14610276575f80fd5b8063088c2463146100b55780631e6d9789146100ea5780631fd979c71461010b5780632e43ceb5146101525780638545c16a1461017a575b5f80fd5b6100c86100c3366004611904565b610289565b604080519283526001600160401b039091166020830152015b60405180910390f35b6100fd6100f8366004611904565b61047f565b6040519081526020016100e1565b61014561011936600461193d565b604080515f60208201819052602282015260268082019390935281518082039093018352604601905290565b6040516100e191906119a1565b610165610160366004611904565b61060c565b604080519283529015156020830152016100e1565b6101456101883660046119d5565b604080515f6020820152600360e01b602282015260268101949094526001600160c01b031960c093841b811660468601529190921b16604e830152805180830360360181526056909201905290565b6101456101e5366004611a0e565b6107c8565b6101fd6101f8366004611904565b6109a9565b604080519384526001600160401b0392831660208501529116908201526060016100e1565b610235610230366004611904565b610bff565b6040516100e19190611aaa565b610255610250366004611c45565b61154a565b6040516100e1929190611d41565b610145610271366004611d59565b611737565b610145610284366004611d83565b611781565b5f808251602e146102c457825160405163cc92daa160e01b815263ffffffff9091166004820152602e60248201526044015b60405180910390fd5b5f805b6002811015610313576102db816001611dc9565b6102e6906008611ddc565b61ffff168582815181106102fc576102fc611df3565b016020015160f81c901b91909117906001016102c7565b5061ffff81161561033d5760405163407b587360e01b815261ffff821660048201526024016102bb565b5f805b600481101561039857610354816003611dc9565b61035f906008611ddc565b63ffffffff1686610371836002611e07565b8151811061038157610381611df3565b016020015160f81c901b9190911790600101610340565b5063ffffffff8116156103be57604051635b60892f60e01b815260040160405180910390fd5b5f805b6020811015610413576103d581601f611dc9565b6103e0906008611ddc565b876103ec836006611e07565b815181106103fc576103fc611df3565b016020015160f81c901b91909117906001016103c1565b505f805b60088110156104725761042b816007611dc9565b610436906008611ddc565b6001600160401b03168861044b836026611e07565b8151811061045b5761045b611df3565b016020015160f81c901b9190911790600101610417565b5090969095509350505050565b5f81516026146104b457815160405163cc92daa160e01b815263ffffffff9091166004820152602660248201526044016102bb565b5f805b6002811015610503576104cb816001611dc9565b6104d6906008611ddc565b61ffff168482815181106104ec576104ec611df3565b016020015160f81c901b91909117906001016104b7565b5061ffff81161561052d5760405163407b587360e01b815261ffff821660048201526024016102bb565b5f805b600481101561058857610544816003611dc9565b61054f906008611ddc565b63ffffffff1685610561836002611e07565b8151811061057157610571611df3565b016020015160f81c901b9190911790600101610530565b5063ffffffff8116156105ae57604051635b60892f60e01b815260040160405180910390fd5b5f805b6020811015610603576105c581601f611dc9565b6105d0906008611ddc565b866105dc836006611e07565b815181106105ec576105ec611df3565b016020015160f81c901b91909117906001016105b1565b50949350505050565b5f80825160271461064257825160405163cc92daa160e01b815263ffffffff9091166004820152602760248201526044016102bb565b5f805b600281101561069157610659816001611dc9565b610664906008611ddc565b61ffff1685828151811061067a5761067a611df3565b016020015160f81c901b9190911790600101610645565b5061ffff8116156106bb5760405163407b587360e01b815261ffff821660048201526024016102bb565b5f805b6004811015610716576106d2816003611dc9565b6106dd906008611ddc565b63ffffffff16866106ef836002611e07565b815181106106ff576106ff611df3565b016020015160f81c901b91909117906001016106be565b5063ffffffff811660021461073e57604051635b60892f60e01b815260040160405180910390fd5b5f805b60208110156107935761075581601f611dc9565b610760906008611ddc565b8761076c836006611e07565b8151811061077c5761077c611df3565b016020015160f81c901b9190911790600101610741565b505f866026815181106107a8576107a8611df3565b016020015191976001600160f81b03199092161515965090945050505050565b60605f808335602085013560146107e487870160408901611e1a565b6107f16060890189611e33565b60405160f09790971b6001600160f01b0319166020880152602287019590955250604285019290925260e090811b6001600160e01b0319908116606286015260609290921b6bffffffffffffffffffffffff191660668501529190911b16607a820152607e0160405160208183030381529060405290505f5b6108776060850185611e33565b90508110156109a2578161088e6060860186611e33565b8381811061089e5761089e611df3565b90506020028101906108b09190611e7f565b6108ba9080611e9d565b90506108c96060870187611e33565b848181106108d9576108d9611df3565b90506020028101906108eb9190611e7f565b6108f59080611e9d565b6109026060890189611e33565b8681811061091257610912611df3565b90506020028101906109249190611e7f565b610932906020810190611e9d565b61093f60608b018b611e33565b8881811061094f5761094f611df3565b90506020028101906109619190611e7f565b610972906060810190604001611edf565b6040516020016109889796959493929190611ef8565b60408051601f19818403018152919052915060010161086a565b5092915050565b5f805f83516036146109e057835160405163cc92daa160e01b815263ffffffff9091166004820152603660248201526044016102bb565b5f805b6002811015610a2f576109f7816001611dc9565b610a02906008611ddc565b61ffff16868281518110610a1857610a18611df3565b016020015160f81c901b91909117906001016109e3565b5061ffff811615610a595760405163407b587360e01b815261ffff821660048201526024016102bb565b5f805b6004811015610ab457610a70816003611dc9565b610a7b906008611ddc565b63ffffffff1687610a8d836002611e07565b81518110610a9d57610a9d611df3565b016020015160f81c901b9190911790600101610a5c565b5063ffffffff8116600314610adc57604051635b60892f60e01b815260040160405180910390fd5b5f805b6020811015610b3157610af381601f611dc9565b610afe906008611ddc565b88610b0a836006611e07565b81518110610b1a57610b1a611df3565b016020015160f81c901b9190911790600101610adf565b505f805b6008811015610b9057610b49816007611dc9565b610b54906008611ddc565b6001600160401b031689610b69836026611e07565b81518110610b7957610b79611df3565b016020015160f81c901b9190911790600101610b35565b505f805b6008811015610bef57610ba8816007611dc9565b610bb3906008611ddc565b6001600160401b03168a610bc883602e611e07565b81518110610bd857610bd8611df3565b016020015160f81c901b9190911790600101610b94565b5091989097509095509350505050565b610c076117b1565b5f610c106117b1565b5f805b6002811015610c6e57610c27816001611dc9565b610c32906008611ddc565b61ffff1686610c4763ffffffff871684611e07565b81518110610c5757610c57611df3565b016020015160f81c901b9190911790600101610c13565b5061ffff811615610c985760405163407b587360e01b815261ffff821660048201526024016102bb565b610ca3600284611f61565b9250505f805b6004811015610d0857610cbd816003611dc9565b610cc8906008611ddc565b63ffffffff16868563ffffffff1683610ce19190611e07565b81518110610cf157610cf1611df3565b016020015160f81c901b9190911790600101610ca9565b5063ffffffff8116600114610d3057604051635b60892f60e01b815260040160405180910390fd5b610d3b600484611f61565b9250505f805b6020811015610d9857610d5581601f611dc9565b610d60906008611ddc565b86610d7163ffffffff871684611e07565b81518110610d8157610d81611df3565b016020015160f81c901b9190911790600101610d41565b50808252610da7602084611f61565b9250505f805b6004811015610e0c57610dc1816003611dc9565b610dcc906008611ddc565b63ffffffff16868563ffffffff1683610de59190611e07565b81518110610df557610df5611df3565b016020015160f81c901b9190911790600101610dad565b50610e18600484611f61565b92505f8163ffffffff166001600160401b03811115610e3957610e3961180b565b6040519080825280601f01601f191660200182016040528015610e63576020820181803683370190505b5090505f5b8263ffffffff16811015610ed25786610e8763ffffffff871683611e07565b81518110610e9757610e97611df3565b602001015160f81c60f81b828281518110610eb457610eb4611df3565b60200101906001600160f81b03191690815f1a905350600101610e68565b5060208301819052610ee48285611f61565b604080516030808252606082019092529195505f92506020820181803683370190505090505f5b6030811015610f705786610f2563ffffffff871683611e07565b81518110610f3557610f35611df3565b602001015160f81c60f81b828281518110610f5257610f52611df3565b60200101906001600160f81b03191690815f1a905350600101610f0b565b5060408301819052610f83603085611f61565b9350505f805b6008811015610fe957610f9d816007611dc9565b610fa8906008611ddc565b6001600160401b031687610fc263ffffffff881684611e07565b81518110610fd257610fd2611df3565b016020015160f81c901b9190911790600101610f89565b506001600160401b0381166060840152611004600885611f61565b9350505f805f5b600481101561106a5761101f816003611dc9565b61102a906008611ddc565b63ffffffff16888763ffffffff16836110439190611e07565b8151811061105357611053611df3565b016020015160f81c901b919091179060010161100b565b50611076600486611f61565b94505f5b60048110156110d95761108e816003611dc9565b611099906008611ddc565b63ffffffff16888763ffffffff16836110b29190611e07565b815181106110c2576110c2611df3565b016020015160f81c901b929092179160010161107a565b506110e5600486611f61565b94505f8263ffffffff166001600160401b038111156111065761110661180b565b60405190808252806020026020018201604052801561112f578160200160208202803683370190505b5090505f5b8363ffffffff16811015611217576040805160148082528183019092525f916020820181803683370190505090505f5b60148110156111c9578a61117e63ffffffff8b1683611e07565b8151811061118e5761118e611df3565b602001015160f81c60f81b8282815181106111ab576111ab611df3565b60200101906001600160f81b03191690815f1a905350600101611164565b505f60148201519050808484815181106111e5576111e5611df3565b6001600160a01b039092166020928302919091019091015261120860148a611f61565b98505050806001019050611134565b506040805180820190915263ffffffff9092168252602082015260808401525f80805b60048110156112995761124e816003611dc9565b611259906008611ddc565b63ffffffff16898863ffffffff16836112729190611e07565b8151811061128257611282611df3565b016020015160f81c901b919091179060010161123a565b506112a5600487611f61565b95505f5b6004811015611308576112bd816003611dc9565b6112c8906008611ddc565b63ffffffff16898863ffffffff16836112e19190611e07565b815181106112f1576112f1611df3565b016020015160f81c901b92909217916001016112a9565b50611314600487611f61565b95505f8263ffffffff166001600160401b038111156113355761133561180b565b60405190808252806020026020018201604052801561135e578160200160208202803683370190505b5090505f5b8363ffffffff16811015611446576040805160148082528183019092525f916020820181803683370190505090505f5b60148110156113f8578b6113ad63ffffffff8c1683611e07565b815181106113bd576113bd611df3565b602001015160f81c60f81b8282815181106113da576113da611df3565b60200101906001600160f81b03191690815f1a905350600101611393565b505f601482015190508084848151811061141457611414611df3565b6001600160a01b039092166020928302919091019091015261143760148b611f61565b99505050806001019050611363565b506040805180820190915263ffffffff9092168252602082015260a08501525f6114708284611f61565b61147b906014611f7e565b61148685607a611f61565b6114909190611f61565b90508063ffffffff168851146114cc57875160405163cc92daa160e01b815263ffffffff918216600482015290821660248201526044016102bb565b5f805b600881101561152f576114e3816007611dc9565b6114ee906008611ddc565b6001600160401b03168a61150863ffffffff8b1684611e07565b8151811061151857611518611df3565b016020015160f81c901b91909117906001016114cf565b506001600160401b031660c086015250929695505050505050565b5f60608260400151516030146115735760405163180ffa0d60e01b815260040160405180910390fd5b82516020808501518051604080880151606089015160808a01518051908701515193515f986115b4988a986001989297929690959094909390929101611fa6565b60405160208183030381529060405290505f5b84608001516020015151811015611626578185608001516020015182815181106115f3576115f3611df3565b602002602001015160405160200161160c929190612060565b60408051601f1981840301815291905291506001016115c7565b5060a0840151805160209182015151604051611646938593929101612096565b60405160208183030381529060405290505f5b8460a0015160200151518110156116b857818560a0015160200151828151811061168557611685611df3565b602002602001015160405160200161169e929190612060565b60408051601f198184030181529190529150600101611659565b5060c08401516040516116cf9183916020016120d1565b60405160208183030381529060405290506002816040516116f09190612102565b602060405180830381855afa15801561170b573d5f803e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061172e9190612113565b94909350915050565b6040515f602082018190526022820152602681018390526001600160c01b031960c083901b166046820152606090604e015b60405160208183030381529060405290505b92915050565b6040515f6020820152600160e11b60228201526026810183905281151560f81b6046820152606090604701611769565b6040805160e0810182525f808252606060208084018290528385018290528184018390528451808601865283815280820183905260808501528451808601909552918452908301529060a082019081525f60209091015290565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156118415761184161180b565b60405290565b60405160e081016001600160401b03811182821017156118415761184161180b565b604051601f8201601f191681016001600160401b03811182821017156118915761189161180b565b604052919050565b5f82601f8301126118a8575f80fd5b81356001600160401b038111156118c1576118c161180b565b6118d4601f8201601f1916602001611869565b8181528460208386010111156118e8575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60208284031215611914575f80fd5b81356001600160401b03811115611929575f80fd5b61193584828501611899565b949350505050565b5f6020828403121561194d575f80fd5b5035919050565b5f5b8381101561196e578181015183820152602001611956565b50505f910152565b5f815180845261198d816020860160208601611954565b601f01601f19169290920160200192915050565b602081525f6119b36020830184611976565b9392505050565b80356001600160401b03811681146119d0575f80fd5b919050565b5f805f606084860312156119e7575f80fd5b833592506119f7602085016119ba565b9150611a05604085016119ba565b90509250925092565b5f60208284031215611a1e575f80fd5b81356001600160401b03811115611a33575f80fd5b8201608081850312156119b3575f80fd5b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015611a9f5784516001600160a01b03168252938301936001929092019190830190611a76565b509695505050505050565b60208152815160208201525f602083015160e06040840152611ad0610100840182611976565b90506040840151601f1980858403016060860152611aee8383611976565b92506001600160401b03606087015116608086015260808601519150808584030160a0860152611b1e8383611a44565b925060a08601519150808584030160c086015250611b3c8282611a44565b91505060c0840151611b5960e08501826001600160401b03169052565b509392505050565b80356001600160a01b03811681146119d0575f80fd5b5f60408284031215611b87575f80fd5b611b8f61181f565b9050813563ffffffff81168114611ba4575f80fd5b81526020828101356001600160401b0380821115611bc0575f80fd5b818501915085601f830112611bd3575f80fd5b813581811115611be557611be561180b565b8060051b9150611bf6848301611869565b8181529183018401918481019088841115611c0f575f80fd5b938501935b83851015611c3457611c2585611b61565b82529385019390850190611c14565b808688015250505050505092915050565b5f60208284031215611c55575f80fd5b81356001600160401b0380821115611c6b575f80fd5b9083019060e08286031215611c7e575f80fd5b611c86611847565b82358152602083013582811115611c9b575f80fd5b611ca787828601611899565b602083015250604083013582811115611cbe575f80fd5b611cca87828601611899565b604083015250611cdc606084016119ba565b6060820152608083013582811115611cf2575f80fd5b611cfe87828601611b77565b60808301525060a083013582811115611d15575f80fd5b611d2187828601611b77565b60a083015250611d3360c084016119ba565b60c082015295945050505050565b828152604060208201525f6119356040830184611976565b5f8060408385031215611d6a575f80fd5b82359150611d7a602084016119ba565b90509250929050565b5f8060408385031215611d94575f80fd5b8235915060208301358015158114611daa575f80fd5b809150509250929050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561177b5761177b611db5565b808202811582820484141761177b5761177b611db5565b634e487b7160e01b5f52603260045260245ffd5b8082018082111561177b5761177b611db5565b5f60208284031215611e2a575f80fd5b6119b382611b61565b5f808335601e19843603018112611e48575f80fd5b8301803591506001600160401b03821115611e61575f80fd5b6020019150600581901b3603821315611e78575f80fd5b9250929050565b5f8235605e19833603018112611e93575f80fd5b9190910192915050565b5f808335601e19843603018112611eb2575f80fd5b8301803591506001600160401b03821115611ecb575f80fd5b602001915036819003821315611e78575f80fd5b5f60208284031215611eef575f80fd5b6119b3826119ba565b5f8851611f09818460208d01611954565b60e089901b6001600160e01b031916908301908152868860048301378681019050600481015f8152858782375060c09390931b6001600160c01b0319166004939094019283019390935250600c019695505050505050565b63ffffffff8181168382160190808211156109a2576109a2611db5565b63ffffffff818116838216028082169190828114611f9e57611f9e611db5565b505092915050565b61ffff60f01b8a60f01b1681525f63ffffffff60e01b808b60e01b166002840152896006840152808960e01b166026840152508651611fec81602a850160208b01611954565b86519083019061200381602a840160208b01611954565b60c087901b6001600160c01b031916602a9290910191820152612035603282018660e01b6001600160e01b0319169052565b61204e603682018560e01b6001600160e01b0319169052565b603a019b9a5050505050505050505050565b5f8351612071818460208801611954565b60609390931b6bffffffffffffffffffffffff19169190920190815260140192915050565b5f84516120a7818460208901611954565b6001600160e01b031960e095861b8116919093019081529290931b16600482015260080192915050565b5f83516120e2818460208801611954565b60c09390931b6001600160c01b0319169190920190815260080192915050565b5f8251611e93818460208701611954565b5f60208284031215612123575f80fd5b505191905056fea2646970667358221220c6870ddd810b3f0c78fe8c26d4cde70582b0e1033651ad0d4b083666233a715064736f6c63430008190033",
}

// ValidatorMessagesABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorMessagesMetaData.ABI instead.
var ValidatorMessagesABI = ValidatorMessagesMetaData.ABI

// ValidatorMessagesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorMessagesMetaData.Bin instead.
var ValidatorMessagesBin = ValidatorMessagesMetaData.Bin

// DeployValidatorMessages deploys a new Ethereum contract, binding an instance of ValidatorMessages to it.
func DeployValidatorMessages(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidatorMessages, error) {
	parsed, err := ValidatorMessagesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorMessagesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorMessages{ValidatorMessagesCaller: ValidatorMessagesCaller{contract: contract}, ValidatorMessagesTransactor: ValidatorMessagesTransactor{contract: contract}, ValidatorMessagesFilterer: ValidatorMessagesFilterer{contract: contract}}, nil
}

// ValidatorMessages is an auto generated Go binding around an Ethereum contract.
type ValidatorMessages struct {
	ValidatorMessagesCaller     // Read-only binding to the contract
	ValidatorMessagesTransactor // Write-only binding to the contract
	ValidatorMessagesFilterer   // Log filterer for contract events
}

// ValidatorMessagesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorMessagesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMessagesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorMessagesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMessagesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorMessagesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorMessagesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorMessagesSession struct {
	Contract     *ValidatorMessages // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ValidatorMessagesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorMessagesCallerSession struct {
	Contract *ValidatorMessagesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ValidatorMessagesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorMessagesTransactorSession struct {
	Contract     *ValidatorMessagesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ValidatorMessagesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorMessagesRaw struct {
	Contract *ValidatorMessages // Generic contract binding to access the raw methods on
}

// ValidatorMessagesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorMessagesCallerRaw struct {
	Contract *ValidatorMessagesCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorMessagesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorMessagesTransactorRaw struct {
	Contract *ValidatorMessagesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorMessages creates a new instance of ValidatorMessages, bound to a specific deployed contract.
func NewValidatorMessages(address common.Address, backend bind.ContractBackend) (*ValidatorMessages, error) {
	contract, err := bindValidatorMessages(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorMessages{ValidatorMessagesCaller: ValidatorMessagesCaller{contract: contract}, ValidatorMessagesTransactor: ValidatorMessagesTransactor{contract: contract}, ValidatorMessagesFilterer: ValidatorMessagesFilterer{contract: contract}}, nil
}

// NewValidatorMessagesCaller creates a new read-only instance of ValidatorMessages, bound to a specific deployed contract.
func NewValidatorMessagesCaller(address common.Address, caller bind.ContractCaller) (*ValidatorMessagesCaller, error) {
	contract, err := bindValidatorMessages(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorMessagesCaller{contract: contract}, nil
}

// NewValidatorMessagesTransactor creates a new write-only instance of ValidatorMessages, bound to a specific deployed contract.
func NewValidatorMessagesTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorMessagesTransactor, error) {
	contract, err := bindValidatorMessages(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorMessagesTransactor{contract: contract}, nil
}

// NewValidatorMessagesFilterer creates a new log filterer instance of ValidatorMessages, bound to a specific deployed contract.
func NewValidatorMessagesFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorMessagesFilterer, error) {
	contract, err := bindValidatorMessages(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorMessagesFilterer{contract: contract}, nil
}

// bindValidatorMessages binds a generic wrapper to an already deployed contract.
func bindValidatorMessages(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorMessagesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorMessages *ValidatorMessagesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorMessages.Contract.ValidatorMessagesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorMessages *ValidatorMessagesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMessages.Contract.ValidatorMessagesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorMessages *ValidatorMessagesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorMessages.Contract.ValidatorMessagesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorMessages *ValidatorMessagesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorMessages.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorMessages *ValidatorMessagesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorMessages.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorMessages *ValidatorMessagesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorMessages.Contract.contract.Transact(opts, method, params...)
}

// PackRegisterSubnetValidatorMessage is a free data retrieval call binding the contract method 0x01bbec74.
//
// Solidity: function packRegisterSubnetValidatorMessage((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64) validationPeriod) pure returns(bytes32, bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackRegisterSubnetValidatorMessage(opts *bind.CallOpts, validationPeriod ValidatorMessagesValidationPeriod) ([32]byte, []byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packRegisterSubnetValidatorMessage", validationPeriod)

	if err != nil {
		return *new([32]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// PackRegisterSubnetValidatorMessage is a free data retrieval call binding the contract method 0x01bbec74.
//
// Solidity: function packRegisterSubnetValidatorMessage((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64) validationPeriod) pure returns(bytes32, bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackRegisterSubnetValidatorMessage(validationPeriod ValidatorMessagesValidationPeriod) ([32]byte, []byte, error) {
	return _ValidatorMessages.Contract.PackRegisterSubnetValidatorMessage(&_ValidatorMessages.CallOpts, validationPeriod)
}

// PackRegisterSubnetValidatorMessage is a free data retrieval call binding the contract method 0x01bbec74.
//
// Solidity: function packRegisterSubnetValidatorMessage((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64) validationPeriod) pure returns(bytes32, bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackRegisterSubnetValidatorMessage(validationPeriod ValidatorMessagesValidationPeriod) ([32]byte, []byte, error) {
	return _ValidatorMessages.Contract.PackRegisterSubnetValidatorMessage(&_ValidatorMessages.CallOpts, validationPeriod)
}

// PackSubnetConversionData is a free data retrieval call binding the contract method 0xf65e4b33.
//
// Solidity: function packSubnetConversionData((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackSubnetConversionData(opts *bind.CallOpts, subnetConversionData SubnetConversionData) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packSubnetConversionData", subnetConversionData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackSubnetConversionData is a free data retrieval call binding the contract method 0xf65e4b33.
//
// Solidity: function packSubnetConversionData((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackSubnetConversionData(subnetConversionData SubnetConversionData) ([]byte, error) {
	return _ValidatorMessages.Contract.PackSubnetConversionData(&_ValidatorMessages.CallOpts, subnetConversionData)
}

// PackSubnetConversionData is a free data retrieval call binding the contract method 0xf65e4b33.
//
// Solidity: function packSubnetConversionData((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackSubnetConversionData(subnetConversionData SubnetConversionData) ([]byte, error) {
	return _ValidatorMessages.Contract.PackSubnetConversionData(&_ValidatorMessages.CallOpts, subnetConversionData)
}

// PackSubnetConversionMessage is a free data retrieval call binding the contract method 0x1fd979c7.
//
// Solidity: function packSubnetConversionMessage(bytes32 subnetConversionID) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackSubnetConversionMessage(opts *bind.CallOpts, subnetConversionID [32]byte) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packSubnetConversionMessage", subnetConversionID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackSubnetConversionMessage is a free data retrieval call binding the contract method 0x1fd979c7.
//
// Solidity: function packSubnetConversionMessage(bytes32 subnetConversionID) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackSubnetConversionMessage(subnetConversionID [32]byte) ([]byte, error) {
	return _ValidatorMessages.Contract.PackSubnetConversionMessage(&_ValidatorMessages.CallOpts, subnetConversionID)
}

// PackSubnetConversionMessage is a free data retrieval call binding the contract method 0x1fd979c7.
//
// Solidity: function packSubnetConversionMessage(bytes32 subnetConversionID) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackSubnetConversionMessage(subnetConversionID [32]byte) ([]byte, error) {
	return _ValidatorMessages.Contract.PackSubnetConversionMessage(&_ValidatorMessages.CallOpts, subnetConversionID)
}

// PackSubnetValidatorRegistrationMessage is a free data retrieval call binding the contract method 0xfa1f8dfb.
//
// Solidity: function packSubnetValidatorRegistrationMessage(bytes32 validationID, bool valid) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackSubnetValidatorRegistrationMessage(opts *bind.CallOpts, validationID [32]byte, valid bool) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packSubnetValidatorRegistrationMessage", validationID, valid)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackSubnetValidatorRegistrationMessage is a free data retrieval call binding the contract method 0xfa1f8dfb.
//
// Solidity: function packSubnetValidatorRegistrationMessage(bytes32 validationID, bool valid) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackSubnetValidatorRegistrationMessage(validationID [32]byte, valid bool) ([]byte, error) {
	return _ValidatorMessages.Contract.PackSubnetValidatorRegistrationMessage(&_ValidatorMessages.CallOpts, validationID, valid)
}

// PackSubnetValidatorRegistrationMessage is a free data retrieval call binding the contract method 0xfa1f8dfb.
//
// Solidity: function packSubnetValidatorRegistrationMessage(bytes32 validationID, bool valid) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackSubnetValidatorRegistrationMessage(validationID [32]byte, valid bool) ([]byte, error) {
	return _ValidatorMessages.Contract.PackSubnetValidatorRegistrationMessage(&_ValidatorMessages.CallOpts, validationID, valid)
}

// PackSubnetValidatorWeightMessage is a free data retrieval call binding the contract method 0x8545c16a.
//
// Solidity: function packSubnetValidatorWeightMessage(bytes32 validationID, uint64 nonce, uint64 weight) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackSubnetValidatorWeightMessage(opts *bind.CallOpts, validationID [32]byte, nonce uint64, weight uint64) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packSubnetValidatorWeightMessage", validationID, nonce, weight)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackSubnetValidatorWeightMessage is a free data retrieval call binding the contract method 0x8545c16a.
//
// Solidity: function packSubnetValidatorWeightMessage(bytes32 validationID, uint64 nonce, uint64 weight) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackSubnetValidatorWeightMessage(validationID [32]byte, nonce uint64, weight uint64) ([]byte, error) {
	return _ValidatorMessages.Contract.PackSubnetValidatorWeightMessage(&_ValidatorMessages.CallOpts, validationID, nonce, weight)
}

// PackSubnetValidatorWeightMessage is a free data retrieval call binding the contract method 0x8545c16a.
//
// Solidity: function packSubnetValidatorWeightMessage(bytes32 validationID, uint64 nonce, uint64 weight) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackSubnetValidatorWeightMessage(validationID [32]byte, nonce uint64, weight uint64) ([]byte, error) {
	return _ValidatorMessages.Contract.PackSubnetValidatorWeightMessage(&_ValidatorMessages.CallOpts, validationID, nonce, weight)
}

// PackValidationUptimeMessage is a free data retrieval call binding the contract method 0xe1d68f30.
//
// Solidity: function packValidationUptimeMessage(bytes32 validationID, uint64 uptime) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCaller) PackValidationUptimeMessage(opts *bind.CallOpts, validationID [32]byte, uptime uint64) ([]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "packValidationUptimeMessage", validationID, uptime)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackValidationUptimeMessage is a free data retrieval call binding the contract method 0xe1d68f30.
//
// Solidity: function packValidationUptimeMessage(bytes32 validationID, uint64 uptime) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesSession) PackValidationUptimeMessage(validationID [32]byte, uptime uint64) ([]byte, error) {
	return _ValidatorMessages.Contract.PackValidationUptimeMessage(&_ValidatorMessages.CallOpts, validationID, uptime)
}

// PackValidationUptimeMessage is a free data retrieval call binding the contract method 0xe1d68f30.
//
// Solidity: function packValidationUptimeMessage(bytes32 validationID, uint64 uptime) pure returns(bytes)
func (_ValidatorMessages *ValidatorMessagesCallerSession) PackValidationUptimeMessage(validationID [32]byte, uptime uint64) ([]byte, error) {
	return _ValidatorMessages.Contract.PackValidationUptimeMessage(&_ValidatorMessages.CallOpts, validationID, uptime)
}

// UnpackRegisterSubnetValidatorMessage is a free data retrieval call binding the contract method 0xa5233770.
//
// Solidity: function unpackRegisterSubnetValidatorMessage(bytes input) pure returns((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64))
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackRegisterSubnetValidatorMessage(opts *bind.CallOpts, input []byte) (ValidatorMessagesValidationPeriod, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackRegisterSubnetValidatorMessage", input)

	if err != nil {
		return *new(ValidatorMessagesValidationPeriod), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorMessagesValidationPeriod)).(*ValidatorMessagesValidationPeriod)

	return out0, err

}

// UnpackRegisterSubnetValidatorMessage is a free data retrieval call binding the contract method 0xa5233770.
//
// Solidity: function unpackRegisterSubnetValidatorMessage(bytes input) pure returns((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64))
func (_ValidatorMessages *ValidatorMessagesSession) UnpackRegisterSubnetValidatorMessage(input []byte) (ValidatorMessagesValidationPeriod, error) {
	return _ValidatorMessages.Contract.UnpackRegisterSubnetValidatorMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackRegisterSubnetValidatorMessage is a free data retrieval call binding the contract method 0xa5233770.
//
// Solidity: function unpackRegisterSubnetValidatorMessage(bytes input) pure returns((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64))
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackRegisterSubnetValidatorMessage(input []byte) (ValidatorMessagesValidationPeriod, error) {
	return _ValidatorMessages.Contract.UnpackRegisterSubnetValidatorMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackSubnetConversionMessage is a free data retrieval call binding the contract method 0x1e6d9789.
//
// Solidity: function unpackSubnetConversionMessage(bytes input) pure returns(bytes32)
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackSubnetConversionMessage(opts *bind.CallOpts, input []byte) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackSubnetConversionMessage", input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UnpackSubnetConversionMessage is a free data retrieval call binding the contract method 0x1e6d9789.
//
// Solidity: function unpackSubnetConversionMessage(bytes input) pure returns(bytes32)
func (_ValidatorMessages *ValidatorMessagesSession) UnpackSubnetConversionMessage(input []byte) ([32]byte, error) {
	return _ValidatorMessages.Contract.UnpackSubnetConversionMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackSubnetConversionMessage is a free data retrieval call binding the contract method 0x1e6d9789.
//
// Solidity: function unpackSubnetConversionMessage(bytes input) pure returns(bytes32)
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackSubnetConversionMessage(input []byte) ([32]byte, error) {
	return _ValidatorMessages.Contract.UnpackSubnetConversionMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackSubnetValidatorRegistrationMessage is a free data retrieval call binding the contract method 0x2e43ceb5.
//
// Solidity: function unpackSubnetValidatorRegistrationMessage(bytes input) pure returns(bytes32, bool)
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackSubnetValidatorRegistrationMessage(opts *bind.CallOpts, input []byte) ([32]byte, bool, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackSubnetValidatorRegistrationMessage", input)

	if err != nil {
		return *new([32]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// UnpackSubnetValidatorRegistrationMessage is a free data retrieval call binding the contract method 0x2e43ceb5.
//
// Solidity: function unpackSubnetValidatorRegistrationMessage(bytes input) pure returns(bytes32, bool)
func (_ValidatorMessages *ValidatorMessagesSession) UnpackSubnetValidatorRegistrationMessage(input []byte) ([32]byte, bool, error) {
	return _ValidatorMessages.Contract.UnpackSubnetValidatorRegistrationMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackSubnetValidatorRegistrationMessage is a free data retrieval call binding the contract method 0x2e43ceb5.
//
// Solidity: function unpackSubnetValidatorRegistrationMessage(bytes input) pure returns(bytes32, bool)
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackSubnetValidatorRegistrationMessage(input []byte) ([32]byte, bool, error) {
	return _ValidatorMessages.Contract.UnpackSubnetValidatorRegistrationMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackSubnetValidatorWeightMessage is a free data retrieval call binding the contract method 0x9de23d40.
//
// Solidity: function unpackSubnetValidatorWeightMessage(bytes input) pure returns(bytes32, uint64, uint64)
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackSubnetValidatorWeightMessage(opts *bind.CallOpts, input []byte) ([32]byte, uint64, uint64, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackSubnetValidatorWeightMessage", input)

	if err != nil {
		return *new([32]byte), *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return out0, out1, out2, err

}

// UnpackSubnetValidatorWeightMessage is a free data retrieval call binding the contract method 0x9de23d40.
//
// Solidity: function unpackSubnetValidatorWeightMessage(bytes input) pure returns(bytes32, uint64, uint64)
func (_ValidatorMessages *ValidatorMessagesSession) UnpackSubnetValidatorWeightMessage(input []byte) ([32]byte, uint64, uint64, error) {
	return _ValidatorMessages.Contract.UnpackSubnetValidatorWeightMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackSubnetValidatorWeightMessage is a free data retrieval call binding the contract method 0x9de23d40.
//
// Solidity: function unpackSubnetValidatorWeightMessage(bytes input) pure returns(bytes32, uint64, uint64)
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackSubnetValidatorWeightMessage(input []byte) ([32]byte, uint64, uint64, error) {
	return _ValidatorMessages.Contract.UnpackSubnetValidatorWeightMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackValidationUptimeMessage is a free data retrieval call binding the contract method 0x088c2463.
//
// Solidity: function unpackValidationUptimeMessage(bytes input) pure returns(bytes32, uint64)
func (_ValidatorMessages *ValidatorMessagesCaller) UnpackValidationUptimeMessage(opts *bind.CallOpts, input []byte) ([32]byte, uint64, error) {
	var out []interface{}
	err := _ValidatorMessages.contract.Call(opts, &out, "unpackValidationUptimeMessage", input)

	if err != nil {
		return *new([32]byte), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// UnpackValidationUptimeMessage is a free data retrieval call binding the contract method 0x088c2463.
//
// Solidity: function unpackValidationUptimeMessage(bytes input) pure returns(bytes32, uint64)
func (_ValidatorMessages *ValidatorMessagesSession) UnpackValidationUptimeMessage(input []byte) ([32]byte, uint64, error) {
	return _ValidatorMessages.Contract.UnpackValidationUptimeMessage(&_ValidatorMessages.CallOpts, input)
}

// UnpackValidationUptimeMessage is a free data retrieval call binding the contract method 0x088c2463.
//
// Solidity: function unpackValidationUptimeMessage(bytes input) pure returns(bytes32, uint64)
func (_ValidatorMessages *ValidatorMessagesCallerSession) UnpackValidationUptimeMessage(input []byte) ([32]byte, uint64, error) {
	return _ValidatorMessages.Contract.UnpackValidationUptimeMessage(&_ValidatorMessages.CallOpts, input)
}
