package utils

import "math/big"

var (
	DefaultERC20RequiredGasLimit       = big.NewInt(100_000)
	DefaultNativeTokenRequiredGasLimit = big.NewInt(150_000)
)
