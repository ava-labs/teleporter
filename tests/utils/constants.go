package utils

import "math/big"

var (
	DefaultERC20RequiredGasLimit       = big.NewInt(300_000)
	DefaultNativeTokenRequiredGasLimit = big.NewInt(400_000)
)
