package queries

import (
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/queries/order"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("orders").GetName(),base.NewQueries(
		order.Query,
	).Get("orders").GetName())
}
