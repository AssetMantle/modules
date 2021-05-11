package metas

import (
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/block"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/queries"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/simulator"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/transactions"
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
