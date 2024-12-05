// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package argumentsexternal

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

// InitializeEndValidationArgs is an auto generated low-level Go binding around an user-defined struct.
type InitializeEndValidationArgs struct {
	Force              bool
	IncludeUptimeProof bool
	MessageIndex       uint32
	RewardRecipient    common.Address
}

// InitializeValidatorRegistrationArgs is an auto generated low-level Go binding around an user-defined struct.
type InitializeValidatorRegistrationArgs struct {
	DelegationFeeBips uint16
	MinStakeDuration  uint64
}

// ArgumentsExternalMetaData contains all meta data concerning the ArgumentsExternal contract.
var ArgumentsExternalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"decodeInitializeEndValidationArgs\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"rewardRecipient\",\"type\":\"address\"}],\"internalType\":\"structInitializeEndValidationArgs\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"decodeInitializeValidatorRegistrationArgs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"}],\"internalType\":\"structInitializeValidatorRegistrationArgs\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506102f58061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c8063868fba8e14610038578063ae41703014610079575b5f80fd5b61004b610046366004610127565b6100d2565b60408051825161ffff16815260209283015167ffffffffffffffff1692810192909252015b60405180910390f35b61008c610087366004610127565b6100f8565b604051610070919081511515815260208083015115159082015260408083015163ffffffff16908201526060918201516001600160a01b03169181019190915260800190565b604080518082019091525f80825260208201526100f182840184610193565b9392505050565b604080516080810182525f8082526020820181905291810182905260608101919091526100f182840184610221565b5f8060208385031215610138575f80fd5b823567ffffffffffffffff8082111561014f575f80fd5b818501915085601f830112610162575f80fd5b813581811115610170575f80fd5b866020828501011115610181575f80fd5b60209290920196919550909350505050565b5f604082840312156101a3575f80fd5b6040516040810167ffffffffffffffff82821081831117156101d357634e487b7160e01b5f52604160045260245ffd5b816040528435915061ffff821682146101ea575f80fd5b9082526020840135908082168214610200575f80fd5b5060208201529392505050565b8035801515811461021c575f80fd5b919050565b5f60808284031215610231575f80fd5b6040516080810181811067ffffffffffffffff8211171561026057634e487b7160e01b5f52604160045260245ffd5b60405261026c8361020d565b815261027a6020840161020d565b6020820152604083013563ffffffff81168114610295575f80fd5b604082015260608301356001600160a01b03811681146102b3575f80fd5b6060820152939250505056fea2646970667358221220b8d6f6dbc9b54c93fe0d3326345e1f4200f8fa64b8e14054556bbb83799fde4864736f6c63430008190033",
}

// ArgumentsExternalABI is the input ABI used to generate the binding from.
// Deprecated: Use ArgumentsExternalMetaData.ABI instead.
var ArgumentsExternalABI = ArgumentsExternalMetaData.ABI

// ArgumentsExternalBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ArgumentsExternalMetaData.Bin instead.
var ArgumentsExternalBin = ArgumentsExternalMetaData.Bin

// DeployArgumentsExternal deploys a new Ethereum contract, binding an instance of ArgumentsExternal to it.
func DeployArgumentsExternal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArgumentsExternal, error) {
	parsed, err := ArgumentsExternalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArgumentsExternalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArgumentsExternal{ArgumentsExternalCaller: ArgumentsExternalCaller{contract: contract}, ArgumentsExternalTransactor: ArgumentsExternalTransactor{contract: contract}, ArgumentsExternalFilterer: ArgumentsExternalFilterer{contract: contract}}, nil
}

// ArgumentsExternal is an auto generated Go binding around an Ethereum contract.
type ArgumentsExternal struct {
	ArgumentsExternalCaller     // Read-only binding to the contract
	ArgumentsExternalTransactor // Write-only binding to the contract
	ArgumentsExternalFilterer   // Log filterer for contract events
}

// ArgumentsExternalCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArgumentsExternalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArgumentsExternalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArgumentsExternalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArgumentsExternalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArgumentsExternalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArgumentsExternalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArgumentsExternalSession struct {
	Contract     *ArgumentsExternal // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ArgumentsExternalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArgumentsExternalCallerSession struct {
	Contract *ArgumentsExternalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ArgumentsExternalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArgumentsExternalTransactorSession struct {
	Contract     *ArgumentsExternalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ArgumentsExternalRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArgumentsExternalRaw struct {
	Contract *ArgumentsExternal // Generic contract binding to access the raw methods on
}

// ArgumentsExternalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArgumentsExternalCallerRaw struct {
	Contract *ArgumentsExternalCaller // Generic read-only contract binding to access the raw methods on
}

// ArgumentsExternalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArgumentsExternalTransactorRaw struct {
	Contract *ArgumentsExternalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArgumentsExternal creates a new instance of ArgumentsExternal, bound to a specific deployed contract.
func NewArgumentsExternal(address common.Address, backend bind.ContractBackend) (*ArgumentsExternal, error) {
	contract, err := bindArgumentsExternal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArgumentsExternal{ArgumentsExternalCaller: ArgumentsExternalCaller{contract: contract}, ArgumentsExternalTransactor: ArgumentsExternalTransactor{contract: contract}, ArgumentsExternalFilterer: ArgumentsExternalFilterer{contract: contract}}, nil
}

