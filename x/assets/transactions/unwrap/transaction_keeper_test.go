// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unwrap

import (
	"cosmossdk.io/math"
	storeTypes "cosmossdk.io/store/types"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/key"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
	"github.com/AssetMantle/modules/x/assets/record"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/burn"
	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/parameters/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	"github.com/cometbft/cometbft/crypto/ed25519"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strconv"
	"testing"
)

type testSetup struct {
	Context                            sdkTypes.Context
	genesisAddress                     sdkTypes.AccAddress
	TransactionKeeper                  transactionKeeper
	BankKeeper                         bankKeeper.BaseKeeper
	parameterManager                   helpers.ParameterManager
	authenticateAuxiliaryFailureAddress sdkTypes.AccAddress
	burnAuxiliaryFailureDenom          string
	newCollectionFaliure               string
}

func setupTest(t *testing.T) *testSetup {
	t.Helper()

	moduleStoreKey := storeTypes.NewKVStoreKey(constants.ModuleName)
	ctx, keepers, genesisAddress := testutil.NewTestContextWithBankAuth(t, constants.ModuleName, moduleStoreKey)

	// Send coins from genesis to module account for unwrap tests
	coinSupply := sdkTypes.NewCoins(sdkTypes.NewCoin(testutil.Denom, math.NewInt(testutil.GenesisSupply)))
	// genesisAddress already has the coins from NewTestContextWithBankAuth, send to module
	keepers.BankKeeper.SendCoinsFromAccountToModule(ctx, genesisAddress, constants.ModuleName, coinSupply)

	newCollectionFaliure := "notfound"

	authenticateAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	authenticateAuxiliaryFailureAddress := testutil.TestAddress()
	authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(&Message{From: authenticateAuxiliaryFailureAddress.String(), FromID: baseIDs.PrototypeIdentityID().(*baseIDs.IdentityID)})).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	authenticateAuxiliary := new(testutil.MockAuxiliary)
	authenticateAuxiliary.On("GetKeeper").Return(authenticateAuxiliaryKeeper)

	burnAuxiliaryFailureDenom := "burn"
	burnAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	burnAuxiliaryKeeper.On("Help", mock.Anything, burn.NewAuxiliaryRequest(baseIDs.PrototypeIdentityID(), baseDocuments.NewCoinAsset(burnAuxiliaryFailureDenom).GetCoinAssetID(), math.OneInt())).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	burnAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	burnAuxiliary := new(testutil.MockAuxiliary)
	burnAuxiliary.On("GetKeeper").Return(burnAuxiliaryKeeper)

	parameterManager, _ := parameters.Prototype().Initialize(moduleStoreKey).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(testutil.Denom))))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(testutil.Denom))))).
		Update(ctx)

	TransactionKeeper := transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey), parameterManager, keepers.BankKeeper, burnAuxiliary, authenticateAuxiliary}

	TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(ctx)).Add(record.NewRecord(baseDocuments.NewCoinAsset(testutil.Denom)))

	return &testSetup{
		Context:                            ctx,
		genesisAddress:                     genesisAddress,
		TransactionKeeper:                  TransactionKeeper,
		BankKeeper:                         keepers.BankKeeper,
		parameterManager:                   parameterManager,
		authenticateAuxiliaryFailureAddress: authenticateAuxiliaryFailureAddress,
		burnAuxiliaryFailureDenom:          burnAuxiliaryFailureDenom,
		newCollectionFaliure:               newCollectionFaliure,
	}
}

