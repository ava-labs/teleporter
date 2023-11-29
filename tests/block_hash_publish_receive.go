package tests

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/teleporter/tests/network"
	localUtils "github.com/ava-labs/teleporter/tests/utils/local-network-utils"
	. "github.com/onsi/gomega"
)

func BlockHashPublishReceive(network network.Network) {
	var (
		subnets     = network.GetSubnetsInfo()
		subnetAInfo = subnets[0]
		subnetBInfo = subnets[1]

		_, fundedKey = network.GetFundedAccountInfo()

		ctx = context.Background()

		publisherAddress, publisher = localUtils.DeployBlockHashPublisher(
			ctx,
			fundedKey,
			subnetAInfo,
		)
		receiverAddress, receiver = localUtils.DeployBlockHashReceiver(
			ctx,
			fundedKey,
			subnetBInfo,
			publisherAddress,
			subnetAInfo.BlockchainID,
		)
	)

	// gather expectations

	expectedBlockNumberU64, err := subnetAInfo.ChainRPCClient.BlockNumber(ctx)
	Expect(err).Should(BeNil())
	expectedBlockNumber := big.NewInt(0).SetUint64(expectedBlockNumberU64)

	block, err := subnetAInfo.ChainRPCClient.BlockByNumber(
		ctx, expectedBlockNumber)
	Expect(err).Should(BeNil())
	expectedBlockHash := block.Hash()

	// publish latest block hash

	tx_opts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, subnetAInfo.ChainIDInt)
	Expect(err).Should(BeNil())

	tx, err := publisher.PublishLatestBlockHash(
		tx_opts, subnetBInfo.BlockchainID, receiverAddress)
	Expect(err).Should(BeNil())

	receipt, err := bind.WaitMined(ctx, subnetAInfo.ChainRPCClient, tx)
	Expect(err).Should(BeNil())

	// relay publication

	network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)

	// receive publication

	blockNumber, blockHash, err := receiver.GetLatestBlockInfo(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// verify expectations

	Expect(blockNumber).Should(BeEquivalentTo(expectedBlockNumber))
	Expect(blockHash).Should(BeEquivalentTo(expectedBlockHash))
}
