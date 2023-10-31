package network

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
)

// Implements Network, pointing to the network setup in network_setup.go
type LocalNetwork struct{}

func (g *LocalNetwork) GetSubnetsInfo() []utils.SubnetTestInfo {
	return utils.GetSubnetsInfo()
}

func (g *LocalNetwork) GetTeleporterContractAddress() common.Address {
	return utils.GetTeleporterContractAddress()
}

func (g *LocalNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	return utils.GetFundedAccountInfo()
}

func (g *LocalNetwork) RelayMessage(ctx context.Context,
	sourceBlockHash common.Hash,
	sourceBlockNumber *big.Int,
	source utils.SubnetTestInfo,
	destination utils.SubnetTestInfo) *types.Receipt {
	return utils.RelayMessage(ctx, sourceBlockHash, sourceBlockNumber, source, destination)
}
