/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"testing"

	"github.com/persistenceOne/persistenceSDK/constants/ids"
	"github.com/persistenceOne/persistenceSDK/schema/mappables/qualified"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/qualified"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func Test_Order_Methods(t *testing.T) {

	classificationID := base.NewID("classificationID")
	makerOwnableID := base.NewID("makerOwnableID")
	takerOwnableID := base.NewID("takerOwnableID")
	makerID := base.NewID("makerID")
	rateID := base.NewID(sdkTypes.OneDec().String())
	creationID := base.NewID("100")

	takerIDImmutableProperty := base.NewProperty(ids.TakerIDProperty, base.NewStringData("takerIDImmutableProperty"))
	exchangeRateImmutableProperty := base.NewMetaProperty(ids.ExchangeRateProperty, base.NewDecData(sdkTypes.OneDec()))
	creationImmutableProperty := base.NewMetaProperty(ids.CreationProperty, base.NewHeightData(base.NewHeight(100)))
	expiryImmutableProperty := base.NewProperty(ids.ExpiryProperty, base.NewStringData("expiryImmutableProperty"))
	makerOwnableSplitImmutableProperty := base.NewProperty(ids.MakerOwnableSplitProperty, base.NewStringData("makerOwnableSplitImmutableProperty"))

	takerIDMutableProperty := base.NewProperty(ids.TakerIDProperty, base.NewStringData("takerIDMutableProperty"))
	exchangeRateMutableProperty := base.NewProperty(ids.ExchangeRateProperty, base.NewDecData(sdkTypes.OneDec()))
	creationMutableProperty := base.NewProperty(ids.CreationProperty, base.NewHeightData(base.NewHeight(100)))
	expiryMutableProperty := base.NewProperty(ids.ExpiryProperty, base.NewStringData("expiryMutableProperty"))
	makerOwnableSplitMutableProperty := base.NewProperty(ids.MakerOwnableSplitProperty, base.NewStringData("makerOwnableSplitMutableProperty"))

	immutableProperties := base.NewProperties(takerIDImmutableProperty, exchangeRateImmutableProperty.RemoveData(), creationImmutableProperty.RemoveData(), expiryImmutableProperty, makerOwnableSplitImmutableProperty)
	mutableProperties := base.NewProperties(takerIDMutableProperty, exchangeRateMutableProperty, creationMutableProperty, expiryMutableProperty, makerOwnableSplitMutableProperty)
	testOrderID := key.NewOrderID(classificationID, makerOwnableID, takerOwnableID, rateID, creationID, makerID, immutableProperties)
	testOrder := NewOrder(testOrderID, immutableProperties, base.NewProperties()).(order)
	testOrder2 := NewOrder(testOrderID, base.NewProperties(), mutableProperties).(order)
	testOrder3 := NewOrder(testOrderID, base.NewProperties(), base.NewProperties()).(order)

	data, _ := base.ReadIDData("")
	defaultTakerProperty := base.NewProperty(ids.TakerIDProperty, data)
	defaultExchangeRateProperty := base.NewProperty(ids.ExchangeRateProperty, base.NewDecData(sdkTypes.OneDec()))
	data, _ = base.ReadHeightData("100")
	defaultCreationProperty := base.NewProperty(ids.CreationProperty, data)
	data, _ = base.ReadHeightData("-1")
	defaultExpiryProperty := base.NewProperty(ids.ExpiryProperty, data)
	data, _ = base.ReadDecData("")
	defaultMakerOwnableSplitProperty := base.NewProperty(ids.MakerOwnableSplitProperty, data)

	require.Equal(t, order{Document: qualified.Document{ID: testOrderID, HasImmutables: baseTraits.HasImmutables{Properties: immutableProperties}, HasMutables: baseTraits.HasMutables{Properties: base.NewProperties()}}}, testOrder)
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

	require.Equal(t, immutableProperties, testOrder.GetImmutableProperties())
	require.Equal(t, mutableProperties, testOrder2.GetMutableProperties())
	require.Equal(t, testOrderID, testOrder2.GetID())
	require.Equal(t, testOrderID, testOrder2.GetKey())

}
