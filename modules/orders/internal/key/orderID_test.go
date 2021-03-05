/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"strings"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_OrderID_Methods(t *testing.T) {

	classificationID := base.NewID("classificationID")
	makerOwnableID := base.NewID("makerOwnableID")
	takerOwnableID := base.NewID("takerOwnableID")
	makerID := base.NewID("makerID")
	defaultImmutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("ImmutableData"))))
	rateID := base.NewID(sdkTypes.OneDec().String())
	creationID := base.NewID("100")
	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("ImmutableData"))))

	testOrderID := NewOrderID(classificationID, makerOwnableID, takerOwnableID, rateID, creationID, makerID, immutableProperties).(orderID)
	testOrderID2 := NewOrderID(classificationID, makerOwnableID, takerOwnableID, base.NewID(sdkTypes.MustNewDecFromStr("2.3").String()), base.NewID("creation"), makerID, base.NewProperties()).(orderID)
	require.Equal(t, testOrderID, orderID{ClassificationID: classificationID, MakerOwnableID: makerOwnableID, TakerOwnableID: takerOwnableID, RateID: rateID, CreationID: creationID, MakerID: makerID, HashID: baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID()})
	require.Equal(t, true, testOrderID.Equals(testOrderID))
	require.Equal(t, false, testOrderID.Equals(base.NewID("")))
	require.Equal(t, strings.Join([]string{classificationID.String(), makerOwnableID.String(), takerOwnableID.String(), rateID.String(), creationID.String(), makerID.String(), baseTraits.HasImmutables{Properties: defaultImmutableProperties}.GenerateHashID().String()}, constants.SecondOrderCompositeIDSeparator), testOrderID.String())
	require.Equal(t, false, testOrderID.IsPartial())
	require.Equal(t, true, testOrderID.Matches(testOrderID))
	require.Equal(t, false, testOrderID.Matches(testOrderID2))
	require.Equal(t, false, testOrderID.Matches(FromID(base.NewID(""))))
	require.Equal(t, testOrderID, FromID(testOrderID))
	require.Equal(t, orderID{ClassificationID: base.NewID(""), MakerOwnableID: base.NewID(""), TakerOwnableID: base.NewID(""), RateID: base.NewID(""), CreationID: base.NewID(""), MakerID: base.NewID(""), HashID: base.NewID("")}, FromID(base.NewID("")))
	require.Equal(t, testOrderID, FromID(base.NewID(classificationID.String()+constants.SecondOrderCompositeIDSeparator+makerOwnableID.String()+constants.SecondOrderCompositeIDSeparator+takerOwnableID.String()+constants.SecondOrderCompositeIDSeparator+rateID.String()+constants.SecondOrderCompositeIDSeparator+creationID.String()+constants.SecondOrderCompositeIDSeparator+makerID.String()+constants.SecondOrderCompositeIDSeparator+baseTraits.HasImmutables{Properties: defaultImmutableProperties}.GenerateHashID().String())))
	require.Equal(t, classificationID, ReadClassificationID(testOrderID))
	require.Equal(t, makerOwnableID, ReadMakerOwnableID(testOrderID))
	require.Equal(t, takerOwnableID, ReadTakerOwnableID(testOrderID))
	require.Equal(t, rateID, ReadRateID(testOrderID))
	require.Equal(t, creationID, ReadCreationID(testOrderID))
	require.Equal(t, makerID, ReadMakerID(testOrderID))
	require.Equal(t, true, FromID(base.NewID("")).IsPartial())

}
