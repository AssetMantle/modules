package define

import (
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/super"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
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

func CreateTestInputContext(t *testing.T) (sdkTypes.Context, TestKeepers) {
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

	scrubAuxiliary := scrub.AuxiliaryMock.Initialize(Mapper, Parameters)
	defineAuxiliary := define.AuxiliaryMock.Initialize(Mapper, Parameters)
	superAuxiliary := super.AuxiliaryMock.Initialize(Mapper, Parameters)
	keepers := TestKeepers{
		IdentitiesKeeper: keeperPrototype().Initialize(Mapper, Parameters,
			[]interface{}{scrubAuxiliary,
				defineAuxiliary, superAuxiliary}).(helpers.TransactionKeeper),
	}

	return context, keepers
}

func Test_keeperPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.TransactionKeeper
	}{
		// TODO: Add test cases.
		{"+ve", transactionKeeper{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keeperPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keeperPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Initialize(t *testing.T) {
	//context, keepers := CreateTestInputContext(t)

	type fields struct {
		mapper          helpers.Mapper
		defineAuxiliary helpers.Auxiliary
		scrubAuxiliary  helpers.Auxiliary
		superAuxiliary  helpers.Auxiliary
		verifyAuxiliary helpers.Auxiliary
	}
	type args struct {
		mapper      helpers.Mapper
		in1         helpers.Parameters
		auxiliaries []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}{
		// TODO: Add test cases.
		//{"+ve", fields{}, args{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:          tt.fields.mapper,
				defineAuxiliary: tt.fields.defineAuxiliary,
				scrubAuxiliary:  tt.fields.scrubAuxiliary,
				superAuxiliary:  tt.fields.superAuxiliary,
				verifyAuxiliary: tt.fields.verifyAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.auxiliaries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	type fields struct {
		mapper          helpers.Mapper
		defineAuxiliary helpers.Auxiliary
		scrubAuxiliary  helpers.Auxiliary
		superAuxiliary  helpers.Auxiliary
		verifyAuxiliary helpers.Auxiliary
	}
	type args struct {
		context sdkTypes.Context
		msg     sdkTypes.Msg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.TransactionResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:          tt.fields.mapper,
				defineAuxiliary: tt.fields.defineAuxiliary,
				scrubAuxiliary:  tt.fields.scrubAuxiliary,
				superAuxiliary:  tt.fields.superAuxiliary,
				verifyAuxiliary: tt.fields.verifyAuxiliary,
			}
			if got := transactionKeeper.Transact(tt.args.context, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
