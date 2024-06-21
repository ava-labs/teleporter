package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	nativetokenhome "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHome/NativeTokenHome"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/gomega"
)

/**
 * Deploy a NativeTokenHome on the primary network
 * Deploys ERC20TokenRemote to Subnet A and Subnet B
 * Bridges C-Chain native tokens to Subnet A
 * Bridge tokens from Subnet A to Subnet B through multi-hop
 * Brige back tokens from Subnet B to Subnet A through multi-hop
 */
func NativeTokenHomeERC20TokenRemoteMultiHop(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, subnetBInfo := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy a NativeTokenHome on the primary network
	wavaxAddress, wavax := utils.DeployWrappedNativeToken(
		ctx,
		fundedKey,
		cChainInfo,
		"AVAX",
	)
	nativeTokenHomeAddress, nativeTokenHome := utils.DeployNativeTokenHome(
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

	// Deploy an ERC20TokenRemote on Subnet A
	erc20TokenRemoteAddressA, erc20TokenRemoteA := utils.DeployERC20TokenRemote(
		ctx,
		fundedKey,
		subnetAInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenHomeAddress,
		18,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Deploy an ERC20TokenRemote on Subnet B
	erc20TokenRemoteAddressB, erc20TokenRemoteB := utils.DeployERC20TokenRemote(
		ctx,
		fundedKey,
		subnetBInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenHomeAddress,
		18,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Register both ERC20Destinations on the NativeTokenHome
	utils.RegisterERC20TokenRemoteOnHome(
		ctx,
		network,
		cChainInfo,
		nativeTokenHomeAddress,
		subnetAInfo,
		erc20TokenRemoteAddressA,
	)

	utils.RegisterERC20TokenRemoteOnHome(
		ctx,
		network,
		cChainInfo,
		nativeTokenHomeAddress,
		subnetBInfo,
		erc20TokenRemoteAddressB,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to recipient on subnet A
	input := nativetokenhome.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: erc20TokenRemoteAddressA,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   wavaxAddress,
		PrimaryFee:               big.NewInt(1e10),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultERC20RequiredGas,
	}

	// Send the tokens and verify expected events
	amount := big.NewInt(2e18)
	receipt, bridgedAmount := utils.SendNativeTokenHome(
		ctx,
		cChainInfo,
		nativeTokenHome,
		nativeTokenHomeAddress,
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

	// Send tokens from subnet A to recipient on subnet B through a multi-hop
	secondaryFeeAmount := new(big.Int).Div(bridgedAmount, big.NewInt(4))
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
}
