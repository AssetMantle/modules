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

func readAssetID(assetIDString string) types.ID {
	idList := strings.Split(assetIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return assetID{
			ClassificationID: base.NewID(idList[0]),
			HashID:           base.NewID(idList[1]),
		}
	}

	return assetID{ClassificationID: base.NewID(""), HashID: base.NewID("")}
}
func assetIDFromInterface(i interface{}) assetID {
	switch value := i.(type) {
	case assetID:
		return value
	case types.ID:
		return assetIDFromInterface(readAssetID(value.String()))
	default:
		panic(i)
	}
}

func ReadClassificationID(assetID types.ID) types.ID {
	return assetIDFromInterface(assetID).ClassificationID
}

func FromID(id types.ID) helpers.Key {
	return assetIDFromInterface(id)
}
