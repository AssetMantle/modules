// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base_test

import (
	"testing"

	storeTypes "cosmossdk.io/store/types"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	"github.com/AssetMantle/modules/x/metas/parameters/reveal_enabled"
)

func Test_parameter_manager_initialize(t *testing.T) {
	pm := baseHelpers.NewParameterManager(reveal_enabled.ValidatableParameter)
	require.NotNil(t, pm)

	storeKey := storeTypes.NewKVStoreKey("params")
	initialized := pm.Initialize(storeKey)
	require.NotNil(t, initialized)
}

func Test_parameter_manager_get_default(t *testing.T) {
	pm := baseHelpers.NewParameterManager(reveal_enabled.ValidatableParameter)
	storeKey := storeTypes.NewKVStoreKey("params")
	pm = pm.Initialize(storeKey)

	defaultList := pm.GetDefaultParameterList()
	require.NotNil(t, defaultList)

	params := defaultList.Get()
	require.Len(t, params, 1, "should have one default parameter")
}

func Test_parameter_manager_set(t *testing.T) {
	pm := baseHelpers.NewParameterManager(reveal_enabled.ValidatableParameter)
	storeKey := storeTypes.NewKVStoreKey("params")
	pm = pm.Initialize(storeKey)

	param := reveal_enabled.Parameter
	pm = pm.Set(param)

	got := pm.Get()
	require.NotNil(t, got)
	assert.Len(t, got.Get(), 1)
}

func Test_parameter_manager_validate(t *testing.T) {
	pm := baseHelpers.NewParameterManager(reveal_enabled.ValidatableParameter)
	storeKey := storeTypes.NewKVStoreKey("params")
	pm = pm.Initialize(storeKey)
	pm = pm.Set(reveal_enabled.Parameter)

	err := pm.Validate()
	assert.NoError(t, err)
}

func Test_parameter_manager_update_persists(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("params")
	ctx := testutil.NewTestContext(t, storeKey)

	pm := baseHelpers.NewParameterManager(reveal_enabled.ValidatableParameter)
	pm = pm.Initialize(storeKey)
	pm = pm.Set(reveal_enabled.Parameter)

	updated, err := pm.Update(ctx)
	require.NoError(t, err)
	require.NotNil(t, updated)

	// Fetch from store and verify
	fetched := pm.Fetch(ctx)
	require.NotNil(t, fetched)
	fetchedList := fetched.Get()
	require.NotNil(t, fetchedList)
	assert.Len(t, fetchedList.Get(), 1)
}

func Test_parameter_manager_set_with_different_value(t *testing.T) {
	pm := baseHelpers.NewParameterManager(reveal_enabled.ValidatableParameter)
	storeKey := storeTypes.NewKVStoreKey("params")
	pm = pm.Initialize(storeKey)

	pm = pm.Set(reveal_enabled.Parameter)

	falseParam := reveal_enabled.Parameter.Mutate(baseData.NewBooleanData(false).ToAnyData())
	pm = pm.Set(falseParam)

	err := pm.Validate()
	assert.NoError(t, err)
}
