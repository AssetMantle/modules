// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

func Test_ClassificationID_Methods(t *testing.T) {
	immutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	mutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData")))

	testClassificationID := NewClassificationID(immutableProperties, mutableProperties).(classificationID)
	require.NotPanics(t, func() {
		require.Equal(t, classificationID{Hash: baseIDs.NewStringID(stringUtilities.Hash(stringUtilities.Hash("ID1"), stringUtilities.Hash("ID2"), baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID().String()))}, testClassificationID)
		require.Equal(t, false, testClassificationID.Equals(nil))
		require.Equal(t, false, testClassificationID.Compare(baseIDs.NewStringID("id")) == 0)
		require.Equal(t, true, testClassificationID.Equals(testClassificationID))
		require.Equal(t, false, testClassificationID.IsPartial())
		require.Equal(t, testClassificationID, FromID(testClassificationID))
	})
}
