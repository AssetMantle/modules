package classifications

import (
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/block"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/queries"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/simulator"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/transactions"
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
