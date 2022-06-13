// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

import (
	"reflect"
	"testing"

	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/mappable"
	"github.com/AssetMantle/modules/modules/orders/internal/parameters"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/mint"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/transfer"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
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
	immutableProperties, Error := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, Error)
	mutableProperties, Error := utilities.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, Error)
	mutablePropertiesUpdated, Error := utilities.ReadMetaProperties("defaultMutable1:S|defaultMutable2")
	require.Equal(t, nil, Error)
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")
	defaultAddr := sdkTypes.AccAddress("addr")
	defaultIdentityID := baseIDs.NewID("fromID")
	classificationID := baseIDs.NewID("classificationID")
	makerOwnableID := baseIDs.NewID("makerOwnableID")
	rateID := baseIDs.NewID(sdkTypes.OneDec().String())
	updatedRate := sdkTypes.MustNewDecFromStr("0.002")
	creationID := baseIDs.NewID("100")
	// makerID := baseIDs.NewID("makerID")
	takerOwnableID := baseIDs.NewID("takerOwnableID")
	makerOwnableSplit := sdkTypes.SmallestDec().MulInt64(2)
	orderID := key.NewOrderID(classificationID, makerOwnableID, takerOwnableID, rateID, creationID, defaultIdentityID, immutableProperties)
	keepers.OrdersKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewOrder(orderID, immutableProperties, mutableProperties))

	t.Run("PositiveCase Modifying Order - Changing TakerOwnableSplit, mutableProperty and subtracting makerOwnableSplit", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, orderID,
			updatedRate, makerOwnableSplit.Sub(sdkTypes.SmallestDec()), baseTypes.NewHeight(100),
			base.NewMetaProperties(), mutablePropertiesUpdated.ToPropertyList())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	keepers.OrdersKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewOrder(orderID, immutableProperties, mutableProperties))
	t.Run("PositiveCase Modifying Order - Changing TakerOwnableSplit, mutableProperty and adding makerOwnableSplit", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, orderID,
			updatedRate, makerOwnableSplit.Add(sdkTypes.SmallestDec()), baseTypes.NewHeight(100),
			base.NewMetaProperties(), mutablePropertiesUpdated.ToPropertyList())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase Modifying Order - Order not found", func(t *testing.T) {
		want := newTransactionResponse(errors.EntityNotFound)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, baseIDs.NewID("orderID"),
			updatedRate, makerOwnableSplit, baseTypes.NewHeight(100),
			base.NewMetaProperties(), mutablePropertiesUpdated.ToPropertyList())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - Identity mock error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(verifyMockErrorAddress, defaultIdentityID, orderID,
			updatedRate, makerOwnableSplit, baseTypes.NewHeight(100), base.NewMetaProperties(),
			mutablePropertiesUpdated.ToPropertyList())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	orderID = key.NewOrderID(classificationID, baseIDs.NewID("transferError"), takerOwnableID, rateID, creationID, defaultIdentityID, immutableProperties)
	keepers.OrdersKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewOrder(orderID, immutableProperties, mutableProperties))
	t.Run("NegativeCase Modifying Order - transferError", func(t *testing.T) {
		want := newTransactionResponse(errors.MockError)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, orderID,
			updatedRate, makerOwnableSplit.Sub(sdkTypes.SmallestDec()), baseTypes.NewHeight(100),
			base.NewMetaProperties(), mutablePropertiesUpdated.ToPropertyList())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase Modifying Order - Changing TakerOwnableSplit, mutableProperty and adding makerOwnableSplit", func(t *testing.T) {
		want := newTransactionResponse(errors.MockError)
		if got := keepers.OrdersKeeper.Transact(context, newMessage(defaultAddr, defaultIdentityID, orderID,
			updatedRate, makerOwnableSplit.Add(sdkTypes.SmallestDec()), baseTypes.NewHeight(100),
			base.NewMetaProperties(), mutablePropertiesUpdated.ToPropertyList())); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
