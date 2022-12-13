/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
)

func readSplitID(splitIDString string) SplitID {
	idList := strings.Split(splitIDString, constants.SecondOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return SplitID{
			OwnerID:   base.NewID(idList[0]),
			OwnableID: base.NewID(idList[1]),
		}
	}

	return SplitID{OwnerID: base.NewID(""), OwnableID: base.NewID("")}
}

func splitIDFromInterface(i interface{}) SplitID {
	switch value := i.(type) {
	case SplitID:
		return value
	case types.ID:
		return splitIDFromInterface(readSplitID(value.String()))
	default:
		panic(i)
	}
}

func ReadOwnableID(id types.ID) types.ID {
	return splitIDFromInterface(id).OwnableID
}

func ReadOwnerID(id types.ID) types.ID {
	return splitIDFromInterface(id).OwnerID
}

func FromID(id types.ID) helpers.Key {
	return splitIDFromInterface(id)
}

func ToID(key helpers.Key) types.ID {
	id := splitIDFromInterface(key)
	return &id
}
