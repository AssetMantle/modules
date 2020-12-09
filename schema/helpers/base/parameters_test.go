package base

import (
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	baseTestUilities "github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParameters(t *testing.T) {

	codec := baseTestUilities.MakeCodec()
	_, storeKeys := baseTestUilities.SetupTest(t)
	validatorFunction := func(interface{}) error { return nil }
	Parameter := base.NewParameter(base.NewID("testParameter1"), base.NewStringData("testData1"), validatorFunction)
	Parameters := NewParameters(Parameter)

	Parameters = Parameters.Initialize(params.NewSubspace(codec, storeKeys, nil, "test"))

	require.NotNil(t, Parameters.ParamSetPairs())

	require.NotNil(t, Parameters.GetKeyTable())

	require.Equal(t, true, Parameters.Equal(Parameters))

	require.Equal(t, true, Parameters.GetList()[0].Equal(Parameter))
	require.Equal(t, `{"id":{"idString":"testParameter1"},"data":{"value":"testData1"}}`, Parameters.String())

	Error := Parameters.Validate()
	require.Nil(t, Error)

	//TODO Parameters.Fetch(context, Parameters.Get(base.NewID("testParameter1")).GetID())
	//TODO Parameters.Mutate(context, Parameters.Get(base.NewID("")))
}
