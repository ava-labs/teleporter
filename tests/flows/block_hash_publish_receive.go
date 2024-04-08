package flows

import (
	"context"
	"math/big"

	coreEthClient "github.com/ava-labs/coreth/ethclient"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

func BlockHashPublishReceive(network interfaces.Network) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := utils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	publisherAddress, publisher := utils.DeployBlockHashPublisher(
		ctx,
		fundedKey,
		subnetAInfo,
	)
	receiverAddress, receiver := utils.DeployBlockHashReceiver(
		ctx,
		fundedKey,
		fundedAddress,
		subnetBInfo,
		publisherAddress,
		subnetAInfo.BlockchainID,
	)

	// coreth and subnet-evm have different Block implementations,
	// which means that the block hashes will be different when queried from different clients
	// To workaround this, we need to query the block hash from the coreth client
	// TODO: Design a unified interface that accounts for this different
	rpcUri := utils.HttpToRPCURI(subnetAInfo.NodeURIs[1], utils.CChainPathSpecifier)
	rpcClient, err := coreEthClient.Dial(rpcUri)
	Expect(err).Should(BeNil())
	expectedBlockNumberU64, err := rpcClient.BlockNumber(ctx)
	Expect(err).Should(BeNil())
	expectedBlockNumber := big.NewInt(0).SetUint64(expectedBlockNumberU64)

	block, err := rpcClient.BlockByNumber(
		ctx, expectedBlockNumber)
	Expect(err).Should(BeNil())
	expectedBlockHash := block.Hash()

	// publish latest block hash
	tx_opts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, subnetAInfo.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := publisher.PublishLatestBlockHash(
		tx_opts, subnetBInfo.BlockchainID, receiverAddress)
	Expect(err).Should(BeNil())

	receipt := utils.WaitForTransactionSuccess(ctx, subnetAInfo, tx.Hash())

	// relay publication
	network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)

	// receive publication
	blockNumber, blockHash, err := receiver.GetLatestBlockInfo(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// verify expectations
	Expect(blockNumber.Uint64()).Should(Equal(expectedBlockNumber.Uint64()))
	Expect(blockHash[:]).Should(Equal(expectedBlockHash[:]))
}
