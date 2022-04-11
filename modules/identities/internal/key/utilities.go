// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"strings"

	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/schema/types"
)

func readIdentityID(identityIDString string) types.ID {
	idList := strings.Split(identityIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return identityID{
			ClassificationID: baseIDs.NewID(idList[0]),
			HashID:           baseIDs.NewID(idList[1]),
		}
	}

	return identityID{ClassificationID: baseIDs.NewID(""), HashID: baseIDs.NewID("")}
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
