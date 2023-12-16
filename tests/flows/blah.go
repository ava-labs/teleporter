package flows

import (
	"github.com/ava-labs/teleporter/tests/interfaces"
	// . "github.com/onsi/gomega"
)

func Blah(network interfaces.LocalNetwork) {
	// client := network.GetANRClient()
	// chains, err := client.ListBlockchains(context.Background())
	// Expect(err).Should(BeNil())

	// subnets, err := client.ListSubnets(context.Background())
	// Expect(err).Should(BeNil())

	// log.Info("subnets")
	// for _, subnet := range subnets {
	// 	log.Info("subnet", subnet)
	// }
	// log.Info("chains")
	// for _, chain := range chains {
	// 	log.Info("chain", chain.ChainId, "name", chain.ChainName)
	// }

	// subnetAInfo, _, _ := utils.GetThreeSubnets(network)
	// pChainClient := platformvm.NewClient(subnetAInfo.NodeURIs[0])
	// blockChains, err := pChainClient.GetBlockchains(context.Background())
	// Expect(err).Should(BeNil())

	// log.Info("blockchains")

	// for _, chain := range blockChains {
	// 	log.Info("chain", chain.ID.String(), "name", chain.Name)
	// }

	// // Produce signatures

}
