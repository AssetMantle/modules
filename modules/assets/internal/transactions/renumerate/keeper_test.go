/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package renumerate

import (
	"reflect"
	"testing"

	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/persistenceOne/persistenceSDK/schema/applications"
	tendermintProto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
)

type TestKeepers struct {
	AssetsKeeper helpers.TransactionKeeper
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
	mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
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
	immutableProperties, Error := base.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, Error)
	mutableProperties, Error := base.ReadProperties("")
	require.Equal(t, nil, Error)
	supplementError, Error := base.ReadMetaProperties("supplementError:S|mockError")
	require.Equal(t, nil, Error)
	defaultAddr := sdkTypes.AccAddress("addr")
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")
	defaultIdentityID := base.NewID("fromIdentityID")
	classificationID := base.NewID("ClassificationID")
	assetID := key.NewAssetID(classificationID, base.NewImmutables(immutableProperties))
	assetID2 := key.NewAssetID(base.NewID("ClassificationID2"), base.NewImmutables(immutableProperties))
	assetID3 := key.NewAssetID(base.NewID("ClassificationID3"), base.NewImmutables(immutableProperties))
	keepers.AssetsKeeper.(transactionKeeper).mapper.NewCollection(ctx).Add(mappable.NewAsset(assetID,
		base.NewImmutables(immutableProperties), base.NewMutables(mutableProperties)))
	keepers.AssetsKeeper.(transactionKeeper).mapper.NewCollection(ctx).Add(mappable.NewAsset(assetID2,
		base.NewImmutables(immutableProperties), base.NewMutables(supplementError.RemoveData())))
	keepers.AssetsKeeper.(transactionKeeper).mapper.NewCollection(ctx).Add(mappable.NewAsset(assetID3,
		base.NewImmutables(immutableProperties), base.NewMutables(mutableProperties)))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.AssetsKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, assetID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - verify identity mock error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
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
		want := newTransactionResponse(errors.MockError)
		if got := keepers.AssetsKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, assetID2)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
