package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEventCmd(t *testing.T) {
	var tests = []struct {
		name string
		args []string
		err  error
		out  string
	}{
		{
			name: "no args",
			args: []string{"event"},
			err:  fmt.Errorf("required flag(s) \"topics\" not set"),
		},
		{
			name: "help",
			args: []string{"event", "--help"},
			err:  nil,
			out:  "Given the topics and data of a Teleporter log",
		},
		{
			name: "topics",
			args: []string{"event", "--topics", "0x6b013241f9192863bc66c1f1e9a01dc592c94592bfed5e1ed380808525679575,0x39bec71f2a981a5d5dc824471788df002194dee5db051c7902b51960c466ca81,0x0000000000000000000000000000000000000000000000000000000000000251,0x00000000000000000000000066fc895081cf815aa25bf4853cc377dfc8cc5871"},
			err:  nil,
			out:  "Event command ran successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := execute(t, rootCmd, tt.args...)
			if tt.err != nil {
				require.ErrorContains(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
				require.Contains(t, out, tt.out)
			}
		})
	}
}

func TestEventTopics(t *testing.T) {
	var tests = []struct {
		name string
		args []string
		err  error
		out  string
	}{
		{
			name: "topics",
			args: []string{"event", "--topics", "0x6b013241f9192863bc66c1f1e9a01dc592c94592bfed5e1ed380808525679575,0x39bec71f2a981a5d5dc824471788df002194dee5db051c7902b51960c466ca81,0x0000000000000000000000000000000000000000000000000000000000000251,0x00000000000000000000000066fc895081cf815aa25bf4853cc377dfc8cc5871"},
			err:  nil,
			out:  "Event command ran successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := execute(t, rootCmd, tt.args...)
			if tt.err != nil {
				require.ErrorContains(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
				require.Contains(t, out, tt.out)
			}
		})
	}
}
