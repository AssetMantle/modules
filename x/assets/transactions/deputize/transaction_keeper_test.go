// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	storeTypes "cosmossdk.io/store/types"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
	permissionHelper "github.com/AssetMantle/modules/x/assets/utilities"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/deputize"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/parameters/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type testSetup struct {
	Context                            sdkTypes.Context
	genesisAddress                     sdkTypes.AccAddress
	TransactionKeeper                  transactionKeeper
	authenticateAuxiliaryFailureAddress sdkTypes.AccAddress
}

func setupTest(t *testing.T) *testSetup {
	t.Helper()

	moduleStoreKey := storeTypes.NewKVStoreKey(constants.ModuleName)
	ctx, _, genesisAddress := testutil.NewTestContextWithBankAuth(t, constants.ModuleName, moduleStoreKey)

	authenticateAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	authenticateAuxiliaryFailureAddress := testutil.TestAddress()
	authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(&Message{From: authenticateAuxiliaryFailureAddress.String(), FromID: baseIDs.PrototypeIdentityID().(*baseIDs.IdentityID)})).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	authenticateAuxiliary := new(testutil.MockAuxiliary)
	authenticateAuxiliary.On("GetKeeper").Return(authenticateAuxiliaryKeeper)

	deputizeAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	deputizeAuxiliaryKeeper.On("Help", mock.Anything, deputize.NewAuxiliaryRequest(baseIDs.PrototypeIdentityID(), baseIDs.PrototypeIdentityID(), baseIDs.PrototypeClassificationID(), baseLists.NewPropertyList(), true, true, true, permissionHelper.SetModulePermissions(false, true, true)...)).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	deputizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	deputizeAuxiliary := new(testutil.MockAuxiliary)
	deputizeAuxiliary.On("GetKeeper").Return(deputizeAuxiliaryKeeper)

	parameterManager := parameters.Prototype().Initialize(moduleStoreKey).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(testutil.Denom))))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(testutil.Denom)))))

	TransactionKeeper := transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey), parameterManager, authenticateAuxiliary, deputizeAuxiliary}

	return &testSetup{
		Context:                            ctx,
		genesisAddress:                     genesisAddress,
		TransactionKeeper:                  TransactionKeeper,
		authenticateAuxiliaryFailureAddress: authenticateAuxiliaryFailureAddress,
	}
}

func TestTransactionKeeperTransact(t *testing.T) {
	s := setupTest(t)

	type args struct {
		from                sdkTypes.AccAddress
		fromID              ids.IdentityID
		toID                ids.IdentityID
		canMintAsset        bool
		canRenumerateAsset  bool
		canBurnAsset        bool
		canAddMaintainer    bool
		CanRemoveMaintainer bool
		CanMutateMaintainer bool
	}

	tests := []struct {
		name    string
		args    args
		setup   func()
		want    *TransactionResponse
		wantErr helpers.Error
	}{
		{
			name: "DeputizeTransactionKeeperSuccess",
			args: args{
				from:                s.genesisAddress,
				fromID:              baseIDs.PrototypeIdentityID(),
				toID:                baseIDs.PrototypeIdentityID(),
				canMintAsset:        true,
				canRenumerateAsset:  true,
				canBurnAsset:        true,
				canAddMaintainer:    true,
				CanRemoveMaintainer: true,
				CanMutateMaintainer: true,
			},
			setup: func() {
			},
			want:    newTransactionResponse(),
			wantErr: nil,
		},
		{
			name: "AuthenticationFailure",
			args: args{
				from:                s.authenticateAuxiliaryFailureAddress,
				fromID:              baseIDs.PrototypeIdentityID(),
				toID:                baseIDs.PrototypeIdentityID(),
				canMintAsset:        true,
				canRenumerateAsset:  true,
				canBurnAsset:        true,
				canAddMaintainer:    true,
				CanRemoveMaintainer: true,
				CanMutateMaintainer: true,
			},
			setup: func() {
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "DeputizeAuxiliaryFailure",
			args: args{
				from:                s.genesisAddress,
				fromID:              baseIDs.PrototypeIdentityID(),
				toID:                baseIDs.PrototypeIdentityID(),
				canMintAsset:        false,
				canRenumerateAsset:  true,
				canBurnAsset:        true,
				canAddMaintainer:    true,
				CanRemoveMaintainer: true,
				CanMutateMaintainer: true,
			},
			setup: func() {
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			got, err := s.TransactionKeeper.Transact(sdkTypes.WrapSDKContext(s.Context),
				NewMessage(tt.args.from,
					tt.args.fromID,
					tt.args.toID,
					baseIDs.PrototypeClassificationID(),
					baseLists.NewPropertyList(),
					tt.args.canMintAsset,
					tt.args.canRenumerateAsset,
					tt.args.canBurnAsset,
					tt.args.canAddMaintainer,
					tt.args.CanRemoveMaintainer,
					tt.args.CanMutateMaintainer).(helpers.Message))

			if tt.wantErr != nil {
				require.Error(t, err)
				require.Equal(t, tt.wantErr, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
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
			got := keeperPrototype()
			assert.Equal(t, tt.want, got, "keeperPrototype()")
		})
	}
}
