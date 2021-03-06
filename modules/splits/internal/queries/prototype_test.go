package queries

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/queries/ownable"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/queries/split"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.Panics(t, func() {
		require.Equal(t, Prototype().Get("splits").GetName(), base.NewQueries(
			split.Query,
			ownable.Query,
		).Get("splits").GetName())
		require.Equal(t, Prototype().Get("ownable").GetName(), base.NewQueries(
			split.Query,
			ownable.Query,
		).Get("ownable").GetName())
	})
}
