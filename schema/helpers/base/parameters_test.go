/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	baseTestUtilities "github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParameters(t *testing.T) {

	context, storeKey, transientStoreKey := baseTestUtilities.SetupTest(t)
	codec := baseTestUtilities.MakeCodec()
	Parameter := base.NewParameter(base.NewID("testParameter"), base.NewStringData("testData"), func(interface{}) error { return nil })
	ParameterList := []types.Parameter{Parameter}
	Parameters := NewParameters(ParameterList...)
	subspace := params.NewSubspace(codec, storeKey, transientStoreKey, "test").WithKeyTable(Parameters.GetKeyTable())
	subspace.SetParamSet(context, Parameters)
	Parameters = Parameters.Initialize(subspace).(parameters)

	require.NotNil(t, Parameters.ParamSetPairs())

	require.NotNil(t, Parameters.GetKeyTable())

	require.Equal(t, true, Parameters.Equal(Parameters))

	require.Equal(t, true, Parameters.GetList()[0].Equal(Parameter))
	require.Equal(t, `{"id":{"idString":"testParameter"},"data":{"value":"testData"}}`, Parameters.String())

	Error := Parameters.Validate()
	require.Nil(t, Error)

	require.NotPanics(t, func() {
		Parameters.Fetch(context, base.NewID("testParameter"))
	})

	require.Equal(t, "testData123", Parameters.Mutate(context,
		base.NewParameter(base.NewID("testParameter"), base.NewStringData("testData123"), func(interface{}) error { return nil })).Get(base.NewID("testParameter")).GetData().String())
}
