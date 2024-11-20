package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMessageCmd(t *testing.T) {
	var tests = []struct {
		name string
		args []string
		err  error
		out  string
	}{
		{
			name: "no args",
			args: []string{"message"},
			err:  fmt.Errorf("accepts 1 arg(s), received 0"),
		},
		{
			name: "help",
			args: []string{"message", "--help"},
			err:  nil,
			out:  "Given the hex encoded bytes of a TeleporterMessenger message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := executeTestCmd(t, rootCmd, tt.args...)
			if tt.err != nil {
				require.ErrorContains(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
				require.Contains(t, out, tt.out)
			}
		})
	}
}
