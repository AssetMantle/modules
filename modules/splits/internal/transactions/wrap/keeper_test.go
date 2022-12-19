// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrap

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/splits/internal/key"
	"github.com/AssetMantle/modules/modules/splits/internal/mappable"
	"github.com/AssetMantle/modules/modules/splits/internal/module"
	"github.com/AssetMantle/modules/modules/splits/internal/parameters"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIds "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

var (
	authenticateAuxiliary helpers.Auxiliary
	delPk1                = ed25519.GenPrivKey().PubKey()
	delAddr1              = sdkTypes.AccAddress(delPk1.Address())

	// test addresses
	TestAddrs = []sdkTypes.AccAddress{
		delAddr1,
	}
)

type TestKeepers struct {
	WrapKeeper helpers.TransactionKeeper
}

func createTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.Parameters, bankKeeper.Keeper) {
	var legacyAmino = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	keyAcc := sdkTypes.NewKVStoreKey(authTypes.StoreKey)
	keyDistr := sdkTypes.NewKVStoreKey(types.StoreKey)
	keyStaking := sdkTypes.NewKVStoreKey(stakingTypes.StoreKey)
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := paramsKeeper.NewKeeper(
		legacyAmino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters := parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(keyDistr, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(keyStaking, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(keyAcc, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	authenticateAuxiliary = authenticate.AuxiliaryMock.Initialize(Mapper, Parameters)

	context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	feeCollectorAcc := authTypes.NewEmptyModuleAccount(auth.FeeCollectorName)
	notBondedPool := authTypes.NewEmptyModuleAccount(staking.NotBondedPoolName, supply.Burner, supply.Staking)
	bondPool := authTypes.NewEmptyModuleAccount(staking.BondedPoolName, supply.Burner, supply.Staking)
	distrAcc := authTypes.NewEmptyModuleAccount(types.ModuleName)
	splitAcc := authTypes.NewEmptyModuleAccount(module.Name)

	blacklistedAddrs := make(map[string]bool)
	blacklistedAddrs[feeCollectorAcc.GetAddress().String()] = true
	blacklistedAddrs[notBondedPool.GetAddress().String()] = true
	blacklistedAddrs[bondPool.GetAddress().String()] = true
	blacklistedAddrs[distrAcc.GetAddress().String()] = true
	blacklistedAddrs[splitAcc.GetAddress().String()] = true

	accountKeeper := keeper.NewAccountKeeper(legacyAmino, keyAcc, paramsKeeper.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bankKeeper := bankKeeper.NewBaseKeeper(accountKeeper, paramsKeeper.Subspace(bank.DefaultParamspace), blacklistedAddrs)

	maccPerms := map[string][]string{
		auth.FeeCollectorName:          nil,
		types.ModuleName:               nil,
		module.Name:                    nil,
		stakingTypes.NotBondedPoolName: {supply.Burner, supply.Staking},
		stakingTypes.BondedPoolName:    {supply.Burner, supply.Staking},
	}

	supplyKeeper := supply.NewKeeper(legacyAmino, storeKey, accountKeeper, bankKeeper, maccPerms)

	sk := staking.NewKeeper(legacyAmino, keyStaking, supplyKeeper, paramsKeeper.Subspace(staking.DefaultParamspace))
	sk.SetParams(context, staking.DefaultParams())
	intToken := sdkTypes.TokensFromConsensusPower(100000000)
	initCoins := sdkTypes.NewCoins(sdkTypes.NewCoin(sk.BondDenom(context), intToken))
	totalSupply := sdkTypes.NewCoins(sdkTypes.NewCoin(sk.BondDenom(context), intToken.MulRaw(int64(len(TestAddrs)))))
	supplyKeeper.SetSupply(context, supply.NewSupply(totalSupply))

	for _, addr := range TestAddrs {
		_, err := bankKeeper.AddCoins(context, addr, initCoins)
		require.Nil(t, err)
	}

	// set module accounts
	supplyKeeper.SetModuleAccount(context, feeCollectorAcc)
	supplyKeeper.SetModuleAccount(context, notBondedPool)
	supplyKeeper.SetModuleAccount(context, bondPool)
	supplyKeeper.SetModuleAccount(context, distrAcc)
	supplyKeeper.SetModuleAccount(context, splitAcc)

	keepers := TestKeepers{
		WrapKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.TransactionKeeper),
	}

	return context, keepers, Mapper, Parameters, supplyKeeper
}

func Test_keeperPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.TransactionKeeper
	}{
		{"+ve", transactionKeeper{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keeperPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keeperPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Initialize(t *testing.T) {
	_, _, Mapper, Parameters, _ := createTestInput(t)
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	paramsKeeper := paramsKeeper.NewKeeper.NewKeeper(
		codec.NewLegacyAmino(),
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	accountKeeper := auth.NewAccountKeeper(
		codec.NewLegacyAmino(), // amino legacyAmino
		paramsStoreKey,         // target store
		paramsKeeper.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount, // prototype
	)
	feeCollectorAcc := supply.NewEmptyModuleAccount(auth.FeeCollectorName)
	notBondedPool := supply.NewEmptyModuleAccount(staking.NotBondedPoolName, supply.Burner, supply.Staking)
	bondPool := supply.NewEmptyModuleAccount(staking.BondedPoolName, supply.Burner, supply.Staking)
	distrAcc := supply.NewEmptyModuleAccount(types.ModuleName)
	blacklistedAddrs := make(map[string]bool)
	blacklistedAddrs[feeCollectorAcc.GetAddress().String()] = true
	blacklistedAddrs[notBondedPool.GetAddress().String()] = true
	blacklistedAddrs[bondPool.GetAddress().String()] = true
	blacklistedAddrs[distrAcc.GetAddress().String()] = true
	bankKeeper := bank.NewBaseKeeper(accountKeeper, paramsKeeper.Subspace(bank.DefaultParamspace), blacklistedAddrs)
	maccPerms := map[string][]string{
		auth.FeeCollectorName:     nil,
		types.ModuleName:          nil,
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
	}
	supplyKeeper := supply.NewKeeper(codec.Cdc, sdkTypes.NewKVStoreKey(supply.StoreKey), accountKeeper, bankKeeper, maccPerms)
	type fields struct {
		mapper                helpers.Mapper
		parameters            helpers.Parameters
		supplyKeeper          supply.Keeper
		authenticateAuxiliary helpers.Auxiliary
	}
	type args struct {
		mapper      helpers.Mapper
		parameters  helpers.Parameters
		auxiliaries []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}{
		{"+ve", fields{Mapper, Parameters, supplyKeeper, authenticateAuxiliary}, args{Mapper, Parameters, []interface{}{}}, transactionKeeper{Mapper, Parameters, supplyKeeper, authenticateAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameters:            tt.fields.parameters,
				bankKeeper:            tt.fields.supplyKeeper,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.parameters, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers, Mapper, Parameters, supplyKeeper := createTestInput(t)
	testRate1 := sdkTypes.NewCoins(sdkTypes.NewInt64Coin("stake", 1))
	testOwnableID := baseIds.NewOwnableID(baseIds.NewStringID("stake"))
	split := baseTypes.NewSplit(fromID, testOwnableID, sdkTypes.NewDec(1))
	keepers.WrapKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewMappable(split))
	type fields struct {
		mapper                helpers.Mapper
		parameters            helpers.Parameters
		supplyKeeper          supply.Keeper
		authenticateAuxiliary helpers.Auxiliary
	}
	type args struct {
		context sdkTypes.Context
		msg     sdkTypes.Msg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.TransactionResponse
	}{
		{"+ve", fields{Mapper, Parameters, supplyKeeper, authenticateAuxiliary}, args{context, newMessage(delAddr1, fromID, testRate1)}, newTransactionResponse(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameters:            tt.fields.parameters,
				bankKeeper:            tt.fields.supplyKeeper,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
			}
			if got := transactionKeeper.Transact(tt.args.context, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
