// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"strings"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"

	"github.com/AssetMantle/modules/constants"
)

func readAssetID(assetIDString string) ids.ID {
	idList := strings.Split(assetIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return assetID{
			ClassificationID: baseIDs.NewID(idList[0]),
			HashID:           baseIDs.NewID(idList[1]),
		}
	}

	return assetID{ClassificationID: baseIDs.NewID(""), HashID: baseIDs.NewID("")}
}

// TODO remove panic and add error
func assetIDFromInterface(i interface{}) (assetID, error) {
	switch value := i.(type) {
	case assetID:
		return value, nil
		// TODO remove this use case
	case ids.ID:
		return assetIDFromInterface(readAssetID(value.String()))
	default:
		return assetID{}, errorConstants.MetaDataError
	}
}

func ReadClassificationID(id ids.ID) ids.ID {
	if assetID, err := assetIDFromInterface(id); err != nil {
		panic(err)
	} else {
		return assetID.ClassificationID
	}
}

func FromID(id ids.ID) helpers.Key {
	if assetID, err := assetIDFromInterface(id); err != nil {
		panic(err)
	} else {
		return assetID
	}
}
