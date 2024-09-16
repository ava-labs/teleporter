package staking

import (
	"context"
	"crypto/sha256"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	poavalidatormanager "github.com/ava-labs/teleporter/abi-bindings/go/staking/PoAValidatorManager"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/*
 * Register a PoA validator manager on a L1. The steps are as follows:
 * - Generate random address to be the owner address
 * - Fund native assets to the owner address
 * - Deploy the PoAValidatorManager contract
 * - Attempt to initiate with non owner and check that it fails
 * - Initiate validator registration
 * - Deliver the Warp message to the P-Chain (not implemented)
 * - Aggregate P-Chain signatures on the response Warp message
 * - Deliver the Warp message to the L1
 * - Verify that the validator is registered in the validator manager contract
 *
 * Delists the validator from the L1. The steps are as follows:
 * - Attempt to initiate with non owner and check that it fails
 * - Initiate validator delisting
 * - Deliver the Warp message to the P-Chain (not implemented)
 * - Aggregate P-Chain signatures on the response Warp message
 * - Deliver the Warp message to the L1
 * - Verify that the validator is delisted from the validator manager contract
 */
func PoAValidatorManager(network interfaces.LocalNetwork) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := utils.GetTwoSubnets(network)
	_, fundedKey := network.GetFundedAccountInfo()
	pChainInfo := utils.GetPChainInfo(cChainInfo)

	signatureAggregator := utils.NewSignatureAggregator(
		cChainInfo.NodeURIs[0],
		[]ids.ID{
			subnetAInfo.SubnetID,
			ids.Empty, // Primary network subnet ID
		},
	)

	// Generate random address to be the owner address
	ownerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	ownerAddress := crypto.PubkeyToAddress(ownerKey.PublicKey)

	// Transfer native assets to the owner account
	ctx := context.Background()
	fundAmount := big.NewInt(1e18) // 1avax
	utils.SendNativeTransfer(
		ctx,
		subnetAInfo,
		fundedKey,
		ownerAddress,
		fundAmount,
	)

	validatorManagerAddress, validatorManager := utils.DeployAndInitializePoAValidatorManager(
		ctx,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		ownerAddress,
	)

	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetAInfo.EVMChainID)
	Expect(err).Should(BeNil())

	convertSubnetTxId := ids.GenerateTestID()
	subnetConversionData := poavalidatormanager.SubnetConversionData{
		ConvertSubnetTxID:       convertSubnetTxId,
		BlockchainID:            subnetAInfo.BlockchainID,
		ValidatorManagerAddress: validatorManagerAddress.Bytes(),
		InitialValidators: []poavalidatormanager.InitialValidator{
			{
				NodeID:       ids.GenerateTestID(),
				Weight:       1,
				BlsPublicKey: []byte{},
			},
		},
	}
	unsignedMsg := createOffChainSubnetConversionMessage(subnetConversionData, pChainInfo, network.GetNetworkID())

	unsignedMsgs := []avalancheWarp.UnsignedMessage{*unsignedMsg}
	warpConfigWithMsg := utils.GetChainConfigWithOffChainMessages(unsignedMsgs)
	chainConfigs := make(utils.ChainConfigMap)
	chainConfigs.Add(subnetAInfo, warpConfigWithMsg)

	network.SetChainConfigs(chainConfigs)
	// network.RestartNodes(context.Background(), network.GetAllNodeIDs())

	_, err = validatorManager.InitializeValidatorSet(opts, subnetConversionData, 0)
	Expect(err).Should(BeNil())

	var validationID ids.ID // To be used in the delisting step
	nodeID := ids.GenerateTestID()
	blsPublicKey := [bls.PublicKeyLen]byte{}
	weight := uint64(1)

	{
		// Try to call with invalid owner
		opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetAInfo.EVMChainID)
		Expect(err).Should(BeNil())

		_, err = validatorManager.InitializeValidatorRegistration(
			opts,
			poavalidatormanager.ValidatorRegistrationInput{
				NodeID:             nodeID,
				RegistrationExpiry: uint64(time.Now().Add(24 * time.Hour).Unix()),
				BlsPublicKey:       blsPublicKey[:],
			},
			weight,
		)
		Expect(err).ShouldNot(BeNil())

		// Initiate validator registration
		validationID = utils.InitializeAndCompletePoAValidatorRegistration(
			network,
			signatureAggregator,
			ownerKey,
			fundedKey,
			subnetAInfo,
			pChainInfo,
			validatorManager,
			validatorManagerAddress,
			weight,
			nodeID,
			blsPublicKey,
		)
	}

	//
	// Delist the validator
	//
	{
		// Try with invalid validator owner
		opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetAInfo.EVMChainID)
		Expect(err).Should(BeNil())
		_, err = validatorManager.InitializeEndValidation(
			opts,
			validationID,
		)
		Expect(err).ShouldNot(BeNil())

		utils.InitializeAndCompleteEndPoAValidation(
			network,
			signatureAggregator,
			ownerKey,
			fundedKey,
			subnetAInfo,
			pChainInfo,
			validatorManager,
			validatorManagerAddress,
			validationID,
			weight,
			1,
		)
	}
}

func createOffChainSubnetConversionMessage(
	subnetConversionData poavalidatormanager.SubnetConversionData,
	pChainInfo interfaces.SubnetTestInfo,
	networkID uint32,
) *avalancheWarp.UnsignedMessage {
	subnetConversionBytes, err := subnetConversionData.Pack()
	Expect(err).Should(BeNil())
	subnetConversionID := sha256.Sum256(subnetConversionBytes)
	sourceAddress := []byte{}

	payloadBytes := make([]byte, 38)
	copy(payloadBytes[:6], subnetConversionID[:])
	addressedPayload, err := payload.NewAddressedCall(sourceAddress, payloadBytes)
	Expect(err).Should(BeNil())

	unsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		networkID,
		pChainInfo.BlockchainID,
		addressedPayload.Bytes(),
	)
	Expect(err).Should(BeNil())

	return unsignedMessage
}
