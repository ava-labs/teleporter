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
	TokenHomeBlockchainID     [32]byte
	TokenHomeAddress          common.Address
	TokenHomeDecimals         uint8
}

// NativeTokenRemoteUpgradeableMetaData contains all meta data concerning the NativeTokenRemoteUpgradeable contract.
var NativeTokenRemoteUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"init\",\"type\":\"uint8\",\"internalType\":\"enumInitializable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"BURNED_FOR_TRANSFER_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BURNED_TX_FEES_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"HOME_CHAIN_BURN_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MULTI_HOP_CALL_GAS_PER_WORD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MULTI_HOP_CALL_REQUIRED_GAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MULTI_HOP_SEND_REQUIRED_GAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NATIVE_MINTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINativeMinter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NATIVE_TOKEN_REMOTE_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTER_REMOTE_REQUIRED_GAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TELEPORTER_REGISTRY_APP_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_REMOTE_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateNumWords\",\"inputs\":[{\"name\":\"payloadSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getBlockchainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInitialReserveImbalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIsCollateralized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinTeleporterVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMultiplyOnRemote\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenHomeAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenHomeBlockchainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenMultiplier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalMinted\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"settings\",\"type\":\"tuple\",\"internalType\":\"structTokenRemoteSettings\",\"components\":[{\"name\":\"teleporterRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"teleporterManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenHomeBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenHomeAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenHomeDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"nativeAssetSymbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"initialReserveImbalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"burnedFeesReportingRewardPercentage_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isTeleporterAddressPaused\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseTeleporterAddress\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"receiveTeleporterMessage\",\"inputs\":[{\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originSenderAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerWithHome\",\"inputs\":[{\"name\":\"feeInfo\",\"type\":\"tuple\",\"internalType\":\"structTeleporterFeeInfo\",\"components\":[{\"name\":\"feeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reportBurnedTxFees\",\"inputs\":[{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"send\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structSendTokensInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"sendAndCall\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structSendAndCallInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipientGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalNativeAssetSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseTeleporterAddress\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMinTeleporterVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"recipientContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSucceeded\",\"inputs\":[{\"name\":\"recipientContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinTeleporterVersionUpdated\",\"inputs\":[{\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReportBurnedTxFees\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"feesBurned\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TeleporterAddressPaused\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TeleporterAddressUnpaused\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensAndCallSent\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"input\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipientGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensSent\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"input\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161507238038061507283398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b614f258061014d5f395ff3fe608060405260043610610280575f3560e01c8063715018a61161014e578063c3cd6927116100c0578063e0fd9cb811610079578063e0fd9cb814610736578063ed0ae4b014610452578063ef793e2a1461074a578063f2fde38b1461075e578063f3f981d81461077d578063f56c363d1461079c5761028f565b8063c3cd6927146106b9578063c452165e146106cd578063c868efaa146106e4578063d0e30db01461028f578063d2cc7a7014610703578063dd62ed3e146107175761028f565b8063909a6ac011610112578063909a6ac01461060857806395d89b4114610628578063973142971461063c578063a5717bc01461065b578063a9059cbb1461067b578063b8a46d021461069a5761028f565b8063715018a61461057b57806371717c181461058f5780637ee3779a146105a55780638bf2fa94146105b95780638da5cb5b146105cc5761028f565b80632e1a7d4d116101f25780634511243e116101ab5780634511243e146104b55780635507f3d1146104d457806355538c8b146104ea5780635eb99514146105095780636e6eef8d1461052857806370a082311461053b5761028f565b80632e1a7d4d146103e6578063313ce56714610405578063329c3e1214610420578063347212c41461045257806335cac1591461046e5780634213cf78146104a15761028f565b806315beb59f1161024457806315beb59f1461035557806318160ddd1461036a5780631906529c1461037e57806323b872dd14610392578063254ac160146103b15780632b0d8f18146103c75761028f565b806302a30c7d1461029757806306fdde03146102c05780630733c8c8146102e1578063095ea7b3146103035780630ca1c5c9146103225761028f565b3661028f5761028d6107bb565b005b61028d6107bb565b3480156102a2575f80fd5b506102ab6107fc565b60405190151581526020015b60405180910390f35b3480156102cb575f80fd5b506102d4610813565b6040516102b79190613e94565b3480156102ec575f80fd5b506102f56108d3565b6040519081526020016102b7565b34801561030e575f80fd5b506102ab61031d366004613eca565b6108e7565b34801561032d575f80fd5b507f69a5f7616543528c4fbe43f410b1034bd6da4ba06c25bedf04617268014cf501546102f5565b348015610360575f80fd5b506102f56105dc81565b348015610375575f80fd5b506102f5610900565b348015610389575f80fd5b506102f561091b565b34801561039d575f80fd5b506102ab6103ac366004613ef4565b610972565b3480156103bc575f80fd5b506102f56201fbd081565b3480156103d2575f80fd5b5061028d6103e1366004613f32565b610997565b3480156103f1575f80fd5b5061028d610400366004613f4d565b610a99565b348015610410575f80fd5b50604051601281526020016102b7565b34801561042b575f80fd5b5061043a6001600160991b0181565b6040516001600160a01b0390911681526020016102b7565b34801561045d575f80fd5b5061043a62010203600160981b0181565b348015610479575f80fd5b506102f57f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef0081565b3480156104ac575f80fd5b506102f5610ae5565b3480156104c0575f80fd5b5061028d6104cf366004613f32565b610af6565b3480156104df575f80fd5b506102f56205302081565b3480156104f5575f80fd5b5061028d610504366004613f4d565b610be5565b348015610514575f80fd5b5061028d610523366004613f4d565b610ef7565b61028d610536366004613f64565b610f08565b348015610546575f80fd5b506102f5610555366004613f32565b6001600160a01b03165f9081525f80516020614e99833981519152602052604090205490565b348015610586575f80fd5b5061028d610f36565b34801561059a575f80fd5b506102f56205573081565b3480156105b0575f80fd5b506102ab610f47565b61028d6105c7366004613f9b565b610f5e565b3480156105d7575f80fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661043a565b348015610613575f80fd5b506102f55f80516020614eb983398151915281565b348015610633575f80fd5b506102d4610f8c565b348015610647575f80fd5b506102ab610656366004613f32565b610fca565b348015610666575f80fd5b506102f55f80516020614ef983398151915281565b348015610686575f80fd5b506102ab610695366004613eca565b610fe3565b3480156106a5575f80fd5b5061028d6106b4366004613fb2565b610ff0565b3480156106c4575f80fd5b5061043a611164565b3480156106d8575f80fd5b5061043a600160981b81565b3480156106ef575f80fd5b5061028d6106fe366004613fc2565b611181565b34801561070e575f80fd5b506102f561133e565b348015610722575f80fd5b506102f5610731366004614043565b611353565b348015610741575f80fd5b506102f561139c565b348015610755575f80fd5b506102f56113b0565b348015610769575f80fd5b5061028d610778366004613f32565b6113c4565b348015610788575f80fd5b506102f5610797366004613f4d565b6113fe565b3480156107a7575f80fd5b5061028d6107b63660046141a3565b611414565b60405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a26107fa3334611526565b565b5f8061080661155e565b6006015460ff1692915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0380546060915f80516020614e998339815191529161085190614260565b80601f016020809104026020016040519081016040528092919081815260200182805461087d90614260565b80156108c85780601f1061089f576101008083540402835291602001916108c8565b820191905f5260205f20905b8154815290600101906020018083116108ab57829003601f168201915b505050505091505090565b5f806108dd61155e565b6003015492915050565b5f336108f4818585611582565b60019150505b92915050565b5f805f80516020614e998339815191525b6002015492915050565b5f5f80516020614ef98339815191528161094462010203600160981b0131600160981b316142a6565b90505f61094f6113b0565b836001015461095e91906142a6565b905061096a82826142b9565b935050505090565b5f3361097f858285611594565b61098a8585856115f1565b60019150505b9392505050565b5f80516020614eb98339815191526109ad61164e565b6001600160a01b0382166109dc5760405162461bcd60e51b81526004016109d3906142cc565b60405180910390fd5b6109e68183611656565b15610a495760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b60648201526084016109d3565b6001600160a01b0382165f81815260018381016020526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a25050565b60405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a2610ad83382611677565b610ae233826116ab565b50565b5f80610aef61155e565b5492915050565b5f80516020614eb9833981519152610b0c61164e565b6001600160a01b038216610b325760405162461bcd60e51b81526004016109d3906142cc565b610b3c8183611656565b610b9a5760405162461bcd60e51b815260206004820152602960248201527f54656c65706f7274657252656769737472794170703a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b60648201526084016109d3565b6001600160a01b0382165f818152600183016020526040808220805460ff19169055517f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c39190a25050565b5f80516020614ed98339815191528054600114610c145760405162461bcd60e51b81526004016109d39061431a565b600281557f69a5f7616543528c4fbe43f410b1034bd6da4ba06c25bedf04617268014cf502545f80516020614ef983398151915290600160981b31908111610cd25760405162461bcd60e51b8152602060048201526044602482018190527f4e6174697665546f6b656e52656d6f74653a206275726e206164647265737320908201527f62616c616e6365206e6f742067726561746572207468616e206c6173742072656064820152631c1bdc9d60e21b608482015260a4016109d3565b5f826002015482610ce391906142b9565b90505f6064845f015483610cf7919061435e565b610d019190614375565b90505f610d0e82846142b9565b6002860185905590508115610d3157610d27308361173e565b610d313083611526565b5f610d4b610d3d6108d3565b610d45610f47565b846117eb565b11610db55760405162461bcd60e51b815260206004820152603460248201527f4e6174697665546f6b656e52656d6f74653a207a65726f207363616c6564206160448201527336b7bab73a103a37903932b837b93a10313ab93760611b60648201526084016109d3565b604080518082018252600181528151808301835262010203600160981b018152602080820185905292515f9380840192610df1929091016143a8565b60405160208183030381529060405281525090505f610eac6040518060c00160405280610e1c61139c565b8152602001610e29611164565b6001600160a01b0316815260408051808201825230815260208181018a905283015281018c90526060015f5b604051908082528060200260200182016040528015610e7e578160200160208202803683370190505b50815260200184604051602001610e9591906143c8565b604051602081830303815290604052815250611800565b9050807f0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a584604051610ee091815260200190565b60405180910390a250506001909555505050505050565b610eff61164e565b610ae28161191b565b610f106107fc565b610f2c5760405162461bcd60e51b81526004016109d39061440a565b610ae28134611ab3565b610f3e611b2a565b6107fa5f611b85565b5f80610f5161155e565b6004015460ff1692915050565b610f666107fc565b610f825760405162461bcd60e51b81526004016109d39061440a565b610ae28134611bf5565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060915f80516020614e998339815191529161085190614260565b5f5f80516020614eb98339815191526109908184611656565b5f336108f48185856115f1565b5f610ff961155e565b6006810154909150610100900460ff16156110565760405162461bcd60e51b815260206004820152601f60248201527f546f6b656e52656d6f74653a20616c726561647920726567697374657265640060448201526064016109d3565b604080516060808201835260058401548252600284015460ff600160a01b820481166020808601918252600160a81b9093048216858701908152865180880188525f808252885188518188015293518516848a015291519093168286015286518083039095018552608090910190955280820192909252919290916110eb906110e190870187613f32565b8660200135611c61565b6040805160c0810182526001870154815260028701546001600160a01b03166020808301919091528251808401845293945061115c939192830191908190611135908b018b613f32565b6001600160a01b0316815260209081018690529082526201fbd0908201526040015f610e55565b505050505050565b5f8061116e61155e565b600201546001600160a01b031692915050565b611189611ca8565b5f5f80516020614eb983398151915260028101548154919250906001600160a01b0316634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa1580156111f4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112189190614459565b101561127f5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b60648201526084016109d3565b6112898133611656565b156112ef5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b60648201526084016109d3565b61132f858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250611cf292505050565b50611338611f09565b50505050565b5f805f80516020614eb9833981519152610911565b6001600160a01b039182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b5f806113a661155e565b6001015492915050565b5f806113ba61155e565b6005015492915050565b6113cc611b2a565b6001600160a01b0381166113f557604051631e4fbdf760e01b81525f60048201526024016109d3565b610ae281611b85565b5f600561140c83601f6142a6565b901c92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156114585750825b90505f826001600160401b031660011480156114735750303b155b905081158015611481575080155b1561149f5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156114c957845460ff60401b1916600160401b1785555b6114d589898989611f33565b831561151b57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b6001600160a01b03821661154f5760405163ec442f0560e01b81525f60048201526024016109d3565b61155a5f8383611fc3565b5050565b7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef0090565b61158f83838360016120fc565b505050565b5f61159f8484611353565b90505f19811461133857818110156115e357604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016109d3565b61133884848484035f6120fc565b6001600160a01b03831661161a57604051634b637e8f60e11b81525f60048201526024016109d3565b6001600160a01b0382166116435760405163ec442f0560e01b81525f60048201526024016109d3565b61158f838383611fc3565b6107fa611b2a565b6001600160a01b03165f908152600191909101602052604090205460ff1690565b6001600160a01b0382166116a057604051634b637e8f60e11b81525f60048201526024016109d3565b61155a825f83611fc3565b804710156116ce5760405163cd78605960e01b81523060048201526024016109d3565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114611717576040519150601f19603f3d011682016040523d82523d5f602084013e61171c565b606091505b505090508061158f57604051630a12f52160e11b815260040160405180910390fd5b7f69a5f7616543528c4fbe43f410b1034bd6da4ba06c25bedf04617268014cf50180545f80516020614ef98339815191529183915f9061177f9084906142a6565b90915550506040516327ad555d60e11b81526001600160a01b0384166004820152602481018390526001600160991b0190634f5aaaba906044015f604051808303815f87803b1580156117d0575f80fd5b505af11580156117e2573d5f803e3d5ffd5b50505050505050565b5f6117f88484845f6121df565b949350505050565b5f8061180a61220f565b604084015160200151909150156118af576040830151516001600160a01b031661188c5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a207a65726f206665652060448201526c746f6b656e206164647265737360981b60648201526084016109d3565b6040830151602081015190516118af916001600160a01b039091169083906122ff565b604051630624488560e41b81526001600160a01b038216906362448850906118db908690600401614470565b6020604051808303815f875af11580156118f7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109909190614459565b5f80516020614eb983398151915280546040805163301fd1f560e21b815290515f926001600160a01b03169163c07f47d49160048083019260209291908290030181865afa15801561196f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119939190614459565b600283015490915081841115611a055760405162461bcd60e51b815260206004820152603160248201527f54656c65706f7274657252656769737472794170703a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b60648201526084016109d3565b808411611a7a5760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e0060648201526084016109d3565b60028301849055604051849082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d905f90a350505050565b5f80516020614ed98339815191528054600114611ae25760405162461bcd60e51b81526004016109d39061431a565b600281555f611aef61155e565b9050611afa84612386565b6001810154843503611b1657611b1084846125c0565b50611b22565b611b1084846127e4565b505b600190555050565b33611b5c7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146107fa5760405163118cdaa760e01b81523360048201526024016109d3565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f80516020614ed98339815191528054600114611c245760405162461bcd60e51b81526004016109d39061431a565b600281555f611c3161155e565b9050611c3c84612a85565b6001810154843503611c5757611c528484612b70565b611b20565b611b208484612cdf565b5f815f03611c7057505f6108fa565b306001600160a01b03841603611c9d57611c8b333084611594565b611c963330846115f1565b50806108fa565b610990833384612ea5565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00805460011901611cec57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f611cfb61155e565b905080600101548414611d625760405162461bcd60e51b815260206004820152602960248201527f546f6b656e52656d6f74653a20696e76616c696420736f7572636520626c6f636044820152681ad8da185a5b88125160ba1b60648201526084016109d3565b60028101546001600160a01b03848116911614611dd45760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e52656d6f74653a20696e76616c6964206f726967696e2073656e646044820152696572206164647265737360b01b60648201526084016109d3565b5f82806020019051810190611de99190614569565b6006830154909150610100900460ff161580611e0a5750600682015460ff16155b15611e215760068201805461ffff19166101011790555b600181516004811115611e3657611e36614394565b03611e6d575f8160200151806020019051810190611e5491906145f1565b9050611e67815f01518260200151613008565b50611f02565b600281516004811115611e8257611e82614394565b03611eb0575f8160200151806020019051810190611ea09190614629565b9050611e67818260800151613055565b60405162461bcd60e51b815260206004820152602160248201527f546f6b656e52656d6f74653a20696e76616c6964206d657373616765207479706044820152606560f81b60648201526084016109d3565b5050505050565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005b6001905550565b611f3b61317f565b815f03611fa45760405162461bcd60e51b815260206004820152603160248201527f4e6174697665546f6b656e52656d6f74653a207a65726f20696e697469616c206044820152707265736572766520696d62616c616e636560781b60648201526084016109d3565b611fae83846131c8565b611fba848360126131da565b61133881613206565b5f80516020614e998339815191526001600160a01b038416611ffd5781816002015f828254611ff291906142a6565b9091555061206d9050565b6001600160a01b0384165f908152602082905260409020548281101561204f5760405163391434e360e21b81526001600160a01b038616600482015260248101829052604481018490526064016109d3565b6001600160a01b0385165f9081526020839052604090209083900390555b6001600160a01b03831661208b5760028101805483900390556120a9565b6001600160a01b0383165f9081526020829052604090208054830190555b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516120ee91815260200190565b60405180910390a350505050565b5f80516020614e998339815191526001600160a01b0385166121335760405163e602df0560e01b81525f60048201526024016109d3565b6001600160a01b03841661215c57604051634a1406b160e11b81525f60048201526024016109d3565b6001600160a01b038086165f90815260018301602090815260408083209388168352929052208390558115611f0257836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925856040516121d091815260200190565b60405180910390a35050505050565b5f811515841515036121fc576121f5858461435e565b90506117f8565b6122068584614375565b95945050505050565b5f80516020614eb983398151915280546040805163d820e64f60e01b815290515f939284926001600160a01b039091169163d820e64f916004808201926020929091908290030181865afa158015612269573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061228d91906146f3565b90506122998282611656565b156108fa5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b60648201526084016109d3565b604051636eb1769f60e11b81523060048201526001600160a01b0383811660248301525f919085169063dd62ed3e90604401602060405180830381865afa15801561234c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123709190614459565b9050611338848461238185856142a6565b61327d565b80356123a45760405162461bcd60e51b81526004016109d39061470e565b5f6123b56040830160208401613f32565b6001600160a01b0316036123db5760405162461bcd60e51b81526004016109d390614759565b5f6123ec6060830160408401613f32565b6001600160a01b0316036124575760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e7420636f6e7460448201526b72616374206164647265737360a01b60648201526084016109d3565b5f81608001351161247a5760405162461bcd60e51b81526004016109d3906147b6565b5f8160a00135116124db5760405162461bcd60e51b815260206004820152602560248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e7420676173206044820152641b1a5b5a5d60da1b60648201526084016109d3565b80608001358160a00135106125435760405162461bcd60e51b815260206004820152602860248201527f546f6b656e52656d6f74653a20696e76616c696420726563697069656e742067604482015267185cc81b1a5b5a5d60c21b60648201526084016109d3565b5f612555610100830160e08401613f32565b6001600160a01b031603610ae25760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a207a65726f2066616c6c6261636b20726563697060448201526b69656e74206164647265737360a01b60648201526084016109d3565b5f6125c961155e565b90506125f96125de6040850160208601613f32565b6101408501356125f460e0870160c08801613f32565b61333a565b5f6126218361261061012087016101008801613f32565b866101200135876101400135613437565b6040805180820190915291945091505f908060028152602001604051806101000160405280865f01548152602001306001600160a01b031681526020016126653390565b6001600160a01b0316815260200161268360608a0160408b01613f32565b6001600160a01b03168152602081018890526040016126a560608a018a6147fa565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525050509082525060a089013560208201526040016126f96101008a0160e08b01613f32565b6001600160a01b03169052604051612714919060200161483c565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b03166020820152815180830183529293505f92612796928201908061276e6101208c016101008d01613f32565b6001600160a01b03168152602090810188905290825260808a0135908201526040015f610e55565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b1688886040516127d4929190614943565b60405180910390a3505050505050565b5f6127ed61155e565b905061281983356128046040860160208701613f32565b61281460e0870160c08801613f32565b6134fa565b5f6128308361261061012087016101008801613f32565b6040805180820190915291945091505f90806004815260200160405180610160016040528061285c3390565b6001600160a01b03168152602001885f013581526020018860200160208101906128869190613f32565b6001600160a01b031681526020016128a460608a0160408b01613f32565b6001600160a01b03168152602081018890526040016128c660608a018a6147fa565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525050509082525060a0890135602082015260400161291a6101008a0160e08b01613f32565b6001600160a01b031681526080890135602082015260400161294260e08a0160c08b01613f32565b6001600160a01b03168152610140890135602091820152604051612967929101614a51565b60408051601f19818403018152919052905290505f6105dc61299661298f60608901896147fa565b90506113fe565b6129a0919061435e565b6129ad90620557306142a6565b6040805160c0810182526001870154815260028701546001600160a01b03166020820152815180830183529293505f92612a3692820190806129f76101208d016101008e01613f32565b6001600160a01b031681526020908101899052908252818101869052604080515f815280830182528184015251606090920191610e95918891016143c8565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b168989604051612a74929190614943565b60405180910390a350505050505050565b5f612a966060830160408401613f32565b6001600160a01b031603612af85760405162461bcd60e51b815260206004820152602360248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e74206164647260448201526265737360e81b60648201526084016109d3565b5f8160c0013511612b1b5760405162461bcd60e51b81526004016109d3906147b6565b8035612b395760405162461bcd60e51b81526004016109d39061470e565b5f612b4a6040830160208401613f32565b6001600160a01b031603610ae25760405162461bcd60e51b81526004016109d390614759565b5f612b7961155e565b9050612ba4612b8e6040850160208601613f32565b60a08501356125f4610100870160e08801613f32565b5f612bc883612bb96080870160608801613f32565b86608001358760a00135613437565b6040805180820190915291945091505f9080600181526020016040518060400160405280886040016020810190612bff9190613f32565b6001600160a01b0316815260200187815250604051602001612c2191906143a8565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b03166020820152815180830183529293505f92612ca19282019080612c7960808c0160608d01613f32565b6001600160a01b03168152602090810188905290825260c08a0135908201526040015f610e55565b9050336001600160a01b0316817f93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb5288886040516127d4929190614b2e565b5f612ce861155e565b9050612d108335612cff6040860160208701613f32565b612814610100870160e08801613f32565b5f612d2583612bb96080870160608801613f32565b60408051808201825260038152815160e081018352883581529396509193505f9260208084019282820191612d5e918b01908b01613f32565b6001600160a01b03168152602001612d7c60608a0160408b01613f32565b6001600160a01b031681526020810188905260a0890135604082015260c08901356060820152608001612db66101008a0160e08b01613f32565b6001600160a01b03169052604051612e269190602001815181526020808301516001600160a01b0390811691830191909152604080840151821690830152606080840151908301526080808401519083015260a0808401519083015260c092830151169181019190915260e00190565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b03166020820152815180830183529293505f92612ca19282019080612e7e60808c0160608d01613f32565b6001600160a01b03168152602090810188905290825262053020908201526040015f610e55565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa158015612eeb573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f0f9190614459565b9050612f266001600160a01b038616853086613598565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa158015612f6a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f8e9190614459565b9050818111612ff45760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b60648201526084016109d3565b612ffe82826142b9565b9695505050505050565b816001600160a01b03167f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b8260405161304391815260200190565b60405180910390a261155a828261173e565b61305f308261173e565b5f825f0151836020015184604001518560a001516040516024016130869493929190614bcd565b60408051601f198184030181529190526020810180516001600160e01b031663161b12ff60e11b17905260c084015160608501519192505f916130cc91908590856135d1565b905080156131205783606001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff48460405161311391815260200190565b60405180910390a2611338565b83606001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb08460405161315f91815260200190565b60405180910390a260e0840151611338906001600160a01b0316846116ab565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166107fa57604051631afcd79f60e31b815260040160405180910390fd5b6131d061317f565b61155a82826136a1565b6131e261317f565b6131f3835f015184602001516136f1565b6131fb61370b565b61158f83838361371b565b61320e61317f565b6064811061326c5760405162461bcd60e51b815260206004820152602560248201527f4e6174697665546f6b656e52656d6f74653a20696e76616c69642070657263656044820152646e7461676560d81b60648201526084016109d3565b5f80516020614ef983398151915255565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b1790526132ce8482613a3f565b611338576040516001600160a01b0384811660248301525f604483015261333091869182169063095ea7b3906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613adc565b6113388482613adc565b5f61334361155e565b60028101549091506001600160a01b038581169116146133755760405162461bcd60e51b81526004016109d390614bfe565b82156133cf5760405162461bcd60e51b815260206004820152602360248201527f546f6b656e52656d6f74653a206e6f6e2d7a65726f207365636f6e646172792060448201526266656560e81b60648201526084016109d3565b6001600160a01b038216156113385760405162461bcd60e51b815260206004820152602860248201527f546f6b656e52656d6f74653a206e6f6e2d7a65726f206d756c74692d686f702060448201526766616c6c6261636b60c01b60648201526084016109d3565b5f805f61344261155e565b905061344d87613b3d565b96506134598686611c61565b600382015460048301549196506134739160ff16866117eb565b6003820154600483015461348b919060ff168a6117eb565b116134ed5760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a20696e73756666696369656e7420746f6b656e7360448201526b103a37903a3930b739b332b960a11b60648201526084016109d3565b5094959294509192505050565b5f61350361155e565b8054909150840361353657306001600160a01b038416036135365760405162461bcd60e51b81526004016109d390614bfe565b6001600160a01b0382166113385760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a207a65726f206d756c74692d686f702066616c6c6044820152636261636b60e01b60648201526084016109d3565b6040516001600160a01b0384811660248301528381166044830152606482018390526113389186918216906323b872dd906084016132fe565b5f845a10156136225760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e7420676173000000000060448201526064016109d3565b834710156136725760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c756500000060448201526064016109d3565b826001600160a01b03163b5f0361368a57505f6117f8565b5f805f84516020860188888bf19695505050505050565b6136a961317f565b5f80516020614e998339815191527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace036136e28482614c9f565b50600481016113388382614c9f565b6136f961317f565b61370282613b55565b61155a81613b76565b61371361317f565b6107fa613b87565b61372361317f565b5f61372c61155e565b90506005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015613771573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906137959190614459565b815560408401516137fb5760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e52656d6f74653a207a65726f20746f6b656e20686f6d6520626c6f60448201526918dad8da185a5b88125160b21b60648201526084016109d3565b80546040850151036138755760405162461bcd60e51b815260206004820152603b60248201527f546f6b656e52656d6f74653a2063616e6e6f74206465706c6f7920746f20736160448201527f6d6520626c6f636b636861696e20617320746f6b656e20686f6d65000000000060648201526084016109d3565b60608401516001600160a01b03166138db5760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a207a65726f20746f6b656e20686f6d65206164646044820152637265737360e01b60648201526084016109d3565b6012846080015160ff1611156139455760405162461bcd60e51b815260206004820152602960248201527f546f6b656e52656d6f74653a20746f6b656e20686f6d6520646563696d616c73604482015268040e8dede40d0d2ced60bb1b60648201526084016109d3565b60128260ff1611156139a55760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a20746f6b656e20646563696d616c7320746f6f206044820152630d0d2ced60e31b60648201526084016109d3565b60408401516001820155606084015160028201805460058401869055600684018054871560ff1990911617905560808701516001600160a01b039093166001600160a81b031990911617600160a01b60ff808516919091029190911760ff60a81b1916600160a81b91861691909102179055613a219083613b9b565b60048301805460ff1916911515919091179055600390910155505050565b5f805f846001600160a01b031684604051613a5a9190614d5a565b5f604051808303815f865af19150503d805f8114613a93576040519150601f19603f3d011682016040523d82523d5f602084013e613a98565b606091505b5091509150818015613ac2575080511580613ac2575080806020019051810190613ac29190614d75565b80156122065750505050506001600160a01b03163b151590565b5f613af06001600160a01b03841683613be5565b905080515f14158015613b14575080806020019051810190613b129190614d75565b155b1561158f57604051635274afe760e01b81526001600160a01b03841660048201526024016109d3565b5f613b5162010203600160981b01836116ab565b5090565b613b5d61317f565b613b65613bf2565b613b6d613c02565b610ae281613c0a565b613b7e61317f565b610ae281613d23565b5f5f80516020614ed9833981519152611f2c565b5f8060ff808516908416118181613bbe57613bb68587614d94565b60ff16613bcc565b613bc88686614d94565b60ff165b613bd790600a614e8d565b9350909150505b9250929050565b606061099083835f613d2b565b613bfa61317f565b6107fa613dba565b6107fa61317f565b613c1261317f565b6001600160a01b038116613c8e5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f7274657252656769737472794170703a207a65726f2074656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084016109d3565b5f80516020614eb983398151915280546001600160a01b0319166001600160a01b03831690811782556040805163301fd1f560e21b8152905184929163c07f47d49160048083019260209291908290030181865afa158015613cf2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613d169190614459565b8260020181905550505050565b6113cc61317f565b606081471015613d505760405163cd78605960e01b81523060048201526024016109d3565b5f80856001600160a01b03168486604051613d6b9190614d5a565b5f6040518083038185875af1925050503d805f8114613da5576040519150601f19603f3d011682016040523d82523d5f602084013e613daa565b606091505b5091509150612ffe868383613dc2565b611f0961317f565b606082613dd757613dd282613e1e565b610990565b8151158015613dee57506001600160a01b0384163b155b15613e1757604051639996b31560e01b81526001600160a01b03851660048201526024016109d3565b5080610990565b805115613e2e5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f5b83811015613e61578181015183820152602001613e49565b50505f910152565b5f8151808452613e80816020860160208601613e47565b601f01601f19169290920160200192915050565b602081525f6109906020830184613e69565b6001600160a01b0381168114610ae2575f80fd5b8035613ec581613ea6565b919050565b5f8060408385031215613edb575f80fd5b8235613ee681613ea6565b946020939093013593505050565b5f805f60608486031215613f06575f80fd5b8335613f1181613ea6565b92506020840135613f2181613ea6565b929592945050506040919091013590565b5f60208284031215613f42575f80fd5b813561099081613ea6565b5f60208284031215613f5d575f80fd5b5035919050565b5f60208284031215613f74575f80fd5b81356001600160401b03811115613f89575f80fd5b82016101608185031215610990575f80fd5b5f6101008284031215613fac575f80fd5b50919050565b5f60408284031215613fac575f80fd5b5f805f8060608587031215613fd5575f80fd5b843593506020850135613fe781613ea6565b925060408501356001600160401b0380821115614002575f80fd5b818701915087601f830112614015575f80fd5b813581811115614023575f80fd5b886020828501011115614034575f80fd5b95989497505060200194505050565b5f8060408385031215614054575f80fd5b823561405f81613ea6565b9150602083013561406f81613ea6565b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b60405160a081016001600160401b03811182821017156140b0576140b061407a565b60405290565b604080519081016001600160401b03811182821017156140b0576140b061407a565b60405161010081016001600160401b03811182821017156140b0576140b061407a565b604051601f8201601f191681016001600160401b03811182821017156141235761412361407a565b604052919050565b5f6001600160401b038211156141435761414361407a565b50601f01601f191660200190565b5f82601f830112614160575f80fd5b813561417361416e8261412b565b6140fb565b818152846020838601011115614187575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f808486036101008112156141b8575f80fd5b60a08112156141c5575f80fd5b506141ce61408e565b85356141d981613ea6565b815260208601356141e981613ea6565b602082015260408681013590820152606086013561420681613ea6565b6060820152608086013560ff8116811461421e575f80fd5b6080820152935060a08501356001600160401b0381111561423d575f80fd5b61424987828801614151565b949794965050505060c08301359260e00135919050565b600181811c9082168061427457607f821691505b602082108103613fac57634e487b7160e01b5f52602260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b808201808211156108fa576108fa614292565b818103818111156108fa576108fa614292565b6020808252602e908201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b80820281158282048414176108fa576108fa614292565b5f8261438f57634e487b7160e01b5f52601260045260245ffd5b500490565b634e487b7160e01b5f52602160045260245ffd5b81516001600160a01b0316815260208083015190820152604081016108fa565b602081525f8251600581106143eb57634e487b7160e01b5f52602160045260245ffd5b8060208401525060208301516040808401526117f86060840182613e69565b6020808252602f908201527f4e6174697665546f6b656e52656d6f74653a20636f6e747261637420756e646560408201526e1c98dbdb1b185d195c985b1a5e9959608a1b606082015260800190565b5f60208284031215614469575f80fd5b5051919050565b6020808252825182820152828101516001600160a01b039081166040808501919091528401518051821660608501528083015160808501525f929161010085019190606087015160a0870152608087015160e060c08801528051938490528401925f92506101208701905b808410156144fd578451831682529385019360019390930192908501906144db565b5060a0880151878203601f190160e0890152945061451b8186613e69565b98975050505050505050565b5f82601f830112614536575f80fd5b815161454461416e8261412b565b818152846020838601011115614558575f80fd5b6117f8826020830160208701613e47565b5f60208284031215614579575f80fd5b81516001600160401b038082111561458f575f80fd5b90830190604082860312156145a2575f80fd5b6145aa6140b6565b8251600581106145b8575f80fd5b81526020830151828111156145cb575f80fd5b6145d787828601614527565b60208301525095945050505050565b8051613ec581613ea6565b5f60408284031215614601575f80fd5b6146096140b6565b825161461481613ea6565b81526020928301519281019290925250919050565b5f60208284031215614639575f80fd5b81516001600160401b038082111561464f575f80fd5b908301906101008286031215614663575f80fd5b61466b6140d8565b8251815261467b602084016145e6565b602082015261468c604084016145e6565b604082015261469d606084016145e6565b60608201526080830151608082015260a0830151828111156146bd575f80fd5b6146c987828601614527565b60a08301525060c083015160c08201526146e560e084016145e6565b60e082015295945050505050565b5f60208284031215614703575f80fd5b815161099081613ea6565b6020808252602b908201527f546f6b656e52656d6f74653a207a65726f2064657374696e6174696f6e20626c60408201526a1bd8dad8da185a5b88125160aa1b606082015260800190565b60208082526037908201527f546f6b656e52656d6f74653a207a65726f2064657374696e6174696f6e20746f60408201527f6b656e207472616e736665727265722061646472657373000000000000000000606082015260800190565b60208082526024908201527f546f6b656e52656d6f74653a207a65726f20726571756972656420676173206c6040820152631a5b5a5d60e21b606082015260800190565b5f808335601e1984360301811261480f575f80fd5b8301803591506001600160401b03821115614828575f80fd5b602001915036819003821315613bde575f80fd5b60208152815160208201525f602083015160018060a01b0380821660408501528060408601511660608501525050606083015161488460808401826001600160a01b03169052565b50608083015160a083015260a08301516101008060c08501526148ab610120850183613e69565b915060c085015160e085015260e08501516148d0828601826001600160a01b03169052565b5090949350505050565b5f808335601e198436030181126148ef575f80fd5b83016020810192503590506001600160401b0381111561490d575f80fd5b803603821315613bde575f80fd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60408152823560408201525f61495b60208501613eba565b6001600160a01b0316606083015261497560408501613eba565b6001600160a01b0316608083015261499060608501856148da565b6101608060a08601526149a86101a08601838561491b565b9250608087013560c086015260a087013560e08601526149ca60c08801613eba565b91506101006149e3818701846001600160a01b03169052565b6149ef60e08901613eba565b9250610120614a08818801856001600160a01b03169052565b614a13828a01613eba565b93506101409150614a2e828801856001600160a01b03169052565b880135918601919091529095013561018084015260209092019290925292915050565b60208152614a6b6020820183516001600160a01b03169052565b602082015160408201525f6040830151614a9060608401826001600160a01b03169052565b5060608301516001600160a01b038116608084015250608083015160a083015260a08301516101608060c0850152614acc610180850183613e69565b915060c085015160e085015260e0850151610100614af4818701836001600160a01b03169052565b860151610120868101919091528601519050610140614b1d818701836001600160a01b03169052565b959095015193019290925250919050565b8235815261012081016020840135614b4581613ea6565b6001600160a01b039081166020840152604085013590614b6482613ea6565b166040830152614b7660608501613eba565b6001600160a01b0381166060840152506080840135608083015260a084013560a083015260c084013560c0830152614bb060e08501613eba565b6001600160a01b031660e083015261010090910191909152919050565b8481526001600160a01b038481166020830152831660408201526080606082018190525f90612ffe90830184613e69565b6020808252603a908201527f546f6b656e52656d6f74653a20696e76616c69642064657374696e6174696f6e60408201527f20746f6b656e207472616e736665727265722061646472657373000000000000606082015260800190565b601f82111561158f57805f5260205f20601f840160051c81016020851015614c805750805b601f840160051c820191505b81811015611f02575f8155600101614c8c565b81516001600160401b03811115614cb857614cb861407a565b614ccc81614cc68454614260565b84614c5b565b602080601f831160018114614cff575f8415614ce85750858301515b5f19600386901b1c1916600185901b17855561115c565b5f85815260208120601f198616915b82811015614d2d57888601518255948401946001909101908401614d0e565b5085821015614d4a57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f8251614d6b818460208701613e47565b9190910192915050565b5f60208284031215614d85575f80fd5b81518015158114610990575f80fd5b60ff82811682821603908111156108fa576108fa614292565b600181815b80851115614de757815f1904821115614dcd57614dcd614292565b80851615614dda57918102915b93841c9390800290614db2565b509250929050565b5f82614dfd575060016108fa565b81614e0957505f6108fa565b8160018114614e1f5760028114614e2957614e45565b60019150506108fa565b60ff841115614e3a57614e3a614292565b50506001821b6108fa565b5060208310610133831016604e8410600b8410161715614e68575081810a6108fa565b614e728383614dad565b805f1904821115614e8557614e85614292565b029392505050565b5f6109908383614def56fe52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace002a141db39aa10c95f68b87d187b82be26cf679a21fe482d47c0b39775a4a5200d2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c750069a5f7616543528c4fbe43f410b1034bd6da4ba06c25bedf04617268014cf500a164736f6c6343000819000a",
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

