package local

import (
	"context"
	"crypto/ecdsa"

	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ethereum/go-ethereum/common"
)

var _ network.Network = &LocalNetwork{}

// Implements Network, pointing to the network setup in local_network_setup.go
type LocalNetwork struct{}

func (n *LocalNetwork) GetSubnetInfo() (network.SubnetTestInfo, network.SubnetTestInfo) {
	return getSubnetInfo()
}

func (n *LocalNetwork) GetTeleporterContractAddress() common.Address {
	return getTeleporterContractAddress()
}

func (n *LocalNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	return getFundedAccountInfo()
}

func (n *LocalNetwork) RelayMessage(ctx context.Context,
	sourceReceipt *types.Receipt,
	source network.SubnetTestInfo,
	destination network.SubnetTestInfo,
	expectSuccess bool) *types.Receipt {
	return relayMessage(ctx, sourceReceipt, source, destination, expectSuccess)
}
