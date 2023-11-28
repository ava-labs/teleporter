package local

import (
	"context"
	"crypto/ecdsa"

	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ethereum/go-ethereum/common"
)

var _ interfaces.Network = &LocalNetwork{}

// Implements Network, pointing to the network setup in local_network_setup.go
type LocalNetwork struct{}

func (n *LocalNetwork) GetSubnetsInfo() []interfaces.SubnetTestInfo {
	return getSubnetsInfo()
}

func (n *LocalNetwork) GetTeleporterContractAddress() common.Address {
	return getTeleporterContractAddress()
}

func (n *LocalNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	return getFundedAccountInfo()
}

func (n *LocalNetwork) RelayMessage(ctx context.Context,
	sourceReceipt *types.Receipt,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
	expectSuccess bool) *types.Receipt {
	return relayMessage(ctx, sourceReceipt, source, destination, expectSuccess)
}
