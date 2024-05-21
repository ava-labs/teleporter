package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20source "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Source"
	"github.com/ava-labs/teleporter-token-bridge/tests/errors"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploys an ERC20Source contract on the C-Chain
 * Deploys an ERC20Destination contract on Subnet A
 * Check sending to non registered destination fails
 * Register the ERC20Destination to source contract
 * Check sending to non collateralized destination fails
 * Collateralize the destination
 * Check sending to collateralized destination succeeds and withdraws with correct scale.
 */
func RegistrationAndCollateralCheck(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
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

	// Deploy a NativeTokenDestination to Subnet A
	nativeTokenDestinationAddressA, _ := utils.DeployNativeTokenDestination(
		ctx,
		subnetAInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20SourceAddress,
		initialReserveImbalance,
		decimalsShift,
		multiplyOnSpoke,
		burnedFeesReportingRewardPercentage,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to Subnet A
	input := erc20source.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenDestinationAddressA,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   sourceTokenAddress,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
	}

	amount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10))

	initialBalance, err := sourceToken.BalanceOf(&bind.CallOpts{}, erc20SourceAddress)
	Expect(err).Should(BeNil())

	// Send the tokens and expect for failure since destination bridge is not registered.
	optsA, err := bind.NewKeyedTransactorWithChainID(fundedKey, cChainInfo.EVMChainID)
	Expect(err).Should(BeNil())
	_, err = erc20Source.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(Not(BeNil()))
	Expect(err.Error()).Should(ContainSubstring(errors.ErrDestinationNotRegistered))

	// Check the balance of the ERC20Source to ensure it was not changed
	balance, err := sourceToken.BalanceOf(&bind.CallOpts{}, erc20SourceAddress)
	Expect(err).Should(BeNil())
	teleporterUtils.ExpectBigEqual(balance, initialBalance)

	// Register the NativeTokenDestination to the ERC20Source
	collateralNeeded := utils.RegisterTokenDestinationOnSource(
		ctx,
		network,
		cChainInfo,
		erc20SourceAddress,
		subnetAInfo,
		nativeTokenDestinationAddressA,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		multiplyOnSpoke,
	)

	// Try sending again and expect failure since destination is not collateralized
	_, err = erc20Source.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(Not(BeNil()))
	Expect(err.Error()).Should(ContainSubstring(errors.ErrNonZeroCollateralNeeded))

	// Check the balance of the ERC20Source to ensure it was not changed
	balance, err = sourceToken.BalanceOf(&bind.CallOpts{}, erc20SourceAddress)
	Expect(err).Should(BeNil())
	teleporterUtils.ExpectBigEqual(balance, initialBalance)

	// Add collateral to the ERC20Source
	utils.AddCollateralToERC20Source(
		ctx,
		cChainInfo,
		erc20Source,
		erc20SourceAddress,
		sourceToken,
		subnetAInfo.BlockchainID,
		nativeTokenDestinationAddressA,
		collateralNeeded,
		fundedKey,
	)

	// Check the balance of the ERC20Source to ensure it was increased by the collateral amount.
	// Also set the new initial balance before sending tokens
	balance, err = sourceToken.BalanceOf(&bind.CallOpts{}, erc20SourceAddress)
	Expect(err).Should(BeNil())
	teleporterUtils.ExpectBigEqual(balance, initialBalance.Add(initialBalance, collateralNeeded))

	// Send the tokens and expect success now that collateral is added
	teleporterUtils.ERC20Approve(
		ctx,
		sourceToken,
		erc20SourceAddress,
		big.NewInt(0).Add(amount, input.PrimaryFee),
		cChainInfo,
		fundedKey,
	)
	tx, err := erc20Source.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, cChainInfo, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20Source.ParseTokensSent)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(fundedKey.PublicKey)))

	// Compute the scaled amount
	scaledAmount := utils.GetScaledAmountFromERC20Source(
		erc20Source,
		input.DestinationBlockchainID,
		input.DestinationBridgeAddress,
		amount,
	)
	teleporterUtils.ExpectBigEqual(event.Amount, scaledAmount)

	// Check the balance of the ERC20Source increased by the bridged amount
	balance, err = sourceToken.BalanceOf(&bind.CallOpts{}, erc20SourceAddress)
	Expect(err).Should(BeNil())
	teleporterUtils.ExpectBigEqual(balance, big.NewInt(0).Add(initialBalance, amount))

	// Relay the message to subnet A and check for a native token mint withdrawal
	network.RelayMessage(
		ctx,
		receipt,
		cChainInfo,
		subnetAInfo,
		true,
	)

	// Verify the recipient received the tokens
	teleporterUtils.CheckBalance(ctx, recipientAddress, scaledAmount, subnetAInfo.RPCClient)
}
