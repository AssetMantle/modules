package unwrap

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authVesting "github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/schema"
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
	SplitsKeeper  helpers.TransactionKeeper
	AccountKeeper auth.AccountKeeper
	BankKeeper    bank.Keeper
	SupplyKeeper  supply.Keeper
}

func MakeCodec() *codec.Codec {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	authVesting.RegisterCodec(Codec)
	supply.RegisterCodec(Codec)
	params.RegisterCodec(Codec)
	auth.RegisterCodec(Codec)
	return Codec
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {

	keyIdentity := mapper.Mapper.GetKVStoreKey()
	keyParams := sdkTypes.NewKVStoreKey(params.StoreKey)
	keyAuth := sdkTypes.NewKVStoreKey(auth.StoreKey)
	keySupply := sdkTypes.NewKVStoreKey(supply.StoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyIdentity, sdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, sdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyAuth, sdkTypes.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keySupply, sdkTypes.StoreTypeIAVL, db)
	err := ms.LoadLatestVersion()
	require.Nil(t, err)
	transientStoreKeys := sdkTypes.NewTransientStoreKeys(params.TStoreKey)

	ctx := sdkTypes.NewContext(ms, abci.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	Codec := MakeCodec()
	paramsKeeper := params.NewKeeper(Codec, keyParams, transientStoreKeys[params.TStoreKey])
	accountKeeper := auth.NewAccountKeeper(Codec, keyAuth, paramsKeeper.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)

	bankKeeper := bank.NewBaseKeeper(accountKeeper, paramsKeeper.Subspace(bank.DefaultParamspace), make(map[string]bool))
	supplyKeeper := supply.NewKeeper(Codec, keySupply, accountKeeper, bankKeeper, map[string][]string{mapper.ModuleName: nil})
	verify.AuxiliaryMock.InitializeKeeper(mapper.Mapper, parameters.Prototype)
	keepers := TestKeepers{
		SplitsKeeper: initializeTransactionKeeper(mapper.Mapper, parameters.Prototype,
			[]interface{}{verify.AuxiliaryMock, supplyKeeper}),
		AccountKeeper: accountKeeper,
		BankKeeper:    bankKeeper,
		SupplyKeeper:  supplyKeeper,
	}

	return ctx, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {

	ctx, keepers := CreateTestInput(t)
	defaultAddr := sdkTypes.AccAddress("addr")
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")
	ownableID := base.NewID("stake")
	fromID := base.NewID("fromID")
	coins := func(amount int64) sdkTypes.Coins {
		return sdkTypes.NewCoins(sdkTypes.NewCoin("stake", sdkTypes.NewInt(amount)))
	}
	Error := keepers.BankKeeper.SetCoins(ctx, defaultAddr, coins(1000))
	require.Equal(t, nil, Error)
	Error = keepers.SupplyKeeper.SendCoinsFromAccountToModule(ctx, defaultAddr, mapper.ModuleName, coins(1000))
	require.Equal(t, nil, Error)
	mapper.NewSplits(mapper.Mapper, ctx).Add(mapper.NewSplit(mapper.NewSplitID(fromID, ownableID), sdkTypes.NewDec(1000)))

	t.Run("PositiveCase- Send All", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.SplitsKeeper.Transact(ctx, newMessage(defaultAddr, fromID, ownableID, sdkTypes.NewDec(1000))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	Error = keepers.SupplyKeeper.SendCoinsFromAccountToModule(ctx, defaultAddr, mapper.ModuleName, coins(1000))
	require.Equal(t, nil, Error)
	mapper.NewSplits(mapper.Mapper, ctx).Add(mapper.NewSplit(mapper.NewSplitID(fromID, ownableID), sdkTypes.NewDec(1000)))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.SplitsKeeper.Transact(ctx, newMessage(defaultAddr, fromID, ownableID, sdkTypes.NewDec(10))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Verify Identity Failure", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.SplitsKeeper.Transact(ctx, newMessage(verifyMockErrorAddress, fromID, ownableID, sdkTypes.NewDec(10))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Send Negative Balance", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.NotAuthorized)
		if got := keepers.SplitsKeeper.Transact(ctx, newMessage(defaultAddr, fromID, ownableID, sdkTypes.NewDec(-10))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Send More than own Balance", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.InsufficientBalance)
		if got := keepers.SplitsKeeper.Transact(ctx, newMessage(defaultAddr, fromID, ownableID, sdkTypes.NewDec(10000))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Split Not found", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.EntityNotFound)
		if got := keepers.SplitsKeeper.Transact(ctx, newMessage(defaultAddr, base.NewID("id"), ownableID, sdkTypes.NewDec(10))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	Error = keepers.SupplyKeeper.SendCoinsFromModuleToAccount(ctx, mapper.ModuleName, defaultAddr, coins(900))
	require.Equal(t, nil, Error)
	t.Run("NegativeCase-Module does not have enough coins", func(t *testing.T) {
		want := newTransactionResponse(errors.InsufficientBalance)
		if got := keepers.SplitsKeeper.Transact(ctx, newMessage(defaultAddr, fromID, ownableID, sdkTypes.NewDec(200))); !reflect.DeepEqual(got.IsSuccessful(), want.IsSuccessful()) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
