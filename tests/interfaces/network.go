package interfaces

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

// Defines the interface for the network setup functions used in the E2E tests
type Network interface {
	// Returns information about the primary network
	GetPrimaryNetworkInfo() SubnetTestInfo

	// Returns all of the subnets support by this network, excluding the primary network
	GetSubnetsInfo() []SubnetTestInfo

	// Returns all of the subnets support by this network in no particular order, including the primary network.
	// If the distinction between primary network and other subnets needs to be made,
	// use GetPrimaryNetworkInfo() and GetSubnetsInfo() instead.
	GetAllSubnetsInfo() []SubnetTestInfo

	// An address and corresponding key that has native tokens on each of the subnets in this network.
	GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey)

	// Whether or not the Avalanche network being used is available outside the scope of the test application.
	IsExternalNetwork() bool

	// Whether or not the funded wallet is capable of relaying messages between subnets in this network.
	// Intended to be true for local networks where all nodes are querable by the test application for their
	// BLS signatures, and false for testnet networks where test application does not necessarily have
	// connections with each validator.
	SupportsIndependentRelaying() bool
}
