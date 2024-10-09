package interfaces

import (
	"context"
	"crypto/ecdsa"

	"github.com/ava-labs/avalanchego/ids"
	pwallet "github.com/ava-labs/avalanchego/wallet/chain/p/wallet"
	"github.com/ethereum/go-ethereum/common"
)

type LocalNetwork interface {
	Network
	AddSubnetValidators(ctx context.Context, subnetID ids.ID, count uint)
	GetAllNodeIDs() []ids.NodeID
	SetChainConfigs(chainConfigs map[string]string)
	RestartNodes(ctx context.Context, nodeIDs []ids.NodeID)
	DeployTeleporterContractToCChain(
		transactionBytes []byte,
		deployerAddress common.Address,
		contractAddress common.Address,
		fundedKey *ecdsa.PrivateKey,
	)
	DeployTeleporterContractToAllChains(
		transactionBytes []byte,
		deployerAddress common.Address,
		contractAddress common.Address,
		fundedKey *ecdsa.PrivateKey,
	)
	GetNetworkID() uint32
	GetPChainWallet() pwallet.Wallet
}
