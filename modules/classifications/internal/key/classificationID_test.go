// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	baseTraits "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	metaUtilities "github.com/AssetMantle/modules/utilities/meta"
)

func Test_ClassificationID_Methods(t *testing.T) {
	chainID := baseIDs.NewID("chainID")
	immutableProperties := base.NewPropertyList(baseTypes.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("ImmutableData")))
	mutableProperties := base.NewPropertyList(baseTypes.NewProperty(baseIDs.NewID("ID2"), baseData.NewStringData("MutableData")))

	testClassificationID := NewClassificationID(chainID, immutableProperties, mutableProperties).(classificationID)
	require.NotPanics(t, func() {
		require.Equal(t, classificationID{ChainID: chainID, HashID: baseIDs.NewID(metaUtilities.Hash(metaUtilities.Hash("ID1"), metaUtilities.Hash("ID2"), baseTraits.Immutables{Properties: immutableProperties}.GenerateHashID().String()))}, testClassificationID)
		require.Equal(t, strings.Join([]string{chainID.String(), baseIDs.NewID(metaUtilities.Hash(metaUtilities.Hash("ID1"), metaUtilities.Hash("ID2"), baseTraits.Immutables{Properties: immutableProperties}.GenerateHashID().String())).String()}, constants.IDSeparator), testClassificationID.String())
		require.Equal(t, false, testClassificationID.Equals(classificationID{ChainID: baseIDs.NewID("chainID"), HashID: baseIDs.NewID("hashID")}))
		require.Equal(t, false, testClassificationID.Equals(nil))
		require.Equal(t, false, testClassificationID.Compare(baseIDs.NewID("id")) == 0)
		require.Equal(t, true, testClassificationID.Equals(testClassificationID))
		require.Equal(t, false, testClassificationID.IsPartial())
		require.Equal(t, true, classificationID{ChainID: chainID, HashID: baseIDs.NewID("")}.IsPartial())
		require.Equal(t, testClassificationID, FromID(testClassificationID))
		require.Equal(t, classificationID{ChainID: baseIDs.NewID(""), HashID: baseIDs.NewID("")}, FromID(baseIDs.NewID("tempID")))
		require.Equal(t, testClassificationID, readClassificationID(testClassificationID.String()))
	})

}
