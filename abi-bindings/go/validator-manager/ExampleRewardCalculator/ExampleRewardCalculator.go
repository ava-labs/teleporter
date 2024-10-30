// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package examplerewardcalculator

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

// ExampleRewardCalculatorMetaData contains all meta data concerning the ExampleRewardCalculator contract.
var ExampleRewardCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"rewardBasisPoints_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BIPS_CONVERSION_FACTOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SECONDS_IN_YEAR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateReward\",\"inputs\":[{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stakingStartTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stakingEndTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x60a0604052348015600e575f80fd5b506040516102d43803806102d4833981016040819052602b91603b565b6001600160401b03166080526066565b5f60208284031215604a575f80fd5b81516001600160401b0381168114605f575f80fd5b9392505050565b6080516102516100835f395f818160a0015260fa01526102515ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c80634f22429f1461004e5780635dcc939114610074578063a9778a7a1461007f578063bb65b2421461009b575b5f80fd5b61006161005c366004610172565b6100db565b6040519081526020015b60405180910390f35b6100616301e1338081565b61008861271081565b60405161ffff909116815260200161006b565b6100c27f000000000000000000000000000000000000000000000000000000000000000081565b60405167ffffffffffffffff909116815260200161006b565b5f6127106301e133806100ee86866101e0565b67ffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168961012e9190610208565b6101389190610208565b6101429190610225565b61014c9190610225565b9695505050505050565b803567ffffffffffffffff8116811461016d575f80fd5b919050565b5f805f805f60a08688031215610186575f80fd5b8535945061019660208701610156565b93506101a460408701610156565b92506101b260608701610156565b91506101c060808701610156565b90509295509295909350565b634e487b7160e01b5f52601160045260245ffd5b67ffffffffffffffff828116828216039080821115610201576102016101cc565b5092915050565b808202811582820484141761021f5761021f6101cc565b92915050565b5f8261023f57634e487b7160e01b5f52601260045260245ffd5b50049056fea164736f6c6343000819000a",
}

// ExampleRewardCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleRewardCalculatorMetaData.ABI instead.
var ExampleRewardCalculatorABI = ExampleRewardCalculatorMetaData.ABI

// ExampleRewardCalculatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExampleRewardCalculatorMetaData.Bin instead.
var ExampleRewardCalculatorBin = ExampleRewardCalculatorMetaData.Bin

// DeployExampleRewardCalculator deploys a new Ethereum contract, binding an instance of ExampleRewardCalculator to it.
func DeployExampleRewardCalculator(auth *bind.TransactOpts, backend bind.ContractBackend, rewardBasisPoints_ uint64) (common.Address, *types.Transaction, *ExampleRewardCalculator, error) {
	parsed, err := ExampleRewardCalculatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExampleRewardCalculatorBin), backend, rewardBasisPoints_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExampleRewardCalculator{ExampleRewardCalculatorCaller: ExampleRewardCalculatorCaller{contract: contract}, ExampleRewardCalculatorTransactor: ExampleRewardCalculatorTransactor{contract: contract}, ExampleRewardCalculatorFilterer: ExampleRewardCalculatorFilterer{contract: contract}}, nil
}

// ExampleRewardCalculator is an auto generated Go binding around an Ethereum contract.
type ExampleRewardCalculator struct {
	ExampleRewardCalculatorCaller     // Read-only binding to the contract
	ExampleRewardCalculatorTransactor // Write-only binding to the contract
	ExampleRewardCalculatorFilterer   // Log filterer for contract events
}

// ExampleRewardCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleRewardCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleRewardCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleRewardCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleRewardCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleRewardCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleRewardCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleRewardCalculatorSession struct {
	Contract     *ExampleRewardCalculator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ExampleRewardCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleRewardCalculatorCallerSession struct {
	Contract *ExampleRewardCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ExampleRewardCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleRewardCalculatorTransactorSession struct {
	Contract     *ExampleRewardCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ExampleRewardCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleRewardCalculatorRaw struct {
	Contract *ExampleRewardCalculator // Generic contract binding to access the raw methods on
}

// ExampleRewardCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleRewardCalculatorCallerRaw struct {
	Contract *ExampleRewardCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// ExampleRewardCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleRewardCalculatorTransactorRaw struct {
	Contract *ExampleRewardCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExampleRewardCalculator creates a new instance of ExampleRewardCalculator, bound to a specific deployed contract.
func NewExampleRewardCalculator(address common.Address, backend bind.ContractBackend) (*ExampleRewardCalculator, error) {
	contract, err := bindExampleRewardCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExampleRewardCalculator{ExampleRewardCalculatorCaller: ExampleRewardCalculatorCaller{contract: contract}, ExampleRewardCalculatorTransactor: ExampleRewardCalculatorTransactor{contract: contract}, ExampleRewardCalculatorFilterer: ExampleRewardCalculatorFilterer{contract: contract}}, nil
}

// NewExampleRewardCalculatorCaller creates a new read-only instance of ExampleRewardCalculator, bound to a specific deployed contract.
func NewExampleRewardCalculatorCaller(address common.Address, caller bind.ContractCaller) (*ExampleRewardCalculatorCaller, error) {
	contract, err := bindExampleRewardCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleRewardCalculatorCaller{contract: contract}, nil
}

// NewExampleRewardCalculatorTransactor creates a new write-only instance of ExampleRewardCalculator, bound to a specific deployed contract.
func NewExampleRewardCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ExampleRewardCalculatorTransactor, error) {
	contract, err := bindExampleRewardCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleRewardCalculatorTransactor{contract: contract}, nil
}

// NewExampleRewardCalculatorFilterer creates a new log filterer instance of ExampleRewardCalculator, bound to a specific deployed contract.
func NewExampleRewardCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ExampleRewardCalculatorFilterer, error) {
	contract, err := bindExampleRewardCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleRewardCalculatorFilterer{contract: contract}, nil
}

