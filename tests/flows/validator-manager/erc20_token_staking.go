package staking

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	exampleerc20 "github.com/ava-labs/teleporter/abi-bindings/go/mocks/ExampleERC20"
	erc20stakingmanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/ERC20TokenStakingManager"
	validatormanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/ValidatorManager"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

/*
 * Registers a erc20 token staking validator on a subnet. The steps are as follows:
 * - Deploy the ERCTokenStakingManager
 * - Initiate validator registration
 * - Deliver the Warp message to the P-Chain (not implemented)
 * - Aggregate P-Chain signatures on the response Warp message
 * - Deliver the Warp message to the subnet
 * - Verify that the validator is registered in the staking contract
 *
 * Delists the validator from the subnet. The steps are as follows:
 * - Initiate validator delisting
 * - Deliver the Warp message to the P-Chain (not implemented)
 * - Aggregate P-Chain signatures on the response Warp message
 * - Deliver the Warp message to the subnet
 * - Verify that the validator is delisted from the staking contract
 */
func ERC20TokenStakingManager(network *localnetwork.LocalNetwork) {
	// Get the subnets info
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := network.GetTwoSubnets()
	_, fundedKey := network.GetFundedAccountInfo()
	pChainInfo := utils.GetPChainInfo(cChainInfo)

	signatureAggregator := utils.NewSignatureAggregator(
		cChainInfo.NodeURIs[0],
		[]ids.ID{
			subnetAInfo.SubnetID,
		},
	)
	ctx := context.Background()

	nodes, initialValidationIDs, _ := network.ConvertSubnet(
		ctx,
		subnetAInfo,
		utils.ERC20TokenStakingManager,
		[]uint64{units.Schmeckle, 1000 * units.Schmeckle}, // Choose weights to avoid validator churn limits
		fundedKey,
		false,
	)
	validatorManagerAddress := network.GetValidatorManager(subnetAInfo.SubnetID)
	validatorManager, err := validatormanager.NewValidatorManager(
		validatorManagerAddress,
		subnetAInfo.RPCClient,
	)
	Expect(err).Should(BeNil())
	securityModuleAddress, err := validatorManager.GetSecurityModule(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	securityModule, err := erc20stakingmanager.NewERC20TokenStakingManager(securityModuleAddress, subnetAInfo.RPCClient)
	Expect(err).Should(BeNil())

	erc20Address, err := securityModule.Erc20(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	erc20, err := exampleerc20.NewExampleERC20(erc20Address, subnetAInfo.RPCClient)
	Expect(err).Should(BeNil())

	//
	// Delist one initial validator
	//
	utils.InitializeAndCompleteEndInitialACP99PoSValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		securityModule,
		securityModuleAddress,
		validatorManager,
		initialValidationIDs[0],
		0,
		nodes[0].Weight,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	//
	// Register the validator as PoS
	//
	expiry := uint64(time.Now().Add(24 * time.Hour).Unix())
	validationID := utils.InitializeAndCompleteACP99ERC20ValidatorRegistration(
		ctx,
		signatureAggregator,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		validatorManager,
		erc20,
		expiry,
		nodes[0],
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)
	validatorStartTime := time.Now()
	time.Sleep(5 * time.Second)

	//
	// Register a delegator
	//
	// var delegationID ids.ID
	// {
	// 	log.Println("Registering delegator")
	// 	delegatorStake, err := erc20StakingManager.WeightToValue(
	// 		&bind.CallOpts{},
	// 		nodes[0].Weight,
	// 	)
	// 	Expect(err).Should(BeNil())
	// 	delegatorStake.Div(delegatorStake, big.NewInt(10))
	// 	delegatorWeight, err := erc20StakingManager.ValueToWeight(
	// 		&bind.CallOpts{},
	// 		delegatorStake,
	// 	)
	// 	Expect(err).Should(BeNil())
	// 	newValidatorWeight := nodes[0].Weight + delegatorWeight

	// 	nonce := uint64(1)

	// 	receipt := utils.InitializeERC20DelegatorRegistration(
	// 		ctx,
	// 		fundedKey,
	// 		subnetAInfo,
	// 		validationID,
	// 		delegatorStake,
	// 		erc20,
	// 		stakingManagerAddress,
	// 		erc20StakingManager,
	// 	)
	// 	initRegistrationEvent, err := utils.GetEventFromLogs(
	// 		receipt.Logs,
	// 		erc20StakingManager.ParseDelegatorAdded,
	// 	)
	// 	Expect(err).Should(BeNil())
	// 	delegationID = initRegistrationEvent.DelegationID

	// 	// Gather subnet-evm Warp signatures for the SubnetValidatorWeightUpdateMessage & relay to the P-Chain
	// 	signedWarpMessage := utils.ConstructSignedWarpMessage(
	// 		context.Background(),
	// 		receipt,
	// 		subnetAInfo,
	// 		pChainInfo,
	// 		nil,
	// 		network.GetSignatureAggregator(),
	// 	)

	// 	// Issue a tx to update the validator's weight on the P-Chain
	// 	network.GetPChainWallet().IssueSetL1ValidatorWeightTx(signedWarpMessage.Bytes())
	// 	utils.PChainProposerVMWorkaround(network.GetPChainWallet())
	// 	utils.AdvanceProposerVM(ctx, subnetAInfo, fundedKey, 5)

	// 	// Construct a SubnetValidatorWeightUpdateMessage Warp message from the P-Chain
	// 	registrationSignedMessage := utils.ConstructSubnetValidatorWeightUpdateMessage(
	// 		validationID,
	// 		nonce,
	// 		newValidatorWeight,
	// 		subnetAInfo,
	// 		pChainInfo,
	// 		signatureAggregator,
	// 		network.GetNetworkID(),
	// 	)

	// 	// Deliver the Warp message to the subnet
	// 	receipt = utils.CompleteDelegatorRegistration(
	// 		ctx,
	// 		fundedKey,
	// 		delegationID,
	// 		subnetAInfo,
	// 		stakingManagerAddress,
	// 		registrationSignedMessage,
	// 	)
	// 	// Check that the validator is registered in the staking contract
	// 	registrationEvent, err := utils.GetEventFromLogs(
	// 		receipt.Logs,
	// 		erc20StakingManager.ParseDelegatorRegistered,
	// 	)
	// 	Expect(err).Should(BeNil())
	// 	Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
	// 	Expect(registrationEvent.DelegationID[:]).Should(Equal(delegationID[:]))
	// }

	// //
	// // Delist the delegator
	// //
	// {
	// 	log.Println("Delisting delegator")
	// 	nonce := uint64(2)
	// 	receipt := utils.InitializeEndDelegation(
	// 		ctx,
	// 		fundedKey,
	// 		subnetAInfo,
	// 		stakingManagerAddress,
	// 		delegationID,
	// 	)
	// 	delegatorRemovalEvent, err := utils.GetEventFromLogs(
	// 		receipt.Logs,
	// 		erc20StakingManager.ParseDelegatorRemovalInitialized,
	// 	)
	// 	Expect(err).Should(BeNil())
	// 	Expect(delegatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
	// 	Expect(delegatorRemovalEvent.DelegationID[:]).Should(Equal(delegationID[:]))

	// 	// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
	// 	// (Sending to the P-Chain will be skipped for now)
	// 	signedWarpMessage := utils.ConstructSignedWarpMessage(
	// 		context.Background(),
	// 		receipt,
	// 		subnetAInfo,
	// 		pChainInfo,
	// 		nil,
	// 		network.GetSignatureAggregator(),
	// 	)
	// 	Expect(err).Should(BeNil())

	// 	// Issue a tx to update the validator's weight on the P-Chain
	// 	network.GetPChainWallet().IssueSetL1ValidatorWeightTx(signedWarpMessage.Bytes())
	// 	utils.PChainProposerVMWorkaround(network.GetPChainWallet())
	// 	utils.AdvanceProposerVM(ctx, subnetAInfo, fundedKey, 5)

	// 	// Construct a SubnetValidatorWeightUpdateMessage Warp message from the P-Chain
	// 	signedMessage := utils.ConstructSubnetValidatorWeightUpdateMessage(
	// 		validationID,
	// 		nonce,
	// 		nodes[0].Weight,
	// 		subnetAInfo,
	// 		pChainInfo,
	// 		signatureAggregator,
	// 		network.GetNetworkID(),
	// 	)

	// 	// Deliver the Warp message to the subnet
	// 	receipt = utils.CompleteEndDelegation(
	// 		ctx,
	// 		fundedKey,
	// 		delegationID,
	// 		subnetAInfo,
	// 		stakingManagerAddress,
	// 		signedMessage,
	// 	)

	// 	// Check that the delegator has been delisted from the staking contract
	// 	registrationEvent, err := utils.GetEventFromLogs(
	// 		receipt.Logs,
	// 		erc20StakingManager.ParseDelegationEnded,
	// 	)
	// 	Expect(err).Should(BeNil())
	// 	Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
	// 	Expect(registrationEvent.DelegationID[:]).Should(Equal(delegationID[:]))
	// }

	//
	// Delist the validator
	//
	utils.InitializeAndCompleteEndACP99PoSValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		securityModule,
		securityModuleAddress,
		validatorManager,
		validationID,
		expiry,
		nodes[0],
		1,
		true,
		validatorStartTime,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)
}
