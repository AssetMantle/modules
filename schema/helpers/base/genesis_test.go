// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	parameters2 "github.com/AssetMantle/modules/schema/parameters"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
	baseTestUtilities "github.com/AssetMantle/modules/utilities/test/schema/helpers/base"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenesis(t *testing.T) {
	context, storeKey, transientStoreKey := baseTestUtilities.SetupTest(t)
	codec := baseTestUtilities.MakeCodec()
	Mapper := NewMapper(baseTestUtilities.KeyPrototype, baseTestUtilities.MappablePrototype).Initialize(storeKey)

	mappableList := []helpers.Mappable{baseTestUtilities.NewMappable("test", "testValue")}
	ParameterList := []parameters2.Parameter{baseTypes.NewParameter(baseIDs.NewStringID("testParameter"), baseData.NewStringData("testData"), func(interface{}) error { return nil })}
	Parameters := NewParameters(ParameterList...)
	subspace := paramsTypes.NewSubspace(codec, storeKey, transientStoreKey, "test").WithKeyTable(Parameters.GetKeyTable())
	Parameters = Parameters.Initialize(subspace)

	Genesis := NewGenesis(baseTestUtilities.KeyPrototype, baseTestUtilities.MappablePrototype, mappableList, ParameterList).Initialize(mappableList, ParameterList).(genesis)

	err := Genesis.Validate()
	require.Nil(t, err)

	require.Equal(t, mappableList, Genesis.Default().(genesis).MappableList)
	require.Equal(t, ParameterList, Genesis.Default().(genesis).defaultParameterList)

	require.Equal(t, mappableList, Genesis.GetMappableList())
	require.Equal(t, ParameterList, Genesis.GetParameterList())

	require.Equal(t, Genesis.Encode(), Genesis.Decode(Genesis.Encode()).Encode())

	require.NotPanics(t, func() {
		Genesis.Import(context, Mapper, Parameters)
	})
	require.NotPanics(t, func() {
		err := Genesis.Export(context, Mapper, Parameters).Validate()
		require.Nil(t, err)
	})
}
