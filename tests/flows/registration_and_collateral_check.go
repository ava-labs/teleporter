package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20tokenhome "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHome/ERC20TokenHome"
	"github.com/ava-labs/teleporter-token-bridge/tests/errors"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploys an ERC20TokenHome contract on the C-Chain
 * Deploys an ERC20TokenRemote contract on Subnet A
 * Check sending to  unregistered remote fails
 * Register the ERC20TokenRemote to home contract
 * Check sending to non-collateralized remote fails
 * Collateralize the remote
 * Check sending to collateralized remote succeeds and withdraws with correct scale.
 */
func RegistrationAndCollateralCheck(network interfaces.Network) {
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

	// Create an ERC20TokenHome for bridging the ERC20 token
	erc20TokenHomeAddress, erc20TokenHome := utils.DeployERC20TokenHome(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		exampleERC20Address,
		erc20TokenHomeDecimals,
	)

	// Deploy a NativeTokenRemote to Subnet A
	nativeTokenRemoteAddressA, _ := utils.DeployNativeTokenRemote(
		ctx,
		subnetAInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHomeAddress,
		erc20TokenHomeDecimals,
		initialReserveImbalance,
		multiplyOnRemote,
		burnedFeesReportingRewardPercentage,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
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

	initialBalance, err := exampleERC20.BalanceOf(&bind.CallOpts{}, erc20TokenHomeAddress)
	Expect(err).Should(BeNil())

	// Send the tokens and expect for failure since TokenRemote instance is not registered.
	optsA, err := bind.NewKeyedTransactorWithChainID(fundedKey, cChainInfo.EVMChainID)
	Expect(err).Should(BeNil())
	_, err = erc20TokenHome.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(Not(BeNil()))
	Expect(err.Error()).Should(ContainSubstring(errors.ErrRemoteNotRegistered))

	// Check the balance of the ERC20TokenHome to ensure it was not changed
	balance, err := exampleERC20.BalanceOf(&bind.CallOpts{}, erc20TokenHomeAddress)
	Expect(err).Should(BeNil())
	teleporterUtils.ExpectBigEqual(balance, initialBalance)

	// Register the NativeTokenRemote to the ERC20TokenHome
	collateralNeeded := utils.RegisterTokenRemoteOnHome(
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

	// Try sending again and expect failure since remote is not collateralized
	_, err = erc20TokenHome.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(Not(BeNil()))
	Expect(err.Error()).Should(ContainSubstring(errors.ErrNonZeroCollateralNeeded))

	// Check the balance of the ERC20TokenHome to ensure it was not changed
	balance, err = exampleERC20.BalanceOf(&bind.CallOpts{}, erc20TokenHomeAddress)
	Expect(err).Should(BeNil())
	teleporterUtils.ExpectBigEqual(balance, initialBalance)

	// Add collateral to the ERC20TokenHome
	utils.AddCollateralToERC20TokenHome(
		ctx,
		cChainInfo,
		erc20TokenHome,
		erc20TokenHomeAddress,
		exampleERC20,
		subnetAInfo.BlockchainID,
		nativeTokenRemoteAddressA,
		collateralNeeded,
		fundedKey,
	)

	// Check the balance of the ERC20TokenHome to ensure it was increased by the collateral amount.
	// Also set the new initial balance before sending tokens
	balance, err = exampleERC20.BalanceOf(&bind.CallOpts{}, erc20TokenHomeAddress)
	Expect(err).Should(BeNil())
	teleporterUtils.ExpectBigEqual(balance, initialBalance.Add(initialBalance, collateralNeeded))

	// Send the tokens and expect success now that collateral is added
	utils.ERC20Approve(
		ctx,
		exampleERC20,
		erc20TokenHomeAddress,
		big.NewInt(0).Add(amount, input.PrimaryFee),
		cChainInfo,
		fundedKey,
	)
	tx, err := erc20TokenHome.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, cChainInfo, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20TokenHome.ParseTokensSent)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(fundedKey.PublicKey)))

	// Compute the scaled amount
	scaledAmount := utils.GetScaledAmountFromERC20TokenHome(
		erc20TokenHome,
		input.DestinationBlockchainID,
		input.DestinationBridgeAddress,
		amount,
	)
	teleporterUtils.ExpectBigEqual(event.Amount, scaledAmount)

	// Check the balance of the ERC20TokenHome increased by the bridged amount
	balance, err = exampleERC20.BalanceOf(&bind.CallOpts{}, erc20TokenHomeAddress)
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
