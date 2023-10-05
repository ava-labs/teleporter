// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teleportermessenger

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

// TeleportermessengerMetaData contains all meta data concerning the Teleportermessenger contract.
var TeleportermessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BalanceNotIncreased\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdditionalFeeAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDestinationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDestinationChainID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFeeAssetContractAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMessageHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOriginSenderAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRelayerRewardAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWarpMessage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageAlreadyDelivered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageRetryExecutionFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoRelayerRewardToRedeem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiptNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverReentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderReentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedRelayer\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"updatedFeeInfo\",\"type\":\"tuple\"}],\"name\":\"AddFeeAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"FailedMessageExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"MessageExecutionRetried\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ReceiveCrossChainMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"SendCrossChainMessage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAXIMUM_RECEIPT_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_REQUIRED_CALL_DATA_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REQUIRED_ORIGIN_CHAIN_ID_START_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalFeeAmount\",\"type\":\"uint256\"}],\"name\":\"addFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delivererAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayers\",\"type\":\"address[]\"}],\"name\":\"checkIsAllowedRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"checkRelayerRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"getFeeInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chainID\",\"type\":\"bytes32\"}],\"name\":\"getNextMessageID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"getRelayerRewardAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"latestMessageIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"messageReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"delivered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"outstandingReceipts\",\"outputs\":[{\"internalType\":\"contractReceiptQueue\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"name\":\"receiveCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"receivedFailedMessageHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"redeemRelayerRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"relayerRewardAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"relayerRewardAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retryMessageExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"messageIDs\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"retryReceipts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retrySendCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageInput\",\"name\":\"messageInput\",\"type\":\"tuple\"}],\"name\":\"sendCrossChainMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sentMessageInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TeleportermessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleportermessengerMetaData.ABI instead.
var TeleportermessengerABI = TeleportermessengerMetaData.ABI

// Teleportermessenger is an auto generated Go binding around an Ethereum contract.
type Teleportermessenger struct {
	TeleportermessengerCaller     // Read-only binding to the contract
	TeleportermessengerTransactor // Write-only binding to the contract
	TeleportermessengerFilterer   // Log filterer for contract events
}

// TeleportermessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeleportermessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleportermessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeleportermessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleportermessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeleportermessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleportermessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeleportermessengerSession struct {
	Contract     *Teleportermessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TeleportermessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeleportermessengerCallerSession struct {
	Contract *TeleportermessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TeleportermessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeleportermessengerTransactorSession struct {
	Contract     *TeleportermessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TeleportermessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeleportermessengerRaw struct {
	Contract *Teleportermessenger // Generic contract binding to access the raw methods on
}

// TeleportermessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeleportermessengerCallerRaw struct {
	Contract *TeleportermessengerCaller // Generic read-only contract binding to access the raw methods on
}

// TeleportermessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeleportermessengerTransactorRaw struct {
	Contract *TeleportermessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeleportermessenger creates a new instance of Teleportermessenger, bound to a specific deployed contract.
func NewTeleportermessenger(address common.Address, backend bind.ContractBackend) (*Teleportermessenger, error) {
	contract, err := bindTeleportermessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Teleportermessenger{TeleportermessengerCaller: TeleportermessengerCaller{contract: contract}, TeleportermessengerTransactor: TeleportermessengerTransactor{contract: contract}, TeleportermessengerFilterer: TeleportermessengerFilterer{contract: contract}}, nil
}

// NewTeleportermessengerCaller creates a new read-only instance of Teleportermessenger, bound to a specific deployed contract.
func NewTeleportermessengerCaller(address common.Address, caller bind.ContractCaller) (*TeleportermessengerCaller, error) {
	contract, err := bindTeleportermessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeleportermessengerCaller{contract: contract}, nil
}

// NewTeleportermessengerTransactor creates a new write-only instance of Teleportermessenger, bound to a specific deployed contract.
func NewTeleportermessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*TeleportermessengerTransactor, error) {
	contract, err := bindTeleportermessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeleportermessengerTransactor{contract: contract}, nil
}

// NewTeleportermessengerFilterer creates a new log filterer instance of Teleportermessenger, bound to a specific deployed contract.
func NewTeleportermessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*TeleportermessengerFilterer, error) {
	contract, err := bindTeleportermessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeleportermessengerFilterer{contract: contract}, nil
}

// bindTeleportermessenger binds a generic wrapper to an already deployed contract.
func bindTeleportermessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeleportermessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Teleportermessenger *TeleportermessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Teleportermessenger.Contract.TeleportermessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Teleportermessenger *TeleportermessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.TeleportermessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Teleportermessenger *TeleportermessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.TeleportermessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Teleportermessenger *TeleportermessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Teleportermessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Teleportermessenger *TeleportermessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Teleportermessenger *TeleportermessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.contract.Transact(opts, method, params...)
}

// MAXIMUMRECEIPTCOUNT is a free data retrieval call binding the contract method 0x10534371.
//
// Solidity: function MAXIMUM_RECEIPT_COUNT() view returns(uint256)
func (_Teleportermessenger *TeleportermessengerCaller) MAXIMUMRECEIPTCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "MAXIMUM_RECEIPT_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXIMUMRECEIPTCOUNT is a free data retrieval call binding the contract method 0x10534371.
//
// Solidity: function MAXIMUM_RECEIPT_COUNT() view returns(uint256)
func (_Teleportermessenger *TeleportermessengerSession) MAXIMUMRECEIPTCOUNT() (*big.Int, error) {
	return _Teleportermessenger.Contract.MAXIMUMRECEIPTCOUNT(&_Teleportermessenger.CallOpts)
}

