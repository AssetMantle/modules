// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrap

import (
	"context"
	"cosmossdk.io/math"
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/key"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/mint"
	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/parameters/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	cosmosDB "github.com/cosmos/cosmos-db"
	"github.com/cometbft/cometbft/crypto/ed25519"
	"cosmossdk.io/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"cosmossdk.io/store"
	storeMetrics "cosmossdk.io/store/metrics"
	storeTypes "cosmossdk.io/store/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	addressCodec "github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/mock"
	"math/rand"
	"reflect"
	"strconv"
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

func (mockAuxiliaryKeeper *MockAuxiliaryKeeper) Help(context context.Context, auxiliaryRequest helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	args := mockAuxiliaryKeeper.Called(context, auxiliaryRequest)
	return args.Get(0).(helpers.AuxiliaryResponse), args.Error(1)
}
func (mockAuxiliaryKeeper *MockAuxiliaryKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, i []interface{}) helpers.Keeper {
	args := mockAuxiliaryKeeper.Called(mapper, parameterManager, i)
	return args.Get(0).(helpers.Keeper)
}

const (
	TestMinterModuleName = "testMinter"
	Denom                = "stake"
	ChainID              = "testChain"
	GenesisSupply        = 1000000000000
)

var (
	moduleStoreKey = storeTypes.NewKVStoreKey(constants.ModuleName)

	authenticateAuxiliaryKeeper         = new(MockAuxiliaryKeeper)
	authenticateAuxiliaryFailureAddress = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	_                                   = authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(&Message{From: authenticateAuxiliaryFailureAddress.String(), FromID: baseIDs.PrototypeIdentityID().(*baseIDs.IdentityID)})).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	_                                   = authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)

	mintAuxiliaryKeeper       = new(MockAuxiliaryKeeper)
	mintAuxiliaryFailureDenom = "mint"
	_                         = mintAuxiliaryKeeper.On("Help", mock.Anything, mint.NewAuxiliaryRequest(baseIDs.PrototypeIdentityID(), baseDocuments.NewCoinAsset(mintAuxiliaryFailureDenom).GetCoinAssetID(), math.OneInt())).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	_                         = mintAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)

	authenticateAuxiliary = new(MockAuxiliary)
	_                     = authenticateAuxiliary.On("GetKeeper").Return(authenticateAuxiliaryKeeper)

	mintAuxiliary = new(MockAuxiliary)
	_             = mintAuxiliary.On("GetKeeper").Return(mintAuxiliaryKeeper)

	paramsStoreKey           = storeTypes.NewKVStoreKey(paramsTypes.StoreKey)
	paramsTransientStoreKeys = storeTypes.NewTransientStoreKey(paramsTypes.TStoreKey)

	codec                    = baseHelpers.TestCodec()
	authStoreKey             = storeTypes.NewKVStoreKey(authTypes.StoreKey)
	moduleAccountPermissions = map[string][]string{TestMinterModuleName: {authTypes.Minter}, constants.ModuleName: nil}
	AuthKeeper               = authKeeper.NewAccountKeeper(codec, runtime.NewKVStoreService(authStoreKey), authTypes.ProtoBaseAccount, moduleAccountPermissions, addressCodec.NewBech32Codec(sdkTypes.GetConfig().GetBech32AccountAddrPrefix()), sdkTypes.GetConfig().GetBech32AccountAddrPrefix(), authTypes.NewModuleAddress(govTypes.ModuleName).String())

	bankStoreKey         = storeTypes.NewKVStoreKey(bankTypes.StoreKey)
	blacklistedAddresses = map[string]bool{authTypes.NewModuleAddress(TestMinterModuleName).String(): false, authTypes.NewModuleAddress(constants.ModuleName).String(): false}
	BankKeeper           = bankKeeper.NewBaseKeeper(codec, runtime.NewKVStoreService(bankStoreKey), AuthKeeper, blacklistedAddresses, authTypes.NewModuleAddress(govTypes.ModuleName).String(), log.NewNopLogger())

	Context = setContext()

	coinSupply = sdkTypes.NewCoins(sdkTypes.NewCoin(Denom, math.NewInt(GenesisSupply)))
	_          = BankKeeper.MintCoins(Context, TestMinterModuleName, coinSupply)

	genesisAddress = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	_              = BankKeeper.SendCoinsFromModuleToAccount(Context, TestMinterModuleName, genesisAddress, coinSupply)

	parameterManager, _ = parameters.Prototype().Initialize(moduleStoreKey).
				Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(Denom))))).
				Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
				Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
				Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
				Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(Denom))))).
				Update(Context)

	TransactionKeeper = transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey), parameterManager, BankKeeper, authenticateAuxiliary, mintAuxiliary}
)

