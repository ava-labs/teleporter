package network

import (
	"context"
	"crypto/ecdsa"

	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/teleporter/tests/utils"
	localUtils "github.com/ava-labs/teleporter/tests/utils/local-network-utils"
	"github.com/ethereum/go-ethereum/common"
)

var _ Network = &LocalNetwork{}

// Implements Network, pointing to the network setup in local_network_setup.go
type LocalNetwork struct{}

func (n *LocalNetwork) GetSubnetsInfo() []utils.SubnetTestInfo {
	return localUtils.GetSubnetsInfo()
}

func (n *LocalNetwork) GetTeleporterContractAddress() common.Address {
	return localUtils.GetTeleporterContractAddress()
}

func (n *LocalNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	return localUtils.GetFundedAccountInfo()
}

func (n *LocalNetwork) RelayMessage(ctx context.Context,
	sourceReceipt *types.Receipt,
	source utils.SubnetTestInfo,
	destination utils.SubnetTestInfo,
	alterMessage bool,
	expectSuccess bool) *types.Receipt {
	return localUtils.RelayMessage(ctx, sourceReceipt, source, destination, alterMessage, expectSuccess)
}
