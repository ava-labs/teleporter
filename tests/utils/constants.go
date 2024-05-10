package utils

import "math/big"

var (
	DefaultERC20RequiredGas       = big.NewInt(85_000)
	DefaultNativeTokenRequiredGas = big.NewInt(120_000)
)
