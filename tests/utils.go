// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"os"
	"strings"

	. "github.com/onsi/gomega"
)

func ReadHexTextFile(filename string) string {
	fileData, err := os.ReadFile(filename)
	Expect(err).Should(BeNil())
	return strings.TrimRight(string(fileData), "\n")
}

// SanitizeHexString removes the "0x" prefix from a hex string if it exists.
// Otherwise, returns the original string.
func SanitizeHexString(hex string) string {
	return strings.TrimPrefix(hex, "0x")
}
