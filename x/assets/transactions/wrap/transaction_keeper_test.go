// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrap

import (
	"context"
	"fmt"
	"github.com/AssetMantle/modules/x/splits/record"
	baseData "github.com/AssetMantle/schema/go/data/base"
	baseDocuments "github.com/AssetMantle/schema/go/documents/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"reflect"
	"testing"

	baseTypes "github.com/AssetMantle/schema/go/types/base"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/splits/constants"
	"github.com/AssetMantle/modules/x/splits/parameters"
)

var (
	authenticateAuxiliary helpers.Auxiliary
	mintAuxiliary         helpers.Auxiliary
	delPk1                = ed25519.GenPrivKey().PubKey()
	delAddr1              = sdkTypes.AccAddress(delPk1.Address())
)

type TestKeepers struct {
	WrapKeeper helpers.TransactionKeeper
}

func createTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager, bankKeeper.Keeper) {
	var legacyAmino = baseHelpers.CodecPrototype().GetLegacyAmino()

	Codec := baseHelpers.CodecPrototype()

	storeKey := sdkTypes.NewKVStoreKey("test")
	keyAcc := sdkTypes.NewKVStoreKey(authTypes.StoreKey)
	keyDistr := sdkTypes.NewKVStoreKey(types.StoreKey)
	keyStaking := sdkTypes.NewKVStoreKey(stakingTypes.StoreKey)
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := mapper.Prototype().Initialize(storeKey)
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
	commitMultiStore.MountStoreWithDB(keyDistr, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(keyStaking, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(keyAcc, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	authenticateAuxiliary = authenticate.Auxiliary.Initialize(Mapper, parameterManager)

	Context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	feeCollectorAcc := authTypes.NewEmptyModuleAccount(authTypes.FeeCollectorName)
	notBondedPool := authTypes.NewEmptyModuleAccount(stakingTypes.NotBondedPoolName, authTypes.Burner, authTypes.Staking)
	bondPool := authTypes.NewEmptyModuleAccount(stakingTypes.BondedPoolName, authTypes.Burner, authTypes.Staking)
	distrAcc := authTypes.NewEmptyModuleAccount(types.ModuleName)
	splitAcc := authTypes.NewEmptyModuleAccount(constants.ModuleName)

	blacklistedAddrs := make(map[string]bool)
	blacklistedAddrs[feeCollectorAcc.GetAddress().String()] = true
	blacklistedAddrs[notBondedPool.GetAddress().String()] = true
	blacklistedAddrs[bondPool.GetAddress().String()] = true
	blacklistedAddrs[distrAcc.GetAddress().String()] = true
	blacklistedAddrs[splitAcc.GetAddress().String()] = true

	accountKeeper := keeper.NewAccountKeeper(Codec.GetProtoCodec(), keyAcc, ParamsKeeper.Subspace(authTypes.ModuleName), authTypes.ProtoBaseAccount, nil)
	BankKeeper := bankKeeper.NewBaseKeeper(Codec.GetProtoCodec(), keyAcc, accountKeeper, ParamsKeeper.Subspace(bankTypes.ModuleName), blacklistedAddrs)

	sk := stakingKeeper.NewKeeper(Codec.GetProtoCodec(), keyStaking, accountKeeper, BankKeeper, ParamsKeeper.Subspace(stakingTypes.ModuleName))
	sk.SetParams(Context, stakingTypes.DefaultParams())
	//intToken := sdkTypes.TokensFromConsensusPower(1000000, sdkTypes.NewInt(100))
	//initCoins := sdkTypes.NewCoins(sdkTypes.NewCoin(sk.BondDenom(Context), intToken))
	//totalSupply := sdkTypes.NewCoins(sdkTypes.NewCoin(sk.BondDenom(Context), intToken.MulRaw(int64(len(TestAddrs)))))

	// set module accounts
	accountKeeper.SetModuleAccount(Context, feeCollectorAcc)
	accountKeeper.SetModuleAccount(Context, notBondedPool)
	accountKeeper.SetModuleAccount(Context, bondPool)
	accountKeeper.SetModuleAccount(Context, distrAcc)
	accountKeeper.SetModuleAccount(Context, splitAcc)

	keepers := TestKeepers{
		WrapKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.TransactionKeeper),
	}

	return Context, keepers, Mapper, parameterManager, BankKeeper
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
	_, _, Mapper, parameterManager, _ := createTestInput(t)
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	encodingConfig := simapp.MakeTestEncodingConfig()
	appCodec := encodingConfig.Marshaler
	ParamsKeeper := paramsKeeper.NewKeeper(
		appCodec,
		codec.NewLegacyAmino(),
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	accountKeeper := authKeeper.NewAccountKeeper(
		codec.NewProtoCodec(nil),
		sdkTypes.NewKVStoreKey(authTypes.StoreKey),
		ParamsKeeper.Subspace(authTypes.ModuleName),
		authTypes.ProtoBaseAccount,
		nil,
	)
	feeCollectorAcc := authTypes.NewEmptyModuleAccount(authTypes.FeeCollectorName)
	notBondedPool := authTypes.NewEmptyModuleAccount(stakingTypes.NotBondedPoolName, authTypes.Burner, authTypes.Staking)
	bondPool := authTypes.NewEmptyModuleAccount(stakingTypes.BondedPoolName, authTypes.Burner, authTypes.Staking)
	distrAcc := authTypes.NewEmptyModuleAccount(types.ModuleName)
	blacklistedAddrs := make(map[string]bool)
	blacklistedAddrs[feeCollectorAcc.GetAddress().String()] = true
	blacklistedAddrs[notBondedPool.GetAddress().String()] = true
	blacklistedAddrs[bondPool.GetAddress().String()] = true
	blacklistedAddrs[distrAcc.GetAddress().String()] = true
	BankKeeper := bankKeeper.NewBaseKeeper(appCodec, sdkTypes.NewKVStoreKey(bankTypes.StoreKey), accountKeeper, ParamsKeeper.Subspace(bankTypes.ModuleName), blacklistedAddrs)
	type fields struct {
		mapper                helpers.Mapper
		parameterManager      helpers.ParameterManager
		bankKeeper            bankKeeper.Keeper
		authenticateAuxiliary helpers.Auxiliary
		mintAuxiliary         helpers.Auxiliary
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
		{"+ve", fields{Mapper, parameterManager, BankKeeper, authenticateAuxiliary, mintAuxiliary}, args{Mapper, parameterManager, []interface{}{}}, transactionKeeper{Mapper, parameterManager, BankKeeper, authenticateAuxiliary, mintAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameterManager:      tt.fields.parameterManager,
				bankKeeper:            tt.fields.bankKeeper,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
				mintAuxiliary:         tt.fields.mintAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.parameterManager, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	Context, keepers, Mapper, parameterManager, _ := createTestInput(t)
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testOwnerIdentityID := baseIDs.NewIdentityID(classificationID, immutables)
	testRate1 := sdkTypes.NewCoins(sdkTypes.NewInt64Coin("stake", 1))
	testAssetID := baseDocuments.NewCoinAsset("stake").GetCoinAssetID()
	split := baseTypes.NewSplit(sdkTypes.NewInt(1))
	keepers.WrapKeeper.(transactionKeeper).mapper.NewCollection(Context.Context()).Add(record.NewRecord(baseIDs.NewSplitID(testAssetID, testOwnerIdentityID), split))
	feeCollectorAcc := authTypes.NewEmptyModuleAccount(authTypes.FeeCollectorName)
	notBondedPool := authTypes.NewEmptyModuleAccount(stakingTypes.NotBondedPoolName, authTypes.Burner, authTypes.Staking)
	bondPool := authTypes.NewEmptyModuleAccount(stakingTypes.BondedPoolName, authTypes.Burner, authTypes.Staking)
	distrAcc := authTypes.NewEmptyModuleAccount(types.ModuleName)
	blacklistedAddrs := make(map[string]bool)
	blacklistedAddrs[feeCollectorAcc.GetAddress().String()] = true
	blacklistedAddrs[notBondedPool.GetAddress().String()] = true
	blacklistedAddrs[bondPool.GetAddress().String()] = true
	blacklistedAddrs[distrAcc.GetAddress().String()] = true
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")

	ParamsKeeper := paramsKeeper.NewKeeper(
		baseHelpers.CodecPrototype(),
		codec.NewLegacyAmino(),
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	accountKeeper := authKeeper.NewAccountKeeper(
		codec.NewProtoCodec(nil),
		sdkTypes.NewKVStoreKey(authTypes.StoreKey),
		ParamsKeeper.Subspace(authTypes.ModuleName),
		authTypes.ProtoBaseAccount,
		nil,
	)
	BankKeeper := bankKeeper.NewBaseKeeper(baseHelpers.CodecPrototype().GetProtoCodec(), sdkTypes.NewKVStoreKey(bankTypes.StoreKey), accountKeeper, ParamsKeeper.Subspace(bankTypes.ModuleName), blacklistedAddrs)

	type fields struct {
		mapper                helpers.Mapper
		parameterManager      helpers.ParameterManager
		bankKeeper            bankKeeper.Keeper
		authenticateAuxiliary helpers.Auxiliary
		mintAuxiliary         helpers.Auxiliary
	}
	type args struct {
		context context.Context
		message sdkTypes.Msg
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.TransactionResponse
		wantErr bool
	}{
		{"+ve", fields{Mapper, parameterManager, BankKeeper, authenticateAuxiliary, mintAuxiliary}, args{Context.Context(), NewMessage(delAddr1, fromID, testRate1)}, newTransactionResponse(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				parameterManager:      tt.fields.parameterManager,
				bankKeeper:            tt.fields.bankKeeper,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
				mintAuxiliary:         tt.fields.mintAuxiliary,
			}
			got, err := transactionKeeper.Transact(tt.args.context, tt.args.message.(helpers.Message))
			if (err != nil) != tt.wantErr {
				t.Errorf("Transact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() got = %v, want %v", got, tt.want)
			}
		})
	}
}
