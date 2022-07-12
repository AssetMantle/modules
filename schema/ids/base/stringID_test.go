// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ID(t *testing.T) {
	testID := NewStringID("ID")

	require.Equal(t, stringID{IDString: "ID"}, testID)
	require.Equal(t, "ID", testID.String())
	require.Equal(t, true, testID.Compare(testID) == 0)
	require.Equal(t, false, testID.Compare(NewStringID("ID2")) == 0)
	require.Equal(t, []byte("ID"), testID.Bytes())
}
