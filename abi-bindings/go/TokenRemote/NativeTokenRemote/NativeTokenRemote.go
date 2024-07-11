// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokenremote

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

// NativeTokenRemoteMetaData contains all meta data concerning the NativeTokenRemote contract.
var NativeTokenRemoteMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"BURNED_FOR_TRANSFER_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BURNED_TX_FEES_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"HOME_CHAIN_BURN_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MULTI_HOP_CALL_GAS_PER_WORD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MULTI_HOP_CALL_REQUIRED_GAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MULTI_HOP_SEND_REQUIRED_GAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NATIVE_MINTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINativeMinter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTER_REMOTE_REQUIRED_GAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateNumWords\",\"inputs\":[{\"name\":\"payloadSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getBlockchainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInitialReserveImbalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIsCollateralized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinTeleporterVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMultiplyOnRemote\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenHomeAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenHomeBlockchainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenMultiplier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalMinted\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"settings\",\"type\":\"tuple\",\"internalType\":\"structTokenRemoteSettings\",\"components\":[{\"name\":\"teleporterRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"teleporterManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenHomeBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenHomeAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenHomeDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"nativeAssetSymbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"initialReserveImbalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"burnedFeesReportingRewardPercentage_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isTeleporterAddressPaused\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseTeleporterAddress\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"receiveTeleporterMessage\",\"inputs\":[{\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"originSenderAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerWithHome\",\"inputs\":[{\"name\":\"feeInfo\",\"type\":\"tuple\",\"internalType\":\"structTeleporterFeeInfo\",\"components\":[{\"name\":\"feeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reportBurnedTxFees\",\"inputs\":[{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"send\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structSendTokensInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"sendAndCall\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structSendAndCallInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipientGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"teleporterRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractTeleporterRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalNativeAssetSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseTeleporterAddress\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMinTeleporterVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"recipientContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSucceeded\",\"inputs\":[{\"name\":\"recipientContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinTeleporterVersionUpdated\",\"inputs\":[{\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReportBurnedTxFees\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"feesBurned\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TeleporterAddressPaused\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TeleporterAddressUnpaused\",\"inputs\":[{\"name\":\"teleporterAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensAndCallSent\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"input\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipientPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipientGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensSent\",\"inputs\":[{\"name\":\"teleporterMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"input\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"components\":[{\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"primaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"secondaryFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"multiHopFallback\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b5061533d806100206000396000f3fe6080604052600436106102815760003560e01c806370a082311161014f578063c3cd6927116100c1578063e0fd9cb81161007a578063e0fd9cb814610735578063ed0ae4b014610481578063ef793e2a1461074a578063f2fde38b1461075f578063f3f981d81461077f578063f56c363d1461079f57610290565b8063c3cd6927146106b3578063c452165e146106c8578063c868efaa146106e0578063d0e30db014610290578063d2cc7a7014610700578063dd62ed3e1461071557610290565b80638da5cb5b116101135780638da5cb5b146105e757806395d89b4114610605578063973142971461061a578063a457c2d714610653578063a9059cbb14610673578063b8a46d021461069357610290565b806370a082311461055d578063715018a61461059357806371717c18146105a85780637ee3779a146105bf5780638bf2fa94146105d457610290565b80632b0d8f18116101f35780634213cf78116101ac5780634213cf78146104be5780634511243e146104d35780635507f3d1146104f357806355538c8b1461050a5780635eb995141461052a5780636e6eef8d1461054a57610290565b80632b0d8f181461040a5780632e1a7d4d1461042a578063313ce5671461044a578063329c3e1214610466578063347212c414610481578063395093511461049e57610290565b806315beb59f1161024557806315beb59f1461035b57806318160ddd146103715780631906529c146103865780631a7f5bec1461039b57806323b872dd146103d3578063254ac160146103f357610290565b806302a30c7d1461029857806306fdde03146102c25780630733c8c8146102e4578063095ea7b3146103075780630ca1c5c91461032757610290565b366102905761028e6107bf565b005b61028e6107bf565b3480156102a457600080fd5b506102ad610800565b60405190151581526020015b60405180910390f35b3480156102ce57600080fd5b506102d7610818565b6040516102b99190614216565b3480156102f057600080fd5b506102f96108aa565b6040519081526020016102b9565b34801561031357600080fd5b506102ad610322366004614249565b6108bf565b34801561033357600080fd5b507f914a9547f6c3ddce1d5efbd9e687708f0d1d408ce129e8e1a88bce4f40e29501546102f9565b34801561036757600080fd5b506102f96105dc81565b34801561037d57600080fd5b506035546102f9565b34801561039257600080fd5b506102f96108d9565b3480156103a757600080fd5b506097546103bb906001600160a01b031681565b6040516001600160a01b0390911681526020016102b9565b3480156103df57600080fd5b506102ad6103ee366004614275565b610933565b3480156103ff57600080fd5b506102f96201fbd081565b34801561041657600080fd5b5061028e6104253660046142b6565b610957565b34801561043657600080fd5b5061028e6104453660046142d3565b610a59565b34801561045657600080fd5b50604051601281526020016102b9565b34801561047257600080fd5b506103bb6001600160991b0181565b34801561048d57600080fd5b506103bb62010203600160981b0181565b3480156104aa57600080fd5b506102ad6104b9366004614249565b610aa5565b3480156104ca57600080fd5b506102f9610ac7565b3480156104df57600080fd5b5061028e6104ee3660046142b6565b610ad9565b3480156104ff57600080fd5b506102f96205302081565b34801561051657600080fd5b5061028e6105253660046142d3565b610bd6565b34801561053657600080fd5b5061028e6105453660046142d3565b610ef2565b61028e6105583660046142ec565b610f03565b34801561056957600080fd5b506102f96105783660046142b6565b6001600160a01b031660009081526033602052604090205490565b34801561059f57600080fd5b5061028e610f31565b3480156105b457600080fd5b506102f96205573081565b3480156105cb57600080fd5b506102ad610f43565b61028e6105e2366004614327565b610f5b565b3480156105f357600080fd5b50609a546001600160a01b03166103bb565b34801561061157600080fd5b506102d7610f89565b34801561062657600080fd5b506102ad6106353660046142b6565b6001600160a01b031660009081526098602052604090205460ff1690565b34801561065f57600080fd5b506102ad61066e366004614249565b610f98565b34801561067f57600080fd5b506102ad61068e366004614249565b611013565b34801561069f57600080fd5b5061028e6106ae366004614340565b611021565b3480156106bf57600080fd5b506103bb611198565b3480156106d457600080fd5b506103bb600160981b81565b3480156106ec57600080fd5b5061028e6106fb366004614352565b6111b6565b34801561070c57600080fd5b506099546102f9565b34801561072157600080fd5b506102f96107303660046143da565b611362565b34801561074157600080fd5b506102f961138d565b34801561075657600080fd5b506102f96113a2565b34801561076b57600080fd5b5061028e61077a3660046142b6565b6113b7565b34801561078b57600080fd5b506102f961079a3660046142d3565b61142d565b3480156107ab57600080fd5b5061028e6107ba366004614543565b611444565b60405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a26107fe333461155d565b565b60008061080b61161f565b6006015460ff1692915050565b60606036805461082790614606565b80601f016020809104026020016040519081016040528092919081815260200182805461085390614606565b80156108a05780601f10610875576101008083540402835291602001916108a0565b820191906000526020600020905b81548152906001019060200180831161088357829003601f168201915b5050505050905090565b6000806108b561161f565b6003015492915050565b6000336108cd818585611643565b60019150505b92915050565b60006000805160206152e88339815191528161090462010203600160981b0131600160981b31614650565b905060006109106113a2565b836001015461091f9190614650565b905061092b8282614663565b935050505090565b600033610941858285611768565b61094c8585856117dc565b506001949350505050565b61095f611987565b6001600160a01b03811661098e5760405162461bcd60e51b815260040161098590614676565b60405180910390fd5b6001600160a01b03811660009081526098602052604090205460ff1615610a0d5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608401610985565b6001600160a01b038116600081815260986020526040808220805460ff19166001179055517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b60405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a2610a98338261198f565b610aa23382611ac0565b50565b6000336108cd818585610ab88383611362565b610ac29190614650565b611643565b600080610ad261161f565b5492915050565b610ae1611987565b6001600160a01b038116610b075760405162461bcd60e51b815260040161098590614676565b6001600160a01b03811660009081526098602052604090205460ff16610b815760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608401610985565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152609860205260409020805460ff19169055565b6000805160206152c88339815191528054600114610c065760405162461bcd60e51b8152600401610985906146c4565b600281557f914a9547f6c3ddce1d5efbd9e687708f0d1d408ce129e8e1a88bce4f40e29502546000805160206152e883398151915290600160981b31908111610cc55760405162461bcd60e51b8152602060048201526044602482018190527f4e6174697665546f6b656e52656d6f74653a206275726e206164647265737320908201527f62616c616e6365206e6f742067726561746572207468616e206c6173742072656064820152631c1bdc9d60e21b608482015260a401610985565b6000826002015482610cd79190614663565b905060006064846000015483610ced9190614708565b610cf7919061471f565b90506000610d058284614663565b6002860185905590508115610d2857610d1e3083611bd9565b610d28308361155d565b6000610d43610d356108aa565b610d3d610f43565b84611c8d565b11610dad5760405162461bcd60e51b815260206004820152603460248201527f4e6174697665546f6b656e52656d6f74653a207a65726f207363616c6564206160448201527336b7bab73a103a37903932b837b93a10313ab93760611b6064820152608401610985565b604080518082018252600181528151808301835262010203600160981b0181526020808201859052925160009380840192610dea92909101614757565b60405160208183030381529060405281525090506000610ea76040518060c00160405280610e1661138d565b8152602001610e23611198565b6001600160a01b0316815260408051808201825230815260208181018a905283015281018c905260600160005b604051908082528060200260200182016040528015610e79578160200160208202803683370190505b50815260200184604051602001610e909190614777565b604051602081830303815290604052815250611ca4565b9050807f0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a584604051610edb91815260200190565b60405180910390a250506001909555505050505050565b610efa611987565b610aa281611dca565b610f0b610800565b610f275760405162461bcd60e51b8152600401610985906147bc565b610aa28134611f54565b610f39611fcd565b6107fe6000612027565b600080610f4e61161f565b6004015460ff1692915050565b610f63610800565b610f7f5760405162461bcd60e51b8152600401610985906147bc565b610aa28134612079565b60606037805461082790614606565b60003381610fa68286611362565b9050838110156110065760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152608401610985565b61094c8286868403611643565b6000336108cd8185856117dc565b600061102b61161f565b6006810154909150610100900460ff16156110885760405162461bcd60e51b815260206004820152601f60248201527f546f6b656e52656d6f74653a20616c72656164792072656769737465726564006044820152606401610985565b604080516060808201835260058401548252600284015460ff600160a01b820481166020808601918252600160a81b9093048216858701908152865180880188526000808252885188518188015293518516848a0152915190931682860152865180830390950185526080909101909552808201929092529192909161111e90611114908701876142b6565b86602001356120e7565b6040805160c0810182526001870154815260028701546001600160a01b031660208083019190915282518084018452939450611190939192830191908190611168908b018b6142b6565b6001600160a01b0316815260209081018690529082526201fbd0908201526040016000610e50565b505050505050565b6000806111a361161f565b600201546001600160a01b031692915050565b6111be612131565b6099546097546001600160a01b0316634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015611217573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061123b919061480b565b10156112a25760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608401610985565b6112ab33610635565b156113115760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608401610985565b611352848484848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061218a92505050565b61135c6001606555565b50505050565b6001600160a01b03918216600090815260346020908152604080832093909416825291909152205490565b60008061139861161f565b6001015492915050565b6000806113ad61161f565b6005015492915050565b6113bf611fcd565b6001600160a01b0381166114245760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610985565b610aa281612027565b6000600561143c83601f614650565b901c92915050565b600054610100900460ff16158080156114645750600054600160ff909116105b8061147e5750303b15801561147e575060005460ff166001145b6114e15760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610985565b6000805460ff191660011790558015611504576000805461ff0019166101001790555b611510858585856123a6565b8015611556576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b6001600160a01b0382166115b35760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610985565b80603560008282546115c59190614650565b90915550506001600160a01b0382166000818152603360209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35b5050565b7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef0090565b6001600160a01b0383166116a55760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610985565b6001600160a01b0382166117065760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610985565b6001600160a01b0383811660008181526034602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b60006117748484611362565b9050600019811461135c57818110156117cf5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610985565b61135c8484848403611643565b6001600160a01b0383166118405760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610985565b6001600160a01b0382166118a25760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610985565b6001600160a01b0383166000908152603360205260409020548181101561191a5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610985565b6001600160a01b0380851660008181526033602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9061197a9086815260200190565b60405180910390a361135c565b6107fe611fcd565b6001600160a01b0382166119ef5760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610985565b6001600160a01b03821660009081526033602052604090205481811015611a635760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610985565b6001600160a01b03831660008181526033602090815260408083208686039055603580548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910161175b565b505050565b80471015611b105760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610985565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114611b5d576040519150601f19603f3d011682016040523d82523d6000602084013e611b62565b606091505b5050905080611abb5760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610985565b7f914a9547f6c3ddce1d5efbd9e687708f0d1d408ce129e8e1a88bce4f40e2950180546000805160206152e8833981519152918391600090611c1c908490614650565b90915550506040516327ad555d60e11b81526001600160a01b0384166004820152602481018390526001600160991b0190634f5aaaba90604401600060405180830381600087803b158015611c7057600080fd5b505af1158015611c84573d6000803e3d6000fd5b50505050505050565b6000611c9c8484846000612456565b949350505050565b600080611caf612487565b60408401516020015190915015611d54576040830151516001600160a01b0316611d315760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608401610985565b604083015160208101519051611d54916001600160a01b03909116908390612590565b604051630624488560e41b81526001600160a01b03821690636244885090611d80908690600401614824565b6020604051808303816000875af1158015611d9f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dc3919061480b565b9392505050565b6097546040805163301fd1f560e21b815290516000926001600160a01b03169163c07f47d49160048083019260209291908290030181865afa158015611e14573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e38919061480b565b60995490915081831115611ea85760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608401610985565b808311611f1d5760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608401610985565b6099839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b6000805160206152c88339815191528054600114611f845760405162461bcd60e51b8152600401610985906146c4565b600281556000611f9261161f565b9050611f9d84612675565b6001810154843503611fb957611fb384846128b4565b50611fc5565b611fb38484612adf565b505b600190555050565b609a546001600160a01b031633146107fe5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610985565b609a80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805160206152c883398151915280546001146120a95760405162461bcd60e51b8152600401610985906146c4565b6002815560006120b761161f565b90506120c284612d88565b60018101548435036120dd576120d88484612e76565b611fc3565b611fc38484612fea565b6000816000036120f9575060006108d3565b306001600160a01b0384160361212657612114333084611768565b61211f3330846117dc565b50806108d3565b611dc38333846131b5565b6002606554036121835760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610985565b6002606555565b600061219461161f565b9050806001015484146121fb5760405162461bcd60e51b815260206004820152602960248201527f546f6b656e52656d6f74653a20696e76616c696420736f7572636520626c6f636044820152681ad8da185a5b88125160ba1b6064820152608401610985565b60028101546001600160a01b0384811691161461226d5760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e52656d6f74653a20696e76616c6964206f726967696e2073656e646044820152696572206164647265737360b01b6064820152608401610985565b6000828060200190518101906122839190614922565b6006830154909150610100900460ff1615806122a45750600682015460ff16155b156122bb5760068201805461ffff19166101011790555b6001815160048111156122d0576122d0614741565b0361230957600081602001518060200190518101906122ef91906149b0565b90506123038160000151826020015161331e565b50611556565b60028151600481111561231e5761231e614741565b0361234d576000816020015180602001905181019061233d91906149ea565b905061230381826080015161336b565b60405162461bcd60e51b815260206004820152602160248201527f546f6b656e52656d6f74653a20696e76616c6964206d657373616765207479706044820152606560f81b6064820152608401610985565b6001606555565b600054610100900460ff166123cd5760405162461bcd60e51b815260040161098590614ab9565b816000036124375760405162461bcd60e51b815260206004820152603160248201527f4e6174697665546f6b656e52656d6f74653a207a65726f20696e697469616c206044820152707265736572766520696d62616c616e636560781b6064820152608401610985565b6124418384613498565b61244d848360126134c9565b61135c81613515565b6000811515841515036124745761246d8584614708565b9050611c9c565b61247e858461471f565b95945050505050565b600080609760009054906101000a90046001600160a01b03166001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156124dd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125019190614b04565b9050612525816001600160a01b031660009081526098602052604090205460ff1690565b1561258b5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608401610985565b919050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa1580156125e1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612605919061480b565b61260f9190614650565b6040516001600160a01b03851660248201526044810182905290915061135c90859063095ea7b360e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526135ac565b80356126935760405162461bcd60e51b815260040161098590614b21565b60006126a560408301602084016142b6565b6001600160a01b0316036126cb5760405162461bcd60e51b815260040161098590614b6c565b60006126dd60608301604084016142b6565b6001600160a01b0316036127485760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e7420636f6e7460448201526b72616374206164647265737360a01b6064820152608401610985565b600081608001351161276c5760405162461bcd60e51b815260040161098590614bc9565b60008160a00135116127ce5760405162461bcd60e51b815260206004820152602560248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e7420676173206044820152641b1a5b5a5d60da1b6064820152608401610985565b80608001358160a00135106128365760405162461bcd60e51b815260206004820152602860248201527f546f6b656e52656d6f74653a20696e76616c696420726563697069656e742067604482015267185cc81b1a5b5a5d60c21b6064820152608401610985565b6000612849610100830160e084016142b6565b6001600160a01b031603610aa25760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a207a65726f2066616c6c6261636b20726563697060448201526b69656e74206164647265737360a01b6064820152608401610985565b60006128be61161f565b90506128ee6128d360408501602086016142b6565b6101408501356128e960e0870160c088016142b6565b61367e565b600061291783612906610120870161010088016142b6565b86610120013587610140013561377c565b604080518082019091529194509150600090806002815260200160405180610100016040528086600001548152602001306001600160a01b0316815260200161295d3390565b6001600160a01b0316815260200161297b60608a0160408b016142b6565b6001600160a01b031681526020810188905260400161299d60608a018a614c0d565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060a089013560208201526040016129f26101008a0160e08b016142b6565b6001600160a01b03169052604051612a0d9190602001614c53565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b0316602082015281518083018352929350600092612a919282019080612a686101208c016101008d016142b6565b6001600160a01b03168152602090810188905290825260808a0135908201526040016000610e50565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b168888604051612acf929190614d60565b60405180910390a3505050505050565b6000612ae961161f565b9050612b158335612b0060408601602087016142b6565b612b1060e0870160c088016142b6565b613841565b6000612b2d83612906610120870161010088016142b6565b6040805180820190915291945091506000908060048152602001604051806101600160405280612b5a3390565b6001600160a01b0316815260200188600001358152602001886020016020810190612b8591906142b6565b6001600160a01b03168152602001612ba360608a0160408b016142b6565b6001600160a01b0316815260208101889052604001612bc560608a018a614c0d565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060a08901356020820152604001612c1a6101008a0160e08b016142b6565b6001600160a01b0316815260808901356020820152604001612c4260e08a0160c08b016142b6565b6001600160a01b03168152610140890135602091820152604051612c67929101614e6f565b60408051601f198184030181529190529052905060006105dc612c97612c906060890189614c0d565b905061142d565b612ca19190614708565b612cae9062055730614650565b6040805160c0810182526001870154815260028701546001600160a01b0316602082015281518083018352929350600092612d399282019080612cf96101208d016101008e016142b6565b6001600160a01b031681526020908101899052908252818101869052604080516000815280830182528184015251606090920191610e9091889101614777565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b168989604051612d77929190614d60565b60405180910390a350505050505050565b6000612d9a60608301604084016142b6565b6001600160a01b031603612dfc5760405162461bcd60e51b815260206004820152602360248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e74206164647260448201526265737360e81b6064820152608401610985565b60008160c0013511612e205760405162461bcd60e51b815260040161098590614bc9565b8035612e3e5760405162461bcd60e51b815260040161098590614b21565b6000612e5060408301602084016142b6565b6001600160a01b031603610aa25760405162461bcd60e51b815260040161098590614b6c565b6000612e8061161f565b9050612eab612e9560408501602086016142b6565b60a08501356128e9610100870160e088016142b6565b6000612ed083612ec160808701606088016142b6565b86608001358760a0013561377c565b60408051808201909152919450915060009080600181526020016040518060400160405280886040016020810190612f0891906142b6565b6001600160a01b0316815260200187815250604051602001612f2a9190614757565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b0316602082015281518083018352929350600092612fac9282019080612f8360808c0160608d016142b6565b6001600160a01b03168152602090810188905290825260c08a0135908201526040016000610e50565b9050336001600160a01b0316817f93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb528888604051612acf929190614f4d565b6000612ff461161f565b905061301c833561300b60408601602087016142b6565b612b10610100870160e088016142b6565b600061303283612ec160808701606088016142b6565b60408051808201825260038152815160e081018352883581529396509193506000926020808401928282019161306c918b01908b016142b6565b6001600160a01b0316815260200161308a60608a0160408b016142b6565b6001600160a01b031681526020810188905260a0890135604082015260c089013560608201526080016130c46101008a0160e08b016142b6565b6001600160a01b031690526040516131349190602001815181526020808301516001600160a01b0390811691830191909152604080840151821690830152606080840151908301526080808401519083015260a0808401519083015260c092830151169181019190915260e00190565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b0316602082015281518083018352929350600092612fac928201908061318d60808c0160608d016142b6565b6001600160a01b03168152602090810188905290825262053020908201526040016000610e50565b6040516370a0823160e01b815230600482015260009081906001600160a01b038616906370a0823190602401602060405180830381865afa1580156131fe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613222919061480b565b90506132396001600160a01b0386168530866138e0565b6040516370a0823160e01b81523060048201526000906001600160a01b038716906370a0823190602401602060405180830381865afa158015613280573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132a4919061480b565b905081811161330a5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610985565b6133148282614663565b9695505050505050565b816001600160a01b03167f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b8260405161335991815260200190565b60405180910390a261161b8282611bd9565b6133753082611bd9565b60008260000151836020015184604001518560a0015160405160240161339e9493929190614fec565b60408051601f198184030181529190526020810180516001600160e01b031663161b12ff60e11b17905260c084015160608501519192506000916133e59190859085613918565b905080156134395783606001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff48460405161342c91815260200190565b60405180910390a261135c565b83606001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb08460405161347891815260200190565b60405180910390a260e084015161135c906001600160a01b031684611ac0565b600054610100900460ff166134bf5760405162461bcd60e51b815260040161098590614ab9565b61161b82826139ed565b600054610100900460ff166134f05760405162461bcd60e51b815260040161098590614ab9565b61350283600001518460200151613a2d565b61350a613a6e565b611abb838383613aac565b600054610100900460ff1661353c5760405162461bcd60e51b815260040161098590614ab9565b6000805160206152e8833981519152606482106135a95760405162461bcd60e51b815260206004820152602560248201527f4e6174697665546f6b656e52656d6f74653a20696e76616c69642070657263656044820152646e7461676560d81b6064820152608401610985565b55565b6000613601826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613df29092919063ffffffff16565b805190915015611abb578080602001905181019061361f919061501e565b611abb5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610985565b600061368861161f565b60028101549091506001600160a01b038581169116146136ba5760405162461bcd60e51b815260040161098590615040565b82156137145760405162461bcd60e51b815260206004820152602360248201527f546f6b656e52656d6f74653a206e6f6e2d7a65726f207365636f6e646172792060448201526266656560e81b6064820152608401610985565b6001600160a01b0382161561135c5760405162461bcd60e51b815260206004820152602860248201527f546f6b656e52656d6f74653a206e6f6e2d7a65726f206d756c74692d686f702060448201526766616c6c6261636b60c01b6064820152608401610985565b600080600061378961161f565b905061379586866120e7565b94506137a087613e01565b600382015460048301549198506137ba9160ff1686611c8d565b600382015460048301546137d2919060ff168a611c8d565b116138345760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a20696e73756666696369656e7420746f6b656e7360448201526b103a37903a3930b739b332b960a11b6064820152608401610985565b5094959294509192505050565b600061384b61161f565b8054909150840361387e57306001600160a01b0384160361387e5760405162461bcd60e51b815260040161098590615040565b6001600160a01b03821661135c5760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a207a65726f206d756c74692d686f702066616c6c6044820152636261636b60e01b6064820152608401610985565b6040516001600160a01b038085166024830152831660448201526064810182905261135c9085906323b872dd60e01b9060840161263e565b6000845a101561396a5760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e742067617300000000006044820152606401610985565b834710156139ba5760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c75650000006044820152606401610985565b826001600160a01b03163b6000036139d457506000611c9c565b600080600084516020860188888bf19695505050505050565b600054610100900460ff16613a145760405162461bcd60e51b815260040161098590614ab9565b6036613a2083826150e3565b506037611abb82826150e3565b600054610100900460ff16613a545760405162461bcd60e51b815260040161098590614ab9565b613a5d82613e1a565b613a65613f48565b61161b816113b7565b600054610100900460ff16613a955760405162461bcd60e51b815260040161098590614ab9565b6107fe60016000805160206152c883398151915255565b600054610100900460ff16613ad35760405162461bcd60e51b815260040161098590614ab9565b6000613add61161f565b90506005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015613b24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b48919061480b565b81556040840151613bae5760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e52656d6f74653a207a65726f20746f6b656e20686f6d6520626c6f60448201526918dad8da185a5b88125160b21b6064820152608401610985565b8054604085015103613c285760405162461bcd60e51b815260206004820152603b60248201527f546f6b656e52656d6f74653a2063616e6e6f74206465706c6f7920746f20736160448201527f6d6520626c6f636b636861696e20617320746f6b656e20686f6d6500000000006064820152608401610985565b60608401516001600160a01b0316613c8e5760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a207a65726f20746f6b656e20686f6d65206164646044820152637265737360e01b6064820152608401610985565b6012846080015160ff161115613cf85760405162461bcd60e51b815260206004820152602960248201527f546f6b656e52656d6f74653a20746f6b656e20686f6d6520646563696d616c73604482015268040e8dede40d0d2ced60bb1b6064820152608401610985565b60128260ff161115613d585760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a20746f6b656e20646563696d616c7320746f6f206044820152630d0d2ced60e31b6064820152608401610985565b60408401516001820155606084015160028201805460058401869055600684018054871560ff1990911617905560808701516001600160a01b039093166001600160a81b031990911617600160a01b60ff808516919091029190911760ff60a81b1916600160a81b91861691909102179055613dd49083613f77565b60048301805460ff1916911515919091179055600390910155505050565b6060611c9c8484600085613fc2565b6000613e1662010203600160981b0183611ac0565b5090565b600054610100900460ff16613e415760405162461bcd60e51b815260040161098590614ab9565b613e4961409d565b6001600160a01b038116613ec55760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f7274657220726567697374727920616464726573730000000000000000006064820152608401610985565b609780546001600160a01b0319166001600160a01b0383169081179091556040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa158015613f1e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613f42919061480b565b60995550565b600054610100900460ff16613f6f5760405162461bcd60e51b815260040161098590614ab9565b6107fe6140cc565b60008060ff808516908416118181613f9b57613f9385876151a2565b60ff16613fa9565b613fa586866151a2565b60ff165b613fb490600a61529f565b9350909150505b9250929050565b6060824710156140235760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610985565b600080866001600160a01b0316858760405161403f91906152ab565b60006040518083038185875af1925050503d806000811461407c576040519150601f19603f3d011682016040523d82523d6000602084013e614081565b606091505b5091509150614092878383876140fc565b979650505050505050565b600054610100900460ff166140c45760405162461bcd60e51b815260040161098590614ab9565b6107fe614175565b600054610100900460ff166140f35760405162461bcd60e51b815260040161098590614ab9565b6107fe33612027565b6060831561416b578251600003614164576001600160a01b0385163b6141645760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610985565b5081611c9c565b611c9c838361419c565b600054610100900460ff1661239f5760405162461bcd60e51b815260040161098590614ab9565b8151156141ac5781518083602001fd5b8060405162461bcd60e51b81526004016109859190614216565b60005b838110156141e15781810151838201526020016141c9565b50506000910152565b600081518084526142028160208601602086016141c6565b601f01601f19169290920160200192915050565b602081526000611dc360208301846141ea565b6001600160a01b0381168114610aa257600080fd5b803561258b81614229565b6000806040838503121561425c57600080fd5b823561426781614229565b946020939093013593505050565b60008060006060848603121561428a57600080fd5b833561429581614229565b925060208401356142a581614229565b929592945050506040919091013590565b6000602082840312156142c857600080fd5b8135611dc381614229565b6000602082840312156142e557600080fd5b5035919050565b6000602082840312156142fe57600080fd5b81356001600160401b0381111561431457600080fd5b82016101608185031215611dc357600080fd5b6000610100828403121561433a57600080fd5b50919050565b60006040828403121561433a57600080fd5b6000806000806060858703121561436857600080fd5b84359350602085013561437a81614229565b925060408501356001600160401b038082111561439657600080fd5b818701915087601f8301126143aa57600080fd5b8135818111156143b957600080fd5b8860208285010111156143cb57600080fd5b95989497505060200194505050565b600080604083850312156143ed57600080fd5b82356143f881614229565b9150602083013561440881614229565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b038111828210171561444b5761444b614413565b60405290565b604080519081016001600160401b038111828210171561444b5761444b614413565b60405161010081016001600160401b038111828210171561444b5761444b614413565b604051601f8201601f191681016001600160401b03811182821017156144be576144be614413565b604052919050565b60006001600160401b038211156144df576144df614413565b50601f01601f191660200190565b600082601f8301126144fe57600080fd5b813561451161450c826144c6565b614496565b81815284602083860101111561452657600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008084860361010081121561455b57600080fd5b60a081121561456957600080fd5b50614572614429565b853561457d81614229565b8152602086013561458d81614229565b60208201526040868101359082015260608601356145aa81614229565b6060820152608086013560ff811681146145c357600080fd5b6080820152935060a08501356001600160401b038111156145e357600080fd5b6145ef878288016144ed565b949794965050505060c08301359260e00135919050565b600181811c9082168061461a57607f821691505b60208210810361433a57634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b808201808211156108d3576108d361463a565b818103818111156108d3576108d361463a565b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b80820281158282048414176108d3576108d361463a565b60008261473c57634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052602160045260246000fd5b81516001600160a01b0316815260208083015190820152604081016108d3565b60208152600082516005811061479d57634e487b7160e01b600052602160045260246000fd5b806020840152506020830151604080840152611c9c60608401826141ea565b6020808252602f908201527f4e6174697665546f6b656e52656d6f74653a20636f6e747261637420756e646560408201526e1c98dbdb1b185d195c985b1a5e9959608a1b606082015260800190565b60006020828403121561481d57600080fd5b5051919050565b6020808252825182820152828101516001600160a01b039081166040808501919091528401518051821660608501528083015160808501526000929161010085019190606087015160a0870152608087015160e060c0880152805193849052840192600092506101208701905b808410156148b357845183168252938501936001939093019290850190614891565b5060a0880151878203601f190160e089015294506148d181866141ea565b98975050505050505050565b600082601f8301126148ee57600080fd5b81516148fc61450c826144c6565b81815284602083860101111561491157600080fd5b611c9c8260208301602087016141c6565b60006020828403121561493457600080fd5b81516001600160401b038082111561494b57600080fd5b908301906040828603121561495f57600080fd5b614967614451565b82516005811061497657600080fd5b815260208301518281111561498a57600080fd5b614996878286016148dd565b60208301525095945050505050565b805161258b81614229565b6000604082840312156149c257600080fd5b6149ca614451565b82516149d581614229565b81526020928301519281019290925250919050565b6000602082840312156149fc57600080fd5b81516001600160401b0380821115614a1357600080fd5b908301906101008286031215614a2857600080fd5b614a30614473565b82518152614a40602084016149a5565b6020820152614a51604084016149a5565b6040820152614a62606084016149a5565b60608201526080830151608082015260a083015182811115614a8357600080fd5b614a8f878286016148dd565b60a08301525060c083015160c0820152614aab60e084016149a5565b60e082015295945050505050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b600060208284031215614b1657600080fd5b8151611dc381614229565b6020808252602b908201527f546f6b656e52656d6f74653a207a65726f2064657374696e6174696f6e20626c60408201526a1bd8dad8da185a5b88125160aa1b606082015260800190565b60208082526037908201527f546f6b656e52656d6f74653a207a65726f2064657374696e6174696f6e20746f60408201527f6b656e207472616e736665727265722061646472657373000000000000000000606082015260800190565b60208082526024908201527f546f6b656e52656d6f74653a207a65726f20726571756972656420676173206c6040820152631a5b5a5d60e21b606082015260800190565b6000808335601e19843603018112614c2457600080fd5b8301803591506001600160401b03821115614c3e57600080fd5b602001915036819003821315613fbb57600080fd5b60208152815160208201526000602083015160018060a01b03808216604085015280604086015116606085015250506060830151614c9c60808401826001600160a01b03169052565b50608083015160a083015260a08301516101008060c0850152614cc36101208501836141ea565b915060c085015160e085015260e0850151614ce8828601826001600160a01b03169052565b5090949350505050565b6000808335601e19843603018112614d0957600080fd5b83016020810192503590506001600160401b03811115614d2857600080fd5b803603821315613fbb57600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60408152823560408201526000614d796020850161423e565b6001600160a01b03166060830152614d936040850161423e565b6001600160a01b03166080830152614dae6060850185614cf2565b6101608060a0860152614dc66101a086018385614d37565b9250608087013560c086015260a087013560e0860152614de860c0880161423e565b9150610100614e01818701846001600160a01b03169052565b614e0d60e0890161423e565b9250610120614e26818801856001600160a01b03169052565b614e31828a0161423e565b93506101409150614e4c828801856001600160a01b03169052565b880135918601919091529095013561018084015260209092019290925292915050565b60208152614e896020820183516001600160a01b03169052565b6020820151604082015260006040830151614eaf60608401826001600160a01b03169052565b5060608301516001600160a01b038116608084015250608083015160a083015260a08301516101608060c0850152614eeb6101808501836141ea565b915060c085015160e085015260e0850151610100614f13818701836001600160a01b03169052565b860151610120868101919091528601519050610140614f3c818701836001600160a01b03169052565b959095015193019290925250919050565b8235815261012081016020840135614f6481614229565b6001600160a01b039081166020840152604085013590614f8382614229565b166040830152614f956060850161423e565b6001600160a01b0381166060840152506080840135608083015260a084013560a083015260c084013560c0830152614fcf60e0850161423e565b6001600160a01b031660e083015261010090910191909152919050565b8481526001600160a01b03848116602083015283166040820152608060608201819052600090613314908301846141ea565b60006020828403121561503057600080fd5b81518015158114611dc357600080fd5b6020808252603a908201527f546f6b656e52656d6f74653a20696e76616c69642064657374696e6174696f6e60408201527f20746f6b656e207472616e736665727265722061646472657373000000000000606082015260800190565b601f821115611abb57600081815260208120601f850160051c810160208610156150c45750805b601f850160051c820191505b81811015611190578281556001016150d0565b81516001600160401b038111156150fc576150fc614413565b6151108161510a8454614606565b8461509d565b602080601f831160018114615145576000841561512d5750858301515b600019600386901b1c1916600185901b178555611190565b600085815260208120601f198616915b8281101561517457888601518255948401946001909101908401615155565b50858210156151925787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60ff82811682821603908111156108d3576108d361463a565b600181815b808511156151f65781600019048211156151dc576151dc61463a565b808516156151e957918102915b93841c93908002906151c0565b509250929050565b60008261520d575060016108d3565b8161521a575060006108d3565b8160018114615230576002811461523a57615256565b60019150506108d3565b60ff84111561524b5761524b61463a565b50506001821b6108d3565b5060208310610133831016604e8410600b8410161715615279575081810a6108d3565b61528383836151bb565b80600019048211156152975761529761463a565b029392505050565b6000611dc383836151fe565b600082516152bd8184602087016141c6565b919091019291505056fed2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c7500914a9547f6c3ddce1d5efbd9e687708f0d1d408ce129e8e1a88bce4f40e29500a26469706673582212207557e2b303df06846fb8ab74a00641a8920fae4b873f9bfcc0df54383a6c841464736f6c63430008120033",
}

// NativeTokenRemoteABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenRemoteMetaData.ABI instead.
var NativeTokenRemoteABI = NativeTokenRemoteMetaData.ABI

// NativeTokenRemoteBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenRemoteMetaData.Bin instead.
var NativeTokenRemoteBin = NativeTokenRemoteMetaData.Bin

// DeployNativeTokenRemote deploys a new Ethereum contract, binding an instance of NativeTokenRemote to it.
func DeployNativeTokenRemote(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NativeTokenRemote, error) {
	parsed, err := NativeTokenRemoteMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenRemoteBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenRemote{NativeTokenRemoteCaller: NativeTokenRemoteCaller{contract: contract}, NativeTokenRemoteTransactor: NativeTokenRemoteTransactor{contract: contract}, NativeTokenRemoteFilterer: NativeTokenRemoteFilterer{contract: contract}}, nil
}

// NativeTokenRemote is an auto generated Go binding around an Ethereum contract.
type NativeTokenRemote struct {
	NativeTokenRemoteCaller     // Read-only binding to the contract
	NativeTokenRemoteTransactor // Write-only binding to the contract
	NativeTokenRemoteFilterer   // Log filterer for contract events
}

// NativeTokenRemoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenRemoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenRemoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenRemoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenRemoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenRemoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenRemoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenRemoteSession struct {
	Contract     *NativeTokenRemote // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NativeTokenRemoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenRemoteCallerSession struct {
	Contract *NativeTokenRemoteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// NativeTokenRemoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenRemoteTransactorSession struct {
	Contract     *NativeTokenRemoteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NativeTokenRemoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenRemoteRaw struct {
	Contract *NativeTokenRemote // Generic contract binding to access the raw methods on
}

// NativeTokenRemoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenRemoteCallerRaw struct {
	Contract *NativeTokenRemoteCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenRemoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenRemoteTransactorRaw struct {
	Contract *NativeTokenRemoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenRemote creates a new instance of NativeTokenRemote, bound to a specific deployed contract.
func NewNativeTokenRemote(address common.Address, backend bind.ContractBackend) (*NativeTokenRemote, error) {
	contract, err := bindNativeTokenRemote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemote{NativeTokenRemoteCaller: NativeTokenRemoteCaller{contract: contract}, NativeTokenRemoteTransactor: NativeTokenRemoteTransactor{contract: contract}, NativeTokenRemoteFilterer: NativeTokenRemoteFilterer{contract: contract}}, nil
}

// NewNativeTokenRemoteCaller creates a new read-only instance of NativeTokenRemote, bound to a specific deployed contract.
func NewNativeTokenRemoteCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenRemoteCaller, error) {
	contract, err := bindNativeTokenRemote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteCaller{contract: contract}, nil
}

// NewNativeTokenRemoteTransactor creates a new write-only instance of NativeTokenRemote, bound to a specific deployed contract.
func NewNativeTokenRemoteTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenRemoteTransactor, error) {
	contract, err := bindNativeTokenRemote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteTransactor{contract: contract}, nil
}

