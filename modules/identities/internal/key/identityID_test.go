/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"testing"

	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_IdentityID_Methods(t *testing.T) {

	classificationID := base.NewID("classificationID")
	immutableProperties, _ := base.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	emptyImmutableProperties, _ := base.ReadProperties("")
	testIdentityID := NewIdentityID(classificationID, immutableProperties)
	testIdentityID2 := NewIdentityID(classificationID, emptyImmutableProperties)
	key := FromID(testIdentityID)

	require.NotPanics(t, func() {
		require.Equal(t, testIdentityID, identityID{ClassificationID: classificationID, HashID: baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID()})
		require.Equal(t, FromID(testIdentityID), identityIDFromInterface(testIdentityID))
		require.Equal(t, false, testIdentityID.(identityID).IsPartial())
		require.Equal(t, true, testIdentityID2.(identityID).IsPartial())
		require.Equal(t, false, testIdentityID2.(identityID).Equals(testIdentityID))
		require.Equal(t, true, testIdentityID.(identityID).Matches(key))
		require.Equal(t, false, testIdentityID.(identityID).Matches(FromID(base.NewID("id"))))
		require.Equal(t, false, testIdentityID.(identityID).Matches(nil))
		require.Equal(t, testIdentityID.(identityID).Bytes(), append(classificationID.Bytes(), baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID().Bytes()...))
		require.Equal(t, readIdentityID(testIdentityID.(identityID).String()), testIdentityID)
		require.Equal(t, identityIDFromInterface(testIdentityID.(identityID)), testIdentityID.(identityID))
		require.Equal(t, identityIDFromInterface(base.NewID("id")), identityID{ClassificationID: base.NewID(""), HashID: base.NewID("")})
	})
}
