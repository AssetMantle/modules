// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/modules/assets/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func Test_Asset_Methods(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	immutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("ImmutableData")))
	mutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID2"), baseData.NewStringData("MutableData")))

	assetID := key.NewAssetID(classificationID, immutableProperties)
	testAsset := NewAsset(assetID, immutableProperties, mutableProperties).(asset)

	require.Equal(t, asset{Document: baseQualified.Document{ID: assetID, Immutables: baseQualified.Immutables{Properties: immutableProperties}, Mutables: baseQualified.Mutables{Properties: mutableProperties}}}, testAsset)
	require.Equal(t, assetID, testAsset.GetID())
	require.Equal(t, classificationID, testAsset.GetClassificationID())
	require.Equal(t, immutableProperties, testAsset.GetImmutableProperties())
	require.Equal(t, mutableProperties, testAsset.GetMutableProperties())
	data, _ := baseData.ReadHeightData("-1")
	require.Equal(t, testAsset.GetBurn(), baseTypes.NewProperty(ids.BurnProperty, data))
	require.Equal(t, baseTypes.NewProperty(ids.BurnProperty, baseData.NewStringData("BurnImmutableData")), asset{baseQualified.Document{ID: assetID, Immutables: baseQualified.Immutables{Properties: baseTypes.NewProperties(baseTypes.NewProperty(ids.BurnProperty, baseData.NewStringData("BurnImmutableData")))}, Mutables: baseQualified.Mutables{Properties: mutableProperties}}}.GetBurn())
	require.Equal(t, baseTypes.NewProperty(ids.BurnProperty, baseData.NewStringData("BurnMutableData")), asset{baseQualified.Document{ID: assetID, Immutables: baseQualified.Immutables{Properties: immutableProperties}, Mutables: baseQualified.Mutables{Properties: baseTypes.NewProperties(baseTypes.NewProperty(ids.BurnProperty, baseData.NewStringData("BurnMutableData")))}}}.GetBurn())
	require.Equal(t, baseTypes.NewProperty(ids.LockProperty, data), testAsset.GetLock())
	require.Equal(t, baseTypes.NewProperty(ids.LockProperty, baseData.NewStringData("LockImmutableData")), asset{baseQualified.Document{ID: assetID, Immutables: baseQualified.Immutables{Properties: baseTypes.NewProperties(baseTypes.NewProperty(ids.LockProperty, baseData.NewStringData("LockImmutableData")))}, Mutables: baseQualified.Mutables{Properties: mutableProperties}}}.GetLock())
	require.Equal(t, baseTypes.NewProperty(ids.LockProperty, baseData.NewStringData("LockMutableData")), asset{baseQualified.Document{ID: assetID, Immutables: baseQualified.Immutables{Properties: immutableProperties}, Mutables: baseQualified.Mutables{Properties: baseTypes.NewProperties(baseTypes.NewProperty(ids.LockProperty, baseData.NewStringData("LockMutableData")))}}}.GetLock())
	require.Equal(t, assetID, testAsset.GetKey())

}
