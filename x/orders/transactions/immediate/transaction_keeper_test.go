// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package immediate

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
	documentsBase "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	propertyConstants "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	baseTypes "github.com/AssetMantle/schema/types/base"
)

type testSetup struct {
	Context                     sdkTypes.Context
	TransactionKeeper           transactionKeeper
	authenticateAuxiliaryKeeper *testutil.MockAuxiliaryKeeper
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
	conformAuxiliary, conformAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	supplementAuxiliary, supplementAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()
	transferAuxiliary, transferAuxiliaryKeeper := testutil.NewMockAuxiliaryPair()

	TransactionKeeper := transactionKeeper{mapper.Prototype().Initialize(moduleStoreKey), parameterManager, conformAuxiliary, supplementAuxiliary, transferAuxiliary, authenticateAuxiliary}

	return &testSetup{
		Context:                     ctx,
		TransactionKeeper:           TransactionKeeper,
		authenticateAuxiliaryKeeper: authenticateAuxiliaryKeeper,
		conformAuxiliaryKeeper:      conformAuxiliaryKeeper,
		supplementAuxiliaryKeeper:   supplementAuxiliaryKeeper,
		transferAuxiliaryKeeper:     transferAuxiliaryKeeper,
	}
}

// newTestImmediateMessage builds an immediate order message with fresh property
// lists so that keeper-side list mutations cannot leak between test cases.
// TakerSplit is revealed in the mutable meta properties (order.ValidateBasic
// requires it). The maker split must be at least the taker split: the
// keeper-computed exchange rate (takerSplit/makerSplit scaled by 1e18) panics in
// LegacySortableDecBytes when it exceeds the sortable bound of 1e18.
func newTestImmediateMessage(makerAsset, takerAsset ids.AssetID, makerSplitValue int64, expiresIn int64) helpers.Message {
	mutableMetaProperties := baseLists.NewPropertyList(
		baseProperties.NewMetaProperty(propertyConstants.TakerSplitProperty.GetKey(), baseData.NewNumberData(math.NewInt(1))),
	)

	return NewMessage(
		fromAccAddress,
		testFromID,
		testClassificationID,
		baseIDs.PrototypeIdentityID(),
		makerAsset,
		takerAsset,
		baseTypes.NewHeight(expiresIn),
		math.NewInt(makerSplitValue),
		math.NewInt(1),
		baseLists.NewPropertyList(),
		baseLists.NewPropertyList(),
		mutableMetaProperties,
		baseLists.NewPropertyList(),
	).(helpers.Message)
}

// restingExchangeRate mirrors the keeper's stored exchange rate:
// takerSplit / 1e-18 / makerSplit.
func restingExchangeRate(makerSplitValue, takerSplitValue int64) math.LegacyDec {
	return math.NewInt(takerSplitValue).ToLegacyDec().QuoTruncate(math.LegacySmallestDec()).QuoTruncate(math.NewInt(makerSplitValue).ToLegacyDec())
}

// orderIDFor mirrors the keeper's order ID derivation for a message built by
// newTestImmediateMessage at block height zero (taker split 1).
func orderIDFor(makerAsset, takerAsset ids.AssetID, makerSplitValue int64) ids.OrderID {
	return orderIDForRate(makerAsset, takerAsset, makerSplitValue, 1)
}

func orderIDForRate(makerAsset, takerAsset ids.AssetID, makerSplitValue, takerSplitValue int64) ids.OrderID {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(
		baseProperties.NewMetaProperty(propertyConstants.ExchangeRateProperty.GetKey(), baseData.NewDecData(restingExchangeRate(makerSplitValue, takerSplitValue))),
		baseProperties.NewMetaProperty(propertyConstants.CreationHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(0))),
		baseProperties.NewMetaProperty(propertyConstants.MakerAssetIDProperty.GetKey(), baseData.NewIDData(makerAsset)),
		baseProperties.NewMetaProperty(propertyConstants.TakerAssetIDProperty.GetKey(), baseData.NewIDData(takerAsset)),
		baseProperties.NewMetaProperty(propertyConstants.MakerIDProperty.GetKey(), baseData.NewIDData(testFromID)),
		baseProperties.NewMetaProperty(propertyConstants.TakerIDProperty.GetKey(), baseData.NewIDData(baseIDs.PrototypeIdentityID())),
	))
	return baseIDs.NewOrderID(testClassificationID, immutables)
}

