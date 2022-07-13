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
	classificationID := baseIDs.NewStringID("classificationID")
	makerOwnableID := baseIDs.NewStringID("makerOwnableID")
	takerOwnableID := baseIDs.NewStringID("takerOwnableID")
	makerID := baseIDs.NewStringID("makerID")
	defaultImmutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	rateID := baseIDs.NewStringID(sdkTypes.OneDec().String())
	creationID := baseIDs.NewStringID("100")
	immutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))

	testOrderID := NewOrderID(classificationID, makerOwnableID, takerOwnableID, rateID, creationID, makerID, immutableProperties).(orderID)
	testOrderID2 := NewOrderID(classificationID, makerOwnableID, takerOwnableID, baseIDs.NewStringID(sdkTypes.MustNewDecFromStr("2.3").String()), baseIDs.NewStringID("creation"), makerID, base.NewPropertyList()).(orderID)
	require.Equal(t, testOrderID, orderID{ClassificationID: classificationID, MakerOwnableID: makerOwnableID, TakerOwnableID: takerOwnableID, RateID: rateID, CreationID: creationID, MakerID: makerID, Hash: baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID()})
	require.Equal(t, true, testOrderID.Equals(testOrderID))
	require.Equal(t, false, testOrderID.Compare(baseIDs.NewStringID("")) == 0)
	require.Equal(t, strings.Join([]string{classificationID.String(), makerOwnableID.String(), takerOwnableID.String(), rateID.String(), creationID.String(), makerID.String(), baseQualified.Immutables{PropertyList: defaultImmutableProperties}.GenerateHashID().String()}, "."), testOrderID.String())
	require.Equal(t, false, testOrderID.IsPartial())
	require.Equal(t, true, testOrderID.Equals(testOrderID))
	require.Equal(t, false, testOrderID.Equals(testOrderID2))
	require.Equal(t, false, testOrderID.Equals(FromID(baseIDs.NewStringID(""))))
	require.Equal(t, testOrderID, FromID(testOrderID))
	require.Equal(t, orderID{ClassificationID: baseIDs.NewStringID(""), MakerOwnableID: baseIDs.NewStringID(""), TakerOwnableID: baseIDs.NewStringID(""), RateID: baseIDs.NewStringID(""), CreationID: baseIDs.NewStringID(""), MakerID: baseIDs.NewStringID(""), Hash: baseIDs.NewStringID("")}, FromID(baseIDs.NewStringID("")))
	require.Equal(t, testOrderID, FromID(baseIDs.NewStringID(classificationID.String()+"."+makerOwnableID.String()+"."+takerOwnableID.String()+"."+rateID.String()+"."+creationID.String()+"."+makerID.String()+"."+baseQualified.Immutables{PropertyList: defaultImmutableProperties}.GenerateHashID().String())))
	require.Equal(t, classificationID, ReadClassificationID(testOrderID))
	require.Equal(t, makerOwnableID, ReadMakerOwnableID(testOrderID))
	require.Equal(t, takerOwnableID, ReadTakerOwnableID(testOrderID))
	require.Equal(t, rateID, ReadRateID(testOrderID))
	require.Equal(t, creationID, ReadCreationID(testOrderID))
	require.Equal(t, makerID, ReadMakerID(testOrderID))
	require.Equal(t, true, FromID(baseIDs.NewStringID("")).IsPartial())

}
