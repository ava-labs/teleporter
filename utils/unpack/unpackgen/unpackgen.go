// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// The unpackgen binary generates Solidity code to invert `abi.encodePacked()`
// for known, unambiguous output sizes. See the [unpack] package for details.
package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/pflag"

	"github.com/ava-labs/teleporter/utils/unpack"
)

func main() {
	implPath := pflag.String("out", "./Unpack.sol", "Output path to which the Solidity implementations are written")
	testPath := pflag.String("test_out", "./Unpack.t.sol", "Output path to which the Solidity tests are written")
	bytes := new(bytesFlag)
	pflag.Var(bytes, "byte_sizes", "Semicolon-delimited list of groups of comma-delimited byte sizes. Use 0 for (at most one) dynamically sized byte array per group.")

	pflag.Parse()

	if err := run(*implPath, *testPath, bytes.val); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run creates files at both paths and then calls `unpack.Generate()`.
func run(implPath, testPath string, bytes [][]uint) error {
	impl, err := os.Create(implPath)
	if err != nil {
		return err
	}
	test, err := os.Create(testPath)
	if err != nil {
		return err
	}

	rel, err := filepath.Rel(filepath.Dir(testPath), implPath)
	if err != nil {
		return fmt.Errorf("determine path from test to implementation: %v", err)
	}

	if err := unpack.Generate(impl, test, rel, bytes...); err != nil {
		return err
	}
	return errors.Join(impl.Close(), test.Close())
}

type bytesFlag struct {
	val [][]uint
}

var _ pflag.Value = (*bytesFlag)(nil)

// Set splits `s` by semicolons then splits each result by commas, parsing all
// resulting values as uints. It trims whitespace. If a parsed value is 0 it is
// replaced with unpack.Dynamic.
func (f *bytesFlag) Set(s string) error {
	groups := strings.Split(s, ";")
	var val [][]uint

	for _, g := range groups {
		var bytes []uint
		for _, bStr := range strings.Split(g, ",") {
			b, err := strconv.ParseUint(strings.TrimSpace(bStr), 10, 32)
			if err != nil {
				return err
			}
			if b == 0 {
				b = unpack.Dynamic
			}
			bytes = append(bytes, uint(b))
		}
		val = append(val, bytes)
	}

	f.val = val
	return nil
}

// String is the inverse of `bytesFlag.Set()`.
func (f *bytesFlag) String() string {
	var parts []string
	for _, grp := range f.val {
		var vs []string
		for _, v := range grp {
			if v == unpack.Dynamic {
				v = 0
			}
			vs = append(vs, fmt.Sprintf("%d", v))
		}
		parts = append(parts, strings.Join(vs, ","))
	}
	return strings.Join(parts, ";")
}

func (f *bytesFlag) Type() string {
	return fmt.Sprintf("%T", *f)
}
