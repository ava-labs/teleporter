package staking

import (
	"context"
	"log"
	"math/big"
	"sort"
	"time"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/teleporter/tests/interfaces"
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
func ERC20TokenStakingManager(network interfaces.LocalNetwork) {
	// Get the subnets info
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := utils.GetTwoSubnets(network)
	_, fundedKey := network.GetFundedAccountInfo()
	pChainInfo := utils.GetPChainInfo(cChainInfo)

	signatureAggregator := utils.NewSignatureAggregator(
		cChainInfo.NodeURIs[0],
		[]ids.ID{
			subnetAInfo.SubnetID,
		},
	)
	ctx := context.Background()

	// Deploy the staking manager contract
	stakingManagerAddress, stakingManager, _, erc20 := utils.DeployAndInitializeERC20TokenStakingManager(
		ctx,
		fundedKey,
		subnetAInfo,
		pChainInfo,
	)

	var nodes []utils.Node

	// Remove the current validators before converting the subnet
	for _, uri := range subnetAInfo.NodeURIs {
		infoClient := info.NewClient(uri)
		nodeID, nodePoP, err := infoClient.GetNodeID(ctx)
		Expect(err).Should(BeNil())
		nodes = append(nodes, utils.Node{
			NodeID:  nodeID,
			NodePoP: nodePoP,
		})

		_, err = network.GetPChainWallet().IssueRemoveSubnetValidatorTx(
			nodeID,
			subnetAInfo.SubnetID,
		)
		Expect(err).Should(BeNil())
	}

	// Sort the nodeIDs so that the subnet conversion ID matches the P-Chain
	sort.Slice(nodes, func(i, j int) bool {
		return string(nodes[i].NodeID.Bytes()) < string(nodes[j].NodeID.Bytes())
	})

	weights := make([]uint64, len(nodes))
	totalWeight := uint64(len(nodes)-1) * units.Schmeckle
	for i := 0; i < len(nodes)-1; i++ {
		weights[i] = units.Schmeckle
		totalWeight += units.Schmeckle
	}
	// Set the last node's weight such that removing any other node will not violate the churn limit
	weights[len(nodes)-1] = 4 * totalWeight

	// Construct the convert subnet info
	destAddr, err := address.ParseToID(utils.DefaultPChainAddress)
	Expect(err).Should(BeNil())
	vdrs := make([]*txs.ConvertSubnetValidator, len(nodes))
	for i, node := range nodes {
		vdrs[i] = &txs.ConvertSubnetValidator{
			NodeID:  node.NodeID.Bytes(),
			Weight:  weights[i],
			Balance: units.Avax * 100,
			Signer:  *node.NodePoP,
			RemainingBalanceOwner: message.PChainOwner{
				Threshold: 1,
				Addresses: []ids.ShortID{destAddr},
			},
			DeactivationOwner: message.PChainOwner{
				Threshold: 1,
				Addresses: []ids.ShortID{destAddr},
			},
		}
	}

	log.Println("Issuing ConvertSubnetTx")
	_, err = network.GetPChainWallet().IssueConvertSubnetTx(
		subnetAInfo.SubnetID,
		subnetAInfo.BlockchainID,
		stakingManagerAddress[:],
		vdrs,
	)
	Expect(err).Should(BeNil())

	utils.PChainProposerVMWorkaround(network)
	utils.AdvanceProposerVM(ctx, subnetAInfo, fundedKey, 5)

	// Initialize the validator set on the subnet
	log.Println("Initializing validator set")
	initialValidationIDs := utils.InitializeERC20TokenValidatorSet(
		ctx,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		stakingManager,
		stakingManagerAddress,
		network,
		signatureAggregator,
		nodes,
		weights,
	)

	//
	// Delist one initial validators
	//
	utils.InitializeAndCompleteEndInitialERC20Validation(
		ctx,
		network,
		signatureAggregator,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		stakingManager,
		stakingManagerAddress,
		initialValidationIDs[0],
		0,
		weights[0],
	)

	//
	// Register the validator as PoS
	//
	stakeAmount, err := stakingManager.WeightToValue(
		&bind.CallOpts{},
		weights[0],
	)
	Expect(err).Should(BeNil())
	expiry := uint64(time.Now().Add(24 * time.Hour).Unix())
	validationID := utils.InitializeAndCompleteERC20ValidatorRegistration(
		ctx,
		network,
		signatureAggregator,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		stakingManager,
		stakingManagerAddress,
		erc20,
		stakeAmount,
		weights[0],
		expiry,
		nodes[0],
	)

	//
	// Register a delegator
	//
	var delegationID ids.ID
	{
		log.Println("Registering delegator")
		delegatorStake, err := stakingManager.WeightToValue(
			&bind.CallOpts{},
			weights[0],
		)
		Expect(err).Should(BeNil())
		delegatorStake.Div(delegatorStake, big.NewInt(10))
		delegatorWeight, err := stakingManager.ValueToWeight(
			&bind.CallOpts{},
			delegatorStake,
		)
		Expect(err).Should(BeNil())
		newValidatorWeight := weights[0] + delegatorWeight

		nonce := uint64(1)

		receipt := utils.InitializeERC20DelegatorRegistration(
			ctx,
			fundedKey,
			subnetAInfo,
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
		// (Sending to the P-Chain will be skipped for now)
		signedWarpMessage := network.ConstructSignedWarpMessage(context.Background(), receipt, subnetAInfo, pChainInfo)

		// Issue a tx to update the validator's weight on the P-Chain
		network.GetPChainWallet().IssueSetSubnetValidatorWeightTx(signedWarpMessage.Bytes())
		utils.PChainProposerVMWorkaround(network)
		utils.AdvanceProposerVM(ctx, subnetAInfo, fundedKey, 5)

		// Construct a SubnetValidatorWeightUpdateMessage Warp message from the P-Chain
		registrationSignedMessage := utils.ConstructSubnetValidatorWeightUpdateMessage(
			validationID,
			nonce,
			newValidatorWeight,
			subnetAInfo,
			pChainInfo,
			network,
			signatureAggregator,
		)

		// Deliver the Warp message to the subnet
		receipt = utils.CompleteERC20DelegatorRegistration(
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
			subnetAInfo,
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
		signedWarpMessage := network.ConstructSignedWarpMessage(context.Background(), receipt, subnetAInfo, pChainInfo)
		Expect(err).Should(BeNil())

		// Issue a tx to update the validator's weight on the P-Chain
		network.GetPChainWallet().IssueSetSubnetValidatorWeightTx(signedWarpMessage.Bytes())

		utils.PChainProposerVMWorkaround(network)
		utils.AdvanceProposerVM(ctx, subnetAInfo, fundedKey, 5)

		// Construct a SubnetValidatorWeightUpdateMessage Warp message from the P-Chain
		signedMessage := utils.ConstructSubnetValidatorWeightUpdateMessage(
			validationID,
			nonce,
			weights[0],
			subnetAInfo,
			pChainInfo,
			network,
			signatureAggregator,
		)

		// Deliver the Warp message to the subnet
		receipt = utils.CompleteEndERC20Delegation(
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
		network,
		signatureAggregator,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		stakingManager,
		stakingManagerAddress,
		validationID,
		expiry,
		nodes[0],
		weights[0],
		1,
	)
}
