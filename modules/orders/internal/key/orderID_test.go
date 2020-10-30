package key

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func Test_OrderID_Methods(t *testing.T) {

	classificationID := base.NewID("classificationID")
	makerOwnableID := base.NewID("makerOwnableID")
	takerOwnableID := base.NewID("takerOwnableID")
	makerID := base.NewID("makerID")
	immutables := base.NewImmutables(base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("ImmutableData")))))

	testOrderID := NewOrderID(classificationID, makerOwnableID, takerOwnableID, makerID, immutables).(orderID)
	testOrderID2 := NewOrderID(classificationID, makerOwnableID, takerOwnableID, makerID, base.NewImmutables(base.NewProperties())).(orderID)
	require.Equal(t, testOrderID, orderID{ClassificationID: classificationID, MakerOwnableID: makerOwnableID, TakerOwnableID: takerOwnableID, MakerID: makerID, HashID: immutables.GetHashID()})
	require.Equal(t, true, testOrderID.Equals(testOrderID))
	require.Equal(t, false, testOrderID.Equals(base.NewID("")))
	require.Equal(t, strings.Join([]string{classificationID.String(), makerOwnableID.String(), takerOwnableID.String(), makerID.String(), immutables.GetHashID().String()}, constants.SecondOrderCompositeIDSeparator), testOrderID.String())
	require.Equal(t, false, testOrderID.IsPartial())
	require.Equal(t, true, testOrderID.Matches(testOrderID))
	require.Equal(t, false, testOrderID.Matches(testOrderID2))
	require.Equal(t, testOrderID, New(testOrderID))
	require.Equal(t, orderID{ClassificationID: base.NewID(""), MakerOwnableID: base.NewID(""), TakerOwnableID: base.NewID(""), MakerID: base.NewID(""), HashID: base.NewID("")}, New(base.NewID("")))
	require.Equal(t, testOrderID, New(base.NewID(classificationID.String()+constants.SecondOrderCompositeIDSeparator+makerOwnableID.String()+constants.SecondOrderCompositeIDSeparator+takerOwnableID.String()+constants.SecondOrderCompositeIDSeparator+makerID.String()+constants.SecondOrderCompositeIDSeparator+immutables.GetHashID().String())))
	require.Equal(t, classificationID, ReadClassificationID(testOrderID))
	require.Equal(t, makerOwnableID, ReadMakerOwnableID(testOrderID))
	require.Equal(t, takerOwnableID, ReadTakerOwnableID(testOrderID))
	require.Equal(t, makerID, ReadMakerID(testOrderID))

}
