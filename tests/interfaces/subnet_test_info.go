package interfaces

import (
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/ethclient"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	teleporterregistry "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/upgrades/TeleporterRegistry"
	"github.com/ethereum/go-ethereum/common"
)

// Tracks information about a test subnet used for executing tests against.
type SubnetTestInfo struct {
	SubnetID            ids.ID
	BlockchainID        ids.ID
	NodeURIs            []string
	WSClient            ethclient.Client
	RPCClient           ethclient.Client
	EVMChainID          *big.Int
	TeleporterRegistry  *teleporterregistry.TeleporterRegistry
	TeleporterMessenger *teleportermessenger.TeleporterMessenger

	// TeleporterRegistryAddress is unique across subnets.
	TeleporterRegistryAddress common.Address
}
