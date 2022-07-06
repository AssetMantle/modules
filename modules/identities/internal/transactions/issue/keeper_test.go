// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

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

	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
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

	scrubAuxiliary := scrub.AuxiliaryMock.Initialize(Mapper, Parameters)
	conformAuxiliary := conform.AuxiliaryMock.Initialize(Mapper, Parameters)
	keepers := TestKeepers{
		IdentitiesKeeper: keeperPrototype().Initialize(Mapper, Parameters,
			[]interface{}{scrubAuxiliary,
				conformAuxiliary}).(helpers.TransactionKeeper),
	}

	return context, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers := CreateTestInput(t)

	immutableMetaProperties, err := utilities.ReadMetaPropertyList("defaultImmutableMeta1:S|defaultImmutableMeta1")
	require.Equal(t, nil, err)

	var immutableProperties lists.PropertyList
	immutableProperties, err = utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)

	var mutableMetaProperties lists.MetaPropertyList
	mutableMetaProperties, err = utilities.ReadMetaPropertyList("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, err)

	var mutableProperties lists.PropertyList
	mutableProperties, err = utilities.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, err)

	var scrubMockErrorProperties lists.MetaPropertyList
	scrubMockErrorProperties, err = utilities.ReadMetaPropertyList("scrubError:S|mockError")
	require.Equal(t, nil, err)

	var conformMockErrorProperties lists.MetaPropertyList
	conformMockErrorProperties, err = utilities.ReadMetaPropertyList("conformError:S|mockError")
	require.Equal(t, nil, err)

	defaultAddr := sdkTypes.AccAddress("addr")
	defaultClassificationID := baseIDs.NewID("test.cGn3HMW8M3t5gMDv-wXa9sseHnA=")
	defaultIdentityID := key.NewIdentityID(defaultClassificationID, immutableProperties)
	keepers.IdentitiesKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewIdentity(defaultIdentityID, base.NewPropertyList(), base.NewPropertyList()))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.IdentitiesKeeper.Transact(context, newMessage(defaultAddr, defaultAddr, defaultIdentityID, defaultClassificationID,
			immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-DuplicateIdentity", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(constants.EntityAlreadyExists)
		if got := keepers.IdentitiesKeeper.Transact(context, newMessage(defaultAddr, defaultAddr, defaultIdentityID, defaultClassificationID,
			immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-ImmutableScrub Error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(constants.MockError)
		if got := keepers.IdentitiesKeeper.Transact(context, newMessage(defaultAddr, defaultAddr, defaultIdentityID, defaultClassificationID,
			scrubMockErrorProperties, immutableProperties, mutableMetaProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-MutableScrubError", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(constants.MockError)
		if got := keepers.IdentitiesKeeper.Transact(context, newMessage(defaultAddr, defaultAddr, defaultIdentityID, baseIDs.NewID("newClassificationID"),
			immutableMetaProperties, immutableProperties, scrubMockErrorProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-ConformError", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(constants.MockError)
		if got := keepers.IdentitiesKeeper.Transact(context, newMessage(defaultAddr, defaultAddr, defaultIdentityID, baseIDs.NewID("newClassificationID"),
			immutableMetaProperties, immutableProperties, conformMockErrorProperties, mutableProperties)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
}
