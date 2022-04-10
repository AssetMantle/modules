// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/constants/test"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"

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
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

type TestKeepers struct {
	IdentitiesKeeper helpers.TransactionKeeper
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

	scrubAuxiliary := scrub.AuxiliaryMock.Initialize(mapper, Parameters)
	maintainAuxiliary := maintain.AuxiliaryMock.Initialize(mapper, Parameters)
	conformAuxiliary := conform.AuxiliaryMock.Initialize(mapper, Parameters)
	verifyAuxiliary := verify.AuxiliaryMock.Initialize(mapper, Parameters)
	keepers := TestKeepers{
		IdentitiesKeeper: keeperPrototype().Initialize(mapper, Parameters,
			[]interface{}{scrubAuxiliary, verifyAuxiliary,
				maintainAuxiliary, conformAuxiliary}).(helpers.TransactionKeeper),
	}

	return context, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers := CreateTestInput(t)

	immutableProperties, err := baseTypes.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)
	mutableMetaProperties, err := baseTypes.ReadMetaProperties("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, err)
	mutableProperties, err := baseTypes.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, err)
	scrubMockErrorProperties, err := baseTypes.ReadMetaProperties("scrubError:S|mockError")
	require.Equal(t, nil, err)
	conformMockErrorProperties, err := baseTypes.ReadMetaProperties("conformError:S|mockError")
	require.Equal(t, nil, err)
	defaultAddr := sdkTypes.AccAddress("addr")
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")
	defaultIdentityID := baseIDs.NewID("fromIdentityID")
	maintainIdentityMockError := baseIDs.NewID("maintainError")
	classificationID := baseIDs.NewID("ClassificationID")
	identityID := key.NewIdentityID(classificationID, immutableProperties)
	keepers.IdentitiesKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewIdentity(identityID, immutableProperties, mutableProperties))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		require.Panics(t, func() {
			if got := keepers.IdentitiesKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, identityID,
				mutableMetaProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
				t.Errorf("Transact() = %v, want %v", got, want)
			}
		})
	})

	t.Run("NegativeCase - verify identity mock error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(test.MockError)
		if got := keepers.IdentitiesKeeper.Transact(context, newMessage(verifyMockErrorAddress, defaultIdentityID, identityID,
			mutableMetaProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - UnMinted identity", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.EntityNotFound)
		if got := keepers.IdentitiesKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, baseIDs.NewID("identityID"),
			mutableMetaProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - scrub error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(test.MockError)
		if got := keepers.IdentitiesKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, identityID,
			scrubMockErrorProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
	t.Run("NegativeCase - conform error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(test.MockError)
		require.Panics(t, func() {
			if got := keepers.IdentitiesKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, identityID,
				conformMockErrorProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
				t.Errorf("Transact() = %v, want %v", got, want)
			}
		})
	})
	t.Run("NegativeCase - maintain Error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(test.MockError)
		require.Panics(t, func() {
			if got := keepers.IdentitiesKeeper.Transact(context, newMessage(defaultAddr, maintainIdentityMockError, identityID,
				mutableMetaProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
				t.Errorf("Transact() = %v, want %v", got, want)
			}
		})
	})

}
