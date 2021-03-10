/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package wrap

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/test"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/applications"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tendermintProto "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"reflect"
	"testing"
)

type TestKeepers struct {
	SplitsKeeper  helpers.TransactionKeeper
	AccountKeeper authKeeper.AccountKeeper
	BankKeeper    bankKeeper.Keeper
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
	authStoreKey := sdkTypes.NewKVStoreKey("testAuth")
	bankStoreKey := sdkTypes.NewKVStoreKey("testBank")
	mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
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
	commitMultiStore.MountStoreWithDB(authStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(bankStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	Error := commitMultiStore.LoadLatestVersion()
	require.Nil(t, Error)

	context := sdkTypes.NewContext(commitMultiStore, tendermintProto.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	accountKeeper := authKeeper.NewAccountKeeper(encodingConfig.Marshaler, authStoreKey, paramsKeeper.Subspace(authTypes.ModuleName), authTypes.ProtoBaseAccount, make(map[string][]string))

	bankKeeper := bankKeeper.NewBaseKeeper(encodingConfig.Marshaler, bankStoreKey, accountKeeper, paramsKeeper.Subspace(bankTypes.ModuleName), make(map[string]bool))
	verifyAuxiliary := verify.AuxiliaryMock.Initialize(mapper, Parameters)
	keepers := TestKeepers{
		SplitsKeeper: keeperPrototype().Initialize(mapper, Parameters,
			[]interface{}{verifyAuxiliary}).(helpers.TransactionKeeper),
		AccountKeeper: accountKeeper,
		BankKeeper:    bankKeeper,
	}

	return context, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {

	ctx, keepers := CreateTestInput(t)
	defaultAddr := sdkTypes.AccAddress("addr")
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")

	fromID := base.NewID("fromID")
	coins := func(amount int64) sdkTypes.Coins {
		return sdkTypes.NewCoins(sdkTypes.NewCoin("stake", sdkTypes.NewInt(amount)))
	}

	Error := keepers.BankKeeper.SetBalances(ctx, defaultAddr, coins(1000))
	require.Equal(t, nil, Error)
	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.SplitsKeeper.Transact(ctx, newMessage(defaultAddr, fromID, coins(100))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("PositiveCase- reAdd", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(nil)
		if got := keepers.SplitsKeeper.Transact(ctx, newMessage(defaultAddr, fromID, coins(100))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Verify Identity Failure", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(test.MockError)
		if got := keepers.SplitsKeeper.Transact(ctx, newMessage(verifyMockErrorAddress, fromID, coins(100))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Wrap Negative coins", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.InsufficientBalance)
		if got := keepers.SplitsKeeper.Transact(ctx, newMessage(defaultAddr, fromID, coins(10000))); !reflect.DeepEqual(got.IsSuccessful(), want.IsSuccessful()) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
