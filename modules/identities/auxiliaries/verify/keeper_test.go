package verify

import (
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
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

type TestKeepers1 struct {
	IdentitiesKeeper helpers.AuxiliaryKeeper
}

func CreateTestInput1(t *testing.T) (sdkTypes.Context, TestKeepers1) {
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

	keepers := TestKeepers1{
		IdentitiesKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.AuxiliaryKeeper),
	}

	return context, keepers
}

func Test_auxiliaryKeeper_Help(t *testing.T) {
	context, keepers := CreateTestInput1(t)
	defaultAddr := sdkTypes.AccAddress("addr")
	immutableProperties, _ := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	defaultClassificationID := baseIDs.NewID("test.cGn3HMW8M3t5gMDv-wXa9sseHnA=")
	defaultIdentityID := key.NewIdentityID(defaultClassificationID, immutableProperties)
	keepers.IdentitiesKeeper.(auxiliaryKeeper).mapper.NewCollection(context).Add(mappable.NewIdentity(defaultIdentityID, baseLists.NewPropertyList(), baseLists.NewPropertyList()))
	type fields struct {
		mapper              helpers.Mapper
		parameters          helpers.Parameters
		supplementAuxiliary helpers.Auxiliary
	}
	type args struct {
		context sdkTypes.Context
		request helpers.AuxiliaryRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.AuxiliaryResponse
	}{
		// TODO: Add test cases.
		{"+ve", fields{keepers.IdentitiesKeeper.(auxiliaryKeeper).mapper, keepers.IdentitiesKeeper.(auxiliaryKeeper).parameters, keepers.IdentitiesKeeper.(auxiliaryKeeper).supplementAuxiliary}, args{context, NewAuxiliaryRequest(defaultAddr, defaultIdentityID)}, newAuxiliaryResponse(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryKeeper := auxiliaryKeeper{
				mapper:              tt.fields.mapper,
				parameters:          tt.fields.parameters,
				supplementAuxiliary: tt.fields.supplementAuxiliary,
			}
			if got := auxiliaryKeeper.Help(tt.args.context, tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Help() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_auxiliaryKeeper_Initialize(t *testing.T) {
//
//	type fields struct {
//		mapper              helpers.Mapper
//		parameters          helpers.Parameters
//		supplementAuxiliary helpers.Auxiliary
//	}
//	type args struct {
//		mapper      helpers.Mapper
//		parameters  helpers.Parameters
//		auxiliaries []interface{}
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   helpers.Keeper
//	}{
//		// TODO: Add test cases.
//		{"+ve", fields{}, args{}, newAuxiliaryResponse(nil)},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			auxiliaryKeeper := auxiliaryKeeper{
//				mapper:              tt.fields.mapper,
//				parameters:          tt.fields.parameters,
//				supplementAuxiliary: tt.fields.supplementAuxiliary,
//			}
//			if got := auxiliaryKeeper.Initialize(tt.args.mapper, tt.args.parameters, tt.args.auxiliaries); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Initialize() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_keeperPrototype(t *testing.T) {
//	tests := []struct {
//		name string
//		want helpers.AuxiliaryKeeper
//	}{
//		// TODO: Add test cases.
//		{"+ve", newAuxiliaryResponse(nil)},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := keeperPrototype(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("keeperPrototype() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
