// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"testing"

	"github.com/AssetMantle/modules/constants/ids"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/schema/types/base"
)

func Test_Asset_Methods(t *testing.T) {
	classificationID := base.NewID("classificationID")
	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), baseData.NewStringData("ImmutableData")))
	mutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID2"), baseData.NewStringData("MutableData")))

	assetID := key.NewAssetID(classificationID, immutableProperties)
	testAsset := NewAsset(assetID, immutableProperties, mutableProperties).(asset)

	require.Equal(t, asset{Document: baseQualified.Document{ID: assetID, HasImmutables: baseQualified.HasImmutables{Properties: immutableProperties}, HasMutables: baseQualified.HasMutables{Properties: mutableProperties}}}, testAsset)
	require.Equal(t, assetID, testAsset.GetID())
	require.Equal(t, classificationID, testAsset.GetClassificationID())
	require.Equal(t, immutableProperties, testAsset.GetImmutableProperties())
	require.Equal(t, mutableProperties, testAsset.GetMutableProperties())
	data, _ := baseData.ReadHeightData("-1")
	require.Equal(t, testAsset.GetBurn(), base.NewProperty(ids.BurnProperty, data))
	require.Equal(t, base.NewProperty(ids.BurnProperty, baseData.NewStringData("BurnImmutableData")), asset{baseQualified.Document{ID: assetID, HasImmutables: baseQualified.HasImmutables{Properties: base.NewProperties(base.NewProperty(ids.BurnProperty, baseData.NewStringData("BurnImmutableData")))}, HasMutables: baseQualified.HasMutables{Properties: mutableProperties}}}.GetBurn())
	require.Equal(t, base.NewProperty(ids.BurnProperty, baseData.NewStringData("BurnMutableData")), asset{baseQualified.Document{ID: assetID, HasImmutables: baseQualified.HasImmutables{Properties: immutableProperties}, HasMutables: baseQualified.HasMutables{Properties: base.NewProperties(base.NewProperty(ids.BurnProperty, baseData.NewStringData("BurnMutableData")))}}}.GetBurn())
	require.Equal(t, base.NewProperty(ids.LockProperty, data), testAsset.GetLock())
	require.Equal(t, base.NewProperty(ids.LockProperty, baseData.NewStringData("LockImmutableData")), asset{baseQualified.Document{ID: assetID, HasImmutables: baseQualified.HasImmutables{Properties: base.NewProperties(base.NewProperty(ids.LockProperty, baseData.NewStringData("LockImmutableData")))}, HasMutables: baseQualified.HasMutables{Properties: mutableProperties}}}.GetLock())
	require.Equal(t, base.NewProperty(ids.LockProperty, baseData.NewStringData("LockMutableData")), asset{baseQualified.Document{ID: assetID, HasImmutables: baseQualified.HasImmutables{Properties: immutableProperties}, HasMutables: baseQualified.HasMutables{Properties: base.NewProperties(base.NewProperty(ids.LockProperty, baseData.NewStringData("LockMutableData")))}}}.GetLock())
	require.Equal(t, assetID, testAsset.GetKey())

}
