// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teleportertokensource

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

// SendAndCallInput is an auto generated low-level Go binding around an user-defined struct.
type SendAndCallInput struct {
	DestinationBlockchainID  [32]byte
	DestinationBridgeAddress common.Address
	RecipientContract        common.Address
	RecipientPayload         []byte
	RequiredGasLimit         *big.Int
	RecipientGasLimit        *big.Int
	MultiHopFallback         common.Address
	FallbackRecipient        common.Address
	PrimaryFeeTokenAddress   common.Address
	PrimaryFee               *big.Int
	SecondaryFee             *big.Int
}

// SendTokensInput is an auto generated low-level Go binding around an user-defined struct.
type SendTokensInput struct {
	DestinationBlockchainID  [32]byte
	DestinationBridgeAddress common.Address
	Recipient                common.Address
	PrimaryFeeTokenAddress   common.Address
	PrimaryFee               *big.Int
	SecondaryFee             *big.Int
	RequiredGasLimit         *big.Int
	MultiHopFallback         common.Address
}

// TeleporterTokenSourceMetaData contains all meta data concerning the TeleporterTokenSource contract.
var TeleporterTokenSourceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"name\":\"CollateralAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialCollateralNeeded\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenMultiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"multiplyOnDestination\",\"type\":\"bool\"}],\"name\":\"DestinationRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallRouted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensRouted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_TOKEN_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"}],\"name\":\"bridgedBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"}],\"name\":\"registeredDestinations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralNeeded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"multiplyOnDestination\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeleporterTokenSourceABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleporterTokenSourceMetaData.ABI instead.
var TeleporterTokenSourceABI = TeleporterTokenSourceMetaData.ABI

// TeleporterTokenSource is an auto generated Go binding around an Ethereum contract.
type TeleporterTokenSource struct {
	TeleporterTokenSourceCaller     // Read-only binding to the contract
	TeleporterTokenSourceTransactor // Write-only binding to the contract
	TeleporterTokenSourceFilterer   // Log filterer for contract events
}

// TeleporterTokenSourceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeleporterTokenSourceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterTokenSourceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeleporterTokenSourceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterTokenSourceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeleporterTokenSourceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterTokenSourceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeleporterTokenSourceSession struct {
	Contract     *TeleporterTokenSource // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TeleporterTokenSourceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeleporterTokenSourceCallerSession struct {
	Contract *TeleporterTokenSourceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// TeleporterTokenSourceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeleporterTokenSourceTransactorSession struct {
	Contract     *TeleporterTokenSourceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// TeleporterTokenSourceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeleporterTokenSourceRaw struct {
	Contract *TeleporterTokenSource // Generic contract binding to access the raw methods on
}

// TeleporterTokenSourceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeleporterTokenSourceCallerRaw struct {
	Contract *TeleporterTokenSourceCaller // Generic read-only contract binding to access the raw methods on
}

// TeleporterTokenSourceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeleporterTokenSourceTransactorRaw struct {
	Contract *TeleporterTokenSourceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeleporterTokenSource creates a new instance of TeleporterTokenSource, bound to a specific deployed contract.
