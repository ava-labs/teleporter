// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"errors"
	"math/big"

	"github.com/ava-labs/avalanchego/utils/math"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
)

const (
	// The total gas cost for a single call to receiveCrossChainMessage is ~500_000. Add 300_000 to the
	// static gas cost charged by the Warp precompile per signature verification (200_000) to arrive at this number.
	ReceiveCrossChainMessageStaticGasCost uint64 = 300_000 + warp.GasCostPerSignatureVerification // Totals 500_000
	ReceiveMessageGasLimitBufferAmount    uint64 = 100_000

	BaseFeeFactor        = 2
	MaxPriorityFeePerGas = 2500000000 // 2.5 gwei
)

// CalculateReceiveMessageGasLimit calculates the estimated gas amount used by a single call
// to receiveCrossChainMessage for the given message and validator bit vector. The result amount
// depends on the required limit for the message execution, the number of validator signatures
// included in the aggregate signature, the static gas cost defined by the precompile, and an
// extra buffer amount defined here to ensure the call doesn't run out of gas.
func CalculateReceiveMessageGasLimit(numSigners int, executionRequiredGasLimit *big.Int) (uint64, error) {
	if !executionRequiredGasLimit.IsUint64() {
		return 0, errors.New("required gas limit too high")
	}

	gasAmounts := []uint64{
		executionRequiredGasLimit.Uint64(),
		ReceiveCrossChainMessageStaticGasCost,
		uint64(numSigners) * warp.GasCostPerWarpSigner,
		ReceiveMessageGasLimitBufferAmount,
	}

	res := gasAmounts[0]
	var err error
	for i := 1; i < len(gasAmounts); i++ {
		res, err = math.Add64(res, gasAmounts[i])
		if err != nil {
			return 0, err
		}
	}

	return res, nil
}
