// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokenremoteupgradeable

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
	DestinationBlockchainID            [32]byte
	DestinationTokenTransferrerAddress common.Address
	RecipientContract                  common.Address
	RecipientPayload                   []byte
	RequiredGasLimit                   *big.Int
	RecipientGasLimit                  *big.Int
	MultiHopFallback                   common.Address
	FallbackRecipient                  common.Address
	PrimaryFeeTokenAddress             common.Address
	PrimaryFee                         *big.Int
	SecondaryFee                       *big.Int
}

// SendTokensInput is an auto generated low-level Go binding around an user-defined struct.
type SendTokensInput struct {
	DestinationBlockchainID            [32]byte
	DestinationTokenTransferrerAddress common.Address
	Recipient                          common.Address
	PrimaryFeeTokenAddress             common.Address
	PrimaryFee                         *big.Int
	SecondaryFee                       *big.Int
	RequiredGasLimit                   *big.Int
	MultiHopFallback                   common.Address
}

// TeleporterFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type TeleporterFeeInfo struct {
	FeeTokenAddress common.Address
	Amount          *big.Int
}

// TokenRemoteSettings is an auto generated low-level Go binding around an user-defined struct.
type TokenRemoteSettings struct {
	TeleporterRegistryAddress common.Address
	TeleporterManager         common.Address
	MinTeleporterVersion      *big.Int
	TokenHomeBlockchainID     [32]byte
	TokenHomeAddress          common.Address
	TokenHomeDecimals         uint8
}

