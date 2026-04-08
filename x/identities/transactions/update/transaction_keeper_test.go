// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package update

import (
	"github.com/AssetMantle/modules/x/identities/mapper"
	"github.com/AssetMantle/modules/x/identities/record"
	storeTypes "cosmossdk.io/store/types"
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
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/member"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/identities/parameters"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/maintain"
)

var (
	authenticateAuxiliary helpers.Auxiliary
	maintainAuxiliary     helpers.Auxiliary
	memberAuxiliary       helpers.Auxiliary
)

type TestKeepers struct {
	MutateKeeper helpers.TransactionKeeper
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
	maintainAuxiliary = maintain.Auxiliary.Initialize(Mapper, parameterManager)
	memberAuxiliary = member.Auxiliary.Initialize(Mapper, parameterManager)

	Context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	parameterManager, _ = parameterManager.Set().Update(sdkTypes.WrapSDKContext(Context))

	keepers := TestKeepers{
		MutateKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{maintainAuxiliary, memberAuxiliary}).(helpers.TransactionKeeper),
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
			got := keeperPrototype()
			assert.Equal(t, tt.want, got, "keeperPrototype()")
		})
	}
}

func Test_transactionKeeper_Initialize(t *testing.T) {
	_, _, Mapper, parameterManager := CreateTestInput(t)
	type fields struct {
		mapper                helpers.Mapper
		authenticateAuxiliary helpers.Auxiliary
		maintainAuxiliary     helpers.Auxiliary
		memberAuxiliary       helpers.Auxiliary
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
		{"valid", fields{Mapper, authenticateAuxiliary, maintainAuxiliary, memberAuxiliary}, args{Mapper, parameterManager, []interface{}{}}, transactionKeeper{Mapper, maintainAuxiliary, memberAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:            tt.fields.mapper,
				maintainAuxiliary: tt.fields.maintainAuxiliary,
				memberAuxiliary:   tt.fields.memberAuxiliary,
			}
			got := transactionKeeper.Initialize(tt.args.mapper, tt.args.parameterManager, tt.args.auxiliaries)
			assert.NotNil(t, got)
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	Context, _, Mapper, _ := CreateTestInput(t)

	maintAux, maintAuxKeeper := testutil.NewMockAuxiliaryPair()
	maintAuxKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	memAux, memAuxKeeper := testutil.NewMockAuxiliaryPair()
	memAuxKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)

	mutableProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()))
	immutableProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData()))
	toMutateMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData(baseData.NewStringData("Test"))))
	immutables := baseQualified.NewImmutables(immutableProperties)
	mutables := baseQualified.NewMutables(mutableProperties)
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	fromAccAddress, err := sdkTypes.AccAddressFromBech32("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	require.NoError(t, err)
	testIdentity := baseDocuments.NewIdentity(testClassificationID, immutables, mutables)
	testIdentity = testIdentity.ProvisionAddress([]sdkTypes.AccAddress{fromAccAddress}...)

	tk := transactionKeeper{mapper: Mapper, maintainAuxiliary: maintAux, memberAuxiliary: memAux}
	tk.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(record.NewRecord(testIdentity))

	t.Run("valid", func(t *testing.T) {
		got, err := tk.Transact(sdkTypes.WrapSDKContext(Context), NewMessage(fromAccAddress, testFromID, testFromID, mutableProperties, toMutateMetaProperties).(*Message))
		require.NoError(t, err)
		require.NotNil(t, got)
	})
}