func setContext() sdkTypes.Context {
	memDB := cosmosDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB, log.NewNopLogger(), storeMetrics.NewNoOpMetrics())
	commitMultiStore.MountStoreWithDB(moduleStoreKey, storeTypes.StoreTypeIAVL, nil)
	commitMultiStore.MountStoreWithDB(authStoreKey, storeTypes.StoreTypeIAVL, nil)
	commitMultiStore.MountStoreWithDB(bankStoreKey, storeTypes.StoreTypeIAVL, nil)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, storeTypes.StoreTypeIAVL, nil)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, storeTypes.StoreTypeTransient, memDB)
	_ = commitMultiStore.LoadLatestVersion()
	return sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{ChainID: ChainID}, false, log.NewNopLogger())
}

func TestTransactionKeeperTransact(t *testing.T) {
	type args struct {
		from   sdkTypes.AccAddress
		denom  string
		amount int
	}
	tests := []struct {
		name    string
		args    args
		setup   func()
		want    *TransactionResponse
		wantErr helpers.Error
	}{
		{"wrapOne",
			args{genesisAddress, Denom, 1},
			func() {},
			newTransactionResponse(),
			nil,
		},
		{"wrapRandom",
			args{genesisAddress, Denom, rand.Intn(GenesisSupply)},
			func() {},
			newTransactionResponse(),
			nil,
		},
		{"wrapOneMoreThanSupply",
			args{genesisAddress, Denom, GenesisSupply + 1},
			func() {},
			nil,
			sdkErrors.ErrInsufficientFunds,
		},
		{
			"wrapNegative",
			args{genesisAddress, Denom, -1},
			func() {},
			nil,
			errorConstants.InvalidParameter,
		},
		{
			"wrapInvalidDenom",
			args{genesisAddress, random.GenerateUniqueIdentifier(), 1},
			func() {},
			nil,
			errorConstants.InvalidParameter,
		},
		{
			"identityAuthenticationFailure",
			args{authenticateAuxiliaryFailureAddress, Denom, 1},
			func() {},
			nil,
			errorConstants.MockError,
		},
		{
			"wrapZero",
			args{genesisAddress, Denom, 0},
			func() {},
			newTransactionResponse(),
			nil,
		},
		{
			"wrapCoinNotPresent",
			args{genesisAddress, "coinNotPresent", 1},
			func() {},
			nil,
			errorConstants.NotAuthorized,
		},
		{
			"wrapCoinNotAuthorized",
			args{genesisAddress, "unauthorizedCoin", 1},
			func() {
				coinSupply = sdkTypes.NewCoins(sdkTypes.NewCoin("unauthorizedCoin", math.NewInt(GenesisSupply)))
				_ = BankKeeper.MintCoins(Context, TestMinterModuleName, coinSupply)
				_ = BankKeeper.SendCoinsFromModuleToAccount(Context, TestMinterModuleName, genesisAddress, coinSupply)
			},
			nil,
			errorConstants.NotAuthorized,
		},
		{
			"mintAuxiliaryFailure",
			args{genesisAddress, mintAuxiliaryFailureDenom, 1},
			func() {
				parameterManager.Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(mintAuxiliaryFailureDenom), baseData.NewStringData(Denom))))).Update(sdkTypes.WrapSDKContext(Context))
			},
			nil,
			errorConstants.MockError,
		},
		{
			"wrapInMultiCoinScenario",
			args{genesisAddress, Denom, 1},
			func() {
				for i := 0; i < 1000; i++ {
					coinSupply = sdkTypes.NewCoins(sdkTypes.NewCoin(Denom+strconv.Itoa(i), math.NewInt(GenesisSupply)))
					_ = BankKeeper.MintCoins(Context, TestMinterModuleName, coinSupply)
					_ = BankKeeper.SendCoinsFromModuleToAccount(Context, TestMinterModuleName, genesisAddress, coinSupply)
				}
			},
			newTransactionResponse(),
			nil,
		},
		{
			"wrapInMultiCoinMultipleAddressScenario",
			args{genesisAddress, Denom, 1},
			func() {
				for i := 0; i < 1000; i++ {
					coinSupply = sdkTypes.NewCoins(sdkTypes.NewCoin(Denom+strconv.Itoa(i), math.NewInt(GenesisSupply)))
					_ = BankKeeper.MintCoins(Context, TestMinterModuleName, coinSupply)
					_ = BankKeeper.SendCoinsFromModuleToAccount(Context, TestMinterModuleName, sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address()), coinSupply)
				}
			},
			newTransactionResponse(),
			nil,
		},
		{
			"wrapInMultiAssetScenario",
			args{genesisAddress, Denom, 1},
			func() {
				wrapAllowedDenoms := baseData.NewListData(baseData.NewStringData(Denom))
				wrapCoins := sdkTypes.NewCoins()
				for i := 0; i < 1000; i++ {
					coinSupply = sdkTypes.NewCoins(sdkTypes.NewCoin(Denom+strconv.Itoa(i), math.NewInt(GenesisSupply)))
					_ = BankKeeper.MintCoins(Context, TestMinterModuleName, coinSupply)
					_ = BankKeeper.SendCoinsFromModuleToAccount(Context, TestMinterModuleName, genesisAddress, coinSupply)
					wrapAllowedDenoms = wrapAllowedDenoms.Add(baseData.NewStringData(Denom + strconv.Itoa(i)))
					wrapCoins = wrapCoins.Add(sdkTypes.NewCoin(Denom+strconv.Itoa(i), math.NewInt(GenesisSupply)))
				}
				parameterManager.Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), wrapAllowedDenoms))).Update(sdkTypes.WrapSDKContext(Context))
				// NOTE: This Transact may fail if parameter update validation rejects the large list
				TransactionKeeper.Transact(sdkTypes.WrapSDKContext(Context), NewMessage(genesisAddress, baseIDs.PrototypeIdentityID(), wrapCoins).(helpers.Message))
			},
			newTransactionResponse(),
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.setup()

			var initialSupply, initialAddressBalance, finalSupply, finalAddressBalance math.Int
			if sdkTypes.ValidateDenom(tt.args.denom) == nil {
				initialSupply = BankKeeper.GetSupply(Context, tt.args.denom).Amount
				initialAddressBalance = BankKeeper.GetBalance(Context, genesisAddress, tt.args.denom).Amount
			}
			got, err := TransactionKeeper.Transact(sdkTypes.WrapSDKContext(Context), NewMessage(tt.args.from, baseIDs.PrototypeIdentityID(), sdkTypes.Coins{sdkTypes.Coin{Denom: tt.args.denom, Amount: math.NewInt(int64(tt.args.amount))}}).(helpers.Message))

			if sdkTypes.ValidateDenom(tt.args.denom) == nil {
				finalSupply = BankKeeper.GetSupply(Context, tt.args.denom).Amount
				if !initialSupply.Sub(finalSupply).IsZero() {
					t.Error("supply should not change")
				}

				finalAddressBalance = BankKeeper.GetBalance(Context, genesisAddress, tt.args.denom).Amount
				if tt.wantErr == nil && !initialAddressBalance.Sub(finalAddressBalance).Equal(math.NewInt(int64(tt.args.amount))) {
					t.Error("unexpected address balance")
				}
			}

			if tt.wantErr == nil {
				if Mappable := TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Fetch(key.NewKey(baseDocuments.NewCoinAsset(tt.args.denom).GetCoinAssetID())).GetMappable(key.NewKey(baseDocuments.NewCoinAsset(tt.args.denom).GetCoinAssetID())); Mappable == nil {
					t.Error("coin asset should have been created")
				}
			}

			if tt.wantErr != nil && !initialAddressBalance.Equal(finalAddressBalance) {
				t.Error("address balance should not have changed")

			}

			if (err != nil) && !tt.wantErr.Is(err) {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Error("unexpected response")
			}
		})
	}
}
