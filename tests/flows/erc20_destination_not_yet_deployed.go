package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20source "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Source"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy an ERC20 token source on the primary network
 * Defers deploying ERC20Destination to Subnet A
 * Bridges C-Chain example ERC20 tokens to Subnet A
 * Confirms failure to execute message, since destination hasn't been deployed.
 * Deploys ERC20Destination to Subnet A
 * Invokes a retry of message execution
 * Confirms message execution
 */
func ERC20DestinationNotYetDeployed(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on the primary network as the source token to be bridged
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

	// Token representation on subnet A will have same name, symbol, and decimals
	tokenName, err := sourceToken.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := sourceToken.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := sourceToken.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Pre-determine the deployment address of the ERC20Destination, but
	// don't deploy it yet. the address is based on 1 + the current nonce,
	// since the fundedAddress account will perform one more transaction
	// (relaying the message) prior to actually deploying the
	// ERC20Destination.
	erc20DestinationAddressNonce, err := subnetAInfo.RPCClient.NonceAt(
		ctx,
		fundedAddress,
		nil,
	)
	Expect(err).Should(BeNil())
	erc20DestinationAddressNonce += 1
	erc20DestinationAddress := crypto.CreateAddress(
		fundedAddress,
		erc20DestinationAddressNonce,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to recipient on subnet A
	input := erc20source.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: erc20DestinationAddress,
		Recipient:                recipientAddress,
		FeeTokenAddress:          sourceTokenAddress,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultERC20RequiredGas,
	}
	amount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(13))

	receipt, bridgedAmount := utils.SendERC20Source(
		ctx,
		cChainInfo,
		erc20Source,
		erc20SourceAddress,
		sourceToken,
		input,
		amount,
		fundedKey,
	)

	msgSentEvent, err := teleporterUtils.GetEventFromLogs(
		receipt.Logs,
		cChainInfo.TeleporterMessenger.ParseSendCrossChainMessage,
	)
	Expect(err).Should(BeNil())
	msgSent := msgSentEvent.Message

	// Relay the message to Subnet A and check for message delivery
	receipt = network.RelayMessage(
		ctx,
		receipt,
		cChainInfo,
		subnetAInfo,
		true,
	)

	// Confirm that the relayer's subnetA transaction emitted the
	// MessageExecutionFailed event
	_, err = teleporterUtils.GetEventFromLogs(
		receipt.Logs,
		subnetAInfo.TeleporterMessenger.ParseMessageExecutionFailed,
	)
	Expect(err).Should(BeNil())

	// Deploy an ERC20Destination to Subnet A
	Expect( // Just to ensure no more subnetA fundedAddr tx's have occurred
		subnetAInfo.RPCClient.NonceAt(ctx, fundedAddress, nil),
	).Should(Equal(erc20DestinationAddressNonce))
	actualERC20DestDeploymentAddress, erc20Destination := utils.DeployERC20Destination(
		ctx,
		fundedKey,
		subnetAInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20SourceAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)
	Expect(actualERC20DestDeploymentAddress).Should(Equal(erc20DestinationAddress))

	// Retry the message execution on the destination TeleporterMessenger
	// instance now that the ERC20Destination contract is deployed.
	opts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey,
		subnetAInfo.EVMChainID,
	)
	Expect(err).Should(BeNil())
	tx, err := subnetAInfo.TeleporterMessenger.RetryMessageExecution(
		opts,
		cChainInfo.BlockchainID,
		msgSent,
	)
	Expect(err).Should(BeNil())
	receipt = teleporterUtils.WaitForTransactionSuccess(ctx, subnetAInfo, tx.Hash())

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
}
