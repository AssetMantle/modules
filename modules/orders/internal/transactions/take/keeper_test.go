/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package take

import (
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/parameters"
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

	scrub.AuxiliaryMock.InitializeKeeper(mapper.Mapper, parameters.Prototype)
	transfer.AuxiliaryMock.InitializeKeeper(mapper.Mapper, parameters.Prototype)
	verify.AuxiliaryMock.InitializeKeeper(mapper.Mapper, parameters.Prototype)
	supplement.AuxiliaryMock.InitializeKeeper(mapper.Mapper, parameters.Prototype)
	keepers := TestKeepers{
		OrdersKeeper: initializeTransactionKeeper(mapper.Mapper, parameters.Prototype,
			[]interface{}{scrub.AuxiliaryMock, verify.AuxiliaryMock,
				transfer.AuxiliaryMock, supplement.AuxiliaryMock}),
	}

	return ctx, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {

	ctx, keepers := CreateTestInput(t)
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")
	defaultAddr := sdkTypes.AccAddress("addr")
	defaultIdentityID := base.NewID("fromID")
	classificationID := base.NewID("classificationID")
	makerOwnableID := base.NewID("makerOwnableID")
	takerOwnableID := base.NewID("takerOwnableID")
	orderID := mapper.NewOrderID(classificationID, makerOwnableID,
		takerOwnableID, defaultIdentityID, base.NewImmutables(base.NewProperties()))
	nontakingOrderID := mapper.NewOrderID(base.NewID(""), makerOwnableID,
		takerOwnableID, defaultIdentityID, base.NewImmutables(base.NewProperties()))
	metaProperties, Error := base.ReadMetaProperties(properties.MakerOwnableSplit + ":D|0.000000000000000001" +
		"," + properties.TakerID + ":I|fromID" + "," +
		properties.ExchangeRate + ":D|0.000000000000000001")
	require.Equal(t, nil, Error)

	mapper.NewOrders(mapper.Mapper, ctx).Add(mapper.NewOrder(nontakingOrderID, base.NewImmutables(base.NewProperties()), base.NewMutables(metaProperties)))

	t.Run("PositiveCase", func(t *testing.T) {
		metaProperties, Error := base.ReadMetaProperties(properties.MakerOwnableSplit + ":D|0.000000000000000001" +
			"," + properties.TakerID + ":I|fromID" + "," +
			properties.ExchangeRate + ":D|0.000000000000000001")
		require.Equal(t, nil, Error)
		mapper.NewOrders(mapper.Mapper, ctx).Add(mapper.NewOrder(orderID, base.NewImmutables(base.NewProperties()), base.NewMutables(metaProperties)))

		want := newTransactionResponse(nil)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, sdkTypes.SmallestDec(),
			orderID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - Identity mock error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.EntityNotFound)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(verifyMockErrorAddress, defaultIdentityID, sdkTypes.SmallestDec(),
			nontakingOrderID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - No order", func(t *testing.T) {
		t.Parallel()

		want := newTransactionResponse(errors.EntityNotFound)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, sdkTypes.SmallestDec(),
			base.NewID("orderID"))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - transfer mock fail", func(t *testing.T) {
		t.Parallel()
		transferErrorID := mapper.NewOrderID(classificationID, makerOwnableID,
			base.NewID("transferError"), defaultIdentityID, base.NewImmutables(base.NewProperties()))
		metaProperties, Error := base.ReadMetaProperties(properties.MakerOwnableSplit + ":D|0.000000000000000001" +
			"," + properties.TakerID + ":I|fromID" + "," +
			properties.ExchangeRate + ":D|0.000000000000000001")
		require.Equal(t, nil, Error)

		mapper.NewOrders(mapper.Mapper, ctx).Add(mapper.NewOrder(transferErrorID, base.NewImmutables(base.NewProperties()), base.NewMutables(metaProperties)))

		want := newTransactionResponse(errors.MockError)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, sdkTypes.SmallestDec(),
			transferErrorID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - transfer mock fail", func(t *testing.T) {
		t.Parallel()
		transferErrorID := mapper.NewOrderID(classificationID, base.NewID("transferError"),
			takerOwnableID, defaultIdentityID, base.NewImmutables(base.NewProperties()))
		metaProperties, Error := base.ReadMetaProperties(properties.MakerOwnableSplit + ":D|0.000000000000000001" +
			"," + properties.TakerID + ":I|fromID" + "," +
			properties.ExchangeRate + ":D|0.000000000000000001")
		require.Equal(t, nil, Error)

		mapper.NewOrders(mapper.Mapper, ctx).Add(mapper.NewOrder(transferErrorID, base.NewImmutables(base.NewProperties()), base.NewMutables(metaProperties)))

		want := newTransactionResponse(errors.MockError)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, sdkTypes.SmallestDec(),
			transferErrorID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - transfer mock fail", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.NotAuthorized)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, base.NewID("id"), sdkTypes.SmallestDec(),
			nontakingOrderID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - take more than make order", func(t *testing.T) {
		t.Parallel()
		orderID := mapper.NewOrderID(classificationID, makerOwnableID,
			takerOwnableID, defaultIdentityID, base.NewImmutables(base.NewProperties()))
		metaProperties, Error := base.ReadMetaProperties(properties.MakerOwnableSplit + ":D|0.000000000000000001" +
			"," + properties.TakerID + ":I|fromID" + "," +
			properties.ExchangeRate + ":D|0.000000000000000001")
		require.Equal(t, nil, Error)

		mapper.NewOrders(mapper.Mapper, ctx).Add(mapper.NewOrder(orderID, base.NewImmutables(base.NewProperties()), base.NewMutables(metaProperties)))

		want := newTransactionResponse(nil)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, sdkTypes.SmallestDec().MulInt64(1),
			nontakingOrderID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}

		want = newTransactionResponse(nil)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, sdkTypes.SmallestDec().MulInt64(10),
			nontakingOrderID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
