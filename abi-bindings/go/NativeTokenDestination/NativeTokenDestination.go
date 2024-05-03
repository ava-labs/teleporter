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
	MultiplyOnReceive                   bool
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
	FallbackRecipient        common.Address
	PrimaryFee               *big.Int
	SecondaryFee             *big.Int
}

// SendTokensInput is an auto generated low-level Go binding around an user-defined struct.
type SendTokensInput struct {
	DestinationBlockchainID  [32]byte
	DestinationBridgeAddress common.Address
	Recipient                common.Address
	PrimaryFee               *big.Int
	SecondaryFee             *big.Int
	RequiredGasLimit         *big.Int
}

// NativeTokenDestinationMetaData contains all meta data concerning the NativeTokenDestination contract.
var NativeTokenDestinationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"nativeAssetSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenSourceAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialReserveImbalance\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimalsShift\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"multiplyOnReceive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"burnedFeesReportingRewardPercentage\",\"type\":\"uint256\"}],\"internalType\":\"structNativeTokenDestinationSettings\",\"name\":\"settings\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"name\":\"CollateralAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesBurned\",\"type\":\"uint256\"}],\"name\":\"ReportBurnedTxFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallRouted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensRouted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"BURNED_FOR_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURNED_TX_FEES_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_CALL_GAS_PER_WORD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_MINTER\",\"outputs\":[{\"internalType\":\"contractINativeMinter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_BURN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnedFeesReportingRewardPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"payloadSize\",\"type\":\"uint256\"}],\"name\":\"calculateNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentReserveImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialReserveImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastestBurnedFeesReported\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplyOnReceive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"name\":\"reportBurnedTxFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isReceive\",\"type\":\"bool\"}],\"name\":\"scaleTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"sendAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenSourceAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalNativeAssetSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101806040523480156200001257600080fd5b5060405162004c7538038062004c7583398101604081905262000035916200079a565b80602001518160400151826060015183608001518460c001518560e0015185858189600001516040516020016200006d91906200088f565b60408051601f198184030181529190528a5160036200008d838262000950565b5060046200009c828262000950565b50506001600555506001600160a01b038116620001265760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084015b60405180910390fd5b6001600160a01b03811660808190526040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa15801562000171573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000197919062000a1c565b60075550620001a63362000532565b620001b18162000584565b505060016009819055507302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200020e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000234919062000a1c565b60a052836200029b5760405162461bcd60e51b8152602060048201526035602482015260008051602062004c5583398151915260448201527f20736f7572636520626c6f636b636861696e204944000000000000000000000060648201526084016200011d565b60a0518403620003235760405162461bcd60e51b815260206004820152604660248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a2063616e6e60448201527f6f74206465706c6f7920746f2073616d6520626c6f636b636861696e20617320606482015265736f7572636560d01b608482015260a4016200011d565b6001600160a01b038316620003905760405162461bcd60e51b8152602060048201526035602482015260008051602062004c5583398151915260448201527f20746f6b656e20736f757263652061646472657373000000000000000000000060648201526084016200011d565b60128260ff161115620004005760405162461bcd60e51b815260206004820152603160248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a20696e76616044820152701b1a5908191958da5b585b1cd4da1a599d607a1b60648201526084016200011d565b60c08490526001600160a01b03831660e0526200041f82600a62000b4b565b610100521515610120525050505060a08201516000039050620004ab5760405162461bcd60e51b815260206004820152603660248201527f4e6174697665546f6b656e44657374696e6174696f6e3a207a65726f20696e6960448201527f7469616c207265736572766520696d62616c616e63650000000000000000000060648201526084016200011d565b60a0810180516101405251600a55610100810151606411620005235760405162461bcd60e51b815260206004820152602a60248201527f4e6174697665546f6b656e44657374696e6174696f6e3a20696e76616c69642060448201526970657263656e7461676560b01b60648201526084016200011d565b61010001516101605262000b63565b600880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6200058e62000603565b6001600160a01b038116620005f55760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016200011d565b620006008162000532565b50565b6008546001600160a01b031633146200065f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016200011d565b565b634e487b7160e01b600052604160045260246000fd5b60405161012081016001600160401b03811182821017156200069d576200069d62000661565b60405290565b60005b83811015620006c0578181015183820152602001620006a6565b50506000910152565b600082601f830112620006db57600080fd5b81516001600160401b0380821115620006f857620006f862000661565b604051601f8301601f19908116603f0116810190828211818310171562000723576200072362000661565b816040528381528660208588010111156200073d57600080fd5b62000750846020830160208901620006a3565b9695505050505050565b80516001600160a01b03811681146200077257600080fd5b919050565b805160ff811681146200077257600080fd5b805180151581146200077257600080fd5b600060208284031215620007ad57600080fd5b81516001600160401b0380821115620007c557600080fd5b908301906101208286031215620007db57600080fd5b620007e562000677565b825182811115620007f557600080fd5b6200080387828601620006c9565b82525062000814602084016200075a565b602082015262000827604084016200075a565b60408201526060830151606082015262000844608084016200075a565b608082015260a083015160a08201526200086160c0840162000777565b60c08201526200087460e0840162000789565b60e08201526101009283015192810192909252509392505050565b6702bb930b83832b2160c51b815260008251620008b4816008850160208701620006a3565b9190910160080192915050565b600181811c90821680620008d657607f821691505b602082108103620008f757634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200094b57600081815260208120601f850160051c81016020861015620009265750805b601f850160051c820191505b81811015620009475782815560010162000932565b5050505b505050565b81516001600160401b038111156200096c576200096c62000661565b62000984816200097d8454620008c1565b84620008fd565b602080601f831160018114620009bc5760008415620009a35750858301515b600019600386901b1c1916600185901b17855562000947565b600085815260208120601f198616915b82811015620009ed57888601518255948401946001909101908401620009cc565b508582101562000a0c5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60006020828403121562000a2f57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600181815b8085111562000a8d57816000190482111562000a715762000a7162000a36565b8085161562000a7f57918102915b93841c939080029062000a51565b509250929050565b60008262000aa65750600162000b45565b8162000ab55750600062000b45565b816001811462000ace576002811462000ad95762000af9565b600191505062000b45565b60ff84111562000aed5762000aed62000a36565b50506001821b62000b45565b5060208310610133831016604e8410600b841016171562000b1e575081810a62000b45565b62000b2a838362000a4c565b806000190482111562000b415762000b4162000a36565b0290505b92915050565b600062000b5c60ff84168362000a95565b9392505050565b60805160a05160c05160e05161010051610120516101405161016051613feb62000c6a600039600081816104f50152610dc501526000818161061d01526109c50152600081816103c301526110b5015260008181610733015281816110e0015261111101526000818161085e01528181610f1e015281816116e20152818161198801528181612043015281816122d3015261285301526000818161041701528181610efb015281816116b2015281816119620152818161201c015281816122ad01526127dd0152600081816107b501528181611820015281816120bc0152612154015260008181610377015281816111520152818161253a0152612bf00152613feb6000f3fe6080604052600436106102805760003560e01c80635eb995141161014f578063ba3f5a12116100c1578063d2cc7a701161007a578063d2cc7a70146107d7578063dd62ed3e146107ec578063ecd4ed1b14610537578063f2fde38b1461080c578063f3f981d81461082c578063f5ea06031461084c5761028f565b8063ba3f5a1214610721578063c452165e14610755578063c868efaa1461076d578063d0e30db01461028f578063d10a5b8c1461078d578063d127dc9b146107a35761028f565b806395d89b411161011357806395d89b411461065d5780639731429714610672578063a2309ff8146106ab578063a457c2d7146106c1578063a9059cbb146106e1578063b9448587146107015761028f565b80635eb99514146105a057806370a08231146105c0578063715018a6146105f65780638ac7dd201461060b5780638da5cb5b1461063f5761028f565b80632b0d8f18116101f35780633a23dfe2116101ac5780633a23dfe2146104e35780634511243e1461051757806347a9a22c1461053757806349e3284e14610554578063525975e61461056957806355538c8b146105805761028f565b80632b0d8f18146104395780632e1a7d4d1461045957806330079bff14610479578063313ce5671461048c578063329c3e12146104a857806339509351146104c35761028f565b806318160ddd1161024557806318160ddd1461033b5780631906529c146103505780631a7f5bec146103655780631ce22075146103b157806323b872dd146103e557806329b7b3fd146104055761028f565b8062d872ae1461029757806306fdde03146102c0578063095ea7b3146102e2578063146dfd3c1461031257806315beb59f146103255761028f565b3661028f5761028d610880565b005b61028d610880565b3480156102a357600080fd5b506102ad600a5481565b6040519081526020015b60405180910390f35b3480156102cc57600080fd5b506102d56108c1565b6040516102b7919061338f565b3480156102ee57600080fd5b506103026102fd3660046133c2565b610953565b60405190151581526020016102b7565b61028d6103203660046133ee565b61096d565b34801561033157600080fd5b506102ad61213481565b34801561034757600080fd5b506002546102ad565b34801561035c57600080fd5b506102ad6109a3565b34801561037157600080fd5b506103997f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016102b7565b3480156103bd57600080fd5b506103027f000000000000000000000000000000000000000000000000000000000000000081565b3480156103f157600080fd5b50610302610400366004613406565b610a04565b34801561041157600080fd5b506102ad7f000000000000000000000000000000000000000000000000000000000000000081565b34801561044557600080fd5b5061028d610454366004613447565b610a28565b34801561046557600080fd5b5061028d610474366004613464565b610b21565b61028d61048736600461347d565b610b91565b34801561049857600080fd5b50604051601281526020016102b7565b3480156104b457600080fd5b506103996001600160991b0181565b3480156104cf57600080fd5b506103026104de3660046133c2565b610bc3565b3480156104ef57600080fd5b506102ad7f000000000000000000000000000000000000000000000000000000000000000081565b34801561052357600080fd5b5061028d610532366004613447565b610be5565b34801561054357600080fd5b5061039962010203600160981b0181565b34801561056057600080fd5b50610302610ce2565b34801561057557600080fd5b506102ad62035b6081565b34801561058c57600080fd5b5061028d61059b366004613464565b610cf3565b3480156105ac57600080fd5b5061028d6105bb366004613464565b610ff3565b3480156105cc57600080fd5b506102ad6105db366004613447565b6001600160a01b031660009081526020819052604090205490565b34801561060257600080fd5b5061028d611004565b34801561061757600080fd5b506102ad7f000000000000000000000000000000000000000000000000000000000000000081565b34801561064b57600080fd5b506008546001600160a01b0316610399565b34801561066957600080fd5b506102d5611016565b34801561067e57600080fd5b5061030261068d366004613447565b6001600160a01b031660009081526006602052604090205460ff1690565b3480156106b757600080fd5b506102ad600b5481565b3480156106cd57600080fd5b506103026106dc3660046133c2565b611025565b3480156106ed57600080fd5b506103026106fc3660046133c2565b6110a0565b34801561070d57600080fd5b506102ad61071c3660046134c7565b6110ae565b34801561072d57600080fd5b506102ad7f000000000000000000000000000000000000000000000000000000000000000081565b34801561076157600080fd5b50610399600160981b81565b34801561077957600080fd5b5061028d6107883660046134f7565b61113d565b34801561079957600080fd5b506102ad600c5481565b3480156107af57600080fd5b506102ad7f000000000000000000000000000000000000000000000000000000000000000081565b3480156107e357600080fd5b506007546102ad565b3480156107f857600080fd5b506102ad610807366004613580565b611307565b34801561081857600080fd5b5061028d610827366004613447565b611332565b34801561083857600080fd5b506102ad610847366004613464565b6113a8565b34801561085857600080fd5b506103997f000000000000000000000000000000000000000000000000000000000000000081565b60405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a26108bf33346113bf565b565b6060600380546108d0906135ae565b80601f01602080910402602001604051908101604052809291908181526020018280546108fc906135ae565b80156109495780601f1061091e57610100808354040283529160200191610949565b820191906000526020600020905b81548152906001019060200180831161092c57829003601f168201915b5050505050905090565b60003361096181858561147e565b60019150505b92915050565b600a54156109965760405162461bcd60e51b815260040161098d906135e2565b60405180910390fd5b6109a081346115a3565b50565b6000806109bf62010203600160981b0131600160981b3161364c565b905060007f0000000000000000000000000000000000000000000000000000000000000000600b546109f1919061364c565b90506109fd828261365f565b9250505090565b600033610a12858285611a8e565b610a1d858585611b02565b506001949350505050565b610a30611ca6565b6001600160a01b038116610a565760405162461bcd60e51b815260040161098d90613672565b6001600160a01b03811660009081526006602052604090205460ff1615610ad55760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b606482015260840161098d565b6001600160a01b038116600081815260066020526040808220805460ff19166001179055517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b60405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a2610b603382611cae565b604051339082156108fc029083906000818181858888f19350505050158015610b8d573d6000803e3d6000fd5b5050565b600a5415610bb15760405162461bcd60e51b815260040161098d906135e2565b6109a0610bbd826137d2565b34611ddd565b600033610961818585610bd68383611307565b610be0919061364c565b61147e565b610bed611ca6565b6001600160a01b038116610c135760405162461bcd60e51b815260040161098d90613672565b6001600160a01b03811660009081526006602052604090205460ff16610c8d5760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b606482015260840161098d565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152600660205260409020805460ff19169055565b6000610cee600a541590565b905090565b600160095414610d155760405162461bcd60e51b815260040161098d90613880565b6002600955600c54600160981b31908111610daa5760405162461bcd60e51b815260206004820152604960248201527f4e6174697665546f6b656e44657374696e6174696f6e3a206275726e2061646460448201527f726573732062616c616e6365206e6f742067726561746572207468616e206c616064820152681cdd081c995c1bdc9d60ba1b608482015260a40161098d565b6000600c5482610dba919061365f565b905060006064610dea7f0000000000000000000000000000000000000000000000000000000000000000846138c4565b610df491906138db565b90506000610e02828461365f565b600c85905590508115610e2457610e193083612385565b610e2282612407565b505b6000610e318260006110ae565b905060008111610ea95760405162461bcd60e51b815260206004820152603960248201527f4e6174697665546f6b656e44657374696e6174696f6e3a207a65726f2073636160448201527f6c656420616d6f756e7420746f207265706f7274206275726e00000000000000606482015260840161098d565b6040805160608082018352600080835260208084018690528451808201865262010203600160981b0190819052855180830191909152855180820383018152908601865284860152845160c0810186527f000000000000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681830152855180870187523081528083018a9052818701529283018b905284518281529081019094529192610fa891906080820190815260200184604051602001610f919190613913565b604051602081830303815290604052815250612417565b9050807f0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a585604051610fdc91815260200190565b60405180910390a250506001600955505050505050565b610ffb611ca6565b6109a081612536565b61100c6126d6565b6108bf6000612730565b6060600480546108d0906135ae565b600033816110338286611307565b9050838110156110935760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b606482015260840161098d565b610a1d828686840361147e565b600033610961818585611b02565b60008115157f000000000000000000000000000000000000000000000000000000000000000015150361110c576111057f0000000000000000000000000000000000000000000000000000000000000000846138c4565b9050610967565b6111367f0000000000000000000000000000000000000000000000000000000000000000846138db565b9392505050565b611145612782565b6007546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa1580156111bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111e09190613962565b10156112475760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b606482015260840161098d565b6112503361068d565b156112b65760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b606482015260840161098d565b6112f7848484848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506127db92505050565b6113016001600555565b50505050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b61133a6126d6565b6001600160a01b03811661139f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161098d565b6109a081612730565b600060056113b783601f61364c565b901c92915050565b6001600160a01b0382166114155760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161098d565b8060026000828254611427919061364c565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6001600160a01b0383166114e05760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b606482015260840161098d565b6001600160a01b0382166115415760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b606482015260840161098d565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6001600954146115c55760405162461bcd60e51b815260040161098d90613880565b600260095560006115dc6060840160408501613447565b6001600160a01b03160361163b5760405162461bcd60e51b81526020600482015260326024820152600080516020613f7683398151915260448201527120726563697069656e74206164647265737360701b606482015260840161098d565b60008260a001351161165f5760405162461bcd60e51b815260040161098d9061397b565b61168482356116746040850160208601613447565b83856060013586608001356129f1565b90506116ab6040805160608101909152806000815260200160008152602001606081525090565b60a08301357f000000000000000000000000000000000000000000000000000000000000000084350361181e576001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166117126040860160208701613447565b6001600160a01b0316146117385760405162461bcd60e51b815260040161098d906139bc565b6080840135156117a55760405162461bcd60e51b815260206004820152603260248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a206e6f6e2d6044820152717a65726f207365636f6e646172792066656560701b606482015260840161098d565b6040805160608101909152806000815260200184815260200160405180602001604052808760400160208101906117dc9190613447565b6001600160a01b03169052604051611805919060200190516001600160a01b0316815260200190565b6040516020818303038152906040528152509150611950565b7f000000000000000000000000000000000000000000000000000000000000000084350361187d57306118576040860160208701613447565b6001600160a01b03160361187d5760405162461bcd60e51b815260040161098d906139bc565b604080516060810190915280600281526020018481526020016040518060a00160405280876000013581526020018760200160208101906118be9190613447565b6001600160a01b031681526020016118dc6060890160408a01613447565b6001600160a01b03908116825260808981013560208085019190915260a0808c01356040958601528451865181840152918601518416828601528585015190931660608083019190915285015181830152930151838201528151808403909101815260c090920190529052915062035b6090505b6000611a3c6040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020016040518060400160405280306001600160a01b0316815260200189606001358152508152602001848152602001600067ffffffffffffffff8111156119fc576119fc6136c0565b604051908082528060200260200182016040528015611a25578160200160208202803683370190505b50815260200185604051602001610f919190613913565b9050336001600160a01b0316817f78488d924de07bf96852578ad434a6c920f0835e97f9b302a77e1a77757c640b8787604051611a7a929190613a07565b60405180910390a350506001600955505050565b6000611a9a8484611307565b905060001981146113015781811015611af55760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161098d565b611301848484840361147e565b6001600160a01b038316611b665760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b606482015260840161098d565b6001600160a01b038216611bc85760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b606482015260840161098d565b6001600160a01b03831660009081526020819052604090205481811015611c405760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b606482015260840161098d565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3611301565b6108bf6126d6565b6001600160a01b038216611d0e5760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b606482015260840161098d565b6001600160a01b03821660009081526020819052604090205481811015611d825760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b606482015260840161098d565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101611596565b505050565b600160095414611dff5760405162461bcd60e51b815260040161098d90613880565b600260095560408201516001600160a01b0316611e725760405162461bcd60e51b815260206004820152603b6024820152600080516020613f7683398151915260448201527f20726563697069656e7420636f6e747261637420616464726573730000000000606482015260840161098d565b6000826080015111611e965760405162461bcd60e51b815260040161098d9061397b565b60008260a0015111611ef55760405162461bcd60e51b81526020600482015260346024820152600080516020613f76833981519152604482015273081c9958da5c1a595b9d0819d85cc81b1a5b5a5d60621b606482015260840161098d565b81608001518260a0015110611f605760405162461bcd60e51b81526020600482015260376024820152600080516020613f9683398151915260448201527f6c696420726563697069656e7420676173206c696d6974000000000000000000606482015260840161098d565b60c08201516001600160a01b0316611fce5760405162461bcd60e51b815260206004820152603b6024820152600080516020613f7683398151915260448201527f2066616c6c6261636b20726563697069656e7420616464726573730000000000606482015260840161098d565b611fec82600001518360200151838560e001518661010001516129f1565b90506120136040805160608101909152806000815260200160008152602001606081525090565b608083015183517f00000000000000000000000000000000000000000000000000000000000000009003612150577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031684602001516001600160a01b0316146120965760405162461bcd60e51b815260040161098d906139bc565b604080516060810190915280600181526020018481526020016040518060c001604052807f00000000000000000000000000000000000000000000000000000000000000008152602001336001600160a01b0316815260200187604001516001600160a01b03168152602001876060015181526020018760a0015181526020018760c001516001600160a01b03168152506040516020016121379190613a6a565b604051602081830303815290604052815250915061229b565b83517f000000000000000000000000000000000000000000000000000000000000000090036121ae57306001600160a01b031684602001516001600160a01b0316036121ae5760405162461bcd60e51b815260040161098d906139bc565b60408051606081019091528060038152602001848152602001604051806101200160405280336001600160a01b031681526020018760000151815260200187602001516001600160a01b0316815260200187604001516001600160a01b03168152602001876060015181526020018760a0015181526020018760c001516001600160a01b031681526020018760800151815260200187610100015181525060405160200161225c9190613ad5565b60405160208183030381529060405281525091506121346122818560600151516113a8565b61228b91906138c4565b6122989062035b6061364c565b90505b60006123476040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020016040518060400160405280306001600160a01b031681526020018960e001518152508152602001848152602001600067ffffffffffffffff8111156119fc576119fc6136c0565b9050336001600160a01b0316817f76b18d78fd0b0c8a046526d2a500e1e5ced780f056df0acc4932088d10e665628787604051611a7a929190613b8e565b80600b6000828254612397919061364c565b90915550506040516327ad555d60e11b81526001600160a01b0383166004820152602481018290526001600160991b0190634f5aaaba90604401600060405180830381600087803b1580156123eb57600080fd5b505af11580156123ff573d6000803e3d6000fd5b505050505050565b600061241330836113bf565b5090565b600080612422612beb565b604084015160200151909150156124c7576040830151516001600160a01b03166124a45760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a207a65726f206665652060448201526c746f6b656e206164647265737360981b606482015260840161098d565b6040830151602081015190516124c7916001600160a01b03909116908390612cff565b604051630624488560e41b81526001600160a01b038216906362448850906124f3908690600401613c7b565b6020604051808303816000875af1158015612512573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111369190613962565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015612596573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125ba9190613962565b6007549091508183111561262a5760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b606482015260840161098d565b80831161269f5760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e00606482015260840161098d565b6007839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b6008546001600160a01b031633146108bf5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161098d565b600880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6002600554036127d45760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161098d565b6002600555565b7f000000000000000000000000000000000000000000000000000000000000000083146128515760405162461bcd60e51b81526020600482015260306024820152600080516020613f9683398151915260448201526f3634b21039b7bab931b29031b430b4b760811b606482015260840161098d565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b0316146128e65760405162461bcd60e51b81526020600482015260386024820152600080516020613f9683398151915260448201527f6c696420746f6b656e20736f7572636520616464726573730000000000000000606482015260840161098d565b6000818060200190518101906128fc9190613d47565b9050600061290f826020015160016110ae565b9050600082516003811115612926576129266138fd565b0361295b57600082604001518060200190518101906129459190613deb565b9050612955816000015183612dd3565b506129ea565b600182516003811115612970576129706138fd565b0361299b576000826040015180602001905181019061298f9190613e37565b90506129558183612ed2565b60405162461bcd60e51b81526020600482015260306024820152600080516020613f9683398151915260448201526f6c6964206d657373616765207479706560801b606482015260840161098d565b5050505050565b600085612a545760405162461bcd60e51b815260206004820152603a6024820152600080516020613f7683398151915260448201527f2064657374696e6174696f6e20626c6f636b636861696e204944000000000000606482015260840161098d565b6001600160a01b038516612abe5760405162461bcd60e51b815260206004820152603b6024820152600080516020613f7683398151915260448201527f2064657374696e6174696f6e2062726964676520616464726573730000000000606482015260840161098d565b612ac784612407565b9350612ad3828461364c565b8411612b475760405162461bcd60e51b815260206004820152603d60248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a20696e737560448201527f6666696369656e7420616d6f756e7420746f20636f7665722066656573000000606482015260840161098d565b612b51838561365f565b9350612b5c84613033565b6000612b698560006110ae565b905060008111612be15760405162461bcd60e51b815260206004820152603b60248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a20696e737560448201527f6666696369656e7420746f6b656e7320746f207472616e736665720000000000606482015260840161098d565b9695505050505050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612c4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c709190613ef5565b9050612c94816001600160a01b031660009081526006602052604090205460ff1690565b15612cfa5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b606482015260840161098d565b919050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa158015612d50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d749190613962565b612d7e919061364c565b604080516001600160a01b038616602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052909150611301908590613073565b600a5481908015612e855780831115612e355760408051828152600060208201527f244160b15e69cc411f041d94ae7fab6f6bba85dade8403216c05ff4b920d5449910160405180910390a1612e29818361365f565b6000600a559150612e85565b6000612e41848361365f565b60408051868152602081018390529192507f244160b15e69cc411f041d94ae7fab6f6bba85dade8403216c05ff4b920d5449910160405180910390a1600a55600091505b836001600160a01b03167f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b83604051612ec091815260200190565b60405180910390a26113018483612385565b600a5415612ee857610b8d8260a0015182612dd3565b612ef23082612385565b6000826000015183602001518460600151604051602401612f1593929190613f12565b60408051601f198184030181529181526020820180516001600160e01b0316630d57eddf60e21b179052608085015190850151919250600091612f5b9190859085613145565b90508015612faf5783604001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff484604051612fa291815260200190565b60405180910390a2611301565b83604001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb084604051612fee91815260200190565b60405180910390a28360a001516001600160a01b03166108fc849081150290604051600060405180830381858888f193505050501580156129ea573d6000803e3d6000fd5b61303d3082611cae565b60405162010203600160981b019082156108fc029083906000818181858888f19350505050158015610b8d573d6000803e3d6000fd5b60006130c8826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661321c9092919063ffffffff16565b805190915015611dd857808060200190518101906130e69190613f3c565b611dd85760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161098d565b6000845a10156131975760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e74206761730000000000604482015260640161098d565b834710156131e75760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c7565000000604482015260640161098d565b826001600160a01b03163b60000361320157506000613214565b600080600084516020860188888bf19150505b949350505050565b6060613214848460008585600080866001600160a01b031685876040516132439190613f59565b60006040518083038185875af1925050503d8060008114613280576040519150601f19603f3d011682016040523d82523d6000602084013e613285565b606091505b5091509150613296878383876132a1565b979650505050505050565b60608315613310578251600003613309576001600160a01b0385163b6133095760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161098d565b5081613214565b61321483838151156133255781518083602001fd5b8060405162461bcd60e51b815260040161098d919061338f565b60005b8381101561335a578181015183820152602001613342565b50506000910152565b6000815180845261337b81602086016020860161333f565b601f01601f19169290920160200192915050565b6020815260006111366020830184613363565b6001600160a01b03811681146109a057600080fd5b8035612cfa816133a2565b600080604083850312156133d557600080fd5b82356133e0816133a2565b946020939093013593505050565b600060c0828403121561340057600080fd5b50919050565b60008060006060848603121561341b57600080fd5b8335613426816133a2565b92506020840135613436816133a2565b929592945050506040919091013590565b60006020828403121561345957600080fd5b8135611136816133a2565b60006020828403121561347657600080fd5b5035919050565b60006020828403121561348f57600080fd5b813567ffffffffffffffff8111156134a657600080fd5b8201610120818503121561113657600080fd5b80151581146109a057600080fd5b600080604083850312156134da57600080fd5b8235915060208301356134ec816134b9565b809150509250929050565b6000806000806060858703121561350d57600080fd5b84359350602085013561351f816133a2565b9250604085013567ffffffffffffffff8082111561353c57600080fd5b818701915087601f83011261355057600080fd5b81358181111561355f57600080fd5b88602082850101111561357157600080fd5b95989497505060200194505050565b6000806040838503121561359357600080fd5b823561359e816133a2565b915060208301356134ec816133a2565b600181811c908216806135c257607f821691505b60208210810361340057634e487b7160e01b600052602260045260246000fd5b60208082526034908201527f4e6174697665546f6b656e44657374696e6174696f6e3a20636f6e7472616374604082015273081d5b99195c98dbdb1b185d195c985b1a5e995960621b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b8082018082111561096757610967613636565b8181038181111561096757610967613636565b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b634e487b7160e01b600052604160045260246000fd5b604051610120810167ffffffffffffffff811182821017156136fa576136fa6136c0565b60405290565b60405160c0810167ffffffffffffffff811182821017156136fa576136fa6136c0565b604051601f8201601f1916810167ffffffffffffffff8111828210171561374c5761374c6136c0565b604052919050565b600067ffffffffffffffff82111561376e5761376e6136c0565b50601f01601f191660200190565b600082601f83011261378d57600080fd5b81356137a061379b82613754565b613723565b8181528460208386010111156137b557600080fd5b816020850160208301376000918101602001919091529392505050565b600061012082360312156137e557600080fd5b6137ed6136d6565b823581526137fd602084016133b7565b602082015261380e604084016133b7565b6040820152606083013567ffffffffffffffff81111561382d57600080fd5b6138393682860161377c565b6060830152506080830135608082015260a083013560a082015261385f60c084016133b7565b60c082015260e0838101359082015261010092830135928101929092525090565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b808202811582820484141761096757610967613636565b6000826138f857634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052602160045260246000fd5b60208152600082516004811061393957634e487b7160e01b600052602160045260246000fd5b806020840152506020830151604083015260408301516060808401526132146080840182613363565b60006020828403121561397457600080fd5b5051919050565b6020808252603390820152600080516020613f76833981519152604082015272081c995c5d5a5c99590819d85cc81b1a5b5a5d606a1b606082015260800190565b6020808252603e90820152600080516020613f9683398151915260408201527f6c69642064657374696e6174696f6e2062726964676520616464726573730000606082015260800190565b8235815260e081016020840135613a1d816133a2565b6001600160a01b039081166020840152604085013590613a3c826133a2565b166040830152606084810135908301526080808501359083015260a0938401359382019390935260c0015290565b60208152815160208201526000602083015160018060a01b0380821660408501528060408601511660608501526060850151915060c06080850152613ab260e0850183613363565b9150608085015160a08501528060a08601511660c0850152508091505092915050565b60208152613aef6020820183516001600160a01b03169052565b6020820151604082015260006040830151613b1560608401826001600160a01b03169052565b5060608301516001600160a01b03811660808401525060808301516101208060a0850152613b47610140850183613363565b915060a085015160c085015260c0850151613b6d60e08601826001600160a01b03169052565b5060e085015161010085810191909152909401519390920192909252919050565b604081528251604082015260006020840151613bb560608401826001600160a01b03169052565b5060408401516001600160a01b03166080830152606084015161012060a08401819052613be6610160850183613363565b9150608086015160c085015260a086015160e085015260c0860151610100613c18818701836001600160a01b03169052565b60e0880151928601929092525090940151610140830152506020015290565b600081518084526020808501945080840160005b83811015613c705781516001600160a01b031687529582019590820190600101613c4b565b509495945050505050565b60208152815160208201526000602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c0840152613cdc610100840182613c37565b905060a0840151601f198483030160e0850152613cf98282613363565b95945050505050565b600082601f830112613d1357600080fd5b8151613d2161379b82613754565b818152846020838601011115613d3657600080fd5b61321482602083016020870161333f565b600060208284031215613d5957600080fd5b815167ffffffffffffffff80821115613d7157600080fd5b9083019060608286031215613d8557600080fd5b604051606081018181108382111715613da057613da06136c0565b604052825160048110613db257600080fd5b815260208381015190820152604083015182811115613dd057600080fd5b613ddc87828601613d02565b60408301525095945050505050565b600060208284031215613dfd57600080fd5b6040516020810181811067ffffffffffffffff82111715613e2057613e206136c0565b6040528251613e2e816133a2565b81529392505050565b600060208284031215613e4957600080fd5b815167ffffffffffffffff80821115613e6157600080fd5b9083019060c08286031215613e7557600080fd5b613e7d613700565b825181526020830151613e8f816133a2565b60208201526040830151613ea2816133a2565b6040820152606083015182811115613eb957600080fd5b613ec587828601613d02565b6060830152506080830151608082015260a08301519250613ee5836133a2565b60a0810192909252509392505050565b600060208284031215613f0757600080fd5b8151611136816133a2565b8381526001600160a01b0383166020820152606060408201819052600090613cf990830184613363565b600060208284031215613f4e57600080fd5b8151611136816134b9565b60008251613f6b81846020870161333f565b919091019291505056fe54656c65706f72746572546f6b656e44657374696e6174696f6e3a207a65726f54656c65706f72746572546f6b656e44657374696e6174696f6e3a20696e7661a264697066735822122055025661e48e6c862cfbff4caac63d876f057b6cf66b5edc890f4c9059414a7d64736f6c6343000812003354656c65706f72746572546f6b656e44657374696e6174696f6e3a207a65726f",
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

