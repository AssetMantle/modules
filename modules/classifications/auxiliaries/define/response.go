// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
)

type auxiliaryResponse struct {
	Success          bool   `json:"success"`
	Error            error  `json:"error"`
	ClassificationID ids.ID `json:"classificationID"`
}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func (auxiliaryResponse auxiliaryResponse) IsSuccessful() bool {
	return auxiliaryResponse.Success
}
func (auxiliaryResponse auxiliaryResponse) GetError() error {
	return auxiliaryResponse.Error
}

func newAuxiliaryResponse(classificationID ids.ID, error error) helpers.AuxiliaryResponse {
	if error != nil {
		return auxiliaryResponse{
			Success:          false,
			Error:            error,
			ClassificationID: classificationID,
		}
	}

	return auxiliaryResponse{
		Success:          true,
		ClassificationID: classificationID,
	}
}

func GetClassificationIDFromResponse(response helpers.AuxiliaryResponse) (ids.ID, error) {
	switch value := response.(type) {
	case auxiliaryResponse:
		if value.IsSuccessful() {
			return value.ClassificationID, nil
		}

		return value.ClassificationID, value.GetError()
	default:
		return nil, constants.InvalidRequest
	}
}
