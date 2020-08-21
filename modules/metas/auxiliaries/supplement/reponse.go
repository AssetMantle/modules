/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package supplement

import (
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type AuxiliaryResponse struct {
	Success        bool                 `json:"success"`
	Error          error                `json:"error"`
	MetaProperties types.MetaProperties `json:"metaProperties"`
}

var _ helpers.AuxiliaryResponse = (*AuxiliaryResponse)(nil)

func (auxiliaryResponse AuxiliaryResponse) IsSuccessful() bool {
	return auxiliaryResponse.Success
}
func (auxiliaryResponse AuxiliaryResponse) GetError() error {
	return auxiliaryResponse.Error
}
func newAuxiliaryResponse(metaProperties types.MetaProperties, error error) helpers.AuxiliaryResponse {
	if error != nil {
		return AuxiliaryResponse{
			Success: false,
			Error:   error,
		}
	} else {
		return AuxiliaryResponse{
			Success:        true,
			MetaProperties: metaProperties,
		}
	}

}

func ValidateResponse(auxiliaryResponse helpers.AuxiliaryResponse) (AuxiliaryResponse, error) {
	switch value := auxiliaryResponse.(type) {
	case AuxiliaryResponse:
		if auxiliaryResponse.IsSuccessful() {
			return value, nil
		} else {
			return AuxiliaryResponse{}, auxiliaryResponse.GetError()
		}
	default:
		return AuxiliaryResponse{}, errors.NotAuthorized
	}
}
