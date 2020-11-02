package key

import (
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_IdentityID_Methods(t *testing.T) {

	classificationID := base.NewID("classificationID")
	hashID := base.NewID("hashID")
	testIdentityID := NewIdentityID(classificationID, hashID)
	testIdentityID2 := NewIdentityID(classificationID, base.NewID(""))
	key := New(testIdentityID)

	require.Equal(t, testIdentityID, identityID{ClassificationID: classificationID, HashID: hashID})
	require.Equal(t, New(testIdentityID), identityIDFromInterface(testIdentityID))
	require.Equal(t, testIdentityID.(identityID).IsPartial(), false)
	require.Equal(t, testIdentityID2.(identityID).IsPartial(), true)
	require.Equal(t, testIdentityID2.(identityID).Equals(testIdentityID), false)
	require.Equal(t, testIdentityID.(identityID).Matches(key), true)
	require.Equal(t, testIdentityID.(identityID).Matches(New(base.NewID("id"))), false)
	require.Equal(t, testIdentityID.(identityID).Bytes(), append(classificationID.Bytes(), hashID.Bytes()...))
	require.Equal(t, readIdentityID(testIdentityID.(identityID).String()), testIdentityID)
	require.Equal(t, identityIDFromInterface(testIdentityID.(identityID)), testIdentityID.(identityID))
	require.Equal(t, identityIDFromInterface(base.NewID("id")), identityID{ClassificationID: base.NewID(""), HashID: base.NewID("")})
}