// CurrentReserveImbalance is a free data retrieval call binding the contract method 0x00d872ae.
//
// Solidity: function currentReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) CurrentReserveImbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "currentReserveImbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentReserveImbalance is a free data retrieval call binding the contract method 0x00d872ae.
//
// Solidity: function currentReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) CurrentReserveImbalance() (*big.Int, error) {
	return _NativeTokenDestination.Contract.CurrentReserveImbalance(&_NativeTokenDestination.CallOpts)
}

// CurrentReserveImbalance is a free data retrieval call binding the contract method 0x00d872ae.
//
// Solidity: function currentReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) CurrentReserveImbalance() (*big.Int, error) {
	return _NativeTokenDestination.Contract.CurrentReserveImbalance(&_NativeTokenDestination.CallOpts)
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

// MultiplyOnReceive is a free data retrieval call binding the contract method 0x1ce22075.
//
// Solidity: function multiplyOnReceive() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCaller) MultiplyOnReceive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "multiplyOnReceive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MultiplyOnReceive is a free data retrieval call binding the contract method 0x1ce22075.
//
// Solidity: function multiplyOnReceive() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) MultiplyOnReceive() (bool, error) {
	return _NativeTokenDestination.Contract.MultiplyOnReceive(&_NativeTokenDestination.CallOpts)
}

