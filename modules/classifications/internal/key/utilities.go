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

func readClassificationID(classificationIDString string) types.ID {
	idList := strings.Split(classificationIDString, constants.IDSeparator)
	if len(idList) == 2 {
		return classificationID{
			ChainID: base.NewID(idList[0]),
			HashID:  base.NewID(idList[1]),
		}
	}
	return classificationID{ChainID: base.NewID(""), HashID: base.NewID("")}
}
func classificationIDFromInterface(id types.ID) classificationID {
	switch value := id.(type) {
	case classificationID:
		return value
	default:
		return classificationIDFromInterface(readClassificationID(id.String()))
	}
}
