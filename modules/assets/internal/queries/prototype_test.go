package queries

import (
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/queries/asset"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Prototype(t *testing.T) {
	require.Equal(t, base.NewQueries(asset.Query), Prototype())
}
