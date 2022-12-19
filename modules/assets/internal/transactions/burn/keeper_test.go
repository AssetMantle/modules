// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/assets/internal/mappable"
	"github.com/AssetMantle/modules/modules/assets/internal/parameters"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/renumerate"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

var (
	renumerateAuxiliary   helpers.Auxiliary
	maintainAuxiliary     helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
)

type TestKeepers struct {
	BurnKeeper helpers.TransactionKeeper
}

func createTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.Parameters) {
	var legacyAmino = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := params.NewKeeper(
		legacyAmino,
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

	renumerateAuxiliary = renumerate.AuxiliaryMock.Initialize(Mapper, Parameters)
	maintainAuxiliary = maintain.AuxiliaryMock.Initialize(Mapper, Parameters)
	supplementAuxiliary = supplement.AuxiliaryMock.Initialize(Mapper, Parameters)
	authenticateAuxiliary = authenticate.AuxiliaryMock.Initialize(Mapper, Parameters)

	keepers := TestKeepers{
		BurnKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.TransactionKeeper),
	}

	return context, keepers, Mapper, Parameters
}

func Test_keeperPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.TransactionKeeper
	}{
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
	_, _, Mapper, Parameters := createTestInput(t)
	type fields struct {
		mapper                helpers.Mapper
		renumerateAuxiliary   helpers.Auxiliary
		maintainAuxiliary     helpers.Auxiliary
		supplementAuxiliary   helpers.Auxiliary
		authenticateAuxiliary helpers.Auxiliary
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
		{"+ve", fields{Mapper, renumerateAuxiliary, maintainAuxiliary, supplementAuxiliary, authenticateAuxiliary}, args{Mapper, Parameters, []interface{}{}}, transactionKeeper{Mapper, renumerateAuxiliary, maintainAuxiliary, supplementAuxiliary, authenticateAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				renumerateAuxiliary:   tt.fields.renumerateAuxiliary,
				maintainAuxiliary:     tt.fields.maintainAuxiliary,
				supplementAuxiliary:   tt.fields.supplementAuxiliary,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers, Mapper, _ := createTestInput(t)
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData(baseLists.NewDataList()))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testAssetID := baseIDs.NewAssetID(classificationID, immutables)
	testAsset := base.NewAsset(classificationID, immutables, mutables)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	fromID := baseIDs.NewIdentityID(classificationID, immutables)
	keepers.BurnKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewMappable(testAsset))
	type fields struct {
		mapper                helpers.Mapper
		renumerateAuxiliary   helpers.Auxiliary
		maintainAuxiliary     helpers.Auxiliary
		supplementAuxiliary   helpers.Auxiliary
		authenticateAuxiliary helpers.Auxiliary
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
		{"+ve", fields{Mapper, renumerateAuxiliary, maintainAuxiliary, supplementAuxiliary, authenticateAuxiliary}, args{context, newMessage(fromAccAddress, fromID, testAssetID)}, newTransactionResponse(nil)},
		{"+ve", fields{Mapper, renumerateAuxiliary, maintainAuxiliary, supplementAuxiliary, authenticateAuxiliary}, args{context, newMessage(fromAccAddress, fromID, baseIDs.PrototypeAssetID())}, newTransactionResponse(errorConstants.EntityNotFound)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				renumerateAuxiliary:   tt.fields.renumerateAuxiliary,
				maintainAuxiliary:     tt.fields.maintainAuxiliary,
				supplementAuxiliary:   tt.fields.supplementAuxiliary,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
			}
			if got := transactionKeeper.Transact(tt.args.context, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
