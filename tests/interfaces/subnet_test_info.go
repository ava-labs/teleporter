package interfaces

import (
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/ethclient"
)

// Tracks information about a test subnet used for executing tests against.
type SubnetTestInfo struct {
	SubnetID     ids.ID
	BlockchainID ids.ID
	NodeURIs     []string
	WSClient     ethclient.Client
	RPCClient    ethclient.Client
	EVMChainID   *big.Int
}