// MultiplyOnReceive is a free data retrieval call binding the contract method 0x1ce22075.
//
// Solidity: function multiplyOnReceive() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) MultiplyOnReceive() (bool, error) {
	return _NativeTokenDestination.Contract.MultiplyOnReceive(&_NativeTokenDestination.CallOpts)
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

// ScaleTokens is a free data retrieval call binding the contract method 0xb9448587.
//
// Solidity: function scaleTokens(uint256 value, bool isReceive) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) ScaleTokens(opts *bind.CallOpts, value *big.Int, isReceive bool) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "scaleTokens", value, isReceive)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaleTokens is a free data retrieval call binding the contract method 0xb9448587.
//
// Solidity: function scaleTokens(uint256 value, bool isReceive) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) ScaleTokens(value *big.Int, isReceive bool) (*big.Int, error) {
	return _NativeTokenDestination.Contract.ScaleTokens(&_NativeTokenDestination.CallOpts, value, isReceive)
}

// ScaleTokens is a free data retrieval call binding the contract method 0xb9448587.
//
// Solidity: function scaleTokens(uint256 value, bool isReceive) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) ScaleTokens(value *big.Int, isReceive bool) (*big.Int, error) {
	return _NativeTokenDestination.Contract.ScaleTokens(&_NativeTokenDestination.CallOpts, value, isReceive)
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

