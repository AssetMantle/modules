// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/types/base"
)

func TestFromID(t *testing.T) {
	classificationID := base.NewID("classificationID")
	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), baseData.NewStringData("ImmutableData")))
	newAssetID := NewAssetID(classificationID, immutableProperties)
	require.Equal(t, assetIDFromInterface(newAssetID), FromID(newAssetID))

	id := base.NewID("")
	testAssetID := assetID{ClassificationID: base.NewID(""), HashID: base.NewID("")}
	require.Equal(t, FromID(id), testAssetID)

	testString1 := "string1"
	testString2 := "string2"
	id2 := base.NewID(testString1 + constants.FirstOrderCompositeIDSeparator + testString2)
	testAssetID2 := assetID{ClassificationID: base.NewID(testString1), HashID: base.NewID(testString2)}
	require.Equal(t, FromID(id2), testAssetID2)
}

func TestReadClassificationID(t *testing.T) {
	classificationID := base.NewID("classificationID")
	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), baseData.NewStringData("ImmutableData")))
	assetID := NewAssetID(classificationID, immutableProperties)

	require.Equal(t, assetIDFromInterface(assetID).ClassificationID, ReadClassificationID(assetID))
}
