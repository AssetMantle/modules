// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"context"
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/go/data/utilities"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/metas/key"
	"github.com/AssetMantle/modules/x/metas/mappable"
	"github.com/AssetMantle/modules/x/metas/parameters"
)

type TestKeepers struct {
	MetasKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (context.Context, TestKeepers) {
	var legacyAmino = baseHelpers.CodecPrototype().GetLegacyAmino()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	encodingConfig := simapp.MakeTestEncodingConfig()
	appCodec := encodingConfig.Marshaler
	ParamsKeeper := paramsKeeper.NewKeeper(
		appCodec,
		legacyAmino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	parameterManager := parameters.Prototype().Initialize(ParamsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		MetasKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.TransactionKeeper),
	}

	return sdkTypes.WrapSDKContext(context), keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers := CreateTestInput(t)
	defaultAddr := sdkTypes.AccAddress("addr")
	data, err := utilities.ReadData("S|default")
	require.Equal(t, nil, err)
	newFact, err := utilities.ReadData("S|newFact")
	require.Equal(t, nil, err)
	keepers.MetasKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewMappable(data))
	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse()
		if got := keepers.MetasKeeper.Transact(context, NewMessage(defaultAddr, newFact).(*Message)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Reveal metas again", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse()
		if got := keepers.MetasKeeper.Transact(context, NewMessage(defaultAddr, data).(*Message)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
