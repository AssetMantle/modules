// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"

	"github.com/AssetMantle/modules/helpers"
)

type auxiliaryResponse struct {
	ids.ClassificationID
}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func NewAuxiliaryResponse(classificationID ids.ClassificationID) helpers.AuxiliaryResponse {
	return auxiliaryResponse{
		ClassificationID: classificationID,
	}
}

func GetClassificationIDFromResponse(response helpers.AuxiliaryResponse) ids.ClassificationID {
	switch value := response.(type) {
	case auxiliaryResponse:
		return value.ClassificationID
	default:
		panic(errorConstants.InvalidRequest.Wrapf("invalid response type %T", value))
	}
}
