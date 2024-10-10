package flows

import (
	"context"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

func PauseTeleporter(n *network.LocalNetwork) {
	subnetAInfo := n.GetPrimaryNetworkInfo()
	subnetBInfo, _ := n.GetTwoSubnets()
	fundedAddress, fundedKey := n.GetFundedAccountInfo()

	//
	// Deploy TestMessenger to Subnets A and B
	//
	ctx := context.Background()
	teleporterAddress := n.GetTeleporterContractAddress()
	_, testMessengerA := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		subnetAInfo,
	)
	testMessengerAddressB, testMessengerB := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		subnetBInfo,
	)

	// Pause Teleporter on subnet B
	opts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, subnetBInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := testMessengerB.PauseTeleporterAddress(opts, teleporterAddress)
	Expect(err).Should(BeNil())

	receipt := utils.WaitForTransactionSuccess(ctx, subnetBInfo, tx.Hash())
	pauseTeleporterEvent, err := utils.GetEventFromLogs(receipt.Logs, testMessengerB.ParseTeleporterAddressPaused)
	Expect(err).Should(BeNil())
	Expect(pauseTeleporterEvent.TeleporterAddress).Should(Equal(teleporterAddress))

	isPaused, err := testMessengerB.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeTrue())

	// Send a message from subnet A to subnet B, which should fail
	network.SendExampleCrossChainMessageAndVerify(
		ctx,
		n,
		subnetAInfo,
		testMessengerA,
		subnetBInfo,
		testMessengerAddressB,
		testMessengerB,
		fundedKey,
		"message_1",
		false,
	)

	// Unpause Teleporter on subnet B
	tx, err = testMessengerB.UnpauseTeleporterAddress(opts, teleporterAddress)
	Expect(err).Should(BeNil())

	receipt = utils.WaitForTransactionSuccess(ctx, subnetBInfo, tx.Hash())
	unpauseTeleporterEvent, err := utils.GetEventFromLogs(receipt.Logs, testMessengerB.ParseTeleporterAddressUnpaused)
	Expect(err).Should(BeNil())
	Expect(unpauseTeleporterEvent.TeleporterAddress).Should(Equal(teleporterAddress))

	isPaused, err = testMessengerB.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeFalse())

	// Send a message from subnet A to subnet B again, which should now succeed
	network.SendExampleCrossChainMessageAndVerify(
		ctx,
		n,
		subnetAInfo,
		testMessengerA,
		subnetBInfo,
		testMessengerAddressB,
		testMessengerB,
		fundedKey,
		"message_2",
		true,
	)
}
