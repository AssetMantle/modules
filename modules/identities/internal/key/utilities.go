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

func identityIDFromInterface(id types.ID) identityID {
	switch value := id.(type) {
	case identityID:
		return value
	default:
		return identityIDFromInterface(readIdentityID(id.String()))
	}
}
