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
 * Registers a erc20 token staking validator on a L1. The steps are as follows:
 * - Deploy the ERCTokenStakingManager
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
func ERC20TokenStakingManager(network *localnetwork.LocalNetwork) {
	// Get the L1s info
	cChainInfo := network.GetPrimaryNetworkInfo()
	l1AInfo, _ := network.GetTwoL1s()
	_, fundedKey := network.GetFundedAccountInfo()
	pChainInfo := utils.GetPChainInfo(cChainInfo)

	signatureAggregator := utils.NewSignatureAggregator(
		cChainInfo.NodeURIs[0],
		[]ids.ID{
			l1AInfo.L1ID,
		},
	)
	ctx := context.Background()

	// Deploy the staking manager contract
	stakingManagerAddress, stakingManager, _, erc20 := utils.DeployAndInitializeERC20TokenStakingManager(
		ctx,
		fundedKey,
		l1AInfo,
		pChainInfo,
	)

	nodes := utils.ConvertSubnet(
		ctx,
		l1AInfo,
		network.GetPChainWallet(),
		stakingManagerAddress,
		fundedKey,
	)

	// Initialize the validator set on the L1
	log.Println("Initializing validator set")
	initialValidationIDs := utils.InitializeERC20TokenValidatorSet(
		ctx,
		fundedKey,
		l1AInfo,
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
	utils.InitializeAndCompleteEndInitialERC20Validation(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
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
	validationID := utils.InitializeAndCompleteERC20ValidatorRegistration(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		stakingManager,
		stakingManagerAddress,
		erc20,
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

		receipt := utils.InitializeERC20DelegatorRegistration(
			ctx,
			fundedKey,
			l1AInfo,
			validationID,
			delegatorStake,
			erc20,
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
		signedWarpMessage := utils.ConstructSignedWarpMessage(context.Background(), receipt, l1AInfo, pChainInfo)

		// Issue a tx to update the validator's weight on the P-Chain
		network.GetPChainWallet().IssueSetSubnetValidatorWeightTx(signedWarpMessage.Bytes())
		utils.PChainProposerVMWorkaround(network.GetPChainWallet())
		utils.AdvanceProposerVM(ctx, l1AInfo, fundedKey, 5)

		// Construct a SubnetValidatorWeightUpdateMessage Warp message from the P-Chain
		registrationSignedMessage := utils.ConstructSubnetValidatorWeightUpdateMessage(
			validationID,
			nonce,
			newValidatorWeight,
			l1AInfo,
			pChainInfo,
			signatureAggregator,
			network.GetNetworkID(),
		)

		// Deliver the Warp message to the L1
		receipt = utils.CompleteERC20DelegatorRegistration(
			ctx,
			fundedKey,
			delegationID,
			l1AInfo,
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
		receipt := utils.InitializeEndERC20Delegation(
			ctx,
			fundedKey,
			l1AInfo,
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
		signedWarpMessage := utils.ConstructSignedWarpMessage(context.Background(), receipt, l1AInfo, pChainInfo)
		Expect(err).Should(BeNil())

		// Issue a tx to update the validator's weight on the P-Chain
		network.GetPChainWallet().IssueSetSubnetValidatorWeightTx(signedWarpMessage.Bytes())
		utils.PChainProposerVMWorkaround(network.GetPChainWallet())
		utils.AdvanceProposerVM(ctx, l1AInfo, fundedKey, 5)

		// Construct a SubnetValidatorWeightUpdateMessage Warp message from the P-Chain
		signedMessage := utils.ConstructSubnetValidatorWeightUpdateMessage(
			validationID,
			nonce,
			nodes[0].Weight,
			l1AInfo,
			pChainInfo,
			signatureAggregator,
			network.GetNetworkID(),
		)

		// Deliver the Warp message to the L1
		receipt = utils.CompleteEndERC20Delegation(
			ctx,
			fundedKey,
			delegationID,
			l1AInfo,
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
	utils.InitializeAndCompleteEndERC20Validation(
		ctx,
		signatureAggregator,
		fundedKey,
		l1AInfo,
		pChainInfo,
		stakingManager,
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
