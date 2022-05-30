// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	base2 "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func Test_Order_Methods(t *testing.T) {

	classificationID := baseIDs.NewID("classificationID")
	makerOwnableID := baseIDs.NewID("makerOwnableID")
	takerOwnableID := baseIDs.NewID("takerOwnableID")
	makerID := baseIDs.NewID("makerID")
	rateID := baseIDs.NewID(sdkTypes.OneDec().String())
	creationID := baseIDs.NewID("100")

	takerIDImmutableProperty := base2.NewProperty(ids.TakerIDProperty, baseData.NewStringData("takerIDImmutableProperty"))
	exchangeRateImmutableProperty := base2.NewMetaProperty(ids.ExchangeRateProperty, baseData.NewDecData(sdkTypes.OneDec()))
	creationImmutableProperty := base2.NewMetaProperty(ids.CreationProperty, baseData.NewHeightData(baseTypes.NewHeight(100)))
	expiryImmutableProperty := base2.NewProperty(ids.ExpiryProperty, baseData.NewStringData("expiryImmutableProperty"))
	makerOwnableSplitImmutableProperty := base2.NewProperty(ids.MakerOwnableSplitProperty, baseData.NewStringData("makerOwnableSplitImmutableProperty"))

	takerIDMutableProperty := base2.NewProperty(ids.TakerIDProperty, baseData.NewStringData("takerIDMutableProperty"))
	exchangeRateMutableProperty := base2.NewProperty(ids.ExchangeRateProperty, baseData.NewDecData(sdkTypes.OneDec()))
	creationMutableProperty := base2.NewProperty(ids.CreationProperty, baseData.NewHeightData(baseTypes.NewHeight(100)))
	expiryMutableProperty := base2.NewProperty(ids.ExpiryProperty, baseData.NewStringData("expiryMutableProperty"))
	makerOwnableSplitMutableProperty := base2.NewProperty(ids.MakerOwnableSplitProperty, baseData.NewStringData("makerOwnableSplitMutableProperty"))

	immutableProperties := base.NewPropertyList(takerIDImmutableProperty, exchangeRateImmutableProperty.RemoveData(), creationImmutableProperty.RemoveData(), expiryImmutableProperty, makerOwnableSplitImmutableProperty)
	mutableProperties := base.NewPropertyList(takerIDMutableProperty, exchangeRateMutableProperty, creationMutableProperty, expiryMutableProperty, makerOwnableSplitMutableProperty)
	testOrderID := key.NewOrderID(classificationID, makerOwnableID, takerOwnableID, rateID, creationID, makerID, immutableProperties)
	testOrder := NewOrder(testOrderID, immutableProperties, base.NewPropertyList()).(order)
	testOrder2 := NewOrder(testOrderID, base.NewPropertyList(), mutableProperties).(order)
	testOrder3 := NewOrder(testOrderID, base.NewPropertyList(), base.NewPropertyList()).(order)

	data, _ := baseData.ReadIDData("")
	defaultTakerProperty := base2.NewProperty(ids.TakerIDProperty, data)
	defaultExchangeRateProperty := base2.NewProperty(ids.ExchangeRateProperty, baseData.NewDecData(sdkTypes.OneDec()))
	data, _ = baseData.ReadHeightData("100")
	defaultCreationProperty := base2.NewProperty(ids.CreationProperty, data)
	data, _ = baseData.ReadHeightData("-1")
	defaultExpiryProperty := base2.NewProperty(ids.ExpiryProperty, data)
	data, _ = baseData.ReadDecData("")
	defaultMakerOwnableSplitProperty := base2.NewProperty(ids.MakerOwnableSplitProperty, data)

	require.Equal(t, order{Document: baseQualified.Document{ID: testOrderID, Immutables: baseQualified.Immutables{PropertyList: immutableProperties}, Mutables: baseQualified.Mutables{Properties: base.NewPropertyList()}}}, testOrder)
	require.Equal(t, testOrderID, testOrder.GetID())
	require.Equal(t, testOrderID, testOrder.GetKey())
	require.Equal(t, classificationID, testOrder.GetClassificationID())
	require.Equal(t, makerOwnableID, testOrder.GetMakerOwnableID())
	require.Equal(t, takerOwnableID, testOrder.GetTakerOwnableID())
	require.Equal(t, makerID, testOrder.GetMakerID())

	// GetTakerID
	require.Equal(t, takerIDImmutableProperty, testOrder.GetTakerID())
	require.Equal(t, takerIDMutableProperty, testOrder2.GetTakerID())
	require.Equal(t, defaultTakerProperty, testOrder3.GetTakerID())
	// GetExchangeRate
	require.Equal(t, exchangeRateImmutableProperty, testOrder.GetExchangeRate())
	require.Equal(t, exchangeRateMutableProperty, testOrder2.GetExchangeRate().RemoveData())
	require.Equal(t, defaultExchangeRateProperty, testOrder3.GetExchangeRate().RemoveData())

	// GetCreation
	require.Equal(t, creationImmutableProperty, testOrder.GetCreation())
	require.Equal(t, creationMutableProperty, testOrder2.GetCreation().RemoveData())
	require.Equal(t, defaultCreationProperty, testOrder3.GetCreation().RemoveData())

	// GetExpiry
	require.Equal(t, expiryImmutableProperty, testOrder.GetExpiry())
	require.Equal(t, expiryMutableProperty, testOrder2.GetExpiry())
	require.Equal(t, defaultExpiryProperty, testOrder3.GetExpiry())

	// GetMakerOwnableSplit
	require.Equal(t, makerOwnableSplitImmutableProperty, testOrder.GetMakerOwnableSplit())
	require.Equal(t, makerOwnableSplitMutableProperty, testOrder2.GetMakerOwnableSplit())
	require.Equal(t, defaultMakerOwnableSplitProperty, testOrder3.GetMakerOwnableSplit())

	require.Equal(t, immutableProperties, testOrder.GetImmutablePropertyList())
	require.Equal(t, mutableProperties, testOrder2.GetMutablePropertyList())
	require.Equal(t, testOrderID, testOrder2.GetID())
	require.Equal(t, testOrderID, testOrder2.GetKey())

}
