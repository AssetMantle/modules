// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

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
	tendermintDB "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/deputize"
	"github.com/AssetMantle/modules/x/orders/parameters"
)

var (
	authenticateAuxiliary helpers.Auxiliary
	deputizeAuxiliary     helpers.Auxiliary
)

type TestKeepers struct {
	DeputizeKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
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
	deputizeAuxiliary = deputize.Auxiliary.Initialize(Mapper, parameterManager)

	Context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		DeputizeKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.TransactionKeeper),
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
		authenticateAuxiliary helpers.Auxiliary
		deputizeAuxiliary     helpers.Auxiliary
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
		{"+ve", fields{Mapper, parameterManager, authenticateAuxiliary, deputizeAuxiliary}, args{Mapper, parameterManager, []interface{}{}}, transactionKeeper{Mapper, parameterManager, authenticateAuxiliary, deputizeAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameterManager:      tt.fields.parameterManager,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
				deputizeAuxiliary:     tt.fields.deputizeAuxiliary,
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
	)
	immutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData()))
	immutablesMeta := baseQualified.NewImmutables(immutableMetaProperties)
	mutablesMeta := baseQualified.NewMutables(mutableMetaProperties)
	testClassificationID := baseIDs.NewClassificationID(immutablesMeta, mutablesMeta)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutablesMeta)
	mutableMetaProperties.Mutate(
		baseProperties.NewMetaProperty(baseIDs.NewStringID("makerID"), baseData.NewIDData(testFromID)),
		baseProperties.NewMetaProperty(baseIDs.NewStringID("makerID"), baseData.NewIDData(testFromID)))
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	testIdentity := baseDocuments.NewIdentity(testClassificationID, immutablesMeta, mutablesMeta)
	testIdentity.ProvisionAddress([]sdkTypes.AccAddress{fromAccAddress}...)
	testOrder := baseDocuments.NewOrder(testClassificationID, immutablesMeta, mutablesMeta)
	keepers.DeputizeKeeper.(transactionKeeper).mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(record.NewRecord(testOrder))
	type fields struct {
		mapper                helpers.Mapper
		parameterManager      helpers.ParameterManager
		authenticateAuxiliary helpers.Auxiliary
		deputizeAuxiliary     helpers.Auxiliary
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
		{"+ve", fields{Mapper, parameterManager, authenticateAuxiliary, deputizeAuxiliary}, args{Context.Context(), NewMessage(fromAccAddress, testFromID, testFromID, testClassificationID, mutableMetaProperties, true, true, true, true, true).(*Message)}, newTransactionResponse(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameterManager:      tt.fields.parameterManager,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
				deputizeAuxiliary:     tt.fields.deputizeAuxiliary,
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
