package flows

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	bridgetoken "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/examples/ERC20Bridge/BridgeToken"
	erc20bridge "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/examples/ERC20Bridge/ERC20Bridge"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func ERC20BridgeMultihop(network interfaces.Network) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, subnetCInfo := utils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	// Deploy an ERC20 to subnet A
	nativeERC20Address, nativeERC20 := utils.DeployExampleERC20(
		context.Background(),
		fundedKey,
		subnetAInfo,
	)

	// Deploy the ERC20 bridge to subnet A
	erc20BridgeAddressA, erc20BridgeA := utils.DeployERC20Bridge(
		ctx,
		fundedKey,
		fundedAddress,
		subnetAInfo,
	)
	// Deploy the ERC20 bridge to subnet B
	erc20BridgeAddressB, erc20BridgeB := utils.DeployERC20Bridge(
		ctx,
		fundedKey,
		fundedAddress,
		subnetBInfo,
	)
	// Deploy the ERC20 bridge to subnet C
	erc20BridgeAddressC, erc20BridgeC := utils.DeployERC20Bridge(
		ctx,
		fundedKey,
		fundedAddress,
		subnetCInfo,
	)

	amount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10000000000000))
	utils.ERC20Approve(
		ctx,
		nativeERC20,
		erc20BridgeAddressA,
		amount,
		subnetAInfo,
		fundedKey,
	)

	// Send a transaction on Subnet A to add support for the the ERC20 token to the bridge on Subnet B
	receipt, messageID := submitCreateBridgeToken(
		ctx,
		subnetAInfo,
		subnetBInfo.BlockchainID,
		erc20BridgeAddressB,
		nativeERC20Address,
		nativeERC20Address,
		big.NewInt(0),
		fundedAddress,
		fundedKey,
		erc20BridgeA,
		subnetAInfo.TeleporterMessenger,
	)

	// Relay message
	network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)

	// Check Teleporter message received on the destination
	delivered, err := subnetBInfo.TeleporterMessenger.MessageReceived(
		&bind.CallOpts{},
		messageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check the bridge token was added on Subnet B
	bridgeTokenSubnetBAddress, err := erc20BridgeB.NativeToWrappedTokens(
		&bind.CallOpts{},
		subnetAInfo.BlockchainID,
		erc20BridgeAddressA,
		nativeERC20Address,
	)
	Expect(err).Should(BeNil())
	Expect(bridgeTokenSubnetBAddress).ShouldNot(Equal(common.Address{}))
	bridgeTokenB, err := bridgetoken.NewBridgeToken(bridgeTokenSubnetBAddress, subnetBInfo.RPCClient)
	Expect(err).Should(BeNil())

	// Check all the settings of the new bridge token are correct.
	actualNativeChainID, err := bridgeTokenB.NativeBlockchainID(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(actualNativeChainID[:]).Should(Equal(subnetAInfo.BlockchainID[:]))

	actualNativeBridgeAddress, err := bridgeTokenB.NativeBridge(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(actualNativeBridgeAddress).Should(Equal(erc20BridgeAddressA))

	actualNativeAssetAddress, err := bridgeTokenB.NativeAsset(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(actualNativeAssetAddress).Should(Equal(nativeERC20Address))

	actualName, err := bridgeTokenB.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(actualName).Should(Equal("Mock Token"))

	actualSymbol, err := bridgeTokenB.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(actualSymbol).Should(Equal("EXMP"))

	actualDecimals, err := bridgeTokenB.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(actualDecimals).Should(Equal(uint8(18)))

	// Send a transaction on Subnet A to add support for the the ERC20 token to the bridge on Subnet C
	receipt, messageID = submitCreateBridgeToken(
		ctx,
		subnetAInfo,
		subnetCInfo.BlockchainID,
		erc20BridgeAddressC,
		nativeERC20Address,
		nativeERC20Address,
		big.NewInt(0),
		fundedAddress,
		fundedKey,
		erc20BridgeA,
		subnetAInfo.TeleporterMessenger,
	)

	// Relay message
	network.RelayMessage(ctx, receipt, subnetAInfo, subnetCInfo, true)

	// Check Teleporter message received on the destination
	delivered, err =
		subnetCInfo.TeleporterMessenger.MessageReceived(
			&bind.CallOpts{},
			messageID,
		)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check the bridge token was added on Subnet C
	bridgeTokenSubnetCAddress, err := erc20BridgeC.NativeToWrappedTokens(
		&bind.CallOpts{},
		subnetAInfo.BlockchainID,
		erc20BridgeAddressA,
		nativeERC20Address,
	)
	Expect(err).Should(BeNil())
	Expect(bridgeTokenSubnetCAddress).ShouldNot(Equal(common.Address{}))
	bridgeTokenC, err := bridgetoken.NewBridgeToken(bridgeTokenSubnetCAddress, subnetCInfo.RPCClient)
	Expect(err).Should(BeNil())

	// Send a bridge transfer for the newly added token from subnet A to subnet B
	totalAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(13))
	primaryFeeAmount := big.NewInt(1e18)
	receipt, messageID = bridgeToken(
		ctx,
		subnetAInfo,
		subnetBInfo.BlockchainID,
		erc20BridgeAddressB,
		nativeERC20Address,
		fundedAddress,
		totalAmount,
		primaryFeeAmount,
		big.NewInt(0),
		fundedAddress,
		fundedKey,
		erc20BridgeA,
		true,
		subnetAInfo.BlockchainID,
		subnetAInfo.TeleporterMessenger,
	)

	// Relay message
	deliveryReceipt := network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)
	receiveEvent, err := utils.GetEventFromLogs(
		deliveryReceipt.Logs,
		subnetBInfo.TeleporterMessenger.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())

	// Check Teleporter message received on the destination
	delivered, err =
		subnetBInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check the recipient balance of the new bridge token.
	actualRecipientBalance, err := bridgeTokenB.BalanceOf(&bind.CallOpts{}, fundedAddress)
	Expect(err).Should(BeNil())
	Expect(actualRecipientBalance).Should(Equal(totalAmount.Sub(totalAmount, primaryFeeAmount)))

	// Approve the bridge contract on subnet B to spend the wrapped tokens in the user account.
	approveBridgeToken(
		ctx,
		subnetBInfo,
		bridgeTokenSubnetBAddress,
		bridgeTokenB,
		amount,
		erc20BridgeAddressB,
		fundedAddress,
		fundedKey,
	)

	// Check the initial relayer reward amount on SubnetA.
	currentRewardAmount, err := subnetAInfo.TeleporterMessenger.CheckRelayerRewardAmount(
		&bind.CallOpts{},
		receiveEvent.RewardRedeemer,
		nativeERC20Address)
	Expect(err).Should(BeNil())

	// Unwrap bridged tokens back to subnet A, then wrap tokens to final destination on subnet C
	totalAmount = big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(11))
	secondaryFeeAmount := big.NewInt(1e18)
	receipt, messageID = bridgeToken(
		ctx,
		subnetBInfo,
		subnetCInfo.BlockchainID,
		erc20BridgeAddressC,
		bridgeTokenSubnetBAddress,
		fundedAddress,
		totalAmount,
		primaryFeeAmount,
		secondaryFeeAmount,
		fundedAddress,
		fundedKey,
		erc20BridgeB,
		false,
		subnetAInfo.BlockchainID,
		subnetBInfo.TeleporterMessenger,
	)

	// Relay message from SubnetB to SubnetA
	// The receipt of transaction that delivers the message will also have the "second hop"
	// message sent from subnet A to subnet C.
	receipt = network.RelayMessage(ctx, receipt, subnetBInfo, subnetAInfo, true)

	// Check Teleporter message received on the destination
	delivered, err =
		subnetAInfo.TeleporterMessenger.MessageReceived(
			&bind.CallOpts{},
			messageID,
		)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Get the sendCrossChainMessage event from SubnetA to SubnetC, which should be present in
	// the receipt of the transaction that delivered the first message from SubnetB to SubnetA.
	event, err := utils.GetEventFromLogs(receipt.Logs,
		subnetAInfo.TeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(subnetCInfo.BlockchainID[:]))
	messageID = event.MessageID

	// Check the redeemable reward balance of the relayer if the relayer address was set.
	// If this is an external network, skip this check since it depends on the initial state of the receipt
	// queue prior to the test run.
	if !network.IsExternalNetwork() {
		updatedRewardAmount, err :=
			subnetAInfo.TeleporterMessenger.CheckRelayerRewardAmount(
				&bind.CallOpts{},
				receiveEvent.RewardRedeemer,
				nativeERC20Address,
			)
		Expect(err).Should(BeNil())
		Expect(updatedRewardAmount).Should(Equal(new(big.Int).Add(currentRewardAmount, primaryFeeAmount)))
	}

	// Relay message from SubnetA to SubnetC
	deliveryReceipt = network.RelayMessage(ctx, receipt, subnetAInfo, subnetCInfo, true)
	receiveEvent, err = utils.GetEventFromLogs(
		deliveryReceipt.Logs,
		subnetCInfo.TeleporterMessenger.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())

	// Check Teleporter message received on the destination
	delivered, err =
		subnetCInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	actualRecipientBalance, err = bridgeTokenC.BalanceOf(&bind.CallOpts{}, fundedAddress)
	Expect(err).Should(BeNil())
	expectedAmount := totalAmount.Sub(totalAmount, primaryFeeAmount).Sub(totalAmount, secondaryFeeAmount)
	Expect(actualRecipientBalance).Should(Equal(expectedAmount))

	// Approve the bridge contract on Subnet C to spend the bridge tokens from the user account
	approveBridgeToken(
		ctx,
		subnetCInfo,
		bridgeTokenSubnetCAddress,
		bridgeTokenC,
		amount,
		erc20BridgeAddressC,
		fundedAddress,
		fundedKey)

	// Get the current relayer reward amount on SubnetA.
	currentRewardAmount, err = subnetAInfo.TeleporterMessenger.CheckRelayerRewardAmount(
		&bind.CallOpts{},
		receiveEvent.RewardRedeemer,
		nativeERC20Address)
	Expect(err).Should(BeNil())

	// Send a transaction to unwrap tokens from Subnet C back to Subnet A
	totalAmount = big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(8))
	receipt, messageID = bridgeToken(
		ctx,
		subnetCInfo,
		subnetAInfo.BlockchainID,
		erc20BridgeAddressA,
		bridgeTokenSubnetCAddress,
		fundedAddress,
		totalAmount,
		primaryFeeAmount,
		big.NewInt(0),
		fundedAddress,
		fundedKey,
		erc20BridgeC,
		false,
		subnetAInfo.BlockchainID,
		subnetCInfo.TeleporterMessenger,
	)

	// Relay message from SubnetC to SubnetA
	network.RelayMessage(ctx, receipt, subnetCInfo, subnetAInfo, true)

	// Check Teleporter message received on the destination
	delivered, err =
		subnetAInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check the balance of the native token after the unwrap
	actualNativeTokenDefaultAccountBalance, err := nativeERC20.BalanceOf(&bind.CallOpts{}, fundedAddress)
	Expect(err).Should(BeNil())
	expectedAmount = big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(9999999994))
	Expect(actualNativeTokenDefaultAccountBalance).Should(Equal(expectedAmount))

	// Check the balance of the native token for the relayer, which should have received the fee rewards
	// If this is an external network, skip this check since it depends on the initial state of the receipt
	// queue prior to the test run.
	if !network.IsExternalNetwork() {
		updatedRewardAmount, err :=
			subnetAInfo.TeleporterMessenger.CheckRelayerRewardAmount(
				&bind.CallOpts{},
				receiveEvent.RewardRedeemer,
				nativeERC20Address,
			)
		Expect(err).Should(BeNil())
		Expect(updatedRewardAmount).Should(Equal(new(big.Int).Add(currentRewardAmount, secondaryFeeAmount)))
	}
}

