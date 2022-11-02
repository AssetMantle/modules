// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/lists"
)

type auxiliaryResponse struct {
	Success            bool  `json:"success"`
	Error              error `json:"error"`
	lists.PropertyList `json:"propertyList"`
}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func (auxiliaryResponse auxiliaryResponse) IsSuccessful() bool {
	return auxiliaryResponse.Success
}
func (auxiliaryResponse auxiliaryResponse) GetError() error {
	return auxiliaryResponse.Error
}
func newAuxiliaryResponse(properties lists.PropertyList, error error) helpers.AuxiliaryResponse {
	if error != nil {
		return auxiliaryResponse{
			Success: false,
			Error:   error,
		}
	}

	return auxiliaryResponse{
		Success:      true,
		PropertyList: properties,
	}
}

func GetPropertiesFromResponse(response helpers.AuxiliaryResponse) (lists.PropertyList, error) {
	switch value := response.(type) {
	case auxiliaryResponse:
		if value.IsSuccessful() {
			return value.PropertyList, nil
		}

		return nil, value.GetError()
	default:
		return nil, constants.NotAuthorized
	}
}