// NewArgumentsExternalCaller creates a new read-only instance of ArgumentsExternal, bound to a specific deployed contract.
func NewArgumentsExternalCaller(address common.Address, caller bind.ContractCaller) (*ArgumentsExternalCaller, error) {
	contract, err := bindArgumentsExternal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArgumentsExternalCaller{contract: contract}, nil
}

// NewArgumentsExternalTransactor creates a new write-only instance of ArgumentsExternal, bound to a specific deployed contract.
func NewArgumentsExternalTransactor(address common.Address, transactor bind.ContractTransactor) (*ArgumentsExternalTransactor, error) {
	contract, err := bindArgumentsExternal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArgumentsExternalTransactor{contract: contract}, nil
}

// NewArgumentsExternalFilterer creates a new log filterer instance of ArgumentsExternal, bound to a specific deployed contract.
func NewArgumentsExternalFilterer(address common.Address, filterer bind.ContractFilterer) (*ArgumentsExternalFilterer, error) {
	contract, err := bindArgumentsExternal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArgumentsExternalFilterer{contract: contract}, nil
}

// bindArgumentsExternal binds a generic wrapper to an already deployed contract.
func bindArgumentsExternal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArgumentsExternalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArgumentsExternal *ArgumentsExternalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArgumentsExternal.Contract.ArgumentsExternalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArgumentsExternal *ArgumentsExternalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArgumentsExternal.Contract.ArgumentsExternalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArgumentsExternal *ArgumentsExternalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArgumentsExternal.Contract.ArgumentsExternalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArgumentsExternal *ArgumentsExternalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArgumentsExternal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArgumentsExternal *ArgumentsExternalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArgumentsExternal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArgumentsExternal *ArgumentsExternalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArgumentsExternal.Contract.contract.Transact(opts, method, params...)
}

// DecodeInitializeEndValidationArgs is a free data retrieval call binding the contract method 0xae417030.
//
// Solidity: function decodeInitializeEndValidationArgs(bytes args) pure returns((bool,bool,uint32,address))
func (_ArgumentsExternal *ArgumentsExternalCaller) DecodeInitializeEndValidationArgs(opts *bind.CallOpts, args []byte) (InitializeEndValidationArgs, error) {
	var out []interface{}
	err := _ArgumentsExternal.contract.Call(opts, &out, "decodeInitializeEndValidationArgs", args)

	if err != nil {
		return *new(InitializeEndValidationArgs), err
	}

	out0 := *abi.ConvertType(out[0], new(InitializeEndValidationArgs)).(*InitializeEndValidationArgs)

	return out0, err

}

// DecodeInitializeEndValidationArgs is a free data retrieval call binding the contract method 0xae417030.
//
// Solidity: function decodeInitializeEndValidationArgs(bytes args) pure returns((bool,bool,uint32,address))
func (_ArgumentsExternal *ArgumentsExternalSession) DecodeInitializeEndValidationArgs(args []byte) (InitializeEndValidationArgs, error) {
	return _ArgumentsExternal.Contract.DecodeInitializeEndValidationArgs(&_ArgumentsExternal.CallOpts, args)
}

// DecodeInitializeEndValidationArgs is a free data retrieval call binding the contract method 0xae417030.
//
// Solidity: function decodeInitializeEndValidationArgs(bytes args) pure returns((bool,bool,uint32,address))
func (_ArgumentsExternal *ArgumentsExternalCallerSession) DecodeInitializeEndValidationArgs(args []byte) (InitializeEndValidationArgs, error) {
	return _ArgumentsExternal.Contract.DecodeInitializeEndValidationArgs(&_ArgumentsExternal.CallOpts, args)
}

// DecodeInitializeValidatorRegistrationArgs is a free data retrieval call binding the contract method 0x868fba8e.
//
// Solidity: function decodeInitializeValidatorRegistrationArgs(bytes args) pure returns((uint16,uint64))
func (_ArgumentsExternal *ArgumentsExternalCaller) DecodeInitializeValidatorRegistrationArgs(opts *bind.CallOpts, args []byte) (InitializeValidatorRegistrationArgs, error) {
	var out []interface{}
	err := _ArgumentsExternal.contract.Call(opts, &out, "decodeInitializeValidatorRegistrationArgs", args)

	if err != nil {
		return *new(InitializeValidatorRegistrationArgs), err
	}

	out0 := *abi.ConvertType(out[0], new(InitializeValidatorRegistrationArgs)).(*InitializeValidatorRegistrationArgs)

	return out0, err

}

// DecodeInitializeValidatorRegistrationArgs is a free data retrieval call binding the contract method 0x868fba8e.
//
// Solidity: function decodeInitializeValidatorRegistrationArgs(bytes args) pure returns((uint16,uint64))
func (_ArgumentsExternal *ArgumentsExternalSession) DecodeInitializeValidatorRegistrationArgs(args []byte) (InitializeValidatorRegistrationArgs, error) {
	return _ArgumentsExternal.Contract.DecodeInitializeValidatorRegistrationArgs(&_ArgumentsExternal.CallOpts, args)
}

// DecodeInitializeValidatorRegistrationArgs is a free data retrieval call binding the contract method 0x868fba8e.
//
// Solidity: function decodeInitializeValidatorRegistrationArgs(bytes args) pure returns((uint16,uint64))
func (_ArgumentsExternal *ArgumentsExternalCallerSession) DecodeInitializeValidatorRegistrationArgs(args []byte) (InitializeValidatorRegistrationArgs, error) {
	return _ArgumentsExternal.Contract.DecodeInitializeValidatorRegistrationArgs(&_ArgumentsExternal.CallOpts, args)
}
