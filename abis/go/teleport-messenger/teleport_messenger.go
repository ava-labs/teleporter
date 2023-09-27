// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teleport_messenger

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

// TeleporterFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type TeleporterFeeInfo struct {
	ContractAddress common.Address
	Amount          *big.Int
}

// TeleporterMessage is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessage struct {
	MessageID               *big.Int
	SenderAddress           common.Address
	DestinationAddress      common.Address
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Receipts                []TeleporterMessageReceipt
	Message                 []byte
}

// TeleporterMessageInput is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageInput struct {
	DestinationChainID      [32]byte
	DestinationAddress      common.Address
	FeeInfo                 TeleporterFeeInfo
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Message                 []byte
}

// TeleporterMessageReceipt is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageReceipt struct {
	ReceivedMessageID    *big.Int
	RelayerRewardAddress common.Address
}

// TeleportMessengerMetaData contains all meta data concerning the TeleportMessenger contract.
var TeleportMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BalanceNotIncreased\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdditionalFeeAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDestinationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDestinationChainID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFeeAssetContractAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMessageHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOriginSenderAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRelayerRewardAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWarpMessage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageAlreadyDelivered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageRetryExecutionFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoRelayerRewardToRedeem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiptNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverReentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderReentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedRelayer\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"updatedFeeInfo\",\"type\":\"tuple\"}],\"name\":\"AddFeeAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"FailedMessageExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"MessageExecutionRetried\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ReceiveCrossChainMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"SendCrossChainMessage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAXIMUM_RECEIPT_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_REQUIRED_CALL_DATA_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REQUIRED_ORIGIN_CHAIN_ID_START_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalFeeAmount\",\"type\":\"uint256\"}],\"name\":\"addFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delivererAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayers\",\"type\":\"address[]\"}],\"name\":\"checkIsAllowedRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"checkRelayerRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"getFeeInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chainID\",\"type\":\"bytes32\"}],\"name\":\"getNextMessageID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"getRelayerRewardAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"latestMessageIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"messageReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"delivered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"outstandingReceipts\",\"outputs\":[{\"internalType\":\"contractReceiptQueue\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"name\":\"receiveCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"receivedFailedMessageHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"redeemRelayerRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"relayerRewardAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"relayerRewardAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retryMessageExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"messageIDs\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"retryReceipts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retrySendCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageInput\",\"name\":\"messageInput\",\"type\":\"tuple\"}],\"name\":\"sendCrossChainMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sentMessageInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TeleportMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleportMessengerMetaData.ABI instead.
var TeleportMessengerABI = TeleportMessengerMetaData.ABI

// TeleportMessenger is an auto generated Go binding around an Ethereum contract.
type TeleportMessenger struct {
	TeleportMessengerCaller     // Read-only binding to the contract
	TeleportMessengerTransactor // Write-only binding to the contract
	TeleportMessengerFilterer   // Log filterer for contract events
}

// TeleportMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeleportMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleportMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeleportMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleportMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeleportMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleportMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeleportMessengerSession struct {
	Contract     *TeleportMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TeleportMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeleportMessengerCallerSession struct {
	Contract *TeleportMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TeleportMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeleportMessengerTransactorSession struct {
	Contract     *TeleportMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TeleportMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeleportMessengerRaw struct {
	Contract *TeleportMessenger // Generic contract binding to access the raw methods on
}

// TeleportMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeleportMessengerCallerRaw struct {
	Contract *TeleportMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// TeleportMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeleportMessengerTransactorRaw struct {
	Contract *TeleportMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeleportMessenger creates a new instance of TeleportMessenger, bound to a specific deployed contract.
func NewTeleportMessenger(address common.Address, backend bind.ContractBackend) (*TeleportMessenger, error) {
	contract, err := bindTeleportMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeleportMessenger{TeleportMessengerCaller: TeleportMessengerCaller{contract: contract}, TeleportMessengerTransactor: TeleportMessengerTransactor{contract: contract}, TeleportMessengerFilterer: TeleportMessengerFilterer{contract: contract}}, nil
}

// NewTeleportMessengerCaller creates a new read-only instance of TeleportMessenger, bound to a specific deployed contract.
func NewTeleportMessengerCaller(address common.Address, caller bind.ContractCaller) (*TeleportMessengerCaller, error) {
	contract, err := bindTeleportMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeleportMessengerCaller{contract: contract}, nil
}

// NewTeleportMessengerTransactor creates a new write-only instance of TeleportMessenger, bound to a specific deployed contract.
func NewTeleportMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*TeleportMessengerTransactor, error) {
	contract, err := bindTeleportMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeleportMessengerTransactor{contract: contract}, nil
}

// NewTeleportMessengerFilterer creates a new log filterer instance of TeleportMessenger, bound to a specific deployed contract.
func NewTeleportMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*TeleportMessengerFilterer, error) {
	contract, err := bindTeleportMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeleportMessengerFilterer{contract: contract}, nil
}

// bindTeleportMessenger binds a generic wrapper to an already deployed contract.
func bindTeleportMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeleportMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleportMessenger *TeleportMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleportMessenger.Contract.TeleportMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleportMessenger *TeleportMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.TeleportMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleportMessenger *TeleportMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.TeleportMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleportMessenger *TeleportMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleportMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleportMessenger *TeleportMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleportMessenger *TeleportMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.contract.Transact(opts, method, params...)
}

// MAXIMUMRECEIPTCOUNT is a free data retrieval call binding the contract method 0x10534371.
//
// Solidity: function MAXIMUM_RECEIPT_COUNT() view returns(uint256)
func (_TeleportMessenger *TeleportMessengerCaller) MAXIMUMRECEIPTCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "MAXIMUM_RECEIPT_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXIMUMRECEIPTCOUNT is a free data retrieval call binding the contract method 0x10534371.
//
// Solidity: function MAXIMUM_RECEIPT_COUNT() view returns(uint256)
func (_TeleportMessenger *TeleportMessengerSession) MAXIMUMRECEIPTCOUNT() (*big.Int, error) {
	return _TeleportMessenger.Contract.MAXIMUMRECEIPTCOUNT(&_TeleportMessenger.CallOpts)
}

