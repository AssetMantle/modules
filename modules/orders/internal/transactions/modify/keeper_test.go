/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package modify

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
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/mint"
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
	}, false, log.NewNopLogger())

	scrubAuxiliary := scrub.AuxiliaryMock.Initialize(Mapper, Parameters)
	conformAuxiliary := conform.AuxiliaryMock.Initialize(Mapper, Parameters)
	transferAuxiliary := transfer.AuxiliaryMock.Initialize(Mapper, Parameters)
	verifyAuxiliary := verify.AuxiliaryMock.Initialize(Mapper, Parameters)
	mintAuxiliary := mint.AuxiliaryMock.Initialize(Mapper, Parameters)
	supplementAuxiliary := supplement.AuxiliaryMock.Initialize(Mapper, Parameters)
	keepers := TestKeepers{
		OrdersKeeper: keeperPrototype().Initialize(Mapper, Parameters,
			[]interface{}{scrubAuxiliary, verifyAuxiliary,
				conformAuxiliary, transferAuxiliary, mintAuxiliary, supplementAuxiliary}).(helpers.TransactionKeeper),
	}

	return context, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {

	context, keepers := CreateTestInput(t)
	immutableProperties, Error := base.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, Error)
	mutableProperties, Error := base.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, Error)
	mutablePropertiesUpdated, Error := base.ReadMetaProperties("defaultMutable1:S|defaultMutable2")
	require.Equal(t, nil, Error)
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")
	defaultAddr := sdkTypes.AccAddress("addr")
	defaultIdentityID := base.NewID("fromID")
	classificationID := base.NewID("classificationID")
	makerOwnableID := base.NewID("makerOwnableID")
	rateID := base.NewID(sdkTypes.OneDec().String())
	updatedRate := sdkTypes.MustNewDecFromStr("0.002")
	creationID := base.NewID("100")
	//makerID := base.NewID("makerID")
	takerOwnableID := base.NewID("takerOwnableID")
	makerOwnableSplit := sdkTypes.SmallestDec().MulInt64(2)
	orderID := key.NewOrderID(classificationID, makerOwnableID, takerOwnableID, rateID, creationID, defaultIdentityID, immutableProperties)
	keepers.OrdersKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewOrder(orderID, immutableProperties, mutableProperties))

	t.Run("PositiveCase Modifying Order - Changing TakerOwnableSplit, mutableProperty and subtracting makerOwnableSplit", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, orderID,
			updatedRate, makerOwnableSplit.Sub(sdkTypes.SmallestDec()), base.NewHeight(100),
			base.NewMetaProperties(), mutablePropertiesUpdated.RemoveData())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	keepers.OrdersKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewOrder(orderID, immutableProperties, mutableProperties))
	t.Run("PositiveCase Modifying Order - Changing TakerOwnableSplit, mutableProperty and adding makerOwnableSplit", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, orderID,
			updatedRate, makerOwnableSplit.Add(sdkTypes.SmallestDec()), base.NewHeight(100),
			base.NewMetaProperties(), mutablePropertiesUpdated.RemoveData())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase Modifying Order - Order not found", func(t *testing.T) {
		want := newTransactionResponse(errors.EntityNotFound)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, base.NewID("orderID"),
			updatedRate, makerOwnableSplit, base.NewHeight(100),
			base.NewMetaProperties(), mutablePropertiesUpdated.RemoveData())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - Identity mock error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(test.MockError)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(verifyMockErrorAddress, defaultIdentityID, orderID,
			updatedRate, makerOwnableSplit, base.NewHeight(100), base.NewMetaProperties(),
			mutablePropertiesUpdated.RemoveData())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	orderID = key.NewOrderID(classificationID, base.NewID("transferError"), takerOwnableID, rateID, creationID, defaultIdentityID, immutableProperties)
	keepers.OrdersKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewOrder(orderID, immutableProperties, mutableProperties))
	t.Run("NegativeCase Modifying Order - transferError", func(t *testing.T) {
		want := newTransactionResponse(test.MockError)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, orderID,
			updatedRate, makerOwnableSplit.Sub(sdkTypes.SmallestDec()), base.NewHeight(100),
			base.NewMetaProperties(), mutablePropertiesUpdated.RemoveData())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase Modifying Order - Changing TakerOwnableSplit, mutableProperty and adding makerOwnableSplit", func(t *testing.T) {
		want := newTransactionResponse(test.MockError)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, orderID,
			updatedRate, makerOwnableSplit.Add(sdkTypes.SmallestDec()), base.NewHeight(100),
			base.NewMetaProperties(), mutablePropertiesUpdated.RemoveData())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
