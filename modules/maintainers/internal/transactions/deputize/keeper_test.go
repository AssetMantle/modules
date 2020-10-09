/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package deputize

import (
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/parameters"
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
	MaintainersKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {

	keyMaintainers := mapper.Mapper.GetKVStoreKey()

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyMaintainers, sdkTypes.StoreTypeIAVL, db)

	err := ms.LoadLatestVersion()
	require.Nil(t, err)

	ctx := sdkTypes.NewContext(ms, abci.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	conform.AuxiliaryMock.InitializeKeeper(mapper.Mapper, parameters.Prototype)
	verify.AuxiliaryMock.InitializeKeeper(mapper.Mapper, parameters.Prototype)
	keepers := TestKeepers{
		MaintainersKeeper: initializeTransactionKeeper(mapper.Mapper, parameters.Prototype,
			[]interface{}{verify.AuxiliaryMock, conform.AuxiliaryMock}),
	}

	return ctx, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {

	ctx, keepers := CreateTestInput(t)
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

	mapper.NewMaintainers(mapper.Mapper, ctx).Add(mapper.NewMaintainer(mapper.NewMaintainerID(classificationID, defaultIdentityID),
		base.NewMutables(maintainedTraits), true, true, true))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.MaintainersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, toID, classificationID,
			maintainedTraits, true, true, true)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - non-maintainer adding maintainer", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.NotAuthorized)
		if got := keepers.MaintainersKeeper.Transact(ctx, newMessage(defaultAddr, fakeFromID, toID, classificationID,
			maintainedTraits, true, true, true)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - verify identity fail", func(t *testing.T) {
		want := newTransactionResponse(errors.MockError)
		if got := keepers.MaintainersKeeper.Transact(ctx, newMessage(verifyMockErrorAddress, defaultIdentityID, toID, classificationID,
			maintainedTraits, true, true, true)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
	t.Run("NegativeCase - ReAdd same maintainer", func(t *testing.T) {
		want := newTransactionResponse(errors.EntityAlreadyExists)
		if got := keepers.MaintainersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, toID, classificationID,
			maintainedTraits, true, true, true)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
	t.Run("NegativeCase - conform mock error", func(t *testing.T) {
		want := newTransactionResponse(errors.MockError)
		if got := keepers.MaintainersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, toID2, classificationID,
			conformMockErrorTraits, true, true, true)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
}
