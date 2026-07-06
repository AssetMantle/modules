// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package take

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
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/burn"
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
	// oneToOneExchangeRate is the rate the make keeper computes for equal
	// maker and taker splits: 1/1e-18 = 1e18.
	oneToOneExchangeRate = math.LegacyOneDec().Quo(math.LegacySmallestDec())
)

// newStoredTestOrder builds an order document the way the make keeper
// persists it: maker-side identifiers and the exchange rate in the immutable
// meta properties, splits and the bond amount in the mutable meta properties.
// The name property keeps each order's ID distinct.
func newStoredTestOrder(name string, makerSplitValue int64, takerID ids.IdentityID, withBondAmount bool) (documents.Order, ids.OrderID) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(
		baseProperties.NewMetaProperty(baseIDs.NewStringID("orderName"), baseData.NewStringData(name)),
		baseProperties.NewMetaProperty(propertyConstants.MakerIDProperty.GetKey(), baseData.NewIDData(testFromID)),
		baseProperties.NewMetaProperty(propertyConstants.MakerAssetIDProperty.GetKey(), baseData.NewIDData(testMakerAssetID)),
		baseProperties.NewMetaProperty(propertyConstants.TakerAssetIDProperty.GetKey(), baseData.NewIDData(testTakerAssetID)),
		baseProperties.NewMetaProperty(propertyConstants.TakerIDProperty.GetKey(), baseData.NewIDData(takerID)),
		baseProperties.NewMetaProperty(propertyConstants.ExchangeRateProperty.GetKey(), baseData.NewDecData(oneToOneExchangeRate)),
	))

	mutableMetaProperties := baseLists.NewPropertyList(
		baseProperties.NewMetaProperty(propertyConstants.MakerSplitProperty.GetKey(), baseData.NewNumberData(math.NewInt(makerSplitValue))),
		baseProperties.NewMetaProperty(propertyConstants.TakerSplitProperty.GetKey(), baseData.NewNumberData(math.NewInt(1))),
		baseProperties.NewMetaProperty(propertyConstants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(100))),
	)
	if withBondAmount {
		mutableMetaProperties = mutableMetaProperties.Add(
			baseProperties.NewMetaProperty(propertyConstants.BondAmountProperty.GetKey(), baseData.NewNumberData(math.NewInt(1))),
		)
	}
	mutables := baseQualified.NewMutables(mutableMetaProperties)

	order := baseDocuments.NewOrder(testClassificationID, immutables, mutables)
	return order, baseIDs.NewOrderID(testClassificationID, immutables)
}

type testSetup struct {
	Context                     sdkTypes.Context
	TransactionKeeper           transactionKeeper
	authenticateAuxiliaryKeeper *testutil.MockAuxiliaryKeeper
	burnAuxiliaryKeeper         *testutil.MockAuxiliaryKeeper
	supplementAuxiliaryKeeper   *testutil.MockAuxiliaryKeeper
	transferAuxiliaryKeeper     *testutil.MockAuxiliaryKeeper
	fullOrderID                 ids.OrderID
	partialOrderID              ids.OrderID
	privateOrderID              ids.OrderID
	noBondOrderID               ids.OrderID
}

