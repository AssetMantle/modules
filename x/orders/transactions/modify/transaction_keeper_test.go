// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

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
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/orders/constants"
	"github.com/AssetMantle/modules/x/orders/key"
	"github.com/AssetMantle/modules/x/orders/mappable"
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
// persists it, with an initial maker split of 100.
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

func newTestModifyMessage(orderID ids.OrderID, makerSplitValue int64) helpers.Message {
	return NewMessage(
		fromAccAddress,
		testFromID,
		orderID,
		math.NewInt(1),
		math.NewInt(makerSplitValue),
		baseTypes.NewHeight(10),
		baseLists.NewPropertyList(),
		baseLists.NewPropertyList(),
	).(helpers.Message)
}

type testSetup struct {
	Context                     sdkTypes.Context
	TransactionKeeper           transactionKeeper
	authenticateAuxiliaryKeeper *testutil.MockAuxiliaryKeeper
	conformAuxiliaryKeeper      *testutil.MockAuxiliaryKeeper
	supplementAuxiliaryKeeper   *testutil.MockAuxiliaryKeeper
	transferAuxiliaryKeeper     *testutil.MockAuxiliaryKeeper
	modifyOrderID               ids.OrderID
}

func setupTest(t *testing.T) *testSetup {
	t.Helper()

	moduleStoreKey := storeTypes.NewKVStoreKey(constants.ModuleName)
	ctx := testutil.NewTestContext(t, moduleStoreKey)

	parameterManager, err := parameters.Prototype().Initialize(moduleStoreKey).Set().Update(sdkTypes.WrapSDKContext(ctx))
	require.NoError(t, err)

	authenticateAuxiliary, authenticateAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	conformAuxiliary, conformAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	supplementAuxiliary, supplementAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	transferAuxiliary, transferAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()

	TransactionKeeper := transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey), parameterManager, conformAuxiliary, supplementAuxiliary, transferAuxiliary, authenticateAuxiliary}

	modifyOrder, modifyOrderID := newStoredTestOrder("modifyOrder")

	TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(ctx)).Add(record.NewRecord(modifyOrder))

	return &testSetup{
		Context:                     ctx,
		TransactionKeeper:           TransactionKeeper,
		authenticateAuxiliaryKeeper: authenticateAuxiliaryKeeper,
		conformAuxiliaryKeeper:      conformAuxiliaryKeeper,
		supplementAuxiliaryKeeper:   supplementAuxiliaryKeeper,
		transferAuxiliaryKeeper:     transferAuxiliaryKeeper,
		modifyOrderID:               modifyOrderID,
	}
}

func TestTransactionKeeperTransact(t *testing.T) {
	s := setupTest(t)

	tests := []struct {
		name    string
		message func() helpers.Message
		setup   func()
		check   func(t *testing.T)
		wantErr helpers.Error
	}{
		{
			// Keeping the maker split unchanged skips the escrow transfer and
			// only refreshes the mutables.
			name:    "modifyOrderSameSplit",
			message: func() helpers.Message { return newTestModifyMessage(s.modifyOrderID, 100) },
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			check: func(t *testing.T) {
				orders := s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Fetch(key.NewKey(s.modifyOrderID))
				Mappable := orders.GetMappable(key.NewKey(s.modifyOrderID))
				require.NotNil(t, Mappable, "modified order must remain in the store")
				assert.True(t, mappable.GetOrder(Mappable).GetMakerSplit().Equal(math.NewInt(100)))
			},
		},
		{
			// Increasing the maker split escrows the difference with one
			// transfer and records the new split.
			name:    "modifyOrderIncreaseSplit",
			message: func() helpers.Message { return newTestModifyMessage(s.modifyOrderID, 150) },
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			check: func(t *testing.T) {
				orders := s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Fetch(key.NewKey(s.modifyOrderID))
				Mappable := orders.GetMappable(key.NewKey(s.modifyOrderID))
				require.NotNil(t, Mappable, "modified order must remain in the store")
				assert.True(t, mappable.GetOrder(Mappable).GetMakerSplit().Equal(math.NewInt(150)), "increased maker split must be recorded")
			},
		},
		{
			name:    "orderNotFound",
			message: func() helpers.Message { return newTestModifyMessage(testOrderID, 100) },
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.EntityNotFound,
		},
		{
			name: "invalidMakerSplit",
			message: func() helpers.Message {
				message := newTestModifyMessage(s.modifyOrderID, 100)
				message.(*Message).MakerSplit = "invalidSplit"
				return message
			},
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.IncorrectFormat,
		},
		{
			name:    "conformFailure",
			message: func() helpers.Message { return newTestModifyMessage(s.modifyOrderID, 150) },
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			wantErr: errorConstants.MockError,
		},
		{
			name:    "authenticationFailure",
			message: func() helpers.Message { return newTestModifyMessage(s.modifyOrderID, 100) },
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			wantErr: errorConstants.MockError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			got, err := s.TransactionKeeper.Transact(sdkTypes.WrapSDKContext(s.Context), tt.message())

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
	conformAuxiliary, _ := testutil.NewNamedMockAuxiliaryPair(conform.Auxiliary.GetName())
	supplementAuxiliary, _ := testutil.NewNamedMockAuxiliaryPair(supplement.Auxiliary.GetName())
	transferAuxiliary, _ := testutil.NewNamedMockAuxiliaryPair(transfer.Auxiliary.GetName())

	keeper := keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{authenticateAuxiliary, conformAuxiliary, supplementAuxiliary, transferAuxiliary})
	require.NotNil(t, keeper)
}
