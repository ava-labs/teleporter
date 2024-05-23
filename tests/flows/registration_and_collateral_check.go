package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20tokenhub "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHub/ERC20TokenHub"
	"github.com/ava-labs/teleporter-token-bridge/tests/errors"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploys an ERC20TokenHub contract on the C-Chain
 * Deploys an ERC20TokenSpoke contract on Subnet A
 * Check sending to  unregistered spoke fails
 * Register the ERC20TokenSpoke to hub contract
 * Check sending to non-collateralized spoke fails
 * Collateralize the spoke
 * Check sending to collateralized spoke succeeds and withdraws with correct scale.
 */
func RegistrationAndCollateralCheck(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
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

	// Deploy a NativeTokenSpoke to Subnet A
	nativeTokenSpokeAddressA, _ := utils.DeployNativeTokenSpoke(
		ctx,
		subnetAInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHubAddress,
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
	input := erc20tokenhub.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenSpokeAddressA,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   exampleERC20Address,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
	}

	amount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10))

	initialBalance, err := exampleERC20.BalanceOf(&bind.CallOpts{}, erc20TokenHubAddress)
	Expect(err).Should(BeNil())

	// Send the tokens and expect for failure since spoke instance is not registered.
	optsA, err := bind.NewKeyedTransactorWithChainID(fundedKey, cChainInfo.EVMChainID)
	Expect(err).Should(BeNil())
	_, err = erc20TokenHub.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(Not(BeNil()))
	Expect(err.Error()).Should(ContainSubstring(errors.ErrSpokeNotRegistered))

	// Check the balance of the ERC20TokenHub to ensure it was not changed
	balance, err := exampleERC20.BalanceOf(&bind.CallOpts{}, erc20TokenHubAddress)
	Expect(err).Should(BeNil())
	teleporterUtils.ExpectBigEqual(balance, initialBalance)

	// Register the NativeTokenSpoke to the ERC20TokenHub
	collateralNeeded := utils.RegisterTokenSpokeOnHub(
		ctx,
		network,
		cChainInfo,
		erc20TokenHubAddress,
		subnetAInfo,
		nativeTokenSpokeAddressA,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		multiplyOnSpoke,
	)

	// Try sending again and expect failure since spoke is not collateralized
	_, err = erc20TokenHub.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(Not(BeNil()))
	Expect(err.Error()).Should(ContainSubstring(errors.ErrNonZeroCollateralNeeded))

	// Check the balance of the ERC20TokenHub to ensure it was not changed
	balance, err = exampleERC20.BalanceOf(&bind.CallOpts{}, erc20TokenHubAddress)
	Expect(err).Should(BeNil())
	teleporterUtils.ExpectBigEqual(balance, initialBalance)

	// Add collateral to the ERC20TokenHub
	utils.AddCollateralToERC20TokenHub(
		ctx,
		cChainInfo,
		erc20TokenHub,
		erc20TokenHubAddress,
		exampleERC20,
		subnetAInfo.BlockchainID,
		nativeTokenSpokeAddressA,
		collateralNeeded,
		fundedKey,
	)

	// Check the balance of the ERC20TokenHub to ensure it was increased by the collateral amount.
	// Also set the new initial balance before sending tokens
	balance, err = exampleERC20.BalanceOf(&bind.CallOpts{}, erc20TokenHubAddress)
	Expect(err).Should(BeNil())
	teleporterUtils.ExpectBigEqual(balance, initialBalance.Add(initialBalance, collateralNeeded))

	// Send the tokens and expect success now that collateral is added
	teleporterUtils.ERC20Approve(
		ctx,
		exampleERC20,
		erc20TokenHubAddress,
		big.NewInt(0).Add(amount, input.PrimaryFee),
		cChainInfo,
		fundedKey,
	)
	tx, err := erc20TokenHub.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, cChainInfo, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20TokenHub.ParseTokensSent)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(fundedKey.PublicKey)))

	// Compute the scaled amount
	scaledAmount := utils.GetScaledAmountFromERC20TokenHub(
		erc20TokenHub,
		input.DestinationBlockchainID,
		input.DestinationBridgeAddress,
		amount,
	)
	teleporterUtils.ExpectBigEqual(event.Amount, scaledAmount)

	// Check the balance of the ERC20TokenHub increased by the bridged amount
	balance, err = exampleERC20.BalanceOf(&bind.CallOpts{}, erc20TokenHubAddress)
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
