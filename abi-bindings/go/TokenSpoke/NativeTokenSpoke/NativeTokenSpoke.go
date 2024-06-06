// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokenspoke

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

// TeleporterFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type TeleporterFeeInfo struct {
	FeeTokenAddress common.Address
	Amount          *big.Int
}

// TokenSpokeSettings is an auto generated low-level Go binding around an user-defined struct.
type TokenSpokeSettings struct {
	TeleporterRegistryAddress common.Address
	TeleporterManager         common.Address
	TokenHubBlockchainID      [32]byte
	TokenHubAddress           common.Address
	TokenHubDecimals          uint8
}

// NativeTokenSpokeMetaData contains all meta data concerning the NativeTokenSpoke contract.
var NativeTokenSpokeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesBurned\",\"type\":\"uint256\"}],\"name\":\"ReportBurnedTxFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"BURNED_FOR_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURNED_TX_FEES_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HUB_CHAIN_BURN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_CALL_GAS_PER_WORD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_CALL_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_SEND_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_MINTER\",\"outputs\":[{\"internalType\":\"contractINativeMinter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTER_SPOKE_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnedFeesReportingRewardPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"payloadSize\",\"type\":\"uint256\"}],\"name\":\"calculateNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hubTokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialReserveImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tokenHubBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenHubAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenHubDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenSpokeSettings\",\"name\":\"settings\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"nativeAssetSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialReserveImbalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burnedFeesReportingRewardPercentage_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastestBurnedFeesReported\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplyOnSpoke\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"registerWithHub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"name\":\"reportBurnedTxFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"sendAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenHubAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenHubBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalNativeAssetSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615118806100206000396000f3fe6080604052600436106102e85760003560e01c806371717c1811610190578063c868efaa116100dc578063eae0544f11610095578063f3f981d81161006f578063f3f981d81461083d578063f56c363d1461085d578063fdd7546f1461087d578063feea2b721461089e576102f7565b8063eae0544f14610533578063eb3da90914610807578063f2fde38b1461081d576102f7565b8063c868efaa14610786578063d0e30db0146102f7578063d10a5b8c146107a6578063d127dc9b146107bc578063d2cc7a70146107d2578063dd62ed3e146107e7576102f7565b8063a2309ff811610149578063ba3f5a1211610123578063ba3f5a1214610721578063bf48d36b14610737578063c03f401314610757578063c452165e1461076e576102f7565b8063a2309ff8146106cb578063a457c2d7146106e1578063a9059cbb14610701576102f7565b806371717c181461061f5780638ac7dd20146106365780638bf2fa941461064c5780638da5cb5b1461065f57806395d89b411461067d5780639731429714610692576102f7565b8063329c3e121161024f57806349e3284e116102085780635eb99514116101e25780635eb99514146105a15780636e6eef8d146105c157806370a08231146105d4578063715018a61461060a576102f7565b806349e3284e146105505780635507f3d11461056a57806355538c8b14610581576102f7565b8063329c3e12146104a157806339509351146104bc5780633a23dfe2146104dc5780633b97e856146104f25780634511243e1461051357806347a9a22c14610533576102f7565b806322366844116102a157806322366844146103e057806323b872dd146103ff578063244276291461041f5780632b0d8f181461043f5780632e1a7d4d1461045f578063313ce5671461047f576102f7565b806306fdde03146102ff578063095ea7b31461032a57806315beb59f1461035a57806318160ddd1461037e5780631906529c146103935780631a7f5bec146103a8576102f7565b366102f7576102f56108b8565b005b6102f56108b8565b34801561030b57600080fd5b506103146108f9565b6040516103219190614018565b60405180910390f35b34801561033657600080fd5b5061034a61034536600461404b565b61098b565b6040519015158152602001610321565b34801561036657600080fd5b506103706105dc81565b604051908152602001610321565b34801561038a57600080fd5b5060a454610370565b34801561039f57600080fd5b506103706109a5565b3480156103b457600080fd5b506065546103c8906001600160a01b031681565b6040516001600160a01b039091168152602001610321565b3480156103ec57600080fd5b5060a15461034a90610100900460ff1681565b34801561040b57600080fd5b5061034a61041a366004614077565b6109e8565b34801561042b57600080fd5b50609d546103c8906001600160a01b031681565b34801561044b57600080fd5b506102f561045a3660046140b8565b610a0c565b34801561046b57600080fd5b506102f561047a3660046140d5565b610b0e565b34801561048b57600080fd5b5060125b60405160ff9091168152602001610321565b3480156104ad57600080fd5b506103c86001600160991b0181565b3480156104c857600080fd5b5061034a6104d736600461404b565b610b5a565b3480156104e857600080fd5b5061037060d45481565b3480156104fe57600080fd5b50609d5461048f90600160a81b900460ff1681565b34801561051f57600080fd5b506102f561052e3660046140b8565b610b7c565b34801561053f57600080fd5b506103c862010203600160981b0181565b34801561055c57600080fd5b5060a15461034a9060ff1681565b34801561057657600080fd5b506103706205302081565b34801561058d57600080fd5b506102f561059c3660046140d5565b610c79565b3480156105ad57600080fd5b506102f56105bc3660046140d5565b610f47565b6102f56105cf3660046140ee565b610f58565b3480156105e057600080fd5b506103706105ef3660046140b8565b6001600160a01b0316600090815260a2602052604090205490565b34801561061657600080fd5b506102f5610f84565b34801561062b57600080fd5b506103706205573081565b34801561064257600080fd5b5061037060a05481565b6102f561065a366004614129565b610f96565b34801561066b57600080fd5b506068546001600160a01b03166103c8565b34801561068957600080fd5b50610314610fc2565b34801561069e57600080fd5b5061034a6106ad3660046140b8565b6001600160a01b031660009081526066602052604090205460ff1690565b3480156106d757600080fd5b5061037060d55481565b3480156106ed57600080fd5b5061034a6106fc36600461404b565b610fd1565b34801561070d57600080fd5b5061034a61071c36600461404b565b61104c565b34801561072d57600080fd5b50610370609e5481565b34801561074357600080fd5b506102f5610752366004614142565b61105a565b34801561076357600080fd5b506103706201c13881565b34801561077a57600080fd5b506103c8600160981b81565b34801561079257600080fd5b506102f56107a1366004614154565b6111b9565b3480156107b257600080fd5b5061037060d65481565b3480156107c857600080fd5b50610370609b5481565b3480156107de57600080fd5b50606754610370565b3480156107f357600080fd5b506103706108023660046141dc565b611365565b34801561081357600080fd5b50610370609c5481565b34801561082957600080fd5b506102f56108383660046140b8565b611390565b34801561084957600080fd5b506103706108583660046140d5565b611406565b34801561086957600080fd5b506102f5610878366004614345565b61141d565b34801561088957600080fd5b50609d5461048f90600160a01b900460ff1681565b3480156108aa57600080fd5b50609f5461034a9060ff1681565b60405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a26108f73334611628565b565b606060a5805461090890614408565b80601f016020809104026020016040519081016040528092919081815260200182805461093490614408565b80156109815780601f1061095657610100808354040283529160200191610981565b820191906000526020600020905b81548152906001019060200180831161096457829003601f168201915b5050505050905090565b6000336109998185856116ea565b60019150505b92915050565b6000806109c162010203600160981b0131600160981b31614452565b9050600060a05460d5546109d59190614452565b90506109e18282614465565b9250505090565b6000336109f685828561180f565b610a01858585611883565b506001949350505050565b610a14611a2e565b6001600160a01b038116610a435760405162461bcd60e51b8152600401610a3a90614478565b60405180910390fd5b6001600160a01b03811660009081526066602052604090205460ff1615610ac25760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608401610a3a565b6001600160a01b038116600081815260666020526040808220805460ff19166001179055517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b60405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a2610b4d3382611a36565b610b573382611b67565b50565b600033610999818585610b6d8383611365565b610b779190614452565b6116ea565b610b84611a2e565b6001600160a01b038116610baa5760405162461bcd60e51b8152600401610a3a90614478565b6001600160a01b03811660009081526066602052604090205460ff16610c245760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608401610a3a565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152606660205260409020805460ff19169055565b6001609a5414610c9b5760405162461bcd60e51b8152600401610a3a906144c6565b6002609a5560d654600160981b31908111610d2a5760405162461bcd60e51b815260206004820152604360248201527f4e6174697665546f6b656e53706f6b653a206275726e2061646472657373206260448201527f616c616e6365206e6f742067726561746572207468616e206c617374207265706064820152621bdc9d60ea1b608482015260a401610a3a565b600060d65482610d3a9190614465565b90506000606460d45483610d4e919061450a565b610d589190614521565b90506000610d668284614465565b60d685905590508115610d8857610d7d3083611c80565b610d8682611d02565b505b609e54609f54600091610d9e9160ff1684611d12565b11610e075760405162461bcd60e51b815260206004820152603360248201527f4e6174697665546f6b656e53706f6b653a207a65726f207363616c656420616d60448201527237bab73a103a37903932b837b93a10313ab93760691b6064820152608401610a3a565b604080518082018252600181528151808301835262010203600160981b0181526020808201859052925160009380840192610e4492909101614559565b60408051601f198184030181529181529152805160c081018252609c548152609d546001600160a01b0316602082810191909152825180840184523081529081018790529181019190915260608101889052909150600090610efd9060808101835b604051908082528060200260200182016040528015610ecf578160200160208202803683370190505b50815260200184604051602001610ee69190614579565b604051602081830303815290604052815250611d29565b9050807f0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a584604051610f3191815260200190565b60405180910390a250506001609a555050505050565b610f4f611a2e565b610b5781611e4f565b60a15460ff16610f7a5760405162461bcd60e51b8152600401610a3a906145be565b610b578134611fd9565b610f8c612035565b6108f7600061208f565b60a15460ff16610fb85760405162461bcd60e51b8152600401610a3a906145be565b610b5781346120e1565b606060a6805461090890614408565b60003381610fdf8286611365565b90508381101561103f5760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152608401610a3a565b610a0182868684036116ea565b600033610999818585611883565b60a154610100900460ff16156110b25760405162461bcd60e51b815260206004820152601e60248201527f546f6b656e53706f6b653a20616c7265616479207265676973746572656400006044820152606401610a3a565b604080516060808201835260a0548252609d5460ff600160a01b820481166020808601918252600160a81b9093048216858701908152865180880188526000808252885188518188015293518516848a015291519093168286015286518083039095018552608090910190955280820192909252919290916111449061113a908601866140b8565b856020013561212f565b6040805160c081018252609c548152609d546001600160a01b0316602080830191909152825180840184529394506111b293919283019190819061118a908a018a6140b8565b6001600160a01b0316815260209081018690529082526201c138908201526040016000610ea6565b5050505050565b6111c1612178565b6067546065546001600160a01b0316634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa15801561121a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061123e919061460c565b10156112a55760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608401610a3a565b6112ae336106ad565b156113145760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608401610a3a565b611355848484848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506121d192505050565b61135f6001603355565b50505050565b6001600160a01b03918216600090815260a36020908152604080832093909416825291909152205490565b611398612035565b6001600160a01b0381166113fd5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610a3a565b610b578161208f565b6000600561141583601f614452565b901c92915050565b600054610100900460ff161580801561143d5750600054600160ff909116105b806114575750303b158015611457575060005460ff166001145b6114ba5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610a3a565b6000805460ff1916600117905580156114dd576000805461ff0019166101001790555b611506846040516020016114f19190614625565b604051602081830303815290604052856123cb565b611512858460126123fc565b8260000361157b5760405162461bcd60e51b815260206004820152603060248201527f4e6174697665546f6b656e53706f6b653a207a65726f20696e697469616c207260448201526f65736572766520696d62616c616e636560801b6064820152608401610a3a565b606482106115d75760405162461bcd60e51b8152602060048201526024808201527f4e6174697665546f6b656e53706f6b653a20696e76616c69642070657263656e6044820152637461676560e01b6064820152608401610a3a565b60d482905580156111b2576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050505050565b6001600160a01b03821661167e5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610a3a565b8060a460008282546116909190614452565b90915550506001600160a01b038216600081815260a260209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35b5050565b6001600160a01b03831661174c5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610a3a565b6001600160a01b0382166117ad5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610a3a565b6001600160a01b03838116600081815260a3602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b600061181b8484611365565b9050600019811461135f57818110156118765760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610a3a565b61135f84848484036116ea565b6001600160a01b0383166118e75760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610a3a565b6001600160a01b0382166119495760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610a3a565b6001600160a01b038316600090815260a26020526040902054818110156119c15760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610a3a565b6001600160a01b03808516600081815260a2602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90611a219086815260200190565b60405180910390a361135f565b6108f7612035565b6001600160a01b038216611a965760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610a3a565b6001600160a01b038216600090815260a2602052604090205481811015611b0a5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610a3a565b6001600160a01b038316600081815260a260209081526040808320868603905560a480548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101611802565b505050565b80471015611bb75760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610a3a565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114611c04576040519150601f19603f3d011682016040523d82523d6000602084013e611c09565b606091505b5050905080611b625760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610a3a565b8060d56000828254611c929190614452565b90915550506040516327ad555d60e11b81526001600160a01b0383166004820152602481018290526001600160991b0190634f5aaaba90604401600060405180830381600087803b158015611ce657600080fd5b505af1158015611cfa573d6000803e3d6000fd5b505050505050565b6000611d0e3083611628565b5090565b6000611d218484846000612748565b949350505050565b600080611d34612779565b60408401516020015190915015611dd9576040830151516001600160a01b0316611db65760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608401610a3a565b604083015160208101519051611dd9916001600160a01b03909116908390612882565b604051630624488560e41b81526001600160a01b03821690636244885090611e05908690600401614655565b6020604051808303816000875af1158015611e24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e48919061460c565b9392505050565b6065546040805163301fd1f560e21b815290516000926001600160a01b03169163c07f47d49160048083019260209291908290030181865afa158015611e99573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ebd919061460c565b60675490915081831115611f2d5760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608401610a3a565b808311611fa25760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608401610a3a565b6067839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b6001609a5414611ffb5760405162461bcd60e51b8152600401610a3a906144c6565b6002609a5561200982612967565b609c548235036120225761201d8282612ba1565b61202c565b61201d8282612db9565b50506001609a55565b6068546001600160a01b031633146108f75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610a3a565b606880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001609a54146121035760405162461bcd60e51b8152600401610a3a906144c6565b6002609a5561211182613051565b609c548235036121255761201d828261313e565b61202c82826132a2565b6000816000036121415750600061099f565b306001600160a01b0384160361216e5761215c33308461180f565b612167333084611883565b508061099f565b611e48838361345d565b6002603354036121ca5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610a3a565b6002603355565b609c5483146122335760405162461bcd60e51b815260206004820152602860248201527f546f6b656e53706f6b653a20696e76616c696420736f7572636520626c6f636b60448201526718da185a5b88125160c21b6064820152608401610a3a565b609d546001600160a01b038381169116146122a25760405162461bcd60e51b815260206004820152602960248201527f546f6b656e53706f6b653a20696e76616c6964206f726967696e2073656e646560448201526872206164647265737360b81b6064820152608401610a3a565b6000818060200190518101906122b89190614753565b60a154909150610100900460ff1615806122d5575060a15460ff16155b156122ea5760a1805461ffff19166101011790555b6001815160048111156122ff576122ff614543565b03612338576000816020015180602001905181019061231e91906147e1565b9050612332816000015182602001516135bc565b5061135f565b60028151600481111561234d5761234d614543565b0361237c576000816020015180602001905181019061236c919061481b565b9050612332818260800151613609565b60405162461bcd60e51b815260206004820181905260248201527f546f6b656e53706f6b653a20696e76616c6964206d65737361676520747970656044820152606401610a3a565b6001603355565b600054610100900460ff166123f25760405162461bcd60e51b8152600401610a3a906148ea565b6116e68282613736565b600054610100900460ff166124235760405162461bcd60e51b8152600401610a3a906148ea565b61243583600001518460200151613776565b61243d6137b7565b6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015612482573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124a6919061460c565b609b55604083015161250b5760405162461bcd60e51b815260206004820152602860248201527f546f6b656e53706f6b653a207a65726f20746f6b656e2068756220626c6f636b60448201526718da185a5b88125160c21b6064820152608401610a3a565b609b548360400151036125865760405162461bcd60e51b815260206004820152603960248201527f546f6b656e53706f6b653a2063616e6e6f74206465706c6f7920746f2073616d60448201527f6520626c6f636b636861696e20617320746f6b656e20687562000000000000006064820152608401610a3a565b60608301516001600160a01b03166125eb5760405162461bcd60e51b815260206004820152602260248201527f546f6b656e53706f6b653a207a65726f20746f6b656e20687562206164647265604482015261737360f01b6064820152608401610a3a565b6012836080015160ff1611156126535760405162461bcd60e51b815260206004820152602760248201527f546f6b656e53706f6b653a20746f6b656e2068756220646563696d616c7320746044820152660dede40d0d2ced60cb1b6064820152608401610a3a565b60128160ff1611156126b35760405162461bcd60e51b815260206004820152602360248201527f546f6b656e53706f6b653a20746f6b656e20646563696d616c7320746f6f20686044820152620d2ced60eb1b6064820152608401610a3a565b6040830151609c556060830151609d805460a085905560a18054861560ff1990911617905560808601516001600160a01b039093166001600160a81b031990911617600160a01b60ff93841681029190911760ff60a81b1916600160a81b8585168102919091179283905561272f9391830482169204166137e5565b609f805460ff1916911515919091179055609e55505050565b6000811515841515036127665761275f858461450a565b9050611d21565b6127708584614521565b95945050505050565b600080606560009054906101000a90046001600160a01b03166001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156127cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127f39190614935565b9050612817816001600160a01b031660009081526066602052604090205460ff1690565b1561287d5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608401610a3a565b919050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa1580156128d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128f7919061460c565b6129019190614452565b6040516001600160a01b03851660248201526044810182905290915061135f90859063095ea7b360e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613829565b80356129855760405162461bcd60e51b8152600401610a3a90614952565b600061299760408301602084016140b8565b6001600160a01b0316036129bd5760405162461bcd60e51b8152600401610a3a9061499c565b60006129cf60608301604084016140b8565b6001600160a01b031603612a395760405162461bcd60e51b815260206004820152602b60248201527f546f6b656e53706f6b653a207a65726f20726563697069656e7420636f6e747260448201526a616374206164647265737360a81b6064820152608401610a3a565b6000816080013511612a5d5760405162461bcd60e51b8152600401610a3a906149e7565b60008160a0013511612abd5760405162461bcd60e51b8152602060048201526024808201527f546f6b656e53706f6b653a207a65726f20726563697069656e7420676173206c6044820152631a5b5a5d60e21b6064820152608401610a3a565b80608001358160a0013510612b245760405162461bcd60e51b815260206004820152602760248201527f546f6b656e53706f6b653a20696e76616c696420726563697069656e742067616044820152661cc81b1a5b5a5d60ca1b6064820152608401610a3a565b6000612b37610100830160e084016140b8565b6001600160a01b031603610b575760405162461bcd60e51b815260206004820152602b60248201527f546f6b656e53706f6b653a207a65726f2066616c6c6261636b2072656369706960448201526a656e74206164647265737360a81b6064820152608401610a3a565b612bcf612bb460408401602085016140b8565b610140840135612bca60e0860160c087016140b8565b6138fb565b6000612bf882612be7610120860161010087016140b8565b8561012001358661014001356139e8565b6040805180820190915291935091506000908060028152602001604051806101000160405280609b548152602001306001600160a01b03168152602001612c3c3390565b6001600160a01b03168152602001612c5a6060890160408a016140b8565b6001600160a01b0316815260208101879052604001612c7c6060890189614a2a565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060a08801356020820152604001612cd1610100890160e08a016140b8565b6001600160a01b03169052604051612cec9190602001614a70565b60408051601f198184030181529181529152805160c081018252609c548152609d546001600160a01b0316602082015281518083018352929350600092612d6c9282019080612d436101208b016101008c016140b8565b6001600160a01b0316815260209081018890529082526080890135908201526040016000610ea6565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b168787604051612daa929190614b7d565b60405180910390a35050505050565b612de38235612dce60408501602086016140b8565b612dde60e0860160c087016140b8565b613aa0565b6000612dfb82612be7610120860161010087016140b8565b6040805180820190915291935091506000908060048152602001604051806101600160405280612e283390565b6001600160a01b0316815260200187600001358152602001876020016020810190612e5391906140b8565b6001600160a01b03168152602001612e716060890160408a016140b8565b6001600160a01b0316815260208101879052604001612e936060890189614a2a565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060a08801356020820152604001612ee8610100890160e08a016140b8565b6001600160a01b0316815260808801356020820152604001612f1060e0890160c08a016140b8565b6001600160a01b03168152610140880135602091820152604051612f35929101614c8c565b60408051601f198184030181529190529052905060006105dc612f65612f5e6060880188614a2a565b9050611406565b612f6f919061450a565b612f7c9062055730614452565b6040805160c081018252609c548152609d546001600160a01b03166020820152815180830183529293506000926130039282019080612fc36101208c016101008d016140b8565b6001600160a01b031681526020908101899052908252818101869052604080516000815280830182528184015251606090920191610ee691889101614579565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b168888604051613041929190614b7d565b60405180910390a3505050505050565b600061306360608301604084016140b8565b6001600160a01b0316036130c45760405162461bcd60e51b815260206004820152602260248201527f546f6b656e53706f6b653a207a65726f20726563697069656e74206164647265604482015261737360f01b6064820152608401610a3a565b60008160c00135116130e85760405162461bcd60e51b8152600401610a3a906149e7565b80356131065760405162461bcd60e51b8152600401610a3a90614952565b600061311860408301602084016140b8565b6001600160a01b031603610b575760405162461bcd60e51b8152600401610a3a9061499c565b61316761315160408401602085016140b8565b60a0840135612bca610100860160e087016140b8565b600061318c8261317d60808601606087016140b8565b85608001358660a001356139e8565b604080518082019091529193509150600090806001815260200160405180604001604052808760400160208101906131c491906140b8565b6001600160a01b03168152602001868152506040516020016131e69190614559565b60408051601f198184030181529181529152805160c081018252609c548152609d546001600160a01b0316602082015281518083018352929350600092613264928201908061323b60808b0160608c016140b8565b6001600160a01b03168152602090810188905290825260c0890135908201526040016000610ea6565b9050336001600160a01b0316817f93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb528787604051612daa929190614d6a565b6132c882356132b760408501602086016140b8565b612dde610100860160e087016140b8565b60006132de8261317d60808601606087016140b8565b60408051808201825260038152815160e0810183528735815293955091935060009260208084019282820191613318918a01908a016140b8565b6001600160a01b031681526020016133366060890160408a016140b8565b6001600160a01b031681526020810187905260a0880135604082015260c08801356060820152608001613370610100890160e08a016140b8565b6001600160a01b031690526040516133e09190602001815181526020808301516001600160a01b0390811691830191909152604080840151821690830152606080840151908301526080808401519083015260a0808401519083015260c092830151169181019190915260e00190565b60408051601f198184030181529181529152805160c081018252609c548152609d546001600160a01b0316602082015281518083018352929350600092613264928201908061343560808b0160608c016140b8565b6001600160a01b03168152602090810188905290825262053020908201526040016000610ea6565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa1580156134a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134ca919061460c565b90506134e16001600160a01b038516333086613b33565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015613528573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061354c919061460c565b90508181116135b25760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610a3a565b6127708282614465565b816001600160a01b03167f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b826040516135f791815260200190565b60405180910390a26116e68282611c80565b6136133082611c80565b60008260000151836020015184604001518560a0015160405160240161363c9493929190614e09565b60408051601f198184030181529190526020810180516001600160e01b031663161b12ff60e11b17905260c084015160608501519192506000916136839190859085613b6b565b905080156136d75783606001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4846040516136ca91815260200190565b60405180910390a261135f565b83606001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb08460405161371691815260200190565b60405180910390a260e084015161135f906001600160a01b031684611b67565b600054610100900460ff1661375d5760405162461bcd60e51b8152600401610a3a906148ea565b60a56137698382614e8b565b5060a6611b628282614e8b565b600054610100900460ff1661379d5760405162461bcd60e51b8152600401610a3a906148ea565b6137a682613c40565b6137ae613d6e565b6116e681611390565b600054610100900460ff166137de5760405162461bcd60e51b8152600401610a3a906148ea565b6001609a55565b60008060ff808516908416118181613806576138018587614f4a565b613810565b6138108686614f4a565b61381b90600a615047565b9350909150505b9250929050565b600061387e826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613d9d9092919063ffffffff16565b805190915015611b62578080602001905181019061389c9190615056565b611b625760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610a3a565b609d546001600160a01b038481169116146139285760405162461bcd60e51b8152600401610a3a90615078565b81156139815760405162461bcd60e51b815260206004820152602260248201527f546f6b656e53706f6b653a206e6f6e2d7a65726f207365636f6e646172792066604482015261656560f01b6064820152608401610a3a565b6001600160a01b03811615611b625760405162461bcd60e51b815260206004820152602760248201527f546f6b656e53706f6b653a206e6f6e2d7a65726f206d756c74692d686f702066604482015266616c6c6261636b60c81b6064820152608401610a3a565b6000806139f486611d02565b9550613a00858561212f565b9350613a0b86613dac565b609e54609f54613a1f919060ff1685611d12565b609e54609f54613a33919060ff1689611d12565b11613a945760405162461bcd60e51b815260206004820152602b60248201527f546f6b656e53706f6b653a20696e73756666696369656e7420746f6b656e732060448201526a3a37903a3930b739b332b960a91b6064820152608401610a3a565b50939491935090915050565b609b548303613ad157306001600160a01b03831603613ad15760405162461bcd60e51b8152600401610a3a90615078565b6001600160a01b038116611b625760405162461bcd60e51b815260206004820152602360248201527f546f6b656e53706f6b653a207a65726f206d756c74692d686f702066616c6c6260448201526261636b60e81b6064820152608401610a3a565b6040516001600160a01b038085166024830152831660448201526064810182905261135f9085906323b872dd60e01b90608401612930565b6000845a1015613bbd5760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e742067617300000000006044820152606401610a3a565b83471015613c0d5760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c75650000006044820152606401610a3a565b826001600160a01b03163b600003613c2757506000611d21565b600080600084516020860188888bf19695505050505050565b600054610100900460ff16613c675760405162461bcd60e51b8152600401610a3a906148ea565b613c6f613dc9565b6001600160a01b038116613ceb5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f7274657220726567697374727920616464726573730000000000000000006064820152608401610a3a565b606580546001600160a01b0319166001600160a01b0383169081179091556040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa158015613d44573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d68919061460c565b60675550565b600054610100900460ff16613d955760405162461bcd60e51b8152600401610a3a906148ea565b6108f7613df8565b6060611d218484600085613e28565b613db63082611a36565b610b5762010203600160981b0182611b67565b600054610100900460ff16613df05760405162461bcd60e51b8152600401610a3a906148ea565b6108f7613f03565b600054610100900460ff16613e1f5760405162461bcd60e51b8152600401610a3a906148ea565b6108f73361208f565b606082471015613e895760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610a3a565b600080866001600160a01b03168587604051613ea591906150c6565b60006040518083038185875af1925050503d8060008114613ee2576040519150601f19603f3d011682016040523d82523d6000602084013e613ee7565b606091505b5091509150613ef887838387613f2a565b979650505050505050565b600054610100900460ff166123c45760405162461bcd60e51b8152600401610a3a906148ea565b60608315613f99578251600003613f92576001600160a01b0385163b613f925760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610a3a565b5081611d21565b611d218383815115613fae5781518083602001fd5b8060405162461bcd60e51b8152600401610a3a9190614018565b60005b83811015613fe3578181015183820152602001613fcb565b50506000910152565b60008151808452614004816020860160208601613fc8565b601f01601f19169290920160200192915050565b602081526000611e486020830184613fec565b6001600160a01b0381168114610b5757600080fd5b803561287d8161402b565b6000806040838503121561405e57600080fd5b82356140698161402b565b946020939093013593505050565b60008060006060848603121561408c57600080fd5b83356140978161402b565b925060208401356140a78161402b565b929592945050506040919091013590565b6000602082840312156140ca57600080fd5b8135611e488161402b565b6000602082840312156140e757600080fd5b5035919050565b60006020828403121561410057600080fd5b81356001600160401b0381111561411657600080fd5b82016101608185031215611e4857600080fd5b6000610100828403121561413c57600080fd5b50919050565b60006040828403121561413c57600080fd5b6000806000806060858703121561416a57600080fd5b84359350602085013561417c8161402b565b925060408501356001600160401b038082111561419857600080fd5b818701915087601f8301126141ac57600080fd5b8135818111156141bb57600080fd5b8860208285010111156141cd57600080fd5b95989497505060200194505050565b600080604083850312156141ef57600080fd5b82356141fa8161402b565b9150602083013561420a8161402b565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b038111828210171561424d5761424d614215565b60405290565b604080519081016001600160401b038111828210171561424d5761424d614215565b60405161010081016001600160401b038111828210171561424d5761424d614215565b604051601f8201601f191681016001600160401b03811182821017156142c0576142c0614215565b604052919050565b60006001600160401b038211156142e1576142e1614215565b50601f01601f191660200190565b600082601f83011261430057600080fd5b813561431361430e826142c8565b614298565b81815284602083860101111561432857600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008084860361010081121561435d57600080fd5b60a081121561436b57600080fd5b5061437461422b565b853561437f8161402b565b8152602086013561438f8161402b565b60208201526040868101359082015260608601356143ac8161402b565b6060820152608086013560ff811681146143c557600080fd5b6080820152935060a08501356001600160401b038111156143e557600080fd5b6143f1878288016142ef565b949794965050505060c08301359260e00135919050565b600181811c9082168061441c57607f821691505b60208210810361413c57634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b8082018082111561099f5761099f61443c565b8181038181111561099f5761099f61443c565b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b808202811582820484141761099f5761099f61443c565b60008261453e57634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052602160045260246000fd5b81516001600160a01b03168152602080830151908201526040810161099f565b60208152600082516005811061459f57634e487b7160e01b600052602160045260246000fd5b806020840152506020830151604080840152611d216060840182613fec565b6020808252602e908201527f4e6174697665546f6b656e53706f6b653a20636f6e747261637420756e64657260408201526d18dbdb1b185d195c985b1a5e995960921b606082015260800190565b60006020828403121561461e57600080fd5b5051919050565b6702bb930b83832b2160c51b815260008251614648816008850160208701613fc8565b9190910160080192915050565b6020808252825182820152828101516001600160a01b039081166040808501919091528401518051821660608501528083015160808501526000929161010085019190606087015160a0870152608087015160e060c0880152805193849052840192600092506101208701905b808410156146e4578451831682529385019360019390930192908501906146c2565b5060a0880151878203601f190160e089015294506147028186613fec565b98975050505050505050565b600082601f83011261471f57600080fd5b815161472d61430e826142c8565b81815284602083860101111561474257600080fd5b611d21826020830160208701613fc8565b60006020828403121561476557600080fd5b81516001600160401b038082111561477c57600080fd5b908301906040828603121561479057600080fd5b614798614253565b8251600581106147a757600080fd5b81526020830151828111156147bb57600080fd5b6147c78782860161470e565b60208301525095945050505050565b805161287d8161402b565b6000604082840312156147f357600080fd5b6147fb614253565b82516148068161402b565b81526020928301519281019290925250919050565b60006020828403121561482d57600080fd5b81516001600160401b038082111561484457600080fd5b90830190610100828603121561485957600080fd5b614861614275565b82518152614871602084016147d6565b6020820152614882604084016147d6565b6040820152614893606084016147d6565b60608201526080830151608082015260a0830151828111156148b457600080fd5b6148c08782860161470e565b60a08301525060c083015160c08201526148dc60e084016147d6565b60e082015295945050505050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b60006020828403121561494757600080fd5b8151611e488161402b565b6020808252602a908201527f546f6b656e53706f6b653a207a65726f2064657374696e6174696f6e20626c6f60408201526918dad8da185a5b88125160b21b606082015260800190565b6020808252602b908201527f546f6b656e53706f6b653a207a65726f2064657374696e6174696f6e2062726960408201526a646765206164647265737360a81b606082015260800190565b60208082526023908201527f546f6b656e53706f6b653a207a65726f20726571756972656420676173206c696040820152621b5a5d60ea1b606082015260800190565b6000808335601e19843603018112614a4157600080fd5b8301803591506001600160401b03821115614a5b57600080fd5b60200191503681900382131561382257600080fd5b60208152815160208201526000602083015160018060a01b03808216604085015280604086015116606085015250506060830151614ab960808401826001600160a01b03169052565b50608083015160a083015260a08301516101008060c0850152614ae0610120850183613fec565b915060c085015160e085015260e0850151614b05828601826001600160a01b03169052565b5090949350505050565b6000808335601e19843603018112614b2657600080fd5b83016020810192503590506001600160401b03811115614b4557600080fd5b80360382131561382257600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60408152823560408201526000614b9660208501614040565b6001600160a01b03166060830152614bb060408501614040565b6001600160a01b03166080830152614bcb6060850185614b0f565b6101608060a0860152614be36101a086018385614b54565b9250608087013560c086015260a087013560e0860152614c0560c08801614040565b9150610100614c1e818701846001600160a01b03169052565b614c2a60e08901614040565b9250610120614c43818801856001600160a01b03169052565b614c4e828a01614040565b93506101409150614c69828801856001600160a01b03169052565b880135918601919091529095013561018084015260209092019290925292915050565b60208152614ca66020820183516001600160a01b03169052565b6020820151604082015260006040830151614ccc60608401826001600160a01b03169052565b5060608301516001600160a01b038116608084015250608083015160a083015260a08301516101608060c0850152614d08610180850183613fec565b915060c085015160e085015260e0850151610100614d30818701836001600160a01b03169052565b860151610120868101919091528601519050610140614d59818701836001600160a01b03169052565b959095015193019290925250919050565b8235815261012081016020840135614d818161402b565b6001600160a01b039081166020840152604085013590614da08261402b565b166040830152614db260608501614040565b6001600160a01b0381166060840152506080840135608083015260a084013560a083015260c084013560c0830152614dec60e08501614040565b6001600160a01b031660e083015261010090910191909152919050565b8481526001600160a01b03848116602083015283166040820152608060608201819052600090614e3b90830184613fec565b9695505050505050565b601f821115611b6257600081815260208120601f850160051c81016020861015614e6c5750805b601f850160051c820191505b81811015611cfa57828155600101614e78565b81516001600160401b03811115614ea457614ea4614215565b614eb881614eb28454614408565b84614e45565b602080601f831160018114614eed5760008415614ed55750858301515b600019600386901b1c1916600185901b178555611cfa565b600085815260208120601f198616915b82811015614f1c57888601518255948401946001909101908401614efd565b5085821015614f3a5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60ff828116828216039081111561099f5761099f61443c565b600181815b80851115614f9e578160001904821115614f8457614f8461443c565b80851615614f9157918102915b93841c9390800290614f68565b509250929050565b600082614fb55750600161099f565b81614fc25750600061099f565b8160018114614fd85760028114614fe257614ffe565b600191505061099f565b60ff841115614ff357614ff361443c565b50506001821b61099f565b5060208310610133831016604e8410600b8410161715615021575081810a61099f565b61502b8383614f63565b806000190482111561503f5761503f61443c565b029392505050565b6000611e4860ff841683614fa6565b60006020828403121561506857600080fd5b81518015158114611e4857600080fd5b6020808252602e908201527f546f6b656e53706f6b653a20696e76616c69642064657374696e6174696f6e2060408201526d627269646765206164647265737360901b606082015260800190565b600082516150d8818460208701613fc8565b919091019291505056fea2646970667358221220551cb85c8323f71d68f868b5d25cad9d037328b9fad7b0ce5818a58e61000cd264736f6c63430008120033",
}

// NativeTokenSpokeABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenSpokeMetaData.ABI instead.
var NativeTokenSpokeABI = NativeTokenSpokeMetaData.ABI

// NativeTokenSpokeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenSpokeMetaData.Bin instead.
var NativeTokenSpokeBin = NativeTokenSpokeMetaData.Bin

// DeployNativeTokenSpoke deploys a new Ethereum contract, binding an instance of NativeTokenSpoke to it.
func DeployNativeTokenSpoke(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NativeTokenSpoke, error) {
	parsed, err := NativeTokenSpokeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenSpokeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenSpoke{NativeTokenSpokeCaller: NativeTokenSpokeCaller{contract: contract}, NativeTokenSpokeTransactor: NativeTokenSpokeTransactor{contract: contract}, NativeTokenSpokeFilterer: NativeTokenSpokeFilterer{contract: contract}}, nil
}

// NativeTokenSpoke is an auto generated Go binding around an Ethereum contract.
type NativeTokenSpoke struct {
	NativeTokenSpokeCaller     // Read-only binding to the contract
	NativeTokenSpokeTransactor // Write-only binding to the contract
	NativeTokenSpokeFilterer   // Log filterer for contract events
}

// NativeTokenSpokeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenSpokeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenSpokeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenSpokeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenSpokeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenSpokeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenSpokeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenSpokeSession struct {
	Contract     *NativeTokenSpoke // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NativeTokenSpokeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenSpokeCallerSession struct {
	Contract *NativeTokenSpokeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// NativeTokenSpokeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenSpokeTransactorSession struct {
	Contract     *NativeTokenSpokeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// NativeTokenSpokeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenSpokeRaw struct {
	Contract *NativeTokenSpoke // Generic contract binding to access the raw methods on
}

// NativeTokenSpokeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenSpokeCallerRaw struct {
	Contract *NativeTokenSpokeCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenSpokeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenSpokeTransactorRaw struct {
	Contract *NativeTokenSpokeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenSpoke creates a new instance of NativeTokenSpoke, bound to a specific deployed contract.
func NewNativeTokenSpoke(address common.Address, backend bind.ContractBackend) (*NativeTokenSpoke, error) {
	contract, err := bindNativeTokenSpoke(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpoke{NativeTokenSpokeCaller: NativeTokenSpokeCaller{contract: contract}, NativeTokenSpokeTransactor: NativeTokenSpokeTransactor{contract: contract}, NativeTokenSpokeFilterer: NativeTokenSpokeFilterer{contract: contract}}, nil
}

// NewNativeTokenSpokeCaller creates a new read-only instance of NativeTokenSpoke, bound to a specific deployed contract.
func NewNativeTokenSpokeCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenSpokeCaller, error) {
	contract, err := bindNativeTokenSpoke(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeCaller{contract: contract}, nil
}

// NewNativeTokenSpokeTransactor creates a new write-only instance of NativeTokenSpoke, bound to a specific deployed contract.
func NewNativeTokenSpokeTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenSpokeTransactor, error) {
	contract, err := bindNativeTokenSpoke(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeTransactor{contract: contract}, nil
}

// NewNativeTokenSpokeFilterer creates a new log filterer instance of NativeTokenSpoke, bound to a specific deployed contract.
func NewNativeTokenSpokeFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenSpokeFilterer, error) {
	contract, err := bindNativeTokenSpoke(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeFilterer{contract: contract}, nil
}

// bindNativeTokenSpoke binds a generic wrapper to an already deployed contract.
func bindNativeTokenSpoke(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenSpokeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenSpoke *NativeTokenSpokeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenSpoke.Contract.NativeTokenSpokeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenSpoke *NativeTokenSpokeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.NativeTokenSpokeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenSpoke *NativeTokenSpokeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.NativeTokenSpokeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenSpoke *NativeTokenSpokeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenSpoke.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenSpoke *NativeTokenSpokeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenSpoke *NativeTokenSpokeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.contract.Transact(opts, method, params...)
}

// BURNEDFORBRIDGEADDRESS is a free data retrieval call binding the contract method 0x47a9a22c.
//
// Solidity: function BURNED_FOR_BRIDGE_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) BURNEDFORBRIDGEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "BURNED_FOR_BRIDGE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDFORBRIDGEADDRESS is a free data retrieval call binding the contract method 0x47a9a22c.
//
// Solidity: function BURNED_FOR_BRIDGE_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeSession) BURNEDFORBRIDGEADDRESS() (common.Address, error) {
	return _NativeTokenSpoke.Contract.BURNEDFORBRIDGEADDRESS(&_NativeTokenSpoke.CallOpts)
}

// BURNEDFORBRIDGEADDRESS is a free data retrieval call binding the contract method 0x47a9a22c.
//
// Solidity: function BURNED_FOR_BRIDGE_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) BURNEDFORBRIDGEADDRESS() (common.Address, error) {
	return _NativeTokenSpoke.Contract.BURNEDFORBRIDGEADDRESS(&_NativeTokenSpoke.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) BURNEDTXFEESADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "BURNED_TX_FEES_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenSpoke.Contract.BURNEDTXFEESADDRESS(&_NativeTokenSpoke.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenSpoke.Contract.BURNEDTXFEESADDRESS(&_NativeTokenSpoke.CallOpts)
}

