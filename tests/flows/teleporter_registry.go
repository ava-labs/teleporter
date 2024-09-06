package flows

import (
	"context"
	"time"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

const (
	teleporterByteCodeFile = "./out/TeleporterMessenger.sol/TeleporterMessenger.json"
)

func TeleporterRegistry(network interfaces.LocalNetwork) {
	// Deploy dApp on both chains that use Teleporter Registry
	// Deploy version 2 of Teleporter to both chains
	// Construct AddProtocolVersion txs for both chains
	// Send tx for one of the two chains, verify events emitted
	// Send a message from chain with old version to chain with new version, verify message is received
	// Update minTeleporterVersion on chain with new version
	// Send same message from chain with old version, verify message is not received
	// Send a message from chain with new version to chain with old version, verify message is not received
	// Send tx with new protocol version to other chain, verify events emitted
	// Retry the previously failed message execution, verify message is now able to be delivered to dApp

	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, subnetBInfo := utils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy a test cross chain messenger to both chains
	testMessengerContractC, testMessengerC := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		cChainInfo,
	)
	testMessengerContractB, testMessengerB := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		subnetBInfo,
	)

	// Deploy the new version of Teleporter to both chains
	newTeleporterAddress := utils.DeployNewTeleporterVersion(ctx, network, fundedKey, teleporterByteCodeFile)
	networkID := network.GetNetworkID()
	// Create chain config file with off chain message for each chain
	offchainMessageC, warpEnabledChainConfigC := utils.InitOffChainMessageChainConfig(
		networkID,
		cChainInfo,
		newTeleporterAddress,
		2,
	)
	offchainMessageB, warpEnabledChainConfigB := utils.InitOffChainMessageChainConfig(
		networkID,
		subnetBInfo,
		newTeleporterAddress,
		2,
	)
	offchainMessageA, warpEnabledChainConfigA := utils.InitOffChainMessageChainConfig(
		networkID,
		subnetAInfo,
		newTeleporterAddress,
		2,
	)

	// Create chain config with off chain messages
	chainConfigs := make(utils.ChainConfigMap)
	chainConfigs.Add(cChainInfo, warpEnabledChainConfigC)
	chainConfigs.Add(subnetBInfo, warpEnabledChainConfigB)
	chainConfigs.Add(subnetAInfo, warpEnabledChainConfigA)

	// Restart nodes with new chain config
	network.SetChainConfigs(chainConfigs)
	restartCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	network.RestartNodes(restartCtx, network.GetAllNodeIDs())

	// Call addProtocolVersion on subnetB to register the new Teleporter version
	utils.AddProtocolVersionAndWaitForAcceptance(
		ctx,
		network,
		subnetBInfo,
		newTeleporterAddress,
		fundedKey,
		offchainMessageB,
	)

	// Send a message using old Teleporter version to test messenger using new Teleporter version.
	// Message should be received successfully since we haven't updated mininum Teleporter version yet.
	utils.SendExampleCrossChainMessageAndVerify(
		ctx,
		network,
		cChainInfo,
		testMessengerC,
		subnetBInfo,
		testMessengerContractB,
		testMessengerB,
		fundedKey,
		"message_1",
		true,
	)

	// Update minimum Teleporter version on destination chain
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetBInfo.EVMChainID)
	Expect(err).Should(BeNil())

	latestVersionB, err := subnetBInfo.TeleporterRegistry.LatestVersion(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	minTeleporterVersion, err := testMessengerB.GetMinTeleporterVersion(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tx, err := testMessengerB.UpdateMinTeleporterVersion(opts, latestVersionB)
	Expect(err).Should(BeNil())

	receipt := utils.WaitForTransactionSuccess(ctx, subnetBInfo, tx.Hash())

	// Verify that minTeleporterVersion updated
	minTeleporterVersionUpdatedEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		testMessengerB.ParseMinTeleporterVersionUpdated,
	)
	Expect(err).Should(BeNil())
	Expect(minTeleporterVersionUpdatedEvent.OldMinTeleporterVersion.Cmp(minTeleporterVersion)).Should(Equal(0))
	Expect(minTeleporterVersionUpdatedEvent.NewMinTeleporterVersion.Cmp(latestVersionB)).Should(Equal(0))

	// Send a message using old Teleporter version to test messenger with updated minimum Teleporter version.
	// Message should fail since we updated minimum Teleporter version.
	utils.SendExampleCrossChainMessageAndVerify(
		ctx,
		network,
		cChainInfo,
		testMessengerC,
		subnetBInfo,
		testMessengerContractB,
		testMessengerB,
		fundedKey,
		"message_2",
		false,
	)

	// Update the subnets to use new Teleporter messengers
	network.SetTeleporterContractAddress(newTeleporterAddress)
	cChainInfo = network.GetPrimaryNetworkInfo()
	subnetAInfo, subnetBInfo = utils.GetTwoSubnets(network)

	utils.SendExampleCrossChainMessageAndVerify(
		ctx,
		network,
		subnetBInfo,
		testMessengerB,
		cChainInfo,
		testMessengerContractC,
		testMessengerC,
		fundedKey,
		"message_3",
		false,
	)

	// Call addProtocolVersion on subnetA to register the new Teleporter version
	utils.AddProtocolVersionAndWaitForAcceptance(
		ctx,
		network,
		cChainInfo,
		newTeleporterAddress,
		fundedKey,
		offchainMessageC,
	)

	// Send a message from A->B, which previously failed, but now using the new Teleporter version.
	// Teleporter versions should match, so message should be received successfully.
	utils.SendExampleCrossChainMessageAndVerify(
		ctx,
		network,
		subnetBInfo,
		testMessengerB,
		cChainInfo,
		testMessengerContractC,
		testMessengerC,
		fundedKey,
		"message_4",
		true,
	)

	// To make sure all subnets are using the same Teleporter version, call addProtocolVersion on subnetA
	// to register the new Teleporter version
	utils.AddProtocolVersionAndWaitForAcceptance(
		ctx,
		network,
		subnetAInfo,
		newTeleporterAddress,
		fundedKey,
		offchainMessageA,
	)

	latestVersionA, err := subnetAInfo.TeleporterRegistry.LatestVersion(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(latestVersionA.Cmp(latestVersionB)).Should(Equal(0))

	latestVersionC, err := cChainInfo.TeleporterRegistry.LatestVersion(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(latestVersionC.Cmp(latestVersionB)).Should(Equal(0))
}
