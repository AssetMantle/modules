// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package quash

import (
	"fmt"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/unbond"
	"github.com/AssetMantle/modules/x/identities/mapper"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/AssetMantle/modules/x/identities/record"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/authorize"
	storeTypes "cosmossdk.io/store/types"
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/helpers/base/testutil"

	"github.com/stretchr/testify/mock"

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
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/identities/parameters"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var (
	authorizeAuxiliary  helpers.Auxiliary
	supplementAuxiliary helpers.Auxiliary
	unbondAuxiliary     helpers.Auxiliary
)

type TestKeepers struct {
	QuashKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (types.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {

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

	authorizeAuxiliary = authorize.Auxiliary.Initialize(Mapper, parameterManager)
	supplementAuxiliary = supplement.Auxiliary.Initialize(Mapper, parameterManager)
	unbondAuxiliary = unbond.Auxiliary.Initialize(Mapper, parameterManager, bankKeeper.BaseKeeper{}, &stakingKeeper.Keeper{})

	Context := types.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	parameterManager, _ = parameterManager.Set().Update(sdkTypes.WrapSDKContext(Context))

	keepers := TestKeepers{
		QuashKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{authorizeAuxiliary, supplementAuxiliary, unbondAuxiliary}).(helpers.TransactionKeeper),
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
		mapper              helpers.Mapper
		authorizeAuxiliary  helpers.Auxiliary
		supplementAuxiliary helpers.Auxiliary
		unbondAuxiliary     helpers.Auxiliary
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
		{"valid", fields{Mapper, authorizeAuxiliary, supplementAuxiliary, unbondAuxiliary}, args{Mapper, parameterManager, []interface{}{}}, transactionKeeper{Mapper, parameterManager, authorizeAuxiliary, supplementAuxiliary, unbondAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:              tt.fields.mapper,
				authorizeAuxiliary:  tt.fields.authorizeAuxiliary,
				supplementAuxiliary: tt.fields.supplementAuxiliary,
				unbondAuxiliary:     tt.fields.unbondAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	Context, _, Mapper, parameterManager := CreateTestInput(t)

	authzAux, authzAuxKeeper := testutil.NewMockAuxiliaryPair()
	authzAuxKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	supplAux, supplAuxKeeper := testutil.NewMockAuxiliaryPair()
	supplAuxKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	unbondAux, unbondAuxKeeper := testutil.NewMockAuxiliaryPair()
	unbondAuxKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)

	fromAccAddress, err := types.AccAddressFromBech32("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	require.NoError(t, err)
	mutableProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()))
	immutableProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData()))
	immutables := baseQualified.NewImmutables(immutableProperties)
	mutables := baseQualified.NewMutables(mutableProperties)
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	testIdentity := baseDocuments.NewIdentity(testClassificationID, immutables, mutables)
	testIdentity = testIdentity.ProvisionAddress([]types.AccAddress{fromAccAddress}...)

	tk := transactionKeeper{mapper: Mapper, parameterManager: parameterManager, authorizeAuxiliary: authzAux, supplementAuxiliary: supplAux, unbondAuxiliary: unbondAux}
	tk.mapper.NewCollection(types.WrapSDKContext(Context)).Add(record.NewRecord(testIdentity))

	t.Run("authorize succeeds but bond amount not revealed", func(t *testing.T) {
		_, err := tk.Transact(sdkTypes.WrapSDKContext(Context), NewMessage(fromAccAddress, testFromID, testFromID).(*Message))
		require.Error(t, err)
	})
}