// HUBCHAINBURNADDRESS is a free data retrieval call binding the contract method 0xeae0544f.
//
// Solidity: function HUB_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) HUBCHAINBURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "HUB_CHAIN_BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HUBCHAINBURNADDRESS is a free data retrieval call binding the contract method 0xeae0544f.
//
// Solidity: function HUB_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeSession) HUBCHAINBURNADDRESS() (common.Address, error) {
	return _NativeTokenSpoke.Contract.HUBCHAINBURNADDRESS(&_NativeTokenSpoke.CallOpts)
}

// HUBCHAINBURNADDRESS is a free data retrieval call binding the contract method 0xeae0544f.
//
// Solidity: function HUB_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) HUBCHAINBURNADDRESS() (common.Address, error) {
	return _NativeTokenSpoke.Contract.HUBCHAINBURNADDRESS(&_NativeTokenSpoke.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) MULTIHOPCALLGASPERWORD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "MULTI_HOP_CALL_GAS_PER_WORD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.MULTIHOPCALLGASPERWORD(&_NativeTokenSpoke.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.MULTIHOPCALLGASPERWORD(&_NativeTokenSpoke.CallOpts)
}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) MULTIHOPCALLREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "MULTI_HOP_CALL_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) MULTIHOPCALLREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.MULTIHOPCALLREQUIREDGAS(&_NativeTokenSpoke.CallOpts)
}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) MULTIHOPCALLREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.MULTIHOPCALLREQUIREDGAS(&_NativeTokenSpoke.CallOpts)
}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) MULTIHOPSENDREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "MULTI_HOP_SEND_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) MULTIHOPSENDREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.MULTIHOPSENDREQUIREDGAS(&_NativeTokenSpoke.CallOpts)
}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) MULTIHOPSENDREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.MULTIHOPSENDREQUIREDGAS(&_NativeTokenSpoke.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) NATIVEMINTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "NATIVE_MINTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenSpoke.Contract.NATIVEMINTER(&_NativeTokenSpoke.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenSpoke.Contract.NATIVEMINTER(&_NativeTokenSpoke.CallOpts)
}

