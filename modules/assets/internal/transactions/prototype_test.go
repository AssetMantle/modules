package transactions

import (
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/define"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/deputize"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/mutate"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/renumerate"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/transactions/revoke"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrototype(t *testing.T) {
	want := base.NewTransactions(burn.Transaction,
		define.Transaction,
		deputize.Transaction,
		mint.Transaction,
		mutate.Transaction,
		renumerate.Transaction,
		revoke.Transaction)

	require.Equal(t, Prototype().Get(""), want.Get(""))

}
