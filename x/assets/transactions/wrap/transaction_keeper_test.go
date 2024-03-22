// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrap

import (
	"context"
	"fmt"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
	baseData "github.com/AssetMantle/schema/go/data/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/parameters/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributionTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/mock"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
	"reflect"
	"testing"
)

type MockAuxiliary struct {
	mock.Mock
}

var _ helpers.Auxiliary = (*MockAuxiliary)(nil)

func (mockAuxiliary *MockAuxiliary) GetName() string { panic(mockAuxiliary) }
func (mockAuxiliary *MockAuxiliary) GetKeeper() helpers.AuxiliaryKeeper {
	args := mockAuxiliary.Called()
	return args.Get(0).(helpers.AuxiliaryKeeper)
}
func (mockAuxiliary *MockAuxiliary) Initialize(_ helpers.Mapper, _ helpers.ParameterManager, _ ...interface{}) helpers.Auxiliary {
	panic(mockAuxiliary)
}

type MockAuxiliaryKeeper struct {
	mock.Mock
}

var _ helpers.AuxiliaryKeeper = (*MockAuxiliaryKeeper)(nil)

func (mockAuxiliaryKeeper *MockAuxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	args := mockAuxiliaryKeeper.Called(context, request)
	return args.Get(0).(helpers.AuxiliaryResponse), args.Error(1)
}
func (mockAuxiliaryKeeper *MockAuxiliaryKeeper) Initialize(m2 helpers.Mapper, manager helpers.ParameterManager, i []interface{}) helpers.Keeper {
	args := mockAuxiliaryKeeper.Called(m2, manager, i)
	return args.Get(0).(helpers.Keeper)
}

var (
	encodingConfig = simapp.MakeTestEncodingConfig()

	authenticateAuxiliaryKeeper = new(MockAuxiliaryKeeper)
	mintAuxiliaryKeeper         = new(MockAuxiliaryKeeper)
	authenticateAuxiliary       = new(MockAuxiliary)
	mintAuxiliary               = new(MockAuxiliary)
	moduleStoreKey              = sdkTypes.NewKVStoreKey(moduleName)

	authStoreKey             = sdkTypes.NewKVStoreKey(authTypes.StoreKey)
	bankStoreKey             = sdkTypes.NewKVStoreKey(bankTypes.StoreKey)
	paramsStoreKey           = sdkTypes.NewKVStoreKey(paramsTypes.StoreKey)
	paramsTransientStoreKeys = sdkTypes.NewTransientStoreKey(paramsTypes.TStoreKey)

	genesisCoinSupply = sdkTypes.NewCoins(sdkTypes.NewCoin(Denom, sdkTypes.NewInt(1000000000000)))
	genesisAddress    = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	moduleName        = constants.ModuleName

	ParamsKeeper = paramsKeeper.NewKeeper(encodingConfig.Marshaler, encodingConfig.Amino, paramsStoreKey, paramsTransientStoreKeys)

	moduleAccountPermissions = map[string][]string{TestMinterModuleName: {authTypes.Minter}, moduleName: nil}
	AuthKeeper               = authKeeper.NewAccountKeeper(encodingConfig.Marshaler, authStoreKey, ParamsKeeper.Subspace(authTypes.ModuleName), authTypes.ProtoBaseAccount, moduleAccountPermissions)

	blacklistedAddresses = map[string]bool{authTypes.NewModuleAddress(TestMinterModuleName).String(): false, authTypes.NewModuleAddress(moduleName).String(): false}
	BankKeeper           = bankKeeper.NewBaseKeeper(encodingConfig.Marshaler, bankStoreKey, AuthKeeper, ParamsKeeper.Subspace(bankTypes.ModuleName), blacklistedAddresses)

	Context = setContext()

	_ = BankKeeper.MintCoins(Context, TestMinterModuleName, genesisCoinSupply)
	_ = BankKeeper.SendCoinsFromModuleToAccount(Context, TestMinterModuleName, genesisAddress, genesisCoinSupply)

	parameterManager = setParameterManager(Context, ParamsKeeper)

	moduleMapper = mapper.Prototype().Initialize(moduleStoreKey)
)

const (
	TestMinterModuleName = "testMinter"
	Denom                = "stake"
	ChainID              = "testChain"
)

func setMockBehaviour() {
	authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	authenticateAuxiliary.On("GetKeeper").Return(authenticateAuxiliaryKeeper)
	mintAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	mintAuxiliary.On("GetKeeper").Return(mintAuxiliaryKeeper)
}
func setParameterManager(Context sdkTypes.Context, ParamsKeeper paramsKeeper.Keeper) helpers.ParameterManager {
	parameterManager := parameters.Prototype().Initialize(ParamsKeeper.Subspace(moduleName).WithKeyTable(parameters.Prototype().GetKeyTable()))
	parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData("stake"))))))
	parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true)))))
	parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true)))))
	parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true)))))
	parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData("stake"))))))
	return parameterManager
}
func setContext() sdkTypes.Context {
	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(moduleStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(authStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(bankStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	_ = commitMultiStore.LoadLatestVersion()
	return sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{ChainID: ChainID}, false, log.NewNopLogger())
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
	setMockBehaviour()
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
	distrAcc := authTypes.NewEmptyModuleAccount(distributionTypes.ModuleName)
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
		{"+ve", fields{moduleMapper, parameterManager, BankKeeper, authenticateAuxiliary, mintAuxiliary}, args{moduleMapper, parameterManager, []interface{}{}}, transactionKeeper{moduleMapper, parameterManager, BankKeeper, authenticateAuxiliary, mintAuxiliary}},
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
	setMockBehaviour()
	testCoins := sdkTypes.NewCoins(sdkTypes.NewCoin(Denom, sdkTypes.OneInt()))

	type args struct {
		message sdkTypes.Msg
	}
	tests := []struct {
		name    string
		args    args
		want    helpers.TransactionResponse
		wantErr bool
	}{
		{"+ve", args{NewMessage(genesisAddress, fromID, testCoins)}, newTransactionResponse(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				moduleMapper,
				parameterManager,
				BankKeeper,
				authenticateAuxiliary,
				mintAuxiliary,
			}
			got, err := transactionKeeper.Transact(sdkTypes.WrapSDKContext(Context), tt.args.message.(helpers.Message))
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
