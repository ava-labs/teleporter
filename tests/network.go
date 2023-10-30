package tests

import (
	"crypto/ecdsa"

	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
)

// Defines the interface for the network setup functions used in the E2E tests
type Network interface {
	GetSubnetsInfo() []utils.SubnetTestInfo
	GetTeleporterContractAddress() common.Address
	GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey)
}

// Implements Network, pointing to the network setup in network_setup.go
type ginkgoNetwork struct{}

func (g *ginkgoNetwork) GetSubnetsInfo() []utils.SubnetTestInfo {
	return utils.GetSubnetsInfo()
}

func (g *ginkgoNetwork) GetTeleporterContractAddress() common.Address {
	return utils.GetTeleporterContractAddress()
}

func (g *ginkgoNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	return utils.GetFundedAccountInfo()
}