// restOrder constructs and stores a resting order directly (bypassing the
// keeper), used to seed the order book with counterparties. makerAsset/takerAsset
// name the asset the resting order escrows and the asset it wants; its exchange
// rate is takerSplit/makerSplit scaled by 1e18 and it demands taker split equal
// to takerSplitValue.
func restOrder(t *testing.T, s *testSetup, makerAsset, takerAsset ids.AssetID, makerSplitValue, takerSplitValue int64) documents.Order {
	t.Helper()
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(
		baseProperties.NewMetaProperty(propertyConstants.ExchangeRateProperty.GetKey(), baseData.NewDecData(restingExchangeRate(makerSplitValue, takerSplitValue))),
		baseProperties.NewMetaProperty(propertyConstants.CreationHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(0))),
		baseProperties.NewMetaProperty(propertyConstants.MakerAssetIDProperty.GetKey(), baseData.NewIDData(makerAsset)),
		baseProperties.NewMetaProperty(propertyConstants.TakerAssetIDProperty.GetKey(), baseData.NewIDData(takerAsset)),
		baseProperties.NewMetaProperty(propertyConstants.MakerIDProperty.GetKey(), baseData.NewIDData(testFromID)),
		baseProperties.NewMetaProperty(propertyConstants.TakerIDProperty.GetKey(), baseData.NewIDData(baseIDs.PrototypeIdentityID())),
	))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(
		baseProperties.NewMetaProperty(propertyConstants.TakerSplitProperty.GetKey(), baseData.NewNumberData(math.NewInt(takerSplitValue))),
		baseProperties.NewMetaProperty(propertyConstants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(100))),
		baseProperties.NewMetaProperty(propertyConstants.MakerSplitProperty.GetKey(), baseData.NewNumberData(math.NewInt(makerSplitValue))),
	))
	order := documentsBase.NewOrder(testClassificationID, immutables, mutables)
	s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Add(record.NewRecord(order))
	return order
}

func orderInStore(t *testing.T, s *testSetup, orderID ids.OrderID) documents.Order {
	t.Helper()
	Mappable := s.TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(s.Context)).Fetch(key.NewKey(orderID)).GetMappable(key.NewKey(orderID))
	if Mappable == nil {
		return nil
	}
	return mappable.GetOrder(Mappable)
}

