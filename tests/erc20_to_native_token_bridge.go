package tests

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	erc20tokensource "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/NativeTokenBridge/ERC20TokenSource"
	nativetokendestination "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/NativeTokenBridge/NativeTokenDestination"
	exampleerc20 "github.com/ava-labs/teleporter/abi-bindings/go/Mocks/ExampleERC20"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func ERC20ToNativeTokenBridge(network network.Network) {
	const (
		// Each test file needs a unique deployer that must be funded with tokens to deploy
		deployerKeyStr                     = "ca7269c1fe2a5b86884a119aa516b8d5b641670b83aac0ebf9f2d71affcc12e4"
		ExampleERC20ByteCodeFile           = "./contracts/out/ExampleERC20.sol/ExampleERC20.json"
		ERC20TokenSourceByteCodeFile       = "./contracts/out/ERC20TokenSource.sol/ERC20TokenSource.json"
		NativeTokenDestinationByteCodeFile = "./contracts/out/NativeTokenDestination.sol/NativeTokenDestination.json"
	)
	var (
		initialReserveImbalance      = big.NewInt(0).Mul(big.NewInt(1e15), big.NewInt(1e9))
		valueToSend                  = big.NewInt(0).Div(initialReserveImbalance, big.NewInt(4))
		intermediateReserveImbalance = utils.BigIntSub(initialReserveImbalance, valueToSend)
		valueToReturn                = big.NewInt(0).Div(valueToSend, big.NewInt(4))

		ctx                  = context.Background()
		deployerAddress      = common.HexToAddress("0x539447ab8Be7e927bE8E005663C81ff2AE951337")
		tokenReceiverAddress = common.HexToAddress("0x4444444444444444444444444444444444444444")
		burnedTxFeeAddress   = common.HexToAddress("0x0100000000000000000000000000000000000000")

		emptyDestFeeInfo = nativetokendestination.TeleporterFeeInfo{
			FeeTokenAddress: common.Address{},
			Amount:          common.Big0,
		}
	)

	subnets := network.GetSubnetsInfo()
	subnetA := subnets[0]
	subnetB := subnets[1]
	teleporterContractAddress := network.GetTeleporterContractAddress()

	// Info we need to calculate for the test
	deployerPK, err := crypto.HexToECDSA(deployerKeyStr)
	Expect(err).Should(BeNil())
	bridgeContractAddress, err := deploymentUtils.DeriveEVMContractAddress(deployerAddress, 0)
	Expect(err).Should(BeNil())
	log.Info("Native Token Bridge Contract Address: " + bridgeContractAddress.Hex())
	exampleERC20ContractAddress, err := deploymentUtils.DeriveEVMContractAddress(deployerAddress, 1)
	Expect(err).Should(BeNil())
	log.Info("Example ERC20 Contract Address: " + exampleERC20ContractAddress.Hex())

	// Deploy the contracts
	// Both contracts in this test will be deployed to 0x3405506b3711859c5070949ed9b700c7ba7bf750,
	// though they do not necessarily have to be deployed at the same address, each contract needs
	// to know the address of the other.
	// The nativeTokenDestination contract must be added to "adminAddresses" of "contractNativeMinterConfig"
	// in the genesis file for the subnet. This will allow it to call the native minter precompile.
	erc20TokenSourceAbi, err := erc20tokensource.ERC20TokenSourceMetaData.GetAbi()
	Expect(err).Should(BeNil())
	utils.DeployContract(
		ctx,
		ERC20TokenSourceByteCodeFile,
		deployerPK,
		subnetA,
		erc20TokenSourceAbi,
		teleporterContractAddress,
		subnetB.BlockchainID,
		bridgeContractAddress,
		exampleERC20ContractAddress,
	)

	nativeTokenDestinationAbi, err := nativetokendestination.NativeTokenDestinationMetaData.GetAbi()
	Expect(err).Should(BeNil())
	utils.DeployContract(
		ctx,
		NativeTokenDestinationByteCodeFile,
		deployerPK,
		subnetB,
		nativeTokenDestinationAbi,
		teleporterContractAddress,
		subnetA.BlockchainID,
		bridgeContractAddress,
		initialReserveImbalance,
	)

	exampleERC20Abi, err := exampleerc20.ExampleERC20MetaData.GetAbi()
	Expect(err).Should(BeNil())
	utils.DeployContract(ctx, ExampleERC20ByteCodeFile, deployerPK, subnetA, exampleERC20Abi)

	log.Info("Finished deploying contracts")

	// Create abi objects to call the contract with
	nativeTokenDestination, err := nativetokendestination.NewNativeTokenDestination(
		bridgeContractAddress,
		subnetB.ChainWSClient,
	)
	Expect(err).Should(BeNil())
	erc20TokenSource, err := erc20tokensource.NewERC20TokenSource(
		bridgeContractAddress,
		subnetA.ChainWSClient,
	)
	Expect(err).Should(BeNil())
	exampleERC20, err := exampleerc20.NewExampleERC20(
		exampleERC20ContractAddress,
		subnetA.ChainWSClient,
	)
	Expect(err).Should(BeNil())

	// Helper function
	sendTokensToSource := func(
		valueToSend *big.Int,
		fromKey *ecdsa.PrivateKey,
		toAddress common.Address,
	) *types.Receipt {
		transactor, err := bind.NewKeyedTransactorWithChainID(fromKey, subnetB.ChainIDInt)
		Expect(err).Should(BeNil())
		transactor.Value = valueToSend

		tx, err := nativeTokenDestination.TransferToSource(
			transactor,
			toAddress,
			emptyDestFeeInfo,
			[]common.Address{},
		)
		Expect(err).Should(BeNil())
		log.Info(
			"Sent TransferToSource transaction on destination chain",
			"sourceChainID",
			subnetA.BlockchainID,
			"txHash",
			tx.Hash().Hex(),
		)

		sourceChainReceipt := utils.WaitForTransactionSuccess(ctx, tx.Hash(), subnetB)

		transferEvent, err := utils.GetEventFromLogs(
			sourceChainReceipt.Logs,
			nativeTokenDestination.ParseTransferToSource,
		)
		Expect(err).Should(BeNil())
		utils.ExpectBigEqual(transferEvent.Amount, valueToSend)

		receipt := network.RelayMessage(ctx, sourceChainReceipt, subnetB, subnetA, true)

		return receipt
	}

	// Helper function
	sendTokensToDestination := func(
		valueToSend *big.Int,
		fromKey *ecdsa.PrivateKey,
		toAddress common.Address,
	) *types.Receipt {
		transactor, err := bind.NewKeyedTransactorWithChainID(fromKey, subnetA.ChainIDInt)
		Expect(err).Should(BeNil())

		tx, err := erc20TokenSource.TransferToDestination(
			transactor,
			toAddress,
			valueToSend,
			big.NewInt(0),
			[]common.Address{},
		)
		Expect(err).Should(BeNil())
		log.Info(
			"Sent TransferToDestination transaction on source chain",
			"destinationChainID",
			subnetB.BlockchainID,
			"txHash",
			tx.Hash().Hex(),
		)

		sourceChainReceipt := utils.WaitForTransactionSuccess(ctx, tx.Hash(), subnetA)

		transferEvent, err := utils.GetEventFromLogs(
			sourceChainReceipt.Logs,
			erc20TokenSource.ParseTransferToDestination,
		)
		Expect(err).Should(BeNil())
		utils.ExpectBigEqual(transferEvent.TransferAmount, valueToSend)

		receipt := network.RelayMessage(ctx, sourceChainReceipt, subnetA, subnetB, true)

		return receipt
	}

	{ // Give erc20TokenSource allowance to spend all of the deployer's ERC20 Tokens
		bal, err := exampleERC20.BalanceOf(nil, deployerAddress)
		Expect(err).Should(BeNil())

		transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetA.ChainIDInt)
		Expect(err).Should(BeNil())
		tx, err := exampleERC20.Approve(transactor, bridgeContractAddress, bal)
		Expect(err).Should(BeNil())

		utils.WaitForTransactionSuccess(ctx, tx.Hash(), subnetA)
	}

	{ // Transfer some tokens A -> B
		// Check starting balance is 0

		utils.CheckBalance(ctx, tokenReceiverAddress, common.Big0, subnetB.ChainWSClient)

		checkReserveImbalance(initialReserveImbalance, nativeTokenDestination)

		destChainReceipt := sendTokensToDestination(valueToSend, deployerPK, tokenReceiverAddress)

		collateralEvent, err := utils.GetEventFromLogs(
			destChainReceipt.Logs,
			nativeTokenDestination.ParseCollateralAdded,
		)
		Expect(err).Should(BeNil())
		utils.ExpectBigEqual(collateralEvent.Amount, valueToSend)

		_, err = utils.GetEventFromLogs(
			destChainReceipt.Logs,
			nativeTokenDestination.ParseNativeTokensMinted,
		)
		Expect(err).ShouldNot(BeNil())

		checkReserveImbalance(intermediateReserveImbalance, nativeTokenDestination)

		// Check intermediate balance, no tokens should be minted because we haven't collateralized
		utils.CheckBalance(ctx, tokenReceiverAddress, common.Big0, subnetB.ChainWSClient)
	}

	{ // Fail to Transfer tokens B -> A because bridge is not collateralized
		// Check starting balance is 0
		utils.CheckBalance(ctx, tokenReceiverAddress, common.Big0, subnetB.ChainWSClient)

		transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetB.ChainIDInt)
		Expect(err).Should(BeNil())
		transactor.Value = valueToSend

		_, err = nativeTokenDestination.TransferToSource(
			transactor,
			tokenReceiverAddress,
			emptyDestFeeInfo,
			[]common.Address{},
		)
		Expect(err).ShouldNot(BeNil())

		checkReserveImbalance(intermediateReserveImbalance, nativeTokenDestination)

		// Check we failed to send because we're not collateralized
		utils.CheckBalance(ctx, tokenReceiverAddress, common.Big0, subnetB.ChainWSClient)
	}

	{ // Transfer more tokens A -> B to collateralize the bridge
		// Check starting balance is 0
		utils.CheckBalance(ctx, tokenReceiverAddress, common.Big0, subnetB.ChainWSClient)

		checkReserveImbalance(intermediateReserveImbalance, nativeTokenDestination)

		destChainReceipt := sendTokensToDestination(
			initialReserveImbalance,
			deployerPK,
			tokenReceiverAddress,
		)

		collateralEvent, err := utils.GetEventFromLogs(
			destChainReceipt.Logs,
			nativeTokenDestination.ParseCollateralAdded,
		)
		Expect(err).Should(BeNil())
		Expect(collateralEvent.Amount).Should(Equal(intermediateReserveImbalance))

		mintEvent, err := utils.GetEventFromLogs(
			destChainReceipt.Logs,
			nativeTokenDestination.ParseNativeTokensMinted,
		)
		Expect(err).Should(BeNil())
		utils.ExpectBigEqual(mintEvent.Amount, valueToSend)

		checkReserveImbalance(common.Big0, nativeTokenDestination)

		// We should have minted the excess coins after checking the collateral
		utils.CheckBalance(ctx, tokenReceiverAddress, valueToSend, subnetB.ChainWSClient)
	}

	{ // Transfer tokens B -> A
		sourceChainReceipt := sendTokensToSource(valueToReturn, deployerPK, tokenReceiverAddress)

		checkUnlockERC20Event(
			sourceChainReceipt.Logs,
			erc20TokenSource,
			tokenReceiverAddress,
			valueToReturn,
		)

		bal, err := exampleERC20.BalanceOf(nil, tokenReceiverAddress)
		Expect(err).Should(BeNil())
		utils.ExpectBigEqual(bal, valueToReturn)
	}

	{ // Check reporting of burned tx fees to Source Chain
		burnedTxFeesBalanceDest, err := subnetB.ChainWSClient.BalanceAt(
			ctx,
			burnedTxFeeAddress,
			nil,
		)
		Expect(err).Should(BeNil())
		Expect(burnedTxFeesBalanceDest.Cmp(common.Big0) > 0).Should(BeTrue())

		transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetB.ChainIDInt)
		Expect(err).Should(BeNil())
		tx, err := nativeTokenDestination.ReportTotalBurnedTxFees(
			transactor,
			emptyDestFeeInfo,
			[]common.Address{},
		)
		Expect(err).Should(BeNil())

		destChainReceipt := utils.WaitForTransactionSuccess(ctx, tx.Hash(), subnetB)

		reportEvent, err := utils.GetEventFromLogs(
			destChainReceipt.Logs,
			nativeTokenDestination.ParseReportTotalBurnedTxFees,
		)
		Expect(err).Should(BeNil())
		utils.ExpectBigEqual(reportEvent.BurnAddressBalance, burnedTxFeesBalanceDest)

		burnedTxFeesBalanceSource, err := exampleERC20.BalanceOf(nil, burnedTxFeeAddress)
		Expect(err).Should(BeNil())
		utils.ExpectBigEqual(burnedTxFeesBalanceSource, common.Big0)

		sourceChainReceipt := network.RelayMessage(ctx, destChainReceipt, subnetB, subnetA, true)

		burnEvent, err := utils.GetEventFromLogs(
			sourceChainReceipt.Logs,
			erc20TokenSource.ParseBurnTokens,
		)
		Expect(err).Should(BeNil())
		utils.ExpectBigEqual(burnedTxFeesBalanceDest, burnEvent.Amount)

		burnedTxFeesBalanceSource2, err := exampleERC20.BalanceOf(nil, burnedTxFeeAddress)
		Expect(err).Should(BeNil())
		utils.ExpectBigEqual(burnedTxFeesBalanceSource2, burnEvent.Amount)
	}
}

func checkUnlockERC20Event(
	logs []*types.Log,
	erc20TokenSource *erc20tokensource.ERC20TokenSource,
	recipient common.Address,
	value *big.Int,
) {
	unlockEvent, err := utils.GetEventFromLogs(logs, erc20TokenSource.ParseUnlockTokens)
	Expect(err).Should(BeNil())
	Expect(unlockEvent.Recipient).Should(Equal(recipient))
	utils.ExpectBigEqual(unlockEvent.Amount, value)
}
