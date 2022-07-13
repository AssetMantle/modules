// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

func Test_MetaID_Methods(t *testing.T) {
	typeID := baseIDs.NewStringID("I")
	hashID := baseIDs.NewStringID("hashID")
	testMetaID := NewMetaID(typeID, hashID).(metaID)

	require.NotPanics(t, func() {
		require.Equal(t, typeID.String()+"."+hashID.String(), testMetaID.String())
		require.Equal(t, true, testMetaID.Equals(testMetaID))
		require.Equal(t, false, testMetaID.Equals(metaID{Type: baseIDs.NewStringID("tempID"), Hash: baseIDs.NewStringID("tempHash")}))
		require.Equal(t, false, testMetaID.IsPartial())
		require.Equal(t, true, metaID{Hash: baseIDs.NewStringID("")}.IsPartial())
		require.Equal(t, true, testMetaID.Equals(testMetaID))
		require.Equal(t, false, testMetaID.Equals(nil))
		require.Equal(t, testMetaID, FromID(testMetaID))
		require.Equal(t, NewMetaID(baseIDs.NewStringID(""), baseIDs.NewStringID("")), FromID(baseIDs.NewStringID("")))
	})
}
