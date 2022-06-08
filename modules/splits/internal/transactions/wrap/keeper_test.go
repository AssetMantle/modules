// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrap

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/splits/internal/key"
	"github.com/AssetMantle/modules/modules/splits/internal/mappable"
	"github.com/AssetMantle/modules/modules/splits/internal/module"
	"github.com/AssetMantle/modules/modules/splits/internal/parameters"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

type TestKeepers struct {
	SplitsKeeper  helpers.TransactionKeeper
	AccountKeeper auth.AccountKeeper
	BankKeeper    bank.Keeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	supply.RegisterCodec(Codec)
	params.RegisterCodec(Codec)
	auth.RegisterCodec(Codec)
	Codec.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	authStoreKey := sdkTypes.NewKVStoreKey("testAuth")
	supplyStoreKey := sdkTypes.NewKVStoreKey("testSupply")
	mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
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
	commitMultiStore.MountStoreWithDB(authStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(supplyStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	accountKeeper := auth.NewAccountKeeper(Codec, authStoreKey, paramsKeeper.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)

	bankKeeper := bank.NewBaseKeeper(accountKeeper, paramsKeeper.Subspace(bank.DefaultParamspace), make(map[string]bool))
	supplyKeeper := supply.NewKeeper(Codec, supplyStoreKey, accountKeeper, bankKeeper, map[string][]string{module.Name: nil})
	verifyAuxiliary := verify.AuxiliaryMock.Initialize(mapper, Parameters)
	keepers := TestKeepers{
		SplitsKeeper: keeperPrototype().Initialize(mapper, Parameters,
			[]interface{}{verifyAuxiliary, supplyKeeper}).(helpers.TransactionKeeper),
		AccountKeeper: accountKeeper,
		BankKeeper:    bankKeeper,
	}

	return context, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {
	ctx, keepers := CreateTestInput(t)
	defaultAddr := sdkTypes.AccAddress("addr")
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")

	fromID := baseIDs.NewID("fromID")
	coins := func(amount int64) sdkTypes.Coins {
		return sdkTypes.NewCoins(sdkTypes.NewCoin("stake", sdkTypes.NewInt(amount)))
	}

	err := keepers.BankKeeper.SetCoins(ctx, defaultAddr, coins(1000))
	require.Equal(t, nil, err)
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
		want := newTransactionResponse(errors.MockError)
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
