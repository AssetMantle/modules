// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"strings"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"

	"github.com/AssetMantle/modules/constants"
)

func readMetaID(metaIDString string) ids.ID {
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
	case ids.ID:
		return metaIDFromInterface(readMetaID(value.String()))
	default:
		panic(i)
	}
}

func GenerateMetaID(data data.Data) ids.ID {
	return metaID{
		TypeID: data.GetType(),
		HashID: data.GenerateHash(),
	}
}

func FromID(id ids.ID) helpers.Key {
	return metaIDFromInterface(id)
}
