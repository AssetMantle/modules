package queries

import (
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/queries/asset"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Prototype(t *testing.T) {

	prototype := Prototype()
	require.Equal(t, asset.Query.GetName(), prototype.Get("assets").GetName())
}
