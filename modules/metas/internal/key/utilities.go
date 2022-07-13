// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
)

func readMetaID(metaIDString string) ids.ID {
	idList := stringUtilities.SplitCompositeIDString(metaIDString)
	if len(idList) == 2 {
		return metaID{
			Type: baseIDs.NewStringID(idList[0]),
			Hash: baseIDs.NewStringID(idList[1]),
		}
	}

	return metaID{Type: baseIDs.NewStringID(""), Hash: baseIDs.NewStringID("")}
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

func GenerateMetaID(data data.Data) ids.MetaID {
	return metaID{
		Type: data.GetType(),
		Hash: data.GenerateHash(),
	}
}

func FromID(id ids.ID) helpers.Key {
	return metaIDFromInterface(id)
}
