package ictt

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	nativetokenhome "github.com/ava-labs/teleporter/abi-bindings/go/ictt/TokenHome/NativeTokenHome"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/gomega"
)

/**
 * Deploy a NativeTokenHome on the primary network
 * Deploys ERC20TokenRemote to Subnet A and Subnet B
 * Transfers C-Chain native tokens to Subnet A
 * Transfer tokens from Subnet A to Subnet B through multi-hop
 * Brige back tokens from Subnet B to Subnet A through multi-hop
 */
func NativeTokenHomeERC20TokenRemoteMultiHop(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, subnetBInfo := network.GetTwoSubnets()
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
		teleporter,
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
		teleporter,
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
		teleporter,
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
		teleporter,
		cChainInfo,
		nativeTokenHomeAddress,
		subnetAInfo,
		erc20TokenRemoteAddressA,
		fundedKey,
	)

	utils.RegisterERC20TokenRemoteOnHome(
		ctx,
		teleporter,
		cChainInfo,
		nativeTokenHomeAddress,
		subnetBInfo,
		erc20TokenRemoteAddressB,
		fundedKey,
	)

	// Generate new recipient to receive transferred tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to recipient on subnet A
	input := nativetokenhome.SendTokensInput{
		DestinationBlockchainID:            subnetAInfo.BlockchainID,
		DestinationTokenTransferrerAddress: erc20TokenRemoteAddressA,
		Recipient:                          recipientAddress,
		PrimaryFeeTokenAddress:             wavaxAddress,
		PrimaryFee:                         big.NewInt(1e10),
		SecondaryFee:                       big.NewInt(0),
		RequiredGasLimit:                   utils.DefaultERC20RequiredGas,
	}

	// Send the tokens and verify expected events
	amount := big.NewInt(2e18)
	receipt, transferredAmount := utils.SendNativeTokenHome(
		ctx,
		cChainInfo,
		nativeTokenHome,
		nativeTokenHomeAddress,
		wavax,
		input,
		utils.BigIntSub(amount, input.PrimaryFee),
		fundedKey,
	)

	// Relay the message to subnet B and check for message delivery
	receipt = teleporter.RelayTeleporterMessage(
		ctx,
		receipt,
		cChainInfo,
		subnetAInfo,
		true,
		fundedKey,
	)

	utils.CheckERC20TokenRemoteWithdrawal(
		ctx,
		erc20TokenRemoteA,
		receipt,
		recipientAddress,
		transferredAmount,
	)

	// Check that the recipient received the tokens
	balance, err := erc20TokenRemoteA.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(transferredAmount))

	// Send tokens from subnet A to recipient on subnet B through a multi-hop
	secondaryFeeAmount := new(big.Int).Div(transferredAmount, big.NewInt(4))
	utils.SendERC20TokenMultiHopAndVerify(
		ctx,
		teleporter,
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
		transferredAmount,
		secondaryFeeAmount,
	)
}
