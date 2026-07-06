// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package put

import (
	"testing"

	"cosmossdk.io/math"
	storeTypes "cosmossdk.io/store/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/orders/constants"
	"github.com/AssetMantle/modules/x/orders/key"
	"github.com/AssetMantle/modules/x/orders/mapper"
	"github.com/AssetMantle/modules/x/orders/parameters"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	baseParameters "github.com/AssetMantle/schema/parameters/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseTypes "github.com/AssetMantle/schema/types/base"
)

type testSetup struct {
	Context                     sdkTypes.Context
	TransactionKeeper           transactionKeeper
	parameterManager            helpers.ParameterManager
	authenticateAuxiliaryKeeper *testutil.MockAuxiliaryKeeper
	supplementAuxiliaryKeeper   *testutil.MockAuxiliaryKeeper
	transferAuxiliaryKeeper     *testutil.MockAuxiliaryKeeper
}

func setupTest(t *testing.T) *testSetup {
	t.Helper()

	moduleStoreKey := storeTypes.NewKVStoreKey(constants.ModuleName)
	ctx := testutil.NewTestContext(t, moduleStoreKey)

	parameterManager, err := parameters.Prototype().Initialize(moduleStoreKey).Set().Update(sdkTypes.WrapSDKContext(ctx))
	require.NoError(t, err)

	authenticateAuxiliary, authenticateAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	supplementAuxiliary, supplementAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	transferAuxiliary, transferAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()

	TransactionKeeper := transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey), parameterManager, authenticateAuxiliary, supplementAuxiliary, transferAuxiliary}

	return &testSetup{
		Context:                     ctx,
		TransactionKeeper:           TransactionKeeper,
		parameterManager:            parameterManager,
		authenticateAuxiliaryKeeper: authenticateAuxiliaryKeeper,
		supplementAuxiliaryKeeper:   supplementAuxiliaryKeeper,
		transferAuxiliaryKeeper:     transferAuxiliaryKeeper,
	}
}

func newTestPutMessage(makerSplitValue int64, takerSplitValue int64, expiryHeightValue int64) helpers.Message {
	return NewMessage(
		fromAccAddress,
		testFromID,
		makerAssetID,
		takerAssetID,
		math.NewInt(makerSplitValue),
		math.NewInt(takerSplitValue),
		baseTypes.NewHeight(expiryHeightValue),
	).(helpers.Message)
}

func TestTransactionKeeperTransact(t *testing.T) {
	s := setupTest(t)

	tests := []struct {
		name    string
		message helpers.Message
		setup   func()
		check   func(t *testing.T, got helpers.TransactionResponse)
		wantErr helpers.Error
	}{
		{
			name:    "putDisabled",
			message: newTestPutMessage(100, 100, 100),
			setup: func() {
				_, err := s.parameterManager.Set(baseParameters.NewParameter(baseProperties.NewMetaProperty(constantProperties.PutEnabledProperty.GetKey(), baseData.NewBooleanData(false)))).Update(sdkTypes.WrapSDKContext(s.Context))
				require.NoError(t, err)
			},
			wantErr: errorConstants.NotAuthorized,
		},
		{
			name:    "putOrder",
			message: newTestPutMessage(100, 100, 100),
			setup: func() {
				_, err := s.parameterManager.Set(baseParameters.NewParameter(baseProperties.NewMetaProperty(constantProperties.PutEnabledProperty.GetKey(), baseData.NewBooleanData(true)))).Update(sdkTypes.WrapSDKContext(s.Context))
				require.NoError(t, err)
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			check: func(t *testing.T, got helpers.TransactionResponse) {
				expectedPutOrderID := baseDocuments.NewPutOrder(testFromID, makerAssetID, takerAssetID, math.NewInt(100), math.NewInt(100), baseTypes.NewHeight(100)).GetPutOrderID()
				assert.Equal(t, newTransactionResponse(expectedPutOrderID), got)

				orders := s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Fetch(key.NewKey(expectedPutOrderID))
				assert.NotNil(t, orders.GetMappable(key.NewKey(expectedPutOrderID)), "put order must be added to the store")
			},
		},
		{
			name:    "duplicateOrder",
			message: newTestPutMessage(100, 100, 100),
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.EntityAlreadyExists,
		},
		{
			name:    "negativeMakerSplit",
			message: newTestPutMessage(-1, 100, 100),
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.IncorrectFormat,
		},
		{
			name:    "negativeTakerSplit",
			message: newTestPutMessage(100, -1, 100),
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.IncorrectFormat,
		},
		{
			name:    "expiryInPast",
			message: newTestPutMessage(100, 100, 0),
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.InvalidRequest,
		},
		{
			name:    "expiryExceedsMaxOrderLife",
			message: newTestPutMessage(100, 100, 50000),
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.InvalidRequest,
		},
		{
			name:    "authenticationFailure",
			message: newTestPutMessage(100, 100, 100),
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			wantErr: errorConstants.MockError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			got, err := s.TransactionKeeper.Transact(sdkTypes.WrapSDKContext(s.Context), tt.message)

			if tt.wantErr != nil {
				assert.True(t, tt.wantErr.Is(err), "Transact() error = %v, want %v", err, tt.wantErr)
				assert.Nil(t, got)
			} else {
				require.NoError(t, err)
				require.NotNil(t, got)
				if tt.check != nil {
					tt.check(t, got)
				}
			}
		})
	}

	t.Run("nilMessagePanics", func(t *testing.T) {
		require.Panics(t, func() {
			_, _ = s.TransactionKeeper.Transact(sdkTypes.WrapSDKContext(s.Context), nil)
		})
	})
}

func Test_keeperPrototype(t *testing.T) {
	assert.Equal(t, transactionKeeper{}, keeperPrototype())
}

func Test_transactionKeeper_Initialize(t *testing.T) {
	moduleStoreKey := storeTypes.NewKVStoreKey(constants.ModuleName)
	Mapper := mapper.Prototype().Initialize(moduleStoreKey)
	parameterManager := parameters.Prototype().Initialize(moduleStoreKey)

	authenticateAuxiliary, _ := testutil.NewNamedMockAuxiliaryPair(authenticate.Auxiliary.GetName())
	supplementAuxiliary, _ := testutil.NewNamedMockAuxiliaryPair(supplement.Auxiliary.GetName())
	transferAuxiliary, _ := testutil.NewNamedMockAuxiliaryPair(transfer.Auxiliary.GetName())

	keeper := keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{authenticateAuxiliary, supplementAuxiliary, transferAuxiliary})
	require.NotNil(t, keeper)
}
