package utils

import (
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20tokenhub "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHub/ERC20TokenHub"
	nativetokenhub "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHub/NativeTokenHub"
	"github.com/ethereum/go-ethereum/common"

	. "github.com/onsi/gomega"
)

// ApplyTokenScaling applies token scaling to the given amount of hub tokens.
// Token scaling is applied when sending tokens from the hub to the spoke instances.
func ApplyTokenScaling(
	tokenMultiplier *big.Int,
	multiplyOnSpoke bool,
	hubTokenAmount *big.Int,
) *big.Int {
	return scaleTokens(tokenMultiplier, multiplyOnSpoke, hubTokenAmount, true)
}

// RemoveTokenScaling removes token scaling from the given amount of spoke tokens.
// Token scaling is removed when sending tokens from the spoke back to the hub instance.
func RemoveTokenScaling(
	tokenMultiplier *big.Int,
	multiplyOnSpoke bool,
	spokeTokenAmount *big.Int,
) *big.Int {
	return scaleTokens(tokenMultiplier, multiplyOnSpoke, spokeTokenAmount, false)
}

func scaleTokens(
	tokenMultiplier *big.Int,
	multiplyOnSpoke bool,
	amount *big.Int,
	isSendToSpoke bool,
) *big.Int {
	// Multiply when multiplyOnSpoke and isSendToSpoke are
	// both true or both false.
	if multiplyOnSpoke == isSendToSpoke {
		return big.NewInt(0).Mul(amount, tokenMultiplier)
	}

	return big.NewInt(0).Div(amount, tokenMultiplier)
}

// GetScaledAmountFromERC20TokenHub returns the scaled amount of spoke tokens that
// will be sent to the spoke bridge for an amount of hub tokens.
func GetScaledAmountFromERC20TokenHub(
	erc20TokenHub *erc20tokenhub.ERC20TokenHub,
	spokeBlockchainID ids.ID,
	spokeAddress common.Address,
	hubTokenAmount *big.Int,
) *big.Int {
	spokeSettings, err := erc20TokenHub.RegisteredSpokes(
		&bind.CallOpts{},
		spokeBlockchainID,
		spokeAddress,
	)
	Expect(err).Should(BeNil())

	return ApplyTokenScaling(
		spokeSettings.TokenMultiplier,
		spokeSettings.MultiplyOnSpoke,
		hubTokenAmount,
	)
}

// GetScaledAmountFromNativeTokenHub returns the scaled amount of tokens that will be sent to
// the spoke bridge for corresponding amount of hub tokens.
func GetScaledAmountFromNativeTokenHub(
	nativeTokenHub *nativetokenhub.NativeTokenHub,
	spokeBlockchainID ids.ID,
	spokeAddress common.Address,
	amount *big.Int,
) *big.Int {
	destinationSettings, err := nativeTokenHub.RegisteredSpokes(
		&bind.CallOpts{},
		spokeBlockchainID,
		spokeAddress,
	)
	Expect(err).Should(BeNil())

	return ApplyTokenScaling(
		destinationSettings.TokenMultiplier,
		destinationSettings.MultiplyOnSpoke,
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
