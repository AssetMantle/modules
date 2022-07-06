// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
)

func readSplitID(splitIDString string) ids.ID {
	idList := stringUtilities.SplitCompositeIDString(splitIDString)
	if len(idList) == 2 {
		return splitID{
			OwnerID:   baseIDs.NewID(idList[0]),
			OwnableID: baseIDs.NewID(idList[1]),
		}
	}

	return splitID{OwnerID: baseIDs.NewID(""), OwnableID: baseIDs.NewID("")}
}

func splitIDFromInterface(i interface{}) splitID {
	switch value := i.(type) {
	case splitID:
		return value
	case ids.ID:
		return splitIDFromInterface(readSplitID(value.String()))
	default:
		panic(i)
	}
}

func ReadOwnableID(id ids.ID) ids.ID {
	return splitIDFromInterface(id).OwnableID
}

func ReadOwnerID(id ids.ID) ids.ID {
	return splitIDFromInterface(id).OwnerID
}

func FromID(id ids.ID) helpers.Key {
	return splitIDFromInterface(id)
}

func ToID(key helpers.Key) ids.ID {
	return splitIDFromInterface(key)
}
