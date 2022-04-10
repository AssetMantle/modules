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

func readAssetID(assetIDString string) types.ID {
	idList := strings.Split(assetIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return assetID{
			Classification: baseIDs.NewID(idList[0]),
			Hash:           baseIDs.NewID(idList[1]),
		}
	}

	return assetID{Classification: baseIDs.NewID(""), Hash: baseIDs.NewID("")}
}

// TODO remove panic and add error
func assetIDFromInterface(i interface{}) assetID {
	switch value := i.(type) {
	case assetID:
		return value
		// TODO remove this use case
	case types.ID:
		return assetIDFromInterface(readAssetID(value.String()))
	default:
		panic(i)
	}
}

func ReadClassificationID(assetID types.ID) types.ID {
	return assetIDFromInterface(assetID).Classification
}

func FromID(id types.ID) helpers.Key {
	return assetIDFromInterface(id)
}
