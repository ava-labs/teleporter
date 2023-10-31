package tests

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
		destination utils.SubnetTestInfo) *types.Receipt
}

// Implements Network, pointing to the network setup in network_setup.go
type ginkgoNetwork struct{}

func (g *ginkgoNetwork) GetSubnetsInfo() []utils.SubnetTestInfo {
	return utils.GetSubnetsInfo()
}

func (g *ginkgoNetwork) GetTeleporterContractAddress() common.Address {
	return utils.GetTeleporterContractAddress()
}

func (g *ginkgoNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	return utils.GetFundedAccountInfo()
}

func (g *ginkgoNetwork) RelayMessage(ctx context.Context,
	sourceBlockHash common.Hash,
	sourceBlockNumber *big.Int,
	source utils.SubnetTestInfo,
	destination utils.SubnetTestInfo) *types.Receipt {
	return utils.RelayMessage(ctx, sourceBlockHash, sourceBlockNumber, source, destination)
}
