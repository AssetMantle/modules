package genesis

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.Panics(t, func() {
		require.Equal(t, Prototype(), base.NewGenesis(key.Prototype, mappable.Prototype, []helpers.Mappable{}, parameters.Prototype().GetList()))
	})
}
