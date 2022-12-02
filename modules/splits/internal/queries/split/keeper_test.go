// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package split

import (
	"fmt"
	"github.com/AssetMantle/modules/modules/splits/internal/key"
	"github.com/AssetMantle/modules/modules/splits/internal/mappable"
	"github.com/AssetMantle/modules/modules/splits/internal/parameters"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIds "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
	"reflect"
	"testing"
)

type TestKeepers struct {
	SplitKeeper helpers.QueryKeeper
}

func createTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.Parameters) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := params.NewKeeper(
		Codec,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters := parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		SplitKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.QueryKeeper),
	}

	return context, keepers, Mapper, Parameters
}

func Test_keeperPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.QueryKeeper
	}{
		{"+ve", queryKeeper{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keeperPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keeperPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryKeeper_Enquire(t *testing.T) {
	context, keepers, Mapper, _ := createTestInput(t)
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIds.NewClassificationID(immutables, mutables)
	testOwnerIdentityID := baseIds.NewIdentityID(classificationID, immutables)
	testOwnableID := baseIds.NewOwnableID(baseIds.NewStringID("OwnerID"))
	splitID := baseIds.NewSplitID(testOwnerIdentityID, testOwnableID)
	testRate := sdkTypes.NewDec(1)
	split := baseTypes.NewSplit(testOwnerIdentityID, testOwnableID, testRate)
	type fields struct {
		mapper helpers.Mapper
	}
	type args struct {
		context      sdkTypes.Context
		queryRequest helpers.QueryRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.QueryResponse
	}{
		{"+ve", fields{Mapper}, args{context, newQueryRequest(splitID)}, newQueryResponse(keepers.SplitKeeper.(queryKeeper).mapper.NewCollection(context).Add(mappable.NewMappable(split)), nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryKeeper := queryKeeper{
				mapper: tt.fields.mapper,
			}
			if got := queryKeeper.Enquire(tt.args.context, tt.args.queryRequest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enquire() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryKeeper_Initialize(t *testing.T) {
	_, _, Mapper, Parameters := createTestInput(t)
	type fields struct {
		mapper helpers.Mapper
	}
	type args struct {
		mapper helpers.Mapper
		in1    helpers.Parameters
		in2    []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}{
		{"+ve", fields{Mapper}, args{Mapper, Parameters, []interface{}{}}, queryKeeper{Mapper}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryKeeper := queryKeeper{
				mapper: tt.fields.mapper,
			}
			if got := queryKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.in2); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
