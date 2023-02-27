// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/lists"
)

type auxiliaryResponse struct {
	lists.PropertyList
}

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func newAuxiliaryResponse(metaProperties lists.PropertyList) helpers.AuxiliaryResponse {
	return auxiliaryResponse{
		PropertyList: metaProperties,
	}
}

func GetMetaPropertiesFromResponse(response helpers.AuxiliaryResponse) lists.PropertyList {
	switch value := response.(type) {
	case auxiliaryResponse:
		return value.PropertyList
	default:
		panic(errorConstants.InvalidRequest.Wrapf("invalid response type %T", value))
	}
}
