package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20source "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Source"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a ERC20 token source on the primary network
 * Deploys ERC20Destination to Subnet A and Subnet B
 * Bridges C-Chain example ERC20 tokens to Subnet A
 * Bridge tokens from Subnet A to Subnet B through multi-hop
 * Bridge back tokens from Subnet B to Subnet A through multi-hop
 */
func ERC20SourceERC20DestinationMultihop(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, subnetBInfo := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on subnet A as the source token to be bridged
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

	// Token representation on subnets A and B will have same name, symbol, and decimals
	tokenName, err := sourceToken.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := sourceToken.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := sourceToken.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20Destination for Subnet A
	erc20DestinationAddress_A, erc20Destination_A := utils.DeployERC20Destination(
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

	// Deploy an ERC20Destination for Subnet B
	erc20DestinationAddress_B, erc20Destination_B := utils.DeployERC20Destination(
		ctx,
		fundedKey,
		subnetBInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20SourceAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Register both ERC20Destinations on the ERC20Source
	utils.RegisterERC20DestinationOnSource(
		ctx,
		network,
		cChainInfo,
		erc20SourceAddress,
		subnetAInfo,
		erc20DestinationAddress_A,
	)
	utils.RegisterERC20DestinationOnSource(
		ctx,
		network,
		cChainInfo,
		erc20SourceAddress,
		subnetBInfo,
		erc20DestinationAddress_B,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to Subnet A
	input := erc20source.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: erc20DestinationAddress_A,
		Recipient:                recipientAddress,
		FeeTokenAddress:          sourceTokenAddress,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultERC20RequiredGas,
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

	// Relay the message to subnet A and check for ERC20Destination withdrawal
	receipt = network.RelayMessage(
		ctx,
		receipt,
		cChainInfo,
		subnetAInfo,
		true,
	)

	utils.CheckERC20DestinationWithdrawal(
		ctx,
		erc20Destination_A,
		receipt,
		recipientAddress,
		bridgedAmount,
	)

	// Check that the recipient received the tokens
	balance, err := erc20Destination_A.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(bridgedAmount))

	bridgedAmount = big.NewInt(0).Div(bridgedAmount, big.NewInt(2))
	// Multi-hop transfer to Subnet B
	utils.SendERC20MultihopAndVerify(
		ctx,
		network,
		fundedKey,
		recipientKey,
		recipientAddress,
		subnetAInfo,
		erc20Destination_A,
		erc20DestinationAddress_A,
		subnetBInfo,
		erc20Destination_B,
		erc20DestinationAddress_B,
		cChainInfo,
		bridgedAmount,
	)

	// Multi-hop transfer back to Subnet A
	utils.SendERC20MultihopAndVerify(
		ctx,
		network,
		fundedKey,
		recipientKey,
		recipientAddress,
		subnetBInfo,
		erc20Destination_B,
		erc20DestinationAddress_B,
		subnetAInfo,
		erc20Destination_A,
		erc20DestinationAddress_A,
		cChainInfo,
		bridgedAmount,
	)
}
