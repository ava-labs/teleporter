package interfaces

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	pwallet "github.com/ava-labs/avalanchego/wallet/chain/p/wallet"
)

type LocalNetwork interface {
	Network
	AddSubnetValidators(ctx context.Context, subnetID ids.ID, count uint)
	SetChainConfigs(chainConfigs map[string]string)
	RestartNodes(ctx context.Context, nodeIDs []ids.NodeID)
	GetNetworkID() uint32
	GetPChainWallet() pwallet.Wallet
}
