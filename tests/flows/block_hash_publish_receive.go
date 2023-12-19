package flows

import (
	"context"
	"encoding/hex"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

func BlockHashPublishReceive(network interfaces.Network) {
	// subnetAInfo := network.GetPrimaryNetworkInfo() TODO: Integrate the C-Chain
	subnetAInfo, subnetBInfo := utils.GetTwoSubnets(network)
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

	// gather expectations

	expectedBlockNumberU64, err := subnetAInfo.RPCClient.BlockNumber(ctx)
	Expect(err).Should(BeNil())
	expectedBlockNumber := big.NewInt(0).SetUint64(expectedBlockNumberU64)

	block, err := subnetAInfo.RPCClient.BlockByNumber(
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

	receipt := utils.WaitForTransactionSuccess(ctx, subnetAInfo, tx)

	publishLog, err := utils.GetEventFromLogs(receipt.Logs, publisher.ParsePublishBlockHash)
	Expect(err).Should(BeNil())
	Expect(publishLog.BlockHeight.Uint64()).Should(Equal(expectedBlockNumberU64))
	Expect(publishLog.BlockHash[:]).Should(Equal(expectedBlockHash[:]))

	// relay publication
	network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)

	// receive publication
	blockNumber, blockHash, err := receiver.GetLatestBlockInfo(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// verify expectations
	Expect(blockNumber.Uint64()).Should(Equal(expectedBlockNumber.Uint64()))
	Expect(hex.EncodeToString(blockHash[:])).Should(Equal(hex.EncodeToString(expectedBlockHash[:])))
}
