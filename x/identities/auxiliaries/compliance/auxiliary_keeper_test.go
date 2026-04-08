// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package compliance

import (
	"testing"

	"cosmossdk.io/math"
	storeTypes "cosmossdk.io/store/types"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	identitiesMapper "github.com/AssetMantle/modules/x/identities/mapper"
	identitiesParameters "github.com/AssetMantle/modules/x/identities/parameters"
	identitiesRecord "github.com/AssetMantle/modules/x/identities/record"
)

func setupKeeper(t *testing.T) (helpers.AuxiliaryKeeper, helpers.Mapper, helpers.ParameterManager, sdkTypes.Context) {
	t.Helper()
	storeKey := storeTypes.NewKVStoreKey("test")
	paramStoreKey := storeTypes.NewKVStoreKey("testParams")
	ctx := testutil.NewTestContext(t, storeKey, paramStoreKey)

	m := identitiesMapper.Prototype().Initialize(storeKey)
	pm := identitiesParameters.Prototype().Initialize(paramStoreKey)

	keeper := keeperPrototype().Initialize(m, pm, []interface{}{}).(helpers.AuxiliaryKeeper)
	return keeper, m, pm, ctx
}

func Test_auxiliaryKeeper_Help_invalid_request(t *testing.T) {
	keeper, _, _, ctx := setupKeeper(t)

	_, err := keeper.Help(ctx, nil)
	assert.Error(t, err, "nil request should return error")
}

func Test_auxiliaryKeeper_Help_identity_not_found(t *testing.T) {
	keeper, _, _, ctx := setupKeeper(t)

	identityID := testutil.TestIdentityID()
	request := NewAuxiliaryRequest(identityID, math.ZeroInt(), "", false)

	_, err := keeper.Help(ctx, request)
	assert.Error(t, err, "non-existent identity should return not found error")
	assert.Contains(t, err.Error(), "not found")
}

func Test_auxiliaryKeeper_Help_happy_path(t *testing.T) {
	keeper, m, _, ctx := setupKeeper(t)

	// Create a module identity and add it to the store
	moduleIdentity := baseDocuments.NewIdentityFromDocument(baseDocuments.NewModuleIdentity("testModule"))
	identityID := baseIDs.NewIdentityID(moduleIdentity.GetClassificationID(), moduleIdentity.GetImmutables())

	coll := m.NewCollection(ctx)
	coll.Add(identitiesRecord.NewRecord(moduleIdentity))

	// Module identity has default compliance tier 0, no jurisdiction, not sanctions cleared
	// With MinTier=0, no jurisdiction check, no sanctions check => should pass
	request := NewAuxiliaryRequest(identityID, math.ZeroInt(), "", false)
	response, err := keeper.Help(ctx, request)
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.True(t, GetComplianceFromResponse(response))
}

func Test_auxiliaryKeeper_Help_insufficient_tier(t *testing.T) {
	keeper, m, _, ctx := setupKeeper(t)

	moduleIdentity := baseDocuments.NewIdentityFromDocument(baseDocuments.NewModuleIdentity("testModule2"))
	identityID := baseIDs.NewIdentityID(moduleIdentity.GetClassificationID(), moduleIdentity.GetImmutables())

	coll := m.NewCollection(ctx)
	coll.Add(identitiesRecord.NewRecord(moduleIdentity))

	// Module identity has tier 0, require tier 1 => should fail
	request := NewAuxiliaryRequest(identityID, math.NewInt(1), "", false)
	_, err := keeper.Help(ctx, request)
	assert.Error(t, err, "insufficient compliance tier should fail")
	assert.Contains(t, err.Error(), "compliance tier")
}

func Test_auxiliaryKeeper_Help_sanctions_required(t *testing.T) {
	keeper, m, _, ctx := setupKeeper(t)

	moduleIdentity := baseDocuments.NewIdentityFromDocument(baseDocuments.NewModuleIdentity("testModule3"))
	identityID := baseIDs.NewIdentityID(moduleIdentity.GetClassificationID(), moduleIdentity.GetImmutables())

	coll := m.NewCollection(ctx)
	coll.Add(identitiesRecord.NewRecord(moduleIdentity))

	// Module identity defaults to sanctionsCleared=false, require sanctions => should fail
	request := NewAuxiliaryRequest(identityID, math.ZeroInt(), "", true)
	_, err := keeper.Help(ctx, request)
	assert.Error(t, err, "sanctions requirement should fail for default identity")
	assert.Contains(t, err.Error(), "sanctions")
}

func Test_keeperPrototype(t *testing.T) {
	got := keeperPrototype()
	assert.Equal(t, auxiliaryKeeper{}, got)
}

func Test_auxiliaryKeeper_Initialize(t *testing.T) {
	keeper, _, pm, _ := setupKeeper(t)
	_ = pm

	require.NotNil(t, keeper)
	// setupKeeper already calls keeperPrototype().Initialize(m, pm, []interface{}{})
	// so this validates the Initialize path
}
