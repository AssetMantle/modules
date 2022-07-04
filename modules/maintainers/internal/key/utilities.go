// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

func readMaintainerID(maintainerIDString string) ids.ID {
	idList := stringUtilities.SplitCompositeIDString(maintainerIDString)
	if len(idList) == 2 {
		return maintainerID{
			ClassificationID: baseIDs.NewID(idList[0]),
			IdentityID:       baseIDs.NewID(idList[1]),
		}
	}

	return maintainerID{IdentityID: baseIDs.NewID(""), ClassificationID: baseIDs.NewID("")}
}
func maintainerIDFromInterface(i interface{}) maintainerID {
	switch value := i.(type) {
	case maintainerID:
		return value
	case ids.ID:
		return maintainerIDFromInterface(readMaintainerID(value.String()))
	default:
		panic(i)
	}
}

func ReadClassificationID(assetID ids.ID) ids.ID {
	return maintainerIDFromInterface(assetID).ClassificationID
}

func ReadIdentityID(assetID ids.ID) ids.ID {
	return maintainerIDFromInterface(assetID).IdentityID
}

func FromID(id ids.ID) helpers.Key {
	return maintainerIDFromInterface(id)
}
