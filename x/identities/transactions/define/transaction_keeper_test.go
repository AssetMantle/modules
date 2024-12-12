// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"context"
	"fmt"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/identities/mapper"
	"github.com/AssetMantle/modules/x/identities/parameters"
	"github.com/AssetMantle/modules/x/identities/record"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/super"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
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
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type TestKeepers struct {
	DefineKeeper helpers.TransactionKeeper
}

var (
	authenticateAuxiliary helpers.Auxiliary
	defineAuxiliary       helpers.Auxiliary
	superAuxiliary        helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	parameterManager      helpers.ParameterManager
)

func createTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper) {

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := mapper.Prototype().Initialize(storeKey)

	parameterManager = parameters.Prototype().Initialize(storeKey)

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, storeTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	Context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	authenticateAuxiliary = authenticate.Auxiliary.Initialize(Mapper, parameterManager)
	defineAuxiliary = define.Auxiliary.Initialize(Mapper, parameterManager)
	superAuxiliary = super.Auxiliary.Initialize(Mapper, parameterManager)
	supplementAuxiliary = supplement.Auxiliary.Initialize(Mapper, parameterManager)
	keepers := TestKeepers{
		DefineKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{authenticateAuxiliary}).(helpers.TransactionKeeper),
	}

	return Context, keepers, Mapper
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
		mapper              helpers.Mapper
		defineAuxiliary     helpers.Auxiliary
		superAuxiliary      helpers.Auxiliary
		supplementAuxiliary helpers.Auxiliary
	}
	type args struct {
		mapper      helpers.Mapper
		in1         helpers.ParameterManager
		auxiliaries []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}{
		{"+ve with nil", fields{}, args{}, transactionKeeper{}},
		{"+ve", fields{Mapper, defineAuxiliary, superAuxiliary, supplementAuxiliary}, args{Mapper, parameterManager, []interface{}{}}, transactionKeeper{Mapper, defineAuxiliary, superAuxiliary, supplementAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:              tt.fields.mapper,
				defineAuxiliary:     tt.fields.defineAuxiliary,
				superAuxiliary:      tt.fields.superAuxiliary,
				supplementAuxiliary: tt.fields.supplementAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	Context, keepers, Mapper := createTestInput(t)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	fromAddress2 := "cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"
	fromAccAddress2, err := sdkTypes.AccAddressFromBech32(fromAddress2)
	require.Nil(t, err)
	immutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	immutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID11"), baseData.NewStringData("ImmutableData")))
	mutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()))
	mutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("authentication"), baseData.NewStringData("MutableData")))
	immutables := baseQualified.NewImmutables(immutableMetaProperties)
	mutables := baseQualified.NewMutables(mutableMetaProperties)
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	identity := baseDocuments.NewIdentity(classificationID, immutables, mutables)
	identity = identity.ProvisionAddress([]sdkTypes.AccAddress{fromAccAddress}...)
	keepers.DefineKeeper.(transactionKeeper).mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(record.NewRecord(identity))

	type fields struct {
		mapper              helpers.Mapper
		defineAuxiliary     helpers.Auxiliary
		superAuxiliary      helpers.Auxiliary
		supplementAuxiliary helpers.Auxiliary
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
		{"+ve", fields{Mapper, defineAuxiliary, superAuxiliary, supplementAuxiliary}, args{Context.Context(), NewMessage(fromAccAddress, baseIDs.NewIdentityID(classificationID, immutables), immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).(*Message)}, newTransactionResponse(classificationID), false},
		{"-ve", fields{Mapper, defineAuxiliary, superAuxiliary, supplementAuxiliary}, args{Context.Context(), NewMessage(fromAccAddress2, baseIDs.NewIdentityID(classificationID, immutables), immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).(*Message)}, newTransactionResponse(nil), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:              tt.fields.mapper,
				defineAuxiliary:     tt.fields.defineAuxiliary,
				superAuxiliary:      tt.fields.superAuxiliary,
				supplementAuxiliary: tt.fields.supplementAuxiliary,
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
