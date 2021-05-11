package transactions

import (
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/transactions/reveal"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("reveal").GetName(), base.NewTransactions(
		reveal.Transaction,
	).Get("reveal").GetName())
}
