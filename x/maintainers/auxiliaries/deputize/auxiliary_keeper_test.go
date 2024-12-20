// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"context"
	"fmt"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	tendermintDB "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/member"
	"github.com/AssetMantle/modules/x/maintainers/mapper"
	"github.com/AssetMantle/modules/x/maintainers/parameters"
	"github.com/AssetMantle/modules/x/maintainers/record"
	maintainerUtilities "github.com/AssetMantle/modules/x/maintainers/utilities"
)

type TestKeepers struct {
	DeputizeKeeper helpers.AuxiliaryKeeper
}

var (
	memberAuxiliary         helpers.Auxiliary
	immutables              = baseQualified.NewImmutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables                = baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID    = baseIDs.NewClassificationID(immutables, mutables)
	testFromID              = baseIDs.NewIdentityID(testClassificationID, immutables)
	maintainedProperty      = "maintainedProperty:S|maintainedProperty"
	maintainedProperties, _ = base.NewPropertyList().FromMetaPropertiesString(maintainedProperty)
	permissions             = maintainerUtilities.SetModulePermissions(true, true, true)
)

func createTestInput(t *testing.T) (types.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {

	storeKey := types.NewKVStoreKey("test")
	paramsStoreKey := types.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := types.NewTransientStoreKey("testParamsTransient")
	Mapper := mapper.Prototype().Initialize(storeKey)

	parameterManager := parameters.Prototype().Initialize(storeKey)

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, storeTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	Context := types.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	memberAuxiliary = member.Auxiliary.Initialize(Mapper, parameterManager)
	keepers := TestKeepers{
		DeputizeKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.AuxiliaryKeeper),
	}

	return Context, keepers, Mapper, parameterManager
}

func Test_auxiliaryKeeper_Help(t *testing.T) {
	Context, keepers, Mapper, parameterManager := createTestInput(t)
	keepers.DeputizeKeeper.(auxiliaryKeeper).mapper.NewCollection(Context.Context()).Add(record.NewRecord(baseDocuments.NewMaintainer(testFromID, testClassificationID, maintainedProperties.GetPropertyIDList(), permissions)))
	type fields struct {
		mapper           helpers.Mapper
		parameterManager helpers.ParameterManager
		memberAuxiliary  helpers.Auxiliary
	}
	type args struct {
		context context.Context
		request helpers.AuxiliaryRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.AuxiliaryResponse
		wantErr bool
	}{
		{"+ve", fields{Mapper, parameterManager, memberAuxiliary}, args{Context.Context(), NewAuxiliaryRequest(testFromID, testFromID, testClassificationID, maintainedProperties, true, true, true)}, newAuxiliaryResponse(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryKeeper := auxiliaryKeeper{
				mapper:           tt.fields.mapper,
				parameterManager: tt.fields.parameterManager,
				memberAuxiliary:  tt.fields.memberAuxiliary,
			}
			got, err := auxiliaryKeeper.Help(tt.args.context, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Help() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Help() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliaryKeeper_Initialize(t *testing.T) {
	_, _, Mapper, parameterManager := createTestInput(t)
	type fields struct {
		mapper          helpers.Mapper
		memberAuxiliary helpers.Auxiliary
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
		{"+ve", fields{Mapper, memberAuxiliary}, args{Mapper, parameterManager, []interface{}{}}, auxiliaryKeeper{Mapper, parameterManager, memberAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryKeeper := auxiliaryKeeper{
				mapper:          tt.fields.mapper,
				memberAuxiliary: tt.fields.memberAuxiliary,
			}
			if got := auxiliaryKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keeperPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.AuxiliaryKeeper
	}{
		{"+ve", auxiliaryKeeper{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keeperPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keeperPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
