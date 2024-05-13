// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokendestination

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

// NativeTokenDestinationSettings is an auto generated low-level Go binding around an user-defined struct.
type NativeTokenDestinationSettings struct {
	NativeAssetSymbol                   string
	TeleporterRegistryAddress           common.Address
	TeleporterManager                   common.Address
	SourceBlockchainID                  [32]byte
	TokenSourceAddress                  common.Address
	InitialReserveImbalance             *big.Int
	DecimalsShift                       uint8
	MultiplyOnDestination               bool
	BurnedFeesReportingRewardPercentage *big.Int
}

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

// NativeTokenDestinationMetaData contains all meta data concerning the NativeTokenDestination contract.
var NativeTokenDestinationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"nativeAssetSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenSourceAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialReserveImbalance\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimalsShift\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"multiplyOnDestination\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"burnedFeesReportingRewardPercentage\",\"type\":\"uint256\"}],\"internalType\":\"structNativeTokenDestinationSettings\",\"name\":\"settings\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesBurned\",\"type\":\"uint256\"}],\"name\":\"ReportBurnedTxFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"BURNED_FOR_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURNED_TX_FEES_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_CALL_GAS_PER_WORD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_MINTER\",\"outputs\":[{\"internalType\":\"contractINativeMinter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTER_DESTINATION_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_BURN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnedFeesReportingRewardPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"payloadSize\",\"type\":\"uint256\"}],\"name\":\"calculateNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialReserveImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastestBurnedFeesReported\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplyOnDestination\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"registerWithSource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"name\":\"reportBurnedTxFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"sendAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenSourceAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalNativeAssetSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101806040523480156200001257600080fd5b5060405162005469380380620054698339810160408190526200003591620007a6565b80602001518160400151826060015183608001518460a001518560c001518660e001518686818a600001516040516020016200007291906200089b565b60408051601f198184030181529190528b5160036200009283826200095c565b506004620000a182826200095c565b50506001600555506001600160a01b0381166200012b5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084015b60405180910390fd5b6001600160a01b03811660808190526040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa15801562000176573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200019c919062000a28565b60075550620001ab336200053e565b620001b68162000590565b505060016009819055507302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000213573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000239919062000a28565b60a05284620002a05760405162461bcd60e51b815260206004820152603560248201526000805160206200544983398151915260448201527f20736f7572636520626c6f636b636861696e2049440000000000000000000000606482015260840162000122565b60a0518503620003285760405162461bcd60e51b815260206004820152604660248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a2063616e6e60448201527f6f74206465706c6f7920746f2073616d6520626c6f636b636861696e20617320606482015265736f7572636560d01b608482015260a40162000122565b6001600160a01b038416620003955760405162461bcd60e51b815260206004820152603560248201526000805160206200544983398151915260448201527f20746f6b656e20736f7572636520616464726573730000000000000000000000606482015260840162000122565b60128260ff161115620004055760405162461bcd60e51b815260206004820152603160248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a20696e76616044820152701b1a5908191958da5b585b1cd4da1a599d607a1b606482015260840162000122565b60c08590526001600160a01b03841660e052610140839052600a805460ff191684151781556200043790839062000b57565b6101005215156101205250505060a08401516000039250620004c59150505760405162461bcd60e51b815260206004820152603660248201527f4e6174697665546f6b656e44657374696e6174696f6e3a207a65726f20696e6960448201527f7469616c207265736572766520696d62616c616e636500000000000000000000606482015260840162000122565b6064816101000151106200052f5760405162461bcd60e51b815260206004820152602a60248201527f4e6174697665546f6b656e44657374696e6174696f6e3a20696e76616c69642060448201526970657263656e7461676560b01b606482015260840162000122565b61010001516101605262000b6f565b600880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6200059a6200060f565b6001600160a01b038116620006015760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840162000122565b6200060c816200053e565b50565b6008546001600160a01b031633146200066b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640162000122565b565b634e487b7160e01b600052604160045260246000fd5b60405161012081016001600160401b0381118282101715620006a957620006a96200066d565b60405290565b60005b83811015620006cc578181015183820152602001620006b2565b50506000910152565b600082601f830112620006e757600080fd5b81516001600160401b03808211156200070457620007046200066d565b604051601f8301601f19908116603f011681019082821181831017156200072f576200072f6200066d565b816040528381528660208588010111156200074957600080fd5b6200075c846020830160208901620006af565b9695505050505050565b80516001600160a01b03811681146200077e57600080fd5b919050565b805160ff811681146200077e57600080fd5b805180151581146200077e57600080fd5b600060208284031215620007b957600080fd5b81516001600160401b0380821115620007d157600080fd5b908301906101208286031215620007e757600080fd5b620007f162000683565b8251828111156200080157600080fd5b6200080f87828601620006d5565b825250620008206020840162000766565b6020820152620008336040840162000766565b604082015260608301516060820152620008506080840162000766565b608082015260a083015160a08201526200086d60c0840162000783565b60c08201526200088060e0840162000795565b60e08201526101009283015192810192909252509392505050565b6702bb930b83832b2160c51b815260008251620008c0816008850160208701620006af565b9190910160080192915050565b600181811c90821680620008e257607f821691505b6020821081036200090357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200095757600081815260208120601f850160051c81016020861015620009325750805b601f850160051c820191505b8181101562000953578281556001016200093e565b5050505b505050565b81516001600160401b038111156200097857620009786200066d565b6200099081620009898454620008cd565b8462000909565b602080601f831160018114620009c85760008415620009af5750858301515b600019600386901b1c1916600185901b17855562000953565b600085815260208120601f198616915b82811015620009f957888601518255948401946001909101908401620009d8565b508582101562000a185787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60006020828403121562000a3b57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600181815b8085111562000a9957816000190482111562000a7d5762000a7d62000a42565b8085161562000a8b57918102915b93841c939080029062000a5d565b509250929050565b60008262000ab25750600162000b51565b8162000ac15750600062000b51565b816001811462000ada576002811462000ae55762000b05565b600191505062000b51565b60ff84111562000af95762000af962000a42565b50506001821b62000b51565b5060208310610133831016604e8410600b841016171562000b2a575081810a62000b51565b62000b36838362000a58565b806000190482111562000b4d5762000b4d62000a42565b0290505b92915050565b600062000b6860ff84168362000aa1565b9392505050565b60805160a05160c05160e0516101005161012051610140516101605161479b62000cae600039600081816105080152610f6b01526000818161065f01528181610a230152610bd20152600081816103ca01528181610a6f01528181610ff201528181613248015261329301526000818161076801528181610a4601528181610fd101528181613227015261327201526000818161089301528181610b0401528181611113015281816121f2015281816125f6015281816128c501528181612bce0152612d2801526000818161043d01528181610ad9015281816110ed015281816121c2015281816125d00152818161289501528181612ba80152612cb20152600081816107ea015281816122c2015281816123ed0152612a1c01526000818161037e0152818161133501528181611dbc0152612ee9015261479b6000f3fe60806040526004361061028c5760003560e01c80635eb995141161015a578063ba3f5a12116100c1578063d2cc7a701161007a578063d2cc7a701461080c578063dd62ed3e14610821578063ecd4ed1b1461054a578063f2fde38b14610841578063f3f981d814610861578063f5ea0603146108815761029b565b8063ba3f5a1214610756578063c452165e1461078a578063c868efaa146107a2578063d0e30db01461029b578063d10a5b8c146107c2578063d127dc9b146107d85761029b565b80638da5cb5b116101135780638da5cb5b1461069457806395d89b41146106b257806397314297146106c7578063a2309ff814610700578063a457c2d714610716578063a9059cbb146107365761029b565b80635eb99514146105cf5780636e6eef8d146105ef57806370a0823114610602578063715018a6146106385780638ac7dd201461064d5780638bf2fa94146106815761029b565b80632b0d8f18116101fe5780634511243e116101b75780634511243e1461052a57806347a9a22c1461054a57806349e3284e14610567578063525975e61461058157806355538c8b1461059857806355c26725146105b85761029b565b80632b0d8f181461045f5780632e1a7d4d1461047f578063313ce5671461049f578063329c3e12146104bb57806339509351146104d65780633a23dfe2146104f65761029b565b80631906529c116102505780631906529c146103575780631a7f5bec1461036c5780632001eff6146103b857806322366844146103ec57806323b872dd1461040b57806329b7b3fd1461042b5761029b565b806306fdde03146102a3578063095ea7b3146102ce578063151637f6146102fe57806315beb59f1461031e57806318160ddd146103425761029b565b3661029b576102996108b5565b005b6102996108b5565b3480156102af57600080fd5b506102b86108f6565b6040516102c591906139df565b60405180910390f35b3480156102da57600080fd5b506102ee6102e9366004613a12565b610988565b60405190151581526020016102c5565b34801561030a57600080fd5b50610299610319366004613a3e565b6109a2565b34801561032a57600080fd5b5061033461213481565b6040519081526020016102c5565b34801561034e57600080fd5b50600254610334565b34801561036357600080fd5b50610334610bb0565b34801561037857600080fd5b506103a07f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016102c5565b3480156103c457600080fd5b506102ee7f000000000000000000000000000000000000000000000000000000000000000081565b3480156103f857600080fd5b50600a546102ee90610100900460ff1681565b34801561041757600080fd5b506102ee610426366004613a56565b610c11565b34801561043757600080fd5b506103347f000000000000000000000000000000000000000000000000000000000000000081565b34801561046b57600080fd5b5061029961047a366004613a97565b610c35565b34801561048b57600080fd5b5061029961049a366004613ab4565b610d2e565b3480156104ab57600080fd5b50604051601281526020016102c5565b3480156104c757600080fd5b506103a06001600160991b0181565b3480156104e257600080fd5b506102ee6104f1366004613a12565b610d7a565b34801561050257600080fd5b506103347f000000000000000000000000000000000000000000000000000000000000000081565b34801561053657600080fd5b50610299610545366004613a97565b610d9c565b34801561055657600080fd5b506103a062010203600160981b0181565b34801561057357600080fd5b50600a546102ee9060ff1681565b34801561058d57600080fd5b506103346203a98081565b3480156105a457600080fd5b506102996105b3366004613ab4565b610e99565b3480156105c457600080fd5b50610334620249f081565b3480156105db57600080fd5b506102996105ea366004613ab4565b61120d565b6102996105fd366004613acd565b61121e565b34801561060e57600080fd5b5061033461061d366004613a97565b6001600160a01b031660009081526020819052604090205490565b34801561064457600080fd5b5061029961124a565b34801561065957600080fd5b506103347f000000000000000000000000000000000000000000000000000000000000000081565b61029961068f366004613b09565b61125c565b3480156106a057600080fd5b506008546001600160a01b03166103a0565b3480156106be57600080fd5b506102b8611288565b3480156106d357600080fd5b506102ee6106e2366004613a97565b6001600160a01b031660009081526006602052604090205460ff1690565b34801561070c57600080fd5b50610334600b5481565b34801561072257600080fd5b506102ee610731366004613a12565b611297565b34801561074257600080fd5b506102ee610751366004613a12565b611312565b34801561076257600080fd5b506103347f000000000000000000000000000000000000000000000000000000000000000081565b34801561079657600080fd5b506103a0600160981b81565b3480156107ae57600080fd5b506102996107bd366004613b1c565b611320565b3480156107ce57600080fd5b50610334600c5481565b3480156107e457600080fd5b506103347f000000000000000000000000000000000000000000000000000000000000000081565b34801561081857600080fd5b50600754610334565b34801561082d57600080fd5b5061033461083c366004613ba5565b6114e4565b34801561084d57600080fd5b5061029961085c366004613a97565b61150f565b34801561086d57600080fd5b5061033461087c366004613ab4565b611585565b34801561088d57600080fd5b506103a07f000000000000000000000000000000000000000000000000000000000000000081565b60405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a26108f4333461159c565b565b60606003805461090590613bde565b80601f016020809104026020016040519081016040528092919081815260200182805461093190613bde565b801561097e5780601f106109535761010080835404028352916020019161097e565b820191906000526020600020905b81548152906001019060200180831161096157829003601f168201915b5050505050905090565b60003361099681858561165c565b60019150505b92915050565b600a54610100900460ff1615610a165760405162461bcd60e51b815260206004820152602e60248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a20616c726560448201526d18591e481c9959da5cdd195c995960921b60648201526084015b60405180910390fd5b60408051606080820183527f000000000000000000000000000000000000000000000000000000000000000082527f000000000000000000000000000000000000000000000000000000000000000060208084019182527f00000000000000000000000000000000000000000000000000000000000000001515848601908152855180870187526000815286518651818501529351848801529051151583850152855180840390940184526080909201855281810192909252835160c0810185527f000000000000000000000000000000000000000000000000000000000000000081526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928101929092529192610baa91908101610b4336879003870187613c8a565b8152620249f060208201526040016000604051908082528060200260200182016040528015610b7c578160200160208202803683370190505b50815260200183604051602001610b939190613cc4565b604051602081830303815290604052815250611781565b50505050565b600080610bcc62010203600160981b0131600160981b31613d1f565b905060007f0000000000000000000000000000000000000000000000000000000000000000600b54610bfe9190613d1f565b9050610c0a8282613d32565b9250505090565b600033610c1f8582856118a7565b610c2a85858561191b565b506001949350505050565b610c3d611abf565b6001600160a01b038116610c635760405162461bcd60e51b8152600401610a0d90613d45565b6001600160a01b03811660009081526006602052604090205460ff1615610ce25760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608401610a0d565b6001600160a01b038116600081815260066020526040808220805460ff19166001179055517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b60405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a2610d6d3382611ac7565b610d773382611bf6565b50565b600033610996818585610d8d83836114e4565b610d979190613d1f565b61165c565b610da4611abf565b6001600160a01b038116610dca5760405162461bcd60e51b8152600401610a0d90613d45565b6001600160a01b03811660009081526006602052604090205460ff16610e445760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608401610a0d565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152600660205260409020805460ff19169055565b600160095414610ebb5760405162461bcd60e51b8152600401610a0d90613d93565b6002600955600c54600160981b31908111610f505760405162461bcd60e51b815260206004820152604960248201527f4e6174697665546f6b656e44657374696e6174696f6e3a206275726e2061646460448201527f726573732062616c616e6365206e6f742067726561746572207468616e206c616064820152681cdd081c995c1bdc9d60ba1b608482015260a401610a0d565b6000600c5482610f609190613d32565b905060006064610f907f000000000000000000000000000000000000000000000000000000000000000084613dd7565b610f9a9190613dee565b90506000610fa88284613d32565b600c85905590508115610fca57610fbf3083611d0f565b610fc882611d91565b505b60006110177f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000084611da1565b1161108a5760405162461bcd60e51b815260206004820152603960248201527f4e6174697665546f6b656e44657374696e6174696f6e3a207a65726f2073636160448201527f6c656420616d6f756e7420746f207265706f7274206275726e000000000000006064820152608401610a0d565b604080518082018252600181528151808301835262010203600160981b01815260208082018590529251600093808401926110c792909101613e10565b604051602081830303815290604052815250905060006111c36040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020016040518060400160405280306001600160a01b03168152602001888152508152602001898152602001600067ffffffffffffffff81111561118357611183613c28565b6040519080825280602002602001820160405280156111ac578160200160208202803683370190505b50815260200184604051602001610b939190613cc4565b9050807f0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5846040516111f791815260200190565b60405180910390a2505060016009555050505050565b611215611abf565b610d7781611db8565b600a5460ff166112405760405162461bcd60e51b8152600401610a0d90613e30565b610d778134611f58565b6112526126d5565b6108f4600061272f565b600a5460ff1661127e5760405162461bcd60e51b8152600401610a0d90613e30565b610d778134612781565b60606004805461090590613bde565b600033816112a582866114e4565b9050838110156113055760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152608401610a0d565b610c2a828686840361165c565b60003361099681858561191b565b611328612c57565b6007546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa15801561139f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113c39190613e84565b101561142a5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608401610a0d565b611433336106e2565b156114995760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608401610a0d565b6114da848484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612cb092505050565b610baa6001600555565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6115176126d5565b6001600160a01b03811661157c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610a0d565b610d778161272f565b6000600561159483601f613d1f565b901c92915050565b6001600160a01b0382166115f25760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610a0d565b80600260008282546116049190613d1f565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35b5050565b6001600160a01b0383166116be5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610a0d565b6001600160a01b03821661171f5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610a0d565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b60008061178c612ee4565b60408401516020015190915015611831576040830151516001600160a01b031661180e5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608401610a0d565b604083015160208101519051611831916001600160a01b03909116908390612ff8565b604051630624488560e41b81526001600160a01b0382169063624488509061185d908690600401613e9d565b6020604051808303816000875af115801561187c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118a09190613e84565b9392505050565b60006118b384846114e4565b90506000198114610baa578181101561190e5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610a0d565b610baa848484840361165c565b6001600160a01b03831661197f5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610a0d565b6001600160a01b0382166119e15760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610a0d565b6001600160a01b03831660009081526020819052604090205481811015611a595760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610a0d565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610baa565b6108f46126d5565b6001600160a01b038216611b275760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610a0d565b6001600160a01b03821660009081526020819052604090205481811015611b9b5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610a0d565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101611774565b505050565b80471015611c465760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610a0d565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114611c93576040519150601f19603f3d011682016040523d82523d6000602084013e611c98565b606091505b5050905080611bf15760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610a0d565b80600b6000828254611d219190613d1f565b90915550506040516327ad555d60e11b81526001600160a01b0383166004820152602481018290526001600160991b0190634f5aaaba90604401600060405180830381600087803b158015611d7557600080fd5b505af1158015611d89573d6000803e3d6000fd5b505050505050565b6000611d9d308361159c565b5090565b6000611db084848460006130dd565b949350505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015611e18573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e3c9190613e84565b60075490915081831115611eac5760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608401610a0d565b808311611f215760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608401610a0d565b6007839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b600160095414611f7a5760405162461bcd60e51b8152600401610a0d90613d93565b60026009556000611f916060840160408501613a97565b6001600160a01b031603611ffb5760405162461bcd60e51b815260206004820152603b602482015260008051602061472683398151915260448201527f20726563697069656e7420636f6e7472616374206164647265737300000000006064820152608401610a0d565b600082608001351161201f5760405162461bcd60e51b8152600401610a0d90613f56565b60008260a001351161207e5760405162461bcd60e51b81526020600482015260346024820152600080516020614726833981519152604482015273081c9958da5c1a595b9d0819d85cc81b1a5b5a5d60621b6064820152608401610a0d565b81608001358260a00135106120e95760405162461bcd60e51b8152602060048201526037602482015260008051602061474683398151915260448201527f6c696420726563697069656e7420676173206c696d69740000000000000000006064820152608401610a0d565b60006120fc610100840160e08501613a97565b6001600160a01b0316036121665760405162461bcd60e51b815260206004820152603b602482015260008051602061472683398151915260448201527f2066616c6c6261636b20726563697069656e74206164647265737300000000006064820152608401610a0d565b60006121a1833561217d6040860160208701613a97565b8461219061012088016101008901613a97565b87610120013588610140013561310e565b604080518082019091526000815260606020820152919350915060808401357f00000000000000000000000000000000000000000000000000000000000000008535036123b3576001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166122226040870160208801613a97565b6001600160a01b0316146122485760405162461bcd60e51b8152600401610a0d90613f97565b6101408501351561226b5760405162461bcd60e51b8152600401610a0d90613fe2565b600061227d60e0870160c08801613a97565b6001600160a01b0316146122a35760405162461bcd60e51b8152600401610a0d90614034565b6040805180820190915280600281526020016040518060e001604052807f00000000000000000000000000000000000000000000000000000000000000008152602001336001600160a01b031681526020018860400160208101906123089190613a97565b6001600160a01b031681526020810188905260400161232a60608a018a614091565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060a0890135602082015260400161237f6101008a0160e08b01613a97565b6001600160a01b0316905260405161239a91906020016140df565b60405160208183030381529060405281525091506125be565b60006123c560e0870160c08801613a97565b6001600160a01b0316036123eb5760405162461bcd60e51b8152600401610a0d90614155565b7f000000000000000000000000000000000000000000000000000000000000000085350361244a57306124246040870160208801613a97565b6001600160a01b03160361244a5760405162461bcd60e51b8152600401610a0d90613f97565b604080518082019091528060048152602001604051806101600160405280336001600160a01b03168152602001886000013581526020018860200160208101906124949190613a97565b6001600160a01b031681526020016124b260608a0160408b01613a97565b6001600160a01b03168152602081018890526040016124d460608a018a614091565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060a089013560208201526040016125296101008a0160e08b01613a97565b6001600160a01b031681526080890135602082015260400161255160e08a0160c08b01613a97565b6001600160a01b03168152610140890135602091820152604051612576929101614196565b60408051601f19818403018152919052905291506121346125a461259d6060880188614091565b9050611585565b6125ae9190613dd7565b6125bb906203a980613d1f565b90505b60006126826040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200160405180604001604052808a6101000160208101906126429190613a97565b6001600160a01b031681526020908101899052908252818101869052604080516000815280830182528184015251606090920191610b9391889101613cc4565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b1688886040516126c09291906142e3565b60405180910390a35050600160095550505050565b6008546001600160a01b031633146108f45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610a0d565b600880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600954146127a35760405162461bcd60e51b8152600401610a0d90613d93565b600260095560006127ba6060840160408501613a97565b6001600160a01b0316036128195760405162461bcd60e51b8152602060048201526032602482015260008051602061472683398151915260448201527120726563697069656e74206164647265737360701b6064820152608401610a0d565b60008260c001351161283d5760405162461bcd60e51b8152600401610a0d90613f56565b600061287483356128546040860160208701613a97565b846128656080880160608901613a97565b87608001358860a0013561310e565b604080518082019091526000815260606020820152919350915060c08401357f00000000000000000000000000000000000000000000000000000000000000008535036129e1576001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166128f56040870160208801613a97565b6001600160a01b03161461291b5760405162461bcd60e51b8152600401610a0d90613f97565b60a08501351561293d5760405162461bcd60e51b8152600401610a0d90613fe2565b6000612950610100870160e08801613a97565b6001600160a01b0316146129765760405162461bcd60e51b8152600401610a0d90614034565b60408051808201909152806001815260200160405180604001604052808860400160208101906129a69190613a97565b6001600160a01b03168152602001878152506040516020016129c89190613e10565b6040516020818303038152906040528152509150612b96565b60006129f4610100870160e08801613a97565b6001600160a01b031603612a1a5760405162461bcd60e51b8152600401610a0d90614155565b7f0000000000000000000000000000000000000000000000000000000000000000853503612a795730612a536040870160208801613a97565b6001600160a01b031603612a795760405162461bcd60e51b8152600401610a0d90613f97565b6040805180820190915280600381526020016040518060e0016040528088600001358152602001886020016020810190612ab39190613a97565b6001600160a01b03168152602001612ad160608a0160408b01613a97565b6001600160a01b031681526020810188905260a0890135604082015260c08901356060820152608001612b0b6101008a0160e08b01613a97565b6001600160a01b03169052604051612b7b9190602001815181526020808301516001600160a01b0390811691830191909152604080840151821690830152606080840151908301526080808401519083015260a0808401519083015260c092830151169181019190915260e00190565b60405160208183030381529060405281525091506203a98090505b6000612c196040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200160405180604001604052808a60600160208101906126429190613a97565b9050336001600160a01b0316817f93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb5288886040516126c09291906143f2565b600260055403612ca95760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610a0d565b6002600555565b7f00000000000000000000000000000000000000000000000000000000000000008314612d265760405162461bcd60e51b8152602060048201526030602482015260008051602061474683398151915260448201526f3634b21039b7bab931b29031b430b4b760811b6064820152608401610a0d565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b031614612dbb5760405162461bcd60e51b8152602060048201526038602482015260008051602061474683398151915260448201527f6c696420746f6b656e20736f75726365206164647265737300000000000000006064820152608401610a0d565b600081806020019051810190612dd19190614519565b600a54909150610100900460ff161580612dee5750600a5460ff16155b15612e0357600a805461ffff19166101011790555b600181516004811115612e1857612e18613c12565b03612e515760008160200151806020019051810190612e3791906145a8565b9050612e4b81600001518260200151613339565b50610baa565b600281516004811115612e6657612e66613c12565b03612e955760008160200151806020019051810190612e8591906145e2565b9050612e4b818260600151613386565b60405162461bcd60e51b8152602060048201526030602482015260008051602061474683398151915260448201526f6c6964206d657373616765207479706560801b6064820152608401610a0d565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f6991906146a0565b9050612f8d816001600160a01b031660009081526006602052604090205460ff1690565b15612ff35760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608401610a0d565b919050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa158015613049573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061306d9190613e84565b6130779190613d1f565b6040516001600160a01b038516602482015260448101829052909150610baa90859063095ea7b360e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526134ac565b6000811515841515036130fb576130f48584613dd7565b9050611db0565b6131058584613dee565b95945050505050565b600080876131725760405162461bcd60e51b815260206004820152603a602482015260008051602061472683398151915260448201527f2064657374696e6174696f6e20626c6f636b636861696e2049440000000000006064820152608401610a0d565b6001600160a01b0387166131dc5760405162461bcd60e51b815260206004820152603b602482015260008051602061472683398151915260448201527f2064657374696e6174696f6e20627269646765206164647265737300000000006064820152608401610a0d565b6131e586611d91565b9550831561321957306001600160a01b0386160361320c5761320684611d91565b50613219565b613216858561357e565b93505b613222866136dd565b61326d7f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000085611da1565b6132b87f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000089611da1565b1161332b5760405162461bcd60e51b815260206004820152603b60248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a20696e737560448201527f6666696369656e7420746f6b656e7320746f207472616e7366657200000000006064820152608401610a0d565b509396919550909350505050565b816001600160a01b03167f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b8260405161337491815260200190565b60405180910390a26116588282611d0f565b6133903082611d0f565b60008260000151836020015184608001516040516024016133b3939291906146bd565b60408051601f198184030181529181526020820180516001600160e01b0316630d57eddf60e21b17905260a0850151908501519192506000916133f991908590856136fa565b9050801561344d5783604001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff48460405161344091815260200190565b60405180910390a2610baa565b83604001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb08460405161348c91815260200190565b60405180910390a260c0840151610baa906001600160a01b031684611bf6565b6000613501826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166137cf9092919063ffffffff16565b805190915015611bf1578080602001905181019061351f91906146e7565b611bf15760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610a0d565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa1580156135c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135eb9190613e84565b90506136026001600160a01b0385163330866137de565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015613649573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061366d9190613e84565b90508181116136d35760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610a0d565b6131058282613d32565b6136e73082611ac7565b610d7762010203600160981b0182611bf6565b6000845a101561374c5760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e742067617300000000006044820152606401610a0d565b8347101561379c5760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c75650000006044820152606401610a0d565b826001600160a01b03163b6000036137b657506000611db0565b600080600084516020860188888bf19695505050505050565b6060611db08484600085613816565b6040516001600160a01b0380851660248301528316604482015260648101829052610baa9085906323b872dd60e01b906084016130a6565b6060824710156138775760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610a0d565b600080866001600160a01b031685876040516138939190614709565b60006040518083038185875af1925050503d80600081146138d0576040519150601f19603f3d011682016040523d82523d6000602084013e6138d5565b606091505b50915091506138e6878383876138f1565b979650505050505050565b60608315613960578251600003613959576001600160a01b0385163b6139595760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610a0d565b5081611db0565b611db083838151156139755781518083602001fd5b8060405162461bcd60e51b8152600401610a0d91906139df565b60005b838110156139aa578181015183820152602001613992565b50506000910152565b600081518084526139cb81602086016020860161398f565b601f01601f19169290920160200192915050565b6020815260006118a060208301846139b3565b6001600160a01b0381168114610d7757600080fd5b8035612ff3816139f2565b60008060408385031215613a2557600080fd5b8235613a30816139f2565b946020939093013593505050565b600060408284031215613a5057600080fd5b50919050565b600080600060608486031215613a6b57600080fd5b8335613a76816139f2565b92506020840135613a86816139f2565b929592945050506040919091013590565b600060208284031215613aa957600080fd5b81356118a0816139f2565b600060208284031215613ac657600080fd5b5035919050565b600060208284031215613adf57600080fd5b813567ffffffffffffffff811115613af657600080fd5b820161016081850312156118a057600080fd5b60006101008284031215613a5057600080fd5b60008060008060608587031215613b3257600080fd5b843593506020850135613b44816139f2565b9250604085013567ffffffffffffffff80821115613b6157600080fd5b818701915087601f830112613b7557600080fd5b813581811115613b8457600080fd5b886020828501011115613b9657600080fd5b95989497505060200194505050565b60008060408385031215613bb857600080fd5b8235613bc3816139f2565b91506020830135613bd3816139f2565b809150509250929050565b600181811c90821680613bf257607f821691505b602082108103613a5057634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613c6157613c61613c28565b60405290565b60405160e0810167ffffffffffffffff81118282101715613c6157613c61613c28565b600060408284031215613c9c57600080fd5b613ca4613c3e565b8235613caf816139f2565b81526020928301359281019290925250919050565b602081526000825160058110613cea57634e487b7160e01b600052602160045260246000fd5b806020840152506020830151604080840152611db060608401826139b3565b634e487b7160e01b600052601160045260246000fd5b8082018082111561099c5761099c613d09565b8181038181111561099c5761099c613d09565b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b808202811582820484141761099c5761099c613d09565b600082613e0b57634e487b7160e01b600052601260045260246000fd5b500490565b81516001600160a01b03168152602080830151908201526040810161099c565b60208082526034908201527f4e6174697665546f6b656e44657374696e6174696f6e3a20636f6e7472616374604082015273081d5b99195c98dbdb1b185d195c985b1a5e995960621b606082015260800190565b600060208284031215613e9657600080fd5b5051919050565b6020808252825182820152828101516001600160a01b039081166040808501919091528401518051821660608501528083015160808501526000929161010085019190606087015160a0870152608087015160e060c0880152805193849052840192600092506101208701905b80841015613f2c57845183168252938501936001939093019290850190613f0a565b5060a0880151878203601f190160e08901529450613f4a81866139b3565b98975050505050505050565b6020808252603390820152600080516020614726833981519152604082015272081c995c5d5a5c99590819d85cc81b1a5b5a5d606a1b606082015260800190565b6020808252603e9082015260008051602061474683398151915260408201527f6c69642064657374696e6174696f6e2062726964676520616464726573730000606082015260800190565b60208082526032908201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a206e6f6e2d6040820152717a65726f207365636f6e646172792066656560701b606082015260800190565b60208082526037908201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a206e6f6e2d60408201527f7a65726f206d756c74692d686f702066616c6c6261636b000000000000000000606082015260800190565b6000808335601e198436030181126140a857600080fd5b83018035915067ffffffffffffffff8211156140c357600080fd5b6020019150368190038213156140d857600080fd5b9250929050565b60208152815160208201526000602083015160018060a01b038082166040850152806040860151166060850152606085015160808501526080850151915060e060a08501526141326101008501836139b3565b915060a085015160c08501528060c08601511660e0850152508091505092915050565b6020808252603390820152600080516020614726833981519152604082015272206d756c74692d686f702066616c6c6261636b60681b606082015260800190565b602081526141b06020820183516001600160a01b03169052565b60208201516040820152600060408301516141d660608401826001600160a01b03169052565b5060608301516001600160a01b038116608084015250608083015160a083015260a08301516101608060c08501526142126101808501836139b3565b915060c085015160e085015260e085015161010061423a818701836001600160a01b03169052565b860151610120868101919091528601519050610140614263818701836001600160a01b03169052565b959095015193019290925250919050565b6000808335601e1984360301811261428b57600080fd5b830160208101925035905067ffffffffffffffff8111156142ab57600080fd5b8036038213156140d857600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b604081528235604082015260006142fc60208501613a07565b6001600160a01b0316606083015261431660408501613a07565b6001600160a01b031660808301526143316060850185614274565b6101608060a08601526143496101a0860183856142ba565b9250608087013560c086015260a087013560e086015261436b60c08801613a07565b9150610100614384818701846001600160a01b03169052565b61439060e08901613a07565b92506101206143a9818801856001600160a01b03169052565b6143b4828a01613a07565b935061014091506143cf828801856001600160a01b03169052565b880135918601919091529095013561018084015260209092019290925292915050565b8235815261012081016020840135614409816139f2565b6001600160a01b039081166020840152604085013590614428826139f2565b16604083015261443a60608501613a07565b6001600160a01b0381166060840152506080840135608083015260a084013560a083015260c084013560c083015261447460e08501613a07565b6001600160a01b031660e083015261010090910191909152919050565b600082601f8301126144a257600080fd5b815167ffffffffffffffff808211156144bd576144bd613c28565b604051601f8301601f19908116603f011681019082821181831017156144e5576144e5613c28565b816040528381528660208588010111156144fe57600080fd5b61450f84602083016020890161398f565b9695505050505050565b60006020828403121561452b57600080fd5b815167ffffffffffffffff8082111561454357600080fd5b908301906040828603121561455757600080fd5b61455f613c3e565b82516005811061456e57600080fd5b815260208301518281111561458257600080fd5b61458e87828601614491565b60208301525095945050505050565b8051612ff3816139f2565b6000604082840312156145ba57600080fd5b6145c2613c3e565b82516145cd816139f2565b81526020928301519281019290925250919050565b6000602082840312156145f457600080fd5b815167ffffffffffffffff8082111561460c57600080fd5b9083019060e0828603121561462057600080fd5b614628613c67565b825181526146386020840161459d565b60208201526146496040840161459d565b60408201526060830151606082015260808301518281111561466a57600080fd5b61467687828601614491565b60808301525060a083015160a082015261469260c0840161459d565b60c082015295945050505050565b6000602082840312156146b257600080fd5b81516118a0816139f2565b8381526001600160a01b0383166020820152606060408201819052600090613105908301846139b3565b6000602082840312156146f957600080fd5b815180151581146118a057600080fd5b6000825161471b81846020870161398f565b919091019291505056fe54656c65706f72746572546f6b656e44657374696e6174696f6e3a207a65726f54656c65706f72746572546f6b656e44657374696e6174696f6e3a20696e7661a26469706673582212206c5128b884319f91e93f46f4f5efd34e11e2fde4ac9c01cf89e8fc9ec90f191364736f6c6343000812003354656c65706f72746572546f6b656e44657374696e6174696f6e3a207a65726f",
}

// NativeTokenDestinationABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenDestinationMetaData.ABI instead.
var NativeTokenDestinationABI = NativeTokenDestinationMetaData.ABI

// NativeTokenDestinationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenDestinationMetaData.Bin instead.
var NativeTokenDestinationBin = NativeTokenDestinationMetaData.Bin

// DeployNativeTokenDestination deploys a new Ethereum contract, binding an instance of NativeTokenDestination to it.
func DeployNativeTokenDestination(auth *bind.TransactOpts, backend bind.ContractBackend, settings NativeTokenDestinationSettings) (common.Address, *types.Transaction, *NativeTokenDestination, error) {
	parsed, err := NativeTokenDestinationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenDestinationBin), backend, settings)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenDestination{NativeTokenDestinationCaller: NativeTokenDestinationCaller{contract: contract}, NativeTokenDestinationTransactor: NativeTokenDestinationTransactor{contract: contract}, NativeTokenDestinationFilterer: NativeTokenDestinationFilterer{contract: contract}}, nil
}

// NativeTokenDestination is an auto generated Go binding around an Ethereum contract.
type NativeTokenDestination struct {
	NativeTokenDestinationCaller     // Read-only binding to the contract
	NativeTokenDestinationTransactor // Write-only binding to the contract
	NativeTokenDestinationFilterer   // Log filterer for contract events
}

// NativeTokenDestinationCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenDestinationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenDestinationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenDestinationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenDestinationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenDestinationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenDestinationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenDestinationSession struct {
	Contract     *NativeTokenDestination // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NativeTokenDestinationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenDestinationCallerSession struct {
	Contract *NativeTokenDestinationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// NativeTokenDestinationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenDestinationTransactorSession struct {
	Contract     *NativeTokenDestinationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// NativeTokenDestinationRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenDestinationRaw struct {
	Contract *NativeTokenDestination // Generic contract binding to access the raw methods on
}

// NativeTokenDestinationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenDestinationCallerRaw struct {
	Contract *NativeTokenDestinationCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenDestinationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenDestinationTransactorRaw struct {
	Contract *NativeTokenDestinationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenDestination creates a new instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestination(address common.Address, backend bind.ContractBackend) (*NativeTokenDestination, error) {
	contract, err := bindNativeTokenDestination(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestination{NativeTokenDestinationCaller: NativeTokenDestinationCaller{contract: contract}, NativeTokenDestinationTransactor: NativeTokenDestinationTransactor{contract: contract}, NativeTokenDestinationFilterer: NativeTokenDestinationFilterer{contract: contract}}, nil
}

// NewNativeTokenDestinationCaller creates a new read-only instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestinationCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenDestinationCaller, error) {
	contract, err := bindNativeTokenDestination(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationCaller{contract: contract}, nil
}

// NewNativeTokenDestinationTransactor creates a new write-only instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestinationTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenDestinationTransactor, error) {
	contract, err := bindNativeTokenDestination(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTransactor{contract: contract}, nil
}

// NewNativeTokenDestinationFilterer creates a new log filterer instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestinationFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenDestinationFilterer, error) {
	contract, err := bindNativeTokenDestination(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationFilterer{contract: contract}, nil
}

// bindNativeTokenDestination binds a generic wrapper to an already deployed contract.
func bindNativeTokenDestination(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenDestinationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenDestination *NativeTokenDestinationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenDestination.Contract.NativeTokenDestinationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenDestination *NativeTokenDestinationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.NativeTokenDestinationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenDestination *NativeTokenDestinationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.NativeTokenDestinationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenDestination *NativeTokenDestinationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenDestination.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenDestination *NativeTokenDestinationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenDestination *NativeTokenDestinationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.contract.Transact(opts, method, params...)
}

// BURNEDFORBRIDGEADDRESS is a free data retrieval call binding the contract method 0x47a9a22c.
//
// Solidity: function BURNED_FOR_BRIDGE_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BURNEDFORBRIDGEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "BURNED_FOR_BRIDGE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDFORBRIDGEADDRESS is a free data retrieval call binding the contract method 0x47a9a22c.
//
// Solidity: function BURNED_FOR_BRIDGE_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) BURNEDFORBRIDGEADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNEDFORBRIDGEADDRESS(&_NativeTokenDestination.CallOpts)
}

// BURNEDFORBRIDGEADDRESS is a free data retrieval call binding the contract method 0x47a9a22c.
//
// Solidity: function BURNED_FOR_BRIDGE_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BURNEDFORBRIDGEADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNEDFORBRIDGEADDRESS(&_NativeTokenDestination.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BURNEDTXFEESADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "BURNED_TX_FEES_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNEDTXFEESADDRESS(&_NativeTokenDestination.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNEDTXFEESADDRESS(&_NativeTokenDestination.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) MULTIHOPCALLGASPERWORD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "MULTI_HOP_CALL_GAS_PER_WORD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _NativeTokenDestination.Contract.MULTIHOPCALLGASPERWORD(&_NativeTokenDestination.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _NativeTokenDestination.Contract.MULTIHOPCALLGASPERWORD(&_NativeTokenDestination.CallOpts)
}