// bindExampleRewardCalculator binds a generic wrapper to an already deployed contract.
func bindExampleRewardCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExampleRewardCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleRewardCalculator *ExampleRewardCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleRewardCalculator.Contract.ExampleRewardCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleRewardCalculator *ExampleRewardCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleRewardCalculator.Contract.ExampleRewardCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleRewardCalculator *ExampleRewardCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleRewardCalculator.Contract.ExampleRewardCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleRewardCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleRewardCalculator *ExampleRewardCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleRewardCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleRewardCalculator *ExampleRewardCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleRewardCalculator.Contract.contract.Transact(opts, method, params...)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCaller) BIPSCONVERSIONFACTOR(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ExampleRewardCalculator.contract.Call(opts, &out, "BIPS_CONVERSION_FACTOR")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ExampleRewardCalculator *ExampleRewardCalculatorSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _ExampleRewardCalculator.Contract.BIPSCONVERSIONFACTOR(&_ExampleRewardCalculator.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _ExampleRewardCalculator.Contract.BIPSCONVERSIONFACTOR(&_ExampleRewardCalculator.CallOpts)
}

// SECONDSINYEAR is a free data retrieval call binding the contract method 0x5dcc9391.
//
// Solidity: function SECONDS_IN_YEAR() view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCaller) SECONDSINYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExampleRewardCalculator.contract.Call(opts, &out, "SECONDS_IN_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSINYEAR is a free data retrieval call binding the contract method 0x5dcc9391.
//
// Solidity: function SECONDS_IN_YEAR() view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorSession) SECONDSINYEAR() (*big.Int, error) {
	return _ExampleRewardCalculator.Contract.SECONDSINYEAR(&_ExampleRewardCalculator.CallOpts)
}

// SECONDSINYEAR is a free data retrieval call binding the contract method 0x5dcc9391.
//
// Solidity: function SECONDS_IN_YEAR() view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerSession) SECONDSINYEAR() (*big.Int, error) {
	return _ExampleRewardCalculator.Contract.SECONDSINYEAR(&_ExampleRewardCalculator.CallOpts)
}

// CalculateReward is a free data retrieval call binding the contract method 0x4f22429f.
//
// Solidity: function calculateReward(uint256 stakeAmount, uint64 , uint64 stakingStartTime, uint64 stakingEndTime, uint64 ) view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCaller) CalculateReward(opts *bind.CallOpts, stakeAmount *big.Int, arg1 uint64, stakingStartTime uint64, stakingEndTime uint64, arg4 uint64) (*big.Int, error) {
	var out []interface{}
	err := _ExampleRewardCalculator.contract.Call(opts, &out, "calculateReward", stakeAmount, arg1, stakingStartTime, stakingEndTime, arg4)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateReward is a free data retrieval call binding the contract method 0x4f22429f.
//
// Solidity: function calculateReward(uint256 stakeAmount, uint64 , uint64 stakingStartTime, uint64 stakingEndTime, uint64 ) view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorSession) CalculateReward(stakeAmount *big.Int, arg1 uint64, stakingStartTime uint64, stakingEndTime uint64, arg4 uint64) (*big.Int, error) {
	return _ExampleRewardCalculator.Contract.CalculateReward(&_ExampleRewardCalculator.CallOpts, stakeAmount, arg1, stakingStartTime, stakingEndTime, arg4)
}

// CalculateReward is a free data retrieval call binding the contract method 0x4f22429f.
//
// Solidity: function calculateReward(uint256 stakeAmount, uint64 , uint64 stakingStartTime, uint64 stakingEndTime, uint64 ) view returns(uint256)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerSession) CalculateReward(stakeAmount *big.Int, arg1 uint64, stakingStartTime uint64, stakingEndTime uint64, arg4 uint64) (*big.Int, error) {
	return _ExampleRewardCalculator.Contract.CalculateReward(&_ExampleRewardCalculator.CallOpts, stakeAmount, arg1, stakingStartTime, stakingEndTime, arg4)
}

// RewardBasisPoints is a free data retrieval call binding the contract method 0xbb65b242.
//
// Solidity: function rewardBasisPoints() view returns(uint64)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCaller) RewardBasisPoints(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ExampleRewardCalculator.contract.Call(opts, &out, "rewardBasisPoints")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RewardBasisPoints is a free data retrieval call binding the contract method 0xbb65b242.
//
// Solidity: function rewardBasisPoints() view returns(uint64)
func (_ExampleRewardCalculator *ExampleRewardCalculatorSession) RewardBasisPoints() (uint64, error) {
	return _ExampleRewardCalculator.Contract.RewardBasisPoints(&_ExampleRewardCalculator.CallOpts)
}

// RewardBasisPoints is a free data retrieval call binding the contract method 0xbb65b242.
//
// Solidity: function rewardBasisPoints() view returns(uint64)
func (_ExampleRewardCalculator *ExampleRewardCalculatorCallerSession) RewardBasisPoints() (uint64, error) {
	return _ExampleRewardCalculator.Contract.RewardBasisPoints(&_ExampleRewardCalculator.CallOpts)
}
