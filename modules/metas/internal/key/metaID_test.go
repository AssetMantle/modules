// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/schema/types/base"
)

func Test_MetaID_Methods(t *testing.T) {
	typeID := base.NewID("I")
	hashID := base.NewID("hashID")
	testMetaID := NewMetaID(typeID, hashID).(metaID)

	require.NotPanics(t, func() {
		require.Equal(t, typeID.String()+constants.FirstOrderCompositeIDSeparator+hashID.String(), testMetaID.String())
		require.Equal(t, true, testMetaID.Equals(testMetaID))
		require.Equal(t, false, testMetaID.Equals(metaID{TypeID: base.NewID("tempID"), HashID: base.NewID("tempHash")}))
		require.Equal(t, false, testMetaID.IsPartial())
		require.Equal(t, true, metaID{HashID: base.NewID("")}.IsPartial())
		require.Equal(t, true, testMetaID.Equals(testMetaID))
		require.Equal(t, false, testMetaID.Equals(nil))
		require.Equal(t, testMetaID, FromID(testMetaID))
		require.Equal(t, NewMetaID(base.NewID(""), base.NewID("")), FromID(base.NewID("")))
	})
}