func NewTeleporterTokenSource(address common.Address, backend bind.ContractBackend) (*TeleporterTokenSource, error) {
	contract, err := bindTeleporterTokenSource(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSource{TeleporterTokenSourceCaller: TeleporterTokenSourceCaller{contract: contract}, TeleporterTokenSourceTransactor: TeleporterTokenSourceTransactor{contract: contract}, TeleporterTokenSourceFilterer: TeleporterTokenSourceFilterer{contract: contract}}, nil
}

// NewTeleporterTokenSourceCaller creates a new read-only instance of TeleporterTokenSource, bound to a specific deployed contract.
func NewTeleporterTokenSourceCaller(address common.Address, caller bind.ContractCaller) (*TeleporterTokenSourceCaller, error) {
	contract, err := bindTeleporterTokenSource(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSourceCaller{contract: contract}, nil
}

// NewTeleporterTokenSourceTransactor creates a new write-only instance of TeleporterTokenSource, bound to a specific deployed contract.
func NewTeleporterTokenSourceTransactor(address common.Address, transactor bind.ContractTransactor) (*TeleporterTokenSourceTransactor, error) {
	contract, err := bindTeleporterTokenSource(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSourceTransactor{contract: contract}, nil
}

// NewTeleporterTokenSourceFilterer creates a new log filterer instance of TeleporterTokenSource, bound to a specific deployed contract.
func NewTeleporterTokenSourceFilterer(address common.Address, filterer bind.ContractFilterer) (*TeleporterTokenSourceFilterer, error) {
	contract, err := bindTeleporterTokenSource(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSourceFilterer{contract: contract}, nil
}

// bindTeleporterTokenSource binds a generic wrapper to an already deployed contract.
func bindTeleporterTokenSource(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeleporterTokenSourceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterTokenSource *TeleporterTokenSourceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterTokenSource.Contract.TeleporterTokenSourceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterTokenSource *TeleporterTokenSourceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterTokenSource.Contract.TeleporterTokenSourceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterTokenSource *TeleporterTokenSourceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterTokenSource.Contract.TeleporterTokenSourceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterTokenSource *TeleporterTokenSourceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterTokenSource.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterTokenSource *TeleporterTokenSourceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterTokenSource.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterTokenSource *TeleporterTokenSourceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterTokenSource.Contract.contract.Transact(opts, method, params...)
}

// MAXTOKENMULTIPLIER is a free data retrieval call binding the contract method 0x008ad777.
//
// Solidity: function MAX_TOKEN_MULTIPLIER() view returns(uint256)
func (_TeleporterTokenSource *TeleporterTokenSourceCaller) MAXTOKENMULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterTokenSource.contract.Call(opts, &out, "MAX_TOKEN_MULTIPLIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOKENMULTIPLIER is a free data retrieval call binding the contract method 0x008ad777.
//
// Solidity: function MAX_TOKEN_MULTIPLIER() view returns(uint256)
func (_TeleporterTokenSource *TeleporterTokenSourceSession) MAXTOKENMULTIPLIER() (*big.Int, error) {
	return _TeleporterTokenSource.Contract.MAXTOKENMULTIPLIER(&_TeleporterTokenSource.CallOpts)
}

// MAXTOKENMULTIPLIER is a free data retrieval call binding the contract method 0x008ad777.
//
// Solidity: function MAX_TOKEN_MULTIPLIER() view returns(uint256)
func (_TeleporterTokenSource *TeleporterTokenSourceCallerSession) MAXTOKENMULTIPLIER() (*big.Int, error) {
	return _TeleporterTokenSource.Contract.MAXTOKENMULTIPLIER(&_TeleporterTokenSource.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterTokenSource *TeleporterTokenSourceCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterTokenSource.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterTokenSource *TeleporterTokenSourceSession) BlockchainID() ([32]byte, error) {
	return _TeleporterTokenSource.Contract.BlockchainID(&_TeleporterTokenSource.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterTokenSource *TeleporterTokenSourceCallerSession) BlockchainID() ([32]byte, error) {
	return _TeleporterTokenSource.Contract.BlockchainID(&_TeleporterTokenSource.CallOpts)
}

// BridgedBalances is a free data retrieval call binding the contract method 0x02ee3e9c.
//
// Solidity: function bridgedBalances(bytes32 destinationBlockchainID, address destinationBridgeAddress) view returns(uint256 balance)
func (_TeleporterTokenSource *TeleporterTokenSourceCaller) BridgedBalances(opts *bind.CallOpts, destinationBlockchainID [32]byte, destinationBridgeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterTokenSource.contract.Call(opts, &out, "bridgedBalances", destinationBlockchainID, destinationBridgeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgedBalances is a free data retrieval call binding the contract method 0x02ee3e9c.
//
// Solidity: function bridgedBalances(bytes32 destinationBlockchainID, address destinationBridgeAddress) view returns(uint256 balance)
func (_TeleporterTokenSource *TeleporterTokenSourceSession) BridgedBalances(destinationBlockchainID [32]byte, destinationBridgeAddress common.Address) (*big.Int, error) {
	return _TeleporterTokenSource.Contract.BridgedBalances(&_TeleporterTokenSource.CallOpts, destinationBlockchainID, destinationBridgeAddress)
}

// BridgedBalances is a free data retrieval call binding the contract method 0x02ee3e9c.
//
// Solidity: function bridgedBalances(bytes32 destinationBlockchainID, address destinationBridgeAddress) view returns(uint256 balance)
func (_TeleporterTokenSource *TeleporterTokenSourceCallerSession) BridgedBalances(destinationBlockchainID [32]byte, destinationBridgeAddress common.Address) (*big.Int, error) {
	return _TeleporterTokenSource.Contract.BridgedBalances(&_TeleporterTokenSource.CallOpts, destinationBlockchainID, destinationBridgeAddress)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_TeleporterTokenSource *TeleporterTokenSourceCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterTokenSource.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_TeleporterTokenSource *TeleporterTokenSourceSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _TeleporterTokenSource.Contract.GetMinTeleporterVersion(&_TeleporterTokenSource.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_TeleporterTokenSource *TeleporterTokenSourceCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _TeleporterTokenSource.Contract.GetMinTeleporterVersion(&_TeleporterTokenSource.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_TeleporterTokenSource *TeleporterTokenSourceCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _TeleporterTokenSource.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_TeleporterTokenSource *TeleporterTokenSourceSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _TeleporterTokenSource.Contract.IsTeleporterAddressPaused(&_TeleporterTokenSource.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_TeleporterTokenSource *TeleporterTokenSourceCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _TeleporterTokenSource.Contract.IsTeleporterAddressPaused(&_TeleporterTokenSource.CallOpts, teleporterAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TeleporterTokenSource *TeleporterTokenSourceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterTokenSource.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TeleporterTokenSource *TeleporterTokenSourceSession) Owner() (common.Address, error) {
	return _TeleporterTokenSource.Contract.Owner(&_TeleporterTokenSource.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TeleporterTokenSource *TeleporterTokenSourceCallerSession) Owner() (common.Address, error) {
	return _TeleporterTokenSource.Contract.Owner(&_TeleporterTokenSource.CallOpts)
}

// RegisteredDestinations is a free data retrieval call binding the contract method 0xfcfb5ff9.
//
// Solidity: function registeredDestinations(bytes32 destinationBlockchainID, address destinationBridgeAddress) view returns(bool registered, uint256 collateralNeeded, uint256 tokenMultiplier, bool multiplyOnDestination)
func (_TeleporterTokenSource *TeleporterTokenSourceCaller) RegisteredDestinations(opts *bind.CallOpts, destinationBlockchainID [32]byte, destinationBridgeAddress common.Address) (struct {
	Registered            bool
	CollateralNeeded      *big.Int
	TokenMultiplier       *big.Int
	MultiplyOnDestination bool
}, error) {
	var out []interface{}
	err := _TeleporterTokenSource.contract.Call(opts, &out, "registeredDestinations", destinationBlockchainID, destinationBridgeAddress)

	outstruct := new(struct {
		Registered            bool
		CollateralNeeded      *big.Int
		TokenMultiplier       *big.Int
		MultiplyOnDestination bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Registered = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CollateralNeeded = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TokenMultiplier = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MultiplyOnDestination = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// RegisteredDestinations is a free data retrieval call binding the contract method 0xfcfb5ff9.
//
// Solidity: function registeredDestinations(bytes32 destinationBlockchainID, address destinationBridgeAddress) view returns(bool registered, uint256 collateralNeeded, uint256 tokenMultiplier, bool multiplyOnDestination)
func (_TeleporterTokenSource *TeleporterTokenSourceSession) RegisteredDestinations(destinationBlockchainID [32]byte, destinationBridgeAddress common.Address) (struct {
	Registered            bool
	CollateralNeeded      *big.Int
	TokenMultiplier       *big.Int
	MultiplyOnDestination bool
}, error) {
	return _TeleporterTokenSource.Contract.RegisteredDestinations(&_TeleporterTokenSource.CallOpts, destinationBlockchainID, destinationBridgeAddress)
}

// RegisteredDestinations is a free data retrieval call binding the contract method 0xfcfb5ff9.
//
// Solidity: function registeredDestinations(bytes32 destinationBlockchainID, address destinationBridgeAddress) view returns(bool registered, uint256 collateralNeeded, uint256 tokenMultiplier, bool multiplyOnDestination)
func (_TeleporterTokenSource *TeleporterTokenSourceCallerSession) RegisteredDestinations(destinationBlockchainID [32]byte, destinationBridgeAddress common.Address) (struct {
	Registered            bool
	CollateralNeeded      *big.Int
	TokenMultiplier       *big.Int
	MultiplyOnDestination bool
}, error) {
	return _TeleporterTokenSource.Contract.RegisteredDestinations(&_TeleporterTokenSource.CallOpts, destinationBlockchainID, destinationBridgeAddress)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_TeleporterTokenSource *TeleporterTokenSourceCaller) TeleporterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterTokenSource.contract.Call(opts, &out, "teleporterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_TeleporterTokenSource *TeleporterTokenSourceSession) TeleporterRegistry() (common.Address, error) {
	return _TeleporterTokenSource.Contract.TeleporterRegistry(&_TeleporterTokenSource.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_TeleporterTokenSource *TeleporterTokenSourceCallerSession) TeleporterRegistry() (common.Address, error) {
	return _TeleporterTokenSource.Contract.TeleporterRegistry(&_TeleporterTokenSource.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_TeleporterTokenSource *TeleporterTokenSourceCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterTokenSource.contract.Call(opts, &out, "tokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_TeleporterTokenSource *TeleporterTokenSourceSession) TokenAddress() (common.Address, error) {
	return _TeleporterTokenSource.Contract.TokenAddress(&_TeleporterTokenSource.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_TeleporterTokenSource *TeleporterTokenSourceCallerSession) TokenAddress() (common.Address, error) {
	return _TeleporterTokenSource.Contract.TokenAddress(&_TeleporterTokenSource.CallOpts)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_TeleporterTokenSource *TeleporterTokenSourceTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _TeleporterTokenSource.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_TeleporterTokenSource *TeleporterTokenSourceSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _TeleporterTokenSource.Contract.PauseTeleporterAddress(&_TeleporterTokenSource.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_TeleporterTokenSource *TeleporterTokenSourceTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _TeleporterTokenSource.Contract.PauseTeleporterAddress(&_TeleporterTokenSource.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_TeleporterTokenSource *TeleporterTokenSourceTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _TeleporterTokenSource.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_TeleporterTokenSource *TeleporterTokenSourceSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _TeleporterTokenSource.Contract.ReceiveTeleporterMessage(&_TeleporterTokenSource.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_TeleporterTokenSource *TeleporterTokenSourceTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _TeleporterTokenSource.Contract.ReceiveTeleporterMessage(&_TeleporterTokenSource.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TeleporterTokenSource *TeleporterTokenSourceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterTokenSource.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TeleporterTokenSource *TeleporterTokenSourceSession) RenounceOwnership() (*types.Transaction, error) {
	return _TeleporterTokenSource.Contract.RenounceOwnership(&_TeleporterTokenSource.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TeleporterTokenSource *TeleporterTokenSourceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TeleporterTokenSource.Contract.RenounceOwnership(&_TeleporterTokenSource.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TeleporterTokenSource *TeleporterTokenSourceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TeleporterTokenSource.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TeleporterTokenSource *TeleporterTokenSourceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TeleporterTokenSource.Contract.TransferOwnership(&_TeleporterTokenSource.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TeleporterTokenSource *TeleporterTokenSourceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TeleporterTokenSource.Contract.TransferOwnership(&_TeleporterTokenSource.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_TeleporterTokenSource *TeleporterTokenSourceTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _TeleporterTokenSource.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_TeleporterTokenSource *TeleporterTokenSourceSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _TeleporterTokenSource.Contract.UnpauseTeleporterAddress(&_TeleporterTokenSource.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_TeleporterTokenSource *TeleporterTokenSourceTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _TeleporterTokenSource.Contract.UnpauseTeleporterAddress(&_TeleporterTokenSource.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_TeleporterTokenSource *TeleporterTokenSourceTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _TeleporterTokenSource.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_TeleporterTokenSource *TeleporterTokenSourceSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _TeleporterTokenSource.Contract.UpdateMinTeleporterVersion(&_TeleporterTokenSource.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_TeleporterTokenSource *TeleporterTokenSourceTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _TeleporterTokenSource.Contract.UpdateMinTeleporterVersion(&_TeleporterTokenSource.TransactOpts, version)
}

// TeleporterTokenSourceCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceCallFailedIterator struct {
	Event *TeleporterTokenSourceCallFailed // Event containing the contract specifics and raw log

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
func (it *TeleporterTokenSourceCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterTokenSourceCallFailed)
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
		it.Event = new(TeleporterTokenSourceCallFailed)
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
func (it *TeleporterTokenSourceCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterTokenSourceCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterTokenSourceCallFailed represents a CallFailed event raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceCallFailed struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) FilterCallFailed(opts *bind.FilterOpts, recipientContract []common.Address) (*TeleporterTokenSourceCallFailedIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.FilterLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSourceCallFailedIterator{contract: _TeleporterTokenSource.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *TeleporterTokenSourceCallFailed, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.WatchLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterTokenSourceCallFailed)
				if err := _TeleporterTokenSource.contract.UnpackLog(event, "CallFailed", log); err != nil {
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

// ParseCallFailed is a log parse operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) ParseCallFailed(log types.Log) (*TeleporterTokenSourceCallFailed, error) {
	event := new(TeleporterTokenSourceCallFailed)
	if err := _TeleporterTokenSource.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterTokenSourceCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceCallSucceededIterator struct {
	Event *TeleporterTokenSourceCallSucceeded // Event containing the contract specifics and raw log

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
func (it *TeleporterTokenSourceCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterTokenSourceCallSucceeded)
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
		it.Event = new(TeleporterTokenSourceCallSucceeded)
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
func (it *TeleporterTokenSourceCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterTokenSourceCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterTokenSourceCallSucceeded represents a CallSucceeded event raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceCallSucceeded struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) FilterCallSucceeded(opts *bind.FilterOpts, recipientContract []common.Address) (*TeleporterTokenSourceCallSucceededIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.FilterLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSourceCallSucceededIterator{contract: _TeleporterTokenSource.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *TeleporterTokenSourceCallSucceeded, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.WatchLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterTokenSourceCallSucceeded)
				if err := _TeleporterTokenSource.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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

// ParseCallSucceeded is a log parse operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) ParseCallSucceeded(log types.Log) (*TeleporterTokenSourceCallSucceeded, error) {
	event := new(TeleporterTokenSourceCallSucceeded)
	if err := _TeleporterTokenSource.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterTokenSourceCollateralAddedIterator is returned from FilterCollateralAdded and is used to iterate over the raw logs and unpacked data for CollateralAdded events raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceCollateralAddedIterator struct {
	Event *TeleporterTokenSourceCollateralAdded // Event containing the contract specifics and raw log

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
func (it *TeleporterTokenSourceCollateralAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterTokenSourceCollateralAdded)
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
		it.Event = new(TeleporterTokenSourceCollateralAdded)
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
func (it *TeleporterTokenSourceCollateralAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterTokenSourceCollateralAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterTokenSourceCollateralAdded represents a CollateralAdded event raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceCollateralAdded struct {
	DestinationBlockchainID  [32]byte
	DestinationBridgeAddress common.Address
	Amount                   *big.Int
	Remaining                *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterCollateralAdded is a free log retrieval operation binding the contract event 0x6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6.
//
// Solidity: event CollateralAdded(bytes32 indexed destinationBlockchainID, address indexed destinationBridgeAddress, uint256 amount, uint256 remaining)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) FilterCollateralAdded(opts *bind.FilterOpts, destinationBlockchainID [][32]byte, destinationBridgeAddress []common.Address) (*TeleporterTokenSourceCollateralAddedIterator, error) {

	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var destinationBridgeAddressRule []interface{}
	for _, destinationBridgeAddressItem := range destinationBridgeAddress {
		destinationBridgeAddressRule = append(destinationBridgeAddressRule, destinationBridgeAddressItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.FilterLogs(opts, "CollateralAdded", destinationBlockchainIDRule, destinationBridgeAddressRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSourceCollateralAddedIterator{contract: _TeleporterTokenSource.contract, event: "CollateralAdded", logs: logs, sub: sub}, nil
}

// WatchCollateralAdded is a free log subscription operation binding the contract event 0x6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6.
//
// Solidity: event CollateralAdded(bytes32 indexed destinationBlockchainID, address indexed destinationBridgeAddress, uint256 amount, uint256 remaining)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) WatchCollateralAdded(opts *bind.WatchOpts, sink chan<- *TeleporterTokenSourceCollateralAdded, destinationBlockchainID [][32]byte, destinationBridgeAddress []common.Address) (event.Subscription, error) {

	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var destinationBridgeAddressRule []interface{}
	for _, destinationBridgeAddressItem := range destinationBridgeAddress {
		destinationBridgeAddressRule = append(destinationBridgeAddressRule, destinationBridgeAddressItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.WatchLogs(opts, "CollateralAdded", destinationBlockchainIDRule, destinationBridgeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterTokenSourceCollateralAdded)
				if err := _TeleporterTokenSource.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
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

// ParseCollateralAdded is a log parse operation binding the contract event 0x6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6.
//
// Solidity: event CollateralAdded(bytes32 indexed destinationBlockchainID, address indexed destinationBridgeAddress, uint256 amount, uint256 remaining)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) ParseCollateralAdded(log types.Log) (*TeleporterTokenSourceCollateralAdded, error) {
	event := new(TeleporterTokenSourceCollateralAdded)
	if err := _TeleporterTokenSource.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterTokenSourceDestinationRegisteredIterator is returned from FilterDestinationRegistered and is used to iterate over the raw logs and unpacked data for DestinationRegistered events raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceDestinationRegisteredIterator struct {
	Event *TeleporterTokenSourceDestinationRegistered // Event containing the contract specifics and raw log

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
func (it *TeleporterTokenSourceDestinationRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterTokenSourceDestinationRegistered)
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
		it.Event = new(TeleporterTokenSourceDestinationRegistered)
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
func (it *TeleporterTokenSourceDestinationRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterTokenSourceDestinationRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterTokenSourceDestinationRegistered represents a DestinationRegistered event raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceDestinationRegistered struct {
	DestinationBlockchainID  [32]byte
	DestinationBridgeAddress common.Address
	InitialCollateralNeeded  *big.Int
	TokenMultiplier          *big.Int
	MultiplyOnDestination    bool
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterDestinationRegistered is a free log retrieval operation binding the contract event 0x34e0f289c6952b50ea58fa31d8c3732fb85091cfc3d5ded7767d7562ce0bfb22.
//
// Solidity: event DestinationRegistered(bytes32 indexed destinationBlockchainID, address indexed destinationBridgeAddress, uint256 initialCollateralNeeded, uint256 tokenMultiplier, bool multiplyOnDestination)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) FilterDestinationRegistered(opts *bind.FilterOpts, destinationBlockchainID [][32]byte, destinationBridgeAddress []common.Address) (*TeleporterTokenSourceDestinationRegisteredIterator, error) {

	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var destinationBridgeAddressRule []interface{}
	for _, destinationBridgeAddressItem := range destinationBridgeAddress {
		destinationBridgeAddressRule = append(destinationBridgeAddressRule, destinationBridgeAddressItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.FilterLogs(opts, "DestinationRegistered", destinationBlockchainIDRule, destinationBridgeAddressRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSourceDestinationRegisteredIterator{contract: _TeleporterTokenSource.contract, event: "DestinationRegistered", logs: logs, sub: sub}, nil
}

// WatchDestinationRegistered is a free log subscription operation binding the contract event 0x34e0f289c6952b50ea58fa31d8c3732fb85091cfc3d5ded7767d7562ce0bfb22.
//
// Solidity: event DestinationRegistered(bytes32 indexed destinationBlockchainID, address indexed destinationBridgeAddress, uint256 initialCollateralNeeded, uint256 tokenMultiplier, bool multiplyOnDestination)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) WatchDestinationRegistered(opts *bind.WatchOpts, sink chan<- *TeleporterTokenSourceDestinationRegistered, destinationBlockchainID [][32]byte, destinationBridgeAddress []common.Address) (event.Subscription, error) {

	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var destinationBridgeAddressRule []interface{}
	for _, destinationBridgeAddressItem := range destinationBridgeAddress {
		destinationBridgeAddressRule = append(destinationBridgeAddressRule, destinationBridgeAddressItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.WatchLogs(opts, "DestinationRegistered", destinationBlockchainIDRule, destinationBridgeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterTokenSourceDestinationRegistered)
				if err := _TeleporterTokenSource.contract.UnpackLog(event, "DestinationRegistered", log); err != nil {
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

// ParseDestinationRegistered is a log parse operation binding the contract event 0x34e0f289c6952b50ea58fa31d8c3732fb85091cfc3d5ded7767d7562ce0bfb22.
//
// Solidity: event DestinationRegistered(bytes32 indexed destinationBlockchainID, address indexed destinationBridgeAddress, uint256 initialCollateralNeeded, uint256 tokenMultiplier, bool multiplyOnDestination)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) ParseDestinationRegistered(log types.Log) (*TeleporterTokenSourceDestinationRegistered, error) {
	event := new(TeleporterTokenSourceDestinationRegistered)
	if err := _TeleporterTokenSource.contract.UnpackLog(event, "DestinationRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterTokenSourceMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceMinTeleporterVersionUpdatedIterator struct {
	Event *TeleporterTokenSourceMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *TeleporterTokenSourceMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterTokenSourceMinTeleporterVersionUpdated)
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
		it.Event = new(TeleporterTokenSourceMinTeleporterVersionUpdated)
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
func (it *TeleporterTokenSourceMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterTokenSourceMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterTokenSourceMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*TeleporterTokenSourceMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSourceMinTeleporterVersionUpdatedIterator{contract: _TeleporterTokenSource.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *TeleporterTokenSourceMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterTokenSourceMinTeleporterVersionUpdated)
				if err := _TeleporterTokenSource.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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

// ParseMinTeleporterVersionUpdated is a log parse operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*TeleporterTokenSourceMinTeleporterVersionUpdated, error) {
	event := new(TeleporterTokenSourceMinTeleporterVersionUpdated)
	if err := _TeleporterTokenSource.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterTokenSourceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceOwnershipTransferredIterator struct {
	Event *TeleporterTokenSourceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TeleporterTokenSourceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterTokenSourceOwnershipTransferred)
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
		it.Event = new(TeleporterTokenSourceOwnershipTransferred)
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
func (it *TeleporterTokenSourceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterTokenSourceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterTokenSourceOwnershipTransferred represents a OwnershipTransferred event raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TeleporterTokenSourceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSourceOwnershipTransferredIterator{contract: _TeleporterTokenSource.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TeleporterTokenSourceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterTokenSourceOwnershipTransferred)
				if err := _TeleporterTokenSource.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) ParseOwnershipTransferred(log types.Log) (*TeleporterTokenSourceOwnershipTransferred, error) {
	event := new(TeleporterTokenSourceOwnershipTransferred)
	if err := _TeleporterTokenSource.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterTokenSourceTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceTeleporterAddressPausedIterator struct {
	Event *TeleporterTokenSourceTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *TeleporterTokenSourceTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterTokenSourceTeleporterAddressPaused)
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
		it.Event = new(TeleporterTokenSourceTeleporterAddressPaused)
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
func (it *TeleporterTokenSourceTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterTokenSourceTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterTokenSourceTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*TeleporterTokenSourceTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSourceTeleporterAddressPausedIterator{contract: _TeleporterTokenSource.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *TeleporterTokenSourceTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterTokenSourceTeleporterAddressPaused)
				if err := _TeleporterTokenSource.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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

// ParseTeleporterAddressPaused is a log parse operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) ParseTeleporterAddressPaused(log types.Log) (*TeleporterTokenSourceTeleporterAddressPaused, error) {
	event := new(TeleporterTokenSourceTeleporterAddressPaused)
	if err := _TeleporterTokenSource.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterTokenSourceTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceTeleporterAddressUnpausedIterator struct {
	Event *TeleporterTokenSourceTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *TeleporterTokenSourceTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterTokenSourceTeleporterAddressUnpaused)
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
		it.Event = new(TeleporterTokenSourceTeleporterAddressUnpaused)
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
func (it *TeleporterTokenSourceTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterTokenSourceTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterTokenSourceTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*TeleporterTokenSourceTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSourceTeleporterAddressUnpausedIterator{contract: _TeleporterTokenSource.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *TeleporterTokenSourceTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterTokenSourceTeleporterAddressUnpaused)
				if err := _TeleporterTokenSource.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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

// ParseTeleporterAddressUnpaused is a log parse operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*TeleporterTokenSourceTeleporterAddressUnpaused, error) {
	event := new(TeleporterTokenSourceTeleporterAddressUnpaused)
	if err := _TeleporterTokenSource.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterTokenSourceTokensAndCallRoutedIterator is returned from FilterTokensAndCallRouted and is used to iterate over the raw logs and unpacked data for TokensAndCallRouted events raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceTokensAndCallRoutedIterator struct {
	Event *TeleporterTokenSourceTokensAndCallRouted // Event containing the contract specifics and raw log

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
func (it *TeleporterTokenSourceTokensAndCallRoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterTokenSourceTokensAndCallRouted)
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
		it.Event = new(TeleporterTokenSourceTokensAndCallRouted)
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
func (it *TeleporterTokenSourceTokensAndCallRoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterTokenSourceTokensAndCallRoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterTokenSourceTokensAndCallRouted represents a TokensAndCallRouted event raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceTokensAndCallRouted struct {
	TeleporterMessageID [32]byte
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallRouted is a free log retrieval operation binding the contract event 0x42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb30.
//
// Solidity: event TokensAndCallRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) FilterTokensAndCallRouted(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*TeleporterTokenSourceTokensAndCallRoutedIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.FilterLogs(opts, "TokensAndCallRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSourceTokensAndCallRoutedIterator{contract: _TeleporterTokenSource.contract, event: "TokensAndCallRouted", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallRouted is a free log subscription operation binding the contract event 0x42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb30.
//
// Solidity: event TokensAndCallRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) WatchTokensAndCallRouted(opts *bind.WatchOpts, sink chan<- *TeleporterTokenSourceTokensAndCallRouted, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.WatchLogs(opts, "TokensAndCallRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterTokenSourceTokensAndCallRouted)
				if err := _TeleporterTokenSource.contract.UnpackLog(event, "TokensAndCallRouted", log); err != nil {
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

// ParseTokensAndCallRouted is a log parse operation binding the contract event 0x42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb30.
//
// Solidity: event TokensAndCallRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) ParseTokensAndCallRouted(log types.Log) (*TeleporterTokenSourceTokensAndCallRouted, error) {
	event := new(TeleporterTokenSourceTokensAndCallRouted)
	if err := _TeleporterTokenSource.contract.UnpackLog(event, "TokensAndCallRouted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterTokenSourceTokensAndCallSentIterator is returned from FilterTokensAndCallSent and is used to iterate over the raw logs and unpacked data for TokensAndCallSent events raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceTokensAndCallSentIterator struct {
	Event *TeleporterTokenSourceTokensAndCallSent // Event containing the contract specifics and raw log

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
func (it *TeleporterTokenSourceTokensAndCallSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterTokenSourceTokensAndCallSent)
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
		it.Event = new(TeleporterTokenSourceTokensAndCallSent)
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
func (it *TeleporterTokenSourceTokensAndCallSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterTokenSourceTokensAndCallSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterTokenSourceTokensAndCallSent represents a TokensAndCallSent event raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceTokensAndCallSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) FilterTokensAndCallSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*TeleporterTokenSourceTokensAndCallSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.FilterLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSourceTokensAndCallSentIterator{contract: _TeleporterTokenSource.contract, event: "TokensAndCallSent", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) WatchTokensAndCallSent(opts *bind.WatchOpts, sink chan<- *TeleporterTokenSourceTokensAndCallSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.WatchLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterTokenSourceTokensAndCallSent)
				if err := _TeleporterTokenSource.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
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

// ParseTokensAndCallSent is a log parse operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) ParseTokensAndCallSent(log types.Log) (*TeleporterTokenSourceTokensAndCallSent, error) {
	event := new(TeleporterTokenSourceTokensAndCallSent)
	if err := _TeleporterTokenSource.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterTokenSourceTokensRoutedIterator is returned from FilterTokensRouted and is used to iterate over the raw logs and unpacked data for TokensRouted events raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceTokensRoutedIterator struct {
	Event *TeleporterTokenSourceTokensRouted // Event containing the contract specifics and raw log

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
func (it *TeleporterTokenSourceTokensRoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterTokenSourceTokensRouted)
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
		it.Event = new(TeleporterTokenSourceTokensRouted)
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
func (it *TeleporterTokenSourceTokensRoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterTokenSourceTokensRoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterTokenSourceTokensRouted represents a TokensRouted event raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceTokensRouted struct {
	TeleporterMessageID [32]byte
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensRouted is a free log retrieval operation binding the contract event 0x825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf0.
//
// Solidity: event TokensRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) FilterTokensRouted(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*TeleporterTokenSourceTokensRoutedIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.FilterLogs(opts, "TokensRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSourceTokensRoutedIterator{contract: _TeleporterTokenSource.contract, event: "TokensRouted", logs: logs, sub: sub}, nil
}

// WatchTokensRouted is a free log subscription operation binding the contract event 0x825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf0.
//
// Solidity: event TokensRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) WatchTokensRouted(opts *bind.WatchOpts, sink chan<- *TeleporterTokenSourceTokensRouted, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.WatchLogs(opts, "TokensRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterTokenSourceTokensRouted)
				if err := _TeleporterTokenSource.contract.UnpackLog(event, "TokensRouted", log); err != nil {
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

// ParseTokensRouted is a log parse operation binding the contract event 0x825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf0.
//
// Solidity: event TokensRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) ParseTokensRouted(log types.Log) (*TeleporterTokenSourceTokensRouted, error) {
	event := new(TeleporterTokenSourceTokensRouted)
	if err := _TeleporterTokenSource.contract.UnpackLog(event, "TokensRouted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterTokenSourceTokensSentIterator is returned from FilterTokensSent and is used to iterate over the raw logs and unpacked data for TokensSent events raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceTokensSentIterator struct {
	Event *TeleporterTokenSourceTokensSent // Event containing the contract specifics and raw log

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
func (it *TeleporterTokenSourceTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterTokenSourceTokensSent)
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
		it.Event = new(TeleporterTokenSourceTokensSent)
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
func (it *TeleporterTokenSourceTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterTokenSourceTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterTokenSourceTokensSent represents a TokensSent event raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensSent is a free log retrieval operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) FilterTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*TeleporterTokenSourceTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.FilterLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSourceTokensSentIterator{contract: _TeleporterTokenSource.contract, event: "TokensSent", logs: logs, sub: sub}, nil
}

// WatchTokensSent is a free log subscription operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) WatchTokensSent(opts *bind.WatchOpts, sink chan<- *TeleporterTokenSourceTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.WatchLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterTokenSourceTokensSent)
				if err := _TeleporterTokenSource.contract.UnpackLog(event, "TokensSent", log); err != nil {
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

// ParseTokensSent is a log parse operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) ParseTokensSent(log types.Log) (*TeleporterTokenSourceTokensSent, error) {
	event := new(TeleporterTokenSourceTokensSent)
	if err := _TeleporterTokenSource.contract.UnpackLog(event, "TokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterTokenSourceTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceTokensWithdrawnIterator struct {
	Event *TeleporterTokenSourceTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *TeleporterTokenSourceTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterTokenSourceTokensWithdrawn)
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
		it.Event = new(TeleporterTokenSourceTokensWithdrawn)
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
func (it *TeleporterTokenSourceTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterTokenSourceTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterTokenSourceTokensWithdrawn represents a TokensWithdrawn event raised by the TeleporterTokenSource contract.
type TeleporterTokenSourceTokensWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*TeleporterTokenSourceTokensWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.FilterLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterTokenSourceTokensWithdrawnIterator{contract: _TeleporterTokenSource.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *TeleporterTokenSourceTokensWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TeleporterTokenSource.contract.WatchLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterTokenSourceTokensWithdrawn)
				if err := _TeleporterTokenSource.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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

// ParseTokensWithdrawn is a log parse operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_TeleporterTokenSource *TeleporterTokenSourceFilterer) ParseTokensWithdrawn(log types.Log) (*TeleporterTokenSourceTokensWithdrawn, error) {
	event := new(TeleporterTokenSourceTokensWithdrawn)
	if err := _TeleporterTokenSource.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
