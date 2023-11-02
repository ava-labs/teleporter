package network

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
)

// Defines the interface for the network setup functions used in the E2E tests
type Network interface {
	GetSubnetsInfo() []utils.SubnetTestInfo
	GetTeleporterContractAddress() common.Address
	GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey)
	RelayMessage(ctx context.Context,
		sourceBlockHash common.Hash,
		sourceBlockNumber *big.Int,
		source utils.SubnetTestInfo,
		destination utils.SubnetTestInfo,
		expectSuccess bool) *types.Receipt
}
