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
			name: "base",
			args: []string{"event"},
			err:  fmt.Errorf("required flag(s) \"topics\" not set"),
		},
		{
			name: "help",
			args: []string{"event", "--help"},
			err:  nil,
			out:  "Given the topics and data of a Teleporter log",
		},
		// {
		// 	name: "invalid",
		// 	args: []string{"invalid"},
		// 	err:  fmt.Errorf("unknown command"),
		// },
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
