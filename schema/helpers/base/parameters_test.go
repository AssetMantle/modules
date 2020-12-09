package base

import (
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	baseTestUtilities "github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParameters(t *testing.T) {

	codec := baseTestUtilities.MakeCodec()
	context, storeKeys := baseTestUtilities.SetupTest(t)
	validatorFunction := func(interface{}) error { return nil }
	Parameter := base.NewParameter(base.NewID("testParameter1"), base.NewStringData("testData1"), validatorFunction)
	Parameters := NewParameters(Parameter)

	paramsSubspace := params.NewSubspace(codec, storeKeys, nil, "test").
		WithKeyTable(Parameters.GetKeyTable())
	initializedParameters := Parameters.Initialize(paramsSubspace).(parameters)
	//initializedParameters.paramsSubspace.SetParamSet(context, Parameters.ParamSetPairs())
	require.NotNil(t, initializedParameters.ParamSetPairs())

	require.NotNil(t, initializedParameters.GetKeyTable())

	require.Equal(t, true, initializedParameters.Equal(initializedParameters))

	require.Equal(t, true, initializedParameters.GetList()[0].Equal(Parameter))
	require.Equal(t, `{"id":{"idString":"testParameter1"},"data":{"value":"testData1"}}`, initializedParameters.String())

	Error := initializedParameters.Validate()
	require.Nil(t, Error)

	require.NotPanics(t, func() {
		initializedParameters.Fetch(context, base.NewID("testParameter1"))
	})
	//TODO initializedParameters.Mutate(context, initializedParameters.Get(base.NewID("")))
}
