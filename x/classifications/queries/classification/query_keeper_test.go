// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package classification

import (
	storeTypes "cosmossdk.io/store/types"
	"testing"

	"github.com/AssetMantle/modules/x/classifications/record"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	"github.com/AssetMantle/modules/x/classifications/key"
	"github.com/AssetMantle/modules/x/classifications/mapper"
	"github.com/AssetMantle/modules/x/classifications/parameters"
)

func CreateTestInput2(t *testing.T) (sdkTypes.Context, helpers.Keeper) {

	storeKey := storeTypes.NewKVStoreKey("test")
	context := testutil.NewTestContext(t, storeKey)

	Mapper := mapper.Prototype().Initialize(storeKey)
	parameterManager := parameters.Prototype().Initialize(storeKey)
	testQueryKeeper := keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{})

	return context, testQueryKeeper
}

func Test_Query_Keeper_Classification(t *testing.T) {
	context, keepers := CreateTestInput2(t)
	immutableProperties, err := base.NewPropertyList().FromMetaPropertiesString("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)
	mutableProperties, Error2 := base.NewPropertyList().FromMetaPropertiesString("burn:S|100")
	require.Equal(t, nil, Error2)

	classificationID := baseIDs.NewClassificationID(baseQualified.NewImmutables(immutableProperties), baseQualified.NewMutables(mutableProperties))
	keepers.(queryKeeper).mapper.NewCollection(sdkTypes.WrapSDKContext(context)).Add(record.NewRecord(baseDocuments.NewClassification(baseQualified.NewImmutables(immutableProperties), baseQualified.NewMutables(mutableProperties))))

	testQueryRequest := newQueryRequest(classificationID)
	queryResponse, err := keepers.(queryKeeper).Enquire(sdkTypes.WrapSDKContext(context), testQueryRequest)
	require.Equal(t, &QueryResponse{Record: keepers.(queryKeeper).mapper.NewCollection(sdkTypes.WrapSDKContext(context)).Fetch(key.NewKey(classificationID)).FetchRecord(key.NewKey(classificationID)).(*record.Record)}, queryResponse, testQueryRequest)

}
