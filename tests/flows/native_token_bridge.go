package flows

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	nativetokendestination "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/examples/NativeTokenBridge/NativeTokenDestination"
	nativetokensource "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/examples/NativeTokenBridge/NativeTokenSource"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func NativeTokenBridge(network interfaces.LocalNetwork) {
	const (
		// This test needs a unique deployer key, whose nonce 0 is used to deploy the bridge contract
		// on each chain. The address of the resulting contract has been added to the genesis file as
		// an admin for the Native Minter precompile.
		deployerKeyStr                     = "aad7440febfc8f9d73a58c3cb1f1754779a566978f9ebffcd4f4698e9b043985"
		NativeTokenSourceByteCodeFile      = "./contracts/out/NativeTokenSource.sol/NativeTokenSource.json"
		NativeTokenDestinationByteCodeFile = "./contracts/out/NativeTokenDestination.sol/NativeTokenDestination.json"
	)
	var (
		initialReserveImbalance = big.NewInt(0).Mul(big.NewInt(1e15), big.NewInt(1e9))
		valueToSend             = big.NewInt(0).Div(initialReserveImbalance, big.NewInt(4))
		valueToReturn           = big.NewInt(0).Div(valueToSend, big.NewInt(4))
		ctx                     = context.Background()
		deployerAddress         = common.HexToAddress("0x1337cfd2dCff6270615B90938aCB1efE79801704")
		tokenReceiverAddress    = common.HexToAddress("0x0123456789012345678901234567890123456789")
		burnedTxFeeAddressDest  = common.HexToAddress("0x0100000000000000000000000000000000000000")
		burnAddressSource       = common.HexToAddress("0x0100000000000000000000000000000000010203")

		emptyDestFeeInfo = nativetokendestination.TeleporterFeeInfo{
			FeeTokenAddress: common.Address{},
			Amount:          common.Big0,
		}
		emptySourceFeeInfo = nativetokensource.TeleporterFeeInfo{
			FeeTokenAddress: common.Address{},
			Amount:          common.Big0,
		}
	)

	sourceSubnet := network.GetPrimaryNetworkInfo() // TODO: Integrate the C-Chain
	_, destSubnet := utils.GetTwoSubnets(network)
	_, fundedKey := network.GetFundedAccountInfo()

	// Info we need to calculate for the test
	deployerPK, err := crypto.HexToECDSA(deployerKeyStr)
	Expect(err).Should(BeNil())
	bridgeContractAddress, err := deploymentUtils.DeriveEVMContractAddress(deployerAddress, 0)
	Expect(err).Should(BeNil())
	log.Info("Native Token Bridge Contract Address: " + bridgeContractAddress.Hex())

	{
		// The deployer is used to also send native transfers so we don't have to create an additional address.
		// The deployer will send 1.25*initialReserveImbalance over the course of the test. Send 2*initialReserveImbalance
		// native tokens so that it also has enough to cover transaction fees, including deploying the contract.
		utils.SendNativeTransfer(
			ctx,
			sourceSubnet,
			fundedKey,
			deployerAddress,
			utils.BigIntMul(initialReserveImbalance, big.NewInt(2)),
		)

		utils.SendNativeTransfer(
			ctx,
			destSubnet,
			fundedKey,
			deployerAddress,
			utils.BigIntMul(initialReserveImbalance, big.NewInt(2)),
		)
	}

	{
		// Deploy the contracts
		// Both contracts in this test will be deployed to 0xAcB633F5B00099c7ec187eB00156c5cd9D854b5B,
		// though they do not necessarily have to be deployed at the same address, each contract needs
		// to know the address of the other.
		// The nativeTokenDestination contract must be added to "adminAddresses" of "contractNativeMinterConfig"
		// in the genesis file for the subnet. This will allow it to call the native minter precompile.
		erc20TokenSourceAbi, err := nativetokensource.NativeTokenSourceMetaData.GetAbi()
		Expect(err).Should(BeNil())
		utils.DeployContract(
			ctx,
			NativeTokenSourceByteCodeFile,
			deployerPK,
			sourceSubnet,
			erc20TokenSourceAbi,
			sourceSubnet.TeleporterRegistryAddress,
			deployerAddress,
			destSubnet.BlockchainID,
			bridgeContractAddress,
		)
		Expect(err).Should(BeNil())

		nativeTokenDestinationAbi, err := nativetokendestination.NativeTokenDestinationMetaData.GetAbi()
		Expect(err).Should(BeNil())
		utils.DeployContract(
			ctx,
			NativeTokenDestinationByteCodeFile,
			deployerPK,
			destSubnet,
			nativeTokenDestinationAbi,
			destSubnet.TeleporterRegistryAddress,
			deployerAddress,
			sourceSubnet.BlockchainID,
			bridgeContractAddress,
			initialReserveImbalance,
		)

		log.Info("Finished deploying Bridge contracts")
	}

	// Create abi objects to call the contract with
	nativeTokenDestination, err := nativetokendestination.NewNativeTokenDestination(
		bridgeContractAddress,
		destSubnet.RPCClient,
	)
	Expect(err).Should(BeNil())
	nativeTokenSource, err := nativetokensource.NewNativeTokenSource(
		bridgeContractAddress,
		sourceSubnet.RPCClient,
	)
	Expect(err).Should(BeNil())

	{
		// Transfer some tokens A -> B
		// Check starting balance is 0
		utils.CheckBalance(ctx, tokenReceiverAddress, common.Big0, destSubnet.RPCClient)

		checkReserveImbalance(initialReserveImbalance, nativeTokenDestination)

		destChainReceipt := sendNativeTokensToDestination(
			ctx,
			network,
			valueToSend,
			deployerPK,
			tokenReceiverAddress,
			sourceSubnet,
			destSubnet,
			nativeTokenSource,
			emptySourceFeeInfo,
		)

		checkCollateralEvent(
			destChainReceipt.Logs,
			nativeTokenDestination,
			valueToSend,
			big.NewInt(0).Sub(initialReserveImbalance, valueToSend),
		)
		checkReserveImbalance(
			big.NewInt(0).Sub(initialReserveImbalance, valueToSend),
			nativeTokenDestination,
		)

		_, err = utils.GetEventFromLogs(
			destChainReceipt.Logs,
			nativeTokenDestination.ParseNativeTokensMinted,
		)
		Expect(err).ShouldNot(BeNil())

		// Check intermediate balance, no tokens should be minted because we haven't collateralized
		utils.CheckBalance(ctx, tokenReceiverAddress, common.Big0, destSubnet.RPCClient)
	}

	{
		// Fail to Transfer tokens B -> A because bridge is not collateralized
		// Check starting balance is 0
		utils.CheckBalance(ctx, tokenReceiverAddress, common.Big0, sourceSubnet.RPCClient)

		transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, destSubnet.EVMChainID)
		Expect(err).Should(BeNil())
		transactor.Value = valueToSend

		// This transfer should revert because the bridge isn't collateralized
		_, err = nativeTokenDestination.TransferToSource(
			transactor,
			tokenReceiverAddress,
			emptyDestFeeInfo,
			[]common.Address{},
		)
		Expect(err).ShouldNot(BeNil())

		// Check we should fail to send because we're not collateralized
		utils.CheckBalance(ctx, tokenReceiverAddress, common.Big0, sourceSubnet.RPCClient)
	}

	{
		// Transfer more tokens A -> B to collateralize the bridge
		// Check starting balance is 0
		utils.CheckBalance(ctx, tokenReceiverAddress, common.Big0, destSubnet.RPCClient)
		checkReserveImbalance(
			big.NewInt(0).Sub(initialReserveImbalance, valueToSend),
			nativeTokenDestination,
		)

		destChainReceipt := sendNativeTokensToDestination(
			ctx,
			network,
			initialReserveImbalance,
			deployerPK,
			tokenReceiverAddress,
			sourceSubnet,
			destSubnet,
			nativeTokenSource,
			emptySourceFeeInfo,
		)

		checkCollateralEvent(
			destChainReceipt.Logs,
			nativeTokenDestination,
			big.NewInt(0).Sub(initialReserveImbalance, valueToSend),
			common.Big0,
		)
		checkMintEvent(
			destChainReceipt.Logs,
			nativeTokenDestination,
			tokenReceiverAddress,
			valueToSend,
		)
		checkReserveImbalance(common.Big0, nativeTokenDestination)

		// We should have minted the excess coins after checking the collateral
		utils.CheckBalance(ctx, tokenReceiverAddress, valueToSend, destSubnet.RPCClient)
	}

	{
		// Transfer tokens B -> A
		sourceChainReceipt := sendTokensToSource(
			ctx,
			network,
			valueToReturn,
			deployerPK,
			tokenReceiverAddress,
			sourceSubnet,
			destSubnet,
			nativeTokenDestination,
			emptyDestFeeInfo,
		)

		checkUnlockNativeEvent(
			sourceChainReceipt.Logs,
			nativeTokenSource,
			tokenReceiverAddress,
			valueToReturn,
		)

		utils.CheckBalance(ctx, tokenReceiverAddress, valueToReturn, sourceSubnet.RPCClient)
	}

	{
		// Check reporting of burned tx fees to Source Chain
		burnedTxFeesBalanceDest, err := destSubnet.RPCClient.BalanceAt(
			ctx,
			burnedTxFeeAddressDest,
			nil,
		)
		Expect(err).Should(BeNil())
		Expect(burnedTxFeesBalanceDest.Cmp(common.Big0) > 0).Should(BeTrue())

		transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, destSubnet.EVMChainID)
		Expect(err).Should(BeNil())
		tx, err := nativeTokenDestination.ReportTotalBurnedTxFees(
			transactor,
			emptyDestFeeInfo,
			[]common.Address{},
		)
		Expect(err).Should(BeNil())

		destChainReceipt := utils.WaitForTransactionSuccess(ctx, destSubnet, tx.Hash())

		reportEvent, err := utils.GetEventFromLogs(
			destChainReceipt.Logs,
			nativeTokenDestination.ParseReportTotalBurnedTxFees,
		)
		Expect(err).Should(BeNil())
		utils.ExpectBigEqual(reportEvent.BurnAddressBalance, burnedTxFeesBalanceDest)

		burnedTxFeesBalanceSource, err := sourceSubnet.RPCClient.BalanceAt(
			ctx,
			burnAddressSource,
			nil,
		)
		Expect(err).Should(BeNil())
		Expect(burnedTxFeesBalanceSource.Uint64()).Should(BeZero())

		sourceChainReceipt := network.RelayMessage(ctx, destChainReceipt, destSubnet, sourceSubnet, true)

		burnEvent := utils.GetEventFromLogsOrTrace(
			ctx,
			sourceChainReceipt,
			sourceSubnet,
			nativeTokenSource.ParseBurnTokens,
		)
		utils.ExpectBigEqual(burnedTxFeesBalanceDest, burnEvent.Amount)

		burnedTxFeesBalanceSource2, err := sourceSubnet.RPCClient.BalanceAt(
			ctx,
			burnAddressSource,
			nil,
		)
		Expect(err).Should(BeNil())
		Expect(
			burnedTxFeesBalanceSource2.Cmp(
				big.NewInt(0).Add(burnedTxFeesBalanceSource, burnEvent.Amount),
			) >= 0,
		).Should(BeTrue())
	}
}

