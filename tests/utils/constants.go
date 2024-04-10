package utils

import "math/big"

var (
	DefaultERC20RequiredGasLimit       = big.NewInt(70_000)
	DefaultNativeTokenRequiredGasLimit = big.NewInt(90_000)
)
