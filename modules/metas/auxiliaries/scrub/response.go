// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/types"
)

type auxiliaryResponse struct {
	Success    bool             `json:"success"`
	Error      error            `json:"error"`
	Properties types.Properties `json:"properties"`
}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func (auxiliaryResponse auxiliaryResponse) IsSuccessful() bool {
	return auxiliaryResponse.Success
}
func (auxiliaryResponse auxiliaryResponse) GetError() error {
	return auxiliaryResponse.Error
}
func newAuxiliaryResponse(properties types.Properties, error error) helpers.AuxiliaryResponse {
	if error != nil {
		return auxiliaryResponse{
			Success: false,
			Error:   error,
		}
	}

	return auxiliaryResponse{
		Success:    true,
		Properties: properties,
	}
}

func GetPropertiesFromResponse(response helpers.AuxiliaryResponse) (types.Properties, error) {
	switch value := response.(type) {
	case auxiliaryResponse:
		if value.IsSuccessful() {
			return value.Properties, nil
		}

		return nil, value.GetError()
	default:
		return nil, errors.NotAuthorized
	}
}
