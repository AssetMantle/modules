// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	storeTypes "cosmossdk.io/store/types"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
	permissionHelper "github.com/AssetMantle/modules/x/assets/utilities"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/super"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"

	dataHelper "github.com/AssetMantle/modules/simulation/schema/types/base"
)

type testSetup struct {
	Context                            sdkTypes.Context
	genesisAddress                     sdkTypes.AccAddress
	TransactionKeeper                  transactionKeeper
	authenticateAuxiliaryFailureAddress sdkTypes.AccAddress
	defineAuxiliaryKeeperFailureAddress sdkTypes.AccAddress
	superAuxiliaryMutablesFailure       lists.PropertyList
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

	defineAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	defineAuxiliaryKeeperFailureAddress := testutil.TestAddress()
	defineAuxiliaryKeeper.On("Help", mock.Anything, define.NewAuxiliaryRequest(defineAuxiliaryKeeperFailureAddress, baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList()))).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	defineAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(define.NewAuxiliaryResponse(baseIDs.PrototypeClassificationID()), nil)
	defineAuxiliary := new(testutil.MockAuxiliary)
	defineAuxiliary.On("GetKeeper").Return(defineAuxiliaryKeeper)

	superAuxiliaryKeeper := new(testutil.MockAuxiliaryKeeper)
	superAuxiliaryMutablesFailure := dataHelper.GenerateRandomMetaPropertyListWithoutData(rand.New(rand.NewSource(99)))
	superAuxiliaryKeeper.On("Help", mock.Anything, super.NewAuxiliaryRequest(baseIDs.PrototypeClassificationID(), baseIDs.PrototypeIdentityID(), baseQualified.NewMutables(superAuxiliaryMutablesFailure), permissionHelper.SetModulePermissions(true, true, true)...)).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	superAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	superAuxiliary := new(testutil.MockAuxiliary)
	superAuxiliary.On("GetKeeper").Return(superAuxiliaryKeeper)

	TransactionKeeper := transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey),
		defineAuxiliary,
		superAuxiliary,
		authenticateAuxiliary,
	}

	return &testSetup{
		Context:                            ctx,
		genesisAddress:                     genesisAddress,
		TransactionKeeper:                  TransactionKeeper,
		authenticateAuxiliaryFailureAddress: authenticateAuxiliaryFailureAddress,
		defineAuxiliaryKeeperFailureAddress: defineAuxiliaryKeeperFailureAddress,
		superAuxiliaryMutablesFailure:       superAuxiliaryMutablesFailure,
	}
}

func TestTransactionKeeperTransact(t *testing.T) {
	s := setupTest(t)

	type args struct {
		from               sdkTypes.AccAddress
		fromID             ids.IdentityID
		immutableMetaProps lists.PropertyList
		immutableProps     lists.PropertyList
		mutableMetaProps   lists.PropertyList
		mutableProps       lists.PropertyList
	}

	tests := []struct {
		name    string
		args    args
		setup   func()
		want    *TransactionResponse
		wantErr helpers.Error
	}{
		{
			name: "DefineTransactionKeeperSuccess",
			args: args{
				from:               s.genesisAddress,
				fromID:             baseIDs.PrototypeIdentityID(),
				immutableMetaProps: baseLists.NewPropertyList(),
				immutableProps:     baseLists.NewPropertyList(),
				mutableMetaProps:   baseLists.NewPropertyList(),
				mutableProps:       baseLists.NewPropertyList(),
			},
			setup: func() {
			},
			want:    newTransactionResponse(baseIDs.PrototypeClassificationID()),
			wantErr: nil,
		},
		{
			name: "AuthenticationFailure",
			args: args{
				from:               s.authenticateAuxiliaryFailureAddress,
				fromID:             baseIDs.PrototypeIdentityID(),
				immutableMetaProps: baseLists.NewPropertyList(),
				immutableProps:     baseLists.NewPropertyList(),
				mutableMetaProps:   baseLists.NewPropertyList(),
				mutableProps:       baseLists.NewPropertyList(),
			},
			setup: func() {
			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "DefineAuxiliaryFailure",
			args: args{
				from:               s.defineAuxiliaryKeeperFailureAddress,
				fromID:             baseIDs.PrototypeIdentityID(),
				immutableMetaProps: baseLists.NewPropertyList(),
				immutableProps:     baseLists.NewPropertyList(),
				mutableMetaProps:   baseLists.NewPropertyList(),
				mutableProps:       baseLists.NewPropertyList(),
			},
			setup: func() {

			},
			want:    nil,
			wantErr: errorConstants.MockError,
		},
		{
			name: "SuperAuxiliaryFailure",
			args: args{
				from:               s.genesisAddress,
				fromID:             baseIDs.PrototypeIdentityID(),
				immutableMetaProps: baseLists.NewPropertyList(),
				immutableProps:     baseLists.NewPropertyList(),
				mutableMetaProps:   baseLists.NewPropertyList(),
				mutableProps:       s.superAuxiliaryMutablesFailure,
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
					tt.args.immutableMetaProps.(lists.PropertyList),
					tt.args.immutableProps.(lists.PropertyList),
					tt.args.mutableMetaProps.(lists.PropertyList),
					tt.args.mutableProps.(lists.PropertyList)).(helpers.Message))

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
	got := keeperPrototype()
	assert.Equal(t, transactionKeeper{}, got)
}

func Test_transactionKeeper_Initialize(t *testing.T) {
	moduleStoreKey := storeTypes.NewKVStoreKey(constants.ModuleName)
	m := mapper.Prototype().Initialize(moduleStoreKey)
	pm := parameters.Prototype().Initialize(moduleStoreKey)

	defineAux, _ := testutil.NewNamedMockAuxiliaryPair(define.Auxiliary.GetName())
	superAux, _ := testutil.NewNamedMockAuxiliaryPair(super.Auxiliary.GetName())
	authenticateAux, _ := testutil.NewNamedMockAuxiliaryPair(authenticate.Auxiliary.GetName())

	keeper := keeperPrototype().Initialize(m, pm, []interface{}{defineAux, superAux, authenticateAux})
	require.NotNil(t, keeper)
}
