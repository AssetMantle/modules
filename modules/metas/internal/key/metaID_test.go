package key

import (
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_MetaID_Methods(t *testing.T) {
	hashID := base.NewID("hashID")
	testMetaID := NewMetaID(hashID).(metaID)

	require.Equal(t, hashID.String(), testMetaID.String())
	require.Equal(t, true, testMetaID.Equals(testMetaID))
	require.Equal(t, false, testMetaID.Equals(metaID{HashID: base.NewID("tempHash")}))
	require.Equal(t, false, testMetaID.IsPartial())
	require.Equal(t, true, metaID{HashID: base.NewID("")}.IsPartial())
	require.Equal(t, true, testMetaID.Matches(testMetaID))
	require.Equal(t, false, testMetaID.Matches(metaID{HashID: base.NewID("tempHash")}))
	require.Equal(t, testMetaID, New(testMetaID))
	require.Equal(t, NewMetaID(base.NewID("")), New(base.NewID("")))
}
