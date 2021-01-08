/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
)

func readMetaID(metaIDString string) types.ID {
	idList := strings.Split(metaIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return metaID{
			TypeID: base.NewID(idList[0]),
			HashID: base.NewID(idList[1]),
		}
	}
	return metaID{TypeID: base.NewID(""), HashID: base.NewID("")}
}
func metaIDFromInterface(id types.ID) metaID {
	switch value := id.(type) {
	case metaID:
		return value
	default:
		return metaIDFromInterface(readMetaID(id.String()))
	}
}
