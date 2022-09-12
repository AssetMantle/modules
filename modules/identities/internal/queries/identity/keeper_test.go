// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
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

func CreateTestInput2(t *testing.T) (sdkTypes.Context, helpers.Keeper) {
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

	mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := params.NewKeeper(
		Codec,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters := parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))

	testQueryKeeper := keeperPrototype().Initialize(mapper, Parameters, []interface{}{})

	return context, testQueryKeeper
}

func Test_keeperPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.QueryKeeper
	}{
		// TODO: Add test cases.
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
	context, testQueryKeeper := CreateTestInput2(t)
	immutableProperties, err := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)
	mutableProperties, Error2 := utilities.ReadProperties("burn:S|100")
	require.Equal(t, nil, Error2)
	classificationID := baseIDs.NewID("ClassificationID")
	identityID := key.NewIdentityID(classificationID, immutableProperties)
	testQueryKeeper.(queryKeeper).mapper.NewCollection(context).Add(mappable.NewIdentity(identityID, immutableProperties, mutableProperties))

	type fields struct {
		mapper helpers.Keeper
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
		// TODO: Add test cases.
		{"+ve", fields{testQueryKeeper}, args{context: context, queryRequest: newQueryRequest(baseIDs.NewID("ClassificationID"))}, queryResponse{Success: true, Error: nil, List: testQueryKeeper.(queryKeeper).mapper.NewCollection(context).Fetch(key.FromID(identityID)).GetList()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//queryKeeper := queryKeeper{
			//	mapper: tt.fields.mapper,
			//}
			if got := tt.fields.mapper.(queryKeeper).Enquire(tt.args.context, tt.args.queryRequest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Enquire() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryKeeper_Initialize(t *testing.T) {
	context, testQueryKeeper := CreateTestInput2(t)
	immutableProperties, err := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)
	mutableProperties, Error2 := utilities.ReadProperties("burn:S|100")
	require.Equal(t, nil, Error2)
	classificationID := baseIDs.NewID("ClassificationID")
	identityID := key.NewIdentityID(classificationID, immutableProperties)
	testQueryKeeper.(queryKeeper).mapper.NewCollection(context).Add(mappable.NewIdentity(identityID, immutableProperties, mutableProperties))
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
		// TODO: Add more test cases.
		//{"+ve", fields{testQueryKeeper.(queryKeeper).mapper}, args{testQueryKeeper.(queryKeeper).mapper,dummy.Parameter}, queryKeeper{testQueryKeeper.(queryKeeper).mapper}},
		{"+ve with nil", fields{}, args{}, queryKeeper{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryKeeper := queryKeeper{
				mapper: tt.fields.mapper,
			}
			if got := queryKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.in2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
