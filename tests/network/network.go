package network

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

type SubnetTestInfo struct {
	SubnetID                  ids.ID
	BlockchainID              ids.ID
	NodeURIs                  []string
	NodeNames                 []string
	WSClient                  ethclient.Client
	RPCClient                 ethclient.Client
	EVMChainID                *big.Int
	TeleporterRegistryAddress common.Address
}

// Defines the interface for the network setup functions used in the E2E tests
type Network interface {
	GetSubnetInfo() (SubnetTestInfo, SubnetTestInfo)
	GetTeleporterContractAddress() common.Address
	GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey)

	RelayMessage(ctx context.Context,
		sourceReceipt *types.Receipt,
		source SubnetTestInfo,
		destination SubnetTestInfo,
		alterMessage bool,
		expectSuccess bool) *types.Receipt
}
