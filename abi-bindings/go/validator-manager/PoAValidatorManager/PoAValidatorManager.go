// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poavalidatormanager

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

// PoAValidatorManagerMetaData contains all meta data concerning the PoAValidatorManager contract.
var PoAValidatorManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumICMInitializable\",\"name\":\"init\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"handleCompleteEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"handleCompleteValidatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"handleCompleteValidatorWeightChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"handleInitializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"handleInitializeValidatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"handleInitializeValidatorWeightChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516106f43803806106f483398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b6105a78061014d5f395ff3fe608060405234801561000f575f80fd5b506004361061009b575f3560e01c8063c4d66de811610063578063c4d66de814610120578063d8761f4c14610133578063effb55b41461010d578063f2fde38b14610141578063f7795a8c1461009f575f80fd5b806356ab447d1461009f578063715018a6146100b25780638da5cb5b146100ba57806398d9fa7c146100f8578063bddb5a141461010d575b5f80fd5b6100b06100ad3660046103ed565b50565b005b6100b0610154565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054604080516001600160a01b039092168252519081900360200190f35b6100b0610106366004610465565b5050505050565b6100b061011b3660046104cf565b505050565b6100b061012e36600461052d565b610167565b6100b061010636600461054d565b6100b061014f36600461052d565b610275565b61015c6102b4565b6101655f61030f565b565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f811580156101ac5750825b90505f8267ffffffffffffffff1660011480156101c85750303b155b9050811580156101d6575080155b156101f45760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561021e57845460ff60401b1916600160401b1785555b6102278661037f565b831561026d57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b61027d6102b4565b6001600160a01b0381166102ab57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b6100ad8161030f565b336102e67f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146101655760405163118cdaa760e01b81523360048201526024016102a2565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b610387610390565b6100ad816103d9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661016557604051631afcd79f60e31b815260040160405180910390fd5b6103e1610390565b6100ad8161027d610390565b5f602082840312156103fd575f80fd5b5035919050565b803567ffffffffffffffff8116811461041b575f80fd5b919050565b5f8083601f840112610430575f80fd5b50813567ffffffffffffffff811115610447575f80fd5b60208301915083602082850101111561045e575f80fd5b9250929050565b5f805f805f60808688031215610479575f80fd5b8535945061048960208701610404565b935061049760408701610404565b9250606086013567ffffffffffffffff8111156104b2575f80fd5b6104be88828901610420565b969995985093965092949392505050565b5f805f604084860312156104e1575f80fd5b83359250602084013567ffffffffffffffff8111156104fe575f80fd5b61050a86828701610420565b9497909650939450505050565b80356001600160a01b038116811461041b575f80fd5b5f6020828403121561053d575f80fd5b61054682610517565b9392505050565b5f805f805f60808688031215610561575f80fd5b853594506104896020870161051756fea2646970667358221220b569989b7d5b5408c504449f77e2ca6c675b871cbf569a0d293f196d3c2ae5b664736f6c63430008190033",
}

// PoAValidatorManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PoAValidatorManagerMetaData.ABI instead.
var PoAValidatorManagerABI = PoAValidatorManagerMetaData.ABI

// PoAValidatorManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoAValidatorManagerMetaData.Bin instead.
var PoAValidatorManagerBin = PoAValidatorManagerMetaData.Bin

// DeployPoAValidatorManager deploys a new Ethereum contract, binding an instance of PoAValidatorManager to it.
func DeployPoAValidatorManager(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *PoAValidatorManager, error) {
	parsed, err := PoAValidatorManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoAValidatorManagerBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoAValidatorManager{PoAValidatorManagerCaller: PoAValidatorManagerCaller{contract: contract}, PoAValidatorManagerTransactor: PoAValidatorManagerTransactor{contract: contract}, PoAValidatorManagerFilterer: PoAValidatorManagerFilterer{contract: contract}}, nil
}

// PoAValidatorManager is an auto generated Go binding around an Ethereum contract.
type PoAValidatorManager struct {
	PoAValidatorManagerCaller     // Read-only binding to the contract
	PoAValidatorManagerTransactor // Write-only binding to the contract
	PoAValidatorManagerFilterer   // Log filterer for contract events
}

// PoAValidatorManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoAValidatorManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoAValidatorManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoAValidatorManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoAValidatorManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoAValidatorManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoAValidatorManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoAValidatorManagerSession struct {
	Contract     *PoAValidatorManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PoAValidatorManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoAValidatorManagerCallerSession struct {
	Contract *PoAValidatorManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PoAValidatorManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoAValidatorManagerTransactorSession struct {
	Contract     *PoAValidatorManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PoAValidatorManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoAValidatorManagerRaw struct {
	Contract *PoAValidatorManager // Generic contract binding to access the raw methods on
}

// PoAValidatorManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoAValidatorManagerCallerRaw struct {
	Contract *PoAValidatorManagerCaller // Generic read-only contract binding to access the raw methods on
}

// PoAValidatorManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoAValidatorManagerTransactorRaw struct {
	Contract *PoAValidatorManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoAValidatorManager creates a new instance of PoAValidatorManager, bound to a specific deployed contract.
func NewPoAValidatorManager(address common.Address, backend bind.ContractBackend) (*PoAValidatorManager, error) {
	contract, err := bindPoAValidatorManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManager{PoAValidatorManagerCaller: PoAValidatorManagerCaller{contract: contract}, PoAValidatorManagerTransactor: PoAValidatorManagerTransactor{contract: contract}, PoAValidatorManagerFilterer: PoAValidatorManagerFilterer{contract: contract}}, nil
}

// NewPoAValidatorManagerCaller creates a new read-only instance of PoAValidatorManager, bound to a specific deployed contract.
func NewPoAValidatorManagerCaller(address common.Address, caller bind.ContractCaller) (*PoAValidatorManagerCaller, error) {
	contract, err := bindPoAValidatorManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerCaller{contract: contract}, nil
}

// NewPoAValidatorManagerTransactor creates a new write-only instance of PoAValidatorManager, bound to a specific deployed contract.
func NewPoAValidatorManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PoAValidatorManagerTransactor, error) {
	contract, err := bindPoAValidatorManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerTransactor{contract: contract}, nil
}

// NewPoAValidatorManagerFilterer creates a new log filterer instance of PoAValidatorManager, bound to a specific deployed contract.
func NewPoAValidatorManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PoAValidatorManagerFilterer, error) {
	contract, err := bindPoAValidatorManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerFilterer{contract: contract}, nil
}

// bindPoAValidatorManager binds a generic wrapper to an already deployed contract.
func bindPoAValidatorManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoAValidatorManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoAValidatorManager *PoAValidatorManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoAValidatorManager.Contract.PoAValidatorManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoAValidatorManager *PoAValidatorManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.PoAValidatorManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoAValidatorManager *PoAValidatorManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.PoAValidatorManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoAValidatorManager *PoAValidatorManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoAValidatorManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoAValidatorManager *PoAValidatorManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoAValidatorManager *PoAValidatorManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoAValidatorManager *PoAValidatorManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoAValidatorManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoAValidatorManager *PoAValidatorManagerSession) Owner() (common.Address, error) {
	return _PoAValidatorManager.Contract.Owner(&_PoAValidatorManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoAValidatorManager *PoAValidatorManagerCallerSession) Owner() (common.Address, error) {
	return _PoAValidatorManager.Contract.Owner(&_PoAValidatorManager.CallOpts)
}

// HandleCompleteEndValidation is a paid mutator transaction binding the contract method 0xf7795a8c.
//
// Solidity: function handleCompleteEndValidation(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) HandleCompleteEndValidation(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "handleCompleteEndValidation", validationID)
}

// HandleCompleteEndValidation is a paid mutator transaction binding the contract method 0xf7795a8c.
//
// Solidity: function handleCompleteEndValidation(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) HandleCompleteEndValidation(validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.HandleCompleteEndValidation(&_PoAValidatorManager.TransactOpts, validationID)
}

// HandleCompleteEndValidation is a paid mutator transaction binding the contract method 0xf7795a8c.
//
// Solidity: function handleCompleteEndValidation(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) HandleCompleteEndValidation(validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.HandleCompleteEndValidation(&_PoAValidatorManager.TransactOpts, validationID)
}

