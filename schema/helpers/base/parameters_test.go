// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	parametersSchema "github.com/AssetMantle/modules/schema/parameters"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
	"github.com/AssetMantle/modules/utilities/test"
)

func TestParameters(t *testing.T) {
	context, storeKey, transientStoreKey := test.SetupTest(t)
	var legacyAmino = sdkCodec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()
	Parameter := baseTypes.NewParameter(baseIDs.NewStringID("testParameter"), baseData.NewStringData("testData"), func(interface{}) error { return nil })
	ParameterList := []parametersSchema.Parameter{Parameter}
	Parameters := NewParameters(ParameterList...)
	subspace := paramsTypes.NewSubspace(legacyAmino, storeKey, transientStoreKey, "test").WithKeyTable(Parameters.GetKeyTable())
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