// REGISTERSPOKEREQUIREDGAS is a free data retrieval call binding the contract method 0xc03f4013.
//
// Solidity: function REGISTER_SPOKE_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) REGISTERSPOKEREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "REGISTER_SPOKE_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGISTERSPOKEREQUIREDGAS is a free data retrieval call binding the contract method 0xc03f4013.
//
// Solidity: function REGISTER_SPOKE_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) REGISTERSPOKEREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.REGISTERSPOKEREQUIREDGAS(&_NativeTokenSpoke.CallOpts)
}

// REGISTERSPOKEREQUIREDGAS is a free data retrieval call binding the contract method 0xc03f4013.
//
// Solidity: function REGISTER_SPOKE_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) REGISTERSPOKEREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.REGISTERSPOKEREQUIREDGAS(&_NativeTokenSpoke.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NativeTokenSpoke.Contract.Allowance(&_NativeTokenSpoke.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NativeTokenSpoke.Contract.Allowance(&_NativeTokenSpoke.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NativeTokenSpoke.Contract.BalanceOf(&_NativeTokenSpoke.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NativeTokenSpoke.Contract.BalanceOf(&_NativeTokenSpoke.CallOpts, account)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenSpoke *NativeTokenSpokeSession) BlockchainID() ([32]byte, error) {
	return _NativeTokenSpoke.Contract.BlockchainID(&_NativeTokenSpoke.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) BlockchainID() ([32]byte, error) {
	return _NativeTokenSpoke.Contract.BlockchainID(&_NativeTokenSpoke.CallOpts)
}

