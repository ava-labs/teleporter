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
}

// NativeTokenSpokeMetaData contains all meta data concerning the NativeTokenSpoke contract.
var NativeTokenSpokeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tokenHubBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenHubAddress\",\"type\":\"address\"}],\"internalType\":\"structTokenSpokeSettings\",\"name\":\"settings\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"nativeAssetSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialReserveImbalance\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimalsShift\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"multiplyOnSpoke\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"burnedFeesReportingRewardPercentage_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesBurned\",\"type\":\"uint256\"}],\"name\":\"ReportBurnedTxFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"BURNED_FOR_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURNED_TX_FEES_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HUB_CHAIN_BURN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_CALL_GAS_PER_WORD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_MINTER\",\"outputs\":[{\"internalType\":\"contractINativeMinter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTER_SPOKE_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnedFeesReportingRewardPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"payloadSize\",\"type\":\"uint256\"}],\"name\":\"calculateNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialReserveImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastestBurnedFeesReported\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplyOnSpoke\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"registerWithHub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"name\":\"reportBurnedTxFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"sendAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenHubAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenHubBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalNativeAssetSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101806040523480156200001257600080fd5b50604051620054fb380380620054fb833981016040819052620000359162000753565b846040516020016200004891906200082d565b60408051601f198184030181529190528651602088015160016000558791899188918891889190816001600160a01b038116620000f25760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084015b60405180910390fd5b6001600160a01b03811660808190526040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa1580156200013d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200016391906200085f565b600255506200017233620004ec565b6200017d816200053e565b505060016004819055507302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200020091906200085f565b60a0526040840151620002675760405162461bcd60e51b815260206004820152602860248201527f546f6b656e53706f6b653a207a65726f20746f6b656e2068756220626c6f636b60448201526718da185a5b88125160c21b6064820152608401620000e9565b60a051846040015103620002e45760405162461bcd60e51b815260206004820152603960248201527f546f6b656e53706f6b653a2063616e6e6f74206465706c6f7920746f2073616d60448201527f6520626c6f636b636861696e20617320746f6b656e20687562000000000000006064820152608401620000e9565b60608401516001600160a01b03166200034b5760405162461bcd60e51b815260206004820152602260248201527f546f6b656e53706f6b653a207a65726f20746f6b656e20687562206164647265604482015261737360f01b6064820152608401620000e9565b60128260ff161115620003ab5760405162461bcd60e51b815260206004820152602160248201527f546f6b656e53706f6b653a20696e76616c696420646563696d616c73536869666044820152601d60fa1b6064820152608401620000e9565b604084015160c05260608401516001600160a01b031660e0526101408390526005805460ff19168415179055620003e482600a6200098e565b610100521515610120525060099150620004019050838262000a35565b50600a62000410828262000a35565b505050836000036200047e5760405162461bcd60e51b815260206004820152603060248201527f4e6174697665546f6b656e53706f6b653a207a65726f20696e697469616c207260448201526f65736572766520696d62616c616e636560801b6064820152608401620000e9565b60648110620004dc5760405162461bcd60e51b8152602060048201526024808201527f4e6174697665546f6b656e53706f6b653a20696e76616c69642070657263656e6044820152637461676560e01b6064820152608401620000e9565b610160525062000b019350505050565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b62000548620005bd565b6001600160a01b038116620005af5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401620000e9565b620005ba81620004ec565b50565b6003546001600160a01b03163314620006195760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401620000e9565b565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b03811182821017156200065657620006566200061b565b60405290565b80516001600160a01b03811681146200067457600080fd5b919050565b60005b83811015620006965781810151838201526020016200067c565b50506000910152565b600082601f830112620006b157600080fd5b81516001600160401b0380821115620006ce57620006ce6200061b565b604051601f8301601f19908116603f01168101908282118183101715620006f957620006f96200061b565b816040528381528660208588010111156200071357600080fd5b6200072684602083016020890162000679565b9695505050505050565b805160ff811681146200067457600080fd5b805180151581146200067457600080fd5b6000806000806000808688036101208112156200076f57600080fd5b60808112156200077e57600080fd5b506200078962000631565b62000794886200065c565b8152620007a4602089016200065c565b602082015260408801516040820152620007c1606089016200065c565b606082015260808801519096506001600160401b03811115620007e357600080fd5b620007f189828a016200069f565b95505060a087015193506200080960c0880162000730565b92506200081960e0880162000742565b915061010087015190509295509295509295565b6702bb930b83832b2160c51b8152600082516200085281600885016020870162000679565b9190910160080192915050565b6000602082840312156200087257600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600181815b80851115620008d0578160001904821115620008b457620008b462000879565b80851615620008c257918102915b93841c939080029062000894565b509250929050565b600082620008e95750600162000988565b81620008f85750600062000988565b81600181146200091157600281146200091c576200093c565b600191505062000988565b60ff84111562000930576200093062000879565b50506001821b62000988565b5060208310610133831016604e8410600b841016171562000961575081810a62000988565b6200096d83836200088f565b806000190482111562000984576200098462000879565b0290505b92915050565b60006200099f60ff841683620008d8565b9392505050565b600181811c90821680620009bb57607f821691505b602082108103620009dc57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111562000a3057600081815260208120601f850160051c8101602086101562000a0b5750805b601f850160051c820191505b8181101562000a2c5782815560010162000a17565b5050505b505050565b81516001600160401b0381111562000a515762000a516200061b565b62000a698162000a628454620009a6565b84620009e2565b602080601f83116001811462000aa1576000841562000a885750858301515b600019600386901b1c1916600185901b17855562000a2c565b600085815260208120601f198616915b8281101562000ad25788860151825594840194600190910190840162000ab1565b508582101562000af15787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c05160e051610100516101205161014051610160516148ad62000c4e600039600081816104b40152610d630152600081816105f4015281816109c4015261118a01526000818161089301528181610dea015281816111d6015281816136f3015261373e0152600081816106fd01528181610dc9015281816111ad015281816136d2015261371d0152600081816103e901528181610f010152818161128d015281816121f9015281816129c501528181612c8401528181612f3f0152818161317b015261359b01526000818161081f01528181610edb0152818161126701528181611f82015281816120a8015281816121790152818161299f01528181612c5e01528181612f1901526131550152600081816107b60152818161288f01526137d201526000818161035e0152818161131c01528181611db401526123e201526148ad6000f3fe60806040526004361061028c5760003560e01c8063715018a61161015a578063c452165e116100c1578063dd62ed3e1161007a578063dd62ed3e146107ed578063eae0544f146104f6578063eb3da9091461080d578063f2fde38b14610841578063f3f981d814610861578063feea2b72146108815761029b565b8063c452165e14610756578063c868efaa1461076e578063d0e30db01461029b578063d10a5b8c1461078e578063d127dc9b146107a4578063d2cc7a70146107d85761029b565b8063a2309ff811610113578063a2309ff814610695578063a457c2d7146106ab578063a9059cbb146106cb578063ba3f5a12146106eb578063bf48d36b1461071f578063c03f40131461073f5761029b565b8063715018a6146105cd5780638ac7dd20146105e25780638bf2fa94146106165780638da5cb5b1461062957806395d89b4114610647578063973142971461065c5761029b565b8063313ce567116101fe57806349e3284e116101b757806349e3284e14610513578063525975e61461052d57806355538c8b146105445780635eb99514146105645780636e6eef8d1461058457806370a08231146105975761029b565b8063313ce5671461044b578063329c3e121461046757806339509351146104825780633a23dfe2146104a25780634511243e146104d657806347a9a22c146104f65761029b565b80631a7f5bec116102505780631a7f5bec1461034c578063223668441461039857806323b872dd146103b757806324427629146103d75780632b0d8f181461040b5780632e1a7d4d1461042b5761029b565b806306fdde03146102a3578063095ea7b3146102ce57806315beb59f146102fe57806318160ddd146103225780631906529c146103375761029b565b3661029b576102996108b5565b005b6102996108b5565b3480156102af57600080fd5b506102b86108f6565b6040516102c59190613b83565b60405180910390f35b3480156102da57600080fd5b506102ee6102e9366004613bb6565b610988565b60405190151581526020016102c5565b34801561030a57600080fd5b5061031461213481565b6040519081526020016102c5565b34801561032e57600080fd5b50600854610314565b34801561034357600080fd5b506103146109a2565b34801561035857600080fd5b506103807f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016102c5565b3480156103a457600080fd5b506005546102ee90610100900460ff1681565b3480156103c357600080fd5b506102ee6103d2366004613be2565b610a03565b3480156103e357600080fd5b506103807f000000000000000000000000000000000000000000000000000000000000000081565b34801561041757600080fd5b50610299610426366004613c23565b610a27565b34801561043757600080fd5b50610299610446366004613c40565b610b2c565b34801561045757600080fd5b50604051601281526020016102c5565b34801561047357600080fd5b506103806001600160991b0181565b34801561048e57600080fd5b506102ee61049d366004613bb6565b610b78565b3480156104ae57600080fd5b506103147f000000000000000000000000000000000000000000000000000000000000000081565b3480156104e257600080fd5b506102996104f1366004613c23565b610b9a565b34801561050257600080fd5b5061038062010203600160981b0181565b34801561051f57600080fd5b506005546102ee9060ff1681565b34801561053957600080fd5b506103146203a98081565b34801561055057600080fd5b5061029961055f366004613c40565b610c97565b34801561057057600080fd5b5061029961057f366004613c40565b611012565b610299610592366004613c59565b611023565b3480156105a357600080fd5b506103146105b2366004613c23565b6001600160a01b031660009081526006602052604090205490565b3480156105d957600080fd5b5061029961104f565b3480156105ee57600080fd5b506103147f000000000000000000000000000000000000000000000000000000000000000081565b610299610624366004613c95565b611061565b34801561063557600080fd5b506003546001600160a01b0316610380565b34801561065357600080fd5b506102b861108d565b34801561066857600080fd5b506102ee610677366004613c23565b6001600160a01b031660009081526001602052604090205460ff1690565b3480156106a157600080fd5b50610314600b5481565b3480156106b757600080fd5b506102ee6106c6366004613bb6565b61109c565b3480156106d757600080fd5b506102ee6106e6366004613bb6565b611117565b3480156106f757600080fd5b506103147f000000000000000000000000000000000000000000000000000000000000000081565b34801561072b57600080fd5b5061029961073a366004613cae565b611125565b34801561074b57600080fd5b50610314620249f081565b34801561076257600080fd5b50610380600160981b81565b34801561077a57600080fd5b50610299610789366004613cc0565b611307565b34801561079a57600080fd5b50610314600c5481565b3480156107b057600080fd5b506103147f000000000000000000000000000000000000000000000000000000000000000081565b3480156107e457600080fd5b50600254610314565b3480156107f957600080fd5b50610314610808366004613d49565b6114d1565b34801561081957600080fd5b506103147f000000000000000000000000000000000000000000000000000000000000000081565b34801561084d57600080fd5b5061029961085c366004613c23565b6114fc565b34801561086d57600080fd5b5061031461087c366004613c40565b611572565b34801561088d57600080fd5b506102ee7f000000000000000000000000000000000000000000000000000000000000000081565b60405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a26108f43334611589565b565b60606009805461090590613d82565b80601f016020809104026020016040519081016040528092919081815260200182805461093190613d82565b801561097e5780601f106109535761010080835404028352916020019161097e565b820191906000526020600020905b81548152906001019060200180831161096157829003601f168201915b5050505050905090565b60003361099681858561164b565b60019150505b92915050565b6000806109be62010203600160981b0131600160981b31613dcc565b905060007f0000000000000000000000000000000000000000000000000000000000000000600b546109f09190613dcc565b90506109fc8282613ddf565b9250505090565b600033610a11858285611770565b610a1c8585856117e4565b506001949350505050565b610a2f61198f565b6001600160a01b038116610a5e5760405162461bcd60e51b8152600401610a5590613df2565b60405180910390fd5b6001600160a01b03811660009081526001602052604090205460ff1615610add5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608401610a55565b6001600160a01b0381166000818152600160208190526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b60405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a2610b6b3382611997565b610b753382611ac8565b50565b600033610996818585610b8b83836114d1565b610b959190613dcc565b61164b565b610ba261198f565b6001600160a01b038116610bc85760405162461bcd60e51b8152600401610a5590613df2565b6001600160a01b03811660009081526001602052604090205460ff16610c425760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608401610a55565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152600160205260409020805460ff19169055565b600160045414610cb95760405162461bcd60e51b8152600401610a5590613e40565b6002600455600c54600160981b31908111610d485760405162461bcd60e51b815260206004820152604360248201527f4e6174697665546f6b656e53706f6b653a206275726e2061646472657373206260448201527f616c616e6365206e6f742067726561746572207468616e206c617374207265706064820152621bdc9d60ea1b608482015260a401610a55565b6000600c5482610d589190613ddf565b905060006064610d887f000000000000000000000000000000000000000000000000000000000000000084613e84565b610d929190613e9b565b90506000610da08284613ddf565b600c85905590508115610dc257610db73083611be1565b610dc082611c63565b505b6000610e0f7f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000084611c73565b11610e785760405162461bcd60e51b815260206004820152603360248201527f4e6174697665546f6b656e53706f6b653a207a65726f207363616c656420616d60448201527237bab73a103a37903932b837b93a10313ab93760691b6064820152608401610a55565b604080518082018252600181528151808301835262010203600160981b0181526020808201859052925160009380840192610eb592909101613ed3565b60405160208183030381529060405281525090506000610fc86040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020016040518060400160405280306001600160a01b03168152602001888152508152602001898152602001600067ffffffffffffffff811115610f7157610f71613ef3565b604051908082528060200260200182016040528015610f9a578160200160208202803683370190505b50815260200184604051602001610fb19190613f09565b604051602081830303815290604052815250611c8a565b9050807f0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a584604051610ffc91815260200190565b60405180910390a2505060016004555050505050565b61101a61198f565b610b7581611db0565b60055460ff166110455760405162461bcd60e51b8152600401610a5590613f4e565b610b758134611f50565b611057611fca565b6108f46000612024565b60055460ff166110835760405162461bcd60e51b8152600401610a5590613f4e565b610b758134612076565b6060600a805461090590613d82565b600033816110aa82866114d1565b90508381101561110a5760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152608401610a55565b610a1c828686840361164b565b6000336109968185856117e4565b600554610100900460ff161561117d5760405162461bcd60e51b815260206004820152601e60248201527f546f6b656e53706f6b653a20616c7265616479207265676973746572656400006044820152606401610a55565b60408051606080820183527f000000000000000000000000000000000000000000000000000000000000000082527f000000000000000000000000000000000000000000000000000000000000000060208084019182527f0000000000000000000000000000000000000000000000000000000000000000151584860190815285518087018752600080825287518751818601529451858901529151151584860152865180850390950185526080909301909552818101929092529192906112559061124b90860186613c23565b85602001356120e2565b90506113006040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200160405180604001604052808860000160208101906112d89190613c23565b6001600160a01b031681526020908101869052908252620249f0908201526040016000610f71565b5050505050565b61130f61211e565b6002546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015611386573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113aa9190613f9c565b10156114115760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608401610a55565b61141a33610677565b156114805760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608401610a55565b6114c1848484848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061217792505050565b6114cb6001600055565b50505050565b6001600160a01b03918216600090815260076020908152604080832093909416825291909152205490565b611504611fca565b6001600160a01b0381166115695760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610a55565b610b7581612024565b6000600561158183601f613dcc565b901c92915050565b6001600160a01b0382166115df5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610a55565b80600860008282546115f19190613dcc565b90915550506001600160a01b0382166000818152600660209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35b5050565b6001600160a01b0383166116ad5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610a55565b6001600160a01b03821661170e5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610a55565b6001600160a01b0383811660008181526007602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b600061177c84846114d1565b905060001981146114cb57818110156117d75760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610a55565b6114cb848484840361164b565b6001600160a01b0383166118485760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610a55565b6001600160a01b0382166118aa5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610a55565b6001600160a01b038316600090815260066020526040902054818110156119225760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610a55565b6001600160a01b0380851660008181526006602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906119829086815260200190565b60405180910390a36114cb565b6108f4611fca565b6001600160a01b0382166119f75760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610a55565b6001600160a01b03821660009081526006602052604090205481811015611a6b5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610a55565b6001600160a01b03831660008181526006602090815260408083208686039055600880548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101611763565b505050565b80471015611b185760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610a55565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114611b65576040519150601f19603f3d011682016040523d82523d6000602084013e611b6a565b606091505b5050905080611ac35760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610a55565b80600b6000828254611bf39190613dcc565b90915550506040516327ad555d60e11b81526001600160a01b0383166004820152602481018290526001600160991b0190634f5aaaba90604401600060405180830381600087803b158015611c4757600080fd5b505af1158015611c5b573d6000803e3d6000fd5b505050505050565b6000611c6f3083611589565b5090565b6000611c8284848460006123ac565b949350505050565b600080611c956123dd565b60408401516020015190915015611d3a576040830151516001600160a01b0316611d175760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608401610a55565b604083015160208101519051611d3a916001600160a01b039091169083906124f1565b604051630624488560e41b81526001600160a01b03821690636244885090611d66908690600401613fb5565b6020604051808303816000875af1158015611d85573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611da99190613f9c565b9392505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015611e10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e349190613f9c565b60025490915081831115611ea45760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608401610a55565b808311611f195760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608401610a55565b6002839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b600160045414611f725760405162461bcd60e51b8152600401610a5590613e40565b6002600455611f80826125d6565b7f0000000000000000000000000000000000000000000000000000000000000000823503611fb757611fb28282612810565b611fc1565b611fb28282612a87565b50506001600455565b6003546001600160a01b031633146108f45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610a55565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600454146120985760405162461bcd60e51b8152600401610a5590613e40565b60026004556120a682612d5e565b7f00000000000000000000000000000000000000000000000000000000000000008235036120d857611fb28282612e4b565b611fc18282612ff1565b6000816000036120f45750600061099c565b306001600160a01b038416036121145761210d82611c63565b905061099c565b611da983836131ee565b6002600054036121705760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610a55565b6002600055565b7f000000000000000000000000000000000000000000000000000000000000000083146121f75760405162461bcd60e51b815260206004820152602860248201527f546f6b656e53706f6b653a20696e76616c696420736f7572636520626c6f636b60448201526718da185a5b88125160c21b6064820152608401610a55565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b03161461228a5760405162461bcd60e51b815260206004820152602960248201527f546f6b656e53706f6b653a20696e76616c6964206f726967696e2073656e646560448201526872206164647265737360b81b6064820152608401610a55565b6000818060200190518101906122a09190614143565b600554909150610100900460ff1615806122bd575060055460ff16155b156122d2576005805461ffff19166101011790555b6001815160048111156122e7576122e7613ebd565b03612320576000816020015180602001905181019061230691906141d2565b905061231a8160000151826020015161334d565b506114cb565b60028151600481111561233557612335613ebd565b036123645760008160200151806020019051810190612354919061420c565b905061231a81826080015161339a565b60405162461bcd60e51b815260206004820181905260248201527f546f6b656e53706f6b653a20696e76616c6964206d65737361676520747970656044820152606401610a55565b6000811515841515036123ca576123c38584613e84565b9050611c82565b6123d48584613e9b565b95945050505050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561243e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061246291906142dc565b9050612486816001600160a01b031660009081526001602052604090205460ff1690565b156124ec5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608401610a55565b919050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa158015612542573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125669190613f9c565b6125709190613dcc565b6040516001600160a01b0385166024820152604481018290529091506114cb90859063095ea7b360e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526134c7565b80356125f45760405162461bcd60e51b8152600401610a55906142f9565b60006126066040830160208401613c23565b6001600160a01b03160361262c5760405162461bcd60e51b8152600401610a5590614343565b600061263e6060830160408401613c23565b6001600160a01b0316036126a85760405162461bcd60e51b815260206004820152602b60248201527f546f6b656e53706f6b653a207a65726f20726563697069656e7420636f6e747260448201526a616374206164647265737360a81b6064820152608401610a55565b60008160800135116126cc5760405162461bcd60e51b8152600401610a559061438e565b60008160a001351161272c5760405162461bcd60e51b8152602060048201526024808201527f546f6b656e53706f6b653a207a65726f20726563697069656e7420676173206c6044820152631a5b5a5d60e21b6064820152608401610a55565b80608001358160a00135106127935760405162461bcd60e51b815260206004820152602760248201527f546f6b656e53706f6b653a20696e76616c696420726563697069656e742067616044820152661cc81b1a5b5a5d60ca1b6064820152608401610a55565b60006127a6610100830160e08401613c23565b6001600160a01b031603610b755760405162461bcd60e51b815260206004820152602b60248201527f546f6b656e53706f6b653a207a65726f2066616c6c6261636b2072656369706960448201526a656e74206164647265737360a81b6064820152608401610a55565b61283e6128236040840160208501613c23565b61014084013561283960e0860160c08701613c23565b613599565b60006128678261285661012086016101008701613c23565b8561012001358661014001356136aa565b60408051808201909152919350915060009080600281526020016040518061010001604052807f00000000000000000000000000000000000000000000000000000000000000008152602001306001600160a01b031681526020016128c93390565b6001600160a01b031681526020016128e76060890160408a01613c23565b6001600160a01b031681526020810187905260400161290960608901896143d1565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060a0880135602082015260400161295e610100890160e08a01613c23565b6001600160a01b03169052604051612979919060200161441f565b60405160208183030381529060405281525090506000612a3a6040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602001604051806040016040528089610100016020810190612a119190613c23565b6001600160a01b0316815260209081018890529082526080890135908201526040016000610f71565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b168787604051612a7892919061452d565b60405180910390a35050505050565b612ab18235612a9c6040850160208601613c23565b612aac60e0860160c08701613c23565b6137d0565b6000612ac98261285661012086016101008701613c23565b6040805180820190915291935091506000908060048152602001604051806101600160405280612af63390565b6001600160a01b0316815260200187600001358152602001876020016020810190612b219190613c23565b6001600160a01b03168152602001612b3f6060890160408a01613c23565b6001600160a01b0316815260208101879052604001612b6160608901896143d1565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060a08801356020820152604001612bb6610100890160e08a01613c23565b6001600160a01b0316815260808801356020820152604001612bde60e0890160c08a01613c23565b6001600160a01b03168152610140880135602091820152604051612c0392910161463c565b60408051601f19818403018152919052905290506000612134612c33612c2c60608801886143d1565b9050611572565b612c3d9190613e84565b612c4a906203a980613dcc565b90506000612d106040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200160405180604001604052808a610100016020810190612cd09190613c23565b6001600160a01b031681526020908101899052908252818101869052604080516000815280830182528184015251606090920191610fb191889101613f09565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b168888604051612d4e92919061452d565b60405180910390a3505050505050565b6000612d706060830160408401613c23565b6001600160a01b031603612dd15760405162461bcd60e51b815260206004820152602260248201527f546f6b656e53706f6b653a207a65726f20726563697069656e74206164647265604482015261737360f01b6064820152608401610a55565b60008160c0013511612df55760405162461bcd60e51b8152600401610a559061438e565b8035612e135760405162461bcd60e51b8152600401610a55906142f9565b6000612e256040830160208401613c23565b6001600160a01b031603610b755760405162461bcd60e51b8152600401610a5590614343565b612e74612e5e6040840160208501613c23565b60a0840135612839610100860160e08701613c23565b6000612e9982612e8a6080860160608701613c23565b85608001358660a001356136aa565b60408051808201909152919350915060009080600181526020016040518060400160405280876040016020810190612ed19190613c23565b6001600160a01b0316815260200186815250604051602001612ef39190613ed3565b60405160208183030381529060405281525090506000612fb36040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020016040518060400160405280896060016020810190612f8a9190613c23565b6001600160a01b03168152602090810188905290825260c0890135908201526040016000610f71565b9050336001600160a01b0316817f93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb528787604051612a7892919061471a565b61301782356130066040850160208601613c23565b612aac610100860160e08701613c23565b600061302d82612e8a6080860160608701613c23565b60408051808201825260038152815160e0810183528735815293955091935060009260208084019282820191613067918a01908a01613c23565b6001600160a01b031681526020016130856060890160408a01613c23565b6001600160a01b031681526020810187905260a0880135604082015260c088013560608201526080016130bf610100890160e08a01613c23565b6001600160a01b0316905260405161312f9190602001815181526020808301516001600160a01b0390811691830191909152604080840151821690830152606080840151908301526080808401519083015260a0808401519083015260c092830151169181019190915260e00190565b60405160208183030381529060405281525090506000612fb36040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200160405180604001604052808960600160208101906131c69190613c23565b6001600160a01b0316815260209081018890529082526203a980908201526040016000610f71565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa158015613237573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061325b9190613f9c565b90506132726001600160a01b038516333086613881565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa1580156132b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132dd9190613f9c565b90508181116133435760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610a55565b6123d48282613ddf565b816001600160a01b03167f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b8260405161338891815260200190565b60405180910390a26116478282611be1565b6133a43082611be1565b60008260000151836020015184604001518560a001516040516024016133cd94939291906147b9565b60408051601f198184030181529190526020810180516001600160e01b031663161b12ff60e11b17905260c0840151606085015191925060009161341491908590856138b9565b905080156134685783606001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff48460405161345b91815260200190565b60405180910390a26114cb565b83606001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0846040516134a791815260200190565b60405180910390a260e08401516114cb906001600160a01b031684611ac8565b600061351c826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661398e9092919063ffffffff16565b805190915015611ac3578080602001905181019061353a91906147eb565b611ac35760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610a55565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b0316146135ea5760405162461bcd60e51b8152600401610a559061480d565b81156136435760405162461bcd60e51b815260206004820152602260248201527f546f6b656e53706f6b653a206e6f6e2d7a65726f207365636f6e646172792066604482015261656560f01b6064820152608401610a55565b6001600160a01b03811615611ac35760405162461bcd60e51b815260206004820152602760248201527f546f6b656e53706f6b653a206e6f6e2d7a65726f206d756c74692d686f702066604482015266616c6c6261636b60c81b6064820152608401610a55565b6000806136b686611c63565b95506136c285856120e2565b93506136cd8661399d565b6137187f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000085611c73565b6137637f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000089611c73565b116137c45760405162461bcd60e51b815260206004820152602b60248201527f546f6b656e53706f6b653a20696e73756666696369656e7420746f6b656e732060448201526a3a37903a3930b739b332b960a91b6064820152608401610a55565b50939491935090915050565b7f0000000000000000000000000000000000000000000000000000000000000000830361381f57306001600160a01b0383160361381f5760405162461bcd60e51b8152600401610a559061480d565b6001600160a01b038116611ac35760405162461bcd60e51b815260206004820152602360248201527f546f6b656e53706f6b653a207a65726f206d756c74692d686f702066616c6c6260448201526261636b60e81b6064820152608401610a55565b6040516001600160a01b03808516602483015283166044820152606481018290526114cb9085906323b872dd60e01b9060840161259f565b6000845a101561390b5760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e742067617300000000006044820152606401610a55565b8347101561395b5760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c75650000006044820152606401610a55565b826001600160a01b03163b60000361397557506000611c82565b600080600084516020860188888bf19695505050505050565b6060611c8284846000856139ba565b6139a73082611997565b610b7562010203600160981b0182611ac8565b606082471015613a1b5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610a55565b600080866001600160a01b03168587604051613a37919061485b565b60006040518083038185875af1925050503d8060008114613a74576040519150601f19603f3d011682016040523d82523d6000602084013e613a79565b606091505b5091509150613a8a87838387613a95565b979650505050505050565b60608315613b04578251600003613afd576001600160a01b0385163b613afd5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610a55565b5081611c82565b611c828383815115613b195781518083602001fd5b8060405162461bcd60e51b8152600401610a559190613b83565b60005b83811015613b4e578181015183820152602001613b36565b50506000910152565b60008151808452613b6f816020860160208601613b33565b601f01601f19169290920160200192915050565b602081526000611da96020830184613b57565b6001600160a01b0381168114610b7557600080fd5b80356124ec81613b96565b60008060408385031215613bc957600080fd5b8235613bd481613b96565b946020939093013593505050565b600080600060608486031215613bf757600080fd5b8335613c0281613b96565b92506020840135613c1281613b96565b929592945050506040919091013590565b600060208284031215613c3557600080fd5b8135611da981613b96565b600060208284031215613c5257600080fd5b5035919050565b600060208284031215613c6b57600080fd5b813567ffffffffffffffff811115613c8257600080fd5b82016101608185031215611da957600080fd5b60006101008284031215613ca857600080fd5b50919050565b600060408284031215613ca857600080fd5b60008060008060608587031215613cd657600080fd5b843593506020850135613ce881613b96565b9250604085013567ffffffffffffffff80821115613d0557600080fd5b818701915087601f830112613d1957600080fd5b813581811115613d2857600080fd5b886020828501011115613d3a57600080fd5b95989497505060200194505050565b60008060408385031215613d5c57600080fd5b8235613d6781613b96565b91506020830135613d7781613b96565b809150509250929050565b600181811c90821680613d9657607f821691505b602082108103613ca857634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b8082018082111561099c5761099c613db6565b8181038181111561099c5761099c613db6565b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b808202811582820484141761099c5761099c613db6565b600082613eb857634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052602160045260246000fd5b81516001600160a01b03168152602080830151908201526040810161099c565b634e487b7160e01b600052604160045260246000fd5b602081526000825160058110613f2f57634e487b7160e01b600052602160045260246000fd5b806020840152506020830151604080840152611c826060840182613b57565b6020808252602e908201527f4e6174697665546f6b656e53706f6b653a20636f6e747261637420756e64657260408201526d18dbdb1b185d195c985b1a5e995960921b606082015260800190565b600060208284031215613fae57600080fd5b5051919050565b6020808252825182820152828101516001600160a01b039081166040808501919091528401518051821660608501528083015160808501526000929161010085019190606087015160a0870152608087015160e060c0880152805193849052840192600092506101208701905b8084101561404457845183168252938501936001939093019290850190614022565b5060a0880151878203601f190160e089015294506140628186613b57565b98975050505050505050565b6040805190810167ffffffffffffffff8111828210171561409157614091613ef3565b60405290565b604051610100810167ffffffffffffffff8111828210171561409157614091613ef3565b600082601f8301126140cc57600080fd5b815167ffffffffffffffff808211156140e7576140e7613ef3565b604051601f8301601f19908116603f0116810190828211818310171561410f5761410f613ef3565b8160405283815286602085880101111561412857600080fd5b614139846020830160208901613b33565b9695505050505050565b60006020828403121561415557600080fd5b815167ffffffffffffffff8082111561416d57600080fd5b908301906040828603121561418157600080fd5b61418961406e565b82516005811061419857600080fd5b81526020830151828111156141ac57600080fd5b6141b8878286016140bb565b60208301525095945050505050565b80516124ec81613b96565b6000604082840312156141e457600080fd5b6141ec61406e565b82516141f781613b96565b81526020928301519281019290925250919050565b60006020828403121561421e57600080fd5b815167ffffffffffffffff8082111561423657600080fd5b90830190610100828603121561424b57600080fd5b614253614097565b82518152614263602084016141c7565b6020820152614274604084016141c7565b6040820152614285606084016141c7565b60608201526080830151608082015260a0830151828111156142a657600080fd5b6142b2878286016140bb565b60a08301525060c083015160c08201526142ce60e084016141c7565b60e082015295945050505050565b6000602082840312156142ee57600080fd5b8151611da981613b96565b6020808252602a908201527f546f6b656e53706f6b653a207a65726f2064657374696e6174696f6e20626c6f60408201526918dad8da185a5b88125160b21b606082015260800190565b6020808252602b908201527f546f6b656e53706f6b653a207a65726f2064657374696e6174696f6e2062726960408201526a646765206164647265737360a81b606082015260800190565b60208082526023908201527f546f6b656e53706f6b653a207a65726f20726571756972656420676173206c696040820152621b5a5d60ea1b606082015260800190565b6000808335601e198436030181126143e857600080fd5b83018035915067ffffffffffffffff82111561440357600080fd5b60200191503681900382131561441857600080fd5b9250929050565b60208152815160208201526000602083015160018060a01b0380821660408501528060408601511660608501525050606083015161446860808401826001600160a01b03169052565b50608083015160a083015260a08301516101008060c085015261448f610120850183613b57565b915060c085015160e085015260e08501516144b4828601826001600160a01b03169052565b5090949350505050565b6000808335601e198436030181126144d557600080fd5b830160208101925035905067ffffffffffffffff8111156144f557600080fd5b80360382131561441857600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6040815282356040820152600061454660208501613bab565b6001600160a01b0316606083015261456060408501613bab565b6001600160a01b0316608083015261457b60608501856144be565b6101608060a08601526145936101a086018385614504565b9250608087013560c086015260a087013560e08601526145b560c08801613bab565b91506101006145ce818701846001600160a01b03169052565b6145da60e08901613bab565b92506101206145f3818801856001600160a01b03169052565b6145fe828a01613bab565b93506101409150614619828801856001600160a01b03169052565b880135918601919091529095013561018084015260209092019290925292915050565b602081526146566020820183516001600160a01b03169052565b602082015160408201526000604083015161467c60608401826001600160a01b03169052565b5060608301516001600160a01b038116608084015250608083015160a083015260a08301516101608060c08501526146b8610180850183613b57565b915060c085015160e085015260e08501516101006146e0818701836001600160a01b03169052565b860151610120868101919091528601519050610140614709818701836001600160a01b03169052565b959095015193019290925250919050565b823581526101208101602084013561473181613b96565b6001600160a01b03908116602084015260408501359061475082613b96565b16604083015261476260608501613bab565b6001600160a01b0381166060840152506080840135608083015260a084013560a083015260c084013560c083015261479c60e08501613bab565b6001600160a01b031660e083015261010090910191909152919050565b8481526001600160a01b0384811660208301528316604082015260806060820181905260009061413990830184613b57565b6000602082840312156147fd57600080fd5b81518015158114611da957600080fd5b6020808252602e908201527f546f6b656e53706f6b653a20696e76616c69642064657374696e6174696f6e2060408201526d627269646765206164647265737360901b606082015260800190565b6000825161486d818460208701613b33565b919091019291505056fea2646970667358221220404ba4ac1885b0d8a88442deb9e940c6d991bd6a70c9f49bb440b4269e32bdee64736f6c63430008120033",
}

// NativeTokenSpokeABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenSpokeMetaData.ABI instead.
var NativeTokenSpokeABI = NativeTokenSpokeMetaData.ABI

// NativeTokenSpokeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenSpokeMetaData.Bin instead.
var NativeTokenSpokeBin = NativeTokenSpokeMetaData.Bin

// DeployNativeTokenSpoke deploys a new Ethereum contract, binding an instance of NativeTokenSpoke to it.
func DeployNativeTokenSpoke(auth *bind.TransactOpts, backend bind.ContractBackend, settings TokenSpokeSettings, nativeAssetSymbol string, initialReserveImbalance *big.Int, decimalsShift uint8, multiplyOnSpoke bool, burnedFeesReportingRewardPercentage_ *big.Int) (common.Address, *types.Transaction, *NativeTokenSpoke, error) {
	parsed, err := NativeTokenSpokeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenSpokeBin), backend, settings, nativeAssetSymbol, initialReserveImbalance, decimalsShift, multiplyOnSpoke, burnedFeesReportingRewardPercentage_)
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

// MULTIHOPREQUIREDGAS is a free data retrieval call binding the contract method 0x525975e6.
//
// Solidity: function MULTI_HOP_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCaller) MULTIHOPREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSpoke.contract.Call(opts, &out, "MULTI_HOP_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPREQUIREDGAS is a free data retrieval call binding the contract method 0x525975e6.
//
// Solidity: function MULTI_HOP_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeSession) MULTIHOPREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.MULTIHOPREQUIREDGAS(&_NativeTokenSpoke.CallOpts)
}

// MULTIHOPREQUIREDGAS is a free data retrieval call binding the contract method 0x525975e6.
//
// Solidity: function MULTI_HOP_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSpoke *NativeTokenSpokeCallerSession) MULTIHOPREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSpoke.Contract.MULTIHOPREQUIREDGAS(&_NativeTokenSpoke.CallOpts)
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
