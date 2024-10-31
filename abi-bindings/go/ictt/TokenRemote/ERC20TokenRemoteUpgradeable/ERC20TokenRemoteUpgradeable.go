// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20tokenremoteupgradeable

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

// ERC20TokenRemoteUpgradeableMetaData contains all meta data concerning the ERC20TokenRemoteUpgradeable contract.
var ERC20TokenRemoteUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"init\",\"type\":\"uint8\",\"internalType\":\"enumICMInitializable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ERC20_TOKEN_REMOTE_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MULTI_HOP_CALL_GAS_PER_WORD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MULTI_HOP_CALL_REQUIRED_GAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MULTI_HOP_SEND_REQUIRED_GAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTER_REMOTE_REQUIRED_GAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TELEPORTER_REGISTRY_APP_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKEN_REMOTE_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateNumWords\",\"inputs\":[{\"name\":\"payloadSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlockchainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInitialReserveImbalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIsCollateralized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinTeleporterVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMultiplyOnRemote\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenHomeAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenHomeBlockchainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenMultiplier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"settings\",\"type\":\"tuple\",\"internalType\":\"structTokenRemoteSettings\",\"components\":[{\"name\":\"teleporterRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"teleporterManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minTeleporterVersion\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenHomeBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenHomeAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenHomeDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"tokenName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isTeleporterAddressPaused\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseTeleporterAddress\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"receiveTeleporterMessage\",\"inputs\":[{\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originSenderAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerWithHome\",\"inputs\":[{\"name\":\"feeInfo\",\"type\":\"tuple\",\"internalType\":\"structTeleporterFeeInfo\",\"components\":[{\"name\":\"feeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"send\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structSendTokensInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendAndCall\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structSendAndCallInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipientGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseTeleporterAddress\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMinTeleporterVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"recipientContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSucceeded\",\"inputs\":[{\"name\":\"recipientContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinTeleporterVersionUpdated\",\"inputs\":[{\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TeleporterAddressPaused\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TeleporterAddressUnpaused\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensAndCallSent\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"input\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipientGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensSent\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"input\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161473638038061473683398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b6145e98061014d5f395ff3fe608060405234801561000f575f80fd5b50600436106101f1575f3560e01c806370a082311161011b578063b8a46d02116100b4578063dd62ed3e11610079578063dd62ed3e1461042d578063e0fd9cb814610440578063ef793e2a14610448578063f2fde38b14610450578063f3f981d814610463575f80fd5b8063b8a46d02146103e4578063c3cd6927146103f7578063c868efaa146103ff578063c9fe4ddf14610412578063d2cc7a7014610425575f80fd5b806370a082311461034d578063715018a61461036057806371717c18146103685780637ee3779a146103725780638da5cb5b1461037a578063909a6ac01461038f57806395d89b41146103b657806397314297146103be578063a9059cbb146103d1575f80fd5b8063313ce5671161018d578063313ce5671461029357806335cac159146102b45780634213cf78146102db5780634511243e146102e35780635507f3d1146102f65780635d16225d146103005780635eb995141461031357806362431a6514610326578063656900381461033a575f80fd5b806302a30c7d146101f557806306fdde03146102125780630733c8c814610227578063095ea7b31461023d57806315beb59f1461025057806318160ddd1461025957806323b872dd14610261578063254ac160146102745780632b0d8f181461027e575b5f80fd5b6101fd610476565b60405190151581526020015b60405180910390f35b61021a61048d565b604051610209919061358e565b61022f61052b565b604051908152602001610209565b6101fd61024b3660046135c4565b61053f565b61022f6105dc81565b61022f610558565b6101fd61026f3660046135ee565b61056c565b61022f6201fbd081565b61029161028c36600461362c565b610591565b005b5f8051602061459d8339815191525460405160ff9091168152602001610209565b61022f7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef0081565b61022f610690565b6102916102f136600461362c565b6106a1565b61022f6205302081565b61029161030e366004613647565b61078d565b610291610321366004613676565b61079b565b61022f5f8051602061459d83398151915281565b61029161034836600461368d565b6107af565b61022f61035b36600461362c565b6107b9565b6102916107e2565b61022f6205573081565b6101fd6107f5565b61038261080c565b60405161020991906136d2565b61022f7fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d0081565b61021a61083a565b6101fd6103cc36600461362c565b610856565b6101fd6103df3660046135c4565b61086c565b6102916103f23660046136e6565b610879565b610382610a40565b61029161040d3660046136fc565b610a5d565b6102916104203660046138b6565b610bf6565b61022f610d08565b61022f61043b36600461399f565b610d12565b61022f610d4c565b61022f610d60565b61029161045e36600461362c565b610d74565b61022f610471366004613676565b610dae565b5f80610480610dc4565b6006015460ff1692915050565b60605f610498610de8565b90508060030180546104a9906139d6565b80601f01602080910402602001604051908101604052809291908181526020018280546104d5906139d6565b80156105205780601f106104f757610100808354040283529160200191610520565b820191905f5260205f20905b81548152906001019060200180831161050357829003601f168201915b505050505091505090565b5f80610535610dc4565b6003015492915050565b5f3361054c818585610e0c565b60019150505b92915050565b5f80610562610de8565b6002015492915050565b5f33610579858285610e1e565b610584858585610e68565b60019150505b9392505050565b5f61059a610ec5565b90506105a4610ee9565b6001600160a01b0382166105d35760405162461bcd60e51b81526004016105ca90613a08565b60405180910390fd5b6105dd8183610ef1565b156106405760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b60648201526084016105ca565b6001600160a01b0382165f81815260018381016020526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a25050565b5f8061069a610dc4565b5492915050565b5f6106aa610ec5565b90506106b4610ee9565b6001600160a01b0382166106da5760405162461bcd60e51b81526004016105ca90613a08565b6106e48183610ef1565b6107425760405162461bcd60e51b815260206004820152602960248201527f54656c65706f7274657252656769737472794170703a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b60648201526084016105ca565b6001600160a01b0382165f818152600183016020526040808220805460ff19169055517f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c39190a25050565b6107978282610f12565b5050565b6107a3610ee9565b6107ac81610f85565b50565b6107978282611113565b5f806107c3610de8565b6001600160a01b039093165f9081526020939093525050604090205490565b6107ea611186565b6107f35f6111b8565b565b5f806107ff610dc4565b6004015460ff1692915050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b60605f610845610de8565b90508060040180546104a9906139d6565b5f80610860610ec5565b905061058a8184610ef1565b5f3361054c818585610e68565b5f610882610dc4565b6006810154909150610100900460ff16156108df5760405162461bcd60e51b815260206004820152601f60248201527f546f6b656e52656d6f74653a20616c726561647920726567697374657265640060448201526064016105ca565b604080516060808201835260058401548252600284015460ff600160a01b820481166020808601918252600160a81b9093048216858701908152865180880188525f808252885188518188015293518516848a015291519093168286015286518083039095018552608090910190955280820192909252919290916109749061096a9087018761362c565b8660200135611228565b6040805160c0810182526001870154815260028701546001600160a01b031660208083019190915282518084018452939450610a389391928301919081906109be908b018b61362c565b6001600160a01b0316815260209081018690529082526201fbd0908201526040015f5b604051908082528060200260200182016040528015610a0a578160200160208202803683370190505b50815260200184604051602001610a219190613a6a565b604051602081830303815290604052815250611270565b505050505050565b5f80610a4a610dc4565b600201546001600160a01b031692915050565b610a6561138b565b5f610a6e610ec5565b60028101548154919250906001600160a01b0316634c1f08ce336040518263ffffffff1660e01b8152600401610aa491906136d2565b602060405180830381865afa158015610abf573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ae39190613aac565b1015610b375760405162461bcd60e51b815260206004820152603060248201525f805160206145bd83398151915260448201526f32b632b837b93a32b91039b2b73232b960811b60648201526084016105ca565b610b418133610ef1565b15610ba75760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b60648201526084016105ca565b610be7858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506113d592505050565b50610bf06115ec565b50505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015610c3a5750825b90505f826001600160401b03166001148015610c555750303b155b905081158015610c63575080155b15610c815760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610cab57845460ff60401b1916600160401b1785555b610cb789898989611616565b8315610cfd57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b5f80610562610ec5565b5f80610d1c610de8565b6001600160a01b039485165f90815260019190910160209081526040808320959096168252939093525050205490565b5f80610d56610dc4565b6001015492915050565b5f80610d6a610dc4565b6005015492915050565b610d7c611186565b6001600160a01b038116610da5575f604051631e4fbdf760e01b81526004016105ca91906136d2565b6107ac816111b8565b5f6005610dbc83601f613ad7565b901c92915050565b7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef0090565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0090565b610e198383836001611653565b505050565b5f610e298484610d12565b90505f198114610bf05781811015610e5a57828183604051637dc7a0d960e11b81526004016105ca93929190613aea565b610bf084848484035f611653565b6001600160a01b038316610e91575f604051634b637e8f60e11b81526004016105ca91906136d2565b6001600160a01b038216610eba575f60405163ec442f0560e01b81526004016105ca91906136d2565b610e19838383611733565b7fde77a4dc7391f6f8f2d9567915d687d3aee79e7a1fc7300392f2727e9a0f1d0090565b6107f3611186565b6001600160a01b03165f908152600191909101602052604090205460ff1690565b5f610f1b611856565b8054909150600114610f3f5760405162461bcd60e51b81526004016105ca90613b0b565b600281555f610f4c610dc4565b9050610f578461187a565b6001810154843503610f7257610f6d8484611965565b610f7c565b610f7c8484611ae9565b50600190555050565b5f610f8e610ec5565b90505f815f015f9054906101000a90046001600160a01b03166001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fe2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110069190613aac565b6002830154909150818411156110655760405162461bcd60e51b815260206004820152603160248201525f805160206145bd83398151915260448201527032b632b837b93a32b9103b32b939b4b7b760791b60648201526084016105ca565b8084116110da5760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f7274657252656769737472794170703a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e0060648201526084016105ca565b60028301849055604051849082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d905f90a350505050565b5f61111c611856565b80549091506001146111405760405162461bcd60e51b81526004016105ca90613b0b565b600281555f61114d610dc4565b905061115884611cb4565b60018101548435036111745761116e8484611eee565b5061117e565b61116e84846120fd565b600190555050565b3361118f61080c565b6001600160a01b0316146107f3573360405163118cdaa760e01b81526004016105ca91906136d2565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f815f0361123757505f610552565b306001600160a01b0384160361126557611253335b3084610e1e565b61125e333084610e68565b5080610552565b61058a833384612399565b5f8061127a612504565b6040840151602001519091501561131f576040830151516001600160a01b03166112fc5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f7274657252656769737472794170703a207a65726f206665652060448201526c746f6b656e206164647265737360981b60648201526084016105ca565b60408301516020810151905161131f916001600160a01b039091169083906125f8565b604051630624488560e41b81526001600160a01b0382169063624488509061134b908690600401613b4f565b6020604051808303815f875af1158015611367573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061058a9190613aac565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f008054600119016113cf57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f6113de610dc4565b9050806001015484146114455760405162461bcd60e51b815260206004820152602960248201527f546f6b656e52656d6f74653a20696e76616c696420736f7572636520626c6f636044820152681ad8da185a5b88125160ba1b60648201526084016105ca565b60028101546001600160a01b038481169116146114b75760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e52656d6f74653a20696e76616c6964206f726967696e2073656e646044820152696572206164647265737360b01b60648201526084016105ca565b5f828060200190518101906114cc9190613c48565b6006830154909150610100900460ff1615806114ed5750600682015460ff16155b156115045760068201805461ffff19166101011790555b60018151600481111561151957611519613a56565b03611550575f81602001518060200190518101906115379190613cd0565b905061154a815f0151826020015161267f565b506115e5565b60028151600481111561156557611565613a56565b03611593575f81602001518060200190518101906115839190613d08565b905061154a8182608001516126cc565b60405162461bcd60e51b815260206004820152602160248201527f546f6b656e52656d6f74653a20696e76616c6964206d657373616765207479706044820152606560f81b60648201526084016105ca565b5050505050565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005b6001905550565b61161e612825565b611628838361286e565b611633845f83612880565b5f8051602061459d833981519152805460ff191660ff8316179055610bf0565b5f61165c610de8565b90506001600160a01b038516611687575f60405163e602df0560e01b81526004016105ca91906136d2565b6001600160a01b0384166116b0575f604051634a1406b160e11b81526004016105ca91906136d2565b6001600160a01b038086165f908152600183016020908152604080832093881683529290522083905581156115e557836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258560405161172491815260200190565b60405180910390a35050505050565b5f61173c610de8565b90506001600160a01b03841661176a5781816002015f82825461175f9190613ad7565b909155506117c79050565b6001600160a01b0384165f90815260208290526040902054828110156117a95784818460405163391434e360e21b81526004016105ca93929190613aea565b6001600160a01b0385165f9081526020839052604090209083900390555b6001600160a01b0383166117e5576002810180548390039055611803565b6001600160a01b0383165f9081526020829052604090208054830190555b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161184891815260200190565b60405180910390a350505050565b7fd2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c750090565b5f61188b606083016040840161362c565b6001600160a01b0316036118ed5760405162461bcd60e51b815260206004820152602360248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e74206164647260448201526265737360e81b60648201526084016105ca565b5f8160c00135116119105760405162461bcd60e51b81526004016105ca90613dd2565b803561192e5760405162461bcd60e51b81526004016105ca90613e16565b5f61193f604083016020840161362c565b6001600160a01b0316036107ac5760405162461bcd60e51b81526004016105ca90613e61565b5f61196e610dc4565b905061199e611983604085016020860161362c565b60a0850135611999610100870160e0880161362c565b6128b1565b5f6119c2836119b3608087016060880161362c565b86608001358760a001356129ae565b6040805180820190915291945091505f90806001815260200160405180604001604052808860400160208101906119f9919061362c565b6001600160a01b0316815260200187815250604051602001611a1b9190613eb8565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b03166020820152815180830183529293505f92611a9b9282019080611a7360808c0160608d0161362c565b6001600160a01b03168152602090810188905290825260c08a0135908201526040015f6109e1565b9050336001600160a01b0316817f93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb528888604051611ad9929190613ed8565b60405180910390a3505050505050565b5f611af2610dc4565b9050611b1f8335611b09604086016020870161362c565b611b1a610100870160e0880161362c565b612a71565b5f611b34836119b3608087016060880161362c565b60408051808201825260038152815160e081018352883581529396509193505f9260208084019282820191611b6d918b01908b0161362c565b6001600160a01b03168152602001611b8b60608a0160408b0161362c565b6001600160a01b031681526020810188905260a0890135604082015260c08901356060820152608001611bc56101008a0160e08b0161362c565b6001600160a01b03169052604051611c359190602001815181526020808301516001600160a01b0390811691830191909152604080840151821690830152606080840151908301526080808401519083015260a0808401519083015260c092830151169181019190915260e00190565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b03166020820152815180830183529293505f92611a9b9282019080611c8d60808c0160608d0161362c565b6001600160a01b03168152602090810188905290825262053020908201526040015f6109e1565b8035611cd25760405162461bcd60e51b81526004016105ca90613e16565b5f611ce3604083016020840161362c565b6001600160a01b031603611d095760405162461bcd60e51b81526004016105ca90613e61565b5f611d1a606083016040840161362c565b6001600160a01b031603611d855760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e7420636f6e7460448201526b72616374206164647265737360a01b60648201526084016105ca565b5f816080013511611da85760405162461bcd60e51b81526004016105ca90613dd2565b5f8160a0013511611e095760405162461bcd60e51b815260206004820152602560248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e7420676173206044820152641b1a5b5a5d60da1b60648201526084016105ca565b80608001358160a0013510611e715760405162461bcd60e51b815260206004820152602860248201527f546f6b656e52656d6f74653a20696e76616c696420726563697069656e742067604482015267185cc81b1a5b5a5d60c21b60648201526084016105ca565b5f611e83610100830160e0840161362c565b6001600160a01b0316036107ac5760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a207a65726f2066616c6c6261636b20726563697060448201526b69656e74206164647265737360a01b60648201526084016105ca565b5f611ef7610dc4565b9050611f22611f0c604085016020860161362c565b61014085013561199960e0870160c0880161362c565b5f611f4a83611f396101208701610100880161362c565b8661012001358761014001356129ae565b6040805180820190915291945091505f908060028152602001604051806101000160405280865f01548152602001306001600160a01b03168152602001611f8e3390565b6001600160a01b03168152602001611fac60608a0160408b0161362c565b6001600160a01b0316815260208101889052604001611fce60608a018a613f74565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525050509082525060a089013560208201526040016120226101008a0160e08b0161362c565b6001600160a01b0316905260405161203d9190602001613fb6565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b03166020820152815180830183529293505f926120bf92820190806120976101208c016101008d0161362c565b6001600160a01b03168152602090810188905290825260808a0135908201526040015f6109e1565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b168888604051611ad99291906140ad565b5f612106610dc4565b905061212d833561211d604086016020870161362c565b611b1a60e0870160c0880161362c565b5f61214483611f396101208701610100880161362c565b6040805180820190915291945091505f9080600481526020016040518061016001604052806121703390565b6001600160a01b03168152602001885f0135815260200188602001602081019061219a919061362c565b6001600160a01b031681526020016121b860608a0160408b0161362c565b6001600160a01b03168152602081018890526040016121da60608a018a613f74565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525050509082525060a0890135602082015260400161222e6101008a0160e08b0161362c565b6001600160a01b031681526080890135602082015260400161225660e08a0160c08b0161362c565b6001600160a01b0316815261014089013560209182015260405161227b9291016141a3565b60408051601f19818403018152919052905290505f6105dc6122aa6122a36060890189613f74565b9050610dae565b6122b4919061425e565b6122c19062055730613ad7565b6040805160c0810182526001870154815260028701546001600160a01b03166020820152815180830183529293505f9261234a928201908061230b6101208d016101008e0161362c565b6001600160a01b031681526020908101899052908252818101869052604080515f815280830182528184015251606090920191610a2191889101613a6a565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b1689896040516123889291906140ad565b60405180910390a350505050505050565b5f80846001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016123c791906136d2565b602060405180830381865afa1580156123e2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906124069190613aac565b905061241d6001600160a01b038616853086612b0f565b6040516370a0823160e01b81525f906001600160a01b038716906370a082319061244b9030906004016136d2565b602060405180830381865afa158015612466573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061248a9190613aac565b90508181116124f05760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b60648201526084016105ca565b6124fa8282614275565b9695505050505050565b5f8061250e610ec5565b90505f815f015f9054906101000a90046001600160a01b03166001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612562573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906125869190614288565b90506125928282610ef1565b156105525760405162461bcd60e51b815260206004820152603060248201527f54656c65706f7274657252656769737472794170703a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b60648201526084016105ca565b604051636eb1769f60e11b81523060048201526001600160a01b0383811660248301525f919085169063dd62ed3e90604401602060405180830381865afa158015612645573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906126699190613aac565b9050610bf0848461267a8585613ad7565b612b76565b816001600160a01b03167f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b826040516126ba91815260200190565b60405180910390a26107978282612c05565b6126d63082612c05565b6126e530836060015183610e0c565b5f825f01518360200151846040015130858760a00151604051602401612710969594939291906142a3565b60408051601f198184030181529190526020810180516001600160e01b03166394395edd60e01b17905260c084015160608501519192505f91612754919084612c39565b90505f612765308660600151610d12565b90506127763086606001515f610e0c565b81156127c85784606001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4856040516127bb91815260200190565b60405180910390a2612810565b84606001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb08560405161280791815260200190565b60405180910390a25b80156115e5576115e5308660e0015183610e68565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166107f357604051631afcd79f60e31b815260040160405180910390fd5b612876612825565b6107978282612c4e565b612888612825565b61289e835f015184602001518560400151612c7e565b6128a6612c99565b610e19838383612ca9565b5f6128ba610dc4565b60028101549091506001600160a01b038581169116146128ec5760405162461bcd60e51b81526004016105ca906142e3565b82156129465760405162461bcd60e51b815260206004820152602360248201527f546f6b656e52656d6f74653a206e6f6e2d7a65726f207365636f6e646172792060448201526266656560e81b60648201526084016105ca565b6001600160a01b03821615610bf05760405162461bcd60e51b815260206004820152602860248201527f546f6b656e52656d6f74653a206e6f6e2d7a65726f206d756c74692d686f702060448201526766616c6c6261636b60c01b60648201526084016105ca565b5f805f6129b9610dc4565b90506129c487612fcd565b96506129d08686611228565b600382015460048301549196506129ea9160ff1686612fe5565b60038201546004830154612a02919060ff168a612fe5565b11612a645760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a20696e73756666696369656e7420746f6b656e7360448201526b103a37903a3930b739b332b960a11b60648201526084016105ca565b5094959294509192505050565b5f612a7a610dc4565b80549091508403612aad57306001600160a01b03841603612aad5760405162461bcd60e51b81526004016105ca906142e3565b6001600160a01b038216610bf05760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a207a65726f206d756c74692d686f702066616c6c6044820152636261636b60e01b60648201526084016105ca565b6040516001600160a01b038481166024830152838116604483015260648201839052610bf09186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050612ff2565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052612bc7848261304a565b610bf0576040516001600160a01b0384811660248301525f6044830152612bfb91869182169063095ea7b390606401612b44565b610bf08482612ff2565b6001600160a01b038216612c2e575f60405163ec442f0560e01b81526004016105ca91906136d2565b6107975f8383611733565b5f612c46845f85856130eb565b949350505050565b612c56612825565b5f612c5f610de8565b905060038101612c6f8482614384565b5060048101610bf08382614384565b612c86612825565b612c9083826131bb565b610e19826131dd565b612ca1612825565b6107f36131ee565b612cb1612825565b5f612cba610dc4565b90506005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015612cff573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612d239190613aac565b81556060840151612d895760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e52656d6f74653a207a65726f20746f6b656e20686f6d6520626c6f60448201526918dad8da185a5b88125160b21b60648201526084016105ca565b8054606085015103612e035760405162461bcd60e51b815260206004820152603b60248201527f546f6b656e52656d6f74653a2063616e6e6f74206465706c6f7920746f20736160448201527f6d6520626c6f636b636861696e20617320746f6b656e20686f6d65000000000060648201526084016105ca565b60808401516001600160a01b0316612e695760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a207a65726f20746f6b656e20686f6d65206164646044820152637265737360e01b60648201526084016105ca565b60128460a0015160ff161115612ed35760405162461bcd60e51b815260206004820152602960248201527f546f6b656e52656d6f74653a20746f6b656e20686f6d6520646563696d616c73604482015268040e8dede40d0d2ced60bb1b60648201526084016105ca565b60128260ff161115612f335760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a20746f6b656e20646563696d616c7320746f6f206044820152630d0d2ced60e31b60648201526084016105ca565b60608401516001820155608084015160028201805460058401869055600684018054871560ff1990911617905560a08701516001600160a01b039093166001600160a81b031990911617600160a01b60ff808516919091029190911760ff60a81b1916600160a81b91861691909102179055612faf90836131f7565b60048301805460ff1916911515919091179055600390910155505050565b5f612fd73361124c565b612fe13383613241565b5090565b5f612c468484845f613275565b5f6130066001600160a01b0384168361329c565b905080515f1415801561302a575080806020019051810190613028919061443f565b155b15610e195782604051635274afe760e01b81526004016105ca91906136d2565b5f805f846001600160a01b031684604051613065919061445e565b5f604051808303815f865af19150503d805f811461309e576040519150601f19603f3d011682016040523d82523d5f602084013e6130a3565b606091505b50915091508180156130cd5750805115806130cd5750808060200190518101906130cd919061443f565b80156130e257505f856001600160a01b03163b115b95945050505050565b5f845a101561313c5760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e7420676173000000000060448201526064016105ca565b8347101561318c5760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c756500000060448201526064016105ca565b826001600160a01b03163b5f036131a457505f612c46565b5f805f84516020860188888bf19695505050505050565b6131c3612825565b6131cb6132a9565b6131d36132b9565b61079782826132c1565b6131e5612825565b6107ac81613426565b5f61160f611856565b5f8060ff80851690841611818161321a576132128587614479565b60ff16613228565b6132248686614479565b60ff165b61323390600a614572565b9350909150505b9250929050565b6001600160a01b03821661326a575f604051634b637e8f60e11b81526004016105ca91906136d2565b610797825f83611733565b5f811515841515036132925761328b858461425e565b9050612c46565b6130e2858461457d565b606061058a83835f61342e565b6132b1612825565b6107f36134bd565b6107f3612825565b6132c9612825565b6001600160a01b03821661333f5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c65604482015276706f72746572207265676973747279206164647265737360481b60648201526084016105ca565b5f613348610ec5565b90505f8390505f816001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa15801561338b573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906133af9190613aac565b116134045760405162461bcd60e51b815260206004820152603260248201525f805160206145bd833981519152604482015271656c65706f7274657220726567697374727960701b60648201526084016105ca565b81546001600160a01b0319166001600160a01b038216178255610bf083610f85565b610d7c612825565b606081471015613453573060405163cd78605960e01b81526004016105ca91906136d2565b5f80856001600160a01b0316848660405161346e919061445e565b5f6040518083038185875af1925050503d805f81146134a8576040519150601f19603f3d011682016040523d82523d5f602084013e6134ad565b606091505b50915091506124fa8683836134c5565b6115ec612825565b6060826134da576134d582613518565b61058a565b81511580156134f157506001600160a01b0384163b155b156135115783604051639996b31560e01b81526004016105ca91906136d2565b508061058a565b8051156135285780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f5b8381101561355b578181015183820152602001613543565b50505f910152565b5f815180845261357a816020860160208601613541565b601f01601f19169290920160200192915050565b602081525f61058a6020830184613563565b6001600160a01b03811681146107ac575f80fd5b80356135bf816135a0565b919050565b5f80604083850312156135d5575f80fd5b82356135e0816135a0565b946020939093013593505050565b5f805f60608486031215613600575f80fd5b833561360b816135a0565b9250602084013561361b816135a0565b929592945050506040919091013590565b5f6020828403121561363c575f80fd5b813561058a816135a0565b5f8082840361012081121561365a575f80fd5b61010080821215613669575f80fd5b9395938601359450505050565b5f60208284031215613686575f80fd5b5035919050565b5f806040838503121561369e575f80fd5b82356001600160401b038111156136b3575f80fd5b830161016081860312156135e0575f80fd5b6001600160a01b03169052565b6001600160a01b0391909116815260200190565b5f604082840312156136f6575f80fd5b50919050565b5f805f806060858703121561370f575f80fd5b843593506020850135613721816135a0565b925060408501356001600160401b038082111561373c575f80fd5b818701915087601f83011261374f575f80fd5b81358181111561375d575f80fd5b88602082850101111561376e575f80fd5b95989497505060200194505050565b634e487b7160e01b5f52604160045260245ffd5b60405160c081016001600160401b03811182821017156137b3576137b361377d565b60405290565b604080519081016001600160401b03811182821017156137b3576137b361377d565b60405161010081016001600160401b03811182821017156137b3576137b361377d565b604051601f8201601f191681016001600160401b03811182821017156138265761382661377d565b604052919050565b803560ff811681146135bf575f80fd5b5f6001600160401b038211156138565761385661377d565b50601f01601f191660200190565b5f82601f830112613873575f80fd5b81356138866138818261383e565b6137fe565b81815284602083860101111561389a575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f808486036101208112156138cb575f80fd5b60c08112156138d8575f80fd5b506138e1613791565b85356138ec816135a0565b815260208601356138fc816135a0565b8060208301525060408601356040820152606086013560608201526080860135613925816135a0565b608082015261393660a0870161382e565b60a0820152935060c08501356001600160401b0380821115613956575f80fd5b61396288838901613864565b945060e0870135915080821115613977575f80fd5b5061398487828801613864565b925050613994610100860161382e565b905092959194509250565b5f80604083850312156139b0575f80fd5b82356139bb816135a0565b915060208301356139cb816135a0565b809150509250929050565b600181811c908216806139ea57607f821691505b6020821081036136f657634e487b7160e01b5f52602260045260245ffd5b6020808252602e908201527f54656c65706f7274657252656769737472794170703a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b634e487b7160e01b5f52602160045260245ffd5b602081525f825160058110613a8d57634e487b7160e01b5f52602160045260245ffd5b806020840152506020830151604080840152612c466060840182613563565b5f60208284031215613abc575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561055257610552613ac3565b6001600160a01b039390931683526020830191909152604082015260600190565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b6020808252825182820152828101516001600160a01b039081166040808501919091528401518051821660608501528083015160808501525f929161010085019190606087015160a0870152608087015160e060c08801528051938490528401925f92506101208701905b80841015613bdc57845183168252938501936001939093019290850190613bba565b5060a0880151878203601f190160e08901529450613bfa8186613563565b98975050505050505050565b5f82601f830112613c15575f80fd5b8151613c236138818261383e565b818152846020838601011115613c37575f80fd5b612c46826020830160208701613541565b5f60208284031215613c58575f80fd5b81516001600160401b0380821115613c6e575f80fd5b9083019060408286031215613c81575f80fd5b613c896137b9565b825160058110613c97575f80fd5b8152602083015182811115613caa575f80fd5b613cb687828601613c06565b60208301525095945050505050565b80516135bf816135a0565b5f60408284031215613ce0575f80fd5b613ce86137b9565b8251613cf3816135a0565b81526020928301519281019290925250919050565b5f60208284031215613d18575f80fd5b81516001600160401b0380821115613d2e575f80fd5b908301906101008286031215613d42575f80fd5b613d4a6137db565b82518152613d5a60208401613cc5565b6020820152613d6b60408401613cc5565b6040820152613d7c60608401613cc5565b60608201526080830151608082015260a083015182811115613d9c575f80fd5b613da887828601613c06565b60a08301525060c083015160c0820152613dc460e08401613cc5565b60e082015295945050505050565b60208082526024908201527f546f6b656e52656d6f74653a207a65726f20726571756972656420676173206c6040820152631a5b5a5d60e21b606082015260800190565b6020808252602b908201527f546f6b656e52656d6f74653a207a65726f2064657374696e6174696f6e20626c60408201526a1bd8dad8da185a5b88125160aa1b606082015260800190565b60208082526037908201527f546f6b656e52656d6f74653a207a65726f2064657374696e6174696f6e20746f6040820152766b656e207472616e73666572726572206164647265737360481b606082015260800190565b81516001600160a01b031681526020808301519082015260408101610552565b8235815261012081016020840135613eef816135a0565b6001600160a01b039081166020840152604085013590613f0e826135a0565b166040830152613f20606085016135b4565b613f2d60608401826136c5565b506080840135608083015260a084013560a083015260c084013560c0830152613f5860e085016135b4565b613f6560e08401826136c5565b50826101008301529392505050565b5f808335601e19843603018112613f89575f80fd5b8301803591506001600160401b03821115613fa2575f80fd5b60200191503681900382131561323a575f80fd5b60208152815160208201525f602083015160018060a01b03808216604085015280604086015116606085015250506060830151613ff660808401826136c5565b50608083015160a083015260a08301516101008060c085015261401d610120850183613563565b915060c085015160e085015260e085015161403a828601826136c5565b5090949350505050565b5f808335601e19843603018112614059575f80fd5b83016020810192503590506001600160401b03811115614077575f80fd5b80360382131561323a575f80fd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60408152823560408201525f6140c5602085016135b4565b6140d260608401826136c5565b506140df604085016135b4565b6140ec60808401826136c5565b506140fa6060850185614044565b6101608060a08601526141126101a086018385614085565b9250608087013560c086015260a087013560e086015261413460c088016135b4565b9150610100614145818701846136c5565b61415160e089016135b4565b9250610120614162818801856136c5565b61416d828a016135b4565b93506101409150614180828801856136c5565b880135918601919091529095013561018084015260209092019290925292915050565b602081526141b56020820183516136c5565b602082015160408201525f60408301516141d260608401826136c5565b5060608301516141e560808401826136c5565b50608083015160a083015260a08301516101608060c085015261420c610180850183613563565b915060c085015160e085015260e085015161010061422c818701836136c5565b86015161012086810191909152860151905061014061424d818701836136c5565b959095015193019290925250919050565b808202811582820484141761055257610552613ac3565b8181038181111561055257610552613ac3565b5f60208284031215614298575f80fd5b815161058a816135a0565b8681526001600160a01b0386811660208301528581166040830152841660608201526080810183905260c060a082018190525f90613bfa90830184613563565b6020808252603a908201527f546f6b656e52656d6f74653a20696e76616c69642064657374696e6174696f6e60408201527f20746f6b656e207472616e736665727265722061646472657373000000000000606082015260800190565b601f821115610e1957805f5260205f20601f840160051c810160208510156143655750805b601f840160051c820191505b818110156115e5575f8155600101614371565b81516001600160401b0381111561439d5761439d61377d565b6143b1816143ab84546139d6565b84614340565b602080601f8311600181146143e4575f84156143cd5750858301515b5f19600386901b1c1916600185901b178555610a38565b5f85815260208120601f198616915b82811015614412578886015182559484019460019091019084016143f3565b508582101561442f57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f6020828403121561444f575f80fd5b8151801515811461058a575f80fd5b5f825161446f818460208701613541565b9190910192915050565b60ff828116828216039081111561055257610552613ac3565b600181815b808511156144cc57815f19048211156144b2576144b2613ac3565b808516156144bf57918102915b93841c9390800290614497565b509250929050565b5f826144e257506001610552565b816144ee57505f610552565b8160018114614504576002811461450e5761452a565b6001915050610552565b60ff84111561451f5761451f613ac3565b50506001821b610552565b5060208310610133831016604e8410600b841016171561454d575081810a610552565b6145578383614492565b805f190482111561456a5761456a613ac3565b029392505050565b5f61058a83836144d4565b5f8261459757634e487b7160e01b5f52601260045260245ffd5b50049056fe9b9029a3537fcf0e984763da4ac33bbf592a3462819171bf424e91cf6262230054656c65706f7274657252656769737472794170703a20696e76616c69642054a164736f6c6343000819000a",
}

