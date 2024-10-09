package flows

import (
	"context"
	"math/big"
	"time"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	validatorsetsig "github.com/ava-labs/teleporter/abi-bindings/go/governance/ValidatorSetSig"
	exampleerc20 "github.com/ava-labs/teleporter/abi-bindings/go/mocks/ExampleERC20"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

func ValidatorSetSig(network interfaces.LocalNetwork) {
	// ************************************************************************************************
	// Setup
	// ************************************************************************************************
	// Deploy ValidatorSetSig expecting signatures from subnetB instances to both subnets
	// Deploy exampleERC20 instance to both subnets
	// Construct ValidatorSetSig message with mock ERC20 as the target contract
	// Create off-chain Warp messages using the ValidatorSetSig message to be signed by the subnetB
	// ************************************************************************************************
	// Test Case 1: validatorChain (subnetB) != targetChain (subnetA)
	// ************************************************************************************************
	// Send the off-chain message to subnetA instance of ValidatorSetSig and confirm it is accepted.
	// Confirm the event is emitted
	// Retry the same message and confirm it fails
	// Send a new message with incremented nonce and confirm it is accepted.
	// ************************************************************************************************
	// Test Case 2: validatorChain (subnetB) == targetChain (subnetB)
	// ************************************************************************************************
	// Send a new message to subnetB instance of ValidatorSetSig and confirm it is accepted

	// ************************************************************************************************
	// Setup
	// ************************************************************************************************
	subnetA, subnetB := utils.GetTwoSubnets(network)
	_, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy a ValidatorSetSigContract to subnetA
	validatorSetSigContractAddress, validatorSetSig := utils.DeployValidatorSetSig(
		ctx,
		fundedKey,
		subnetA,
		subnetB,
	)
	// Deploy a ValidatorSetSigContract to subnetB
	validatorSetSigContractAddress2, validatorSetSig2 := utils.DeployValidatorSetSig(
		ctx,
		fundedKey,
		subnetB,
		subnetB,
	)

	// Deploy a mock ERC20 contract to subnetA
	exampleERC20ContractAddressA, exampleERC20A := utils.DeployExampleERC20(
		ctx,
		fundedKey,
		subnetA,
	)

	// Deploy a new example ERC20 contract this time to the same subnet as the validator.
	exampleERC20ContractAddressB, exampleERC20B := utils.DeployExampleERC20(
		ctx,
		fundedKey,
		subnetB,
	)

	erc20ABI, err := exampleerc20.ExampleERC20MetaData.GetAbi()
	Expect(err).Should(BeNil())

	// Confirm that the validatorContract has a balance of 0 on the example erc20 contracts on both subnets
	startingBalanceA, err := exampleERC20A.BalanceOf(
		&bind.CallOpts{}, validatorSetSigContractAddress)
	Expect(err).Should(BeNil())
	Expect(startingBalanceA.Cmp(big.NewInt(0))).Should(BeZero())

	startingBalanceB, err := exampleERC20B.BalanceOf(
		&bind.CallOpts{}, validatorSetSigContractAddress2)
	Expect(err).Should(BeNil())
	Expect(startingBalanceB.Cmp(big.NewInt(0))).Should(BeZero())

	// Construct a ValidatorSetSig message with mock ERC20 as the target contract
	// and mint 100 tokens as the TxPayload
	callData, err := erc20ABI.Pack("mint", big.NewInt(100))
	Expect(err).Should(BeNil())

	vssMessage1 := validatorsetsig.ValidatorSetSigMessage{
		ValidatorSetSigAddress: validatorSetSigContractAddress,
		TargetContractAddress:  exampleERC20ContractAddressA,
		TargetBlockchainID:     subnetA.BlockchainID,
		Nonce:                  big.NewInt(0),
		Value:                  big.NewInt(0),
		Payload:                callData,
	}

	// Construct a second ValidatorSetSig message with mock ERC20 as the target contract
	// and mint 50 tokens as the TxPayload
	callData2, err := erc20ABI.Pack("mint", big.NewInt(50))
	Expect(err).Should(BeNil())
	vssMessage2 := validatorsetsig.ValidatorSetSigMessage{
		ValidatorSetSigAddress: validatorSetSigContractAddress,
		TargetContractAddress:  exampleERC20ContractAddressA,
		TargetBlockchainID:     subnetA.BlockchainID,
		Nonce:                  big.NewInt(1),
		Value:                  big.NewInt(0),
		Payload:                callData2,
	}

	// Create a message for the case where validatorSetSig contract and targetContract are on the same chain.
	vssMessage3 := validatorsetsig.ValidatorSetSigMessage{
		ValidatorSetSigAddress: validatorSetSigContractAddress2,
		TargetContractAddress:  exampleERC20ContractAddressB,
		TargetBlockchainID:     subnetB.BlockchainID,
		Nonce:                  big.NewInt(0),
		Value:                  big.NewInt(0),
		Payload:                callData,
	}

	// Create chain config file with off-chain validatorsetsig message
	networkID := network.GetNetworkID()
	offchainMessages, warpEnabledChainConfigWithMsg := utils.InitOffChainMessageChainConfigValidatorSetSig(
		networkID,
		subnetB,
		validatorSetSigContractAddress,
		[]validatorsetsig.ValidatorSetSigMessage{vssMessage1, vssMessage2, vssMessage3},
	)

	// Create chain config with off-chain messages
	chainConfigs := make(utils.ChainConfigMap)
	chainConfigs.Add(subnetB, warpEnabledChainConfigWithMsg)

	// Restart nodes with new chain config
	network.SetChainConfigs(chainConfigs)
	restartCtx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	network.RestartNodes(restartCtx, network.GetAllNodeIDs())

	// ************************************************************************************************
	// Test Case 1: validatorChain (subnetB) != targetChain (subnetA)
	// ************************************************************************************************

	// Execute the ValidatorSetSig executeCall and wait for acceptance
	receipt := utils.ExecuteValidatorSetSigCallAndVerify(
		ctx,
		network,
		subnetB,
		subnetA,
		validatorSetSigContractAddress,
		fundedKey,
		&offchainMessages[0],
		true,
	)

	// Confirm that the Delivered event is emitted and that validatorSetSig contract has a balance of 100 of ERC20
	deliveredEvent, err := utils.GetEventFromLogs(receipt.Logs, validatorSetSig.ParseDelivered)
	Expect(err).Should(BeNil())
	Expect(deliveredEvent.TargetContractAddress).Should(Equal(exampleERC20ContractAddressA))
	Expect(deliveredEvent.Nonce.Cmp(big.NewInt(0))).Should(BeZero())

	endingBalance, err := exampleERC20A.BalanceOf(nil, validatorSetSigContractAddress)
	Expect(err).Should(BeNil())
	Expect(endingBalance).Should(Equal(big.NewInt(100)))

	// Resend the same message again and it should fail due to nonce being consumed

	_ = utils.ExecuteValidatorSetSigCallAndVerify(
		ctx,
		network,
		subnetB,
		subnetA,
		validatorSetSigContractAddress,
		fundedKey,
		&offchainMessages[0],
		false,
	)

	// Confirm that the validatorSetSig contract still has a balance of 100 of ERC20
	endingBalance, err = exampleERC20A.BalanceOf(nil, validatorSetSigContractAddress)
	Expect(err).Should(BeNil())
	Expect(endingBalance).Should(Equal(big.NewInt(100)))

	// Send another valid transaction with the incremented nonce
	receipt2 := utils.ExecuteValidatorSetSigCallAndVerify(
		ctx,
		network,
		subnetB,
		subnetA,
		validatorSetSigContractAddress,
		fundedKey,
		&offchainMessages[1],
		true,
	)

	// Confirm that the Delivered event is emitted and that validatorSetSig contract has a balance of 150 of ERC20
	deliveredEvent2, err := utils.GetEventFromLogs(receipt2.Logs, validatorSetSig.ParseDelivered)
	Expect(err).Should(BeNil())
	Expect(deliveredEvent2.TargetContractAddress).Should(Equal(exampleERC20ContractAddressA))
	Expect(deliveredEvent2.Nonce).Should(Equal(big.NewInt(1)))

	endingBalance, err = exampleERC20A.BalanceOf(nil, validatorSetSigContractAddress)
	Expect(err).Should(BeNil())
	Expect(endingBalance).Should(Equal(big.NewInt(150)))

	startingBalanceB, err = exampleERC20B.BalanceOf(
		&bind.CallOpts{}, validatorSetSigContractAddress)
	Expect(err).Should(BeNil())
	Expect(startingBalanceB.Cmp(big.NewInt(0))).Should(BeZero())

	// ************************************************************************************************
	// Test Case 2: validatorChain (subnetB) == targetChain (subnetB)
	// ************************************************************************************************

	// Send the third transaction where the validatorSetSig contract expects validator signatures
	// from the same chain that it is deployed on.
	receipt3 := utils.ExecuteValidatorSetSigCallAndVerify(
		ctx,
		network,
		subnetB,
		subnetB,
		validatorSetSigContractAddress2,
		fundedKey,
		&offchainMessages[2],
		true,
	)

	// Confirm that the Delivered event is emitted and that validatorSetSig2 contract has a balance
	// of 100 of ERC20
	deliveredEvent3, err := utils.GetEventFromLogs(receipt3.Logs, validatorSetSig2.ParseDelivered)
	Expect(err).Should(BeNil())
	Expect(deliveredEvent3.TargetContractAddress).Should(Equal(exampleERC20ContractAddressB))
	Expect(deliveredEvent.Nonce.Cmp(big.NewInt(0))).Should(BeZero())

	endingBalance, err = exampleERC20B.BalanceOf(nil, validatorSetSigContractAddress2)
	Expect(err).Should(BeNil())
	Expect(endingBalance).Should(Equal(big.NewInt(100)))
}
