/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
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

func splitIDFromInterface(id types.ID) splitID {
	switch value := id.(type) {
	case splitID:
		return value
	default:
		return splitIDFromInterface(readSplitID(id.String()))
	}
}

func ReadOwnableID(assetID types.ID) types.ID {
	return splitIDFromInterface(assetID).OwnableID
}

func ReadOwnerID(assetID types.ID) types.ID {
	return splitIDFromInterface(assetID).OwnerID
}
