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
	Version         *big.Int
	ProtocolAddress common.Address
}

// TeleporterFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type TeleporterFeeInfo struct {
	FeeTokenAddress common.Address
	Amount          *big.Int
}

// TeleporterMessage is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessage struct {
	MessageNonce            *big.Int
	OriginSenderAddress     common.Address
	DestinationBlockchainID [32]byte
	DestinationAddress      common.Address
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Receipts                []TeleporterMessageReceipt
	Message                 []byte
}

// TeleporterMessageInput is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageInput struct {
	DestinationBlockchainID [32]byte
	DestinationAddress      common.Address
	FeeInfo                 TeleporterFeeInfo
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Message                 []byte
}

// TeleporterMessageReceipt is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageReceipt struct {
	ReceivedMessageNonce *big.Int
	RelayerRewardAddress common.Address
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

// ITeleporterMessengerMetaData contains all meta data concerning the ITeleporterMessenger contract.
var ITeleporterMessengerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"updatedFeeInfo\",\"type\":\"tuple\"}],\"name\":\"AddFeeAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"}],\"name\":\"BlockchainIDInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"MessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MessageExecutionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"ReceiptReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deliverer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardRedeemer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ReceiveCrossChainMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RelayerRewardsRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"SendCrossChainMessage\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalFeeAmount\",\"type\":\"uint256\"}],\"name\":\"addFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"}],\"name\":\"checkRelayerRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getFeeInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getNextMessageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getReceiptAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getReceiptQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getRelayerRewardAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"messageReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"name\":\"receiveCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"}],\"name\":\"redeemRelayerRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retryMessageExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retrySendCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageInput\",\"name\":\"messageInput\",\"type\":\"tuple\"}],\"name\":\"sendCrossChainMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"messageIDs\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"sendSpecifiedReceipts\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITeleporterMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use ITeleporterMessengerMetaData.ABI instead.
var ITeleporterMessengerABI = ITeleporterMessengerMetaData.ABI

// ITeleporterMessenger is an auto generated Go binding around an Ethereum contract.
type ITeleporterMessenger struct {
	ITeleporterMessengerCaller     // Read-only binding to the contract
	ITeleporterMessengerTransactor // Write-only binding to the contract
	ITeleporterMessengerFilterer   // Log filterer for contract events
}

// ITeleporterMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITeleporterMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITeleporterMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITeleporterMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITeleporterMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITeleporterMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITeleporterMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITeleporterMessengerSession struct {
	Contract     *ITeleporterMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ITeleporterMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITeleporterMessengerCallerSession struct {
	Contract *ITeleporterMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ITeleporterMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITeleporterMessengerTransactorSession struct {
	Contract     *ITeleporterMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ITeleporterMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITeleporterMessengerRaw struct {
	Contract *ITeleporterMessenger // Generic contract binding to access the raw methods on
}

// ITeleporterMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITeleporterMessengerCallerRaw struct {
	Contract *ITeleporterMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// ITeleporterMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITeleporterMessengerTransactorRaw struct {
	Contract *ITeleporterMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITeleporterMessenger creates a new instance of ITeleporterMessenger, bound to a specific deployed contract.
func NewITeleporterMessenger(address common.Address, backend bind.ContractBackend) (*ITeleporterMessenger, error) {
	contract, err := bindITeleporterMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessenger{ITeleporterMessengerCaller: ITeleporterMessengerCaller{contract: contract}, ITeleporterMessengerTransactor: ITeleporterMessengerTransactor{contract: contract}, ITeleporterMessengerFilterer: ITeleporterMessengerFilterer{contract: contract}}, nil
}

// NewITeleporterMessengerCaller creates a new read-only instance of ITeleporterMessenger, bound to a specific deployed contract.
func NewITeleporterMessengerCaller(address common.Address, caller bind.ContractCaller) (*ITeleporterMessengerCaller, error) {
	contract, err := bindITeleporterMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerCaller{contract: contract}, nil
}

// NewITeleporterMessengerTransactor creates a new write-only instance of ITeleporterMessenger, bound to a specific deployed contract.
func NewITeleporterMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*ITeleporterMessengerTransactor, error) {
	contract, err := bindITeleporterMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerTransactor{contract: contract}, nil
}

// NewITeleporterMessengerFilterer creates a new log filterer instance of ITeleporterMessenger, bound to a specific deployed contract.
func NewITeleporterMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*ITeleporterMessengerFilterer, error) {
	contract, err := bindITeleporterMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerFilterer{contract: contract}, nil
}

