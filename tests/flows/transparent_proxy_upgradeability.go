package flows

import (
	"context"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	transparentupgradeableproxy "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TransparentUpgradeableProxy"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

func TransparentProxyUpgradeability(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	opts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey,
		cChainInfo.EVMChainID,
	)
	Expect(err).To(BeNil())

	utils.DeployERC20TokenHub(ctx, fundedKey, cChainInfo, fundedAddress)

	transparentupgradeableproxy.DeployTransparentUpgradeableProxy(opts, cChainInfo.RPCClient)

}
