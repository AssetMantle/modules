/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"strings"
	"testing"

	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/qualified"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
)

func Test_ClassificationID_Methods(t *testing.T) {
	chainID := base.NewID("chainID")
	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewStringData("ImmutableData")))
	mutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID2"), base.NewStringData("MutableData")))

	testClassificationID := NewClassificationID(chainID, immutableProperties, mutableProperties).(classificationID)
	require.NotPanics(t, func() {
		require.Equal(t, classificationID{ChainID: chainID, HashID: base.NewID(metaUtilities.Hash(metaUtilities.Hash("ID1"), metaUtilities.Hash("ID2"), baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID().String()))}, testClassificationID)
		require.Equal(t, strings.Join([]string{chainID.String(), base.NewID(metaUtilities.Hash(metaUtilities.Hash("ID1"), metaUtilities.Hash("ID2"), baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID().String())).String()}, constants.IDSeparator), testClassificationID.String())
		require.Equal(t, false, testClassificationID.Equals(classificationID{ChainID: base.NewID("chainID"), HashID: base.NewID("hashID")}))
		require.Equal(t, false, testClassificationID.Equals(nil))
		require.Equal(t, false, testClassificationID.Compare(base.NewID("id")) == 0)
		require.Equal(t, true, testClassificationID.Equals(testClassificationID))
		require.Equal(t, false, testClassificationID.IsPartial())
		require.Equal(t, true, classificationID{ChainID: chainID, HashID: base.NewID("")}.IsPartial())
		require.Equal(t, testClassificationID, FromID(testClassificationID))
		require.Equal(t, classificationID{ChainID: base.NewID(""), HashID: base.NewID("")}, FromID(base.NewID("tempID")))
		require.Equal(t, testClassificationID, readClassificationID(testClassificationID.String()))
	})

}