// HandleCompleteValidatorRegistration is a paid mutator transaction binding the contract method 0x56ab447d.
//
// Solidity: function handleCompleteValidatorRegistration(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) HandleCompleteValidatorRegistration(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "handleCompleteValidatorRegistration", validationID)
}

// HandleCompleteValidatorRegistration is a paid mutator transaction binding the contract method 0x56ab447d.
//
// Solidity: function handleCompleteValidatorRegistration(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) HandleCompleteValidatorRegistration(validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.HandleCompleteValidatorRegistration(&_PoAValidatorManager.TransactOpts, validationID)
}

// HandleCompleteValidatorRegistration is a paid mutator transaction binding the contract method 0x56ab447d.
//
// Solidity: function handleCompleteValidatorRegistration(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) HandleCompleteValidatorRegistration(validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.HandleCompleteValidatorRegistration(&_PoAValidatorManager.TransactOpts, validationID)
}

// HandleCompleteValidatorWeightChange is a paid mutator transaction binding the contract method 0xeffb55b4.
//
// Solidity: function handleCompleteValidatorWeightChange(bytes32 validationID, bytes args) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) HandleCompleteValidatorWeightChange(opts *bind.TransactOpts, validationID [32]byte, args []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "handleCompleteValidatorWeightChange", validationID, args)
}

// HandleCompleteValidatorWeightChange is a paid mutator transaction binding the contract method 0xeffb55b4.
//
// Solidity: function handleCompleteValidatorWeightChange(bytes32 validationID, bytes args) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) HandleCompleteValidatorWeightChange(validationID [32]byte, args []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.HandleCompleteValidatorWeightChange(&_PoAValidatorManager.TransactOpts, validationID, args)
}

// HandleCompleteValidatorWeightChange is a paid mutator transaction binding the contract method 0xeffb55b4.
//
// Solidity: function handleCompleteValidatorWeightChange(bytes32 validationID, bytes args) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) HandleCompleteValidatorWeightChange(validationID [32]byte, args []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.HandleCompleteValidatorWeightChange(&_PoAValidatorManager.TransactOpts, validationID, args)
}

// HandleInitializeEndValidation is a paid mutator transaction binding the contract method 0xbddb5a14.
//
// Solidity: function handleInitializeEndValidation(bytes32 validationID, bytes args) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) HandleInitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, args []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "handleInitializeEndValidation", validationID, args)
}

// HandleInitializeEndValidation is a paid mutator transaction binding the contract method 0xbddb5a14.
//
// Solidity: function handleInitializeEndValidation(bytes32 validationID, bytes args) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) HandleInitializeEndValidation(validationID [32]byte, args []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.HandleInitializeEndValidation(&_PoAValidatorManager.TransactOpts, validationID, args)
}

// HandleInitializeEndValidation is a paid mutator transaction binding the contract method 0xbddb5a14.
//
// Solidity: function handleInitializeEndValidation(bytes32 validationID, bytes args) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) HandleInitializeEndValidation(validationID [32]byte, args []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.HandleInitializeEndValidation(&_PoAValidatorManager.TransactOpts, validationID, args)
}

// HandleInitializeValidatorRegistration is a paid mutator transaction binding the contract method 0xd8761f4c.
//
// Solidity: function handleInitializeValidatorRegistration(bytes32 validationID, address sender, uint64 weight, bytes args) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) HandleInitializeValidatorRegistration(opts *bind.TransactOpts, validationID [32]byte, sender common.Address, weight uint64, args []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "handleInitializeValidatorRegistration", validationID, sender, weight, args)
}

// HandleInitializeValidatorRegistration is a paid mutator transaction binding the contract method 0xd8761f4c.
//
// Solidity: function handleInitializeValidatorRegistration(bytes32 validationID, address sender, uint64 weight, bytes args) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) HandleInitializeValidatorRegistration(validationID [32]byte, sender common.Address, weight uint64, args []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.HandleInitializeValidatorRegistration(&_PoAValidatorManager.TransactOpts, validationID, sender, weight, args)
}

