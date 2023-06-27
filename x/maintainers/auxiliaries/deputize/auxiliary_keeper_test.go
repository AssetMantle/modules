// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"fmt"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/go/data/base"
	baseDocuments "github.com/AssetMantle/schema/go/documents/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/lists/utilities"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/member"
	"github.com/AssetMantle/modules/x/maintainers/key"
	"github.com/AssetMantle/modules/x/maintainers/mappable"
	"github.com/AssetMantle/modules/x/maintainers/parameters"
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
	maintainedProperties, _ = utilities.ReadMetaPropertyList(maintainedProperty)
	permissions             = maintainerUtilities.SetModulePermissions(true, true, true)
)

func createTestInput(t *testing.T) (types.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {
	var legacyAmino = baseHelpers.CodecPrototype().GetLegacyAmino()

	storeKey := types.NewKVStoreKey("test")
	paramsStoreKey := types.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := types.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	encodingConfig := simapp.MakeTestEncodingConfig()
	appCodec := encodingConfig.Marshaler
	ParamsKeeper := paramsKeeper.NewKeeper(
		appCodec,
		legacyAmino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	parameterManager := parameters.Prototype().Initialize(ParamsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, types.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, types.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, types.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := types.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	memberAuxiliary = member.Auxiliary.Initialize(Mapper, parameterManager)
	keepers := TestKeepers{
		DeputizeKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.AuxiliaryKeeper),
	}

	return context, keepers, Mapper, parameterManager
}

func Test_auxiliaryKeeper_Help(t *testing.T) {
	context, keepers, Mapper, _ := createTestInput(t)
	keepers.DeputizeKeeper.(auxiliaryKeeper).mapper.NewCollection(types.WrapSDKContext(context)).Add(mappable.NewMappable(baseDocuments.NewMaintainer(testFromID, testClassificationID, maintainedProperties.GetPropertyIDList(), permissions)))
	type fields struct {
		mapper          helpers.Mapper
		memberAuxiliary helpers.Auxiliary
	}
	type args struct {
		context types.Context
		request helpers.AuxiliaryRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.AuxiliaryResponse
	}{
		{"+ve", fields{Mapper, memberAuxiliary}, args{context, NewAuxiliaryRequest(testFromID, testFromID, testClassificationID, maintainedProperties, true, true, true)}, newAuxiliaryResponse()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryKeeper := auxiliaryKeeper{
				mapper:          tt.fields.mapper,
				memberAuxiliary: tt.fields.memberAuxiliary,
			}
			if got := auxiliaryKeeper.Help(types.WrapSDKContext(tt.args.context), tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Help() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliaryKeeper_Initialize(t *testing.T) {
	_, _, mapper, parameterManager := createTestInput(t)
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
		{"+ve", fields{mapper, memberAuxiliary}, args{mapper, parameterManager, []interface{}{}}, auxiliaryKeeper{mapper, parameterManager, memberAuxiliary}},
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
