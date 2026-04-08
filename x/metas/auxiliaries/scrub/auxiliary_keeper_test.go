// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	storeTypes "cosmossdk.io/store/types"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	"github.com/AssetMantle/modules/x/metas/mapper"
	"github.com/AssetMantle/modules/x/metas/parameters"
)

type TestKeepers struct {
	MetasKeeper helpers.AuxiliaryKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {
	storeKey := storeTypes.NewKVStoreKey("test")
	Context := testutil.NewTestContext(t, storeKey)

	Mapper := mapper.Prototype().Initialize(storeKey)
	parameterManager := parameters.Prototype().Initialize(storeKey)

	keepers := TestKeepers{
		MetasKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.AuxiliaryKeeper),
	}

	return Context, keepers
}

func Test_auxiliaryKeeper_Help_invalid_request(t *testing.T) {
	ctx, keepers := CreateTestInput(t)

	// Pass a wrong request type (use nil-ish invalid type)
	_, err := keepers.MetasKeeper.Help(ctx, nil)
	assert.Error(t, err, "nil request should return error")
}

func Test_auxiliaryKeeper_Help_empty_property_list(t *testing.T) {
	ctx, keepers := CreateTestInput(t)

	// Empty property list should succeed
	request := NewAuxiliaryRequest(baseLists.NewPropertyList())
	response, err := keepers.MetasKeeper.Help(ctx, request)
	require.NoError(t, err)
	assert.NotNil(t, response)

	propertyList := GetPropertiesFromResponse(response)
	assert.Empty(t, propertyList.Get())
}

func Test_auxiliaryKeeper_Help_meta_property(t *testing.T) {
	ctx, keepers := CreateTestInput(t)

	// Create a meta property with actual data
	metaProperty := baseProperties.NewMetaProperty(baseIDs.NewStringID("testKey"), baseData.NewStringData("testValue"))
	request := NewAuxiliaryRequest(baseLists.NewPropertyList(metaProperty))

	response, err := keepers.MetasKeeper.Help(ctx, request)
	require.NoError(t, err)
	assert.NotNil(t, response)

	propertyList := GetPropertiesFromResponse(response)
	require.Len(t, propertyList.Get(), 1)

	// The returned property should be scrubbed (mesa, not meta)
	scrubbedProp := propertyList.Get()[0]
	assert.False(t, scrubbedProp.IsMeta(), "scrubbed property should not be meta")
}

func Test_auxiliaryKeeper_Help_mesa_property_passthrough(t *testing.T) {
	ctx, keepers := CreateTestInput(t)

	// Create a mesa (non-meta) property
	mesaProperty := baseProperties.NewMesaProperty(baseIDs.NewStringID("mesaKey"), baseData.NewStringData("mesaValue"))
	request := NewAuxiliaryRequest(baseLists.NewPropertyList(mesaProperty))

	response, err := keepers.MetasKeeper.Help(ctx, request)
	require.NoError(t, err)
	assert.NotNil(t, response)

	propertyList := GetPropertiesFromResponse(response)
	require.Len(t, propertyList.Get(), 1)
}

func Test_auxiliaryKeeper_Help_mixed_properties(t *testing.T) {
	ctx, keepers := CreateTestInput(t)

	metaProperty := baseProperties.NewMetaProperty(baseIDs.NewStringID("key1"), baseData.NewStringData("value1"))
	mesaProperty := baseProperties.NewMesaProperty(baseIDs.NewStringID("key2"), baseData.NewStringData("value2"))
	request := NewAuxiliaryRequest(baseLists.NewPropertyList(metaProperty, mesaProperty))

	response, err := keepers.MetasKeeper.Help(ctx, request)
	require.NoError(t, err)
	assert.NotNil(t, response)

	propertyList := GetPropertiesFromResponse(response)
	assert.Len(t, propertyList.Get(), 2)
}
