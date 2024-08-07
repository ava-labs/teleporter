package local

import (
	"context"
	"crypto/ecdsa"
	"slices"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
	"github.com/ava-labs/subnet-evm/ethclient"
	subnetEvmUtils "github.com/ava-labs/subnet-evm/tests/utils"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/log"

	. "github.com/onsi/gomega"
)

// Issues txs to activate the proposer VM fork on the specified subnet index in the manager
func setupProposerVM(ctx context.Context, fundedKey *ecdsa.PrivateKey, network *tmpnet.Network, subnetID ids.ID) {
	subnetDetails := network.Subnets[slices.IndexFunc(
		network.Subnets,
		func(s *tmpnet.Subnet) bool { return s.SubnetID == subnetID },
	)]

	chainID := subnetDetails.Chains[0].ChainID

	nodeURI, err := network.GetURIForNodeID(subnetDetails.ValidatorIDs[0])
	Expect(err).Should(BeNil())
	uri := utils.HttpToWebsocketURI(nodeURI, chainID.String())

	client, err := ethclient.Dial(uri)
	Expect(err).Should(BeNil())
	chainIDInt, err := client.ChainID(ctx)
	Expect(err).Should(BeNil())

	err = subnetEvmUtils.IssueTxsToActivateProposerVMFork(ctx, chainIDInt, fundedKey, client)
	Expect(err).Should(BeNil())
}

// Blocks until all validators specified in nodeURIs have reached the specified block height
func waitForAllValidatorsToAcceptBlock(ctx context.Context, nodeURIs []string, blockchainID ids.ID, height uint64) {
	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	for i, uri := range nodeURIs {
		chainAWSURI := utils.HttpToWebsocketURI(uri, blockchainID.String())
		log.Debug("Creating ethclient for blockchain", "blockchainID", blockchainID.String(), "wsURI", chainAWSURI)
		client, err := ethclient.Dial(chainAWSURI)
		Expect(err).Should(BeNil())
		defer client.Close()

		// Loop until each node has advanced to >= the height of the block that emitted the warp log
		for {
			block, err := client.BlockByNumber(cctx, nil)
			Expect(err).Should(BeNil())
			if block.NumberU64() >= height {
				log.Debug("Client accepted the block containing SendWarpMessage", "client", i, "height", block.NumberU64())
				break
			}
		}
	}
}
