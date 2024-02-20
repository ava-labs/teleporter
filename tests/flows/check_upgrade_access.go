package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

func CheckUpgradeAccess(network interfaces.Network) {
	subnetInfo := network.GetPrimaryNetworkInfo()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Deploy ExampleMessenger to the subnet
	//
	ctx := context.Background()
	teleporterAddress := network.GetTeleporterContractAddress()
	_, exampleMessenger := utils.DeployExampleCrossChainMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		subnetInfo,
	)

	// Check that owner is the funded address
	owner, err := exampleMessenger.Owner(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(owner).Should(Equal(fundedAddress))

	// Try to call updateMinTeleporterVersion from a non owner account
	nonOwnerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	nonOwnerAddress := crypto.PubkeyToAddress(nonOwnerKey.PublicKey)

	// Transfer native assets to the non owner account
	fundAmount := big.NewInt(0.1e18) // 0.1avax
	utils.SendNativeTransfer(
		ctx,
		subnetInfo,
		fundedKey,
		nonOwnerAddress,
		fundAmount,
	)

	// Check that access is not granted to the non owner and has no effect
	nonOwnerOpts, err := bind.NewKeyedTransactorWithChainID(
		nonOwnerKey, subnetInfo.EVMChainID)
	Expect(err).Should(BeNil())
	_, err = exampleMessenger.PauseTeleporterAddress(nonOwnerOpts, teleporterAddress)
	Expect(err).ShouldNot(BeNil())
	Expect(err.Error()).Should(ContainSubstring(errCallerNotOwnerStr))

	// Check that the teleporter address is not paused, because previous call should have failed
	isPaused, err := exampleMessenger.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeFalse())

	// Check that the owner is able to pause the Teleporter address
	ownerOpts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, subnetInfo.EVMChainID)
	Expect(err).Should(BeNil())
	// Try to call pauseTeleporterAddress from the owner account
	tx, err := exampleMessenger.PauseTeleporterAddress(ownerOpts, teleporterAddress)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, subnetInfo, tx.Hash())
	blockNum := receipt.BlockNumber
	pauseTeleporterEvent, err := utils.GetEventFromLogs(receipt.Logs, exampleMessenger.ParseTeleporterAddressPaused)
	Expect(err).Should(BeNil())
	Expect(pauseTeleporterEvent.TeleporterAddress).Should(Equal(teleporterAddress))

	// We pass the block number to the call options because the Fuji C-chain endpoint is load balanced,
	// and between consecutive calls we might be reaching different nodes with different states.
	isPaused, err = exampleMessenger.IsTeleporterAddressPaused(&bind.CallOpts{BlockNumber: blockNum}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeTrue())

	// Transfer ownership to the non owner account
	tx, err = exampleMessenger.TransferOwnership(ownerOpts, nonOwnerAddress)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, subnetInfo, tx.Hash())

	// Try to call unpauseTeleporterAddress from the previous owner account
	_, err = exampleMessenger.UnpauseTeleporterAddress(ownerOpts, teleporterAddress)
	Expect(err).ShouldNot(BeNil())
	Expect(err.Error()).Should(ContainSubstring(errCallerNotOwnerStr))

	// Make sure the teleporter address is still paused
	isPaused, err = exampleMessenger.IsTeleporterAddressPaused(&bind.CallOpts{}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeTrue())

	// Try to call unpauseTeleporterAddress from the non owner account now
	tx, err = exampleMessenger.UnpauseTeleporterAddress(nonOwnerOpts, teleporterAddress)
	Expect(err).Should(BeNil())
	receipt = utils.WaitForTransactionSuccess(ctx, subnetInfo, tx.Hash())
	blockNum = receipt.BlockNumber
	unpauseTeleporterEvent, err := utils.GetEventFromLogs(receipt.Logs, exampleMessenger.ParseTeleporterAddressUnpaused)
	Expect(err).Should(BeNil())
	Expect(unpauseTeleporterEvent.TeleporterAddress).Should(Equal(teleporterAddress))

	// We pass the block number to the call options because the Fuji C-chain endpoint is load balanced,
	// and between consecutive calls we might be reaching different nodes with different states.
	isPaused, err = exampleMessenger.IsTeleporterAddressPaused(&bind.CallOpts{BlockNumber: blockNum}, teleporterAddress)
	Expect(err).Should(BeNil())
	Expect(isPaused).Should(BeFalse())
}
