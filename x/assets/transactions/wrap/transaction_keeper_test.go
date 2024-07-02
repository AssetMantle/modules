// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrap

import (
	"context"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/AssetMantle/modules/helpers"
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
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/parameters/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	tendermintDB "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/cometbft/cometbft/libs/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
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

func (mockAuxiliaryKeeper *MockAuxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	args := mockAuxiliaryKeeper.Called(context, request)
	return args.Get(0).(helpers.AuxiliaryResponse), args.Error(1)
}
func (mockAuxiliaryKeeper *MockAuxiliaryKeeper) Initialize(m2 helpers.Mapper, manager helpers.ParameterManager, i []interface{}) helpers.Keeper {
	args := mockAuxiliaryKeeper.Called(m2, manager, i)
	return args.Get(0).(helpers.Keeper)
}

const (
	TestMinterModuleName = "testMinter"
	Denom                = "stake"
	ChainID              = "testChain"
	GenesisSupply        = 1000000000000
)

var (
	moduleStoreKey = sdkTypes.NewKVStoreKey(constants.ModuleName)

	authenticateAuxiliaryKeeper         = new(MockAuxiliaryKeeper)
	authenticateAuxiliaryFailureAddress = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	_                                   = authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(authenticateAuxiliaryFailureAddress, baseIDs.PrototypeIdentityID())).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	_                                   = authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)

	mintAuxiliaryKeeper       = new(MockAuxiliaryKeeper)
	mintAuxiliaryFailureDenom = "mint"
	_                         = mintAuxiliaryKeeper.On("Help", mock.Anything, mint.NewAuxiliaryRequest(baseIDs.PrototypeIdentityID(), baseDocuments.NewCoinAsset(mintAuxiliaryFailureDenom).GetCoinAssetID(), sdkTypes.OneInt())).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	_                         = mintAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)

	authenticateAuxiliary = new(MockAuxiliary)
	_                     = authenticateAuxiliary.On("GetKeeper").Return(authenticateAuxiliaryKeeper)

	mintAuxiliary = new(MockAuxiliary)
	_             = mintAuxiliary.On("GetKeeper").Return(mintAuxiliaryKeeper)

	encodingConfig = testutil.MakeTestEncodingConfig()

	paramsStoreKey           = sdkTypes.NewKVStoreKey(paramsTypes.StoreKey)
	paramsTransientStoreKeys = sdkTypes.NewTransientStoreKey(paramsTypes.TStoreKey)
	ParamsKeeper             = paramsKeeper.NewKeeper(encodingConfig.Codec, encodingConfig.Amino, paramsStoreKey, paramsTransientStoreKeys)

	authStoreKey             = sdkTypes.NewKVStoreKey(authTypes.StoreKey)
	moduleAccountPermissions = map[string][]string{TestMinterModuleName: {authTypes.Minter}, constants.ModuleName: nil}
	AuthKeeper               = authKeeper.NewAccountKeeper(encodingConfig.Codec, authStoreKey, authTypes.ProtoBaseAccount, moduleAccountPermissions, sdkTypes.GetConfig().GetBech32AccountAddrPrefix(), authTypes.NewModuleAddress(govTypes.ModuleName).String())

	bankStoreKey         = sdkTypes.NewKVStoreKey(bankTypes.StoreKey)
	blacklistedAddresses = map[string]bool{authTypes.NewModuleAddress(TestMinterModuleName).String(): false, authTypes.NewModuleAddress(constants.ModuleName).String(): false}
	BankKeeper           = bankKeeper.NewBaseKeeper(encodingConfig.Codec, bankStoreKey, AuthKeeper, blacklistedAddresses, authTypes.NewModuleAddress(govTypes.ModuleName).String())

	Context = setContext()

	coinSupply = sdkTypes.NewCoins(sdkTypes.NewCoin(Denom, sdkTypes.NewInt(GenesisSupply)))
	_          = BankKeeper.MintCoins(Context, TestMinterModuleName, coinSupply)

	genesisAddress = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	_              = BankKeeper.SendCoinsFromModuleToAccount(Context, TestMinterModuleName, genesisAddress, coinSupply)

	parameterManager = parameters.Prototype().Initialize(ParamsKeeper.Subspace(constants.ModuleName).WithKeyTable(parameters.Prototype().GetKeyTable())).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(Denom)))))).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true))))).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true))))).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true))))).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(Denom))))))

	TransactionKeeper = transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey), parameterManager, BankKeeper, authenticateAuxiliary, mintAuxiliary}
)

