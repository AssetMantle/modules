// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	storeTypes "cosmossdk.io/store/types"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/revoke"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/parameters/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

type testSetup struct {
	Context                            sdkTypes.Context
	genesisAddress                     sdkTypes.AccAddress
	TransactionKeeper                  transactionKeeper
	authenticateAuxiliaryFailureAddress sdkTypes.AccAddress
	revokeAuxiliaryFailureAddress      sdkTypes.AccAddress
	revokeAuxiliaryKeeperFailureID     ids.IdentityID
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

	revokeAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	revokeAuxiliaryFailureAddress := testutil.TestAddress()
	revokeAuxiliaryKeeperFailureID := baseIDs.NewIdentityID(baseIDs.NewClassificationID(immutables, mutables), immutables)
	revokeAuxiliaryKeeper.On("Help", mock.Anything, revoke.NewAuxiliaryRequest(revokeAuxiliaryKeeperFailureID, baseIDs.PrototypeIdentityID(), baseIDs.PrototypeClassificationID())).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	revokeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(define.NewAuxiliaryResponse(baseIDs.PrototypeClassificationID()), nil)
	revokeAuxiliary := new(testutil.MockAuxiliary)
	revokeAuxiliary.On("GetKeeper").Return(revokeAuxiliaryKeeper)

	parameterManager := parameters.Prototype().Initialize(moduleStoreKey).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(testutil.Denom))))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).
		Set(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(testutil.Denom)))))

	TransactionKeeper := transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey),
		parameterManager,
		authenticateAuxiliary,
		revokeAuxiliary,
	}

	return &testSetup{
		Context:                            ctx,
		genesisAddress:                     genesisAddress,
		TransactionKeeper:                  TransactionKeeper,
		authenticateAuxiliaryFailureAddress: authenticateAuxiliaryFailureAddress,
		revokeAuxiliaryFailureAddress:      revokeAuxiliaryFailureAddress,
		revokeAuxiliaryKeeperFailureID:     revokeAuxiliaryKeeperFailureID,
	}
}

func TestTransactionKeeperTransact(t *testing.T) {
	s := setupTest(t)

	type args struct {
		from             sdkTypes.AccAddress
		fromID           ids.IdentityID
		toID             ids.IdentityID
		classificationID ids.ClassificationID
	}

	tests := []struct {
		name    string
		args    args
		setup   func()
		want    *TransactionResponse
		wantErr helpers.Error
	}{
		{
			name: "RevokeTransactionKeeperSuccess",
			args: args{
				from:             s.genesisAddress,
				fromID:           baseIDs.PrototypeIdentityID(),
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
			},
			setup: func() {
			},
			want:    newTransactionResponse(),
			wantErr: nil,
		},
		{
			name: "AuthenticationFailure",
			args: args{
				from:             s.authenticateAuxiliaryFailureAddress,
				fromID:           baseIDs.PrototypeIdentityID(),
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
			},
			setup: func() {
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "RevokeAuxiliaryFailure",
			args: args{
				from:             s.revokeAuxiliaryFailureAddress,
				fromID:           s.revokeAuxiliaryKeeperFailureID,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
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
					tt.args.classificationID).(helpers.Message))

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
