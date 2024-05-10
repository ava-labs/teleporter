package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20destination "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Destination"
	erc20source "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Source"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy an ERC20 token source on the primary network
 * Deploys ERC20Destination to Subnet A
 * Bridges C-Chain example ERC20 tokens to Subnet A
 * Bridge tokens from Subnet A to C-Chain
 */
func ERC20SourceERC20Destination(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on the primary network as the source token to be bridged
	sourceTokenAddress, sourceToken := teleporterUtils.DeployExampleERC20(
		ctx,
		fundedKey,
		cChainInfo,
	)

	// Create an ERC20Source for bridging the source token
	erc20SourceAddress, erc20Source := utils.DeployERC20Source(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		sourceTokenAddress,
	)

	// Token representation on subnet A will have same name, symbol, and decimals
	tokenName, err := sourceToken.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := sourceToken.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := sourceToken.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20Destination to Subnet A
	erc20DestinationAddress, erc20Destination := utils.DeployERC20Destination(
		ctx,
		fundedKey,
		subnetAInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20SourceAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	utils.RegisterERC20DestinationOnSource(
		ctx,
		network,
		cChainInfo,
		erc20SourceAddress,
		subnetAInfo,
		erc20DestinationAddress,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to recipient on subnet A
	input := erc20source.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: erc20DestinationAddress,
		Recipient:                recipientAddress,
		FeeTokenAddress:          sourceTokenAddress,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultERC20RequiredGasLimit,
	}
	amount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(13))

	receipt, bridgedAmount := utils.SendERC20Source(
		ctx,
		cChainInfo,
		erc20Source,
		erc20SourceAddress,
		sourceToken,
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

	utils.CheckERC20DestinationWithdrawal(
		ctx,
		erc20Destination,
		receipt,
		recipientAddress,
		bridgedAmount,
	)

	// Check that the recipient received the tokens
	balance, err := erc20Destination.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(bridgedAmount))

	// Bridge back to source
	// Fund recipient with gas tokens on subnet A
	teleporterUtils.SendNativeTransfer(
		ctx,
		subnetAInfo,
		fundedKey,
		recipientAddress,
		big.NewInt(1e18),
	)
	inputB := erc20destination.SendTokensInput{
		DestinationBlockchainID:  cChainInfo.BlockchainID,
		DestinationBridgeAddress: erc20SourceAddress,
		Recipient:                recipientAddress,
		FeeTokenAddress:          erc20DestinationAddress,
		PrimaryFee:               big.NewInt(1e10),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultERC20RequiredGasLimit,
	}

	receipt, bridgedAmount = utils.SendERC20Destination(
		ctx,
		subnetAInfo,
		erc20Destination,
		erc20DestinationAddress,
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

	utils.CheckERC20SourceWithdrawal(
		ctx,
		erc20SourceAddress,
		sourceToken,
		receipt,
		recipientAddress,
		bridgedAmount,
	)

	// Check that the recipient received the tokens
	balance, err = sourceToken.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(bridgedAmount))
}
