// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate

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

	"github.com/AssetMantle/modules/schema/errors/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"

	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/assets/internal/mappable"
	"github.com/AssetMantle/modules/modules/assets/internal/parameters"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

type TestKeepers struct {
	AssetsKeeper helpers.TransactionKeeper
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
	mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
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

	supplementAuxiliary := supplement.AuxiliaryMock.Initialize(mapper, Parameters)
	verifyAuxiliary := verify.AuxiliaryMock.Initialize(mapper, Parameters)
	keepers := TestKeepers{
		AssetsKeeper: keeperPrototype().Initialize(mapper, Parameters,
			[]interface{}{supplementAuxiliary,
				verifyAuxiliary}).(helpers.TransactionKeeper),
	}

	return context, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {
	ctx, keepers := CreateTestInput(t)
	ctx = ctx.WithBlockHeight(2)
	immutableProperties, err := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)
	mutableProperties, err := utilities.ReadProperties("")
	require.Equal(t, nil, err)
	supplementError, err := utilities.ReadMetaPropertyList("supplementError:S|mockError")
	require.Equal(t, nil, err)
	defaultAddr := sdkTypes.AccAddress("addr")
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")
	defaultIdentityID := baseIDs.NewStringID("fromIdentityID")
	classificationID := baseIDs.NewStringID("ClassificationID")
	assetID := key.NewAssetID(classificationID, immutableProperties)
	assetID2 := key.NewAssetID(baseIDs.NewStringID("ClassificationID2"), immutableProperties)
	assetID3 := key.NewAssetID(baseIDs.NewStringID("ClassificationID3"), immutableProperties)
	keepers.AssetsKeeper.(transactionKeeper).mapper.NewCollection(ctx).Add(mappable.NewAsset(assetID, immutableProperties, mutableProperties))
	keepers.AssetsKeeper.(transactionKeeper).mapper.NewCollection(ctx).Add(mappable.NewAsset(assetID2, immutableProperties, supplementError.ToPropertyList()))
	keepers.AssetsKeeper.(transactionKeeper).mapper.NewCollection(ctx).Add(mappable.NewAsset(assetID3, immutableProperties, mutableProperties))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		require.Panics(t, func() {
			if got := keepers.AssetsKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, assetID)); !reflect.DeepEqual(got, want) {
				t.Errorf("Transact() = %v, want %v", got, want)
			}
		})
	})

	t.Run("NegativeCase - verify identity mock error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(constants.MockError)
		require.NotPanics(t, func() {
			if got := keepers.AssetsKeeper.Transact(ctx, newMessage(verifyMockErrorAddress, defaultIdentityID, assetID)); !reflect.DeepEqual(got, want) {
				t.Errorf("Transact() = %v, want %v", got, want)
			}
		})
	})

	t.Run("NegativeCase - unMinted asset", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(constants.EntityNotFound)
		require.NotPanics(t, func() {
			if got := keepers.AssetsKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, baseIDs.NewStringID(""))); !reflect.DeepEqual(got, want) {
				t.Errorf("Transact() = %v, want %v", got, want)
			}
		})
	})

	t.Run("NegativeCase - supplement mock error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(constants.MockError)
		require.Panics(t, func() {
			if got := keepers.AssetsKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, assetID2)); !reflect.DeepEqual(got, want) {
				t.Errorf("Transact() = %v, want %v", got, want)
			}
		})
	})

}