func TestTransactionKeeperTransact(t *testing.T) {
	tests := []struct {
		name    string
		message helpers.Message
		seed    func(s *testSetup)
		setup   func(s *testSetup)
		check   func(t *testing.T, s *testSetup)
		wantErr helpers.Error
	}{
		{
			// An immediate order with no compatible counterparty in the book must
			// NOT self-execute; it rests with its full maker split, and only the
			// escrow transfer fires. This is the regression guard for the
			// self-match bug where the matcher iterated only the just-inserted
			// order and executed against itself.
			name:    "restsWhenBookEmpty",
			message: newTestImmediateMessage(makerAssetID, takerAssetID, 1, 10),
			setup: func(s *testSetup) {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			check: func(t *testing.T, s *testSetup) {
				order := orderInStore(t, s, orderIDFor(makerAssetID, takerAssetID, 1))
				require.NotNil(t, order, "order with no counterparty must rest in the store")
				assert.True(t, order.GetMakerSplit().Equal(math.NewInt(1)), "resting order keeps its full maker split")
				s.transferAuxiliaryKeeper.AssertNumberOfCalls(t, "Help", 1)
			},
		},
		{
			// A resting counterparty (mirror asset pair, crossing 1:1 rate) is
			// matched: escrow transfer plus two settlement transfers, the incoming
			// order fully fills and is removed, and the counterparty is consumed.
			name:    "executesAgainstCounterparty",
			message: newTestImmediateMessage(makerAssetID, takerAssetID, 1, 10),
			seed: func(s *testSetup) {
				// mirror asset pair at a distinct rate (maker split 2) so it is a
				// genuine counterparty rather than a same-ID duplicate; it demands
				// taker split 1, exactly filling the incoming 1-unit order.
				restOrder(t, s, takerAssetID, makerAssetID, 2, 1)
			},
			setup: func(s *testSetup) {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Times(3)
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			check: func(t *testing.T, s *testSetup) {
				assert.Nil(t, orderInStore(t, s, orderIDFor(makerAssetID, takerAssetID, 1)), "fully filled incoming order must be removed")
				assert.Nil(t, orderInStore(t, s, orderIDFor(takerAssetID, makerAssetID, 2)), "fully consumed counterparty must be removed")
				s.transferAuxiliaryKeeper.AssertNumberOfCalls(t, "Help", 3)
			},
		},
		{
			// Case B: the incoming order (leftover 1) is smaller than the resting
			// counterparty's demand (taker split 2), so the incoming order is fully
			// consumed and the counterparty is PARTIALLY filled. The buyer must
			// receive maker split × makerSplit/takerSplit = 1 × 4/2 = 2 units of the
			// counter-asset, reducing the resting maker split from 4 to 2. The
			// pre-rewrite formula omitted the 1e-18 scaling and paid the buyer 0,
			// leaving the resting maker split at 4 — this asserts the corrected math.
			name:    "partialFillPaysBuyerScaledAmount",
			message: newTestImmediateMessage(makerAssetID, takerAssetID, 1, 10),
			seed: func(s *testSetup) {
				restOrder(t, s, takerAssetID, makerAssetID, 4, 2)
			},
			setup: func(s *testSetup) {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Times(3)
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			check: func(t *testing.T, s *testSetup) {
				assert.Nil(t, orderInStore(t, s, orderIDFor(makerAssetID, takerAssetID, 1)), "fully consumed incoming order must be removed")
				counterparty := orderInStore(t, s, orderIDForRate(takerAssetID, makerAssetID, 4, 2))
				require.NotNil(t, counterparty, "partially filled counterparty must remain")
				assert.True(t, counterparty.GetMakerSplit().Equal(math.NewInt(2)), "counterparty maker split must drop from 4 to 2 (buyer paid 2, not 0); got %s", counterparty.GetMakerSplit())
			},
		},
		{
			// A resting order with the SAME asset pair is not a counterparty (it
			// wants the asset this order is selling). The incoming order must
			// ignore it and rest, doing only the escrow transfer.
			name:    "ignoresIncompatibleAssetPair",
			message: newTestImmediateMessage(makerAssetID, takerAssetID, 1, 10),
			seed: func(s *testSetup) {
				// same maker/taker asset direction as the incoming order → not a counterparty; distinct order ID via maker split 2
				restOrder(t, s, makerAssetID, takerAssetID, 2, 1)
			},
			setup: func(s *testSetup) {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			check: func(t *testing.T, s *testSetup) {
				require.NotNil(t, orderInStore(t, s, orderIDFor(makerAssetID, takerAssetID, 1)), "order must rest when the only resting order is not a counterparty")
				s.transferAuxiliaryKeeper.AssertNumberOfCalls(t, "Help", 1)
			},
		},
		{
			name:    "duplicateOrder",
			message: newTestImmediateMessage(makerAssetID, takerAssetID, 2, 10),
			seed: func(s *testSetup) {
				restOrder(t, s, makerAssetID, takerAssetID, 2, 1)
			},
			setup: func(s *testSetup) {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.EntityAlreadyExists,
		},
		{
			name:    "authenticationFailure",
			message: newTestImmediateMessage(makerAssetID, takerAssetID, 1, 10),
			setup: func(s *testSetup) {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
			},
			wantErr: errorConstants.MockError,
		},
		{
			name: "invalidMakerSplit",
			message: func() helpers.Message {
				message := newTestImmediateMessage(makerAssetID, takerAssetID, 1, 10)
				message.(*Message).MakerSplit = "invalidRate"
				return message
			}(),
			setup: func(s *testSetup) {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.InvalidParameter,
		},
		{
			name:    "expiryExceedsMaxOrderLife",
			message: newTestImmediateMessage(makerAssetID, takerAssetID, 3, 50000),
			setup: func(s *testSetup) {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.InvalidRequest,
		},
		{
			// A settlement transfer failure must surface as an error, not a panic.
			name:    "settlementTransferFailureReturnsError",
			message: newTestImmediateMessage(makerAssetID, takerAssetID, 1, 10),
			seed: func(s *testSetup) {
				restOrder(t, s, takerAssetID, makerAssetID, 2, 1)
			},
			setup: func(s *testSetup) {
				s.authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				// escrow transfer succeeds, first settlement transfer fails
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
				s.transferAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError).Once()
				s.conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil).Once()
			},
			wantErr: errorConstants.MockError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := setupTest(t)
			if tt.seed != nil {
				tt.seed(s)
			}
			tt.setup(s)

			got, err := s.TransactionKeeper.Transact(sdkTypes.WrapSDKContext(s.Context), tt.message)

			if tt.wantErr != nil {
				assert.True(t, tt.wantErr.Is(err), "Transact() error = %v, want %v", err, tt.wantErr)
				assert.Nil(t, got)
			} else {
				require.NoError(t, err)
				assert.Equal(t, newTransactionResponse(), got)
				if tt.check != nil {
					tt.check(t, s)
				}
			}
		})
	}

	t.Run("nilMessagePanics", func(t *testing.T) {
		s := setupTest(t)
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
