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

func readIdentityID(identityIDString string) types.ID {
	idList := strings.Split(identityIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return identityID{
			ClassificationID: base.NewID(idList[0]),
			HashID:           base.NewID(idList[1]),
		}
	}

	return identityID{ClassificationID: base.NewID(""), HashID: base.NewID("")}
}

func identityIDFromInterface(i interface{}) identityID {
	switch value := i.(type) {
	case identityID:
		return value
	case types.ID:
		return identityIDFromInterface(readIdentityID(value.String()))
	default:
		panic(i)
	}
}

func FromID(id types.ID) helpers.Key {
	return identityIDFromInterface(id)
}

func ReadClassificationID(identityID types.ID) types.ID {
	return identityIDFromInterface(identityID).ClassificationID
}