// MAXIMUMRECEIPTCOUNT is a free data retrieval call binding the contract method 0x10534371.
//
// Solidity: function MAXIMUM_RECEIPT_COUNT() view returns(uint256)
func (_Teleportermessenger *TeleportermessengerCallerSession) MAXIMUMRECEIPTCOUNT() (*big.Int, error) {
	return _Teleportermessenger.Contract.MAXIMUMRECEIPTCOUNT(&_Teleportermessenger.CallOpts)
}

// MINIMUMREQUIREDCALLDATALENGTH is a free data retrieval call binding the contract method 0x8f12376f.
//
// Solidity: function MINIMUM_REQUIRED_CALL_DATA_LENGTH() view returns(uint256)
func (_Teleportermessenger *TeleportermessengerCaller) MINIMUMREQUIREDCALLDATALENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "MINIMUM_REQUIRED_CALL_DATA_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMREQUIREDCALLDATALENGTH is a free data retrieval call binding the contract method 0x8f12376f.
//
// Solidity: function MINIMUM_REQUIRED_CALL_DATA_LENGTH() view returns(uint256)
func (_Teleportermessenger *TeleportermessengerSession) MINIMUMREQUIREDCALLDATALENGTH() (*big.Int, error) {
	return _Teleportermessenger.Contract.MINIMUMREQUIREDCALLDATALENGTH(&_Teleportermessenger.CallOpts)
}

// MINIMUMREQUIREDCALLDATALENGTH is a free data retrieval call binding the contract method 0x8f12376f.
//
// Solidity: function MINIMUM_REQUIRED_CALL_DATA_LENGTH() view returns(uint256)
func (_Teleportermessenger *TeleportermessengerCallerSession) MINIMUMREQUIREDCALLDATALENGTH() (*big.Int, error) {
	return _Teleportermessenger.Contract.MINIMUMREQUIREDCALLDATALENGTH(&_Teleportermessenger.CallOpts)
}

// REQUIREDORIGINCHAINIDSTARTINDEX is a free data retrieval call binding the contract method 0x5bf91119.
//
// Solidity: function REQUIRED_ORIGIN_CHAIN_ID_START_INDEX() view returns(uint256)
func (_Teleportermessenger *TeleportermessengerCaller) REQUIREDORIGINCHAINIDSTARTINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "REQUIRED_ORIGIN_CHAIN_ID_START_INDEX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REQUIREDORIGINCHAINIDSTARTINDEX is a free data retrieval call binding the contract method 0x5bf91119.
//
// Solidity: function REQUIRED_ORIGIN_CHAIN_ID_START_INDEX() view returns(uint256)
func (_Teleportermessenger *TeleportermessengerSession) REQUIREDORIGINCHAINIDSTARTINDEX() (*big.Int, error) {
	return _Teleportermessenger.Contract.REQUIREDORIGINCHAINIDSTARTINDEX(&_Teleportermessenger.CallOpts)
}

