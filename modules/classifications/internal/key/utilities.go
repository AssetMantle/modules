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

func readClassificationID(classificationIDString string) ClassificationID {
	idList := strings.Split(classificationIDString, constants.IDSeparator)
	if len(idList) == 2 {
		return ClassificationID{
			ChainID: base.NewID(idList[0]),
			HashID:  base.NewID(idList[1]),
		}
	}

	return ClassificationID{ChainID: base.NewID(""), HashID: base.NewID("")}
}
func classificationIDFromInterface(i interface{}) ClassificationID {
	switch value := i.(type) {
	case ClassificationID:
		return value
	case types.ID:
		return classificationIDFromInterface(readClassificationID(value.String()))
	default:
		panic(i)
	}
}

func FromID(id types.ID) helpers.Key {
	return classificationIDFromInterface(id)
}
