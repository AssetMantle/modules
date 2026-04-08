// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"fmt"
	"github.com/AssetMantle/modules/x/orders/mapper"
	storeTypes "cosmossdk.io/store/types"
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/helpers/base/testutil"

	"github.com/stretchr/testify/mock"

	baseData "github.com/AssetMantle/schema/data/base"
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
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/super"
	"github.com/AssetMantle/modules/x/orders/parameters"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

var (
	defineAuxiliary       helpers.Auxiliary
	superAuxiliary        helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
)

type TestKeepers struct {
	DefineKeeper helpers.TransactionKeeper
}

func createTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {

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
	defineAuxiliary = define.Auxiliary.Initialize(Mapper, parameterManager, bankKeeper.BaseKeeper{}, &stakingKeeper.Keeper{})
	superAuxiliary = super.Auxiliary.Initialize(Mapper, parameterManager)

	Context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	parameterManager, _ = parameterManager.Set().Update(sdkTypes.WrapSDKContext(Context))

	keepers := TestKeepers{
		DefineKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{defineAuxiliary, superAuxiliary, authenticateAuxiliary}).(helpers.TransactionKeeper),
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
	_, _, Mapper, parameterManger := createTestInput(t)
	type fields struct {
		mapper                helpers.Mapper
		parameterManager      helpers.ParameterManager
		defineAuxiliary       helpers.Auxiliary
		superAuxiliary        helpers.Auxiliary
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
		{"valid", fields{Mapper, parameterManger, defineAuxiliary, superAuxiliary, authenticateAuxiliary}, args{Mapper, parameterManger, []interface{}{}}, transactionKeeper{Mapper, parameterManger, defineAuxiliary, superAuxiliary, authenticateAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameterManager:      tt.fields.parameterManager,
				defineAuxiliary:       tt.fields.defineAuxiliary,
				superAuxiliary:        tt.fields.superAuxiliary,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.parameterManager, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	Context, _, Mapper, parameterManager := createTestInput(t)

	authAux, authAuxKeeper := testutil.NewMockAuxiliaryPair()
	authAuxKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	defAux, defAuxKeeper := testutil.NewMockAuxiliaryPair()
	defAuxKeeper.On("Help", mock.Anything, mock.Anything).Return(define.NewAuxiliaryResponse(baseIDs.PrototypeClassificationID()), nil)
	supAux, supAuxKeeper := testutil.NewMockAuxiliaryPair()
	supAuxKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)

	fromAccAddress, err := sdkTypes.AccAddressFromBech32("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	require.NoError(t, err)
	immutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData()))
	immutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData()))
	mutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()))
	mutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()))
	immutablesMeta := baseQualified.NewImmutables(immutableMetaProperties)
	testClassificationID := baseIDs.NewClassificationID(immutablesMeta, baseQualified.NewMutables(mutableMetaProperties))
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutablesMeta)

	tk := transactionKeeper{mapper: Mapper, parameterManager: parameterManager, defineAuxiliary: defAux, superAuxiliary: supAux, authenticateAuxiliary: authAux}

	t.Run("valid", func(t *testing.T) {
		got, err := tk.Transact(sdkTypes.WrapSDKContext(Context), NewMessage(fromAccAddress, testFromID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).(*Message))
		require.NoError(t, err)
		require.NotNil(t, got)
	})
}
