package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenesis(t *testing.T) {

	Genesis := NewGenesis(base.KeyPrototype, base.MappablePrototype, []helpers.Mappable{}, []types.Parameter{})
	Genesis.Initialize(nil, nil)
	Error := Genesis.Validate()
	require.Nil(t, Error)
	Genesis.Default()
	Genesis.Encode()
	Genesis.Decode(Genesis.Encode())
	//Genesis.Import()
	//Genesis.Export()
}
