package transactions

import (
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/define"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/deputize"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/issue"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/nub"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/provision"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/revoke"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions/unprovision"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("unprovision").GetName(), base.NewTransactions(
		define.Transaction,
		deputize.Transaction,
		issue.Transaction,
		nub.Transaction,
		provision.Transaction,
		revoke.Transaction,
		unprovision.Transaction,
	).Get("unprovision").GetName())
}
