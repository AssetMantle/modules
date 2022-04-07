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

func readMetaID(metaIDString string) types.ID {
	idList := strings.Split(metaIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return metaID{
			TypeID: baseIDs.NewID(idList[0]),
			HashID: baseIDs.NewID(idList[1]),
		}
	}

	return metaID{TypeID: baseIDs.NewID(""), HashID: baseIDs.NewID("")}
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
		TypeID: data.GetType(),
		HashID: data.GenerateHash(),
	}
}

func FromID(id types.ID) helpers.Key {
	return metaIDFromInterface(id)
}
