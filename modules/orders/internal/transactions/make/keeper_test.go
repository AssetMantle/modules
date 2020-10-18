/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package make

import (
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/mint"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/transfer"
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
	OrdersKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {

	keyOrder := mapper.Mapper.GetKVStoreKey()

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyOrder, sdkTypes.StoreTypeIAVL, db)
	err := ms.LoadLatestVersion()
	require.Nil(t, err)

	ctx := sdkTypes.NewContext(ms, abci.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	scrub.AuxiliaryMock.Initialize(mapper.Mapper, parameters.Prototype)
	conform.AuxiliaryMock.Initialize(mapper.Mapper, parameters.Prototype)
	transfer.AuxiliaryMock.Initialize(mapper.Mapper, parameters.Prototype)
	verify.AuxiliaryMock.Initialize(mapper.Mapper, parameters.Prototype)
	mint.AuxiliaryMock.Initialize(mapper.Mapper, parameters.Prototype)
	supplement.AuxiliaryMock.Initialize(mapper.Mapper, parameters.Prototype)
	keepers := TestKeepers{
		OrdersKeeper: initializeTransactionKeeper(mapper.Mapper, parameters.Prototype,
			[]interface{}{scrub.AuxiliaryMock, verify.AuxiliaryMock,
				conform.AuxiliaryMock, transfer.AuxiliaryMock, mint.AuxiliaryMock, supplement.AuxiliaryMock}),
	}

	return ctx, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {

	ctx, keepers := CreateTestInput(t)
	immutableMetaTraits, Error := base.ReadMetaProperties("defaultImmutableMeta1:S|defaultImmutableMeta1")
	require.Equal(t, nil, Error)
	immutableTraits, Error := base.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, Error)
	mutableMetaTraits, Error := base.ReadMetaProperties("makerOwnableSplit:D|1")
	require.Equal(t, nil, Error)
	mutableTraits, Error := base.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, Error)
	conformMockErrorTraits, Error := base.ReadMetaProperties("conformError:S|mockError")
	require.Equal(t, nil, Error)
	scrubMockErrorTraits, Error := base.ReadMetaProperties("scrubError:S|mockError")
	require.Equal(t, nil, Error)
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")
	defaultAddr := sdkTypes.AccAddress("addr")
	defaultIdentityID := base.NewID("fromID")
	classificationID := base.NewID("classificationID")
	makerOwnableID := base.NewID("makerOwnableID")
	takerOwnableID := base.NewID("takerOwnableID")
	orderID := key.NewOrderID(classificationID, makerOwnableID,
		takerOwnableID, defaultIdentityID, base.NewImmutables(base.NewProperties()))
	mapper.NewOrders(mapper.Mapper, ctx).Add(mappable.NewOrder(orderID, base.NewImmutables(base.NewProperties()), base.NewMutables(base.NewProperties())))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, classificationID,
			makerOwnableID, takerOwnableID, base.NewHeight(0), sdkTypes.SmallestDec(),
			immutableMetaTraits, immutableTraits, mutableMetaTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
	t.Run("PositiveCase- ReAdd order", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, classificationID,
			makerOwnableID, takerOwnableID, base.NewHeight(0), sdkTypes.SmallestDec(),
			immutableMetaTraits, immutableTraits, mutableMetaTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - Identity mock error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(verifyMockErrorAddress, defaultIdentityID, classificationID,
			makerOwnableID, takerOwnableID, base.NewHeight(0), sdkTypes.SmallestDec(),
			immutableMetaTraits, immutableTraits, mutableMetaTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - conform mock fail", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, classificationID,
			makerOwnableID, takerOwnableID, base.NewHeight(0), sdkTypes.SmallestDec(),
			immutableMetaTraits, immutableTraits, mutableMetaTraits, conformMockErrorTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - mutables scrub mock fail", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, classificationID,
			makerOwnableID, takerOwnableID, base.NewHeight(0), sdkTypes.SmallestDec(),
			immutableMetaTraits, immutableTraits, scrubMockErrorTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - immutables scrub mock fail", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, classificationID,
			makerOwnableID, takerOwnableID, base.NewHeight(0), sdkTypes.SmallestDec(),
			scrubMockErrorTraits, immutableTraits, mutableMetaTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - Transfer mock fail", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, classificationID,
			base.NewID("transferError"), takerOwnableID, base.NewHeight(0), sdkTypes.SmallestDec(),
			immutableMetaTraits, immutableTraits, mutableMetaTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
