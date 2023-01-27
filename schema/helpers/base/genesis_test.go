// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/utilities/test"
	baseTestUtilities "github.com/AssetMantle/modules/utilities/test/schema/helpers/base"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	parametersSchema "github.com/AssetMantle/modules/schema/parameters"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
)

func TestGenesis(t *testing.T) {
	context, storeKey, transientStoreKey, _ := test.SetupTest(t)
	codec := CodecPrototype()

	Mapper := NewMapper(baseTestUtilities.KeyPrototype, baseTestUtilities.MappablePrototype).Initialize(storeKey)

	mappableList := []helpers.Mappable{baseTestUtilities.NewMappable("test", "testValue")}
	ParameterList := []parametersSchema.Parameter{baseTypes.NewParameter(baseIDs.NewStringID("testParameter"), baseData.NewStringData("testData"), func(interface{}) error { return nil })}
	Parameters := NewParameters(ParameterList...)
	subspace := paramsTypes.NewSubspace(codec.GetProtoCodec(), codec.GetLegacyAmino(), storeKey, transientStoreKey, "test").WithKeyTable(Parameters.GetKeyTable())
	Parameters = Parameters.Initialize(subspace)

	GenesisState := baseTestUtilities.PrototypeGenesisState()

	Genesis := NewGenesis(baseTestUtilities.KeyPrototype, GenesisState).Initialize(mappableList, ParameterList)

	err := Genesis.Validate()
	require.Nil(t, err)

	require.Equal(t, mappableList, Genesis.Default().(genesis).GetMappableList())
	require.Equal(t, ParameterList, Genesis.Default().(genesis).GetParameterList())

	require.Equal(t, mappableList, Genesis.GetMappableList())
	require.Equal(t, ParameterList, Genesis.GetParameterList())

	rawMessage := Genesis.Encode(codec.GetProtoCodec())
	require.Equal(t, rawMessage, Genesis.Decode(codec.GetProtoCodec(), rawMessage).Encode(codec.GetProtoCodec()))

	require.NotPanics(t, func() {
		Genesis.Import(context.Context(), Mapper, Parameters)
	})
	require.NotPanics(t, func() {
		err := Genesis.Export(context.Context(), Mapper, Parameters).Validate()
		require.Nil(t, err)
	})
}
