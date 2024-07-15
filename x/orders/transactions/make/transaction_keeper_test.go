// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	"context"
	"fmt"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/bond"
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
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/authorize"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/orders/parameters"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
)

var (
	authenticateAuxiliary helpers.Auxiliary
	authorizeAuxiliary    helpers.Auxiliary
	bondAuxiliary         helpers.Auxiliary
	conformAuxiliary      helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	transferAuxiliary     helpers.Auxiliary
)

type TestKeepers struct {
	MakeKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {
	var legacyAmino = baseHelpers.CodecPrototype().GetLegacyAmino()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := mapper.Prototype().Initialize(storeKey)
	codec := baseHelpers.TestCodec()
	ParamsKeeper := paramsKeeper.NewKeeper(
		codec,
		legacyAmino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	parameterManager := parameters.Prototype().Initialize(ParamsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, storeTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	authenticateAuxiliary = authenticate.Auxiliary.Initialize(Mapper, parameterManager)
	authorizeAuxiliary = authorize.Auxiliary.Initialize(Mapper, parameterManager)
	bondAuxiliary = bond.Auxiliary.Initialize(Mapper, parameterManager)
	conformAuxiliary = conform.Auxiliary.Initialize(Mapper, parameterManager)
	supplementAuxiliary = supplement.Auxiliary.Initialize(Mapper, parameterManager)
	transferAuxiliary = transfer.Auxiliary.Initialize(Mapper, parameterManager)

	Context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		MakeKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.TransactionKeeper),
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
		conformAuxiliary      helpers.Auxiliary
		supplementAuxiliary   helpers.Auxiliary
		transferAuxiliary     helpers.Auxiliary
		authenticateAuxiliary helpers.Auxiliary
		authorizeAuxiliary    helpers.Auxiliary
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
		{"+ve", fields{Mapper, parameterManager, conformAuxiliary, supplementAuxiliary, transferAuxiliary, authenticateAuxiliary, authorizeAuxiliary}, args{Mapper, parameterManager, []interface{}{}}, transactionKeeper{Mapper, parameterManager, authenticateAuxiliary, conformAuxiliary, supplementAuxiliary, transferAuxiliary, authenticateAuxiliary, authorizeAuxiliary}},
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
				authorizeAuxiliary:    tt.fields.authorizeAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.parameterManager, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	Context, keepers, Mapper, parameterManager := CreateTestInput(t)
	mutableMetaProperties := baseLists.NewPropertyList(
		baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("exchangeRate"), baseData.NewDecData(sdkTypes.NewDec(10))),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("makerAssetID"), baseData.NewIDData(baseDocuments.NewCoinAsset("makerID").GetCoinAssetID())),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("creationHeight"), baseData.NewHeightData(baseTypes.NewHeight(1))),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("takerAssetID"), baseData.NewIDData(baseDocuments.NewCoinAsset("takerID").GetCoinAssetID())),
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
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	testIdentity := baseDocuments.NewIdentity(testClassificationID, immutablesMeta, mutablesMeta)
	testIdentity.ProvisionAddress([]sdkTypes.AccAddress{fromAccAddress}...)
	testOrder := baseDocuments.NewOrder(testClassificationID, immutablesMeta, mutablesMeta)
	// testOrderID := mappable.NewMappable(testOrder).GenerateKey()
	testMakerAssetID := baseDocuments.NewCoinAsset("makerID").GetCoinAssetID()
	testTakerAssetID := baseDocuments.NewCoinAsset("takerID").GetCoinAssetID()
	testRate := sdkTypes.NewInt(10)
	testHeight := baseTypes.NewHeight(1)
	// testOrderID := baseIDs.NewOrderID(testClassificationID, testMakerAssetID, testTakerAssetID, testRate, testHeight, testFromID, immutablesMeta)
	// testOrderID2 := baseIDs.NewOrderID(testClassificationID, testTakerAssetID, testTakerAssetID, testRate, testHeight, testFromID, immutablesMeta)
	keepers.MakeKeeper.(transactionKeeper).mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(record.NewRecord(testOrder))

	type fields struct {
		mapper                helpers.Mapper
		parameterManager      helpers.ParameterManager
		authenticateAuxiliary helpers.Auxiliary
		authorizeAuxiliary    helpers.Auxiliary
		bondAuxiliary         helpers.Auxiliary
		conformAuxiliary      helpers.Auxiliary
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
		{"+ve", fields{Mapper, parameterManager, authenticateAuxiliary, authorizeAuxiliary, bondAuxiliary, conformAuxiliary, supplementAuxiliary, transferAuxiliary}, args{Context.Context(), NewMessage(fromAccAddress, testFromID, testClassificationID, testTakerID, testMakerAssetID, testTakerAssetID, testHeight, testRate, testRate, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).(*Message)}, newTransactionResponse(nil), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameterManager:      tt.fields.parameterManager,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
				authorizeAuxiliary:    tt.fields.authorizeAuxiliary,
				bondAuxiliary:         tt.fields.bondAuxiliary,
				conformAuxiliary:      tt.fields.conformAuxiliary,
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