func setupTest(t *testing.T) *testSetup {
	t.Helper()

	moduleStoreKey := storeTypes.NewKVStoreKey(constants.ModuleName)
	ctx := testutil.NewTestContext(t, moduleStoreKey)

	parameterManager, err := parameters.Prototype().Initialize(moduleStoreKey).Set().Update(sdkTypes.WrapSDKContext(ctx))
	require.NoError(t, err)

	authenticateAuxiliary, authenticateAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	burnAuxiliary, burnAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	supplementAuxiliary, supplementAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	transferAuxiliary, transferAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()

	TransactionKeeper := transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey), parameterManager, authenticateAuxiliary, burnAuxiliary, supplementAuxiliary, transferAuxiliary}

	fullOrder, fullOrderID := newStoredTestOrder("fullOrder", 100, baseIDs.PrototypeIdentityID(), true)
	partialOrder, partialOrderID := newStoredTestOrder("partialOrder", 100, baseIDs.PrototypeIdentityID(), true)
	privateOrder, privateOrderID := newStoredTestOrder("privateOrder", 100, testutil.TestIdentityID(), true)
	noBondOrder, noBondOrderID := newStoredTestOrder("noBondOrder", 100, baseIDs.PrototypeIdentityID(), false)

	TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(ctx)).
		Add(record.NewRecord(fullOrder)).
		Add(record.NewRecord(partialOrder)).
		Add(record.NewRecord(privateOrder)).
		Add(record.NewRecord(noBondOrder))

	return &testSetup{
		Context:                     ctx,
		TransactionKeeper:           TransactionKeeper,
		authenticateAuxiliaryKeeper: authenticateAuxiliaryKeeper,
		burnAuxiliaryKeeper:         burnAuxiliaryKeeper,
		supplementAuxiliaryKeeper:   supplementAuxiliaryKeeper,
		transferAuxiliaryKeeper:     transferAuxiliaryKeeper,
		fullOrderID:                 fullOrderID,
		partialOrderID:              partialOrderID,
		privateOrderID:              privateOrderID,
		noBondOrderID:               noBondOrderID,
	}
}

func TestTransactionKeeperTransact(t *testing.T) {
	s := setupTest(t)

	tests := []struct {
		name       string
		takerSplit int64
		orderID    func() ids.OrderID
		setup      func()
		check      func(t *testing.T)
		wantErr    helpers.Error
	}{
		{
			// Taking the full maker split removes the order and settles both
			// sides: two transfers plus the bond burn.
			name:       "takeOrderFully",
			takerSplit: 100,
			orderID:    func() ids.OrderID { return s.fullOrderID },
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Twice()
				s.burnAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			check: func(t *testing.T) {
				orders := s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Fetch(key.NewKey(s.fullOrderID))
				assert.Nil(t, orders.GetMappable(key.NewKey(s.fullOrderID)), "fully taken order must be removed from the store")
			},
		},
		{
			// A partial take mutates the order, leaving the remaining maker
			// split on the book.
			name:       "takeOrderPartially",
			takerSplit: 40,
			orderID:    func() ids.OrderID { return s.partialOrderID },
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Twice()
				s.burnAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			check: func(t *testing.T) {
				orders := s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Fetch(key.NewKey(s.partialOrderID))
				Mappable := orders.GetMappable(key.NewKey(s.partialOrderID))
				require.NotNil(t, Mappable, "partially taken order must remain in the store")
				assert.True(t, mappable.GetOrder(Mappable).GetMakerSplit().Equal(math.NewInt(60)), "remaining maker split must be recorded")
			},
		},
		{
			name:       "orderNotFound",
			takerSplit: 100,
			orderID:    func() ids.OrderID { return testOrderID },
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.EntityNotFound,
		},
		{
			name:       "privateOrderTakerMismatch",
			takerSplit: 100,
			orderID:    func() ids.OrderID { return s.privateOrderID },
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.NotAuthorized,
		},
		{
			name:       "bondAmountNotRevealed",
			takerSplit: 100,
			orderID:    func() ids.OrderID { return s.noBondOrderID },
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Twice()
			},
			wantErr: errorConstants.MetaDataError,
		},
		{
			name:       "authenticationFailure",
			takerSplit: 100,
			orderID:    func() ids.OrderID { return s.fullOrderID },
			setup: func() {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			wantErr: errorConstants.MockError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			message := NewMessage(fromAccAddress, testFromID, math.NewInt(tt.takerSplit), tt.orderID()).(helpers.Message)
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
	burnAuxiliary, _ := testutil.NewNamedMockAuxiliaryPair(burn.Auxiliary.GetName())
	supplementAuxiliary, _ := testutil.NewNamedMockAuxiliaryPair(supplement.Auxiliary.GetName())
	transferAuxiliary, _ := testutil.NewNamedMockAuxiliaryPair(transfer.Auxiliary.GetName())

	keeper := keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{authenticateAuxiliary, burnAuxiliary, supplementAuxiliary, transferAuxiliary})
	require.NotNil(t, keeper)
}
