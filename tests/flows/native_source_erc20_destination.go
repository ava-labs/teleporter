package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20destination "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Destination"
	nativetokensource "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/NativeTokenSource"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

// NativeSourceERC20Destination deploys NativeTokenSource to subnet A and ERC20Destination to subnet B.
// Then it sends native tokens from subnet A for an erc20 token on subnet B, and transfers back to
// subnet A for its original native tokens.
func NativeSourceERC20Destination(network interfaces.Network) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy a native token source on subnet A
	// TODO: for now use empty fee token address, later on need a token like WAVAX
	wavaxAddress, wavax := utils.DeployExampleWAVAX(
		ctx,
		fundedKey,
		subnetAInfo,
	)

	nativeTokenSourceAddress, nativeTokenSource := utils.DeployNativeTokenSource(
		ctx,
		fundedKey,
		subnetAInfo,
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

	// Deploy an ERC20Destination for the token source on subnet A
	erc20DestinationAddress, erc20Destination := utils.DeployERC20Destination(
		ctx,
		fundedKey,
		subnetBInfo,
		fundedAddress,
		subnetAInfo.BlockchainID,
		nativeTokenSourceAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from subnet A to recipient on subnet B
	input := nativetokensource.SendTokensInput{
		DestinationBlockchainID:  subnetBInfo.BlockchainID,
		DestinationBridgeAddress: erc20DestinationAddress,
		Recipient:                recipientAddress,
		PrimaryFee:               big.NewInt(0),
		SecondaryFee:             big.NewInt(0),
		AllowedRelayerAddresses:  []common.Address{},
	}

	// Send the tokens and verify expected events
	amount := big.NewInt(2e18)
	receipt, bridgedAmount := utils.SendNativeTokenSource(
		ctx,
		subnetAInfo,
		nativeTokenSource,
		nativeTokenSourceAddress,
		input,
		amount,
		fundedKey,
	)

	// Relay the message to subnet B and check for message delivery
	receipt = network.RelayMessage(
		ctx,
		receipt,
		subnetAInfo,
		subnetBInfo,
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

	// Fund recipient with gas tokens on subnet B
	teleporterUtils.SendNativeTransfer(
		ctx,
		subnetBInfo,
		fundedKey,
		recipientAddress,
		big.NewInt(1e18),
	)
	inputB := erc20destination.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenSourceAddress,
		Recipient:                recipientAddress,
		PrimaryFee:               big.NewInt(0),
		SecondaryFee:             big.NewInt(0),
		AllowedRelayerAddresses:  []common.Address{},
	}

	// Send tokens on subnet B back for native tokens on subnet A
	receipt, bridgedAmount = utils.SendERC20Destination(
		ctx,
		subnetBInfo,
		erc20Destination,
		erc20DestinationAddress,
		inputB,
		bridgedAmount,
		recipientKey,
	)

	receipt = network.RelayMessage(
		ctx,
		receipt,
		subnetBInfo,
		subnetAInfo,
		true,
	)

	// Check that the recipient received the tokens
	utils.CheckNativeTokenSourceWithdrawal(
		ctx,
		nativeTokenSourceAddress,
		wavax,
		receipt,
		bridgedAmount,
	)

	teleporterUtils.CheckBalance(ctx, recipientAddress, bridgedAmount, subnetAInfo.RPCClient)
}
