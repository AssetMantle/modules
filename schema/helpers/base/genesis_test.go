// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
	baseTestUtilities "github.com/AssetMantle/modules/utilities/test/schema/helpers/base"
)

func TestGenesis(t *testing.T) {
	context, storeKey, transientStoreKey := baseTestUtilities.SetupTest(t)
	codec := baseTestUtilities.MakeCodec()
	Mapper := NewMapper(baseTestUtilities.KeyPrototype, baseTestUtilities.MappablePrototype).Initialize(storeKey)

	mappableList := []helpers.Mappable{baseTestUtilities.NewMappable("test", "testValue")}
	ParameterList := []types.Parameter{base.NewParameter(base.NewID("testParameter"), base.NewStringData("testData"), func(interface{}) error { return nil })}
	Parameters := NewParameters(ParameterList...)
	subspace := params.NewSubspace(codec, storeKey, transientStoreKey, "test").WithKeyTable(Parameters.GetKeyTable())
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