// MULTIHOPREQUIREDGAS is a free data retrieval call binding the contract method 0x525975e6.
//
// Solidity: function MULTI_HOP_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) MULTIHOPREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "MULTI_HOP_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPREQUIREDGAS is a free data retrieval call binding the contract method 0x525975e6.
//
// Solidity: function MULTI_HOP_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) MULTIHOPREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.MULTIHOPREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// MULTIHOPREQUIREDGAS is a free data retrieval call binding the contract method 0x525975e6.
//
// Solidity: function MULTI_HOP_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) MULTIHOPREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.MULTIHOPREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) NATIVEMINTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "NATIVE_MINTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenDestination.Contract.NATIVEMINTER(&_NativeTokenDestination.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenDestination.Contract.NATIVEMINTER(&_NativeTokenDestination.CallOpts)
}

// REGISTERDESTINATIONREQUIREDGAS is a free data retrieval call binding the contract method 0x55c26725.
//
// Solidity: function REGISTER_DESTINATION_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) REGISTERDESTINATIONREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "REGISTER_DESTINATION_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGISTERDESTINATIONREQUIREDGAS is a free data retrieval call binding the contract method 0x55c26725.
//
// Solidity: function REGISTER_DESTINATION_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) REGISTERDESTINATIONREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.REGISTERDESTINATIONREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// REGISTERDESTINATIONREQUIREDGAS is a free data retrieval call binding the contract method 0x55c26725.
//
// Solidity: function REGISTER_DESTINATION_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) REGISTERDESTINATIONREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.REGISTERDESTINATIONREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// SOURCECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xecd4ed1b.
//
// Solidity: function SOURCE_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) SOURCECHAINBURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "SOURCE_CHAIN_BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SOURCECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xecd4ed1b.
//
// Solidity: function SOURCE_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) SOURCECHAINBURNADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.SOURCECHAINBURNADDRESS(&_NativeTokenDestination.CallOpts)
}

// SOURCECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xecd4ed1b.
//
// Solidity: function SOURCE_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) SOURCECHAINBURNADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.SOURCECHAINBURNADDRESS(&_NativeTokenDestination.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NativeTokenDestination.Contract.Allowance(&_NativeTokenDestination.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NativeTokenDestination.Contract.Allowance(&_NativeTokenDestination.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NativeTokenDestination.Contract.BalanceOf(&_NativeTokenDestination.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NativeTokenDestination.Contract.BalanceOf(&_NativeTokenDestination.CallOpts, account)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationSession) BlockchainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.BlockchainID(&_NativeTokenDestination.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BlockchainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.BlockchainID(&_NativeTokenDestination.CallOpts)
}

// BurnedFeesReportingRewardPercentage is a free data retrieval call binding the contract method 0x3a23dfe2.
//
// Solidity: function burnedFeesReportingRewardPercentage() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BurnedFeesReportingRewardPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "burnedFeesReportingRewardPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnedFeesReportingRewardPercentage is a free data retrieval call binding the contract method 0x3a23dfe2.
//
// Solidity: function burnedFeesReportingRewardPercentage() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) BurnedFeesReportingRewardPercentage() (*big.Int, error) {
	return _NativeTokenDestination.Contract.BurnedFeesReportingRewardPercentage(&_NativeTokenDestination.CallOpts)
}

