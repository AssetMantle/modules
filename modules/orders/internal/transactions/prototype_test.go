package transactions

import (
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/cancel"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/define"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/immediate"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/make"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/modify"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/transactions/take"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("cancel").GetName(), base.NewTransactions(
		cancel.Transaction,
		define.Transaction,
		immediate.Transaction,
		make.Transaction,
		modify.Transaction,
		take.Transaction,
	).Get("cancel").GetName())
	require.Equal(t, Prototype().Get("define").GetName(), base.NewTransactions(
		cancel.Transaction,
		define.Transaction,
		immediate.Transaction,
		make.Transaction,
		modify.Transaction,
		take.Transaction,
	).Get("define").GetName())
	require.Equal(t, Prototype().Get("immediate").GetName(), base.NewTransactions(
		cancel.Transaction,
		define.Transaction,
		immediate.Transaction,
		make.Transaction,
		modify.Transaction,
		take.Transaction,
	).Get("immediate").GetName())
	require.Equal(t, Prototype().Get("make").GetName(), base.NewTransactions(
		cancel.Transaction,
		define.Transaction,
		immediate.Transaction,
		make.Transaction,
		modify.Transaction,
		take.Transaction,
	).Get("make").GetName())
	require.Equal(t, Prototype().Get("modify").GetName(), base.NewTransactions(
		cancel.Transaction,
		define.Transaction,
		immediate.Transaction,
		make.Transaction,
		modify.Transaction,
		take.Transaction,
	).Get("modify").GetName())
	require.Equal(t, Prototype().Get("take").GetName(), base.NewTransactions(
		cancel.Transaction,
		define.Transaction,
		immediate.Transaction,
		make.Transaction,
		modify.Transaction,
		take.Transaction,
	).Get("take").GetName())

}
