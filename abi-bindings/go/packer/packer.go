// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package packer

// Interface to be implemented manually by the user for all types that need to be packed
type ABIPacker interface {
	Pack() ([]byte, error)
	Unpack([]byte) error
}
