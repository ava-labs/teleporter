package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20tokenhome "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHome/ERC20TokenHome"
	nativetokenremote "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenRemote/NativeTokenRemote"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

var (
	decimalsShift           = uint8(1)
	tokenMultiplier         = utils.GetTokenMultiplier(decimalsShift)
	initialReserveImbalance = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e6))

	// These two should be changed together
	multiplyOnRemote       = true
	erc20TokenHomeDecimals = utils.NativeTokenDecimals - decimalsShift

	burnedFeesReportingRewardPercentage = big.NewInt(1)
)

/**
 * Deploy a ERC20Token on the primary network
 * Deploys NativeTokenRemote to Subnet A and Subnet B
 * Bridges C-Chain example ERC20 tokens to Subnet A as Subnet A's native token
 * Bridge back tokens from Subnet A to C-Chain
 */
func ERC20TokenHomeNativeTokenRemote(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on subnet A as the token to be bridged
	exampleERC20Address, exampleERC20 := utils.DeployExampleERC20(
		ctx,
		fundedKey,
		cChainInfo,
		erc20TokenHomeDecimals,
	)

	exampleERC20Decimals, err := exampleERC20.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Create an ERC20TokenHome for bridging the ERC20 token
	erc20TokenHomeAddress, erc20TokenHome := utils.DeployERC20TokenHome(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		exampleERC20Address,
		exampleERC20Decimals,
	)

	// Deploy a NativeTokenRemote to Subnet A
	nativeTokenRemoteAddressA, nativeTokenRemoteA := utils.DeployNativeTokenRemote(
		ctx,
		subnetAInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHomeAddress,
		exampleERC20Decimals,
		initialReserveImbalance,
		multiplyOnRemote,
		burnedFeesReportingRewardPercentage,
	)

	collateralAmount := utils.RegisterTokenRemoteOnHome(
		ctx,
		network,
		cChainInfo,
		erc20TokenHomeAddress,
		subnetAInfo,
		nativeTokenRemoteAddressA,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		multiplyOnRemote,
	)

	utils.AddCollateralToERC20TokenHome(
		ctx,
		cChainInfo,
		erc20TokenHome,
		erc20TokenHomeAddress,
		exampleERC20,
		subnetAInfo.BlockchainID,
		nativeTokenRemoteAddressA,
		collateralAmount,
		fundedKey,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	recipientKey.ECDH()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to Subnet A
	input := erc20tokenhome.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenRemoteAddressA,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   exampleERC20Address,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
	}

	amount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10))
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

	// Relay the message to subnet A and check for a native token mint withdrawal
	network.RelayMessage(
		ctx,
		receipt,
		cChainInfo,
		subnetAInfo,
		true,
	)

	// Verify the recipient received the tokens
	teleporterUtils.CheckBalance(ctx, recipientAddress, bridgedAmount, subnetAInfo.RPCClient)

	// Send back to the home chain and check that ERC20TokenHome received the tokens
	input_A := nativetokenremote.SendTokensInput{
		DestinationBlockchainID:  cChainInfo.BlockchainID,
		DestinationBridgeAddress: erc20TokenHomeAddress,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   nativeTokenRemoteAddressA,
		PrimaryFee:               big.NewInt(1e10),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
	}
	// Send half of the received amount to account for gas expenses
	amountToSendA := new(big.Int).Div(bridgedAmount, big.NewInt(2))
	receipt, bridgedAmount = utils.SendNativeTokenRemote(
		ctx,
		subnetAInfo,
		nativeTokenRemoteA,
		nativeTokenRemoteAddressA,
		input_A,
		amountToSendA,
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
	scaledAmount := utils.RemoveTokenScaling(tokenMultiplier, multiplyOnRemote, bridgedAmount)
	utils.CheckERC20TokenHomeWithdrawal(
		ctx,
		erc20TokenHomeAddress,
		exampleERC20,
		receipt,
		recipientAddress,
		scaledAmount,
	)

	// Check that the recipient received the tokens
	balance, err := exampleERC20.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(scaledAmount))
}
