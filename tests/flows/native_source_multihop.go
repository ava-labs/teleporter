package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	nativetokensource "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/NativeTokenSource"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/gomega"
)

/**
 * Deploy a native token source on the primary network
 * Deploys ERC20Destination to Subnet A and Subnet B
 * Bridges C-Chain native tokens to Subnet A
 * Bridge tokens from Subnet A to Subnet B through multihop
 * Brige back tokens from Subnet B to Subnet A through multihop
 */
func NativeTokenSourceMultihop(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, subnetBInfo := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy a native token source on the primary network
	wavaxAddress, wavax := utils.DeployExampleWAVAX(
		ctx,
		fundedKey,
		cChainInfo,
	)

	nativeTokenSourceAddress, nativeTokenSource := utils.DeployNativeTokenSource(
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

	// Deploy an ERC20Destination on Subnet A
	erc20DestinationAddress_A, erc20Destination_A := utils.DeployERC20Destination(
		ctx,
		fundedKey,
		subnetAInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenSourceAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Deploy an ERC20Destination on Subnet B
	erc20DestinationAddress_B, erc20Destination_B := utils.DeployERC20Destination(
		ctx,
		fundedKey,
		subnetBInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
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
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: erc20DestinationAddress_A,
		Recipient:                recipientAddress,
		PrimaryFee:               big.NewInt(0),
		SecondaryFee:             big.NewInt(0),
		AllowedRelayerAddresses:  []common.Address{},
	}

	// Send the tokens and verify expected events
	amount := big.NewInt(2e18)
	receipt, bridgedAmount := utils.SendNativeTokenSource(
		ctx,
		cChainInfo,
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

	// Send tokens from subnet A to recipient on subnet B through a multihop
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
}