func submitCreateBridgeToken(
	ctx context.Context,
	source interfaces.SubnetTestInfo,
	destinationBlockchainID ids.ID,
	destinationBridgeAddress common.Address,
	nativeToken common.Address,
	messageFeeAsset common.Address,
	messageFeeAmount *big.Int,
	fundedAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	transactor *erc20bridge.ERC20Bridge,
	teleporterMessenger *teleportermessenger.TeleporterMessenger,
) (*types.Receipt, ids.ID) {
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := transactor.SubmitCreateBridgeToken(
		opts,
		destinationBlockchainID,
		destinationBridgeAddress,
		nativeToken,
		messageFeeAsset,
		messageFeeAmount,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt := utils.WaitForTransactionSuccess(ctx, source, tx.Hash())

	event, err := utils.GetEventFromLogs(receipt.Logs, teleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(destinationBlockchainID[:]))

	log.Info("Successfully SubmitCreateBridgeToken",
		"txHash", tx.Hash().Hex(),
		"messageID", hex.EncodeToString(event.MessageID[:]))

	return receipt, event.MessageID
}

func bridgeToken(
	ctx context.Context,
	source interfaces.SubnetTestInfo,
	destinationBlockchainID ids.ID,
	destinationBridgeAddress common.Address,
	token common.Address,
	recipient common.Address,
	totalAmount *big.Int,
	primaryFeeAmount *big.Int,
	secondaryFeeAmount *big.Int,
	fundedAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	transactor *erc20bridge.ERC20Bridge,
	isNative bool,
	nativeTokenChainID ids.ID,
	teleporterMessenger *teleportermessenger.TeleporterMessenger,
) (*types.Receipt, ids.ID) {
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := transactor.BridgeTokens(
		opts,
		destinationBlockchainID,
		destinationBridgeAddress,
		token,
		recipient,
		totalAmount,
		primaryFeeAmount,
		secondaryFeeAmount,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt := utils.WaitForTransactionSuccess(ctx, source, tx.Hash())

	event, err := utils.GetEventFromLogs(receipt.Logs, teleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	if isNative {
		Expect(event.DestinationBlockchainID[:]).Should(Equal(destinationBlockchainID[:]))
	} else {
		Expect(event.DestinationBlockchainID[:]).Should(Equal(nativeTokenChainID[:]))
	}

	return receipt, event.MessageID
}

func approveBridgeToken(
	ctx context.Context,
	source interfaces.SubnetTestInfo,
	bridgeTokenAddress common.Address,
	transactor *bridgetoken.BridgeToken,
	amount *big.Int,
	spender common.Address,
	fundedAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
) {
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := transactor.Approve(opts, spender, amount)
	Expect(err).Should(BeNil())

	utils.WaitForTransactionSuccess(ctx, source, tx.Hash())
}