// bindITeleporterMessenger binds a generic wrapper to an already deployed contract.
func bindITeleporterMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITeleporterMessenger *ITeleporterMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITeleporterMessenger.Contract.ITeleporterMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITeleporterMessenger *ITeleporterMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.ITeleporterMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITeleporterMessenger *ITeleporterMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.ITeleporterMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITeleporterMessenger *ITeleporterMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITeleporterMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITeleporterMessenger *ITeleporterMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITeleporterMessenger *ITeleporterMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.contract.Transact(opts, method, params...)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeTokenAddress) view returns(uint256)
func (_ITeleporterMessenger *ITeleporterMessengerCaller) CheckRelayerRewardAmount(opts *bind.CallOpts, relayer common.Address, feeTokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ITeleporterMessenger.contract.Call(opts, &out, "checkRelayerRewardAmount", relayer, feeTokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeTokenAddress) view returns(uint256)
func (_ITeleporterMessenger *ITeleporterMessengerSession) CheckRelayerRewardAmount(relayer common.Address, feeTokenAddress common.Address) (*big.Int, error) {
	return _ITeleporterMessenger.Contract.CheckRelayerRewardAmount(&_ITeleporterMessenger.CallOpts, relayer, feeTokenAddress)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeTokenAddress) view returns(uint256)
func (_ITeleporterMessenger *ITeleporterMessengerCallerSession) CheckRelayerRewardAmount(relayer common.Address, feeTokenAddress common.Address) (*big.Int, error) {
	return _ITeleporterMessenger.Contract.CheckRelayerRewardAmount(&_ITeleporterMessenger.CallOpts, relayer, feeTokenAddress)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_ITeleporterMessenger *ITeleporterMessengerCaller) GetFeeInfo(opts *bind.CallOpts, messageID [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ITeleporterMessenger.contract.Call(opts, &out, "getFeeInfo", messageID)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_ITeleporterMessenger *ITeleporterMessengerSession) GetFeeInfo(messageID [32]byte) (common.Address, *big.Int, error) {
	return _ITeleporterMessenger.Contract.GetFeeInfo(&_ITeleporterMessenger.CallOpts, messageID)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_ITeleporterMessenger *ITeleporterMessengerCallerSession) GetFeeInfo(messageID [32]byte) (common.Address, *big.Int, error) {
	return _ITeleporterMessenger.Contract.GetFeeInfo(&_ITeleporterMessenger.CallOpts, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerCaller) GetMessageHash(opts *bind.CallOpts, messageID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ITeleporterMessenger.contract.Call(opts, &out, "getMessageHash", messageID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerSession) GetMessageHash(messageID [32]byte) ([32]byte, error) {
	return _ITeleporterMessenger.Contract.GetMessageHash(&_ITeleporterMessenger.CallOpts, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerCallerSession) GetMessageHash(messageID [32]byte) ([32]byte, error) {
	return _ITeleporterMessenger.Contract.GetMessageHash(&_ITeleporterMessenger.CallOpts, messageID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerCaller) GetNextMessageID(opts *bind.CallOpts, destinationBlockchainID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ITeleporterMessenger.contract.Call(opts, &out, "getNextMessageID", destinationBlockchainID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerSession) GetNextMessageID(destinationBlockchainID [32]byte) ([32]byte, error) {
	return _ITeleporterMessenger.Contract.GetNextMessageID(&_ITeleporterMessenger.CallOpts, destinationBlockchainID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerCallerSession) GetNextMessageID(destinationBlockchainID [32]byte) ([32]byte, error) {
	return _ITeleporterMessenger.Contract.GetNextMessageID(&_ITeleporterMessenger.CallOpts, destinationBlockchainID)
}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_ITeleporterMessenger *ITeleporterMessengerCaller) GetReceiptAtIndex(opts *bind.CallOpts, sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	var out []interface{}
	err := _ITeleporterMessenger.contract.Call(opts, &out, "getReceiptAtIndex", sourceBlockchainID, index)

	if err != nil {
		return *new(TeleporterMessageReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(TeleporterMessageReceipt)).(*TeleporterMessageReceipt)

	return out0, err

}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_ITeleporterMessenger *ITeleporterMessengerSession) GetReceiptAtIndex(sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	return _ITeleporterMessenger.Contract.GetReceiptAtIndex(&_ITeleporterMessenger.CallOpts, sourceBlockchainID, index)
}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_ITeleporterMessenger *ITeleporterMessengerCallerSession) GetReceiptAtIndex(sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	return _ITeleporterMessenger.Contract.GetReceiptAtIndex(&_ITeleporterMessenger.CallOpts, sourceBlockchainID, index)
}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_ITeleporterMessenger *ITeleporterMessengerCaller) GetReceiptQueueSize(opts *bind.CallOpts, sourceBlockchainID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ITeleporterMessenger.contract.Call(opts, &out, "getReceiptQueueSize", sourceBlockchainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_ITeleporterMessenger *ITeleporterMessengerSession) GetReceiptQueueSize(sourceBlockchainID [32]byte) (*big.Int, error) {
	return _ITeleporterMessenger.Contract.GetReceiptQueueSize(&_ITeleporterMessenger.CallOpts, sourceBlockchainID)
}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_ITeleporterMessenger *ITeleporterMessengerCallerSession) GetReceiptQueueSize(sourceBlockchainID [32]byte) (*big.Int, error) {
	return _ITeleporterMessenger.Contract.GetReceiptQueueSize(&_ITeleporterMessenger.CallOpts, sourceBlockchainID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_ITeleporterMessenger *ITeleporterMessengerCaller) GetRelayerRewardAddress(opts *bind.CallOpts, messageID [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ITeleporterMessenger.contract.Call(opts, &out, "getRelayerRewardAddress", messageID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_ITeleporterMessenger *ITeleporterMessengerSession) GetRelayerRewardAddress(messageID [32]byte) (common.Address, error) {
	return _ITeleporterMessenger.Contract.GetRelayerRewardAddress(&_ITeleporterMessenger.CallOpts, messageID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_ITeleporterMessenger *ITeleporterMessengerCallerSession) GetRelayerRewardAddress(messageID [32]byte) (common.Address, error) {
	return _ITeleporterMessenger.Contract.GetRelayerRewardAddress(&_ITeleporterMessenger.CallOpts, messageID)
}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_ITeleporterMessenger *ITeleporterMessengerCaller) MessageReceived(opts *bind.CallOpts, messageID [32]byte) (bool, error) {
	var out []interface{}
	err := _ITeleporterMessenger.contract.Call(opts, &out, "messageReceived", messageID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_ITeleporterMessenger *ITeleporterMessengerSession) MessageReceived(messageID [32]byte) (bool, error) {
	return _ITeleporterMessenger.Contract.MessageReceived(&_ITeleporterMessenger.CallOpts, messageID)
}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_ITeleporterMessenger *ITeleporterMessengerCallerSession) MessageReceived(messageID [32]byte) (bool, error) {
	return _ITeleporterMessenger.Contract.MessageReceived(&_ITeleporterMessenger.CallOpts, messageID)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactor) AddFeeAmount(opts *bind.TransactOpts, messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _ITeleporterMessenger.contract.Transact(opts, "addFeeAmount", messageID, feeTokenAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_ITeleporterMessenger *ITeleporterMessengerSession) AddFeeAmount(messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.AddFeeAmount(&_ITeleporterMessenger.TransactOpts, messageID, feeTokenAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactorSession) AddFeeAmount(messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.AddFeeAmount(&_ITeleporterMessenger.TransactOpts, messageID, feeTokenAddress, additionalFeeAmount)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactor) ReceiveCrossChainMessage(opts *bind.TransactOpts, messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.contract.Transact(opts, "receiveCrossChainMessage", messageIndex, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_ITeleporterMessenger *ITeleporterMessengerSession) ReceiveCrossChainMessage(messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.ReceiveCrossChainMessage(&_ITeleporterMessenger.TransactOpts, messageIndex, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactorSession) ReceiveCrossChainMessage(messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.ReceiveCrossChainMessage(&_ITeleporterMessenger.TransactOpts, messageIndex, relayerRewardAddress)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeTokenAddress) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactor) RedeemRelayerRewards(opts *bind.TransactOpts, feeTokenAddress common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.contract.Transact(opts, "redeemRelayerRewards", feeTokenAddress)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeTokenAddress) returns()
func (_ITeleporterMessenger *ITeleporterMessengerSession) RedeemRelayerRewards(feeTokenAddress common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.RedeemRelayerRewards(&_ITeleporterMessenger.TransactOpts, feeTokenAddress)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeTokenAddress) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactorSession) RedeemRelayerRewards(feeTokenAddress common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.RedeemRelayerRewards(&_ITeleporterMessenger.TransactOpts, feeTokenAddress)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xfc2d6197.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactor) RetryMessageExecution(opts *bind.TransactOpts, sourceBlockchainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _ITeleporterMessenger.contract.Transact(opts, "retryMessageExecution", sourceBlockchainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xfc2d6197.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ITeleporterMessenger *ITeleporterMessengerSession) RetryMessageExecution(sourceBlockchainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.RetryMessageExecution(&_ITeleporterMessenger.TransactOpts, sourceBlockchainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xfc2d6197.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactorSession) RetryMessageExecution(sourceBlockchainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.RetryMessageExecution(&_ITeleporterMessenger.TransactOpts, sourceBlockchainID, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x8245a1b0.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactor) RetrySendCrossChainMessage(opts *bind.TransactOpts, message TeleporterMessage) (*types.Transaction, error) {
	return _ITeleporterMessenger.contract.Transact(opts, "retrySendCrossChainMessage", message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x8245a1b0.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ITeleporterMessenger *ITeleporterMessengerSession) RetrySendCrossChainMessage(message TeleporterMessage) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.RetrySendCrossChainMessage(&_ITeleporterMessenger.TransactOpts, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x8245a1b0.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactorSession) RetrySendCrossChainMessage(message TeleporterMessage) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.RetrySendCrossChainMessage(&_ITeleporterMessenger.TransactOpts, message)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerTransactor) SendCrossChainMessage(opts *bind.TransactOpts, messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _ITeleporterMessenger.contract.Transact(opts, "sendCrossChainMessage", messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.SendCrossChainMessage(&_ITeleporterMessenger.TransactOpts, messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerTransactorSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.SendCrossChainMessage(&_ITeleporterMessenger.TransactOpts, messageInput)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerTransactor) SendSpecifiedReceipts(opts *bind.TransactOpts, sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.contract.Transact(opts, "sendSpecifiedReceipts", sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerSession) SendSpecifiedReceipts(sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.SendSpecifiedReceipts(&_ITeleporterMessenger.TransactOpts, sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerTransactorSession) SendSpecifiedReceipts(sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.SendSpecifiedReceipts(&_ITeleporterMessenger.TransactOpts, sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// ITeleporterMessengerAddFeeAmountIterator is returned from FilterAddFeeAmount and is used to iterate over the raw logs and unpacked data for AddFeeAmount events raised by the ITeleporterMessenger contract.
type ITeleporterMessengerAddFeeAmountIterator struct {
	Event *ITeleporterMessengerAddFeeAmount // Event containing the contract specifics and raw log

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
func (it *ITeleporterMessengerAddFeeAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITeleporterMessengerAddFeeAmount)
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
		it.Event = new(ITeleporterMessengerAddFeeAmount)
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
func (it *ITeleporterMessengerAddFeeAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITeleporterMessengerAddFeeAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITeleporterMessengerAddFeeAmount represents a AddFeeAmount event raised by the ITeleporterMessenger contract.
type ITeleporterMessengerAddFeeAmount struct {
	MessageID      [32]byte
	UpdatedFeeInfo TeleporterFeeInfo
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddFeeAmount is a free log retrieval operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) FilterAddFeeAmount(opts *bind.FilterOpts, messageID [][32]byte) (*ITeleporterMessengerAddFeeAmountIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.FilterLogs(opts, "AddFeeAmount", messageIDRule)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerAddFeeAmountIterator{contract: _ITeleporterMessenger.contract, event: "AddFeeAmount", logs: logs, sub: sub}, nil
}

// WatchAddFeeAmount is a free log subscription operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) WatchAddFeeAmount(opts *bind.WatchOpts, sink chan<- *ITeleporterMessengerAddFeeAmount, messageID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.WatchLogs(opts, "AddFeeAmount", messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITeleporterMessengerAddFeeAmount)
				if err := _ITeleporterMessenger.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
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

// ParseAddFeeAmount is a log parse operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) ParseAddFeeAmount(log types.Log) (*ITeleporterMessengerAddFeeAmount, error) {
	event := new(ITeleporterMessengerAddFeeAmount)
	if err := _ITeleporterMessenger.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITeleporterMessengerBlockchainIDInitializedIterator is returned from FilterBlockchainIDInitialized and is used to iterate over the raw logs and unpacked data for BlockchainIDInitialized events raised by the ITeleporterMessenger contract.
type ITeleporterMessengerBlockchainIDInitializedIterator struct {
	Event *ITeleporterMessengerBlockchainIDInitialized // Event containing the contract specifics and raw log

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
func (it *ITeleporterMessengerBlockchainIDInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITeleporterMessengerBlockchainIDInitialized)
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
		it.Event = new(ITeleporterMessengerBlockchainIDInitialized)
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
func (it *ITeleporterMessengerBlockchainIDInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITeleporterMessengerBlockchainIDInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITeleporterMessengerBlockchainIDInitialized represents a BlockchainIDInitialized event raised by the ITeleporterMessenger contract.
type ITeleporterMessengerBlockchainIDInitialized struct {
	BlockchainID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBlockchainIDInitialized is a free log retrieval operation binding the contract event 0x1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104.
//
// Solidity: event BlockchainIDInitialized(bytes32 indexed blockchainID)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) FilterBlockchainIDInitialized(opts *bind.FilterOpts, blockchainID [][32]byte) (*ITeleporterMessengerBlockchainIDInitializedIterator, error) {

	var blockchainIDRule []interface{}
	for _, blockchainIDItem := range blockchainID {
		blockchainIDRule = append(blockchainIDRule, blockchainIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.FilterLogs(opts, "BlockchainIDInitialized", blockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerBlockchainIDInitializedIterator{contract: _ITeleporterMessenger.contract, event: "BlockchainIDInitialized", logs: logs, sub: sub}, nil
}

// WatchBlockchainIDInitialized is a free log subscription operation binding the contract event 0x1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104.
//
// Solidity: event BlockchainIDInitialized(bytes32 indexed blockchainID)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) WatchBlockchainIDInitialized(opts *bind.WatchOpts, sink chan<- *ITeleporterMessengerBlockchainIDInitialized, blockchainID [][32]byte) (event.Subscription, error) {

	var blockchainIDRule []interface{}
	for _, blockchainIDItem := range blockchainID {
		blockchainIDRule = append(blockchainIDRule, blockchainIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.WatchLogs(opts, "BlockchainIDInitialized", blockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITeleporterMessengerBlockchainIDInitialized)
				if err := _ITeleporterMessenger.contract.UnpackLog(event, "BlockchainIDInitialized", log); err != nil {
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

// ParseBlockchainIDInitialized is a log parse operation binding the contract event 0x1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104.
//
// Solidity: event BlockchainIDInitialized(bytes32 indexed blockchainID)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) ParseBlockchainIDInitialized(log types.Log) (*ITeleporterMessengerBlockchainIDInitialized, error) {
	event := new(ITeleporterMessengerBlockchainIDInitialized)
	if err := _ITeleporterMessenger.contract.UnpackLog(event, "BlockchainIDInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITeleporterMessengerMessageExecutedIterator is returned from FilterMessageExecuted and is used to iterate over the raw logs and unpacked data for MessageExecuted events raised by the ITeleporterMessenger contract.
type ITeleporterMessengerMessageExecutedIterator struct {
	Event *ITeleporterMessengerMessageExecuted // Event containing the contract specifics and raw log

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
func (it *ITeleporterMessengerMessageExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITeleporterMessengerMessageExecuted)
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
		it.Event = new(ITeleporterMessengerMessageExecuted)
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
func (it *ITeleporterMessengerMessageExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITeleporterMessengerMessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITeleporterMessengerMessageExecuted represents a MessageExecuted event raised by the ITeleporterMessenger contract.
type ITeleporterMessengerMessageExecuted struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageExecuted is a free log retrieval operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) FilterMessageExecuted(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte) (*ITeleporterMessengerMessageExecutedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.FilterLogs(opts, "MessageExecuted", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerMessageExecutedIterator{contract: _ITeleporterMessenger.contract, event: "MessageExecuted", logs: logs, sub: sub}, nil
}

// WatchMessageExecuted is a free log subscription operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) WatchMessageExecuted(opts *bind.WatchOpts, sink chan<- *ITeleporterMessengerMessageExecuted, messageID [][32]byte, sourceBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.WatchLogs(opts, "MessageExecuted", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITeleporterMessengerMessageExecuted)
				if err := _ITeleporterMessenger.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
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

// ParseMessageExecuted is a log parse operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) ParseMessageExecuted(log types.Log) (*ITeleporterMessengerMessageExecuted, error) {
	event := new(ITeleporterMessengerMessageExecuted)
	if err := _ITeleporterMessenger.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITeleporterMessengerMessageExecutionFailedIterator is returned from FilterMessageExecutionFailed and is used to iterate over the raw logs and unpacked data for MessageExecutionFailed events raised by the ITeleporterMessenger contract.
type ITeleporterMessengerMessageExecutionFailedIterator struct {
	Event *ITeleporterMessengerMessageExecutionFailed // Event containing the contract specifics and raw log

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
func (it *ITeleporterMessengerMessageExecutionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITeleporterMessengerMessageExecutionFailed)
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
		it.Event = new(ITeleporterMessengerMessageExecutionFailed)
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
func (it *ITeleporterMessengerMessageExecutionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITeleporterMessengerMessageExecutionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITeleporterMessengerMessageExecutionFailed represents a MessageExecutionFailed event raised by the ITeleporterMessenger contract.
type ITeleporterMessengerMessageExecutionFailed struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Message            TeleporterMessage
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageExecutionFailed is a free log retrieval operation binding the contract event 0x4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) FilterMessageExecutionFailed(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte) (*ITeleporterMessengerMessageExecutionFailedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.FilterLogs(opts, "MessageExecutionFailed", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerMessageExecutionFailedIterator{contract: _ITeleporterMessenger.contract, event: "MessageExecutionFailed", logs: logs, sub: sub}, nil
}

// WatchMessageExecutionFailed is a free log subscription operation binding the contract event 0x4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) WatchMessageExecutionFailed(opts *bind.WatchOpts, sink chan<- *ITeleporterMessengerMessageExecutionFailed, messageID [][32]byte, sourceBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.WatchLogs(opts, "MessageExecutionFailed", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITeleporterMessengerMessageExecutionFailed)
				if err := _ITeleporterMessenger.contract.UnpackLog(event, "MessageExecutionFailed", log); err != nil {
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

// ParseMessageExecutionFailed is a log parse operation binding the contract event 0x4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) ParseMessageExecutionFailed(log types.Log) (*ITeleporterMessengerMessageExecutionFailed, error) {
	event := new(ITeleporterMessengerMessageExecutionFailed)
	if err := _ITeleporterMessenger.contract.UnpackLog(event, "MessageExecutionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITeleporterMessengerReceiptReceivedIterator is returned from FilterReceiptReceived and is used to iterate over the raw logs and unpacked data for ReceiptReceived events raised by the ITeleporterMessenger contract.
type ITeleporterMessengerReceiptReceivedIterator struct {
	Event *ITeleporterMessengerReceiptReceived // Event containing the contract specifics and raw log

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
func (it *ITeleporterMessengerReceiptReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITeleporterMessengerReceiptReceived)
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
		it.Event = new(ITeleporterMessengerReceiptReceived)
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
func (it *ITeleporterMessengerReceiptReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITeleporterMessengerReceiptReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITeleporterMessengerReceiptReceived represents a ReceiptReceived event raised by the ITeleporterMessenger contract.
type ITeleporterMessengerReceiptReceived struct {
	MessageID               [32]byte
	DestinationBlockchainID [32]byte
	RelayerRewardAddress    common.Address
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterReceiptReceived is a free log retrieval operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) FilterReceiptReceived(opts *bind.FilterOpts, messageID [][32]byte, destinationBlockchainID [][32]byte, relayerRewardAddress []common.Address) (*ITeleporterMessengerReceiptReceivedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var relayerRewardAddressRule []interface{}
	for _, relayerRewardAddressItem := range relayerRewardAddress {
		relayerRewardAddressRule = append(relayerRewardAddressRule, relayerRewardAddressItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.FilterLogs(opts, "ReceiptReceived", messageIDRule, destinationBlockchainIDRule, relayerRewardAddressRule)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerReceiptReceivedIterator{contract: _ITeleporterMessenger.contract, event: "ReceiptReceived", logs: logs, sub: sub}, nil
}

// WatchReceiptReceived is a free log subscription operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) WatchReceiptReceived(opts *bind.WatchOpts, sink chan<- *ITeleporterMessengerReceiptReceived, messageID [][32]byte, destinationBlockchainID [][32]byte, relayerRewardAddress []common.Address) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var relayerRewardAddressRule []interface{}
	for _, relayerRewardAddressItem := range relayerRewardAddress {
		relayerRewardAddressRule = append(relayerRewardAddressRule, relayerRewardAddressItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.WatchLogs(opts, "ReceiptReceived", messageIDRule, destinationBlockchainIDRule, relayerRewardAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITeleporterMessengerReceiptReceived)
				if err := _ITeleporterMessenger.contract.UnpackLog(event, "ReceiptReceived", log); err != nil {
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

// ParseReceiptReceived is a log parse operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) ParseReceiptReceived(log types.Log) (*ITeleporterMessengerReceiptReceived, error) {
	event := new(ITeleporterMessengerReceiptReceived)
	if err := _ITeleporterMessenger.contract.UnpackLog(event, "ReceiptReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITeleporterMessengerReceiveCrossChainMessageIterator is returned from FilterReceiveCrossChainMessage and is used to iterate over the raw logs and unpacked data for ReceiveCrossChainMessage events raised by the ITeleporterMessenger contract.
type ITeleporterMessengerReceiveCrossChainMessageIterator struct {
	Event *ITeleporterMessengerReceiveCrossChainMessage // Event containing the contract specifics and raw log

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
func (it *ITeleporterMessengerReceiveCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITeleporterMessengerReceiveCrossChainMessage)
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
		it.Event = new(ITeleporterMessengerReceiveCrossChainMessage)
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
func (it *ITeleporterMessengerReceiveCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITeleporterMessengerReceiveCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITeleporterMessengerReceiveCrossChainMessage represents a ReceiveCrossChainMessage event raised by the ITeleporterMessenger contract.
type ITeleporterMessengerReceiveCrossChainMessage struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Deliverer          common.Address
	RewardRedeemer     common.Address
	Message            TeleporterMessage
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReceiveCrossChainMessage is a free log retrieval operation binding the contract event 0x292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e34.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) FilterReceiveCrossChainMessage(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte, deliverer []common.Address) (*ITeleporterMessengerReceiveCrossChainMessageIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var delivererRule []interface{}
	for _, delivererItem := range deliverer {
		delivererRule = append(delivererRule, delivererItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.FilterLogs(opts, "ReceiveCrossChainMessage", messageIDRule, sourceBlockchainIDRule, delivererRule)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerReceiveCrossChainMessageIterator{contract: _ITeleporterMessenger.contract, event: "ReceiveCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchReceiveCrossChainMessage is a free log subscription operation binding the contract event 0x292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e34.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) WatchReceiveCrossChainMessage(opts *bind.WatchOpts, sink chan<- *ITeleporterMessengerReceiveCrossChainMessage, messageID [][32]byte, sourceBlockchainID [][32]byte, deliverer []common.Address) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var delivererRule []interface{}
	for _, delivererItem := range deliverer {
		delivererRule = append(delivererRule, delivererItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.WatchLogs(opts, "ReceiveCrossChainMessage", messageIDRule, sourceBlockchainIDRule, delivererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITeleporterMessengerReceiveCrossChainMessage)
				if err := _ITeleporterMessenger.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
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

// ParseReceiveCrossChainMessage is a log parse operation binding the contract event 0x292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e34.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) ParseReceiveCrossChainMessage(log types.Log) (*ITeleporterMessengerReceiveCrossChainMessage, error) {
	event := new(ITeleporterMessengerReceiveCrossChainMessage)
	if err := _ITeleporterMessenger.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITeleporterMessengerRelayerRewardsRedeemedIterator is returned from FilterRelayerRewardsRedeemed and is used to iterate over the raw logs and unpacked data for RelayerRewardsRedeemed events raised by the ITeleporterMessenger contract.
type ITeleporterMessengerRelayerRewardsRedeemedIterator struct {
	Event *ITeleporterMessengerRelayerRewardsRedeemed // Event containing the contract specifics and raw log

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
func (it *ITeleporterMessengerRelayerRewardsRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITeleporterMessengerRelayerRewardsRedeemed)
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
		it.Event = new(ITeleporterMessengerRelayerRewardsRedeemed)
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
func (it *ITeleporterMessengerRelayerRewardsRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITeleporterMessengerRelayerRewardsRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITeleporterMessengerRelayerRewardsRedeemed represents a RelayerRewardsRedeemed event raised by the ITeleporterMessenger contract.
type ITeleporterMessengerRelayerRewardsRedeemed struct {
	Redeemer common.Address
	Asset    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRelayerRewardsRedeemed is a free log retrieval operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) FilterRelayerRewardsRedeemed(opts *bind.FilterOpts, redeemer []common.Address, asset []common.Address) (*ITeleporterMessengerRelayerRewardsRedeemedIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.FilterLogs(opts, "RelayerRewardsRedeemed", redeemerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerRelayerRewardsRedeemedIterator{contract: _ITeleporterMessenger.contract, event: "RelayerRewardsRedeemed", logs: logs, sub: sub}, nil
}

// WatchRelayerRewardsRedeemed is a free log subscription operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) WatchRelayerRewardsRedeemed(opts *bind.WatchOpts, sink chan<- *ITeleporterMessengerRelayerRewardsRedeemed, redeemer []common.Address, asset []common.Address) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.WatchLogs(opts, "RelayerRewardsRedeemed", redeemerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITeleporterMessengerRelayerRewardsRedeemed)
				if err := _ITeleporterMessenger.contract.UnpackLog(event, "RelayerRewardsRedeemed", log); err != nil {
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

// ParseRelayerRewardsRedeemed is a log parse operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) ParseRelayerRewardsRedeemed(log types.Log) (*ITeleporterMessengerRelayerRewardsRedeemed, error) {
	event := new(ITeleporterMessengerRelayerRewardsRedeemed)
	if err := _ITeleporterMessenger.contract.UnpackLog(event, "RelayerRewardsRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITeleporterMessengerSendCrossChainMessageIterator is returned from FilterSendCrossChainMessage and is used to iterate over the raw logs and unpacked data for SendCrossChainMessage events raised by the ITeleporterMessenger contract.
type ITeleporterMessengerSendCrossChainMessageIterator struct {
	Event *ITeleporterMessengerSendCrossChainMessage // Event containing the contract specifics and raw log

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
func (it *ITeleporterMessengerSendCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITeleporterMessengerSendCrossChainMessage)
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
		it.Event = new(ITeleporterMessengerSendCrossChainMessage)
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
func (it *ITeleporterMessengerSendCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITeleporterMessengerSendCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITeleporterMessengerSendCrossChainMessage represents a SendCrossChainMessage event raised by the ITeleporterMessenger contract.
type ITeleporterMessengerSendCrossChainMessage struct {
	MessageID               [32]byte
	DestinationBlockchainID [32]byte
	Message                 TeleporterMessage
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSendCrossChainMessage is a free log retrieval operation binding the contract event 0x2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df8.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) FilterSendCrossChainMessage(opts *bind.FilterOpts, messageID [][32]byte, destinationBlockchainID [][32]byte) (*ITeleporterMessengerSendCrossChainMessageIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.FilterLogs(opts, "SendCrossChainMessage", messageIDRule, destinationBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerSendCrossChainMessageIterator{contract: _ITeleporterMessenger.contract, event: "SendCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchSendCrossChainMessage is a free log subscription operation binding the contract event 0x2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df8.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) WatchSendCrossChainMessage(opts *bind.WatchOpts, sink chan<- *ITeleporterMessengerSendCrossChainMessage, messageID [][32]byte, destinationBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.WatchLogs(opts, "SendCrossChainMessage", messageIDRule, destinationBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITeleporterMessengerSendCrossChainMessage)
				if err := _ITeleporterMessenger.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
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

// ParseSendCrossChainMessage is a log parse operation binding the contract event 0x2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df8.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) ParseSendCrossChainMessage(log types.Log) (*ITeleporterMessengerSendCrossChainMessage, error) {
	event := new(ITeleporterMessengerSendCrossChainMessage)
	if err := _ITeleporterMessenger.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
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

// TeleporterRegistryMetaData contains all meta data concerning the TeleporterRegistry contract.
var TeleporterRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"internalType\":\"structProtocolRegistryEntry[]\",\"name\":\"initialEntries\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"name\":\"AddProtocolVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newVersion\",\"type\":\"uint256\"}],\"name\":\"LatestVersionUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_VERSION_INCREMENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORS_SOURCE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"addProtocolVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getAddressFromVersion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestTeleporter\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getTeleporterFromVersion\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"name\":\"getVersionFromAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051610fca380380610fca83398101604081905261002e916103e2565b7302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa15801561007e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100a291906104bc565b60805280515f5b818110156100e1576100d98382815181106100c6576100c66104d3565b60200260200101516100e960201b60201c565b6001016100a9565b50505061050c565b80515f0361013e5760405162461bcd60e51b815260206004820181905260248201527f54656c65706f7274657252656769737472793a207a65726f2076657273696f6e60448201526064015b60405180910390fd5b80515f908152600160205260409020546001600160a01b0316156101b75760405162461bcd60e51b815260206004820152602a60248201527f54656c65706f7274657252656769737472793a2076657273696f6e20616c72656044820152696164792065786973747360b01b6064820152608401610135565b60208101516001600160a01b03166102235760405162461bcd60e51b815260206004820152602960248201527f54656c65706f7274657252656769737472793a207a65726f2070726f746f636f6044820152686c206164647265737360b81b6064820152608401610135565b5f546102316101f4826104e7565b825111156102985760405162461bcd60e51b815260206004820152602e60248201527f54656c65706f7274657252656769737472793a2076657273696f6e20696e637260448201526d0cadacadce840e8dede40d0d2ced60931b6064820152608401610135565b6020828101805184515f90815260018452604080822080546001600160a01b0319166001600160a01b03948516179055925190911681526002909252902054825111156102ff5781516020808401516001600160a01b03165f908152600290915260409020555b602082015182516040516001600160a01b03909216917fa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a905f90a381518110156103725781515f81815560405183917f30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b091a35b5050565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156103ac576103ac610376565b60405290565b604051601f8201601f191681016001600160401b03811182821017156103da576103da610376565b604052919050565b5f60208083850312156103f3575f80fd5b82516001600160401b0380821115610409575f80fd5b818501915085601f83011261041c575f80fd5b81518181111561042e5761042e610376565b61043c848260051b016103b2565b818152848101925060069190911b83018401908782111561045b575f80fd5b928401925b818410156104b15760408489031215610477575f80fd5b61047f61038a565b84518152858501516001600160a01b038116811461049b575f80fd5b8187015283526040939093019291840191610460565b979650505050505050565b5f602082840312156104cc575f80fd5b5051919050565b634e487b7160e01b5f52603260045260245ffd5b8082018082111561050657634e487b7160e01b5f52601160045260245ffd5b92915050565b608051610a9f61052b5f395f818161014301526102580152610a9f5ff3fe608060405234801561000f575f80fd5b506004361061009b575f3560e01c8063ac473ac311610063578063ac473ac31461011f578063b771b3bc14610128578063c07f47d414610136578063d127dc9b1461013e578063d820e64f14610165575f80fd5b80630731775d1461009f578063215abce9146100c357806341f34ed9146100d657806346f9ef49146100eb5780634c1f08ce146100fe575b5f80fd5b6100a65f81565b6040516001600160a01b0390911681526020015b60405180910390f35b6100a66100d13660046107aa565b61016d565b6100e96100e43660046107c1565b61017d565b005b6100a66100f93660046107aa565b6103ec565b61011161010c366004610802565b6104ae565b6040519081526020016100ba565b6101116101f481565b6100a66005600160991b0181565b6101115f5481565b6101117f000000000000000000000000000000000000000000000000000000000000000081565b6100a6610554565b5f610177826103ec565b92915050565b6040516306f8253560e41b815263ffffffff821660048201525f9081906005600160991b0190636f825350906024015f60405180830381865afa1580156101c6573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526101ed919081019061089f565b91509150806102545760405162461bcd60e51b815260206004820152602860248201527f54656c65706f7274657252656769737472793a20696e76616c69642077617270604482015267206d65737361676560c01b60648201526084015b60405180910390fd5b81517f0000000000000000000000000000000000000000000000000000000000000000146102d85760405162461bcd60e51b815260206004820152602b60248201527f54656c65706f7274657252656769737472793a20696e76616c696420736f757260448201526a18d94818da185a5b88125160aa1b606482015260840161024b565b60208201516001600160a01b03161561034d5760405162461bcd60e51b815260206004820152603160248201527f54656c65706f7274657252656769737472793a20696e76616c6964206f726967604482015270696e2073656e646572206164647265737360781b606482015260840161024b565b5f8083604001518060200190518101906103679190610988565b90925090506001600160a01b03811630146103dc5760405162461bcd60e51b815260206004820152602f60248201527f54656c65706f7274657252656769737472793a20696e76616c6964206465737460448201526e696e6174696f6e206164647265737360881b606482015260840161024b565b6103e582610564565b5050505050565b5f815f0361043c5760405162461bcd60e51b815260206004820181905260248201527f54656c65706f7274657252656769737472793a207a65726f2076657273696f6e604482015260640161024b565b5f828152600160205260409020546001600160a01b0316806101775760405162461bcd60e51b815260206004820152602560248201527f54656c65706f7274657252656769737472793a2076657273696f6e206e6f7420604482015264199bdd5b9960da1b606482015260840161024b565b5f6001600160a01b0382166104d55760405162461bcd60e51b815260040161024b90610a01565b6001600160a01b0382165f90815260026020526040812054908190036101775760405162461bcd60e51b815260206004820152602e60248201527f54656c65706f7274657252656769737472793a2070726f746f636f6c2061646460448201526d1c995cdcc81b9bdd08199bdd5b9960921b606482015260840161024b565b5f61055f5f546103ec565b905090565b80515f036105b45760405162461bcd60e51b815260206004820181905260248201527f54656c65706f7274657252656769737472793a207a65726f2076657273696f6e604482015260640161024b565b80515f908152600160205260409020546001600160a01b03161561062d5760405162461bcd60e51b815260206004820152602a60248201527f54656c65706f7274657252656769737472793a2076657273696f6e20616c72656044820152696164792065786973747360b01b606482015260840161024b565b60208101516001600160a01b03166106575760405162461bcd60e51b815260040161024b90610a01565b5f546106656101f482610a4a565b825111156106cc5760405162461bcd60e51b815260206004820152602e60248201527f54656c65706f7274657252656769737472793a2076657273696f6e20696e637260448201526d0cadacadce840e8dede40d0d2ced60931b606482015260840161024b565b6020828101805184515f90815260018452604080822080546001600160a01b0319166001600160a01b03948516179055925190911681526002909252902054825111156107335781516020808401516001600160a01b03165f908152600290915260409020555b602082015182516040516001600160a01b03909216917fa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a905f90a381518110156107a65781515f81815560405183917f30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b091a35b5050565b5f602082840312156107ba575f80fd5b5035919050565b5f602082840312156107d1575f80fd5b813563ffffffff811681146107e4575f80fd5b9392505050565b6001600160a01b03811681146107ff575f80fd5b50565b5f60208284031215610812575f80fd5b81356107e4816107eb565b634e487b7160e01b5f52604160045260245ffd5b6040516060810167ffffffffffffffff811182821017156108545761085461081d565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156108835761088361081d565b604052919050565b8051801515811461089a575f80fd5b919050565b5f80604083850312156108b0575f80fd5b825167ffffffffffffffff808211156108c7575f80fd5b90840190606082870312156108da575f80fd5b6108e2610831565b825181526020808401516108f5816107eb565b8282015260408401518381111561090a575f80fd5b80850194505087601f85011261091e575f80fd5b8351838111156109305761093061081d565b610942601f8201601f1916830161085a565b93508084528882828701011115610957575f80fd5b808286018386015e5f8282860101525082604083015281955061097b81880161088b565b9450505050509250929050565b5f80828403606081121561099a575f80fd5b60408112156109a7575f80fd5b506040516040810181811067ffffffffffffffff821117156109cb576109cb61081d565b6040528351815260208401516109e0816107eb565b602082015260408401519092506109f6816107eb565b809150509250929050565b60208082526029908201527f54656c65706f7274657252656769737472793a207a65726f2070726f746f636f6040820152686c206164647265737360b81b606082015260800190565b8082018082111561017757634e487b7160e01b5f52601160045260245ffdfea2646970667358221220fe036b8dd4c0898f583fa51960ec46bcf6307508fd784944cdf1ccf51969ca8564736f6c63430008190033",
}

// TeleporterRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleporterRegistryMetaData.ABI instead.
var TeleporterRegistryABI = TeleporterRegistryMetaData.ABI

// TeleporterRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TeleporterRegistryMetaData.Bin instead.
var TeleporterRegistryBin = TeleporterRegistryMetaData.Bin

// DeployTeleporterRegistry deploys a new Ethereum contract, binding an instance of TeleporterRegistry to it.
func DeployTeleporterRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, initialEntries []ProtocolRegistryEntry) (common.Address, *types.Transaction, *TeleporterRegistry, error) {
	parsed, err := TeleporterRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TeleporterRegistryBin), backend, initialEntries)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TeleporterRegistry{TeleporterRegistryCaller: TeleporterRegistryCaller{contract: contract}, TeleporterRegistryTransactor: TeleporterRegistryTransactor{contract: contract}, TeleporterRegistryFilterer: TeleporterRegistryFilterer{contract: contract}}, nil
}

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

// MAXVERSIONINCREMENT is a free data retrieval call binding the contract method 0xac473ac3.
//
// Solidity: function MAX_VERSION_INCREMENT() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCaller) MAXVERSIONINCREMENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "MAX_VERSION_INCREMENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVERSIONINCREMENT is a free data retrieval call binding the contract method 0xac473ac3.
//
// Solidity: function MAX_VERSION_INCREMENT() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistrySession) MAXVERSIONINCREMENT() (*big.Int, error) {
	return _TeleporterRegistry.Contract.MAXVERSIONINCREMENT(&_TeleporterRegistry.CallOpts)
}

// MAXVERSIONINCREMENT is a free data retrieval call binding the contract method 0xac473ac3.
//
// Solidity: function MAX_VERSION_INCREMENT() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) MAXVERSIONINCREMENT() (*big.Int, error) {
	return _TeleporterRegistry.Contract.MAXVERSIONINCREMENT(&_TeleporterRegistry.CallOpts)
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

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterRegistry *TeleporterRegistryCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterRegistry *TeleporterRegistrySession) BlockchainID() ([32]byte, error) {
	return _TeleporterRegistry.Contract.BlockchainID(&_TeleporterRegistry.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) BlockchainID() ([32]byte, error) {
	return _TeleporterRegistry.Contract.BlockchainID(&_TeleporterRegistry.CallOpts)
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

// LatestVersion is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCaller) LatestVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "latestVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestVersion is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistrySession) LatestVersion() (*big.Int, error) {
	return _TeleporterRegistry.Contract.LatestVersion(&_TeleporterRegistry.CallOpts)
}

// LatestVersion is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) LatestVersion() (*big.Int, error) {
	return _TeleporterRegistry.Contract.LatestVersion(&_TeleporterRegistry.CallOpts)
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

// TeleporterRegistryLatestVersionUpdatedIterator is returned from FilterLatestVersionUpdated and is used to iterate over the raw logs and unpacked data for LatestVersionUpdated events raised by the TeleporterRegistry contract.
type TeleporterRegistryLatestVersionUpdatedIterator struct {
	Event *TeleporterRegistryLatestVersionUpdated // Event containing the contract specifics and raw log

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
func (it *TeleporterRegistryLatestVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterRegistryLatestVersionUpdated)
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
		it.Event = new(TeleporterRegistryLatestVersionUpdated)
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
func (it *TeleporterRegistryLatestVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterRegistryLatestVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterRegistryLatestVersionUpdated represents a LatestVersionUpdated event raised by the TeleporterRegistry contract.
type TeleporterRegistryLatestVersionUpdated struct {
	OldVersion *big.Int
	NewVersion *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLatestVersionUpdated is a free log retrieval operation binding the contract event 0x30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b0.
//
// Solidity: event LatestVersionUpdated(uint256 indexed oldVersion, uint256 indexed newVersion)
func (_TeleporterRegistry *TeleporterRegistryFilterer) FilterLatestVersionUpdated(opts *bind.FilterOpts, oldVersion []*big.Int, newVersion []*big.Int) (*TeleporterRegistryLatestVersionUpdatedIterator, error) {

	var oldVersionRule []interface{}
	for _, oldVersionItem := range oldVersion {
		oldVersionRule = append(oldVersionRule, oldVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _TeleporterRegistry.contract.FilterLogs(opts, "LatestVersionUpdated", oldVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistryLatestVersionUpdatedIterator{contract: _TeleporterRegistry.contract, event: "LatestVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchLatestVersionUpdated is a free log subscription operation binding the contract event 0x30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b0.
//
// Solidity: event LatestVersionUpdated(uint256 indexed oldVersion, uint256 indexed newVersion)
func (_TeleporterRegistry *TeleporterRegistryFilterer) WatchLatestVersionUpdated(opts *bind.WatchOpts, sink chan<- *TeleporterRegistryLatestVersionUpdated, oldVersion []*big.Int, newVersion []*big.Int) (event.Subscription, error) {

	var oldVersionRule []interface{}
	for _, oldVersionItem := range oldVersion {
		oldVersionRule = append(oldVersionRule, oldVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _TeleporterRegistry.contract.WatchLogs(opts, "LatestVersionUpdated", oldVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterRegistryLatestVersionUpdated)
				if err := _TeleporterRegistry.contract.UnpackLog(event, "LatestVersionUpdated", log); err != nil {
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

// ParseLatestVersionUpdated is a log parse operation binding the contract event 0x30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b0.
//
// Solidity: event LatestVersionUpdated(uint256 indexed oldVersion, uint256 indexed newVersion)
func (_TeleporterRegistry *TeleporterRegistryFilterer) ParseLatestVersionUpdated(log types.Log) (*TeleporterRegistryLatestVersionUpdated, error) {
	event := new(TeleporterRegistryLatestVersionUpdated)
	if err := _TeleporterRegistry.contract.UnpackLog(event, "LatestVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