// Initialize is a paid mutator transaction binding the contract method 0xf56c363d.
//
// Solidity: function initialize((address,address,bytes32,address,uint8) settings, string nativeAssetSymbol, uint256 initialReserveImbalance, uint256 burnedFeesReportingRewardPercentage_) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactor) Initialize(opts *bind.TransactOpts, settings TokenRemoteSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage_ *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.contract.Transact(opts, "initialize", settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf56c363d.
//
// Solidity: function initialize((address,address,bytes32,address,uint8) settings, string nativeAssetSymbol, uint256 initialReserveImbalance, uint256 burnedFeesReportingRewardPercentage_) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableSession) Initialize(settings TokenRemoteSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage_ *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Initialize(&_NativeTokenRemoteUpgradeable.TransactOpts, settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf56c363d.
//
// Solidity: function initialize((address,address,bytes32,address,uint8) settings, string nativeAssetSymbol, uint256 initialReserveImbalance, uint256 burnedFeesReportingRewardPercentage_) returns()
func (_NativeTokenRemoteUpgradeable *NativeTokenRemoteUpgradeableTransactorSession) Initialize(settings TokenRemoteSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage_ *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemoteUpgradeable.Contract.Initialize(&_NativeTokenRemoteUpgradeable.TransactOpts, settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage_)
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
