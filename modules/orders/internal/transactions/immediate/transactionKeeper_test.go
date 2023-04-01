// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package immediate

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/mappable"
	"github.com/AssetMantle/modules/modules/orders/internal/parameters"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/transfer"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseDocuments "github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

var (
	conformAuxiliary      helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	transferAuxiliary     helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
)

type TestKeepers struct {
	ImmediateKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (types.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {
	var legacyAmino = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	storeKey := types.NewKVStoreKey("test")
	paramsStoreKey := types.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := types.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	encodingConfig := simapp.MakeTestEncodingConfig()
	appCodec := encodingConfig.Marshaler
	ParamsKeeper := paramsKeeper.NewKeeper(
		appCodec,
		legacyAmino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	parameterManager := parameters.Prototype().Initialize(ParamsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, types.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, types.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, types.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	authenticateAuxiliary = authenticate.Auxiliary.Initialize(Mapper, parameterManager)
	supplementAuxiliary = supplement.Auxiliary.Initialize(Mapper, parameterManager)
	transferAuxiliary = transfer.Auxiliary.Initialize(Mapper, parameterManager)
	conformAuxiliary = conform.Auxiliary.Initialize(Mapper, parameterManager)

	context := types.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		ImmediateKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.TransactionKeeper),
	}

	return context, keepers, Mapper, parameterManager
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
	_, _, mapper, parameterManager := CreateTestInput(t)
	type fields struct {
		mapper                helpers.Mapper
		parameterManager      helpers.ParameterManager
		conformAuxiliary      helpers.Auxiliary
		supplementAuxiliary   helpers.Auxiliary
		transferAuxiliary     helpers.Auxiliary
		authenticateAuxiliary helpers.Auxiliary
	}
	type args struct {
		mapper           helpers.Mapper
		parameterManager helpers.ParameterManager
		auxiliaries      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}{
		{"+ve with nil", fields{}, args{}, transactionKeeper{}},
		{"+ve", fields{mapper, parameterManager, conformAuxiliary, supplementAuxiliary, transferAuxiliary, authenticateAuxiliary}, args{mapper, parameterManager, []interface{}{}}, transactionKeeper{mapper, parameterManager, conformAuxiliary, supplementAuxiliary, transferAuxiliary, authenticateAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameterManager:      tt.fields.parameterManager,
				conformAuxiliary:      tt.fields.conformAuxiliary,
				supplementAuxiliary:   tt.fields.supplementAuxiliary,
				transferAuxiliary:     tt.fields.transferAuxiliary,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.parameterManager, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers, mapper, parameterManager := CreateTestInput(t)
	mutableMetaProperties := baseLists.NewPropertyList(
		baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("exchangeRate"), baseData.NewDecData(types.NewDec(10))),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("makerOwnableID"), baseData.NewIDData(baseIDs.NewCoinID(baseIDs.NewStringID("makerID")))),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("creationHeight"), baseData.NewHeightData(baseTypes.NewHeight(1))),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("takerOwnableID"), baseData.NewIDData(baseIDs.NewCoinID(baseIDs.NewStringID("takerID")))),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("makerID"), baseData.NewIDData(baseIDs.PrototypeIdentityID())),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("makerID"), baseData.NewIDData(baseIDs.PrototypeIdentityID())),
	)
	mutableProperties := baseLists.NewPropertyList(
		baseProperties.NewMesaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()),
	)
	immutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData()))
	immutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData()))
	immutablesMeta := baseQualified.NewImmutables(immutableMetaProperties)
	mutablesMeta := baseQualified.NewMutables(mutableMetaProperties)
	testClassificationID := baseIDs.NewClassificationID(immutablesMeta, mutablesMeta)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutablesMeta)
	testTakerID := baseIDs.PrototypeIdentityID()
	mutableMetaProperties.Mutate(
		baseProperties.NewMetaProperty(baseIDs.NewStringID("makerID"), baseData.NewIDData(testFromID)),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("makerID"), baseData.NewIDData(testFromID)))
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := types.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	testIdentity := baseDocuments.NewIdentity(testClassificationID, immutablesMeta, mutablesMeta)
	testIdentity.ProvisionAddress([]types.AccAddress{fromAccAddress}...)
	testOrder := baseDocuments.NewOrder(testClassificationID, immutablesMeta, mutablesMeta)
	// testOrderID := mappable.NewMappable(testOrder).GetKey()
	testMakerOwnableID := baseIDs.NewCoinID(baseIDs.NewStringID("makerID"))
	testTakerOwnableID := baseIDs.NewCoinID(baseIDs.NewStringID("takerID"))
	testRate := types.NewDec(10)
	testHeight := baseTypes.NewHeight(1)
	// testOrderID := baseIDs.NewOrderID(testClassificationID, testMakerOwnableID, testTakerOwnableID, testRate, testHeight, testFromID, immutablesMeta)
	// testOrderID2 := baseIDs.NewOrderID(testClassificationID, testTakerOwnableID, testTakerOwnableID, testRate, testHeight, testFromID, immutablesMeta)
	keepers.ImmediateKeeper.(transactionKeeper).mapper.NewCollection(types.WrapSDKContext(context)).Add(mappable.NewMappable(testOrder))
	type fields struct {
		mapper                helpers.Mapper
		parameterManager      helpers.ParameterManager
		conformAuxiliary      helpers.Auxiliary
		supplementAuxiliary   helpers.Auxiliary
		transferAuxiliary     helpers.Auxiliary
		authenticateAuxiliary helpers.Auxiliary
	}
	type args struct {
		context types.Context
		msg     helpers.Message
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.TransactionResponse
	}{
		{"+ve", fields{mapper, parameterManager, conformAuxiliary, supplementAuxiliary, transferAuxiliary, authenticateAuxiliary}, args{context, newMessage(fromAccAddress, testFromID, testClassificationID, testTakerID, testMakerOwnableID, testTakerOwnableID, testHeight, testRate, testRate, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).(*Message)}, newTransactionResponse(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameterManager:      tt.fields.parameterManager,
				conformAuxiliary:      tt.fields.conformAuxiliary,
				supplementAuxiliary:   tt.fields.supplementAuxiliary,
				transferAuxiliary:     tt.fields.transferAuxiliary,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
			}
			if got := transactionKeeper.Transact(types.WrapSDKContext(context), tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