// ERC20TokenRemoteUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TokenRemoteUpgradeableMetaData.ABI instead.
var ERC20TokenRemoteUpgradeableABI = ERC20TokenRemoteUpgradeableMetaData.ABI

// ERC20TokenRemoteUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20TokenRemoteUpgradeableMetaData.Bin instead.
var ERC20TokenRemoteUpgradeableBin = ERC20TokenRemoteUpgradeableMetaData.Bin

// DeployERC20TokenRemoteUpgradeable deploys a new Ethereum contract, binding an instance of ERC20TokenRemoteUpgradeable to it.
func DeployERC20TokenRemoteUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *ERC20TokenRemoteUpgradeable, error) {
	parsed, err := ERC20TokenRemoteUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20TokenRemoteUpgradeableBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20TokenRemoteUpgradeable{ERC20TokenRemoteUpgradeableCaller: ERC20TokenRemoteUpgradeableCaller{contract: contract}, ERC20TokenRemoteUpgradeableTransactor: ERC20TokenRemoteUpgradeableTransactor{contract: contract}, ERC20TokenRemoteUpgradeableFilterer: ERC20TokenRemoteUpgradeableFilterer{contract: contract}}, nil
}

// ERC20TokenRemoteUpgradeable is an auto generated Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeable struct {
	ERC20TokenRemoteUpgradeableCaller     // Read-only binding to the contract
	ERC20TokenRemoteUpgradeableTransactor // Write-only binding to the contract
	ERC20TokenRemoteUpgradeableFilterer   // Log filterer for contract events
}

