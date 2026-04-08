// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supply

import (
	"cosmossdk.io/math"
	"context"
	storeTypes "cosmossdk.io/store/types"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	baseTypes "github.com/AssetMantle/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	"github.com/AssetMantle/modules/x/splits/parameters"
	"github.com/AssetMantle/modules/x/splits/record"
)

type TestKeepers struct {
	QueryKeeper helpers.QueryKeeper
}

func createTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {

	storeKey := storeTypes.NewKVStoreKey("test")
	Context := testutil.NewTestContext(t, storeKey)

	Mapper := baseHelpers.NewMapper(record.Prototype).Initialize(storeKey)
	parameterManager := parameters.Prototype().Initialize(storeKey)

	keepers := TestKeepers{
		QueryKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.QueryKeeper),
	}

	return Context, keepers, Mapper, parameterManager
}

func Test_keeperPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.QueryKeeper
	}{
		{"valid", queryKeeper{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keeperPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keeperPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryKeeper_Initialize(t *testing.T) {
	_, _, mapper, parameterManager := createTestInput(t)
	type fields struct {
		mapper helpers.Mapper
	}
	type args struct {
		mapper helpers.Mapper
		in1    helpers.ParameterManager
		in2    []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}{
		{"valid", fields{mapper}, args{mapper, parameterManager, []interface{}{}}, queryKeeper{mapper}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryKeeper := queryKeeper{
				mapper: tt.fields.mapper,
			}
			if got := queryKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.in2); got == nil {
				t.Errorf("Initialize() = nil, want non-nil")
			}
		})
	}
}

func Test_queryKeeper_Enquire(t *testing.T) {
	Context, keepers, Mapper, _ := createTestInput(t)
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testOwnerIdentityID := baseIDs.NewIdentityID(classificationID, immutables)
	testAssetID := baseDocuments.NewCoinAsset("OwnerID").GetCoinAssetID()
	testRate := math.OneInt()
	split := baseTypes.NewSplit(testRate)
	keepers.QueryKeeper.(queryKeeper).mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(record.NewRecord(baseIDs.NewSplitID(testAssetID, testOwnerIdentityID), split))

	type fields struct {
		mapper helpers.Mapper
	}
	type args struct {
		context      context.Context
		queryRequest helpers.QueryRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.QueryResponse
		wantErr bool
	}{
		{"valid", fields{Mapper}, args{sdkTypes.WrapSDKContext(Context), newQueryRequest(testAssetID)}, newQueryResponse(testRate), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryKeeper := queryKeeper{
				mapper: tt.fields.mapper,
			}
			got, err := queryKeeper.Enquire(tt.args.context, tt.args.queryRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Enquire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enquire() got = %v, want %v", got, tt.want)
			}
		})
	}
}