// MAXIMUMRECEIPTCOUNT is a free data retrieval call binding the contract method 0x10534371.
//
// Solidity: function MAXIMUM_RECEIPT_COUNT() view returns(uint256)
func (_TeleportMessenger *TeleportMessengerCallerSession) MAXIMUMRECEIPTCOUNT() (*big.Int, error) {
	return _TeleportMessenger.Contract.MAXIMUMRECEIPTCOUNT(&_TeleportMessenger.CallOpts)
}

// MINIMUMREQUIREDCALLDATALENGTH is a free data retrieval call binding the contract method 0x8f12376f.
//
// Solidity: function MINIMUM_REQUIRED_CALL_DATA_LENGTH() view returns(uint256)
func (_TeleportMessenger *TeleportMessengerCaller) MINIMUMREQUIREDCALLDATALENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "MINIMUM_REQUIRED_CALL_DATA_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMREQUIREDCALLDATALENGTH is a free data retrieval call binding the contract method 0x8f12376f.
//
// Solidity: function MINIMUM_REQUIRED_CALL_DATA_LENGTH() view returns(uint256)
func (_TeleportMessenger *TeleportMessengerSession) MINIMUMREQUIREDCALLDATALENGTH() (*big.Int, error) {
	return _TeleportMessenger.Contract.MINIMUMREQUIREDCALLDATALENGTH(&_TeleportMessenger.CallOpts)
}

// MINIMUMREQUIREDCALLDATALENGTH is a free data retrieval call binding the contract method 0x8f12376f.
//
// Solidity: function MINIMUM_REQUIRED_CALL_DATA_LENGTH() view returns(uint256)
func (_TeleportMessenger *TeleportMessengerCallerSession) MINIMUMREQUIREDCALLDATALENGTH() (*big.Int, error) {
	return _TeleportMessenger.Contract.MINIMUMREQUIREDCALLDATALENGTH(&_TeleportMessenger.CallOpts)
}

// REQUIREDORIGINCHAINIDSTARTINDEX is a free data retrieval call binding the contract method 0x5bf91119.
//
// Solidity: function REQUIRED_ORIGIN_CHAIN_ID_START_INDEX() view returns(uint256)
func (_TeleportMessenger *TeleportMessengerCaller) REQUIREDORIGINCHAINIDSTARTINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "REQUIRED_ORIGIN_CHAIN_ID_START_INDEX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REQUIREDORIGINCHAINIDSTARTINDEX is a free data retrieval call binding the contract method 0x5bf91119.
//
// Solidity: function REQUIRED_ORIGIN_CHAIN_ID_START_INDEX() view returns(uint256)
func (_TeleportMessenger *TeleportMessengerSession) REQUIREDORIGINCHAINIDSTARTINDEX() (*big.Int, error) {
	return _TeleportMessenger.Contract.REQUIREDORIGINCHAINIDSTARTINDEX(&_TeleportMessenger.CallOpts)
}