// NativeTokenRemoteUpgradeableMetaData contains all meta data concerning the NativeTokenRemoteUpgradeable contract.
var NativeTokenRemoteUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"init\",\"type\":\"uint8\",\"internalType\":\"enumICMInitializable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"BURNED_FOR_TRANSFER_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BURNED_TX_FEES_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"HOME_CHAIN_BURN_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MULTI_HOP_CALL_GAS_PER_WORD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MULTI_HOP_CALL_REQUIRED_GAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MULTI_HOP_SEND_REQUIRED_GAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NATIVE_MINTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINativeMinter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NATIVE_TOKEN_REMOTE_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTER_REMOTE_REQUIRED_GAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TELEPORTER_REGISTRY_APP_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_REMOTE_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateNumWords\",\"inputs\":[{\"name\":\"payloadSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getBlockchainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInitialReserveImbalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIsCollateralized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinTeleporterVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMultiplyOnRemote\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenHomeAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenHomeBlockchainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenMultiplier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalMinted\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"settings\",\"type\":\"tuple\",\"internalType\":\"structTokenRemoteSettings\",\"components\":[{\"name\":\"teleporterRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"teleporterManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minTeleporterVersion\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenHomeBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenHomeAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenHomeDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"nativeAssetSymbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"initialReserveImbalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"burnedFeesReportingRewardPercentage\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isTeleporterAddressPaused\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseTeleporterAddress\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"receiveTeleporterMessage\",\"inputs\":[{\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originSenderAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerWithHome\",\"inputs\":[{\"name\":\"feeInfo\",\"type\":\"tuple\",\"internalType\":\"structTeleporterFeeInfo\",\"components\":[{\"name\":\"feeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reportBurnedTxFees\",\"inputs\":[{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"send\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structSendTokensInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"sendAndCall\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structSendAndCallInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipientGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalNativeAssetSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseTeleporterAddress\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMinTeleporterVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"recipientContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSucceeded\",\"inputs\":[{\"name\":\"recipientContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinTeleporterVersionUpdated\",\"inputs\":[{\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReportBurnedTxFees\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"feesBurned\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TeleporterAddressPaused\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TeleporterAddressUnpaused\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensAndCallSent\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"input\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipientGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensSent\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"input\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051614fb0380380614fb083398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b614e638061014d5f395ff3fe608060405260043610610280575f3560e01c8063715018a61161014e578063b8a46d02116100c0578063dd62ed3e11610079578063dd62ed3e146106e9578063e0fd9cb814610708578063ed0ae4b014610428578063ef793e2a1461071c578063f2fde38b14610730578063f3f981d81461074f5761028f565b8063b8a46d021461066c578063c3cd69271461068b578063c452165e1461069f578063c868efaa146106b6578063d0e30db01461028f578063d2cc7a70146106d55761028f565b80638f6cec88116101125780638f6cec8814610595578063909a6ac0146105b457806395d89b41146105e757806397314297146105fb578063a5717bc01461061a578063a9059cbb1461064d5761028f565b8063715018a61461053057806371717c18146105445780637ee3779a1461055a5780638bf2fa941461056e5780638da5cb5b146105815761028f565b80632e1a7d4d116101f25780634511243e116101ab5780634511243e1461048b5780635507f3d1146104aa57806355538c8b146104c05780635eb99514146104df5780636e6eef8d146104fe57806370a08231146105115761028f565b80632e1a7d4d146103c7578063313ce567146103e6578063329c3e1214610401578063347212c41461042857806335cac159146104445780634213cf78146104775761028f565b806315beb59f1161024457806315beb59f1461033657806318160ddd1461034b5780631906529c1461035f57806323b872dd14610373578063254ac160146103925780632b0d8f18146103a85761028f565b806302a30c7d1461029757806306fdde03146102c05780630733c8c8146102e1578063095ea7b3146103035780630ca1c5c9146103225761028f565b3661028f5761028d61076e565b005b61028d61076e565b3480156102a2575f80fd5b506102ab6107af565b60405190151581526020015b60405180910390f35b3480156102cb575f80fd5b506102d46107c6565b6040516102b79190613e1d565b3480156102ec575f80fd5b506102f5610864565b6040519081526020016102b7565b34801561030e575f80fd5b506102ab61031d366004613e53565b610878565b34801561032d575f80fd5b506102f5610891565b348015610341575f80fd5b506102f56105dc81565b348015610356575f80fd5b506102f56108a3565b34801561036a575f80fd5b506102f56108b7565b34801561037e575f80fd5b506102ab61038d366004613e7d565b61090b565b34801561039d575f80fd5b506102f56201fbd081565b3480156103b3575f80fd5b5061028d6103c2366004613ebb565b610930565b3480156103d2575f80fd5b5061028d6103e1366004613ed6565b610a2f565b3480156103f1575f80fd5b50604051601281526020016102b7565b34801561040c575f80fd5b5061041b6001600160991b0181565b6040516102b79190613eed565b348015610433575f80fd5b5061041b62010203600160981b0181565b34801561044f575f80fd5b506102f57f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef0081565b348015610482575f80fd5b506102f5610a7b565b348015610496575f80fd5b5061028d6104a5366004613ebb565b610a8c565b3480156104b5575f80fd5b506102f56205302081565b3480156104cb575f80fd5b5061028d6104da366004613ed6565b610b78565b3480156104ea575f80fd5b5061028d6104f9366004613ed6565b610e68565b61028d61050c366004613f0e565b610e79565b34801561051c575f80fd5b506102f561052b366004613ebb565b610ea7565b34801561053b575f80fd5b5061028d610ed0565b34801561054f575f80fd5b506102f56205573081565b348015610565575f80fd5b506102ab610ee1565b61028d61057c366004613f45565b610ef8565b34801561058c575f80fd5b5061041b610f26565b3480156105a0575f80fd5b5061028d6105af366004614085565b610f54565b3480156105bf575f80fd5b506102f57fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d0081565b3480156105f2575f80fd5b506102d4611066565b348015610606575f80fd5b506102ab610615366004613ebb565b611082565b348015610625575f80fd5b506102f57f69a5f7616543528c4fbe43f410b1034bd6da4ba06c25bedf04617268014cf50081565b348015610658575f80fd5b506102ab610667366004613e53565b611098565b348015610677575f80fd5b5061028d61068636600461414f565b6110a5565b348015610696575f80fd5b5061041b611219565b3480156106aa575f80fd5b5061041b600160981b81565b3480156106c1575f80fd5b5061028d6106d036600461415f565b611236565b3480156106e0575f80fd5b506102f56113cf565b3480156106f4575f80fd5b506102f56107033660046141e0565b6113d9565b348015610713575f80fd5b506102f5611413565b348015610727575f80fd5b506102f5611427565b34801561073b575f80fd5b5061028d61074a366004613ebb565b61143b565b34801561075a575f80fd5b506102f5610769366004613ed6565b611475565b60405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a26107ad333461148b565b565b5f806107b96114c3565b6006015460ff1692915050565b60605f6107d16114e7565b90508060030180546107e290614217565b80601f016020809104026020016040519081016040528092919081815260200182805461080e90614217565b80156108595780601f1061083057610100808354040283529160200191610859565b820191905f5260205f20905b81548152906001019060200180831161083c57829003601f168201915b505050505091505090565b5f8061086e6114c3565b6003015492915050565b5f3361088581858561150b565b60019150505b92915050565b5f61089a61151d565b60010154905090565b5f806108ad6114e7565b6002015492915050565b5f806108c161151d565b90505f6108dd62010203600160981b0131600160981b3161425d565b90505f6108e8611427565b83600101546108f7919061425d565b90506109038282614270565b935050505090565b5f33610918858285611541565b61092385858561158b565b60019150505b9392505050565b5f6109396115e8565b905061094361160c565b6001600160a01b0382166109725760405162461bcd60e51b815260040161096990614283565b60405180910390fd5b61097c8183611614565b156109df5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608401610969565b6001600160a01b0382165f81815260018381016020526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a25050565b60405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a2610a6e3382611635565b610a783382611669565b50565b5f80610a856114c3565b5492915050565b5f610a956115e8565b9050610a9f61160c565b6001600160a01b038216610ac55760405162461bcd60e51b815260040161096990614283565b610acf8183611614565b610b2d5760405162461bcd60e51b815260206004820152602960248201527f54656c65706f7274657252656769737472794170703a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608401610969565b6001600160a01b0382165f818152600183016020526040808220805460ff19169055517f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c39190a25050565b5f610b816116fc565b8054909150600114610ba55760405162461bcd60e51b8152600401610969906142d1565b600281555f610bb261151d565b6002810154909150600160981b31908111610c435760405162461bcd60e51b8152602060048201526044602482018190527f4e6174697665546f6b656e52656d6f74653a206275726e206164647265737320908201527f62616c616e6365206e6f742067726561746572207468616e206c6173742072656064820152631c1bdc9d60e21b608482015260a401610969565b5f826002015482610c549190614270565b90505f6064845f015483610c689190614315565b610c72919061432c565b90505f610c7f8284614270565b6002860185905590508115610ca257610c983083611720565b610ca2308361148b565b5f610cbc610cae610864565b610cb6610ee1565b846117a1565b11610d265760405162461bcd60e51b815260206004820152603460248201527f4e6174697665546f6b656e52656d6f74653a207a65726f207363616c6564206160448201527336b7bab73a103a37903932b837b93a10313ab93760611b6064820152608401610969565b604080518082018252600181528151808301835262010203600160981b018152602080820185905292515f9380840192610d629290910161435f565b60405160208183030381529060405281525090505f610e1d6040518060c00160405280610d8d611413565b8152602001610d9a611219565b6001600160a01b0316815260408051808201825230815260208181018a905283015281018c90526060015f5b604051908082528060200260200182016040528015610def578160200160208202803683370190505b50815260200184604051602001610e06919061437f565b6040516020818303038152906040528152506117b6565b9050807f0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a584604051610e5191815260200190565b60405180910390a250506001909555505050505050565b610e7061160c565b610a78816118d1565b610e816107af565b610e9d5760405162461bcd60e51b8152600401610969906143c1565b610a788134611a5f565b5f80610eb16114e7565b6001600160a01b039093165f9081526020939093525050604090205490565b610ed8611ad4565b6107ad5f611b06565b5f80610eeb6114c3565b6004015460ff1692915050565b610f006107af565b610f1c5760405162461bcd60e51b8152600401610969906143c1565b610a788134611b76565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015610f985750825b90505f826001600160401b03166001148015610fb35750303b155b905081158015610fc1575080155b15610fdf5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561100957845460ff60401b1916600160401b1785555b61101589898989611be0565b831561105b57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b60605f6110716114e7565b90508060040180546107e290614217565b5f8061108c6115e8565b90506109298184611614565b5f3361088581858561158b565b5f6110ae6114c3565b6006810154909150610100900460ff161561110b5760405162461bcd60e51b815260206004820152601f60248201527f546f6b656e52656d6f74653a20616c72656164792072656769737465726564006044820152606401610969565b604080516060808201835260058401548252600284015460ff600160a01b820481166020808601918252600160a81b9093048216858701908152865180880188525f808252885188518188015293518516848a015291519093168286015286518083039095018552608090910190955280820192909252919290916111a09061119690870187613ebb565b8660200135611c70565b6040805160c0810182526001870154815260028701546001600160a01b0316602080830191909152825180840184529394506112119391928301919081906111ea908b018b613ebb565b6001600160a01b0316815260209081018690529082526201fbd0908201526040015f610dc6565b505050505050565b5f806112236114c3565b600201546001600160a01b031692915050565b61123e611cb7565b5f6112476115e8565b60028101548154919250906001600160a01b0316634c1f08ce336040518263ffffffff1660e01b815260040161127d9190613eed565b602060405180830381865afa158015611298573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112bc9190614410565b10156113105760405162461bcd60e51b815260206004820152603060248201525f80516020614e3783398151915260448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608401610969565b61131a8133611614565b156113805760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608401610969565b6113c0858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250611d0192505050565b506113c9611f18565b50505050565b5f806108ad6115e8565b5f806113e36114e7565b6001600160a01b039485165f90815260019190910160209081526040808320959096168252939093525050205490565b5f8061141d6114c3565b6001015492915050565b5f806114316114c3565b6005015492915050565b611443611ad4565b6001600160a01b03811661146c575f604051631e4fbdf760e01b81526004016109699190613eed565b610a7881611b06565b5f600561148383601f61425d565b901c92915050565b6001600160a01b0382166114b4575f60405163ec442f0560e01b81526004016109699190613eed565b6114bf5f8383611f42565b5050565b7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef0090565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0090565b6115188383836001612065565b505050565b7f69a5f7616543528c4fbe43f410b1034bd6da4ba06c25bedf04617268014cf50090565b5f61154c84846113d9565b90505f1981146113c9578181101561157d57828183604051637dc7a0d960e11b815260040161096993929190614427565b6113c984848484035f612065565b6001600160a01b0383166115b4575f604051634b637e8f60e11b81526004016109699190613eed565b6001600160a01b0382166115dd575f60405163ec442f0560e01b81526004016109699190613eed565b611518838383611f42565b7fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d0090565b6107ad611ad4565b6001600160a01b03165f908152600191909101602052604090205460ff1690565b6001600160a01b03821661165e575f604051634b637e8f60e11b81526004016109699190613eed565b6114bf825f83611f42565b8047101561168c573060405163cd78605960e01b81526004016109699190613eed565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f81146116d5576040519150601f19603f3d011682016040523d82523d5f602084013e6116da565b606091505b505090508061151857604051630a12f52160e11b815260040160405180910390fd5b7fd2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c750090565b5f61172961151d565b905081816001015f82825461173e919061425d565b90915550506040516327ad555d60e11b81526001600160991b0190634f5aaaba9061176f9086908690600401614448565b5f604051808303815f87803b158015611786575f80fd5b505af1158015611798573d5f803e3d5ffd5b50505050505050565b5f6117ae8484845f612145565b949350505050565b5f806117c0612175565b60408401516020015190915015611865576040830151516001600160a01b03166118425760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608401610969565b604083015160208101519051611865916001600160a01b03909116908390612269565b604051630624488560e41b81526001600160a01b03821690636244885090611891908690600401614461565b6020604051808303815f875af11580156118ad573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109299190614410565b5f6118da6115e8565b90505f815f015f9054906101000a90046001600160a01b03166001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa15801561192e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119529190614410565b6002830154909150818411156119b15760405162461bcd60e51b815260206004820152603160248201525f80516020614e3783398151915260448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608401610969565b808411611a265760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608401610969565b60028301849055604051849082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d905f90a350505050565b5f611a686116fc565b8054909150600114611a8c5760405162461bcd60e51b8152600401610969906142d1565b600281555f611a996114c3565b9050611aa4846122f0565b6001810154843503611ac057611aba848461252a565b50611acc565b611aba848461274e565b505b600190555050565b33611add610f26565b6001600160a01b0316146107ad573360405163118cdaa760e01b81526004016109699190613eed565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f611b7f6116fc565b8054909150600114611ba35760405162461bcd60e51b8152600401610969906142d1565b600281555f611bb06114c3565b9050611bbb846129ef565b6001810154843503611bd657611bd18484612ada565b611aca565b611aca8484612c49565b611be8612e0f565b815f03611c515760405162461bcd60e51b815260206004820152603160248201527f4e6174697665546f6b656e52656d6f74653a207a65726f20696e697469616c206044820152707265736572766520696d62616c616e636560781b6064820152608401610969565b611c5b8384612e58565b611c6784836012612e6a565b6113c981612e9b565b5f815f03611c7f57505f61088b565b306001600160a01b03841603611cac57611c9a333084611541565b611ca533308461158b565b508061088b565b610929833384612f0e565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00805460011901611cfb57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f611d0a6114c3565b905080600101548414611d715760405162461bcd60e51b815260206004820152602960248201527f546f6b656e52656d6f74653a20696e76616c696420736f7572636520626c6f636044820152681ad8da185a5b88125160ba1b6064820152608401610969565b60028101546001600160a01b03848116911614611de35760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e52656d6f74653a20696e76616c6964206f726967696e2073656e646044820152696572206164647265737360b01b6064820152608401610969565b5f82806020019051810190611df8919061455a565b6006830154909150610100900460ff161580611e195750600682015460ff16155b15611e305760068201805461ffff19166101011790555b600181516004811115611e4557611e4561434b565b03611e7c575f8160200151806020019051810190611e6391906145e2565b9050611e76815f01518260200151613079565b50611f11565b600281516004811115611e9157611e9161434b565b03611ebf575f8160200151806020019051810190611eaf919061461a565b9050611e768182608001516130c6565b60405162461bcd60e51b815260206004820152602160248201527f546f6b656e52656d6f74653a20696e76616c6964206d657373616765207479706044820152606560f81b6064820152608401610969565b5050505050565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005b6001905550565b5f611f4b6114e7565b90506001600160a01b038416611f795781816002015f828254611f6e919061425d565b90915550611fd69050565b6001600160a01b0384165f9081526020829052604090205482811015611fb85784818460405163391434e360e21b815260040161096993929190614427565b6001600160a01b0385165f9081526020839052604090209083900390555b6001600160a01b038316611ff4576002810180548390039055612012565b6001600160a01b0383165f9081526020829052604090208054830190555b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161205791815260200190565b60405180910390a350505050565b5f61206e6114e7565b90506001600160a01b038516612099575f60405163e602df0560e01b81526004016109699190613eed565b6001600160a01b0384166120c2575f604051634a1406b160e11b81526004016109699190613eed565b6001600160a01b038086165f90815260018301602090815260408083209388168352929052208390558115611f1157836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258560405161213691815260200190565b60405180910390a35050505050565b5f811515841515036121625761215b8584614315565b90506117ae565b61216c858461432c565b95945050505050565b5f8061217f6115e8565b90505f815f015f9054906101000a90046001600160a01b03166001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156121d3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121f791906146e4565b90506122038282611614565b1561088b5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608401610969565b604051636eb1769f60e11b81523060048201526001600160a01b0383811660248301525f919085169063dd62ed3e90604401602060405180830381865afa1580156122b6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122da9190614410565b90506113c984846122eb858561425d565b6131f0565b803561230e5760405162461bcd60e51b8152600401610969906146ff565b5f61231f6040830160208401613ebb565b6001600160a01b0316036123455760405162461bcd60e51b81526004016109699061474a565b5f6123566060830160408401613ebb565b6001600160a01b0316036123c15760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e7420636f6e7460448201526b72616374206164647265737360a01b6064820152608401610969565b5f8160800135116123e45760405162461bcd60e51b8152600401610969906147a1565b5f8160a00135116124455760405162461bcd60e51b815260206004820152602560248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e7420676173206044820152641b1a5b5a5d60da1b6064820152608401610969565b80608001358160a00135106124ad5760405162461bcd60e51b815260206004820152602860248201527f546f6b656e52656d6f74653a20696e76616c696420726563697069656e742067604482015267185cc81b1a5b5a5d60c21b6064820152608401610969565b5f6124bf610100830160e08401613ebb565b6001600160a01b031603610a785760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a207a65726f2066616c6c6261636b20726563697060448201526b69656e74206164647265737360a01b6064820152608401610969565b5f6125336114c3565b90506125636125486040850160208601613ebb565b61014085013561255e60e0870160c08801613ebb565b6132b2565b5f61258b8361257a61012087016101008801613ebb565b8661012001358761014001356133af565b6040805180820190915291945091505f908060028152602001604051806101000160405280865f01548152602001306001600160a01b031681526020016125cf3390565b6001600160a01b031681526020016125ed60608a0160408b01613ebb565b6001600160a01b031681526020810188905260400161260f60608a018a6147e5565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525050509082525060a089013560208201526040016126636101008a0160e08b01613ebb565b6001600160a01b0316905260405161267e9190602001614827565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b03166020820152815180830183529293505f9261270092820190806126d86101208c016101008d01613ebb565b6001600160a01b03168152602090810188905290825260808a0135908201526040015f610dc6565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16888860405161273e92919061491e565b60405180910390a3505050505050565b5f6127576114c3565b9050612783833561276e6040860160208701613ebb565b61277e60e0870160c08801613ebb565b613472565b5f61279a8361257a61012087016101008801613ebb565b6040805180820190915291945091505f9080600481526020016040518061016001604052806127c63390565b6001600160a01b03168152602001885f013581526020018860200160208101906127f09190613ebb565b6001600160a01b0316815260200161280e60608a0160408b01613ebb565b6001600160a01b031681526020810188905260400161283060608a018a6147e5565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525050509082525060a089013560208201526040016128846101008a0160e08b01613ebb565b6001600160a01b03168152608089013560208201526040016128ac60e08a0160c08b01613ebb565b6001600160a01b031681526101408901356020918201526040516128d1929101614a14565b60408051601f19818403018152919052905290505f6105dc6129006128f960608901896147e5565b9050611475565b61290a9190614315565b612917906205573061425d565b6040805160c0810182526001870154815260028701546001600160a01b03166020820152815180830183529293505f926129a092820190806129616101208d016101008e01613ebb565b6001600160a01b031681526020908101899052908252818101869052604080515f815280830182528184015251606090920191610e069188910161437f565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b1689896040516129de92919061491e565b60405180910390a350505050505050565b5f612a006060830160408401613ebb565b6001600160a01b031603612a625760405162461bcd60e51b815260206004820152602360248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e74206164647260448201526265737360e81b6064820152608401610969565b5f8160c0013511612a855760405162461bcd60e51b8152600401610969906147a1565b8035612aa35760405162461bcd60e51b8152600401610969906146ff565b5f612ab46040830160208401613ebb565b6001600160a01b031603610a785760405162461bcd60e51b81526004016109699061474a565b5f612ae36114c3565b9050612b0e612af86040850160208601613ebb565b60a085013561255e610100870160e08801613ebb565b5f612b3283612b236080870160608801613ebb565b86608001358760a001356133af565b6040805180820190915291945091505f9080600181526020016040518060400160405280886040016020810190612b699190613ebb565b6001600160a01b0316815260200187815250604051602001612b8b919061435f565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b03166020820152815180830183529293505f92612c0b9282019080612be360808c0160608d01613ebb565b6001600160a01b03168152602090810188905290825260c08a0135908201526040015f610dc6565b9050336001600160a01b0316817f93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52888860405161273e929190614acf565b5f612c526114c3565b9050612c7a8335612c696040860160208701613ebb565b61277e610100870160e08801613ebb565b5f612c8f83612b236080870160608801613ebb565b60408051808201825260038152815160e081018352883581529396509193505f9260208084019282820191612cc8918b01908b01613ebb565b6001600160a01b03168152602001612ce660608a0160408b01613ebb565b6001600160a01b031681526020810188905260a0890135604082015260c08901356060820152608001612d206101008a0160e08b01613ebb565b6001600160a01b03169052604051612d909190602001815181526020808301516001600160a01b0390811691830191909152604080840151821690830152606080840151908301526080808401519083015260a0808401519083015260c092830151169181019190915260e00190565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b03166020820152815180830183529293505f92612c0b9282019080612de860808c0160608d01613ebb565b6001600160a01b03168152602090810188905290825262053020908201526040015f610dc6565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166107ad57604051631afcd79f60e31b815260040160405180910390fd5b612e60612e0f565b6114bf8282613510565b612e72612e0f565b612e88835f015184602001518560400151613540565b612e9061355b565b61151883838361356b565b612ea3612e0f565b60648110612f015760405162461bcd60e51b815260206004820152602560248201527f4e6174697665546f6b656e52656d6f74653a20696e76616c69642070657263656044820152646e7461676560d81b6064820152608401610969565b80612f0a61151d565b5550565b5f80846001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401612f3c9190613eed565b602060405180830381865afa158015612f57573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f7b9190614410565b9050612f926001600160a01b03861685308661388f565b6040516370a0823160e01b81525f906001600160a01b038716906370a0823190612fc0903090600401613eed565b602060405180830381865afa158015612fdb573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612fff9190614410565b90508181116130655760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610969565b61306f8282614270565b9695505050505050565b816001600160a01b03167f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b826040516130b491815260200190565b60405180910390a26114bf8282611720565b6130d03082611720565b5f825f0151836020015184604001518560a001516040516024016130f79493929190614b6b565b60408051601f198184030181529190526020810180516001600160e01b031663161b12ff60e11b17905260c084015160608501519192505f9161313d91908590856138c8565b905080156131915783606001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff48460405161318491815260200190565b60405180910390a26113c9565b83606001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0846040516131d091815260200190565b60405180910390a260e08401516113c9906001600160a01b031684611669565b5f836001600160a01b031663095ea7b38484604051602401613213929190614448565b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050905061324c8482613998565b6113c9576132a884856001600160a01b031663095ea7b3865f604051602401613276929190614448565b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613a35565b6113c98482613a35565b5f6132bb6114c3565b60028101549091506001600160a01b038581169116146132ed5760405162461bcd60e51b815260040161096990614b9c565b82156133475760405162461bcd60e51b815260206004820152602360248201527f546f6b656e52656d6f74653a206e6f6e2d7a65726f207365636f6e646172792060448201526266656560e81b6064820152608401610969565b6001600160a01b038216156113c95760405162461bcd60e51b815260206004820152602860248201527f546f6b656e52656d6f74653a206e6f6e2d7a65726f206d756c74692d686f702060448201526766616c6c6261636b60c01b6064820152608401610969565b5f805f6133ba6114c3565b90506133c587613a8d565b96506133d18686611c70565b600382015460048301549196506133eb9160ff16866117a1565b60038201546004830154613403919060ff168a6117a1565b116134655760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a20696e73756666696369656e7420746f6b656e7360448201526b103a37903a3930b739b332b960a11b6064820152608401610969565b5094959294509192505050565b5f61347b6114c3565b805490915084036134ae57306001600160a01b038416036134ae5760405162461bcd60e51b815260040161096990614b9c565b6001600160a01b0382166113c95760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a207a65726f206d756c74692d686f702066616c6c6044820152636261636b60e01b6064820152608401610969565b613518612e0f565b5f6135216114e7565b9050600381016135318482614c3d565b50600481016113c98382614c3d565b613548612e0f565b6135528382613aa5565b61151882613ac7565b613563612e0f565b6107ad613ad8565b613573612e0f565b5f61357c6114c3565b90506005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa1580156135c1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135e59190614410565b8155606084015161364b5760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e52656d6f74653a207a65726f20746f6b656e20686f6d6520626c6f60448201526918dad8da185a5b88125160b21b6064820152608401610969565b80546060850151036136c55760405162461bcd60e51b815260206004820152603b60248201527f546f6b656e52656d6f74653a2063616e6e6f74206465706c6f7920746f20736160448201527f6d6520626c6f636b636861696e20617320746f6b656e20686f6d6500000000006064820152608401610969565b60808401516001600160a01b031661372b5760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a207a65726f20746f6b656e20686f6d65206164646044820152637265737360e01b6064820152608401610969565b60128460a0015160ff1611156137955760405162461bcd60e51b815260206004820152602960248201527f546f6b656e52656d6f74653a20746f6b656e20686f6d6520646563696d616c73604482015268040e8dede40d0d2ced60bb1b6064820152608401610969565b60128260ff1611156137f55760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a20746f6b656e20646563696d616c7320746f6f206044820152630d0d2ced60e31b6064820152608401610969565b60608401516001820155608084015160028201805460058401869055600684018054871560ff1990911617905560a08701516001600160a01b039093166001600160a81b031990911617600160a01b60ff808516919091029190911760ff60a81b1916600160a81b918616919091021790556138719083613ae1565b60048301805460ff1916911515919091179055600390910155505050565b6040516001600160a01b0384811660248301528381166044830152606482018390526113c99186918216906323b872dd90608401613276565b5f845a10156139195760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e742067617300000000006044820152606401610969565b834710156139695760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c75650000006044820152606401610969565b826001600160a01b03163b5f0361398157505f6117ae565b5f805f84516020860188888bf19695505050505050565b5f805f846001600160a01b0316846040516139b39190614cf8565b5f604051808303815f865af19150503d805f81146139ec576040519150601f19603f3d011682016040523d82523d5f602084013e6139f1565b606091505b5091509150818015613a1b575080511580613a1b575080806020019051810190613a1b9190614d13565b801561216c5750505050506001600160a01b03163b151590565b5f613a496001600160a01b03841683613b2b565b905080515f14158015613a6d575080806020019051810190613a6b9190614d13565b155b156115185782604051635274afe760e01b81526004016109699190613eed565b5f613aa162010203600160981b0183611669565b5090565b613aad612e0f565b613ab5613b38565b613abd613b48565b6114bf8282613b50565b613acf612e0f565b610a7881613cb5565b5f611f3b6116fc565b5f8060ff808516908416118181613b0457613afc8587614d32565b60ff16613b12565b613b0e8686614d32565b60ff165b613b1d90600a614e2b565b9350909150505b9250929050565b606061092983835f613cbd565b613b40612e0f565b6107ad613d4c565b6107ad612e0f565b613b58612e0f565b6001600160a01b038216613bce5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c65604482015276706f72746572207265676973747279206164647265737360481b6064820152608401610969565b5f613bd76115e8565b90505f8390505f816001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015613c1a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613c3e9190614410565b11613c935760405162461bcd60e51b815260206004820152603260248201525f80516020614e37833981519152604482015271656c65706f7274657220726567697374727960701b6064820152608401610969565b81546001600160a01b0319166001600160a01b0382161782556113c9836118d1565b611443612e0f565b606081471015613ce2573060405163cd78605960e01b81526004016109699190613eed565b5f80856001600160a01b03168486604051613cfd9190614cf8565b5f6040518083038185875af1925050503d805f8114613d37576040519150601f19603f3d011682016040523d82523d5f602084013e613d3c565b606091505b509150915061306f868383613d54565b611f18612e0f565b606082613d6957613d6482613da7565b610929565b8151158015613d8057506001600160a01b0384163b155b15613da05783604051639996b31560e01b81526004016109699190613eed565b5080610929565b805115613db75780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f5b83811015613dea578181015183820152602001613dd2565b50505f910152565b5f8151808452613e09816020860160208601613dd0565b601f01601f19169290920160200192915050565b602081525f6109296020830184613df2565b6001600160a01b0381168114610a78575f80fd5b8035613e4e81613e2f565b919050565b5f8060408385031215613e64575f80fd5b8235613e6f81613e2f565b946020939093013593505050565b5f805f60608486031215613e8f575f80fd5b8335613e9a81613e2f565b92506020840135613eaa81613e2f565b929592945050506040919091013590565b5f60208284031215613ecb575f80fd5b813561092981613e2f565b5f60208284031215613ee6575f80fd5b5035919050565b6001600160a01b0391909116815260200190565b6001600160a01b03169052565b5f60208284031215613f1e575f80fd5b81356001600160401b03811115613f33575f80fd5b82016101608185031215610929575f80fd5b5f6101008284031215613f56575f80fd5b50919050565b634e487b7160e01b5f52604160045260245ffd5b60405160c081016001600160401b0381118282101715613f9257613f92613f5c565b60405290565b604080519081016001600160401b0381118282101715613f9257613f92613f5c565b60405161010081016001600160401b0381118282101715613f9257613f92613f5c565b604051601f8201601f191681016001600160401b038111828210171561400557614005613f5c565b604052919050565b5f6001600160401b0382111561402557614025613f5c565b50601f01601f191660200190565b5f82601f830112614042575f80fd5b81356140556140508261400d565b613fdd565b818152846020838601011115614069575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f8084860361012081121561409a575f80fd5b60c08112156140a7575f80fd5b506140b0613f70565b85356140bb81613e2f565b815260208601356140cb81613e2f565b80602083015250604086013560408201526060860135606082015260808601356140f481613e2f565b608082015260a086013560ff8116811461410c575f80fd5b60a0820152935060c08501356001600160401b0381111561412b575f80fd5b61413787828801614033565b949794965050505060e0830135926101000135919050565b5f60408284031215613f56575f80fd5b5f805f8060608587031215614172575f80fd5b84359350602085013561418481613e2f565b925060408501356001600160401b038082111561419f575f80fd5b818701915087601f8301126141b2575f80fd5b8135818111156141c0575f80fd5b8860208285010111156141d1575f80fd5b95989497505060200194505050565b5f80604083850312156141f1575f80fd5b82356141fc81613e2f565b9150602083013561420c81613e2f565b809150509250929050565b600181811c9082168061422b57607f821691505b602082108103613f5657634e487b7160e01b5f52602260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b8082018082111561088b5761088b614249565b8181038181111561088b5761088b614249565b6020808252602e908201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b808202811582820484141761088b5761088b614249565b5f8261434657634e487b7160e01b5f52601260045260245ffd5b500490565b634e487b7160e01b5f52602160045260245ffd5b81516001600160a01b03168152602080830151908201526040810161088b565b602081525f8251600581106143a257634e487b7160e01b5f52602160045260245ffd5b8060208401525060208301516040808401526117ae6060840182613df2565b6020808252602f908201527f4e6174697665546f6b656e52656d6f74653a20636f6e747261637420756e646560408201526e1c98dbdb1b185d195c985b1a5e9959608a1b606082015260800190565b5f60208284031215614420575f80fd5b5051919050565b6001600160a01b039390931683526020830191909152604082015260600190565b6001600160a01b03929092168252602082015260400190565b6020808252825182820152828101516001600160a01b039081166040808501919091528401518051821660608501528083015160808501525f929161010085019190606087015160a0870152608087015160e060c08801528051938490528401925f92506101208701905b808410156144ee578451831682529385019360019390930192908501906144cc565b5060a0880151878203601f190160e0890152945061450c8186613df2565b98975050505050505050565b5f82601f830112614527575f80fd5b81516145356140508261400d565b818152846020838601011115614549575f80fd5b6117ae826020830160208701613dd0565b5f6020828403121561456a575f80fd5b81516001600160401b0380821115614580575f80fd5b9083019060408286031215614593575f80fd5b61459b613f98565b8251600581106145a9575f80fd5b81526020830151828111156145bc575f80fd5b6145c887828601614518565b60208301525095945050505050565b8051613e4e81613e2f565b5f604082840312156145f2575f80fd5b6145fa613f98565b825161460581613e2f565b81526020928301519281019290925250919050565b5f6020828403121561462a575f80fd5b81516001600160401b0380821115614640575f80fd5b908301906101008286031215614654575f80fd5b61465c613fba565b8251815261466c602084016145d7565b602082015261467d604084016145d7565b604082015261468e606084016145d7565b60608201526080830151608082015260a0830151828111156146ae575f80fd5b6146ba87828601614518565b60a08301525060c083015160c08201526146d660e084016145d7565b60e082015295945050505050565b5f602082840312156146f4575f80fd5b815161092981613e2f565b6020808252602b908201527f546f6b656e52656d6f74653a207a65726f2064657374696e6174696f6e20626c60408201526a1bd8dad8da185a5b88125160aa1b606082015260800190565b60208082526037908201527f546f6b656e52656d6f74653a207a65726f2064657374696e6174696f6e20746f6040820152766b656e207472616e73666572726572206164647265737360481b606082015260800190565b60208082526024908201527f546f6b656e52656d6f74653a207a65726f20726571756972656420676173206c6040820152631a5b5a5d60e21b606082015260800190565b5f808335601e198436030181126147fa575f80fd5b8301803591506001600160401b03821115614813575f80fd5b602001915036819003821315613b24575f80fd5b60208152815160208201525f602083015160018060a01b038082166040850152806040860151166060850152505060608301516148676080840182613f01565b50608083015160a083015260a08301516101008060c085015261488e610120850183613df2565b915060c085015160e085015260e08501516148ab82860182613f01565b5090949350505050565b5f808335601e198436030181126148ca575f80fd5b83016020810192503590506001600160401b038111156148e8575f80fd5b803603821315613b24575f80fd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60408152823560408201525f61493660208501613e43565b6149436060840182613f01565b5061495060408501613e43565b61495d6080840182613f01565b5061496b60608501856148b5565b6101608060a08601526149836101a0860183856148f6565b9250608087013560c086015260a087013560e08601526149a560c08801613e43565b91506101006149b681870184613f01565b6149c260e08901613e43565b92506101206149d381880185613f01565b6149de828a01613e43565b935061014091506149f182880185613f01565b880135918601919091529095013561018084015260209092019290925292915050565b60208152614a26602082018351613f01565b602082015160408201525f6040830151614a436060840182613f01565b506060830151614a566080840182613f01565b50608083015160a083015260a08301516101608060c0850152614a7d610180850183613df2565b915060c085015160e085015260e0850151610100614a9d81870183613f01565b860151610120868101919091528601519050610140614abe81870183613f01565b959095015193019290925250919050565b8235815261012081016020840135614ae681613e2f565b6001600160a01b039081166020840152604085013590614b0582613e2f565b166040830152614b1760608501613e43565b614b246060840182613f01565b506080840135608083015260a084013560a083015260c084013560c0830152614b4f60e08501613e43565b614b5c60e0840182613f01565b50826101008301529392505050565b8481526001600160a01b038481166020830152831660408201526080606082018190525f9061306f90830184613df2565b6020808252603a908201527f546f6b656e52656d6f74653a20696e76616c69642064657374696e6174696f6e60408201527f20746f6b656e207472616e736665727265722061646472657373000000000000606082015260800190565b601f82111561151857805f5260205f20601f840160051c81016020851015614c1e5750805b601f840160051c820191505b81811015611f11575f8155600101614c2a565b81516001600160401b03811115614c5657614c56613f5c565b614c6a81614c648454614217565b84614bf9565b602080601f831160018114614c9d575f8415614c865750858301515b5f19600386901b1c1916600185901b178555611211565b5f85815260208120601f198616915b82811015614ccb57888601518255948401946001909101908401614cac565b5085821015614ce857878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f8251614d09818460208701613dd0565b9190910192915050565b5f60208284031215614d23575f80fd5b81518015158114610929575f80fd5b60ff828116828216039081111561088b5761088b614249565b600181815b80851115614d8557815f1904821115614d6b57614d6b614249565b80851615614d7857918102915b93841c9390800290614d50565b509250929050565b5f82614d9b5750600161088b565b81614da757505f61088b565b8160018114614dbd5760028114614dc757614de3565b600191505061088b565b60ff841115614dd857614dd8614249565b50506001821b61088b565b5060208310610133831016604e8410600b8410161715614e06575081810a61088b565b614e108383614d4b565b805f1904821115614e2357614e23614249565b029392505050565b5f6109298383614d8d56fe54656c65706f7274657252656769737472794170703a20696e76616c69642054a164736f6c6343000819000a",
}

// NativeTokenRemoteUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenRemoteUpgradeableMetaData.ABI instead.
var NativeTokenRemoteUpgradeableABI = NativeTokenRemoteUpgradeableMetaData.ABI

// NativeTokenRemoteUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenRemoteUpgradeableMetaData.Bin instead.
var NativeTokenRemoteUpgradeableBin = NativeTokenRemoteUpgradeableMetaData.Bin

// DeployNativeTokenRemoteUpgradeable deploys a new Ethereum contract, binding an instance of NativeTokenRemoteUpgradeable to it.
func DeployNativeTokenRemoteUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *NativeTokenRemoteUpgradeable, error) {
	parsed, err := NativeTokenRemoteUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenRemoteUpgradeableBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenRemoteUpgradeable{NativeTokenRemoteUpgradeableCaller: NativeTokenRemoteUpgradeableCaller{contract: contract}, NativeTokenRemoteUpgradeableTransactor: NativeTokenRemoteUpgradeableTransactor{contract: contract}, NativeTokenRemoteUpgradeableFilterer: NativeTokenRemoteUpgradeableFilterer{contract: contract}}, nil
}

// NativeTokenRemoteUpgradeable is an auto generated Go binding around an Ethereum contract.
type NativeTokenRemoteUpgradeable struct {
	NativeTokenRemoteUpgradeableCaller     // Read-only binding to the contract
	NativeTokenRemoteUpgradeableTransactor // Write-only binding to the contract
	NativeTokenRemoteUpgradeableFilterer   // Log filterer for contract events
}

// NativeTokenRemoteUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenRemoteUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenRemoteUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenRemoteUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenRemoteUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenRemoteUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenRemoteUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenRemoteUpgradeableSession struct {
	Contract     *NativeTokenRemoteUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// NativeTokenRemoteUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenRemoteUpgradeableCallerSession struct {
	Contract *NativeTokenRemoteUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// NativeTokenRemoteUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenRemoteUpgradeableTransactorSession struct {
	Contract     *NativeTokenRemoteUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// NativeTokenRemoteUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenRemoteUpgradeableRaw struct {
	Contract *NativeTokenRemoteUpgradeable // Generic contract binding to access the raw methods on
}

// NativeTokenRemoteUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenRemoteUpgradeableCallerRaw struct {
	Contract *NativeTokenRemoteUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenRemoteUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenRemoteUpgradeableTransactorRaw struct {
	Contract *NativeTokenRemoteUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenRemoteUpgradeable creates a new instance of NativeTokenRemoteUpgradeable, bound to a specific deployed contract.
func NewNativeTokenRemoteUpgradeable(address common.Address, backend bind.ContractBackend) (*NativeTokenRemoteUpgradeable, error) {
	contract, err := bindNativeTokenRemoteUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeable{NativeTokenRemoteUpgradeableCaller: NativeTokenRemoteUpgradeableCaller{contract: contract}, NativeTokenRemoteUpgradeableTransactor: NativeTokenRemoteUpgradeableTransactor{contract: contract}, NativeTokenRemoteUpgradeableFilterer: NativeTokenRemoteUpgradeableFilterer{contract: contract}}, nil
}

// NewNativeTokenRemoteUpgradeableCaller creates a new read-only instance of NativeTokenRemoteUpgradeable, bound to a specific deployed contract.
func NewNativeTokenRemoteUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenRemoteUpgradeableCaller, error) {
	contract, err := bindNativeTokenRemoteUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableCaller{contract: contract}, nil
}

