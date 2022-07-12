// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"reflect"
	"testing"

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
	"github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

type TestKeepers struct {
	ClassificationsKeeper helpers.AuxiliaryKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {
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

	return context, keepers
}

func Test_Auxiliary_Keeper_Help(t *testing.T) {
	context, keepers := CreateTestInput(t)

	immutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2")))
	mutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")))

	classificationID := key.NewClassificationID(immutableProperties, mutableProperties)

	testClassificationID := key.NewClassificationID(base.NewPropertyList(), base.NewPropertyList())

	keepers.ClassificationsKeeper.(auxiliaryKeeper).mapper.NewCollection(context).Add(mappable.NewClassification(base.NewPropertyList(), base.NewPropertyList()))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newAuxiliaryResponse(baseIDs.NewStringID(classificationID.String()), nil)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(immutableProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Classification already present", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(baseIDs.NewStringID(testClassificationID.String()), constants.EntityAlreadyExists)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(base.NewPropertyList(), base.NewPropertyList())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Max Property Count", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, constants.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")), baseProperties.NewProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2")), baseProperties.NewProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("Data3")), baseProperties.NewProperty(baseIDs.NewStringID("ID4"), baseData.NewStringData("Data4")), baseProperties.NewProperty(baseIDs.NewStringID("ID5"), baseData.NewStringData("Data5")), baseProperties.NewProperty(baseIDs.NewStringID("ID6"), baseData.NewStringData("Data6")), baseProperties.NewProperty(baseIDs.NewStringID("ID7"), baseData.NewStringData("Data7")), baseProperties.NewProperty(baseIDs.NewStringID("ID8"), baseData.NewStringData("Data8")), baseProperties.NewProperty(baseIDs.NewStringID("ID9"), baseData.NewStringData("Data9")), baseProperties.NewProperty(baseIDs.NewStringID("ID10"), baseData.NewStringData("Data10")), baseProperties.NewProperty(baseIDs.NewStringID("ID9"), baseData.NewStringData("Data9")), baseProperties.NewProperty(baseIDs.NewStringID("ID10"), baseData.NewStringData("Data10"))), base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")), baseProperties.NewProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2")), baseProperties.NewProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("Data3")), baseProperties.NewProperty(baseIDs.NewStringID("ID4"), baseData.NewStringData("Data4")), baseProperties.NewProperty(baseIDs.NewStringID("ID5"), baseData.NewStringData("Data5")), baseProperties.NewProperty(baseIDs.NewStringID("ID6"), baseData.NewStringData("Data6")), baseProperties.NewProperty(baseIDs.NewStringID("ID7"), baseData.NewStringData("Data7")), baseProperties.NewProperty(baseIDs.NewStringID("ID8"), baseData.NewStringData("Data8")), baseProperties.NewProperty(baseIDs.NewStringID("ID9"), baseData.NewStringData("Data9")), baseProperties.NewProperty(baseIDs.NewStringID("ID10"), baseData.NewStringData("Data10")), baseProperties.NewProperty(baseIDs.NewStringID("ID9"), baseData.NewStringData("Data9")), baseProperties.NewProperty(baseIDs.NewStringID("ID10"), baseData.NewStringData("Data10"))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Duplicate Immutable Property", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, constants.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")), baseProperties.NewProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data2")), baseProperties.NewProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("Data3")), baseProperties.NewProperty(baseIDs.NewStringID("ID4"), baseData.NewStringData("Data4")), baseProperties.NewProperty(baseIDs.NewStringID("ID5"), baseData.NewStringData("Data5")), baseProperties.NewProperty(baseIDs.NewStringID("ID6"), baseData.NewStringData("Data6")), baseProperties.NewProperty(baseIDs.NewStringID("ID7"), baseData.NewStringData("Data7")), baseProperties.NewProperty(baseIDs.NewStringID("ID8"), baseData.NewStringData("Data8")), baseProperties.NewProperty(baseIDs.NewStringID("ID9"), baseData.NewStringData("Data9")), baseProperties.NewProperty(baseIDs.NewStringID("ID10"), baseData.NewStringData("Data10"))), base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID11"), baseData.NewStringData("Data11")), baseProperties.NewProperty(baseIDs.NewStringID("ID12"), baseData.NewStringData("Data12")), baseProperties.NewProperty(baseIDs.NewStringID("ID13"), baseData.NewStringData("Data13")), baseProperties.NewProperty(baseIDs.NewStringID("ID14"), baseData.NewStringData("Data14")), baseProperties.NewProperty(baseIDs.NewStringID("ID15"), baseData.NewStringData("Data15")), baseProperties.NewProperty(baseIDs.NewStringID("ID16"), baseData.NewStringData("Data16")), baseProperties.NewProperty(baseIDs.NewStringID("ID17"), baseData.NewStringData("Data17")), baseProperties.NewProperty(baseIDs.NewStringID("ID18"), baseData.NewStringData("Data18")), baseProperties.NewProperty(baseIDs.NewStringID("ID19"), baseData.NewStringData("Data19")), baseProperties.NewProperty(baseIDs.NewStringID("ID20"), baseData.NewStringData("Data20"))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Duplicate Immutable and Mutable Property", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(nil, constants.InvalidRequest)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")), baseProperties.NewProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2")), baseProperties.NewProperty(baseIDs.NewStringID("ID3"), baseData.NewStringData("Data3")), baseProperties.NewProperty(baseIDs.NewStringID("ID4"), baseData.NewStringData("Data4")), baseProperties.NewProperty(baseIDs.NewStringID("ID5"), baseData.NewStringData("Data5")), baseProperties.NewProperty(baseIDs.NewStringID("ID6"), baseData.NewStringData("Data6")), baseProperties.NewProperty(baseIDs.NewStringID("ID7"), baseData.NewStringData("Data7")), baseProperties.NewProperty(baseIDs.NewStringID("ID8"), baseData.NewStringData("Data8")), baseProperties.NewProperty(baseIDs.NewStringID("ID9"), baseData.NewStringData("Data9")), baseProperties.NewProperty(baseIDs.NewStringID("ID10"), baseData.NewStringData("Data10"))), base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data11")), baseProperties.NewProperty(baseIDs.NewStringID("ID12"), baseData.NewStringData("Data12")), baseProperties.NewProperty(baseIDs.NewStringID("ID13"), baseData.NewStringData("Data13")), baseProperties.NewProperty(baseIDs.NewStringID("ID14"), baseData.NewStringData("Data14")), baseProperties.NewProperty(baseIDs.NewStringID("ID15"), baseData.NewStringData("Data15")), baseProperties.NewProperty(baseIDs.NewStringID("ID16"), baseData.NewStringData("Data16")), baseProperties.NewProperty(baseIDs.NewStringID("ID17"), baseData.NewStringData("Data17")), baseProperties.NewProperty(baseIDs.NewStringID("ID18"), baseData.NewStringData("Data18")), baseProperties.NewProperty(baseIDs.NewStringID("ID19"), baseData.NewStringData("Data19")), baseProperties.NewProperty(baseIDs.NewStringID("ID20"), baseData.NewStringData("Data20"))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
}
