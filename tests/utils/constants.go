package utils

import "math/big"

var (
	DefaultERC20RequiredGasLimit             = big.NewInt(70_000)
	DefaultNativeTokenRequiredGasLimit       = big.NewInt(90_000)
	DefaultERC20DestinationReserveImbalance  = big.NewInt(0)
	DefaultERC20DestinationTokenMultiplier   = big.NewInt(1)
	DefaultERC20DestinationMultiplyOnReceive = false
)