func checkUnlockNativeEvent(
	logs []*types.Log,
	nativeTokenSource *nativetokensource.NativeTokenSource,
	recipient common.Address,
	value *big.Int,
) {
	unlockEvent, err := utils.GetEventFromLogs(logs, nativeTokenSource.ParseUnlockTokens)
	Expect(err).Should(BeNil())
	Expect(unlockEvent.Recipient).Should(Equal(recipient))
	Expect(unlockEvent.Amount.Cmp(value)).Should(BeZero())
}

func checkCollateralEvent(
	logs []*types.Log,
	nativeTokenDestination *nativetokendestination.NativeTokenDestination,
	collateralAdded *big.Int,
	collateralRemaining *big.Int,
) {
	collateralEvent, err := utils.GetEventFromLogs(
		logs,
		nativeTokenDestination.ParseCollateralAdded,
	)
	Expect(err).Should(BeNil())
	Expect(collateralEvent.Amount.Cmp(collateralAdded)).Should(BeZero())
	Expect(collateralEvent.Remaining.Cmp(collateralEvent.Remaining)).Should(BeZero())
}

func checkMintEvent(
	logs []*types.Log,
	nativeTokenDestination *nativetokendestination.NativeTokenDestination,
	recipient common.Address,
	value *big.Int,
) {
	mintEvent, err := utils.GetEventFromLogs(logs, nativeTokenDestination.ParseNativeTokensMinted)
	Expect(err).Should(BeNil())
	Expect(mintEvent.Recipient).Should(Equal(recipient))
	Expect(mintEvent.Amount.Cmp(value)).Should(BeZero())
}

