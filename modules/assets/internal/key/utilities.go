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

func assetIDFromInterface(id types.ID) assetID {
	switch value := id.(type) {
	case assetID:
		return value
	default:
		return assetIDFromInterface(readAssetID(id.String()))
	}
}

func ReadClassificationID(assetID types.ID) types.ID {
	return assetIDFromInterface(assetID).ClassificationID
}
