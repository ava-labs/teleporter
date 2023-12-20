package flows

import (
	"context"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

func BlockHashPublishReceive(network interfaces.Network) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := utils.GetTwoSubnets(network)
	_, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	publisherAddress, publisher := utils.DeployBlockHashPublisher(
		ctx,
		fundedKey,
		subnetAInfo,
	)
	receiverAddress, receiver := utils.DeployBlockHashReceiver(
		ctx,
		fundedKey,
		subnetBInfo,
		publisherAddress,
		subnetAInfo.BlockchainID,
	)

	// publish latest block hash
	tx_opts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, subnetAInfo.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := publisher.PublishLatestBlockHash(
		tx_opts, subnetBInfo.BlockchainID, receiverAddress)
	Expect(err).Should(BeNil())

	receipt := utils.WaitForTransactionSuccess(ctx, subnetAInfo, tx)

	publishEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		publisher.ParsePublishBlockHash)
	Expect(err).Should(BeNil())
	expectedBlockNumber := publishEvent.BlockHeight
	expectedBlockHash := publishEvent.BlockHash

	// relay publication
	network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)

	// receive publication
	blockNumber, blockHash, err := receiver.GetLatestBlockInfo(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// verify expectations
	Expect(blockNumber.Uint64()).Should(Equal(expectedBlockNumber.Uint64()))
	Expect(blockHash).Should(Equal(expectedBlockHash))
}
