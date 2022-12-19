// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"fmt"
	typesTendermint "github.com/tendermint/tendermint/proto/tendermint/types"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/super"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseDocuments "github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIds "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

type TestKeepers struct {
	DefineKeeper helpers.TransactionKeeper
}

var (
	authenticateAuxiliary helpers.Auxiliary
	defineAuxiliary       helpers.Auxiliary
	superAuxiliary        helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	Parameters            helpers.Parameters
)

func createTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper) {
	var legacyAmino = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	ParamsKeeper := paramsKeeper.NewKeeper(
		legacyAmino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters = parameters.Prototype().Initialize(ParamsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, typesTendermint.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	authenticateAuxiliary = authenticate.AuxiliaryMock.Initialize(Mapper, Parameters)
	defineAuxiliary = define.AuxiliaryMock.Initialize(Mapper, Parameters)
	superAuxiliary = super.AuxiliaryMock.Initialize(Mapper, Parameters)
	supplementAuxiliary = supplement.AuxiliaryMock.Initialize(Mapper, Parameters)
	keepers := TestKeepers{
		DefineKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{authenticateAuxiliary}).(helpers.TransactionKeeper),
	}

	return context, keepers, Mapper
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
	_, _, Mapper := createTestInput(t)
	type fields struct {
		mapper                helpers.Mapper
		authenticateAuxiliary helpers.Auxiliary
		defineAuxiliary       helpers.Auxiliary
		superAuxiliary        helpers.Auxiliary
		supplementAuxiliary   helpers.Auxiliary
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
		{"+ve with nil", fields{}, args{}, transactionKeeper{}},
		{"+ve", fields{Mapper, authenticateAuxiliary, defineAuxiliary, superAuxiliary, supplementAuxiliary}, args{Mapper, Parameters, []interface{}{}}, transactionKeeper{Mapper, authenticateAuxiliary, defineAuxiliary, superAuxiliary, supplementAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
				defineAuxiliary:       tt.fields.defineAuxiliary,
				superAuxiliary:        tt.fields.superAuxiliary,
				supplementAuxiliary:   tt.fields.supplementAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers, mapper := createTestInput(t)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	fromAddress2 := "cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"
	fromAccAddress2, err := sdkTypes.AccAddressFromBech32(fromAddress2)
	require.Nil(t, err)
	immutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	immutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIds.NewStringID("ID11"), baseData.NewStringData("ImmutableData")))
	mutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("authentication"), baseData.NewListData(baseLists.NewDataList())))
	mutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIds.NewStringID("authentication"), baseData.NewStringData("MutableData")))
	immutables := baseQualified.NewImmutables(immutableMetaProperties)
	mutables := baseQualified.NewMutables(mutableMetaProperties)
	classificationID := baseIds.NewClassificationID(immutables, mutables)
	identity := baseDocuments.NewIdentity(classificationID, immutables, mutables)
	identity = identity.ProvisionAddress([]sdkTypes.AccAddress{fromAccAddress}...)
	keepers.DefineKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewMappable(identity))
	type fields struct {
		mapper                helpers.Mapper
		authenticateAuxiliary helpers.Auxiliary
		defineAuxiliary       helpers.Auxiliary
		superAuxiliary        helpers.Auxiliary
		supplementAuxiliary   helpers.Auxiliary
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
		{"+ve", fields{mapper, authenticateAuxiliary, defineAuxiliary, superAuxiliary, supplementAuxiliary}, args{context, newMessage(fromAccAddress, baseIds.NewIdentityID(classificationID, immutables), immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)}, newTransactionResponse(nil)},
		{"-ve", fields{mapper, authenticateAuxiliary, defineAuxiliary, superAuxiliary, supplementAuxiliary}, args{context, newMessage(fromAccAddress2, baseIds.NewIdentityID(classificationID, immutables), immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)}, newTransactionResponse(constants.NotAuthorized)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
				defineAuxiliary:       tt.fields.defineAuxiliary,
				superAuxiliary:        tt.fields.superAuxiliary,
				supplementAuxiliary:   tt.fields.supplementAuxiliary,
			}
			if got := transactionKeeper.Transact(tt.args.context, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
