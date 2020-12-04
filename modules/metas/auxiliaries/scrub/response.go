/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package scrub

import (
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type AuxiliaryResponse struct {
	Success    bool             `json:"success"`
	Error      error            `json:"error"`
	Properties types.Properties `json:"properties"`
}

var _ helpers.AuxiliaryResponse = (*AuxiliaryResponse)(nil)

func (auxiliaryResponse AuxiliaryResponse) IsSuccessful() bool {
	return auxiliaryResponse.Success
}
func (auxiliaryResponse AuxiliaryResponse) GetError() error {
	return auxiliaryResponse.Error
}
func newAuxiliaryResponse(properties types.Properties, error error) helpers.AuxiliaryResponse {
	if error != nil {
		return AuxiliaryResponse{
			Success: false,
			Error:   error,
		}
	} else {
		return AuxiliaryResponse{
			Success:    true,
			Properties: properties,
		}
	}
}

func GetPropertiesFromResponse(auxiliaryResponse helpers.AuxiliaryResponse) (types.Properties, error) {
	switch value := auxiliaryResponse.(type) {
	case AuxiliaryResponse:
		if auxiliaryResponse.IsSuccessful() {
			return value.Properties, nil
		} else {
			return nil, auxiliaryResponse.GetError()
		}
	default:
		return nil, errors.NotAuthorized
	}
}
