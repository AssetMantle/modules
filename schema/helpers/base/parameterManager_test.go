// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseParameters "github.com/AssetMantle/modules/schema/parameters/base"
	parametersSchema "github.com/AssetMantle/modules/schema/parameters/base"
	"github.com/AssetMantle/modules/utilities/test"
)

func TestParameters(t *testing.T) {
	context, storeKey, transientStoreKey, _ := test.SetupTest(t)

	codec := CodecPrototype()

	Parameter := baseParameters.NewParameter(baseIDs.NewStringID("testParameter"), baseData.NewStringData("testData"), func(interface{}) error { return nil })
	parameters := []parametersSchema.Parameter{Parameter}
	ParameterManager := NewParameterManager(parameters...).(*parameterManager)
	subspace := paramsTypes.NewSubspace(codec.GetProtoCodec(), codec.GetLegacyAmino(), storeKey, transientStoreKey, "test").WithKeyTable(ParameterManager.GetKeyTable())
	subspace.SetParamSet(sdkTypes.UnwrapSDKContext(context), ParameterManager)
	ParameterManager = ParameterManager.Initialize(subspace).(*parameterManager)

	require.NotNil(t, ParameterManager.ParamSetPairs())

	require.NotNil(t, ParameterManager.GetKeyTable())

	require.Equal(t, true, ParameterManager.Equal(ParameterManager))

	require.Equal(t, true, ParameterManager.GetList()[0].Equal(Parameter))
	require.Equal(t, `{"id":{"idString":"testParameter"},"data":{"value":"testData"}}`, ParameterManager.String())

	err := ParameterManager.Validate()
	require.Nil(t, err)

	require.NotPanics(t, func() {
		ParameterManager.Fetch(context, baseIDs.NewStringID("testParameter"))
	})

	require.Equal(t, "testData123", ParameterManager.Mutate(context,
		baseParameters.NewParameter(baseIDs.NewStringID("testParameter"), baseData.NewStringData("testData123"), func(interface{}) error { return nil })).Get(baseIDs.NewStringID("testParameter")).GetData().AsString())
}
