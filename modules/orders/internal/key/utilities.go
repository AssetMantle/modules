/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"strings"

	"github.com/persistenceOne/persistenceSDK/schema/helpers"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func readOrderID(orderIDString string) types.ID {
	idList := strings.Split(orderIDString, constants.SecondOrderCompositeIDSeparator)
	if len(idList) == 5 {
		return orderID{
			ClassificationID: base.NewID(idList[0]),
			MakerOwnableID:   base.NewID(idList[1]),
			TakerOwnableID:   base.NewID(idList[2]),
			MakerID:          base.NewID(idList[3]),
			HashID:           base.NewID(idList[4]),
		}
	}

	return orderID{ClassificationID: base.NewID(""), MakerOwnableID: base.NewID(""), TakerOwnableID: base.NewID(""), MakerID: base.NewID(""), HashID: base.NewID("")}
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
