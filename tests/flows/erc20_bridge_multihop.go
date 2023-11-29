package flows

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	bridgetoken "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/ERC20Bridge/BridgeToken"
	erc20bridge "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/ERC20Bridge/ERC20Bridge"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func ERC20BridgeMultihop(network interfaces.Network) {
	subnetAInfo, subnetBInfo, subnetCInfo := utils.GetThreeSubnets(network)
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	subnetATeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress,
		subnetAInfo.RPCClient,
	)
	Expect(err).Should(BeNil())
	subnetBTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress,
		subnetBInfo.RPCClient,
	)
	Expect(err).Should(BeNil())
	subnetCTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress,
		subnetCInfo.RPCClient,
	)
	Expect(err).Should(BeNil())

	// Deploy an ERC20 to subnet A
	nativeERC20Address, nativeERC20 := utils.DeployExampleERC20(
		context.Background(),
		fundedKey,
		subnetAInfo,
	)

	// Deploy the ERC20 bridge to subnet A
	erc20BridgeAddressA, erc20BridgeA := utils.DeployERC20Bridge(ctx, fundedKey, subnetAInfo)
	// Deploy the ERC20 bridge to subnet B
	erc20BridgeAddressB, erc20BridgeB := utils.DeployERC20Bridge(ctx, fundedKey, subnetBInfo)
	// Deploy the ERC20 bridge to subnet C
	erc20BridgeAddressC, erc20BridgeC := utils.DeployERC20Bridge(ctx, fundedKey, subnetCInfo)

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
	createBridgeTokenMessageFeeAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(1))
	receipt, messageID := submitCreateBridgeToken(
		ctx,
		subnetAInfo,
		subnetBInfo.BlockchainID,
		erc20BridgeAddressB,
		nativeERC20Address,
		nativeERC20Address,
		createBridgeTokenMessageFeeAmount,
		fundedAddress,
		fundedKey,
		erc20BridgeA,
		subnetATeleporterMessenger,
	)

	// Relay message
	network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)
	// Check Teleporter message received on the destination
	delivered, err := subnetBTeleporterMessenger.MessageReceived(
		&bind.CallOpts{},
		subnetAInfo.BlockchainID,
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

	// Send a transaction on Subnet B to add support for the the ERC20 token to the bridge on Subnet C
	receipt, messageID = submitCreateBridgeToken(
		ctx,
		subnetAInfo,
		subnetCInfo.BlockchainID,
		erc20BridgeAddressC,
		nativeERC20Address,
		nativeERC20Address,
		createBridgeTokenMessageFeeAmount,
		fundedAddress,
		fundedKey,
		erc20BridgeA,
		subnetATeleporterMessenger,
	)
	// Relay message
	network.RelayMessage(ctx, receipt, subnetAInfo, subnetCInfo, true)
	// Check Teleporter message received on the destination
	delivered, err = subnetCTeleporterMessenger.MessageReceived(
		&bind.CallOpts{},
		subnetAInfo.BlockchainID,
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
	secondaryFeeAmount := big.NewInt(0)
	receipt, messageID = bridgeToken(
		ctx,
		subnetAInfo,
		subnetBInfo.BlockchainID,
		erc20BridgeAddressB,
		nativeERC20Address,
		fundedAddress,
		totalAmount,
		primaryFeeAmount,
		secondaryFeeAmount,
		fundedAddress,
		fundedKey,
		erc20BridgeA,
		true,
		subnetAInfo.BlockchainID,
		subnetATeleporterMessenger,
	)

	// Relay message
	network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)
	// Check Teleporter message received on the destination
	delivered, err = subnetBTeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

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

	// Check the recipient balance of the new bridge token.
	actualRecipientBalance, err := bridgeTokenB.BalanceOf(&bind.CallOpts{}, fundedAddress)
	Expect(err).Should(BeNil())
	Expect(actualRecipientBalance).Should(Equal(totalAmount.Sub(totalAmount, primaryFeeAmount)))

	// Approve the bridge contract on subnet B to spent the wrapped tokens in the user account.
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

	// Unwrap bridged tokens back to subnet A, then wrap tokens to final destination on subnet C
	totalAmount = big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(11))
	primaryFeeAmount = big.NewInt(1e18)
	secondaryFeeAmount = big.NewInt(1e18)
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
		subnetBTeleporterMessenger,
	)

	// Relay message from SubnetB to SubnetA
	receipt = network.RelayMessage(ctx, receipt, subnetBInfo, subnetAInfo, true)
	// Check Teleporter message received on the destination
	delivered, err = subnetATeleporterMessenger.MessageReceived(
		&bind.CallOpts{},
		subnetBInfo.BlockchainID,
		messageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Get the sendCrossChainMessage event from SubnetA to SubnetC
	event, err := utils.GetEventFromLogs(receipt.Logs, subnetATeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(subnetCInfo.BlockchainID[:]))
	messageID = event.Message.MessageID

	// Relay message from SubnetA to SubnetC
	network.RelayMessage(ctx, receipt, subnetAInfo, subnetCInfo, true)

	// Check Teleporter message received on the destination
	delivered, err = subnetCTeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	actualRecipientBalance, err = bridgeTokenC.BalanceOf(&bind.CallOpts{}, fundedAddress)
	Expect(err).Should(BeNil())
	expectedAmount := totalAmount.Sub(totalAmount, primaryFeeAmount).Sub(totalAmount, secondaryFeeAmount)
	Expect(actualRecipientBalance).Should(Equal(expectedAmount))

	// Check the redeemable reward balance of the relayer if the relayer address was set
	actualRelayerRedeemableBalance, err := subnetATeleporterMessenger.CheckRelayerRewardAmount(
		&bind.CallOpts{},
		fundedAddress,
		nativeERC20Address,
	)
	Expect(err).Should(BeNil())
	Expect(actualRelayerRedeemableBalance).Should(Equal(primaryFeeAmount.Add(primaryFeeAmount, secondaryFeeAmount)))

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

	// Send a transaction to unwrap tokens from Subnet C back to Subnet A
	totalAmount = big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(8))
	primaryFeeAmount = big.NewInt(1e18)
	secondaryFeeAmount = big.NewInt(0)
	receipt, messageID = bridgeToken(
		ctx,
		subnetCInfo,
		subnetAInfo.BlockchainID,
		erc20BridgeAddressA,
		bridgeTokenSubnetCAddress,
		fundedAddress,
		totalAmount,
		primaryFeeAmount,
		secondaryFeeAmount,
		fundedAddress,
		fundedKey,
		erc20BridgeC,
		false,
		subnetAInfo.BlockchainID,
		subnetCTeleporterMessenger,
	)

	// Relay message from SubnetC to SubnetA
	network.RelayMessage(ctx, receipt, subnetCInfo, subnetAInfo, true)
	// Check Teleporter message received on the destination
	delivered, err = subnetATeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetCInfo.BlockchainID, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check the balance of the native token after the unwrap
	actualNativeTokenDefaultAccountBalance, err := nativeERC20.BalanceOf(&bind.CallOpts{}, fundedAddress)
	Expect(err).Should(BeNil())
	expectedAmount = big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(9999999992))
	Expect(actualNativeTokenDefaultAccountBalance).Should(Equal(expectedAmount))

	// Check the balance of the native token for the relayer, which should have received the fee rewards
	actualRelayerRedeemableBalance, err = subnetATeleporterMessenger.CheckRelayerRewardAmount(
		&bind.CallOpts{},
		fundedAddress,
		nativeERC20Address,
	)
	Expect(err).Should(BeNil())
	expectedAmount = big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(4))
	Expect(actualRelayerRedeemableBalance).Should(Equal(expectedAmount))
}

