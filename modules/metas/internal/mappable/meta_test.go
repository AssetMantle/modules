package mappable

import (
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Meta_Methods(t *testing.T) {
	//hashID := base.NewID("hashID")
	data := base.NewStringData("Data")

	testMeta := NewMeta(data).(meta)
	require.Equal(t, meta{Data: data}, testMeta)
	require.Equal(t, data, testMeta.GetData())
	require.Equal(t, key.NewMetaID(base.NewID(data.GenerateHash())), testMeta.GetKey())
	require.Equal(t, base.NewID(data.GenerateHash()), testMeta.GetID())

}
