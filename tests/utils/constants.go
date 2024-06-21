package utils

import "math/big"

var (
	DefaultERC20RequiredGas       = big.NewInt(150_000)
	DefaultNativeTokenRequiredGas = big.NewInt(220_000)
)
