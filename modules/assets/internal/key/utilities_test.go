// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func TestFromID(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	immutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("ImmutableData")))
	newAssetID := NewAssetID(classificationID, immutableProperties)
	require.Equal(t, assetIDFromInterface(newAssetID), FromID(newAssetID))

	id := baseIDs.NewID("")
	testAssetID := assetID{Classification: baseIDs.NewID(""), Hash: baseIDs.NewID("")}
	require.Equal(t, FromID(id), testAssetID)

	testString1 := "string1"
	testString2 := "string2"
	id2 := baseIDs.NewID(testString1 + constants.FirstOrderCompositeIDSeparator + testString2)
	testAssetID2 := assetID{Classification: baseIDs.NewID(testString1), Hash: baseIDs.NewID(testString2)}
	require.Equal(t, FromID(id2), testAssetID2)
}

func TestReadClassificationID(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	immutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("ImmutableData")))
	assetID := NewAssetID(classificationID, immutableProperties)

	require.Equal(t, assetIDFromInterface(assetID).Classification, ReadClassificationID(assetID))
}
