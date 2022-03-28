/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ID(t *testing.T) {
	testID := NewID("ID")

	require.Equal(t, ID{IdString: "ID"}, testID)
	require.Equal(t, "ID", testID.String())
	require.Equal(t, true, testID.Compare(testID) == 0)
	require.Equal(t, false, testID.Compare(NewID("ID2")) == 0)
	require.Equal(t, []byte("ID"), testID.Bytes())
}
