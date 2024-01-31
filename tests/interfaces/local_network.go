package interfaces

import (
	"context"
	"crypto/ecdsa"

	runner_sdk "github.com/ava-labs/avalanche-network-runner/client"
	"github.com/ava-labs/avalanchego/ids"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ethereum/go-ethereum/common"
)

type LocalNetwork interface {
	Network
	AddSubnetValidators(ctx context.Context, subnetID ids.ID, nodeNames []string)
	ConstructSignedWarpMessage(
		ctx context.Context,
		sourceReceipt *types.Receipt,
		source SubnetTestInfo,
		destination SubnetTestInfo,
	) *avalancheWarp.Message
	GetAllNodeNames() []string
	RestartNodes(ctx context.Context, nodeNames []string, opts ...runner_sdk.OpOption)
	DeployTeleporterContracts(
		transactionBytes []byte,
		deployerAddress common.Address,
		contractAddress common.Address,
		fundedKey *ecdsa.PrivateKey,
		updateNetworkTeleporter bool)
	GetNetworkID() uint32
}
