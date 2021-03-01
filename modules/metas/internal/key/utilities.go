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
func metaIDFromInterface(i interface{}) metaID {
	switch value := i.(type) {
	case metaID:
		return value
	case types.ID:
		return metaIDFromInterface(readMetaID(value.String()))
	default:
		panic(i)
	}
}

func GenerateMetaID(data types.Data) types.ID {
	return metaID{
		TypeID: data.GetTypeID(),
		HashID: data.GenerateHashID(),
	}
}

func FromID(id types.ID) helpers.Key {
	return metaIDFromInterface(id)
}