func checkReserveImbalance(
	value *big.Int,
	nativeTokenDestination *nativetokendestination.NativeTokenDestination,
) {
	imbalance, err := nativeTokenDestination.CurrentReserveImbalance(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(imbalance, value)
}

func sendTokensToSource(
	ctx context.Context,
	network interfaces.LocalNetwork,
	valueToSend *big.Int,
	fromKey *ecdsa.PrivateKey,
	toAddress common.Address,
	sourceSubnet interfaces.SubnetTestInfo,
	destinationSubnet interfaces.SubnetTestInfo,
	nativeTokenDestination *nativetokendestination.NativeTokenDestination,
	feeInfo nativetokendestination.TeleporterFeeInfo,
) *types.Receipt {
	transactor, err := bind.NewKeyedTransactorWithChainID(fromKey, destinationSubnet.EVMChainID)
	Expect(err).Should(BeNil())
	transactor.Value = valueToSend

	tx, err := nativeTokenDestination.TransferToSource(
		transactor,
		toAddress,
		feeInfo,
		[]common.Address{},
	)
	Expect(err).Should(BeNil())

	destChainReceipt := utils.WaitForTransactionSuccess(ctx, destinationSubnet, tx.Hash())

	transferEvent, err := utils.GetEventFromLogs(
		destChainReceipt.Logs,
		nativeTokenDestination.ParseTransferToSource,
	)
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(transferEvent.Amount, valueToSend)

	receipt := network.RelayMessage(ctx, destChainReceipt, destinationSubnet, sourceSubnet, true)

	return receipt
}

func sendNativeTokensToDestination(
	ctx context.Context,
	network interfaces.LocalNetwork,
	valueToSend *big.Int,
	fromKey *ecdsa.PrivateKey,
	toAddress common.Address,
	sourceSubnet interfaces.SubnetTestInfo,
	destinationSubnet interfaces.SubnetTestInfo,
	nativeTokenSource *nativetokensource.NativeTokenSource,
	feeInfo nativetokensource.TeleporterFeeInfo,
) *types.Receipt {
	transactor, err := bind.NewKeyedTransactorWithChainID(fromKey, sourceSubnet.EVMChainID)
	Expect(err).Should(BeNil())
	transactor.Value = valueToSend

	tx, err := nativeTokenSource.TransferToDestination(
		transactor,
		toAddress,
		feeInfo,
		[]common.Address{},
	)
	Expect(err).Should(BeNil())

	sourceChainReceipt := utils.WaitForTransactionSuccess(ctx, sourceSubnet, tx.Hash())

	transferEvent, err := utils.GetEventFromLogs(
		sourceChainReceipt.Logs,
		nativeTokenSource.ParseTransferToDestination,
	)
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(transferEvent.Amount, valueToSend)

	receipt := network.RelayMessage(ctx, sourceChainReceipt, sourceSubnet, destinationSubnet, true)

	return receipt
}
