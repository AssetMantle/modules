package queries

import (
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/queries/meta"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("metas").GetName(),base.NewQueries(
		meta.Query,
	).Get("metas").GetName())
}
