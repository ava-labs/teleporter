package registry

import (
	"context"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

func PauseTeleporter(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	L1AInfo := network.GetPrimaryNetworkInfo()
	L1BInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Deploy TestMessenger to L1s A and B
	//
	ctx := context.Background()
	teleporterAddress := teleporter.TeleporterMessengerAddress(L1AInfo)
	_, testMessengerA := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(L1AInfo),
		L1AInfo,
	)
	testMessengerAddressB, testMessengerB := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(L1BInfo),
		L1BInfo,
	)

	// Pause Teleporter on L1 B
	opts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, L1BInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := testMessengerB.PauseTeleporterAddress(opts, teleporterAddress)
	Expect(err).Should(BeNil())

	receipt := utils.WaitForTransactionSuccess(ctx, L1BInfo, tx.Hash())
	pauseTeleporterEvent, err := utils.GetEventFromLogs(receipt.Logs, testMessengerB.ParseTeleporterAddressPaused)
	Expect(err).Should(BeNil())
	Expect(pauseTeleporterEvent.TeleporterAddress).Should(Equal(teleporterAddress))

	isPaused, err := testMessengerB.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeTrue())

	// Send a message from L1 A to L1 B, which should fail
	network.SendExampleCrossChainMessageAndVerify(
		ctx,
		teleporter,
		L1AInfo,
		testMessengerA,
		L1BInfo,
		testMessengerAddressB,
		testMessengerB,
		fundedKey,
		"message_1",
		false,
	)

	// Unpause Teleporter on L1 B
	tx, err = testMessengerB.UnpauseTeleporterAddress(opts, teleporterAddress)
	Expect(err).Should(BeNil())

	receipt = utils.WaitForTransactionSuccess(ctx, L1BInfo, tx.Hash())
	unpauseTeleporterEvent, err := utils.GetEventFromLogs(receipt.Logs, testMessengerB.ParseTeleporterAddressUnpaused)
	Expect(err).Should(BeNil())
	Expect(unpauseTeleporterEvent.TeleporterAddress).Should(Equal(teleporterAddress))

	isPaused, err = testMessengerB.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeFalse())

	// Send a message from L1 A to L1 B again, which should now succeed
	network.SendExampleCrossChainMessageAndVerify(
		ctx,
		teleporter,
		L1AInfo,
		testMessengerA,
		L1BInfo,
		testMessengerAddressB,
		testMessengerB,
		fundedKey,
		"message_2",
		true,
	)
}
