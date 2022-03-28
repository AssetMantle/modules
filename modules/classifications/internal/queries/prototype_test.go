package queries

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/queries/classification"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.Panics(t, func() {
		require.Equal(t, Prototype().Get("classification").GetName(), base.NewQueries(
			classification.Query,
		).Get("classification").GetName())
	})
}
