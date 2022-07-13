// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"strconv"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
)

func readOrderID(orderIDString string) ids.ID {
	idList := stringUtilities.SplitCompositeIDString(orderIDString)

	if len(idList) == 7 {
		exchangeRate, err := sdkTypes.NewDecFromStr(idList[3])
		if err != nil {
			return orderID{ClassificationID: baseIDs.NewStringID(""), MakerOwnableID: baseIDs.NewStringID(""), TakerOwnableID: baseIDs.NewStringID(""), RateID: baseIDs.NewStringID(""), CreationID: baseIDs.NewStringID(""), MakerID: baseIDs.NewStringID(""), Hash: baseIDs.NewStringID("")}
		}

		height, err := strconv.ParseInt(idList[4], 10, 64)
		if err != nil {
			return orderID{ClassificationID: baseIDs.NewStringID(""), MakerOwnableID: baseIDs.NewStringID(""), TakerOwnableID: baseIDs.NewStringID(""), RateID: baseIDs.NewStringID(""), CreationID: baseIDs.NewStringID(""), MakerID: baseIDs.NewStringID(""), Hash: baseIDs.NewStringID("")}
		}

		return orderID{
			ClassificationID: baseIDs.NewStringID(idList[0]),
			MakerOwnableID:   baseIDs.NewStringID(idList[1]),
			TakerOwnableID:   baseIDs.NewStringID(idList[2]),
			RateID:           baseIDs.NewStringID(exchangeRate.String()),
			CreationID:       baseIDs.NewStringID(strconv.FormatInt(height, 10)),
			MakerID:          baseIDs.NewStringID(idList[5]),
			Hash:             baseIDs.NewStringID(idList[6]),
		}
	}

	return orderID{ClassificationID: baseIDs.NewStringID(""), MakerOwnableID: baseIDs.NewStringID(""), TakerOwnableID: baseIDs.NewStringID(""), RateID: baseIDs.NewStringID(""), CreationID: baseIDs.NewStringID(""), MakerID: baseIDs.NewStringID(""), Hash: baseIDs.NewStringID("")}
}
func orderIDFromInterface(i interface{}) orderID {
	switch value := i.(type) {
	case orderID:
		return value
	case ids.ID:
		return orderIDFromInterface(readOrderID(value.String()))
	default:
		panic(i)
	}
}

func ReadClassificationID(orderID ids.ID) ids.ID {
	return orderIDFromInterface(orderID).ClassificationID
}

func ReadRateID(orderID ids.ID) ids.ID {
	return orderIDFromInterface(orderID).RateID
}

func ReadCreationID(orderID ids.ID) ids.ID {
	return orderIDFromInterface(orderID).CreationID
}

func ReadMakerOwnableID(orderID ids.ID) ids.ID {
	return orderIDFromInterface(orderID).MakerOwnableID
}

func ReadTakerOwnableID(orderID ids.ID) ids.ID {
	return orderIDFromInterface(orderID).TakerOwnableID
}

func ReadMakerID(orderID ids.ID) ids.ID {
	return orderIDFromInterface(orderID).MakerID
}

func FromID(id ids.ID) helpers.Key {
	return orderIDFromInterface(id)
}
