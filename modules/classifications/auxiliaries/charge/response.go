// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package charge

import (
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
)

type auxiliaryResponse struct {
	Success bool   `json:"success"`
	Error   error  `json:"error"`
	Charge  string `json:"collection"`
}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func (auxiliaryResponse auxiliaryResponse) IsSuccessful() bool {
	return auxiliaryResponse.Success
}
func (auxiliaryResponse auxiliaryResponse) GetError() error {
	return auxiliaryResponse.Error
}

func newAuxiliaryResponse(charge string, error error) helpers.AuxiliaryResponse {
	if error != nil {
		return auxiliaryResponse{
			Success: false,
			Error:   error,
			Charge:  charge,
		}
	}

	return auxiliaryResponse{
		Success: true,
		Charge:  charge,
	}
}

func GetChargedValueFromResponse(response helpers.AuxiliaryResponse) (string, error) {
	switch value := response.(type) {
	case auxiliaryResponse:
		if value.IsSuccessful() {
			return value.Charge, nil
		}

		return value.Charge, value.GetError()
	default:
		return "", constants.InvalidRequest
	}
}
