// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package block

import (
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

	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/mappable"
	"github.com/AssetMantle/modules/modules/orders/internal/parameters"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/transfer"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
)

func CreateTestInput(t *testing.T) (sdkTypes.Context, helpers.Mapper, helpers.Auxiliary, helpers.Auxiliary) {
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

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := params.NewKeeper(
		Codec,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters := parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))
	transferAuxiliary := transfer.AuxiliaryMock.Initialize(Mapper, Parameters)
	supplementAuxiliary := supplement.AuxiliaryMock.Initialize(Mapper, Parameters)

	context := sdkTypes.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
		Height:  1000,
	}, false, log.NewNopLogger())

	return context, Mapper, transferAuxiliary, supplementAuxiliary
}

func Test_Block_Methods(t *testing.T) {
	block := Prototype()
	context, mapper, transferAuxiliary, supplementAuxiliary := CreateTestInput(t)
	block = block.Initialize(mapper, parameters.Prototype(), transferAuxiliary, supplementAuxiliary)
	block.Begin(context, abciTypes.RequestBeginBlock{})

	defaultIdentityID := baseIDs.NewStringID("fromID")
	classificationID := baseIDs.NewStringID("classificationID")
	makerOwnableID := baseIDs.NewStringID("makerOwnableID")
	takerOwnableID := baseIDs.NewStringID("takerOwnableID")
	rateID := baseIDs.NewStringID(sdkTypes.OneDec().String())
	creationID := baseIDs.NewStringID("100")
	orderID := key.NewOrderID(classificationID, makerOwnableID, takerOwnableID, rateID, creationID, defaultIdentityID, baseLists.NewPropertyList())
	orders := mapper.NewCollection(context)
	orders.Add(mappable.NewOrder(orderID, baseLists.NewPropertyList(), baseLists.NewPropertyList()))
	block.End(context, abciTypes.RequestEndBlock{})
}
