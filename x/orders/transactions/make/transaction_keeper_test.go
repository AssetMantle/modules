// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

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
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/bond"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/authorize"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/orders/constants"
	"github.com/AssetMantle/modules/x/orders/key"
	"github.com/AssetMantle/modules/x/orders/mapper"
	"github.com/AssetMantle/modules/x/orders/parameters"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	propertyConstants "github.com/AssetMantle/schema/properties/constants"
	baseTypes "github.com/AssetMantle/schema/types/base"
)

type testSetup struct {
	Context                     sdkTypes.Context
	TransactionKeeper           transactionKeeper
	authenticateAuxiliaryKeeper *testutil.MockAuxiliaryKeeper
	authorizeAuxiliaryKeeper    *testutil.MockAuxiliaryKeeper
	bondAuxiliaryKeeper         *testutil.MockAuxiliaryKeeper
	conformAuxiliaryKeeper      *testutil.MockAuxiliaryKeeper
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
	authorizeAuxiliary, authorizeAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	bondAuxiliary, bondAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	conformAuxiliary, conformAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	supplementAuxiliary, supplementAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	transferAuxiliary, transferAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()

	TransactionKeeper := transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey), parameterManager, authenticateAuxiliary, authorizeAuxiliary, bondAuxiliary, conformAuxiliary, supplementAuxiliary, transferAuxiliary}

	return &testSetup{
		Context:                     ctx,
		TransactionKeeper:           TransactionKeeper,
		authenticateAuxiliaryKeeper: authenticateAuxiliaryKeeper,
		authorizeAuxiliaryKeeper:    authorizeAuxiliaryKeeper,
		bondAuxiliaryKeeper:         bondAuxiliaryKeeper,
		conformAuxiliaryKeeper:      conformAuxiliaryKeeper,
		supplementAuxiliaryKeeper:   supplementAuxiliaryKeeper,
		transferAuxiliaryKeeper:     transferAuxiliaryKeeper,
	}
}

// newTestMakeMessage builds a make message with fresh property lists so that
// keeper-side list mutations cannot leak between test cases. TakerSplit is
// revealed in the mutable meta properties (order.ValidateBasic requires it)
// and BondAmount is included when withBondAmount is true (the keeper requires
// a revealed bond amount). The maker split must be at least the taker split:
// the keeper-computed exchange rate (takerSplit/makerSplit scaled by 1e18)
// panics in LegacySortableDecBytes when it exceeds the sortable bound of 1e18.
// Varying makerSplitValue yields distinct exchange rates and thus distinct
// order IDs.
func newTestMakeMessage(makerSplitValue int64, expiresIn int64, withBondAmount bool) helpers.Message {
	mutableMetaProperties := baseLists.NewPropertyList(
		baseProperties.NewMetaProperty(propertyConstants.TakerSplitProperty.GetKey(), baseData.NewNumberData(math.NewInt(1))),
	)
	if withBondAmount {
		mutableMetaProperties = mutableMetaProperties.Add(
			baseProperties.NewMetaProperty(propertyConstants.BondAmountProperty.GetKey(), baseData.NewNumberData(math.NewInt(1))),
		)
	}

	return NewMessage(
		fromAccAddress,
		testFromID,
		testClassificationID,
		baseIDs.PrototypeIdentityID(),
		makerAssetID,
		takerAssetID,
		baseTypes.NewHeight(expiresIn),
		math.NewInt(makerSplitValue),
		math.NewInt(1),
		baseLists.NewPropertyList(),
		baseLists.NewPropertyList(),
		mutableMetaProperties,
		baseLists.NewPropertyList(),
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
			name:    "makeOrder",
			message: newTestMakeMessage(1, 10, true),
			setup: func() {
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Twice()
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.bondAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			check: func(t *testing.T, got helpers.TransactionResponse) {
				orderID := got.(*TransactionResponse).OrderID
				require.NotNil(t, orderID)
				orders := s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Fetch(key.NewKey(orderID))
				assert.NotNil(t, orders.GetMappable(key.NewKey(orderID)), "order must be added to the store")
			},
		},
		{
			name:    "duplicateOrder",
			message: newTestMakeMessage(1, 10, true),
			setup: func() {
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Twice()
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.EntityAlreadyExists,
		},
		{
			name:    "authenticationFailure",
			message: newTestMakeMessage(1, 10, true),
			setup: func() {
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Twice()
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			wantErr: errorConstants.MockError,
		},
		{
			name:    "expiryExceedsMaxOrderLife",
			message: newTestMakeMessage(2, 50000, true),
			setup: func() {
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Twice()
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.InvalidRequest,
		},
		{
			name:    "bondAmountNotRevealed",
			message: newTestMakeMessage(3, 10, false),
			setup: func() {
				s.authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Twice()
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.MetaDataError,
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
	authorizeAuxiliary, _ := testutil.NewNamedMockAuxiliaryPair(authorize.Auxiliary.GetName())
	bondAuxiliary, _ := testutil.NewNamedMockAuxiliaryPair(bond.Auxiliary.GetName())
	conformAuxiliary, _ := testutil.NewNamedMockAuxiliaryPair(conform.Auxiliary.GetName())
	supplementAuxiliary, _ := testutil.NewNamedMockAuxiliaryPair(supplement.Auxiliary.GetName())
	transferAuxiliary, _ := testutil.NewNamedMockAuxiliaryPair(transfer.Auxiliary.GetName())

	keeper := keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{authenticateAuxiliary, authorizeAuxiliary, bondAuxiliary, conformAuxiliary, supplementAuxiliary, transferAuxiliary})
	require.NotNil(t, keeper)
}
