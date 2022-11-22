// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"fmt"
	"reflect"
	"testing"

	baseDocuments "github.com/AssetMantle/modules/schema/documents/base"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/modules/classifications/internal/mappable"
	"github.com/AssetMantle/modules/modules/classifications/internal/parameters"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

type TestKeepers struct {
	ClassificationsKeeper helpers.AuxiliaryKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.Parameters) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := params.NewKeeper(
		Codec,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters := parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		ClassificationsKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.AuxiliaryKeeper),
	}

	return context, keepers, Mapper, Parameters
}

func Test_Auxiliary_Keeper_Help(t *testing.T) {
	context, keepers, _, _ := CreateTestInput(t)

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

	t.Run("NegativeCase-Classification already present", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(testClassificationID, constants.EntityAlreadyExists)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList()))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Max Property Count", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, constants.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("Data3")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID4"), baseData.NewStringData("Data4")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID5"), baseData.NewStringData("Data5")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID6"), baseData.NewStringData("Data6")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID7"), baseData.NewStringData("Data7")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID8"), baseData.NewStringData("Data8")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID9"), baseData.NewStringData("Data9")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID10"), baseData.NewStringData("Data10")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID9"), baseData.NewStringData("Data9")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID10"), baseData.NewStringData("Data10")))), baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("Data3")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID4"), baseData.NewStringData("Data4")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID5"), baseData.NewStringData("Data5")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID6"), baseData.NewStringData("Data6")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID7"), baseData.NewStringData("Data7")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID8"), baseData.NewStringData("Data8")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID9"), baseData.NewStringData("Data9")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID10"), baseData.NewStringData("Data10")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID9"), baseData.NewStringData("Data9")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID10"), baseData.NewStringData("Data10")))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Duplicate Immutable Property", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, constants.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data2")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("Data3")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID4"), baseData.NewStringData("Data4")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID5"), baseData.NewStringData("Data5")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID6"), baseData.NewStringData("Data6")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID7"), baseData.NewStringData("Data7")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID8"), baseData.NewStringData("Data8")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID9"), baseData.NewStringData("Data9")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID10"), baseData.NewStringData("Data10")))), baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID11"), baseData.NewStringData("Data11")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID12"), baseData.NewStringData("Data12")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID13"), baseData.NewStringData("Data13")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID14"), baseData.NewStringData("Data14")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID15"), baseData.NewStringData("Data15")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID16"), baseData.NewStringData("Data16")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID17"), baseData.NewStringData("Data17")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID18"), baseData.NewStringData("Data18")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID19"), baseData.NewStringData("Data19")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID20"), baseData.NewStringData("Data20")))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Duplicate Immutable and Mutable Property", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, constants.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("Data3")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID4"), baseData.NewStringData("Data4")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID5"), baseData.NewStringData("Data5")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID6"), baseData.NewStringData("Data6")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID7"), baseData.NewStringData("Data7")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID8"), baseData.NewStringData("Data8")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID9"), baseData.NewStringData("Data9")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID10"), baseData.NewStringData("Data10")))), baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data11")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID12"), baseData.NewStringData("Data12")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID13"), baseData.NewStringData("Data13")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID14"), baseData.NewStringData("Data14")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID15"), baseData.NewStringData("Data15")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID16"), baseData.NewStringData("Data16")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID17"), baseData.NewStringData("Data17")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID18"), baseData.NewStringData("Data18")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID19"), baseData.NewStringData("Data19")), baseProperties.NewMesaProperty(baseIDs.NewStringID("ID20"), baseData.NewStringData("Data20")))))); !reflect.DeepEqual(got, want) {
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
	_, _, Mapper, Parameters := CreateTestInput(t)
	type fields struct {
		mapper helpers.Mapper
	}
	type args struct {
		mapper helpers.Mapper
		in1    helpers.Parameters
		in2    []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}{
		{"+ve", fields{Mapper}, args{Mapper, Parameters, []interface{}{}}, auxiliaryKeeper{Mapper}},
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
