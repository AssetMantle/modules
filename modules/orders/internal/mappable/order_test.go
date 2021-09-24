/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_Order_Methods(t *testing.T) {

	classificationID := base.NewID("classificationID")
	makerOwnableID := base.NewID("makerOwnableID")
	takerOwnableID := base.NewID("takerOwnableID")
	makerID := base.NewID("makerID")
	rateID := base.NewID(sdkTypes.OneDec().String())
	creationID := base.NewID("100")

	takerIDImmutableProperty := base.NewProperty(base.NewID(properties.TakerID), base.NewFact(base.NewStringData("takerIDImmutableProperty")))
	exchangeRateImmutableProperty := base.NewMetaProperty(base.NewID(properties.ExchangeRate), base.NewMetaFact(base.NewDecData(sdkTypes.OneDec())))
	creationImmutableProperty := base.NewMetaProperty(base.NewID(properties.Creation), base.NewMetaFact(base.NewHeightData(base.NewHeight(100))))
	expiryImmutableProperty := base.NewProperty(base.NewID(properties.Expiry), base.NewFact(base.NewStringData("expiryImmutableProperty")))
	makerOwnableSplitImmutableProperty := base.NewProperty(base.NewID(properties.MakerOwnableSplit), base.NewFact(base.NewStringData("makerOwnableSplitImmutableProperty")))

	takerIDMutableProperty := base.NewProperty(base.NewID(properties.TakerID), base.NewFact(base.NewStringData("takerIDMutableProperty")))
	exchangeRateMutableProperty := base.NewProperty(base.NewID(properties.ExchangeRate), base.NewFact(base.NewDecData(sdkTypes.OneDec())))
	creationMutableProperty := base.NewProperty(base.NewID(properties.Creation), base.NewFact(base.NewHeightData(base.NewHeight(100))))
	expiryMutableProperty := base.NewProperty(base.NewID(properties.Expiry), base.NewFact(base.NewStringData("expiryMutableProperty")))
	makerOwnableSplitMutableProperty := base.NewProperty(base.NewID(properties.MakerOwnableSplit), base.NewFact(base.NewStringData("makerOwnableSplitMutableProperty")))

	immutableProperties := base.NewProperties(takerIDImmutableProperty, exchangeRateImmutableProperty.RemoveData(), creationImmutableProperty.RemoveData(), expiryImmutableProperty, makerOwnableSplitImmutableProperty)
	mutableProperties := base.NewProperties(takerIDMutableProperty, exchangeRateMutableProperty, creationMutableProperty, expiryMutableProperty, makerOwnableSplitMutableProperty)
	testOrderID := key.NewOrderID(classificationID, makerOwnableID, takerOwnableID, rateID, creationID, makerID, immutableProperties)
	testOrder := NewOrder(testOrderID, immutableProperties, base.NewProperties()).(order)
	testOrder2 := NewOrder(testOrderID, base.NewProperties(), mutableProperties).(order)
	testOrder3 := NewOrder(testOrderID, base.NewProperties(), base.NewProperties()).(order)

	data, _ := base.ReadIDData("")
	defaultTakerProperty := base.NewProperty(base.NewID(properties.TakerID), base.NewFact(data))
	defaultExchangeRateProperty := base.NewProperty(base.NewID(properties.ExchangeRate), base.NewFact(base.NewDecData(sdkTypes.OneDec())))
	data, _ = base.ReadHeightData("100")
	defaultCreationProperty := base.NewProperty(base.NewID(properties.Creation), base.NewFact(data))
	data, _ = base.ReadHeightData("-1")
	defaultExpiryProperty := base.NewProperty(base.NewID(properties.Expiry), base.NewFact(data))
	data, _ = base.ReadDecData("")
	defaultMakerOwnableSplitProperty := base.NewProperty(base.NewID(properties.MakerOwnableSplit), base.NewFact(data))

	require.Equal(t, order{ID: testOrderID, HasImmutables: baseTraits.HasImmutables{Properties: immutableProperties}, HasMutables: baseTraits.HasMutables{Properties: base.NewProperties()}}, testOrder)
	require.Equal(t, testOrderID, testOrder.GetID())
	require.Equal(t, testOrderID, testOrder.GetKey())
	require.Equal(t, classificationID, testOrder.GetClassificationID())
	require.Equal(t, makerOwnableID, testOrder.GetMakerOwnableID())
	require.Equal(t, takerOwnableID, testOrder.GetTakerOwnableID())
	require.Equal(t, makerID, testOrder.GetMakerID())

	//GetTakerID
	require.Equal(t, takerIDImmutableProperty, testOrder.GetTakerID())
	require.Equal(t, takerIDMutableProperty, testOrder2.GetTakerID())
	require.Equal(t, defaultTakerProperty, testOrder3.GetTakerID())
	//GetExchangeRate
	require.Equal(t, exchangeRateImmutableProperty, testOrder.GetExchangeRate())
	require.Equal(t, exchangeRateMutableProperty, testOrder2.GetExchangeRate().RemoveData())
	require.Equal(t, defaultExchangeRateProperty, testOrder3.GetExchangeRate().RemoveData())

	//GetCreation
	require.Equal(t, creationImmutableProperty, testOrder.GetCreation())
	require.Equal(t, creationMutableProperty, testOrder2.GetCreation().RemoveData())
	require.Equal(t, defaultCreationProperty, testOrder3.GetCreation().RemoveData())

	//GetExpiry
	require.Equal(t, expiryImmutableProperty, testOrder.GetExpiry())
	require.Equal(t, expiryMutableProperty, testOrder2.GetExpiry())
	require.Equal(t, defaultExpiryProperty, testOrder3.GetExpiry())

	//GetMakerOwnableSplit
	require.Equal(t, makerOwnableSplitImmutableProperty, testOrder.GetMakerOwnableSplit())
	require.Equal(t, makerOwnableSplitMutableProperty, testOrder2.GetMakerOwnableSplit())
	require.Equal(t, defaultMakerOwnableSplitProperty, testOrder3.GetMakerOwnableSplit())

	require.Equal(t, immutableProperties, testOrder.GetImmutableProperties())
	require.Equal(t, mutableProperties, testOrder2.GetMutableProperties())
	require.Equal(t, testOrderID, testOrder2.GetID())
	require.Equal(t, testOrderID, testOrder2.GetKey())

}
