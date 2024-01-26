package flows

import (
	"context"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

func PauseTeleporter(network interfaces.Network) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := utils.GetTwoSubnets(network)
	_, fundedKey := network.GetFundedAccountInfo()

	//
	// Deploy ExampleMessenger to Subnets A and B
	//
	ctx := context.Background()
	teleporterAddress := network.GetTeleporterContractAddress()
	_, exampleMessengerA := utils.DeployExampleCrossChainMessenger(ctx, fundedKey, subnetAInfo)
	exampleMessengerAddressB, exampleMessengerB := utils.DeployExampleCrossChainMessenger(
		ctx, fundedKey, subnetBInfo,
	)

	// Pause Teleporter on subnet B
	opts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, subnetBInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := exampleMessengerB.PauseTeleporterAddress(opts, teleporterAddress)
	Expect(err).Should(BeNil())

	receipt := utils.WaitForTransactionSuccess(ctx, subnetBInfo, tx)
	pauseTeleporterEvent, err := utils.GetEventFromLogs(receipt.Logs, exampleMessengerB.ParseTeleporterAddressPaused)
	Expect(err).Should(BeNil())
	Expect(pauseTeleporterEvent.TeleporterAddress).Should(Equal(teleporterAddress))

	isPaused, err := exampleMessengerB.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeTrue())

	// Send a message from subnet A to subnet B, which should fail
	utils.SendExampleCrossChainMessageAndVerify(
		ctx,
		network,
		subnetAInfo,
		exampleMessengerA,
		subnetBInfo,
		exampleMessengerAddressB,
		exampleMessengerB,
		fundedKey,
		"message_1",
		false)

	// Unpause Teleporter on subnet B
	tx, err = exampleMessengerB.UnpauseTeleporterAddress(opts, teleporterAddress)
	Expect(err).Should(BeNil())

	receipt = utils.WaitForTransactionSuccess(ctx, subnetBInfo, tx)
	unpauseTeleporterEvent, err := utils.GetEventFromLogs(receipt.Logs, exampleMessengerB.ParseTeleporterAddressUnpaused)
	Expect(err).Should(BeNil())
	Expect(unpauseTeleporterEvent.TeleporterAddress).Should(Equal(teleporterAddress))

	isPaused, err = exampleMessengerB.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeFalse())

	// Send a message from subnet A to subnet B again, which should now succeed
	utils.SendExampleCrossChainMessageAndVerify(
		ctx,
		network,
		subnetAInfo,
		exampleMessengerA,
		subnetBInfo,
		exampleMessengerAddressB,
		exampleMessengerB,
		fundedKey,
		"message_2",
		true)
}
