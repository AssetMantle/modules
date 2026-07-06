// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package get

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
	"github.com/AssetMantle/modules/x/orders/record"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	propertyConstants "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	baseTypes "github.com/AssetMantle/schema/types/base"
)

var (
	testMakerAssetID = baseDocuments.NewCoinAsset("makerAsset").GetCoinAssetID()
	testTakerAssetID = baseDocuments.NewCoinAsset("takerAsset").GetCoinAssetID()
)

// newStoredTestOrder builds an order document the way the make keeper
// persists it. The name property keeps each order's ID distinct.
func newStoredTestOrder(name string) (documents.Order, ids.OrderID) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(
		baseProperties.NewMetaProperty(baseIDs.NewStringID("orderName"), baseData.NewStringData(name)),
		baseProperties.NewMetaProperty(propertyConstants.MakerIDProperty.GetKey(), baseData.NewIDData(testFromID)),
		baseProperties.NewMetaProperty(propertyConstants.MakerAssetIDProperty.GetKey(), baseData.NewIDData(testMakerAssetID)),
		baseProperties.NewMetaProperty(propertyConstants.TakerAssetIDProperty.GetKey(), baseData.NewIDData(testTakerAssetID)),
		baseProperties.NewMetaProperty(propertyConstants.TakerIDProperty.GetKey(), baseData.NewIDData(baseIDs.PrototypeIdentityID())),
	))

	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(
		baseProperties.NewMetaProperty(propertyConstants.MakerSplitProperty.GetKey(), baseData.NewNumberData(math.NewInt(100))),
		baseProperties.NewMetaProperty(propertyConstants.TakerSplitProperty.GetKey(), baseData.NewNumberData(math.NewInt(1))),
		baseProperties.NewMetaProperty(propertyConstants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(100))),
	))

	order := baseDocuments.NewOrder(testClassificationID, immutables, mutables)
	return order, baseIDs.NewOrderID(testClassificationID, immutables)
}

type testSetup struct {
	Context                     sdkTypes.Context
	TransactionKeeper           transactionKeeper
	authenticateAuxiliaryKeeper *testutil.MockAuxiliaryKeeper
	supplementAuxiliaryKeeper   *testutil.MockAuxiliaryKeeper
	transferAuxiliaryKeeper     *testutil.MockAuxiliaryKeeper
	getOrderID                  ids.OrderID
	transferFailOrderID         ids.OrderID
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

	getOrder, getOrderID := newStoredTestOrder("getOrder")
	transferFailOrder, transferFailOrderID := newStoredTestOrder("transferFailOrder")

	TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(ctx)).
		Add(record.NewRecord(getOrder)).
		Add(record.NewRecord(transferFailOrder))

	return &testSetup{
		Context:                     ctx,
		TransactionKeeper:           TransactionKeeper,
		authenticateAuxiliaryKeeper: authenticateAuxiliaryKeeper,
		supplementAuxiliaryKeeper:   supplementAuxiliaryKeeper,
		transferAuxiliaryKeeper:     transferAuxiliaryKeeper,
		getOrderID:                  getOrderID,
		transferFailOrderID:         transferFailOrderID,
	}
}

func TestTransactionKeeperTransact(t *testing.T) {
	s := setupTest(t)

	tests := []struct {
		name    string
		orderID func() ids.OrderID
		setup   func()
		check   func(t *testing.T)
		wantErr helpers.Error
	}{
		{
			// Getting an order settles both sides (two transfers) and removes
			// the order from the store.
			name:    "getOrder",
			orderID: func() ids.OrderID { return s.getOrderID },
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Twice()
			},
			check: func(t *testing.T) {
				orders := s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Fetch(key.NewKey(s.getOrderID))
				assert.Nil(t, orders.GetMappable(key.NewKey(s.getOrderID)), "got order must be removed from the store")
			},
		},
		{
			name:    "orderNotFound",
			orderID: func() ids.OrderID { return testOrderID },
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.EntityNotFound,
		},
		{
			name:    "transferFailure",
			orderID: func() ids.OrderID { return s.transferFailOrderID },
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			wantErr: errorConstants.MockError,
		},
		{
			name:    "authenticationFailure",
			orderID: func() ids.OrderID { return s.transferFailOrderID },
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			wantErr: errorConstants.MockError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			message := NewMessage(fromAccAddress, testFromID, tt.orderID()).(helpers.Message)
			got, err := s.TransactionKeeper.Transact(sdkTypes.WrapSDKContext(s.Context), message)

			if tt.wantErr != nil {
				assert.True(t, tt.wantErr.Is(err), "Transact() error = %v, want %v", err, tt.wantErr)
				assert.Nil(t, got)
			} else {
				require.NoError(t, err)
				assert.Equal(t, newTransactionResponse(), got)
				if tt.check != nil {
					tt.check(t)
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