// NewNativeTokenRemoteFilterer creates a new log filterer instance of NativeTokenRemote, bound to a specific deployed contract.
func NewNativeTokenRemoteFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenRemoteFilterer, error) {
	contract, err := bindNativeTokenRemote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteFilterer{contract: contract}, nil
}

// bindNativeTokenRemote binds a generic wrapper to an already deployed contract.
func bindNativeTokenRemote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenRemoteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenRemote *NativeTokenRemoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenRemote.Contract.NativeTokenRemoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenRemote *NativeTokenRemoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.NativeTokenRemoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenRemote *NativeTokenRemoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.NativeTokenRemoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenRemote *NativeTokenRemoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenRemote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenRemote *NativeTokenRemoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenRemote *NativeTokenRemoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.contract.Transact(opts, method, params...)
}

// BURNEDFORTRANSFERADDRESS is a free data retrieval call binding the contract method 0x347212c4.
//
// Solidity: function BURNED_FOR_TRANSFER_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCaller) BURNEDFORTRANSFERADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "BURNED_FOR_TRANSFER_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDFORTRANSFERADDRESS is a free data retrieval call binding the contract method 0x347212c4.
//
// Solidity: function BURNED_FOR_TRANSFER_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteSession) BURNEDFORTRANSFERADDRESS() (common.Address, error) {
	return _NativeTokenRemote.Contract.BURNEDFORTRANSFERADDRESS(&_NativeTokenRemote.CallOpts)
}

