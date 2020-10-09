package cancel

import (
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
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

	transfer.AuxiliaryMock.InitializeKeeper(mapper.Mapper, parameters.Prototype)
	verify.AuxiliaryMock.InitializeKeeper(mapper.Mapper, parameters.Prototype)
	supplement.AuxiliaryMock.InitializeKeeper(mapper.Mapper, parameters.Prototype)
	keepers := TestKeepers{
		OrdersKeeper: initializeTransactionKeeper(mapper.Mapper, parameters.Prototype,
			[]interface{}{verify.AuxiliaryMock,
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
	metaProperties, Error := base.ReadMetaProperties(properties.MakerOwnableSplit + ":D|0.000000000000000001" +
		"," + properties.TakerID + ":I|fromID" + "," +
		properties.ExchangeRate + ":D|0.000000000000000001")
	require.Equal(t, nil, Error)
	orderID := mapper.NewOrderID(classificationID, makerOwnableID,
		takerOwnableID, defaultIdentityID, base.NewImmutables(base.NewProperties()))
	mapper.NewOrders(mapper.Mapper, ctx).Add(mapper.NewOrder(orderID, base.NewImmutables(base.NewProperties()), base.NewMutables(metaProperties)))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, orderID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - Identity mock error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(verifyMockErrorAddress, defaultIdentityID, orderID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - Cancel not existing order", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.EntityNotFound)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, base.NewID("orderID"))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - Cancel with different makerID", func(t *testing.T) {
		t.Parallel()
		mapper.NewOrders(mapper.Mapper, ctx).Add(mapper.NewOrder(orderID, base.NewImmutables(base.NewProperties()), base.NewMutables(metaProperties)))
		want := newTransactionResponse(errors.NotAuthorized)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, base.NewID("id"), orderID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
	t.Run("NegativeCase - transferMock Error", func(t *testing.T) {
		t.Parallel()
		transferErrorID := mapper.NewOrderID(classificationID, base.NewID("transferError"),
			takerOwnableID, defaultIdentityID, base.NewImmutables(base.NewProperties()))
		mapper.NewOrders(mapper.Mapper, ctx).Add(mapper.NewOrder(transferErrorID, base.NewImmutables(base.NewProperties()), base.NewMutables(metaProperties)))

		want := newTransactionResponse(errors.MockError)
		if got := keepers.OrdersKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, transferErrorID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
