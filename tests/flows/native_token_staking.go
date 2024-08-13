package flows

import (
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/awm-relayer/peers"
	"github.com/ava-labs/awm-relayer/signature-aggregator/aggregator"
	sigAggConfig "github.com/ava-labs/awm-relayer/signature-aggregator/config"
	sigAggMetrics "github.com/ava-labs/awm-relayer/signature-aggregator/metrics"
	"github.com/ava-labs/teleporter/tests/interfaces"
	. "github.com/onsi/gomega"
	"github.com/prometheus/client_golang/prometheus"
)

func NativeTokenStakingManager(network interfaces.Network) {
	// Get the subnets info

	// Create the signature aggregator
	logger := logging.NoLog{}
	cfg := sigAggConfig.Config{}
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
	signatureAggregator := aggregator.NewSignatureAggregator(
		appRequestNetwork,
		logger,
		sigAggMetrics.NewSignatureAggregatorMetrics(
			prometheus.DefaultRegisterer,
		),
		messageCreator,
	)

	// Deploy the staking manager contract

	// Iniatiate validator registration

	// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
	// (Sending to the P-Chain will be skipped for now)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	// Query P-Chain validators for the Warp message

	// Deliver the Warp message to the subnet

	// Check that the validator is registered in the staking contract

}
