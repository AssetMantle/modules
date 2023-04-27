// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package scrub

import (
	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/lists"
)

type auxiliaryResponse struct {
	lists.PropertyList `json:"propertyList"`
}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func newAuxiliaryResponse(properties lists.PropertyList) helpers.AuxiliaryResponse {
	return auxiliaryResponse{
		PropertyList: properties,
	}
}

func GetPropertiesFromResponse(response helpers.AuxiliaryResponse) lists.PropertyList {
	switch value := response.(type) {
	case auxiliaryResponse:
		return value.PropertyList
	default:
		panic(errorConstants.InvalidRequest.Wrapf("invalid response type %T", value))
	}
}
