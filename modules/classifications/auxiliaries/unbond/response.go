// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unbond

import (
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
)

type auxiliaryResponse struct {
	Success bool   `json:"success"`
	Error   error  `json:"error"`
	Amount  string `json:"collection"`
}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func (auxiliaryResponse auxiliaryResponse) IsSuccessful() bool {
	return auxiliaryResponse.Success
}
func (auxiliaryResponse auxiliaryResponse) GetError() error {
	return auxiliaryResponse.Error
}

func newAuxiliaryResponse(amount string, error error) helpers.AuxiliaryResponse {
	if error != nil {
		return auxiliaryResponse{
			Success: false,
			Error:   error,
			Amount:  amount,
		}
	}

	return auxiliaryResponse{
		Success: true,
		Amount:  amount,
	}
}

func GetChargedValueFromResponse(response helpers.AuxiliaryResponse) (string, error) {
	switch value := response.(type) {
	case auxiliaryResponse:
		if value.IsSuccessful() {
			return value.Amount, nil
		}

		return value.Amount, value.GetError()
	default:
		return "", constants.InvalidRequest
	}
}
