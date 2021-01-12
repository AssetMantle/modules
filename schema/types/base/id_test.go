/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ID(t *testing.T) {
	testID := NewID("ID")

	require.Equal(t, id{IDString: "ID"}, testID)
	require.Equal(t, "ID", testID.String())
	require.Equal(t, true, testID.Equals(testID))
	require.Equal(t, false, testID.Equals(NewID("ID2")))
	require.Equal(t, []byte("ID"), testID.Bytes())
}
