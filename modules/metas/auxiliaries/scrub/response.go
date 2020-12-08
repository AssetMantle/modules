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
	} else {
		return auxiliaryResponse{
			Success:    true,
			Properties: properties,
		}
	}
}

func GetPropertiesFromResponse(AuxiliaryResponse helpers.AuxiliaryResponse) (types.Properties, error) {
	switch value := AuxiliaryResponse.(type) {
	case auxiliaryResponse:
		if value.IsSuccessful() {
			return value.Properties, nil
		} else {
			return nil, value.GetError()
		}
	default:
		return nil, errors.NotAuthorized
	}
}
