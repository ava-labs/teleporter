package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20tokenhome "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHome/ERC20TokenHome"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a ERC20 token home on the primary network
 * Deploys ERC20 token remote to Subnet A and Subnet B
 * Bridges C-Chain example ERC20 tokens to Subnet A
 * Bridge tokens from Subnet A to Subnet B through multi-hop
 * Bridge back tokens from Subnet B to Subnet A through multi-hop
 */
func ERC20TokenHomeERC20TokenRemoteMultiHop(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, subnetBInfo := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on subnet A as the token to be bridged
	exampleERC20Address, exampleERC20 := utils.DeployExampleERC20(
		ctx,
		fundedKey,
		cChainInfo,
		erc20TokenHomeDecimals,
	)

	homeTokenDecimals, err := exampleERC20.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Create an ERC20TokenHome for bridging the ERC20 token
	erc20TokenHomeAddress, erc20TokenHome := utils.DeployERC20TokenHome(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		exampleERC20Address,
		homeTokenDecimals,
	)

	// Token representation on subnets A and B will have same name, symbol, and decimals
	tokenName, err := exampleERC20.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := exampleERC20.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := exampleERC20.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20TokenRemote tp Subnet A
	erc20TokenRemoteAddressA, erc20TokenRemoteA := utils.DeployERC20TokenRemote(
		ctx,
		fundedKey,
		subnetAInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHomeAddress,
		homeTokenDecimals,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Deploy an ERC20TokenRemote for Subnet B
	erc20TokenRemoteAddressB, erc20TokenRemoteB := utils.DeployERC20TokenRemote(
		ctx,
		fundedKey,
		subnetBInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHomeAddress,
		homeTokenDecimals,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Register both ERC20TokenRemote instances on the ERC20TokenHome
	utils.RegisterERC20TokenRemoteOnHome(
		ctx,
		network,
		cChainInfo,
		erc20TokenHomeAddress,
		subnetAInfo,
		erc20TokenRemoteAddressA,
	)
	utils.RegisterERC20TokenRemoteOnHome(
		ctx,
		network,
		cChainInfo,
		erc20TokenHomeAddress,
		subnetBInfo,
		erc20TokenRemoteAddressB,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to Subnet A
	input := erc20tokenhome.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: erc20TokenRemoteAddressA,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   exampleERC20Address,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultERC20RequiredGas,
	}
	amount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(13))

	receipt, bridgedAmount := utils.SendERC20TokenHome(
		ctx,
		cChainInfo,
		erc20TokenHome,
		erc20TokenHomeAddress,
		exampleERC20,
		input,
		amount,
		fundedKey,
	)

	// Relay the message to subnet A and check for ERC20TokenRemote withdrawal
	receipt = network.RelayMessage(
		ctx,
		receipt,
		cChainInfo,
		subnetAInfo,
		true,
	)

	utils.CheckERC20TokenRemoteWithdrawal(
		ctx,
		erc20TokenRemoteA,
		receipt,
		recipientAddress,
		bridgedAmount,
	)

	// Check that the recipient received the tokens
	balance, err := erc20TokenRemoteA.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(bridgedAmount))

	// Multi-hop transfer to Subnet B
	bridgedAmount = big.NewInt(0).Div(bridgedAmount, big.NewInt(2))
	secondaryFeeAmount := big.NewInt(0).Div(bridgedAmount, big.NewInt(4))
	utils.SendERC20TokenMultiHopAndVerify(
		ctx,
		network,
		fundedKey,
		recipientKey,
		recipientAddress,
		subnetAInfo,
		erc20TokenRemoteA,
		erc20TokenRemoteAddressA,
		subnetBInfo,
		erc20TokenRemoteB,
		erc20TokenRemoteAddressB,
		cChainInfo,
		bridgedAmount,
		secondaryFeeAmount,
	)

	// Multi-hop transfer back to Subnet A
	bridgedAmount = big.NewInt(0).Sub(bridgedAmount, secondaryFeeAmount)
	secondaryFeeAmount = big.NewInt(0).Div(bridgedAmount, big.NewInt(4))
	utils.SendERC20TokenMultiHopAndVerify(
		ctx,
		network,
		fundedKey,
		recipientKey,
		recipientAddress,
		subnetBInfo,
		erc20TokenRemoteB,
		erc20TokenRemoteAddressB,
		subnetAInfo,
		erc20TokenRemoteA,
		erc20TokenRemoteAddressA,
		cChainInfo,
		bridgedAmount,
		secondaryFeeAmount,
	)
}
