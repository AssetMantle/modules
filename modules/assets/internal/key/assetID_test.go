// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/constants"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/qualified/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func Test_AssetID_Methods(t *testing.T) {
	classificationID := base.NewID("classificationID")
	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewStringData("ImmutableData")))

	testAssetID := NewAssetID(classificationID, immutableProperties).(assetID)

	require.NotPanics(t, func() {
		require.Equal(t, assetID{ClassificationID: classificationID, HashID: baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID()}, testAssetID)
		require.Equal(t, strings.Join([]string{classificationID.String(), baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID().String()}, constants.FirstOrderCompositeIDSeparator), testAssetID.String())
		require.Equal(t, false, testAssetID.IsPartial())
		require.Equal(t, true, assetID{ClassificationID: classificationID, HashID: base.NewID("")}.IsPartial())
		require.Equal(t, true, testAssetID.Equals(testAssetID))
		require.Equal(t, false, testAssetID.Equals(assetID{ClassificationID: classificationID, HashID: base.NewID("")}))
		require.Equal(t, true, testAssetID.Equals(testAssetID))
		require.Equal(t, false, testAssetID.Equals(nil))
		require.Equal(t, false, testAssetID.Equals(assetID{ClassificationID: classificationID, HashID: base.NewID("")}))
		require.Equal(t, testAssetID, FromID(testAssetID))
		require.Equal(t, assetID{ClassificationID: base.NewID(""), HashID: base.NewID("")}, FromID(base.NewID("")))
		require.Equal(t, testAssetID, readAssetID(testAssetID.String()))
	})
}
