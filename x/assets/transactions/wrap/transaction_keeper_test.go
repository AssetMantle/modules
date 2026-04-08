// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrap

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
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/mint"
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
	"github.com/stretchr/testify/mock"
	"math/rand"
	"strconv"
	"testing"
	"github.com/stretchr/testify/assert"
)

type testSetup struct {
	Context                            sdkTypes.Context
	genesisAddress                     sdkTypes.AccAddress
	TransactionKeeper                  transactionKeeper
	BankKeeper                         bankKeeper.BaseKeeper
	parameterManager                   helpers.ParameterManager
	authenticateAuxiliaryFailureAddress sdkTypes.AccAddress
	mintAuxiliaryFailureDenom          string
}

func setupTest(t *testing.T) *testSetup {
	t.Helper()

	moduleStoreKey := storeTypes.NewKVStoreKey(constants.ModuleName)
	ctx, keepers, genesisAddress := testutil.NewTestContextWithBankAuth(t, constants.ModuleName, moduleStoreKey)

	authenticateAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	authenticateAuxiliaryFailureAddress := testutil.TestAddress()
	authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(&Message{From: authenticateAuxiliaryFailureAddress.String(), FromID: baseIDs.PrototypeIdentityID().(*baseIDs.IdentityID)})).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)

	mintAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	mintAuxiliaryFailureDenom := "mint"
	mintAuxiliaryKeeper.On("Help", mock.Anything, mint.NewAuxiliaryRequest(baseIDs.PrototypeIdentityID(), baseDocuments.NewCoinAsset(mintAuxiliaryFailureDenom).GetCoinAssetID(), math.OneInt())).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	mintAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)

	authenticateAuxiliary := new(testutil.MockAuxiliary)
	authenticateAuxiliary.On("GetKeeper").Return(authenticateAuxiliaryKeeper)

	mintAuxiliary := new(testutil.MockAuxiliary)
	mintAuxiliary.On("GetKeeper").Return(mintAuxiliaryKeeper)

	parameterManager, _ := parameters.Prototype().Initialize(moduleStoreKey).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(testutil.Denom))))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(testutil.Denom))))).
		Update(ctx)

	TransactionKeeper := transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey), parameterManager, keepers.BankKeeper, authenticateAuxiliary, mintAuxiliary}

	return &testSetup{
		Context:                            ctx,
		genesisAddress:                     genesisAddress,
		TransactionKeeper:                  TransactionKeeper,
		BankKeeper:                         keepers.BankKeeper,
		parameterManager:                   parameterManager,
		authenticateAuxiliaryFailureAddress: authenticateAuxiliaryFailureAddress,
		mintAuxiliaryFailureDenom:          mintAuxiliaryFailureDenom,
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
		{"wrapOne",
			args{s.genesisAddress, testutil.Denom, 1},
			func() {},
			newTransactionResponse(),
			nil,
		},
		{"wrapRandom",
			args{s.genesisAddress, testutil.Denom, rand.Intn(testutil.GenesisSupply)},
			func() {},
			newTransactionResponse(),
			nil,
		},
		{"wrapOneMoreThanSupply",
			args{s.genesisAddress, testutil.Denom, testutil.GenesisSupply + 1},
			func() {},
			nil,
			sdkErrors.ErrInsufficientFunds,
		},
		{
			"wrapNegative",
			args{s.genesisAddress, testutil.Denom, -1},
			func() {},
			nil,
			errorConstants.InvalidParameter,
		},
		{
			"wrapInvalidDenom",
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
			"wrapZero",
			args{s.genesisAddress, testutil.Denom, 0},
			func() {},
			newTransactionResponse(),
			nil,
		},
		{
			"wrapCoinNotPresent",
			args{s.genesisAddress, "coinNotPresent", 1},
			func() {},
			nil,
			errorConstants.NotAuthorized,
		},
		{
			"wrapCoinNotAuthorized",
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
			"mintAuxiliaryFailure",
			args{s.genesisAddress, s.mintAuxiliaryFailureDenom, 1},
			func() {
				s.parameterManager.Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(s.mintAuxiliaryFailureDenom), baseData.NewStringData(testutil.Denom))))).Update(sdkTypes.WrapSDKContext(s.Context))
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
				wrapAllowedDenoms := baseData.NewListData(baseData.NewStringData(testutil.Denom))
				wrapCoins := sdkTypes.NewCoins()
				for i := 0; i < 1000; i++ {
					coinSupply := sdkTypes.NewCoins(sdkTypes.NewCoin(testutil.Denom+strconv.Itoa(i), math.NewInt(testutil.GenesisSupply)))
					s.BankKeeper.MintCoins(s.Context, testutil.TestMinterModuleName, coinSupply)
					s.BankKeeper.SendCoinsFromModuleToAccount(s.Context, testutil.TestMinterModuleName, s.genesisAddress, coinSupply)
					wrapAllowedDenoms = wrapAllowedDenoms.Add(baseData.NewStringData(testutil.Denom + strconv.Itoa(i)))
					wrapCoins = wrapCoins.Add(sdkTypes.NewCoin(testutil.Denom+strconv.Itoa(i), math.NewInt(testutil.GenesisSupply)))
				}
				s.parameterManager.Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), wrapAllowedDenoms))).Update(sdkTypes.WrapSDKContext(s.Context))
				// NOTE: This Transact may fail if parameter update validation rejects the large list
				s.TransactionKeeper.Transact(sdkTypes.WrapSDKContext(s.Context), NewMessage(s.genesisAddress, baseIDs.PrototypeIdentityID(), wrapCoins).(helpers.Message))
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
				if tt.wantErr == nil && !initialAddressBalance.Sub(finalAddressBalance).Equal(math.NewInt(int64(tt.args.amount))) {
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
