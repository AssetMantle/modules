// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/modules/classifications/internal/mappable"
	"github.com/AssetMantle/modules/modules/classifications/internal/parameters"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseDocuments "github.com/AssetMantle/modules/schema/documents/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

type TestKeepers struct {
	ClassificationsKeeper helpers.AuxiliaryKeeper
}

func createTestInput(t *testing.T) (context.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {
	var legacyAmino = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
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
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		ClassificationsKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.AuxiliaryKeeper),
	}

	return sdkTypes.WrapSDKContext(context), keepers, Mapper, parameterManager
}

func Test_Auxiliary_Keeper_Help(t *testing.T) {
	context, keepers, _, _ := createTestInput(t)

	immutableProperties := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutableProperties := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))

	classificationID := baseIDs.NewClassificationID(immutableProperties, mutableProperties)

	testClassificationID := baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList()))

	keepers.ClassificationsKeeper.(auxiliaryKeeper).mapper.NewCollection(context).Add(mappable.NewMappable(baseDocuments.NewClassification(baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList()))))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newAuxiliaryResponse(classificationID, nil)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(immutableProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("PositiveCase with max Properties", func(t *testing.T) {
		want := newAuxiliaryResponse(baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList(createTestProperties(12, 0, "immutable")...)), baseQualified.NewMutables(baseLists.NewPropertyList(createTestProperties(0, 10, "mutable")...))), nil)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(baseQualified.NewImmutables(baseLists.NewPropertyList(createTestProperties(12, 0, "immutable")...)), baseQualified.NewMutables(baseLists.NewPropertyList(createTestProperties(0, 10, "mutable")...)))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Classification already present", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(testClassificationID, errorConstants.EntityAlreadyExists)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList()))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Max Property Count", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, errorConstants.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(baseQualified.NewImmutables(baseLists.NewPropertyList(createTestProperties(10, 5, "immutable")...)), baseQualified.NewMutables(baseLists.NewPropertyList(createTestProperties(5, 5, "mutable")...)))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Duplicate Immutable Property", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, errorConstants.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(baseQualified.NewImmutables(baseLists.NewPropertyList(append(createTestProperties(5, 0, "immutable"), createTestProperties(1, 0, "immutable")[0])...)), baseQualified.NewMutables(baseLists.NewPropertyList(createTestProperties(2, 0, "mutable")...)))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Duplicate Immutable and Mutable Property", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, errorConstants.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(baseQualified.NewImmutables(baseLists.NewPropertyList(append(createTestProperties(5, 0, "immutable"), createTestProperties(1, 0, "immutable")[0])...)), baseQualified.NewMutables(baseLists.NewPropertyList(append(createTestProperties(2, 0, "mutable"), createTestProperties(1, 0, "mutable")...)...)))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
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

func Test_auxiliaryKeeper_Initialize(t *testing.T) {
	_, _, mapper, parameterManager := createTestInput(t)
	type fields struct {
		mapper helpers.Mapper
	}
	type args struct {
		mapper helpers.Mapper
		in1    helpers.ParameterManager
		in2    []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}{
		{"+ve", fields{mapper}, args{mapper, parameterManager, []interface{}{}}, auxiliaryKeeper{mapper, parameterManager}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := auxiliaryKeeper{
				mapper: tt.fields.mapper,
			}
			if got := au.Initialize(tt.args.mapper, tt.args.in1, tt.args.in2); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createTestProperties(noOfMeta int, nOfMesa int, propertyType string) []properties.Property {
	var propertyLists []properties.Property
	for i := 0; i < noOfMeta; i++ {
		propertyLists = append(propertyLists, baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"+propertyType+fmt.Sprint(i)), baseData.NewStringData("DataMeta"+fmt.Sprint(i))))
	}
	for i := 0; i < nOfMesa; i++ {
		propertyLists = append(propertyLists, baseProperties.NewMesaProperty(baseIDs.NewStringID("ID"+propertyType+fmt.Sprint(i)), baseData.NewStringData("DataMesa"+fmt.Sprint(i))))
	}
	return propertyLists
}
