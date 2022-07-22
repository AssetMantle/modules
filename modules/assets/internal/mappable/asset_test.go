// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

func Test_Asset_Methods(t *testing.T) {
	classificationID := baseIDs.NewStringID("classificationID")
	immutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	mutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData")))

	assetID := baseIDs.NewAssetID(classificationID, immutableProperties)
	testAsset := NewAsset(assetID, immutableProperties, mutableProperties).(asset)

	require.Equal(t, asset{Document: baseQualified.Document{ID: assetID, Immutables: baseQualified.Immutables{PropertyList: immutableProperties}, Mutables: baseQualified.Mutables{PropertyList: mutableProperties}}}, testAsset)
	require.Equal(t, assetID, testAsset.GetID())
	require.Equal(t, classificationID, testAsset.GetClassificationID())
	require.Equal(t, immutableProperties, testAsset.GetImmutablePropertyList())
	require.Equal(t, mutableProperties, testAsset.GetMutablePropertyList())
	require.Equal(t, baseProperties.NewProperty(constants.BurnProperty, baseData.NewStringData("BurnImmutableData")), asset{baseQualified.Document{ID: assetID, Immutables: baseQualified.Immutables{PropertyList: base.NewPropertyList(baseProperties.NewProperty(constants.BurnProperty, baseData.NewStringData("BurnImmutableData")))}, Mutables: baseQualified.Mutables{PropertyList: mutableProperties}}}.GetBurn())
	require.Equal(t, baseProperties.NewProperty(constants.BurnProperty, baseData.NewStringData("BurnMutableData")), asset{baseQualified.Document{ID: assetID, Immutables: baseQualified.Immutables{PropertyList: immutableProperties}, Mutables: baseQualified.Mutables{PropertyList: base.NewPropertyList(baseProperties.NewProperty(constants.BurnProperty, baseData.NewStringData("BurnMutableData")))}}}.GetBurn())
	require.Equal(t, baseProperties.NewProperty(constants.LockProperty, baseData.NewStringData("LockImmutableData")), asset{baseQualified.Document{ID: assetID, Immutables: baseQualified.Immutables{PropertyList: base.NewPropertyList(baseProperties.NewProperty(constants.LockProperty, baseData.NewStringData("LockImmutableData")))}, Mutables: baseQualified.Mutables{PropertyList: mutableProperties}}}.GetLock())
	require.Equal(t, baseProperties.NewProperty(constants.LockProperty, baseData.NewStringData("LockMutableData")), asset{baseQualified.Document{ID: assetID, Immutables: baseQualified.Immutables{PropertyList: immutableProperties}, Mutables: baseQualified.Mutables{PropertyList: base.NewPropertyList(baseProperties.NewProperty(constants.LockProperty, baseData.NewStringData("LockMutableData")))}}}.GetLock())
	require.Equal(t, assetID, testAsset.GetKey())

}
