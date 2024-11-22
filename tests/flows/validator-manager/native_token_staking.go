package staking

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	nativetokenstakingmanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/NativeTokenStakingManager"
	iposvalidatormanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/interfaces/IPoSValidatorManager"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

/*
 * Registers a native token staking validator on a subnet. The steps are as follows:
 * - Deploy the NativeTokenStakingManager
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
func NativeTokenStakingManager(network *localnetwork.LocalNetwork) {
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
		utils.NativeTokenStakingManager,
		[]uint64{units.Schmeckle, 1000 * units.Schmeckle}, // Choose weights to avoid validator churn limits
		fundedKey,
		false,
	)
	stakingManagerAddress := network.GetValidatorManager(subnetAInfo.SubnetID)
	nativeStakingManager, err := nativetokenstakingmanager.NewNativeTokenStakingManager(
		stakingManagerAddress,
		subnetAInfo.RPCClient,
	)
	Expect(err).Should(BeNil())
	utils.AddNativeMinterAdmin(ctx, subnetAInfo, fundedKey, stakingManagerAddress)

	//
	// Delist one initial validator
	//
	posStakingManager, err := iposvalidatormanager.NewIPoSValidatorManager(stakingManagerAddress, subnetAInfo.RPCClient)
	Expect(err).Should(BeNil())
	utils.InitializeAndCompleteEndInitialPoSValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		posStakingManager,
		stakingManagerAddress,
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
	validationID := utils.InitializeAndCompleteNativeValidatorRegistration(
		ctx,
		signatureAggregator,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		nativeStakingManager,
		stakingManagerAddress,
		expiry,
		nodes[0],
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)
	validatorStartTime := time.Now()

	//
	// Register a delegator
	//
	var delegationID ids.ID
	{
		log.Println("Registering delegator")
		delegatorStake, err := nativeStakingManager.WeightToValue(
			&bind.CallOpts{},
			nodes[0].Weight,
		)
		Expect(err).Should(BeNil())
		delegatorStake.Div(delegatorStake, big.NewInt(10))
		delegatorWeight, err := nativeStakingManager.ValueToWeight(
			&bind.CallOpts{},
			delegatorStake,
		)
		Expect(err).Should(BeNil())
		newValidatorWeight := nodes[0].Weight + delegatorWeight

		nonce := uint64(1)

		receipt := utils.InitializeNativeDelegatorRegistration(
			ctx,
			fundedKey,
			subnetAInfo,
			validationID,
			delegatorStake,
			stakingManagerAddress,
			nativeStakingManager,
		)
		initRegistrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			nativeStakingManager.ParseDelegatorAdded,
		)
		Expect(err).Should(BeNil())
		delegationID = initRegistrationEvent.DelegationID

		aggregator := network.GetSignatureAggregator()
		defer aggregator.Shutdown()

		// Gather subnet-evm Warp signatures for the SubnetValidatorWeightUpdateMessage & relay to the P-Chain
		signedWarpMessage := utils.ConstructSignedWarpMessage(
			context.Background(),
			receipt,
			subnetAInfo,
			pChainInfo,
			nil,
			aggregator,
		)

		// Issue a tx to update the validator's weight on the P-Chain
		network.GetPChainWallet().IssueSetL1ValidatorWeightTx(signedWarpMessage.Bytes())
		utils.PChainProposerVMWorkaround(network.GetPChainWallet())
		utils.AdvanceProposerVM(ctx, subnetAInfo, fundedKey, 5)

		// Construct a SubnetValidatorWeightUpdateMessage Warp message from the P-Chain
		registrationSignedMessage := utils.ConstructSubnetValidatorWeightUpdateMessage(
			validationID,
			nonce,
			newValidatorWeight,
			subnetAInfo,
			pChainInfo,
			signatureAggregator,
			network.GetNetworkID(),
		)

		// Deliver the Warp message to the subnet
		receipt = utils.CompleteDelegatorRegistration(
			ctx,
			fundedKey,
			delegationID,
			subnetAInfo,
			stakingManagerAddress,
			registrationSignedMessage,
		)
		// Check that the validator is registered in the staking contract
		registrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			nativeStakingManager.ParseDelegatorRegistered,
		)
		Expect(err).Should(BeNil())
		Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
		Expect(registrationEvent.DelegationID[:]).Should(Equal(delegationID[:]))
	}
	//
	// Delist the delegator
	//
	{
		log.Println("Delisting delegator")
		nonce := uint64(2)
		receipt := utils.InitializeEndDelegation(
			ctx,
			fundedKey,
			subnetAInfo,
			stakingManagerAddress,
			delegationID,
		)
		delegatorRemovalEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			nativeStakingManager.ParseDelegatorRemovalInitialized,
		)
		Expect(err).Should(BeNil())
		Expect(delegatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
		Expect(delegatorRemovalEvent.DelegationID[:]).Should(Equal(delegationID[:]))

		aggregator := network.GetSignatureAggregator()
		defer aggregator.Shutdown()

		// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
		// (Sending to the P-Chain will be skipped for now)
		signedWarpMessage := utils.ConstructSignedWarpMessage(
			context.Background(),
			receipt,
			subnetAInfo,
			pChainInfo,
			nil,
			aggregator,
		)
		Expect(err).Should(BeNil())

		// Issue a tx to update the validator's weight on the P-Chain
		network.GetPChainWallet().IssueSetL1ValidatorWeightTx(signedWarpMessage.Bytes())
		utils.PChainProposerVMWorkaround(network.GetPChainWallet())
		utils.AdvanceProposerVM(ctx, subnetAInfo, fundedKey, 5)

		// Construct a SubnetValidatorWeightUpdateMessage Warp message from the P-Chain
		signedMessage := utils.ConstructSubnetValidatorWeightUpdateMessage(
			validationID,
			nonce,
			nodes[0].Weight,
			subnetAInfo,
			pChainInfo,
			signatureAggregator,
			network.GetNetworkID(),
		)

		// Deliver the Warp message to the subnet
		receipt = utils.CompleteEndDelegation(
			ctx,
			fundedKey,
			delegationID,
			subnetAInfo,
			stakingManagerAddress,
			signedMessage,
		)

		// Check that the delegator has been delisted from the staking contract
		registrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			nativeStakingManager.ParseDelegationEnded,
		)
		Expect(err).Should(BeNil())
		Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
		Expect(registrationEvent.DelegationID[:]).Should(Equal(delegationID[:]))
	}

	//
	// Delist the validator
	//
	utils.InitializeAndCompleteEndPoSValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		posStakingManager,
		stakingManagerAddress,
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
