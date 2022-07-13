// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
)

func readIdentityID(identityIDString string) ids.ID {
	idList := stringUtilities.SplitCompositeIDString(identityIDString)
	if len(idList) == 2 {
		return identityID{
			ClassificationID: baseIDs.NewStringID(idList[0]),
			Hash:             baseIDs.NewStringID(idList[1]),
		}
	}

	return identityID{ClassificationID: baseIDs.NewStringID(""), Hash: baseIDs.NewStringID("")}
}

func identityIDFromInterface(i interface{}) identityID {
	switch value := i.(type) {
	case identityID:
		return value
	case ids.ID:
		return identityIDFromInterface(readIdentityID(value.String()))
	default:
		panic(i)
	}
}

func FromID(id ids.ID) helpers.Key {
	return identityIDFromInterface(id)
}
