package interfaces

import (
	"context"
	"crypto/ecdsa"

	"github.com/ava-labs/avalanchego/ids"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ethereum/go-ethereum/common"
)

// Defines the interface for the network setup functions used in the E2E tests
type Network interface {
	// Returns information about the primary network
	GetPrimaryNetworkInfo() SubnetTestInfo

	// Returns all of the subnets support by this network, excluding the primary network
	GetSubnetsInfo() []SubnetTestInfo

	// Returns all of the subnets support by this network in no particular order, including the primary network.
	// If the distinction between primary network and other subnets needs to be made,
	// use GetPrimaryNetworkInfo() and GetSubnetsInfo() instead.
	GetAllSubnetsInfo() []SubnetTestInfo

	// Returns the Teleporter contract address for all subnets in this network.
	GetTeleporterContractAddress() common.Address

	// Sets the Teleporter contract address for all subnets in this network.
	SetTeleporterContractAddress(address common.Address)

	// An address and corresponding key that has native tokens on each of the subnets in this network.
	GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey)

	// GetSignedMessage returns the signed Warp message for the specified Warp message ID.
	GetSignedMessage(
		ctx context.Context,
		source SubnetTestInfo,
		destination SubnetTestInfo,
		messageID ids.ID,
	) *avalancheWarp.Message

	// Relays the specified message between the two subnets,
	// and returns the receipt of the transaction the message was delivered in.
	RelayMessage(
		ctx context.Context,
		sourceReceipt *types.Receipt,
		source SubnetTestInfo,
		destination SubnetTestInfo,
		expectSuccess bool,
	) *types.Receipt
}
