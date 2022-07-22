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
	classificationID := baseIDs.NewStringID("classificationID")
	immutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))

	testAssetID := baseIDs.NewAssetID(classificationID, immutableProperties).(key)

	require.NotPanics(t, func() {
		require.Equal(t, strings.Join([]string{classificationID.String(), baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID().String()}, "."), testAssetID.String())
		require.Equal(t, false, testAssetID.IsPartial())
		require.Equal(t, true, testAssetID.Equals(testAssetID))
		require.Equal(t, true, testAssetID.Equals(testAssetID))
		require.Equal(t, false, testAssetID.Equals(nil))
		require.Equal(t, testAssetID, FromID(testAssetID))
		require.Equal(t, testAssetID, readAssetID(testAssetID.String()))
	})
}