// BURNEDFORTRANSFERADDRESS is a free data retrieval call binding the contract method 0x347212c4.
//
// Solidity: function BURNED_FOR_TRANSFER_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) BURNEDFORTRANSFERADDRESS() (common.Address, error) {
	return _NativeTokenRemote.Contract.BURNEDFORTRANSFERADDRESS(&_NativeTokenRemote.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCaller) BURNEDTXFEESADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "BURNED_TX_FEES_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenRemote.Contract.BURNEDTXFEESADDRESS(&_NativeTokenRemote.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenRemote.Contract.BURNEDTXFEESADDRESS(&_NativeTokenRemote.CallOpts)
}

// HOMECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xed0ae4b0.
//
// Solidity: function HOME_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCaller) HOMECHAINBURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "HOME_CHAIN_BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HOMECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xed0ae4b0.
//
// Solidity: function HOME_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteSession) HOMECHAINBURNADDRESS() (common.Address, error) {
	return _NativeTokenRemote.Contract.HOMECHAINBURNADDRESS(&_NativeTokenRemote.CallOpts)
}

// HOMECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xed0ae4b0.
//
// Solidity: function HOME_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) HOMECHAINBURNADDRESS() (common.Address, error) {
	return _NativeTokenRemote.Contract.HOMECHAINBURNADDRESS(&_NativeTokenRemote.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) MULTIHOPCALLGASPERWORD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "MULTI_HOP_CALL_GAS_PER_WORD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _NativeTokenRemote.Contract.MULTIHOPCALLGASPERWORD(&_NativeTokenRemote.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _NativeTokenRemote.Contract.MULTIHOPCALLGASPERWORD(&_NativeTokenRemote.CallOpts)
}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) MULTIHOPCALLREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "MULTI_HOP_CALL_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) MULTIHOPCALLREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemote.Contract.MULTIHOPCALLREQUIREDGAS(&_NativeTokenRemote.CallOpts)
}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) MULTIHOPCALLREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemote.Contract.MULTIHOPCALLREQUIREDGAS(&_NativeTokenRemote.CallOpts)
}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) MULTIHOPSENDREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "MULTI_HOP_SEND_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) MULTIHOPSENDREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemote.Contract.MULTIHOPSENDREQUIREDGAS(&_NativeTokenRemote.CallOpts)
}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) MULTIHOPSENDREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemote.Contract.MULTIHOPSENDREQUIREDGAS(&_NativeTokenRemote.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCaller) NATIVEMINTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "NATIVE_MINTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenRemote.Contract.NATIVEMINTER(&_NativeTokenRemote.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenRemote.Contract.NATIVEMINTER(&_NativeTokenRemote.CallOpts)
}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) REGISTERREMOTEREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "REGISTER_REMOTE_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) REGISTERREMOTEREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemote.Contract.REGISTERREMOTEREQUIREDGAS(&_NativeTokenRemote.CallOpts)
}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) REGISTERREMOTEREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenRemote.Contract.REGISTERREMOTEREQUIREDGAS(&_NativeTokenRemote.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NativeTokenRemote.Contract.Allowance(&_NativeTokenRemote.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NativeTokenRemote.Contract.Allowance(&_NativeTokenRemote.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NativeTokenRemote.Contract.BalanceOf(&_NativeTokenRemote.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NativeTokenRemote.Contract.BalanceOf(&_NativeTokenRemote.CallOpts, account)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) CalculateNumWords(opts *bind.CallOpts, payloadSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "calculateNumWords", payloadSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _NativeTokenRemote.Contract.CalculateNumWords(&_NativeTokenRemote.CallOpts, payloadSize)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _NativeTokenRemote.Contract.CalculateNumWords(&_NativeTokenRemote.CallOpts, payloadSize)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenRemote *NativeTokenRemoteCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenRemote *NativeTokenRemoteSession) Decimals() (uint8, error) {
	return _NativeTokenRemote.Contract.Decimals(&_NativeTokenRemote.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) Decimals() (uint8, error) {
	return _NativeTokenRemote.Contract.Decimals(&_NativeTokenRemote.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetBlockchainID() ([32]byte, error) {
	return _NativeTokenRemote.Contract.GetBlockchainID(&_NativeTokenRemote.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetBlockchainID() ([32]byte, error) {
	return _NativeTokenRemote.Contract.GetBlockchainID(&_NativeTokenRemote.CallOpts)
}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetInitialReserveImbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getInitialReserveImbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetInitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenRemote.Contract.GetInitialReserveImbalance(&_NativeTokenRemote.CallOpts)
}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetInitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenRemote.Contract.GetInitialReserveImbalance(&_NativeTokenRemote.CallOpts)
}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetIsCollateralized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getIsCollateralized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetIsCollateralized() (bool, error) {
	return _NativeTokenRemote.Contract.GetIsCollateralized(&_NativeTokenRemote.CallOpts)
}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetIsCollateralized() (bool, error) {
	return _NativeTokenRemote.Contract.GetIsCollateralized(&_NativeTokenRemote.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenRemote.Contract.GetMinTeleporterVersion(&_NativeTokenRemote.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenRemote.Contract.GetMinTeleporterVersion(&_NativeTokenRemote.CallOpts)
}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetMultiplyOnRemote(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getMultiplyOnRemote")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetMultiplyOnRemote() (bool, error) {
	return _NativeTokenRemote.Contract.GetMultiplyOnRemote(&_NativeTokenRemote.CallOpts)
}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetMultiplyOnRemote() (bool, error) {
	return _NativeTokenRemote.Contract.GetMultiplyOnRemote(&_NativeTokenRemote.CallOpts)
}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetTokenHomeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getTokenHomeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetTokenHomeAddress() (common.Address, error) {
	return _NativeTokenRemote.Contract.GetTokenHomeAddress(&_NativeTokenRemote.CallOpts)
}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetTokenHomeAddress() (common.Address, error) {
	return _NativeTokenRemote.Contract.GetTokenHomeAddress(&_NativeTokenRemote.CallOpts)
}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetTokenHomeBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getTokenHomeBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetTokenHomeBlockchainID() ([32]byte, error) {
	return _NativeTokenRemote.Contract.GetTokenHomeBlockchainID(&_NativeTokenRemote.CallOpts)
}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetTokenHomeBlockchainID() ([32]byte, error) {
	return _NativeTokenRemote.Contract.GetTokenHomeBlockchainID(&_NativeTokenRemote.CallOpts)
}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetTokenMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getTokenMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetTokenMultiplier() (*big.Int, error) {
	return _NativeTokenRemote.Contract.GetTokenMultiplier(&_NativeTokenRemote.CallOpts)
}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetTokenMultiplier() (*big.Int, error) {
	return _NativeTokenRemote.Contract.GetTokenMultiplier(&_NativeTokenRemote.CallOpts)
}

// GetTotalMinted is a free data retrieval call binding the contract method 0x0ca1c5c9.
//
// Solidity: function getTotalMinted() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) GetTotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "getTotalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalMinted is a free data retrieval call binding the contract method 0x0ca1c5c9.
//
// Solidity: function getTotalMinted() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) GetTotalMinted() (*big.Int, error) {
	return _NativeTokenRemote.Contract.GetTotalMinted(&_NativeTokenRemote.CallOpts)
}

// GetTotalMinted is a free data retrieval call binding the contract method 0x0ca1c5c9.
//
// Solidity: function getTotalMinted() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) GetTotalMinted() (*big.Int, error) {
	return _NativeTokenRemote.Contract.GetTotalMinted(&_NativeTokenRemote.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenRemote.Contract.IsTeleporterAddressPaused(&_NativeTokenRemote.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenRemote.Contract.IsTeleporterAddressPaused(&_NativeTokenRemote.CallOpts, teleporterAddress)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenRemote *NativeTokenRemoteCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenRemote *NativeTokenRemoteSession) Name() (string, error) {
	return _NativeTokenRemote.Contract.Name(&_NativeTokenRemote.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) Name() (string, error) {
	return _NativeTokenRemote.Contract.Name(&_NativeTokenRemote.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteSession) Owner() (common.Address, error) {
	return _NativeTokenRemote.Contract.Owner(&_NativeTokenRemote.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) Owner() (common.Address, error) {
	return _NativeTokenRemote.Contract.Owner(&_NativeTokenRemote.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenRemote *NativeTokenRemoteCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenRemote *NativeTokenRemoteSession) Symbol() (string, error) {
	return _NativeTokenRemote.Contract.Symbol(&_NativeTokenRemote.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) Symbol() (string, error) {
	return _NativeTokenRemote.Contract.Symbol(&_NativeTokenRemote.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCaller) TeleporterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "teleporterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteSession) TeleporterRegistry() (common.Address, error) {
	return _NativeTokenRemote.Contract.TeleporterRegistry(&_NativeTokenRemote.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) TeleporterRegistry() (common.Address, error) {
	return _NativeTokenRemote.Contract.TeleporterRegistry(&_NativeTokenRemote.CallOpts)
}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) TotalNativeAssetSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "totalNativeAssetSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) TotalNativeAssetSupply() (*big.Int, error) {
	return _NativeTokenRemote.Contract.TotalNativeAssetSupply(&_NativeTokenRemote.CallOpts)
}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) TotalNativeAssetSupply() (*big.Int, error) {
	return _NativeTokenRemote.Contract.TotalNativeAssetSupply(&_NativeTokenRemote.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenRemote.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenRemote.Contract.TotalSupply(&_NativeTokenRemote.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenRemote *NativeTokenRemoteCallerSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenRemote.Contract.TotalSupply(&_NativeTokenRemote.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Approve(&_NativeTokenRemote.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Approve(&_NativeTokenRemote.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.DecreaseAllowance(&_NativeTokenRemote.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.DecreaseAllowance(&_NativeTokenRemote.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) Deposit() (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Deposit(&_NativeTokenRemote.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) Deposit() (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Deposit(&_NativeTokenRemote.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.IncreaseAllowance(&_NativeTokenRemote.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.IncreaseAllowance(&_NativeTokenRemote.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xf56c363d.
//
// Solidity: function initialize((address,address,bytes32,address,uint8) settings, string nativeAssetSymbol, uint256 initialReserveImbalance, uint256 burnedFeesReportingRewardPercentage_) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) Initialize(opts *bind.TransactOpts, settings TokenRemoteSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage_ *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "initialize", settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf56c363d.
//
// Solidity: function initialize((address,address,bytes32,address,uint8) settings, string nativeAssetSymbol, uint256 initialReserveImbalance, uint256 burnedFeesReportingRewardPercentage_) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) Initialize(settings TokenRemoteSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage_ *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Initialize(&_NativeTokenRemote.TransactOpts, settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf56c363d.
//
// Solidity: function initialize((address,address,bytes32,address,uint8) settings, string nativeAssetSymbol, uint256 initialReserveImbalance, uint256 burnedFeesReportingRewardPercentage_) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) Initialize(settings TokenRemoteSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage_ *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Initialize(&_NativeTokenRemote.TransactOpts, settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage_)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.PauseTeleporterAddress(&_NativeTokenRemote.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.PauseTeleporterAddress(&_NativeTokenRemote.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.ReceiveTeleporterMessage(&_NativeTokenRemote.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.ReceiveTeleporterMessage(&_NativeTokenRemote.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) RegisterWithHome(opts *bind.TransactOpts, feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "registerWithHome", feeInfo)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) RegisterWithHome(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.RegisterWithHome(&_NativeTokenRemote.TransactOpts, feeInfo)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) RegisterWithHome(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.RegisterWithHome(&_NativeTokenRemote.TransactOpts, feeInfo)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.RenounceOwnership(&_NativeTokenRemote.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.RenounceOwnership(&_NativeTokenRemote.TransactOpts)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) ReportBurnedTxFees(opts *bind.TransactOpts, requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "reportBurnedTxFees", requiredGasLimit)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) ReportBurnedTxFees(requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.ReportBurnedTxFees(&_NativeTokenRemote.TransactOpts, requiredGasLimit)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) ReportBurnedTxFees(requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.ReportBurnedTxFees(&_NativeTokenRemote.TransactOpts, requiredGasLimit)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) Send(opts *bind.TransactOpts, input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Send(&_NativeTokenRemote.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Send(&_NativeTokenRemote.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) SendAndCall(opts *bind.TransactOpts, input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "sendAndCall", input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.SendAndCall(&_NativeTokenRemote.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.SendAndCall(&_NativeTokenRemote.TransactOpts, input)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Transfer(&_NativeTokenRemote.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Transfer(&_NativeTokenRemote.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.TransferFrom(&_NativeTokenRemote.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.TransferFrom(&_NativeTokenRemote.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.TransferOwnership(&_NativeTokenRemote.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.TransferOwnership(&_NativeTokenRemote.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.UnpauseTeleporterAddress(&_NativeTokenRemote.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.UnpauseTeleporterAddress(&_NativeTokenRemote.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.UpdateMinTeleporterVersion(&_NativeTokenRemote.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.UpdateMinTeleporterVersion(&_NativeTokenRemote.TransactOpts, version)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Withdraw(&_NativeTokenRemote.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Withdraw(&_NativeTokenRemote.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Fallback(&_NativeTokenRemote.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Fallback(&_NativeTokenRemote.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenRemote.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteSession) Receive() (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Receive(&_NativeTokenRemote.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenRemote *NativeTokenRemoteTransactorSession) Receive() (*types.Transaction, error) {
	return _NativeTokenRemote.Contract.Receive(&_NativeTokenRemote.TransactOpts)
}

// NativeTokenRemoteApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NativeTokenRemote contract.
type NativeTokenRemoteApprovalIterator struct {
	Event *NativeTokenRemoteApproval // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteApproval)
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
		it.Event = new(NativeTokenRemoteApproval)
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
func (it *NativeTokenRemoteApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteApproval represents a Approval event raised by the NativeTokenRemote contract.
type NativeTokenRemoteApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NativeTokenRemoteApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteApprovalIterator{contract: _NativeTokenRemote.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteApproval)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseApproval(log types.Log) (*NativeTokenRemoteApproval, error) {
	event := new(NativeTokenRemoteApproval)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the NativeTokenRemote contract.
type NativeTokenRemoteCallFailedIterator struct {
	Event *NativeTokenRemoteCallFailed // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteCallFailed)
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
		it.Event = new(NativeTokenRemoteCallFailed)
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
func (it *NativeTokenRemoteCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteCallFailed represents a CallFailed event raised by the NativeTokenRemote contract.
type NativeTokenRemoteCallFailed struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterCallFailed(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenRemoteCallFailedIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteCallFailedIterator{contract: _NativeTokenRemote.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteCallFailed, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteCallFailed)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "CallFailed", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseCallFailed(log types.Log) (*NativeTokenRemoteCallFailed, error) {
	event := new(NativeTokenRemoteCallFailed)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the NativeTokenRemote contract.
type NativeTokenRemoteCallSucceededIterator struct {
	Event *NativeTokenRemoteCallSucceeded // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteCallSucceeded)
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
		it.Event = new(NativeTokenRemoteCallSucceeded)
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
func (it *NativeTokenRemoteCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteCallSucceeded represents a CallSucceeded event raised by the NativeTokenRemote contract.
type NativeTokenRemoteCallSucceeded struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterCallSucceeded(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenRemoteCallSucceededIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteCallSucceededIterator{contract: _NativeTokenRemote.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteCallSucceeded, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteCallSucceeded)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseCallSucceeded(log types.Log) (*NativeTokenRemoteCallSucceeded, error) {
	event := new(NativeTokenRemoteCallSucceeded)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the NativeTokenRemote contract.
type NativeTokenRemoteDepositIterator struct {
	Event *NativeTokenRemoteDeposit // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteDeposit)
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
		it.Event = new(NativeTokenRemoteDeposit)
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
func (it *NativeTokenRemoteDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteDeposit represents a Deposit event raised by the NativeTokenRemote contract.
type NativeTokenRemoteDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenRemoteDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteDepositIterator{contract: _NativeTokenRemote.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteDeposit)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseDeposit(log types.Log) (*NativeTokenRemoteDeposit, error) {
	event := new(NativeTokenRemoteDeposit)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NativeTokenRemote contract.
type NativeTokenRemoteInitializedIterator struct {
	Event *NativeTokenRemoteInitialized // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteInitialized)
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
		it.Event = new(NativeTokenRemoteInitialized)
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
func (it *NativeTokenRemoteInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteInitialized represents a Initialized event raised by the NativeTokenRemote contract.
type NativeTokenRemoteInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterInitialized(opts *bind.FilterOpts) (*NativeTokenRemoteInitializedIterator, error) {

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteInitializedIterator{contract: _NativeTokenRemote.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteInitialized) (event.Subscription, error) {

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteInitialized)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseInitialized(log types.Log) (*NativeTokenRemoteInitialized, error) {
	event := new(NativeTokenRemoteInitialized)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the NativeTokenRemote contract.
type NativeTokenRemoteMinTeleporterVersionUpdatedIterator struct {
	Event *NativeTokenRemoteMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteMinTeleporterVersionUpdated)
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
		it.Event = new(NativeTokenRemoteMinTeleporterVersionUpdated)
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
func (it *NativeTokenRemoteMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the NativeTokenRemote contract.
type NativeTokenRemoteMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*NativeTokenRemoteMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteMinTeleporterVersionUpdatedIterator{contract: _NativeTokenRemote.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteMinTeleporterVersionUpdated)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*NativeTokenRemoteMinTeleporterVersionUpdated, error) {
	event := new(NativeTokenRemoteMinTeleporterVersionUpdated)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NativeTokenRemote contract.
type NativeTokenRemoteOwnershipTransferredIterator struct {
	Event *NativeTokenRemoteOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteOwnershipTransferred)
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
		it.Event = new(NativeTokenRemoteOwnershipTransferred)
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
func (it *NativeTokenRemoteOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteOwnershipTransferred represents a OwnershipTransferred event raised by the NativeTokenRemote contract.
type NativeTokenRemoteOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NativeTokenRemoteOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteOwnershipTransferredIterator{contract: _NativeTokenRemote.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteOwnershipTransferred)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseOwnershipTransferred(log types.Log) (*NativeTokenRemoteOwnershipTransferred, error) {
	event := new(NativeTokenRemoteOwnershipTransferred)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteReportBurnedTxFeesIterator is returned from FilterReportBurnedTxFees and is used to iterate over the raw logs and unpacked data for ReportBurnedTxFees events raised by the NativeTokenRemote contract.
type NativeTokenRemoteReportBurnedTxFeesIterator struct {
	Event *NativeTokenRemoteReportBurnedTxFees // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteReportBurnedTxFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteReportBurnedTxFees)
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
		it.Event = new(NativeTokenRemoteReportBurnedTxFees)
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
func (it *NativeTokenRemoteReportBurnedTxFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteReportBurnedTxFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteReportBurnedTxFees represents a ReportBurnedTxFees event raised by the NativeTokenRemote contract.
type NativeTokenRemoteReportBurnedTxFees struct {
	TeleporterMessageID [32]byte
	FeesBurned          *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReportBurnedTxFees is a free log retrieval operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterReportBurnedTxFees(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*NativeTokenRemoteReportBurnedTxFeesIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "ReportBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteReportBurnedTxFeesIterator{contract: _NativeTokenRemote.contract, event: "ReportBurnedTxFees", logs: logs, sub: sub}, nil
}

// WatchReportBurnedTxFees is a free log subscription operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchReportBurnedTxFees(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteReportBurnedTxFees, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "ReportBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteReportBurnedTxFees)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "ReportBurnedTxFees", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseReportBurnedTxFees(log types.Log) (*NativeTokenRemoteReportBurnedTxFees, error) {
	event := new(NativeTokenRemoteReportBurnedTxFees)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "ReportBurnedTxFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the NativeTokenRemote contract.
type NativeTokenRemoteTeleporterAddressPausedIterator struct {
	Event *NativeTokenRemoteTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteTeleporterAddressPaused)
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
		it.Event = new(NativeTokenRemoteTeleporterAddressPaused)
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
func (it *NativeTokenRemoteTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the NativeTokenRemote contract.
type NativeTokenRemoteTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenRemoteTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteTeleporterAddressPausedIterator{contract: _NativeTokenRemote.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteTeleporterAddressPaused)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseTeleporterAddressPaused(log types.Log) (*NativeTokenRemoteTeleporterAddressPaused, error) {
	event := new(NativeTokenRemoteTeleporterAddressPaused)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the NativeTokenRemote contract.
type NativeTokenRemoteTeleporterAddressUnpausedIterator struct {
	Event *NativeTokenRemoteTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteTeleporterAddressUnpaused)
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
		it.Event = new(NativeTokenRemoteTeleporterAddressUnpaused)
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
func (it *NativeTokenRemoteTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the NativeTokenRemote contract.
type NativeTokenRemoteTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenRemoteTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteTeleporterAddressUnpausedIterator{contract: _NativeTokenRemote.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteTeleporterAddressUnpaused)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*NativeTokenRemoteTeleporterAddressUnpaused, error) {
	event := new(NativeTokenRemoteTeleporterAddressUnpaused)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteTokensAndCallSentIterator is returned from FilterTokensAndCallSent and is used to iterate over the raw logs and unpacked data for TokensAndCallSent events raised by the NativeTokenRemote contract.
type NativeTokenRemoteTokensAndCallSentIterator struct {
	Event *NativeTokenRemoteTokensAndCallSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteTokensAndCallSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteTokensAndCallSent)
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
		it.Event = new(NativeTokenRemoteTokensAndCallSent)
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
func (it *NativeTokenRemoteTokensAndCallSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteTokensAndCallSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteTokensAndCallSent represents a TokensAndCallSent event raised by the NativeTokenRemote contract.
type NativeTokenRemoteTokensAndCallSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterTokensAndCallSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenRemoteTokensAndCallSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteTokensAndCallSentIterator{contract: _NativeTokenRemote.contract, event: "TokensAndCallSent", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchTokensAndCallSent(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteTokensAndCallSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteTokensAndCallSent)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseTokensAndCallSent(log types.Log) (*NativeTokenRemoteTokensAndCallSent, error) {
	event := new(NativeTokenRemoteTokensAndCallSent)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteTokensSentIterator is returned from FilterTokensSent and is used to iterate over the raw logs and unpacked data for TokensSent events raised by the NativeTokenRemote contract.
type NativeTokenRemoteTokensSentIterator struct {
	Event *NativeTokenRemoteTokensSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteTokensSent)
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
		it.Event = new(NativeTokenRemoteTokensSent)
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
func (it *NativeTokenRemoteTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteTokensSent represents a TokensSent event raised by the NativeTokenRemote contract.
type NativeTokenRemoteTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensSent is a free log retrieval operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenRemoteTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteTokensSentIterator{contract: _NativeTokenRemote.contract, event: "TokensSent", logs: logs, sub: sub}, nil
}

// WatchTokensSent is a free log subscription operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchTokensSent(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteTokensSent)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "TokensSent", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseTokensSent(log types.Log) (*NativeTokenRemoteTokensSent, error) {
	event := new(NativeTokenRemoteTokensSent)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "TokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the NativeTokenRemote contract.
type NativeTokenRemoteTokensWithdrawnIterator struct {
	Event *NativeTokenRemoteTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteTokensWithdrawn)
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
		it.Event = new(NativeTokenRemoteTokensWithdrawn)
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
func (it *NativeTokenRemoteTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteTokensWithdrawn represents a TokensWithdrawn event raised by the NativeTokenRemote contract.
type NativeTokenRemoteTokensWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*NativeTokenRemoteTokensWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteTokensWithdrawnIterator{contract: _NativeTokenRemote.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteTokensWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteTokensWithdrawn)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseTokensWithdrawn(log types.Log) (*NativeTokenRemoteTokensWithdrawn, error) {
	event := new(NativeTokenRemoteTokensWithdrawn)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NativeTokenRemote contract.
type NativeTokenRemoteTransferIterator struct {
	Event *NativeTokenRemoteTransfer // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteTransfer)
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
		it.Event = new(NativeTokenRemoteTransfer)
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
func (it *NativeTokenRemoteTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteTransfer represents a Transfer event raised by the NativeTokenRemote contract.
type NativeTokenRemoteTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTokenRemoteTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteTransferIterator{contract: _NativeTokenRemote.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteTransfer)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseTransfer(log types.Log) (*NativeTokenRemoteTransfer, error) {
	event := new(NativeTokenRemoteTransfer)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenRemoteWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the NativeTokenRemote contract.
type NativeTokenRemoteWithdrawalIterator struct {
	Event *NativeTokenRemoteWithdrawal // Event containing the contract specifics and raw log

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
func (it *NativeTokenRemoteWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenRemoteWithdrawal)
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
		it.Event = new(NativeTokenRemoteWithdrawal)
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
func (it *NativeTokenRemoteWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenRemoteWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenRemoteWithdrawal represents a Withdrawal event raised by the NativeTokenRemote contract.
type NativeTokenRemoteWithdrawal struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) FilterWithdrawal(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenRemoteWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.FilterLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenRemoteWithdrawalIterator{contract: _NativeTokenRemote.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenRemote *NativeTokenRemoteFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *NativeTokenRemoteWithdrawal, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenRemote.contract.WatchLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenRemoteWithdrawal)
				if err := _NativeTokenRemote.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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
func (_NativeTokenRemote *NativeTokenRemoteFilterer) ParseWithdrawal(log types.Log) (*NativeTokenRemoteWithdrawal, error) {
	event := new(NativeTokenRemoteWithdrawal)
	if err := _NativeTokenRemote.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
