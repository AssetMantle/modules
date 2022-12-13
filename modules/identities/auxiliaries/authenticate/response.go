// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package authenticate

import "github.com/AssetMantle/modules/schema/helpers"

type auxiliaryResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error"`
}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func (auxiliaryResponse auxiliaryResponse) IsSuccessful() bool {
	return auxiliaryResponse.Success
}
func (auxiliaryResponse auxiliaryResponse) GetError() error {
	return auxiliaryResponse.Error
}
func newAuxiliaryResponse(error error) helpers.AuxiliaryResponse {
	success := true
	if error != nil {
		success = false
	}

	return auxiliaryResponse{
		Success: success,
		Error:   error,
	}
}
