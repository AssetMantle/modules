// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/assets/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

func Test_Asset_Methods(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	immutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("ImmutableData")))
	mutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewID("ID2"), baseData.NewStringData("MutableData")))

	assetID := key.NewAssetID(classificationID, immutableProperties)
	testAsset := NewAsset(assetID, immutableProperties, mutableProperties).(asset)

	require.Equal(t, asset{Document: baseQualified.Document{ID: assetID, Immutables: baseQualified.Immutables{PropertyList: immutableProperties}, Mutables: baseQualified.Mutables{Properties: mutableProperties}}}, testAsset)
	require.Equal(t, assetID, testAsset.GetID())
	require.Equal(t, classificationID, testAsset.GetClassificationID())
	require.Equal(t, immutableProperties, testAsset.GetImmutablePropertyList())
	require.Equal(t, mutableProperties, testAsset.GetMutablePropertyList())
	require.Equal(t, baseProperties.NewProperty(constants.BurnProperty, baseData.NewStringData("BurnImmutableData")), asset{baseQualified.Document{ID: assetID, Immutables: baseQualified.Immutables{PropertyList: base.NewPropertyList(baseProperties.NewProperty(constants.BurnProperty, baseData.NewStringData("BurnImmutableData")))}, Mutables: baseQualified.Mutables{Properties: mutableProperties}}}.GetBurn())
	require.Equal(t, baseProperties.NewProperty(constants.BurnProperty, baseData.NewStringData("BurnMutableData")), asset{baseQualified.Document{ID: assetID, Immutables: baseQualified.Immutables{PropertyList: immutableProperties}, Mutables: baseQualified.Mutables{Properties: base.NewPropertyList(baseProperties.NewProperty(constants.BurnProperty, baseData.NewStringData("BurnMutableData")))}}}.GetBurn())
	require.Equal(t, baseProperties.NewProperty(constants.LockProperty, baseData.NewStringData("LockImmutableData")), asset{baseQualified.Document{ID: assetID, Immutables: baseQualified.Immutables{PropertyList: base.NewPropertyList(baseProperties.NewProperty(constants.LockProperty, baseData.NewStringData("LockImmutableData")))}, Mutables: baseQualified.Mutables{Properties: mutableProperties}}}.GetLock())
	require.Equal(t, baseProperties.NewProperty(constants.LockProperty, baseData.NewStringData("LockMutableData")), asset{baseQualified.Document{ID: assetID, Immutables: baseQualified.Immutables{PropertyList: immutableProperties}, Mutables: baseQualified.Mutables{Properties: base.NewPropertyList(baseProperties.NewProperty(constants.LockProperty, baseData.NewStringData("LockMutableData")))}}}.GetLock())
	require.Equal(t, assetID, testAsset.GetKey())

}