// Send is a paid mutator transaction binding the contract method 0x146dfd3c.
//
// Solidity: function send((bytes32,address,address,uint256,uint256,uint256) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Send(opts *bind.TransactOpts, input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0x146dfd3c.
//
// Solidity: function send((bytes32,address,address,uint256,uint256,uint256) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Send(&_NativeTokenDestination.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0x146dfd3c.
//
// Solidity: function send((bytes32,address,address,uint256,uint256,uint256) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Send(&_NativeTokenDestination.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x30079bff.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,uint256,uint256) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) SendAndCall(opts *bind.TransactOpts, input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "sendAndCall", input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x30079bff.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,uint256,uint256) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.SendAndCall(&_NativeTokenDestination.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x30079bff.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,uint256,uint256) input) payable returns()
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

// NativeTokenDestinationCollateralAddedIterator is returned from FilterCollateralAdded and is used to iterate over the raw logs and unpacked data for CollateralAdded events raised by the NativeTokenDestination contract.
type NativeTokenDestinationCollateralAddedIterator struct {
	Event *NativeTokenDestinationCollateralAdded // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationCollateralAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationCollateralAdded)
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
		it.Event = new(NativeTokenDestinationCollateralAdded)
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
func (it *NativeTokenDestinationCollateralAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationCollateralAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationCollateralAdded represents a CollateralAdded event raised by the NativeTokenDestination contract.
type NativeTokenDestinationCollateralAdded struct {
	Amount    *big.Int
	Remaining *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollateralAdded is a free log retrieval operation binding the contract event 0x244160b15e69cc411f041d94ae7fab6f6bba85dade8403216c05ff4b920d5449.
//
// Solidity: event CollateralAdded(uint256 amount, uint256 remaining)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterCollateralAdded(opts *bind.FilterOpts) (*NativeTokenDestinationCollateralAddedIterator, error) {

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "CollateralAdded")
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationCollateralAddedIterator{contract: _NativeTokenDestination.contract, event: "CollateralAdded", logs: logs, sub: sub}, nil
}

// WatchCollateralAdded is a free log subscription operation binding the contract event 0x244160b15e69cc411f041d94ae7fab6f6bba85dade8403216c05ff4b920d5449.
//
// Solidity: event CollateralAdded(uint256 amount, uint256 remaining)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchCollateralAdded(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationCollateralAdded) (event.Subscription, error) {

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "CollateralAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationCollateralAdded)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
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

// ParseCollateralAdded is a log parse operation binding the contract event 0x244160b15e69cc411f041d94ae7fab6f6bba85dade8403216c05ff4b920d5449.
//
// Solidity: event CollateralAdded(uint256 amount, uint256 remaining)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseCollateralAdded(log types.Log) (*NativeTokenDestinationCollateralAdded, error) {
	event := new(NativeTokenDestinationCollateralAdded)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
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

// NativeTokenDestinationTokensAndCallRoutedIterator is returned from FilterTokensAndCallRouted and is used to iterate over the raw logs and unpacked data for TokensAndCallRouted events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensAndCallRoutedIterator struct {
	Event *NativeTokenDestinationTokensAndCallRouted // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationTokensAndCallRoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTokensAndCallRouted)
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
		it.Event = new(NativeTokenDestinationTokensAndCallRouted)
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
func (it *NativeTokenDestinationTokensAndCallRoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTokensAndCallRoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTokensAndCallRouted represents a TokensAndCallRouted event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensAndCallRouted struct {
	TeleporterMessageID [32]byte
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallRouted is a free log retrieval operation binding the contract event 0xf10e091ede0da7ecb035d842d6990fd01ceebf01f51b4903f7c34de0a4e48f97.
//
// Solidity: event TokensAndCallRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,bytes,uint256,uint256,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTokensAndCallRouted(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*NativeTokenDestinationTokensAndCallRoutedIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TokensAndCallRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTokensAndCallRoutedIterator{contract: _NativeTokenDestination.contract, event: "TokensAndCallRouted", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallRouted is a free log subscription operation binding the contract event 0xf10e091ede0da7ecb035d842d6990fd01ceebf01f51b4903f7c34de0a4e48f97.
//
// Solidity: event TokensAndCallRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,bytes,uint256,uint256,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTokensAndCallRouted(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTokensAndCallRouted, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TokensAndCallRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTokensAndCallRouted)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensAndCallRouted", log); err != nil {
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

// ParseTokensAndCallRouted is a log parse operation binding the contract event 0xf10e091ede0da7ecb035d842d6990fd01ceebf01f51b4903f7c34de0a4e48f97.
//
// Solidity: event TokensAndCallRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,bytes,uint256,uint256,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTokensAndCallRouted(log types.Log) (*NativeTokenDestinationTokensAndCallRouted, error) {
	event := new(NativeTokenDestinationTokensAndCallRouted)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensAndCallRouted", log); err != nil {
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

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0x76b18d78fd0b0c8a046526d2a500e1e5ced780f056df0acc4932088d10e66562.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,uint256,uint256) input, uint256 amount)
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

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0x76b18d78fd0b0c8a046526d2a500e1e5ced780f056df0acc4932088d10e66562.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,uint256,uint256) input, uint256 amount)
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