// REQUIREDORIGINCHAINIDSTARTINDEX is a free data retrieval call binding the contract method 0x5bf91119.
//
// Solidity: function REQUIRED_ORIGIN_CHAIN_ID_START_INDEX() view returns(uint256)
func (_TeleportMessenger *TeleportMessengerCallerSession) REQUIREDORIGINCHAINIDSTARTINDEX() (*big.Int, error) {
	return _TeleportMessenger.Contract.REQUIREDORIGINCHAINIDSTARTINDEX(&_TeleportMessenger.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleportMessenger *TeleportMessengerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleportMessenger *TeleportMessengerSession) WARPMESSENGER() (common.Address, error) {
	return _TeleportMessenger.Contract.WARPMESSENGER(&_TeleportMessenger.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleportMessenger *TeleportMessengerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _TeleportMessenger.Contract.WARPMESSENGER(&_TeleportMessenger.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleportMessenger *TeleportMessengerCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleportMessenger *TeleportMessengerSession) BlockchainID() ([32]byte, error) {
	return _TeleportMessenger.Contract.BlockchainID(&_TeleportMessenger.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleportMessenger *TeleportMessengerCallerSession) BlockchainID() ([32]byte, error) {
	return _TeleportMessenger.Contract.BlockchainID(&_TeleportMessenger.CallOpts)
}

// CheckIsAllowedRelayer is a free data retrieval call binding the contract method 0x65171908.
//
// Solidity: function checkIsAllowedRelayer(address delivererAddress, address[] allowedRelayers) pure returns(bool)
func (_TeleportMessenger *TeleportMessengerCaller) CheckIsAllowedRelayer(opts *bind.CallOpts, delivererAddress common.Address, allowedRelayers []common.Address) (bool, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "checkIsAllowedRelayer", delivererAddress, allowedRelayers)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckIsAllowedRelayer is a free data retrieval call binding the contract method 0x65171908.
//
// Solidity: function checkIsAllowedRelayer(address delivererAddress, address[] allowedRelayers) pure returns(bool)
func (_TeleportMessenger *TeleportMessengerSession) CheckIsAllowedRelayer(delivererAddress common.Address, allowedRelayers []common.Address) (bool, error) {
	return _TeleportMessenger.Contract.CheckIsAllowedRelayer(&_TeleportMessenger.CallOpts, delivererAddress, allowedRelayers)
}

// CheckIsAllowedRelayer is a free data retrieval call binding the contract method 0x65171908.
//
// Solidity: function checkIsAllowedRelayer(address delivererAddress, address[] allowedRelayers) pure returns(bool)
func (_TeleportMessenger *TeleportMessengerCallerSession) CheckIsAllowedRelayer(delivererAddress common.Address, allowedRelayers []common.Address) (bool, error) {
	return _TeleportMessenger.Contract.CheckIsAllowedRelayer(&_TeleportMessenger.CallOpts, delivererAddress, allowedRelayers)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleportMessenger *TeleportMessengerCaller) CheckRelayerRewardAmount(opts *bind.CallOpts, relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "checkRelayerRewardAmount", relayer, feeAsset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleportMessenger *TeleportMessengerSession) CheckRelayerRewardAmount(relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	return _TeleportMessenger.Contract.CheckRelayerRewardAmount(&_TeleportMessenger.CallOpts, relayer, feeAsset)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleportMessenger *TeleportMessengerCallerSession) CheckRelayerRewardAmount(relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	return _TeleportMessenger.Contract.CheckRelayerRewardAmount(&_TeleportMessenger.CallOpts, relayer, feeAsset)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0x82f2c43a.
//
// Solidity: function getFeeInfo(bytes32 destinationChainID, uint256 messageID) view returns(address feeAsset, uint256 feeAmount)
func (_TeleportMessenger *TeleportMessengerCaller) GetFeeInfo(opts *bind.CallOpts, destinationChainID [32]byte, messageID *big.Int) (struct {
	FeeAsset  common.Address
	FeeAmount *big.Int
}, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "getFeeInfo", destinationChainID, messageID)

	outstruct := new(struct {
		FeeAsset  common.Address
		FeeAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeAsset = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.FeeAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetFeeInfo is a free data retrieval call binding the contract method 0x82f2c43a.
//
// Solidity: function getFeeInfo(bytes32 destinationChainID, uint256 messageID) view returns(address feeAsset, uint256 feeAmount)
func (_TeleportMessenger *TeleportMessengerSession) GetFeeInfo(destinationChainID [32]byte, messageID *big.Int) (struct {
	FeeAsset  common.Address
	FeeAmount *big.Int
}, error) {
	return _TeleportMessenger.Contract.GetFeeInfo(&_TeleportMessenger.CallOpts, destinationChainID, messageID)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0x82f2c43a.
//
// Solidity: function getFeeInfo(bytes32 destinationChainID, uint256 messageID) view returns(address feeAsset, uint256 feeAmount)
func (_TeleportMessenger *TeleportMessengerCallerSession) GetFeeInfo(destinationChainID [32]byte, messageID *big.Int) (struct {
	FeeAsset  common.Address
	FeeAmount *big.Int
}, error) {
	return _TeleportMessenger.Contract.GetFeeInfo(&_TeleportMessenger.CallOpts, destinationChainID, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x220c9568.
//
// Solidity: function getMessageHash(bytes32 destinationChainID, uint256 messageID) view returns(bytes32 messageHash)
func (_TeleportMessenger *TeleportMessengerCaller) GetMessageHash(opts *bind.CallOpts, destinationChainID [32]byte, messageID *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "getMessageHash", destinationChainID, messageID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x220c9568.
//
// Solidity: function getMessageHash(bytes32 destinationChainID, uint256 messageID) view returns(bytes32 messageHash)
func (_TeleportMessenger *TeleportMessengerSession) GetMessageHash(destinationChainID [32]byte, messageID *big.Int) ([32]byte, error) {
	return _TeleportMessenger.Contract.GetMessageHash(&_TeleportMessenger.CallOpts, destinationChainID, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x220c9568.
//
// Solidity: function getMessageHash(bytes32 destinationChainID, uint256 messageID) view returns(bytes32 messageHash)
func (_TeleportMessenger *TeleportMessengerCallerSession) GetMessageHash(destinationChainID [32]byte, messageID *big.Int) ([32]byte, error) {
	return _TeleportMessenger.Contract.GetMessageHash(&_TeleportMessenger.CallOpts, destinationChainID, messageID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 chainID) view returns(uint256 messageID)
func (_TeleportMessenger *TeleportMessengerCaller) GetNextMessageID(opts *bind.CallOpts, chainID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "getNextMessageID", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 chainID) view returns(uint256 messageID)
func (_TeleportMessenger *TeleportMessengerSession) GetNextMessageID(chainID [32]byte) (*big.Int, error) {
	return _TeleportMessenger.Contract.GetNextMessageID(&_TeleportMessenger.CallOpts, chainID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 chainID) view returns(uint256 messageID)
func (_TeleportMessenger *TeleportMessengerCallerSession) GetNextMessageID(chainID [32]byte) (*big.Int, error) {
	return _TeleportMessenger.Contract.GetNextMessageID(&_TeleportMessenger.CallOpts, chainID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x33e890fe.
//
// Solidity: function getRelayerRewardAddress(bytes32 originChainID, uint256 messageID) view returns(address relayerRewardAddress)
func (_TeleportMessenger *TeleportMessengerCaller) GetRelayerRewardAddress(opts *bind.CallOpts, originChainID [32]byte, messageID *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "getRelayerRewardAddress", originChainID, messageID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x33e890fe.
//
// Solidity: function getRelayerRewardAddress(bytes32 originChainID, uint256 messageID) view returns(address relayerRewardAddress)
func (_TeleportMessenger *TeleportMessengerSession) GetRelayerRewardAddress(originChainID [32]byte, messageID *big.Int) (common.Address, error) {
	return _TeleportMessenger.Contract.GetRelayerRewardAddress(&_TeleportMessenger.CallOpts, originChainID, messageID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x33e890fe.
//
// Solidity: function getRelayerRewardAddress(bytes32 originChainID, uint256 messageID) view returns(address relayerRewardAddress)
func (_TeleportMessenger *TeleportMessengerCallerSession) GetRelayerRewardAddress(originChainID [32]byte, messageID *big.Int) (common.Address, error) {
	return _TeleportMessenger.Contract.GetRelayerRewardAddress(&_TeleportMessenger.CallOpts, originChainID, messageID)
}

// LatestMessageIDs is a free data retrieval call binding the contract method 0x29ec9beb.
//
// Solidity: function latestMessageIDs(bytes32 ) view returns(uint256)
func (_TeleportMessenger *TeleportMessengerCaller) LatestMessageIDs(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "latestMessageIDs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestMessageIDs is a free data retrieval call binding the contract method 0x29ec9beb.
//
// Solidity: function latestMessageIDs(bytes32 ) view returns(uint256)
func (_TeleportMessenger *TeleportMessengerSession) LatestMessageIDs(arg0 [32]byte) (*big.Int, error) {
	return _TeleportMessenger.Contract.LatestMessageIDs(&_TeleportMessenger.CallOpts, arg0)
}

// LatestMessageIDs is a free data retrieval call binding the contract method 0x29ec9beb.
//
// Solidity: function latestMessageIDs(bytes32 ) view returns(uint256)
func (_TeleportMessenger *TeleportMessengerCallerSession) LatestMessageIDs(arg0 [32]byte) (*big.Int, error) {
	return _TeleportMessenger.Contract.LatestMessageIDs(&_TeleportMessenger.CallOpts, arg0)
}

// MessageReceived is a free data retrieval call binding the contract method 0xe03555df.
//
// Solidity: function messageReceived(bytes32 originChainID, uint256 messageID) view returns(bool delivered)
func (_TeleportMessenger *TeleportMessengerCaller) MessageReceived(opts *bind.CallOpts, originChainID [32]byte, messageID *big.Int) (bool, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "messageReceived", originChainID, messageID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MessageReceived is a free data retrieval call binding the contract method 0xe03555df.
//
// Solidity: function messageReceived(bytes32 originChainID, uint256 messageID) view returns(bool delivered)
func (_TeleportMessenger *TeleportMessengerSession) MessageReceived(originChainID [32]byte, messageID *big.Int) (bool, error) {
	return _TeleportMessenger.Contract.MessageReceived(&_TeleportMessenger.CallOpts, originChainID, messageID)
}

// MessageReceived is a free data retrieval call binding the contract method 0xe03555df.
//
// Solidity: function messageReceived(bytes32 originChainID, uint256 messageID) view returns(bool delivered)
func (_TeleportMessenger *TeleportMessengerCallerSession) MessageReceived(originChainID [32]byte, messageID *big.Int) (bool, error) {
	return _TeleportMessenger.Contract.MessageReceived(&_TeleportMessenger.CallOpts, originChainID, messageID)
}

// OutstandingReceipts is a free data retrieval call binding the contract method 0x781f9744.
//
// Solidity: function outstandingReceipts(bytes32 ) view returns(address)
func (_TeleportMessenger *TeleportMessengerCaller) OutstandingReceipts(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "outstandingReceipts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OutstandingReceipts is a free data retrieval call binding the contract method 0x781f9744.
//
// Solidity: function outstandingReceipts(bytes32 ) view returns(address)
func (_TeleportMessenger *TeleportMessengerSession) OutstandingReceipts(arg0 [32]byte) (common.Address, error) {
	return _TeleportMessenger.Contract.OutstandingReceipts(&_TeleportMessenger.CallOpts, arg0)
}

// OutstandingReceipts is a free data retrieval call binding the contract method 0x781f9744.
//
// Solidity: function outstandingReceipts(bytes32 ) view returns(address)
func (_TeleportMessenger *TeleportMessengerCallerSession) OutstandingReceipts(arg0 [32]byte) (common.Address, error) {
	return _TeleportMessenger.Contract.OutstandingReceipts(&_TeleportMessenger.CallOpts, arg0)
}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0xc9bb1143.
//
// Solidity: function receivedFailedMessageHashes(bytes32 , uint256 ) view returns(bytes32)
func (_TeleportMessenger *TeleportMessengerCaller) ReceivedFailedMessageHashes(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "receivedFailedMessageHashes", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0xc9bb1143.
//
// Solidity: function receivedFailedMessageHashes(bytes32 , uint256 ) view returns(bytes32)
func (_TeleportMessenger *TeleportMessengerSession) ReceivedFailedMessageHashes(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _TeleportMessenger.Contract.ReceivedFailedMessageHashes(&_TeleportMessenger.CallOpts, arg0, arg1)
}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0xc9bb1143.
//
// Solidity: function receivedFailedMessageHashes(bytes32 , uint256 ) view returns(bytes32)
func (_TeleportMessenger *TeleportMessengerCallerSession) ReceivedFailedMessageHashes(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _TeleportMessenger.Contract.ReceivedFailedMessageHashes(&_TeleportMessenger.CallOpts, arg0, arg1)
}

// RelayerRewardAddresses is a free data retrieval call binding the contract method 0x21f18054.
//
// Solidity: function relayerRewardAddresses(bytes32 , uint256 ) view returns(address)
func (_TeleportMessenger *TeleportMessengerCaller) RelayerRewardAddresses(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "relayerRewardAddresses", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RelayerRewardAddresses is a free data retrieval call binding the contract method 0x21f18054.
//
// Solidity: function relayerRewardAddresses(bytes32 , uint256 ) view returns(address)
func (_TeleportMessenger *TeleportMessengerSession) RelayerRewardAddresses(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _TeleportMessenger.Contract.RelayerRewardAddresses(&_TeleportMessenger.CallOpts, arg0, arg1)
}

// RelayerRewardAddresses is a free data retrieval call binding the contract method 0x21f18054.
//
// Solidity: function relayerRewardAddresses(bytes32 , uint256 ) view returns(address)
func (_TeleportMessenger *TeleportMessengerCallerSession) RelayerRewardAddresses(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _TeleportMessenger.Contract.RelayerRewardAddresses(&_TeleportMessenger.CallOpts, arg0, arg1)
}

// RelayerRewardAmounts is a free data retrieval call binding the contract method 0x6192762c.
//
// Solidity: function relayerRewardAmounts(address , address ) view returns(uint256)
func (_TeleportMessenger *TeleportMessengerCaller) RelayerRewardAmounts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "relayerRewardAmounts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayerRewardAmounts is a free data retrieval call binding the contract method 0x6192762c.
//
// Solidity: function relayerRewardAmounts(address , address ) view returns(uint256)
func (_TeleportMessenger *TeleportMessengerSession) RelayerRewardAmounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TeleportMessenger.Contract.RelayerRewardAmounts(&_TeleportMessenger.CallOpts, arg0, arg1)
}

// RelayerRewardAmounts is a free data retrieval call binding the contract method 0x6192762c.
//
// Solidity: function relayerRewardAmounts(address , address ) view returns(uint256)
func (_TeleportMessenger *TeleportMessengerCallerSession) RelayerRewardAmounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TeleportMessenger.Contract.RelayerRewardAmounts(&_TeleportMessenger.CallOpts, arg0, arg1)
}

// SentMessageInfo is a free data retrieval call binding the contract method 0x66533d12.
//
// Solidity: function sentMessageInfo(bytes32 , uint256 ) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleportMessenger *TeleportMessengerCaller) SentMessageInfo(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	var out []interface{}
	err := _TeleportMessenger.contract.Call(opts, &out, "sentMessageInfo", arg0, arg1)

	outstruct := new(struct {
		MessageHash [32]byte
		FeeInfo     TeleporterFeeInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MessageHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.FeeInfo = *abi.ConvertType(out[1], new(TeleporterFeeInfo)).(*TeleporterFeeInfo)

	return *outstruct, err

}

// SentMessageInfo is a free data retrieval call binding the contract method 0x66533d12.
//
// Solidity: function sentMessageInfo(bytes32 , uint256 ) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleportMessenger *TeleportMessengerSession) SentMessageInfo(arg0 [32]byte, arg1 *big.Int) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	return _TeleportMessenger.Contract.SentMessageInfo(&_TeleportMessenger.CallOpts, arg0, arg1)
}

// SentMessageInfo is a free data retrieval call binding the contract method 0x66533d12.
//
// Solidity: function sentMessageInfo(bytes32 , uint256 ) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleportMessenger *TeleportMessengerCallerSession) SentMessageInfo(arg0 [32]byte, arg1 *big.Int) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	return _TeleportMessenger.Contract.SentMessageInfo(&_TeleportMessenger.CallOpts, arg0, arg1)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x19570c74.
//
// Solidity: function addFeeAmount(bytes32 destinationChainID, uint256 messageID, address feeContractAddress, uint256 additionalFeeAmount) returns()
func (_TeleportMessenger *TeleportMessengerTransactor) AddFeeAmount(opts *bind.TransactOpts, destinationChainID [32]byte, messageID *big.Int, feeContractAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleportMessenger.contract.Transact(opts, "addFeeAmount", destinationChainID, messageID, feeContractAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x19570c74.
//
// Solidity: function addFeeAmount(bytes32 destinationChainID, uint256 messageID, address feeContractAddress, uint256 additionalFeeAmount) returns()
func (_TeleportMessenger *TeleportMessengerSession) AddFeeAmount(destinationChainID [32]byte, messageID *big.Int, feeContractAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.AddFeeAmount(&_TeleportMessenger.TransactOpts, destinationChainID, messageID, feeContractAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x19570c74.
//
// Solidity: function addFeeAmount(bytes32 destinationChainID, uint256 messageID, address feeContractAddress, uint256 additionalFeeAmount) returns()
func (_TeleportMessenger *TeleportMessengerTransactorSession) AddFeeAmount(destinationChainID [32]byte, messageID *big.Int, feeContractAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.AddFeeAmount(&_TeleportMessenger.TransactOpts, destinationChainID, messageID, feeContractAddress, additionalFeeAmount)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0x10085181.
//
// Solidity: function receiveCrossChainMessage(address relayerRewardAddress) returns()
func (_TeleportMessenger *TeleportMessengerTransactor) ReceiveCrossChainMessage(opts *bind.TransactOpts, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleportMessenger.contract.Transact(opts, "receiveCrossChainMessage", relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0x10085181.
//
// Solidity: function receiveCrossChainMessage(address relayerRewardAddress) returns()
func (_TeleportMessenger *TeleportMessengerSession) ReceiveCrossChainMessage(relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.ReceiveCrossChainMessage(&_TeleportMessenger.TransactOpts, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0x10085181.
//
// Solidity: function receiveCrossChainMessage(address relayerRewardAddress) returns()
func (_TeleportMessenger *TeleportMessengerTransactorSession) ReceiveCrossChainMessage(relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.ReceiveCrossChainMessage(&_TeleportMessenger.TransactOpts, relayerRewardAddress)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleportMessenger *TeleportMessengerTransactor) RedeemRelayerRewards(opts *bind.TransactOpts, feeAsset common.Address) (*types.Transaction, error) {
	return _TeleportMessenger.contract.Transact(opts, "redeemRelayerRewards", feeAsset)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleportMessenger *TeleportMessengerSession) RedeemRelayerRewards(feeAsset common.Address) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.RedeemRelayerRewards(&_TeleportMessenger.TransactOpts, feeAsset)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleportMessenger *TeleportMessengerTransactorSession) RedeemRelayerRewards(feeAsset common.Address) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.RedeemRelayerRewards(&_TeleportMessenger.TransactOpts, feeAsset)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xcd3f2daa.
//
// Solidity: function retryMessageExecution(bytes32 originChainID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleportMessenger *TeleportMessengerTransactor) RetryMessageExecution(opts *bind.TransactOpts, originChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleportMessenger.contract.Transact(opts, "retryMessageExecution", originChainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xcd3f2daa.
//
// Solidity: function retryMessageExecution(bytes32 originChainID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleportMessenger *TeleportMessengerSession) RetryMessageExecution(originChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.RetryMessageExecution(&_TeleportMessenger.TransactOpts, originChainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xcd3f2daa.
//
// Solidity: function retryMessageExecution(bytes32 originChainID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleportMessenger *TeleportMessengerTransactorSession) RetryMessageExecution(originChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.RetryMessageExecution(&_TeleportMessenger.TransactOpts, originChainID, message)
}

// RetryReceipts is a paid mutator transaction binding the contract method 0x9a496900.
//
// Solidity: function retryReceipts(bytes32 originChainID, uint256[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(uint256 messageID)
func (_TeleportMessenger *TeleportMessengerTransactor) RetryReceipts(opts *bind.TransactOpts, originChainID [32]byte, messageIDs []*big.Int, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleportMessenger.contract.Transact(opts, "retryReceipts", originChainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// RetryReceipts is a paid mutator transaction binding the contract method 0x9a496900.
//
// Solidity: function retryReceipts(bytes32 originChainID, uint256[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(uint256 messageID)
func (_TeleportMessenger *TeleportMessengerSession) RetryReceipts(originChainID [32]byte, messageIDs []*big.Int, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.RetryReceipts(&_TeleportMessenger.TransactOpts, originChainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// RetryReceipts is a paid mutator transaction binding the contract method 0x9a496900.
//
// Solidity: function retryReceipts(bytes32 originChainID, uint256[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(uint256 messageID)
func (_TeleportMessenger *TeleportMessengerTransactorSession) RetryReceipts(originChainID [32]byte, messageIDs []*big.Int, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.RetryReceipts(&_TeleportMessenger.TransactOpts, originChainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x1af671f8.
//
// Solidity: function retrySendCrossChainMessage(bytes32 destinationChainID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleportMessenger *TeleportMessengerTransactor) RetrySendCrossChainMessage(opts *bind.TransactOpts, destinationChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleportMessenger.contract.Transact(opts, "retrySendCrossChainMessage", destinationChainID, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x1af671f8.
//
// Solidity: function retrySendCrossChainMessage(bytes32 destinationChainID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleportMessenger *TeleportMessengerSession) RetrySendCrossChainMessage(destinationChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.RetrySendCrossChainMessage(&_TeleportMessenger.TransactOpts, destinationChainID, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x1af671f8.
//
// Solidity: function retrySendCrossChainMessage(bytes32 destinationChainID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleportMessenger *TeleportMessengerTransactorSession) RetrySendCrossChainMessage(destinationChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.RetrySendCrossChainMessage(&_TeleportMessenger.TransactOpts, destinationChainID, message)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(uint256 messageID)
func (_TeleportMessenger *TeleportMessengerTransactor) SendCrossChainMessage(opts *bind.TransactOpts, messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleportMessenger.contract.Transact(opts, "sendCrossChainMessage", messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(uint256 messageID)
func (_TeleportMessenger *TeleportMessengerSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.SendCrossChainMessage(&_TeleportMessenger.TransactOpts, messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(uint256 messageID)
func (_TeleportMessenger *TeleportMessengerTransactorSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleportMessenger.Contract.SendCrossChainMessage(&_TeleportMessenger.TransactOpts, messageInput)
}

// TeleportMessengerAddFeeAmountIterator is returned from FilterAddFeeAmount and is used to iterate over the raw logs and unpacked data for AddFeeAmount events raised by the TeleportMessenger contract.
type TeleportMessengerAddFeeAmountIterator struct {
	Event *TeleportMessengerAddFeeAmount // Event containing the contract specifics and raw log

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
func (it *TeleportMessengerAddFeeAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleportMessengerAddFeeAmount)
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
		it.Event = new(TeleportMessengerAddFeeAmount)
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
func (it *TeleportMessengerAddFeeAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleportMessengerAddFeeAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleportMessengerAddFeeAmount represents a AddFeeAmount event raised by the TeleportMessenger contract.
type TeleportMessengerAddFeeAmount struct {
	DestinationChainID [32]byte
	MessageID          *big.Int
	UpdatedFeeInfo     TeleporterFeeInfo
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAddFeeAmount is a free log retrieval operation binding the contract event 0x28fe05eedf0479c9159e5b6dd2a28c93fa1a408eba22dc801fd9bc493a7fc0c2.
//
// Solidity: event AddFeeAmount(bytes32 indexed destinationChainID, uint256 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleportMessenger *TeleportMessengerFilterer) FilterAddFeeAmount(opts *bind.FilterOpts, destinationChainID [][32]byte, messageID []*big.Int) (*TeleportMessengerAddFeeAmountIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleportMessenger.contract.FilterLogs(opts, "AddFeeAmount", destinationChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleportMessengerAddFeeAmountIterator{contract: _TeleportMessenger.contract, event: "AddFeeAmount", logs: logs, sub: sub}, nil
}

// WatchAddFeeAmount is a free log subscription operation binding the contract event 0x28fe05eedf0479c9159e5b6dd2a28c93fa1a408eba22dc801fd9bc493a7fc0c2.
//
// Solidity: event AddFeeAmount(bytes32 indexed destinationChainID, uint256 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleportMessenger *TeleportMessengerFilterer) WatchAddFeeAmount(opts *bind.WatchOpts, sink chan<- *TeleportMessengerAddFeeAmount, destinationChainID [][32]byte, messageID []*big.Int) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleportMessenger.contract.WatchLogs(opts, "AddFeeAmount", destinationChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleportMessengerAddFeeAmount)
				if err := _TeleportMessenger.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
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

// ParseAddFeeAmount is a log parse operation binding the contract event 0x28fe05eedf0479c9159e5b6dd2a28c93fa1a408eba22dc801fd9bc493a7fc0c2.
//
// Solidity: event AddFeeAmount(bytes32 indexed destinationChainID, uint256 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleportMessenger *TeleportMessengerFilterer) ParseAddFeeAmount(log types.Log) (*TeleportMessengerAddFeeAmount, error) {
	event := new(TeleportMessengerAddFeeAmount)
	if err := _TeleportMessenger.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleportMessengerFailedMessageExecutionIterator is returned from FilterFailedMessageExecution and is used to iterate over the raw logs and unpacked data for FailedMessageExecution events raised by the TeleportMessenger contract.
type TeleportMessengerFailedMessageExecutionIterator struct {
	Event *TeleportMessengerFailedMessageExecution // Event containing the contract specifics and raw log

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
func (it *TeleportMessengerFailedMessageExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleportMessengerFailedMessageExecution)
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
		it.Event = new(TeleportMessengerFailedMessageExecution)
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
func (it *TeleportMessengerFailedMessageExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleportMessengerFailedMessageExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleportMessengerFailedMessageExecution represents a FailedMessageExecution event raised by the TeleportMessenger contract.
type TeleportMessengerFailedMessageExecution struct {
	OriginChainID [32]byte
	MessageID     *big.Int
	Message       TeleporterMessage
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFailedMessageExecution is a free log retrieval operation binding the contract event 0x50e5f3de5b0b3e6c82553b7e0f5f7080a291be07aa30a311084a9457f4a956bc.
//
// Solidity: event FailedMessageExecution(bytes32 indexed originChainID, uint256 indexed messageID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleportMessenger *TeleportMessengerFilterer) FilterFailedMessageExecution(opts *bind.FilterOpts, originChainID [][32]byte, messageID []*big.Int) (*TeleportMessengerFailedMessageExecutionIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleportMessenger.contract.FilterLogs(opts, "FailedMessageExecution", originChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleportMessengerFailedMessageExecutionIterator{contract: _TeleportMessenger.contract, event: "FailedMessageExecution", logs: logs, sub: sub}, nil
}

// WatchFailedMessageExecution is a free log subscription operation binding the contract event 0x50e5f3de5b0b3e6c82553b7e0f5f7080a291be07aa30a311084a9457f4a956bc.
//
// Solidity: event FailedMessageExecution(bytes32 indexed originChainID, uint256 indexed messageID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleportMessenger *TeleportMessengerFilterer) WatchFailedMessageExecution(opts *bind.WatchOpts, sink chan<- *TeleportMessengerFailedMessageExecution, originChainID [][32]byte, messageID []*big.Int) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleportMessenger.contract.WatchLogs(opts, "FailedMessageExecution", originChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleportMessengerFailedMessageExecution)
				if err := _TeleportMessenger.contract.UnpackLog(event, "FailedMessageExecution", log); err != nil {
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

// ParseFailedMessageExecution is a log parse operation binding the contract event 0x50e5f3de5b0b3e6c82553b7e0f5f7080a291be07aa30a311084a9457f4a956bc.
//
// Solidity: event FailedMessageExecution(bytes32 indexed originChainID, uint256 indexed messageID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleportMessenger *TeleportMessengerFilterer) ParseFailedMessageExecution(log types.Log) (*TeleportMessengerFailedMessageExecution, error) {
	event := new(TeleportMessengerFailedMessageExecution)
	if err := _TeleportMessenger.contract.UnpackLog(event, "FailedMessageExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleportMessengerMessageExecutionRetriedIterator is returned from FilterMessageExecutionRetried and is used to iterate over the raw logs and unpacked data for MessageExecutionRetried events raised by the TeleportMessenger contract.
type TeleportMessengerMessageExecutionRetriedIterator struct {
	Event *TeleportMessengerMessageExecutionRetried // Event containing the contract specifics and raw log

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
func (it *TeleportMessengerMessageExecutionRetriedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleportMessengerMessageExecutionRetried)
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
		it.Event = new(TeleportMessengerMessageExecutionRetried)
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
func (it *TeleportMessengerMessageExecutionRetriedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleportMessengerMessageExecutionRetriedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleportMessengerMessageExecutionRetried represents a MessageExecutionRetried event raised by the TeleportMessenger contract.
type TeleportMessengerMessageExecutionRetried struct {
	OriginChainID [32]byte
	MessageID     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMessageExecutionRetried is a free log retrieval operation binding the contract event 0x3d5f30e93c1e27cda0e05a7b9e51144613a816cd90561f8493393bbcf4e00358.
//
// Solidity: event MessageExecutionRetried(bytes32 indexed originChainID, uint256 indexed messageID)
func (_TeleportMessenger *TeleportMessengerFilterer) FilterMessageExecutionRetried(opts *bind.FilterOpts, originChainID [][32]byte, messageID []*big.Int) (*TeleportMessengerMessageExecutionRetriedIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleportMessenger.contract.FilterLogs(opts, "MessageExecutionRetried", originChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleportMessengerMessageExecutionRetriedIterator{contract: _TeleportMessenger.contract, event: "MessageExecutionRetried", logs: logs, sub: sub}, nil
}

// WatchMessageExecutionRetried is a free log subscription operation binding the contract event 0x3d5f30e93c1e27cda0e05a7b9e51144613a816cd90561f8493393bbcf4e00358.
//
// Solidity: event MessageExecutionRetried(bytes32 indexed originChainID, uint256 indexed messageID)
func (_TeleportMessenger *TeleportMessengerFilterer) WatchMessageExecutionRetried(opts *bind.WatchOpts, sink chan<- *TeleportMessengerMessageExecutionRetried, originChainID [][32]byte, messageID []*big.Int) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleportMessenger.contract.WatchLogs(opts, "MessageExecutionRetried", originChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleportMessengerMessageExecutionRetried)
				if err := _TeleportMessenger.contract.UnpackLog(event, "MessageExecutionRetried", log); err != nil {
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

// ParseMessageExecutionRetried is a log parse operation binding the contract event 0x3d5f30e93c1e27cda0e05a7b9e51144613a816cd90561f8493393bbcf4e00358.
//
// Solidity: event MessageExecutionRetried(bytes32 indexed originChainID, uint256 indexed messageID)
func (_TeleportMessenger *TeleportMessengerFilterer) ParseMessageExecutionRetried(log types.Log) (*TeleportMessengerMessageExecutionRetried, error) {
	event := new(TeleportMessengerMessageExecutionRetried)
	if err := _TeleportMessenger.contract.UnpackLog(event, "MessageExecutionRetried", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleportMessengerReceiveCrossChainMessageIterator is returned from FilterReceiveCrossChainMessage and is used to iterate over the raw logs and unpacked data for ReceiveCrossChainMessage events raised by the TeleportMessenger contract.
type TeleportMessengerReceiveCrossChainMessageIterator struct {
	Event *TeleportMessengerReceiveCrossChainMessage // Event containing the contract specifics and raw log

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
func (it *TeleportMessengerReceiveCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleportMessengerReceiveCrossChainMessage)
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
		it.Event = new(TeleportMessengerReceiveCrossChainMessage)
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
func (it *TeleportMessengerReceiveCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleportMessengerReceiveCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleportMessengerReceiveCrossChainMessage represents a ReceiveCrossChainMessage event raised by the TeleportMessenger contract.
type TeleportMessengerReceiveCrossChainMessage struct {
	OriginChainID [32]byte
	MessageID     *big.Int
	Message       TeleporterMessage
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceiveCrossChainMessage is a free log retrieval operation binding the contract event 0x522ce4da81e9fb994bf6a122282939647b6e69817be443f7d1c152e76a082cc7.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed originChainID, uint256 indexed messageID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleportMessenger *TeleportMessengerFilterer) FilterReceiveCrossChainMessage(opts *bind.FilterOpts, originChainID [][32]byte, messageID []*big.Int) (*TeleportMessengerReceiveCrossChainMessageIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleportMessenger.contract.FilterLogs(opts, "ReceiveCrossChainMessage", originChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleportMessengerReceiveCrossChainMessageIterator{contract: _TeleportMessenger.contract, event: "ReceiveCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchReceiveCrossChainMessage is a free log subscription operation binding the contract event 0x522ce4da81e9fb994bf6a122282939647b6e69817be443f7d1c152e76a082cc7.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed originChainID, uint256 indexed messageID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleportMessenger *TeleportMessengerFilterer) WatchReceiveCrossChainMessage(opts *bind.WatchOpts, sink chan<- *TeleportMessengerReceiveCrossChainMessage, originChainID [][32]byte, messageID []*big.Int) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleportMessenger.contract.WatchLogs(opts, "ReceiveCrossChainMessage", originChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleportMessengerReceiveCrossChainMessage)
				if err := _TeleportMessenger.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
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

// ParseReceiveCrossChainMessage is a log parse operation binding the contract event 0x522ce4da81e9fb994bf6a122282939647b6e69817be443f7d1c152e76a082cc7.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed originChainID, uint256 indexed messageID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleportMessenger *TeleportMessengerFilterer) ParseReceiveCrossChainMessage(log types.Log) (*TeleportMessengerReceiveCrossChainMessage, error) {
	event := new(TeleportMessengerReceiveCrossChainMessage)
	if err := _TeleportMessenger.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleportMessengerSendCrossChainMessageIterator is returned from FilterSendCrossChainMessage and is used to iterate over the raw logs and unpacked data for SendCrossChainMessage events raised by the TeleportMessenger contract.
type TeleportMessengerSendCrossChainMessageIterator struct {
	Event *TeleportMessengerSendCrossChainMessage // Event containing the contract specifics and raw log

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
func (it *TeleportMessengerSendCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleportMessengerSendCrossChainMessage)
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
		it.Event = new(TeleportMessengerSendCrossChainMessage)
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
func (it *TeleportMessengerSendCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleportMessengerSendCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleportMessengerSendCrossChainMessage represents a SendCrossChainMessage event raised by the TeleportMessenger contract.
type TeleportMessengerSendCrossChainMessage struct {
	DestinationChainID [32]byte
	MessageID          *big.Int
	Message            TeleporterMessage
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSendCrossChainMessage is a free log retrieval operation binding the contract event 0x7491aecf1f3e24837ce48fd97fc8729fc036cebb3e5078643f3301b72852aa07.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed destinationChainID, uint256 indexed messageID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleportMessenger *TeleportMessengerFilterer) FilterSendCrossChainMessage(opts *bind.FilterOpts, destinationChainID [][32]byte, messageID []*big.Int) (*TeleportMessengerSendCrossChainMessageIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleportMessenger.contract.FilterLogs(opts, "SendCrossChainMessage", destinationChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleportMessengerSendCrossChainMessageIterator{contract: _TeleportMessenger.contract, event: "SendCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchSendCrossChainMessage is a free log subscription operation binding the contract event 0x7491aecf1f3e24837ce48fd97fc8729fc036cebb3e5078643f3301b72852aa07.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed destinationChainID, uint256 indexed messageID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleportMessenger *TeleportMessengerFilterer) WatchSendCrossChainMessage(opts *bind.WatchOpts, sink chan<- *TeleportMessengerSendCrossChainMessage, destinationChainID [][32]byte, messageID []*big.Int) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleportMessenger.contract.WatchLogs(opts, "SendCrossChainMessage", destinationChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleportMessengerSendCrossChainMessage)
				if err := _TeleportMessenger.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
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

// ParseSendCrossChainMessage is a log parse operation binding the contract event 0x7491aecf1f3e24837ce48fd97fc8729fc036cebb3e5078643f3301b72852aa07.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed destinationChainID, uint256 indexed messageID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleportMessenger *TeleportMessengerFilterer) ParseSendCrossChainMessage(log types.Log) (*TeleportMessengerSendCrossChainMessage, error) {
	event := new(TeleportMessengerSendCrossChainMessage)
	if err := _TeleportMessenger.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
