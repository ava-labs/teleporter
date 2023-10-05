package tests

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	nativetokendestination "github.com/ava-labs/teleporter/abis/NativeTokenDestination"
	nativetokensource "github.com/ava-labs/teleporter/abis/NativeTokenSource"
	deploymentUtils "github.com/ava-labs/teleporter/contract-deployment/utils"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func NativeTokenBridge() {
	const (
		valueToSend   = int64(1000_000_000_000_000_000)
		valueToReturn = valueToSend / 4

		bridgeDeployerKeyStr               = "aad7440febfc8f9d73a58c3cb1f1754779a566978f9ebffcd4f4698e9b043985"
		NativeTokenSourceByteCodeFile      = "./contracts/out/NativeTokenSource.sol/NativeTokenSource.json"
		NativeTokenDestinationByteCodeFile = "./contracts/out/NativeTokenDestination.sol/NativeTokenDestination.json"
	)
	var (
		ctx                       = context.Background()
		nativeTokenBridgeDeployer = common.HexToAddress("0x1337cfd2dCff6270615B90938aCB1efE79801704")
		tokenReceiverAddress      = common.HexToAddress("0x0123456789012345678901234567890123456789")
	)

	subnetA := utils.GetSubnetATestInfo()
	subnetB := utils.GetSubnetBTestInfo()
	teleporterContractAddress := utils.GetTeleporterContractAddress()

	nativeTokenBridgeDeployerPK, err := crypto.HexToECDSA(bridgeDeployerKeyStr)
	Expect(err).Should(BeNil())
	nativeTokenBridgeContractAddress, err := deploymentUtils.DeriveEVMContractAddress(nativeTokenBridgeDeployer, 0)
	Expect(err).Should(BeNil())
	log.Info("Native Token Bridge Contract Address: " + nativeTokenBridgeContractAddress.Hex())

	{ // Deploy the contracts
		nativeTokenSourceBytecode, err := deploymentUtils.ExtractByteCode(NativeTokenSourceByteCodeFile)
		Expect(err).Should(BeNil())
		chainATransactor, err := bind.NewKeyedTransactorWithChainID(nativeTokenBridgeDeployerPK, subnetA.ChainID)
		Expect(err).Should(BeNil())
		nativeTokenSourceAbi, err := nativetokensource.NativetokensourceMetaData.GetAbi()
		Expect(err).Should(BeNil())
		_, txA, _, err := bind.DeployContract(chainATransactor, *nativeTokenSourceAbi, nativeTokenSourceBytecode, subnetA.WSClient, teleporterContractAddress, subnetB.BlockchainID, nativeTokenBridgeContractAddress)
		Expect(err).Should(BeNil())

		nativeTokenDestinationBytecode, err := deploymentUtils.ExtractByteCode(NativeTokenDestinationByteCodeFile)
		Expect(err).Should(BeNil())
		chainBTransactor, err := bind.NewKeyedTransactorWithChainID(nativeTokenBridgeDeployerPK, subnetB.ChainID)
		Expect(err).Should(BeNil())
		nativeTokenDestinationAbi, err := nativetokendestination.NativetokendestinationMetaData.GetAbi()
		Expect(err).Should(BeNil())
		_, txB, _, err := bind.DeployContract(chainBTransactor, *nativeTokenDestinationAbi, nativeTokenDestinationBytecode, subnetB.WSClient, teleporterContractAddress, subnetA.BlockchainID, nativeTokenBridgeContractAddress, common.Big0)
		Expect(err).Should(BeNil())

		// Wait for transaction, then check code was deployed
		utils.WaitForTransaction(ctx, txA.Hash(), subnetA.WSClient)
		bridgeCodeA, err := subnetA.WSClient.CodeAt(ctx, nativeTokenBridgeContractAddress, nil)
		Expect(err).Should(BeNil())
		Expect(len(bridgeCodeA)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode

		// Wait for transaction, then check code was deployed
		utils.WaitForTransaction(ctx, txB.Hash(), subnetB.WSClient)
		bridgeCodeB, err := subnetB.WSClient.CodeAt(ctx, nativeTokenBridgeContractAddress, nil)
		Expect(err).Should(BeNil())
		Expect(len(bridgeCodeB)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode

		log.Info("Finished deploying Bridge contracts")
	}

	// Create abi objects to call the contract with
	nativeTokenDestination, err := nativetokendestination.NewNativetokendestination(nativeTokenBridgeContractAddress, subnetB.WSClient)
	Expect(err).Should(BeNil())
	nativeTokenSource, err := nativetokensource.NewNativetokensource(nativeTokenBridgeContractAddress, subnetA.WSClient)
	Expect(err).Should(BeNil())

	{ // Transfer tokens A -> B
		transactor, err := bind.NewKeyedTransactorWithChainID(nativeTokenBridgeDeployerPK, subnetA.ChainID)
		Expect(err).Should(BeNil())
		transactor.Value = big.NewInt(valueToSend)

		tx, err := nativeTokenSource.TransferToDestination(transactor, tokenReceiverAddress, tokenReceiverAddress, big.NewInt(0))
		Expect(err).Should(BeNil())
		log.Info("Sent TransferToDestination transaction on source chain", "destinationChainID", subnetB.BlockchainID, "txHash", tx.Hash().Hex())

		receipt := utils.WaitForTransaction(ctx, tx.Hash(), subnetA.WSClient)
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

		// Check starting balance is 0
		bal, err := subnetB.WSClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Int64()).Should(Equal(common.Big0.Int64()))

		utils.RelayMessage(ctx, receipt.BlockHash, receipt.BlockNumber, subnetA, subnetB)

		// Check final balance
		bal, err = subnetB.WSClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Int64()).Should(Equal(valueToSend))
	}

	// Transfer tokens B -> A
	{
		transactor, err := bind.NewKeyedTransactorWithChainID(nativeTokenBridgeDeployerPK, subnetB.ChainID)
		Expect(err).Should(BeNil())
		transactor.Value = big.NewInt(valueToReturn)

		tx, err := nativeTokenDestination.TransferToSource(transactor, tokenReceiverAddress, tokenReceiverAddress, big.NewInt(0))
		Expect(err).Should(BeNil())
		log.Info("Sent TransferToSource transaction on destination chain", "sourceChainID", subnetA.BlockchainID, "txHash", tx.Hash().Hex())

		receipt := utils.WaitForTransaction(ctx, tx.Hash(), subnetB.WSClient)
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

		// Check starting balance is 0
		bal, err := subnetA.WSClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Int64()).Should(Equal(common.Big0.Int64()))

		utils.RelayMessage(ctx, receipt.BlockHash, receipt.BlockNumber, subnetB, subnetA)

		// Check final balance
		bal, err = subnetA.WSClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Int64()).Should(Equal(valueToReturn))
	}
}
