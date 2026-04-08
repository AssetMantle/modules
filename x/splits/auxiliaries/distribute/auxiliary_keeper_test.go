// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package distribute

import (
	"testing"

	"cosmossdk.io/math"
	storeTypes "cosmossdk.io/store/types"
	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/AssetMantle/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	splitsMapper "github.com/AssetMantle/modules/x/splits/mapper"
	splitsParameters "github.com/AssetMantle/modules/x/splits/parameters"
	splitsRecord "github.com/AssetMantle/modules/x/splits/record"
)

func setupDistributeKeeper(t *testing.T) (helpers.AuxiliaryKeeper, helpers.Mapper, sdkTypes.Context) {
	t.Helper()
	storeKey := storeTypes.NewKVStoreKey("test")
	paramStoreKey := storeTypes.NewKVStoreKey("testParams")
	ctx := testutil.NewTestContext(t, storeKey, paramStoreKey)

	m := splitsMapper.Prototype().Initialize(storeKey)
	pm := splitsParameters.Prototype().Initialize(paramStoreKey)

	keeper := keeperPrototype().Initialize(m, pm, []interface{}{}).(helpers.AuxiliaryKeeper)
	return keeper, m, ctx
}

// helper to create a unique asset ID
func testAssetID(name string) *baseIDs.AssetID {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(
		baseProperties.NewMesaProperty(baseIDs.NewStringID("assetName"), baseData.NewStringData(name)),
	))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList())
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	return baseIDs.NewAssetID(classificationID, immutables).(*baseIDs.AssetID)
}

func Test_auxiliaryKeeper_Help_invalid_request(t *testing.T) {
	keeper, _, ctx := setupDistributeKeeper(t)

	_, err := keeper.Help(ctx, nil)
	assert.Error(t, err, "nil request should return error")
}

func Test_auxiliaryKeeper_Help_zero_amount(t *testing.T) {
	keeper, _, ctx := setupDistributeKeeper(t)

	assetID := testAssetID("asset1")
	distID := testAssetID("dist1")
	fromID := testutil.TestIdentityID()

	request := NewAuxiliaryRequest(assetID, distID, fromID, math.ZeroInt())
	_, err := keeper.Help(ctx, request)
	assert.Error(t, err, "zero amount should fail")
	assert.Contains(t, err.Error(), "positive")
}

func Test_auxiliaryKeeper_Help_no_holders(t *testing.T) {
	keeper, m, ctx := setupDistributeKeeper(t)

	assetID := testAssetID("asset2")
	distID := testAssetID("dist2")
	fromID := testutil.TestIdentityID()

	// Add funder's distribution asset balance
	coll := m.NewCollection(ctx)
	splitID := baseIDs.NewSplitID(distID, fromID)
	coll.Add(splitsRecord.NewRecord(splitID, base.NewSplit(math.NewInt(1000))))

	request := NewAuxiliaryRequest(assetID, distID, fromID, math.NewInt(100))
	_, err := keeper.Help(ctx, request)
	assert.Error(t, err, "no holders should fail")
	assert.Contains(t, err.Error(), "no holders")
}

func Test_auxiliaryKeeper_Help_happy_path_single_holder(t *testing.T) {
	keeper, m, ctx := setupDistributeKeeper(t)

	assetID := testAssetID("asset3")
	distID := testAssetID("dist3")
	fromID := testutil.TestIdentityID()
	holderID := testutil.TestIdentityID()

	coll := m.NewCollection(ctx)

	// Give funder distribution tokens
	funderSplitID := baseIDs.NewSplitID(distID, fromID)
	coll.Add(splitsRecord.NewRecord(funderSplitID, base.NewSplit(math.NewInt(1000))))

	// Give holder some of the target asset
	holderSplitID := baseIDs.NewSplitID(assetID, holderID)
	coll.Add(splitsRecord.NewRecord(holderSplitID, base.NewSplit(math.NewInt(500))))

	request := NewAuxiliaryRequest(assetID, distID, fromID, math.NewInt(100))
	response, err := keeper.Help(ctx, request)
	require.NoError(t, err)
	assert.NotNil(t, response)
}

func Test_auxiliaryKeeper_Help_insufficient_funder_balance(t *testing.T) {
	keeper, m, ctx := setupDistributeKeeper(t)

	assetID := testAssetID("asset4")
	distID := testAssetID("dist4")
	fromID := testutil.TestIdentityID()
	holderID := testutil.TestIdentityID()

	coll := m.NewCollection(ctx)

	// Funder has only 10 tokens but tries to distribute 100
	funderSplitID := baseIDs.NewSplitID(distID, fromID)
	coll.Add(splitsRecord.NewRecord(funderSplitID, base.NewSplit(math.NewInt(10))))

	holderSplitID := baseIDs.NewSplitID(assetID, holderID)
	coll.Add(splitsRecord.NewRecord(holderSplitID, base.NewSplit(math.NewInt(500))))

	request := NewAuxiliaryRequest(assetID, distID, fromID, math.NewInt(100))
	_, err := keeper.Help(ctx, request)
	assert.Error(t, err, "insufficient funder balance should fail")
}

func Test_keeperPrototype(t *testing.T) {
	got := keeperPrototype()
	assert.Equal(t, auxiliaryKeeper{}, got)
}

func Test_auxiliaryKeeper_Initialize(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	paramStoreKey := storeTypes.NewKVStoreKey("testParams")

	m := splitsMapper.Prototype().Initialize(storeKey)
	pm := splitsParameters.Prototype().Initialize(paramStoreKey)

	keeper := keeperPrototype().Initialize(m, pm, []interface{}{})
	require.NotNil(t, keeper)
}