func TestTransactionKeeperTransact(t *testing.T) {
	s := setupTest(t)

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
		{"unwrapOne",
			args{s.genesisAddress, testutil.Denom, 1},
			func() {},
			newTransactionResponse(),
			nil,
		},
		{"unwrapRandom",
			args{s.genesisAddress, testutil.Denom, rand.Intn(testutil.GenesisSupply)},
			func() {},
			newTransactionResponse(),
			nil,
		},
		{"unwrapOneMoreThanSupply",
			args{s.genesisAddress, testutil.Denom, testutil.GenesisSupply + 1},
			func() {},
			nil,
			sdkErrors.ErrInsufficientFunds,
		},
		{
			"unwrapNegative",
			args{s.genesisAddress, testutil.Denom, -1},
			func() {},
			nil,
			errorConstants.InvalidParameter,
		},
		{
			"unwrapInvalidDenom",
			args{s.genesisAddress, random.GenerateUniqueIdentifier(), 1},
			func() {},
			nil,
			errorConstants.InvalidParameter,
		},
		{
			"identityAuthenticationFailure",
			args{s.authenticateAuxiliaryFailureAddress, testutil.Denom, 1},
			func() {},
			nil,
			errorConstants.MockError,
		},
		{
			"unwrapZero",
			args{s.genesisAddress, testutil.Denom, 0},
			func() {},
			newTransactionResponse(),
			nil,
		},
		{
			"unwrapCoinNotPresent",
			args{s.genesisAddress, "coinNotPresent", 1},
			func() {},
			nil,
			errorConstants.NotAuthorized,
		},
		{
			"unwrapCoinNotAuthorized",
			args{s.genesisAddress, "unauthorizedCoin", 1},
			func() {
				coinSupply := sdkTypes.NewCoins(sdkTypes.NewCoin("unauthorizedCoin", math.NewInt(testutil.GenesisSupply)))
				s.BankKeeper.MintCoins(s.Context, testutil.TestMinterModuleName, coinSupply)
				s.BankKeeper.SendCoinsFromModuleToAccount(s.Context, testutil.TestMinterModuleName, s.genesisAddress, coinSupply)
			},
			nil,
			errorConstants.NotAuthorized,
		},
		{
			"EntityNotFound",
			args{s.genesisAddress, s.burnAuxiliaryFailureDenom, 1},
			func() {
				s.parameterManager.Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(s.newCollectionFaliure), baseData.NewStringData(s.burnAuxiliaryFailureDenom))))).Update(sdkTypes.WrapSDKContext(s.Context))
			},
			nil,
			errorConstants.EntityNotFound,
		},
		{
			"burnAuxiliaryFailure",
			args{s.genesisAddress, s.burnAuxiliaryFailureDenom, 1},
			func() {
				s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(record.NewRecord(baseDocuments.NewCoinAsset(s.burnAuxiliaryFailureDenom)))
				s.parameterManager.Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(s.burnAuxiliaryFailureDenom), baseData.NewStringData(testutil.Denom))))).Update(sdkTypes.WrapSDKContext(s.Context))
			},
			nil,
			errorConstants.MockError,
		},
		{
			"wrapInMultiCoinScenario",
			args{s.genesisAddress, testutil.Denom, 1},
			func() {
				for i := 0; i < 1000; i++ {
					coinSupply := sdkTypes.NewCoins(sdkTypes.NewCoin(testutil.Denom+strconv.Itoa(i), math.NewInt(testutil.GenesisSupply)))
					s.BankKeeper.MintCoins(s.Context, testutil.TestMinterModuleName, coinSupply)
					s.BankKeeper.SendCoinsFromModuleToAccount(s.Context, testutil.TestMinterModuleName, s.genesisAddress, coinSupply)
				}
			},
			newTransactionResponse(),
			nil,
		},
		{
			"wrapInMultiCoinMultipleAddressScenario",
			args{s.genesisAddress, testutil.Denom, 1},
			func() {
				for i := 0; i < 1000; i++ {
					coinSupply := sdkTypes.NewCoins(sdkTypes.NewCoin(testutil.Denom+strconv.Itoa(i), math.NewInt(testutil.GenesisSupply)))
					s.BankKeeper.MintCoins(s.Context, testutil.TestMinterModuleName, coinSupply)
					s.BankKeeper.SendCoinsFromModuleToAccount(s.Context, testutil.TestMinterModuleName, sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address()), coinSupply)
				}
			},
			newTransactionResponse(),
			nil,
		},
		{
			"wrapInMultiAssetScenario",
			args{s.genesisAddress, testutil.Denom, 1},
			func() {
				unwrapAllowedDenoms := baseData.NewListData(baseData.NewStringData(testutil.Denom))
				unwrapCoins := sdkTypes.NewCoins()

				for i := 0; i < 1000; i++ {
					coinSupply := sdkTypes.NewCoins(sdkTypes.NewCoin(testutil.Denom+strconv.Itoa(i), math.NewInt(testutil.GenesisSupply)))
					s.BankKeeper.MintCoins(s.Context, testutil.TestMinterModuleName, coinSupply)
					s.BankKeeper.SendCoinsFromModuleToModule(s.Context, testutil.TestMinterModuleName, constants.ModuleName, coinSupply)

					s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(record.NewRecord(baseDocuments.NewCoinAsset(testutil.Denom + strconv.Itoa(i))))

					unwrapAllowedDenoms = unwrapAllowedDenoms.Add(baseData.NewStringData(testutil.Denom + strconv.Itoa(i)))
					unwrapCoins = unwrapCoins.Add(sdkTypes.NewCoin(testutil.Denom+strconv.Itoa(i), math.NewInt(testutil.GenesisSupply)))
				}
				s.parameterManager.Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), unwrapAllowedDenoms))).Update(sdkTypes.WrapSDKContext(s.Context))
				// NOTE: This Transact may fail if parameter update validation rejects the large list; that's acceptable for the test scenario
				s.TransactionKeeper.Transact(sdkTypes.WrapSDKContext(s.Context), NewMessage(s.genesisAddress, baseIDs.PrototypeIdentityID(), unwrapCoins).(helpers.Message))
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
				initialSupply = s.BankKeeper.GetSupply(s.Context, tt.args.denom).Amount
				initialAddressBalance = s.BankKeeper.GetBalance(s.Context, s.genesisAddress, tt.args.denom).Amount
			}
			got, err := s.TransactionKeeper.Transact(sdkTypes.WrapSDKContext(s.Context), NewMessage(tt.args.from, baseIDs.PrototypeIdentityID(), sdkTypes.Coins{sdkTypes.Coin{Denom: tt.args.denom, Amount: math.NewInt(int64(tt.args.amount))}}).(helpers.Message))

			if sdkTypes.ValidateDenom(tt.args.denom) == nil {
				finalSupply = s.BankKeeper.GetSupply(s.Context, tt.args.denom).Amount
				if !initialSupply.Sub(finalSupply).IsZero() {
					t.Error("supply should not change")
				}

				finalAddressBalance = s.BankKeeper.GetBalance(s.Context, s.genesisAddress, tt.args.denom).Amount
				if tt.wantErr == nil && !finalAddressBalance.Sub(initialAddressBalance).Equal(math.NewInt(int64(tt.args.amount))) {
					t.Error("unexpected address balance")
				}
			}

			if tt.wantErr == nil {
				if Mappable := s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Fetch(key.NewKey(baseDocuments.NewCoinAsset(tt.args.denom).GetCoinAssetID())).GetMappable(key.NewKey(baseDocuments.NewCoinAsset(tt.args.denom).GetCoinAssetID())); Mappable == nil {
					t.Error("coin asset should have been created")
				}
			}

			if tt.wantErr != nil && !initialAddressBalance.Equal(finalAddressBalance) {
				t.Error("address balance should not have changed")

			}

			if (err != nil) && !tt.wantErr.Is(err) {
				t.Errorf("unexpected error: %v", err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_keeperPrototype(t *testing.T) {
	got := keeperPrototype()
	assert.Equal(t, transactionKeeper{}, got)
}

func Test_transactionKeeper_Initialize(t *testing.T) {
	moduleStoreKey := storeTypes.NewKVStoreKey(constants.ModuleName)
	ctx, keepers, _ := testutil.NewTestContextWithBankAuth(t, constants.ModuleName, moduleStoreKey)
	_ = ctx

	m := mapper.Prototype().Initialize(moduleStoreKey)
	pm := parameters.Prototype().Initialize(moduleStoreKey)

	authenticateAux, _ := testutil.NewNamedMockAuxiliaryPair(authenticate.Auxiliary.GetName())
	burnAux, _ := testutil.NewNamedMockAuxiliaryPair(burn.Auxiliary.GetName())

	keeper := keeperPrototype().Initialize(m, pm, []interface{}{keepers.BankKeeper, authenticateAux, burnAux})
	require.NotNil(t, keeper)
}
