// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// Package unpack generates Solidity code to invert `abi.encodePacked()` for
// known, unambiguous output sizes.
package unpack

import (
	"errors"
	"fmt"
	"io"
	"math"
	"strings"
	"text/template"

	_ "embed"
)

var (
	//go:embed Unpack.sol.gotmpl
	solTmplRaw string
	//go:embed Unpack.t.sol.gotmpl
	solTestTmplRaw       string
	solTmpl, solTestTmpl *template.Template
)

func init() {
	solTmpl = template.Must(template.New("sol").Parse(solTmplRaw))
	solTestTmpl = template.Must(template.New("solTest").Parse(solTestTmplRaw))
}

// Errors returned by [Generate].
var (
	ErrZeroBytes       = errors.New("zero bytes")
	ErrMax32Bytes      = errors.New("max 32 bytes")
	ErrMultipleDynamic = errors.New("multiple unsized inputs")
)

// Dynamic is a sentinel value, used instead of a number of bytes, indicating
// that an unpacked type has dynamic length. At most one Dynamic value can be
// used per generated unpack function.
const Dynamic = math.MaxUint

// Generate generates implementations and tests for unpacking Solidity values
// encoded with `abi.encodePacked()`. Each slice of `[]uint` results in the
// generation of one `unpack_...()` function and MAY have at most one [Dynamic]
// size.
func Generate(impl, test io.Writer, relPathFromTestToImpl string, bytes ...[]uint) error {
	us := make([]*unpacker, len(bytes))
	for i, bs := range bytes {
		u, err := newUnpacker(bs...)
		if err != nil {
			return err
		}
		us[i] = u
	}

	data := &tmplData{
		RelPathFromTestToImpl: relPathFromTestToImpl,
		Unpackers:             us,
	}
	return errors.Join(
		solTmpl.Execute(impl, data),
		solTestTmpl.Execute(test, data),
	)
}

type tmplData struct {
	RelPathFromTestToImpl string
	Unpackers             []*unpacker
}

// An unpacker defines all values required to generate a single `unpack_...()`
// function.
type unpacker struct {
	Vars       []output
	HasDynamic bool
	UseMCopy   bool
}

// An output defines a single output returned by an `unpack_...()` function.
type output struct {
	end     end
	Offset  uint // relative to `end`
	Len     uint // non-zero if `!Dynamic`
	Dynamic bool
}

type end uint

const (
	front end = iota + 1
	back
)

func newUnpacker(bytes ...uint) (*unpacker, error) {
	u := &unpacker{
		Vars: make([]output, len(bytes)),
	}

	for i, b := range bytes {
		switch {
		case b == 0:
			return nil, ErrZeroBytes
		case b == Dynamic:
			if u.HasDynamic {
				return nil, ErrMultipleDynamic
			}
			u.HasDynamic = true
		case b > 32:
			return nil, ErrMax32Bytes
		}

		v := &u.Vars[i]
		switch {
		case b == Dynamic:
			v.end = front
			v.Offset = u.BytesUntilDynamic()
			v.Dynamic = true

		case !u.HasDynamic:
			v.end = front
			v.Offset = u.BytesUntilDynamic()

		case u.HasDynamic:
			v.end = back
		}
		v.Len = b
	}
	if !u.HasDynamic {
		return u, nil
	}

	// Once a dynamically sized output exists, we can no longer calculate
	// offsets from the front.
	var offset uint
	for i := len(u.Vars) - 1; i >= 0; i-- {
		v := &u.Vars[i]
		if v.Dynamic {
			break
		}
		offset += v.Len
		v.Offset = offset
	}
	return u, nil
}

// UsesMCopy returns whether at least one unpack function uses the MCOPY op
// code.
func (d *tmplData) UsesMCopy() bool {
	for _, u := range d.Unpackers {
		if u.UseMCopy {
			return true
		}
	}
	return false
}

// BytesUntilDynamic returns the total number of bytes before the first
// dynamically sized output, or until the end of the input buffer if all sizes
// are known in advance.
//
// While an `unpacker` is being constructed, this function returns only the
// already-parsed output sizes, which can be used for setting the offset of the
// next value.
func (u *unpacker) BytesUntilDynamic() uint {
	var l uint
	for _, v := range u.Vars {
		if v.Dynamic {
			break
		}
		l += v.Len
	}
	return l
}

// Bytes is an alias of [BytesUntilDynamic] to aid in readability of the
// template.
func (u *unpacker) Bytes() uint {
	return u.BytesUntilDynamic()
}

// MinLength returns the minimum number of bytes that an input buffer MUST have
// to be unpacked. This is the sum of all non-dynamically sized outputs.
func (u *unpacker) MinLength() uint {
	var n uint
	for _, v := range u.Vars {
		if !v.Dynamic {
			n += v.Len
		}
	}
	return n
}

// mapVar maps u.Vars -> []string.
func (u *unpacker) mapVars(fn func(int, output) string) []string {
	out := make([]string, len(u.Vars))
	for i, v := range u.Vars {
		out[i] = fn(i, v)
	}
	return out
}

// FuncNameSuffix returns the suffix appended to `unpack_` when naming the
// generated function. It concatenates all byte sizes, using the keyword "Dyn"
// for dynamically sized outputs.
//
// If the generated code doesn't use MCOPY it takes ownership of the input
// buffer in a destructive fashion. All such functions have the "_Destructive"
// suffix appended.
func (u *unpacker) FuncNameSuffix() string {
	base := strings.Join(
		u.mapVars(func(_ int, o output) string {
			if o.Dynamic {
				return "Dyn"
			}
			return fmt.Sprintf("%d", o.Len)
		}),
		"_",
	)
	if !u.HasDynamic || u.UseMCopy {
		return base
	}
	return fmt.Sprintf("%s_Destructive", base)
}

func (u *unpacker) IsLastVar(i int) bool {
	return i >= len(u.Vars)-1
}

// varPrefix is the prefix of all numerically identified variables in the
// generated code; i.e. for prefix v, the variables are v0, v1, v2...
const varPrefix = "v"

// EncodePackedArgs returns the arguments to be passed to `abi.encodePacked()`
// in the generated test.
func (u *unpacker) EncodePackedArgs() string {
	return strings.Join(
		u.mapVars(func(i int, _ output) string {
			return fmt.Sprintf("%s%d", varPrefix, i)
		}),
		", ",
	)
}

// Returns returns the return values of the generated unpack function.
func (u *unpacker) Returns() string {
	return u.typedVars(varPrefix)
}

// TestVars returns the variables under test in the generated test function.
func (u *unpacker) TestVars() string {
	return u.typedVars("got")
}

func (u *unpacker) typedVars(prefix string) string {
	return strings.Join(
		u.mapVars(func(i int, o output) string {
			return fmt.Sprintf("%s %s%d", o.Type(), prefix, i)
		}),
		", ",
	)
}

// Type returns the Solidity type of the output variable.
func (o *output) Type() string {
	if o.Dynamic {
		return "bytes memory"
	}
	return fmt.Sprintf("bytes%d", o.Len)
}

func (o *output) Front() bool {
	switch o.end {
	case front:
		return true
	case back:
		return false
	default:
		panic(fmt.Sprintf("BUG: %T.end not set", o))
	}
}

func (o *output) OffsetHex() string {
	off := o.Offset
	if o.end == front {
		off += 32
	}
	return fmt.Sprintf("%#02x", off)
}
