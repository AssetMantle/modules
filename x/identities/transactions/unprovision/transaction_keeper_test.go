// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unprovision

import (
	"fmt"
	storeTypes "cosmossdk.io/store/types"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	"github.com/AssetMantle/modules/x/identities/mapper"
	"github.com/AssetMantle/modules/x/identities/parameters"
	"github.com/AssetMantle/modules/x/identities/record"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
)

type TestKeepers struct {
	UnProvisionKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {

	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	Mapper := mapper.Prototype().Initialize(storeKey)
	parameterManager := parameters.Prototype().Initialize(storeKey)
	parameterManager, _ = parameterManager.Set().Update(sdkTypes.WrapSDKContext(ctx))

	supplementAuxiliary := supplement.Auxiliary.Initialize(Mapper, parameterManager)
	keepers := TestKeepers{
		UnProvisionKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{supplementAuxiliary}).(helpers.TransactionKeeper),
	}

	return ctx, keepers, Mapper, parameterManager
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
	supplementAuxiliary := supplement.Auxiliary.Initialize(Mapper, parameterManager)
	type fields struct {
		mapper              helpers.Mapper
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
		{"valid", fields{Mapper, supplementAuxiliary}, args{Mapper, parameterManager, []interface{}{supplementAuxiliary}}, transactionKeeper{Mapper, supplementAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:              tt.fields.mapper,
				supplementAuxiliary: tt.fields.supplementAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	Mapper := mapper.Prototype().Initialize(storeKey)
	parameterManager := parameters.Prototype().Initialize(storeKey)
	parameterManager, _ = parameterManager.Set().Update(sdkTypes.WrapSDKContext(ctx))

	supplementAux, supplementKeeper := testutil.NewMockAuxiliaryPair()
	supplementKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)

	tk := transactionKeeper{
		mapper:              Mapper,
		supplementAuxiliary: supplementAux,
	}

	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData())))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData())))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	fromAccAddress, err := sdkTypes.AccAddressFromBech32("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	require.NoError(t, err)
	toAccAddress, err := sdkTypes.AccAddressFromBech32("cosmos1u6xn6rv07p2yzzj2rm8st04x54xe5ur0t9nl5j")
	require.NoError(t, err)
	testIdentity := baseDocuments.NewIdentity(testClassificationID, immutables, mutables)
	testIdentity.ProvisionAddress([]sdkTypes.AccAddress{fromAccAddress, toAccAddress}...)
	tk.mapper.NewCollection(sdkTypes.WrapSDKContext(ctx)).Add(record.NewRecord(testIdentity))

	t.Run("not authorized", func(t *testing.T) {
		thirdAddr, err2 := sdkTypes.AccAddressFromBech32("cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef")
		require.NoError(t, err2)
		_, err2 = tk.Transact(sdkTypes.WrapSDKContext(ctx), NewMessage(thirdAddr, toAccAddress, testFromID).(*Message))
		require.Error(t, err2)
	})
	t.Run("valid unprovision", func(t *testing.T) {
		_, err := tk.Transact(sdkTypes.WrapSDKContext(ctx), NewMessage(fromAccAddress, toAccAddress, testFromID).(*Message))
		require.NoError(t, err)
	})
}
