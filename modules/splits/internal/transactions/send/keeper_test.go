// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package send

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

	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/splits/internal/key"
	"github.com/AssetMantle/modules/modules/splits/internal/mappable"
	"github.com/AssetMantle/modules/modules/splits/internal/parameters"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

type TestKeepers struct {
	SplitsKeeper helpers.TransactionKeeper
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

	verifyAuxiliary := verify.AuxiliaryMock.Initialize(Mapper, Parameters)
	keepers := TestKeepers{
		SplitsKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{verifyAuxiliary}).(helpers.TransactionKeeper),
	}

	return context, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers := CreateTestInput(t)
	defaultAddr := sdkTypes.AccAddress("addr")
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")

	fromID := baseIDs.NewID("fromID")
	toID := baseIDs.NewID("toID")
	ownableID := baseIDs.NewID("stake")

	keepers.SplitsKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewSplit(key.NewSplitID(fromID, ownableID), sdkTypes.NewDec(100)))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.SplitsKeeper.Transact(context, newMessage(defaultAddr, fromID, toID, ownableID, sdkTypes.NewDec(1))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("Positive Case-Send All splits", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.SplitsKeeper.Transact(context, newMessage(defaultAddr, toID, fromID, ownableID, sdkTypes.NewDec(1))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Verify Identity Failure", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(constants.MockError)
		if got := keepers.SplitsKeeper.Transact(context, newMessage(verifyMockErrorAddress, fromID, toID, ownableID, sdkTypes.NewDec(1))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Negative Value exchange", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(constants.NotAuthorized)
		if got := keepers.SplitsKeeper.Transact(context, newMessage(defaultAddr, fromID, toID, ownableID, sdkTypes.NewDec(-1))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Value not found", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(constants.EntityNotFound)
		if got := keepers.SplitsKeeper.Transact(context, newMessage(defaultAddr, baseIDs.NewID("fakeFromID"), toID, ownableID, sdkTypes.NewDec(1))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Send More than available splits", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(constants.NotAuthorized)
		if got := keepers.SplitsKeeper.Transact(context, newMessage(defaultAddr, fromID, toID, ownableID, sdkTypes.NewDec(101))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
