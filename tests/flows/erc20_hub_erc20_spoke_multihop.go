package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20tokenhub "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHub/ERC20TokenHub"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a ERC20 token hub on the primary network
 * Deploys ERC20 token spoke to Subnet A and Subnet B
 * Bridges C-Chain example ERC20 tokens to Subnet A
 * Bridge tokens from Subnet A to Subnet B through multi-hop
 * Bridge back tokens from Subnet B to Subnet A through multi-hop
 */
func ERC20TokenHubERC20TokenSpokeMultiHop(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, subnetBInfo := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on subnet A as the token to be bridged
	exampleERC20Address, exampleERC20 := teleporterUtils.DeployExampleERC20(
		ctx,
		fundedKey,
		cChainInfo,
	)

	// Create an ERC20TokenHub for bridging the ERC20 token
	erc20TokenHubAddress, erc20TokenHub := utils.DeployERC20TokenHub(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		exampleERC20Address,
	)

	// Token representation on subnets A and B will have same name, symbol, and decimals
	tokenName, err := exampleERC20.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := exampleERC20.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := exampleERC20.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20TokenSpoke tp Subnet A
	erc20TokenSpokeAddressA, erc20TokenSpokeA := utils.DeployERC20TokenSpoke(
		ctx,
		fundedKey,
		subnetAInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHubAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Deploy an ERC20TokenSpoke for Subnet B
	erc20TokenSpokeAddressB, erc20TokenSpokeB := utils.DeployERC20TokenSpoke(
		ctx,
		fundedKey,
		subnetBInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHubAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Register both ERC20TokenSpoke instances on the ERC20TokenHub
	utils.RegisterERC20TokenSpokeOnHub(
		ctx,
		network,
		cChainInfo,
		erc20TokenHubAddress,
		subnetAInfo,
		erc20TokenSpokeAddressA,
	)
	utils.RegisterERC20TokenSpokeOnHub(
		ctx,
		network,
		cChainInfo,
		erc20TokenHubAddress,
		subnetBInfo,
		erc20TokenSpokeAddressB,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to Subnet A
	input := erc20tokenhub.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: erc20TokenSpokeAddressA,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   exampleERC20Address,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultERC20RequiredGas,
	}
	amount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(13))

	receipt, bridgedAmount := utils.SendERC20TokenHub(
		ctx,
		cChainInfo,
		erc20TokenHub,
		erc20TokenHubAddress,
		exampleERC20,
		input,
		amount,
		fundedKey,
	)

	// Relay the message to subnet A and check for ERC20TokenSpoke withdrawal
	receipt = network.RelayMessage(
		ctx,
		receipt,
		cChainInfo,
		subnetAInfo,
		true,
	)

	utils.CheckERC20TokenSpokeWithdrawal(
		ctx,
		erc20TokenSpokeA,
		receipt,
		recipientAddress,
		bridgedAmount,
	)

	// Check that the recipient received the tokens
	balance, err := erc20TokenSpokeA.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(bridgedAmount))

	bridgedAmount = big.NewInt(0).Div(bridgedAmount, big.NewInt(2))
	// Multi-hop transfer to Subnet B
	utils.SendERC20TokenMultiHopAndVerify(
		ctx,
		network,
		fundedKey,
		recipientKey,
		recipientAddress,
		subnetAInfo,
		erc20TokenSpokeA,
		erc20TokenSpokeAddressA,
		subnetBInfo,
		erc20TokenSpokeB,
		erc20TokenSpokeAddressB,
		cChainInfo,
		bridgedAmount,
	)

	// Multi-hop transfer back to Subnet A
	utils.SendERC20TokenMultiHopAndVerify(
		ctx,
		network,
		fundedKey,
		recipientKey,
		recipientAddress,
		subnetBInfo,
		erc20TokenSpokeB,
		erc20TokenSpokeAddressB,
		subnetAInfo,
		erc20TokenSpokeA,
		erc20TokenSpokeAddressA,
		cChainInfo,
		bridgedAmount,
	)
}
