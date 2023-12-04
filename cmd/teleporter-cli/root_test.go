package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func executeTestCmd(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()
	return strings.TrimSpace(buf.String()), err
}

func TestRootCmd(t *testing.T) {
	var tests = []struct {
		name string
		args []string
		err  error
		out  string
	}{
		{
			name: "base",
			args: []string{},
			err:  nil,
			out:  "A CLI that integrates with the Teleporter protocol",
		},
		{
			name: "help",
			args: []string{"--help"},
			err:  nil,
			out:  "A CLI that integrates with the Teleporter protocol",
		},
		{
			name: "invalid",
			args: []string{"invalid"},
			err:  fmt.Errorf("unknown command"),
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
