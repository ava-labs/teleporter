package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	nativetokenhub "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHub/NativeTokenHub"
	erc20tokenspoke "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenSpoke/ERC20TokenSpoke"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a NativeTokenHub on the primary network
 * Deploys ERC20TokenSpoke to Subnet A
 * Bridges C-Chain native tokens to Subnet A
 * Bridge back tokens from Subnet A to C-Chain
 */
func NativeTokenHubERC20TokenSpoke(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an example WAVAX on the primary network
	wavaxAddress, wavax := utils.DeployWrappedNativeToken(
		ctx,
		fundedKey,
		cChainInfo,
		"AVAX",
	)

	// Create a NativeTokenHub for bridging the native token
	nativeTokenHubAddress, nativeTokenHub := utils.DeployNativeTokenHub(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		wavaxAddress,
	)

	// Token representation on subnet A will have same name, symbol, and decimals
	tokenName, err := wavax.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := wavax.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := wavax.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20TokenSpoke to Subnet A
	erc20TokenSpokeAddress, erc20TokenSpoke := utils.DeployERC20TokenSpoke(
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

	utils.RegisterERC20TokenSpokeOnHub(
		ctx,
		network,
		cChainInfo,
		nativeTokenHubAddress,
		subnetAInfo,
		erc20TokenSpokeAddress,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to recipient on subnet A
	input := nativetokenhub.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: erc20TokenSpokeAddress,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   wavaxAddress,
		PrimaryFee:               big.NewInt(1e18),
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
		amount,
		fundedKey,
	)

	// Relay the message to Subnet A and check for message delivery
	receipt = network.RelayMessage(
		ctx,
		receipt,
		cChainInfo,
		subnetAInfo,
		true,
	)

	utils.CheckERC20TokenSpokeWithdrawal(
		ctx,
		erc20TokenSpoke,
		receipt,
		recipientAddress,
		bridgedAmount,
	)

	// Check that the recipient received the tokens
	balance, err := erc20TokenSpoke.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(bridgedAmount))

	// Fund recipient with gas tokens on subnet A
	teleporterUtils.SendNativeTransfer(
		ctx,
		subnetAInfo,
		fundedKey,
		recipientAddress,
		big.NewInt(1e18),
	)
	inputA := erc20tokenspoke.SendTokensInput{
		DestinationBlockchainID:  cChainInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenHubAddress,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   erc20TokenSpokeAddress,
		PrimaryFee:               big.NewInt(1e10),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
	}

	// Send tokens on Subnet A back for native tokens on C-Chain
	receipt, bridgedAmount = utils.SendERC20TokenSpoke(
		ctx,
		subnetAInfo,
		erc20TokenSpoke,
		erc20TokenSpokeAddress,
		inputA,
		teleporterUtils.BigIntSub(bridgedAmount, inputA.PrimaryFee),
		recipientKey,
	)

	receipt = network.RelayMessage(
		ctx,
		receipt,
		subnetAInfo,
		cChainInfo,
		true,
	)

	// Check that the recipient received the tokens
	utils.CheckNativeTokenHubWithdrawal(
		ctx,
		nativeTokenHubAddress,
		wavax,
		receipt,
		bridgedAmount,
	)

	teleporterUtils.CheckBalance(ctx, recipientAddress, bridgedAmount, cChainInfo.RPCClient)
}