// ParseTokensAndCallSent is a log parse operation binding the contract event 0x76b18d78fd0b0c8a046526d2a500e1e5ced780f056df0acc4932088d10e66562.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTokensAndCallSent(log types.Log) (*NativeTokenDestinationTokensAndCallSent, error) {
	event := new(NativeTokenDestinationTokensAndCallSent)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTokensRoutedIterator is returned from FilterTokensRouted and is used to iterate over the raw logs and unpacked data for TokensRouted events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensRoutedIterator struct {
	Event *NativeTokenDestinationTokensRouted // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationTokensRoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTokensRouted)
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
		it.Event = new(NativeTokenDestinationTokensRouted)
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
func (it *NativeTokenDestinationTokensRoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTokensRoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTokensRouted represents a TokensRouted event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensRouted struct {
	TeleporterMessageID [32]byte
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensRouted is a free log retrieval operation binding the contract event 0x5e2ff0b224d3b77f77e08ff220d9f882286bbf27218a51c330d2a9663c976cb1.
//
// Solidity: event TokensRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,uint256,uint256,uint256) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTokensRouted(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*NativeTokenDestinationTokensRoutedIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TokensRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTokensRoutedIterator{contract: _NativeTokenDestination.contract, event: "TokensRouted", logs: logs, sub: sub}, nil
}

// WatchTokensRouted is a free log subscription operation binding the contract event 0x5e2ff0b224d3b77f77e08ff220d9f882286bbf27218a51c330d2a9663c976cb1.
//
// Solidity: event TokensRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,uint256,uint256,uint256) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTokensRouted(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTokensRouted, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TokensRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTokensRouted)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensRouted", log); err != nil {
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

// ParseTokensRouted is a log parse operation binding the contract event 0x5e2ff0b224d3b77f77e08ff220d9f882286bbf27218a51c330d2a9663c976cb1.
//
// Solidity: event TokensRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,uint256,uint256,uint256) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTokensRouted(log types.Log) (*NativeTokenDestinationTokensRouted, error) {
	event := new(NativeTokenDestinationTokensRouted)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensRouted", log); err != nil {
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

// FilterTokensSent is a free log retrieval operation binding the contract event 0x78488d924de07bf96852578ad434a6c920f0835e97f9b302a77e1a77757c640b.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,uint256,uint256,uint256) input, uint256 amount)
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

// WatchTokensSent is a free log subscription operation binding the contract event 0x78488d924de07bf96852578ad434a6c920f0835e97f9b302a77e1a77757c640b.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,uint256,uint256,uint256) input, uint256 amount)
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

// ParseTokensSent is a log parse operation binding the contract event 0x78488d924de07bf96852578ad434a6c920f0835e97f9b302a77e1a77757c640b.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,uint256,uint256,uint256) input, uint256 amount)
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
