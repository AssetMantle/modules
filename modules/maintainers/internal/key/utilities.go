// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"strings"

	"github.com/AssetMantle/modules/schema/helpers"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
)

func readMaintainerID(maintainerIDString string) types.ID {
	idList := strings.Split(maintainerIDString, constants.SecondOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return maintainerID{
			ClassificationID: base.NewID(idList[0]),
			IdentityID:       base.NewID(idList[1]),
		}
	}

	return maintainerID{IdentityID: base.NewID(""), ClassificationID: base.NewID("")}
}
func maintainerIDFromInterface(i interface{}) maintainerID {
	switch value := i.(type) {
	case maintainerID:
		return value
	case types.ID:
		return maintainerIDFromInterface(readMaintainerID(value.String()))
	default:
		panic(i)
	}
}

func ReadClassificationID(assetID types.ID) types.ID {
	return maintainerIDFromInterface(assetID).ClassificationID
}

func ReadIdentityID(assetID types.ID) types.ID {
	return maintainerIDFromInterface(assetID).IdentityID
}

func FromID(id types.ID) helpers.Key {
	return maintainerIDFromInterface(id)
}