// HandleInitializeValidatorRegistration is a paid mutator transaction binding the contract method 0xd8761f4c.
//
// Solidity: function handleInitializeValidatorRegistration(bytes32 validationID, address sender, uint64 weight, bytes args) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) HandleInitializeValidatorRegistration(validationID [32]byte, sender common.Address, weight uint64, args []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.HandleInitializeValidatorRegistration(&_PoAValidatorManager.TransactOpts, validationID, sender, weight, args)
}

// HandleInitializeValidatorWeightChange is a paid mutator transaction binding the contract method 0x98d9fa7c.
//
// Solidity: function handleInitializeValidatorWeightChange(bytes32 validationID, uint64 weight, uint64 nonce, bytes args) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) HandleInitializeValidatorWeightChange(opts *bind.TransactOpts, validationID [32]byte, weight uint64, nonce uint64, args []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "handleInitializeValidatorWeightChange", validationID, weight, nonce, args)
}

// HandleInitializeValidatorWeightChange is a paid mutator transaction binding the contract method 0x98d9fa7c.
//
// Solidity: function handleInitializeValidatorWeightChange(bytes32 validationID, uint64 weight, uint64 nonce, bytes args) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) HandleInitializeValidatorWeightChange(validationID [32]byte, weight uint64, nonce uint64, args []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.HandleInitializeValidatorWeightChange(&_PoAValidatorManager.TransactOpts, validationID, weight, nonce, args)
}

// HandleInitializeValidatorWeightChange is a paid mutator transaction binding the contract method 0x98d9fa7c.
//
// Solidity: function handleInitializeValidatorWeightChange(bytes32 validationID, uint64 weight, uint64 nonce, bytes args) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) HandleInitializeValidatorWeightChange(validationID [32]byte, weight uint64, nonce uint64, args []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.HandleInitializeValidatorWeightChange(&_PoAValidatorManager.TransactOpts, validationID, weight, nonce, args)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.Initialize(&_PoAValidatorManager.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.Initialize(&_PoAValidatorManager.TransactOpts, initialOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.RenounceOwnership(&_PoAValidatorManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.RenounceOwnership(&_PoAValidatorManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.TransferOwnership(&_PoAValidatorManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.TransferOwnership(&_PoAValidatorManager.TransactOpts, newOwner)
}

// PoAValidatorManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PoAValidatorManager contract.
type PoAValidatorManagerInitializedIterator struct {
	Event *PoAValidatorManagerInitialized // Event containing the contract specifics and raw log

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
func (it *PoAValidatorManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoAValidatorManagerInitialized)
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
		it.Event = new(PoAValidatorManagerInitialized)
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
func (it *PoAValidatorManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoAValidatorManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoAValidatorManagerInitialized represents a Initialized event raised by the PoAValidatorManager contract.
type PoAValidatorManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*PoAValidatorManagerInitializedIterator, error) {

	logs, sub, err := _PoAValidatorManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerInitializedIterator{contract: _PoAValidatorManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PoAValidatorManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _PoAValidatorManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoAValidatorManagerInitialized)
				if err := _PoAValidatorManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PoAValidatorManager *PoAValidatorManagerFilterer) ParseInitialized(log types.Log) (*PoAValidatorManagerInitialized, error) {
	event := new(PoAValidatorManagerInitialized)
	if err := _PoAValidatorManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoAValidatorManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PoAValidatorManager contract.
type PoAValidatorManagerOwnershipTransferredIterator struct {
	Event *PoAValidatorManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PoAValidatorManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoAValidatorManagerOwnershipTransferred)
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
		it.Event = new(PoAValidatorManagerOwnershipTransferred)
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
func (it *PoAValidatorManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoAValidatorManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoAValidatorManagerOwnershipTransferred represents a OwnershipTransferred event raised by the PoAValidatorManager contract.
type PoAValidatorManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PoAValidatorManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerOwnershipTransferredIterator{contract: _PoAValidatorManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PoAValidatorManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoAValidatorManagerOwnershipTransferred)
				if err := _PoAValidatorManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) ParseOwnershipTransferred(log types.Log) (*PoAValidatorManagerOwnershipTransferred, error) {
	event := new(PoAValidatorManagerOwnershipTransferred)
	if err := _PoAValidatorManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
