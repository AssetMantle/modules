// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unwrap

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
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
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/splits/internal/key"
	"github.com/AssetMantle/modules/modules/splits/internal/mappable"
	"github.com/AssetMantle/modules/modules/splits/internal/module"
	"github.com/AssetMantle/modules/modules/splits/internal/parameters"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
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
	UnwrapKeeper helpers.TransactionKeeper
}

func createTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager, bankKeeper.Keeper) {
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
	encodingConfig := simapp.MakeTestEncodingConfig()
	appCodec := encodingConfig.Marshaler
	paramsKeeper := paramsKeeper.NewKeeper(
		appCodec,
		legacyAmino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	parameterManager := parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))

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

	authenticateAuxiliary = authenticate.Auxiliary.Initialize(Mapper, parameterManager)

	context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	feeCollectorAcc := authTypes.NewEmptyModuleAccount(authTypes.FeeCollectorName)
	notBondedPool := authTypes.NewEmptyModuleAccount(stakingTypes.NotBondedPoolName, authTypes.Burner, authTypes.Staking)
	bondPool := authTypes.NewEmptyModuleAccount(stakingTypes.BondedPoolName, authTypes.Burner, authTypes.Staking)
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
		authTypes.FeeCollectorName:     nil,
		types.ModuleName:               nil,
		module.Name:                    nil,
		stakingTypes.NotBondedPoolName: {authTypes.Burner, authTypes.Staking},
		stakingTypes.BondedPoolName:    {authTypes.Burner, authTypes.Staking},
	}
	sk := staking.NewKeeper(legacyAmino, keyStaking, supplyKeeper, paramsKeeper.Subspace(staking.DefaultParamspace))
	sk.SetParams(context, staking.DefaultParams())
	intToken := sdkTypes.TokensFromConsensusPower(100000000)
	initCoins := sdkTypes.NewCoins(sdkTypes.NewCoin(sk.BondDenom(context), intToken))
	testCoin := sdkTypes.NewCoins(sdkTypes.NewCoin("stake", intToken))
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

	_, err = bankKeeper.AddCoins(context, splitAcc.GetAddress(), testCoin)
	require.Nil(t, err)

	keepers := TestKeepers{
		UnwrapKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.TransactionKeeper),
	}

	return context, keepers, Mapper, parameterManager, supplyKeeper
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
	_, _, Mapper, Parameters, supplyKeeper := createTestInput(t)
	type fields struct {
		mapper                helpers.Mapper
		parameterManager      helpers.ParameterManager
		supplyKeeper          supply.Keeper
		authenticateAuxiliary helpers.Auxiliary
	}
	type args struct {
		mapper           helpers.Mapper
		parameterManager helpers.ParameterManager
		auxiliaries      []interface{}
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
				parameterManager:      tt.fields.parameterManager,
				bankKeeper:            tt.fields.supplyKeeper,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.parameterManager, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers, Mapper, Parameters, supplyKeeper := createTestInput(t)
	testOwnableID := baseIDs.NewCoinID(baseIDs.NewStringID("stake"))
	testRate2 := sdkTypes.NewDec(1)
	split := baseTypes.NewSplit(fromID, testOwnableID, testRate2)
	keepers.UnwrapKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewMappable(split))

	type fields struct {
		mapper                helpers.Mapper
		parameterManager      helpers.ParameterManager
		supplyKeeper          supply.Keeper
		authenticateAuxiliary helpers.Auxiliary
	}
	type args struct {
		context sdkTypes.Context
		msg     helpers.Message
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.TransactionResponse
	}{
		{"+ve", fields{Mapper, Parameters, supplyKeeper, authenticateAuxiliary}, args{context, newMessage(delAddr1, fromID, testOwnableID, testRate)}, newTransactionResponse(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameterManager:      tt.fields.parameterManager,
				bankKeeper:            tt.fields.supplyKeeper,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
			}
			if got := transactionKeeper.Transact(sdkTypes.WrapSDKContext(context), tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