// REQUIREDORIGINCHAINIDSTARTINDEX is a free data retrieval call binding the contract method 0x5bf91119.
//
// Solidity: function REQUIRED_ORIGIN_CHAIN_ID_START_INDEX() view returns(uint256)
func (_Teleportermessenger *TeleportermessengerCallerSession) REQUIREDORIGINCHAINIDSTARTINDEX() (*big.Int, error) {
	return _Teleportermessenger.Contract.REQUIREDORIGINCHAINIDSTARTINDEX(&_Teleportermessenger.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_Teleportermessenger *TeleportermessengerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_Teleportermessenger *TeleportermessengerSession) WARPMESSENGER() (common.Address, error) {
	return _Teleportermessenger.Contract.WARPMESSENGER(&_Teleportermessenger.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_Teleportermessenger *TeleportermessengerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _Teleportermessenger.Contract.WARPMESSENGER(&_Teleportermessenger.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_Teleportermessenger *TeleportermessengerCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_Teleportermessenger *TeleportermessengerSession) BlockchainID() ([32]byte, error) {
	return _Teleportermessenger.Contract.BlockchainID(&_Teleportermessenger.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_Teleportermessenger *TeleportermessengerCallerSession) BlockchainID() ([32]byte, error) {
	return _Teleportermessenger.Contract.BlockchainID(&_Teleportermessenger.CallOpts)
}

// CheckIsAllowedRelayer is a free data retrieval call binding the contract method 0x65171908.
//
// Solidity: function checkIsAllowedRelayer(address delivererAddress, address[] allowedRelayers) pure returns(bool)
func (_Teleportermessenger *TeleportermessengerCaller) CheckIsAllowedRelayer(opts *bind.CallOpts, delivererAddress common.Address, allowedRelayers []common.Address) (bool, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "checkIsAllowedRelayer", delivererAddress, allowedRelayers)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckIsAllowedRelayer is a free data retrieval call binding the contract method 0x65171908.
//
// Solidity: function checkIsAllowedRelayer(address delivererAddress, address[] allowedRelayers) pure returns(bool)
func (_Teleportermessenger *TeleportermessengerSession) CheckIsAllowedRelayer(delivererAddress common.Address, allowedRelayers []common.Address) (bool, error) {
	return _Teleportermessenger.Contract.CheckIsAllowedRelayer(&_Teleportermessenger.CallOpts, delivererAddress, allowedRelayers)
}

// CheckIsAllowedRelayer is a free data retrieval call binding the contract method 0x65171908.
//
// Solidity: function checkIsAllowedRelayer(address delivererAddress, address[] allowedRelayers) pure returns(bool)
func (_Teleportermessenger *TeleportermessengerCallerSession) CheckIsAllowedRelayer(delivererAddress common.Address, allowedRelayers []common.Address) (bool, error) {
	return _Teleportermessenger.Contract.CheckIsAllowedRelayer(&_Teleportermessenger.CallOpts, delivererAddress, allowedRelayers)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_Teleportermessenger *TeleportermessengerCaller) CheckRelayerRewardAmount(opts *bind.CallOpts, relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "checkRelayerRewardAmount", relayer, feeAsset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_Teleportermessenger *TeleportermessengerSession) CheckRelayerRewardAmount(relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	return _Teleportermessenger.Contract.CheckRelayerRewardAmount(&_Teleportermessenger.CallOpts, relayer, feeAsset)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_Teleportermessenger *TeleportermessengerCallerSession) CheckRelayerRewardAmount(relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	return _Teleportermessenger.Contract.CheckRelayerRewardAmount(&_Teleportermessenger.CallOpts, relayer, feeAsset)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0x82f2c43a.
//
// Solidity: function getFeeInfo(bytes32 destinationChainID, uint256 messageID) view returns(address feeAsset, uint256 feeAmount)
func (_Teleportermessenger *TeleportermessengerCaller) GetFeeInfo(opts *bind.CallOpts, destinationChainID [32]byte, messageID *big.Int) (struct {
	FeeAsset  common.Address
	FeeAmount *big.Int
}, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "getFeeInfo", destinationChainID, messageID)

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
func (_Teleportermessenger *TeleportermessengerSession) GetFeeInfo(destinationChainID [32]byte, messageID *big.Int) (struct {
	FeeAsset  common.Address
	FeeAmount *big.Int
}, error) {
	return _Teleportermessenger.Contract.GetFeeInfo(&_Teleportermessenger.CallOpts, destinationChainID, messageID)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0x82f2c43a.
//
// Solidity: function getFeeInfo(bytes32 destinationChainID, uint256 messageID) view returns(address feeAsset, uint256 feeAmount)
func (_Teleportermessenger *TeleportermessengerCallerSession) GetFeeInfo(destinationChainID [32]byte, messageID *big.Int) (struct {
	FeeAsset  common.Address
	FeeAmount *big.Int
}, error) {
	return _Teleportermessenger.Contract.GetFeeInfo(&_Teleportermessenger.CallOpts, destinationChainID, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x220c9568.
//
// Solidity: function getMessageHash(bytes32 destinationChainID, uint256 messageID) view returns(bytes32 messageHash)
func (_Teleportermessenger *TeleportermessengerCaller) GetMessageHash(opts *bind.CallOpts, destinationChainID [32]byte, messageID *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "getMessageHash", destinationChainID, messageID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x220c9568.
//
// Solidity: function getMessageHash(bytes32 destinationChainID, uint256 messageID) view returns(bytes32 messageHash)
func (_Teleportermessenger *TeleportermessengerSession) GetMessageHash(destinationChainID [32]byte, messageID *big.Int) ([32]byte, error) {
	return _Teleportermessenger.Contract.GetMessageHash(&_Teleportermessenger.CallOpts, destinationChainID, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x220c9568.
//
// Solidity: function getMessageHash(bytes32 destinationChainID, uint256 messageID) view returns(bytes32 messageHash)
func (_Teleportermessenger *TeleportermessengerCallerSession) GetMessageHash(destinationChainID [32]byte, messageID *big.Int) ([32]byte, error) {
	return _Teleportermessenger.Contract.GetMessageHash(&_Teleportermessenger.CallOpts, destinationChainID, messageID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 chainID) view returns(uint256 messageID)
func (_Teleportermessenger *TeleportermessengerCaller) GetNextMessageID(opts *bind.CallOpts, chainID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "getNextMessageID", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 chainID) view returns(uint256 messageID)
func (_Teleportermessenger *TeleportermessengerSession) GetNextMessageID(chainID [32]byte) (*big.Int, error) {
	return _Teleportermessenger.Contract.GetNextMessageID(&_Teleportermessenger.CallOpts, chainID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 chainID) view returns(uint256 messageID)
func (_Teleportermessenger *TeleportermessengerCallerSession) GetNextMessageID(chainID [32]byte) (*big.Int, error) {
	return _Teleportermessenger.Contract.GetNextMessageID(&_Teleportermessenger.CallOpts, chainID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x33e890fe.
//
// Solidity: function getRelayerRewardAddress(bytes32 originChainID, uint256 messageID) view returns(address relayerRewardAddress)
func (_Teleportermessenger *TeleportermessengerCaller) GetRelayerRewardAddress(opts *bind.CallOpts, originChainID [32]byte, messageID *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "getRelayerRewardAddress", originChainID, messageID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x33e890fe.
//
// Solidity: function getRelayerRewardAddress(bytes32 originChainID, uint256 messageID) view returns(address relayerRewardAddress)
func (_Teleportermessenger *TeleportermessengerSession) GetRelayerRewardAddress(originChainID [32]byte, messageID *big.Int) (common.Address, error) {
	return _Teleportermessenger.Contract.GetRelayerRewardAddress(&_Teleportermessenger.CallOpts, originChainID, messageID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x33e890fe.
//
// Solidity: function getRelayerRewardAddress(bytes32 originChainID, uint256 messageID) view returns(address relayerRewardAddress)
func (_Teleportermessenger *TeleportermessengerCallerSession) GetRelayerRewardAddress(originChainID [32]byte, messageID *big.Int) (common.Address, error) {
	return _Teleportermessenger.Contract.GetRelayerRewardAddress(&_Teleportermessenger.CallOpts, originChainID, messageID)
}

// LatestMessageIDs is a free data retrieval call binding the contract method 0x29ec9beb.
//
// Solidity: function latestMessageIDs(bytes32 ) view returns(uint256)
func (_Teleportermessenger *TeleportermessengerCaller) LatestMessageIDs(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "latestMessageIDs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestMessageIDs is a free data retrieval call binding the contract method 0x29ec9beb.
//
// Solidity: function latestMessageIDs(bytes32 ) view returns(uint256)
func (_Teleportermessenger *TeleportermessengerSession) LatestMessageIDs(arg0 [32]byte) (*big.Int, error) {
	return _Teleportermessenger.Contract.LatestMessageIDs(&_Teleportermessenger.CallOpts, arg0)
}

// LatestMessageIDs is a free data retrieval call binding the contract method 0x29ec9beb.
//
// Solidity: function latestMessageIDs(bytes32 ) view returns(uint256)
func (_Teleportermessenger *TeleportermessengerCallerSession) LatestMessageIDs(arg0 [32]byte) (*big.Int, error) {
	return _Teleportermessenger.Contract.LatestMessageIDs(&_Teleportermessenger.CallOpts, arg0)
}

// MessageReceived is a free data retrieval call binding the contract method 0xe03555df.
//
// Solidity: function messageReceived(bytes32 originChainID, uint256 messageID) view returns(bool delivered)
func (_Teleportermessenger *TeleportermessengerCaller) MessageReceived(opts *bind.CallOpts, originChainID [32]byte, messageID *big.Int) (bool, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "messageReceived", originChainID, messageID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MessageReceived is a free data retrieval call binding the contract method 0xe03555df.
//
// Solidity: function messageReceived(bytes32 originChainID, uint256 messageID) view returns(bool delivered)
func (_Teleportermessenger *TeleportermessengerSession) MessageReceived(originChainID [32]byte, messageID *big.Int) (bool, error) {
	return _Teleportermessenger.Contract.MessageReceived(&_Teleportermessenger.CallOpts, originChainID, messageID)
}

// MessageReceived is a free data retrieval call binding the contract method 0xe03555df.
//
// Solidity: function messageReceived(bytes32 originChainID, uint256 messageID) view returns(bool delivered)
func (_Teleportermessenger *TeleportermessengerCallerSession) MessageReceived(originChainID [32]byte, messageID *big.Int) (bool, error) {
	return _Teleportermessenger.Contract.MessageReceived(&_Teleportermessenger.CallOpts, originChainID, messageID)
}

// OutstandingReceipts is a free data retrieval call binding the contract method 0x781f9744.
//
// Solidity: function outstandingReceipts(bytes32 ) view returns(address)
func (_Teleportermessenger *TeleportermessengerCaller) OutstandingReceipts(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "outstandingReceipts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OutstandingReceipts is a free data retrieval call binding the contract method 0x781f9744.
//
// Solidity: function outstandingReceipts(bytes32 ) view returns(address)
func (_Teleportermessenger *TeleportermessengerSession) OutstandingReceipts(arg0 [32]byte) (common.Address, error) {
	return _Teleportermessenger.Contract.OutstandingReceipts(&_Teleportermessenger.CallOpts, arg0)
}

// OutstandingReceipts is a free data retrieval call binding the contract method 0x781f9744.
//
// Solidity: function outstandingReceipts(bytes32 ) view returns(address)
func (_Teleportermessenger *TeleportermessengerCallerSession) OutstandingReceipts(arg0 [32]byte) (common.Address, error) {
	return _Teleportermessenger.Contract.OutstandingReceipts(&_Teleportermessenger.CallOpts, arg0)
}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0xc9bb1143.
//
// Solidity: function receivedFailedMessageHashes(bytes32 , uint256 ) view returns(bytes32)
func (_Teleportermessenger *TeleportermessengerCaller) ReceivedFailedMessageHashes(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "receivedFailedMessageHashes", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0xc9bb1143.
//
// Solidity: function receivedFailedMessageHashes(bytes32 , uint256 ) view returns(bytes32)
func (_Teleportermessenger *TeleportermessengerSession) ReceivedFailedMessageHashes(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _Teleportermessenger.Contract.ReceivedFailedMessageHashes(&_Teleportermessenger.CallOpts, arg0, arg1)
}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0xc9bb1143.
//
// Solidity: function receivedFailedMessageHashes(bytes32 , uint256 ) view returns(bytes32)
func (_Teleportermessenger *TeleportermessengerCallerSession) ReceivedFailedMessageHashes(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _Teleportermessenger.Contract.ReceivedFailedMessageHashes(&_Teleportermessenger.CallOpts, arg0, arg1)
}

// RelayerRewardAddresses is a free data retrieval call binding the contract method 0x21f18054.
//
// Solidity: function relayerRewardAddresses(bytes32 , uint256 ) view returns(address)
func (_Teleportermessenger *TeleportermessengerCaller) RelayerRewardAddresses(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "relayerRewardAddresses", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RelayerRewardAddresses is a free data retrieval call binding the contract method 0x21f18054.
//
// Solidity: function relayerRewardAddresses(bytes32 , uint256 ) view returns(address)
func (_Teleportermessenger *TeleportermessengerSession) RelayerRewardAddresses(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _Teleportermessenger.Contract.RelayerRewardAddresses(&_Teleportermessenger.CallOpts, arg0, arg1)
}

// RelayerRewardAddresses is a free data retrieval call binding the contract method 0x21f18054.
//
// Solidity: function relayerRewardAddresses(bytes32 , uint256 ) view returns(address)
func (_Teleportermessenger *TeleportermessengerCallerSession) RelayerRewardAddresses(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _Teleportermessenger.Contract.RelayerRewardAddresses(&_Teleportermessenger.CallOpts, arg0, arg1)
}

// RelayerRewardAmounts is a free data retrieval call binding the contract method 0x6192762c.
//
// Solidity: function relayerRewardAmounts(address , address ) view returns(uint256)
func (_Teleportermessenger *TeleportermessengerCaller) RelayerRewardAmounts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "relayerRewardAmounts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayerRewardAmounts is a free data retrieval call binding the contract method 0x6192762c.
//
// Solidity: function relayerRewardAmounts(address , address ) view returns(uint256)
func (_Teleportermessenger *TeleportermessengerSession) RelayerRewardAmounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Teleportermessenger.Contract.RelayerRewardAmounts(&_Teleportermessenger.CallOpts, arg0, arg1)
}

// RelayerRewardAmounts is a free data retrieval call binding the contract method 0x6192762c.
//
// Solidity: function relayerRewardAmounts(address , address ) view returns(uint256)
func (_Teleportermessenger *TeleportermessengerCallerSession) RelayerRewardAmounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Teleportermessenger.Contract.RelayerRewardAmounts(&_Teleportermessenger.CallOpts, arg0, arg1)
}

// SentMessageInfo is a free data retrieval call binding the contract method 0x66533d12.
//
// Solidity: function sentMessageInfo(bytes32 , uint256 ) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_Teleportermessenger *TeleportermessengerCaller) SentMessageInfo(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	var out []interface{}
	err := _Teleportermessenger.contract.Call(opts, &out, "sentMessageInfo", arg0, arg1)

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
func (_Teleportermessenger *TeleportermessengerSession) SentMessageInfo(arg0 [32]byte, arg1 *big.Int) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	return _Teleportermessenger.Contract.SentMessageInfo(&_Teleportermessenger.CallOpts, arg0, arg1)
}

// SentMessageInfo is a free data retrieval call binding the contract method 0x66533d12.
//
// Solidity: function sentMessageInfo(bytes32 , uint256 ) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_Teleportermessenger *TeleportermessengerCallerSession) SentMessageInfo(arg0 [32]byte, arg1 *big.Int) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	return _Teleportermessenger.Contract.SentMessageInfo(&_Teleportermessenger.CallOpts, arg0, arg1)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x19570c74.
//
// Solidity: function addFeeAmount(bytes32 destinationChainID, uint256 messageID, address feeContractAddress, uint256 additionalFeeAmount) returns()
func (_Teleportermessenger *TeleportermessengerTransactor) AddFeeAmount(opts *bind.TransactOpts, destinationChainID [32]byte, messageID *big.Int, feeContractAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _Teleportermessenger.contract.Transact(opts, "addFeeAmount", destinationChainID, messageID, feeContractAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x19570c74.
//
// Solidity: function addFeeAmount(bytes32 destinationChainID, uint256 messageID, address feeContractAddress, uint256 additionalFeeAmount) returns()
func (_Teleportermessenger *TeleportermessengerSession) AddFeeAmount(destinationChainID [32]byte, messageID *big.Int, feeContractAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.AddFeeAmount(&_Teleportermessenger.TransactOpts, destinationChainID, messageID, feeContractAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x19570c74.
//
// Solidity: function addFeeAmount(bytes32 destinationChainID, uint256 messageID, address feeContractAddress, uint256 additionalFeeAmount) returns()
func (_Teleportermessenger *TeleportermessengerTransactorSession) AddFeeAmount(destinationChainID [32]byte, messageID *big.Int, feeContractAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.AddFeeAmount(&_Teleportermessenger.TransactOpts, destinationChainID, messageID, feeContractAddress, additionalFeeAmount)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_Teleportermessenger *TeleportermessengerTransactor) ReceiveCrossChainMessage(opts *bind.TransactOpts, messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _Teleportermessenger.contract.Transact(opts, "receiveCrossChainMessage", messageIndex, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_Teleportermessenger *TeleportermessengerSession) ReceiveCrossChainMessage(messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.ReceiveCrossChainMessage(&_Teleportermessenger.TransactOpts, messageIndex, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_Teleportermessenger *TeleportermessengerTransactorSession) ReceiveCrossChainMessage(messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.ReceiveCrossChainMessage(&_Teleportermessenger.TransactOpts, messageIndex, relayerRewardAddress)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_Teleportermessenger *TeleportermessengerTransactor) RedeemRelayerRewards(opts *bind.TransactOpts, feeAsset common.Address) (*types.Transaction, error) {
	return _Teleportermessenger.contract.Transact(opts, "redeemRelayerRewards", feeAsset)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_Teleportermessenger *TeleportermessengerSession) RedeemRelayerRewards(feeAsset common.Address) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.RedeemRelayerRewards(&_Teleportermessenger.TransactOpts, feeAsset)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_Teleportermessenger *TeleportermessengerTransactorSession) RedeemRelayerRewards(feeAsset common.Address) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.RedeemRelayerRewards(&_Teleportermessenger.TransactOpts, feeAsset)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xcd3f2daa.
//
// Solidity: function retryMessageExecution(bytes32 originChainID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_Teleportermessenger *TeleportermessengerTransactor) RetryMessageExecution(opts *bind.TransactOpts, originChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _Teleportermessenger.contract.Transact(opts, "retryMessageExecution", originChainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xcd3f2daa.
//
// Solidity: function retryMessageExecution(bytes32 originChainID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_Teleportermessenger *TeleportermessengerSession) RetryMessageExecution(originChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.RetryMessageExecution(&_Teleportermessenger.TransactOpts, originChainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xcd3f2daa.
//
// Solidity: function retryMessageExecution(bytes32 originChainID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_Teleportermessenger *TeleportermessengerTransactorSession) RetryMessageExecution(originChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.RetryMessageExecution(&_Teleportermessenger.TransactOpts, originChainID, message)
}

// RetryReceipts is a paid mutator transaction binding the contract method 0x9a496900.
//
// Solidity: function retryReceipts(bytes32 originChainID, uint256[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(uint256 messageID)
func (_Teleportermessenger *TeleportermessengerTransactor) RetryReceipts(opts *bind.TransactOpts, originChainID [32]byte, messageIDs []*big.Int, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _Teleportermessenger.contract.Transact(opts, "retryReceipts", originChainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// RetryReceipts is a paid mutator transaction binding the contract method 0x9a496900.
//
// Solidity: function retryReceipts(bytes32 originChainID, uint256[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(uint256 messageID)
func (_Teleportermessenger *TeleportermessengerSession) RetryReceipts(originChainID [32]byte, messageIDs []*big.Int, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.RetryReceipts(&_Teleportermessenger.TransactOpts, originChainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// RetryReceipts is a paid mutator transaction binding the contract method 0x9a496900.
//
// Solidity: function retryReceipts(bytes32 originChainID, uint256[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(uint256 messageID)
func (_Teleportermessenger *TeleportermessengerTransactorSession) RetryReceipts(originChainID [32]byte, messageIDs []*big.Int, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.RetryReceipts(&_Teleportermessenger.TransactOpts, originChainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x1af671f8.
//
// Solidity: function retrySendCrossChainMessage(bytes32 destinationChainID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_Teleportermessenger *TeleportermessengerTransactor) RetrySendCrossChainMessage(opts *bind.TransactOpts, destinationChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _Teleportermessenger.contract.Transact(opts, "retrySendCrossChainMessage", destinationChainID, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x1af671f8.
//
// Solidity: function retrySendCrossChainMessage(bytes32 destinationChainID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_Teleportermessenger *TeleportermessengerSession) RetrySendCrossChainMessage(destinationChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.RetrySendCrossChainMessage(&_Teleportermessenger.TransactOpts, destinationChainID, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x1af671f8.
//
// Solidity: function retrySendCrossChainMessage(bytes32 destinationChainID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_Teleportermessenger *TeleportermessengerTransactorSession) RetrySendCrossChainMessage(destinationChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.RetrySendCrossChainMessage(&_Teleportermessenger.TransactOpts, destinationChainID, message)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(uint256 messageID)
func (_Teleportermessenger *TeleportermessengerTransactor) SendCrossChainMessage(opts *bind.TransactOpts, messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _Teleportermessenger.contract.Transact(opts, "sendCrossChainMessage", messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(uint256 messageID)
func (_Teleportermessenger *TeleportermessengerSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.SendCrossChainMessage(&_Teleportermessenger.TransactOpts, messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(uint256 messageID)
func (_Teleportermessenger *TeleportermessengerTransactorSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _Teleportermessenger.Contract.SendCrossChainMessage(&_Teleportermessenger.TransactOpts, messageInput)
}

// TeleportermessengerAddFeeAmountIterator is returned from FilterAddFeeAmount and is used to iterate over the raw logs and unpacked data for AddFeeAmount events raised by the Teleportermessenger contract.
type TeleportermessengerAddFeeAmountIterator struct {
	Event *TeleportermessengerAddFeeAmount // Event containing the contract specifics and raw log

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
func (it *TeleportermessengerAddFeeAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleportermessengerAddFeeAmount)
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
		it.Event = new(TeleportermessengerAddFeeAmount)
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
func (it *TeleportermessengerAddFeeAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleportermessengerAddFeeAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleportermessengerAddFeeAmount represents a AddFeeAmount event raised by the Teleportermessenger contract.
type TeleportermessengerAddFeeAmount struct {
	DestinationChainID [32]byte
	MessageID          *big.Int
	UpdatedFeeInfo     TeleporterFeeInfo
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAddFeeAmount is a free log retrieval operation binding the contract event 0x28fe05eedf0479c9159e5b6dd2a28c93fa1a408eba22dc801fd9bc493a7fc0c2.
//
// Solidity: event AddFeeAmount(bytes32 indexed destinationChainID, uint256 indexed messageID, (address,uint256) updatedFeeInfo)
func (_Teleportermessenger *TeleportermessengerFilterer) FilterAddFeeAmount(opts *bind.FilterOpts, destinationChainID [][32]byte, messageID []*big.Int) (*TeleportermessengerAddFeeAmountIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _Teleportermessenger.contract.FilterLogs(opts, "AddFeeAmount", destinationChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleportermessengerAddFeeAmountIterator{contract: _Teleportermessenger.contract, event: "AddFeeAmount", logs: logs, sub: sub}, nil
}

// WatchAddFeeAmount is a free log subscription operation binding the contract event 0x28fe05eedf0479c9159e5b6dd2a28c93fa1a408eba22dc801fd9bc493a7fc0c2.
//
// Solidity: event AddFeeAmount(bytes32 indexed destinationChainID, uint256 indexed messageID, (address,uint256) updatedFeeInfo)
func (_Teleportermessenger *TeleportermessengerFilterer) WatchAddFeeAmount(opts *bind.WatchOpts, sink chan<- *TeleportermessengerAddFeeAmount, destinationChainID [][32]byte, messageID []*big.Int) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _Teleportermessenger.contract.WatchLogs(opts, "AddFeeAmount", destinationChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleportermessengerAddFeeAmount)
				if err := _Teleportermessenger.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
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
func (_Teleportermessenger *TeleportermessengerFilterer) ParseAddFeeAmount(log types.Log) (*TeleportermessengerAddFeeAmount, error) {
	event := new(TeleportermessengerAddFeeAmount)
	if err := _Teleportermessenger.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleportermessengerFailedMessageExecutionIterator is returned from FilterFailedMessageExecution and is used to iterate over the raw logs and unpacked data for FailedMessageExecution events raised by the Teleportermessenger contract.
type TeleportermessengerFailedMessageExecutionIterator struct {
	Event *TeleportermessengerFailedMessageExecution // Event containing the contract specifics and raw log

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
func (it *TeleportermessengerFailedMessageExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleportermessengerFailedMessageExecution)
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
		it.Event = new(TeleportermessengerFailedMessageExecution)
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
func (it *TeleportermessengerFailedMessageExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleportermessengerFailedMessageExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleportermessengerFailedMessageExecution represents a FailedMessageExecution event raised by the Teleportermessenger contract.
type TeleportermessengerFailedMessageExecution struct {
	OriginChainID [32]byte
	MessageID     *big.Int
	Message       TeleporterMessage
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFailedMessageExecution is a free log retrieval operation binding the contract event 0x50e5f3de5b0b3e6c82553b7e0f5f7080a291be07aa30a311084a9457f4a956bc.
//
// Solidity: event FailedMessageExecution(bytes32 indexed originChainID, uint256 indexed messageID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message)
func (_Teleportermessenger *TeleportermessengerFilterer) FilterFailedMessageExecution(opts *bind.FilterOpts, originChainID [][32]byte, messageID []*big.Int) (*TeleportermessengerFailedMessageExecutionIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _Teleportermessenger.contract.FilterLogs(opts, "FailedMessageExecution", originChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleportermessengerFailedMessageExecutionIterator{contract: _Teleportermessenger.contract, event: "FailedMessageExecution", logs: logs, sub: sub}, nil
}

// WatchFailedMessageExecution is a free log subscription operation binding the contract event 0x50e5f3de5b0b3e6c82553b7e0f5f7080a291be07aa30a311084a9457f4a956bc.
//
// Solidity: event FailedMessageExecution(bytes32 indexed originChainID, uint256 indexed messageID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message)
func (_Teleportermessenger *TeleportermessengerFilterer) WatchFailedMessageExecution(opts *bind.WatchOpts, sink chan<- *TeleportermessengerFailedMessageExecution, originChainID [][32]byte, messageID []*big.Int) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _Teleportermessenger.contract.WatchLogs(opts, "FailedMessageExecution", originChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleportermessengerFailedMessageExecution)
				if err := _Teleportermessenger.contract.UnpackLog(event, "FailedMessageExecution", log); err != nil {
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
func (_Teleportermessenger *TeleportermessengerFilterer) ParseFailedMessageExecution(log types.Log) (*TeleportermessengerFailedMessageExecution, error) {
	event := new(TeleportermessengerFailedMessageExecution)
	if err := _Teleportermessenger.contract.UnpackLog(event, "FailedMessageExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleportermessengerMessageExecutionRetriedIterator is returned from FilterMessageExecutionRetried and is used to iterate over the raw logs and unpacked data for MessageExecutionRetried events raised by the Teleportermessenger contract.
type TeleportermessengerMessageExecutionRetriedIterator struct {
	Event *TeleportermessengerMessageExecutionRetried // Event containing the contract specifics and raw log

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
func (it *TeleportermessengerMessageExecutionRetriedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleportermessengerMessageExecutionRetried)
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
		it.Event = new(TeleportermessengerMessageExecutionRetried)
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
func (it *TeleportermessengerMessageExecutionRetriedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleportermessengerMessageExecutionRetriedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleportermessengerMessageExecutionRetried represents a MessageExecutionRetried event raised by the Teleportermessenger contract.
type TeleportermessengerMessageExecutionRetried struct {
	OriginChainID [32]byte
	MessageID     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMessageExecutionRetried is a free log retrieval operation binding the contract event 0x3d5f30e93c1e27cda0e05a7b9e51144613a816cd90561f8493393bbcf4e00358.
//
// Solidity: event MessageExecutionRetried(bytes32 indexed originChainID, uint256 indexed messageID)
func (_Teleportermessenger *TeleportermessengerFilterer) FilterMessageExecutionRetried(opts *bind.FilterOpts, originChainID [][32]byte, messageID []*big.Int) (*TeleportermessengerMessageExecutionRetriedIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _Teleportermessenger.contract.FilterLogs(opts, "MessageExecutionRetried", originChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleportermessengerMessageExecutionRetriedIterator{contract: _Teleportermessenger.contract, event: "MessageExecutionRetried", logs: logs, sub: sub}, nil
}

// WatchMessageExecutionRetried is a free log subscription operation binding the contract event 0x3d5f30e93c1e27cda0e05a7b9e51144613a816cd90561f8493393bbcf4e00358.
//
// Solidity: event MessageExecutionRetried(bytes32 indexed originChainID, uint256 indexed messageID)
func (_Teleportermessenger *TeleportermessengerFilterer) WatchMessageExecutionRetried(opts *bind.WatchOpts, sink chan<- *TeleportermessengerMessageExecutionRetried, originChainID [][32]byte, messageID []*big.Int) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _Teleportermessenger.contract.WatchLogs(opts, "MessageExecutionRetried", originChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleportermessengerMessageExecutionRetried)
				if err := _Teleportermessenger.contract.UnpackLog(event, "MessageExecutionRetried", log); err != nil {
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
func (_Teleportermessenger *TeleportermessengerFilterer) ParseMessageExecutionRetried(log types.Log) (*TeleportermessengerMessageExecutionRetried, error) {
	event := new(TeleportermessengerMessageExecutionRetried)
	if err := _Teleportermessenger.contract.UnpackLog(event, "MessageExecutionRetried", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleportermessengerReceiveCrossChainMessageIterator is returned from FilterReceiveCrossChainMessage and is used to iterate over the raw logs and unpacked data for ReceiveCrossChainMessage events raised by the Teleportermessenger contract.
type TeleportermessengerReceiveCrossChainMessageIterator struct {
	Event *TeleportermessengerReceiveCrossChainMessage // Event containing the contract specifics and raw log

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
func (it *TeleportermessengerReceiveCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleportermessengerReceiveCrossChainMessage)
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
		it.Event = new(TeleportermessengerReceiveCrossChainMessage)
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
func (it *TeleportermessengerReceiveCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleportermessengerReceiveCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleportermessengerReceiveCrossChainMessage represents a ReceiveCrossChainMessage event raised by the Teleportermessenger contract.
type TeleportermessengerReceiveCrossChainMessage struct {
	OriginChainID [32]byte
	MessageID     *big.Int
	Message       TeleporterMessage
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceiveCrossChainMessage is a free log retrieval operation binding the contract event 0x522ce4da81e9fb994bf6a122282939647b6e69817be443f7d1c152e76a082cc7.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed originChainID, uint256 indexed messageID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message)
func (_Teleportermessenger *TeleportermessengerFilterer) FilterReceiveCrossChainMessage(opts *bind.FilterOpts, originChainID [][32]byte, messageID []*big.Int) (*TeleportermessengerReceiveCrossChainMessageIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _Teleportermessenger.contract.FilterLogs(opts, "ReceiveCrossChainMessage", originChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleportermessengerReceiveCrossChainMessageIterator{contract: _Teleportermessenger.contract, event: "ReceiveCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchReceiveCrossChainMessage is a free log subscription operation binding the contract event 0x522ce4da81e9fb994bf6a122282939647b6e69817be443f7d1c152e76a082cc7.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed originChainID, uint256 indexed messageID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message)
func (_Teleportermessenger *TeleportermessengerFilterer) WatchReceiveCrossChainMessage(opts *bind.WatchOpts, sink chan<- *TeleportermessengerReceiveCrossChainMessage, originChainID [][32]byte, messageID []*big.Int) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _Teleportermessenger.contract.WatchLogs(opts, "ReceiveCrossChainMessage", originChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleportermessengerReceiveCrossChainMessage)
				if err := _Teleportermessenger.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
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
func (_Teleportermessenger *TeleportermessengerFilterer) ParseReceiveCrossChainMessage(log types.Log) (*TeleportermessengerReceiveCrossChainMessage, error) {
	event := new(TeleportermessengerReceiveCrossChainMessage)
	if err := _Teleportermessenger.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleportermessengerSendCrossChainMessageIterator is returned from FilterSendCrossChainMessage and is used to iterate over the raw logs and unpacked data for SendCrossChainMessage events raised by the Teleportermessenger contract.
type TeleportermessengerSendCrossChainMessageIterator struct {
	Event *TeleportermessengerSendCrossChainMessage // Event containing the contract specifics and raw log

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
func (it *TeleportermessengerSendCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleportermessengerSendCrossChainMessage)
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
		it.Event = new(TeleportermessengerSendCrossChainMessage)
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
func (it *TeleportermessengerSendCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleportermessengerSendCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleportermessengerSendCrossChainMessage represents a SendCrossChainMessage event raised by the Teleportermessenger contract.
type TeleportermessengerSendCrossChainMessage struct {
	DestinationChainID [32]byte
	MessageID          *big.Int
	Message            TeleporterMessage
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSendCrossChainMessage is a free log retrieval operation binding the contract event 0x7491aecf1f3e24837ce48fd97fc8729fc036cebb3e5078643f3301b72852aa07.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed destinationChainID, uint256 indexed messageID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message)
func (_Teleportermessenger *TeleportermessengerFilterer) FilterSendCrossChainMessage(opts *bind.FilterOpts, destinationChainID [][32]byte, messageID []*big.Int) (*TeleportermessengerSendCrossChainMessageIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _Teleportermessenger.contract.FilterLogs(opts, "SendCrossChainMessage", destinationChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleportermessengerSendCrossChainMessageIterator{contract: _Teleportermessenger.contract, event: "SendCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchSendCrossChainMessage is a free log subscription operation binding the contract event 0x7491aecf1f3e24837ce48fd97fc8729fc036cebb3e5078643f3301b72852aa07.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed destinationChainID, uint256 indexed messageID, (uint256,address,address,uint256,address[],(uint256,address)[],bytes) message)
func (_Teleportermessenger *TeleportermessengerFilterer) WatchSendCrossChainMessage(opts *bind.WatchOpts, sink chan<- *TeleportermessengerSendCrossChainMessage, destinationChainID [][32]byte, messageID []*big.Int) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _Teleportermessenger.contract.WatchLogs(opts, "SendCrossChainMessage", destinationChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleportermessengerSendCrossChainMessage)
				if err := _Teleportermessenger.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
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
func (_Teleportermessenger *TeleportermessengerFilterer) ParseSendCrossChainMessage(log types.Log) (*TeleportermessengerSendCrossChainMessage, error) {
	event := new(TeleportermessengerSendCrossChainMessage)
	if err := _Teleportermessenger.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
