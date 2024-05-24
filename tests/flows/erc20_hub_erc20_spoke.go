package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20tokenhub "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHub/ERC20TokenHub"
	erc20tokenspoke "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenSpoke/ERC20TokenSpoke"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy an ERC20TokenHub on the primary network
 * Deploys ERC20TokenSpoke to Subnet A
 * Bridges C-Chain example ERC20 tokens to Subnet A
 * Bridge tokens from Subnet A to C-Chain
 */
func ERC20TokenHubERC20TokenSpoke(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on the primary network as the token to be bridged
	exampleERC20Address, exampleERC20 := utils.DeployExampleERC20(
		ctx,
		fundedKey,
		cChainInfo,
		erc20TokenHubDecimals,
	)

	exampleERC20Decimals, err := exampleERC20.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Create an ERC20TokenHub for bridging the ERC20 token
	erc20TokenHubAddress, erc20TokenHub := utils.DeployERC20TokenHub(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		exampleERC20Address,
		exampleERC20Decimals,
	)

	// Token representation on subnet A will have same name, symbol, and decimals
	tokenName, err := exampleERC20.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := exampleERC20.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := exampleERC20.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20TokenSpoke to Subnet A
	erc20TokenSpokeAddress, erc20TokenSpoke := utils.DeployERC20TokenSpoke(
		ctx,
		fundedKey,
		subnetAInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHubAddress,
		exampleERC20Decimals,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	utils.RegisterERC20TokenSpokeOnHub(
		ctx,
		network,
		cChainInfo,
		erc20TokenHubAddress,
		subnetAInfo,
		erc20TokenSpokeAddress,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to recipient on subnet A
	input := erc20tokenhub.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: erc20TokenSpokeAddress,
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

	// Bridge back to hub chain
	// Fund recipient with gas tokens on subnet A
	teleporterUtils.SendNativeTransfer(
		ctx,
		subnetAInfo,
		fundedKey,
		recipientAddress,
		big.NewInt(1e18),
	)
	inputB := erc20tokenspoke.SendTokensInput{
		DestinationBlockchainID:  cChainInfo.BlockchainID,
		DestinationBridgeAddress: erc20TokenHubAddress,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   erc20TokenSpokeAddress,
		PrimaryFee:               big.NewInt(1e10),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultERC20RequiredGas,
	}

	receipt, bridgedAmount = utils.SendERC20TokenSpoke(
		ctx,
		subnetAInfo,
		erc20TokenSpoke,
		erc20TokenSpokeAddress,
		inputB,
		teleporterUtils.BigIntSub(bridgedAmount, inputB.PrimaryFee),
		recipientKey,
	)

	receipt = network.RelayMessage(
		ctx,
		receipt,
		subnetAInfo,
		cChainInfo,
		true,
	)

	utils.CheckERC20TokenHubWithdrawal(
		ctx,
		erc20TokenHubAddress,
		exampleERC20,
		receipt,
		recipientAddress,
		bridgedAmount,
	)

	// Check that the recipient received the tokens
	balance, err = exampleERC20.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(bridgedAmount))
}
