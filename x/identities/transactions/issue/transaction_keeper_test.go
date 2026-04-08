// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

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
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/bond"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/identities/parameters"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/authorize"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

type TestKeepers struct {
	IssueKeeper helpers.TransactionKeeper
}

var (
	parameterManager      helpers.ParameterManager
	authenticateAuxiliary helpers.Auxiliary
	bondAuxiliary         helpers.Auxiliary
	conformAuxiliary      helpers.Auxiliary
	authorizeAuxiliary    helpers.Auxiliary
)

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {

	storeKey := storeTypes.NewKVStoreKey("test")
	paramsStoreKey := storeTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := storeTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := mapper.Prototype().Initialize(storeKey)

	parameterManager = parameters.Prototype().Initialize(storeKey)

	memDB := cosmosDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB, log.NewNopLogger(), storeMetrics.NewNoOpMetrics())
	commitMultiStore.MountStoreWithDB(storeKey, storeTypes.StoreTypeIAVL, nil)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, storeTypes.StoreTypeIAVL, nil)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, storeTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	Context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	parameterManager, _ = parameterManager.Set().Update(sdkTypes.WrapSDKContext(Context))

	authenticateAuxiliary = authenticate.Auxiliary.Initialize(Mapper, parameterManager)
	bondAuxiliary = bond.Auxiliary.Initialize(Mapper, parameterManager, bankKeeper.BaseKeeper{}, &stakingKeeper.Keeper{})
	conformAuxiliary = conform.Auxiliary.Initialize(Mapper, parameterManager)
	authorizeAuxiliary = authorize.Auxiliary.Initialize(Mapper, parameterManager)
	keepers := TestKeepers{
		IssueKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{conformAuxiliary, bondAuxiliary, authorizeAuxiliary}).(helpers.TransactionKeeper),
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
		mapper             helpers.Mapper
		parameterManager   helpers.ParameterManager
		conformAuxiliary   helpers.Auxiliary
		bondAuxiliary      helpers.Auxiliary
		authorizeAuxiliary helpers.Auxiliary
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
		{"valid", fields{Mapper, parameterManager, conformAuxiliary, bondAuxiliary, authorizeAuxiliary}, args{Mapper, parameterManager, []interface{}{}}, transactionKeeper{Mapper, parameterManager, conformAuxiliary, bondAuxiliary, authorizeAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:             tt.fields.mapper,
				parameterManager:   tt.fields.parameterManager,
				conformAuxiliary:   tt.fields.conformAuxiliary,
				bondAuxiliary:      tt.fields.bondAuxiliary,
				authorizeAuxiliary: tt.fields.authorizeAuxiliary,
			}
			got := transactionKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.auxiliaries)
			assert.NotNil(t, got)
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	Context, _, Mapper, parameterManager := CreateTestInput(t)

	confAux, confAuxKeeper := testutil.NewMockAuxiliaryPair()
	confAuxKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	bondAux, bondAuxKeeper := testutil.NewMockAuxiliaryPair()
	bondAuxKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	authzAux, authzAuxKeeper := testutil.NewMockAuxiliaryPair()
	authzAuxKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)

	fromAccAddress, err := sdkTypes.AccAddressFromBech32("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	require.NoError(t, err)
	immutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	immutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID11"), baseData.NewStringData("ImmutableData")))
	mutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()))
	mutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("authentication"), baseData.NewStringData("MutableData")))
	immutables := baseQualified.NewImmutables(immutableMetaProperties)
	mutables := baseQualified.NewMutables(mutableMetaProperties)
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	fromIdentityID := baseIDs.NewIdentityID(classificationID, immutables)
	identity := baseDocuments.NewIdentity(classificationID, immutables, mutables)
	identity = identity.ProvisionAddress([]sdkTypes.AccAddress{fromAccAddress}...)

	tk := transactionKeeper{mapper: Mapper, parameterManager: parameterManager, conformAuxiliary: confAux, bondAuxiliary: bondAux, authorizeAuxiliary: authzAux}
	tk.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(record.NewRecord(identity))

	t.Run("authorize succeeds but bond amount not revealed", func(t *testing.T) {
		_, err := tk.Transact(sdkTypes.WrapSDKContext(Context), NewMessage(fromAccAddress, fromIdentityID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).(*Message))
		// The identity has no revealed bond amount, so this fails with meta data error
		require.Error(t, err)
	})
}
