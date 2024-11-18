package staking

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

/*
 * Registers a native token staking validator on a L1. The steps are as follows:
 * - Deploy the NativeTokenStakingManager
 * - Initiate validator registration
 * - Deliver the Warp message to the P-Chain (not implemented)
 * - Aggregate P-Chain signatures on the response Warp message
 * - Deliver the Warp message to the L1
 * - Verify that the validator is registered in the staking contract
 *
 * Delists the validator from the L1. The steps are as follows:
 * - Initiate validator delisting
 * - Deliver the Warp message to the P-Chain (not implemented)
 * - Aggregate P-Chain signatures on the response Warp message
 * - Deliver the Warp message to the L1
 * - Verify that the validator is delisted from the staking contract
 */
func NativeTokenStakingManager(network *localnetwork.LocalNetwork) {
	// Get the L1s info
	cChainInfo := network.GetPrimaryNetworkInfo()
	L1AInfo, _ := network.GetTwoL1s()
	_, fundedKey := network.GetFundedAccountInfo()
	pChainInfo := utils.GetPChainInfo(cChainInfo)

	signatureAggregator := utils.NewSignatureAggregator(
		cChainInfo.NodeURIs[0],
		[]ids.ID{
			L1AInfo.L1ID,
		},
	)
	ctx := context.Background()

	// Deploy the staking manager contract
	stakingManagerAddress, stakingManager := utils.DeployAndInitializeNativeTokenStakingManager(
		ctx,
		fundedKey,
		L1AInfo,
		pChainInfo,
	)

	utils.AddNativeMinterAdmin(ctx, L1AInfo, fundedKey, stakingManagerAddress)

	nodes := utils.ConvertSubnet(
		ctx,
		L1AInfo,
		network.GetPChainWallet(),
		stakingManagerAddress,
		fundedKey,
	)

	// Initialize the validator set on the L1
	log.Println("Initializing validator set")
	initialValidationIDs := utils.InitializeNativeTokenValidatorSet(
		ctx,
		fundedKey,
		L1AInfo,
		pChainInfo,
		stakingManager,
		stakingManagerAddress,
		network.GetNetworkID(),
		signatureAggregator,
		nodes,
	)

	//
	// Delist one initial validator
	//
	utils.InitializeAndCompleteEndInitialNativeValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		L1AInfo,
		pChainInfo,
		stakingManager,
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
		L1AInfo,
		pChainInfo,
		stakingManager,
		stakingManagerAddress,
		expiry,
		nodes[0],
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	//
	// Register a delegator
	//
	var delegationID ids.ID
	{
		log.Println("Registering delegator")
		delegatorStake, err := stakingManager.WeightToValue(
			&bind.CallOpts{},
			nodes[0].Weight,
		)
		Expect(err).Should(BeNil())
		delegatorStake.Div(delegatorStake, big.NewInt(10))
		delegatorWeight, err := stakingManager.ValueToWeight(
			&bind.CallOpts{},
			delegatorStake,
		)
		Expect(err).Should(BeNil())
		newValidatorWeight := nodes[0].Weight + delegatorWeight

		nonce := uint64(1)

		receipt := utils.InitializeNativeDelegatorRegistration(
			ctx,
			fundedKey,
			L1AInfo,
			validationID,
			delegatorStake,
			stakingManagerAddress,
			stakingManager,
		)
		initRegistrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			stakingManager.ParseDelegatorAdded,
		)
		Expect(err).Should(BeNil())
		delegationID = initRegistrationEvent.DelegationID

		// Gather subnet-evm Warp signatures for the SubnetValidatorWeightUpdateMessage & relay to the P-Chain
		signedWarpMessage := utils.ConstructSignedWarpMessage(context.Background(), receipt, L1AInfo, pChainInfo)

		// Issue a tx to update the validator's weight on the P-Chain
		network.GetPChainWallet().IssueSetSubnetValidatorWeightTx(signedWarpMessage.Bytes())
		utils.PChainProposerVMWorkaround(network.GetPChainWallet())
		utils.AdvanceProposerVM(ctx, L1AInfo, fundedKey, 5)

		// Construct a SubnetValidatorWeightUpdateMessage Warp message from the P-Chain
		registrationSignedMessage := utils.ConstructSubnetValidatorWeightUpdateMessage(
			validationID,
			nonce,
			newValidatorWeight,
			L1AInfo,
			pChainInfo,
			signatureAggregator,
			network.GetNetworkID(),
		)

		// Deliver the Warp message to the L1
		receipt = utils.CompleteNativeDelegatorRegistration(
			ctx,
			fundedKey,
			delegationID,
			L1AInfo,
			stakingManagerAddress,
			registrationSignedMessage,
		)
		// Check that the validator is registered in the staking contract
		registrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			stakingManager.ParseDelegatorRegistered,
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
		receipt := utils.InitializeEndNativeDelegation(
			ctx,
			fundedKey,
			L1AInfo,
			stakingManager,
			delegationID,
		)
		delegatorRemovalEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			stakingManager.ParseDelegatorRemovalInitialized,
		)
		Expect(err).Should(BeNil())
		Expect(delegatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
		Expect(delegatorRemovalEvent.DelegationID[:]).Should(Equal(delegationID[:]))

		// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
		// (Sending to the P-Chain will be skipped for now)
		signedWarpMessage := utils.ConstructSignedWarpMessage(context.Background(), receipt, L1AInfo, pChainInfo)
		Expect(err).Should(BeNil())

		// Issue a tx to update the validator's weight on the P-Chain
		network.GetPChainWallet().IssueSetSubnetValidatorWeightTx(signedWarpMessage.Bytes())
		utils.PChainProposerVMWorkaround(network.GetPChainWallet())
		utils.AdvanceProposerVM(ctx, L1AInfo, fundedKey, 5)

		// Construct a SubnetValidatorWeightUpdateMessage Warp message from the P-Chain
		signedMessage := utils.ConstructSubnetValidatorWeightUpdateMessage(
			validationID,
			nonce,
			nodes[0].Weight,
			L1AInfo,
			pChainInfo,
			signatureAggregator,
			network.GetNetworkID(),
		)

		// Deliver the Warp message to the L1
		receipt = utils.CompleteEndNativeDelegation(
			ctx,
			fundedKey,
			delegationID,
			L1AInfo,
			stakingManagerAddress,
			signedMessage,
		)

		// Check that the delegator has been delisted from the staking contract
		registrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			stakingManager.ParseDelegationEnded,
		)
		Expect(err).Should(BeNil())
		Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
		Expect(registrationEvent.DelegationID[:]).Should(Equal(delegationID[:]))
	}

	//
	// Delist the validator
	//
	utils.InitializeAndCompleteEndNativeValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		L1AInfo,
		pChainInfo,
		stakingManager,
		stakingManagerAddress,
		validationID,
		expiry,
		nodes[0],
		1,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)
}