func submitCreateBridgeToken(
	ctx context.Context,
	source interfaces.SubnetTestInfo,
	destinationChainID ids.ID,
	destinationBridgeAddress common.Address,
	nativeToken common.Address,
	messageFeeAsset common.Address,
	messageFeeAmount *big.Int,
	fundedAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	transactor *erc20bridge.ERC20Bridge,
	teleporterMessenger *teleportermessenger.TeleporterMessenger,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := transactor.SubmitCreateBridgeToken(
		opts,
		destinationChainID,
		destinationBridgeAddress,
		nativeToken,
		messageFeeAsset,
		messageFeeAmount,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, source.RPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	event, err := utils.GetEventFromLogs(receipt.Logs, teleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(destinationChainID[:]))

	log.Info("Successfully SubmitCreateBridgeToken",
		"txHash", tx.Hash().Hex(),
		"messageID", event.Message.MessageID)

	return receipt, event.Message.MessageID
}

func bridgeToken(
	ctx context.Context,
	source interfaces.SubnetTestInfo,
	destinationChainID ids.ID,
	destinationBridgeAddress common.Address,
	nativeToken common.Address,
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
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := transactor.BridgeTokens(
		opts,
		destinationChainID,
		destinationBridgeAddress,
		nativeToken,
		recipient,
		totalAmount,
		primaryFeeAmount,
		secondaryFeeAmount,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, source.RPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	event, err := utils.GetEventFromLogs(receipt.Logs, teleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	if isNative {
		Expect(event.DestinationBlockchainID[:]).Should(Equal(destinationChainID[:]))
	} else {
		Expect(event.DestinationBlockchainID[:]).Should(Equal(nativeTokenChainID[:]))
	}

	return receipt, event.Message.MessageID
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

	txn, err := transactor.Approve(opts, spender, amount)
	Expect(err).Should(BeNil())

	receipt, err := bind.WaitMined(ctx, source.RPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
}
