package flows

import (
	"context"
	"encoding/hex"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpMessages "github.com/ava-labs/avalanchego/vms/platformvm/warp/messages"
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	relayerConfig "github.com/ava-labs/awm-relayer/config"
	"github.com/ava-labs/awm-relayer/peers"
	"github.com/ava-labs/awm-relayer/signature-aggregator/aggregator"
	sigAggConfig "github.com/ava-labs/awm-relayer/signature-aggregator/config"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	predicateutils "github.com/ava-labs/subnet-evm/predicate"
	nativetokenstakingmanager "github.com/ava-labs/teleporter/abi-bindings/go/staking/NativeTokenStakingManager"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
	"github.com/prometheus/client_golang/prometheus"
)

// Registers a native token staking validator on a subnet. The steps are as follows:
// - Deploy the NativeTokenStakingManager
// - Initiate validator registration
// - Deliver the Warp message to the P-Chain (not implemented)
// - Aggregate P-Chain signatures on the response Warp message
// - Deliver the Warp message to the subnet
// - Verify that the validator is registered in the staking contract
func NativeTokenStakingManager(network interfaces.LocalNetwork) {
	// Get the subnets info
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := utils.GetTwoSubnets(network)
	_, fundedKey := network.GetFundedAccountInfo()

	// Construct the P-Chain info
	pChainBlockchainID, err := info.NewClient(cChainInfo.NodeURIs[0]).GetBlockchainID(context.Background(), "P")
	Expect(err).Should(BeNil())
	pChainInfo := interfaces.SubnetTestInfo{
		BlockchainID: pChainBlockchainID,
		SubnetID:     ids.Empty,
	}

	// Create the signature aggregator
	logger := logging.NoLog{}
	cfg := sigAggConfig.Config{
		PChainAPI: &relayerConfig.APIConfig{
			BaseURL: cChainInfo.NodeURIs[0],
		},
		InfoAPI: &relayerConfig.APIConfig{
			BaseURL: cChainInfo.NodeURIs[0],
		},
	}
	trackedSubnets := set.NewSet[ids.ID](2)
	trackedSubnets.Add(subnetAInfo.BlockchainID)
	trackedSubnets.Add(ids.Empty) // Primary network subnet ID
	appRequestNetwork, err := peers.NewNetwork(
		logging.Debug,
		prometheus.DefaultRegisterer,
		trackedSubnets,
		&cfg,
	)
	Expect(err).Should(BeNil())

	messageCreator, err := message.NewCreator(
		logger,
		prometheus.DefaultRegisterer,
		constants.DefaultNetworkCompressionType,
		constants.DefaultNetworkMaximumInboundTimeout,
	)
	Expect(err).Should(BeNil())
	signatureAggregator := aggregator.NewSignatureAggregator(
		appRequestNetwork,
		logger,
		messageCreator,
	)
	signatureAggregator.GetSubnetID(subnetAInfo.BlockchainID)

	// Deploy the staking manager contract
	stakingManagerContractAddress, stakingManager := utils.DeployNativeTokenStakingManager(
		context.Background(),
		fundedKey,
		subnetAInfo,
	)
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetAInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.Initialize(
		opts,
		nativetokenstakingmanager.StakingManagerSettings{
			PChainBlockchainID:   pChainInfo.BlockchainID,
			SubnetID:             subnetAInfo.SubnetID,
			MinimumStakeAmount:   big.NewInt(0).SetUint64(1e6),
			MaximumStakeAmount:   big.NewInt(0).SetUint64(10e6),
			MinimumStakeDuration: uint64(24 * time.Hour),
			MaximumHourlyChurn:   1,
			RewardCalculator:     common.Address{},
		},
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(context.Background(), subnetAInfo, tx.Hash())
	// Iniatiate validator registration

	opts.Value = big.NewInt(0).SetUint64(1e18)
	nodeID := ids.GenerateTestID()
	signature := [64]byte{}
	tx, err = stakingManager.InitializeValidatorRegistration(
		opts,
		nodeID,
		uint64(time.Now().Add(24*time.Hour).Unix()),
		signature[:],
	)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(context.Background(), subnetAInfo, tx.Hash())
	registrationInitiatedEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidationPeriodCreated,
	)
	Expect(err).Should(BeNil())
	validationID := ids.ID(registrationInitiatedEvent.ValidationID)

	// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
	// (Sending to the P-Chain will be skipped for now)
	signedWarpMessage := network.ConstructSignedWarpMessage(context.Background(), receipt, subnetAInfo, pChainInfo)
	Expect(err).Should(BeNil())
	log.Info("Constructed signed RegisterSubnetValidator message", "message", signedWarpMessage)
	// Validate the Warp message, (this will be done on the P-Chain in the future)
	// msg, err := warpPayload.ParseAddressedCall(signedWarpMessage.UnsignedMessage.Payload)
	// Expect(err).Should(BeNil())
	// // Check that the addressed call payload is a registered Warp message type
	// var registerValidatorPayload warpMessages.RegisterSubnetValidator
	// ver, err := warpMessages.Codec.Unmarshal(msg.Payload, &registerValidatorPayload)
	// Expect(err).Should(BeNil())
	// Expect(ver).Should(Equal(uint16(warpMessages.CodecVersion)))
	// Expect(registerValidatorPayload.NodeID).Should(Equal(nodeID))
	// Expect(registerValidatorPayload.Weight).Should(Equal(1e18))
	// Expect(registerValidatorPayload.SubnetID).Should(Equal(subnetAInfo.SubnetID))

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	// Query P-Chain validators for the Warp message
	registrationPayload, err := warpMessages.NewSubnetValidatorRegistration(validationID, true)
	Expect(err).Should(BeNil())
	registrationAddressedCall, err := warpPayload.NewAddressedCall(common.Address{}.Bytes(), registrationPayload.Bytes())
	Expect(err).Should(BeNil())
	registrationUnsignedMessage, err := avalancheWarp.NewUnsignedMessage(network.GetNetworkID(), pChainInfo.BlockchainID, registrationAddressedCall.Bytes())
	Expect(err).Should(BeNil())
	log.Info("Constructed unsigned SubnetValidatorRegistration message", "message", hex.EncodeToString(registrationUnsignedMessage.Bytes()))

	registrationSignedMessage, err := signatureAggregator.CreateSignedMessage(
		registrationUnsignedMessage,
		nil,
		subnetAInfo.SubnetID,
		67,
	)
	Expect(err).Should(BeNil())
	log.Info("Constructed signed SubnetValidatorRegistration message", "message", registrationSignedMessage)

	// Deliver the Warp message to the subnet
	abi, err := nativetokenstakingmanager.NativeTokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeValidatorRegistration", uint32(0))
	Expect(err).Should(BeNil())
	gasFeeCap, gasTipCap, nonce := utils.CalculateTxParams(context.Background(), subnetAInfo, utils.PrivateKeyToAddress(fundedKey))
	registrationTx := predicateutils.NewPredicateTx(
		subnetAInfo.EVMChainID,
		nonce,
		&stakingManagerContractAddress,
		2_000_000,
		gasFeeCap,
		gasTipCap,
		big.NewInt(0),
		callData,
		types.AccessList{},
		warp.ContractAddress,
		registrationSignedMessage.Bytes(),
	)
	signedRegistrationTx := utils.SignTransaction(registrationTx, fundedKey, subnetAInfo.EVMChainID)
	receipt = utils.SendTransactionAndWaitForSuccess(context.Background(), subnetAInfo, signedRegistrationTx)

	// Check that the validator is registered in the staking contract
	registrationEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidationPeriodRegistered,
	)
	Expect(err).Should(BeNil())
	Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
}
