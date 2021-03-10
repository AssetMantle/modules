/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package transfer

import (
	"reflect"
	"testing"

	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/persistenceOne/persistenceSDK/schema/applications"
	tendermintProto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
)

type TestKeepers struct {
	SplitsKeeper helpers.AuxiliaryKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {
	var Codec = codec.NewLegacyAmino()
	schema.RegisterCodec(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)

	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)
	std.RegisterLegacyAminoCodec(Codec)
	std.RegisterInterfaces(interfaceRegistry)

	encodingConfig := applications.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             Codec,
	}

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := paramsKeeper.NewKeeper(
		encodingConfig.Marshaler,
		encodingConfig.Amino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters := parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	Error := commitMultiStore.LoadLatestVersion()
	require.Nil(t, Error)

	context := sdkTypes.NewContext(commitMultiStore, tendermintProto.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		SplitsKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.AuxiliaryKeeper),
	}

	return context, keepers

}

func Test_Transfer_Aux_Keeper_Help(t *testing.T) {

	context, keepers := CreateTestInput(t)
	ownerID := base.NewID("ownerID")
	ownableID := base.NewID("ownableID")

	toID := base.NewID("toID")
	defaultSplitID := key.NewSplitID(ownerID, ownableID)
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
		want := newAuxiliaryResponse(errors.NotAuthorized)
		if got := keepers.SplitsKeeper.Help(context, NewAuxiliaryRequest(ownerID, toID, ownableID, sdkTypes.NewDec(0))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-No Value Present", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(errors.EntityNotFound)
		if got := keepers.SplitsKeeper.Help(context, NewAuxiliaryRequest(ownerID, toID, base.NewID("ownableIDNotPresent"), sdkTypes.NewDec(1))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Transfer More than available splits", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(errors.NotAuthorized)
		if got := keepers.SplitsKeeper.Help(context, NewAuxiliaryRequest(ownerID, toID, ownableID, sdkTypes.NewDec(1234))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
}
