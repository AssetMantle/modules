// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
)

func classificationIDFromInterface(i interface{}) (classificationID, error) {
	switch value := i.(type) {
	case classificationID:
		return value, nil
	default:
		return classificationID{}, errorConstants.MetaDataError
	}
}

func FromID(id ids.ID) helpers.Key {
	if classificationID, err := classificationIDFromInterface(id); err != nil {
		// TODO plug all panic scenarios
		panic(err)
	} else {
		return classificationID
	}
}
