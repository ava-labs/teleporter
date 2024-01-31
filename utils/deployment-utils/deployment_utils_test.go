package utils

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDefaultContractCreationGasPrice(t *testing.T) {
	gasPrice := GetDefaultContractCreationGasPrice()
	require.Equal(t, defaultContractCreationGasPrice, gasPrice)

	// Ensure the function returns a copy of the default gas price.
	gasPrice.Add(gasPrice, big.NewInt(1))
	require.NotEqual(t, defaultContractCreationGasPrice, gasPrice)

	// Ensure if the default gas price is changed, the function returns the new value.
	newDefaultGasPrice := big.NewInt(1234)
	defaultContractCreationGasPrice = newDefaultGasPrice
	gasPrice = GetDefaultContractCreationGasPrice()
	require.Equal(t, newDefaultGasPrice, gasPrice)
}
