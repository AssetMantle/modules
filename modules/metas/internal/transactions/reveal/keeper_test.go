/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package reveal

import (
	"reflect"
	"testing"

	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/persistenceOne/persistenceSDK/schema/applications"
	tendermintProto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
)

type TestKeepers struct {
	MetasKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {

	var Codec = codec.NewLegacyAmino()
	schema.RegisterCodec(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)

	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)
	std.RegisterLegacyAminoCodec(Codec)
	std.RegisterInterfaces(interfaceRegistry)

	encodingConfig := applications.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             Codec,
	}

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := paramsKeeper.NewKeeper(
		encodingConfig.Marshaler,
		encodingConfig.Amino,
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

	context := sdkTypes.NewContext(commitMultiStore, tendermintProto.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		MetasKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.TransactionKeeper),
	}

	return context, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {

	context, keepers := CreateTestInput(t)
	defaultAddr := sdkTypes.AccAddress("addr")
	defaultFact, Error := base.ReadMetaFact("S|default")
	require.Equal(t, nil, Error)
	newFact, Error := base.ReadMetaFact("S|newFact")
	require.Equal(t, nil, Error)
	keepers.MetasKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewMeta(defaultFact.GetData()))
	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.MetasKeeper.Transact(context, newMessage(defaultAddr, newFact)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Reveal metas again", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.EntityAlreadyExists)
		if got := keepers.MetasKeeper.Transact(context, newMessage(defaultAddr, defaultFact)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
