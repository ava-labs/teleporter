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
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/stakeable"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
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
	var weights []uint64
	for _, uri := range subnetAInfo.NodeURIs {
		infoClient := info.NewClient(uri)
		nodeID, nodePoP, err := infoClient.GetNodeID(ctx)
		Expect(err).Should(BeNil())
		nodes = append(nodes, utils.Node{
			NodeID:  nodeID,
			NodePoP: nodePoP,
		})
		weights = append(weights, units.Schmeckle)

		// Remove the current validators before converting the subnet
		_, err = network.GetPChainWallet().IssueRemoveSubnetValidatorTx(
			nodeID,
			subnetAInfo.SubnetID,
		)
		Expect(err).Should(BeNil())
	}

	// Sort the nodeIDs so that the subnet conversion ID matches the P-Chain
	log.Println("nodes before sorting")
	log.Println("nodeID1: ", nodes[0].NodeID.Bytes())
	log.Println("nodeID2: ", nodes[1].NodeID.Bytes())
	sort.Slice(nodes, func(i, j int) bool {
		return string(nodes[i].NodeID.Bytes()) < string(nodes[j].NodeID.Bytes())
	})
	log.Println("nodes after sorting")
	log.Println("nodeID1: ", nodes[0].NodeID.Bytes())
	log.Println("nodeID2: ", nodes[1].NodeID.Bytes())

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
	vdrs := []*txs.ConvertSubnetValidator{
		{
			NodeID:  nodes[0].NodeID.Bytes(),
			Weight:  weights[0],
			Balance: units.Avax * 100,
			Signer:  *nodes[0].NodePoP,
			RemainingBalanceOwner: message.PChainOwner{
				Threshold: 1,
				Addresses: []ids.ShortID{destAddr},
			},
			DeactivationOwner: message.PChainOwner{
				Threshold: 1,
				Addresses: []ids.ShortID{destAddr},
			},
		},
		{
			NodeID:  nodes[1].NodeID.Bytes(),
			Weight:  weights[1],
			Balance: units.Avax * 100,
			Signer:  *nodes[1].NodePoP,
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

	// DEBUG
	log.Println("convert subnet tx info")
	data := message.SubnetConversionData{
		SubnetID:       subnetAInfo.SubnetID,
		ManagerChainID: subnetAInfo.BlockchainID,
		ManagerAddress: stakingManagerAddress[:],
		Validators: []message.SubnetConversionValidatorData{
			{
				NodeID:       nodes[0].NodeID.Bytes(),
				BLSPublicKey: nodes[0].NodePoP.PublicKey,
				Weight:       weights[0],
			},
			{
				NodeID:       nodes[1].NodeID.Bytes(),
				BLSPublicKey: nodes[1].NodePoP.PublicKey,
				Weight:       weights[1],
			},
		},
	}
	conversionID, _ := message.SubnetConversionID(data)
	log.Println(data)
	log.Printf("conversionID: %s\n", conversionID.String())
	data = message.SubnetConversionData{
		SubnetID:       subnetAInfo.SubnetID,
		ManagerChainID: subnetAInfo.BlockchainID,
		ManagerAddress: stakingManagerAddress[:],
		Validators: []message.SubnetConversionValidatorData{
			{
				NodeID:       nodes[1].NodeID.Bytes(),
				BLSPublicKey: nodes[1].NodePoP.PublicKey,
				Weight:       weights[1],
			},
			{
				NodeID:       nodes[0].NodeID.Bytes(),
				BLSPublicKey: nodes[0].NodePoP.PublicKey,
				Weight:       weights[0],
			},
		},
	}
	conversionID, _ = message.SubnetConversionID(data)
	log.Println(data)
	log.Printf("conversionID: %s\n", conversionID.String())
	// END DEBUG

	// // Workaround current block map rules
	time.Sleep(30 * time.Second)

	pBuilder := network.GetPChainWallet().Builder()
	pContext := pBuilder.Context()
	avaxAssetID := pContext.AVAXAssetID
	locktime := uint64(time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC).Unix())
	amount := 500 * units.MilliAvax
	_, err = network.GetPChainWallet().IssueBaseTx([]*avax.TransferableOutput{
		{
			Asset: avax.Asset{
				ID: avaxAssetID,
			},
			Out: &stakeable.LockOut{
				Locktime: locktime,
				TransferableOut: &secp256k1fx.TransferOutput{
					Amt: amount,
					OutputOwners: secp256k1fx.OutputOwners{
						Threshold: 1,
						Addrs: []ids.ShortID{
							destAddr,
						},
					},
				},
			},
		},
	})
	Expect(err).Should(BeNil())
	_, err = network.GetPChainWallet().IssueBaseTx([]*avax.TransferableOutput{
		{
			Asset: avax.Asset{
				ID: avaxAssetID,
			},
			Out: &stakeable.LockOut{
				Locktime: locktime,
				TransferableOut: &secp256k1fx.TransferOutput{
					Amt: amount,
					OutputOwners: secp256k1fx.OutputOwners{
						Threshold: 1,
						Addrs: []ids.ShortID{
							destAddr,
						},
					},
				},
			},
		},
	})
	Expect(err).Should(BeNil())
	// End workaround

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
		nodes,
		weights,
	)

	//
	// Delist all initial validators
	//
	Expect(false).Should(BeTrue())

	//
	// Register a validator
	//
	stakeAmount := new(big.Int).SetUint64(utils.DefaultMinStakeAmount)
	weight, err := stakingManager.ValueToWeight(
		&bind.CallOpts{},
		stakeAmount,
	)
	Expect(err).Should(BeNil())
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
		weight,
		1,
	)
}
