package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransactionCmd(t *testing.T) {
	var tests = []struct {
		name string
		args []string
		err  error
		out  string
	}{
		{
			name: "no args",
			args: []string{"transaction"},
			err:  fmt.Errorf("accepts 1 arg(s), received 0"),
		},
		{
			name: "help",
			args: []string{"transaction", "--help"},
			err:  nil,
			out:  "Given a transaction this command looks through the transaction's receipt",
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
