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

func readIdentityID(identityIDString string) IdentityID {
	idList := strings.Split(identityIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return IdentityID{
			ClassificationID: base.NewID(idList[0]),
			HashID:           base.NewID(idList[1]),
		}
	}

	return IdentityID{ClassificationID: base.NewID(""), HashID: base.NewID("")}
}

func identityIDFromInterface(i interface{}) IdentityID {
	switch value := i.(type) {
	case IdentityID:
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
	return base.NewID(identityIDFromInterface(identityID).ClassificationID.String())
}
