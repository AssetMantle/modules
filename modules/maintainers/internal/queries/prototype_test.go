package queries

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/queries/maintainer"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.Panics(t, func() {
		require.Equal(t, Prototype().Get("maintainer").GetName(), base.NewQueries(
			maintainer.Query,
		).Get("maintainer").GetName())
	})
}
