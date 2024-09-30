package flows

import (
	"context"
	"math/big"

	nativetokenhome "github.com/ava-labs/avalanche-interchain-token-transfer/abi-bindings/go/TokenHome/NativeTokenHome"
	erc20tokenremote "github.com/ava-labs/avalanche-interchain-token-transfer/abi-bindings/go/TokenRemote/ERC20TokenRemote"
	"github.com/ava-labs/avalanche-interchain-token-transfer/tests/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a NativeTokenHome on the primary network
 * Deploys ERC20TokenRemote to Subnet A
 * Transfers C-Chain native tokens to Subnet A
 * Transfer back tokens from Subnet A to C-Chain
 */
func NativeTokenHomeERC20TokenRemote(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an example WAVAX on the primary network
	wavaxAddress, wavax := utils.DeployWrappedNativeToken(
		ctx,
		fundedKey,
		cChainInfo,
		"AVAX",
	)

	// Create a NativeTokenHome for transferring the native token
	nativeTokenHomeAddress, nativeTokenHome := utils.DeployNativeTokenHome(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		wavaxAddress,
	)

	// Token representation on subnet A will have same name, symbol, and decimals
	tokenName, err := wavax.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := wavax.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := wavax.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20TokenRemote to Subnet A
	erc20TokenRemoteAddress, erc20TokenRemote := utils.DeployERC20TokenRemote(
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

	utils.RegisterERC20TokenRemoteOnHome(
		ctx,
		network,
		cChainInfo,
		nativeTokenHomeAddress,
		subnetAInfo,
		erc20TokenRemoteAddress,
	)

	// Generate new recipient to receive transferred tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to recipient on subnet A
	input := nativetokenhome.SendTokensInput{
		DestinationBlockchainID:            subnetAInfo.BlockchainID,
		DestinationTokenTransferrerAddress: erc20TokenRemoteAddress,
		Recipient:                          recipientAddress,
		PrimaryFeeTokenAddress:             wavaxAddress,
		PrimaryFee:                         big.NewInt(1e18),
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

	utils.CheckERC20TokenRemoteWithdrawal(
		ctx,
		erc20TokenRemote,
		receipt,
		recipientAddress,
		transferredAmount,
	)

	// Check that the recipient received the tokens
	balance, err := erc20TokenRemote.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(transferredAmount))

	// Fund recipient with gas tokens on subnet A
	teleporterUtils.SendNativeTransfer(
		ctx,
		subnetAInfo,
		fundedKey,
		recipientAddress,
		big.NewInt(1e18),
	)
	inputA := erc20tokenremote.SendTokensInput{
		DestinationBlockchainID:            cChainInfo.BlockchainID,
		DestinationTokenTransferrerAddress: nativeTokenHomeAddress,
		Recipient:                          recipientAddress,
		PrimaryFeeTokenAddress:             erc20TokenRemoteAddress,
		PrimaryFee:                         big.NewInt(1e10),
		SecondaryFee:                       big.NewInt(0),
		RequiredGasLimit:                   utils.DefaultNativeTokenRequiredGas,
	}

	// Send tokens on Subnet A back for native tokens on C-Chain
	receipt, transferredAmount = utils.SendERC20TokenRemote(
		ctx,
		subnetAInfo,
		erc20TokenRemote,
		erc20TokenRemoteAddress,
		inputA,
		teleporterUtils.BigIntSub(transferredAmount, inputA.PrimaryFee),
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
	utils.CheckNativeTokenHomeWithdrawal(
		ctx,
		nativeTokenHomeAddress,
		wavax,
		receipt,
		transferredAmount,
	)

	teleporterUtils.CheckBalance(ctx, recipientAddress, transferredAmount, cChainInfo.RPCClient)
}