// ERC20TokenRemoteUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenRemoteUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenRemoteUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenRemoteUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenRemoteUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenRemoteUpgradeableSession struct {
	Contract     *ERC20TokenRemoteUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ERC20TokenRemoteUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenRemoteUpgradeableCallerSession struct {
	Contract *ERC20TokenRemoteUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// ERC20TokenRemoteUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenRemoteUpgradeableTransactorSession struct {
	Contract     *ERC20TokenRemoteUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ERC20TokenRemoteUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeableRaw struct {
	Contract *ERC20TokenRemoteUpgradeable // Generic contract binding to access the raw methods on
}

// ERC20TokenRemoteUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeableCallerRaw struct {
	Contract *ERC20TokenRemoteUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenRemoteUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeableTransactorRaw struct {
	Contract *ERC20TokenRemoteUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20TokenRemoteUpgradeable creates a new instance of ERC20TokenRemoteUpgradeable, bound to a specific deployed contract.
func NewERC20TokenRemoteUpgradeable(address common.Address, backend bind.ContractBackend) (*ERC20TokenRemoteUpgradeable, error) {
	contract, err := bindERC20TokenRemoteUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeable{ERC20TokenRemoteUpgradeableCaller: ERC20TokenRemoteUpgradeableCaller{contract: contract}, ERC20TokenRemoteUpgradeableTransactor: ERC20TokenRemoteUpgradeableTransactor{contract: contract}, ERC20TokenRemoteUpgradeableFilterer: ERC20TokenRemoteUpgradeableFilterer{contract: contract}}, nil
}

// NewERC20TokenRemoteUpgradeableCaller creates a new read-only instance of ERC20TokenRemoteUpgradeable, bound to a specific deployed contract.
func NewERC20TokenRemoteUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenRemoteUpgradeableCaller, error) {
	contract, err := bindERC20TokenRemoteUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableCaller{contract: contract}, nil
}

// NewERC20TokenRemoteUpgradeableTransactor creates a new write-only instance of ERC20TokenRemoteUpgradeable, bound to a specific deployed contract.
func NewERC20TokenRemoteUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenRemoteUpgradeableTransactor, error) {
	contract, err := bindERC20TokenRemoteUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTransactor{contract: contract}, nil
}