// NewNativeTokenRemoteUpgradeableTransactor creates a new write-only instance of NativeTokenRemoteUpgradeable, bound to a specific deployed contract.
func NewNativeTokenRemoteUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenRemoteUpgradeableTransactor, error) {
	contract, err := bindNativeTokenRemoteUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableTransactor{contract: contract}, nil
}

// NewNativeTokenRemoteUpgradeableFilterer creates a new log filterer instance of NativeTokenRemoteUpgradeable, bound to a specific deployed contract.
func NewNativeTokenRemoteUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenRemoteUpgradeableFilterer, error) {
	contract, err := bindNativeTokenRemoteUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableFilterer{contract: contract}, nil
}

// bindNativeTokenRemoteUpgradeable binds a generic wrapper to an already deployed contract.
func bindNativeTokenRemoteUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenRemoteUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenRemoteUpgradeable.Contract.NativeTokenRemoteUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.NativeTokenRemoteUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.NativeTokenRemoteUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenRemoteUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// BURNEDFORTRANSFERADDRESS is a free data retrieval call binding the contract method 0x347212c4.
//
// Solidity: function BURNED_FOR_TRANSFER_ADDRESS() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) BURNEDFORTRANSFERADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "BURNED_FOR_TRANSFER_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDFORTRANSFERADDRESS is a free data retrieval call binding the contract method 0x347212c4.
//
// Solidity: function BURNED_FOR_TRANSFER_ADDRESS() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) BURNEDFORTRANSFERADDRESS() (common.Address, error) {
	return _NativeTokenRemoteUpgradeable.Contract.BURNEDFORTRANSFERADDRESS(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// BURNEDFORTRANSFERADDRESS is a free data retrieval call binding the contract method 0x347212c4.
//
// Solidity: function BURNED_FOR_TRANSFER_ADDRESS() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) BURNEDFORTRANSFERADDRESS() (common.Address, error) {
	return _NativeTokenRemoteUpgradeable.Contract.BURNEDFORTRANSFERADDRESS(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) BURNEDTXFEESADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "BURNED_TX_FEES_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenRemoteUpgradeable.Contract.BURNEDTXFEESADDRESS(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenRemoteUpgradeable.Contract.BURNEDTXFEESADDRESS(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// HOMECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xed0ae4b0.
//
// Solidity: function HOME_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) HOMECHAINBURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "HOME_CHAIN_BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HOMECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xed0ae4b0.
//
// Solidity: function HOME_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) HOMECHAINBURNADDRESS() (common.Address, error) {
	return _NativeTokenRemoteUpgradeable.Contract.HOMECHAINBURNADDRESS(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// HOMECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xed0ae4b0.
//
// Solidity: function HOME_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) HOMECHAINBURNADDRESS() (common.Address, error) {
	return _NativeTokenRemoteUpgradeable.Contract.HOMECHAINBURNADDRESS(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) MULTIHOPCALLGASPERWORD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "MULTI_HOP_CALL_GAS_PER_WORD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.MULTIHOPCALLGASPERWORD(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.MULTIHOPCALLGASPERWORD(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) MULTIHOPCALLREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "MULTI_HOP_CALL_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) MULTIHOPCALLREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.MULTIHOPCALLREQUIREDGAS(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) MULTIHOPCALLREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.MULTIHOPCALLREQUIREDGAS(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) MULTIHOPSENDREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "MULTI_HOP_SEND_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) MULTIHOPSENDREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.MULTIHOPSENDREQUIREDGAS(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) MULTIHOPSENDREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.MULTIHOPSENDREQUIREDGAS(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) NATIVEMINTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "NATIVE_MINTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenRemoteUpgradeable.Contract.NATIVEMINTER(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenRemoteUpgradeable.Contract.NATIVEMINTER(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// NATIVETOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0xa5717bc0.
//
// Solidity: function NATIVE_TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) NATIVETOKENREMOTESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "NATIVE_TOKEN_REMOTE_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NATIVETOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0xa5717bc0.
//
// Solidity: function NATIVE_TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) NATIVETOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenRemoteUpgradeable.Contract.NATIVETOKENREMOTESTORAGELOCATION(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// NATIVETOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0xa5717bc0.
//
// Solidity: function NATIVE_TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) NATIVETOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenRemoteUpgradeable.Contract.NATIVETOKENREMOTESTORAGELOCATION(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) REGISTERREMOTEREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "REGISTER_REMOTE_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) REGISTERREMOTEREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.REGISTERREMOTEREQUIREDGAS(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) REGISTERREMOTEREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.REGISTERREMOTEREQUIREDGAS(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) TELEPORTERREGISTRYAPPSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "TELEPORTER_REGISTRY_APP_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenRemoteUpgradeable.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenRemoteUpgradeable.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x35cac159.
//
// Solidity: function TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) TOKENREMOTESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "TOKEN_REMOTE_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x35cac159.
//
// Solidity: function TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) TOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenRemoteUpgradeable.Contract.TOKENREMOTESTORAGELOCATION(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x35cac159.
//
// Solidity: function TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) TOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenRemoteUpgradeable.Contract.TOKENREMOTESTORAGELOCATION(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Allowance(&_NativeTokenRemoteUpgradeable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Allowance(&_NativeTokenRemoteUpgradeable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.BalanceOf(&_NativeTokenRemoteUpgradeable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.BalanceOf(&_NativeTokenRemoteUpgradeable.CallOpts, account)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) CalculateNumWords(opts *bind.CallOpts, payloadSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "calculateNumWords", payloadSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.CalculateNumWords(&_NativeTokenRemoteUpgradeable.CallOpts, payloadSize)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.CalculateNumWords(&_NativeTokenRemoteUpgradeable.CallOpts, payloadSize)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) Decimals() (uint8, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Decimals(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) Decimals() (uint8, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Decimals(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) GetBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "getBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) GetBlockchainID() ([32]byte, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetBlockchainID(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) GetBlockchainID() ([32]byte, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetBlockchainID(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) GetInitialReserveImbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "getInitialReserveImbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) GetInitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetInitialReserveImbalance(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) GetInitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetInitialReserveImbalance(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) GetIsCollateralized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "getIsCollateralized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) GetIsCollateralized() (bool, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetIsCollateralized(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) GetIsCollateralized() (bool, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetIsCollateralized(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetMinTeleporterVersion(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetMinTeleporterVersion(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) GetMultiplyOnRemote(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "getMultiplyOnRemote")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) GetMultiplyOnRemote() (bool, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetMultiplyOnRemote(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) GetMultiplyOnRemote() (bool, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetMultiplyOnRemote(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) GetTokenHomeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "getTokenHomeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) GetTokenHomeAddress() (common.Address, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetTokenHomeAddress(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) GetTokenHomeAddress() (common.Address, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetTokenHomeAddress(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) GetTokenHomeBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "getTokenHomeBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) GetTokenHomeBlockchainID() ([32]byte, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetTokenHomeBlockchainID(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) GetTokenHomeBlockchainID() ([32]byte, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetTokenHomeBlockchainID(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) GetTokenMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "getTokenMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) GetTokenMultiplier() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetTokenMultiplier(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) GetTokenMultiplier() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetTokenMultiplier(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetTotalMinted is a free data retrieval call binding the contract method 0x0ca1c5c9.
//
// Solidity: function getTotalMinted() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) GetTotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "getTotalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalMinted is a free data retrieval call binding the contract method 0x0ca1c5c9.
//
// Solidity: function getTotalMinted() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) GetTotalMinted() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetTotalMinted(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// GetTotalMinted is a free data retrieval call binding the contract method 0x0ca1c5c9.
//
// Solidity: function getTotalMinted() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) GetTotalMinted() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.GetTotalMinted(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenRemoteUpgradeable.Contract.IsTeleporterAddressPaused(&_NativeTokenRemoteUpgradeable.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenRemoteUpgradeable.Contract.IsTeleporterAddressPaused(&_NativeTokenRemoteUpgradeable.CallOpts, teleporterAddress)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) Name() (string, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Name(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) Name() (string, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Name(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) Owner() (common.Address, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Owner(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) Owner() (common.Address, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Owner(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) Symbol() (string, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Symbol(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) Symbol() (string, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Symbol(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) TotalNativeAssetSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "totalNativeAssetSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) TotalNativeAssetSupply() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.TotalNativeAssetSupply(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) TotalNativeAssetSupply() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.TotalNativeAssetSupply(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemoteUpgradeable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.TotalSupply(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableCallerSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenRemoteUpgradeable.Contract.TotalSupply(&_NativeTokenRemoteUpgradeable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Approve(&_NativeTokenRemoteUpgradeable.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Approve(&_NativeTokenRemoteUpgradeable.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) Deposit() (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Deposit(&_NativeTokenRemoteUpgradeable.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) Deposit() (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Deposit(&_NativeTokenRemoteUpgradeable.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f6cec88.
//
// Solidity: function initialize((address,address,uint256,bytes32,address,uint8) settings, string nativeAssetSymbol, uint256 initialReserveImbalance, uint256 burnedFeesReportingRewardPercentage) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) Initialize(opts *bind.TransactOpts, settings TokenRemoteSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "initialize", settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f6cec88.
//
// Solidity: function initialize((address,address,uint256,bytes32,address,uint8) settings, string nativeAssetSymbol, uint256 initialReserveImbalance, uint256 burnedFeesReportingRewardPercentage) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) Initialize(settings TokenRemoteSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Initialize(&_NativeTokenRemoteUpgradeable.TransactOpts, settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f6cec88.
//
// Solidity: function initialize((address,address,uint256,bytes32,address,uint8) settings, string nativeAssetSymbol, uint256 initialReserveImbalance, uint256 burnedFeesReportingRewardPercentage) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) Initialize(settings TokenRemoteSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Initialize(&_NativeTokenRemoteUpgradeable.TransactOpts, settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.PauseTeleporterAddress(&_NativeTokenRemoteUpgradeable.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.PauseTeleporterAddress(&_NativeTokenRemoteUpgradeable.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.ReceiveTeleporterMessage(&_NativeTokenRemoteUpgradeable.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.ReceiveTeleporterMessage(&_NativeTokenRemoteUpgradeable.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) RegisterWithHome(opts *bind.TransactOpts, feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "registerWithHome", feeInfo)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) RegisterWithHome(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.RegisterWithHome(&_NativeTokenRemoteUpgradeable.TransactOpts, feeInfo)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) RegisterWithHome(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.RegisterWithHome(&_NativeTokenRemoteUpgradeable.TransactOpts, feeInfo)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.RenounceOwnership(&_NativeTokenRemoteUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.RenounceOwnership(&_NativeTokenRemoteUpgradeable.TransactOpts)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) ReportBurnedTxFees(opts *bind.TransactOpts, requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "reportBurnedTxFees", requiredGasLimit)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) ReportBurnedTxFees(requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.ReportBurnedTxFees(&_NativeTokenRemoteUpgradeable.TransactOpts, requiredGasLimit)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) ReportBurnedTxFees(requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.ReportBurnedTxFees(&_NativeTokenRemoteUpgradeable.TransactOpts, requiredGasLimit)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) Send(opts *bind.TransactOpts, input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Send(&_NativeTokenRemoteUpgradeable.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Send(&_NativeTokenRemoteUpgradeable.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) SendAndCall(opts *bind.TransactOpts, input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "sendAndCall", input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.SendAndCall(&_NativeTokenRemoteUpgradeable.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.SendAndCall(&_NativeTokenRemoteUpgradeable.TransactOpts, input)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Transfer(&_NativeTokenRemoteUpgradeable.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Transfer(&_NativeTokenRemoteUpgradeable.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.TransferFrom(&_NativeTokenRemoteUpgradeable.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.TransferFrom(&_NativeTokenRemoteUpgradeable.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.TransferOwnership(&_NativeTokenRemoteUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.TransferOwnership(&_NativeTokenRemoteUpgradeable.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.UnpauseTeleporterAddress(&_NativeTokenRemoteUpgradeable.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.UnpauseTeleporterAddress(&_NativeTokenRemoteUpgradeable.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.UpdateMinTeleporterVersion(&_NativeTokenRemoteUpgradeable.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.UpdateMinTeleporterVersion(&_NativeTokenRemoteUpgradeable.TransactOpts, version)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Withdraw(&_NativeTokenRemoteUpgradeable.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Withdraw(&_NativeTokenRemoteUpgradeable.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Fallback(&_NativeTokenRemoteUpgradeable.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Fallback(&_NativeTokenRemoteUpgradeable.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) Receive() (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Receive(&_NativeTokenRemoteUpgradeable.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) Receive() (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Receive(&_NativeTokenRemoteUpgradeable.TransactOpts)
}

// NativeTokenRemoteUpgradeableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableApprovalIterator struct {
	Event *NativeTokenRemoteUpgradeableApproval // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteUpgradeableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteUpgradeableApproval)
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
		it.Event = new(NativeTokenRemoteUpgradeableApproval)
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
func (it *NativeTokenRemoteUpgradeableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteUpgradeableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteUpgradeableApproval represents a Approval event raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NativeTokenRemoteUpgradeableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableApprovalIterator{contract: _NativeTokenRemoteUpgradeable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteUpgradeableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteUpgradeableApproval)
				if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) ParseApproval(log types.Log) (*NativeTokenRemoteUpgradeableApproval, error) {
	event := new(NativeTokenRemoteUpgradeableApproval)
	if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteUpgradeableCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableCallFailedIterator struct {
	Event *NativeTokenRemoteUpgradeableCallFailed // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteUpgradeableCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteUpgradeableCallFailed)
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
		it.Event = new(NativeTokenRemoteUpgradeableCallFailed)
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
func (it *NativeTokenRemoteUpgradeableCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteUpgradeableCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteUpgradeableCallFailed represents a CallFailed event raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableCallFailed struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) FilterCallFailed(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenRemoteUpgradeableCallFailedIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.FilterLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableCallFailedIterator{contract: _NativeTokenRemoteUpgradeable.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteUpgradeableCallFailed, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.WatchLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteUpgradeableCallFailed)
				if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "CallFailed", log); err != nil {
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
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) ParseCallFailed(log types.Log) (*NativeTokenRemoteUpgradeableCallFailed, error) {
	event := new(NativeTokenRemoteUpgradeableCallFailed)
	if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteUpgradeableCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableCallSucceededIterator struct {
	Event *NativeTokenRemoteUpgradeableCallSucceeded // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteUpgradeableCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteUpgradeableCallSucceeded)
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
		it.Event = new(NativeTokenRemoteUpgradeableCallSucceeded)
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
func (it *NativeTokenRemoteUpgradeableCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteUpgradeableCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteUpgradeableCallSucceeded represents a CallSucceeded event raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableCallSucceeded struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) FilterCallSucceeded(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenRemoteUpgradeableCallSucceededIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.FilterLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableCallSucceededIterator{contract: _NativeTokenRemoteUpgradeable.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteUpgradeableCallSucceeded, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.WatchLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteUpgradeableCallSucceeded)
				if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) ParseCallSucceeded(log types.Log) (*NativeTokenRemoteUpgradeableCallSucceeded, error) {
	event := new(NativeTokenRemoteUpgradeableCallSucceeded)
	if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteUpgradeableDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableDepositIterator struct {
	Event *NativeTokenRemoteUpgradeableDeposit // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteUpgradeableDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteUpgradeableDeposit)
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
		it.Event = new(NativeTokenRemoteUpgradeableDeposit)
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
func (it *NativeTokenRemoteUpgradeableDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteUpgradeableDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteUpgradeableDeposit represents a Deposit event raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenRemoteUpgradeableDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableDepositIterator{contract: _NativeTokenRemoteUpgradeable.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteUpgradeableDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteUpgradeableDeposit)
				if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) ParseDeposit(log types.Log) (*NativeTokenRemoteUpgradeableDeposit, error) {
	event := new(NativeTokenRemoteUpgradeableDeposit)
	if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableInitializedIterator struct {
	Event *NativeTokenRemoteUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteUpgradeableInitialized)
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
		it.Event = new(NativeTokenRemoteUpgradeableInitialized)
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
func (it *NativeTokenRemoteUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteUpgradeableInitialized represents a Initialized event raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*NativeTokenRemoteUpgradeableInitializedIterator, error) {

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableInitializedIterator{contract: _NativeTokenRemoteUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteUpgradeableInitialized)
				if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) ParseInitialized(log types.Log) (*NativeTokenRemoteUpgradeableInitialized, error) {
	event := new(NativeTokenRemoteUpgradeableInitialized)
	if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator struct {
	Event *NativeTokenRemoteUpgradeableMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteUpgradeableMinTeleporterVersionUpdated)
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
		it.Event = new(NativeTokenRemoteUpgradeableMinTeleporterVersionUpdated)
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
func (it *NativeTokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteUpgradeableMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*NativeTokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator{contract: _NativeTokenRemoteUpgradeable.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteUpgradeableMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteUpgradeableMinTeleporterVersionUpdated)
				if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*NativeTokenRemoteUpgradeableMinTeleporterVersionUpdated, error) {
	event := new(NativeTokenRemoteUpgradeableMinTeleporterVersionUpdated)
	if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableOwnershipTransferredIterator struct {
	Event *NativeTokenRemoteUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteUpgradeableOwnershipTransferred)
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
		it.Event = new(NativeTokenRemoteUpgradeableOwnershipTransferred)
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
func (it *NativeTokenRemoteUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NativeTokenRemoteUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableOwnershipTransferredIterator{contract: _NativeTokenRemoteUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteUpgradeableOwnershipTransferred)
				if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*NativeTokenRemoteUpgradeableOwnershipTransferred, error) {
	event := new(NativeTokenRemoteUpgradeableOwnershipTransferred)
	if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteUpgradeableReportBurnedTxFeesIterator is returned from FilterReportBurnedTxFees and is used to iterate over the raw logs and unpacked data for ReportBurnedTxFees events raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableReportBurnedTxFeesIterator struct {
	Event *NativeTokenRemoteUpgradeableReportBurnedTxFees // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteUpgradeableReportBurnedTxFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteUpgradeableReportBurnedTxFees)
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
		it.Event = new(NativeTokenRemoteUpgradeableReportBurnedTxFees)
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
func (it *NativeTokenRemoteUpgradeableReportBurnedTxFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteUpgradeableReportBurnedTxFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteUpgradeableReportBurnedTxFees represents a ReportBurnedTxFees event raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableReportBurnedTxFees struct {
	TeleporterMessageID [32]byte
	FeesBurned          *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReportBurnedTxFees is a free log retrieval operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) FilterReportBurnedTxFees(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*NativeTokenRemoteUpgradeableReportBurnedTxFeesIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.FilterLogs(opts, "ReportBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableReportBurnedTxFeesIterator{contract: _NativeTokenRemoteUpgradeable.contract, event: "ReportBurnedTxFees", logs: logs, sub: sub}, nil
}

// WatchReportBurnedTxFees is a free log subscription operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) WatchReportBurnedTxFees(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteUpgradeableReportBurnedTxFees, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.WatchLogs(opts, "ReportBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteUpgradeableReportBurnedTxFees)
				if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "ReportBurnedTxFees", log); err != nil {
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

// ParseReportBurnedTxFees is a log parse operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) ParseReportBurnedTxFees(log types.Log) (*NativeTokenRemoteUpgradeableReportBurnedTxFees, error) {
	event := new(NativeTokenRemoteUpgradeableReportBurnedTxFees)
	if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "ReportBurnedTxFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteUpgradeableTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableTeleporterAddressPausedIterator struct {
	Event *NativeTokenRemoteUpgradeableTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteUpgradeableTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteUpgradeableTeleporterAddressPaused)
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
		it.Event = new(NativeTokenRemoteUpgradeableTeleporterAddressPaused)
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
func (it *NativeTokenRemoteUpgradeableTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteUpgradeableTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteUpgradeableTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenRemoteUpgradeableTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableTeleporterAddressPausedIterator{contract: _NativeTokenRemoteUpgradeable.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteUpgradeableTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteUpgradeableTeleporterAddressPaused)
				if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) ParseTeleporterAddressPaused(log types.Log) (*NativeTokenRemoteUpgradeableTeleporterAddressPaused, error) {
	event := new(NativeTokenRemoteUpgradeableTeleporterAddressPaused)
	if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteUpgradeableTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableTeleporterAddressUnpausedIterator struct {
	Event *NativeTokenRemoteUpgradeableTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteUpgradeableTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteUpgradeableTeleporterAddressUnpaused)
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
		it.Event = new(NativeTokenRemoteUpgradeableTeleporterAddressUnpaused)
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
func (it *NativeTokenRemoteUpgradeableTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteUpgradeableTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteUpgradeableTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenRemoteUpgradeableTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableTeleporterAddressUnpausedIterator{contract: _NativeTokenRemoteUpgradeable.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteUpgradeableTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteUpgradeableTeleporterAddressUnpaused)
				if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*NativeTokenRemoteUpgradeableTeleporterAddressUnpaused, error) {
	event := new(NativeTokenRemoteUpgradeableTeleporterAddressUnpaused)
	if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteUpgradeableTokensAndCallSentIterator is returned from FilterTokensAndCallSent and is used to iterate over the raw logs and unpacked data for TokensAndCallSent events raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableTokensAndCallSentIterator struct {
	Event *NativeTokenRemoteUpgradeableTokensAndCallSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteUpgradeableTokensAndCallSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteUpgradeableTokensAndCallSent)
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
		it.Event = new(NativeTokenRemoteUpgradeableTokensAndCallSent)
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
func (it *NativeTokenRemoteUpgradeableTokensAndCallSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteUpgradeableTokensAndCallSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteUpgradeableTokensAndCallSent represents a TokensAndCallSent event raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableTokensAndCallSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) FilterTokensAndCallSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenRemoteUpgradeableTokensAndCallSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.FilterLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableTokensAndCallSentIterator{contract: _NativeTokenRemoteUpgradeable.contract, event: "TokensAndCallSent", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) WatchTokensAndCallSent(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteUpgradeableTokensAndCallSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.WatchLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteUpgradeableTokensAndCallSent)
				if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
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
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) ParseTokensAndCallSent(log types.Log) (*NativeTokenRemoteUpgradeableTokensAndCallSent, error) {
	event := new(NativeTokenRemoteUpgradeableTokensAndCallSent)
	if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteUpgradeableTokensSentIterator is returned from FilterTokensSent and is used to iterate over the raw logs and unpacked data for TokensSent events raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableTokensSentIterator struct {
	Event *NativeTokenRemoteUpgradeableTokensSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteUpgradeableTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteUpgradeableTokensSent)
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
		it.Event = new(NativeTokenRemoteUpgradeableTokensSent)
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
func (it *NativeTokenRemoteUpgradeableTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteUpgradeableTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteUpgradeableTokensSent represents a TokensSent event raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensSent is a free log retrieval operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) FilterTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenRemoteUpgradeableTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.FilterLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableTokensSentIterator{contract: _NativeTokenRemoteUpgradeable.contract, event: "TokensSent", logs: logs, sub: sub}, nil
}

// WatchTokensSent is a free log subscription operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) WatchTokensSent(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteUpgradeableTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.WatchLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteUpgradeableTokensSent)
				if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "TokensSent", log); err != nil {
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
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) ParseTokensSent(log types.Log) (*NativeTokenRemoteUpgradeableTokensSent, error) {
	event := new(NativeTokenRemoteUpgradeableTokensSent)
	if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "TokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteUpgradeableTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableTokensWithdrawnIterator struct {
	Event *NativeTokenRemoteUpgradeableTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteUpgradeableTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteUpgradeableTokensWithdrawn)
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
		it.Event = new(NativeTokenRemoteUpgradeableTokensWithdrawn)
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
func (it *NativeTokenRemoteUpgradeableTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteUpgradeableTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteUpgradeableTokensWithdrawn represents a TokensWithdrawn event raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableTokensWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*NativeTokenRemoteUpgradeableTokensWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.FilterLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableTokensWithdrawnIterator{contract: _NativeTokenRemoteUpgradeable.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteUpgradeableTokensWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.WatchLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteUpgradeableTokensWithdrawn)
				if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) ParseTokensWithdrawn(log types.Log) (*NativeTokenRemoteUpgradeableTokensWithdrawn, error) {
	event := new(NativeTokenRemoteUpgradeableTokensWithdrawn)
	if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteUpgradeableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableTransferIterator struct {
	Event *NativeTokenRemoteUpgradeableTransfer // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteUpgradeableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteUpgradeableTransfer)
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
		it.Event = new(NativeTokenRemoteUpgradeableTransfer)
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
func (it *NativeTokenRemoteUpgradeableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteUpgradeableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteUpgradeableTransfer represents a Transfer event raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTokenRemoteUpgradeableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableTransferIterator{contract: _NativeTokenRemoteUpgradeable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteUpgradeableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteUpgradeableTransfer)
				if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) ParseTransfer(log types.Log) (*NativeTokenRemoteUpgradeableTransfer, error) {
	event := new(NativeTokenRemoteUpgradeableTransfer)
	if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteUpgradeableWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableWithdrawalIterator struct {
	Event *NativeTokenRemoteUpgradeableWithdrawal // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteUpgradeableWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteUpgradeableWithdrawal)
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
		it.Event = new(NativeTokenRemoteUpgradeableWithdrawal)
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
func (it *NativeTokenRemoteUpgradeableWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteUpgradeableWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteUpgradeableWithdrawal represents a Withdrawal event raised by the NativeTokenRemoteUpgradeable contract.
type NativeTokenRemoteUpgradeableWithdrawal struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) FilterWithdrawal(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenRemoteUpgradeableWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.FilterLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteUpgradeableWithdrawalIterator{contract: _NativeTokenRemoteUpgradeable.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteUpgradeableWithdrawal, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemoteUpgradeable.contract.WatchLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteUpgradeableWithdrawal)
				if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableFilterer) ParseWithdrawal(log types.Log) (*NativeTokenRemoteUpgradeableWithdrawal, error) {
	event := new(NativeTokenRemoteUpgradeableWithdrawal)
	if err := _NativeTokenRemoteUpgradeable.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
