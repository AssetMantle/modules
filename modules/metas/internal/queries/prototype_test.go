package queries

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/queries/meta"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("metas").GetName(), base.NewQueries(
		meta.Query,
	).Get("metas").GetName())
}
