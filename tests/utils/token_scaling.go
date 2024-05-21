package utils

import (
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20source "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Source"
	nativetokensource "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/NativeTokenSource"
	"github.com/ethereum/go-ethereum/common"

	. "github.com/onsi/gomega"
)

// ApplyTokenScaling applies token scaling to the given amount of source tokens.
// Token scaling is applied when sending tokens from the source to the destination bridge,
func ApplyTokenScaling(
	tokenMultiplier *big.Int,
	multiplyOnSpoke bool,
	sourceTokenAmount *big.Int,
) *big.Int {
	return scaleTokens(tokenMultiplier, multiplyOnSpoke, sourceTokenAmount, true)
}

// RemoveTokenScaling removes token scaling from the given amount of destination tokens.
// Token scaling is removed when sending tokens from the destination bridge back to the source.
func RemoveTokenScaling(
	tokenMultiplier *big.Int,
	multiplyOnSpoke bool,
	destinationTokenAmount *big.Int,
) *big.Int {
	return scaleTokens(tokenMultiplier, multiplyOnSpoke, destinationTokenAmount, false)
}

func scaleTokens(
	tokenMultiplier *big.Int,
	multiplyOnSpoke bool,
	amount *big.Int,
	isSendToDestination bool,
) *big.Int {
	// Multiply when multiplyOnSpoke and isSendToDestination are
	// both true or both false.
	if multiplyOnSpoke == isSendToDestination {
		return big.NewInt(0).Mul(amount, tokenMultiplier)
	}

	return big.NewInt(0).Div(amount, tokenMultiplier)
}

// GetScaledAmountFromERC20Source returns the scaled amount of destination tokens that
// will be sent to the destination bridge for an amount of source tokens.
func GetScaledAmountFromERC20Source(
	erc20Source *erc20source.ERC20Source,
	destinationBlockchainID ids.ID,
	destinationBridgeAddress common.Address,
	sourceTokenAmount *big.Int,
) *big.Int {
	destinationSettings, err := erc20Source.RegisteredDestinations(
		&bind.CallOpts{},
		destinationBlockchainID,
		destinationBridgeAddress,
	)
	Expect(err).Should(BeNil())

	return ApplyTokenScaling(
		destinationSettings.TokenMultiplier,
		destinationSettings.multiplyOnSpoke,
		sourceTokenAmount,
	)
}

// GetScaledAmountFromNativeTokenSource returns the scaled amount of tokens that will be sent to the destination bridge
// for corresponding amount of source tokens.
func GetScaledAmountFromNativeTokenSource(
	nativeTokenSource *nativetokensource.NativeTokenSource,
	destinationBlockchainID ids.ID,
	destinationBridgeAddress common.Address,
	amount *big.Int,
) *big.Int {
	destinationSettings, err := nativeTokenSource.RegisteredDestinations(
		&bind.CallOpts{},
		destinationBlockchainID,
		destinationBridgeAddress,
	)
	Expect(err).Should(BeNil())

	return ApplyTokenScaling(
		destinationSettings.TokenMultiplier,
		destinationSettings.multiplyOnSpoke,
		amount,
	)
}

func calculateCollateralNeeded(
	initialReserveImbalance *big.Int,
	tokenMultiplier *big.Int,
	multiplyOnSpoke bool,
) *big.Int {
	collateralNeeded := RemoveTokenScaling(tokenMultiplier, multiplyOnSpoke, initialReserveImbalance)

	remainder := big.NewInt(0).Mod(collateralNeeded, tokenMultiplier)
	if multiplyOnSpoke && (remainder.Cmp(big.NewInt(0)) != 0) {
		collateralNeeded.Add(collateralNeeded, big.NewInt(1))
	}
	return collateralNeeded
}
