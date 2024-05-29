// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"context"
	"errors"
	"math/big"

	"github.com/ava-labs/coreth/precompile/contracts/warp"
	predicateutils "github.com/ava-labs/subnet-evm/predicate"
	subnetEvmUtils "github.com/ava-labs/subnet-evm/utils"

	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
)

const (
	ReceiveCrossChainMessageStaticGasCost uint64 = 500_000

	BaseFeeFactor        = 2
	MaxPriorityFeePerGas = 2500000000 // 2.5 gwei
)

// CalculateReceiveMessageGasLimit calculates the estimated gas amount used by a single call
// to receiveCrossChainMessage for the given message and validator bit vector. The result amount
// depends on the required limit for the message execution, the number of validator signatures
// included in the aggregate signature, the static gas cost defined by the precompile, and an
// extra buffer amount defined here to ensure the call doesn't run out of gas.
func CalculateReceiveMessageGasLimit(
	ctx context.Context,
	client ethclient.Client,
	sender common.Address,
	teleporterContractAddress common.Address,
	destinationAddress common.Address,
	executionRequiredGasLimit *big.Int,
	teleporterCallData []byte,
	destinationCallData []byte,
	predicateBytes []byte) (uint64, error) {
	if !executionRequiredGasLimit.IsUint64() {
		return 0, errors.New("required gas limit too high")
	}

	accessList := types.AccessList{}
	accessList = append(accessList, types.AccessTuple{
		Address:     warp.ContractAddress,
		StorageKeys: subnetEvmUtils.BytesToHashSlice(predicateutils.PackPredicate(predicateBytes)),
	})
	gas1, err := client.EstimateGas(ctx, interfaces.CallMsg{
		From:       sender,
		To:         &teleporterContractAddress,
		Data:       teleporterCallData,
		AccessList: accessList,
	})
	if err != nil {
		return 0, err
	}

	gas2, err := client.EstimateGas(ctx, interfaces.CallMsg{
		From: teleporterContractAddress,
		To:   &destinationAddress,
		Data: destinationCallData,
	})
	if err != nil {
		return 0, err
	}

	res := gas1 - gas2 + executionRequiredGasLimit.Uint64()
	return res, nil
}
