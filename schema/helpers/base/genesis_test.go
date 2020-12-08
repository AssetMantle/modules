package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenesis(t *testing.T) {

	Genesis := NewGenesis(base.KeyPrototype, base.MappablePrototype, []helpers.Mappable{}, []types.Parameter{}).Initialize(nil, nil).(genesis)

	Error := Genesis.Validate()
	require.Nil(t, Error)

	require.Equal(t, []helpers.Mappable{}, Genesis.Default().(genesis).MappableList)
	require.Equal(t, []types.Parameter{}, Genesis.Default().(genesis).defaultParameterList)

	require.Equal(t, Genesis.Encode(), Genesis.Decode(Genesis.Encode()).Encode())

	//TODO Genesis.Import()
	//TODO Genesis.Export()
}
