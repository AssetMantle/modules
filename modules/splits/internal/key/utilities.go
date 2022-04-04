// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"strings"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
)

func readSplitID(splitIDString string) types.ID {
	idList := strings.Split(splitIDString, constants.SecondOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return splitID{
			OwnerID:   base.NewID(idList[0]),
			OwnableID: base.NewID(idList[1]),
		}
	}

	return splitID{OwnerID: base.NewID(""), OwnableID: base.NewID("")}
}

func splitIDFromInterface(i interface{}) splitID {
	switch value := i.(type) {
	case splitID:
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
	return splitIDFromInterface(key)
}