func setContext() sdkTypes.Context {
	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(moduleStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(authStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(bankStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, storeTypes.StoreTypeIAVL, memDB)
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
				coinSupply = sdkTypes.NewCoins(sdkTypes.NewCoin("unauthorizedCoin", sdkTypes.NewInt(GenesisSupply)))
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
				parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(mintAuxiliaryFailureDenom), baseData.NewStringData(Denom))))))
			},
			nil,
			errorConstants.MockError,
		},
		{
			"wrapInMultiCoinScenario",
			args{genesisAddress, Denom, 1},
			func() {
				for i := 0; i < 1000; i++ {
					coinSupply = sdkTypes.NewCoins(sdkTypes.NewCoin(Denom+strconv.Itoa(i), sdkTypes.NewInt(GenesisSupply)))
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
					coinSupply = sdkTypes.NewCoins(sdkTypes.NewCoin(Denom+strconv.Itoa(i), sdkTypes.NewInt(GenesisSupply)))
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
					coinSupply = sdkTypes.NewCoins(sdkTypes.NewCoin(Denom+strconv.Itoa(i), sdkTypes.NewInt(GenesisSupply)))
					_ = BankKeeper.MintCoins(Context, TestMinterModuleName, coinSupply)
					_ = BankKeeper.SendCoinsFromModuleToAccount(Context, TestMinterModuleName, genesisAddress, coinSupply)
					wrapAllowedDenoms = wrapAllowedDenoms.Add(baseData.NewStringData(Denom + strconv.Itoa(i)))
					wrapCoins = wrapCoins.Add(sdkTypes.NewCoin(Denom+strconv.Itoa(i), sdkTypes.NewInt(GenesisSupply)))
				}
				parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), wrapAllowedDenoms))))
				_, err := TransactionKeeper.Transact(sdkTypes.WrapSDKContext(Context), NewMessage(genesisAddress, baseIDs.PrototypeIdentityID(), wrapCoins).(helpers.Message))
				if err != nil {
					t.Error("unexpected error")
				}
			},
			newTransactionResponse(),
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.setup()

			var initialSupply, initialAddressBalance, finalSupply, finalAddressBalance sdkTypes.Int
			if sdkTypes.ValidateDenom(tt.args.denom) == nil {
				initialSupply = BankKeeper.GetSupply(Context, tt.args.denom).Amount
				initialAddressBalance = BankKeeper.GetBalance(Context, genesisAddress, tt.args.denom).Amount
			}
			got, err := TransactionKeeper.Transact(sdkTypes.WrapSDKContext(Context), NewMessage(tt.args.from, baseIDs.PrototypeIdentityID(), sdkTypes.Coins{sdkTypes.Coin{Denom: tt.args.denom, Amount: sdkTypes.NewInt(int64(tt.args.amount))}}).(helpers.Message))

			if sdkTypes.ValidateDenom(tt.args.denom) == nil {
				finalSupply = BankKeeper.GetSupply(Context, tt.args.denom).Amount
				if !initialSupply.Sub(finalSupply).IsZero() {
					t.Error("supply should not change")
				}

				finalAddressBalance = BankKeeper.GetBalance(Context, genesisAddress, tt.args.denom).Amount
				if tt.wantErr == nil && !initialAddressBalance.Sub(finalAddressBalance).Equal(sdkTypes.NewInt(int64(tt.args.amount))) {
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
