/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package deputize

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
	"reflect"
	"testing"
)

type TestKeepers struct {
	MaintainersKeeper helpers.TransactionKeeper
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
	Error := commitMultiStore.LoadLatestVersion()
	require.Nil(t, Error)

	context := sdkTypes.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	conformAuxiliary := conform.AuxiliaryMock.Initialize(Mapper, Parameters)
	verifyAuxiliary := verify.AuxiliaryMock.Initialize(Mapper, Parameters)
	keepers := TestKeepers{
		MaintainersKeeper: keeperPrototype().Initialize(Mapper, Parameters,
			[]interface{}{verifyAuxiliary, conformAuxiliary}).(helpers.TransactionKeeper),
	}

	return context, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {

	context, keepers := CreateTestInput(t)
	maintainedTraits, Error := base.ReadProperties("maintainedTraits:S|maintainedTraits")
	require.Equal(t, nil, Error)
	conformMockErrorTraits, Error := base.ReadProperties("conformError:S|mockError")
	require.Equal(t, nil, Error)
	defaultAddr := sdkTypes.AccAddress("addr")
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")
	defaultIdentityID := base.NewID("fromIdentityID")
	fakeFromID := base.NewID("fakeFromID")
	toID := base.NewID("toID")
	toID2 := base.NewID("toID2")
	classificationID := base.NewID("ClassificationID")

	keepers.MaintainersKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewMaintainer(key.NewMaintainerID(classificationID, defaultIdentityID),
		base.NewMutables(maintainedTraits), true, true, true))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.MaintainersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, toID, classificationID,
			maintainedTraits, true, true, true)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - non-maintainer adding maintainer", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.NotAuthorized)
		if got := keepers.MaintainersKeeper.Transact(context, newMessage(defaultAddr, fakeFromID, toID, classificationID,
			maintainedTraits, true, true, true)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - verify identity fail", func(t *testing.T) {
		want := newTransactionResponse(errors.MockError)
		if got := keepers.MaintainersKeeper.Transact(context, newMessage(verifyMockErrorAddress, defaultIdentityID, toID, classificationID,
			maintainedTraits, true, true, true)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
	t.Run("NegativeCase - ReAdd same maintainer", func(t *testing.T) {
		want := newTransactionResponse(errors.EntityAlreadyExists)
		if got := keepers.MaintainersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, toID, classificationID,
			maintainedTraits, true, true, true)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
	t.Run("NegativeCase - conform mock error", func(t *testing.T) {
		want := newTransactionResponse(errors.MockError)
		if got := keepers.MaintainersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, toID2, classificationID,
			conformMockErrorTraits, true, true, true)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
}
