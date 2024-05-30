// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"errors"
	"math/big"

	"github.com/ava-labs/avalanchego/utils/math"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
)

const (
	ReceiveCrossChainMessageBaseGasCost uint64 = 250_000
	MarkMessageReceiptGasCost           uint64 = 2500
	DecodeMessageGasCostPerByte         uint64 = 35

	BaseFeeFactor        = 2
	MaxPriorityFeePerGas = 2500000000 // 2.5 gwei
)

// CalculateReceiveMessageGasLimit calculates the estimated gas amount used by a single call
// to receiveCrossChainMessage for the given message and validator bit vector. The result amount
// depends on the required limit for the message execution, the number of validator signatures
// included in the aggregate signature, the static gas cost defined by the precompile, and an
// extra buffer amount defined here to ensure the call doesn't run out of gas.
func CalculateReceiveMessageGasLimit(
	numSigners int,
	executionRequiredGasLimit *big.Int,
	warpMessageSize int,
	teleporterMessage teleportermessenger.TeleporterMessage) (uint64, error) {
	if !executionRequiredGasLimit.IsUint64() {
		return 0, errors.New("required gas limit too high")
	}

	gasAmounts := []uint64{
		executionRequiredGasLimit.Uint64(),
		// The variable gas on message bytes is accounted for both when used in predicate verification
		// and also when used in `getVerifiedWarpMessage`
		uint64(warpMessageSize) * warp.GasCostPerWarpMessageBytes * 2,
		// Take into the variable gas cost for decoding the Teleporter message
		// and marking the receipts as received.
		uint64(len(teleporterMessage.Message)) * DecodeMessageGasCostPerByte,
		uint64(len(teleporterMessage.Receipts)) * MarkMessageReceiptGasCost,
		ReceiveCrossChainMessageBaseGasCost,
		uint64(numSigners) * warp.GasCostPerWarpSigner,
		warp.GasCostPerSignatureVerification,
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
