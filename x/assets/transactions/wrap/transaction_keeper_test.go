// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrap

import (
	"context"
	"fmt"
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
	"github.com/AssetMantle/modules/x/assets/record"
	"github.com/AssetMantle/modules/x/classifications"
	"github.com/AssetMantle/modules/x/identities"
	"github.com/AssetMantle/modules/x/maintainers"
	"github.com/AssetMantle/modules/x/metas"
	"github.com/AssetMantle/modules/x/orders"
	"github.com/AssetMantle/modules/x/splits"
	"github.com/AssetMantle/modules/x/splits/constants"
	baseData "github.com/AssetMantle/schema/go/data/base"
	baseDocuments "github.com/AssetMantle/schema/go/documents/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/parameters/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authzModule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	distributionClient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distributionTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	feegrantModule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintTypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsClient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeClient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	ica "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts"
	icaTypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	"github.com/cosmos/ibc-go/v4/modules/apps/transfer"
	ibcTransferTypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v4/modules/core"
	ibcClientClient "github.com/cosmos/ibc-go/v4/modules/core/02-client/client"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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

func (m *MockAuxiliary) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (m *MockAuxiliary) GetKeeper() helpers.AuxiliaryKeeper {
	args := m.Called()
	return args.Get(0).(helpers.AuxiliaryKeeper)
}

func (m *MockAuxiliary) Initialize(_ helpers.Mapper, _ helpers.ParameterManager, _ ...interface{}) helpers.Auxiliary {
	//TODO implement me
	panic("implement me")
}

type MockAuxiliaryKeeper struct {
	mock.Mock
}

var _ helpers.AuxiliaryKeeper = (*MockAuxiliaryKeeper)(nil)

func (m *MockAuxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	args := m.Called(context, request)
	return args.Get(0).(helpers.AuxiliaryResponse), args.Error(1)
}

func (m *MockAuxiliaryKeeper) Initialize(m2 helpers.Mapper, manager helpers.ParameterManager, i []interface{}) helpers.Keeper {
	args := m.Called(m2, manager, i)
	return args.Get(0).(helpers.Keeper)
}

var (
	authenticateAuxiliary helpers.Auxiliary
	mintAuxiliary         helpers.Auxiliary
	delPk1                = ed25519.GenPrivKey().PubKey()
	delAddr1              = sdkTypes.AccAddress(delPk1.Address())
)

func createTestInput(t *testing.T) (sdkTypes.Context, helpers.TransactionKeeper, helpers.Mapper, helpers.ParameterManager, bankKeeper.Keeper) {
	var legacyAmino = baseHelpers.CodecPrototype().GetLegacyAmino()
	var ModuleBasicManager = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distribution.AppModuleBasic{},
		gov.NewAppModuleBasic(
			paramsClient.ProposalHandler,
			distributionClient.ProposalHandler,
			upgradeClient.ProposalHandler,
			upgradeClient.CancelProposalHandler,
			ibcClientClient.UpdateClientProposalHandler,
			ibcClientClient.UpgradeProposalHandler,
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantModule.AppModuleBasic{},
		authzModule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		ica.AppModuleBasic{},

		classifications.Prototype(),
		identities.Prototype(),
		maintainers.Prototype(),
		metas.Prototype(),
		orders.Prototype(),
		splits.Prototype(),
	)

	Codec := baseHelpers.CodecPrototype().Initialize(ModuleBasicManager)

	storeKey := sdkTypes.NewKVStoreKey("test")
	keyAcc := sdkTypes.NewKVStoreKey(authTypes.StoreKey)
	keyDistr := sdkTypes.NewKVStoreKey(distributionTypes.StoreKey)
	keyStaking := sdkTypes.NewKVStoreKey(stakingTypes.StoreKey)
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := mapper.Prototype().Initialize(storeKey)
	encodingConfig := simapp.MakeTestEncodingConfig()
	ParamsKeeper := paramsKeeper.NewKeeper(
		encodingConfig.Marshaler,
		legacyAmino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)

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

	authenticateAuxiliaryKeeper := new(MockAuxiliaryKeeper)
	authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	authenticateAuxiliary = new(MockAuxiliary)
	authenticateAuxiliary.(*MockAuxiliary).On("GetKeeper").Return(authenticateAuxiliaryKeeper)

	mintAuxiliaryKeeper := new(MockAuxiliaryKeeper)
	mintAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	mintAuxiliary = new(MockAuxiliary)
	mintAuxiliary.(*MockAuxiliary).On("GetKeeper").Return(mintAuxiliaryKeeper)

	Context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	parameterManager := parameters.Prototype().Initialize(ParamsKeeper.Subspace("test").WithKeyTable(parameters.Prototype().GetKeyTable()))
	parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData("stake"))))))
	parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true)))))
	parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true)))))
	parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true)))))
	parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData("stake"))))))

	feeCollectorAcc := authTypes.NewEmptyModuleAccount(authTypes.FeeCollectorName)
	notBondedPool := authTypes.NewEmptyModuleAccount(stakingTypes.NotBondedPoolName, authTypes.Burner, authTypes.Staking)
	bondPool := authTypes.NewEmptyModuleAccount(stakingTypes.BondedPoolName, authTypes.Burner, authTypes.Staking)
	distrAcc := authTypes.NewEmptyModuleAccount(distributionTypes.ModuleName)
	splitAcc := authTypes.NewEmptyModuleAccount(constants.ModuleName)

	var ModuleAccountPermissions = map[string][]string{
		authTypes.FeeCollectorName:         nil,
		distributionTypes.ModuleName:       nil,
		icaTypes.ModuleName:                nil,
		mintTypes.ModuleName:               {authTypes.Minter},
		stakingTypes.BondedPoolName:        {authTypes.Burner, authTypes.Staking},
		stakingTypes.NotBondedPoolName:     {authTypes.Burner, authTypes.Staking},
		govTypes.ModuleName:                {authTypes.Burner},
		ibcTransferTypes.ModuleName:        {authTypes.Minter, authTypes.Burner},
		"test":                             {authTypes.Minter},
		"assets":                           nil,
		classifications.Prototype().Name(): {authTypes.Burner},
	}

	var TokenReceiveAllowedModules = map[string]bool{
		distributionTypes.ModuleName: true,
	}

	blacklistedAddresses := make(map[string]bool)
	for account := range ModuleAccountPermissions {
		blacklistedAddresses[authTypes.NewModuleAddress(account).String()] = !TokenReceiveAllowedModules[account]
	}

	accountKeeper := keeper.NewAccountKeeper(Codec.GetProtoCodec(), keyAcc, ParamsKeeper.Subspace(authTypes.ModuleName), authTypes.ProtoBaseAccount, ModuleAccountPermissions)
	BankKeeper := bankKeeper.NewBaseKeeper(Codec.GetProtoCodec(), keyAcc, accountKeeper, ParamsKeeper.Subspace(bankTypes.ModuleName), blacklistedAddresses)

	testCoins := sdkTypes.NewCoins(sdkTypes.NewCoin("stake", sdkTypes.NewInt(1000000000000)))
	if err := BankKeeper.MintCoins(Context, "test", testCoins); err != nil {
		panic(err)
	}

	if err := BankKeeper.SendCoinsFromModuleToAccount(Context, "test", delAddr1, testCoins); err != nil {
		panic(err)
	}

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

	wrapKeeper := keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{}).(helpers.TransactionKeeper)

	return Context, wrapKeeper, Mapper, parameterManager, BankKeeper
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
	Context, wrapKeepers, Mapper, parameterManager, BankKeeper := createTestInput(t)
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testAsset := baseDocuments.NewAsset(classificationID, immutables, mutables)
	testCoins := sdkTypes.NewCoins(sdkTypes.NewCoin("stake", sdkTypes.OneInt()))

	wrapKeepers.(transactionKeeper).mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(record.NewRecord(testAsset))

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
		{"+ve", fields{Mapper, parameterManager, BankKeeper, authenticateAuxiliary, mintAuxiliary}, args{sdkTypes.WrapSDKContext(Context), NewMessage(delAddr1, fromID, testCoins)}, newTransactionResponse(), false},
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
