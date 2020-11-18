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
	require.Equal(t, false, testIdentityID.(identityID).IsPartial())
	require.Equal(t, true, testIdentityID2.(identityID).IsPartial())
	require.Equal(t, false, testIdentityID2.(identityID).Equals(testIdentityID))
	require.Equal(t, true, testIdentityID.(identityID).Matches(key))
	require.Equal(t, false, testIdentityID.(identityID).Matches(New(base.NewID("id"))))
	require.Equal(t, false, testIdentityID.(identityID).Matches(nil))
	require.Equal(t, testIdentityID.(identityID).Bytes(), append(classificationID.Bytes(), hashID.Bytes()...))
	require.Equal(t, readIdentityID(testIdentityID.(identityID).String()), testIdentityID)
	require.Equal(t, identityIDFromInterface(testIdentityID.(identityID)), testIdentityID.(identityID))
	require.Equal(t, identityIDFromInterface(base.NewID("id")), identityID{ClassificationID: base.NewID(""), HashID: base.NewID("")})
}