// NewERC20TokenRemoteUpgradeableFilterer creates a new log filterer instance of ERC20TokenRemoteUpgradeable, bound to a specific deployed contract.
func NewERC20TokenRemoteUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenRemoteUpgradeableFilterer, error) {
	contract, err := bindERC20TokenRemoteUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableFilterer{contract: contract}, nil
}

// bindERC20TokenRemoteUpgradeable binds a generic wrapper to an already deployed contract.
func bindERC20TokenRemoteUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20TokenRemoteUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenRemoteUpgradeable.Contract.ERC20TokenRemoteUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ERC20TokenRemoteUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ERC20TokenRemoteUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenRemoteUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ERC20TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x62431a65.
//
// Solidity: function ERC20_TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) ERC20TOKENREMOTESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "ERC20_TOKEN_REMOTE_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ERC20TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x62431a65.
//
// Solidity: function ERC20_TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) ERC20TOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ERC20TOKENREMOTESTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// ERC20TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x62431a65.
//
// Solidity: function ERC20_TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) ERC20TOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ERC20TOKENREMOTESTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) MULTIHOPCALLGASPERWORD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "MULTI_HOP_CALL_GAS_PER_WORD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPCALLGASPERWORD(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPCALLGASPERWORD(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) MULTIHOPCALLREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "MULTI_HOP_CALL_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) MULTIHOPCALLREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPCALLREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) MULTIHOPCALLREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPCALLREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) MULTIHOPSENDREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "MULTI_HOP_SEND_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) MULTIHOPSENDREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPSENDREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) MULTIHOPSENDREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPSENDREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) REGISTERREMOTEREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "REGISTER_REMOTE_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) REGISTERREMOTEREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.REGISTERREMOTEREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) REGISTERREMOTEREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.REGISTERREMOTEREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) TELEPORTERREGISTRYAPPSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "TELEPORTER_REGISTRY_APP_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TELEPORTERREGISTRYAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0x909a6ac0.
//
// Solidity: function TELEPORTER_REGISTRY_APP_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) TELEPORTERREGISTRYAPPSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TELEPORTERREGISTRYAPPSTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x35cac159.
//
// Solidity: function TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) TOKENREMOTESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "TOKEN_REMOTE_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x35cac159.
//
// Solidity: function TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) TOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TOKENREMOTESTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x35cac159.
//
// Solidity: function TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) TOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TOKENREMOTESTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Allowance(&_ERC20TokenRemoteUpgradeable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Allowance(&_ERC20TokenRemoteUpgradeable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.BalanceOf(&_ERC20TokenRemoteUpgradeable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.BalanceOf(&_ERC20TokenRemoteUpgradeable.CallOpts, account)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) CalculateNumWords(opts *bind.CallOpts, payloadSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "calculateNumWords", payloadSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.CalculateNumWords(&_ERC20TokenRemoteUpgradeable.CallOpts, payloadSize)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.CalculateNumWords(&_ERC20TokenRemoteUpgradeable.CallOpts, payloadSize)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Decimals() (uint8, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Decimals(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) Decimals() (uint8, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Decimals(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetBlockchainID() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetBlockchainID(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetBlockchainID() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetBlockchainID(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetInitialReserveImbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getInitialReserveImbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetInitialReserveImbalance() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetInitialReserveImbalance(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetInitialReserveImbalance() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetInitialReserveImbalance(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetIsCollateralized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getIsCollateralized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetIsCollateralized() (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetIsCollateralized(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetIsCollateralized() (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetIsCollateralized(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetMinTeleporterVersion(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetMinTeleporterVersion(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetMultiplyOnRemote(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getMultiplyOnRemote")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetMultiplyOnRemote() (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetMultiplyOnRemote(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetMultiplyOnRemote() (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetMultiplyOnRemote(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetTokenHomeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getTokenHomeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetTokenHomeAddress() (common.Address, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenHomeAddress(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetTokenHomeAddress() (common.Address, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenHomeAddress(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetTokenHomeBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getTokenHomeBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetTokenHomeBlockchainID() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenHomeBlockchainID(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetTokenHomeBlockchainID() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenHomeBlockchainID(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetTokenMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getTokenMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetTokenMultiplier() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenMultiplier(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetTokenMultiplier() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenMultiplier(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.IsTeleporterAddressPaused(&_ERC20TokenRemoteUpgradeable.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.IsTeleporterAddressPaused(&_ERC20TokenRemoteUpgradeable.CallOpts, teleporterAddress)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Name() (string, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Name(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) Name() (string, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Name(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Owner() (common.Address, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Owner(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) Owner() (common.Address, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Owner(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Symbol() (string, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Symbol(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) Symbol() (string, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Symbol(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) TotalSupply() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TotalSupply(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TotalSupply(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Approve(&_ERC20TokenRemoteUpgradeable.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Approve(&_ERC20TokenRemoteUpgradeable.TransactOpts, spender, value)
}

