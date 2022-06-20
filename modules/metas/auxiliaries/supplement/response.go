// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/lists"
)

type auxiliaryResponse struct {
	Success        bool                   `json:"success"`
	Error          error                  `json:"error"`
	MetaProperties lists.MetaPropertyList `json:"metaProperties"`
}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func (auxiliaryResponse auxiliaryResponse) IsSuccessful() bool {
	return auxiliaryResponse.Success
}
func (auxiliaryResponse auxiliaryResponse) GetError() error {
	return auxiliaryResponse.Error
}
func newAuxiliaryResponse(metaProperties lists.MetaPropertyList, error error) helpers.AuxiliaryResponse {
	if error != nil {
		return auxiliaryResponse{
			Success: false,
			Error:   error,
		}
	}

	return auxiliaryResponse{
		Success:        true,
		MetaProperties: metaProperties,
	}
}

func GetMetaPropertiesFromResponse(response helpers.AuxiliaryResponse) (lists.MetaPropertyList, error) {
	switch value := response.(type) {
	case auxiliaryResponse:
		if value.IsSuccessful() {
			return value.MetaProperties, nil
		}

		return nil, value.GetError()
	default:
		return nil, constants.InvalidRequest
	}
}
