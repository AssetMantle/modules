/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"strconv"
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func readOrderID(orderIDString string) types.ID {
	idList := strings.Split(orderIDString, constants.SecondOrderCompositeIDSeparator)

	if len(idList) == 7 {
		exchangeRateID := base.NewID("")
		exchangeRate, Error := sdkTypes.NewDecFromStr(idList[3])

		if Error == nil && !exchangeRate.Equal(sdkTypes.ZeroDec()) {
			exchangeRateID = base.NewDecID(exchangeRate)
		}

		heightID := base.NewID("")
		height, Error := strconv.ParseInt(idList[4], 10, 64)

		if Error == nil && height != 0 {
			heightID = base.NewHeightID(height)
		}

		return orderID{
			ClassificationID: base.NewID(idList[0]),
			MakerOwnableID:   base.NewID(idList[1]),
			TakerOwnableID:   base.NewID(idList[2]),
			ExchangeRate:     exchangeRateID,
			CreationHeight:   heightID,
			MakerID:          base.NewID(idList[5]),
			HashID:           base.NewID(idList[6]),
		}
	}

	return orderID{ClassificationID: base.NewID(""), MakerOwnableID: base.NewID(""), TakerOwnableID: base.NewID(""), ExchangeRate: base.NewID(""), CreationHeight: base.NewID(""), MakerID: base.NewID(""), HashID: base.NewID("")}
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

func ReadExchangeRate(orderID types.ID) types.ID {
	return orderIDFromInterface(orderID).ExchangeRate
}

func ReadCreationHeight(orderID types.ID) types.ID {
	return orderIDFromInterface(orderID).CreationHeight
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
