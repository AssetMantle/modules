package identities

import (
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/block"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/queries"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/simulator"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/transactions"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Name(), base.NewModule(
		module.Name,
		auxiliaries.Prototype,
		genesis.Prototype,
		mapper.Prototype,
		parameters.Prototype,
		queries.Prototype,
		simulator.Prototype,
		transactions.Prototype,
		block.Prototype,
	).Name())
}
