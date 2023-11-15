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
	"github.com/ava-labs/teleporter/tests/utils"
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func ERC20ToNativeTokenBridge() {
	const (
		initialReserveImbalance = uint64(1e15)
		valueToSend1            = initialReserveImbalance / 4
		valueToReturn           = valueToSend1 / 4

		// Each test file needs a unique deployer that must be funded with tokens to deploy
		deployerKeyStr                     = "ca7269c1fe2a5b86884a119aa516b8d5b641670b83aac0ebf9f2d71affcc12e4"
		ExampleERC20ByteCodeFile           = "./contracts/out/ExampleERC20.sol/ExampleERC20.json"
		ERC20TokenSourceByteCodeFile       = "./contracts/out/ERC20TokenSource.sol/ERC20TokenSource.json"
		NativeTokenDestinationByteCodeFile = "./contracts/out/NativeTokenDestination.sol/NativeTokenDestination.json"
	)
	var (
		ctx                  = context.Background()
		deployerAddress      = common.HexToAddress("0x539447ab8Be7e927bE8E005663C81ff2AE951337")
		tokenReceiverAddress = common.HexToAddress("0x4444444444444444444444444444444444444444")

		emptyDestFeeInfo = nativetokendestination.TeleporterFeeInfo{
			ContractAddress: common.Address{},
			Amount:          common.Big0,
		}
	)

	subnetA := utils.GetSubnetATestInfo()
	subnetB := utils.GetSubnetBTestInfo()
	teleporterContractAddress := utils.GetTeleporterContractAddress()

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
	utils.DeployContract(ctx, ERC20TokenSourceByteCodeFile, deployerPK, subnetA, erc20TokenSourceAbi, teleporterContractAddress, subnetB.BlockchainID, bridgeContractAddress, exampleERC20ContractAddress)

	nativeTokenDestinationAbi, err := nativetokendestination.NativeTokenDestinationMetaData.GetAbi()
	Expect(err).Should(BeNil())
	utils.DeployContract(ctx, NativeTokenDestinationByteCodeFile, deployerPK, subnetB, nativeTokenDestinationAbi, teleporterContractAddress, subnetA.BlockchainID, bridgeContractAddress, new(big.Int).SetUint64(initialReserveImbalance))

	exampleERC20Abi, err := exampleerc20.ExampleERC20MetaData.GetAbi()
	Expect(err).Should(BeNil())
	utils.DeployContract(ctx, ExampleERC20ByteCodeFile, deployerPK, subnetA, exampleERC20Abi)

	log.Info("Finished deploying contracts")

	// Create abi objects to call the contract with
	nativeTokenDestination, err := nativetokendestination.NewNativeTokenDestination(bridgeContractAddress, subnetB.WSClient)
	Expect(err).Should(BeNil())
	erc20TokenSource, err := erc20tokensource.NewERC20TokenSource(bridgeContractAddress, subnetA.WSClient)
	Expect(err).Should(BeNil())
	exampleERC20, err := exampleerc20.NewExampleERC20(exampleERC20ContractAddress, subnetA.WSClient)
	Expect(err).Should(BeNil())

	// Helper function
	sendTokensToSource := func(valueToSend uint64, fromKey *ecdsa.PrivateKey, toAddress common.Address) *types.Receipt {
		transactor, err := bind.NewKeyedTransactorWithChainID(fromKey, subnetB.ChainID)
		Expect(err).Should(BeNil())
		transactor.Value = new(big.Int).SetUint64(valueToSend)

		tx, err := nativeTokenDestination.TransferToSource(transactor, toAddress, emptyDestFeeInfo, []common.Address{})
		Expect(err).Should(BeNil())
		log.Info("Sent TransferToSource transaction on destination chain", "sourceChainID", subnetA.BlockchainID, "txHash", tx.Hash().Hex())

		sourceChainReceipt := utils.WaitForTransaction(ctx, tx.Hash(), subnetB.WSClient)
		Expect(err).Should(BeNil())
		Expect(sourceChainReceipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

		transferEvent, err := utils.GetEventFromLogs(sourceChainReceipt.Logs, nativeTokenDestination.ParseTransferToSource)
		Expect(err).Should(BeNil())
		Expect(transferEvent.Amount.Uint64()).Should(Equal(valueToSend))

		receipt, err := utils.RelayMessage(ctx, sourceChainReceipt.BlockHash, sourceChainReceipt.BlockNumber, subnetB, subnetA)
		Expect(err).Should(BeNil())

		return receipt
	}

	// Helper function
	sendTokensToDestination := func(valueToSend uint64, fromKey *ecdsa.PrivateKey, toAddress common.Address) *types.Receipt {
		transactor, err := bind.NewKeyedTransactorWithChainID(fromKey, subnetA.ChainID)
		Expect(err).Should(BeNil())

		tx, err := erc20TokenSource.TransferToDestination(transactor, toAddress, new(big.Int).SetUint64(valueToSend), big.NewInt(0), []common.Address{})
		Expect(err).Should(BeNil())
		log.Info("Sent TransferToDestination transaction on source chain", "destinationChainID", subnetB.BlockchainID, "txHash", tx.Hash().Hex())

		sourceChainReceipt := utils.WaitForTransaction(ctx, tx.Hash(), subnetA.WSClient)
		Expect(err).Should(BeNil())
		Expect(sourceChainReceipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

		transferEvent, err := utils.GetEventFromLogs(sourceChainReceipt.Logs, erc20TokenSource.ParseTransferToDestination)
		Expect(err).Should(BeNil())
		Expect(transferEvent.TransferAmount.Uint64()).Should(Equal(valueToSend))

		receipt, err := utils.RelayMessage(ctx, sourceChainReceipt.BlockHash, sourceChainReceipt.BlockNumber, subnetA, subnetB)
		Expect(err).Should(BeNil())

		return receipt
	}

	{ // Give erc20TokenSource allowance to spend all of the deployer's ERC20 Tokens
		bal, err := exampleERC20.BalanceOf(nil, deployerAddress)
		Expect(err).Should(BeNil())
		Expect(bal.Div(bal, big.NewInt(10)).Uint64() > initialReserveImbalance).Should(BeTrue())

		transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetA.ChainID)
		Expect(err).Should(BeNil())
		tx, err := exampleERC20.Approve(transactor, bridgeContractAddress, bal)
		Expect(err).Should(BeNil())

		receipt := utils.WaitForTransaction(ctx, tx.Hash(), subnetA.WSClient)
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	}

	{ // Transfer some tokens A -> B
		// Check starting balance is 0

		utils.CheckBalance(ctx, tokenReceiverAddress, 0, subnetB.WSClient)

		CheckReserveImbalance(initialReserveImbalance, nativeTokenDestination)

		destChainReceipt := sendTokensToDestination(valueToSend1, deployerPK, tokenReceiverAddress)

		collateralEvent, err := utils.GetEventFromLogs(destChainReceipt.Logs, nativeTokenDestination.ParseCollateralAdded)
		Expect(err).Should(BeNil())
		Expect(collateralEvent.Amount.Uint64()).Should(Equal(valueToSend1))

		_, err = utils.GetEventFromLogs(destChainReceipt.Logs, nativeTokenDestination.ParseNativeTokensMinted)
		Expect(err).ShouldNot(BeNil())

		CheckReserveImbalance(initialReserveImbalance-valueToSend1, nativeTokenDestination)

		// Check intermediate balance, no tokens should be minted because we haven't collateralized
		utils.CheckBalance(ctx, tokenReceiverAddress, 0, subnetB.WSClient)
	}

	{ // Fail to Transfer tokens B -> A because bridge is not collateralized
		// Check starting balance is 0
		utils.CheckBalance(ctx, tokenReceiverAddress, 0, subnetB.WSClient)

		transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetB.ChainID)
		Expect(err).Should(BeNil())
		transactor.Value = new(big.Int).SetUint64(valueToSend1)

		_, err = nativeTokenDestination.TransferToSource(transactor, tokenReceiverAddress, emptyDestFeeInfo, []common.Address{})
		Expect(err).ShouldNot(BeNil())

		CheckReserveImbalance(initialReserveImbalance-valueToSend1, nativeTokenDestination)

		// Check we failed to send because we're not collateralized
		utils.CheckBalance(ctx, tokenReceiverAddress, 0, subnetB.WSClient)
	}

	{ // Transfer more tokens A -> B to collateralize the bridge
		// Check starting balance is 0
		utils.CheckBalance(ctx, tokenReceiverAddress, 0, subnetB.WSClient)

		CheckReserveImbalance(initialReserveImbalance-valueToSend1, nativeTokenDestination)

		destChainReceipt := sendTokensToDestination(initialReserveImbalance, deployerPK, tokenReceiverAddress)

		collateralEvent, err := utils.GetEventFromLogs(destChainReceipt.Logs, nativeTokenDestination.ParseCollateralAdded)
		Expect(err).Should(BeNil())
		Expect(collateralEvent.Amount.Uint64()).Should(Equal(initialReserveImbalance - valueToSend1))

		mintEvent, err := utils.GetEventFromLogs(destChainReceipt.Logs, nativeTokenDestination.ParseNativeTokensMinted)
		Expect(err).Should(BeNil())
		Expect(mintEvent.Amount.Uint64()).Should(Equal(valueToSend1))

		CheckReserveImbalance(0, nativeTokenDestination)

		// We should have minted the excess coins after checking the collateral
		utils.CheckBalance(ctx, tokenReceiverAddress, valueToSend1, subnetB.WSClient)
	}

	{ // Transfer tokens B -> A
		sourceChainReceipt := sendTokensToSource(valueToReturn, deployerPK, tokenReceiverAddress)

		unlockEvent, err := utils.GetEventFromLogs(sourceChainReceipt.Logs, erc20TokenSource.ParseUnlockTokens)
		Expect(err).Should(BeNil())
		Expect(unlockEvent.Amount.Uint64()).Should(Equal(valueToReturn))

		bal, err := exampleERC20.BalanceOf(nil, tokenReceiverAddress)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(Equal(valueToReturn))
	}
}

func CheckReserveImbalance(value uint64, nativeTokenDestination *nativetokendestination.NativeTokenDestination) {
	imbalance, err := nativeTokenDestination.CurrentReserveImbalance(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(imbalance.Uint64()).Should(Equal(value))
}
