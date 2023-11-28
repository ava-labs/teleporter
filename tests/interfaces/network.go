package interfaces

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ethereum/go-ethereum/common"
)

type SubnetTestInfo struct {
	SubnetID                  ids.ID
	BlockchainID              ids.ID
	ChainNodeURIs             []string
	ChainWSClient             ethclient.Client
	ChainRPCClient            ethclient.Client
	ChainIDInt                *big.Int
	TeleporterRegistryAddress common.Address
	TeleporterMessenger       *teleportermessenger.TeleporterMessenger
}

// Defines the interface for the network setup functions used in the E2E tests
type Network interface {
	GetSubnetsInfo() []SubnetTestInfo
	GetTeleporterContractAddress() common.Address
	GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey)
	RelayMessage(ctx context.Context,
		sourceReceipt *types.Receipt,
		source SubnetTestInfo,
		destination SubnetTestInfo,
		expectSuccess bool) *types.Receipt
}