// Initialize is a paid mutator transaction binding the contract method 0xc9fe4ddf.
//
// Solidity: function initialize((address,address,uint256,bytes32,address,uint8) settings, string tokenName, string tokenSymbol, uint8 tokenDecimals) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) Initialize(opts *bind.TransactOpts, settings TokenRemoteSettings, tokenName string, tokenSymbol string, tokenDecimals uint8) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "initialize", settings, tokenName, tokenSymbol, tokenDecimals)
}

// Initialize is a paid mutator transaction binding the contract method 0xc9fe4ddf.
//
// Solidity: function initialize((address,address,uint256,bytes32,address,uint8) settings, string tokenName, string tokenSymbol, uint8 tokenDecimals) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Initialize(settings TokenRemoteSettings, tokenName string, tokenSymbol string, tokenDecimals uint8) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Initialize(&_ERC20TokenRemoteUpgradeable.TransactOpts, settings, tokenName, tokenSymbol, tokenDecimals)
}

// Initialize is a paid mutator transaction binding the contract method 0xc9fe4ddf.
//
// Solidity: function initialize((address,address,uint256,bytes32,address,uint8) settings, string tokenName, string tokenSymbol, uint8 tokenDecimals) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) Initialize(settings TokenRemoteSettings, tokenName string, tokenSymbol string, tokenDecimals uint8) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Initialize(&_ERC20TokenRemoteUpgradeable.TransactOpts, settings, tokenName, tokenSymbol, tokenDecimals)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.PauseTeleporterAddress(&_ERC20TokenRemoteUpgradeable.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.PauseTeleporterAddress(&_ERC20TokenRemoteUpgradeable.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ReceiveTeleporterMessage(&_ERC20TokenRemoteUpgradeable.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ReceiveTeleporterMessage(&_ERC20TokenRemoteUpgradeable.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) RegisterWithHome(opts *bind.TransactOpts, feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "registerWithHome", feeInfo)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) RegisterWithHome(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.RegisterWithHome(&_ERC20TokenRemoteUpgradeable.TransactOpts, feeInfo)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) RegisterWithHome(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.RegisterWithHome(&_ERC20TokenRemoteUpgradeable.TransactOpts, feeInfo)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.RenounceOwnership(&_ERC20TokenRemoteUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.RenounceOwnership(&_ERC20TokenRemoteUpgradeable.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x5d16225d.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) Send(opts *bind.TransactOpts, input SendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "send", input, amount)
}

