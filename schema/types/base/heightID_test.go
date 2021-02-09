/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HeightID(t *testing.T) {
	testID := NewHeightID(100)

	require.Equal(t, heightID{Value: 100}, testID)
	require.Equal(t, "100", testID.String())
	require.Equal(t, true, testID.Equals(testID))
	require.Equal(t, false, testID.Equals(NewID("1")))
	require.Equal(t, append([]byte{uint8(3)}, []byte("100")...), testID.Bytes())
}
