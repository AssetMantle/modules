/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package take

import (
	"reflect"
	"testing"

	"github.com/persistenceOne/persistenceSDK/constants/test"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/transfer"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
)

type TestKeepers struct {
	OrdersKeeper helpers.TransactionKeeper
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
		Height:  100,
	}, false, log.NewNopLogger())

	scrubAuxiliary := scrub.AuxiliaryMock.Initialize(Mapper, Parameters)
	transferAuxiliary := transfer.AuxiliaryMock.Initialize(Mapper, Parameters)
	verifyAuxiliary := verify.AuxiliaryMock.Initialize(Mapper, Parameters)
	supplementAuxiliary := supplement.AuxiliaryMock.Initialize(Mapper, Parameters)
	keepers := TestKeepers{
		OrdersKeeper: keeperPrototype().Initialize(Mapper, Parameters,
			[]interface{}{scrubAuxiliary, verifyAuxiliary,
				transferAuxiliary, supplementAuxiliary}).(helpers.TransactionKeeper),
	}

	return context, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {

	context, keepers := CreateTestInput(t)
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")
	defaultAddr := sdkTypes.AccAddress("addr")
	defaultIdentityID := base.NewID("fromID")
	classificationID := base.NewID("classificationID")
	makerOwnableID := base.NewID("makerOwnableID")
	takerOwnableID := base.NewID("takerOwnableID")
	rateID := base.NewID(sdkTypes.OneDec().MulInt64(2).Quo(sdkTypes.SmallestDec()).Quo(sdkTypes.SmallestDec()).String())
	creationID := base.NewID("100")
	orderID := key.NewOrderID(classificationID, makerOwnableID,
		takerOwnableID, rateID, creationID, defaultIdentityID, base.NewProperties())
	metaProperties, Error := base.ReadMetaProperties(properties.MakerOwnableSplit + ":D|0.000000000000000001" +
		"," + properties.TakerID + ":I|fromID")
	require.Equal(t, nil, Error)

	keepers.OrdersKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewOrder(orderID, base.NewProperties(), metaProperties.RemoveData()))

	t.Run("PositiveCase", func(t *testing.T) {
		metaProperties, Error := base.ReadMetaProperties(properties.MakerOwnableSplit + ":D|0.000000000000000001" +
			"," + properties.TakerID + ":I|fromID")
		require.Equal(t, nil, Error)
		keepers.OrdersKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewOrder(orderID, base.NewProperties(), metaProperties.RemoveData()))

		want := newTransactionResponse(nil)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, sdkTypes.OneDec().MulInt64(2),
			orderID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - Identity mock error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.EntityNotFound)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(verifyMockErrorAddress, defaultIdentityID, sdkTypes.SmallestDec(), orderID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - No order", func(t *testing.T) {
		t.Parallel()

		want := newTransactionResponse(errors.EntityNotFound)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, sdkTypes.SmallestDec(),
			base.NewID("orderID"))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - transfer mock fail", func(t *testing.T) {
		t.Parallel()
		transferErrorID := key.NewOrderID(classificationID, makerOwnableID,
			base.NewID("transferError"), rateID, creationID, defaultIdentityID, base.NewProperties())
		metaProperties, Error := base.ReadMetaProperties(properties.MakerOwnableSplit + ":D|0.000000000000000001" +
			"," + properties.TakerID + ":I|fromID")
		require.Equal(t, nil, Error)

		keepers.OrdersKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewOrder(transferErrorID, base.NewProperties(), metaProperties.RemoveData()))

		want := newTransactionResponse(test.MockError)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, sdkTypes.SmallestDec(),
			transferErrorID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - transfer mock fail", func(t *testing.T) {
		t.Parallel()
		transferErrorID := key.NewOrderID(classificationID, base.NewID("transferError"),
			takerOwnableID, rateID, creationID, defaultIdentityID, base.NewProperties())
		metaProperties, Error := base.ReadMetaProperties(properties.MakerOwnableSplit + ":D|0.000000000000000001" +
			"," + properties.TakerID + ":I|fromID" + "," +
			properties.ExchangeRate + ":D|1")
		require.Equal(t, nil, Error)

		keepers.OrdersKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewOrder(transferErrorID, base.NewProperties(), metaProperties.RemoveData()))

		want := newTransactionResponse(test.MockError)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, sdkTypes.SmallestDec(),
			transferErrorID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - transfer mock fail", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.NotAuthorized)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, base.NewID("id"), sdkTypes.SmallestDec(), orderID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("Positive Case - take more than make order", func(t *testing.T) {
		t.Parallel()
		orderID := key.NewOrderID(classificationID, makerOwnableID,
			takerOwnableID, rateID, creationID, defaultIdentityID, base.NewProperties())
		metaProperties, Error := base.ReadMetaProperties(properties.MakerOwnableSplit + ":D|0.000000000000000001" +
			"," + properties.TakerID + ":I|fromID" + "," +
			properties.ExchangeRate + ":D|1")
		require.Equal(t, nil, Error)

		keepers.OrdersKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewOrder(orderID, base.NewProperties(), metaProperties.RemoveData()))

		want := newTransactionResponse(nil)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, sdkTypes.SmallestDec().MulInt64(1),
			orderID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
		keepers.OrdersKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewOrder(orderID, base.NewProperties(), metaProperties.RemoveData()))
		want = newTransactionResponse(nil)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, sdkTypes.SmallestDec().MulInt64(10),
			orderID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
