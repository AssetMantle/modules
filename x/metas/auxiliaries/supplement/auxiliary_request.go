// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supplement

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/go/properties"
)

type auxiliaryRequest struct {
	PropertyList []properties.Property
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	for _, property := range auxiliaryRequest.PropertyList {
		if property != nil {
			if err := property.ValidateBasic(); err != nil {
				return constants.InvalidRequest.Wrapf("invalid property %s: %s", property.GetKey().AsString(), err.Error())
			}
		}
	}

	return nil
}

func NewAuxiliaryRequest(propertyList ...properties.Property) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		PropertyList: propertyList,
	}
}
