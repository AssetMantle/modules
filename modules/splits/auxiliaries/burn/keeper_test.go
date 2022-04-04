// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

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

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/modules/splits/internal/key"
	"github.com/AssetMantle/modules/modules/splits/internal/mappable"
	"github.com/AssetMantle/modules/modules/splits/internal/parameters"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/types/base"
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

func Test_Burn_Aux_Keeper_Help(t *testing.T) {
	context, keepers := CreateTestInput(t)

	ownerID := base.NewID("ownerID")
	ownableID := base.NewID("ownableID")

	ownerID2 := base.NewID("ownerID2")
	ownableID2 := base.NewID("ownableID2")

	splitID := key.NewSplitID(ownerID, ownableID)
	splitID2 := key.NewSplitID(ownerID2, ownableID2)
	splits := sdkTypes.NewDec(10)

	keepers.SplitsKeeper.(auxiliaryKeeper).mapper.NewCollection(context).Add(mappable.NewSplit(splitID, splits)).Add(mappable.NewSplit(splitID2, splits))

	t.Run("PositiveCase- mutate split", func(t *testing.T) {
		want := newAuxiliaryResponse(nil)
		if got := keepers.SplitsKeeper.Help(context, NewAuxiliaryRequest(ownerID, ownableID, sdkTypes.NewDec(1))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("PositiveCase- remove split", func(t *testing.T) {
		want := newAuxiliaryResponse(nil)
		if got := keepers.SplitsKeeper.Help(context, NewAuxiliaryRequest(ownerID2, ownableID2, sdkTypes.NewDec(10))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Nil Value", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(errors.EntityNotFound)
		if got := keepers.SplitsKeeper.Help(context, NewAuxiliaryRequest(base.NewID("negativeTestOwner"), base.NewID("negativeTestOwnable"), sdkTypes.NewDec(1))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Insufficient Balance", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(errors.InsufficientBalance)
		if got := keepers.SplitsKeeper.Help(context, NewAuxiliaryRequest(ownerID, ownableID, sdkTypes.NewDec(1234))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
