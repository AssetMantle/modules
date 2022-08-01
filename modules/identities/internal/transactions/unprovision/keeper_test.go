// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unprovision

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

	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	listsUtilities "github.com/AssetMantle/modules/schema/lists/utilities"
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

	keepers := TestKeepers{
		IdentitiesKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.TransactionKeeper),
	}

	return context, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers := CreateTestInput(t)
	defaultAddr := sdkTypes.AccAddress("addr")
	provisionedAddr2 := sdkTypes.AccAddress("addr2")
	unprovisionedAddr := sdkTypes.AccAddress("unProvisionedAddr")
	immutableProperties, _ := listsUtilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	defaultClassificationID := baseIDs.NewStringID("test.cGn3HMW8M3t5gMDv-wXa9sseHnA=")
	defaultIdentityID := baseIDs.NewIdentityID(defaultClassificationID, immutableProperties)

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.IdentitiesKeeper.Transact(context, newMessage(defaultAddr, provisionedAddr2, defaultIdentityID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Nil Identity", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(constants.EntityNotFound)
		if got := keepers.IdentitiesKeeper.Transact(context, newMessage(defaultAddr, provisionedAddr2, baseIDs.NewStringID("id"))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-UnProvisioning unprovisioned Address", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(constants.DeletionNotAllowed)
		if got := keepers.IdentitiesKeeper.Transact(context, newMessage(defaultAddr, unprovisionedAddr, defaultIdentityID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-UnProvisioning random Address", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(constants.EntityNotFound)
		if got := keepers.IdentitiesKeeper.Transact(context, newMessage(defaultAddr, sdkTypes.AccAddress("randomAddr"), defaultIdentityID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Identity From Address mismatch", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(constants.NotAuthorized)
		if got := keepers.IdentitiesKeeper.Transact(context, newMessage(sdkTypes.AccAddress("randomAddr"), unprovisionedAddr, defaultIdentityID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
