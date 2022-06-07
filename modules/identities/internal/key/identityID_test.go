// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

func Test_IdentityID_Methods(t *testing.T) {

	classificationID := baseIDs.NewID("classificationID")
	immutableProperties, _ := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	emptyImmutableProperties, _ := utilities.ReadProperties("")
	testIdentityID := NewIdentityID(classificationID, immutableProperties)
	testIdentityID2 := NewIdentityID(classificationID, emptyImmutableProperties)
	key := FromID(testIdentityID)

	require.NotPanics(t, func() {
		require.Equal(t, testIdentityID, identityID{ClassificationID: classificationID, HashID: baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()})
		require.Equal(t, FromID(testIdentityID), identityIDFromInterface(testIdentityID))
		require.Equal(t, false, testIdentityID.(identityID).IsPartial())
		require.Equal(t, true, testIdentityID2.(identityID).IsPartial())
		require.Equal(t, false, testIdentityID2.(identityID).Compare(testIdentityID) == 0)
		require.Equal(t, true, testIdentityID.(identityID).Equals(key))
		require.Equal(t, false, testIdentityID.(identityID).Equals(FromID(baseIDs.NewID("id"))))
		require.Equal(t, false, testIdentityID.(identityID).Equals(nil))
		require.Equal(t, testIdentityID.(identityID).Bytes(), append(classificationID.Bytes(), baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID().Bytes()...))
		require.Equal(t, readIdentityID(testIdentityID.(identityID).String()), testIdentityID)
		require.Equal(t, identityIDFromInterface(testIdentityID.(identityID)), testIdentityID.(identityID))
		require.Equal(t, identityIDFromInterface(baseIDs.NewID("id")), identityID{ClassificationID: baseIDs.NewID(""), HashID: baseIDs.NewID("")})
	})
}
