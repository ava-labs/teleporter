package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	nativetokenhub "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHub/NativeTokenHub"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/gomega"
)

/**
 * Deploy a NativeTokenHub on the primary network
 * Deploys ERC20TokenSpoke to Subnet A and Subnet B
 * Bridges C-Chain native tokens to Subnet A
 * Bridge tokens from Subnet A to Subnet B through multi-hop
 * Brige back tokens from Subnet B to Subnet A through multi-hop
 */
func NativeTokenHubERC20TokenSpokeMultiHop(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, subnetBInfo := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy a NativeTokenHub on the primary network
	wavaxAddress, wavax := utils.DeployWrappedNativeToken(
		ctx,
		fundedKey,
		cChainInfo,
		"AVAX",
	)
	nativeTokenHubAddress, nativeTokenHub := utils.DeployNativeTokenHub(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		wavaxAddress,
	)

	// Token representation on subnet B will have same name, symbol, and decimals
	tokenName, err := wavax.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := wavax.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := wavax.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20TokenSpoke on Subnet A
	erc20TokenSpokeAddressA, erc20TokenSpokeA := utils.DeployERC20TokenSpoke(
		ctx,
		fundedKey,
		subnetAInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenHubAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Deploy an ERC20TokenSpoke on Subnet B
	erc20TokenSpokeAddressB, erc20TokenSpokeB := utils.DeployERC20TokenSpoke(
		ctx,
		fundedKey,
		subnetBInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenHubAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Register both ERC20Destinations on the NativeTokenHub
	utils.RegisterERC20TokenSpokeOnHub(
		ctx,
		network,
		cChainInfo,
		nativeTokenHubAddress,
		subnetAInfo,
		erc20TokenSpokeAddressA,
	)

	utils.RegisterERC20TokenSpokeOnHub(
		ctx,
		network,
		cChainInfo,
		nativeTokenHubAddress,
		subnetBInfo,
		erc20TokenSpokeAddressB,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to recipient on subnet A
	input := nativetokenhub.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: erc20TokenSpokeAddressA,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   wavaxAddress,
		PrimaryFee:               big.NewInt(1e10),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultERC20RequiredGas,
	}

	// Send the tokens and verify expected events
	amount := big.NewInt(2e18)
	receipt, bridgedAmount := utils.SendNativeTokenSource(
		ctx,
		cChainInfo,
		nativeTokenHub,
		nativeTokenHubAddress,
		wavax,
		input,
		teleporterUtils.BigIntSub(amount, input.PrimaryFee),
		fundedKey,
	)

	// Relay the message to subnet B and check for message delivery
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

	// Send tokens from subnet A to recipient on subnet B through a multi-hop
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
}
