package flows

import (
	"context"

	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
)

func ExampleMessenger(network interfaces.Network) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := utils.GetTwoSubnets(network)
	_, fundedKey := network.GetFundedAccountInfo()

	//
	// Deploy ExampleMessenger to Subnets A and B
	//
	ctx := context.Background()

	_, exampleMessengerA := utils.DeployExampleCrossChainMessenger(ctx, fundedKey, subnetAInfo)
	exampleMessengerAddressB, exampleMessengerB := utils.DeployExampleCrossChainMessenger(
		ctx, fundedKey, subnetBInfo,
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
