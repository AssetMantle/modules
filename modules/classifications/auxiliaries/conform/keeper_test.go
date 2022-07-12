// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package conform

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
	classificationID := baseIDs.NewStringID("classificationID")
	mutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")))
	immutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2")))

	keepers.ClassificationsKeeper.(auxiliaryKeeper).mapper.NewCollection(context).Add(mappable.NewClassification(immutableProperties, mutableProperties))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newAuxiliaryResponse(nil)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(classificationID, immutableProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("Negative Case - Immutable Data Type mismatch", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(constants.NotAuthorized)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(classificationID, base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID2"), baseData.NewIDData(baseIDs.NewStringID("Data2")))), mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("Negative Case - Mutable Data Type mismatch", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(constants.NotAuthorized)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(classificationID, immutableProperties, base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID1"), baseData.NewIDData(baseIDs.NewStringID("Data1")))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-PropertyList list length mismatch", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(constants.NotAuthorized)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(classificationID, base.NewPropertyList(), mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-PropertyList mismatch", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(constants.NotAuthorized)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(classificationID, base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data3"))), mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-PropertyList list length mismatch", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(constants.NotAuthorized)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(classificationID, immutableProperties, base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID4"), baseData.NewStringData("Data4")), baseProperties.NewProperty(baseIDs.NewStringID("ID5"), baseData.NewStringData("Data5"))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-PropertyList mismatch", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(constants.NotAuthorized)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(classificationID, immutableProperties, base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID6"), baseData.NewStringData("Data3"))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase- Classification EntityNotFound", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(constants.EntityNotFound)
		if got := keepers.ClassificationsKeeper.Help(context, NewAuxiliaryRequest(baseIDs.NewStringID("test.classification"), immutableProperties, base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID6"), baseData.NewStringData("Data3"))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
