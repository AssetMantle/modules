// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
)

type auxiliaryResponse struct {
	ids.ClassificationID
}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func newAuxiliaryResponse(classificationID ids.ClassificationID) helpers.AuxiliaryResponse {
	return auxiliaryResponse{
		ClassificationID: classificationID,
	}
}

func GetClassificationIDFromResponse(response helpers.AuxiliaryResponse) ids.ClassificationID {
	switch value := response.(type) {
	case auxiliaryResponse:
		return value.ClassificationID
	default:
		panic(errorConstants.InvalidRequest)
	}
}
