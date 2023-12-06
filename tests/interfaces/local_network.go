package interfaces

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/core/types"
)

type LocalNetwork interface {
	Network
	AddSubnetValidators(ctx context.Context, subnetID ids.ID, nodeNames []string)
	ConstructSignedWarpMessageBytes(
		ctx context.Context,
		sourceReceipt *types.Receipt,
		source SubnetTestInfo,
		destination SubnetTestInfo,
	) []byte
}
