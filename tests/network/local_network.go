package network

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
)

// Implements Network, pointing to the network setup in local_network_setup.go
type LocalNetwork struct{}

func (n *LocalNetwork) GetSubnetsInfo() []utils.SubnetTestInfo {
	return utils.GetSubnetsInfo()
}

func (n *LocalNetwork) GetTeleporterContractAddress() common.Address {
	return utils.GetTeleporterContractAddress()
}

func (n *LocalNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	return utils.GetFundedAccountInfo()
}

func (n *LocalNetwork) RelayMessage(ctx context.Context,
	sourceBlockHash common.Hash,
	sourceBlockNumber *big.Int,
	source utils.SubnetTestInfo,
	destination utils.SubnetTestInfo,
	expectSuccess bool) *types.Receipt {
	return utils.RelayMessage(ctx, sourceBlockHash, sourceBlockNumber, source, destination, expectSuccess)
}
