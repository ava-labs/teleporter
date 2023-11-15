package network

import (
	"context"
	"crypto/ecdsa"

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
		sourceReceipt *types.Receipt,
		source utils.SubnetTestInfo,
		destination utils.SubnetTestInfo,
		alterMessage bool,
		expectSuccess bool) *types.Receipt
}
