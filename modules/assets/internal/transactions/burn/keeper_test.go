/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package burn

import (
	"reflect"
	"testing"

	"github.com/persistenceOne/persistenceSDK/constants/test"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/burn"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type TestKeepers struct {
	AssetsKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {
	var Codec = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptoCodec.RegisterCrypto(Codec)
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

	burnAuxiliary := burn.AuxiliaryMock.Initialize(mapper, Parameters)
	supplementAuxiliary := supplement.AuxiliaryMock.Initialize(mapper, Parameters)
	verifyAuxiliary := verify.AuxiliaryMock.Initialize(mapper, Parameters)
	keepers := TestKeepers{
		AssetsKeeper: keeperPrototype().Initialize(mapper, Parameters,
			[]interface{}{burnAuxiliary, supplementAuxiliary,
				verifyAuxiliary}).(helpers.TransactionKeeper),
	}

	return context, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {
	ctx, keepers := CreateTestInput(t)
	ctx = ctx.WithBlockHeight(2)
	immutableProperties, err := base.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)
	mutableProperties, err := base.ReadProperties("burn:H|100")
	require.Equal(t, nil, err)
	supplementError, err := base.ReadMetaProperties("supplementError:S|mockError")
	require.Equal(t, nil, err)
	defaultAddr := sdkTypes.AccAddress("addr")
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")
	defaultIdentityID := base.NewID("fromIdentityID")
	burnMockErrorIdentity := base.NewID("burnError")
	classificationID := base.NewID("ClassificationID")
	assetID := key.NewAssetID(classificationID, immutableProperties)
	assetID2 := key.NewAssetID(base.NewID("ClassificationID2"), immutableProperties)
	assetID3 := key.NewAssetID(base.NewID("ClassificationID3"), immutableProperties)
	keepers.AssetsKeeper.(transactionKeeper).mapper.NewCollection(ctx).Add(mappable.NewAsset(assetID, immutableProperties, mutableProperties))
	keepers.AssetsKeeper.(transactionKeeper).mapper.NewCollection(ctx).Add(mappable.NewAsset(assetID2, immutableProperties, supplementError.RemoveData()))
	keepers.AssetsKeeper.(transactionKeeper).mapper.NewCollection(ctx).Add(mappable.NewAsset(assetID3, immutableProperties, mutableProperties))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.AssetsKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, assetID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - verify identity mock error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(test.MockError)
		if got := keepers.AssetsKeeper.Transact(ctx, newMessage(verifyMockErrorAddress, defaultIdentityID, assetID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - unMinted asset", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.EntityNotFound)
		if got := keepers.AssetsKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, base.NewID(""))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - supplement mock error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(test.MockError)
		if got := keepers.AssetsKeeper.Transact(ctx, newMessage(defaultAddr, burnMockErrorIdentity, assetID2)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - burn mock error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(test.MockError)
		if got := keepers.AssetsKeeper.Transact(ctx, newMessage(defaultAddr, burnMockErrorIdentity, assetID3)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
	t.Run("NegativeCase - burn height error", func(t *testing.T) {
		ctx2 := ctx.WithBlockHeight(-20)
		want := newTransactionResponse(errors.NotAuthorized)
		if got := keepers.AssetsKeeper.Transact(ctx2, newMessage(defaultAddr, burnMockErrorIdentity, assetID3)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