// Send is a paid mutator transaction binding the contract method 0x5d16225d.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Send(input SendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Send(&_ERC20TokenRemoteUpgradeable.TransactOpts, input, amount)
}

// Send is a paid mutator transaction binding the contract method 0x5d16225d.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) Send(input SendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Send(&_ERC20TokenRemoteUpgradeable.TransactOpts, input, amount)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x65690038.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) SendAndCall(opts *bind.TransactOpts, input SendAndCallInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "sendAndCall", input, amount)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x65690038.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) SendAndCall(input SendAndCallInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.SendAndCall(&_ERC20TokenRemoteUpgradeable.TransactOpts, input, amount)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x65690038.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) SendAndCall(input SendAndCallInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.SendAndCall(&_ERC20TokenRemoteUpgradeable.TransactOpts, input, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Transfer(&_ERC20TokenRemoteUpgradeable.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Transfer(&_ERC20TokenRemoteUpgradeable.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TransferFrom(&_ERC20TokenRemoteUpgradeable.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TransferFrom(&_ERC20TokenRemoteUpgradeable.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TransferOwnership(&_ERC20TokenRemoteUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TransferOwnership(&_ERC20TokenRemoteUpgradeable.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.UnpauseTeleporterAddress(&_ERC20TokenRemoteUpgradeable.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.UnpauseTeleporterAddress(&_ERC20TokenRemoteUpgradeable.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.UpdateMinTeleporterVersion(&_ERC20TokenRemoteUpgradeable.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.UpdateMinTeleporterVersion(&_ERC20TokenRemoteUpgradeable.TransactOpts, version)
}

// ERC20TokenRemoteUpgradeableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableApprovalIterator struct {
	Event *ERC20TokenRemoteUpgradeableApproval // Event containing the contract specifics and raw log

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
func (it *ERC20TokenRemoteUpgradeableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableApproval)
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
		it.Event = new(ERC20TokenRemoteUpgradeableApproval)
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
func (it *ERC20TokenRemoteUpgradeableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableApproval represents a Approval event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20TokenRemoteUpgradeableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableApprovalIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableApproval)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseApproval(log types.Log) (*ERC20TokenRemoteUpgradeableApproval, error) {
	event := new(ERC20TokenRemoteUpgradeableApproval)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableCallFailedIterator struct {
	Event *ERC20TokenRemoteUpgradeableCallFailed // Event containing the contract specifics and raw log

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
func (it *ERC20TokenRemoteUpgradeableCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableCallFailed)
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
		it.Event = new(ERC20TokenRemoteUpgradeableCallFailed)
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
func (it *ERC20TokenRemoteUpgradeableCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableCallFailed represents a CallFailed event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableCallFailed struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterCallFailed(opts *bind.FilterOpts, recipientContract []common.Address) (*ERC20TokenRemoteUpgradeableCallFailedIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableCallFailedIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableCallFailed, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableCallFailed)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "CallFailed", log); err != nil {
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
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseCallFailed(log types.Log) (*ERC20TokenRemoteUpgradeableCallFailed, error) {
	event := new(ERC20TokenRemoteUpgradeableCallFailed)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableCallSucceededIterator struct {
	Event *ERC20TokenRemoteUpgradeableCallSucceeded // Event containing the contract specifics and raw log

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
func (it *ERC20TokenRemoteUpgradeableCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableCallSucceeded)
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
		it.Event = new(ERC20TokenRemoteUpgradeableCallSucceeded)
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
func (it *ERC20TokenRemoteUpgradeableCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableCallSucceeded represents a CallSucceeded event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableCallSucceeded struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterCallSucceeded(opts *bind.FilterOpts, recipientContract []common.Address) (*ERC20TokenRemoteUpgradeableCallSucceededIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableCallSucceededIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableCallSucceeded, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableCallSucceeded)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseCallSucceeded(log types.Log) (*ERC20TokenRemoteUpgradeableCallSucceeded, error) {
	event := new(ERC20TokenRemoteUpgradeableCallSucceeded)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableInitializedIterator struct {
	Event *ERC20TokenRemoteUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *ERC20TokenRemoteUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableInitialized)
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
		it.Event = new(ERC20TokenRemoteUpgradeableInitialized)
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
func (it *ERC20TokenRemoteUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableInitialized represents a Initialized event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC20TokenRemoteUpgradeableInitializedIterator, error) {

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableInitializedIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableInitialized)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseInitialized(log types.Log) (*ERC20TokenRemoteUpgradeableInitialized, error) {
	event := new(ERC20TokenRemoteUpgradeableInitialized)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator struct {
	Event *ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated)
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
		it.Event = new(ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated)
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
func (it *ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdatedIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated, error) {
	event := new(ERC20TokenRemoteUpgradeableMinTeleporterVersionUpdated)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableOwnershipTransferredIterator struct {
	Event *ERC20TokenRemoteUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20TokenRemoteUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableOwnershipTransferred)
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
		it.Event = new(ERC20TokenRemoteUpgradeableOwnershipTransferred)
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
func (it *ERC20TokenRemoteUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20TokenRemoteUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableOwnershipTransferredIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableOwnershipTransferred)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20TokenRemoteUpgradeableOwnershipTransferred, error) {
	event := new(ERC20TokenRemoteUpgradeableOwnershipTransferred)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTeleporterAddressPausedIterator struct {
	Event *ERC20TokenRemoteUpgradeableTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *ERC20TokenRemoteUpgradeableTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableTeleporterAddressPaused)
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
		it.Event = new(ERC20TokenRemoteUpgradeableTeleporterAddressPaused)
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
func (it *ERC20TokenRemoteUpgradeableTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*ERC20TokenRemoteUpgradeableTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTeleporterAddressPausedIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableTeleporterAddressPaused)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseTeleporterAddressPaused(log types.Log) (*ERC20TokenRemoteUpgradeableTeleporterAddressPaused, error) {
	event := new(ERC20TokenRemoteUpgradeableTeleporterAddressPaused)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTeleporterAddressUnpausedIterator struct {
	Event *ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC20TokenRemoteUpgradeableTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused)
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
		it.Event = new(ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused)
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
func (it *ERC20TokenRemoteUpgradeableTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*ERC20TokenRemoteUpgradeableTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTeleporterAddressUnpausedIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused, error) {
	event := new(ERC20TokenRemoteUpgradeableTeleporterAddressUnpaused)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableTokensAndCallSentIterator is returned from FilterTokensAndCallSent and is used to iterate over the raw logs and unpacked data for TokensAndCallSent events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensAndCallSentIterator struct {
	Event *ERC20TokenRemoteUpgradeableTokensAndCallSent // Event containing the contract specifics and raw log

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
func (it *ERC20TokenRemoteUpgradeableTokensAndCallSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableTokensAndCallSent)
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
		it.Event = new(ERC20TokenRemoteUpgradeableTokensAndCallSent)
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
func (it *ERC20TokenRemoteUpgradeableTokensAndCallSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableTokensAndCallSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableTokensAndCallSent represents a TokensAndCallSent event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensAndCallSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterTokensAndCallSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*ERC20TokenRemoteUpgradeableTokensAndCallSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTokensAndCallSentIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "TokensAndCallSent", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchTokensAndCallSent(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableTokensAndCallSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableTokensAndCallSent)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
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
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseTokensAndCallSent(log types.Log) (*ERC20TokenRemoteUpgradeableTokensAndCallSent, error) {
	event := new(ERC20TokenRemoteUpgradeableTokensAndCallSent)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableTokensSentIterator is returned from FilterTokensSent and is used to iterate over the raw logs and unpacked data for TokensSent events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensSentIterator struct {
	Event *ERC20TokenRemoteUpgradeableTokensSent // Event containing the contract specifics and raw log

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
func (it *ERC20TokenRemoteUpgradeableTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableTokensSent)
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
		it.Event = new(ERC20TokenRemoteUpgradeableTokensSent)
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
func (it *ERC20TokenRemoteUpgradeableTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableTokensSent represents a TokensSent event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensSent is a free log retrieval operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*ERC20TokenRemoteUpgradeableTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTokensSentIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "TokensSent", logs: logs, sub: sub}, nil
}

// WatchTokensSent is a free log subscription operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchTokensSent(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableTokensSent)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensSent", log); err != nil {
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
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseTokensSent(log types.Log) (*ERC20TokenRemoteUpgradeableTokensSent, error) {
	event := new(ERC20TokenRemoteUpgradeableTokensSent)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensWithdrawnIterator struct {
	Event *ERC20TokenRemoteUpgradeableTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *ERC20TokenRemoteUpgradeableTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableTokensWithdrawn)
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
		it.Event = new(ERC20TokenRemoteUpgradeableTokensWithdrawn)
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
func (it *ERC20TokenRemoteUpgradeableTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableTokensWithdrawn represents a TokensWithdrawn event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*ERC20TokenRemoteUpgradeableTokensWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTokensWithdrawnIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableTokensWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableTokensWithdrawn)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseTokensWithdrawn(log types.Log) (*ERC20TokenRemoteUpgradeableTokensWithdrawn, error) {
	event := new(ERC20TokenRemoteUpgradeableTokensWithdrawn)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTransferIterator struct {
	Event *ERC20TokenRemoteUpgradeableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20TokenRemoteUpgradeableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableTransfer)
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
		it.Event = new(ERC20TokenRemoteUpgradeableTransfer)
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
func (it *ERC20TokenRemoteUpgradeableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableTransfer represents a Transfer event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TokenRemoteUpgradeableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTransferIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableTransfer)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseTransfer(log types.Log) (*ERC20TokenRemoteUpgradeableTransfer, error) {
	event := new(ERC20TokenRemoteUpgradeableTransfer)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
