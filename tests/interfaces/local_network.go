package interfaces

import (
	"context"
	"crypto/ecdsa"

	"github.com/ava-labs/avalanchego/ids"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	pwallet "github.com/ava-labs/avalanchego/wallet/chain/p/wallet"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ethereum/go-ethereum/common"
)

type LocalNetwork interface {
	Network
	AddSubnetValidators(ctx context.Context, subnetID ids.ID, count uint)
	ExtractWarpMessageFromLog(
		ctx context.Context,
		sourceReceipt *types.Receipt,
		source SubnetTestInfo,
	) *avalancheWarp.UnsignedMessage
	ConstructSignedWarpMessage(
		ctx context.Context,
		sourceReceipt *types.Receipt,
		source SubnetTestInfo,
		destination SubnetTestInfo,
	) *avalancheWarp.Message
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
	Dir() string
	GetPChainWallet() pwallet.Wallet
}
