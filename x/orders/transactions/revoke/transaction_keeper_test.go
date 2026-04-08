// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"context"
	"fmt"
	"github.com/AssetMantle/modules/x/orders/mapper"
	"github.com/AssetMantle/modules/x/orders/record"
	storeTypes "cosmossdk.io/store/types"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	cosmosDB "github.com/cosmos/cosmos-db"
	"cosmossdk.io/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"cosmossdk.io/store"
	storeMetrics "cosmossdk.io/store/metrics"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/revoke"
	"github.com/AssetMantle/modules/x/orders/parameters"
)

var (
	authenticateAuxiliary helpers.Auxiliary
	revokeAuxiliary       helpers.Auxiliary
)

type TestKeepers struct {
	RevokeKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {

	storeKey := storeTypes.NewKVStoreKey("test")
	paramsStoreKey := storeTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := storeTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := mapper.Prototype().Initialize(storeKey)

	parameterManager := parameters.Prototype().Initialize(storeKey)

	memDB := cosmosDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB, log.NewNopLogger(), storeMetrics.NewNoOpMetrics())
	commitMultiStore.MountStoreWithDB(storeKey, storeTypes.StoreTypeIAVL, nil)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, storeTypes.StoreTypeIAVL, nil)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, storeTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	authenticateAuxiliary = authenticate.Auxiliary.Initialize(Mapper, parameterManager)
	revokeAuxiliary = revoke.Auxiliary.Initialize(Mapper, parameterManager)

	Context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	parameterManager, _ = parameterManager.Set().Update(sdkTypes.WrapSDKContext(Context))

	keepers := TestKeepers{
		RevokeKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{authenticateAuxiliary, revokeAuxiliary}).(helpers.TransactionKeeper),
	}

	return Context, keepers, Mapper, parameterManager
}

func Test_keeperPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.TransactionKeeper
	}{
		{"valid", transactionKeeper{}},
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
		revokeAuxiliary       helpers.Auxiliary
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
		{"valid", fields{Mapper, parameterManager, authenticateAuxiliary, revokeAuxiliary}, args{Mapper, parameterManager, []interface{}{authenticateAuxiliary, revokeAuxiliary}}, transactionKeeper{Mapper, parameterManager, authenticateAuxiliary, revokeAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameterManager:      tt.fields.parameterManager,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
				revokeAuxiliary:       tt.fields.revokeAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.parameterManager, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact1(t *testing.T) {
	t.Skip("test infrastructure shares single store/mapper across modules")
	Context, keepers, Mapper, parameterManager := CreateTestInput(t)
	mutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()))
	immutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData()))
	immutablesMeta := baseQualified.NewImmutables(immutableMetaProperties)
	mutablesMeta := baseQualified.NewMutables(mutableMetaProperties)
	testIdentity := baseDocuments.NewIdentity(testClassificationID, immutablesMeta, mutablesMeta)
	testIdentity.ProvisionAddress([]sdkTypes.AccAddress{fromAccAddress}...)
	testOrder := baseDocuments.NewOrder(testClassificationID, immutablesMeta, mutablesMeta)
	keepers.RevokeKeeper.(transactionKeeper).mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(record.NewRecord(testOrder))
	type fields struct {
		mapper                helpers.Mapper
		parameterManager      helpers.ParameterManager
		authenticateAuxiliary helpers.Auxiliary
		revokeAuxiliary       helpers.Auxiliary
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
		{"valid", fields{Mapper, parameterManager, authenticateAuxiliary, revokeAuxiliary}, args{sdkTypes.WrapSDKContext(Context), NewMessage(fromAccAddress, testFromID, testFromID, testClassificationID).(*Message)}, newTransactionResponse(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameterManager:      tt.fields.parameterManager,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
				revokeAuxiliary:       tt.fields.revokeAuxiliary,
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
