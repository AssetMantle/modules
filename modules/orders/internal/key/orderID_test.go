// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"strings"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

func Test_OrderID_Methods(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	makerOwnableID := baseIDs.NewID("makerOwnableID")
	takerOwnableID := baseIDs.NewID("takerOwnableID")
	makerID := baseIDs.NewID("makerID")
	defaultImmutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("ImmutableData")))
	rateID := baseIDs.NewID(sdkTypes.OneDec().String())
	creationID := baseIDs.NewID("100")
	immutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("ImmutableData")))

	testOrderID := NewOrderID(classificationID, makerOwnableID, takerOwnableID, rateID, creationID, makerID, immutableProperties).(orderID)
	testOrderID2 := NewOrderID(classificationID, makerOwnableID, takerOwnableID, baseIDs.NewID(sdkTypes.MustNewDecFromStr("2.3").String()), baseIDs.NewID("creation"), makerID, base.NewPropertyList()).(orderID)
	require.Equal(t, testOrderID, orderID{ClassificationID: classificationID, MakerOwnableID: makerOwnableID, TakerOwnableID: takerOwnableID, RateID: rateID, CreationID: creationID, MakerID: makerID, HashID: baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()})
	require.Equal(t, true, testOrderID.Equals(testOrderID))
	require.Equal(t, false, testOrderID.Compare(baseIDs.NewID("")) == 0)
	require.Equal(t, strings.Join([]string{classificationID.String(), makerOwnableID.String(), takerOwnableID.String(), rateID.String(), creationID.String(), makerID.String(), baseQualified.Immutables{PropertyList: defaultImmutableProperties}.GenerateHashID().String()}, "."), testOrderID.String())
	require.Equal(t, false, testOrderID.IsPartial())
	require.Equal(t, true, testOrderID.Equals(testOrderID))
	require.Equal(t, false, testOrderID.Equals(testOrderID2))
	require.Equal(t, false, testOrderID.Equals(FromID(baseIDs.NewID(""))))
	require.Equal(t, testOrderID, FromID(testOrderID))
	require.Equal(t, orderID{ClassificationID: baseIDs.NewID(""), MakerOwnableID: baseIDs.NewID(""), TakerOwnableID: baseIDs.NewID(""), RateID: baseIDs.NewID(""), CreationID: baseIDs.NewID(""), MakerID: baseIDs.NewID(""), HashID: baseIDs.NewID("")}, FromID(baseIDs.NewID("")))
	require.Equal(t, testOrderID, FromID(baseIDs.NewID(classificationID.String()+"."+makerOwnableID.String()+"."+takerOwnableID.String()+"."+rateID.String()+"."+creationID.String()+"."+makerID.String()+"."+baseQualified.Immutables{PropertyList: defaultImmutableProperties}.GenerateHashID().String())))
	require.Equal(t, classificationID, ReadClassificationID(testOrderID))
	require.Equal(t, makerOwnableID, ReadMakerOwnableID(testOrderID))
	require.Equal(t, takerOwnableID, ReadTakerOwnableID(testOrderID))
	require.Equal(t, rateID, ReadRateID(testOrderID))
	require.Equal(t, creationID, ReadCreationID(testOrderID))
	require.Equal(t, makerID, ReadMakerID(testOrderID))
	require.Equal(t, true, FromID(baseIDs.NewID("")).IsPartial())

}
