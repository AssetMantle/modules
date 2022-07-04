// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

func Test_AssetID_Methods(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	immutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("ImmutableData")))

	testAssetID := NewAssetID(classificationID, immutableProperties).(assetID)

	require.NotPanics(t, func() {
		require.Equal(t, assetID{ClassificationID: classificationID, HashID: baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()}, testAssetID)
		require.Equal(t, strings.Join([]string{classificationID.String(), baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID().String()}, "."), testAssetID.String())
		require.Equal(t, false, testAssetID.IsPartial())
		require.Equal(t, true, assetID{ClassificationID: classificationID, HashID: baseIDs.NewID("")}.IsPartial())
		require.Equal(t, true, testAssetID.Equals(testAssetID))
		require.Equal(t, false, testAssetID.Equals(assetID{ClassificationID: classificationID, HashID: baseIDs.NewID("")}))
		require.Equal(t, true, testAssetID.Equals(testAssetID))
		require.Equal(t, false, testAssetID.Equals(nil))
		require.Equal(t, false, testAssetID.Equals(assetID{ClassificationID: classificationID, HashID: baseIDs.NewID("")}))
		require.Equal(t, testAssetID, FromID(testAssetID))
		require.Equal(t, assetID{ClassificationID: baseIDs.NewID(""), HashID: baseIDs.NewID("")}, FromID(baseIDs.NewID("")))
		require.Equal(t, testAssetID, readAssetID(testAssetID.String()))
	})
}
