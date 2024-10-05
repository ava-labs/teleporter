package staking

import (
	"context"
	"sort"
	"time"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	subnetEvmUtils "github.com/ava-labs/subnet-evm/tests/utils"
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

	for _, uri := range subnetAInfo.NodeURIs {
		infoClient := info.NewClient(uri)
		nodeID, nodePoP, err := infoClient.GetNodeID(ctx)
		Expect(err).Should(BeNil())
		nodes = append(nodes, utils.Node{
			NodeID:  nodeID,
			NodePoP: nodePoP,
		})

		// Remove the current validators before converting the subnet
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

	// TODONOW: issue a P-Chain ConvertSubnetTx
	destAddrStr := "P-local18jma8ppw3nhx5r4ap8clazz0dps7rv5u00z96u"
	destAddr, err := address.ParseToID(destAddrStr)
	Expect(err).Should(BeNil())
	// vdrs := make([]*txs.ConvertSubnetValidator, len(nodes))
	// for i, node := range nodes {
	// 	vdrs[i] = &txs.ConvertSubnetValidator{
	// 		NodeID:  node.NodeID.Bytes(),
	// 		Weight:  weights[i],
	// 		Balance: units.Avax * 100,
	// 		Signer:  *node.NodePoP,
	// 		RemainingBalanceOwner: message.PChainOwner{
	// 			Threshold: 1,
	// 			Addresses: []ids.ShortID{destAddr},
	// 		},
	// 		DeactivationOwner: message.PChainOwner{
	// 			Threshold: 1,
	// 			Addresses: []ids.ShortID{destAddr},
	// 		},
	// 	}
	// }

	// Set just the last node (with a high weight) to the initial validator list
	vdrs := []*txs.ConvertSubnetValidator{
		{
			NodeID:  nodes[len(nodes)-1].NodeID.Bytes(),
			Weight:  weights[len(nodes)-1],
			Balance: units.Avax * 100,
			Signer:  *nodes[len(nodes)-1].NodePoP,
			RemainingBalanceOwner: message.PChainOwner{
				Threshold: 1,
				Addresses: []ids.ShortID{destAddr},
			},
			DeactivationOwner: message.PChainOwner{
				Threshold: 1,
				Addresses: []ids.ShortID{destAddr},
			},
		},
	}

	_, err = network.GetPChainWallet().IssueConvertSubnetTx(
		subnetAInfo.SubnetID,
		subnetAInfo.BlockchainID,
		stakingManagerAddress[:],
		vdrs,
	)
	Expect(err).Should(BeNil())

	utils.PChainProposerVMWorkaround(network)

	// Issue txs on the subnet to advance the proposer vm
	for i := 0; i < 5; i++ {
		err = subnetEvmUtils.IssueTxsToActivateProposerVMFork(
			ctx, subnetAInfo.EVMChainID, fundedKey, subnetAInfo.WSClient,
		)
		Expect(err).Should(BeNil())
	}

	// Initialize the validator set on the subnet
	_ = utils.InitializeERC20TokenValidatorSet(
		ctx,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		stakingManager,
		stakingManagerAddress,
		network,
		signatureAggregator,
		nodes[len(nodes)-1:],
		weights[len(nodes)-1:],
	)

	// Delisting an initial validator does not work due to the way Warp signatures are requested for inital validators
	//
	// Delist one initial validators
	//
	// utils.InitializeAndCompleteEndERC20Validation(
	// 	ctx,
	// 	network,
	// 	signatureAggregator,
	// 	fundedKey,
	// 	subnetAInfo,
	// 	pChainInfo,
	// 	stakingManager,
	// 	stakingManagerAddress,
	// 	initialValidationIDs[0],
	// 	weights[0],
	// 	1,
	// )

	//
	// Register a validator
	//
	stakeAmount, err := stakingManager.WeightToValue(
		&bind.CallOpts{},
		weights[0],
	)
	Expect(err).Should(BeNil())
	// TODONOW: pass the validatorID pre image to the signature aggregator
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
