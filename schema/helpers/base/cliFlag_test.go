// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func Test_CliFlag(t *testing.T) {
	testCliFlag := NewCLIFlag("name", "value", ",usage")
	require.Panics(t, func() {
		NewCLIFlag("name", struct{}{}, ",usage").Register(&cobra.Command{})
	})
	// GetName method test
	require.Equal(t, "name", testCliFlag.GetName())
	// GetSupply method test
	require.Equal(t, "value", testCliFlag.GetValue())

	// ReadCLIValue method test
	require.Equal(t, "", testCliFlag.ReadCLIValue())
	require.Equal(t, int64(0), NewCLIFlag("name", int64(-1), ",usage").ReadCLIValue())
	require.Equal(t, 0, NewCLIFlag("name", 123, ",usage").ReadCLIValue())
	require.Equal(t, false, NewCLIFlag("name", false, ",usage").ReadCLIValue())
	require.Panics(t, func() {
		NewCLIFlag("name", struct{}{}, ",usage").ReadCLIValue()
	})

}
