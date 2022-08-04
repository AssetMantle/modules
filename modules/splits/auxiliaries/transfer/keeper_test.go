// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transfer

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

	"github.com/AssetMantle/modules/modules/splits/internal/key"
	"github.com/AssetMantle/modules/modules/splits/internal/mappable"
	"github.com/AssetMantle/modules/modules/splits/internal/parameters"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

type TestKeepers struct {
	SplitsKeeper helpers.AuxiliaryKeeper
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
		SplitsKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.AuxiliaryKeeper),
	}

	return context, keepers

}

func Test_Transfer_Aux_Keeper_Help(t *testing.T) {
	context, keepers := CreateTestInput(t)
	ownerID := baseIDs.NewStringID("ownerID")
	ownableID := baseIDs.NewStringID("ownableID")

	toID := baseIDs.NewStringID("toID")
	defaultSplitID := baseIDs.NewSplitID(ownerID, ownableID)
	splits := sdkTypes.NewDec(123)
	keepers.SplitsKeeper.(auxiliaryKeeper).mapper.NewCollection(context).Add(mappable.NewSplit(defaultSplitID, splits))

	t.Run("Positive case-  Value transfer", func(t *testing.T) {
		want := newAuxiliaryResponse(nil)
		if got := keepers.SplitsKeeper.Help(context, NewAuxiliaryRequest(ownerID, toID, ownableID, sdkTypes.NewDec(1))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("Positive case-  Value transfer", func(t *testing.T) {
		want := newAuxiliaryResponse(nil)
		if got := keepers.SplitsKeeper.Help(context, NewAuxiliaryRequest(ownerID, toID, ownableID, sdkTypes.NewDec(1))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("Positive case-  Value transfer", func(t *testing.T) {
		want := newAuxiliaryResponse(nil)
		if got := keepers.SplitsKeeper.Help(context, NewAuxiliaryRequest(toID, ownerID, ownableID, sdkTypes.NewDec(2))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-0 Value transfer", func(t *testing.T) {
		want := newAuxiliaryResponse(constants.NotAuthorized)
		if got := keepers.SplitsKeeper.Help(context, NewAuxiliaryRequest(ownerID, toID, ownableID, sdkTypes.NewDec(0))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-No Value Present", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(constants.EntityNotFound)
		if got := keepers.SplitsKeeper.Help(context, NewAuxiliaryRequest(ownerID, toID, baseIDs.NewStringID("ownableIDNotPresent"), sdkTypes.NewDec(1))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Transfer More than available splits", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(constants.NotAuthorized)
		if got := keepers.SplitsKeeper.Help(context, NewAuxiliaryRequest(ownerID, toID, ownableID, sdkTypes.NewDec(1234))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
}
