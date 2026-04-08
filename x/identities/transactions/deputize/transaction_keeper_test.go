// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/identities/mapper"
	"github.com/AssetMantle/modules/x/identities/parameters"
	"github.com/AssetMantle/modules/x/identities/record"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/member"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/deputize"
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
	storeTypes "cosmossdk.io/store/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/AssetMantle/modules/helpers/base/testutil"

	"github.com/stretchr/testify/mock"
)

type TestKeepers struct {
	DeputizeKeeper helpers.TransactionKeeper
}

var (
	parameterManager      helpers.ParameterManager
	deputizeAuxiliary     helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
)

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper) {

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
	memberAuxiliary := member.Auxiliary.Initialize(Mapper, parameterManager)
	deputizeAuxiliary = deputize.Auxiliary.Initialize(Mapper, parameterManager, memberAuxiliary)
	keepers := TestKeepers{
		DeputizeKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{authenticateAuxiliary, deputizeAuxiliary}).(helpers.TransactionKeeper),
	}

	return Context, keepers, Mapper
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
	_, _, Mapper := CreateTestInput(t)
	type fields struct {
		mapper            helpers.Mapper
		parameterManager  helpers.ParameterManager
		deputizeAuxiliary helpers.Auxiliary
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
		{"valid", fields{Mapper, parameterManager, deputizeAuxiliary}, args{Mapper, parameterManager, []interface{}{deputizeAuxiliary, authenticateAuxiliary}}, transactionKeeper{Mapper, parameterManager, deputizeAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:            tt.fields.mapper,
				parameterManager:  tt.fields.parameterManager,
				deputizeAuxiliary: tt.fields.deputizeAuxiliary,
			}
			got := transactionKeeper.Initialize(tt.args.mapper, tt.args.parameterManager, tt.args.auxiliaries)
			assert.NotNil(t, got)
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	Context, _, Mapper := CreateTestInput(t)

	depAux, depAuxKeeper := testutil.NewMockAuxiliaryPair()
	depAuxKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)

	fromAccAddress, err := sdkTypes.AccAddressFromBech32("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	require.NoError(t, err)
	immutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	maintainedProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("deputize"), baseData.NewListData()))
	mutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()))
	immutables := baseQualified.NewImmutables(immutableMetaProperties)
	mutables := baseQualified.NewMutables(mutableMetaProperties)
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	identity := baseDocuments.NewIdentity(classificationID, immutables, mutables)
	identity = identity.ProvisionAddress([]sdkTypes.AccAddress{fromAccAddress}...)
	fromIdentityID := baseIDs.NewIdentityID(classificationID, immutables)
	toIdentityID := baseIDs.NewIdentityID(classificationID, immutables)

	tk := transactionKeeper{mapper: Mapper, parameterManager: parameterManager, deputizeAuxiliary: depAux}
	tk.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(record.NewRecord(identity))

	t.Run("valid", func(t *testing.T) {
		got, err := tk.Transact(sdkTypes.WrapSDKContext(Context), NewMessage(fromAccAddress, fromIdentityID, toIdentityID, classificationID, maintainedProperties, true, true, true, true, true).(*Message))
		require.NoError(t, err)
		require.NotNil(t, got)
	})
}
