package utils

import "math/big"

var (
	DefaultERC20RequiredGas       = big.NewInt(100_000)
	DefaultNativeTokenRequiredGas = big.NewInt(135_000)
)
