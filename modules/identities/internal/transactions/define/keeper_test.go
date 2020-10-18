/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package define

import (
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/define"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/super"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	"reflect"
	"testing"
)

type TestKeepers struct {
	IdentitiesKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {

	keyIdentity := mapper.Mapper.GetKVStoreKey()

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyIdentity, sdkTypes.StoreTypeIAVL, db)
	err := ms.LoadLatestVersion()
	require.Nil(t, err)

	ctx := sdkTypes.NewContext(ms, abci.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	scrub.AuxiliaryMock.Initialize(mapper.Mapper, parameters.Prototype)
	define.AuxiliaryMock.Initialize(mapper.Mapper, parameters.Prototype)
	super.AuxiliaryMock.Initialize(mapper.Mapper, parameters.Prototype)
	keepers := TestKeepers{
		IdentitiesKeeper: initializeTransactionKeeper(mapper.Mapper, parameters.Prototype,
			[]interface{}{scrub.AuxiliaryMock,
				define.AuxiliaryMock, super.AuxiliaryMock}),
	}

	return ctx, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {

	ctx, keepers := CreateTestInput(t)
	immutableMetaTraits, Error := base.ReadMetaProperties("defaultImmutableMeta1:S|defaultImmutableMeta1")
	require.Equal(t, nil, Error)
	immutableTraits, Error := base.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, Error)
	mutableMetaTraits, Error := base.ReadMetaProperties("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, Error)
	mutableTraits, Error := base.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, Error)
	superMockErrorTraits, Error := base.ReadMetaProperties("superError:S|mockError")
	require.Equal(t, nil, Error)
	scrubMockErrorTraits, Error := base.ReadMetaProperties("scrubError:S|mockError")
	require.Equal(t, nil, Error)
	gt22Traits, Error := base.ReadMetaProperties("0:S|0,1:S|1,2:S|2,3:S|3,4:S|4,5:S|5,6:S|6,7:S|7,8:S|8,9:S|9,10:S|10,11:S|11,12:S|12,13:S|13,14:S|14,15:S|15,16:S|16,17:S|17,18:S|18,19:S|19,20:S|20,21:S|21")
	require.Equal(t, nil, Error)
	defaultAddr := sdkTypes.AccAddress("addr")
	defaultIdentityID := key.NewIdentityID(base.NewID("test.cGn3HMW8M3t5gMDv-wXa9sseHnA="), base.NewID("d0Jhri_bOd3EEPXpyPUpNpGiQ1U="))
	mapper.NewIdentities(mapper.Mapper, ctx).Add(mappable.NewIdentity(defaultIdentityID, []sdkTypes.AccAddress{sdkTypes.AccAddress("addr")},
		[]sdkTypes.AccAddress{}, base.NewImmutables(base.NewProperties()), base.NewMutables(base.NewProperties())))

	t.Run("PositiveCase", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(nil)
		if got := keepers.IdentitiesKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, immutableMetaTraits,
			immutableTraits, mutableMetaTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - IdentityNil", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.EntityNotFound)
		if got := keepers.IdentitiesKeeper.Transact(ctx, newMessage(defaultAddr, base.NewID(""), immutableMetaTraits,
			immutableTraits, mutableMetaTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - Identity unprovisioned address", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.NotAuthorized)
		if got := keepers.IdentitiesKeeper.Transact(ctx, newMessage(sdkTypes.AccAddress("addr1"), defaultIdentityID, immutableMetaTraits,
			immutableTraits, mutableMetaTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - define classification Failure", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.InvalidRequest)
		if got := keepers.IdentitiesKeeper.Transact(ctx, newMessage(sdkTypes.AccAddress("addr"), defaultIdentityID, gt22Traits,
			immutableTraits, mutableMetaTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - Metas Scrub Immutable Failure", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.IdentitiesKeeper.Transact(ctx, newMessage(sdkTypes.AccAddress("addr"), defaultIdentityID, scrubMockErrorTraits,
			immutableTraits, mutableMetaTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - Metas Scrub Mutables Failure", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.IdentitiesKeeper.Transact(ctx, newMessage(sdkTypes.AccAddress("addr"), defaultIdentityID, immutableMetaTraits,
			immutableTraits, scrubMockErrorTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - Maintainer super Failure", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.IdentitiesKeeper.Transact(ctx, newMessage(sdkTypes.AccAddress("addr"), defaultIdentityID, immutableMetaTraits,
			immutableTraits, mutableMetaTraits, superMockErrorTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
