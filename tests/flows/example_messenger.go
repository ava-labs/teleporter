package flows

import (
	"context"

	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
)

func ExampleMessenger(network interfaces.Network) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := utils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Deploy ExampleMessenger to Subnets A and B
	//
	ctx := context.Background()

	_, exampleMessengerA := utils.DeployExampleCrossChainMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		subnetAInfo,
	)
	exampleMessengerAddressB, exampleMessengerB := utils.DeployExampleCrossChainMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		subnetBInfo,
	)

	utils.SendExampleCrossChainMessageAndVerify(
		ctx,
		network,
		subnetAInfo,
		exampleMessengerA,
		subnetBInfo,
		exampleMessengerAddressB,
		exampleMessengerB,
		fundedKey,
		"Hello World!",
		true,
	)
}
