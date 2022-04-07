// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"strconv"
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types"
)

func readOrderID(orderIDString string) types.ID {
	idList := strings.Split(orderIDString, constants.SecondOrderCompositeIDSeparator)

	if len(idList) == 7 {
		exchangeRate, err := sdkTypes.NewDecFromStr(idList[3])
		if err != nil {
			return orderID{ClassificationID: baseIDs.NewID(""), MakerOwnableID: baseIDs.NewID(""), TakerOwnableID: baseIDs.NewID(""), RateID: baseIDs.NewID(""), CreationID: baseIDs.NewID(""), MakerID: baseIDs.NewID(""), HashID: baseIDs.NewID("")}
		}

		height, err := strconv.ParseInt(idList[4], 10, 64)
		if err != nil {
			return orderID{ClassificationID: baseIDs.NewID(""), MakerOwnableID: baseIDs.NewID(""), TakerOwnableID: baseIDs.NewID(""), RateID: baseIDs.NewID(""), CreationID: baseIDs.NewID(""), MakerID: baseIDs.NewID(""), HashID: baseIDs.NewID("")}
		}

		return orderID{
			ClassificationID: baseIDs.NewID(idList[0]),
			MakerOwnableID:   baseIDs.NewID(idList[1]),
			TakerOwnableID:   baseIDs.NewID(idList[2]),
			RateID:           baseIDs.NewID(exchangeRate.String()),
			CreationID:       baseIDs.NewID(strconv.FormatInt(height, 10)),
			MakerID:          baseIDs.NewID(idList[5]),
			HashID:           baseIDs.NewID(idList[6]),
		}
	}

	return orderID{ClassificationID: baseIDs.NewID(""), MakerOwnableID: baseIDs.NewID(""), TakerOwnableID: baseIDs.NewID(""), RateID: baseIDs.NewID(""), CreationID: baseIDs.NewID(""), MakerID: baseIDs.NewID(""), HashID: baseIDs.NewID("")}
}
func orderIDFromInterface(i interface{}) orderID {
	switch value := i.(type) {
	case orderID:
		return value
	case types.ID:
		return orderIDFromInterface(readOrderID(value.String()))
	default:
		panic(i)
	}
}

func ReadClassificationID(orderID types.ID) types.ID {
	return orderIDFromInterface(orderID).ClassificationID
}

func ReadRateID(orderID types.ID) types.ID {
	return orderIDFromInterface(orderID).RateID
}

func ReadCreationID(orderID types.ID) types.ID {
	return orderIDFromInterface(orderID).CreationID
}

func ReadMakerOwnableID(orderID types.ID) types.ID {
	return orderIDFromInterface(orderID).MakerOwnableID
}

func ReadTakerOwnableID(orderID types.ID) types.ID {
	return orderIDFromInterface(orderID).TakerOwnableID
}

func ReadMakerID(orderID types.ID) types.ID {
	return orderIDFromInterface(orderID).MakerID
}

func FromID(id types.ID) helpers.Key {
	return orderIDFromInterface(id)
}
