/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
)

func readMetaID(metaIDString string) MetaID {
	idList := strings.Split(metaIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return MetaID{
			TypeID: *base.NewID(idList[0]),
			HashID: *base.NewID(idList[1]),
		}
	}

	return MetaID{TypeID: *base.NewID(""), HashID: *base.NewID("")}
}
func metaIDFromInterface(i interface{}) MetaID {
	switch value := i.(type) {
	case MetaID:
		return value
	case types.ID:
		return metaIDFromInterface(readMetaID(value.String()))
	default:
		panic(i)
	}
}

func GenerateMetaID(data types.Data) types.ID {
	return &MetaID{
		TypeID: *base.NewID(data.GetTypeID().String()),
		HashID: *base.NewID(data.GenerateHashID().String()),
	}
}

func FromID(id types.ID) helpers.Key {
	return metaIDFromInterface(id)
}