// BurnedFeesReportingRewardPercentage is a free data retrieval call binding the contract method 0x3a23dfe2.
//
// Solidity: function burnedFeesReportingRewardPercentage() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BurnedFeesReportingRewardPercentage() (*big.Int, error) {
	return _NativeTokenDestination.Contract.BurnedFeesReportingRewardPercentage(&_NativeTokenDestination.CallOpts)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) CalculateNumWords(opts *bind.CallOpts, payloadSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "calculateNumWords", payloadSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _NativeTokenDestination.Contract.CalculateNumWords(&_NativeTokenDestination.CallOpts, payloadSize)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _NativeTokenDestination.Contract.CalculateNumWords(&_NativeTokenDestination.CallOpts, payloadSize)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenDestination *NativeTokenDestinationCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenDestination *NativeTokenDestinationSession) Decimals() (uint8, error) {
	return _NativeTokenDestination.Contract.Decimals(&_NativeTokenDestination.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) Decimals() (uint8, error) {
	return _NativeTokenDestination.Contract.Decimals(&_NativeTokenDestination.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenDestination.Contract.GetMinTeleporterVersion(&_NativeTokenDestination.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenDestination.Contract.GetMinTeleporterVersion(&_NativeTokenDestination.CallOpts)
}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) InitialReserveImbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "initialReserveImbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) InitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenDestination.Contract.InitialReserveImbalance(&_NativeTokenDestination.CallOpts)
}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) InitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenDestination.Contract.InitialReserveImbalance(&_NativeTokenDestination.CallOpts)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCaller) IsCollateralized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "isCollateralized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) IsCollateralized() (bool, error) {
	return _NativeTokenDestination.Contract.IsCollateralized(&_NativeTokenDestination.CallOpts)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) IsCollateralized() (bool, error) {
	return _NativeTokenDestination.Contract.IsCollateralized(&_NativeTokenDestination.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0x22366844.
//
// Solidity: function isRegistered() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCaller) IsRegistered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "isRegistered")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0x22366844.
//
// Solidity: function isRegistered() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) IsRegistered() (bool, error) {
	return _NativeTokenDestination.Contract.IsRegistered(&_NativeTokenDestination.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0x22366844.
//
// Solidity: function isRegistered() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) IsRegistered() (bool, error) {
	return _NativeTokenDestination.Contract.IsRegistered(&_NativeTokenDestination.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenDestination.Contract.IsTeleporterAddressPaused(&_NativeTokenDestination.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenDestination.Contract.IsTeleporterAddressPaused(&_NativeTokenDestination.CallOpts, teleporterAddress)
}

// LastestBurnedFeesReported is a free data retrieval call binding the contract method 0xd10a5b8c.
//
// Solidity: function lastestBurnedFeesReported() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) LastestBurnedFeesReported(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "lastestBurnedFeesReported")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastestBurnedFeesReported is a free data retrieval call binding the contract method 0xd10a5b8c.
//
// Solidity: function lastestBurnedFeesReported() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) LastestBurnedFeesReported() (*big.Int, error) {
	return _NativeTokenDestination.Contract.LastestBurnedFeesReported(&_NativeTokenDestination.CallOpts)
}

// LastestBurnedFeesReported is a free data retrieval call binding the contract method 0xd10a5b8c.
//
// Solidity: function lastestBurnedFeesReported() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) LastestBurnedFeesReported() (*big.Int, error) {
	return _NativeTokenDestination.Contract.LastestBurnedFeesReported(&_NativeTokenDestination.CallOpts)
}

// MultiplyOnDestination is a free data retrieval call binding the contract method 0x2001eff6.
//
// Solidity: function multiplyOnDestination() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCaller) MultiplyOnDestination(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "multiplyOnDestination")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MultiplyOnDestination is a free data retrieval call binding the contract method 0x2001eff6.
//
// Solidity: function multiplyOnDestination() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) MultiplyOnDestination() (bool, error) {
	return _NativeTokenDestination.Contract.MultiplyOnDestination(&_NativeTokenDestination.CallOpts)
}

// MultiplyOnDestination is a free data retrieval call binding the contract method 0x2001eff6.
//
// Solidity: function multiplyOnDestination() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) MultiplyOnDestination() (bool, error) {
	return _NativeTokenDestination.Contract.MultiplyOnDestination(&_NativeTokenDestination.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationSession) Name() (string, error) {
	return _NativeTokenDestination.Contract.Name(&_NativeTokenDestination.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) Name() (string, error) {
	return _NativeTokenDestination.Contract.Name(&_NativeTokenDestination.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) Owner() (common.Address, error) {
	return _NativeTokenDestination.Contract.Owner(&_NativeTokenDestination.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) Owner() (common.Address, error) {
	return _NativeTokenDestination.Contract.Owner(&_NativeTokenDestination.CallOpts)
}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCaller) SourceBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "sourceBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationSession) SourceBlockchainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.SourceBlockchainID(&_NativeTokenDestination.CallOpts)
}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) SourceBlockchainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.SourceBlockchainID(&_NativeTokenDestination.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationSession) Symbol() (string, error) {
	return _NativeTokenDestination.Contract.Symbol(&_NativeTokenDestination.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) Symbol() (string, error) {
	return _NativeTokenDestination.Contract.Symbol(&_NativeTokenDestination.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TeleporterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "teleporterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) TeleporterRegistry() (common.Address, error) {
	return _NativeTokenDestination.Contract.TeleporterRegistry(&_NativeTokenDestination.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TeleporterRegistry() (common.Address, error) {
	return _NativeTokenDestination.Contract.TeleporterRegistry(&_NativeTokenDestination.CallOpts)
}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TokenMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "tokenMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TokenMultiplier() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TokenMultiplier(&_NativeTokenDestination.CallOpts)
}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TokenMultiplier() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TokenMultiplier(&_NativeTokenDestination.CallOpts)
}

// TokenSourceAddress is a free data retrieval call binding the contract method 0xf5ea0603.
//
// Solidity: function tokenSourceAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TokenSourceAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "tokenSourceAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenSourceAddress is a free data retrieval call binding the contract method 0xf5ea0603.
//
// Solidity: function tokenSourceAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) TokenSourceAddress() (common.Address, error) {
	return _NativeTokenDestination.Contract.TokenSourceAddress(&_NativeTokenDestination.CallOpts)
}

// TokenSourceAddress is a free data retrieval call binding the contract method 0xf5ea0603.
//
// Solidity: function tokenSourceAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TokenSourceAddress() (common.Address, error) {
	return _NativeTokenDestination.Contract.TokenSourceAddress(&_NativeTokenDestination.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TotalMinted() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalMinted(&_NativeTokenDestination.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TotalMinted() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalMinted(&_NativeTokenDestination.CallOpts)
}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TotalNativeAssetSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "totalNativeAssetSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TotalNativeAssetSupply() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalNativeAssetSupply(&_NativeTokenDestination.CallOpts)
}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TotalNativeAssetSupply() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalNativeAssetSupply(&_NativeTokenDestination.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalSupply(&_NativeTokenDestination.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalSupply(&_NativeTokenDestination.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Approve(&_NativeTokenDestination.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Approve(&_NativeTokenDestination.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.DecreaseAllowance(&_NativeTokenDestination.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.DecreaseAllowance(&_NativeTokenDestination.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) Deposit() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Deposit(&_NativeTokenDestination.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Deposit() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Deposit(&_NativeTokenDestination.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.IncreaseAllowance(&_NativeTokenDestination.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.IncreaseAllowance(&_NativeTokenDestination.TransactOpts, spender, addedValue)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.PauseTeleporterAddress(&_NativeTokenDestination.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.PauseTeleporterAddress(&_NativeTokenDestination.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReceiveTeleporterMessage(&_NativeTokenDestination.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReceiveTeleporterMessage(&_NativeTokenDestination.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RegisterWithSource is a paid mutator transaction binding the contract method 0x151637f6.
//
// Solidity: function registerWithSource((address,uint256) feeInfo) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) RegisterWithSource(opts *bind.TransactOpts, feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "registerWithSource", feeInfo)
}

// RegisterWithSource is a paid mutator transaction binding the contract method 0x151637f6.
//
// Solidity: function registerWithSource((address,uint256) feeInfo) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) RegisterWithSource(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.RegisterWithSource(&_NativeTokenDestination.TransactOpts, feeInfo)
}

// RegisterWithSource is a paid mutator transaction binding the contract method 0x151637f6.
//
// Solidity: function registerWithSource((address,uint256) feeInfo) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) RegisterWithSource(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.RegisterWithSource(&_NativeTokenDestination.TransactOpts, feeInfo)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.RenounceOwnership(&_NativeTokenDestination.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.RenounceOwnership(&_NativeTokenDestination.TransactOpts)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) ReportBurnedTxFees(opts *bind.TransactOpts, requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "reportBurnedTxFees", requiredGasLimit)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) ReportBurnedTxFees(requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReportBurnedTxFees(&_NativeTokenDestination.TransactOpts, requiredGasLimit)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) ReportBurnedTxFees(requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReportBurnedTxFees(&_NativeTokenDestination.TransactOpts, requiredGasLimit)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Send(opts *bind.TransactOpts, input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Send(&_NativeTokenDestination.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Send(&_NativeTokenDestination.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) SendAndCall(opts *bind.TransactOpts, input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "sendAndCall", input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.SendAndCall(&_NativeTokenDestination.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.SendAndCall(&_NativeTokenDestination.TransactOpts, input)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Transfer(&_NativeTokenDestination.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Transfer(&_NativeTokenDestination.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferFrom(&_NativeTokenDestination.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferFrom(&_NativeTokenDestination.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferOwnership(&_NativeTokenDestination.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferOwnership(&_NativeTokenDestination.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.UnpauseTeleporterAddress(&_NativeTokenDestination.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.UnpauseTeleporterAddress(&_NativeTokenDestination.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.UpdateMinTeleporterVersion(&_NativeTokenDestination.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.UpdateMinTeleporterVersion(&_NativeTokenDestination.TransactOpts, version)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Withdraw(&_NativeTokenDestination.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Withdraw(&_NativeTokenDestination.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Fallback(&_NativeTokenDestination.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Fallback(&_NativeTokenDestination.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) Receive() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Receive(&_NativeTokenDestination.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Receive() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Receive(&_NativeTokenDestination.TransactOpts)
}

// NativeTokenDestinationApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NativeTokenDestination contract.
type NativeTokenDestinationApprovalIterator struct {
	Event *NativeTokenDestinationApproval // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationApproval)
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
		it.Event = new(NativeTokenDestinationApproval)
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
func (it *NativeTokenDestinationApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationApproval represents a Approval event raised by the NativeTokenDestination contract.
type NativeTokenDestinationApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NativeTokenDestinationApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationApprovalIterator{contract: _NativeTokenDestination.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationApproval)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseApproval(log types.Log) (*NativeTokenDestinationApproval, error) {
	event := new(NativeTokenDestinationApproval)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the NativeTokenDestination contract.
type NativeTokenDestinationCallFailedIterator struct {
	Event *NativeTokenDestinationCallFailed // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationCallFailed)
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
		it.Event = new(NativeTokenDestinationCallFailed)
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
func (it *NativeTokenDestinationCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationCallFailed represents a CallFailed event raised by the NativeTokenDestination contract.
type NativeTokenDestinationCallFailed struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterCallFailed(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenDestinationCallFailedIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationCallFailedIterator{contract: _NativeTokenDestination.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationCallFailed, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationCallFailed)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "CallFailed", log); err != nil {
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
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseCallFailed(log types.Log) (*NativeTokenDestinationCallFailed, error) {
	event := new(NativeTokenDestinationCallFailed)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the NativeTokenDestination contract.
type NativeTokenDestinationCallSucceededIterator struct {
	Event *NativeTokenDestinationCallSucceeded // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationCallSucceeded)
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
		it.Event = new(NativeTokenDestinationCallSucceeded)
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
func (it *NativeTokenDestinationCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationCallSucceeded represents a CallSucceeded event raised by the NativeTokenDestination contract.
type NativeTokenDestinationCallSucceeded struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterCallSucceeded(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenDestinationCallSucceededIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationCallSucceededIterator{contract: _NativeTokenDestination.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationCallSucceeded, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationCallSucceeded)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseCallSucceeded(log types.Log) (*NativeTokenDestinationCallSucceeded, error) {
	event := new(NativeTokenDestinationCallSucceeded)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the NativeTokenDestination contract.
type NativeTokenDestinationDepositIterator struct {
	Event *NativeTokenDestinationDeposit // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationDeposit)
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
		it.Event = new(NativeTokenDestinationDeposit)
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
func (it *NativeTokenDestinationDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationDeposit represents a Deposit event raised by the NativeTokenDestination contract.
type NativeTokenDestinationDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenDestinationDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationDepositIterator{contract: _NativeTokenDestination.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationDeposit)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseDeposit(log types.Log) (*NativeTokenDestinationDeposit, error) {
	event := new(NativeTokenDestinationDeposit)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the NativeTokenDestination contract.
type NativeTokenDestinationMinTeleporterVersionUpdatedIterator struct {
	Event *NativeTokenDestinationMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationMinTeleporterVersionUpdated)
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
		it.Event = new(NativeTokenDestinationMinTeleporterVersionUpdated)
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
func (it *NativeTokenDestinationMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the NativeTokenDestination contract.
type NativeTokenDestinationMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*NativeTokenDestinationMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationMinTeleporterVersionUpdatedIterator{contract: _NativeTokenDestination.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationMinTeleporterVersionUpdated)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*NativeTokenDestinationMinTeleporterVersionUpdated, error) {
	event := new(NativeTokenDestinationMinTeleporterVersionUpdated)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NativeTokenDestination contract.
type NativeTokenDestinationOwnershipTransferredIterator struct {
	Event *NativeTokenDestinationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationOwnershipTransferred)
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
		it.Event = new(NativeTokenDestinationOwnershipTransferred)
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
func (it *NativeTokenDestinationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationOwnershipTransferred represents a OwnershipTransferred event raised by the NativeTokenDestination contract.
type NativeTokenDestinationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NativeTokenDestinationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationOwnershipTransferredIterator{contract: _NativeTokenDestination.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationOwnershipTransferred)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseOwnershipTransferred(log types.Log) (*NativeTokenDestinationOwnershipTransferred, error) {
	event := new(NativeTokenDestinationOwnershipTransferred)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationReportBurnedTxFeesIterator is returned from FilterReportBurnedTxFees and is used to iterate over the raw logs and unpacked data for ReportBurnedTxFees events raised by the NativeTokenDestination contract.
type NativeTokenDestinationReportBurnedTxFeesIterator struct {
	Event *NativeTokenDestinationReportBurnedTxFees // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationReportBurnedTxFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationReportBurnedTxFees)
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
		it.Event = new(NativeTokenDestinationReportBurnedTxFees)
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
func (it *NativeTokenDestinationReportBurnedTxFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationReportBurnedTxFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationReportBurnedTxFees represents a ReportBurnedTxFees event raised by the NativeTokenDestination contract.
type NativeTokenDestinationReportBurnedTxFees struct {
	TeleporterMessageID [32]byte
	FeesBurned          *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReportBurnedTxFees is a free log retrieval operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterReportBurnedTxFees(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*NativeTokenDestinationReportBurnedTxFeesIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "ReportBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationReportBurnedTxFeesIterator{contract: _NativeTokenDestination.contract, event: "ReportBurnedTxFees", logs: logs, sub: sub}, nil
}

// WatchReportBurnedTxFees is a free log subscription operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchReportBurnedTxFees(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationReportBurnedTxFees, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "ReportBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationReportBurnedTxFees)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "ReportBurnedTxFees", log); err != nil {
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
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseReportBurnedTxFees(log types.Log) (*NativeTokenDestinationReportBurnedTxFees, error) {
	event := new(NativeTokenDestinationReportBurnedTxFees)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "ReportBurnedTxFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTeleporterAddressPausedIterator struct {
	Event *NativeTokenDestinationTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTeleporterAddressPaused)
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
		it.Event = new(NativeTokenDestinationTeleporterAddressPaused)
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
func (it *NativeTokenDestinationTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenDestinationTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTeleporterAddressPausedIterator{contract: _NativeTokenDestination.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTeleporterAddressPaused)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTeleporterAddressPaused(log types.Log) (*NativeTokenDestinationTeleporterAddressPaused, error) {
	event := new(NativeTokenDestinationTeleporterAddressPaused)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTeleporterAddressUnpausedIterator struct {
	Event *NativeTokenDestinationTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTeleporterAddressUnpaused)
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
		it.Event = new(NativeTokenDestinationTeleporterAddressUnpaused)
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
func (it *NativeTokenDestinationTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenDestinationTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTeleporterAddressUnpausedIterator{contract: _NativeTokenDestination.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTeleporterAddressUnpaused)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*NativeTokenDestinationTeleporterAddressUnpaused, error) {
	event := new(NativeTokenDestinationTeleporterAddressUnpaused)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTokensAndCallSentIterator is returned from FilterTokensAndCallSent and is used to iterate over the raw logs and unpacked data for TokensAndCallSent events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensAndCallSentIterator struct {
	Event *NativeTokenDestinationTokensAndCallSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationTokensAndCallSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTokensAndCallSent)
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
		it.Event = new(NativeTokenDestinationTokensAndCallSent)
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
func (it *NativeTokenDestinationTokensAndCallSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTokensAndCallSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTokensAndCallSent represents a TokensAndCallSent event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensAndCallSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTokensAndCallSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenDestinationTokensAndCallSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTokensAndCallSentIterator{contract: _NativeTokenDestination.contract, event: "TokensAndCallSent", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTokensAndCallSent(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTokensAndCallSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTokensAndCallSent)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
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
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTokensAndCallSent(log types.Log) (*NativeTokenDestinationTokensAndCallSent, error) {
	event := new(NativeTokenDestinationTokensAndCallSent)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTokensSentIterator is returned from FilterTokensSent and is used to iterate over the raw logs and unpacked data for TokensSent events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensSentIterator struct {
	Event *NativeTokenDestinationTokensSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTokensSent)
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
		it.Event = new(NativeTokenDestinationTokensSent)
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
func (it *NativeTokenDestinationTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTokensSent represents a TokensSent event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensSent is a free log retrieval operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenDestinationTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTokensSentIterator{contract: _NativeTokenDestination.contract, event: "TokensSent", logs: logs, sub: sub}, nil
}

// WatchTokensSent is a free log subscription operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTokensSent(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTokensSent)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensSent", log); err != nil {
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
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTokensSent(log types.Log) (*NativeTokenDestinationTokensSent, error) {
	event := new(NativeTokenDestinationTokensSent)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensWithdrawnIterator struct {
	Event *NativeTokenDestinationTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTokensWithdrawn)
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
		it.Event = new(NativeTokenDestinationTokensWithdrawn)
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
func (it *NativeTokenDestinationTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTokensWithdrawn represents a TokensWithdrawn event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*NativeTokenDestinationTokensWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTokensWithdrawnIterator{contract: _NativeTokenDestination.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTokensWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTokensWithdrawn)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTokensWithdrawn(log types.Log) (*NativeTokenDestinationTokensWithdrawn, error) {
	event := new(NativeTokenDestinationTokensWithdrawn)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTransferIterator struct {
	Event *NativeTokenDestinationTransfer // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTransfer)
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
		it.Event = new(NativeTokenDestinationTransfer)
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
func (it *NativeTokenDestinationTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTransfer represents a Transfer event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTokenDestinationTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTransferIterator{contract: _NativeTokenDestination.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTransfer)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTransfer(log types.Log) (*NativeTokenDestinationTransfer, error) {
	event := new(NativeTokenDestinationTransfer)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the NativeTokenDestination contract.
type NativeTokenDestinationWithdrawalIterator struct {
	Event *NativeTokenDestinationWithdrawal // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationWithdrawal)
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
		it.Event = new(NativeTokenDestinationWithdrawal)
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
func (it *NativeTokenDestinationWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationWithdrawal represents a Withdrawal event raised by the NativeTokenDestination contract.
type NativeTokenDestinationWithdrawal struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterWithdrawal(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenDestinationWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationWithdrawalIterator{contract: _NativeTokenDestination.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationWithdrawal, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationWithdrawal)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseWithdrawal(log types.Log) (*NativeTokenDestinationWithdrawal, error) {
	event := new(NativeTokenDestinationWithdrawal)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
