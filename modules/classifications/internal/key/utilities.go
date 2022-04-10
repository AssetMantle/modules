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

func readClassificationID(classificationIDString string) types.ID {
	idList := strings.Split(classificationIDString, constants.IDSeparator)
	if len(idList) == 2 {
		return classificationID{
			ChainID: baseIDs.NewID(idList[0]),
			HashID:  baseIDs.NewID(idList[1]),
		}
	}

	return classificationID{ChainID: baseIDs.NewID(""), HashID: baseIDs.NewID("")}
}
func classificationIDFromInterface(i interface{}) classificationID {
	switch value := i.(type) {
	case classificationID:
		return value
	case types.ID:
		// TODO remove this use case
		return classificationIDFromInterface(readClassificationID(value.String()))
	default:
		// TODO remove panic and introduce error
		panic(i)
	}
}

func FromID(id types.ID) helpers.Key {
	return classificationIDFromInterface(id)
}
