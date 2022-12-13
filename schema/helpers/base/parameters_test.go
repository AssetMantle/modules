// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	parameters2 "github.com/AssetMantle/modules/schema/parameters"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
	baseTestUtilities "github.com/AssetMantle/modules/utilities/test/schema/helpers/base"
)

func TestParameters(t *testing.T) {
	context, storeKey, transientStoreKey := baseTestUtilities.SetupTest(t)
	codec := baseTestUtilities.MakeCodec()
	Parameter := baseTypes.NewParameter(baseIDs.NewStringID("testParameter"), baseData.NewStringData("testData"), func(interface{}) error { return nil })
	ParameterList := []parameters2.Parameter{Parameter}
	Parameters := NewParameters(ParameterList...)
	subspace := params.NewSubspace(codec, storeKey, transientStoreKey, "test").WithKeyTable(Parameters.GetKeyTable())
	subspace.SetParamSet(context, Parameters)
	Parameters = Parameters.Initialize(subspace).(parameters)

	require.NotNil(t, Parameters.ParamSetPairs())

	require.NotNil(t, Parameters.GetKeyTable())

	require.Equal(t, true, Parameters.Equal(Parameters))

	require.Equal(t, true, Parameters.GetList()[0].Equal(Parameter))
	require.Equal(t, `{"id":{"idString":"testParameter"},"data":{"value":"testData"}}`, Parameters.String())

	err := Parameters.Validate()
	require.Nil(t, err)

	require.NotPanics(t, func() {
		Parameters.Fetch(context, baseIDs.NewStringID("testParameter"))
	})

	require.Equal(t, "testData123", Parameters.Mutate(context,
		baseTypes.NewParameter(baseIDs.NewStringID("testParameter"), baseData.NewStringData("testData123"), func(interface{}) error { return nil })).Get(baseIDs.NewStringID("testParameter")).GetData().String())
}
