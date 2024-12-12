// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package take

import (
	"context"
	"fmt"
	"github.com/AssetMantle/modules/x/orders/mapper"
	"github.com/AssetMantle/modules/x/orders/record"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	baseTypes "github.com/AssetMantle/schema/types/base"
	tendermintDB "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/orders/parameters"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
)

var (
	supplementAuxiliary   helpers.Auxiliary
	transferAuxiliary     helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
)

type TestKeepers struct {
	TakeKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (types.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {

	storeKey := types.NewKVStoreKey("test")
	paramsStoreKey := types.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := types.NewTransientStoreKey("testParamsTransient")
	Mapper := mapper.Prototype().Initialize(storeKey)

	parameterManager := parameters.Prototype().Initialize(storeKey)

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, storeTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	authenticateAuxiliary = authenticate.Auxiliary.Initialize(Mapper, parameterManager)
	supplementAuxiliary = supplement.Auxiliary.Initialize(Mapper, parameterManager)
	transferAuxiliary = transfer.Auxiliary.Initialize(Mapper, parameterManager)

	Context := types.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		TakeKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.TransactionKeeper),
	}

	return Context, keepers, Mapper, parameterManager
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
	_, _, Mapper, parameterManager := CreateTestInput(t)
	type fields struct {
		mapper                helpers.Mapper
		parameterManager      helpers.ParameterManager
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
		{"+ve", fields{Mapper, parameterManager, supplementAuxiliary, transferAuxiliary, authenticateAuxiliary}, args{Mapper, parameterManager, []interface{}{}}, transactionKeeper{Mapper, parameterManager, authenticateAuxiliary, supplementAuxiliary, transferAuxiliary, authenticateAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameterManager:      tt.fields.parameterManager,
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

func Test_transactionKeeper_Transact1(t *testing.T) {
	Context, keepers, Mapper, parameterManager := CreateTestInput(t)
	mutableMetaProperties := baseLists.NewPropertyList(
		baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("exchangeRate"), baseData.NewDecData(types.NewDec(10))),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("makerAssetID"), baseData.NewIDData(baseDocuments.NewCoinAsset("makerID").GetCoinAssetID())),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("creationHeight"), baseData.NewHeightData(baseTypes.NewHeight(1))),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("takerAssetID"), baseData.NewIDData(baseDocuments.NewCoinAsset("takerID").GetCoinAssetID())),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("makerID"), baseData.NewIDData(baseIDs.PrototypeIdentityID())),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("takerID"), baseData.NewIDData(baseIDs.PrototypeIdentityID())),
	)
	immutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData()))
	immutablesMeta := baseQualified.NewImmutables(immutableMetaProperties)
	mutablesMeta := baseQualified.NewMutables(mutableMetaProperties)
	testClassificationID := baseIDs.NewClassificationID(immutablesMeta, mutablesMeta)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutablesMeta)
	testFromID2 := baseIDs.PrototypeIdentityID()
	mutableMetaProperties.Mutate(
		baseProperties.NewMetaProperty(baseIDs.NewStringID("makerID"), baseData.NewIDData(testFromID)),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("takerID"), baseData.NewIDData(testFromID)))
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := types.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	testIdentity := baseDocuments.NewIdentity(testClassificationID, immutablesMeta, mutablesMeta)
	testIdentity.ProvisionAddress([]types.AccAddress{fromAccAddress}...)
	testOrder := baseDocuments.NewOrder(testClassificationID, immutablesMeta, mutablesMeta)
	// testMakerAssetID := baseDocuments.NewCoinAsset("makerID"))
	// testTakerAssetID := baseDocuments.NewCoinAsset("takerID"))
	testRate := types.NewInt(10)
	// testHeight := baseTypes.NewHeight(1)
	testOrderID := baseIDs.NewOrderID(testClassificationID, immutablesMeta)
	testOrderID2 := baseIDs.NewOrderID(testClassificationID, immutablesMeta)
	keepers.TakeKeeper.(transactionKeeper).mapper.NewCollection(types.WrapSDKContext(Context)).Add(record.NewRecord(testOrder))
	type fields struct {
		mapper                helpers.Mapper
		parameterManager      helpers.ParameterManager
		authenticateAuxiliary helpers.Auxiliary
		burnAuxiliary         helpers.Auxiliary
		supplementAuxiliary   helpers.Auxiliary
		transferAuxiliary     helpers.Auxiliary
	}
	type args struct {
		context context.Context
		message helpers.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.TransactionResponse
		wantErr bool
	}{
		{"+ve Not Authorized", fields{Mapper, parameterManager, authenticateAuxiliary, supplementAuxiliary, transferAuxiliary, authenticateAuxiliary}, args{Context.Context(), NewMessage(fromAccAddress, testFromID2, testRate, testOrderID).(*Message)}, newTransactionResponse(), false},
		{"+ve", fields{Mapper, parameterManager, authenticateAuxiliary, supplementAuxiliary, transferAuxiliary, authenticateAuxiliary}, args{Context.Context(), NewMessage(fromAccAddress, testFromID, testRate, testOrderID).(*Message)}, newTransactionResponse(), false},
		{"+ve Entity Not Found", fields{Mapper, parameterManager, authenticateAuxiliary, supplementAuxiliary, transferAuxiliary, authenticateAuxiliary}, args{Context.Context(), NewMessage(fromAccAddress, testFromID, testRate, testOrderID2).(*Message)}, newTransactionResponse(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameterManager:      tt.fields.parameterManager,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
				burnAuxiliary:         tt.fields.burnAuxiliary,
				supplementAuxiliary:   tt.fields.supplementAuxiliary,
				transferAuxiliary:     tt.fields.transferAuxiliary,
			}
			got, err := transactionKeeper.Transact(tt.args.context, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() got = %v, want %v", got, tt.want)
			}
		})
	}
}