// BurnedFeesReportingRewardPercentage is a free data retrieval call binding the contract method 0x3a23dfe2.
//
// Solidity: function burnedFeesReportingRewardPercentage() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) BurnedFeesReportingRewardPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "burnedFeesReportingRewardPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnedFeesReportingRewardPercentage is a free data retrieval call binding the contract method 0x3a23dfe2.
//
// Solidity: function burnedFeesReportingRewardPercentage() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) BurnedFeesReportingRewardPercentage() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.BurnedFeesReportingRewardPercentage(&_NativeTokenSpoke.CallOpts)
}

// BurnedFeesReportingRewardPercentage is a free data retrieval call binding the contract method 0x3a23dfe2.
//
// Solidity: function burnedFeesReportingRewardPercentage() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) BurnedFeesReportingRewardPercentage() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.BurnedFeesReportingRewardPercentage(&_NativeTokenSpoke.CallOpts)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) CalculateNumWords(opts *bind.CallOpts, payloadSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "calculateNumWords", payloadSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _NativeTokenSpoke.Contract.CalculateNumWords(&_NativeTokenSpoke.CallOpts, payloadSize)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _NativeTokenSpoke.Contract.CalculateNumWords(&_NativeTokenSpoke.CallOpts, payloadSize)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeSession) Decimals() (uint8, error) {
	return _NativeTokenSpoke.Contract.Decimals(&_NativeTokenSpoke.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) Decimals() (uint8, error) {
	return _NativeTokenSpoke.Contract.Decimals(&_NativeTokenSpoke.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.GetMinTeleporterVersion(&_NativeTokenSpoke.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.GetMinTeleporterVersion(&_NativeTokenSpoke.CallOpts)
}

// HubTokenDecimals is a free data retrieval call binding the contract method 0xfdd7546f.
//
// Solidity: function hubTokenDecimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) HubTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "hubTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// HubTokenDecimals is a free data retrieval call binding the contract method 0xfdd7546f.
//
// Solidity: function hubTokenDecimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeSession) HubTokenDecimals() (uint8, error) {
	return _NativeTokenSpoke.Contract.HubTokenDecimals(&_NativeTokenSpoke.CallOpts)
}

// HubTokenDecimals is a free data retrieval call binding the contract method 0xfdd7546f.
//
// Solidity: function hubTokenDecimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) HubTokenDecimals() (uint8, error) {
	return _NativeTokenSpoke.Contract.HubTokenDecimals(&_NativeTokenSpoke.CallOpts)
}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) InitialReserveImbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "initialReserveImbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) InitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.InitialReserveImbalance(&_NativeTokenSpoke.CallOpts)
}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) InitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.InitialReserveImbalance(&_NativeTokenSpoke.CallOpts)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) IsCollateralized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "isCollateralized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) IsCollateralized() (bool, error) {
	return _NativeTokenSpoke.Contract.IsCollateralized(&_NativeTokenSpoke.CallOpts)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) IsCollateralized() (bool, error) {
	return _NativeTokenSpoke.Contract.IsCollateralized(&_NativeTokenSpoke.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0x22366844.
//
// Solidity: function isRegistered() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) IsRegistered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "isRegistered")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0x22366844.
//
// Solidity: function isRegistered() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) IsRegistered() (bool, error) {
	return _NativeTokenSpoke.Contract.IsRegistered(&_NativeTokenSpoke.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0x22366844.
//
// Solidity: function isRegistered() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) IsRegistered() (bool, error) {
	return _NativeTokenSpoke.Contract.IsRegistered(&_NativeTokenSpoke.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenSpoke.Contract.IsTeleporterAddressPaused(&_NativeTokenSpoke.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenSpoke.Contract.IsTeleporterAddressPaused(&_NativeTokenSpoke.CallOpts, teleporterAddress)
}

// LastestBurnedFeesReported is a free data retrieval call binding the contract method 0xd10a5b8c.
//
// Solidity: function lastestBurnedFeesReported() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) LastestBurnedFeesReported(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "lastestBurnedFeesReported")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastestBurnedFeesReported is a free data retrieval call binding the contract method 0xd10a5b8c.
//
// Solidity: function lastestBurnedFeesReported() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) LastestBurnedFeesReported() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.LastestBurnedFeesReported(&_NativeTokenSpoke.CallOpts)
}

// LastestBurnedFeesReported is a free data retrieval call binding the contract method 0xd10a5b8c.
//
// Solidity: function lastestBurnedFeesReported() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) LastestBurnedFeesReported() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.LastestBurnedFeesReported(&_NativeTokenSpoke.CallOpts)
}

// MultiplyOnSpoke is a free data retrieval call binding the contract method 0xfeea2b72.
//
// Solidity: function multiplyOnSpoke() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) MultiplyOnSpoke(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "multiplyOnSpoke")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MultiplyOnSpoke is a free data retrieval call binding the contract method 0xfeea2b72.
//
// Solidity: function multiplyOnSpoke() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) MultiplyOnSpoke() (bool, error) {
	return _NativeTokenSpoke.Contract.MultiplyOnSpoke(&_NativeTokenSpoke.CallOpts)
}

// MultiplyOnSpoke is a free data retrieval call binding the contract method 0xfeea2b72.
//
// Solidity: function multiplyOnSpoke() view returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) MultiplyOnSpoke() (bool, error) {
	return _NativeTokenSpoke.Contract.MultiplyOnSpoke(&_NativeTokenSpoke.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenSpoke *NativeTokenSpokeSession) Name() (string, error) {
	return _NativeTokenSpoke.Contract.Name(&_NativeTokenSpoke.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) Name() (string, error) {
	return _NativeTokenSpoke.Contract.Name(&_NativeTokenSpoke.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeSession) Owner() (common.Address, error) {
	return _NativeTokenSpoke.Contract.Owner(&_NativeTokenSpoke.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) Owner() (common.Address, error) {
	return _NativeTokenSpoke.Contract.Owner(&_NativeTokenSpoke.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenSpoke *NativeTokenSpokeSession) Symbol() (string, error) {
	return _NativeTokenSpoke.Contract.Symbol(&_NativeTokenSpoke.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) Symbol() (string, error) {
	return _NativeTokenSpoke.Contract.Symbol(&_NativeTokenSpoke.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) TeleporterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "teleporterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TeleporterRegistry() (common.Address, error) {
	return _NativeTokenSpoke.Contract.TeleporterRegistry(&_NativeTokenSpoke.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) TeleporterRegistry() (common.Address, error) {
	return _NativeTokenSpoke.Contract.TeleporterRegistry(&_NativeTokenSpoke.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) TokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "tokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TokenDecimals() (uint8, error) {
	return _NativeTokenSpoke.Contract.TokenDecimals(&_NativeTokenSpoke.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) TokenDecimals() (uint8, error) {
	return _NativeTokenSpoke.Contract.TokenDecimals(&_NativeTokenSpoke.CallOpts)
}

// TokenHubAddress is a free data retrieval call binding the contract method 0x24427629.
//
// Solidity: function tokenHubAddress() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) TokenHubAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "tokenHubAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenHubAddress is a free data retrieval call binding the contract method 0x24427629.
//
// Solidity: function tokenHubAddress() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TokenHubAddress() (common.Address, error) {
	return _NativeTokenSpoke.Contract.TokenHubAddress(&_NativeTokenSpoke.CallOpts)
}

// TokenHubAddress is a free data retrieval call binding the contract method 0x24427629.
//
// Solidity: function tokenHubAddress() view returns(address)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) TokenHubAddress() (common.Address, error) {
	return _NativeTokenSpoke.Contract.TokenHubAddress(&_NativeTokenSpoke.CallOpts)
}

// TokenHubBlockchainID is a free data retrieval call binding the contract method 0xeb3da909.
//
// Solidity: function tokenHubBlockchainID() view returns(bytes32)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) TokenHubBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "tokenHubBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenHubBlockchainID is a free data retrieval call binding the contract method 0xeb3da909.
//
// Solidity: function tokenHubBlockchainID() view returns(bytes32)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TokenHubBlockchainID() ([32]byte, error) {
	return _NativeTokenSpoke.Contract.TokenHubBlockchainID(&_NativeTokenSpoke.CallOpts)
}

// TokenHubBlockchainID is a free data retrieval call binding the contract method 0xeb3da909.
//
// Solidity: function tokenHubBlockchainID() view returns(bytes32)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) TokenHubBlockchainID() ([32]byte, error) {
	return _NativeTokenSpoke.Contract.TokenHubBlockchainID(&_NativeTokenSpoke.CallOpts)
}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) TokenMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "tokenMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TokenMultiplier() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.TokenMultiplier(&_NativeTokenSpoke.CallOpts)
}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) TokenMultiplier() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.TokenMultiplier(&_NativeTokenSpoke.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TotalMinted() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.TotalMinted(&_NativeTokenSpoke.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) TotalMinted() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.TotalMinted(&_NativeTokenSpoke.CallOpts)
}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) TotalNativeAssetSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "totalNativeAssetSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TotalNativeAssetSupply() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.TotalNativeAssetSupply(&_NativeTokenSpoke.CallOpts)
}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) TotalNativeAssetSupply() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.TotalNativeAssetSupply(&_NativeTokenSpoke.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.TotalSupply(&_NativeTokenSpoke.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.TotalSupply(&_NativeTokenSpoke.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Approve(&_NativeTokenSpoke.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Approve(&_NativeTokenSpoke.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.DecreaseAllowance(&_NativeTokenSpoke.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.DecreaseAllowance(&_NativeTokenSpoke.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) Deposit() (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Deposit(&_NativeTokenSpoke.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) Deposit() (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Deposit(&_NativeTokenSpoke.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.IncreaseAllowance(&_NativeTokenSpoke.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.IncreaseAllowance(&_NativeTokenSpoke.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xf56c363d.
//
// Solidity: function initialize((address,address,bytes32,address,uint8) settings, string nativeAssetSymbol, uint256 initialReserveImbalance, uint256 burnedFeesReportingRewardPercentage_) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) Initialize(opts *bind.TransactOpts, settings TokenSpokeSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage_ *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "initialize", settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf56c363d.
//
// Solidity: function initialize((address,address,bytes32,address,uint8) settings, string nativeAssetSymbol, uint256 initialReserveImbalance, uint256 burnedFeesReportingRewardPercentage_) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) Initialize(settings TokenSpokeSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage_ *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Initialize(&_NativeTokenSpoke.TransactOpts, settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf56c363d.
//
// Solidity: function initialize((address,address,bytes32,address,uint8) settings, string nativeAssetSymbol, uint256 initialReserveImbalance, uint256 burnedFeesReportingRewardPercentage_) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) Initialize(settings TokenSpokeSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, burnedFeesReportingRewardPercentage_ *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Initialize(&_NativeTokenSpoke.TransactOpts, settings, nativeAssetSymbol, initialReserveImbalance, burnedFeesReportingRewardPercentage_)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.PauseTeleporterAddress(&_NativeTokenSpoke.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.PauseTeleporterAddress(&_NativeTokenSpoke.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.ReceiveTeleporterMessage(&_NativeTokenSpoke.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.ReceiveTeleporterMessage(&_NativeTokenSpoke.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RegisterWithHub is a paid mutator transaction binding the contract method 0xbf48d36b.
//
// Solidity: function registerWithHub((address,uint256) feeInfo) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) RegisterWithHub(opts *bind.TransactOpts, feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "registerWithHub", feeInfo)
}

// RegisterWithHub is a paid mutator transaction binding the contract method 0xbf48d36b.
//
// Solidity: function registerWithHub((address,uint256) feeInfo) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) RegisterWithHub(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.RegisterWithHub(&_NativeTokenSpoke.TransactOpts, feeInfo)
}

// RegisterWithHub is a paid mutator transaction binding the contract method 0xbf48d36b.
//
// Solidity: function registerWithHub((address,uint256) feeInfo) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) RegisterWithHub(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.RegisterWithHub(&_NativeTokenSpoke.TransactOpts, feeInfo)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.RenounceOwnership(&_NativeTokenSpoke.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.RenounceOwnership(&_NativeTokenSpoke.TransactOpts)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) ReportBurnedTxFees(opts *bind.TransactOpts, requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "reportBurnedTxFees", requiredGasLimit)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) ReportBurnedTxFees(requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.ReportBurnedTxFees(&_NativeTokenSpoke.TransactOpts, requiredGasLimit)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) ReportBurnedTxFees(requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.ReportBurnedTxFees(&_NativeTokenSpoke.TransactOpts, requiredGasLimit)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) Send(opts *bind.TransactOpts, input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Send(&_NativeTokenSpoke.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Send(&_NativeTokenSpoke.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) SendAndCall(opts *bind.TransactOpts, input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "sendAndCall", input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.SendAndCall(&_NativeTokenSpoke.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.SendAndCall(&_NativeTokenSpoke.TransactOpts, input)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Transfer(&_NativeTokenSpoke.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Transfer(&_NativeTokenSpoke.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.TransferFrom(&_NativeTokenSpoke.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.TransferFrom(&_NativeTokenSpoke.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.TransferOwnership(&_NativeTokenSpoke.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.TransferOwnership(&_NativeTokenSpoke.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.UnpauseTeleporterAddress(&_NativeTokenSpoke.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.UnpauseTeleporterAddress(&_NativeTokenSpoke.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.UpdateMinTeleporterVersion(&_NativeTokenSpoke.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.UpdateMinTeleporterVersion(&_NativeTokenSpoke.TransactOpts, version)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Withdraw(&_NativeTokenSpoke.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Withdraw(&_NativeTokenSpoke.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Fallback(&_NativeTokenSpoke.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Fallback(&_NativeTokenSpoke.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSpoke.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeSession) Receive() (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Receive(&_NativeTokenSpoke.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenSpoke *NativeTokenSpokeTransactorSession) Receive() (*types.Transaction, error) {
	return _NativeTokenSpoke.Contract.Receive(&_NativeTokenSpoke.TransactOpts)
}

// NativeTokenSpokeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeApprovalIterator struct {
	Event *NativeTokenSpokeApproval // Event containing the contract specifics and raw log

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
func (it *NativeTokenSpokeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeApproval)
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
		it.Event = new(NativeTokenSpokeApproval)
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
func (it *NativeTokenSpokeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeApproval represents a Approval event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NativeTokenSpokeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeApprovalIterator{contract: _NativeTokenSpoke.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeApproval)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseApproval(log types.Log) (*NativeTokenSpokeApproval, error) {
	event := new(NativeTokenSpokeApproval)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeCallFailedIterator struct {
	Event *NativeTokenSpokeCallFailed // Event containing the contract specifics and raw log

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
func (it *NativeTokenSpokeCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeCallFailed)
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
		it.Event = new(NativeTokenSpokeCallFailed)
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
func (it *NativeTokenSpokeCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeCallFailed represents a CallFailed event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeCallFailed struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterCallFailed(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenSpokeCallFailedIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeCallFailedIterator{contract: _NativeTokenSpoke.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeCallFailed, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeCallFailed)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "CallFailed", log); err != nil {
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
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseCallFailed(log types.Log) (*NativeTokenSpokeCallFailed, error) {
	event := new(NativeTokenSpokeCallFailed)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeCallSucceededIterator struct {
	Event *NativeTokenSpokeCallSucceeded // Event containing the contract specifics and raw log

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
func (it *NativeTokenSpokeCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeCallSucceeded)
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
		it.Event = new(NativeTokenSpokeCallSucceeded)
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
func (it *NativeTokenSpokeCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeCallSucceeded represents a CallSucceeded event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeCallSucceeded struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterCallSucceeded(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenSpokeCallSucceededIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeCallSucceededIterator{contract: _NativeTokenSpoke.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeCallSucceeded, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeCallSucceeded)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseCallSucceeded(log types.Log) (*NativeTokenSpokeCallSucceeded, error) {
	event := new(NativeTokenSpokeCallSucceeded)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeDepositIterator struct {
	Event *NativeTokenSpokeDeposit // Event containing the contract specifics and raw log

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
func (it *NativeTokenSpokeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeDeposit)
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
		it.Event = new(NativeTokenSpokeDeposit)
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
func (it *NativeTokenSpokeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeDeposit represents a Deposit event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenSpokeDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeDepositIterator{contract: _NativeTokenSpoke.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeDeposit)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseDeposit(log types.Log) (*NativeTokenSpokeDeposit, error) {
	event := new(NativeTokenSpokeDeposit)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeInitializedIterator struct {
	Event *NativeTokenSpokeInitialized // Event containing the contract specifics and raw log

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
func (it *NativeTokenSpokeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeInitialized)
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
		it.Event = new(NativeTokenSpokeInitialized)
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
func (it *NativeTokenSpokeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeInitialized represents a Initialized event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterInitialized(opts *bind.FilterOpts) (*NativeTokenSpokeInitializedIterator, error) {

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeInitializedIterator{contract: _NativeTokenSpoke.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeInitialized) (event.Subscription, error) {

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeInitialized)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseInitialized(log types.Log) (*NativeTokenSpokeInitialized, error) {
	event := new(NativeTokenSpokeInitialized)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeMinTeleporterVersionUpdatedIterator struct {
	Event *NativeTokenSpokeMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *NativeTokenSpokeMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeMinTeleporterVersionUpdated)
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
		it.Event = new(NativeTokenSpokeMinTeleporterVersionUpdated)
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
func (it *NativeTokenSpokeMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*NativeTokenSpokeMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeMinTeleporterVersionUpdatedIterator{contract: _NativeTokenSpoke.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeMinTeleporterVersionUpdated)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*NativeTokenSpokeMinTeleporterVersionUpdated, error) {
	event := new(NativeTokenSpokeMinTeleporterVersionUpdated)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeOwnershipTransferredIterator struct {
	Event *NativeTokenSpokeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NativeTokenSpokeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeOwnershipTransferred)
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
		it.Event = new(NativeTokenSpokeOwnershipTransferred)
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
func (it *NativeTokenSpokeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeOwnershipTransferred represents a OwnershipTransferred event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NativeTokenSpokeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeOwnershipTransferredIterator{contract: _NativeTokenSpoke.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeOwnershipTransferred)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseOwnershipTransferred(log types.Log) (*NativeTokenSpokeOwnershipTransferred, error) {
	event := new(NativeTokenSpokeOwnershipTransferred)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeReportBurnedTxFeesIterator is returned from FilterReportBurnedTxFees and is used to iterate over the raw logs and unpacked data for ReportBurnedTxFees events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeReportBurnedTxFeesIterator struct {
	Event *NativeTokenSpokeReportBurnedTxFees // Event containing the contract specifics and raw log

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
func (it *NativeTokenSpokeReportBurnedTxFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeReportBurnedTxFees)
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
		it.Event = new(NativeTokenSpokeReportBurnedTxFees)
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
func (it *NativeTokenSpokeReportBurnedTxFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeReportBurnedTxFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeReportBurnedTxFees represents a ReportBurnedTxFees event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeReportBurnedTxFees struct {
	TeleporterMessageID [32]byte
	FeesBurned          *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReportBurnedTxFees is a free log retrieval operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterReportBurnedTxFees(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*NativeTokenSpokeReportBurnedTxFeesIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "ReportBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeReportBurnedTxFeesIterator{contract: _NativeTokenSpoke.contract, event: "ReportBurnedTxFees", logs: logs, sub: sub}, nil
}

// WatchReportBurnedTxFees is a free log subscription operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchReportBurnedTxFees(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeReportBurnedTxFees, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "ReportBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeReportBurnedTxFees)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "ReportBurnedTxFees", log); err != nil {
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
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseReportBurnedTxFees(log types.Log) (*NativeTokenSpokeReportBurnedTxFees, error) {
	event := new(NativeTokenSpokeReportBurnedTxFees)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "ReportBurnedTxFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTeleporterAddressPausedIterator struct {
	Event *NativeTokenSpokeTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenSpokeTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeTeleporterAddressPaused)
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
		it.Event = new(NativeTokenSpokeTeleporterAddressPaused)
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
func (it *NativeTokenSpokeTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenSpokeTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeTeleporterAddressPausedIterator{contract: _NativeTokenSpoke.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeTeleporterAddressPaused)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseTeleporterAddressPaused(log types.Log) (*NativeTokenSpokeTeleporterAddressPaused, error) {
	event := new(NativeTokenSpokeTeleporterAddressPaused)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTeleporterAddressUnpausedIterator struct {
	Event *NativeTokenSpokeTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenSpokeTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeTeleporterAddressUnpaused)
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
		it.Event = new(NativeTokenSpokeTeleporterAddressUnpaused)
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
func (it *NativeTokenSpokeTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenSpokeTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeTeleporterAddressUnpausedIterator{contract: _NativeTokenSpoke.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeTeleporterAddressUnpaused)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*NativeTokenSpokeTeleporterAddressUnpaused, error) {
	event := new(NativeTokenSpokeTeleporterAddressUnpaused)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeTokensAndCallSentIterator is returned from FilterTokensAndCallSent and is used to iterate over the raw logs and unpacked data for TokensAndCallSent events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTokensAndCallSentIterator struct {
	Event *NativeTokenSpokeTokensAndCallSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenSpokeTokensAndCallSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeTokensAndCallSent)
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
		it.Event = new(NativeTokenSpokeTokensAndCallSent)
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
func (it *NativeTokenSpokeTokensAndCallSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeTokensAndCallSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeTokensAndCallSent represents a TokensAndCallSent event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTokensAndCallSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterTokensAndCallSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenSpokeTokensAndCallSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeTokensAndCallSentIterator{contract: _NativeTokenSpoke.contract, event: "TokensAndCallSent", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchTokensAndCallSent(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeTokensAndCallSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeTokensAndCallSent)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
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
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseTokensAndCallSent(log types.Log) (*NativeTokenSpokeTokensAndCallSent, error) {
	event := new(NativeTokenSpokeTokensAndCallSent)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeTokensSentIterator is returned from FilterTokensSent and is used to iterate over the raw logs and unpacked data for TokensSent events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTokensSentIterator struct {
	Event *NativeTokenSpokeTokensSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenSpokeTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeTokensSent)
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
		it.Event = new(NativeTokenSpokeTokensSent)
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
func (it *NativeTokenSpokeTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeTokensSent represents a TokensSent event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensSent is a free log retrieval operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenSpokeTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeTokensSentIterator{contract: _NativeTokenSpoke.contract, event: "TokensSent", logs: logs, sub: sub}, nil
}

// WatchTokensSent is a free log subscription operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchTokensSent(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeTokensSent)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "TokensSent", log); err != nil {
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
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseTokensSent(log types.Log) (*NativeTokenSpokeTokensSent, error) {
	event := new(NativeTokenSpokeTokensSent)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "TokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTokensWithdrawnIterator struct {
	Event *NativeTokenSpokeTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *NativeTokenSpokeTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeTokensWithdrawn)
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
		it.Event = new(NativeTokenSpokeTokensWithdrawn)
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
func (it *NativeTokenSpokeTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeTokensWithdrawn represents a TokensWithdrawn event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTokensWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*NativeTokenSpokeTokensWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeTokensWithdrawnIterator{contract: _NativeTokenSpoke.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeTokensWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeTokensWithdrawn)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseTokensWithdrawn(log types.Log) (*NativeTokenSpokeTokensWithdrawn, error) {
	event := new(NativeTokenSpokeTokensWithdrawn)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTransferIterator struct {
	Event *NativeTokenSpokeTransfer // Event containing the contract specifics and raw log

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
func (it *NativeTokenSpokeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeTransfer)
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
		it.Event = new(NativeTokenSpokeTransfer)
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
func (it *NativeTokenSpokeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeTransfer represents a Transfer event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTokenSpokeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeTransferIterator{contract: _NativeTokenSpoke.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeTransfer)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseTransfer(log types.Log) (*NativeTokenSpokeTransfer, error) {
	event := new(NativeTokenSpokeTransfer)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSpokeWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the NativeTokenSpoke contract.
type NativeTokenSpokeWithdrawalIterator struct {
	Event *NativeTokenSpokeWithdrawal // Event containing the contract specifics and raw log

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
func (it *NativeTokenSpokeWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSpokeWithdrawal)
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
		it.Event = new(NativeTokenSpokeWithdrawal)
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
func (it *NativeTokenSpokeWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSpokeWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSpokeWithdrawal represents a Withdrawal event raised by the NativeTokenSpoke contract.
type NativeTokenSpokeWithdrawal struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) FilterWithdrawal(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenSpokeWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.FilterLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSpokeWithdrawalIterator{contract: _NativeTokenSpoke.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *NativeTokenSpokeWithdrawal, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSpoke.contract.WatchLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSpokeWithdrawal)
				if err := _NativeTokenSpoke.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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
func (_NativeTokenSpoke *NativeTokenSpokeFilterer) ParseWithdrawal(log types.Log) (*NativeTokenSpokeWithdrawal, error) {
	event := new(NativeTokenSpokeWithdrawal)
	if err := _NativeTokenSpoke.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
